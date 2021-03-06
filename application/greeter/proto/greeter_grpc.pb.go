// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package greeter

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

// GreeterServiceClient is the client API for GreeterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterServiceClient interface {
	CreateGreeter(ctx context.Context, in *CreateGreeterRequest, opts ...grpc.CallOption) (*CreateGreeterResponse, error)
	GetGreeterById(ctx context.Context, in *GetGreeterByIdRequest, opts ...grpc.CallOption) (*GetGreeterByIdResponse, error)
	GetGreeterList(ctx context.Context, in *GetGreeterListRequest, opts ...grpc.CallOption) (*GetGreeterListResponse, error)
	UpdateGreeterStatus(ctx context.Context, in *UpdateGreeterStatusRequest, opts ...grpc.CallOption) (*UpdateGreeterStatusResponse, error)
	UpdateGreeterCount(ctx context.Context, in *UpdateGreeterCountRequest, opts ...grpc.CallOption) (*UpdateGreeterCountResponse, error)
	DeleteGreeterById(ctx context.Context, in *DeleteGreeterByIdRequest, opts ...grpc.CallOption) (*DeleteGreeterByIdResponse, error)
	GetGreeterListByStream(ctx context.Context, opts ...grpc.CallOption) (GreeterService_GetGreeterListByStreamClient, error)
}

type greeterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterServiceClient(cc grpc.ClientConnInterface) GreeterServiceClient {
	return &greeterServiceClient{cc}
}

