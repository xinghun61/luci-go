// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/scheduler/appengine/messages/cron.proto

/*
Package messages is a generated protocol buffer package.

It is generated from these files:
	go.chromium.org/luci/scheduler/appengine/messages/cron.proto

It has these top-level messages:
	Acl
	AclSet
	Job
	Trigger
	NoopTask
	GitilesTask
	UrlFetchTask
	SwarmingTask
	BuildbucketTask
	ProjectConfig
	TaskDefWrapper
*/
package messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Acl_Role int32

const (
	// Can do read-only operations, such as listing invocations of a Job.
	Acl_READER Acl_Role = 0
	// Same as READER + can modify state of a Job or Invocation such as
	// triggering or aborting them.
	// LUCI scheduler (this service) is an OWNER of each `Job` and `Trigger`, thus
	// `Trigger`s are allowed to trigger all `Job`s defined in the same
	// project, regardless of their respective ACLs.
	Acl_OWNER Acl_Role = 1
)

var Acl_Role_name = map[int32]string{
	0: "READER",
	1: "OWNER",
}
var Acl_Role_value = map[string]int32{
	"READER": 0,
	"OWNER":  1,
}

func (x Acl_Role) String() string {
	return proto.EnumName(Acl_Role_name, int32(x))
}
func (Acl_Role) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// A single access control rule.
type Acl struct {
	// Role denotes a list of actions that an identity can perform.
	Role Acl_Role `protobuf:"varint,1,opt,name=role,enum=messages.Acl_Role" json:"role,omitempty"`
	// Either email or "group:xyz" or auth service identity string "kind:name".
	GrantedTo string `protobuf:"bytes,2,opt,name=granted_to,json=grantedTo" json:"granted_to,omitempty"`
}

func (m *Acl) Reset()                    { *m = Acl{} }
func (m *Acl) String() string            { return proto.CompactTextString(m) }
func (*Acl) ProtoMessage()               {}
func (*Acl) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Acl) GetRole() Acl_Role {
	if m != nil {
		return m.Role
	}
	return Acl_READER
}

func (m *Acl) GetGrantedTo() string {
	if m != nil {
		return m.GrantedTo
	}
	return ""
}

// A set of Acl messages. Can be referenced in a Job or Trigger by name.
type AclSet struct {
	// A name of the ACL set, unique for a project.
	// Required. Must match regex '^[0-9A-Za-z_\-\.]{1,100}$'.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// List of access control rules.
	// The order does not matter.
	Acls []*Acl `protobuf:"bytes,2,rep,name=acls" json:"acls,omitempty"`
}

func (m *AclSet) Reset()                    { *m = AclSet{} }
func (m *AclSet) String() string            { return proto.CompactTextString(m) }
func (*AclSet) ProtoMessage()               {}
func (*AclSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AclSet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AclSet) GetAcls() []*Acl {
	if m != nil {
		return m.Acls
	}
	return nil
}

