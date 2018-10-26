// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types.proto

package service

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CONNECTION int32

const (
	CONNECTION_UNKNOWN   CONNECTION = 0
	CONNECTION_BROADCAST CONNECTION = 1
	CONNECTION_SBAC      CONNECTION = 2
	CONNECTION_KV        CONNECTION = 3
	CONNECTION_CHECKER   CONNECTION = 4
)

var CONNECTION_name = map[int32]string{
	0: "UNKNOWN",
	1: "BROADCAST",
	2: "SBAC",
	3: "KV",
	4: "CHECKER",
}
var CONNECTION_value = map[string]int32{
	"UNKNOWN":   0,
	"BROADCAST": 1,
	"SBAC":      2,
	"KV":        3,
	"CHECKER":   4,
}

func (x CONNECTION) String() string {
	return proto.EnumName(CONNECTION_name, int32(x))
}
func (CONNECTION) EnumDescriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

type Hello struct {
	Agent     string     `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`
	Payload   []byte     `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature []byte     `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Type      CONNECTION `protobuf:"varint,4,opt,name=type,proto3,enum=service.CONNECTION" json:"type,omitempty"`
}

func (m *Hello) Reset()                    { *m = Hello{} }
func (m *Hello) String() string            { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()               {}
func (*Hello) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{0} }

func (m *Hello) GetAgent() string {
	if m != nil {
		return m.Agent
	}
	return ""
}

func (m *Hello) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Hello) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Hello) GetType() CONNECTION {
	if m != nil {
		return m.Type
	}
	return CONNECTION_UNKNOWN
}

