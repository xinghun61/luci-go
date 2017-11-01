// Copyright 2017 The LUCI Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package buildstore

import (
	"strconv"

	"golang.org/x/net/context"

	"go.chromium.org/gae/service/datastore"
	"go.chromium.org/gae/service/memcache"
	"go.chromium.org/luci/buildbucket"
	bbapi "go.chromium.org/luci/common/api/buildbucket/buildbucket/v1"
	"go.chromium.org/luci/common/clock"
	"go.chromium.org/luci/common/data/strpair"
	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/logging"
	"go.chromium.org/luci/common/sync/parallel"
	"go.chromium.org/luci/milo/api/buildbot"
	api "go.chromium.org/luci/milo/api/proto"
)

// Ternary has 3 defined values: either (zero), yes and no.
type Ternary int

const (
	Either Ternary = iota
	Yes
	No
)

func (t Ternary) filter(q *datastore.Query, fieldName string) *datastore.Query {
	switch t {
	case Yes:
		return q.Eq(fieldName, true)
	case No:
		return q.Eq(fieldName, false)
	default:
		return q
	}
}

// Query is a build query.
type Query struct {
	Master   string
	Builder  string
	Limit    int
	Finished Ternary
	Cursor   string

	// The following fields are tuning parameters specific to a buildstore
	// implementation. Their usage implies understanding of how emulation
	// works.

	// KeyOnly, if true, makes the datastore query keys-only.
	// Loaded Buildbot builds will have only master, builder and number.
	KeyOnly bool // make the data

	// NoAnnotationFetch, if true, will not fetch annotation proto from LogDog.
	// Loaded LUCI builds will not have properties, steps, logs or text.
	NoAnnotationFetch bool

	// NoChangeFetch, if true, will not load change history from Gitiles.
	// Loaded LUCI builds will not have Blame or SourceStamp.Changes.
	NoChangeFetch bool
}

func (q *Query) dsQuery() *datastore.Query {
	dsq := datastore.NewQuery(buildKind)
	if q.Master != "" {
		dsq = dsq.Eq("master", q.Master)
	}
	if q.Builder != "" {
		dsq = dsq.Eq("builder", q.Builder)
	}
	dsq = q.Finished.filter(dsq, "finished")
	if q.Limit > 0 {
		dsq = dsq.Limit(int32(q.Limit))
	}
	if q.KeyOnly {
		dsq = dsq.KeysOnly(true)
	}
	return dsq
}

// QueryResult is a result of running a Query.
type QueryResult struct {
	Builds     []*buildbot.Build // ordered from greater-number to lower-number
	NextCursor string
	PrevCursor string
}

// GetBuilds executes a build query and returns results.
// Does not check access.
func GetBuilds(c context.Context, q Query) (*QueryResult, error) {
	switch {
	case q.Master == "":
		return nil, errors.New("master is required")
	case q.Builder == "":
		return nil, errors.New("builder is required")
	}

	emOptions, err := GetEmulationOptions(c, q.Master, q.Builder)
	if err != nil {
		return nil, err
	}
	if emOptions != nil {
		return getEmulatedBuilds(c, q, *emOptions)
	}
	return getDatastoreBuilds(c, q)
}

