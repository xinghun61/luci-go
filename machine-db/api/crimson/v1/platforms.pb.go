// Code generated by protoc-gen-go. DO NOT EDIT.
// source: go.chromium.org/luci/machine-db/api/crimson/v1/platforms.proto

package crimson

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "go.chromium.org/luci/machine-db/api/common/v1"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// A platform in the database.
type Platform struct {
	// The name of this platform. Uniquely identifies this platform.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A description of this platform.
	Description string `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	// The state of this platform.
	State common.State `protobuf:"varint,3,opt,name=state,enum=common.State" json:"state,omitempty"`
}

func (m *Platform) Reset()                    { *m = Platform{} }
func (m *Platform) String() string            { return proto.CompactTextString(m) }
func (*Platform) ProtoMessage()               {}
func (*Platform) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *Platform) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Platform) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Platform) GetState() common.State {
	if m != nil {
		return m.State
	}
	return common.State_STATE_UNSPECIFIED
}

// A request to list platforms in the database.
type ListPlatformsRequest struct {
	// The names of platforms to retrieve.
	Names []string `protobuf:"bytes,1,rep,name=names" json:"names,omitempty"`
}

func (m *ListPlatformsRequest) Reset()                    { *m = ListPlatformsRequest{} }
func (m *ListPlatformsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListPlatformsRequest) ProtoMessage()               {}
func (*ListPlatformsRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *ListPlatformsRequest) GetNames() []string {
	if m != nil {
		return m.Names
	}
	return nil
}

// A response containing a list of platforms in the database.
type ListPlatformsResponse struct {
	// The platforms matching the request.
	Platforms []*Platform `protobuf:"bytes,1,rep,name=platforms" json:"platforms,omitempty"`
}

func (m *ListPlatformsResponse) Reset()                    { *m = ListPlatformsResponse{} }
func (m *ListPlatformsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListPlatformsResponse) ProtoMessage()               {}
func (*ListPlatformsResponse) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *ListPlatformsResponse) GetPlatforms() []*Platform {
	if m != nil {
		return m.Platforms
	}
	return nil
}

func init() {
	proto.RegisterType((*Platform)(nil), "crimson.Platform")
	proto.RegisterType((*ListPlatformsRequest)(nil), "crimson.ListPlatformsRequest")
	proto.RegisterType((*ListPlatformsResponse)(nil), "crimson.ListPlatformsResponse")
}

func init() {
	proto.RegisterFile("go.chromium.org/luci/machine-db/api/crimson/v1/platforms.proto", fileDescriptor7)
}

var fileDescriptor7 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x41, 0x4f, 0x03, 0x21,
	0x14, 0x84, 0xb3, 0xd6, 0xaa, 0xfb, 0x1a, 0x4d, 0x24, 0x35, 0xd9, 0x78, 0xda, 0xac, 0x97, 0x3d,
	0x28, 0xc4, 0x7a, 0xf3, 0xe0, 0xd9, 0x83, 0x07, 0x83, 0xbf, 0x80, 0x52, 0x6c, 0x49, 0x0a, 0x0f,
	0x79, 0xac, 0xbf, 0xdf, 0x94, 0x65, 0xd5, 0x78, 0xf2, 0x06, 0x33, 0xc3, 0x37, 0x61, 0xe0, 0x69,
	0x8b, 0x5c, 0xef, 0x22, 0x3a, 0x3b, 0x38, 0x8e, 0x71, 0x2b, 0xf6, 0x83, 0xb6, 0xc2, 0x29, 0xbd,
	0xb3, 0xde, 0xdc, 0x6d, 0xd6, 0x42, 0x05, 0x2b, 0x74, 0xb4, 0x8e, 0xd0, 0x8b, 0xcf, 0x7b, 0x11,
	0xf6, 0x2a, 0xbd, 0x63, 0x74, 0xc4, 0x43, 0xc4, 0x84, 0xec, 0xb4, 0x78, 0xd7, 0x8f, 0xff, 0x02,
	0xa1, 0x73, 0x23, 0x87, 0x92, 0x4a, 0xa6, 0x40, 0x3a, 0x03, 0x67, 0xaf, 0x85, 0xcb, 0x18, 0x1c,
	0x7b, 0xe5, 0x4c, 0x53, 0xb5, 0x55, 0x5f, 0xcb, 0x7c, 0x66, 0x2d, 0x2c, 0x36, 0x86, 0x74, 0xb4,
	0x21, 0x59, 0xf4, 0xcd, 0x51, 0xb6, 0x7e, 0x4b, 0xec, 0x06, 0xe6, 0x99, 0xd8, 0xcc, 0xda, 0xaa,
	0xbf, 0x58, 0x9d, 0xf3, 0xb1, 0x89, 0xbf, 0x1d, 0x44, 0x39, 0x7a, 0xdd, 0x2d, 0x2c, 0x5f, 0x2c,
	0xa5, 0xa9, 0x8a, 0xa4, 0xf9, 0x18, 0x0c, 0x25, 0xb6, 0x84, 0xf9, 0xa1, 0x86, 0x9a, 0xaa, 0x9d,
	0xf5, 0xb5, 0x1c, 0x2f, 0xdd, 0x33, 0x5c, 0xfd, 0x49, 0x53, 0x40, 0x4f, 0x86, 0x09, 0xa8, 0xbf,
	0x57, 0xc8, 0x4f, 0x16, 0xab, 0x4b, 0x5e, 0x66, 0xe0, 0x53, 0x5c, 0xfe, 0x64, 0xd6, 0x27, 0xf9,
	0x97, 0x0f, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xb0, 0xfd, 0x75, 0x6c, 0x01, 0x00, 0x00,
}
