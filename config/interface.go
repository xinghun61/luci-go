// Copyright 2015 The LUCI Authors.
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

package config

import (
	"errors"
	"net/url"

	"golang.org/x/net/context"
)

// ErrNoConfig is returned if requested config does not exist.
var ErrNoConfig = errors.New("no such config")

// Meta is metadata about a single configuration file.
type Meta struct {
	// ConfigSet is the config set name (e.g. "projects/<id>") this config
	// belongs to.
	ConfigSet Set `json:"configSet,omitempty"`

	// Path is the filename relative to the root of the config set,
	// without leading slash, e.g. "luci-scheduler.cfg".
	Path string `json:"path,omitempty"`

	// ContentHash can be used to quickly check that content didn't change.
	ContentHash string `json:"contentHash,omitempty"`

	// Revision is git SHA1 of a repository the config was fetched from.
	Revision string `json:"revision,omitempty"`

	// ViewURL is the URL surfaced for viewing the config.
	ViewURL string `json:"view_url,omitempty"`
}

// FormatSpec is a specification for formatted data.
type FormatSpec struct {
	// Formatter is the name of the formatter that produces the content data.
	//
	// An empty string means the original config service format.
	Formatter string

	// Data is additional format data describing the type. It may be empty.
	Data string
}

// Unformatted returns true if fs does not specify a format.
func (fs *FormatSpec) Unformatted() bool { return fs.Formatter == "" }

// Config is a configuration file along with its metadata.
type Config struct {
	Meta

	// FormatSpec, if non-empty, qualifies the format of the Content.
	//
	// It is set only if config was obtained through cfgclient guts that does
	// reformatting (e.g. formats text protos into binary protos).
	FormatSpec FormatSpec `json:"formatSpec,omitempty"`

	// Error is not nil if there where troubles fetching this config. Used only
	// by functions that operate with multiple configs at once, such as
	// GetProjectConfigs and GetRefConfigs.
	Error error `json:"error,omitempty"`

	// Content is the actual body of the config file.
	Content string `json:"content,omitempty"`
}

// RepoType is the type of the repo the Project is stored in.
type RepoType string

const (
	// GitilesRepo means a repo is backed by the Gitiles service.
	GitilesRepo RepoType = "GITILES"

	// UnknownRepo means a repo is backed by an unknown service.
	// It may be an invalid repo.
	UnknownRepo RepoType = "UNKNOWN"
)

// Project is a project registered in the luci-config service.
type Project struct {
	// ID is unique project identifier.
	ID string

	// Name is a short friendly display name of the project.
	Name string

	// RepoType specifies in what kind of storage projects configs are stored.
	RepoType RepoType

	// RepoUrl is the location of this project code repository. May be nil if
	// unknown or cannot be parsed.
	RepoURL *url.URL
}

// Interface represents low-level luci-config service API.
//
// All methods accept context.Context they use for deadlines and for passing to
// callbacks (if the implementation uses any). Contexts here don't necessary
// relate to a "global" package context (used by GetImplementation and
// SetImplementation), though very often they are the same (as is the case when
// using package-level functions like GetConfig).
//
// For example, unit tests may instantiate an implementation of Interface
// directly and don't bother registering it in the context with
// SetImplementation(...).
//
// Transient errors are wrapped in errors.Transient. See common/errors.
type Interface interface {
	// GetConfig returns a config at a path in a config set or ErrNoConfig
	// if missing. If metaOnly is true, returned Config struct has only Meta set
	// (and the call is faster).
	GetConfig(ctx context.Context, configSet Set, path string, metaOnly bool) (*Config, error)

	// GetConfigByHash returns the contents of a config, as identified by its
	// content hash, or ErrNoConfig if missing.
	GetConfigByHash(ctx context.Context, contentHash string) (string, error)

	// GetConfigSetLocation returns the URL location of a config set.
	GetConfigSetLocation(ctx context.Context, configSet Set) (*url.URL, error)

	// GetProjectConfigs returns all the configs at the given path in all
	// projects that have such config. If metaOnly is true, returned Config
	// structs have only Meta set (and the call is faster).
	GetProjectConfigs(ctx context.Context, path string, metaOnly bool) ([]Config, error)

	// GetProjects returns all the registered projects in the configuration
	// service.
	GetProjects(ctx context.Context) ([]Project, error)

	// GetRefConfigs returns the config at the given path in all refs of all
	// projects that have such config. If metaOnly is true, returned Config
	// structs have only Meta set (and the call is faster).
	GetRefConfigs(ctx context.Context, path string, metaOnly bool) ([]Config, error)

	// GetRefs returns the list of refs for a project.
	GetRefs(ctx context.Context, projectID string) ([]string, error)
}