// Job specifies a single regular job belonging to a project.
//
// Such jobs runs on a schedule or can be triggered by some trigger.
type Job struct {
	// Id is a name of the job (unique for the project).
	//
	// Must match '^[0-9A-Za-z_\-\.]{1,100}$'.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Schedule describes when to run the job.
	//
	// Supported kinds of schedules (illustrated by examples):
	//   - "* 0 * * * *": cron-like expression, in a syntax supported by
	//     https://github.com/gorhill/cronexpr (see its docs for full reference).
	//     The cron engine will attempt to start a job at specified moments in
	//     time (based on UTC clock). If when triggering a job, previous
	//     invocation is still running, an overrun will be recorded (and next
	//     attempt to start a job happens based on the schedule, not when the
	//     previous invocation finishes). Some examples:
	//       "0 */3 * * * *" - each 3 hours: at 12:00 AM UTC, 3:00 AM UTC, ...
	//       "0 */3 * * *" - exact same thing (last field is optional)
	//       "0 2,10,18 * * *" - at 2 AM UTC, 10 AM UTC, 6 PM UTC
	//       "0 7 * * *" - at 7 AM UTC, once a day.
	//   - "with 10s interval": runs invocations in a loop, waiting 10s after
	//     finishing invocation before starting a new one. Overruns are not
	//     possible.
	//   - "continuously" is alias for "with 0s interval", meaning the job will
	//     run in a loop without any pauses.
	//   - "triggered" schedule indicates that job is only started via "Run now"
	//     button or via a trigger.
	//
	// Default is "triggered".
	Schedule string `protobuf:"bytes,2,opt,name=schedule" json:"schedule,omitempty"`
	// Disabled is true to disable this job.
	Disabled bool `protobuf:"varint,3,opt,name=disabled" json:"disabled,omitempty"`
	// List of access control rules for the Job.
	// The order does not matter.
	// There can be at most 32 different acls for a Job, including those from
	// acl_sets.
	Acls []*Acl `protobuf:"bytes,5,rep,name=acls" json:"acls,omitempty"`
	// A list of ACL set names. Each ACL in each referenced ACL set will be
	// included in this Job.
	// The order does not matter.
	AclSets []string `protobuf:"bytes,6,rep,name=acl_sets,json=aclSets" json:"acl_sets,omitempty"`
	// Noop is used for testing. It is "do nothing" task.
	Noop *NoopTask `protobuf:"bytes,100,opt,name=noop" json:"noop,omitempty"`
	// UrlFetch can be used to make a simple HTTP call.
	UrlFetch *UrlFetchTask `protobuf:"bytes,101,opt,name=url_fetch,json=urlFetch" json:"url_fetch,omitempty"`
	// SwarmingTask can be used to schedule swarming job.
	Swarming *SwarmingTask `protobuf:"bytes,102,opt,name=swarming" json:"swarming,omitempty"`
	// BuildbucketTask can be used to schedule buildbucket job.
	Buildbucket *BuildbucketTask `protobuf:"bytes,103,opt,name=buildbucket" json:"buildbucket,omitempty"`
}

func (m *Job) Reset()                    { *m = Job{} }
func (m *Job) String() string            { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()               {}
func (*Job) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Job) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Job) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

func (m *Job) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *Job) GetAcls() []*Acl {
	if m != nil {
		return m.Acls
	}
	return nil
}

func (m *Job) GetAclSets() []string {
	if m != nil {
		return m.AclSets
	}
	return nil
}

func (m *Job) GetNoop() *NoopTask {
	if m != nil {
		return m.Noop
	}
	return nil
}

func (m *Job) GetUrlFetch() *UrlFetchTask {
	if m != nil {
		return m.UrlFetch
	}
	return nil
}

func (m *Job) GetSwarming() *SwarmingTask {
	if m != nil {
		return m.Swarming
	}
	return nil
}

func (m *Job) GetBuildbucket() *BuildbucketTask {
	if m != nil {
		return m.Buildbucket
	}
	return nil
}

// Trigger specifies a job that triggers other jobs.
//
// It is a special kind of job that periodically checks the state of the world
// and triggers other jobs.
type Trigger struct {
	// Id is a name of the job (unique for the project).
	//
	// Must match '^[0-9A-Za-z_\-\.]{1,100}$'. It's in the same namespace as
	// regular jobs.
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Schedule describes when to run this triggering job.
	//
	// See Job.schedule fro more info. Default is "with 30s interval".
	Schedule string `protobuf:"bytes,2,opt,name=schedule" json:"schedule,omitempty"`
	// Disabled is true to disable this job.
	Disabled bool `protobuf:"varint,3,opt,name=disabled" json:"disabled,omitempty"`
	// List of access control rules for the Job.
	// The order does not matter.
	// There can be at most 32 different acls for a Job, including those from
	// acl_sets.
	Acls []*Acl `protobuf:"bytes,4,rep,name=acls" json:"acls,omitempty"`
	// A list of ACL set names. Each ACL in each referenced ACL set will be
	// included in this Job.
	// The order does not matter.
	AclSets []string `protobuf:"bytes,5,rep,name=acl_sets,json=aclSets" json:"acl_sets,omitempty"`
	// Noop is used for testing. It is "do nothing" trigger.
	Noop *NoopTask `protobuf:"bytes,100,opt,name=noop" json:"noop,omitempty"`
	// Gitiles is used to trigger jobs for new commits on Gitiles.
	Gitiles *GitilesTask `protobuf:"bytes,101,opt,name=gitiles" json:"gitiles,omitempty"`
	// Triggers are IDs of jobs triggered by this trigger.
	Triggers []string `protobuf:"bytes,200,rep,name=triggers" json:"triggers,omitempty"`
}

