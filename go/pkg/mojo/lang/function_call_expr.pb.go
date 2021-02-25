// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/lang/function_call_expr.proto

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

type FunctionCallExpr struct {
	StartPosition        *Position      `protobuf:"bytes,1,opt,name=start_position,json=startPosition,proto3" json:"startPosition,omitempty"`
	EndPosition          *Position      `protobuf:"bytes,2,opt,name=end_position,json=endPosition,proto3" json:"endPosition,omitempty"`
	Kind                 int32          `protobuf:"varint,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Implicit             bool           `protobuf:"varint,4,opt,name=implicit,proto3" json:"implicit,omitempty"`
	Callee               *Expression    `protobuf:"bytes,10,opt,name=callee,proto3" json:"callee,omitempty"`
	GenericArguments     []*NominalType `protobuf:"bytes,11,rep,name=generic_arguments,json=genericArguments,proto3" json:"genericArguments,omitempty"`
	Arguments            []*Argument    `protobuf:"bytes,12,rep,name=arguments,proto3" json:"arguments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *FunctionCallExpr) Reset()         { *m = FunctionCallExpr{} }
func (m *FunctionCallExpr) String() string { return proto.CompactTextString(m) }
func (*FunctionCallExpr) ProtoMessage()    {}
func (*FunctionCallExpr) Descriptor() ([]byte, []int) {
	return fileDescriptor_d053d9df3400c9ac, []int{0}
}
func (m *FunctionCallExpr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FunctionCallExpr.Unmarshal(m, b)
}
func (m *FunctionCallExpr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FunctionCallExpr.Marshal(b, m, deterministic)
}
func (m *FunctionCallExpr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FunctionCallExpr.Merge(m, src)
}
func (m *FunctionCallExpr) XXX_Size() int {
	return xxx_messageInfo_FunctionCallExpr.Size(m)
}
func (m *FunctionCallExpr) XXX_DiscardUnknown() {
	xxx_messageInfo_FunctionCallExpr.DiscardUnknown(m)
}

var xxx_messageInfo_FunctionCallExpr proto.InternalMessageInfo

func (m *FunctionCallExpr) GetStartPosition() *Position {
	if m != nil {
		return m.StartPosition
	}
	return nil
}

func (m *FunctionCallExpr) GetEndPosition() *Position {
	if m != nil {
		return m.EndPosition
	}
	return nil
}

func (m *FunctionCallExpr) GetKind() int32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *FunctionCallExpr) GetImplicit() bool {
	if m != nil {
		return m.Implicit
	}
	return false
}

func (m *FunctionCallExpr) GetCallee() *Expression {
	if m != nil {
		return m.Callee
	}
	return nil
}

func (m *FunctionCallExpr) GetGenericArguments() []*NominalType {
	if m != nil {
		return m.GenericArguments
	}
	return nil
}

func (m *FunctionCallExpr) GetArguments() []*Argument {
	if m != nil {
		return m.Arguments
	}
	return nil
}

func init() {
	proto.RegisterType((*FunctionCallExpr)(nil), "mojo.lang.FunctionCallExpr")
}

func init() {
	proto.RegisterFile("mojo/lang/function_call_expr.proto", fileDescriptor_d053d9df3400c9ac)
}

