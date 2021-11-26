// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/lang/document.proto

package lang

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	document "github.com/mojo-lang/document/go/pkg/mojo/document"
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

type Document struct {
	StartPosition        *Position          `protobuf:"bytes,1,opt,name=start_position,json=startPosition,proto3" json:"startPosition,omitempty"`
	EndPosition          *Position          `protobuf:"bytes,2,opt,name=end_position,json=endPosition,proto3" json:"endPosition,omitempty"`
	Lines                []*LineDocument    `protobuf:"bytes,10,rep,name=lines,proto3" json:"lines,omitempty"`
	Structured           *document.Document `protobuf:"bytes,15,opt,name=structured,proto3" json:"structured,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Document) Reset()         { *m = Document{} }
func (m *Document) String() string { return proto.CompactTextString(m) }
func (*Document) ProtoMessage()    {}
func (*Document) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4a6b46431335f62, []int{0}
}
func (m *Document) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Document.Unmarshal(m, b)
}
func (m *Document) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Document.Marshal(b, m, deterministic)
}
func (m *Document) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Document.Merge(m, src)
}
func (m *Document) XXX_Size() int {
	return xxx_messageInfo_Document.Size(m)
}
func (m *Document) XXX_DiscardUnknown() {
	xxx_messageInfo_Document.DiscardUnknown(m)
}

var xxx_messageInfo_Document proto.InternalMessageInfo

func (m *Document) GetStartPosition() *Position {
	if m != nil {
		return m.StartPosition
	}
	return nil
}

func (m *Document) GetEndPosition() *Position {
	if m != nil {
		return m.EndPosition
	}
	return nil
}

func (m *Document) GetLines() []*LineDocument {
	if m != nil {
		return m.Lines
	}
	return nil
}

func (m *Document) GetStructured() *document.Document {
	if m != nil {
		return m.Structured
	}
	return nil
}

type LineDocument struct {
	StartPosition        *Position `protobuf:"bytes,1,opt,name=start_position,json=startPosition,proto3" json:"startPosition,omitempty"`
	EndPosition          *Position `protobuf:"bytes,2,opt,name=end_position,json=endPosition,proto3" json:"endPosition,omitempty"`
	Following            bool      `protobuf:"varint,3,opt,name=following,proto3" json:"following,omitempty"`
	Line                 string    `protobuf:"bytes,10,opt,name=line,proto3" json:"line,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LineDocument) Reset()         { *m = LineDocument{} }
func (m *LineDocument) String() string { return proto.CompactTextString(m) }
func (*LineDocument) ProtoMessage()    {}
func (*LineDocument) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4a6b46431335f62, []int{1}
}
func (m *LineDocument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LineDocument.Unmarshal(m, b)
}
func (m *LineDocument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LineDocument.Marshal(b, m, deterministic)
}
func (m *LineDocument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LineDocument.Merge(m, src)
}
func (m *LineDocument) XXX_Size() int {
	return xxx_messageInfo_LineDocument.Size(m)
}
func (m *LineDocument) XXX_DiscardUnknown() {
	xxx_messageInfo_LineDocument.DiscardUnknown(m)
}

var xxx_messageInfo_LineDocument proto.InternalMessageInfo

func (m *LineDocument) GetStartPosition() *Position {
	if m != nil {
		return m.StartPosition
	}
	return nil
}

func (m *LineDocument) GetEndPosition() *Position {
	if m != nil {
		return m.EndPosition
	}
	return nil
}

func (m *LineDocument) GetFollowing() bool {
	if m != nil {
		return m.Following
	}
	return false
}

func (m *LineDocument) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func init() {
	proto.RegisterType((*Document)(nil), "mojo.lang.Document")
	proto.RegisterType((*LineDocument)(nil), "mojo.lang.LineDocument")
}

func init() { proto.RegisterFile("mojo/lang/document.proto", fileDescriptor_d4a6b46431335f62) }

