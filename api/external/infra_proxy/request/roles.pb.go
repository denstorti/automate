// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/infra_proxy/request/roles.proto

package request

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

type Roles struct {
	OrgId                string   `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	ServerId             string   `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Roles) Reset()         { *m = Roles{} }
func (m *Roles) String() string { return proto.CompactTextString(m) }
func (*Roles) ProtoMessage()    {}
func (*Roles) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4febb01f7e33d74, []int{0}
}

func (m *Roles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Roles.Unmarshal(m, b)
}
func (m *Roles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Roles.Marshal(b, m, deterministic)
}
func (m *Roles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Roles.Merge(m, src)
}
func (m *Roles) XXX_Size() int {
	return xxx_messageInfo_Roles.Size(m)
}
func (m *Roles) XXX_DiscardUnknown() {
	xxx_messageInfo_Roles.DiscardUnknown(m)
}

var xxx_messageInfo_Roles proto.InternalMessageInfo

func (m *Roles) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *Roles) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

type Role struct {
	OrgId                string   `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	ServerId             string   `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4febb01f7e33d74, []int{1}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *Role) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Roles)(nil), "chef.automate.api.infra_proxy.request.Roles")
	proto.RegisterType((*Role)(nil), "chef.automate.api.infra_proxy.request.Role")
}

func init() {
	proto.RegisterFile("api/external/infra_proxy/request/roles.proto", fileDescriptor_f4febb01f7e33d74)
}

var fileDescriptor_f4febb01f7e33d74 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0x2c, 0xc8, 0xd4,
	0x4f, 0xad, 0x28, 0x49, 0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0xcf, 0xcc, 0x4b, 0x2b, 0x4a, 0x8c, 0x2f,
	0x28, 0xca, 0xaf, 0xa8, 0xd4, 0x2f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xd1, 0x2f, 0xca, 0xcf,
	0x49, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x4d, 0xce, 0x48, 0x4d, 0xd3, 0x4b,
	0x2c, 0x2d, 0xc9, 0xcf, 0x4d, 0x2c, 0x49, 0xd5, 0x4b, 0x2c, 0xc8, 0xd4, 0x43, 0xd2, 0xa2, 0x07,
	0xd5, 0xa2, 0x64, 0xcd, 0xc5, 0x1a, 0x04, 0xd2, 0x25, 0x24, 0xca, 0xc5, 0x96, 0x5f, 0x94, 0x1e,
	0x9f, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0xc4, 0x9a, 0x5f, 0x94, 0xee, 0x99, 0x22,
	0x24, 0xcd, 0xc5, 0x59, 0x9c, 0x5a, 0x54, 0x96, 0x5a, 0x04, 0x92, 0x61, 0x02, 0xcb, 0x70, 0x40,
	0x04, 0x3c, 0x53, 0x94, 0xfc, 0xb8, 0x58, 0x40, 0x9a, 0xc9, 0xd1, 0x2b, 0x24, 0xc4, 0xc5, 0x92,
	0x97, 0x98, 0x9b, 0x2a, 0xc1, 0x0c, 0x16, 0x07, 0xb3, 0x9d, 0xac, 0xa3, 0x2c, 0xd3, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x41, 0x1e, 0xd0, 0x87, 0x79, 0x40, 0x9f, 0x90,
	0xe7, 0x93, 0xd8, 0xc0, 0xfe, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x20, 0x91, 0xe5,
	0x27, 0x01, 0x00, 0x00,
}