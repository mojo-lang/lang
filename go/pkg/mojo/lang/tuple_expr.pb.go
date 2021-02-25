// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/lang/tuple_expr.proto

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

type TupleExpr struct {
	StartPosition        *Position `protobuf:"bytes,1,opt,name=start_position,json=startPosition,proto3" json:"startPosition,omitempty"`
	EndPosition          *Position `protobuf:"bytes,2,opt,name=end_position,json=endPosition,proto3" json:"endPosition,omitempty"`
	Kind                 int32     `protobuf:"varint,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Implicit             bool      `protobuf:"varint,4,opt,name=implicit,proto3" json:"implicit,omitempty"`
	HasElementNames      bool      `protobuf:"varint,20,opt,name=has_element_names,json=hasElementNames,proto3" json:"hasElementNames,omitempty"`
	HasElementLocations  bool      `protobuf:"varint,21,opt,name=has_element_locations,json=hasElementLocations,proto3" json:"hasElementLocations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TupleExpr) Reset()         { *m = TupleExpr{} }
func (m *TupleExpr) String() string { return proto.CompactTextString(m) }
func (*TupleExpr) ProtoMessage()    {}
func (*TupleExpr) Descriptor() ([]byte, []int) {
	return fileDescriptor_faf81656592390b5, []int{0}
}
func (m *TupleExpr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TupleExpr.Unmarshal(m, b)
}
func (m *TupleExpr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TupleExpr.Marshal(b, m, deterministic)
}
func (m *TupleExpr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TupleExpr.Merge(m, src)
}
func (m *TupleExpr) XXX_Size() int {
	return xxx_messageInfo_TupleExpr.Size(m)
}
func (m *TupleExpr) XXX_DiscardUnknown() {
	xxx_messageInfo_TupleExpr.DiscardUnknown(m)
}

var xxx_messageInfo_TupleExpr proto.InternalMessageInfo

func (m *TupleExpr) GetStartPosition() *Position {
	if m != nil {
		return m.StartPosition
	}
	return nil
}

func (m *TupleExpr) GetEndPosition() *Position {
	if m != nil {
		return m.EndPosition
	}
	return nil
}

func (m *TupleExpr) GetKind() int32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *TupleExpr) GetImplicit() bool {
	if m != nil {
		return m.Implicit
	}
	return false
}

func (m *TupleExpr) GetHasElementNames() bool {
	if m != nil {
		return m.HasElementNames
	}
	return false
}

func (m *TupleExpr) GetHasElementLocations() bool {
	if m != nil {
		return m.HasElementLocations
	}
	return false
}

func init() {
	proto.RegisterType((*TupleExpr)(nil), "mojo.lang.TupleExpr")
}

func init() { proto.RegisterFile("mojo/lang/tuple_expr.proto", fileDescriptor_faf81656592390b5) }

var fileDescriptor_faf81656592390b5 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xdd, 0x4a, 0xc3, 0x30,
	0x14, 0xc7, 0xa9, 0x9b, 0xb2, 0x65, 0x3a, 0x31, 0xfb, 0xb0, 0x9b, 0x8c, 0x4e, 0x2f, 0x64, 0x17,
	0xda, 0xc2, 0xbc, 0xf4, 0xae, 0xb0, 0x0b, 0x41, 0x64, 0x8a, 0x22, 0xe8, 0x45, 0xe9, 0xb6, 0xd0,
	0xc5, 0x35, 0x1f, 0xb4, 0x19, 0xcc, 0x07, 0xf3, 0x75, 0xfa, 0x00, 0x79, 0x0a, 0x49, 0x66, 0xdb,
	0x28, 0xf3, 0x26, 0xd0, 0xf3, 0xfb, 0xe7, 0xd7, 0x70, 0xce, 0x01, 0x7d, 0xc2, 0x3e, 0x98, 0x17,
	0x87, 0x34, 0xf2, 0xc4, 0x9a, 0xc7, 0x28, 0x40, 0x1b, 0x9e, 0xb8, 0x3c, 0x61, 0x82, 0xc1, 0xba,
	0x62, 0xae, 0x62, 0xfd, 0x76, 0x19, 0x2b, 0x03, 0x7d, 0xbb, 0xac, 0x72, 0x96, 0x62, 0x81, 0x19,
	0xdd, 0x92, 0x8b, 0xaf, 0x0a, 0xa8, 0x3f, 0x2b, 0xdf, 0x64, 0xc3, 0x13, 0xf8, 0x0a, 0x9a, 0xa9,
	0x08, 0x13, 0x11, 0xe4, 0x29, 0xdb, 0x1a, 0x5a, 0xa3, 0xc6, 0xb8, 0xe5, 0x16, 0x7f, 0x70, 0xa7,
	0x3f, 0xc8, 0x3f, 0x93, 0x99, 0x73, 0xaa, 0xe3, 0x79, 0xe9, 0x8a, 0x11, 0x2c, 0x10, 0xe1, 0xe2,
	0xf3, 0xe9, 0xe8, 0x17, 0x80, 0x8f, 0xe0, 0x10, 0xd1, 0x45, 0xa9, 0xdd, 0xfb, 0x5f, 0xdb, 0x93,
	0x99, 0xd3, 0x41, 0x74, 0xb1, 0x43, 0xda, 0x30, 0xca, 0xf0, 0x12, 0x54, 0x57, 0x98, 0x2e, 0xec,
	0xca, 0xd0, 0x1a, 0xed, 0xfb, 0x50, 0x66, 0x4e, 0x53, 0x7d, 0x1b, 0x71, 0xcd, 0xe1, 0x18, 0xd4,
	0x30, 0xe1, 0x31, 0x9e, 0x63, 0x61, 0x57, 0x87, 0xd6, 0xa8, 0xe6, 0x77, 0x65, 0xe6, 0xc0, 0xbc,
	0x66, 0xe4, 0x8b, 0x1c, 0xbc, 0x03, 0x27, 0xcb, 0x30, 0x0d, 0x50, 0x8c, 0x08, 0xa2, 0x22, 0xa0,
	0x21, 0x41, 0xa9, 0xdd, 0xd6, 0x97, 0x07, 0x32, 0x73, 0x7a, 0xcb, 0x30, 0x9d, 0x6c, 0xd9, 0x83,
	0x42, 0x86, 0xe3, 0xf8, 0x0f, 0x82, 0x2f, 0xa0, 0x63, 0xaa, 0x62, 0x36, 0x0f, 0xd5, 0xf3, 0x53,
	0xbb, 0xa3, 0x75, 0xe7, 0x32, 0x73, 0x06, 0xe5, 0x9d, 0xfb, 0x1c, 0x1b, 0xca, 0xd6, 0x0e, 0xec,
	0xbf, 0x83, 0x2e, 0x4b, 0x22, 0xdd, 0x3f, 0xdd, 0xbe, 0xa2, 0x91, 0x7e, 0xb3, 0x18, 0xe7, 0x54,
	0x4d, 0x78, 0x6a, 0xbd, 0x79, 0x11, 0x16, 0xcb, 0xf5, 0xcc, 0x9d, 0x33, 0xe2, 0xa9, 0xdc, 0xb5,
	0x5e, 0x04, 0x7d, 0x44, 0xcc, 0xe3, 0xab, 0xc8, 0x2b, 0xd6, 0xe3, 0x56, 0x1d, 0xb3, 0x03, 0xbd,
	0x1b, 0x37, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x27, 0xf4, 0x25, 0xa1, 0x74, 0x02, 0x00, 0x00,
}
