// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/lang/error_stmt.proto

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

type ErrorStmt struct {
	StartPosition        *Position `protobuf:"bytes,1,opt,name=start_position,json=startPosition,proto3" json:"startPosition,omitempty"`
	EndPosition          *Position `protobuf:"bytes,2,opt,name=end_position,json=endPosition,proto3" json:"endPosition,omitempty"`
	Kind                 int64     `protobuf:"varint,4,opt,name=kind,proto3" json:"kind,omitempty"`
	Implicit             bool      `protobuf:"varint,5,opt,name=implicit,proto3" json:"implicit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-" gorm:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte    `json:"-" gorm:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32     `json:"-" gorm:"-" xml:"-" bson:"-"`
}

func (m *ErrorStmt) Reset()         { *m = ErrorStmt{} }
func (m *ErrorStmt) String() string { return proto.CompactTextString(m) }
func (*ErrorStmt) ProtoMessage()    {}
func (*ErrorStmt) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd477d620c9f8fa8, []int{0}
}
func (m *ErrorStmt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorStmt.Unmarshal(m, b)
}
func (m *ErrorStmt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorStmt.Marshal(b, m, deterministic)
}
func (m *ErrorStmt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorStmt.Merge(m, src)
}
func (m *ErrorStmt) XXX_Size() int {
	return xxx_messageInfo_ErrorStmt.Size(m)
}
func (m *ErrorStmt) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorStmt.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorStmt proto.InternalMessageInfo

func (m *ErrorStmt) GetStartPosition() *Position {
	if m != nil {
		return m.StartPosition
	}
	return nil
}

func (m *ErrorStmt) GetEndPosition() *Position {
	if m != nil {
		return m.EndPosition
	}
	return nil
}

func (m *ErrorStmt) GetKind() int64 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *ErrorStmt) GetImplicit() bool {
	if m != nil {
		return m.Implicit
	}
	return false
}

func init() {
	proto.RegisterType((*ErrorStmt)(nil), "mojo.lang.ErrorStmt")
}

func init() { proto.RegisterFile("mojo/lang/error_stmt.proto", fileDescriptor_dd477d620c9f8fa8) }

var fileDescriptor_dd477d620c9f8fa8 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x18, 0x84, 0x49, 0x29, 0xa8, 0x75, 0x21, 0x83, 0x81, 0x12, 0x8a, 0x44, 0x22, 0x06, 0x94, 0x01,
	0x62, 0xa9, 0x8c, 0x6c, 0x91, 0xd8, 0x03, 0x0c, 0x48, 0x2c, 0x51, 0xdb, 0x44, 0xc1, 0xb4, 0xf6,
	0x6f, 0x39, 0x3f, 0x03, 0xcf, 0xc1, 0xc2, 0xd3, 0x65, 0x61, 0xcb, 0x53, 0x20, 0x1b, 0xea, 0x06,
	0x09, 0x96, 0x5f, 0xf2, 0xdd, 0xe7, 0xd3, 0xe9, 0xc8, 0x44, 0xc0, 0x0b, 0xb0, 0xd5, 0x4c, 0x56,
	0xac, 0xd4, 0x1a, 0x74, 0x5e, 0xa3, 0xc0, 0x44, 0x69, 0x40, 0xa0, 0x43, 0xe3, 0x25, 0xc6, 0x9b,
	0x1c, 0x6e, 0x30, 0x73, 0xbe, 0x81, 0xf3, 0xf7, 0x1e, 0x19, 0xde, 0x9a, 0x5f, 0x0f, 0x28, 0x90,
	0x3e, 0x12, 0xbf, 0xc6, 0x99, 0xc6, 0x5c, 0x41, 0xcd, 0x91, 0x83, 0x0c, 0xbc, 0xc8, 0x8b, 0x47,
	0xd3, 0x83, 0xc4, 0xe5, 0x24, 0xd9, 0x8f, 0x95, 0x9e, 0xb6, 0x4d, 0x78, 0x6c, 0xf1, 0xb5, 0x74,
	0x09, 0x82, 0x63, 0x29, 0x14, 0xbe, 0xdd, 0xef, 0xff, 0x32, 0xe8, 0x1d, 0xd9, 0x2b, 0x65, 0xb1,
	0x89, 0xed, 0xfd, 0x1f, 0x7b, 0xd2, 0x36, 0xe1, 0x51, 0x29, 0x8b, 0x3f, 0x42, 0x47, 0x1d, 0x99,
	0x5e, 0x90, 0xfe, 0x92, 0xcb, 0x22, 0xe8, 0x47, 0x5e, 0xbc, 0x9d, 0xd2, 0xb6, 0x09, 0x7d, 0xf3,
	0xee, 0xe0, 0xd6, 0xa7, 0x53, 0x32, 0xe0, 0x42, 0xad, 0xf8, 0x82, 0x63, 0xb0, 0x13, 0x79, 0xf1,
	0x20, 0x1d, 0xb7, 0x4d, 0x48, 0xd7, 0x5a, 0x87, 0x77, 0x5c, 0x9a, 0x7f, 0x7c, 0x9e, 0x6d, 0x91,
	0x31, 0xe8, 0xca, 0x36, 0xb4, 0x05, 0x5d, 0xd5, 0xd4, 0x77, 0x83, 0x65, 0x66, 0xc3, 0xcc, 0x7b,
	0x62, 0x15, 0xc7, 0xe7, 0xd7, 0x79, 0xb2, 0x00, 0xc1, 0x0c, 0x77, 0xe5, 0x66, 0x66, 0x15, 0x30,
	0xb5, 0xac, 0x98, 0x1b, 0xff, 0xc6, 0x9c, 0xf9, 0xae, 0x5d, 0xff, 0xfa, 0x2b, 0x00, 0x00, 0xff,
	0xff, 0xaa, 0x22, 0x7f, 0xd8, 0xbc, 0x01, 0x00, 0x00,
}
