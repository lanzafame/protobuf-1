// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imports/test_b_1/m2.proto

package beta

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

type M2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2) Reset()         { *m = M2{} }
func (m *M2) String() string { return proto.CompactTextString(m) }
func (*M2) ProtoMessage()    {}
func (*M2) Descriptor() ([]byte, []int) {
	return fileDescriptor_m2_0c59cab35ba1b0d8, []int{0}
}
func (m *M2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2.Unmarshal(m, b)
}
func (m *M2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2.Marshal(b, m, deterministic)
}
func (dst *M2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2.Merge(dst, src)
}
func (m *M2) XXX_Size() int {
	return xxx_messageInfo_M2.Size(m)
}
func (m *M2) XXX_DiscardUnknown() {
	xxx_messageInfo_M2.DiscardUnknown(m)
}

var xxx_messageInfo_M2 proto.InternalMessageInfo

func init() {
	proto.RegisterType((*M2)(nil), "test.b.part2.M2")
}

func init() { proto.RegisterFile("imports/test_b_1/m2.proto", fileDescriptor_m2_0c59cab35ba1b0d8) }

var fileDescriptor_m2_0c59cab35ba1b0d8 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x29, 0xd6, 0x2f, 0x49, 0x2d, 0x2e, 0x89, 0x4f, 0x8a, 0x37, 0xd4, 0xcf, 0x35, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x01, 0x09, 0xe9, 0x25, 0xe9, 0x15, 0x24, 0x16, 0x95,
	0x18, 0x29, 0xb1, 0x70, 0x31, 0xf9, 0x1a, 0x39, 0x79, 0x46, 0xb9, 0xa7, 0x67, 0x96, 0x64, 0x94,
	0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xe7, 0x24, 0xe6, 0xa5, 0xeb, 0x83, 0x95, 0x27,
	0x95, 0xa6, 0x41, 0x18, 0xc9, 0xba, 0xe9, 0xa9, 0x79, 0xba, 0xe9, 0xf9, 0x60, 0x13, 0x53, 0x12,
	0x4b, 0x12, 0xf5, 0xd1, 0xad, 0xb0, 0x4e, 0x4a, 0x2d, 0x49, 0x4c, 0x62, 0x03, 0xab, 0x36, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x44, 0x29, 0xbe, 0x6d, 0x82, 0x00, 0x00, 0x00,
}
