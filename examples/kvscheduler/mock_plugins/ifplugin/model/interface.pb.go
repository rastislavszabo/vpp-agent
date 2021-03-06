// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/interface.proto

package mock_interfaces

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

type Interface_Type int32

const (
	Interface_UNDEFINED_TYPE Interface_Type = 0
	Interface_LOOPBACK       Interface_Type = 1
	Interface_TAP            Interface_Type = 2
)

var Interface_Type_name = map[int32]string{
	0: "UNDEFINED_TYPE",
	1: "LOOPBACK",
	2: "TAP",
}

var Interface_Type_value = map[string]int32{
	"UNDEFINED_TYPE": 0,
	"LOOPBACK":       1,
	"TAP":            2,
}

func (x Interface_Type) String() string {
	return proto.EnumName(Interface_Type_name, int32(x))
}

func (Interface_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d354b33e68715588, []int{0, 0}
}

type Interface struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 Interface_Type `protobuf:"varint,2,opt,name=type,proto3,enum=mock.interfaces.Interface_Type" json:"type,omitempty"`
	Enabled              bool           `protobuf:"varint,3,opt,name=enabled,proto3" json:"enabled,omitempty"`
	PhysAddress          string         `protobuf:"bytes,4,opt,name=phys_address,json=physAddress,proto3" json:"phys_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Interface) Reset()         { *m = Interface{} }
func (m *Interface) String() string { return proto.CompactTextString(m) }
func (*Interface) ProtoMessage()    {}
func (*Interface) Descriptor() ([]byte, []int) {
	return fileDescriptor_d354b33e68715588, []int{0}
}

func (m *Interface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Interface.Unmarshal(m, b)
}
func (m *Interface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Interface.Marshal(b, m, deterministic)
}
func (m *Interface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Interface.Merge(m, src)
}
func (m *Interface) XXX_Size() int {
	return xxx_messageInfo_Interface.Size(m)
}
func (m *Interface) XXX_DiscardUnknown() {
	xxx_messageInfo_Interface.DiscardUnknown(m)
}

var xxx_messageInfo_Interface proto.InternalMessageInfo

func (m *Interface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Interface) GetType() Interface_Type {
	if m != nil {
		return m.Type
	}
	return Interface_UNDEFINED_TYPE
}

func (m *Interface) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Interface) GetPhysAddress() string {
	if m != nil {
		return m.PhysAddress
	}
	return ""
}

func init() {
	proto.RegisterEnum("mock.interfaces.Interface_Type", Interface_Type_name, Interface_Type_value)
	proto.RegisterType((*Interface)(nil), "mock.interfaces.Interface")
}

func init() { proto.RegisterFile("model/interface.proto", fileDescriptor_d354b33e68715588) }

var fileDescriptor_d354b33e68715588 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0x4a, 0x4b, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0xcf, 0xcd, 0x4f, 0xce, 0xd6, 0x83, 0x8b, 0x16, 0x2b, 0x1d, 0x61, 0xe4, 0xe2,
	0xf4, 0x84, 0x71, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35,
	0x38, 0x83, 0xc0, 0x6c, 0x21, 0x63, 0x2e, 0x96, 0x92, 0xca, 0x82, 0x54, 0x09, 0x26, 0x05, 0x46,
	0x0d, 0x3e, 0x23, 0x79, 0x3d, 0x34, 0x13, 0xf4, 0xe0, 0xba, 0xf5, 0x42, 0x2a, 0x0b, 0x52, 0x83,
	0xc0, 0x8a, 0x85, 0x24, 0xb8, 0xd8, 0x53, 0xf3, 0x12, 0x93, 0x72, 0x52, 0x53, 0x24, 0x98, 0x15,
	0x18, 0x35, 0x38, 0x82, 0x60, 0x5c, 0x21, 0x45, 0x2e, 0x9e, 0x82, 0x8c, 0xca, 0xe2, 0xf8, 0xc4,
	0x94, 0x94, 0xa2, 0xd4, 0xe2, 0x62, 0x09, 0x16, 0xb0, 0x55, 0xdc, 0x20, 0x31, 0x47, 0x88, 0x90,
	0x92, 0x21, 0x17, 0x0b, 0xc8, 0x28, 0x21, 0x21, 0x2e, 0xbe, 0x50, 0x3f, 0x17, 0x57, 0x37, 0x4f,
	0x3f, 0x57, 0x97, 0xf8, 0x90, 0xc8, 0x00, 0x57, 0x01, 0x06, 0x21, 0x1e, 0x2e, 0x0e, 0x1f, 0x7f,
	0xff, 0x00, 0x27, 0x47, 0x67, 0x6f, 0x01, 0x46, 0x21, 0x76, 0x2e, 0xe6, 0x10, 0xc7, 0x00, 0x01,
	0xa6, 0x24, 0x36, 0xb0, 0xf7, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x59, 0xf7, 0xa4, 0x54,
	0xf7, 0x00, 0x00, 0x00,
}
