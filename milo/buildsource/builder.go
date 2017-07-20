// Copyright 2017 The LUCI Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package buildsource

import (
	"strings"

	"golang.org/x/net/context"

	"github.com/luci/gae/service/datastore"

	"github.com/luci/luci-go/common/errors"
	"github.com/luci/luci-go/milo/api/resp"
	"github.com/luci/luci-go/milo/buildsource/buildbot"
	"github.com/luci/luci-go/milo/buildsource/buildbucket"
	"github.com/luci/luci-go/milo/common"
)

// BuilderID is the universal ID of a builder, and has the form:
//   buildbucket/bucket/builder
//   buildbot/master/builder
type BuilderID string

// Split breaks the BuilderID into pieces.
//   - backend is either 'buildbot' or 'buildbucket'
//   - backendGroup is either the bucket or master name
//   - builderName is the builder name.
//
// Returns an error if the BuilderID is malformed (wrong # slashes) or if any of
// the pieces are empty.
func (b BuilderID) Split() (backend, backendGroup, builderName string, err error) {
	toks := strings.SplitN(string(b), "/", 3)
	if len(toks) != 3 {
		err = errors.Reason("bad BuilderID: not enough tokens: %q", b).
			Tag(common.CodeParameterError).Err()
		return
	}
	backend, backendGroup, builderName = toks[0], toks[1], toks[2]
	switch backend {
	case "buildbot", "buildbucket":
	default:
		err = errors.Reason("bad BuilderID: unknown backend %q", backend).
			Tag(common.CodeParameterError).Err()
		return
	}
	if backendGroup == "" {
		err = errors.New("bad BuilderID: empty backendGroup", common.CodeParameterError)
		return
	}
	if builderName == "" {
		err = errors.New("bad BuilderID: empty builderName", common.CodeParameterError)
		return
	}
	return
}

// Get allows you to obtain the resp.Builder that corresponds with this
// BuilderID.
func (b BuilderID) Get(c context.Context, limit int, cursorStr string) (*resp.Builder, error) {
	// TODO(iannucci): replace these implementations with a BuildSummary query.
	source, group, builder, err := b.Split()
	if err != nil {
		return nil, err
	}

	switch source {
	case "buildbot":
		cursor, err := datastore.DecodeCursor(c, cursorStr)
		if err != nil {
			return nil, errors.Annotate(err, "bad cursor").Err()
		}
		return buildbot.GetBuilder(c, group, builder, limit, cursor)

	case "buildbucket":
		return buildbucket.GetBuilder(c, group, builder, limit)
	}

	panic("impossible")
}