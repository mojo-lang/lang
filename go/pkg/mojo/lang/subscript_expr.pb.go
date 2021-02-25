// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/lang/subscript_expr.proto

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

type SubscriptExpr struct {
	StartPosition        *Position `protobuf:"bytes,1,opt,name=start_position,json=startPosition,proto3" json:"startPosition,omitempty"`
	EndPosition          *Position `protobuf:"bytes,2,opt,name=end_position,json=endPosition,proto3" json:"endPosition,omitempty"`
	Kind                 int32     `protobuf:"varint,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Implicit             bool      `protobuf:"varint,4,opt,name=implicit,proto3" json:"implicit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *SubscriptExpr) Reset()         { *m = SubscriptExpr{} }
func (m *SubscriptExpr) String() string { return proto.CompactTextString(m) }
func (*SubscriptExpr) ProtoMessage()    {}
func (*SubscriptExpr) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b31fde3cf064b31, []int{0}
}
func (m *SubscriptExpr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriptExpr.Unmarshal(m, b)
}
func (m *SubscriptExpr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriptExpr.Marshal(b, m, deterministic)
}
func (m *SubscriptExpr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriptExpr.Merge(m, src)
}
func (m *SubscriptExpr) XXX_Size() int {
	return xxx_messageInfo_SubscriptExpr.Size(m)
}
func (m *SubscriptExpr) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriptExpr.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriptExpr proto.InternalMessageInfo

func (m *SubscriptExpr) GetStartPosition() *Position {
	if m != nil {
		return m.StartPosition
	}
	return nil
}

func (m *SubscriptExpr) GetEndPosition() *Position {
	if m != nil {
		return m.EndPosition
	}
	return nil
}

func (m *SubscriptExpr) GetKind() int32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *SubscriptExpr) GetImplicit() bool {
	if m != nil {
		return m.Implicit
	}
	return false
}

func init() {
	proto.RegisterType((*SubscriptExpr)(nil), "mojo.lang.SubscriptExpr")
}

func init() { proto.RegisterFile("mojo/lang/subscript_expr.proto", fileDescriptor_9b31fde3cf064b31) }

var fileDescriptor_9b31fde3cf064b31 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x49, 0xad, 0x52, 0xa7, 0xb6, 0x8b, 0x51, 0x6b, 0xac, 0x60, 0x82, 0x0b, 0xc9, 0x42,
	0x33, 0x50, 0x97, 0xee, 0x02, 0xee, 0xa3, 0x2e, 0x04, 0x37, 0xa1, 0xf9, 0x61, 0x1c, 0xdb, 0x99,
	0x3b, 0x4c, 0xa6, 0x50, 0x9f, 0xc4, 0xb7, 0xcb, 0x03, 0xe4, 0x29, 0x64, 0x46, 0xf3, 0x53, 0xb0,
	0x9b, 0x0b, 0x73, 0xce, 0xc7, 0x99, 0xcb, 0xb9, 0xe8, 0x9a, 0xc3, 0x27, 0x90, 0xf5, 0x52, 0x50,
	0x52, 0x6e, 0xd2, 0x32, 0x53, 0x4c, 0xea, 0xa4, 0xd8, 0x4a, 0x15, 0x4a, 0x05, 0x1a, 0xf0, 0xb1,
	0xf1, 0x43, 0xe3, 0xcf, 0xcf, 0x3a, 0xb4, 0x03, 0xe6, 0x6e, 0xa7, 0x4a, 0x28, 0x99, 0x66, 0x20,
	0x7e, 0x9d, 0x9b, 0xef, 0x01, 0x9a, 0xbc, 0x36, 0x99, 0x4f, 0x5b, 0xa9, 0xf0, 0x1b, 0x9a, 0x96,
	0x7a, 0xa9, 0x74, 0xd2, 0x90, 0xae, 0xe3, 0x3b, 0xc1, 0x78, 0x71, 0x1a, 0xb6, 0xbf, 0x84, 0xf1,
	0x9f, 0x15, 0x5d, 0xd5, 0x95, 0x77, 0x61, 0xf1, 0x46, 0xba, 0x03, 0xce, 0x74, 0xc1, 0xa5, 0xfe,
	0x7a, 0x99, 0xec, 0x18, 0xf8, 0x19, 0x9d, 0x14, 0x22, 0xef, 0x62, 0x07, 0xfb, 0x63, 0x2f, 0xeb,
	0xca, 0x3b, 0x2f, 0x44, 0xfe, 0x4f, 0xe8, 0xb8, 0x27, 0xe3, 0x5b, 0x34, 0x5c, 0x31, 0x91, 0xbb,
	0x07, 0xbe, 0x13, 0x1c, 0x46, 0xb8, 0xae, 0xbc, 0xa9, 0x79, 0xf7, 0x70, 0xeb, 0xe3, 0x05, 0x1a,
	0x31, 0x2e, 0xd7, 0x2c, 0x63, 0xda, 0x1d, 0xfa, 0x4e, 0x30, 0x8a, 0x66, 0x75, 0xe5, 0xe1, 0x46,
	0xeb, 0xf1, 0x2d, 0x17, 0x25, 0x68, 0x06, 0x8a, 0xda, 0xed, 0xec, 0x72, 0xed, 0x9a, 0x11, 0xde,
	0x29, 0x2c, 0x36, 0x3d, 0xc6, 0xce, 0x3b, 0xa1, 0x4c, 0x7f, 0x6c, 0xd2, 0x30, 0x03, 0x4e, 0x0c,
	0x7b, 0x6f, 0xeb, 0xb6, 0x83, 0x02, 0x91, 0x2b, 0x4a, 0xda, 0x23, 0x3c, 0x9a, 0x91, 0x1e, 0xd9,
	0x0b, 0x3c, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x4b, 0xb2, 0x5a, 0xde, 0x01, 0x00, 0x00,
}