var fileDescriptor_d053d9df3400c9ac = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcb, 0x6a, 0xe3, 0x30,
	0x14, 0x86, 0xf1, 0x24, 0x13, 0x12, 0x39, 0x13, 0x32, 0xca, 0xcd, 0x93, 0x81, 0xda, 0x64, 0x51,
	0xbc, 0x68, 0x6d, 0x48, 0x97, 0x5d, 0xd5, 0xa1, 0x5d, 0x74, 0x51, 0xd2, 0x52, 0x28, 0x74, 0x63,
	0x1c, 0x47, 0x75, 0xd5, 0xd8, 0x92, 0xb0, 0x15, 0x48, 0xde, 0xa2, 0x4f, 0xe8, 0x07, 0xf0, 0x53,
	0x14, 0x39, 0xbe, 0xa8, 0x21, 0xdd, 0x08, 0x7c, 0xce, 0xf7, 0x7f, 0x3a, 0xc8, 0x07, 0xcc, 0x22,
	0xfa, 0x41, 0xed, 0xd0, 0x23, 0x81, 0xfd, 0xb6, 0x25, 0x3e, 0xc7, 0x94, 0xb8, 0xbe, 0x17, 0x86,
	0x2e, 0xda, 0xb1, 0xd8, 0x62, 0x31, 0xe5, 0x14, 0x76, 0x04, 0x63, 0x09, 0x66, 0x3a, 0xac, 0xf1,
	0x1a, 0x90, 0xab, 0xe2, 0x28, 0xaa, 0x5a, 0x5d, 0x65, 0x34, 0xc1, 0x42, 0x7d, 0xe8, 0xcc, 0x3e,
	0x9b, 0xa0, 0x7f, 0x57, 0xdc, 0xb6, 0xf0, 0xc2, 0xf0, 0x76, 0xc7, 0x62, 0xf8, 0x02, 0x7a, 0x09,
	0xf7, 0x62, 0xee, 0x96, 0xb0, 0xa6, 0x18, 0x8a, 0xa9, 0xce, 0x07, 0x56, 0x75, 0xbd, 0xb5, 0x2c,
	0x5a, 0xce, 0xff, 0x2c, 0xd5, 0x27, 0x39, 0x5e, 0x96, 0x2e, 0x68, 0x84, 0x39, 0x8a, 0x18, 0xdf,
	0x3f, 0xfd, 0xf9, 0xd6, 0x80, 0x8f, 0xa0, 0x8b, 0xc8, 0xba, 0xd6, 0xfe, 0xfa, 0x59, 0xfb, 0x2f,
	0x4b, 0xf5, 0x11, 0x22, 0xeb, 0x13, 0x52, 0x55, 0x2a, 0xc3, 0x73, 0xd0, 0xdc, 0x60, 0xb2, 0xd6,
	0x1a, 0x86, 0x62, 0xfe, 0x76, 0x60, 0x96, 0xea, 0x3d, 0xf1, 0x2d, 0xe1, 0x79, 0x1f, 0xce, 0x41,
	0x1b, 0x47, 0x2c, 0xc4, 0x3e, 0xe6, 0x5a, 0xd3, 0x50, 0xcc, 0xb6, 0x33, 0xce, 0x52, 0x1d, 0x96,
	0x35, 0x89, 0xaf, 0x38, 0xb8, 0x00, 0x2d, 0xf1, 0x03, 0x10, 0xd2, 0x40, 0x3e, 0xe8, 0x48, 0x1a,
	0x54, 0x3c, 0x14, 0x4a, 0x12, 0x31, 0xea, 0x30, 0x4b, 0xf5, 0xfe, 0x01, 0x94, 0x34, 0x45, 0x14,
	0xfa, 0xe0, 0x6f, 0x80, 0x08, 0x8a, 0xb1, 0xef, 0x7a, 0x71, 0xb0, 0x8d, 0x10, 0xe1, 0x89, 0xa6,
	0x1a, 0x0d, 0x53, 0x9d, 0x8f, 0x25, 0xdf, 0x03, 0x8d, 0x30, 0xf1, 0xc2, 0xe7, 0x3d, 0x43, 0xce,
	0x59, 0x96, 0xea, 0xd3, 0x22, 0x74, 0x53, 0x66, 0x24, 0x75, 0xff, 0xb8, 0x07, 0xef, 0x41, 0xa7,
	0x96, 0x77, 0x73, 0xb9, 0xfc, 0xaa, 0x25, 0xe8, 0x4c, 0xb2, 0x54, 0x1f, 0x78, 0x27, 0x94, 0x75,
	0xdc, 0x59, 0x81, 0x31, 0x8d, 0x83, 0x3c, 0x9d, 0x87, 0x2b, 0x8d, 0x33, 0x3a, 0xde, 0x94, 0xa5,
	0xd8, 0xa1, 0xa5, 0xf2, 0x6a, 0x07, 0x98, 0xbf, 0x6f, 0x57, 0x96, 0x4f, 0x23, 0x5b, 0xe0, 0x97,
	0xd5, 0x02, 0xda, 0x01, 0xb5, 0xd9, 0x26, 0xb0, 0xab, 0x05, 0xbc, 0x16, 0xc7, 0xaa, 0x95, 0x6f,
	0xdf, 0xd5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5f, 0xbf, 0x0a, 0xdc, 0xf4, 0x02, 0x00, 0x00,
}
