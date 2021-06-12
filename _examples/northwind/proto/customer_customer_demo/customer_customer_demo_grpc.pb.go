// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package customer_customer_demo

import (
	context "context"
	typespb "github.com/walterwanderley/xo-grpc/_examples/northwind/proto/typespb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CustomerCustomerDemoServiceClient is the client API for CustomerCustomerDemoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerCustomerDemoServiceClient interface {
	Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CustomerCustomerDemoByCustomerIDCustomerTypeID(ctx context.Context, in *CustomerCustomerDemoByCustomerIDCustomerTypeIDRequest, opts ...grpc.CallOption) (*typespb.CustomerCustomerDemo, error)
	Customer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*typespb.Customer, error)
	CustomerDemographic(ctx context.Context, in *CustomerDemographicRequest, opts ...grpc.CallOption) (*typespb.CustomerDemographic, error)
}

type customerCustomerDemoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerCustomerDemoServiceClient(cc grpc.ClientConnInterface) CustomerCustomerDemoServiceClient {
	return &customerCustomerDemoServiceClient{cc}
}

func (c *customerCustomerDemoServiceClient) Insert(ctx context.Context, in *InsertRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/customer_customer_demo.CustomerCustomerDemoService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerCustomerDemoServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/customer_customer_demo.CustomerCustomerDemoService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerCustomerDemoServiceClient) CustomerCustomerDemoByCustomerIDCustomerTypeID(ctx context.Context, in *CustomerCustomerDemoByCustomerIDCustomerTypeIDRequest, opts ...grpc.CallOption) (*typespb.CustomerCustomerDemo, error) {
	out := new(typespb.CustomerCustomerDemo)
	err := c.cc.Invoke(ctx, "/customer_customer_demo.CustomerCustomerDemoService/CustomerCustomerDemoByCustomerIDCustomerTypeID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerCustomerDemoServiceClient) Customer(ctx context.Context, in *CustomerRequest, opts ...grpc.CallOption) (*typespb.Customer, error) {
	out := new(typespb.Customer)
	err := c.cc.Invoke(ctx, "/customer_customer_demo.CustomerCustomerDemoService/Customer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerCustomerDemoServiceClient) CustomerDemographic(ctx context.Context, in *CustomerDemographicRequest, opts ...grpc.CallOption) (*typespb.CustomerDemographic, error) {
	out := new(typespb.CustomerDemographic)
	err := c.cc.Invoke(ctx, "/customer_customer_demo.CustomerCustomerDemoService/CustomerDemographic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerCustomerDemoServiceServer is the server API for CustomerCustomerDemoService service.
// All implementations must embed UnimplementedCustomerCustomerDemoServiceServer
// for forward compatibility
type CustomerCustomerDemoServiceServer interface {
	Insert(context.Context, *InsertRequest) (*emptypb.Empty, error)
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	CustomerCustomerDemoByCustomerIDCustomerTypeID(context.Context, *CustomerCustomerDemoByCustomerIDCustomerTypeIDRequest) (*typespb.CustomerCustomerDemo, error)
	Customer(context.Context, *CustomerRequest) (*typespb.Customer, error)
	CustomerDemographic(context.Context, *CustomerDemographicRequest) (*typespb.CustomerDemographic, error)
	mustEmbedUnimplementedCustomerCustomerDemoServiceServer()
}

// UnimplementedCustomerCustomerDemoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerCustomerDemoServiceServer struct {
}

func (UnimplementedCustomerCustomerDemoServiceServer) Insert(context.Context, *InsertRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedCustomerCustomerDemoServiceServer) Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCustomerCustomerDemoServiceServer) CustomerCustomerDemoByCustomerIDCustomerTypeID(context.Context, *CustomerCustomerDemoByCustomerIDCustomerTypeIDRequest) (*typespb.CustomerCustomerDemo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustomerCustomerDemoByCustomerIDCustomerTypeID not implemented")
}
func (UnimplementedCustomerCustomerDemoServiceServer) Customer(context.Context, *CustomerRequest) (*typespb.Customer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Customer not implemented")
}
func (UnimplementedCustomerCustomerDemoServiceServer) CustomerDemographic(context.Context, *CustomerDemographicRequest) (*typespb.CustomerDemographic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CustomerDemographic not implemented")
}
func (UnimplementedCustomerCustomerDemoServiceServer) mustEmbedUnimplementedCustomerCustomerDemoServiceServer() {
}

// UnsafeCustomerCustomerDemoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerCustomerDemoServiceServer will
// result in compilation errors.
type UnsafeCustomerCustomerDemoServiceServer interface {
	mustEmbedUnimplementedCustomerCustomerDemoServiceServer()
}

func RegisterCustomerCustomerDemoServiceServer(s grpc.ServiceRegistrar, srv CustomerCustomerDemoServiceServer) {
	s.RegisterService(&CustomerCustomerDemoService_ServiceDesc, srv)
}

func _CustomerCustomerDemoService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerCustomerDemoServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer_customer_demo.CustomerCustomerDemoService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerCustomerDemoServiceServer).Insert(ctx, req.(*InsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerCustomerDemoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerCustomerDemoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer_customer_demo.CustomerCustomerDemoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerCustomerDemoServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerCustomerDemoService_CustomerCustomerDemoByCustomerIDCustomerTypeID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerCustomerDemoByCustomerIDCustomerTypeIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerCustomerDemoServiceServer).CustomerCustomerDemoByCustomerIDCustomerTypeID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer_customer_demo.CustomerCustomerDemoService/CustomerCustomerDemoByCustomerIDCustomerTypeID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerCustomerDemoServiceServer).CustomerCustomerDemoByCustomerIDCustomerTypeID(ctx, req.(*CustomerCustomerDemoByCustomerIDCustomerTypeIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerCustomerDemoService_Customer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerCustomerDemoServiceServer).Customer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer_customer_demo.CustomerCustomerDemoService/Customer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerCustomerDemoServiceServer).Customer(ctx, req.(*CustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerCustomerDemoService_CustomerDemographic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerDemographicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerCustomerDemoServiceServer).CustomerDemographic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/customer_customer_demo.CustomerCustomerDemoService/CustomerDemographic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerCustomerDemoServiceServer).CustomerDemographic(ctx, req.(*CustomerDemographicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerCustomerDemoService_ServiceDesc is the grpc.ServiceDesc for CustomerCustomerDemoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerCustomerDemoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "customer_customer_demo.CustomerCustomerDemoService",
	HandlerType: (*CustomerCustomerDemoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Insert",
			Handler:    _CustomerCustomerDemoService_Insert_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CustomerCustomerDemoService_Delete_Handler,
		},
		{
			MethodName: "CustomerCustomerDemoByCustomerIDCustomerTypeID",
			Handler:    _CustomerCustomerDemoService_CustomerCustomerDemoByCustomerIDCustomerTypeID_Handler,
		},
		{
			MethodName: "Customer",
			Handler:    _CustomerCustomerDemoService_Customer_Handler,
		},
		{
			MethodName: "CustomerDemographic",
			Handler:    _CustomerCustomerDemoService_CustomerDemographic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer_customer_demo.proto",
}