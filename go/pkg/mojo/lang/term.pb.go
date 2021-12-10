// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mojo/lang/term.proto

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

type Term struct {
	Type                 string        `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value                string        `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Location             *PositionSpan `protobuf:"bytes,10,opt,name=location,proto3" json:"location,omitempty"`
	Terms                []*Term       `protobuf:"bytes,15,rep,name=terms,proto3" json:"terms,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-" gorm:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte        `json:"-" gorm:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32         `json:"-" gorm:"-" xml:"-" bson:"-"`
}

func (m *Term) Reset()         { *m = Term{} }
func (m *Term) String() string { return proto.CompactTextString(m) }
func (*Term) ProtoMessage()    {}
func (*Term) Descriptor() ([]byte, []int) {
	return fileDescriptor_978bceccd11ad85f, []int{0}
}
func (m *Term) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Term.Unmarshal(m, b)
}
func (m *Term) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Term.Marshal(b, m, deterministic)
}
func (m *Term) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Term.Merge(m, src)
}
func (m *Term) XXX_Size() int {
	return xxx_messageInfo_Term.Size(m)
}
func (m *Term) XXX_DiscardUnknown() {
	xxx_messageInfo_Term.DiscardUnknown(m)
}

var xxx_messageInfo_Term proto.InternalMessageInfo

func (m *Term) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Term) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Term) GetLocation() *PositionSpan {
	if m != nil {
		return m.Location
	}
	return nil
}

func (m *Term) GetTerms() []*Term {
	if m != nil {
		return m.Terms
	}
	return nil
}

func init() {
	proto.RegisterType((*Term)(nil), "mojo.lang.Term")
}

func init() { proto.RegisterFile("mojo/lang/term.proto", fileDescriptor_978bceccd11ad85f) }

var fileDescriptor_978bceccd11ad85f = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xad, 0xee, 0x8a, 0xcd, 0x82, 0x85, 0x28, 0x6b, 0xf0, 0x60, 0x8b, 0x07, 0xa9, 0xa0,
	0x09, 0xac, 0x37, 0xbd, 0xf5, 0xea, 0xa5, 0xa8, 0xa7, 0xbd, 0x75, 0x97, 0x10, 0xab, 0x4d, 0x27,
	0x74, 0xb3, 0xc2, 0xfe, 0x13, 0x7f, 0x5d, 0x0f, 0x7a, 0xeb, 0xaf, 0x90, 0x99, 0x6a, 0xed, 0x65,
	0x20, 0x2f, 0xdf, 0x9b, 0xbc, 0x3c, 0x76, 0x6a, 0xe1, 0x0d, 0x54, 0x55, 0xd4, 0x46, 0x79, 0xdd,
	0x58, 0xe9, 0x1a, 0xf0, 0xc0, 0x43, 0x54, 0x25, 0xaa, 0xe7, 0xe2, 0x1f, 0x70, 0xb0, 0x29, 0x7d,
	0x09, 0x75, 0x0f, 0x5d, 0x7e, 0x05, 0x6c, 0xf2, 0xa2, 0x1b, 0xcb, 0xaf, 0xd8, 0xc4, 0xef, 0x9c,
	0x16, 0x41, 0x12, 0xa4, 0x61, 0xc6, 0xbb, 0x36, 0x3e, 0xc6, 0xf3, 0x0d, 0xd8, 0xd2, 0x6b, 0xeb,
	0xfc, 0xee, 0x89, 0xee, 0xf9, 0x35, 0x9b, 0x7e, 0x14, 0xd5, 0x56, 0x8b, 0x7d, 0x02, 0x4f, 0xba,
	0x36, 0x8e, 0x48, 0x18, 0x91, 0x3d, 0xc1, 0x1f, 0xd9, 0x51, 0x05, 0xeb, 0x02, 0x5f, 0x13, 0x2c,
	0x09, 0xd2, 0xd9, 0xe2, 0x4c, 0x0e, 0x99, 0x64, 0xfe, 0x1b, 0xe4, 0xd9, 0x15, 0x75, 0x36, 0xef,
	0xda, 0x98, 0xff, 0xc1, 0xa3, 0x4d, 0xc3, 0x02, 0x7e, 0xcf, 0xa6, 0xf8, 0xb7, 0x8d, 0x88, 0x92,
	0x83, 0x74, 0xb6, 0x88, 0x46, 0x9b, 0x30, 0x7f, 0x1f, 0x84, 0x88, 0x71, 0x10, 0x12, 0xb2, 0xe5,
	0xe7, 0xf7, 0xc5, 0x1e, 0x9b, 0x43, 0x63, 0xc8, 0x45, 0xa6, 0xc1, 0x9e, 0x85, 0xe8, 0xcf, 0xb1,
	0x8d, 0x3c, 0x58, 0x2a, 0x53, 0xfa, 0xd7, 0xed, 0x4a, 0xae, 0xc1, 0x2a, 0x44, 0x6e, 0xa9, 0x34,
	0x1a, 0x06, 0x94, 0x7b, 0x37, 0x6a, 0xa8, 0xf2, 0x01, 0xc7, 0xea, 0x90, 0x7a, 0xbc, 0xfb, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x50, 0x64, 0xb1, 0xf3, 0x84, 0x01, 0x00, 0x00,
}
