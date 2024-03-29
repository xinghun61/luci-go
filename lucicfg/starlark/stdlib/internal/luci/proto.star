# Copyright 2019 The LUCI Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

"""Exports proto modules with messages used in generated LUCI configs.

Prefer using this module over loading "@proto//..." modules directly. Proto
paths may change in a backward incompatible way. Using this module gives more
stability.
"""

load('@stdlib//internal/luci/descpb.star', 'lucitypes_descpb')
lucitypes_descpb.register()

load('@proto//go.chromium.org/luci/buildbucket/proto/project_config.proto', _buildbucket_pb='buildbucket')
load('@proto//go.chromium.org/luci/common/proto/config/project_config.proto', _config_pb='config')
load('@proto//go.chromium.org/luci/cq/api/config/v2/cq.proto', _cq_pb='cq.config')
load('@proto//go.chromium.org/luci/logdog/api/config/svcconfig/project.proto', _logdog_pb='svcconfig')
load('@proto//go.chromium.org/luci/milo/api/config/project.proto', _milo_pb='milo')
load('@proto//go.chromium.org/luci/luci_notify/api/config/notify.proto', _notify_pb='notify')
load('@proto//go.chromium.org/luci/scheduler/appengine/messages/config.proto', _scheduler_pb='scheduler.config')

buildbucket_pb = _buildbucket_pb
config_pb = _config_pb
cq_pb = _cq_pb
logdog_pb = _logdog_pb
milo_pb = _milo_pb
notify_pb = _notify_pb
scheduler_pb = _scheduler_pb
