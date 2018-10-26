// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types.proto

package checker

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import sbac "chainspace.io/prototype/sbac"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Opcode int32

const (
	Opcode_UNKNOWN Opcode = 0
	Opcode_CHECK   Opcode = 1
)

var Opcode_name = map[int32]string{
	0: "UNKNOWN",
	1: "CHECK",
}
var Opcode_value = map[string]int32{
	"UNKNOWN": 0,
	"CHECK":   1,
}

func (x Opcode) String() string {
	return proto.EnumName(Opcode_name, int32(x))
}
func (Opcode) EnumDescriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

type CheckRequest struct {
	Tx *sbac.Transaction `protobuf:"bytes,1,opt,name=tx" json:"tx,omitempty"`
}

func (m *CheckRequest) Reset()                    { *m = CheckRequest{} }
func (m *CheckRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()               {}
func (*CheckRequest) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *CheckRequest) GetTx() *sbac.Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

type CheckResponse struct {
	Ok        bool   `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *CheckResponse) Reset()                    { *m = CheckResponse{} }
func (m *CheckResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()               {}
func (*CheckResponse) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

func (m *CheckResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *CheckResponse) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*CheckRequest)(nil), "checker.CheckRequest")
	proto.RegisterType((*CheckResponse)(nil), "checker.CheckResponse")
	proto.RegisterEnum("checker.Opcode", Opcode_name, Opcode_value)
}

func init() { proto.RegisterFile("checker/types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 232 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xcd, 0x82, 0xad, 0x9d, 0x56, 0xa9, 0xeb, 0xa5, 0x14, 0x0f, 0xb1, 0xa7, 0x20, 0xb8,
	0x41, 0x3d, 0x7b, 0x0a, 0x82, 0x50, 0x48, 0x21, 0x28, 0x9e, 0x37, 0xeb, 0x98, 0x2c, 0xc1, 0x9d,
	0x35, 0xbb, 0x0b, 0xf1, 0xdf, 0x4b, 0xd6, 0x80, 0xbd, 0xcd, 0x7b, 0xf3, 0xbe, 0x79, 0x03, 0x57,
	0xaa, 0x45, 0xd5, 0x61, 0x9f, 0xfb, 0x1f, 0x8b, 0x4e, 0xd8, 0x9e, 0x3c, 0xf1, 0xf9, 0x64, 0x6e,
	0xef, 0x1a, 0xed, 0xdb, 0x50, 0x0b, 0x45, 0x5f, 0x79, 0x43, 0x0d, 0xe5, 0x71, 0x5f, 0x87, 0xcf,
	0xa8, 0xa2, 0x88, 0xd3, 0x1f, 0xb7, 0xcd, 0x54, 0x2b, 0xb5, 0x71, 0x56, 0x2a, 0x14, 0x7a, 0x0a,
	0x8f, 0x77, 0x73, 0x57, 0x4b, 0x75, 0xdc, 0xb0, 0xbb, 0x87, 0x55, 0x31, 0x76, 0x54, 0xf8, 0x1d,
	0xd0, 0x79, 0x7e, 0x03, 0xcc, 0x0f, 0x9b, 0x24, 0x4d, 0xb2, 0xe5, 0xc3, 0xa5, 0x18, 0xe3, 0xe2,
	0xb5, 0x97, 0xc6, 0x49, 0xe5, 0x35, 0x99, 0x8a, 0xf9, 0x61, 0xf7, 0x04, 0xe7, 0x13, 0xe2, 0x2c,
	0x19, 0x87, 0xfc, 0x02, 0x18, 0x75, 0x91, 0x39, 0xab, 0x18, 0x75, 0xfc, 0x1a, 0x16, 0x4e, 0x37,
	0x46, 0xfa, 0xd0, 0xe3, 0x86, 0xa5, 0x49, 0xb6, 0xaa, 0xfe, 0x8d, 0xdb, 0x14, 0x66, 0x07, 0xab,
	0xe8, 0x03, 0xf9, 0x12, 0xe6, 0x6f, 0xe5, 0xbe, 0x3c, 0xbc, 0x97, 0xeb, 0x13, 0xbe, 0x80, 0xd3,
	0xe2, 0xe5, 0xb9, 0xd8, 0xaf, 0x93, 0x7a, 0x16, 0x5f, 0x7b, 0xfc, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0xbb, 0xdf, 0x09, 0xf3, 0x13, 0x01, 0x00, 0x00,
}
