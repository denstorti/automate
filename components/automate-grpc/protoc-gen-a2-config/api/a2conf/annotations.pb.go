// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-grpc/protoc-gen-a2-config/api/a2conf/annotations.proto

package a2conf // import "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ServiceConfig struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceConfig) Reset()         { *m = ServiceConfig{} }
func (m *ServiceConfig) String() string { return proto.CompactTextString(m) }
func (*ServiceConfig) ProtoMessage()    {}
func (*ServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_annotations_0dba191400d6cefb, []int{0}
}
func (m *ServiceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceConfig.Unmarshal(m, b)
}
func (m *ServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceConfig.Marshal(b, m, deterministic)
}
func (dst *ServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceConfig.Merge(dst, src)
}
func (m *ServiceConfig) XXX_Size() int {
	return xxx_messageInfo_ServiceConfig.Size(m)
}
func (m *ServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceConfig proto.InternalMessageInfo

func (m *ServiceConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

var E_Port = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*Port)(nil),
	Field:         51000,
	Name:          "chef.automate.api.port",
	Tag:           "bytes,51000,opt,name=port",
	Filename:      "components/automate-grpc/protoc-gen-a2-config/api/a2conf/annotations.proto",
}

var E_ServiceConfig = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*ServiceConfig)(nil),
	Field:         51000,
	Name:          "chef.automate.api.service_config",
	Tag:           "bytes,51000,opt,name=service_config,json=serviceConfig",
	Filename:      "components/automate-grpc/protoc-gen-a2-config/api/a2conf/annotations.proto",
}

func init() {
	proto.RegisterType((*ServiceConfig)(nil), "chef.automate.api.ServiceConfig")
	proto.RegisterExtension(E_Port)
	proto.RegisterExtension(E_ServiceConfig)
}

func init() {
	proto.RegisterFile("components/automate-grpc/protoc-gen-a2-config/api/a2conf/annotations.proto", fileDescriptor_annotations_0dba191400d6cefb)
}

var fileDescriptor_annotations_0dba191400d6cefb = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd0, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x06, 0x60, 0x16, 0x8b, 0xe0, 0x4a, 0x05, 0x73, 0xb1, 0x14, 0xc4, 0x45, 0x2f, 0xbd, 0x6c,
	0x02, 0xeb, 0xad, 0x47, 0x0b, 0x1e, 0x44, 0x51, 0xb7, 0x37, 0x2f, 0x92, 0x4d, 0x67, 0xd3, 0x48,
	0x37, 0x13, 0x92, 0x59, 0x1f, 0xc3, 0xe7, 0xf1, 0xf1, 0x64, 0x13, 0x2a, 0x6a, 0x3d, 0xf5, 0x96,
	0x81, 0x7f, 0xbe, 0xcc, 0x4c, 0x7e, 0xa7, 0xb0, 0x73, 0x68, 0xc1, 0x52, 0x10, 0xb2, 0x27, 0xec,
	0x24, 0x41, 0xa9, 0xbd, 0x53, 0xc2, 0x79, 0x24, 0x54, 0xa5, 0x06, 0x5b, 0xca, 0xaa, 0x54, 0x68,
	0x5b, 0xa3, 0x85, 0x74, 0x46, 0xc8, 0x6a, 0x28, 0x84, 0xb4, 0x16, 0x49, 0x92, 0x41, 0x1b, 0x78,
	0xcc, 0xb2, 0x53, 0xb5, 0x86, 0x96, 0x6f, 0x15, 0x2e, 0x9d, 0x99, 0x2e, 0xf6, 0xe6, 0x1d, 0x7a,
	0x4a, 0xee, 0xb4, 0xd0, 0x88, 0x7a, 0x03, 0xa9, 0xa5, 0xe9, 0x5b, 0xb1, 0x82, 0xa0, 0xbc, 0x71,
	0x84, 0x3e, 0x25, 0x2e, 0xaf, 0xf2, 0xf1, 0x12, 0xfc, 0xbb, 0x51, 0xb0, 0x88, 0x10, 0x63, 0xf9,
	0xc8, 0xca, 0x0e, 0x26, 0x59, 0x91, 0xcd, 0x8e, 0xea, 0xf8, 0x9e, 0xdf, 0xe7, 0xa3, 0x01, 0x65,
	0xe7, 0x3c, 0x79, 0x7c, 0xeb, 0xf1, 0x5b, 0x03, 0x9b, 0xd5, 0xa3, 0x8b, 0xbb, 0x4c, 0x3e, 0x3f,
	0x0e, 0x8a, 0x6c, 0x76, 0x5c, 0x9d, 0xf1, 0x9d, 0x75, 0xf8, 0x13, 0x7a, 0xaa, 0xa3, 0x32, 0x7f,
	0xcb, 0x4f, 0x42, 0xfa, 0xf2, 0x35, 0x0d, 0xcf, 0x2e, 0x76, 0xdc, 0x07, 0x08, 0x41, 0x6a, 0xf8,
	0x2b, 0x17, 0xff, 0xc8, 0xbf, 0xc6, 0xaf, 0xc7, 0xe1, 0x67, 0x79, 0xb3, 0x7c, 0x79, 0xd6, 0x86,
	0xd6, 0x7d, 0xc3, 0x15, 0x76, 0x62, 0x68, 0xfe, 0x3e, 0xa6, 0xd8, 0xf7, 0xc0, 0xcd, 0x61, 0x0c,
	0x5c, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x26, 0x20, 0x0c, 0x02, 0x02, 0x00, 0x00,
}
