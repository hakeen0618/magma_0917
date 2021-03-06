// Code generated by protoc-gen-go. DO NOT EDIT.
// source: orc8r/cloud/go/services/configurator/protos/southbound.proto

package protos

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protos "magma/orc8r/lib/go/protos"
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

type GetMconfigRequest struct {
	HardwareID           string   `protobuf:"bytes,1,opt,name=hardwareID,proto3" json:"hardwareID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMconfigRequest) Reset()         { *m = GetMconfigRequest{} }
func (m *GetMconfigRequest) String() string { return proto.CompactTextString(m) }
func (*GetMconfigRequest) ProtoMessage()    {}
func (*GetMconfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_185819ac40d93694, []int{0}
}

func (m *GetMconfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMconfigRequest.Unmarshal(m, b)
}
func (m *GetMconfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMconfigRequest.Marshal(b, m, deterministic)
}
func (m *GetMconfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMconfigRequest.Merge(m, src)
}
func (m *GetMconfigRequest) XXX_Size() int {
	return xxx_messageInfo_GetMconfigRequest.Size(m)
}
func (m *GetMconfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMconfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMconfigRequest proto.InternalMessageInfo

func (m *GetMconfigRequest) GetHardwareID() string {
	if m != nil {
		return m.HardwareID
	}
	return ""
}

type GetMconfigResponse struct {
	// configs contains the mconfigs for the requested hardware ID
	// The contained configs_by_key should be str->any.Any, where the any.Any
	// is a BytesValue wrapper containing the the config.
	// Each config is (1) marshaled to an any.Any then (2) JSON-serialized
	// before being placed into the BytesValue wrapper.
	// TODO(#2310): revert configs.configs_by_key to containing each config just marshaled to any.Any (not additionally serialized to JSON)
	Configs *protos.GatewayConfigs `protobuf:"bytes,1,opt,name=configs,proto3" json:"configs,omitempty"`
	// logicalID is the entity ID of the gateway in the configurator graph
	LogicalID            string   `protobuf:"bytes,2,opt,name=logicalID,proto3" json:"logicalID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMconfigResponse) Reset()         { *m = GetMconfigResponse{} }
func (m *GetMconfigResponse) String() string { return proto.CompactTextString(m) }
func (*GetMconfigResponse) ProtoMessage()    {}
func (*GetMconfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_185819ac40d93694, []int{1}
}

func (m *GetMconfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMconfigResponse.Unmarshal(m, b)
}
func (m *GetMconfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMconfigResponse.Marshal(b, m, deterministic)
}
func (m *GetMconfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMconfigResponse.Merge(m, src)
}
func (m *GetMconfigResponse) XXX_Size() int {
	return xxx_messageInfo_GetMconfigResponse.Size(m)
}
func (m *GetMconfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMconfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMconfigResponse proto.InternalMessageInfo

func (m *GetMconfigResponse) GetConfigs() *protos.GatewayConfigs {
	if m != nil {
		return m.Configs
	}
	return nil
}

func (m *GetMconfigResponse) GetLogicalID() string {
	if m != nil {
		return m.LogicalID
	}
	return ""
}

func init() {
	proto.RegisterType((*GetMconfigRequest)(nil), "magma.orc8r.configurator.GetMconfigRequest")
	proto.RegisterType((*GetMconfigResponse)(nil), "magma.orc8r.configurator.GetMconfigResponse")
}

func init() {
	proto.RegisterFile("orc8r/cloud/go/services/configurator/protos/southbound.proto", fileDescriptor_185819ac40d93694)
}

var fileDescriptor_185819ac40d93694 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x57, 0x0f, 0xca, 0x5e, 0x4f, 0xcb, 0x41, 0x6a, 0x15, 0x91, 0x9e, 0x04, 0x25, 0xc1,
	0x15, 0xc1, 0x83, 0x78, 0x70, 0x83, 0xd1, 0x83, 0x97, 0x0a, 0x1e, 0xbc, 0x65, 0x69, 0xec, 0x0a,
	0x6d, 0xde, 0x2d, 0x7f, 0x1c, 0x7e, 0x46, 0xbf, 0x94, 0x90, 0xa8, 0xcb, 0x10, 0x45, 0x4f, 0x81,
	0xe7, 0x7d, 0x7e, 0xef, 0x9f, 0x87, 0xc0, 0x0d, 0x6a, 0x71, 0xad, 0x99, 0xe8, 0xd0, 0xd5, 0xac,
	0x41, 0x66, 0xa4, 0x7e, 0x69, 0x85, 0x34, 0x4c, 0xa0, 0x7a, 0x6e, 0x1b, 0xa7, 0xb9, 0x45, 0xcd,
	0x96, 0x1a, 0x2d, 0x1a, 0x66, 0xd0, 0xd9, 0xc5, 0x1c, 0x9d, 0xaa, 0xa9, 0x57, 0x48, 0xda, 0xf3,
	0xa6, 0xe7, 0xd4, 0xf7, 0xa0, 0x31, 0x91, 0x65, 0xa1, 0xef, 0x07, 0xd8, 0x87, 0x5a, 0xa0, 0xb2,
	0xc3, 0xad, 0x9a, 0xc0, 0xbe, 0x47, 0x15, 0x4a, 0x79, 0x01, 0xa3, 0x99, 0xb4, 0xf7, 0xc1, 0x5e,
	0xc9, 0x95, 0x93, 0xc6, 0x92, 0x13, 0x80, 0x05, 0xd7, 0xf5, 0x9a, 0x6b, 0x59, 0x4e, 0xd3, 0xe4,
	0x34, 0x39, 0x1b, 0x56, 0x91, 0x92, 0xb7, 0x40, 0x62, 0xc8, 0x2c, 0x51, 0x19, 0x49, 0xae, 0x60,
	0x2f, 0x28, 0xc6, 0x23, 0xfb, 0xe3, 0x23, 0x1a, 0x6f, 0x3b, 0xe3, 0x56, 0xae, 0xf9, 0xeb, 0x24,
	0x58, 0xaa, 0x4f, 0x2f, 0x39, 0x86, 0x61, 0x87, 0x4d, 0x2b, 0x78, 0x57, 0x4e, 0xd3, 0x1d, 0x3f,
	0x6b, 0x23, 0x8c, 0xdf, 0x12, 0x38, 0x78, 0xf8, 0x4a, 0x61, 0x12, 0x5d, 0x4c, 0x6e, 0x01, 0x36,
	0x5b, 0x90, 0xd1, 0xd6, 0xb0, 0x47, 0x6c, 0xeb, 0xec, 0xb7, 0xf9, 0xf9, 0x80, 0xac, 0xe2, 0x2b,
	0x4a, 0x65, 0xa5, 0x56, 0xbc, 0x23, 0xe7, 0xf4, 0xa7, 0x88, 0xe9, 0xb7, 0xa0, 0xb2, 0x8b, 0xbf,
	0x99, 0x43, 0x40, 0xf9, 0xe0, 0xae, 0x78, 0xba, 0xf4, 0x00, 0xfb, 0xc7, 0x27, 0x98, 0xef, 0xfa,
	0xb7, 0x78, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x88, 0x16, 0xd1, 0x99, 0x3a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SouthboundConfiguratorClient is the client API for SouthboundConfigurator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SouthboundConfiguratorClient interface {
	GetMconfig(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.GatewayConfigs, error)
	// GetMconfigInternal exists to support the existing streamer mconfig
	// policy. This should be removed when we migrate gateway mconfig updates
	// from streamer to this southbound configurator servicer.
	GetMconfigInternal(ctx context.Context, in *GetMconfigRequest, opts ...grpc.CallOption) (*GetMconfigResponse, error)
}

type southboundConfiguratorClient struct {
	cc grpc.ClientConnInterface
}

func NewSouthboundConfiguratorClient(cc grpc.ClientConnInterface) SouthboundConfiguratorClient {
	return &southboundConfiguratorClient{cc}
}

func (c *southboundConfiguratorClient) GetMconfig(ctx context.Context, in *protos.Void, opts ...grpc.CallOption) (*protos.GatewayConfigs, error) {
	out := new(protos.GatewayConfigs)
	err := c.cc.Invoke(ctx, "/magma.orc8r.configurator.SouthboundConfigurator/GetMconfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *southboundConfiguratorClient) GetMconfigInternal(ctx context.Context, in *GetMconfigRequest, opts ...grpc.CallOption) (*GetMconfigResponse, error) {
	out := new(GetMconfigResponse)
	err := c.cc.Invoke(ctx, "/magma.orc8r.configurator.SouthboundConfigurator/GetMconfigInternal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SouthboundConfiguratorServer is the server API for SouthboundConfigurator service.
type SouthboundConfiguratorServer interface {
	GetMconfig(context.Context, *protos.Void) (*protos.GatewayConfigs, error)
	// GetMconfigInternal exists to support the existing streamer mconfig
	// policy. This should be removed when we migrate gateway mconfig updates
	// from streamer to this southbound configurator servicer.
	GetMconfigInternal(context.Context, *GetMconfigRequest) (*GetMconfigResponse, error)
}

// UnimplementedSouthboundConfiguratorServer can be embedded to have forward compatible implementations.
type UnimplementedSouthboundConfiguratorServer struct {
}

func (*UnimplementedSouthboundConfiguratorServer) GetMconfig(ctx context.Context, req *protos.Void) (*protos.GatewayConfigs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMconfig not implemented")
}
func (*UnimplementedSouthboundConfiguratorServer) GetMconfigInternal(ctx context.Context, req *GetMconfigRequest) (*GetMconfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMconfigInternal not implemented")
}

func RegisterSouthboundConfiguratorServer(s *grpc.Server, srv SouthboundConfiguratorServer) {
	s.RegisterService(&_SouthboundConfigurator_serviceDesc, srv)
}

func _SouthboundConfigurator_GetMconfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(protos.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SouthboundConfiguratorServer).GetMconfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.configurator.SouthboundConfigurator/GetMconfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SouthboundConfiguratorServer).GetMconfig(ctx, req.(*protos.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _SouthboundConfigurator_GetMconfigInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMconfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SouthboundConfiguratorServer).GetMconfigInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/magma.orc8r.configurator.SouthboundConfigurator/GetMconfigInternal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SouthboundConfiguratorServer).GetMconfigInternal(ctx, req.(*GetMconfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SouthboundConfigurator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "magma.orc8r.configurator.SouthboundConfigurator",
	HandlerType: (*SouthboundConfiguratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMconfig",
			Handler:    _SouthboundConfigurator_GetMconfig_Handler,
		},
		{
			MethodName: "GetMconfigInternal",
			Handler:    _SouthboundConfigurator_GetMconfigInternal_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orc8r/cloud/go/services/configurator/protos/southbound.proto",
}
