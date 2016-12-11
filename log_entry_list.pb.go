// Code generated by protoc-gen-go.
// source: log_entry_list.proto
// DO NOT EDIT!

package dynobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type LogEntryList struct {
	Timestamp int64       `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Entries   []*LogEntry `protobuf:"bytes,2,rep,name=entries" json:"entries,omitempty"`
}

func (m *LogEntryList) Reset()                    { *m = LogEntryList{} }
func (m *LogEntryList) String() string            { return proto.CompactTextString(m) }
func (*LogEntryList) ProtoMessage()               {}
func (*LogEntryList) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *LogEntryList) GetEntries() []*LogEntry {
	if m != nil {
		return m.Entries
	}
	return nil
}

func init() {
	proto.RegisterType((*LogEntryList)(nil), "dynobuf.LogEntryList")
}

func init() { proto.RegisterFile("log_entry_list.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0xc9, 0x4f, 0x8f,
	0x4f, 0xcd, 0x2b, 0x29, 0xaa, 0x8c, 0xcf, 0xc9, 0x2c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x4f, 0xa9, 0xcc, 0xcb, 0x4f, 0x2a, 0x4d, 0x93, 0xe2, 0x87, 0x4b, 0x43, 0x64, 0x94,
	0x22, 0xb9, 0x78, 0x7c, 0xf2, 0xd3, 0x5d, 0x41, 0x22, 0x3e, 0x40, 0xf5, 0x42, 0x32, 0x5c, 0x9c,
	0x25, 0x99, 0xb9, 0xa9, 0xc5, 0x25, 0x89, 0xb9, 0x05, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41,
	0x08, 0x01, 0x21, 0x6d, 0x2e, 0x76, 0x90, 0xe6, 0xcc, 0xd4, 0x62, 0x09, 0x26, 0x05, 0x66, 0x0d,
	0x6e, 0x23, 0x41, 0x3d, 0xa8, 0xc9, 0x7a, 0x30, 0x53, 0x82, 0x60, 0x2a, 0x9c, 0xc4, 0xb9, 0x44,
	0x33, 0xf3, 0xc1, 0xf2, 0xb9, 0x89, 0x25, 0x99, 0xc9, 0x10, 0x0b, 0x81, 0x4a, 0x93, 0xd8, 0xc0,
	0x2c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x70, 0x8c, 0xa9, 0xac, 0x00, 0x00, 0x00,
}
