// Copyright 2016 The LUCI Authors.
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

package common

import (
	"golang.org/x/net/context"

	"github.com/luci/luci-go/luci_config/common/cfgtypes"
	"github.com/luci/luci-go/luci_config/server/cfgclient/access"
	"github.com/luci/luci-go/luci_config/server/cfgclient/backend"
	"github.com/luci/luci-go/server/auth"
)

// Helper functions for ACL checking.

// IsAllowed checks to see if the user in the context is allowed to access
// the given project.
func IsAllowed(c context.Context, project string) (bool, error) {
	// Get the project, because that's where the ACLs lie.
	err := access.Check(
		c, backend.AsUser,
		cfgtypes.ProjectConfigSet(cfgtypes.ProjectName(project)))
	switch err {
	case nil:
		return true, nil
	case access.ErrNoAccess:
		return false, nil
	default:
		return false, err
	}
}

// IsAllowedInternal is a shorthand for checking to see if the user is a reader
// of a magic project named "chrome".
func IsAllowedInternal(c context.Context) (bool, error) {
	settings := GetSettings(c)
	if settings.Buildbot.InternalReader == "" {
		return false, nil
	}
	return auth.IsMember(c, settings.Buildbot.InternalReader)
}