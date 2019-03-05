// Code generated by protoc-gen-go. DO NOT EDIT.
// source: control_exchange.proto

package control_exchange

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

type DeleteConnectorReq struct {
	ServerName           string   `protobuf:"bytes,1,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	ServerIp             string   `protobuf:"bytes,2,opt,name=server_ip,json=serverIp,proto3" json:"server_ip,omitempty"`
	ServerPort           uint32   `protobuf:"varint,3,opt,name=server_port,json=serverPort,proto3" json:"server_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteConnectorReq) Reset()         { *m = DeleteConnectorReq{} }
func (m *DeleteConnectorReq) String() string { return proto.CompactTextString(m) }
func (*DeleteConnectorReq) ProtoMessage()    {}
func (*DeleteConnectorReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f59c15b12779ca0, []int{0}
}

func (m *DeleteConnectorReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteConnectorReq.Unmarshal(m, b)
}
func (m *DeleteConnectorReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteConnectorReq.Marshal(b, m, deterministic)
}
func (m *DeleteConnectorReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteConnectorReq.Merge(m, src)
}
func (m *DeleteConnectorReq) XXX_Size() int {
	return xxx_messageInfo_DeleteConnectorReq.Size(m)
}
func (m *DeleteConnectorReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteConnectorReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteConnectorReq proto.InternalMessageInfo

func (m *DeleteConnectorReq) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *DeleteConnectorReq) GetServerIp() string {
	if m != nil {
		return m.ServerIp
	}
	return ""
}

func (m *DeleteConnectorReq) GetServerPort() uint32 {
	if m != nil {
		return m.ServerPort
	}
	return 0
}

type DeleteConnectorAck struct {
	Ret                  uint32   `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteConnectorAck) Reset()         { *m = DeleteConnectorAck{} }
func (m *DeleteConnectorAck) String() string { return proto.CompactTextString(m) }
func (*DeleteConnectorAck) ProtoMessage()    {}
func (*DeleteConnectorAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f59c15b12779ca0, []int{1}
}

func (m *DeleteConnectorAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteConnectorAck.Unmarshal(m, b)
}
func (m *DeleteConnectorAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteConnectorAck.Marshal(b, m, deterministic)
}
func (m *DeleteConnectorAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteConnectorAck.Merge(m, src)
}
func (m *DeleteConnectorAck) XXX_Size() int {
	return xxx_messageInfo_DeleteConnectorAck.Size(m)
}
func (m *DeleteConnectorAck) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteConnectorAck.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteConnectorAck proto.InternalMessageInfo

func (m *DeleteConnectorAck) GetRet() uint32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

type AddConnectorReq struct {
	ServerName           string   `protobuf:"bytes,1,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	ServerIp             string   `protobuf:"bytes,2,opt,name=server_ip,json=serverIp,proto3" json:"server_ip,omitempty"`
	ServerPort           uint32   `protobuf:"varint,3,opt,name=server_port,json=serverPort,proto3" json:"server_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddConnectorReq) Reset()         { *m = AddConnectorReq{} }
func (m *AddConnectorReq) String() string { return proto.CompactTextString(m) }
func (*AddConnectorReq) ProtoMessage()    {}
func (*AddConnectorReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f59c15b12779ca0, []int{2}
}

func (m *AddConnectorReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddConnectorReq.Unmarshal(m, b)
}
func (m *AddConnectorReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddConnectorReq.Marshal(b, m, deterministic)
}
func (m *AddConnectorReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddConnectorReq.Merge(m, src)
}
func (m *AddConnectorReq) XXX_Size() int {
	return xxx_messageInfo_AddConnectorReq.Size(m)
}
func (m *AddConnectorReq) XXX_DiscardUnknown() {
	xxx_messageInfo_AddConnectorReq.DiscardUnknown(m)
}

var xxx_messageInfo_AddConnectorReq proto.InternalMessageInfo

func (m *AddConnectorReq) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

func (m *AddConnectorReq) GetServerIp() string {
	if m != nil {
		return m.ServerIp
	}
	return ""
}

func (m *AddConnectorReq) GetServerPort() uint32 {
	if m != nil {
		return m.ServerPort
	}
	return 0
}

type AddConnectorAck struct {
	Ret                  uint32   `protobuf:"varint,1,opt,name=ret,proto3" json:"ret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddConnectorAck) Reset()         { *m = AddConnectorAck{} }
func (m *AddConnectorAck) String() string { return proto.CompactTextString(m) }
func (*AddConnectorAck) ProtoMessage()    {}
func (*AddConnectorAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f59c15b12779ca0, []int{3}
}

func (m *AddConnectorAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddConnectorAck.Unmarshal(m, b)
}
func (m *AddConnectorAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddConnectorAck.Marshal(b, m, deterministic)
}
func (m *AddConnectorAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddConnectorAck.Merge(m, src)
}
func (m *AddConnectorAck) XXX_Size() int {
	return xxx_messageInfo_AddConnectorAck.Size(m)
}
func (m *AddConnectorAck) XXX_DiscardUnknown() {
	xxx_messageInfo_AddConnectorAck.DiscardUnknown(m)
}

var xxx_messageInfo_AddConnectorAck proto.InternalMessageInfo

func (m *AddConnectorAck) GetRet() uint32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

type GetConnectorReq struct {
	ServerName           string   `protobuf:"bytes,1,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConnectorReq) Reset()         { *m = GetConnectorReq{} }
func (m *GetConnectorReq) String() string { return proto.CompactTextString(m) }
func (*GetConnectorReq) ProtoMessage()    {}
func (*GetConnectorReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f59c15b12779ca0, []int{4}
}

func (m *GetConnectorReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConnectorReq.Unmarshal(m, b)
}
func (m *GetConnectorReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConnectorReq.Marshal(b, m, deterministic)
}
func (m *GetConnectorReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConnectorReq.Merge(m, src)
}
func (m *GetConnectorReq) XXX_Size() int {
	return xxx_messageInfo_GetConnectorReq.Size(m)
}
func (m *GetConnectorReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConnectorReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetConnectorReq proto.InternalMessageInfo

func (m *GetConnectorReq) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

type GetConnectorAck struct {
	ServerAddrs          []string `protobuf:"bytes,1,rep,name=server_addrs,json=serverAddrs,proto3" json:"server_addrs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConnectorAck) Reset()         { *m = GetConnectorAck{} }
func (m *GetConnectorAck) String() string { return proto.CompactTextString(m) }
func (*GetConnectorAck) ProtoMessage()    {}
func (*GetConnectorAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_6f59c15b12779ca0, []int{5}
}

func (m *GetConnectorAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConnectorAck.Unmarshal(m, b)
}
func (m *GetConnectorAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConnectorAck.Marshal(b, m, deterministic)
}
func (m *GetConnectorAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConnectorAck.Merge(m, src)
}
func (m *GetConnectorAck) XXX_Size() int {
	return xxx_messageInfo_GetConnectorAck.Size(m)
}
func (m *GetConnectorAck) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConnectorAck.DiscardUnknown(m)
}

var xxx_messageInfo_GetConnectorAck proto.InternalMessageInfo

func (m *GetConnectorAck) GetServerAddrs() []string {
	if m != nil {
		return m.ServerAddrs
	}
	return nil
}

func init() {
	proto.RegisterType((*DeleteConnectorReq)(nil), "control_exchange.DeleteConnectorReq")
	proto.RegisterType((*DeleteConnectorAck)(nil), "control_exchange.DeleteConnectorAck")
	proto.RegisterType((*AddConnectorReq)(nil), "control_exchange.AddConnectorReq")
	proto.RegisterType((*AddConnectorAck)(nil), "control_exchange.AddConnectorAck")
	proto.RegisterType((*GetConnectorReq)(nil), "control_exchange.getConnectorReq")
	proto.RegisterType((*GetConnectorAck)(nil), "control_exchange.getConnectorAck")
}

func init() { proto.RegisterFile("control_exchange.proto", fileDescriptor_6f59c15b12779ca0) }

var fileDescriptor_6f59c15b12779ca0 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xce, 0xcf, 0x2b,
	0x29, 0xca, 0xcf, 0x89, 0x4f, 0xad, 0x48, 0xce, 0x48, 0xcc, 0x4b, 0x4f, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x40, 0x17, 0x57, 0x2a, 0xe6, 0x12, 0x72, 0x49, 0xcd, 0x49, 0x2d, 0x49,
	0x75, 0xce, 0xcf, 0xcb, 0x4b, 0x4d, 0x2e, 0xc9, 0x2f, 0x0a, 0x4a, 0x2d, 0x14, 0x92, 0xe7, 0xe2,
	0x2e, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0x8a, 0xcf, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0xe2, 0x82, 0x08, 0xf9, 0x25, 0xe6, 0xa6, 0x0a, 0x49, 0x73, 0x71, 0x42, 0x15, 0x64,
	0x16, 0x48, 0x30, 0x81, 0xa5, 0x39, 0x20, 0x02, 0x9e, 0x05, 0x48, 0xba, 0x0b, 0xf2, 0x8b, 0x4a,
	0x24, 0x98, 0x15, 0x18, 0x35, 0x78, 0x61, 0xba, 0x03, 0xf2, 0x8b, 0x4a, 0x94, 0xd4, 0x30, 0x2c,
	0x75, 0x4c, 0xce, 0x16, 0x12, 0xe0, 0x62, 0x2e, 0x4a, 0x2d, 0x01, 0x5b, 0xc6, 0x1b, 0x04, 0x62,
	0x2a, 0x15, 0x70, 0xf1, 0x3b, 0xa6, 0xa4, 0xd0, 0xd3, 0x65, 0xca, 0xa8, 0x36, 0x62, 0x77, 0x96,
	0x11, 0x17, 0x7f, 0x7a, 0x6a, 0x09, 0x49, 0xce, 0x52, 0x32, 0x41, 0xd5, 0x03, 0x32, 0x58, 0x91,
	0x8b, 0x07, 0xaa, 0x27, 0x31, 0x25, 0xa5, 0xa8, 0x58, 0x82, 0x51, 0x81, 0x59, 0x83, 0x33, 0x08,
	0x6a, 0x8e, 0x23, 0x48, 0x28, 0x89, 0x0d, 0x1c, 0x6d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xd3, 0x86, 0xbe, 0x3e, 0xd0, 0x01, 0x00, 0x00,
}