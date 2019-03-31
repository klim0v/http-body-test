// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import httpbody "google.golang.org/genproto/googleapis/api/httpbody"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PongResponse struct {
	Pong                 string   `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PongResponse) Reset()         { *m = PongResponse{} }
func (m *PongResponse) String() string { return proto.CompactTextString(m) }
func (*PongResponse) ProtoMessage()    {}
func (*PongResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_service_db106e9cdca9462f, []int{0}
}
func (m *PongResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongResponse.Unmarshal(m, b)
}
func (m *PongResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongResponse.Marshal(b, m, deterministic)
}
func (dst *PongResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongResponse.Merge(dst, src)
}
func (m *PongResponse) XXX_Size() int {
	return xxx_messageInfo_PongResponse.Size(m)
}
func (m *PongResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PongResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PongResponse proto.InternalMessageInfo

func (m *PongResponse) GetPong() string {
	if m != nil {
		return m.Pong
	}
	return ""
}

func init() {
	proto.RegisterType((*PongResponse)(nil), "pb.PongResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	GetResource(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error)
	Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PongResponse, error)
}

type serviceClient struct {
	cc *grpc.ClientConn
}

func NewServiceClient(cc *grpc.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) GetResource(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*httpbody.HttpBody, error) {
	out := new(httpbody.HttpBody)
	err := c.cc.Invoke(ctx, "/pb.Service/GetResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) Ping(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PongResponse, error) {
	out := new(PongResponse)
	err := c.cc.Invoke(ctx, "/pb.Service/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	GetResource(context.Context, *empty.Empty) (*httpbody.HttpBody, error)
	Ping(context.Context, *empty.Empty) (*PongResponse, error)
}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_GetResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Service/GetResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetResource(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Service/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).Ping(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetResource",
			Handler:    _Service_GetResource_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _Service_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_db106e9cdca9462f) }

var fileDescriptor_service_db106e9cdca9462f = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8e, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x49, 0xa9, 0x96, 0x6e, 0x2d, 0xe8, 0x22, 0xa2, 0xd1, 0x83, 0xe4, 0xe4, 0x69, 0x16,
	0xf4, 0x01, 0x04, 0x41, 0xf4, 0x24, 0x25, 0x3e, 0x41, 0x36, 0x9d, 0xae, 0x0b, 0x75, 0x66, 0xc8,
	0x4e, 0x0b, 0xbd, 0xfa, 0x0a, 0xde, 0x7c, 0x2d, 0x5f, 0xc1, 0x07, 0x11, 0x93, 0x14, 0x72, 0xf1,
	0x36, 0xf3, 0x7f, 0xfc, 0x1f, 0xbf, 0x99, 0x27, 0x6c, 0xb6, 0xb1, 0x46, 0x90, 0x86, 0x95, 0xed,
	0x48, 0x7c, 0x7e, 0x11, 0x98, 0xc3, 0x1a, 0x5d, 0x25, 0xd1, 0xbd, 0xa9, 0x8a, 0xe7, 0xe5, 0xae,
	0xc3, 0xf9, 0xd5, 0x00, 0x55, 0x44, 0xac, 0x95, 0x46, 0xa6, 0xd4, 0xd3, 0xcb, 0x9e, 0xb6, 0x9f,
	0xdf, 0xac, 0x1c, 0xbe, 0x8b, 0xf6, 0xd5, 0xa2, 0x30, 0x47, 0x0b, 0xa6, 0x50, 0x62, 0x12, 0xa6,
	0x84, 0xd6, 0x9a, 0xb1, 0x30, 0x85, 0xf3, 0xec, 0x3a, 0xbb, 0x99, 0x96, 0xed, 0x7d, 0xfb, 0x95,
	0x99, 0xc9, 0x6b, 0xb7, 0xc7, 0xbe, 0x98, 0xd9, 0x13, 0x6a, 0x89, 0x89, 0x37, 0x4d, 0x8d, 0xf6,
	0x0c, 0x3a, 0x39, 0xec, 0xe5, 0xf0, 0xf8, 0x27, 0xcf, 0x4f, 0xf7, 0x79, 0x25, 0x11, 0x9e, 0x55,
	0xe5, 0x81, 0x97, 0xbb, 0xe2, 0xe4, 0xe3, 0xfb, 0xe7, 0x73, 0x34, 0xb3, 0x53, 0xb7, 0x8a, 0x6b,
	0x84, 0x3a, 0x6d, 0xed, 0xbd, 0x19, 0x2f, 0x22, 0x85, 0x7f, 0x45, 0xc7, 0x20, 0x1e, 0x86, 0x0b,
	0x8b, 0x79, 0x2b, 0x99, 0xd8, 0x03, 0x27, 0x91, 0x82, 0x3f, 0x6c, 0x0b, 0x77, 0xbf, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xa1, 0x5d, 0x06, 0x58, 0x32, 0x01, 0x00, 0x00,
}