func (m *Trigger) Reset()                    { *m = Trigger{} }
func (m *Trigger) String() string            { return proto.CompactTextString(m) }
func (*Trigger) ProtoMessage()               {}
func (*Trigger) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Trigger) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Trigger) GetSchedule() string {
	if m != nil {
		return m.Schedule
	}
	return ""
}

func (m *Trigger) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

func (m *Trigger) GetAcls() []*Acl {
	if m != nil {
		return m.Acls
	}
	return nil
}

func (m *Trigger) GetAclSets() []string {
	if m != nil {
		return m.AclSets
	}
	return nil
}

func (m *Trigger) GetNoop() *NoopTask {
	if m != nil {
		return m.Noop
	}
	return nil
}

func (m *Trigger) GetGitiles() *GitilesTask {
	if m != nil {
		return m.Gitiles
	}
	return nil
}

func (m *Trigger) GetTriggers() []string {
	if m != nil {
		return m.Triggers
	}
	return nil
}

// NoopTask is used for testing. It is "do nothing" task that can emit fake
// triggers.
type NoopTask struct {
	SleepMs       int64 `protobuf:"varint,1,opt,name=sleep_ms,json=sleepMs" json:"sleep_ms,omitempty"`
	TriggersCount int64 `protobuf:"varint,2,opt,name=triggers_count,json=triggersCount" json:"triggers_count,omitempty"`
}

func (m *NoopTask) Reset()                    { *m = NoopTask{} }
func (m *NoopTask) String() string            { return proto.CompactTextString(m) }
func (*NoopTask) ProtoMessage()               {}
func (*NoopTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *NoopTask) GetSleepMs() int64 {
	if m != nil {
		return m.SleepMs
	}
	return 0
}

func (m *NoopTask) GetTriggersCount() int64 {
	if m != nil {
		return m.TriggersCount
	}
	return 0
}

// GitilesTask specifies parameters of what repo and which refs to watch for new
// commits.
type GitilesTask struct {
	// Repo is the URL of the Gitiles repository.
	Repo string `protobuf:"bytes,1,opt,name=repo" json:"repo,omitempty"`
	// List of Git references to track.
	Refs []string `protobuf:"bytes,2,rep,name=refs" json:"refs,omitempty"`
}

func (m *GitilesTask) Reset()                    { *m = GitilesTask{} }
func (m *GitilesTask) String() string            { return proto.CompactTextString(m) }
func (*GitilesTask) ProtoMessage()               {}
func (*GitilesTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GitilesTask) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *GitilesTask) GetRefs() []string {
	if m != nil {
		return m.Refs
	}
	return nil
}

// UrlFetchTask specifies parameters for simple HTTP call.
type UrlFetchTask struct {
	// Method is HTTP method to use, such as "GET" or "POST". Default is "GET".
	Method string `protobuf:"bytes,1,opt,name=method" json:"method,omitempty"`
	// Url to send the request to.
	Url string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	// Timeout is how long to wait for request to complete. Default is 60 sec.
	TimeoutSec int32 `protobuf:"varint,3,opt,name=timeout_sec,json=timeoutSec" json:"timeout_sec,omitempty"`
}

func (m *UrlFetchTask) Reset()                    { *m = UrlFetchTask{} }
func (m *UrlFetchTask) String() string            { return proto.CompactTextString(m) }
func (*UrlFetchTask) ProtoMessage()               {}
func (*UrlFetchTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UrlFetchTask) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *UrlFetchTask) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *UrlFetchTask) GetTimeoutSec() int32 {
	if m != nil {
		return m.TimeoutSec
	}
	return 0
}