func getEmulatedBuilds(c context.Context, q Query, emOptions api.EmulationOptions) (*QueryResult, error) {
	if q.Cursor != "" {
		// build query emulation does not support cursors
		logging.Warningf(c, "ignoring cursor %q", q.Cursor)
		q.Cursor = ""
	}

	bb, err := buildbucketClient(c)
	if err != nil {
		return nil, err
	}

	search := bb.Search().
		Bucket(emOptions.Bucket).
		Tag(strpair.Format(buildbucket.TagBuilder, q.Builder)).
		Context(c)
	switch q.Finished {
	case Yes:
		search.Status(bbapi.StatusCompleted)
	case No:
		search.Status(bbapi.StatusFilterIncomplete)
	}

	start := clock.Now(c)
	msgs, err := search.Fetch(q.Limit, nil)
	if err != nil {
		return nil, errors.Annotate(err, "searching on buildbucket").Err()
	}
	logging.Infof(c, "buildbucket search took %s", clock.Since(c, start))

	builds := make([]*buildbot.Build, len(msgs))
	start = clock.Now(c)
	err = parallel.WorkPool(10, func(work chan<- func() error) {
		for i, msg := range msgs {
			i := i
			msg := msg
			work <- func() error {
				// may load annotations from logdog, that's why parallelized.
				b, err := buildFromBuildbucket(c, q.Master, msg, !q.NoAnnotationFetch)
				if err != nil {
					return err
				}
				if b.Number >= int(emOptions.StartFrom) {
					builds[i] = b
				}
				return nil
			}
		}
	})
	if err != nil {
		return nil, err
	}
	logging.Infof(c, "conversion from buildbucket builds took %s", clock.Since(c, start))

	if !q.NoChangeFetch {
		start = clock.Now(c)
		// We need to compute blamelist for multiple builds.
		// 1) We don't have a guarantee that the numbers are contiguous
		// 2) For some builds, we may have cached changes
		// => compute blamelist for each build individually

		// cache build revisions before fetching changes
		// in case build numbers are contiguous.
		caches := make([]memcache.Item, len(builds))
		for i, b := range builds {
			caches[i] = buildRevCache(c, b)
		}
		memcache.Set(c, caches...)

		err = parallel.WorkPool(10, func(work chan<- func() error) {
			for _, b := range builds {
				b := b
				work <- func() error {
					return blame(c, b)
				}
			}
		})
		if err != nil {
			return nil, errors.Annotate(err, "blamelist computation failed").Err()
		}
		logging.Infof(c, "blamelist computation took %s", clock.Since(c, start))
	}

	// Still need to load builds from datastore?
	if q.Limit <= 0 || len(builds) < q.Limit {
		// get builds up to emOptions.StartFrom.
		q.Cursor = strconv.Itoa(int(-emOptions.StartFrom))
		if q.Limit > 0 {
			// don't fetch more than we need
			q.Limit -= len(builds)
		}
		dsRes, err := getDatastoreBuilds(c, q)
		if err != nil {
			return nil, errors.Annotate(err, "running datastore query %#v", q).Err()
		}
		builds = append(builds, dsRes.Builds...)
	}

	return &QueryResult{Builds: builds}, nil
}

func getDatastoreBuilds(c context.Context, q Query) (*QueryResult, error) {
	var builds []*buildEntity
	if q.Limit > 0 {
		builds = make([]*buildEntity, 0, q.Limit)
	}

	dsq := q.dsQuery()

	// CUSTOM CURSOR.
	// This function uses a custom cursor based on build numbers.
	// A cursor is a build number that defines a page boundary.
	// If >=0, it is the inclusive lower boundary.
	//   Example: cursor="10", means return builds ...12, 11, 10.
	// If <0, it is the exclusive upper boundary, negated.
	//   Example: -10, means return builds 9, 8, 7...
	cursorNumber := 0
	order := "-number"
	reverse := false
	hasCursor := false
	if q.Cursor != "" {
		if cursorNumber, err := strconv.Atoi(q.Cursor); err == nil {
			hasCursor = true
			if cursorNumber >= 0 {
				dsq = dsq.Gte("number", cursorNumber)
				order = "number"
				reverse = true
			} else {
				dsq = dsq.Lt("number", -cursorNumber)
			}
		} else {
			return nil, errors.New("bad cursor")
		}
	}
	dsq = dsq.Order(order)

	logging.Debugf(c, "running datastore query: %s", dsq)
	err := datastore.GetAll(c, dsq, &builds)
	if err != nil {
		return nil, err
	}
	if reverse {
		for i, j := 0, len(builds)-1; i < j; i, j = i+1, j-1 {
			builds[i], builds[j] = builds[j], builds[i]
		}
	}
	res := &QueryResult{
		Builds: make([]*buildbot.Build, len(builds)),
	}
	for i, b := range builds {
		res.Builds[i] = (*buildbot.Build)(b)
	}

	// Compute prev and next cursors.
	switch {
	case len(res.Builds) > 0:
		// res.Builds are ordered by numbers descending.

		// previous page must display builds with higher numbers.
		if !hasCursor {
			// do not generate a prev cursor for a non-cursor query
		} else {
			// positive cursors are inclusive
			res.PrevCursor = strconv.Itoa(res.Builds[0].Number + 1)
		}

		// next page must display builds with lower numbers.

		if lastNum := res.Builds[len(res.Builds)-1].Number; lastNum == 0 {
			// this is the first ever build, 0, do not generate a cursor
		} else {
			// negative cursors are exclusive.
			res.NextCursor = strconv.Itoa(-lastNum)
		}

	case cursorNumber > 0:
		// no builds and cursor is the inclusive lower boundary
		// e.g. cursor asks for builds after 10,
		// but there are only 0..5 builds.
		// Make the next cursor for builds <10.
		res.NextCursor = strconv.Itoa(-cursorNumber)

	default:
		// there can't be any builds.
	}

	return res, nil
}