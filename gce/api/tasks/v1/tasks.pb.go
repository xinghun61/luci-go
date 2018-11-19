// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/gce/api/tasks/v1/tasks.proto

package tasks

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A task to expand a VMs config.
type Expansion struct {
	// The ID of the VMs block to expand.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Expansion) Reset()         { *m = Expansion{} }
func (m *Expansion) String() string { return proto.CompactTextString(m) }
func (*Expansion) ProtoMessage()    {}
func (*Expansion) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63d8744087b0bbc, []int{0}
}

func (m *Expansion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expansion.Unmarshal(m, b)
}
func (m *Expansion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expansion.Marshal(b, m, deterministic)
}
func (m *Expansion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expansion.Merge(m, src)
}
func (m *Expansion) XXX_Size() int {
	return xxx_messageInfo_Expansion.Size(m)
}
func (m *Expansion) XXX_DiscardUnknown() {
	xxx_messageInfo_Expansion.DiscardUnknown(m)
}

var xxx_messageInfo_Expansion proto.InternalMessageInfo

func (m *Expansion) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Expansion)(nil), "tasks.Expansion")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/gce/api/tasks/v1/tasks.proto", fileDescriptor_f63d8744087b0bbc)
}

var fileDescriptor_f63d8744087b0bbc = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4c, 0xcf, 0xd7, 0x4b,
	0xce, 0x28, 0xca, 0xcf, 0xcd, 0x2c, 0xcd, 0xd5, 0xcb, 0x2f, 0x4a, 0xd7, 0xcf, 0x29, 0x4d, 0xce,
	0xd4, 0x4f, 0x4f, 0x4e, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x49, 0x2c, 0xce, 0x2e, 0xd6, 0x2f,
	0x33, 0x84, 0x30, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x58, 0xc1, 0x1c, 0x25, 0x69, 0x2e,
	0x4e, 0xd7, 0x8a, 0x82, 0xc4, 0xbc, 0xe2, 0xcc, 0xfc, 0x3c, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xa6, 0xcc, 0x94, 0x24, 0x36, 0xb0, 0x52, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0x03, 0xe4, 0x86, 0x5f, 0x00, 0x00, 0x00,
}