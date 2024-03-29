// Copyright 2019 The LUCI Authors.
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

package pbutil

import (
	"fmt"
	"sort"

	pb "go.chromium.org/luci/resultdb/proto/v1"
)

// StringPair creates a pb.StringPair with the given strings as key/value field values.
func StringPair(k, v string) *pb.StringPair {
	return &pb.StringPair{Key: k, Value: v}
}

// StringPairs creates a slice of pb.StringPair from a list of strings alternating key/value.
//
// Panics if an odd number of tokens is passed.
func StringPairs(pairs ...string) []*pb.StringPair {
	if len(pairs)%2 != 0 {
		panic(fmt.Sprintf("odd number of tokens in %q", pairs))
	}

	strpairs := make([]*pb.StringPair, len(pairs)/2)
	for i := range strpairs {
		strpairs[i] = StringPair(pairs[2*i], pairs[2*i+1])
	}
	return strpairs
}

// StringPairsContain checks if item is present in pairs.
func StringPairsContain(pairs []*pb.StringPair, item *pb.StringPair) bool {
	for _, p := range pairs {
		if p.Key == item.Key && p.Value == item.Value {
			return true
		}
	}
	return false
}

// sortStringPairs sorts in-place the tags slice lexicographically by key, then value.
func sortStringPairs(tags []*pb.StringPair) {
	sort.Slice(tags, func(i, j int) bool {
		if tags[i].Key != tags[j].Key {
			return tags[i].Key < tags[j].Key
		}
		return tags[i].Value < tags[j].Value
	})
}
