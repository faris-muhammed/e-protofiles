// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: address.proto

package address

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserService_AddUserAddress_FullMethodName    = "/address.UserService/AddUserAddress"
	UserService_GetUserAddresses_FullMethodName  = "/address.UserService/GetUserAddresses"
	UserService_EditUserAddress_FullMethodName   = "/address.UserService/EditUserAddress"
	UserService_RemoveUserAddress_FullMethodName = "/address.UserService/RemoveUserAddress"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	AddUserAddress(ctx context.Context, in *AddUserAddressRequest, opts ...grpc.CallOption) (*AddUserAddressResponse, error)
	GetUserAddresses(ctx context.Context, in *GetUserAddressesRequest, opts ...grpc.CallOption) (*GetUserAddressesResponse, error)
	EditUserAddress(ctx context.Context, in *EditUserAddressRequest, opts ...grpc.CallOption) (*EditUserAddressResponse, error)
	RemoveUserAddress(ctx context.Context, in *RemoveUserAddressRequest, opts ...grpc.CallOption) (*RemoveUserAddressResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) AddUserAddress(ctx context.Context, in *AddUserAddressRequest, opts ...grpc.CallOption) (*AddUserAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddUserAddressResponse)
	err := c.cc.Invoke(ctx, UserService_AddUserAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserAddresses(ctx context.Context, in *GetUserAddressesRequest, opts ...grpc.CallOption) (*GetUserAddressesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserAddressesResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserAddresses_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) EditUserAddress(ctx context.Context, in *EditUserAddressRequest, opts ...grpc.CallOption) (*EditUserAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(EditUserAddressResponse)
	err := c.cc.Invoke(ctx, UserService_EditUserAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) RemoveUserAddress(ctx context.Context, in *RemoveUserAddressRequest, opts ...grpc.CallOption) (*RemoveUserAddressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveUserAddressResponse)
	err := c.cc.Invoke(ctx, UserService_RemoveUserAddress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	AddUserAddress(context.Context, *AddUserAddressRequest) (*AddUserAddressResponse, error)
	GetUserAddresses(context.Context, *GetUserAddressesRequest) (*GetUserAddressesResponse, error)
	EditUserAddress(context.Context, *EditUserAddressRequest) (*EditUserAddressResponse, error)
	RemoveUserAddress(context.Context, *RemoveUserAddressRequest) (*RemoveUserAddressResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) AddUserAddress(context.Context, *AddUserAddressRequest) (*AddUserAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserAddress not implemented")
}
func (UnimplementedUserServiceServer) GetUserAddresses(context.Context, *GetUserAddressesRequest) (*GetUserAddressesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserAddresses not implemented")
}
func (UnimplementedUserServiceServer) EditUserAddress(context.Context, *EditUserAddressRequest) (*EditUserAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUserAddress not implemented")
}
func (UnimplementedUserServiceServer) RemoveUserAddress(context.Context, *RemoveUserAddressRequest) (*RemoveUserAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveUserAddress not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_AddUserAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddUserAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_AddUserAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddUserAddress(ctx, req.(*AddUserAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserAddresses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserAddressesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserAddresses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserAddresses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserAddresses(ctx, req.(*GetUserAddressesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_EditUserAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditUserAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EditUserAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_EditUserAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EditUserAddress(ctx, req.(*EditUserAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_RemoveUserAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveUserAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).RemoveUserAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_RemoveUserAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).RemoveUserAddress(ctx, req.(*RemoveUserAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "address.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUserAddress",
			Handler:    _UserService_AddUserAddress_Handler,
		},
		{
			MethodName: "GetUserAddresses",
			Handler:    _UserService_GetUserAddresses_Handler,
		},
		{
			MethodName: "EditUserAddress",
			Handler:    _UserService_EditUserAddress_Handler,
		},
		{
			MethodName: "RemoveUserAddress",
			Handler:    _UserService_RemoveUserAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "address.proto",
}
