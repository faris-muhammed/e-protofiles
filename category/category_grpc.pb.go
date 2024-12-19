// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: category.proto

package category

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
	CategoryService_AddCategory_FullMethodName    = "/category.CategoryService/AddCategory"
	CategoryService_EditCategory_FullMethodName   = "/category.CategoryService/EditCategory"
	CategoryService_DeleteCategory_FullMethodName = "/category.CategoryService/DeleteCategory"
	CategoryService_GetCategories_FullMethodName  = "/category.CategoryService/GetCategories"
)

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryServiceClient interface {
	AddCategory(ctx context.Context, in *AddCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	EditCategory(ctx context.Context, in *EditCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error)
	GetCategories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CategoryListResponse, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) AddCategory(ctx context.Context, in *AddCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, CategoryService_AddCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) EditCategory(ctx context.Context, in *EditCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, CategoryService_EditCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*CategoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryResponse)
	err := c.cc.Invoke(ctx, CategoryService_DeleteCategory_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) GetCategories(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CategoryListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CategoryListResponse)
	err := c.cc.Invoke(ctx, CategoryService_GetCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServiceServer is the server API for CategoryService service.
// All implementations must embed UnimplementedCategoryServiceServer
// for forward compatibility.
type CategoryServiceServer interface {
	AddCategory(context.Context, *AddCategoryRequest) (*CategoryResponse, error)
	EditCategory(context.Context, *EditCategoryRequest) (*CategoryResponse, error)
	DeleteCategory(context.Context, *DeleteCategoryRequest) (*CategoryResponse, error)
	GetCategories(context.Context, *Empty) (*CategoryListResponse, error)
	mustEmbedUnimplementedCategoryServiceServer()
}

// UnimplementedCategoryServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCategoryServiceServer struct{}

func (UnimplementedCategoryServiceServer) AddCategory(context.Context, *AddCategoryRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCategory not implemented")
}
func (UnimplementedCategoryServiceServer) EditCategory(context.Context, *EditCategoryRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditCategory not implemented")
}
func (UnimplementedCategoryServiceServer) DeleteCategory(context.Context, *DeleteCategoryRequest) (*CategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedCategoryServiceServer) GetCategories(context.Context, *Empty) (*CategoryListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (UnimplementedCategoryServiceServer) mustEmbedUnimplementedCategoryServiceServer() {}
func (UnimplementedCategoryServiceServer) testEmbeddedByValue()                         {}

// UnsafeCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServiceServer will
// result in compilation errors.
type UnsafeCategoryServiceServer interface {
	mustEmbedUnimplementedCategoryServiceServer()
}

func RegisterCategoryServiceServer(s grpc.ServiceRegistrar, srv CategoryServiceServer) {
	// If the following call pancis, it indicates UnimplementedCategoryServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CategoryService_ServiceDesc, srv)
}

func _CategoryService_AddCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).AddCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_AddCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).AddCategory(ctx, req.(*AddCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_EditCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).EditCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_EditCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).EditCategory(ctx, req.(*EditCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_DeleteCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).DeleteCategory(ctx, req.(*DeleteCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_GetCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).GetCategories(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryService_ServiceDesc is the grpc.ServiceDesc for CategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "category.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCategory",
			Handler:    _CategoryService_AddCategory_Handler,
		},
		{
			MethodName: "EditCategory",
			Handler:    _CategoryService_EditCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _CategoryService_DeleteCategory_Handler,
		},
		{
			MethodName: "GetCategories",
			Handler:    _CategoryService_GetCategories_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "category.proto",
}
