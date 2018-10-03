// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/appengine/mapper/internal/tasks/tasks.proto

package tasks

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// SplitAndLaunch task splits the key range into shards and kicks off processing
// of each individual shard.
//
// Enqueued transactionally when creating a new mapping job.
type SplitAndLaunch struct {
	JobId                int64    `protobuf:"varint,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SplitAndLaunch) Reset()         { *m = SplitAndLaunch{} }
func (m *SplitAndLaunch) String() string { return proto.CompactTextString(m) }
func (*SplitAndLaunch) ProtoMessage()    {}
func (*SplitAndLaunch) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c79c318740f2dbc, []int{0}
}

func (m *SplitAndLaunch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SplitAndLaunch.Unmarshal(m, b)
}
func (m *SplitAndLaunch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SplitAndLaunch.Marshal(b, m, deterministic)
}
func (m *SplitAndLaunch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SplitAndLaunch.Merge(m, src)
}
func (m *SplitAndLaunch) XXX_Size() int {
	return xxx_messageInfo_SplitAndLaunch.Size(m)
}
func (m *SplitAndLaunch) XXX_DiscardUnknown() {
	xxx_messageInfo_SplitAndLaunch.DiscardUnknown(m)
}

var xxx_messageInfo_SplitAndLaunch proto.InternalMessageInfo

func (m *SplitAndLaunch) GetJobId() int64 {
	if m != nil {
		return m.JobId
	}
	return 0
}

func init() {
	proto.RegisterType((*SplitAndLaunch)(nil), "appengine.mapper.internal.tasks.SplitAndLaunch")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/appengine/mapper/internal/tasks/tasks.proto", fileDescriptor_7c79c318740f2dbc)
}

var fileDescriptor_7c79c318740f2dbc = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x48, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x4f, 0x2c, 0x28, 0x48, 0xcd, 0x4b, 0xcf, 0xcc, 0x4b, 0xd5, 0xcf, 0x05, 0x31, 0x8b, 0xf4,
	0x33, 0xf3, 0x4a, 0x52, 0x8b, 0xf2, 0x12, 0x73, 0xf4, 0x4b, 0x12, 0x8b, 0xb3, 0x8b, 0x21, 0xa4,
	0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x3c, 0x5c, 0xb1, 0x1e, 0x44, 0xb1, 0x1e, 0x4c, 0xb1,
	0x1e, 0x58, 0x99, 0x92, 0x3a, 0x17, 0x5f, 0x70, 0x41, 0x4e, 0x66, 0x89, 0x63, 0x5e, 0x8a, 0x4f,
	0x62, 0x69, 0x5e, 0x72, 0x86, 0x90, 0x28, 0x17, 0x5b, 0x56, 0x7e, 0x52, 0x7c, 0x66, 0x8a, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x6b, 0x56, 0x7e, 0x92, 0x67, 0x8a, 0x13, 0x7b, 0x14, 0x2b,
	0x58, 0x47, 0x12, 0x1b, 0xd8, 0x64, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0xe3, 0x89,
	0x27, 0x9d, 0x00, 0x00, 0x00,
}