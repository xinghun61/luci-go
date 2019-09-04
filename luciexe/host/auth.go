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

package host

import (
	"context"
	"os"

	"go.chromium.org/luci/common/errors"
	"go.chromium.org/luci/common/system/environ"
	"go.chromium.org/luci/lucictx"
)

func startAuthServices(ctx context.Context, opts *Options) (cleanupSlice, error) {
	var myCleanups cleanupSlice
	defer myCleanups.run()

	env := environ.New(nil)

	if err := opts.ExeAuth.Launch(ctx, opts.authDir); err != nil {
		return nil, errors.Annotate(err, "setting up task auth").Err()
	}
	opts.ExeAuth.Report(ctx)
	ctx = opts.ExeAuth.Export(ctx, env)
	myCleanups = append(myCleanups, func() error {
		opts.ExeAuth.Close(ctx)
		return nil
	})

	exported, err := lucictx.ExportInto(ctx, opts.lucictxDir)
	if err != nil {
		return nil, errors.Annotate(err, "exporting LUCI_CONTEXT").Err()
	}
	myCleanups = append(myCleanups, exported.Close)
	exported.SetInEnviron(env)

	if err := env.Iter(os.Setenv); err != nil {
		return nil, errors.Annotate(err, "setting up environment").Err()
	}

	callerCleanups := myCleanups
	myCleanups = nil

	return callerCleanups, nil
}
