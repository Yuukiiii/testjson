// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tcpd_service.proto

package tcpd_service

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

type ServiceReq struct {
	Requestid            uint32    `protobuf:"varint,1,opt,name=requestid,proto3" json:"requestid,omitempty"`
	UserId               string    `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserNetid            uint32    `protobuf:"varint,3,opt,name=user_netid,json=userNetid,proto3" json:"user_netid,omitempty"`
	Data                 string    `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	FieldType            FieldType `protobuf:"varint,5,opt,name=field_type,json=fieldType,proto3,enum=FieldType" json:"field_type,omitempty"`
	DeviceId             string    `protobuf:"bytes,6,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	TcpdId               uint32    `protobuf:"varint,7,opt,name=tcpd_id,json=tcpdId,proto3" json:"tcpd_id,omitempty"`
	ProtocolId           uint32    `protobuf:"varint,8,opt,name=protocol_id,json=protocolId,proto3" json:"protocol_id,omitempty"`
	ClientApp            string    `protobuf:"bytes,9,opt,name=client_app,json=clientApp,proto3" json:"client_app,omitempty"`
	AppVersion           string    `protobuf:"bytes,10,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	Language             string    `protobuf:"bytes,11,opt,name=language,proto3" json:"language,omitempty"`
	Country              string    `protobuf:"bytes,12,opt,name=country,proto3" json:"country,omitempty"`
	ClientIp             uint32    `protobuf:"varint,13,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	PlatformId           uint32    `protobuf:"varint,14,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`
	ProductId            uint32    `protobuf:"varint,15,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ServiceReq) Reset()         { *m = ServiceReq{} }
func (m *ServiceReq) String() string { return proto.CompactTextString(m) }
func (*ServiceReq) ProtoMessage()    {}
func (*ServiceReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d4919d7917e0287, []int{0}
}

func (m *ServiceReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceReq.Unmarshal(m, b)
}
func (m *ServiceReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceReq.Marshal(b, m, deterministic)
}
func (m *ServiceReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceReq.Merge(m, src)
}
func (m *ServiceReq) XXX_Size() int {
	return xxx_messageInfo_ServiceReq.Size(m)
}
func (m *ServiceReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceReq.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceReq proto.InternalMessageInfo

func (m *ServiceReq) GetRequestid() uint32 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *ServiceReq) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ServiceReq) GetUserNetid() uint32 {
	if m != nil {
		return m.UserNetid
	}
	return 0
}

func (m *ServiceReq) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *ServiceReq) GetFieldType() FieldType {
	if m != nil {
		return m.FieldType
	}
	return FieldType_UID
}

func (m *ServiceReq) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *ServiceReq) GetTcpdId() uint32 {
	if m != nil {
		return m.TcpdId
	}
	return 0
}

func (m *ServiceReq) GetProtocolId() uint32 {
	if m != nil {
		return m.ProtocolId
	}
	return 0
}

func (m *ServiceReq) GetClientApp() string {
	if m != nil {
		return m.ClientApp
	}
	return ""
}

func (m *ServiceReq) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

func (m *ServiceReq) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *ServiceReq) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *ServiceReq) GetClientIp() uint32 {
	if m != nil {
		return m.ClientIp
	}
	return 0
}

func (m *ServiceReq) GetPlatformId() uint32 {
	if m != nil {
		return m.PlatformId
	}
	return 0
}

func (m *ServiceReq) GetProductId() uint32 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

type ServiceAck struct {
	Requestid            uint32   `protobuf:"varint,1,opt,name=requestid,proto3" json:"requestid,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserNetid            uint32   `protobuf:"varint,3,opt,name=user_netid,json=userNetid,proto3" json:"user_netid,omitempty"`
	Data                 string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceAck) Reset()         { *m = ServiceAck{} }
func (m *ServiceAck) String() string { return proto.CompactTextString(m) }
func (*ServiceAck) ProtoMessage()    {}
func (*ServiceAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d4919d7917e0287, []int{1}
}

func (m *ServiceAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceAck.Unmarshal(m, b)
}
func (m *ServiceAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceAck.Marshal(b, m, deterministic)
}
func (m *ServiceAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceAck.Merge(m, src)
}
func (m *ServiceAck) XXX_Size() int {
	return xxx_messageInfo_ServiceAck.Size(m)
}
func (m *ServiceAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceAck.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceAck proto.InternalMessageInfo

func (m *ServiceAck) GetRequestid() uint32 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *ServiceAck) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ServiceAck) GetUserNetid() uint32 {
	if m != nil {
		return m.UserNetid
	}
	return 0
}

