// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/lang/stmt.proto

package lang

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Stmt struct {
	StartPosition        *Position `protobuf:"bytes,1,opt,name=start_position,json=startPosition,proto3" json:"startPosition,omitempty"`
	EndPosition          *Position `protobuf:"bytes,2,opt,name=end_position,json=endPosition,proto3" json:"endPosition,omitempty"`
	Kind                 int64     `protobuf:"varint,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Implicit             bool      `protobuf:"varint,4,opt,name=implicit,proto3" json:"implicit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Stmt) Reset()         { *m = Stmt{} }
func (m *Stmt) String() string { return proto.CompactTextString(m) }
func (*Stmt) ProtoMessage()    {}
func (*Stmt) Descriptor() ([]byte, []int) {
	return fileDescriptor_07d0963ccd3c01cf, []int{0}
}
func (m *Stmt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stmt.Unmarshal(m, b)
}
func (m *Stmt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stmt.Marshal(b, m, deterministic)
}
func (m *Stmt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stmt.Merge(m, src)
}
func (m *Stmt) XXX_Size() int {
	return xxx_messageInfo_Stmt.Size(m)
}
func (m *Stmt) XXX_DiscardUnknown() {
	xxx_messageInfo_Stmt.DiscardUnknown(m)
}

var xxx_messageInfo_Stmt proto.InternalMessageInfo

func (m *Stmt) GetStartPosition() *Position {
	if m != nil {
		return m.StartPosition
	}
	return nil
}

func (m *Stmt) GetEndPosition() *Position {
	if m != nil {
		return m.EndPosition
	}
	return nil
}

func (m *Stmt) GetKind() int64 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *Stmt) GetImplicit() bool {
	if m != nil {
		return m.Implicit
	}
	return false
}

func init() {
	proto.RegisterType((*Stmt)(nil), "mojo.lang.Stmt")
}

func init() { proto.RegisterFile("mojo/lang/stmt.proto", fileDescriptor_07d0963ccd3c01cf) }

var fileDescriptor_07d0963ccd3c01cf = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4e, 0xc3, 0x30,
	0x14, 0x44, 0xe5, 0xb6, 0x42, 0xad, 0x0b, 0x5d, 0x18, 0x28, 0x01, 0x16, 0x89, 0x58, 0xa0, 0x2c,
	0xc0, 0x96, 0xca, 0x92, 0x9d, 0x4f, 0x10, 0x40, 0x02, 0x89, 0x0d, 0x6a, 0x9b, 0xc8, 0x98, 0xd6,
	0xfe, 0x56, 0xf2, 0x59, 0x70, 0x09, 0x8e, 0x98, 0x03, 0xe4, 0x14, 0xc8, 0x86, 0xa4, 0x41, 0x82,
	0xcd, 0x97, 0xfe, 0xcc, 0xd3, 0x68, 0x34, 0xf4, 0xc8, 0xc0, 0x1b, 0x88, 0xed, 0xd2, 0x2a, 0x51,
	0xa1, 0x41, 0xee, 0x4a, 0x40, 0x60, 0x13, 0xaf, 0x72, 0xaf, 0x9e, 0x45, 0x3b, 0xc0, 0x41, 0xa5,
	0x51, 0x83, 0xfd, 0x86, 0x2e, 0x3e, 0x07, 0x74, 0xf4, 0x80, 0x06, 0xd9, 0x13, 0x9d, 0x55, 0xb8,
	0x2c, 0xf1, 0xa5, 0x05, 0x22, 0x92, 0x90, 0x74, 0xba, 0x38, 0xe4, 0x5d, 0x0c, 0xcf, 0x7e, 0x2c,
	0x79, 0xde, 0xd4, 0xf1, 0x49, 0xc0, 0x5b, 0xe9, 0x0a, 0x8c, 0xc6, 0xc2, 0x38, 0xfc, 0xb8, 0x3f,
	0xf8, 0x65, 0xb0, 0x3b, 0xba, 0x5f, 0xd8, 0x7c, 0x17, 0x3b, 0xf8, 0x3f, 0xf6, 0xb4, 0xa9, 0xe3,
	0xe3, 0xc2, 0xe6, 0x7f, 0x84, 0x4e, 0x7b, 0x32, 0xbb, 0xa4, 0xa3, 0x8d, 0xb6, 0x79, 0x34, 0x4c,
	0x48, 0x3a, 0x94, 0xac, 0xa9, 0xe3, 0x99, 0xff, 0x7b, 0x78, 0xf0, 0xd9, 0x82, 0x8e, 0xb5, 0x71,
	0x5b, 0xbd, 0xd6, 0x18, 0x8d, 0x12, 0x92, 0x8e, 0xe5, 0xbc, 0xa9, 0x63, 0xd6, 0x6a, 0x3d, 0xbe,
	0xe3, 0xe4, 0x23, 0x9d, 0x43, 0xa9, 0x42, 0xbb, 0x50, 0xae, 0xab, 0x29, 0x27, 0x7e, 0xa7, 0xcc,
	0xaf, 0x96, 0x91, 0x67, 0xa1, 0x34, 0xbe, 0xbe, 0xaf, 0xf8, 0x1a, 0x8c, 0xf0, 0xc8, 0x75, 0x18,
	0x37, 0x1c, 0x05, 0xc2, 0x6d, 0x94, 0xe8, 0x26, 0xbf, 0xf5, 0x67, 0xb5, 0x17, 0xf6, 0xbe, 0xf9,
	0x0a, 0x00, 0x00, 0xff, 0xff, 0xae, 0x74, 0x6b, 0xe3, 0xac, 0x01, 0x00, 0x00,
}
