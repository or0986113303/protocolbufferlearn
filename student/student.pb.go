// Code generated by protoc-gen-go. DO NOT EDIT.
// source: student.proto

package student

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

type Student struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  string   `protobuf:"bytes,2,opt,name=age,proto3" json:"age,omitempty"`
	Gender               string   `protobuf:"bytes,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Number               int32    `protobuf:"varint,4,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Student) Reset()         { *m = Student{} }
func (m *Student) String() string { return proto.CompactTextString(m) }
func (*Student) ProtoMessage()    {}
func (*Student) Descriptor() ([]byte, []int) {
	return fileDescriptor_94a1c1b032ad0c00, []int{0}
}

func (m *Student) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Student.Unmarshal(m, b)
}
func (m *Student) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Student.Marshal(b, m, deterministic)
}
func (m *Student) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Student.Merge(m, src)
}
func (m *Student) XXX_Size() int {
	return xxx_messageInfo_Student.Size(m)
}
func (m *Student) XXX_DiscardUnknown() {
	xxx_messageInfo_Student.DiscardUnknown(m)
}

var xxx_messageInfo_Student proto.InternalMessageInfo

func (m *Student) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Student) GetAge() string {
	if m != nil {
		return m.Age
	}
	return ""
}

func (m *Student) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *Student) GetNumber() int32 {
	if m != nil {
		return m.Number
	}
	return 0
}

func init() {
	proto.RegisterType((*Student)(nil), "Student")
}

func init() { proto.RegisterFile("student.proto", fileDescriptor_94a1c1b032ad0c00) }

var fileDescriptor_94a1c1b032ad0c00 = []byte{
	// 113 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x2e, 0x29, 0x4d,
	0x49, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x8a, 0xe7, 0x62, 0x0f, 0x86, 0x08,
	0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81,
	0xd9, 0x42, 0x02, 0x5c, 0xcc, 0x89, 0xe9, 0xa9, 0x12, 0x4c, 0x60, 0x21, 0x10, 0x53, 0x48, 0x8c,
	0x8b, 0x2d, 0x3d, 0x35, 0x2f, 0x25, 0xb5, 0x48, 0x82, 0x19, 0x2c, 0x08, 0xe5, 0x81, 0xc4, 0xf3,
	0x4a, 0x73, 0x93, 0x52, 0x8b, 0x24, 0x58, 0x14, 0x18, 0x35, 0x58, 0x83, 0xa0, 0xbc, 0x24, 0x36,
	0xb0, 0x3d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x24, 0x3f, 0x17, 0x78, 0x00, 0x00,
	0x00,
}
