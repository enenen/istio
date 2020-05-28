// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/listener/proxy_protocol/v2/proxy_protocol.proto

package envoy_config_filter_listener_proxy_protocol_v2

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type ProxyProtocol struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProxyProtocol) Reset()         { *m = ProxyProtocol{} }
func (m *ProxyProtocol) String() string { return proto.CompactTextString(m) }
func (*ProxyProtocol) ProtoMessage()    {}
func (*ProxyProtocol) Descriptor() ([]byte, []int) {
	return fileDescriptor_614eea5d50c8eea0, []int{0}
}

func (m *ProxyProtocol) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProxyProtocol.Unmarshal(m, b)
}
func (m *ProxyProtocol) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProxyProtocol.Marshal(b, m, deterministic)
}
func (m *ProxyProtocol) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProxyProtocol.Merge(m, src)
}
func (m *ProxyProtocol) XXX_Size() int {
	return xxx_messageInfo_ProxyProtocol.Size(m)
}
func (m *ProxyProtocol) XXX_DiscardUnknown() {
	xxx_messageInfo_ProxyProtocol.DiscardUnknown(m)
}

var xxx_messageInfo_ProxyProtocol proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProxyProtocol)(nil), "envoy.config.filter.listener.proxy_protocol.v2.ProxyProtocol")
}

func init() {
	proto.RegisterFile("envoy/config/filter/listener/proxy_protocol/v2/proxy_protocol.proto", fileDescriptor_614eea5d50c8eea0)
}

var fileDescriptor_614eea5d50c8eea0 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x4e, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xcb, 0xcc, 0x29, 0x49, 0x2d, 0xd2,
	0xcf, 0xc9, 0x2c, 0x2e, 0x49, 0xcd, 0x4b, 0x2d, 0xd2, 0x2f, 0x28, 0xca, 0xaf, 0xa8, 0x8c, 0x2f,
	0x28, 0xca, 0x2f, 0xc9, 0x4f, 0xce, 0xcf, 0xd1, 0x2f, 0x33, 0x42, 0x13, 0xd1, 0x03, 0x33, 0x84,
	0xf4, 0xc0, 0x86, 0xe8, 0x41, 0x0c, 0xd1, 0x83, 0x18, 0xa2, 0x07, 0x33, 0x44, 0x0f, 0x4d, 0x4b,
	0x99, 0x91, 0x94, 0x5c, 0x69, 0x4a, 0x41, 0xa2, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49, 0x62, 0x49,
	0x66, 0x7e, 0x5e, 0xb1, 0x7e, 0x6e, 0x66, 0x7a, 0x51, 0x62, 0x49, 0x2a, 0xc4, 0x3c, 0x29, 0x59,
	0x0c, 0xf9, 0xe2, 0x92, 0xc4, 0x92, 0xd2, 0x62, 0x88, 0xb4, 0x12, 0x3f, 0x17, 0x6f, 0x00, 0xc8,
	0xcc, 0x00, 0xa8, 0x91, 0x4e, 0xd3, 0x19, 0x3f, 0xcd, 0xf8, 0xd7, 0xcf, 0x6a, 0x2a, 0x64, 0x0c,
	0x71, 0x48, 0x6a, 0x45, 0x49, 0x6a, 0x5e, 0x31, 0x48, 0x23, 0xd4, 0x31, 0xc5, 0xb8, 0x5d, 0x63,
	0xbc, 0xab, 0xe1, 0xc4, 0x45, 0x36, 0x26, 0x01, 0x46, 0x2e, 0x9b, 0xcc, 0x7c, 0x88, 0x47, 0xc0,
	0x4a, 0x48, 0xf4, 0x93, 0x93, 0x10, 0x8a, 0x93, 0xc0, 0x74, 0x00, 0x63, 0x12, 0x1b, 0x58, 0x89,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x28, 0x9c, 0x3b, 0x21, 0x67, 0x01, 0x00, 0x00,
}