// SwarmingTask specifies parameters of Swarming-based jobs.
type SwarmingTask struct {
	// Server is URL of the swarming service to use.
	Server string `protobuf:"bytes,1,opt,name=server" json:"server,omitempty"`
	// What to run. Only one of 'command' or 'isolated_ref' must be given.
	Command     []string                  `protobuf:"bytes,2,rep,name=command" json:"command,omitempty"`
	IsolatedRef *SwarmingTask_IsolatedRef `protobuf:"bytes,3,opt,name=isolated_ref,json=isolatedRef" json:"isolated_ref,omitempty"`
	// Additional arguments to pass to isolated command.
	ExtraArgs []string `protobuf:"bytes,4,rep,name=extra_args,json=extraArgs" json:"extra_args,omitempty"`
	// List of "key=value" pairs with additional OS environment variables.
	Env []string `protobuf:"bytes,5,rep,name=env" json:"env,omitempty"`
	// Where to run it. List of "key:value" pairs.
	Dimensions []string `protobuf:"bytes,6,rep,name=dimensions" json:"dimensions,omitempty"`
	// Tags is a list of tags (as "key:value" pairs) to assign to the task.
	Tags []string `protobuf:"bytes,7,rep,name=tags" json:"tags,omitempty"`
	// Priority is task priority (or niceness, lower value - higher priority).
	Priority int32 `protobuf:"varint,8,opt,name=priority" json:"priority,omitempty"`
	// Timeouts. All optional. The scheduler will set reasonable default values.
	ExecutionTimeoutSecs int32 `protobuf:"varint,9,opt,name=execution_timeout_secs,json=executionTimeoutSecs" json:"execution_timeout_secs,omitempty"`
	GracePeriodSecs      int32 `protobuf:"varint,10,opt,name=grace_period_secs,json=gracePeriodSecs" json:"grace_period_secs,omitempty"`
	IoTimeoutSecs        int32 `protobuf:"varint,11,opt,name=io_timeout_secs,json=ioTimeoutSecs" json:"io_timeout_secs,omitempty"`
}

func (m *SwarmingTask) Reset()                    { *m = SwarmingTask{} }
func (m *SwarmingTask) String() string            { return proto.CompactTextString(m) }
func (*SwarmingTask) ProtoMessage()               {}
func (*SwarmingTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SwarmingTask) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *SwarmingTask) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

func (m *SwarmingTask) GetIsolatedRef() *SwarmingTask_IsolatedRef {
	if m != nil {
		return m.IsolatedRef
	}
	return nil
}

func (m *SwarmingTask) GetExtraArgs() []string {
	if m != nil {
		return m.ExtraArgs
	}
	return nil
}

func (m *SwarmingTask) GetEnv() []string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *SwarmingTask) GetDimensions() []string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *SwarmingTask) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *SwarmingTask) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *SwarmingTask) GetExecutionTimeoutSecs() int32 {
	if m != nil {
		return m.ExecutionTimeoutSecs
	}
	return 0
}

func (m *SwarmingTask) GetGracePeriodSecs() int32 {
	if m != nil {
		return m.GracePeriodSecs
	}
	return 0
}

func (m *SwarmingTask) GetIoTimeoutSecs() int32 {
	if m != nil {
		return m.IoTimeoutSecs
	}
	return 0
}