func (m *ServiceAck) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type ServiceMultiAck struct {
	Requestid            uint32   `protobuf:"varint,1,opt,name=requestid,proto3" json:"requestid,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserNetid            uint32   `protobuf:"varint,3,opt,name=user_netid,json=userNetid,proto3" json:"user_netid,omitempty"`
	Data                 string   `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceMultiAck) Reset()         { *m = ServiceMultiAck{} }
func (m *ServiceMultiAck) String() string { return proto.CompactTextString(m) }
func (*ServiceMultiAck) ProtoMessage()    {}
func (*ServiceMultiAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d4919d7917e0287, []int{2}
}

func (m *ServiceMultiAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceMultiAck.Unmarshal(m, b)
}
func (m *ServiceMultiAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceMultiAck.Marshal(b, m, deterministic)
}
func (m *ServiceMultiAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceMultiAck.Merge(m, src)
}
func (m *ServiceMultiAck) XXX_Size() int {
	return xxx_messageInfo_ServiceMultiAck.Size(m)
}
func (m *ServiceMultiAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceMultiAck.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceMultiAck proto.InternalMessageInfo

func (m *ServiceMultiAck) GetRequestid() uint32 {
	if m != nil {
		return m.Requestid
	}
	return 0
}

func (m *ServiceMultiAck) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *ServiceMultiAck) GetUserNetid() uint32 {
	if m != nil {
		return m.UserNetid
	}
	return 0
}

func (m *ServiceMultiAck) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type SynLogout struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DeviceId             string   `protobuf:"bytes,2,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	AppVersion           string   `protobuf:"bytes,3,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	ProtocolId           uint32   `protobuf:"varint,4,opt,name=protocol_id,json=protocolId,proto3" json:"protocol_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SynLogout) Reset()         { *m = SynLogout{} }
func (m *SynLogout) String() string { return proto.CompactTextString(m) }
func (*SynLogout) ProtoMessage()    {}
func (*SynLogout) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d4919d7917e0287, []int{3}
}

func (m *SynLogout) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SynLogout.Unmarshal(m, b)
}
func (m *SynLogout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SynLogout.Marshal(b, m, deterministic)
}
func (m *SynLogout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SynLogout.Merge(m, src)
}
func (m *SynLogout) XXX_Size() int {
	return xxx_messageInfo_SynLogout.Size(m)
}
func (m *SynLogout) XXX_DiscardUnknown() {
	xxx_messageInfo_SynLogout.DiscardUnknown(m)
}

var xxx_messageInfo_SynLogout proto.InternalMessageInfo

func (m *SynLogout) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *SynLogout) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *SynLogout) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

func (m *SynLogout) GetProtocolId() uint32 {
	if m != nil {
		return m.ProtocolId
	}
	return 0
}

func init() {
	proto.RegisterType((*ServiceReq)(nil), "tcpd_service.ServiceReq")
	proto.RegisterType((*ServiceAck)(nil), "tcpd_service.ServiceAck")
	proto.RegisterType((*ServiceMultiAck)(nil), "tcpd_service.ServiceMultiAck")
	proto.RegisterType((*SynLogout)(nil), "tcpd_service.SynLogout")
}

func init() { proto.RegisterFile("tcpd_service.proto", fileDescriptor_5d4919d7917e0287) }

