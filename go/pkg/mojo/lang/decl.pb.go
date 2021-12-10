// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/lang/decl.proto

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

type Decl struct {
	StartPosition        *Position `protobuf:"bytes,1,opt,name=start_position,json=startPosition,proto3" json:"startPosition,omitempty"`
	EndPosition          *Position `protobuf:"bytes,2,opt,name=end_position,json=endPosition,proto3" json:"endPosition,omitempty"`
	Document             *Document `protobuf:"bytes,4,opt,name=document,proto3" json:"document,omitempty"`
	PackageName          string    `protobuf:"bytes,5,opt,name=package_name,json=packageName,proto3" json:"packageName,omitempty"`
	SourceFileName       string    `protobuf:"bytes,6,opt,name=source_file_name,json=sourceFileName,proto3" json:"sourceFileName,omitempty"`
	Implicit             bool      `protobuf:"varint,7,opt,name=implicit,proto3" json:"implicit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-" gorm:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte    `json:"-" gorm:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32     `json:"-" gorm:"-" xml:"-" bson:"-"`
}

func (m *Decl) Reset()         { *m = Decl{} }
func (m *Decl) String() string { return proto.CompactTextString(m) }
func (*Decl) ProtoMessage()    {}
func (*Decl) Descriptor() ([]byte, []int) {
	return fileDescriptor_e62b6616835383ad, []int{0}
}
func (m *Decl) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Decl.Unmarshal(m, b)
}
func (m *Decl) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Decl.Marshal(b, m, deterministic)
}
func (m *Decl) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Decl.Merge(m, src)
}
func (m *Decl) XXX_Size() int {
	return xxx_messageInfo_Decl.Size(m)
}
func (m *Decl) XXX_DiscardUnknown() {
	xxx_messageInfo_Decl.DiscardUnknown(m)
}

var xxx_messageInfo_Decl proto.InternalMessageInfo

func (m *Decl) GetStartPosition() *Position {
	if m != nil {
		return m.StartPosition
	}
	return nil
}

func (m *Decl) GetEndPosition() *Position {
	if m != nil {
		return m.EndPosition
	}
	return nil
}

func (m *Decl) GetDocument() *Document {
	if m != nil {
		return m.Document
	}
	return nil
}

func (m *Decl) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *Decl) GetSourceFileName() string {
	if m != nil {
		return m.SourceFileName
	}
	return ""
}

func (m *Decl) GetImplicit() bool {
	if m != nil {
		return m.Implicit
	}
	return false
}

func init() {
	proto.RegisterType((*Decl)(nil), "mojo.lang.Decl")
}

func init() { proto.RegisterFile("mojo/lang/decl.proto", fileDescriptor_e62b6616835383ad) }

var fileDescriptor_e62b6616835383ad = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x09, 0x94, 0xd2, 0xba, 0xa5, 0x42, 0x06, 0x4a, 0xf8, 0x11, 0xa9, 0x98, 0x3a, 0x40,
	0x22, 0x95, 0x11, 0xa6, 0xa8, 0x2a, 0x1b, 0x2a, 0x2c, 0x48, 0x5d, 0xaa, 0xd4, 0x35, 0xc1, 0x34,
	0xce, 0x8d, 0x52, 0x77, 0xe0, 0x4d, 0x78, 0x11, 0x5e, 0x27, 0x0b, 0x5b, 0x9e, 0x02, 0xd9, 0x89,
	0x83, 0xcb, 0xcf, 0xe2, 0xe1, 0x9c, 0x93, 0x2f, 0xe7, 0xda, 0x17, 0x1d, 0x70, 0x78, 0x05, 0x2f,
	0x0a, 0xe2, 0xd0, 0x9b, 0x53, 0x12, 0xb9, 0x49, 0x0a, 0x02, 0x70, 0x53, 0xaa, 0xae, 0x54, 0x4f,
	0x6c, 0x23, 0x00, 0x64, 0xc5, 0x69, 0x2c, 0x8a, 0x90, 0xe9, 0x24, 0xb0, 0x64, 0x82, 0x41, 0x5c,
	0x38, 0x17, 0x1f, 0x5b, 0xa8, 0x36, 0xa4, 0x24, 0xc2, 0x4f, 0xa8, 0xb3, 0x14, 0x41, 0x2a, 0xa6,
	0x3a, 0x60, 0x5b, 0x3d, 0xab, 0xdf, 0x1a, 0xec, 0xbb, 0xd5, 0x0f, 0xdc, 0x71, 0x69, 0xf9, 0xa7,
	0x79, 0xe6, 0x1c, 0xa9, 0xb8, 0x96, 0x2e, 0x81, 0x33, 0x41, 0x79, 0x22, 0xde, 0x1e, 0x77, 0xd7,
	0x0c, 0xfc, 0x80, 0xda, 0x34, 0x9e, 0x7f, 0x63, 0x37, 0xff, 0xc7, 0x1e, 0xe7, 0x99, 0x73, 0x48,
	0xe3, 0xf9, 0x1f, 0xd0, 0x96, 0x21, 0xe3, 0x3b, 0xd4, 0xd0, 0x03, 0xda, 0xb5, 0x5f, 0xb8, 0x61,
	0x69, 0xf9, 0xdd, 0x3c, 0x73, 0xb0, 0x0e, 0x1a, 0xac, 0xea, 0x63, 0x7c, 0x8b, 0xda, 0x49, 0x40,
	0x16, 0x41, 0x48, 0xa7, 0x71, 0xc0, 0xa9, 0xbd, 0xdd, 0xb3, 0xfa, 0xcd, 0xa2, 0x46, 0xa9, 0xdf,
	0x07, 0x9c, 0x9a, 0x35, 0x0c, 0x19, 0x8f, 0xd0, 0xde, 0x12, 0x56, 0x29, 0xa1, 0xd3, 0x67, 0x16,
	0x95, 0x84, 0xba, 0x22, 0x9c, 0xe5, 0x99, 0x63, 0x17, 0xde, 0x88, 0x45, 0x3f, 0x21, 0x9d, 0x75,
	0x07, 0x0f, 0x50, 0x83, 0xf1, 0x24, 0x62, 0x84, 0x09, 0x7b, 0xa7, 0x67, 0xf5, 0x1b, 0x45, 0x73,
	0xad, 0x99, 0xcd, 0xb5, 0xe6, 0x4f, 0xde, 0x3f, 0xcf, 0x37, 0x50, 0x17, 0xd2, 0x50, 0x4d, 0xae,
	0x06, 0xaf, 0xae, 0xc0, 0x6f, 0xca, 0x27, 0x1d, 0xcb, 0x07, 0x1e, 0x5b, 0x13, 0x2f, 0x64, 0xe2,
	0x65, 0x35, 0x73, 0x09, 0x70, 0x4f, 0x46, 0xae, 0xd4, 0x1e, 0xa8, 0x23, 0x04, 0x2f, 0x59, 0x84,
	0x5e, 0xb5, 0x1d, 0x37, 0xf2, 0x98, 0xd5, 0xd5, 0x6a, 0x5c, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x34, 0xaa, 0xf4, 0xc3, 0x71, 0x02, 0x00, 0x00,
}