// IsolatedRef defines a data tree reference, normally a reference to
// an .isolated file
type SwarmingTask_IsolatedRef struct {
	Isolated       string `protobuf:"bytes,1,opt,name=isolated" json:"isolated,omitempty"`
	IsolatedServer string `protobuf:"bytes,2,opt,name=isolated_server,json=isolatedServer" json:"isolated_server,omitempty"`
	Namespace      string `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
}

func (m *SwarmingTask_IsolatedRef) Reset()                    { *m = SwarmingTask_IsolatedRef{} }
func (m *SwarmingTask_IsolatedRef) String() string            { return proto.CompactTextString(m) }
func (*SwarmingTask_IsolatedRef) ProtoMessage()               {}
func (*SwarmingTask_IsolatedRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7, 0} }

func (m *SwarmingTask_IsolatedRef) GetIsolated() string {
	if m != nil {
		return m.Isolated
	}
	return ""
}

func (m *SwarmingTask_IsolatedRef) GetIsolatedServer() string {
	if m != nil {
		return m.IsolatedServer
	}
	return ""
}

func (m *SwarmingTask_IsolatedRef) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

// BuildbucketTask specifies parameters of Buildbucket-based jobs.
type BuildbucketTask struct {
	// Server is URL of the bulildbucket service to use.
	Server string `protobuf:"bytes,1,opt,name=server" json:"server,omitempty"`
	// Bucket defines what bucket to add the task to.
	Bucket string `protobuf:"bytes,2,opt,name=bucket" json:"bucket,omitempty"`
	// Builder defines what to run.
	Builder string `protobuf:"bytes,3,opt,name=builder" json:"builder,omitempty"`
	// Properties is arbitrary "key:value" pairs describing the task.
	// TODO(tandrii): which properties will be overridden if triggered?
	Properties []string `protobuf:"bytes,4,rep,name=properties" json:"properties,omitempty"`
	// Tags is a list of tags (as "key:value" pairs) to assign to the task.
	Tags []string `protobuf:"bytes,5,rep,name=tags" json:"tags,omitempty"`
}

func (m *BuildbucketTask) Reset()                    { *m = BuildbucketTask{} }
func (m *BuildbucketTask) String() string            { return proto.CompactTextString(m) }
func (*BuildbucketTask) ProtoMessage()               {}
func (*BuildbucketTask) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *BuildbucketTask) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *BuildbucketTask) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *BuildbucketTask) GetBuilder() string {
	if m != nil {
		return m.Builder
	}
	return ""
}

func (m *BuildbucketTask) GetProperties() []string {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *BuildbucketTask) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// ProjectConfig defines a schema for config file that describe jobs belonging
// to some project.
type ProjectConfig struct {
	// Job is a set of jobs defined in the project.
	Job []*Job `protobuf:"bytes,1,rep,name=job" json:"job,omitempty"`
	// Trigger is a set of triggering jobs defined in the project.
	Trigger []*Trigger `protobuf:"bytes,2,rep,name=trigger" json:"trigger,omitempty"`
	// A list of ACL sets. Names must be unique.
	AclSets []*AclSet `protobuf:"bytes,3,rep,name=acl_sets,json=aclSets" json:"acl_sets,omitempty"`
}

func (m *ProjectConfig) Reset()                    { *m = ProjectConfig{} }
func (m *ProjectConfig) String() string            { return proto.CompactTextString(m) }
func (*ProjectConfig) ProtoMessage()               {}
func (*ProjectConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ProjectConfig) GetJob() []*Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *ProjectConfig) GetTrigger() []*Trigger {
	if m != nil {
		return m.Trigger
	}
	return nil
}

func (m *ProjectConfig) GetAclSets() []*AclSet {
	if m != nil {
		return m.AclSets
	}
	return nil
}

// TaskDefWrapper is a union type of all possible tasks known to the scheduler.
//
// It is used internally when storing jobs in the datastore.
type TaskDefWrapper struct {
	Noop        *NoopTask        `protobuf:"bytes,1,opt,name=noop" json:"noop,omitempty"`
	UrlFetch    *UrlFetchTask    `protobuf:"bytes,2,opt,name=url_fetch,json=urlFetch" json:"url_fetch,omitempty"`
	Swarming    *SwarmingTask    `protobuf:"bytes,3,opt,name=swarming" json:"swarming,omitempty"`
	Buildbucket *BuildbucketTask `protobuf:"bytes,4,opt,name=buildbucket" json:"buildbucket,omitempty"`
	Gitiles     *GitilesTask     `protobuf:"bytes,5,opt,name=gitiles" json:"gitiles,omitempty"`
}

func (m *TaskDefWrapper) Reset()                    { *m = TaskDefWrapper{} }
func (m *TaskDefWrapper) String() string            { return proto.CompactTextString(m) }
func (*TaskDefWrapper) ProtoMessage()               {}
func (*TaskDefWrapper) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *TaskDefWrapper) GetNoop() *NoopTask {
	if m != nil {
		return m.Noop
	}
	return nil
}

func (m *TaskDefWrapper) GetUrlFetch() *UrlFetchTask {
	if m != nil {
		return m.UrlFetch
	}
	return nil
}

func (m *TaskDefWrapper) GetSwarming() *SwarmingTask {
	if m != nil {
		return m.Swarming
	}
	return nil
}

func (m *TaskDefWrapper) GetBuildbucket() *BuildbucketTask {
	if m != nil {
		return m.Buildbucket
	}
	return nil
}

func (m *TaskDefWrapper) GetGitiles() *GitilesTask {
	if m != nil {
		return m.Gitiles
	}
	return nil
}

func init() {
	proto.RegisterType((*Acl)(nil), "messages.Acl")
	proto.RegisterType((*AclSet)(nil), "messages.AclSet")
	proto.RegisterType((*Job)(nil), "messages.Job")
	proto.RegisterType((*Trigger)(nil), "messages.Trigger")
	proto.RegisterType((*NoopTask)(nil), "messages.NoopTask")
	proto.RegisterType((*GitilesTask)(nil), "messages.GitilesTask")
	proto.RegisterType((*UrlFetchTask)(nil), "messages.UrlFetchTask")
	proto.RegisterType((*SwarmingTask)(nil), "messages.SwarmingTask")
	proto.RegisterType((*SwarmingTask_IsolatedRef)(nil), "messages.SwarmingTask.IsolatedRef")
	proto.RegisterType((*BuildbucketTask)(nil), "messages.BuildbucketTask")
	proto.RegisterType((*ProjectConfig)(nil), "messages.ProjectConfig")
	proto.RegisterType((*TaskDefWrapper)(nil), "messages.TaskDefWrapper")
	proto.RegisterEnum("messages.Acl_Role", Acl_Role_name, Acl_Role_value)
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/scheduler/appengine/messages/cron.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 920 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x5d, 0x8f, 0xdb, 0x44,
	0x14, 0xc5, 0x71, 0x3e, 0xec, 0x9b, 0xdd, 0x6c, 0x3a, 0x82, 0x95, 0x5b, 0x28, 0x0d, 0x96, 0x28,
	0x11, 0x95, 0x12, 0x69, 0x0b, 0x4f, 0x20, 0xa1, 0xd0, 0x2e, 0x88, 0x15, 0x94, 0x6a, 0xb2, 0xa8,
	0xe2, 0x29, 0x72, 0xec, 0x1b, 0xef, 0x74, 0x6d, 0x8f, 0x35, 0x33, 0x2e, 0xe5, 0x1d, 0x5e, 0x78,
	0xe9, 0xdf, 0xe2, 0x5f, 0xf0, 0x57, 0xd0, 0x8c, 0xc7, 0x8e, 0x77, 0x45, 0xa5, 0xee, 0x8a, 0xb7,
	0xb9, 0xe7, 0x9e, 0x3b, 0x1f, 0xe7, 0xde, 0x1c, 0x07, 0xbe, 0x4e, 0xf9, 0x22, 0xbe, 0x10, 0x3c,
	0x67, 0x55, 0xbe, 0xe0, 0x22, 0x5d, 0x66, 0x55, 0xcc, 0x96, 0x32, 0xbe, 0xc0, 0xa4, 0xca, 0x50,
	0x2c, 0xa3, 0xb2, 0xc4, 0x22, 0x65, 0x05, 0x2e, 0x73, 0x94, 0x32, 0x4a, 0x51, 0x2e, 0x63, 0xc1,
	0x8b, 0x45, 0x29, 0xb8, 0xe2, 0xc4, 0x6b, 0xc0, 0xf0, 0x12, 0xdc, 0x55, 0x9c, 0x91, 0x87, 0xd0,
	0x17, 0x3c, 0xc3, 0xc0, 0x99, 0x39, 0xf3, 0xc9, 0x09, 0x59, 0x34, 0xf9, 0xc5, 0x2a, 0xce, 0x16,
	0x94, 0x67, 0x48, 0x4d, 0x9e, 0xdc, 0x07, 0x48, 0x45, 0x54, 0x28, 0x4c, 0x36, 0x8a, 0x07, 0xbd,
	0x99, 0x33, 0xf7, 0xa9, 0x6f, 0x91, 0x73, 0x1e, 0xde, 0x87, 0xbe, 0x26, 0x13, 0x80, 0x21, 0x3d,
	0x5d, 0x3d, 0x3d, 0xa5, 0xd3, 0xf7, 0x88, 0x0f, 0x83, 0x9f, 0x5f, 0x3c, 0x3b, 0xa5, 0x53, 0x27,
	0xfc, 0x06, 0x86, 0xab, 0x38, 0x5b, 0xa3, 0x22, 0x04, 0xfa, 0x45, 0x94, 0xd7, 0xe7, 0xf9, 0xd4,
	0xac, 0xc9, 0x27, 0xd0, 0x8f, 0xe2, 0x4c, 0x06, 0xbd, 0x99, 0x3b, 0x1f, 0x9f, 0x1c, 0x5e, 0xb9,
	0x03, 0x35, 0xa9, 0xf0, 0x9f, 0x1e, 0xb8, 0x67, 0x7c, 0x4b, 0x26, 0xd0, 0x63, 0x89, 0x2d, 0xee,
	0xb1, 0x84, 0xdc, 0x03, 0xaf, 0x79, 0xba, 0xbd, 0x54, 0x1b, 0xeb, 0x5c, 0xc2, 0x64, 0xb4, 0xcd,
	0x30, 0x09, 0xdc, 0x99, 0x33, 0xf7, 0x68, 0x1b, 0xb7, 0x47, 0x0e, 0xde, 0x7a, 0x24, 0xb9, 0x0b,
	0x5e, 0x14, 0x67, 0x1b, 0x89, 0x4a, 0x06, 0xc3, 0x99, 0x3b, 0xf7, 0xe9, 0x28, 0x32, 0x6f, 0x90,
	0x5a, 0xb4, 0x82, 0xf3, 0x32, 0x48, 0x66, 0xce, 0x7c, 0xdc, 0x15, 0xed, 0x19, 0xe7, 0xe5, 0x79,
	0x24, 0x2f, 0xa9, 0xc9, 0x93, 0xc7, 0xe0, 0x57, 0x22, 0xdb, 0xec, 0x50, 0xc5, 0x17, 0x01, 0x1a,
	0xf2, 0xf1, 0x9e, 0xfc, 0x8b, 0xc8, 0xbe, 0xd3, 0x19, 0x53, 0xe0, 0x55, 0x36, 0x22, 0x27, 0xe0,
	0xc9, 0xdf, 0x22, 0x91, 0xb3, 0x22, 0x0d, 0x76, 0xd7, 0x6b, 0xd6, 0x36, 0x53, 0xd7, 0x34, 0x3c,
	0xf2, 0x15, 0x8c, 0xb7, 0x15, 0xcb, 0x92, 0x6d, 0x15, 0x5f, 0xa2, 0x0a, 0x52, 0x53, 0x76, 0x77,
	0x5f, 0xf6, 0xed, 0x3e, 0x69, 0x2a, 0xbb, 0xec, 0xb3, 0xbe, 0xd7, 0x9f, 0x0e, 0xc2, 0x3f, 0x7b,
	0x30, 0x3a, 0x17, 0x2c, 0x4d, 0x51, 0xfc, 0xef, 0x2a, 0xf7, 0xdf, 0x4d, 0xe5, 0xc1, 0xed, 0x54,
	0x5e, 0xc2, 0x28, 0x65, 0x8a, 0x65, 0x28, 0xad, 0xc6, 0x1f, 0xec, 0xa9, 0xdf, 0xd7, 0x09, 0xc3,
	0x6e, 0x58, 0xe4, 0x43, 0xf0, 0x54, 0xfd, 0x52, 0x19, 0xfc, 0xed, 0x98, 0x43, 0x5b, 0x20, 0xfc,
	0x11, 0xbc, 0x66, 0x7f, 0x7d, 0x39, 0x99, 0x21, 0x96, 0x9b, 0x5c, 0x1a, 0x35, 0x5c, 0x3a, 0x32,
	0xf1, 0x4f, 0x92, 0x7c, 0x0a, 0x93, 0xa6, 0x64, 0x13, 0xf3, 0xaa, 0x50, 0x46, 0x18, 0x97, 0x1e,
	0x36, 0xe8, 0x13, 0x0d, 0x86, 0x5f, 0xc2, 0xb8, 0x73, 0x05, 0x3d, 0xfd, 0x02, 0x4b, 0xde, 0x4c,
	0xbf, 0x5e, 0xd7, 0xd8, 0xae, 0x9e, 0x7e, 0x83, 0xed, 0x64, 0xf8, 0x2b, 0x1c, 0x74, 0xa7, 0x83,
	0x1c, 0xc3, 0x30, 0x47, 0x75, 0xc1, 0x9b, 0xa6, 0xd8, 0x88, 0x4c, 0xc1, 0xad, 0x44, 0x66, 0x7b,
	0xa2, 0x97, 0xe4, 0x01, 0x8c, 0x15, 0xcb, 0x91, 0x57, 0x6a, 0x23, 0x31, 0x36, 0x1d, 0x19, 0x50,
	0xb0, 0xd0, 0x1a, 0xe3, 0xf0, 0x8f, 0x3e, 0x1c, 0x74, 0xa7, 0x48, 0xef, 0x2d, 0x51, 0xbc, 0x42,
	0xd1, 0xec, 0x5d, 0x47, 0x24, 0x80, 0x51, 0xcc, 0xf3, 0x3c, 0x2a, 0x12, 0x7b, 0xb5, 0x26, 0x24,
	0xa7, 0x70, 0xc0, 0x24, 0xcf, 0x22, 0x6d, 0x06, 0x02, 0x77, 0xe6, 0x90, 0xf1, 0x49, 0xf8, 0xdf,
	0x53, 0xba, 0xf8, 0xc1, 0x52, 0x29, 0xee, 0xe8, 0x98, 0xed, 0x03, 0x6d, 0x29, 0xf8, 0x5a, 0x89,
	0x68, 0x13, 0x89, 0xb4, 0x9e, 0x11, 0x9f, 0xfa, 0x06, 0x59, 0x89, 0x54, 0xea, 0xb7, 0x61, 0xf1,
	0xca, 0x0e, 0x85, 0x5e, 0x92, 0x8f, 0x01, 0x12, 0x96, 0x63, 0x21, 0x19, 0x2f, 0x9a, 0xdf, 0x64,
	0x07, 0xd1, 0x4a, 0xaa, 0x28, 0x95, 0xc1, 0xa8, 0x56, 0x52, 0xaf, 0xf5, 0x78, 0x96, 0x82, 0x71,
	0xc1, 0xd4, 0xef, 0x81, 0x67, 0xc4, 0x68, 0x63, 0xf2, 0x05, 0x1c, 0xe3, 0x6b, 0x8c, 0x2b, 0xc5,
	0x78, 0xb1, 0xe9, 0xa8, 0x26, 0x03, 0xdf, 0x30, 0xdf, 0x6f, 0xb3, 0xe7, 0xad, 0x7e, 0x92, 0x7c,
	0x0e, 0x77, 0x52, 0x11, 0xc5, 0xb8, 0x29, 0x51, 0x30, 0x9e, 0xd4, 0x05, 0x60, 0x0a, 0x8e, 0x4c,
	0xe2, 0xb9, 0xc1, 0x0d, 0xf7, 0x21, 0x1c, 0x31, 0x7e, 0x75, 0xeb, 0xb1, 0x61, 0x1e, 0x32, 0xde,
	0xd9, 0xf3, 0x5e, 0x09, 0xe3, 0x8e, 0x4c, 0xfa, 0xd2, 0x8d, 0x50, 0xb6, 0x29, 0x6d, 0x4c, 0x3e,
	0x83, 0xa3, 0x56, 0x7c, 0xdb, 0xb7, 0xba, 0xfd, 0x93, 0x06, 0x5e, 0xd7, 0xfd, 0xfb, 0x08, 0x7c,
	0xed, 0xae, 0xb2, 0x8c, 0x62, 0x34, 0x2d, 0xf2, 0xe9, 0x1e, 0x08, 0xdf, 0x38, 0x70, 0x74, 0xcd,
	0x15, 0xde, 0x3a, 0x09, 0xc7, 0x30, 0xb4, 0xc6, 0x52, 0x9f, 0x64, 0x23, 0x3d, 0x21, 0xc6, 0x47,
	0x50, 0xd8, 0xfd, 0x9b, 0x50, 0x77, 0xaa, 0x14, 0xbc, 0x44, 0xa1, 0x18, 0x36, 0xad, 0xed, 0x20,
	0x6d, 0xa7, 0x06, 0xfb, 0x4e, 0x85, 0x7f, 0x39, 0x70, 0xf8, 0x5c, 0xf0, 0x97, 0x18, 0xab, 0x27,
	0xbc, 0xd8, 0xb1, 0x94, 0x3c, 0x00, 0xf7, 0x25, 0xdf, 0x06, 0xce, 0x75, 0xf7, 0x38, 0xe3, 0x5b,
	0xaa, 0x33, 0xe4, 0x11, 0x8c, 0xec, 0xcf, 0xcd, 0x7e, 0x3b, 0xee, 0xec, 0x49, 0xd6, 0xcb, 0x68,
	0xc3, 0x20, 0x8f, 0x3a, 0x4e, 0xe3, 0x1a, 0xf6, 0xf4, 0x8a, 0x21, 0xad, 0x51, 0xb5, 0xde, 0x13,
	0xbe, 0xe9, 0xc1, 0x44, 0x6b, 0xf2, 0x14, 0x77, 0x2f, 0x84, 0xfe, 0xa4, 0x8a, 0xd6, 0x8e, 0x9c,
	0x9b, 0x98, 0x7e, 0xef, 0x16, 0xa6, 0xef, 0xde, 0xce, 0xf4, 0xfb, 0x37, 0x31, 0xfd, 0xae, 0x69,
	0x0e, 0xde, 0xc5, 0x34, 0xb7, 0x43, 0xf3, 0x07, 0xe2, 0xf1, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x91, 0x86, 0x1c, 0x1c, 0x80, 0x08, 0x00, 0x00,
}
