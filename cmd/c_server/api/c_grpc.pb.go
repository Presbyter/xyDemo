// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CServiceClient is the client API for CService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CServiceClient interface {
	Forward(ctx context.Context, opts ...grpc.CallOption) (CService_ForwardClient, error)
}

type cServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCServiceClient(cc grpc.ClientConnInterface) CServiceClient {
	return &cServiceClient{cc}
}

func (c *cServiceClient) Forward(ctx context.Context, opts ...grpc.CallOption) (CService_ForwardClient, error) {
	stream, err := c.cc.NewStream(ctx, &CService_ServiceDesc.Streams[0], "/main.CService/Forward", opts...)
	if err != nil {
		return nil, err
	}
	x := &cServiceForwardClient{stream}
	return x, nil
}

type CService_ForwardClient interface {
	Send(*ForwardReq) error
	Recv() (*ForwardResp, error)
	grpc.ClientStream
}

type cServiceForwardClient struct {
	grpc.ClientStream
}

func (x *cServiceForwardClient) Send(m *ForwardReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *cServiceForwardClient) Recv() (*ForwardResp, error) {
	m := new(ForwardResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CServiceServer is the server API for CService service.
// All implementations must embed UnimplementedCServiceServer
// for forward compatibility
type CServiceServer interface {
	Forward(CService_ForwardServer) error
	mustEmbedUnimplementedCServiceServer()
}

// UnimplementedCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCServiceServer struct {
}

func (UnimplementedCServiceServer) Forward(CService_ForwardServer) error {
	return status.Errorf(codes.Unimplemented, "method Forward not implemented")
}
func (UnimplementedCServiceServer) mustEmbedUnimplementedCServiceServer() {}

// UnsafeCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CServiceServer will
// result in compilation errors.
type UnsafeCServiceServer interface {
	mustEmbedUnimplementedCServiceServer()
}

func RegisterCServiceServer(s grpc.ServiceRegistrar, srv CServiceServer) {
	s.RegisterService(&CService_ServiceDesc, srv)
}

func _CService_Forward_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CServiceServer).Forward(&cServiceForwardServer{stream})
}

type CService_ForwardServer interface {
	Send(*ForwardResp) error
	Recv() (*ForwardReq, error)
	grpc.ServerStream
}

type cServiceForwardServer struct {
	grpc.ServerStream
}

func (x *cServiceForwardServer) Send(m *ForwardResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *cServiceForwardServer) Recv() (*ForwardReq, error) {
	m := new(ForwardReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CService_ServiceDesc is the grpc.ServiceDesc for CService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "main.CService",
	HandlerType: (*CServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Forward",
			Handler:       _CService_Forward_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "api/c.proto",
}
