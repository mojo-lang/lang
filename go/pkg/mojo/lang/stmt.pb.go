// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/lang/stmt.proto

package lang

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Stmt struct {
	StartPosition        *Position `protobuf:"bytes,1,opt,name=start_position,json=startPosition,proto3" json:"startPosition,omitempty"`
	EndPosition          *Position `protobuf:"bytes,2,opt,name=end_position,json=endPosition,proto3" json:"endPosition,omitempty"`
	Kind                 int64     `protobuf:"varint,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Implicit             bool      `protobuf:"varint,4,opt,name=implicit,proto3" json:"implicit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-" gorm:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte    `json:"-" gorm:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32     `json:"-" gorm:"-" xml:"-" bson:"-"`
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
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc9, 0xcd, 0xcf, 0xca,
	0xd7, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0x2f, 0x2e, 0xc9, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x04, 0x89, 0xea, 0x81, 0x44, 0xa5, 0x24, 0x10, 0x0a, 0x0a, 0xf2, 0x8b, 0x33, 0x4b,
	0x32, 0xf3, 0xf3, 0x20, 0x8a, 0x94, 0xfa, 0x98, 0xb8, 0x58, 0x82, 0x4b, 0x72, 0x4b, 0x84, 0xc2,
	0xb9, 0xf8, 0x8a, 0x4b, 0x12, 0x8b, 0x4a, 0xe2, 0x61, 0x0a, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8,
	0x8d, 0x84, 0xf5, 0xe0, 0xc6, 0xe8, 0x05, 0x40, 0xa5, 0x9c, 0xa4, 0x5f, 0xdd, 0x93, 0x17, 0x07,
	0x2b, 0x87, 0x09, 0xe9, 0xe4, 0xe7, 0x66, 0x96, 0xa4, 0xe6, 0x16, 0x94, 0x54, 0x06, 0xf1, 0xa2,
	0x48, 0x08, 0x05, 0x72, 0xf1, 0xa4, 0xe6, 0xa5, 0x20, 0x8c, 0x65, 0xc2, 0x6d, 0xac, 0xe4, 0xab,
	0x7b, 0xf2, 0xa2, 0xa9, 0x79, 0x29, 0x58, 0x0c, 0xe5, 0x46, 0x12, 0x16, 0x52, 0xe3, 0x62, 0xc9,
	0xce, 0xcc, 0x4b, 0x91, 0x60, 0x56, 0x60, 0xd4, 0x60, 0x76, 0x12, 0x7a, 0x75, 0x4f, 0x9e, 0x0f,
	0xc4, 0x47, 0x52, 0x0e, 0x96, 0x17, 0x32, 0xe2, 0xe2, 0xc8, 0xcc, 0x2d, 0xc8, 0xc9, 0x4c, 0xce,
	0x2c, 0x91, 0x60, 0x51, 0x60, 0xd4, 0xe0, 0x70, 0x12, 0x7b, 0x75, 0x4f, 0x5e, 0x08, 0x26, 0x86,
	0xa4, 0x1e, 0xae, 0xce, 0x29, 0x6a, 0xc6, 0x63, 0x39, 0x06, 0x2e, 0xb1, 0xfc, 0xa2, 0x74, 0xb0,
	0x0b, 0xc1, 0x0e, 0x84, 0x3b, 0xd5, 0x89, 0x13, 0x14, 0x56, 0x01, 0xa0, 0x90, 0x0b, 0x60, 0x8c,
	0xd2, 0x4f, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x07, 0x29, 0xd1, 0x05,
	0x07, 0x30, 0x98, 0x48, 0xcf, 0xd7, 0x2f, 0xc8, 0x4e, 0xd7, 0x87, 0x07, 0xbb, 0x35, 0x88, 0x48,
	0x62, 0x03, 0x87, 0xb9, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x85, 0xba, 0xaf, 0xb0, 0x01,
	0x00, 0x00,
}
