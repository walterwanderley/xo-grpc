// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: category/v1/category.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v1 "northwind/api/typespb/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	CategoryService_CategoryByCategoryID_FullMethodName = "/category.v1.CategoryService/CategoryByCategoryID"
	CategoryService_Delete_FullMethodName               = "/category.v1.CategoryService/Delete"
	CategoryService_Insert_FullMethodName               = "/category.v1.CategoryService/Insert"
	CategoryService_Update_FullMethodName               = "/category.v1.CategoryService/Update"
	CategoryService_Upsert_FullMethodName               = "/category.v1.CategoryService/Upsert"
)

// CategoryServiceClient is the client API for CategoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CategoryServiceClient interface {
	CategoryByCategoryID(ctx context.Context, in *CategoryByCategoryIDRequest, opts ...grpc.CallOption) (*v1.Category, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type categoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCategoryServiceClient(cc grpc.ClientConnInterface) CategoryServiceClient {
	return &categoryServiceClient{cc}
}

func (c *categoryServiceClient) CategoryByCategoryID(ctx context.Context, in *CategoryByCategoryIDRequest, opts ...grpc.CallOption) (*v1.Category, error) {
	out := new(v1.Category)
	err := c.cc.Invoke(ctx, CategoryService_CategoryByCategoryID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CategoryService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CategoryService_Insert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CategoryService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *categoryServiceClient) Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CategoryService_Upsert_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CategoryServiceServer is the server API for CategoryService service.
// All implementations must embed UnimplementedCategoryServiceServer
// for forward compatibility
type CategoryServiceServer interface {
	CategoryByCategoryID(context.Context, *CategoryByCategoryIDRequest) (*v1.Category, error)
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	Insert(context.Context, *InsertRequest) (*emptypb.Empty, error)
	Update(context.Context, *UpdateRequest) (*emptypb.Empty, error)
	Upsert(context.Context, *UpsertRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCategoryServiceServer()
}

// UnimplementedCategoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCategoryServiceServer struct {
}

func (UnimplementedCategoryServiceServer) CategoryByCategoryID(context.Context, *CategoryByCategoryIDRequest) (*v1.Category, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CategoryByCategoryID not implemented")
}
func (UnimplementedCategoryServiceServer) Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCategoryServiceServer) Insert(context.Context, *InsertRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedCategoryServiceServer) Update(context.Context, *UpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCategoryServiceServer) Upsert(context.Context, *UpsertRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (UnimplementedCategoryServiceServer) mustEmbedUnimplementedCategoryServiceServer() {}

// UnsafeCategoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CategoryServiceServer will
// result in compilation errors.
type UnsafeCategoryServiceServer interface {
	mustEmbedUnimplementedCategoryServiceServer()
}

func RegisterCategoryServiceServer(s grpc.ServiceRegistrar, srv CategoryServiceServer) {
	s.RegisterService(&CategoryService_ServiceDesc, srv)
}

func _CategoryService_CategoryByCategoryID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CategoryByCategoryIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).CategoryByCategoryID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_CategoryByCategoryID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).CategoryByCategoryID(ctx, req.(*CategoryByCategoryIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_Insert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Insert(ctx, req.(*InsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CategoryService_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CategoryServiceServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CategoryService_Upsert_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CategoryServiceServer).Upsert(ctx, req.(*UpsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CategoryService_ServiceDesc is the grpc.ServiceDesc for CategoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CategoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "category.v1.CategoryService",
	HandlerType: (*CategoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CategoryByCategoryID",
			Handler:    _CategoryService_CategoryByCategoryID_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CategoryService_Delete_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _CategoryService_Insert_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CategoryService_Update_Handler,
		},
		{
			MethodName: "Upsert",
			Handler:    _CategoryService_Upsert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "category/v1/category.proto",
}
