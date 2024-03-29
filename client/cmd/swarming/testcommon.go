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

package main

import (
	"context"
	"net/http"

	"go.chromium.org/luci/common/api/swarming/swarming/v1"
)

type testService struct {
	newTask        func(context.Context, *swarming.SwarmingRpcsNewTaskRequest) (*swarming.SwarmingRpcsTaskRequestMetadata, error)
	countTasks     func(context.Context, float64, ...string) (*swarming.SwarmingRpcsTasksCount, error)
	listTasks      func(context.Context, float64, ...string) (*swarming.SwarmingRpcsTaskList, error)
	cancelTask     func(context.Context, string, *swarming.SwarmingRpcsTaskCancelRequest) (*swarming.SwarmingRpcsCancelResponse, error)
	getTaskResult  func(context.Context, string, bool) (*swarming.SwarmingRpcsTaskResult, error)
	getTaskOutput  func(context.Context, string) (*swarming.SwarmingRpcsTaskOutput, error)
	getTaskOutputs func(context.Context, string, string, *swarming.SwarmingRpcsFilesRef) ([]string, error)
}

func (s testService) Client() *http.Client {
	return nil
}

func (s testService) NewTask(c context.Context, req *swarming.SwarmingRpcsNewTaskRequest) (*swarming.SwarmingRpcsTaskRequestMetadata, error) {
	return s.newTask(c, req)
}

func (s testService) CountTasks(c context.Context, start float64, tags ...string) (*swarming.SwarmingRpcsTasksCount, error) {
	return s.countTasks(c, start, tags...)
}

func (s testService) ListTasks(c context.Context, start float64, tags ...string) (*swarming.SwarmingRpcsTaskList, error) {
	return s.listTasks(c, start, tags...)
}

func (s testService) CancelTask(c context.Context, taskID string, req *swarming.SwarmingRpcsTaskCancelRequest) (*swarming.SwarmingRpcsCancelResponse, error) {
	return s.cancelTask(c, taskID, req)
}

func (s testService) GetTaskResult(c context.Context, taskID string, perf bool) (*swarming.SwarmingRpcsTaskResult, error) {
	return s.getTaskResult(c, taskID, perf)
}

func (s testService) GetTaskOutput(c context.Context, taskID string) (*swarming.SwarmingRpcsTaskOutput, error) {
	return s.getTaskOutput(c, taskID)
}

func (s testService) GetTaskOutputs(c context.Context, taskID, output string, ref *swarming.SwarmingRpcsFilesRef) ([]string, error) {
	return s.getTaskOutputs(c, taskID, output, ref)
}
