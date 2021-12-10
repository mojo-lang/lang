// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/lang/branch_stmt.proto

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

type BranchStmt struct {
	StartPosition        *Position `protobuf:"bytes,1,opt,name=start_position,json=startPosition,proto3" json:"startPosition,omitempty"`
	EndPosition          *Position `protobuf:"bytes,2,opt,name=end_position,json=endPosition,proto3" json:"endPosition,omitempty"`
	Kind                 int64     `protobuf:"varint,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Implicit             bool      `protobuf:"varint,4,opt,name=implicit,proto3" json:"implicit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-" gorm:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte    `json:"-" gorm:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32     `json:"-" gorm:"-" xml:"-" bson:"-"`
}

func (m *BranchStmt) Reset()         { *m = BranchStmt{} }
func (m *BranchStmt) String() string { return proto.CompactTextString(m) }
func (*BranchStmt) ProtoMessage()    {}
func (*BranchStmt) Descriptor() ([]byte, []int) {
	return fileDescriptor_16b0182d2b1ffeb2, []int{0}
}
func (m *BranchStmt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BranchStmt.Unmarshal(m, b)
}
func (m *BranchStmt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BranchStmt.Marshal(b, m, deterministic)
}
func (m *BranchStmt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BranchStmt.Merge(m, src)
}
func (m *BranchStmt) XXX_Size() int {
	return xxx_messageInfo_BranchStmt.Size(m)
}
func (m *BranchStmt) XXX_DiscardUnknown() {
	xxx_messageInfo_BranchStmt.DiscardUnknown(m)
}

var xxx_messageInfo_BranchStmt proto.InternalMessageInfo

func (m *BranchStmt) GetStartPosition() *Position {
	if m != nil {
		return m.StartPosition
	}
	return nil
}

func (m *BranchStmt) GetEndPosition() *Position {
	if m != nil {
		return m.EndPosition
	}
	return nil
}

func (m *BranchStmt) GetKind() int64 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *BranchStmt) GetImplicit() bool {
	if m != nil {
		return m.Implicit
	}
	return false
}

func init() {
	proto.RegisterType((*BranchStmt)(nil), "mojo.lang.BranchStmt")
}

func init() { proto.RegisterFile("mojo/lang/branch_stmt.proto", fileDescriptor_16b0182d2b1ffeb2) }

var fileDescriptor_16b0182d2b1ffeb2 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x49, 0x5b, 0xa1, 0xe2, 0x42, 0x91, 0x0c, 0x94, 0xd0, 0x4a, 0x24, 0x62, 0x40, 0x19,
	0x20, 0x96, 0xca, 0xc8, 0xe6, 0x27, 0x08, 0x30, 0x20, 0xb1, 0x94, 0xfc, 0x29, 0x35, 0xad, 0x7d,
	0xa3, 0xe4, 0x32, 0xf0, 0x1e, 0x0c, 0x3c, 0x5d, 0x16, 0xb6, 0x3c, 0x05, 0xb2, 0x21, 0x3f, 0x48,
	0xb0, 0x5c, 0xc9, 0xf7, 0x7c, 0xf7, 0xf8, 0xe8, 0x90, 0x85, 0x84, 0x17, 0x60, 0xdb, 0x50, 0x65,
	0x2c, 0x2a, 0x42, 0x15, 0xaf, 0x57, 0x25, 0x4a, 0xf4, 0xf3, 0x02, 0x10, 0xe8, 0x9e, 0x16, 0x7d,
	0x2d, 0xce, 0xed, 0x8e, 0xcb, 0xa1, 0x14, 0x28, 0x40, 0x7d, 0x43, 0xf3, 0xe3, 0x4e, 0xe9, 0x4e,
	0x2f, 0xde, 0x07, 0x84, 0x70, 0x63, 0xf8, 0x80, 0x12, 0xe9, 0x23, 0x99, 0x96, 0x18, 0x16, 0xb8,
	0x6a, 0x8e, 0x6d, 0xcb, 0xb5, 0xbc, 0xc9, 0xf2, 0xc8, 0x6f, 0xbf, 0xf0, 0x83, 0x1f, 0x89, 0x2f,
	0xea, 0xca, 0x39, 0x35, 0x78, 0xb3, 0xba, 0x02, 0x29, 0x30, 0x95, 0x39, 0xbe, 0xdd, 0x1f, 0xfc,
	0x12, 0xe8, 0x1d, 0xd9, 0x4f, 0x55, 0xd2, 0xd9, 0x0e, 0xfe, 0xb7, 0x3d, 0xab, 0x2b, 0xe7, 0x24,
	0x55, 0xc9, 0x1f, 0xa6, 0x93, 0xde, 0x9a, 0x5e, 0x92, 0xd1, 0x46, 0xa8, 0xc4, 0x1e, 0xba, 0x96,
	0x37, 0xe4, 0xb4, 0xae, 0x9c, 0xa9, 0x7e, 0xf7, 0x70, 0xa3, 0xd3, 0x25, 0x19, 0x0b, 0x99, 0x6f,
	0x45, 0x2c, 0xd0, 0x1e, 0xb9, 0x96, 0x37, 0xe6, 0xb3, 0xba, 0x72, 0x68, 0xb3, 0xeb, 0xf1, 0x2d,
	0xc7, 0x9f, 0x3f, 0x3e, 0xcf, 0x77, 0xc8, 0x0c, 0x8a, 0xcc, 0x24, 0x34, 0x01, 0xdb, 0xa8, 0xfc,
	0xb0, 0x6b, 0x2c, 0xd0, 0x2d, 0x06, 0xd6, 0x13, 0xcb, 0x04, 0xae, 0x5f, 0x23, 0x3f, 0x06, 0xc9,
	0x34, 0x78, 0x6d, 0x8a, 0x36, 0x23, 0x03, 0x96, 0x6f, 0x32, 0xd6, 0xd6, 0x7f, 0xab, 0x47, 0xb4,
	0x6b, 0xfa, 0xbf, 0xf9, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x28, 0x9f, 0x9d, 0xd9, 0x01, 0x00,
	0x00,
}