type HelloInfo struct {
	Client    uint64    `protobuf:"varint,1,opt,name=client,proto3" json:"client,omitempty"`
	Nonce     []byte    `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Server    uint64    `protobuf:"varint,3,opt,name=server,proto3" json:"server,omitempty"`
	Timestamp time.Time `protobuf:"bytes,4,opt,name=timestamp,stdtime" json:"timestamp"`
}

func (m *HelloInfo) Reset()                    { *m = HelloInfo{} }
func (m *HelloInfo) String() string            { return proto.CompactTextString(m) }
func (*HelloInfo) ProtoMessage()               {}
func (*HelloInfo) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{1} }

func (m *HelloInfo) GetClient() uint64 {
	if m != nil {
		return m.Client
	}
	return 0
}

func (m *HelloInfo) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *HelloInfo) GetServer() uint64 {
	if m != nil {
		return m.Server
	}
	return 0
}

func (m *HelloInfo) GetTimestamp() time.Time {
	if m != nil {
		return m.Timestamp
	}
	return time.Time{}
}

type Message struct {
	ID      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Opcode  int32  `protobuf:"varint,3,opt,name=opcode,proto3" json:"opcode,omitempty"`
	Payload []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptorTypes, []int{2} }

func (m *Message) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Message) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Message) GetOpcode() int32 {
	if m != nil {
		return m.Opcode
	}
	return 0
}

func (m *Message) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*Hello)(nil), "service.Hello")
	proto.RegisterType((*HelloInfo)(nil), "service.HelloInfo")
	proto.RegisterType((*Message)(nil), "service.Message")
	proto.RegisterEnum("service.CONNECTION", CONNECTION_name, CONNECTION_value)
}

func init() { proto.RegisterFile("service/types.proto", fileDescriptorTypes) }

var fileDescriptorTypes = []byte{
	// 388 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xd1, 0xca, 0x9b, 0x30,
	0x18, 0x86, 0x7f, 0xfd, 0xad, 0xd6, 0xaf, 0xdb, 0x90, 0x74, 0x0c, 0x29, 0x03, 0x4b, 0x4f, 0x56,
	0x06, 0xb3, 0xd0, 0x5d, 0x41, 0xb5, 0x65, 0x2d, 0x65, 0x0a, 0x69, 0xb7, 0x1d, 0x5b, 0x4d, 0xb3,
	0x80, 0x35, 0xa2, 0x76, 0xd0, 0x93, 0x5d, 0xc3, 0x4e, 0x76, 0x4f, 0xbb, 0x8a, 0x0e, 0x76, 0x25,
	0x23, 0x89, 0xce, 0xfd, 0x67, 0x3e, 0xc9, 0x1b, 0xbf, 0xe7, 0x7b, 0x61, 0x5c, 0x93, 0xea, 0x1b,
	0x4b, 0xc9, 0xa2, 0xb9, 0x95, 0xa4, 0xf6, 0xcb, 0x8a, 0x37, 0x1c, 0x59, 0xed, 0xe1, 0xe4, 0x1d,
	0x65, 0xcd, 0xd7, 0xeb, 0xc9, 0x4f, 0xf9, 0x65, 0x41, 0x39, 0xe5, 0x0b, 0x79, 0x7f, 0xba, 0x9e,
	0x25, 0x49, 0x90, 0x5f, 0xea, 0xdd, 0xc4, 0xa3, 0x9c, 0xd3, 0x9c, 0xf4, 0xa9, 0x86, 0x5d, 0x48,
	0xdd, 0x24, 0x97, 0x52, 0x05, 0x66, 0xdf, 0x61, 0xb0, 0x25, 0x79, 0xce, 0xd1, 0x4b, 0x18, 0x24,
	0x94, 0x14, 0x8d, 0xab, 0x4d, 0xb5, 0xb9, 0x8d, 0x15, 0x20, 0x17, 0xac, 0x32, 0xb9, 0xe5, 0x3c,
	0xc9, 0x5c, 0x7d, 0xaa, 0xcd, 0x9f, 0xe1, 0x0e, 0xd1, 0x6b, 0xb0, 0x6b, 0x46, 0x8b, 0xa4, 0xb9,
	0x56, 0xc4, 0x7d, 0x94, 0x77, 0xfd, 0x01, 0x7a, 0x03, 0x86, 0xd0, 0x77, 0x8d, 0xa9, 0x36, 0x7f,
	0xb1, 0x1c, 0xfb, 0xad, 0xbe, 0x1f, 0xc6, 0x51, 0xb4, 0x09, 0x8f, 0xbb, 0x38, 0xc2, 0x32, 0x30,
	0xfb, 0xa9, 0x81, 0x2d, 0x05, 0x76, 0xc5, 0x99, 0xa3, 0x57, 0x60, 0xa6, 0x39, 0xeb, 0x2c, 0x0c,
	0xdc, 0x92, 0x90, 0x2b, 0x78, 0x91, 0x92, 0x56, 0x42, 0x81, 0x48, 0x8b, 0xff, 0x92, 0x4a, 0xce,
	0x37, 0x70, 0x4b, 0x28, 0x00, 0xfb, 0xdf, 0x9a, 0xd2, 0x60, 0xb4, 0x9c, 0xf8, 0xaa, 0x08, 0xbf,
	0x2b, 0xc2, 0x3f, 0x76, 0x89, 0x60, 0xf8, 0xeb, 0xee, 0x3d, 0xfc, 0xf8, 0xed, 0x69, 0xb8, 0x7f,
	0x36, 0x63, 0x60, 0x7d, 0x24, 0x75, 0x9d, 0x50, 0x31, 0x46, 0x67, 0x99, 0x12, 0x0a, 0xcc, 0x3f,
	0x77, 0x4f, 0xdf, 0xad, 0xb1, 0xce, 0x32, 0x21, 0x45, 0xaa, 0x8a, 0x57, 0x52, 0xca, 0xc6, 0x0a,
	0x84, 0x14, 0x2f, 0x53, 0x9e, 0xa9, 0x52, 0x06, 0xb8, 0xa5, 0xff, 0x9b, 0x34, 0x9e, 0x34, 0xf9,
	0xf6, 0x03, 0x40, 0x5f, 0x0b, 0x1a, 0x81, 0xf5, 0x29, 0xda, 0x47, 0xf1, 0x97, 0xc8, 0x79, 0x40,
	0xcf, 0xc1, 0x0e, 0x70, 0xbc, 0x5a, 0x87, 0xab, 0xc3, 0xd1, 0xd1, 0xd0, 0x10, 0x8c, 0x43, 0xb0,
	0x0a, 0x1d, 0x1d, 0x99, 0xa0, 0xef, 0x3f, 0x3b, 0x8f, 0x22, 0x1d, 0x6e, 0x37, 0xe1, 0x7e, 0x83,
	0x1d, 0xe3, 0x64, 0xca, 0xe5, 0xde, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xb5, 0x92, 0x2a,
	0x42, 0x02, 0x00, 0x00,
}
