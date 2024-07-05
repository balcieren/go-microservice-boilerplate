// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.1
// source: pkg/proto/owner.proto

package proto

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

// OwnerServiceClient is the client API for OwnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OwnerServiceClient interface {
	ListOwners(ctx context.Context, in *ListOwnersRequest, opts ...grpc.CallOption) (*ListOwnersResponse, error)
	GetOwner(ctx context.Context, in *GetOwnerRequest, opts ...grpc.CallOption) (*GetOwnerResponse, error)
	CreateOwner(ctx context.Context, in *CreateOwnerRequest, opts ...grpc.CallOption) (*CreateOwnerResponse, error)
	UpdateOwner(ctx context.Context, in *UpdateOwnerRequest, opts ...grpc.CallOption) (*UpdateOwnerResponse, error)
	DeleteOwner(ctx context.Context, in *DeleteOwnerRequest, opts ...grpc.CallOption) (*DeleteOwnerResponse, error)
}

type ownerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOwnerServiceClient(cc grpc.ClientConnInterface) OwnerServiceClient {
	return &ownerServiceClient{cc}
}

func (c *ownerServiceClient) ListOwners(ctx context.Context, in *ListOwnersRequest, opts ...grpc.CallOption) (*ListOwnersResponse, error) {
	out := new(ListOwnersResponse)
	err := c.cc.Invoke(ctx, "/OwnerService/ListOwners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownerServiceClient) GetOwner(ctx context.Context, in *GetOwnerRequest, opts ...grpc.CallOption) (*GetOwnerResponse, error) {
	out := new(GetOwnerResponse)
	err := c.cc.Invoke(ctx, "/OwnerService/GetOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownerServiceClient) CreateOwner(ctx context.Context, in *CreateOwnerRequest, opts ...grpc.CallOption) (*CreateOwnerResponse, error) {
	out := new(CreateOwnerResponse)
	err := c.cc.Invoke(ctx, "/OwnerService/CreateOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownerServiceClient) UpdateOwner(ctx context.Context, in *UpdateOwnerRequest, opts ...grpc.CallOption) (*UpdateOwnerResponse, error) {
	out := new(UpdateOwnerResponse)
	err := c.cc.Invoke(ctx, "/OwnerService/UpdateOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ownerServiceClient) DeleteOwner(ctx context.Context, in *DeleteOwnerRequest, opts ...grpc.CallOption) (*DeleteOwnerResponse, error) {
	out := new(DeleteOwnerResponse)
	err := c.cc.Invoke(ctx, "/OwnerService/DeleteOwner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OwnerServiceServer is the server API for OwnerService service.
// All implementations must embed UnimplementedOwnerServiceServer
// for forward compatibility
type OwnerServiceServer interface {
	ListOwners(context.Context, *ListOwnersRequest) (*ListOwnersResponse, error)
	GetOwner(context.Context, *GetOwnerRequest) (*GetOwnerResponse, error)
	CreateOwner(context.Context, *CreateOwnerRequest) (*CreateOwnerResponse, error)
	UpdateOwner(context.Context, *UpdateOwnerRequest) (*UpdateOwnerResponse, error)
	DeleteOwner(context.Context, *DeleteOwnerRequest) (*DeleteOwnerResponse, error)
	mustEmbedUnimplementedOwnerServiceServer()
}

// UnimplementedOwnerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOwnerServiceServer struct {
}

func (UnimplementedOwnerServiceServer) ListOwners(context.Context, *ListOwnersRequest) (*ListOwnersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOwners not implemented")
}
func (UnimplementedOwnerServiceServer) GetOwner(context.Context, *GetOwnerRequest) (*GetOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOwner not implemented")
}
func (UnimplementedOwnerServiceServer) CreateOwner(context.Context, *CreateOwnerRequest) (*CreateOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOwner not implemented")
}
func (UnimplementedOwnerServiceServer) UpdateOwner(context.Context, *UpdateOwnerRequest) (*UpdateOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOwner not implemented")
}
func (UnimplementedOwnerServiceServer) DeleteOwner(context.Context, *DeleteOwnerRequest) (*DeleteOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOwner not implemented")
}
func (UnimplementedOwnerServiceServer) mustEmbedUnimplementedOwnerServiceServer() {}

// UnsafeOwnerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OwnerServiceServer will
// result in compilation errors.
type UnsafeOwnerServiceServer interface {
	mustEmbedUnimplementedOwnerServiceServer()
}

func RegisterOwnerServiceServer(s grpc.ServiceRegistrar, srv OwnerServiceServer) {
	s.RegisterService(&OwnerService_ServiceDesc, srv)
}

func _OwnerService_ListOwners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOwnersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnerServiceServer).ListOwners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OwnerService/ListOwners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnerServiceServer).ListOwners(ctx, req.(*ListOwnersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnerService_GetOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnerServiceServer).GetOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OwnerService/GetOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnerServiceServer).GetOwner(ctx, req.(*GetOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnerService_CreateOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnerServiceServer).CreateOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OwnerService/CreateOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnerServiceServer).CreateOwner(ctx, req.(*CreateOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnerService_UpdateOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnerServiceServer).UpdateOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OwnerService/UpdateOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnerServiceServer).UpdateOwner(ctx, req.(*UpdateOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OwnerService_DeleteOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OwnerServiceServer).DeleteOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OwnerService/DeleteOwner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OwnerServiceServer).DeleteOwner(ctx, req.(*DeleteOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OwnerService_ServiceDesc is the grpc.ServiceDesc for OwnerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OwnerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OwnerService",
	HandlerType: (*OwnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListOwners",
			Handler:    _OwnerService_ListOwners_Handler,
		},
		{
			MethodName: "GetOwner",
			Handler:    _OwnerService_GetOwner_Handler,
		},
		{
			MethodName: "CreateOwner",
			Handler:    _OwnerService_CreateOwner_Handler,
		},
		{
			MethodName: "UpdateOwner",
			Handler:    _OwnerService_UpdateOwner_Handler,
		},
		{
			MethodName: "DeleteOwner",
			Handler:    _OwnerService_DeleteOwner_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/owner.proto",
}