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

package isolate

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"

	"go.chromium.org/luci/client/archiver"
	"go.chromium.org/luci/common/flag/stringlistflag"
	"go.chromium.org/luci/common/flag/stringmapflag"
	"go.chromium.org/luci/common/isolated"
	"go.chromium.org/luci/common/isolatedclient"
	"go.chromium.org/luci/common/runtime/tracer"
)

// IsolatedGenJSONVersion is used in the batcharchive json format.
//
// TODO(tandrii): Migrate to batch_archive.go.
const IsolatedGenJSONVersion = 1

// ValidVariable is the regexp of valid isolate variable name.
const ValidVariable = "[A-Za-z_][A-Za-z_0-9]*"

var validVariableMatcher = regexp.MustCompile(ValidVariable)
var variableSubstitutionMatcher = regexp.MustCompile("<\\(" + ValidVariable + "\\)")

// IsValidVariable returns true if the variable is a valid symbol name.
func IsValidVariable(variable string) bool {
	return validVariableMatcher.MatchString(variable)
}

// Tree to be isolated.
type Tree struct {
	Cwd  string
	Opts ArchiveOptions
}

// ArchiveOptions for archiving trees.
type ArchiveOptions struct {
	Isolate         string              `json:"isolate"`
	Isolated        string              `json:"isolated"`
	Blacklist       stringlistflag.Flag `json:"blacklist"`
	PathVariables   stringmapflag.Value `json:"path_variables"`
	ExtraVariables  stringmapflag.Value `json:"extra_variables"`
	ConfigVariables stringmapflag.Value `json:"config_variables"`
}

// Init initializes with non-nil values.
func (a *ArchiveOptions) Init() {
	a.Blacklist = stringlistflag.Flag{}
	a.PathVariables = map[string]string{}
	if runtime.GOOS == "windows" {
		a.PathVariables["EXECUTABLE_SUFFIX"] = ".exe"
	} else {
		a.PathVariables["EXECUTABLE_SUFFIX"] = ""
	}
	a.ExtraVariables = map[string]string{}
	a.ConfigVariables = map[string]string{}
}

// PostProcess post-processes the flags to fix any compatibility issue.
func (a *ArchiveOptions) PostProcess(cwd string) {
	// Set default blacklist only if none is set.
	if len(a.Blacklist) == 0 {
		// This cannot be generalized as ".*" as there is known use that require
		// a ".pki" directory to be mapped.
		a.Blacklist = stringlistflag.Flag{
			// Temporary python files.
			"*.pyc",
			// Temporary vim files.
			"*.swp",
			".git",
			".hg",
			".svn",
		}
	}
	if !filepath.IsAbs(a.Isolate) {
		a.Isolate = filepath.Join(cwd, a.Isolate)
	}
	a.Isolate = filepath.Clean(a.Isolate)

	if !filepath.IsAbs(a.Isolated) {
		a.Isolated = filepath.Join(cwd, a.Isolated)
	}
	a.Isolated = filepath.Clean(a.Isolated)

	for k, v := range a.PathVariables {
		// This is due to a Windows + GYP specific issue, where double-quoted paths
		// would get mangled in a way that cannot be resolved unless a space is
		// injected.
		a.PathVariables[k] = strings.TrimSpace(v)
	}
}

// ReplaceVariables replaces any occurrences of '<(FOO)' in 'str' with the
// corresponding variable from 'opts'.
//
// If any substitution refers to a variable that is missing, the returned error will
// refer to the first such variable. In the case of errors, the returned string will
// still contain a valid result for any non-missing substitutions.
func ReplaceVariables(str string, opts *ArchiveOptions) (string, error) {
	var err error
	subst := variableSubstitutionMatcher.ReplaceAllStringFunc(str,
		func(match string) string {
			varName := match[2 : len(match)-1]
			if v, ok := opts.PathVariables[varName]; ok {
				return v
			}
			if v, ok := opts.ExtraVariables[varName]; ok {
				return v
			}
			if v, ok := opts.ConfigVariables[varName]; ok {
				return v
			}
			if err == nil {
				err = errors.New("no value for variable '" + varName + "'")
			}
			return match
		})
	return subst, err
}

// Archive processes a .isolate, generates a .isolated and archive it.
// Returns a *PendingItem to the .isolated.
func Archive(arch *archiver.Archiver, opts *ArchiveOptions) *archiver.PendingItem {
	displayName := filepath.Base(opts.Isolated)
	defer tracer.Span(arch, strings.SplitN(displayName, ".", 2)[0]+":archive", nil)(nil)
	f, err := archive(arch, opts, displayName)
	if err != nil {
		i := &archiver.PendingItem{DisplayName: displayName}
		i.SetErr(err)
		return i
	}
	return f
}

