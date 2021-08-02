// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BifrostGrpcClient is the client API for BifrostGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BifrostGrpcClient interface {
	Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error)
}

type bifrostGrpcClient struct {
	cc *grpc.ClientConn
}

func NewBifrostGrpcClient(cc *grpc.ClientConn) BifrostGrpcClient {
	return &bifrostGrpcClient{cc}
}

func (c *bifrostGrpcClient) Ping(ctx context.Context, in *PingReq, opts ...grpc.CallOption) (*PingResp, error) {
	out := new(PingResp)
	err := c.cc.Invoke(ctx, "/proto.BifrostGrpc/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BifrostGrpcServer is the server API for BifrostGrpc service.
type BifrostGrpcServer interface {
	Ping(context.Context, *PingReq) (*PingResp, error)
}

func RegisterBifrostGrpcServer(s *grpc.Server, srv BifrostGrpcServer) {
	s.RegisterService(&_BifrostGrpc_serviceDesc, srv)
}

func _BifrostGrpc_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BifrostGrpcServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.BifrostGrpc/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BifrostGrpcServer).Ping(ctx, req.(*PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _BifrostGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.BifrostGrpc",
	HandlerType: (*BifrostGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _BifrostGrpc_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_service_84f1b55c1e33a251) }

var fileDescriptor_service_84f1b55c1e33a251 = []byte{
	// 94 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x5c, 0x05,
	0x99, 0x79, 0xe9, 0x10, 0x21, 0x23, 0x33, 0x2e, 0x6e, 0xa7, 0xcc, 0xb4, 0xa2, 0xfc, 0xe2, 0x12,
	0xf7, 0xa2, 0x82, 0x64, 0x21, 0x75, 0x2e, 0x96, 0x80, 0xcc, 0xbc, 0x74, 0x21, 0x3e, 0x88, 0xb4,
	0x1e, 0x88, 0x13, 0x94, 0x5a, 0x28, 0xc5, 0x8f, 0xc2, 0x2f, 0x2e, 0x48, 0x62, 0x03, 0xf3, 0x8d,
	0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xda, 0xd4, 0x47, 0x62, 0x00, 0x00, 0x00,
}
