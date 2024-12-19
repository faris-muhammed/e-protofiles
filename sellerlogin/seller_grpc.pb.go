// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: seller.proto

package sellerlogin

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
	SellerService_Signup_FullMethodName = "/seller.SellerService/Signup"
	SellerService_Login_FullMethodName  = "/seller.SellerService/Login"
)

// SellerServiceClient is the client API for SellerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SellerServiceClient interface {
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type sellerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSellerServiceClient(cc grpc.ClientConnInterface) SellerServiceClient {
	return &sellerServiceClient{cc}
}

func (c *sellerServiceClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, SellerService_Signup_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, SellerService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SellerServiceServer is the server API for SellerService service.
// All implementations must embed UnimplementedSellerServiceServer
// for forward compatibility.
type SellerServiceServer interface {
	Signup(context.Context, *SignupRequest) (*SignupResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	mustEmbedUnimplementedSellerServiceServer()
}

// UnimplementedSellerServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSellerServiceServer struct{}

func (UnimplementedSellerServiceServer) Signup(context.Context, *SignupRequest) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedSellerServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSellerServiceServer) mustEmbedUnimplementedSellerServiceServer() {}
func (UnimplementedSellerServiceServer) testEmbeddedByValue()                       {}

// UnsafeSellerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SellerServiceServer will
// result in compilation errors.
type UnsafeSellerServiceServer interface {
	mustEmbedUnimplementedSellerServiceServer()
}

func RegisterSellerServiceServer(s grpc.ServiceRegistrar, srv SellerServiceServer) {
	// If the following call pancis, it indicates UnimplementedSellerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SellerService_ServiceDesc, srv)
}

func _SellerService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SellerService_Signup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerServiceServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SellerService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SellerService_ServiceDesc is the grpc.ServiceDesc for SellerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SellerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "seller.SellerService",
	HandlerType: (*SellerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _SellerService_Signup_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _SellerService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "seller.proto",
}
