// Code generated by protoc-gen-go.
// source: log_entry.proto
// DO NOT EDIT!

package dynobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LogEntry struct {
	ApplicationId string `protobuf:"bytes,1,opt,name=application_id,json=applicationId" json:"application_id,omitempty"`
	Database      string `protobuf:"bytes,2,opt,name=database" json:"database,omitempty"`
	Process       string `protobuf:"bytes,3,opt,name=process" json:"process,omitempty"`
	Timestamp     int64  `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Data          string `protobuf:"bytes,5,opt,name=data" json:"data,omitempty"`
}

func (m *LogEntry) Reset()                    { *m = LogEntry{} }
func (m *LogEntry) String() string            { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()               {}
func (*LogEntry) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func init() {
	proto.RegisterType((*LogEntry)(nil), "dynobuf.LogEntry")
}

func init() { proto.RegisterFile("log_entry.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0xc9, 0x4f, 0x8f,
	0x4f, 0xcd, 0x2b, 0x29, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xa9, 0xcc,
	0xcb, 0x4f, 0x2a, 0x4d, 0x53, 0x9a, 0xc9, 0xc8, 0xc5, 0xe1, 0x93, 0x9f, 0xee, 0x0a, 0x92, 0x13,
	0x52, 0xe5, 0xe2, 0x4b, 0x2c, 0x28, 0xc8, 0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0x8b, 0xcf,
	0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x45, 0x12, 0xf5, 0x4c, 0x11, 0x92, 0xe2,
	0xe2, 0x48, 0x49, 0x2c, 0x49, 0x4c, 0x4a, 0x2c, 0x4e, 0x95, 0x60, 0x02, 0x2b, 0x80, 0xf3, 0x85,
	0x24, 0xb8, 0xd8, 0x81, 0x36, 0x24, 0xa7, 0x16, 0x17, 0x4b, 0x30, 0x83, 0xa5, 0x60, 0x5c, 0x21,
	0x19, 0x2e, 0xce, 0x92, 0xcc, 0xdc, 0xd4, 0xe2, 0x92, 0xc4, 0xdc, 0x02, 0x09, 0x16, 0xa0, 0x1c,
	0x73, 0x10, 0x42, 0x40, 0x48, 0x88, 0x8b, 0x05, 0x64, 0x86, 0x04, 0x2b, 0x58, 0x13, 0x98, 0xed,
	0x24, 0xce, 0x25, 0x9a, 0x99, 0xaf, 0x07, 0x72, 0x69, 0x2e, 0xd0, 0xea, 0x64, 0x88, 0xd3, 0x81,
	0x8e, 0x4e, 0x62, 0x03, 0xb3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x27, 0xbb, 0xe2, 0xa5,
	0xd7, 0x00, 0x00, 0x00,
}