func (c *greeterServiceClient) CreateGreeter(ctx context.Context, in *CreateGreeterRequest, opts ...grpc.CallOption) (*CreateGreeterResponse, error) {
	out := new(CreateGreeterResponse)
	err := c.cc.Invoke(ctx, "/greeter.GreeterService/CreateGreeter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterServiceClient) GetGreeterById(ctx context.Context, in *GetGreeterByIdRequest, opts ...grpc.CallOption) (*GetGreeterByIdResponse, error) {
	out := new(GetGreeterByIdResponse)
	err := c.cc.Invoke(ctx, "/greeter.GreeterService/GetGreeterById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterServiceClient) GetGreeterList(ctx context.Context, in *GetGreeterListRequest, opts ...grpc.CallOption) (*GetGreeterListResponse, error) {
	out := new(GetGreeterListResponse)
	err := c.cc.Invoke(ctx, "/greeter.GreeterService/GetGreeterList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterServiceClient) UpdateGreeterStatus(ctx context.Context, in *UpdateGreeterStatusRequest, opts ...grpc.CallOption) (*UpdateGreeterStatusResponse, error) {
	out := new(UpdateGreeterStatusResponse)
	err := c.cc.Invoke(ctx, "/greeter.GreeterService/UpdateGreeterStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterServiceClient) UpdateGreeterCount(ctx context.Context, in *UpdateGreeterCountRequest, opts ...grpc.CallOption) (*UpdateGreeterCountResponse, error) {
	out := new(UpdateGreeterCountResponse)
	err := c.cc.Invoke(ctx, "/greeter.GreeterService/UpdateGreeterCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterServiceClient) DeleteGreeterById(ctx context.Context, in *DeleteGreeterByIdRequest, opts ...grpc.CallOption) (*DeleteGreeterByIdResponse, error) {
	out := new(DeleteGreeterByIdResponse)
	err := c.cc.Invoke(ctx, "/greeter.GreeterService/DeleteGreeterById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterServiceClient) GetGreeterListByStream(ctx context.Context, opts ...grpc.CallOption) (GreeterService_GetGreeterListByStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreeterService_ServiceDesc.Streams[0], "/greeter.GreeterService/GetGreeterListByStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterServiceGetGreeterListByStreamClient{stream}
	return x, nil
}

type GreeterService_GetGreeterListByStreamClient interface {
	Send(*GetGreeterListByStreamRequest) error
	Recv() (*GetGreeterListByStreamResponse, error)
	grpc.ClientStream
}

type greeterServiceGetGreeterListByStreamClient struct {
	grpc.ClientStream
}

func (x *greeterServiceGetGreeterListByStreamClient) Send(m *GetGreeterListByStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterServiceGetGreeterListByStreamClient) Recv() (*GetGreeterListByStreamResponse, error) {
	m := new(GetGreeterListByStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServiceServer is the server API for GreeterService service.
// All implementations must embed UnimplementedGreeterServiceServer
// for forward compatibility
type GreeterServiceServer interface {
	CreateGreeter(context.Context, *CreateGreeterRequest) (*CreateGreeterResponse, error)
	GetGreeterById(context.Context, *GetGreeterByIdRequest) (*GetGreeterByIdResponse, error)
	GetGreeterList(context.Context, *GetGreeterListRequest) (*GetGreeterListResponse, error)
	UpdateGreeterStatus(context.Context, *UpdateGreeterStatusRequest) (*UpdateGreeterStatusResponse, error)
	UpdateGreeterCount(context.Context, *UpdateGreeterCountRequest) (*UpdateGreeterCountResponse, error)
	DeleteGreeterById(context.Context, *DeleteGreeterByIdRequest) (*DeleteGreeterByIdResponse, error)
	GetGreeterListByStream(GreeterService_GetGreeterListByStreamServer) error
	mustEmbedUnimplementedGreeterServiceServer()
}

// UnimplementedGreeterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServiceServer struct {
}

func (UnimplementedGreeterServiceServer) CreateGreeter(context.Context, *CreateGreeterRequest) (*CreateGreeterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGreeter not implemented")
}
func (UnimplementedGreeterServiceServer) GetGreeterById(context.Context, *GetGreeterByIdRequest) (*GetGreeterByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGreeterById not implemented")
}
func (UnimplementedGreeterServiceServer) GetGreeterList(context.Context, *GetGreeterListRequest) (*GetGreeterListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGreeterList not implemented")
}
func (UnimplementedGreeterServiceServer) UpdateGreeterStatus(context.Context, *UpdateGreeterStatusRequest) (*UpdateGreeterStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGreeterStatus not implemented")
}
func (UnimplementedGreeterServiceServer) UpdateGreeterCount(context.Context, *UpdateGreeterCountRequest) (*UpdateGreeterCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGreeterCount not implemented")
}
func (UnimplementedGreeterServiceServer) DeleteGreeterById(context.Context, *DeleteGreeterByIdRequest) (*DeleteGreeterByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGreeterById not implemented")
}
func (UnimplementedGreeterServiceServer) GetGreeterListByStream(GreeterService_GetGreeterListByStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetGreeterListByStream not implemented")
}
func (UnimplementedGreeterServiceServer) mustEmbedUnimplementedGreeterServiceServer() {}

// UnsafeGreeterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServiceServer will
// result in compilation errors.
type UnsafeGreeterServiceServer interface {
	mustEmbedUnimplementedGreeterServiceServer()
}

func RegisterGreeterServiceServer(s grpc.ServiceRegistrar, srv GreeterServiceServer) {
	s.RegisterService(&GreeterService_ServiceDesc, srv)
}

func _GreeterService_CreateGreeter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGreeterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).CreateGreeter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeter.GreeterService/CreateGreeter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).CreateGreeter(ctx, req.(*CreateGreeterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreeterService_GetGreeterById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGreeterByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).GetGreeterById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeter.GreeterService/GetGreeterById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).GetGreeterById(ctx, req.(*GetGreeterByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreeterService_GetGreeterList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGreeterListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).GetGreeterList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeter.GreeterService/GetGreeterList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).GetGreeterList(ctx, req.(*GetGreeterListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreeterService_UpdateGreeterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGreeterStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).UpdateGreeterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeter.GreeterService/UpdateGreeterStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).UpdateGreeterStatus(ctx, req.(*UpdateGreeterStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreeterService_UpdateGreeterCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGreeterCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).UpdateGreeterCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeter.GreeterService/UpdateGreeterCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).UpdateGreeterCount(ctx, req.(*UpdateGreeterCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreeterService_DeleteGreeterById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGreeterByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).DeleteGreeterById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greeter.GreeterService/DeleteGreeterById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).DeleteGreeterById(ctx, req.(*DeleteGreeterByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreeterService_GetGreeterListByStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServiceServer).GetGreeterListByStream(&greeterServiceGetGreeterListByStreamServer{stream})
}

type GreeterService_GetGreeterListByStreamServer interface {
	Send(*GetGreeterListByStreamResponse) error
	Recv() (*GetGreeterListByStreamRequest, error)
	grpc.ServerStream
}

type greeterServiceGetGreeterListByStreamServer struct {
	grpc.ServerStream
}

func (x *greeterServiceGetGreeterListByStreamServer) Send(m *GetGreeterListByStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterServiceGetGreeterListByStreamServer) Recv() (*GetGreeterListByStreamRequest, error) {
	m := new(GetGreeterListByStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterService_ServiceDesc is the grpc.ServiceDesc for GreeterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreeterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greeter.GreeterService",
	HandlerType: (*GreeterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGreeter",
			Handler:    _GreeterService_CreateGreeter_Handler,
		},
		{
			MethodName: "GetGreeterById",
			Handler:    _GreeterService_GetGreeterById_Handler,
		},
		{
			MethodName: "GetGreeterList",
			Handler:    _GreeterService_GetGreeterList_Handler,
		},
		{
			MethodName: "UpdateGreeterStatus",
			Handler:    _GreeterService_UpdateGreeterStatus_Handler,
		},
		{
			MethodName: "UpdateGreeterCount",
			Handler:    _GreeterService_UpdateGreeterCount_Handler,
		},
		{
			MethodName: "DeleteGreeterById",
			Handler:    _GreeterService_DeleteGreeterById_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetGreeterListByStream",
			Handler:       _GreeterService_GetGreeterListByStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "greeter.proto",
}
