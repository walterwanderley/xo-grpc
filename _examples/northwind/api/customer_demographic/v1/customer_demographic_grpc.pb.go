// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: customer_demographic/v1/customer_demographic.proto

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

// CustomerDemographicServiceClient is the client API for CustomerDemographicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerDemographicServiceClient interface {
	CustomerDemographicByCustomerTypeID(ctx context.Context, in *CustomerDemographicByCustomerTypeIDRequest, opts ...grpc.CallOption) (*v1.CustomerDemographic, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type customerDemographicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerDemographicServiceClient(cc grpc.ClientConnInterface) CustomerDemographicServiceClient {
	return &customerDemographicServiceClient{cc}
}

func (c *customerDemographicServiceClient) CustomerDemographicByCustomerTypeID(ctx context.Context, in *CustomerDemographicByCustomerTypeIDRequest, opts ...grpc.CallOption) (*v1.CustomerDemographic, error) {
	out := new(v1.CustomerDemographic)
	err := c.cc.Invoke(ctx, "/customer_demographic.v1.CustomerDemographicService/CustomerDemographicByCustomerTypeID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerDemographicServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/customer_demographic.v1.CustomerDemographicService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerDemographicServiceClient) Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/customer_demographic.v1.CustomerDemographicService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerDemographicServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/customer_demographic.v1.CustomerDemographicService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerDemographicServiceClient) Upsert(ctx context.Context, in *UpsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/customer_demographic.v1.CustomerDemographicService/Upsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerDemographicServiceServer is the server API for CustomerDemographicService service.
// All implementations must embed UnimplementedCustomerDemographicServiceServer
// for forward compatibility
type CustomerDemographicServiceServer interface {
	CustomerDemographicByCustomerTypeID(context.Context, *CustomerDemographicByCustomerTypeIDRequest) (*v1.CustomerDemographic, error)
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	Insert(context.Context, *InsertRequest) (*emptypb.Empty, error)
	Update(context.Context, *UpdateRequest) (*emptypb.Empty, error)
	Upsert(context.Context, *UpsertRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCustomerDemographicServiceServer()
}

// UnimplementedCustomerDemographicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerDemographicServiceServer struct {
}

func (UnimplementedCustomerDemographicServiceServer) CustomerDemographicByCustomerTypeID(context.Context, *CustomerDemographicByCustomerTypeIDRequest) (*v1.CustomerDemographic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustomerDemographicByCustomerTypeID not implemented")
}
func (UnimplementedCustomerDemographicServiceServer) Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCustomerDemographicServiceServer) Insert(context.Context, *InsertRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedCustomerDemographicServiceServer) Update(context.Context, *UpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCustomerDemographicServiceServer) Upsert(context.Context, *UpsertRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upsert not implemented")
}
func (UnimplementedCustomerDemographicServiceServer) mustEmbedUnimplementedCustomerDemographicServiceServer() {
}

// UnsafeCustomerDemographicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerDemographicServiceServer will
// result in compilation errors.
type UnsafeCustomerDemographicServiceServer interface {
	mustEmbedUnimplementedCustomerDemographicServiceServer()
}

func RegisterCustomerDemographicServiceServer(s grpc.ServiceRegistrar, srv CustomerDemographicServiceServer) {
	s.RegisterService(&CustomerDemographicService_ServiceDesc, srv)
}

func _CustomerDemographicService_CustomerDemographicByCustomerTypeID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerDemographicByCustomerTypeIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerDemographicServiceServer).CustomerDemographicByCustomerTypeID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer_demographic.v1.CustomerDemographicService/CustomerDemographicByCustomerTypeID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerDemographicServiceServer).CustomerDemographicByCustomerTypeID(ctx, req.(*CustomerDemographicByCustomerTypeIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerDemographicService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerDemographicServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer_demographic.v1.CustomerDemographicService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerDemographicServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerDemographicService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerDemographicServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer_demographic.v1.CustomerDemographicService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerDemographicServiceServer).Insert(ctx, req.(*InsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerDemographicService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerDemographicServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer_demographic.v1.CustomerDemographicService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerDemographicServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerDemographicService_Upsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerDemographicServiceServer).Upsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer_demographic.v1.CustomerDemographicService/Upsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerDemographicServiceServer).Upsert(ctx, req.(*UpsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerDemographicService_ServiceDesc is the grpc.ServiceDesc for CustomerDemographicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerDemographicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "customer_demographic.v1.CustomerDemographicService",
	HandlerType: (*CustomerDemographicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CustomerDemographicByCustomerTypeID",
			Handler:    _CustomerDemographicService_CustomerDemographicByCustomerTypeID_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CustomerDemographicService_Delete_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _CustomerDemographicService_Insert_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CustomerDemographicService_Update_Handler,
		},
		{
			MethodName: "Upsert",
			Handler:    _CustomerDemographicService_Upsert_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer_demographic/v1/customer_demographic.proto",
}
