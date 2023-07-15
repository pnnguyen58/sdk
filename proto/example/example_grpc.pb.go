// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: proto/example/example.proto

package example

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

const (
	ExampleService_CreateExample_FullMethodName = "/go_temporal.ExampleService/CreateExample"
	ExampleService_GetExample_FullMethodName    = "/go_temporal.ExampleService/GetExample"
	ExampleService_GetAllExample_FullMethodName = "/go_temporal.ExampleService/GetAllExample"
	ExampleService_UpdateExample_FullMethodName = "/go_temporal.ExampleService/UpdateExample"
	ExampleService_DeleteExample_FullMethodName = "/go_temporal.ExampleService/DeleteExample"
)

// ExampleServiceClient is the client API for ExampleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExampleServiceClient interface {
	// User request group
	CreateExample(ctx context.Context, in *ExampleCreateRequest, opts ...grpc.CallOption) (*ExampleCreateResponse, error)
	GetExample(ctx context.Context, in *ExampleGetRequest, opts ...grpc.CallOption) (*ExampleGetResponse, error)
	GetAllExample(ctx context.Context, in *ExampleGetAllRequest, opts ...grpc.CallOption) (*ExampleGetAllResponse, error)
	UpdateExample(ctx context.Context, in *ExampleUpdateRequest, opts ...grpc.CallOption) (*ExampleUpdateResponse, error)
	DeleteExample(ctx context.Context, in *ExampleDeleteRequest, opts ...grpc.CallOption) (*ExampleDeleteResponse, error)
}

type exampleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExampleServiceClient(cc grpc.ClientConnInterface) ExampleServiceClient {
	return &exampleServiceClient{cc}
}

func (c *exampleServiceClient) CreateExample(ctx context.Context, in *ExampleCreateRequest, opts ...grpc.CallOption) (*ExampleCreateResponse, error) {
	out := new(ExampleCreateResponse)
	err := c.cc.Invoke(ctx, ExampleService_CreateExample_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) GetExample(ctx context.Context, in *ExampleGetRequest, opts ...grpc.CallOption) (*ExampleGetResponse, error) {
	out := new(ExampleGetResponse)
	err := c.cc.Invoke(ctx, ExampleService_GetExample_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) GetAllExample(ctx context.Context, in *ExampleGetAllRequest, opts ...grpc.CallOption) (*ExampleGetAllResponse, error) {
	out := new(ExampleGetAllResponse)
	err := c.cc.Invoke(ctx, ExampleService_GetAllExample_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) UpdateExample(ctx context.Context, in *ExampleUpdateRequest, opts ...grpc.CallOption) (*ExampleUpdateResponse, error) {
	out := new(ExampleUpdateResponse)
	err := c.cc.Invoke(ctx, ExampleService_UpdateExample_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleServiceClient) DeleteExample(ctx context.Context, in *ExampleDeleteRequest, opts ...grpc.CallOption) (*ExampleDeleteResponse, error) {
	out := new(ExampleDeleteResponse)
	err := c.cc.Invoke(ctx, ExampleService_DeleteExample_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleServiceServer is the server API for ExampleService service.
// All implementations must embed UnimplementedExampleServiceServer
// for forward compatibility
type ExampleServiceServer interface {
	// User request group
	CreateExample(context.Context, *ExampleCreateRequest) (*ExampleCreateResponse, error)
	GetExample(context.Context, *ExampleGetRequest) (*ExampleGetResponse, error)
	GetAllExample(context.Context, *ExampleGetAllRequest) (*ExampleGetAllResponse, error)
	UpdateExample(context.Context, *ExampleUpdateRequest) (*ExampleUpdateResponse, error)
	DeleteExample(context.Context, *ExampleDeleteRequest) (*ExampleDeleteResponse, error)
	mustEmbedUnimplementedExampleServiceServer()
}

// UnimplementedExampleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExampleServiceServer struct {
}

func (UnimplementedExampleServiceServer) CreateExample(context.Context, *ExampleCreateRequest) (*ExampleCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExample not implemented")
}
func (UnimplementedExampleServiceServer) GetExample(context.Context, *ExampleGetRequest) (*ExampleGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExample not implemented")
}
func (UnimplementedExampleServiceServer) GetAllExample(context.Context, *ExampleGetAllRequest) (*ExampleGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllExample not implemented")
}
func (UnimplementedExampleServiceServer) UpdateExample(context.Context, *ExampleUpdateRequest) (*ExampleUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExample not implemented")
}
func (UnimplementedExampleServiceServer) DeleteExample(context.Context, *ExampleDeleteRequest) (*ExampleDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExample not implemented")
}
func (UnimplementedExampleServiceServer) mustEmbedUnimplementedExampleServiceServer() {}

// UnsafeExampleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExampleServiceServer will
// result in compilation errors.
type UnsafeExampleServiceServer interface {
	mustEmbedUnimplementedExampleServiceServer()
}

func RegisterExampleServiceServer(s grpc.ServiceRegistrar, srv ExampleServiceServer) {
	s.RegisterService(&ExampleService_ServiceDesc, srv)
}

func _ExampleService_CreateExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).CreateExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_CreateExample_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).CreateExample(ctx, req.(*ExampleCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_GetExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).GetExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_GetExample_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).GetExample(ctx, req.(*ExampleGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_GetAllExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleGetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).GetAllExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_GetAllExample_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).GetAllExample(ctx, req.(*ExampleGetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_UpdateExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).UpdateExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_UpdateExample_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).UpdateExample(ctx, req.(*ExampleUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExampleService_DeleteExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExampleDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleServiceServer).DeleteExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExampleService_DeleteExample_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleServiceServer).DeleteExample(ctx, req.(*ExampleDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExampleService_ServiceDesc is the grpc.ServiceDesc for ExampleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExampleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "go_temporal.ExampleService",
	HandlerType: (*ExampleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExample",
			Handler:    _ExampleService_CreateExample_Handler,
		},
		{
			MethodName: "GetExample",
			Handler:    _ExampleService_GetExample_Handler,
		},
		{
			MethodName: "GetAllExample",
			Handler:    _ExampleService_GetAllExample_Handler,
		},
		{
			MethodName: "UpdateExample",
			Handler:    _ExampleService_UpdateExample_Handler,
		},
		{
			MethodName: "DeleteExample",
			Handler:    _ExampleService_DeleteExample_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/example/example.proto",
}