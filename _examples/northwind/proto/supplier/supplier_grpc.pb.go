// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package supplier

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	typespb "northwind/proto/typespb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SupplierServiceClient is the client API for SupplierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupplierServiceClient interface {
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SupplierBySupplierID(ctx context.Context, in *SupplierBySupplierIDRequest, opts ...grpc.CallOption) (*typespb.Supplier, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type supplierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSupplierServiceClient(cc grpc.ClientConnInterface) SupplierServiceClient {
	return &supplierServiceClient{cc}
}

func (c *supplierServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/supplier.SupplierService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/supplier.SupplierService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) SupplierBySupplierID(ctx context.Context, in *SupplierBySupplierIDRequest, opts ...grpc.CallOption) (*typespb.Supplier, error) {
	out := new(typespb.Supplier)
	err := c.cc.Invoke(ctx, "/supplier.SupplierService/SupplierBySupplierID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/supplier.SupplierService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/supplier.SupplierService/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupplierServiceServer is the server API for SupplierService service.
// All implementations must embed UnimplementedSupplierServiceServer
// for forward compatibility
type SupplierServiceServer interface {
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	Insert(context.Context, *InsertRequest) (*emptypb.Empty, error)
	SupplierBySupplierID(context.Context, *SupplierBySupplierIDRequest) (*typespb.Supplier, error)
	Update(context.Context, *UpdateRequest) (*emptypb.Empty, error)
	Upsert(context.Context, *UpsertRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedSupplierServiceServer()
}

// UnimplementedSupplierServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSupplierServiceServer struct {
}

func (UnimplementedSupplierServiceServer) Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSupplierServiceServer) Insert(context.Context, *InsertRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedSupplierServiceServer) SupplierBySupplierID(context.Context, *SupplierBySupplierIDRequest) (*typespb.Supplier, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SupplierBySupplierID not implemented")
}
func (UnimplementedSupplierServiceServer) Update(context.Context, *UpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSupplierServiceServer) Upsert(context.Context, *UpsertRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (UnimplementedSupplierServiceServer) mustEmbedUnimplementedSupplierServiceServer() {}

// UnsafeSupplierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupplierServiceServer will
// result in compilation errors.
type UnsafeSupplierServiceServer interface {
	mustEmbedUnimplementedSupplierServiceServer()
}

func RegisterSupplierServiceServer(s grpc.ServiceRegistrar, srv SupplierServiceServer) {
	s.RegisterService(&SupplierService_ServiceDesc, srv)
}

func _SupplierService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supplier.SupplierService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supplier.SupplierService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).Insert(ctx, req.(*InsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_SupplierBySupplierID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupplierBySupplierIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).SupplierBySupplierID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supplier.SupplierService/SupplierBySupplierID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).SupplierBySupplierID(ctx, req.(*SupplierBySupplierIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supplier.SupplierService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supplier.SupplierService/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).Upsert(ctx, req.(*UpsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SupplierService_ServiceDesc is the grpc.ServiceDesc for SupplierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SupplierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "supplier.SupplierService",
	HandlerType: (*SupplierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Delete",
			Handler:    _SupplierService_Delete_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _SupplierService_Insert_Handler,
		},
		{
			MethodName: "SupplierBySupplierID",
			Handler:    _SupplierService_SupplierBySupplierID_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SupplierService_Update_Handler,
		},
		{
			MethodName: "Upsert",
			Handler:    _SupplierService_Upsert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supplier.proto",
}