var fileDescriptor_5d4919d7917e0287 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x95, 0xb6, 0xa4, 0xc9, 0xf5, 0x9f, 0xe4, 0x05, 0xab, 0x80, 0xa8, 0x3a, 0x95, 0xa5,
	0x03, 0x3c, 0x41, 0x17, 0xa4, 0x48, 0xc0, 0x90, 0x22, 0xd6, 0xc8, 0xc4, 0x6e, 0xb1, 0x08, 0xb1,
	0xeb, 0x38, 0x95, 0x22, 0x46, 0x1e, 0x8f, 0x97, 0x42, 0xb6, 0x13, 0x68, 0xc3, 0xde, 0xcd, 0xf7,
	0xfb, 0x74, 0x77, 0xdf, 0xf9, 0x0e, 0x90, 0x4e, 0x25, 0x4d, 0x0a, 0xa6, 0xf6, 0x3c, 0x65, 0x4b,
	0xa9, 0x84, 0x16, 0x68, 0x78, 0xc8, 0xa6, 0xf0, 0xc6, 0x08, 0x75, 0xca, 0xfc, 0xbb, 0x0b, 0xb0,
	0x76, 0x3c, 0x66, 0x3b, 0x74, 0x09, 0xa1, 0x62, 0xbb, 0x92, 0x15, 0x9a, 0x53, 0xec, 0xcd, 0xbc,
	0xc5, 0x28, 0xfe, 0x03, 0xe8, 0x1c, 0xfa, 0x65, 0xc1, 0x54, 0xc2, 0x29, 0xee, 0xcc, 0xbc, 0x45,
	0x18, 0xfb, 0x26, 0x8c, 0x28, 0xba, 0x02, 0xb0, 0x42, 0xce, 0x4c, 0x5e, 0xd7, 0xe5, 0x19, 0xf2,
	0x64, 0x00, 0x42, 0xd0, 0xa3, 0x44, 0x13, 0xdc, 0xb3, 0x49, 0xf6, 0x8d, 0x6e, 0x00, 0x36, 0x9c,
	0x65, 0x34, 0xd1, 0x95, 0x64, 0xf8, 0x6c, 0xe6, 0x2d, 0xc6, 0xb7, 0xb0, 0xbc, 0x37, 0xe8, 0xb9,
	0x92, 0x2c, 0x0e, 0x37, 0xcd, 0x13, 0x5d, 0x40, 0x48, 0x99, 0x71, 0x68, 0x1a, 0xfb, 0xb6, 0x46,
	0xe0, 0x40, 0x64, 0x3d, 0xd9, 0xe1, 0x38, 0xc5, 0x7d, 0xdb, 0xd7, 0x37, 0x61, 0x44, 0xd1, 0x35,
	0x0c, 0xec, 0x88, 0xa9, 0xc8, 0x8c, 0x18, 0x58, 0x11, 0x1a, 0xe4, 0x4c, 0xa7, 0x19, 0x67, 0xb9,
	0x4e, 0x88, 0x94, 0x38, 0xb4, 0x75, 0x43, 0x47, 0x56, 0x52, 0x9a, 0x7c, 0x22, 0x65, 0xb2, 0x67,
	0xaa, 0xe0, 0x22, 0xc7, 0x60, 0x75, 0x20, 0x52, 0xbe, 0x38, 0x82, 0xa6, 0x10, 0x64, 0x24, 0xdf,
	0x96, 0x64, 0xcb, 0xf0, 0xc0, 0xb9, 0x6a, 0x62, 0x84, 0xa1, 0x9f, 0x8a, 0x32, 0xd7, 0xaa, 0xc2,
	0x43, 0x2b, 0x35, 0xa1, 0x19, 0xa6, 0xee, 0xca, 0x25, 0x1e, 0x59, 0x53, 0x81, 0x03, 0x91, 0xed,
	0x29, 0x33, 0xa2, 0x37, 0x42, 0x7d, 0x18, 0xcf, 0xe3, 0xda, 0x73, 0x8d, 0x9c, 0x67, 0xa9, 0x04,
	0x2d, 0x53, 0x6d, 0xf4, 0x89, 0xfb, 0xe8, 0x9a, 0x44, 0x74, 0xbe, 0xff, 0x5d, 0xe6, 0x2a, 0x7d,
	0x3f, 0xdd, 0x32, 0xe7, 0x9f, 0x30, 0xa9, 0xfb, 0x3e, 0x96, 0x99, 0xe6, 0xa7, 0x6d, 0xfe, 0xe5,
	0x41, 0xb8, 0xae, 0xf2, 0x07, 0xb1, 0x15, 0xa5, 0x3e, 0xac, 0xec, 0x1d, 0x55, 0x3e, 0xba, 0xa2,
	0x4e, 0xeb, 0x8a, 0x5a, 0xcb, 0xee, 0xfe, 0x5b, 0x76, 0xeb, 0x9a, 0x7a, 0xed, 0x6b, 0x7a, 0xf5,
	0xed, 0xfb, 0xee, 0x27, 0x00, 0x00, 0xff, 0xff, 0x0f, 0x78, 0x6a, 0xa1, 0x7f, 0x03, 0x00, 0x00,
}
