// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/swarming/proto/config/config.proto

package configpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "go.chromium.org/luci/common/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Schema for settings.cfg service config file in luci-config.
type SettingsCfg struct {
	// id to inject into pages if applicable.
	GoogleAnalytics string `protobuf:"bytes,1,opt,name=google_analytics,json=googleAnalytics,proto3" json:"google_analytics,omitempty"`
	// The number of seconds an old task can be deduped from.
	// Default is one week: 7*24*60*60 = 604800
	ReusableTaskAgeSecs int32 `protobuf:"varint,2,opt,name=reusable_task_age_secs,json=reusableTaskAgeSecs,proto3" json:"reusable_task_age_secs,omitempty"`
	// The amount of time that has to pass before a machine is considered dead.
	// Default is 600 (10 minutes).
	BotDeathTimeoutSecs int32 `protobuf:"varint,3,opt,name=bot_death_timeout_secs,json=botDeathTimeoutSecs,proto3" json:"bot_death_timeout_secs,omitempty"`
	// Enable ts_mon based monitoring.
	EnableTsMonitoring bool `protobuf:"varint,4,opt,name=enable_ts_monitoring,json=enableTsMonitoring,proto3" json:"enable_ts_monitoring,omitempty"`
	// (deprecated, see pools.proto) Configuration for swarming-isolate integration.
	Isolate *IsolateSettings `protobuf:"bytes,5,opt,name=isolate,proto3" json:"isolate,omitempty"`
	// (deprecated, see pools.proto) Configuration for swarming-cipd integration.
	Cipd *CipdSettings `protobuf:"bytes,6,opt,name=cipd,proto3" json:"cipd,omitempty"`
	// Emergency setting to disable bot task reaping. When set, all bots are
	// always put to sleep and are never granted task.
	ForceBotsToSleepAndNotRunTask bool `protobuf:"varint,8,opt,name=force_bots_to_sleep_and_not_run_task,json=forceBotsToSleepAndNotRunTask,proto3" json:"force_bots_to_sleep_and_not_run_task,omitempty"`
	// oauth client id for the ui. This is created in the developer's console
	// under Credentials.
	UiClientId string `protobuf:"bytes,9,opt,name=ui_client_id,json=uiClientId,proto3" json:"ui_client_id,omitempty"`
	// A url to a task display server (e.g. milo).  This should have a %s where
	// a task id can go.
	DisplayServerUrlTemplate string `protobuf:"bytes,11,opt,name=display_server_url_template,json=displayServerUrlTemplate,proto3" json:"display_server_url_template,omitempty"`
	// Sets a maximum sleep time in seconds for bots that limits the exponental
	// backoff. If missing, the task scheduler will provide the default maximum
	// (usually 60s, but see bot_code/task_scheduler.py for details).
	MaxBotSleepTime int32 `protobuf:"varint,12,opt,name=max_bot_sleep_time,json=maxBotSleepTime,proto3" json:"max_bot_sleep_time,omitempty"`
	// Names of the authorization groups used by components/auth.
	Auth *AuthSettings `protobuf:"bytes,13,opt,name=auth,proto3" json:"auth,omitempty"`
	// Sets the default gRPC proxy for the bot's Isolate server calls.
	BotIsolateGrpcProxy string `protobuf:"bytes,14,opt,name=bot_isolate_grpc_proxy,json=botIsolateGrpcProxy,proto3" json:"bot_isolate_grpc_proxy,omitempty"`
	// Sets the default gRPC proxy for the bot's Swarming server calls.
	BotSwarmingGrpcProxy string `protobuf:"bytes,15,opt,name=bot_swarming_grpc_proxy,json=botSwarmingGrpcProxy,proto3" json:"bot_swarming_grpc_proxy,omitempty"`
	// Any extra urls that should be added to frame-src, e.g. anything that
	// will be linked to from the display server.
	// This originally added things to child-src, which was deprecated:
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy/child-src
	ExtraChildSrcCspUrl []string `protobuf:"bytes,16,rep,name=extra_child_src_csp_url,json=extraChildSrcCspUrl,proto3" json:"extra_child_src_csp_url,omitempty"`
	// Whether tasks should be run in FIFO or LIFO order.
	UseLifo              bool     `protobuf:"varint,17,opt,name=use_lifo,json=useLifo,proto3" json:"use_lifo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettingsCfg) Reset()         { *m = SettingsCfg{} }
func (m *SettingsCfg) String() string { return proto.CompactTextString(m) }
func (*SettingsCfg) ProtoMessage()    {}
func (*SettingsCfg) Descriptor() ([]byte, []int) {
	return fileDescriptor_1faffb86b0af733f, []int{0}
}

func (m *SettingsCfg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettingsCfg.Unmarshal(m, b)
}
func (m *SettingsCfg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettingsCfg.Marshal(b, m, deterministic)
}
func (m *SettingsCfg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingsCfg.Merge(m, src)
}
func (m *SettingsCfg) XXX_Size() int {
	return xxx_messageInfo_SettingsCfg.Size(m)
}
func (m *SettingsCfg) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingsCfg.DiscardUnknown(m)
}

var xxx_messageInfo_SettingsCfg proto.InternalMessageInfo

func (m *SettingsCfg) GetGoogleAnalytics() string {
	if m != nil {
		return m.GoogleAnalytics
	}
	return ""
}

func (m *SettingsCfg) GetReusableTaskAgeSecs() int32 {
	if m != nil {
		return m.ReusableTaskAgeSecs
	}
	return 0
}

func (m *SettingsCfg) GetBotDeathTimeoutSecs() int32 {
	if m != nil {
		return m.BotDeathTimeoutSecs
	}
	return 0
}

func (m *SettingsCfg) GetEnableTsMonitoring() bool {
	if m != nil {
		return m.EnableTsMonitoring
	}
	return false
}

func (m *SettingsCfg) GetIsolate() *IsolateSettings {
	if m != nil {
		return m.Isolate
	}
	return nil
}

func (m *SettingsCfg) GetCipd() *CipdSettings {
	if m != nil {
		return m.Cipd
	}
	return nil
}

func (m *SettingsCfg) GetForceBotsToSleepAndNotRunTask() bool {
	if m != nil {
		return m.ForceBotsToSleepAndNotRunTask
	}
	return false
}

func (m *SettingsCfg) GetUiClientId() string {
	if m != nil {
		return m.UiClientId
	}
	return ""
}

func (m *SettingsCfg) GetDisplayServerUrlTemplate() string {
	if m != nil {
		return m.DisplayServerUrlTemplate
	}
	return ""
}

func (m *SettingsCfg) GetMaxBotSleepTime() int32 {
	if m != nil {
		return m.MaxBotSleepTime
	}
	return 0
}

func (m *SettingsCfg) GetAuth() *AuthSettings {
	if m != nil {
		return m.Auth
	}
	return nil
}

func (m *SettingsCfg) GetBotIsolateGrpcProxy() string {
	if m != nil {
		return m.BotIsolateGrpcProxy
	}
	return ""
}

func (m *SettingsCfg) GetBotSwarmingGrpcProxy() string {
	if m != nil {
		return m.BotSwarmingGrpcProxy
	}
	return ""
}

func (m *SettingsCfg) GetExtraChildSrcCspUrl() []string {
	if m != nil {
		return m.ExtraChildSrcCspUrl
	}
	return nil
}

func (m *SettingsCfg) GetUseLifo() bool {
	if m != nil {
		return m.UseLifo
	}
	return false
}

// Configuration for swarming-isolate integration.
type IsolateSettings struct {
	// URL of the default isolate server to use if it is not specified in a
	// task. Must start with "https://" or "http://",
	// e.g. "https://isolateserver.appspot.com"
	DefaultServer string `protobuf:"bytes,1,opt,name=default_server,json=defaultServer,proto3" json:"default_server,omitempty"`
	// Default namespace to use if it is not specified in a task,
	// e.g. "default-gzip"
	DefaultNamespace     string   `protobuf:"bytes,2,opt,name=default_namespace,json=defaultNamespace,proto3" json:"default_namespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsolateSettings) Reset()         { *m = IsolateSettings{} }
func (m *IsolateSettings) String() string { return proto.CompactTextString(m) }
func (*IsolateSettings) ProtoMessage()    {}
func (*IsolateSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_1faffb86b0af733f, []int{1}
}

func (m *IsolateSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsolateSettings.Unmarshal(m, b)
}
func (m *IsolateSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsolateSettings.Marshal(b, m, deterministic)
}
func (m *IsolateSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsolateSettings.Merge(m, src)
}
func (m *IsolateSettings) XXX_Size() int {
	return xxx_messageInfo_IsolateSettings.Size(m)
}
func (m *IsolateSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_IsolateSettings.DiscardUnknown(m)
}

var xxx_messageInfo_IsolateSettings proto.InternalMessageInfo

func (m *IsolateSettings) GetDefaultServer() string {
	if m != nil {
		return m.DefaultServer
	}
	return ""
}

func (m *IsolateSettings) GetDefaultNamespace() string {
	if m != nil {
		return m.DefaultNamespace
	}
	return ""
}

// A CIPD package.
type CipdPackage struct {
	// A template of a full CIPD package name, e.g.
	// "infra/tools/cipd/${platform}"
	// See also cipd.ALL_PARAMS.
	PackageName string `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	// Valid package version for all packages matched by package name.
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CipdPackage) Reset()         { *m = CipdPackage{} }
func (m *CipdPackage) String() string { return proto.CompactTextString(m) }
func (*CipdPackage) ProtoMessage()    {}
func (*CipdPackage) Descriptor() ([]byte, []int) {
	return fileDescriptor_1faffb86b0af733f, []int{2}
}

func (m *CipdPackage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CipdPackage.Unmarshal(m, b)
}
func (m *CipdPackage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CipdPackage.Marshal(b, m, deterministic)
}
func (m *CipdPackage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CipdPackage.Merge(m, src)
}
func (m *CipdPackage) XXX_Size() int {
	return xxx_messageInfo_CipdPackage.Size(m)
}
func (m *CipdPackage) XXX_DiscardUnknown() {
	xxx_messageInfo_CipdPackage.DiscardUnknown(m)
}

var xxx_messageInfo_CipdPackage proto.InternalMessageInfo

func (m *CipdPackage) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *CipdPackage) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Settings for Swarming-CIPD integration.
type CipdSettings struct {
	// URL of the default CIPD server to use if it is not specified in a task.
	// Must start with "https://" or "http://",
	// e.g. "https://chrome-infra-packages.appspot.com".
	DefaultServer string `protobuf:"bytes,1,opt,name=default_server,json=defaultServer,proto3" json:"default_server,omitempty"`
	// Package of the default CIPD client to use if it is not specified in a
	// task.
	DefaultClientPackage *CipdPackage `protobuf:"bytes,2,opt,name=default_client_package,json=defaultClientPackage,proto3" json:"default_client_package,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CipdSettings) Reset()         { *m = CipdSettings{} }
func (m *CipdSettings) String() string { return proto.CompactTextString(m) }
func (*CipdSettings) ProtoMessage()    {}
func (*CipdSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_1faffb86b0af733f, []int{3}
}

func (m *CipdSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CipdSettings.Unmarshal(m, b)
}
func (m *CipdSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CipdSettings.Marshal(b, m, deterministic)
}
func (m *CipdSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CipdSettings.Merge(m, src)
}
func (m *CipdSettings) XXX_Size() int {
	return xxx_messageInfo_CipdSettings.Size(m)
}
func (m *CipdSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_CipdSettings.DiscardUnknown(m)
}

var xxx_messageInfo_CipdSettings proto.InternalMessageInfo

func (m *CipdSettings) GetDefaultServer() string {
	if m != nil {
		return m.DefaultServer
	}
	return ""
}

func (m *CipdSettings) GetDefaultClientPackage() *CipdPackage {
	if m != nil {
		return m.DefaultClientPackage
	}
	return nil
}

type AuthSettings struct {
	// Members of this group have full administrative access.
	//
	// Grants:
	// - config view and edit
	// - delete any bot
	// - all of bot_bootstrap_group membership
	// - all of privileged_users_group membership
	AdminsGroup string `protobuf:"bytes,1,opt,name=admins_group,json=adminsGroup,proto3" json:"admins_group,omitempty"`
	// Members of this group can fetch swarming bot code and bootstrap bots.
	//
	// Grants:
	// - bot create: create a token to anonymously fetch the bot code.
	BotBootstrapGroup string `protobuf:"bytes,2,opt,name=bot_bootstrap_group,json=botBootstrapGroup,proto3" json:"bot_bootstrap_group,omitempty"`
	// Members of this group can schedule tasks and see everyone else's tasks.
	//
	// Grants:
	// - cancel any task
	// - edit (terminate) any bot
	// - all of view_all_bots_group membership
	// - all of view_all_tasks_group membership
	PrivilegedUsersGroup string `protobuf:"bytes,3,opt,name=privileged_users_group,json=privilegedUsersGroup,proto3" json:"privileged_users_group,omitempty"`
	// Members of this group can schedule tasks and see only their own tasks.
	//
	// Grants:
	// - create a task
	// - view and edit own task
	UsersGroup string `protobuf:"bytes,4,opt,name=users_group,json=usersGroup,proto3" json:"users_group,omitempty"`
	// Members of this group can view all bots. This is a read-only group.
	//
	// Grants:
	// - view all bots
	ViewAllBotsGroup string `protobuf:"bytes,5,opt,name=view_all_bots_group,json=viewAllBotsGroup,proto3" json:"view_all_bots_group,omitempty"`
	// Members of this group can view all tasks. This is a read-only group.
	//
	// Grants:
	// - view all tasks
	ViewAllTasksGroup    string   `protobuf:"bytes,6,opt,name=view_all_tasks_group,json=viewAllTasksGroup,proto3" json:"view_all_tasks_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthSettings) Reset()         { *m = AuthSettings{} }
func (m *AuthSettings) String() string { return proto.CompactTextString(m) }
func (*AuthSettings) ProtoMessage()    {}
func (*AuthSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_1faffb86b0af733f, []int{4}
}

func (m *AuthSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthSettings.Unmarshal(m, b)
}
func (m *AuthSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthSettings.Marshal(b, m, deterministic)
}
func (m *AuthSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthSettings.Merge(m, src)
}
func (m *AuthSettings) XXX_Size() int {
	return xxx_messageInfo_AuthSettings.Size(m)
}
func (m *AuthSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthSettings.DiscardUnknown(m)
}

var xxx_messageInfo_AuthSettings proto.InternalMessageInfo

func (m *AuthSettings) GetAdminsGroup() string {
	if m != nil {
		return m.AdminsGroup
	}
	return ""
}

func (m *AuthSettings) GetBotBootstrapGroup() string {
	if m != nil {
		return m.BotBootstrapGroup
	}
	return ""
}

func (m *AuthSettings) GetPrivilegedUsersGroup() string {
	if m != nil {
		return m.PrivilegedUsersGroup
	}
	return ""
}

func (m *AuthSettings) GetUsersGroup() string {
	if m != nil {
		return m.UsersGroup
	}
	return ""
}

func (m *AuthSettings) GetViewAllBotsGroup() string {
	if m != nil {
		return m.ViewAllBotsGroup
	}
	return ""
}

func (m *AuthSettings) GetViewAllTasksGroup() string {
	if m != nil {
		return m.ViewAllTasksGroup
	}
	return ""
}

func init() {
	proto.RegisterType((*SettingsCfg)(nil), "swarming.config.SettingsCfg")
	proto.RegisterType((*IsolateSettings)(nil), "swarming.config.IsolateSettings")
	proto.RegisterType((*CipdPackage)(nil), "swarming.config.CipdPackage")
	proto.RegisterType((*CipdSettings)(nil), "swarming.config.CipdSettings")
	proto.RegisterType((*AuthSettings)(nil), "swarming.config.AuthSettings")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/swarming/proto/config/config.proto", fileDescriptor_1faffb86b0af733f)
}

var fileDescriptor_1faffb86b0af733f = []byte{
	// 860 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4d, 0x6f, 0x1b, 0x37,
	0x10, 0x85, 0x62, 0xd9, 0x96, 0x29, 0x25, 0x92, 0x19, 0xc3, 0xd9, 0x7e, 0x04, 0x55, 0xd5, 0x16,
	0x70, 0x11, 0x44, 0xdb, 0xc6, 0x29, 0x0a, 0xb8, 0xe8, 0x41, 0x56, 0xd1, 0x34, 0x6e, 0x1b, 0x04,
	0x92, 0x7c, 0xe9, 0x85, 0xe0, 0x72, 0xa9, 0x15, 0x61, 0x2e, 0x87, 0x20, 0xb9, 0x8e, 0x7d, 0xec,
	0xb1, 0xff, 0xa0, 0xe8, 0xb1, 0x3f, 0xb4, 0x28, 0xf8, 0xb1, 0xb6, 0x91, 0xa6, 0x40, 0x4e, 0x2b,
	0xcd, 0x7b, 0x6f, 0x38, 0x7c, 0x33, 0xb3, 0x8b, 0xbe, 0xad, 0x60, 0xca, 0x36, 0x06, 0x6a, 0xd1,
	0xd4, 0x53, 0x30, 0x55, 0x2e, 0x1b, 0x26, 0x72, 0xfb, 0x86, 0x9a, 0x5a, 0xa8, 0x2a, 0xd7, 0x06,
	0x1c, 0xe4, 0x0c, 0xd4, 0x5a, 0x54, 0xe9, 0x31, 0x0d, 0x31, 0x3c, 0x6c, 0x39, 0xd3, 0x18, 0xfe,
	0x30, 0x7f, 0x67, 0x26, 0x06, 0x75, 0x0d, 0x2a, 0xe5, 0x01, 0xed, 0x04, 0x28, 0x1b, 0x33, 0x4c,
	0xfe, 0xda, 0x41, 0xfd, 0x25, 0x77, 0x4e, 0xa8, 0xca, 0xce, 0xd7, 0x15, 0xfe, 0x12, 0x8d, 0x2a,
	0x80, 0x4a, 0x72, 0x42, 0x15, 0x95, 0xd7, 0x4e, 0x30, 0x9b, 0x75, 0xc6, 0x9d, 0xa3, 0xbd, 0xc5,
	0x30, 0xc6, 0x67, 0x6d, 0x18, 0x1f, 0xa3, 0x43, 0xc3, 0x1b, 0x4b, 0x0b, 0xc9, 0x89, 0xa3, 0xf6,
	0x82, 0xd0, 0x8a, 0x13, 0xcb, 0x99, 0xcd, 0xee, 0x8d, 0x3b, 0x47, 0xdb, 0x8b, 0x87, 0x2d, 0xba,
	0xa2, 0xf6, 0x62, 0x56, 0xf1, 0x25, 0x8f, 0xa2, 0x02, 0x1c, 0x29, 0x39, 0x75, 0x1b, 0xe2, 0x44,
	0xcd, 0xa1, 0x71, 0x51, 0xb4, 0x15, 0x45, 0x05, 0xb8, 0x1f, 0x3c, 0xb8, 0x8a, 0x58, 0x10, 0x7d,
	0x85, 0x0e, 0xb8, 0x8a, 0xe7, 0x58, 0x52, 0x83, 0x12, 0x0e, 0x8c, 0x50, 0x55, 0xd6, 0x1d, 0x77,
	0x8e, 0x7a, 0x0b, 0x1c, 0xb1, 0x95, 0xfd, 0xf5, 0x06, 0xc1, 0x27, 0x68, 0x57, 0x58, 0x90, 0xd4,
	0xf1, 0x6c, 0x7b, 0xdc, 0x39, 0xea, 0x3f, 0x1b, 0x4f, 0xdf, 0xb2, 0x6a, 0xfa, 0x32, 0xe2, 0xed,
	0xe5, 0x17, 0xad, 0x00, 0x7f, 0x8d, 0xba, 0x4c, 0xe8, 0x32, 0xdb, 0x09, 0xc2, 0xc7, 0xff, 0x11,
	0xce, 0x85, 0x2e, 0x6f, 0x54, 0x81, 0x8a, 0x7f, 0x46, 0x9f, 0xaf, 0xc1, 0x30, 0x4e, 0x0a, 0x70,
	0x96, 0x38, 0x20, 0x56, 0x72, 0xae, 0x09, 0x55, 0x25, 0x51, 0xe0, 0x88, 0x69, 0x54, 0x70, 0x28,
	0xeb, 0x85, 0x82, 0x1f, 0x07, 0xee, 0x29, 0x38, 0xbb, 0x82, 0xa5, 0x27, 0xce, 0x54, 0xf9, 0x0a,
	0xdc, 0xa2, 0x51, 0xde, 0x29, 0x3c, 0x46, 0x83, 0x46, 0x10, 0x26, 0x05, 0x57, 0x8e, 0x88, 0x32,
	0xdb, 0x0b, 0xf6, 0xa3, 0x46, 0xcc, 0x43, 0xe8, 0x65, 0x89, 0xbf, 0x47, 0x1f, 0x95, 0xc2, 0x6a,
	0x49, 0xaf, 0x89, 0xe5, 0xe6, 0x92, 0x1b, 0xd2, 0x18, 0x49, 0x1c, 0xaf, 0x75, 0xb8, 0x71, 0x3f,
	0x08, 0xb2, 0x44, 0x59, 0x06, 0xc6, 0xb9, 0x91, 0xab, 0x84, 0xe3, 0x27, 0x08, 0xd7, 0xf4, 0xca,
	0xd7, 0x9a, 0xea, 0xf4, 0x7d, 0xc8, 0x06, 0xc1, 0xff, 0x61, 0x4d, 0xaf, 0x4e, 0xc1, 0x85, 0xb2,
	0x7c, 0x0b, 0xbc, 0x1b, 0xb4, 0x71, 0x9b, 0xec, 0xfe, 0xff, 0xb8, 0x31, 0x6b, 0xdc, 0xe6, 0xd6,
	0x0d, 0x4f, 0x6d, 0x7b, 0x9c, 0xfc, 0x24, 0x95, 0xd1, 0x8c, 0x68, 0x03, 0x57, 0xd7, 0xd9, 0x83,
	0x50, 0x99, 0xef, 0x71, 0x72, 0xff, 0x85, 0xd1, 0xec, 0xb5, 0x87, 0xf0, 0x37, 0xe8, 0x51, 0x28,
	0x28, 0xa5, 0xbf, 0xab, 0x1a, 0x06, 0xd5, 0x41, 0x01, 0x6e, 0x99, 0xd0, 0x5b, 0xd9, 0x73, 0xf4,
	0x88, 0x5f, 0x39, 0x43, 0x09, 0xdb, 0x08, 0x59, 0x12, 0x6b, 0x18, 0x61, 0x56, 0x7b, 0x3f, 0xb2,
	0xd1, 0x78, 0xcb, 0x1f, 0x16, 0xe0, 0xb9, 0x47, 0x97, 0x86, 0xcd, 0xad, 0x3e, 0x37, 0x12, 0x7f,
	0x80, 0x7a, 0x8d, 0xe5, 0x44, 0x8a, 0x35, 0x64, 0xfb, 0xa1, 0x27, 0xbb, 0x8d, 0xe5, 0xbf, 0x88,
	0x35, 0x9c, 0x75, 0x7b, 0xbb, 0xa3, 0xde, 0x59, 0xb7, 0x87, 0x46, 0xfd, 0x09, 0x47, 0xc3, 0xb7,
	0xa6, 0x04, 0x7f, 0x81, 0x1e, 0x94, 0x7c, 0x4d, 0x1b, 0xe9, 0x92, 0xf5, 0x69, 0x3b, 0xee, 0xa7,
	0x68, 0x74, 0x1b, 0x3f, 0x41, 0xfb, 0x2d, 0x4d, 0xd1, 0x9a, 0x5b, 0x4d, 0x19, 0x0f, 0x6b, 0xb1,
	0xb7, 0x18, 0x25, 0xe0, 0x55, 0x1b, 0x9f, 0x9c, 0xa1, 0xbe, 0x9f, 0xa9, 0xd7, 0x94, 0x5d, 0xd0,
	0x8a, 0xe3, 0x4f, 0xd1, 0x40, 0xc7, 0x9f, 0x41, 0x9b, 0x0e, 0xe8, 0xa7, 0x98, 0x97, 0xe1, 0x0c,
	0xed, 0x5e, 0x72, 0x63, 0x05, 0xa8, 0x94, 0xb4, 0xfd, 0x3b, 0xf9, 0xa3, 0x83, 0x06, 0x77, 0x07,
	0xf4, 0x7d, 0x0b, 0x5e, 0xa0, 0xc3, 0x96, 0x96, 0x26, 0x2f, 0x9d, 0x17, 0x0e, 0xe8, 0x3f, 0xfb,
	0xf8, 0x9d, 0x6b, 0x90, 0x4a, 0x5e, 0x1c, 0x24, 0x6d, 0x9c, 0xd0, 0x14, 0x9d, 0xfc, 0x79, 0x0f,
	0x0d, 0xee, 0x8e, 0x87, 0xbf, 0x19, 0x2d, 0x6b, 0xa1, 0x2c, 0xa9, 0x0c, 0x34, 0xba, 0xbd, 0x59,
	0x8c, 0xbd, 0xf0, 0x21, 0x3c, 0x45, 0x7e, 0x3a, 0x48, 0x01, 0xe0, 0xac, 0x33, 0x54, 0x27, 0x66,
	0xbc, 0xe5, 0x7e, 0x01, 0xee, 0xb4, 0x45, 0x22, 0xff, 0x39, 0x3a, 0xd4, 0x46, 0x5c, 0x0a, 0xc9,
	0x2b, 0x5e, 0x92, 0xc6, 0x72, 0xd3, 0x26, 0xdf, 0x8a, 0x53, 0x73, 0x8b, 0x9e, 0x7b, 0x30, 0xaa,
	0x3e, 0x41, 0xfd, 0xbb, 0xd4, 0x6e, 0xda, 0xb0, 0x5b, 0xc2, 0x53, 0xf4, 0xf0, 0x52, 0xf0, 0x37,
	0x84, 0x4a, 0x19, 0x77, 0x3a, 0x12, 0xb7, 0x63, 0x07, 0x3d, 0x34, 0x93, 0xd2, 0x6f, 0x70, 0xa4,
	0xe7, 0xe8, 0xe0, 0x86, 0xee, 0x17, 0xbd, 0xe5, 0xef, 0xc4, 0xb2, 0x13, 0xdf, 0x6f, 0x77, 0x14,
	0x9c, 0xfe, 0xde, 0xf9, 0xfb, 0x9f, 0xcf, 0x7e, 0x42, 0x3f, 0x6e, 0x9c, 0xd3, 0xf6, 0x24, 0x0f,
	0xaf, 0xe9, 0xa7, 0xc9, 0x58, 0xaa, 0xb5, 0xd5, 0xe0, 0xa6, 0x0c, 0xea, 0xdc, 0xb2, 0x0d, 0xaf,
	0xa9, 0xcd, 0x7d, 0xeb, 0x04, 0xe3, 0xf6, 0xe6, 0xa3, 0x70, 0x62, 0x93, 0xa9, 0x53, 0xb6, 0xae,
	0x7e, 0x3b, 0x7e, 0xff, 0x0f, 0xc8, 0x77, 0xf1, 0xa1, 0x8b, 0x62, 0x27, 0x84, 0x8f, 0xff, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xc7, 0x9f, 0x92, 0x58, 0x7e, 0x06, 0x00, 0x00,
}
