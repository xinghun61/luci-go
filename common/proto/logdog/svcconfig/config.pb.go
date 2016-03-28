// Code generated by protoc-gen-go.
// source: config.proto
// DO NOT EDIT!

/*
Package svcconfig is a generated protocol buffer package.

It is generated from these files:
	config.proto
	storage.proto
	transport.proto

It has these top-level messages:
	Config
	Coordinator
	Collector
	Archivist
	Janitor
	Storage
	Transport
*/
package svcconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/luci/luci-go/common/proto/google"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// Config is the overall instance configuration.
type Config struct {
	// Configuration for the Pub/Sub instances.
	Transport *Transport `protobuf:"bytes,10,opt,name=transport" json:"transport,omitempty"`
	// Configuration for Storage.
	Storage *Storage `protobuf:"bytes,11,opt,name=storage" json:"storage,omitempty"`
	// Coordinator is the coordinator service configuration.
	Coordinator *Coordinator `protobuf:"bytes,20,opt,name=coordinator" json:"coordinator,omitempty"`
	// Collector is the collector fleet configuration.
	Collector *Collector `protobuf:"bytes,21,opt,name=collector" json:"collector,omitempty"`
	// Archivist microservice configuration.
	Archivist *Archivist `protobuf:"bytes,22,opt,name=archivist" json:"archivist,omitempty"`
	// Janitor microservice configuration.
	Janitor *Janitor `protobuf:"bytes,23,opt,name=janitor" json:"janitor,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Config) GetTransport() *Transport {
	if m != nil {
		return m.Transport
	}
	return nil
}

func (m *Config) GetStorage() *Storage {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *Config) GetCoordinator() *Coordinator {
	if m != nil {
		return m.Coordinator
	}
	return nil
}

func (m *Config) GetCollector() *Collector {
	if m != nil {
		return m.Collector
	}
	return nil
}

func (m *Config) GetArchivist() *Archivist {
	if m != nil {
		return m.Archivist
	}
	return nil
}

func (m *Config) GetJanitor() *Janitor {
	if m != nil {
		return m.Janitor
	}
	return nil
}

// Coordinator is the Coordinator service configuration.
type Coordinator struct {
	// Project is the name of the AppEngine Project that the Coordinator belongs
	// to.
	Project string `protobuf:"bytes,1,opt,name=project" json:"project,omitempty"`
	// The name of the authentication group for administrators.
	AdminAuthGroup string `protobuf:"bytes,10,opt,name=admin_auth_group" json:"admin_auth_group,omitempty"`
	// The name of the authentication group for backend services.
	ServiceAuthGroup string `protobuf:"bytes,11,opt,name=service_auth_group" json:"service_auth_group,omitempty"`
	// A list of origin URLs that are allowed to perform CORS RPC calls.
	RpcAllowOrigins []string `protobuf:"bytes,20,rep,name=rpc_allow_origins" json:"rpc_allow_origins,omitempty"`
	// The name of the archive task queue.
	ArchiveTaskQueue string `protobuf:"bytes,30,opt,name=archive_task_queue" json:"archive_task_queue,omitempty"`
	// The amount of time after a log has been terminated before it is candidate
	// for archival.
	//
	// Archival triggered by this delay will NOT succeed if any log entries are
	// missing from intermediate storage.
	//
	// This should be based on a period of time where it's reasonable to expect
	// that all log messages in the transport have arrived for a given log stream.
	// Since the transport doesn't have to guarantee in-order delivery, this
	// should allow for the case where the terminal log entry arrives before some
	// of the intermediate log entries. This will help avoid triggering
	// archive attempts that are doomed to fail because of standard transport lag.
	ArchiveDelay *google_protobuf.Duration `protobuf:"bytes,31,opt,name=archive_delay" json:"archive_delay,omitempty"`
	// The amount of time before a log stream is candidate for archival regardless
	// of whether or not it's been terminated or complete.
	//
	// This endpoint is a failsafe designed to ensure that log streams with
	// missing records or no terminal record (e.g., Butler crashed) are eventually
	// moved out of intermediate storage.
	//
	// This must be >= `archive_delay`, and should be fairly large (days) to allow
	// for the log stream to complete and for all available log entries to be
	// added to intermediate storage.
	ArchiveDelayMax *google_protobuf.Duration `protobuf:"bytes,32,opt,name=archive_delay_max" json:"archive_delay_max,omitempty"`
}

func (m *Coordinator) Reset()                    { *m = Coordinator{} }
func (m *Coordinator) String() string            { return proto.CompactTextString(m) }
func (*Coordinator) ProtoMessage()               {}
func (*Coordinator) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Coordinator) GetArchiveDelay() *google_protobuf.Duration {
	if m != nil {
		return m.ArchiveDelay
	}
	return nil
}

func (m *Coordinator) GetArchiveDelayMax() *google_protobuf.Duration {
	if m != nil {
		return m.ArchiveDelayMax
	}
	return nil
}

// Collector is the set of configuration parameters for Collector instances.
type Collector struct {
	// Workers is the number of ingest workers to run.
	Workers int32 `protobuf:"varint,1,opt,name=workers" json:"workers,omitempty"`
	// The number of transport worker goroutines to run.
	TransportWorkers int32 `protobuf:"varint,2,opt,name=transport_workers" json:"transport_workers,omitempty"`
	// The maximum number of log stream states to cache locally. If <= 0, a
	// default will be used.
	StateCacheSize int32 `protobuf:"varint,3,opt,name=state_cache_size" json:"state_cache_size,omitempty"`
	// The maximum amount of time that cached stream state is valid. If <= 0, a
	// default will be used.
	StateCacheExpiration *google_protobuf.Duration `protobuf:"bytes,4,opt,name=state_cache_expiration" json:"state_cache_expiration,omitempty"`
}

func (m *Collector) Reset()                    { *m = Collector{} }
func (m *Collector) String() string            { return proto.CompactTextString(m) }
func (*Collector) ProtoMessage()               {}
func (*Collector) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Collector) GetStateCacheExpiration() *google_protobuf.Duration {
	if m != nil {
		return m.StateCacheExpiration
	}
	return nil
}

// Configuration for the Archivist microservice.
type Archivist struct {
	// The number of tasks to run at a time. If blank, the archivist will choose a
	// default value.
	Tasks int32 `protobuf:"varint,1,opt,name=tasks" json:"tasks,omitempty"`
	// The name of the Google Storage bucket and optional base path to archive
	// into. For example: gs://foo/bar
	//
	// The bucket name must be included (e.g., "gs://foo"). The remainder of the
	// base path is optional based on desired archive location.
	GsBase string `protobuf:"bytes,10,opt,name=gs_base" json:"gs_base,omitempty"`
	// If not zero, the maximum number of stream indices between index entries.
	StreamIndexRange int32 `protobuf:"varint,11,opt,name=stream_index_range" json:"stream_index_range,omitempty"`
	// If not zero, the maximum number of prefix indices between index entries.
	PrefixIndexRange int32 `protobuf:"varint,12,opt,name=prefix_index_range" json:"prefix_index_range,omitempty"`
	// If not zero, the maximum number of log data bytes between index entries.
	ByteRange int32 `protobuf:"varint,13,opt,name=byte_range" json:"byte_range,omitempty"`
}

func (m *Archivist) Reset()                    { *m = Archivist{} }
func (m *Archivist) String() string            { return proto.CompactTextString(m) }
func (*Archivist) ProtoMessage()               {}
func (*Archivist) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Configuration for the Janitor microservice.
type Janitor struct {
	// The number of tasks to run at a time. If blank, the janitor will choose a
	// default value.
	Tasks int32 `protobuf:"varint,1,opt,name=tasks" json:"tasks,omitempty"`
}

func (m *Janitor) Reset()                    { *m = Janitor{} }
func (m *Janitor) String() string            { return proto.CompactTextString(m) }
func (*Janitor) ProtoMessage()               {}
func (*Janitor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*Config)(nil), "svcconfig.Config")
	proto.RegisterType((*Coordinator)(nil), "svcconfig.Coordinator")
	proto.RegisterType((*Collector)(nil), "svcconfig.Collector")
	proto.RegisterType((*Archivist)(nil), "svcconfig.Archivist")
	proto.RegisterType((*Janitor)(nil), "svcconfig.Janitor")
}

var fileDescriptor0 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x92, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x15, 0x46, 0x57, 0xe5, 0x65, 0xd5, 0x98, 0x35, 0x8a, 0xd7, 0xc3, 0x98, 0xca, 0x01,
	0x24, 0xa4, 0x0c, 0x01, 0x17, 0x8e, 0xa8, 0x9c, 0xb8, 0xc2, 0xdd, 0x72, 0x9d, 0xd7, 0xd4, 0x5b,
	0x1a, 0x07, 0xdb, 0x69, 0x3b, 0xfe, 0x06, 0xfe, 0x5b, 0x6e, 0x9c, 0x70, 0xec, 0xa4, 0xcd, 0xd6,
	0x43, 0x8f, 0x7e, 0xbf, 0xf7, 0xfc, 0x7d, 0xfe, 0x9e, 0xe1, 0x4c, 0xa8, 0x72, 0x21, 0xf3, 0xb4,
	0xd2, 0xca, 0x2a, 0x12, 0x9b, 0xb5, 0x08, 0x85, 0xc9, 0xb9, 0xd5, 0xbc, 0x34, 0x95, 0xd2, 0x36,
	0xb0, 0xc9, 0xc8, 0x58, 0xa5, 0x79, 0x8e, 0xed, 0xf1, 0x3a, 0x57, 0x2a, 0x2f, 0xf0, 0xd6, 0x9f,
	0xe6, 0xf5, 0xe2, 0x36, 0xab, 0x35, 0xb7, 0x52, 0x95, 0x81, 0x4f, 0xff, 0x45, 0x70, 0x3a, 0xf3,
	0x57, 0x91, 0xb7, 0x10, 0xef, 0x2e, 0xa3, 0x70, 0x13, 0xbd, 0x4b, 0x3e, 0x5e, 0xa6, 0x3b, 0xa5,
	0xf4, 0x67, 0xc7, 0xc8, 0x1b, 0x18, 0xb6, 0x22, 0x34, 0xf1, 0x6d, 0xa4, 0xd7, 0xf6, 0x23, 0x10,
	0xf2, 0x1e, 0x12, 0xa1, 0x94, 0xce, 0x64, 0xc9, 0x5d, 0x85, 0x5e, 0xfa, 0xc6, 0x71, 0xaf, 0x71,
	0xb6, 0xa7, 0x8d, 0xb4, 0x50, 0x45, 0x81, 0xa2, 0x69, 0x7d, 0x79, 0x20, 0x3d, 0xeb, 0x58, 0xd3,
	0xc8, 0xb5, 0x58, 0xca, 0xb5, 0x34, 0x96, 0x8e, 0x0f, 0x1a, 0xbf, 0x76, 0xac, 0xf1, 0x78, 0xc7,
	0x4b, 0xd9, 0xdc, 0xf7, 0xea, 0xc0, 0xe3, 0xf7, 0x40, 0xa6, 0x7f, 0x23, 0x48, 0xfa, 0x36, 0xce,
	0x61, 0xe8, 0x52, 0xb9, 0x73, 0x52, 0x34, 0x72, 0x43, 0x31, 0xa1, 0xf0, 0x82, 0x67, 0x2b, 0x59,
	0x32, 0x5e, 0xdb, 0x25, 0xcb, 0xb5, 0xaa, 0x2b, 0x9f, 0x4c, 0x4c, 0x26, 0x40, 0x0c, 0xea, 0xb5,
	0x14, 0xd8, 0x67, 0x89, 0x67, 0x57, 0x70, 0xa1, 0x2b, 0xc1, 0x78, 0x51, 0xa8, 0x0d, 0x53, 0x5a,
	0xe6, 0xb2, 0x34, 0x2e, 0x80, 0x93, 0x30, 0x16, 0xfc, 0x23, 0xb3, 0xdc, 0xdc, 0xb3, 0x5f, 0x35,
	0xd6, 0x48, 0xaf, 0xfd, 0xd8, 0x07, 0x18, 0x75, 0x2c, 0xc3, 0x82, 0x3f, 0xd0, 0xd7, 0xde, 0xf8,
	0x55, 0x1a, 0x56, 0x98, 0x76, 0x2b, 0x4c, 0xbf, 0xb5, 0x2b, 0x24, 0x9f, 0xe1, 0xe2, 0xd1, 0x04,
	0x5b, 0xf1, 0x2d, 0xbd, 0x39, 0x32, 0x35, 0xfd, 0x13, 0x41, 0xbc, 0x4f, 0xd4, 0xbd, 0x79, 0xa3,
	0xf4, 0x3d, 0x6a, 0xe3, 0xdf, 0x3c, 0x68, 0xdc, 0xef, 0xbe, 0x01, 0xeb, 0xd0, 0x33, 0x8f, 0x5c,
	0x1c, 0xc6, 0x72, 0x8b, 0x4c, 0x70, 0xb1, 0x44, 0x66, 0xe4, 0x6f, 0xa4, 0x27, 0x9e, 0x7c, 0x81,
	0x71, 0x9f, 0xe0, 0xb6, 0x92, 0x41, 0x8d, 0x3e, 0x3f, 0x66, 0x67, 0x03, 0xf1, 0x7e, 0x6d, 0x23,
	0x18, 0x34, 0xb9, 0x74, 0x5e, 0x9c, 0xb9, 0xdc, 0xb0, 0x39, 0x37, 0xd8, 0x8b, 0xdd, 0x6a, 0xe4,
	0x2b, 0x26, 0xcb, 0x0c, 0xb7, 0xcc, 0x19, 0x6d, 0x7f, 0xe1, 0xa0, 0x61, 0x95, 0xc6, 0x85, 0xdc,
	0x3e, 0x62, 0x67, 0x9e, 0x11, 0x80, 0xf9, 0x83, 0xb3, 0x17, 0x6a, 0xa3, 0xa6, 0x36, 0xa5, 0x30,
	0x6c, 0x3f, 0xc2, 0x13, 0xd9, 0xf9, 0xa9, 0x77, 0xf9, 0xe9, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xa8, 0xd7, 0xb7, 0x26, 0x76, 0x03, 0x00, 0x00,
}
