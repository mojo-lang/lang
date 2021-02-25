// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/lang/block_stmt.proto

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

type BlockStmt struct {
	StartPosition        *Position `protobuf:"bytes,1,opt,name=start_position,json=startPosition,proto3" json:"startPosition,omitempty"`
	EndPosition          *Position `protobuf:"bytes,2,opt,name=end_position,json=endPosition,proto3" json:"endPosition,omitempty"`
	Kind                 int64     `protobuf:"varint,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Implicit             bool      `protobuf:"varint,4,opt,name=implicit,proto3" json:"implicit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *BlockStmt) Reset()         { *m = BlockStmt{} }
func (m *BlockStmt) String() string { return proto.CompactTextString(m) }
func (*BlockStmt) ProtoMessage()    {}
func (*BlockStmt) Descriptor() ([]byte, []int) {
	return fileDescriptor_54277e2dd8ec7066, []int{0}
}
func (m *BlockStmt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockStmt.Unmarshal(m, b)
}
func (m *BlockStmt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockStmt.Marshal(b, m, deterministic)
}
func (m *BlockStmt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockStmt.Merge(m, src)
}
func (m *BlockStmt) XXX_Size() int {
	return xxx_messageInfo_BlockStmt.Size(m)
}
func (m *BlockStmt) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockStmt.DiscardUnknown(m)
}

var xxx_messageInfo_BlockStmt proto.InternalMessageInfo

func (m *BlockStmt) GetStartPosition() *Position {
	if m != nil {
		return m.StartPosition
	}
	return nil
}

func (m *BlockStmt) GetEndPosition() *Position {
	if m != nil {
		return m.EndPosition
	}
	return nil
}

func (m *BlockStmt) GetKind() int64 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *BlockStmt) GetImplicit() bool {
	if m != nil {
		return m.Implicit
	}
	return false
}

func init() {
	proto.RegisterType((*BlockStmt)(nil), "mojo.lang.BlockStmt")
}

func init() { proto.RegisterFile("mojo/lang/block_stmt.proto", fileDescriptor_54277e2dd8ec7066) }

var fileDescriptor_54277e2dd8ec7066 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4e, 0x84, 0x30,
	0x14, 0x85, 0xc3, 0xcc, 0xc4, 0xcc, 0x74, 0x94, 0x45, 0xd5, 0x11, 0x71, 0x01, 0x71, 0x61, 0x58,
	0x28, 0x4d, 0xc6, 0xa5, 0xbb, 0x3e, 0x01, 0xea, 0xc2, 0x44, 0x17, 0x93, 0xe1, 0x27, 0x58, 0xa1,
	0x5c, 0x02, 0xd7, 0x85, 0xcf, 0xe1, 0xfb, 0xf1, 0x00, 0x3c, 0x85, 0x69, 0x95, 0x82, 0x89, 0x6e,
	0x6e, 0xd2, 0x7b, 0xbe, 0x7b, 0x7a, 0x72, 0x88, 0x2b, 0xe1, 0x0d, 0x58, 0xb9, 0xaf, 0x72, 0x16,
	0x97, 0x90, 0x14, 0xbb, 0x16, 0x25, 0x86, 0x75, 0x03, 0x08, 0x74, 0xa5, 0xb4, 0x50, 0x69, 0xae,
	0x33, 0x62, 0x35, 0xb4, 0x02, 0x05, 0x54, 0xdf, 0x90, 0x7b, 0x32, 0x2a, 0xe3, 0xe9, 0xe5, 0xe7,
	0x8c, 0xac, 0xb8, 0xf2, 0x7b, 0x44, 0x89, 0xf4, 0x89, 0xd8, 0x2d, 0xee, 0x1b, 0xdc, 0x0d, 0xb7,
	0x8e, 0xe5, 0x5b, 0xc1, 0x7a, 0x7b, 0x1c, 0x9a, 0x1f, 0xc2, 0xe8, 0x47, 0xe2, 0x17, 0x7d, 0xe7,
	0x9d, 0x69, 0x7c, 0x58, 0x5d, 0x83, 0x14, 0x98, 0xc9, 0x1a, 0x3f, 0x1e, 0x8e, 0x7e, 0x09, 0xf4,
	0x9e, 0x1c, 0x66, 0x55, 0x3a, 0xda, 0xce, 0xfe, 0xb7, 0x3d, 0xef, 0x3b, 0xef, 0x34, 0xab, 0xd2,
	0x3f, 0x4c, 0xd7, 0x93, 0x35, 0xbd, 0x22, 0x8b, 0x42, 0x54, 0xa9, 0x33, 0xf7, 0xad, 0x60, 0xce,
	0x69, 0xdf, 0x79, 0xb6, 0x7a, 0x4f, 0x70, 0xad, 0xd3, 0x2d, 0x59, 0x0a, 0x59, 0x97, 0x22, 0x11,
	0xe8, 0x2c, 0x7c, 0x2b, 0x58, 0xf2, 0x4d, 0xdf, 0x79, 0x74, 0xd8, 0x4d, 0x78, 0xc3, 0xf1, 0x17,
	0xb2, 0x81, 0x26, 0xd7, 0xe9, 0x74, 0x38, 0x13, 0x93, 0xdb, 0xa6, 0xac, 0x48, 0xf5, 0x17, 0x59,
	0xcf, 0x2c, 0x17, 0xf8, 0xfa, 0x1e, 0x87, 0x09, 0x48, 0xa6, 0xb8, 0x1b, 0x5d, 0xb1, 0x1e, 0x39,
	0xb0, 0xba, 0xc8, 0x99, 0x29, 0xfe, 0x4e, 0x8d, 0xf8, 0x40, 0x37, 0x7f, 0xfb, 0x15, 0x00, 0x00,
	0xff, 0xff, 0x6d, 0x26, 0x7c, 0xcd, 0xd2, 0x01, 0x00, 0x00,
}