var fileDescriptor_d4a6b46431335f62 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x92, 0x4f, 0x4a, 0xc3, 0x40,
	0x14, 0x87, 0x49, 0xab, 0xd2, 0x4e, 0xff, 0xc1, 0x54, 0xed, 0x58, 0x85, 0x84, 0x2e, 0x24, 0x0b,
	0x4d, 0xa0, 0xe2, 0xca, 0xdd, 0xe0, 0xd2, 0x45, 0xeb, 0x46, 0xe8, 0x46, 0xda, 0x66, 0x8c, 0xa3,
	0xc9, 0xbc, 0x90, 0x4c, 0x10, 0x8f, 0xe2, 0xa1, 0xbc, 0x42, 0x0e, 0x90, 0x53, 0xc8, 0x4c, 0x9b,
	0x74, 0x14, 0x3d, 0x80, 0x9b, 0x81, 0xbc, 0xdf, 0x37, 0xdf, 0xbc, 0x3c, 0x1e, 0x22, 0x31, 0xbc,
	0x80, 0x1f, 0x2d, 0x45, 0xe8, 0x07, 0xb0, 0xce, 0x63, 0x26, 0xa4, 0x97, 0xa4, 0x20, 0x01, 0xb7,
	0x55, 0xe2, 0xa9, 0x64, 0x7c, 0xa6, 0xa1, 0x2a, 0xff, 0x01, 0x8e, 0x0d, 0x45, 0x02, 0x19, 0x97,
	0x1c, 0xc4, 0x26, 0x99, 0x7c, 0x36, 0x50, 0xeb, 0x76, 0x0b, 0xe3, 0x07, 0xd4, 0xcf, 0xe4, 0x32,
	0x95, 0x8f, 0x15, 0x44, 0x2c, 0xc7, 0x72, 0x3b, 0xd3, 0xa1, 0x57, 0x3f, 0xe4, 0xcd, 0xb6, 0x11,
	0x3d, 0x2d, 0x0b, 0x7b, 0xa4, 0xf1, 0xaa, 0x74, 0x01, 0x31, 0x97, 0x2c, 0x4e, 0xe4, 0xfb, 0x7d,
	0xef, 0x5b, 0x80, 0xe7, 0xa8, 0xcb, 0x44, 0xb0, 0xd3, 0x36, 0xfe, 0xd6, 0x9e, 0x94, 0x85, 0x7d,
	0xc4, 0x44, 0xf0, 0x8b, 0xb4, 0x63, 0x94, 0x31, 0x45, 0xfb, 0x11, 0x17, 0x2c, 0x23, 0xc8, 0x69,
	0xba, 0x9d, 0xe9, 0xc8, 0x70, 0xdd, 0x71, 0xc1, 0xaa, 0x7f, 0xa2, 0xc3, 0xb2, 0xb0, 0x07, 0x9a,
	0x34, 0x4c, 0x9b, 0xab, 0x78, 0x8e, 0x50, 0x26, 0xd3, 0x7c, 0x2d, 0xf3, 0x94, 0x05, 0x64, 0xa0,
	0x9b, 0xda, 0x8a, 0xea, 0x01, 0xd6, 0x22, 0x52, 0x16, 0xf6, 0xe1, 0x0e, 0x37, 0x6c, 0x86, 0x64,
	0xf2, 0xd1, 0x40, 0x5d, 0xf3, 0xfd, 0x7f, 0x35, 0xd3, 0x6b, 0xd4, 0x7e, 0x82, 0x28, 0x82, 0x37,
	0x2e, 0x42, 0xd2, 0x74, 0x2c, 0xb7, 0x45, 0x47, 0x65, 0x61, 0x0f, 0xeb, 0xa2, 0x71, 0x71, 0x47,
	0xe2, 0x73, 0xb4, 0xa7, 0xe6, 0x49, 0x90, 0x63, 0xb9, 0x6d, 0x8a, 0xcb, 0xc2, 0xee, 0xab, 0x6f,
	0x03, 0xd6, 0x39, 0x5d, 0xa0, 0x63, 0x48, 0x43, 0xdd, 0xa0, 0xee, 0xaf, 0xee, 0x94, 0xf6, 0xaa,
	0x71, 0xcd, 0xd4, 0x52, 0xce, 0xac, 0x85, 0x1f, 0x72, 0xf9, 0x9c, 0xaf, 0xbc, 0x35, 0xc4, 0xbe,
	0xc2, 0x2e, 0xf5, 0xee, 0xea, 0x23, 0x04, 0x3f, 0x79, 0x0d, 0xfd, 0x7a, 0xa3, 0x6f, 0xd4, 0xb1,
	0x3a, 0xd0, 0xeb, 0x7c, 0xf5, 0x15, 0x00, 0x00, 0xff, 0xff, 0x53, 0xed, 0x69, 0xab, 0x2d, 0x03,
	0x00, 0x00,
}