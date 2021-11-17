// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/lang/type_casting_expr.proto

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

type TypeCastingExpr struct {
	StartPosition        *Position `protobuf:"bytes,1,opt,name=start_position,json=startPosition,proto3" json:"startPosition,omitempty"`
	EndPosition          *Position `protobuf:"bytes,2,opt,name=end_position,json=endPosition,proto3" json:"endPosition,omitempty"`
	Kind                 int32     `protobuf:"varint,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Implicit             bool      `protobuf:"varint,4,opt,name=implicit,proto3" json:"implicit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TypeCastingExpr) Reset()         { *m = TypeCastingExpr{} }
func (m *TypeCastingExpr) String() string { return proto.CompactTextString(m) }
func (*TypeCastingExpr) ProtoMessage()    {}
func (*TypeCastingExpr) Descriptor() ([]byte, []int) {
	return fileDescriptor_da83a463f3e9bb10, []int{0}
}
func (m *TypeCastingExpr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TypeCastingExpr.Unmarshal(m, b)
}
func (m *TypeCastingExpr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TypeCastingExpr.Marshal(b, m, deterministic)
}
func (m *TypeCastingExpr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TypeCastingExpr.Merge(m, src)
}
func (m *TypeCastingExpr) XXX_Size() int {
	return xxx_messageInfo_TypeCastingExpr.Size(m)
}
func (m *TypeCastingExpr) XXX_DiscardUnknown() {
	xxx_messageInfo_TypeCastingExpr.DiscardUnknown(m)
}

var xxx_messageInfo_TypeCastingExpr proto.InternalMessageInfo

func (m *TypeCastingExpr) GetStartPosition() *Position {
	if m != nil {
		return m.StartPosition
	}
	return nil
}

func (m *TypeCastingExpr) GetEndPosition() *Position {
	if m != nil {
		return m.EndPosition
	}
	return nil
}

func (m *TypeCastingExpr) GetKind() int32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *TypeCastingExpr) GetImplicit() bool {
	if m != nil {
		return m.Implicit
	}
	return false
}

func init() {
	proto.RegisterType((*TypeCastingExpr)(nil), "mojo.lang.TypeCastingExpr")
}

func init() { proto.RegisterFile("mojo/lang/type_casting_expr.proto", fileDescriptor_da83a463f3e9bb10) }

var fileDescriptor_da83a463f3e9bb10 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x49, 0xad, 0x52, 0xa7, 0x5a, 0x61, 0xac, 0x35, 0xd6, 0x45, 0xa2, 0x0b, 0xc9, 0x42,
	0x33, 0x50, 0x97, 0xee, 0x22, 0xee, 0xa3, 0x08, 0x82, 0x9b, 0x90, 0x26, 0xc3, 0x38, 0xb6, 0x99,
	0x19, 0x92, 0x2b, 0x34, 0x8f, 0xe2, 0xcb, 0xe5, 0x01, 0xf2, 0x14, 0x32, 0x53, 0xf3, 0xa3, 0xe8,
	0xe6, 0xc2, 0x9c, 0xf3, 0x71, 0xe6, 0x72, 0x2e, 0xba, 0xc8, 0xe4, 0xbb, 0x24, 0xeb, 0x58, 0x30,
	0x02, 0xa5, 0xa2, 0x51, 0x12, 0x17, 0xc0, 0x05, 0x8b, 0xe8, 0x46, 0xe5, 0xbe, 0xca, 0x25, 0x48,
	0xbc, 0xaf, 0x11, 0x5f, 0x23, 0xf3, 0x69, 0x47, 0x77, 0xc0, 0xdc, 0xee, 0x54, 0x25, 0x0b, 0x0e,
	0x5c, 0x8a, 0xad, 0x73, 0xf9, 0x39, 0x40, 0x47, 0xcf, 0xa5, 0xa2, 0xf7, 0xdb, 0xd4, 0x87, 0x8d,
	0xca, 0xf1, 0x0b, 0x9a, 0x14, 0x10, 0xe7, 0x10, 0x35, 0xac, 0x6d, 0xb9, 0x96, 0x37, 0x5e, 0x1c,
	0xfb, 0xed, 0x3f, 0x7e, 0xf8, 0x6d, 0x05, 0xe7, 0x75, 0xe5, 0x9c, 0x1a, 0xbc, 0x91, 0xae, 0x65,
	0xc6, 0x81, 0x66, 0x0a, 0xca, 0xa7, 0xc3, 0x1f, 0x06, 0x7e, 0x44, 0x07, 0x54, 0xa4, 0x5d, 0xec,
	0xe0, 0xff, 0xd8, 0xb3, 0xba, 0x72, 0x4e, 0xa8, 0x48, 0xff, 0x08, 0x1d, 0xf7, 0x64, 0x7c, 0x85,
	0x86, 0x2b, 0x2e, 0x52, 0x7b, 0xc7, 0xb5, 0xbc, 0xdd, 0x00, 0xd7, 0x95, 0x33, 0xd1, 0xef, 0x1e,
	0x6e, 0x7c, 0xbc, 0x40, 0x23, 0x9e, 0xa9, 0x35, 0x4f, 0x38, 0xd8, 0x43, 0xd7, 0xf2, 0x46, 0xc1,
	0xac, 0xae, 0x1c, 0xdc, 0x68, 0x3d, 0xbe, 0xe5, 0x82, 0x18, 0xcd, 0x64, 0xce, 0xcc, 0x76, 0x66,
	0xb9, 0x76, 0xcd, 0x60, 0xfa, 0xab, 0xb2, 0x50, 0x77, 0x19, 0x5a, 0xaf, 0x84, 0x71, 0x78, 0xfb,
	0x58, 0xfa, 0x89, 0xcc, 0x88, 0xa6, 0x6f, 0x4c, 0xe5, 0x66, 0x30, 0x49, 0xd4, 0x8a, 0x91, 0xf6,
	0x10, 0x77, 0x7a, 0x2c, 0xf7, 0xcc, 0x15, 0x6e, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x98,
	0xe0, 0xd9, 0xe5, 0x01, 0x00, 0x00,
}
