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

package streamserver

import (
	"context"
	"strings"
	"testing"

	"go.chromium.org/luci/logdog/client/butlerlib/streamclient"

	. "github.com/smartystreets/goconvey/convey"
	. "go.chromium.org/luci/common/testing/assertions"
)

func TestWindowsNamedPipeServer(t *testing.T) {
	t.Parallel()

	// TODO(crbug.com/963705): This test is flaky. Hangs with non-insignificant
	// probability.
	SkipConvey(`A named pipe server`, t, func() {
		ctx := context.Background()

		Convey(`Will refuse to create if there is an empty path.`, func() {
			_, err := newStreamServer(ctx, "")
			So(err, ShouldErrLike, "cannot have empty name")
		})

		Convey(`Will refuse to create if longer than maximum length.`, func() {
			_, err := newStreamServer(ctx, strings.Repeat("A", maxWindowsNamedPipeLength+1))
			So(err, ShouldErrLike, "name exceeds maximum length")
		})

		Convey(`When created and listening.`, func() {
			svr, err := newStreamServer(ctx, "ButlerNamedPipeTest")
			So(err, ShouldBeNil)

			So(svr.Listen(), ShouldBeNil)
			defer svr.Close()

			client, err := streamclient.New(svr.Address(), "")
			So(err, ShouldBeNil)

			testClientServer(t, svr, client)
		})
	})
}
