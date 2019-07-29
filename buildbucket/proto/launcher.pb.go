// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/buildbucket/proto/launcher.proto

package buildbucketpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// A collection of build-related secrets we might pass from Buildbucket to Kitchen.
type BuildSecrets struct {
	// Token to identify RPCs associated with the same build.
	BuildToken           string   `protobuf:"bytes,1,opt,name=build_token,json=buildToken,proto3" json:"build_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildSecrets) Reset()         { *m = BuildSecrets{} }
func (m *BuildSecrets) String() string { return proto.CompactTextString(m) }
func (*BuildSecrets) ProtoMessage()    {}
func (*BuildSecrets) Descriptor() ([]byte, []int) {
	return fileDescriptor_45f9e6fb262a81d2, []int{0}
}

func (m *BuildSecrets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildSecrets.Unmarshal(m, b)
}
func (m *BuildSecrets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildSecrets.Marshal(b, m, deterministic)
}
func (m *BuildSecrets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildSecrets.Merge(m, src)
}
func (m *BuildSecrets) XXX_Size() int {
	return xxx_messageInfo_BuildSecrets.Size(m)
}
func (m *BuildSecrets) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildSecrets.DiscardUnknown(m)
}

var xxx_messageInfo_BuildSecrets proto.InternalMessageInfo

func (m *BuildSecrets) GetBuildToken() string {
	if m != nil {
		return m.BuildToken
	}
	return ""
}

// Arguments for luci_runner command.
// All paths are relateive to the runner process' working directory.
type RunnerArgs struct {
	// Buildbucket service hostname, e.g. "cr-buildbucket.appspot.com".
	BuildbucketHost string `protobuf:"bytes,1,opt,name=buildbucket_host,json=buildbucketHost,proto3" json:"buildbucket_host,omitempty"`
	// LogDog service hostname, e.g. "logs.chromium.org".
	//
	// As a special case, if this is "file://path/to/log/root" the logs will
	// be written to the given local path.
	LogdogHost string `protobuf:"bytes,2,opt,name=logdog_host,json=logdogHost,proto3" json:"logdog_host,omitempty"`
	// Path to the user executable.
	// Required.
	ExecutablePath string `protobuf:"bytes,3,opt,name=executable_path,json=executablePath,proto3" json:"executable_path,omitempty"`
	// Path to a directory where each subdirectory is a cache dir.
	// Managed by Swarming.
	// Required.
	CacheDir string `protobuf:"bytes,4,opt,name=cache_dir,json=cacheDir,proto3" json:"cache_dir,omitempty"`
	// List of Gerrit hosts to force git authentication for.
	//
	// By default public hosts are accessed anonymously, and the anonymous access
	// has very low quota. Context needs to know all such hostnames in advance to
	// be able to force authenticated access to them.
	KnownPublicGerritHosts []string `protobuf:"bytes,5,rep,name=known_public_gerrit_hosts,json=knownPublicGerritHosts,proto3" json:"known_public_gerrit_hosts,omitempty"`
	// Use this LUCI context logical account for system-level operations.
	LuciSystemAccount string `protobuf:"bytes,6,opt,name=luci_system_account,json=luciSystemAccount,proto3" json:"luci_system_account,omitempty"`
	// Initial state of the build, including immutable state such as id and input
	// properties.
	Build                *Build   `protobuf:"bytes,7,opt,name=build,proto3" json:"build,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunnerArgs) Reset()         { *m = RunnerArgs{} }
func (m *RunnerArgs) String() string { return proto.CompactTextString(m) }
func (*RunnerArgs) ProtoMessage()    {}
func (*RunnerArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_45f9e6fb262a81d2, []int{1}
}

func (m *RunnerArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunnerArgs.Unmarshal(m, b)
}
func (m *RunnerArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunnerArgs.Marshal(b, m, deterministic)
}
func (m *RunnerArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunnerArgs.Merge(m, src)
}
func (m *RunnerArgs) XXX_Size() int {
	return xxx_messageInfo_RunnerArgs.Size(m)
}
func (m *RunnerArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_RunnerArgs.DiscardUnknown(m)
}

var xxx_messageInfo_RunnerArgs proto.InternalMessageInfo

func (m *RunnerArgs) GetBuildbucketHost() string {
	if m != nil {
		return m.BuildbucketHost
	}
	return ""
}

func (m *RunnerArgs) GetLogdogHost() string {
	if m != nil {
		return m.LogdogHost
	}
	return ""
}

func (m *RunnerArgs) GetExecutablePath() string {
	if m != nil {
		return m.ExecutablePath
	}
	return ""
}

func (m *RunnerArgs) GetCacheDir() string {
	if m != nil {
		return m.CacheDir
	}
	return ""
}

func (m *RunnerArgs) GetKnownPublicGerritHosts() []string {
	if m != nil {
		return m.KnownPublicGerritHosts
	}
	return nil
}

func (m *RunnerArgs) GetLuciSystemAccount() string {
	if m != nil {
		return m.LuciSystemAccount
	}
	return ""
}

func (m *RunnerArgs) GetBuild() *Build {
	if m != nil {
		return m.Build
	}
	return nil
}

func init() {
	proto.RegisterType((*BuildSecrets)(nil), "buildbucket.v2.BuildSecrets")
	proto.RegisterType((*RunnerArgs)(nil), "buildbucket.v2.RunnerArgs")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/buildbucket/proto/launcher.proto", fileDescriptor_45f9e6fb262a81d2)
}

var fileDescriptor_45f9e6fb262a81d2 = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4d, 0x4b, 0xc3, 0x30,
	0x18, 0xc7, 0xd9, 0xe6, 0xa6, 0xcb, 0x64, 0xd3, 0x88, 0x52, 0xf5, 0xe0, 0xd8, 0xc5, 0x89, 0x90,
	0xc2, 0x7c, 0x01, 0xf1, 0xb4, 0x21, 0xe8, 0x71, 0x74, 0x9e, 0xbc, 0x84, 0x36, 0x0b, 0x6d, 0x58,
	0x9b, 0x94, 0xbc, 0xf8, 0xf2, 0xc5, 0xfc, 0x7c, 0x92, 0x27, 0x82, 0xf5, 0xb6, 0x63, 0x7e, 0xff,
	0x5f, 0xff, 0x4f, 0x93, 0x07, 0xdd, 0xe5, 0x8a, 0xb0, 0x42, 0xab, 0x4a, 0xb8, 0x8a, 0x28, 0x9d,
	0xc7, 0xa5, 0x63, 0x22, 0xce, 0x9c, 0x28, 0xd7, 0x99, 0x63, 0x1b, 0x6e, 0xe3, 0x5a, 0x2b, 0xab,
	0xe2, 0x32, 0x75, 0x92, 0x15, 0x5c, 0x13, 0x38, 0xe2, 0x61, 0xc3, 0x20, 0xef, 0xb3, 0xb3, 0xd9,
	0x96, 0x35, 0x40, 0x42, 0xc7, 0x24, 0x46, 0xfb, 0x0b, 0x7f, 0x5c, 0x71, 0xa6, 0xb9, 0x35, 0xf8,
	0x02, 0x0d, 0x20, 0xa6, 0x56, 0x6d, 0xb8, 0x8c, 0x5a, 0xe3, 0xd6, 0xb4, 0x9f, 0x20, 0x40, 0xaf,
	0x9e, 0x4c, 0xbe, 0xdb, 0x08, 0x25, 0x4e, 0x4a, 0xae, 0xe7, 0x3a, 0x37, 0xf8, 0x0a, 0x1d, 0x34,
	0x06, 0xd0, 0x42, 0x19, 0xfb, 0xfb, 0xd1, 0xa8, 0xc1, 0x5f, 0x94, 0xb1, 0xbe, 0xba, 0x54, 0xf9,
	0x5a, 0xe5, 0xc1, 0x6a, 0x87, 0xea, 0x80, 0x40, 0xb8, 0x44, 0x23, 0xfe, 0xc9, 0x99, 0xb3, 0x69,
	0x56, 0x72, 0x5a, 0xa7, 0xb6, 0x88, 0x3a, 0x20, 0x0d, 0xff, 0xf0, 0x32, 0xb5, 0x05, 0x3e, 0x47,
	0x7d, 0x96, 0xb2, 0x82, 0xd3, 0xb5, 0xd0, 0xd1, 0x0e, 0x28, 0x7b, 0x00, 0x9e, 0x84, 0xc6, 0x0f,
	0xe8, 0x74, 0x23, 0xd5, 0x87, 0xa4, 0xb5, 0xcb, 0x4a, 0xc1, 0x68, 0xce, 0xb5, 0x16, 0xe1, 0xcf,
	0x4c, 0xd4, 0x1d, 0x77, 0xa6, 0xfd, 0xe4, 0x04, 0x84, 0x25, 0xe4, 0xcf, 0x10, 0xfb, 0xf9, 0x06,
	0x13, 0x74, 0xe4, 0x9f, 0x8c, 0x9a, 0x2f, 0x63, 0x79, 0x45, 0x53, 0xc6, 0x94, 0x93, 0x36, 0xea,
	0xc1, 0x84, 0x43, 0x1f, 0xad, 0x20, 0x99, 0x87, 0x00, 0x5f, 0xa3, 0x2e, 0x5c, 0x32, 0xda, 0x1d,
	0xb7, 0xa6, 0x83, 0xd9, 0x31, 0xf9, 0xbf, 0x10, 0x02, 0x2f, 0x9b, 0x04, 0x67, 0x71, 0xff, 0x76,
	0xbb, 0xdd, 0x7e, 0x1e, 0x1b, 0xa4, 0xce, 0xb2, 0x1e, 0xc0, 0x9b, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xdd, 0xd6, 0x47, 0x6b, 0x25, 0x02, 0x00, 0x00,
}
