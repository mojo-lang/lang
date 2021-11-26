// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/lang/error_expr.proto

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

type ErrorExpr struct {
	StartPosition        *Position `protobuf:"bytes,1,opt,name=start_position,json=startPosition,proto3" json:"startPosition,omitempty"`
	EndPosition          *Position `protobuf:"bytes,2,opt,name=end_position,json=endPosition,proto3" json:"endPosition,omitempty"`
	Kind                 int32     `protobuf:"varint,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Implicit             bool      `protobuf:"varint,4,opt,name=implicit,proto3" json:"implicit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ErrorExpr) Reset()         { *m = ErrorExpr{} }
func (m *ErrorExpr) String() string { return proto.CompactTextString(m) }
func (*ErrorExpr) ProtoMessage()    {}
func (*ErrorExpr) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6edf5e430cc806b, []int{0}
}
func (m *ErrorExpr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ErrorExpr.Unmarshal(m, b)
}
func (m *ErrorExpr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ErrorExpr.Marshal(b, m, deterministic)
}
func (m *ErrorExpr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ErrorExpr.Merge(m, src)
}
func (m *ErrorExpr) XXX_Size() int {
	return xxx_messageInfo_ErrorExpr.Size(m)
}
func (m *ErrorExpr) XXX_DiscardUnknown() {
	xxx_messageInfo_ErrorExpr.DiscardUnknown(m)
}

var xxx_messageInfo_ErrorExpr proto.InternalMessageInfo

func (m *ErrorExpr) GetStartPosition() *Position {
	if m != nil {
		return m.StartPosition
	}
	return nil
}

func (m *ErrorExpr) GetEndPosition() *Position {
	if m != nil {
		return m.EndPosition
	}
	return nil
}

func (m *ErrorExpr) GetKind() int32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *ErrorExpr) GetImplicit() bool {
	if m != nil {
		return m.Implicit
	}
	return false
}

func init() {
	proto.RegisterType((*ErrorExpr)(nil), "mojo.lang.ErrorExpr")
}

func init() { proto.RegisterFile("mojo/lang/error_expr.proto", fileDescriptor_c6edf5e430cc806b) }

var fileDescriptor_c6edf5e430cc806b = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0xe9, 0x9c, 0xb2, 0x65, 0xda, 0x43, 0xd4, 0x59, 0xeb, 0xa1, 0xc5, 0x83, 0xf4, 0xa0,
	0x0d, 0xcc, 0xa3, 0xb7, 0xc2, 0xee, 0xd5, 0x8b, 0xa0, 0x87, 0xb1, 0xad, 0x21, 0xc6, 0x2d, 0xf9,
	0x87, 0x2c, 0xc2, 0xfc, 0x1c, 0x7e, 0xbf, 0x7e, 0x80, 0x7e, 0x0a, 0x49, 0xb4, 0x59, 0x05, 0x77,
	0xf9, 0x43, 0xde, 0xfb, 0xf1, 0xf2, 0x78, 0x28, 0x16, 0xf0, 0x0e, 0x64, 0x3d, 0x97, 0x8c, 0x50,
	0xad, 0x41, 0xcf, 0xe8, 0x56, 0xe9, 0x5c, 0x69, 0x30, 0x80, 0x87, 0xd6, 0xcb, 0xad, 0x17, 0x9f,
	0x75, 0x30, 0x0f, 0xc4, 0xd1, 0x4e, 0x55, 0xb0, 0xe1, 0x86, 0x83, 0xfc, 0x71, 0xae, 0xbf, 0x7a,
	0x68, 0x38, 0xb5, 0x79, 0xd3, 0xad, 0xd2, 0xf8, 0x19, 0x85, 0x1b, 0x33, 0xd7, 0x66, 0xd6, 0x52,
	0x51, 0x90, 0x06, 0xd9, 0x68, 0x72, 0x9a, 0xfb, 0x1f, 0xf2, 0xf2, 0xd7, 0x2a, 0xae, 0x9a, 0x3a,
	0xb9, 0x70, 0x78, 0x2b, 0xdd, 0x82, 0xe0, 0x86, 0x0a, 0x65, 0x3e, 0x9f, 0x4e, 0xfe, 0x18, 0xf8,
	0x11, 0x1d, 0x53, 0x59, 0xed, 0x62, 0x7b, 0xfb, 0x63, 0x2f, 0x9b, 0x3a, 0x39, 0xa7, 0xb2, 0xfa,
	0x27, 0x74, 0xd4, 0x91, 0xf1, 0x0d, 0xea, 0xaf, 0xb8, 0xac, 0xa2, 0x83, 0x34, 0xc8, 0x0e, 0x0b,
	0xdc, 0xd4, 0x49, 0x68, 0xdf, 0x1d, 0xdc, 0xf9, 0x78, 0x82, 0x06, 0x5c, 0xa8, 0x35, 0x5f, 0x72,
	0x13, 0xf5, 0xd3, 0x20, 0x1b, 0x14, 0xe3, 0xa6, 0x4e, 0x70, 0xab, 0x75, 0x78, 0xcf, 0x15, 0xaf,
	0x68, 0x0c, 0x9a, 0xb9, 0x76, 0xae, 0x9c, 0xaf, 0x59, 0x84, 0x7e, 0xac, 0xd2, 0xee, 0x57, 0x06,
	0x2f, 0x84, 0x71, 0xf3, 0xf6, 0xb1, 0xc8, 0x97, 0x20, 0x88, 0xe5, 0xee, 0xdc, 0xcc, 0xee, 0x30,
	0x20, 0x6a, 0xc5, 0x88, 0x1f, 0xff, 0xc1, 0x9e, 0xc5, 0x91, 0x5b, 0xfe, 0xfe, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0xae, 0x21, 0x90, 0x3a, 0xd2, 0x01, 0x00, 0x00,
}