// ProcessIsolate parses an isolate file, returning the list of dependencies
// (both files and directories), the root directory and the initial Isolated struct.
func ProcessIsolate(opts *ArchiveOptions) ([]string, string, *isolated.Isolated, error) {
	content, err := ioutil.ReadFile(opts.Isolate)
	if err != nil {
		return nil, "", nil, err
	}
	cmd, deps, readOnly, isolateDir, err := LoadIsolateForConfig(filepath.Dir(opts.Isolate), content, opts.ConfigVariables)
	if err != nil {
		return nil, "", nil, err
	}

	// Expand variables in the commands.
	for i := range cmd {
		if cmd[i], err = ReplaceVariables(cmd[i], opts); err != nil {
			return nil, "", nil, err
		}
	}

	// Expand variables in the deps, and convert each path to an absolute form.
	for i := range deps {
		dep, err := ReplaceVariables(deps[i], opts)
		if err != nil {
			return nil, "", nil, err
		}
		deps[i] = filepath.Join(isolateDir, dep)
	}

	// Find the root directory of all the files (the root might be above isolateDir).
	rootDir := isolateDir
	for _, dep := range deps {
		// Check if the dep is outside isolateDir.
		base := filepath.Dir(dep)
		for {
			rel, err := filepath.Rel(rootDir, base)
			if err != nil {
				return nil, "", nil, err
			}
			if !strings.HasPrefix(rel, "..") {
				break
			}
			newRootDir := filepath.Dir(rootDir)
			if newRootDir == rootDir {
				return nil, "", nil, errors.New("failed to find root dir")
			}
			rootDir = newRootDir
		}
	}
	if rootDir != isolateDir {
		log.Printf("Root: %s", rootDir)
	}

	// Prepare the .isolated struct.
	isol := &isolated.Isolated{
		Algo:     "sha-1",
		Files:    map[string]isolated.File{},
		ReadOnly: readOnly.ToIsolated(),
		Version:  isolated.IsolatedFormatVersion,
	}
	if len(cmd) != 0 {
		isol.Command = cmd
		// Only set RelativeCwd if a command was also specified. This reduce the
		// noise for Swarming tasks where the command is specified as part of the
		// Swarming task request and not through the isolated file.
		if rootDir != isolateDir {
			relPath, err := filepath.Rel(rootDir, isolateDir)
			if err != nil {
				return nil, "", nil, err
			}
			isol.RelativeCwd = relPath
		}
	}
	return deps, rootDir, isol, nil
}

func archive(arch *archiver.Archiver, opts *ArchiveOptions, displayName string) (*archiver.PendingItem, error) {
	end := tracer.Span(arch, strings.SplitN(displayName, ".", 2)[0]+":loading", nil)
	deps, rootDir, i, err := ProcessIsolate(opts)
	end(tracer.Args{"err": err})
	if err != nil {
		return nil, err
	}
	// Handle each dependency, either a file or a directory.
	var fileItems []*archiver.PendingItem
	var dirItems []*archiver.PendingItem
	for _, dep := range deps {
		relPath, err := filepath.Rel(rootDir, dep)
		if err != nil {
			return nil, err
		}
		// Grab the stats right away; this can be used for both checking whether
		// it's a directory and checking whether it's a link.
		info, err := os.Lstat(dep)
		if err != nil {
			return nil, err
		}
		if mode := info.Mode(); mode.IsDir() {
			if relPath, err = filepath.Rel(rootDir, dep); err != nil {
				return nil, err
			}
			dirItems = append(dirItems, archiver.PushDirectory(arch, dep, relPath, opts.Blacklist))
		} else {
			if mode&os.ModeSymlink == os.ModeSymlink {
				l, err := os.Readlink(dep)
				if err != nil {
					// Kill the process: there's no reason to continue if a file is
					// unavailable.
					log.Fatalf("Unable to stat %q: %v", dep, err)
				}
				i.Files[relPath] = isolated.SymLink(l)
			} else {
				i.Files[relPath] = isolated.BasicFile("", int(mode.Perm()), info.Size())
				fileItems = append(fileItems, arch.PushFile(relPath, dep, -info.Size()))
			}
		}
	}

	for _, item := range fileItems {
		item.WaitForHashed()
		if err = item.Error(); err != nil {
			return nil, err
		}
		f := i.Files[item.DisplayName]
		f.Digest = item.Digest()
		i.Files[item.DisplayName] = f
	}
	// Avoid duplicated entries in includes.
	// TODO(tandrii): add test to reproduce the problem.
	includesSet := map[isolated.HexDigest]bool{}
	for _, item := range dirItems {
		item.WaitForHashed()
		if err = item.Error(); err != nil {
			return nil, err
		}
		includesSet[item.Digest()] = true
	}
	for digest := range includesSet {
		i.Includes = append(i.Includes, digest)
	}
	// Make the includes list deterministic.
	sort.Sort(i.Includes)

	raw := &bytes.Buffer{}
	if err = json.NewEncoder(raw).Encode(i); err != nil {
		return nil, err
	}

	if err := ioutil.WriteFile(opts.Isolated, raw.Bytes(), 0644); err != nil {
		return nil, err
	}
	return arch.Push(displayName, isolatedclient.NewBytesSource(raw.Bytes()), 0), nil
}
