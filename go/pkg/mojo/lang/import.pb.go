// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/lang/import.proto

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

type Import struct {
	PackageName          string        `protobuf:"bytes,10,opt,name=package_name,json=packageName,proto3" json:"packageName,omitempty"`
	Identifiers          []*Identifier `protobuf:"bytes,13,rep,name=identifiers,proto3" json:"identifiers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Import) Reset()         { *m = Import{} }
func (m *Import) String() string { return proto.CompactTextString(m) }
func (*Import) ProtoMessage()    {}
func (*Import) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0a82053035e04b1, []int{0}
}
func (m *Import) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Import.Unmarshal(m, b)
}
func (m *Import) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Import.Marshal(b, m, deterministic)
}
func (m *Import) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Import.Merge(m, src)
}
func (m *Import) XXX_Size() int {
	return xxx_messageInfo_Import.Size(m)
}
func (m *Import) XXX_DiscardUnknown() {
	xxx_messageInfo_Import.DiscardUnknown(m)
}

var xxx_messageInfo_Import proto.InternalMessageInfo

func (m *Import) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *Import) GetIdentifiers() []*Identifier {
	if m != nil {
		return m.Identifiers
	}
	return nil
}

func init() {
	proto.RegisterType((*Import)(nil), "mojo.lang.Import")
}

func init() { proto.RegisterFile("mojo/lang/import.proto", fileDescriptor_c0a82053035e04b1) }

var fileDescriptor_c0a82053035e04b1 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcb, 0xcd, 0xcf, 0xca,
	0xd7, 0xcf, 0x49, 0xcc, 0x4b, 0xd7, 0xcf, 0xcc, 0x2d, 0xc8, 0x2f, 0x2a, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x89, 0xeb, 0x81, 0xc4, 0xa5, 0x44, 0x10, 0x4a, 0x40, 0x04, 0x44,
	0x81, 0xd2, 0x2c, 0x46, 0x2e, 0x36, 0x4f, 0xb0, 0x0e, 0x21, 0x1b, 0x2e, 0x9e, 0x82, 0xc4, 0xe4,
	0xec, 0xc4, 0xf4, 0xd4, 0xf8, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x2e, 0x05, 0x46, 0x0d, 0x4e, 0x27,
	0xc9, 0x57, 0xf7, 0xe4, 0x45, 0xa1, 0xe2, 0x7e, 0x89, 0xb9, 0xa9, 0x3a, 0xf9, 0xb9, 0x99, 0x25,
	0xa9, 0xb9, 0x05, 0x25, 0x95, 0x41, 0xdc, 0x48, 0xc2, 0x42, 0x41, 0x5c, 0xdc, 0x99, 0x29, 0xa9,
	0x79, 0x25, 0x99, 0x69, 0x99, 0xa9, 0x45, 0xc5, 0x12, 0xbc, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xa2,
	0x7a, 0x70, 0xfb, 0xf5, 0x3c, 0xe1, 0xb2, 0x10, 0x33, 0x91, 0x54, 0x23, 0x9b, 0x89, 0x24, 0xec,
	0x14, 0xc1, 0x25, 0x96, 0x5f, 0x94, 0x0e, 0x36, 0x03, 0x6c, 0x04, 0xdc, 0x30, 0x27, 0x6e, 0x88,
	0x9b, 0x03, 0x40, 0x7e, 0x08, 0x60, 0x8c, 0xd2, 0x4f, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0x07, 0x29, 0xd2, 0x85, 0x7b, 0x53, 0x3f, 0x3d, 0x5f, 0xbf, 0x20, 0x3b, 0x5d,
	0x1f, 0xee, 0x79, 0x6b, 0x10, 0x91, 0xc4, 0x06, 0xf6, 0xbd, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0xe6, 0xc2, 0x4b, 0xfc, 0x38, 0x01, 0x00, 0x00,
}
