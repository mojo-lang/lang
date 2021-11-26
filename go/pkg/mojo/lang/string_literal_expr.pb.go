// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/lang/string_literal_expr.proto

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

type StringLiteralExpr struct {
	StartPosition        *Position `protobuf:"bytes,1,opt,name=start_position,json=startPosition,proto3" json:"startPosition,omitempty"`
	EndPosition          *Position `protobuf:"bytes,2,opt,name=end_position,json=endPosition,proto3" json:"endPosition,omitempty"`
	Kind                 int32     `protobuf:"varint,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Implicit             bool      `protobuf:"varint,4,opt,name=implicit,proto3" json:"implicit,omitempty"`
	Value                string    `protobuf:"bytes,15,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *StringLiteralExpr) Reset()         { *m = StringLiteralExpr{} }
func (m *StringLiteralExpr) String() string { return proto.CompactTextString(m) }
func (*StringLiteralExpr) ProtoMessage()    {}
func (*StringLiteralExpr) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a702913cb56b0ad, []int{0}
}
func (m *StringLiteralExpr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringLiteralExpr.Unmarshal(m, b)
}
func (m *StringLiteralExpr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringLiteralExpr.Marshal(b, m, deterministic)
}
func (m *StringLiteralExpr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringLiteralExpr.Merge(m, src)
}
func (m *StringLiteralExpr) XXX_Size() int {
	return xxx_messageInfo_StringLiteralExpr.Size(m)
}
func (m *StringLiteralExpr) XXX_DiscardUnknown() {
	xxx_messageInfo_StringLiteralExpr.DiscardUnknown(m)
}

var xxx_messageInfo_StringLiteralExpr proto.InternalMessageInfo

func (m *StringLiteralExpr) GetStartPosition() *Position {
	if m != nil {
		return m.StartPosition
	}
	return nil
}

func (m *StringLiteralExpr) GetEndPosition() *Position {
	if m != nil {
		return m.EndPosition
	}
	return nil
}

func (m *StringLiteralExpr) GetKind() int32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *StringLiteralExpr) GetImplicit() bool {
	if m != nil {
		return m.Implicit
	}
	return false
}

func (m *StringLiteralExpr) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*StringLiteralExpr)(nil), "mojo.lang.StringLiteralExpr")
}

func init() {
	proto.RegisterFile("mojo/lang/string_literal_expr.proto", fileDescriptor_6a702913cb56b0ad)
}

var fileDescriptor_6a702913cb56b0ad = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0xe9, 0xdc, 0x64, 0xcb, 0x74, 0xc3, 0x4c, 0x67, 0x9d, 0x42, 0x8b, 0x82, 0x54, 0xd0,
	0x06, 0xe6, 0xa5, 0x77, 0x01, 0xef, 0xbc, 0x98, 0x7a, 0x21, 0x78, 0x33, 0xba, 0x2d, 0xc4, 0xb8,
	0x36, 0x09, 0x59, 0x26, 0xf3, 0x05, 0x7d, 0x8c, 0x3e, 0x40, 0x9f, 0x42, 0x92, 0xda, 0x3f, 0x32,
	0xbd, 0x39, 0xd0, 0xef, 0xfb, 0x9d, 0xd3, 0x9c, 0xf3, 0x81, 0x8b, 0x44, 0xbc, 0x0b, 0x14, 0x47,
	0x9c, 0xa2, 0x95, 0x56, 0x8c, 0xd3, 0x69, 0xcc, 0x34, 0x51, 0x51, 0x3c, 0x25, 0x1b, 0xa9, 0x42,
	0xa9, 0x84, 0x16, 0xb0, 0x63, 0xa0, 0xd0, 0x40, 0xa3, 0xc3, 0x8a, 0xaf, 0x80, 0xd1, 0x59, 0xa5,
	0x6e, 0xb7, 0x8f, 0xdc, 0xca, 0x95, 0x62, 0xc5, 0x34, 0x13, 0x3c, 0x77, 0xce, 0xbf, 0x1a, 0xe0,
	0xe0, 0xd9, 0xfe, 0xf6, 0x21, 0x6f, 0xbb, 0xdf, 0x48, 0x05, 0x5f, 0x40, 0x6f, 0xa5, 0x23, 0xa5,
	0xa7, 0x05, 0xed, 0x3a, 0xbe, 0x13, 0x74, 0xc7, 0x83, 0xb0, 0x7c, 0x47, 0x38, 0xf9, 0xb1, 0xf0,
	0x69, 0x96, 0x7a, 0xc7, 0x16, 0x2f, 0xa4, 0x6b, 0x91, 0x30, 0x4d, 0x12, 0xa9, 0x3f, 0x9f, 0xf6,
	0x7f, 0x19, 0xf0, 0x11, 0xec, 0x11, 0xbe, 0xa8, 0xc6, 0x36, 0xfe, 0x1f, 0x7b, 0x92, 0xa5, 0xde,
	0x11, 0xe1, 0x8b, 0x3f, 0x86, 0x76, 0x6b, 0x32, 0xbc, 0x04, 0xcd, 0x25, 0xe3, 0x0b, 0x77, 0xc7,
	0x77, 0x82, 0x16, 0x86, 0x59, 0xea, 0xf5, 0xcc, 0x77, 0x0d, 0xb7, 0x3e, 0x1c, 0x83, 0x36, 0x4b,
	0x64, 0xcc, 0xe6, 0x4c, 0xbb, 0x4d, 0xdf, 0x09, 0xda, 0x78, 0x98, 0xa5, 0x1e, 0x2c, 0xb4, 0x1a,
	0x5f, 0x72, 0xf0, 0x0a, 0xb4, 0x3e, 0xa2, 0x78, 0x4d, 0xdc, 0xbe, 0xef, 0x04, 0x1d, 0x3c, 0xc8,
	0x52, 0xaf, 0x6f, 0x85, 0x1a, 0x9d, 0x13, 0x78, 0x0e, 0x86, 0x42, 0x51, 0xbb, 0x88, 0xdd, 0xa3,
	0xdc, 0x08, 0x0f, 0xb7, 0xee, 0x3b, 0x31, 0xa7, 0x9f, 0x38, 0xaf, 0x88, 0x32, 0xfd, 0xb6, 0x9e,
	0x85, 0x73, 0x91, 0x20, 0xc3, 0xdf, 0xe4, 0xf9, 0x99, 0x42, 0x05, 0x92, 0x4b, 0x8a, 0xca, 0xdc,
	0xee, 0x4c, 0x99, 0xed, 0xda, 0xd0, 0x6e, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x3c, 0x21,
	0x0a, 0x34, 0x02, 0x00, 0x00,
}