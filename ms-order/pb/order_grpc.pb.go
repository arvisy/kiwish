// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: order.proto

package pb

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

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	// order
	OrderDirectCreate(ctx context.Context, in *OrderDirectCreateRequest, opts ...grpc.CallOption) (*OrderDirectCreateResponse, error)
	// courier
	AddCourierInfo(ctx context.Context, in *AddCourierInfoRequest, opts ...grpc.CallOption) (*CourierResponse, error)
	// cart
	CartCreate(ctx context.Context, in *CartCreateRequest, opts ...grpc.CallOption) (*CartCreateResponse, error)
	CartGetByID(ctx context.Context, in *CartGetByIDRequest, opts ...grpc.CallOption) (*CartGetByIDResponse, error)
	CartGetAll(ctx context.Context, in *CartGetAllRequest, opts ...grpc.CallOption) (*CartGetAllResponse, error)
	CartUpdate(ctx context.Context, in *CartUpdateRequest, opts ...grpc.CallOption) (*CartUpdateResponse, error)
	CartDeleteOne(ctx context.Context, in *CartDeleteOneRequest, opts ...grpc.CallOption) (*CartDeleteOneResponse, error)
	CartDeleteAll(ctx context.Context, in *CartDeleteAllRequest, opts ...grpc.CallOption) (*CartDeleteAllResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) OrderDirectCreate(ctx context.Context, in *OrderDirectCreateRequest, opts ...grpc.CallOption) (*OrderDirectCreateResponse, error) {
	out := new(OrderDirectCreateResponse)
	err := c.cc.Invoke(ctx, "/OrderService/OrderDirectCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) AddCourierInfo(ctx context.Context, in *AddCourierInfoRequest, opts ...grpc.CallOption) (*CourierResponse, error) {
	out := new(CourierResponse)
	err := c.cc.Invoke(ctx, "/OrderService/AddCourierInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CartCreate(ctx context.Context, in *CartCreateRequest, opts ...grpc.CallOption) (*CartCreateResponse, error) {
	out := new(CartCreateResponse)
	err := c.cc.Invoke(ctx, "/OrderService/CartCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CartGetByID(ctx context.Context, in *CartGetByIDRequest, opts ...grpc.CallOption) (*CartGetByIDResponse, error) {
	out := new(CartGetByIDResponse)
	err := c.cc.Invoke(ctx, "/OrderService/CartGetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CartGetAll(ctx context.Context, in *CartGetAllRequest, opts ...grpc.CallOption) (*CartGetAllResponse, error) {
	out := new(CartGetAllResponse)
	err := c.cc.Invoke(ctx, "/OrderService/CartGetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CartUpdate(ctx context.Context, in *CartUpdateRequest, opts ...grpc.CallOption) (*CartUpdateResponse, error) {
	out := new(CartUpdateResponse)
	err := c.cc.Invoke(ctx, "/OrderService/CartUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CartDeleteOne(ctx context.Context, in *CartDeleteOneRequest, opts ...grpc.CallOption) (*CartDeleteOneResponse, error) {
	out := new(CartDeleteOneResponse)
	err := c.cc.Invoke(ctx, "/OrderService/CartDeleteOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CartDeleteAll(ctx context.Context, in *CartDeleteAllRequest, opts ...grpc.CallOption) (*CartDeleteAllResponse, error) {
	out := new(CartDeleteAllResponse)
	err := c.cc.Invoke(ctx, "/OrderService/CartDeleteAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	// order
	OrderDirectCreate(context.Context, *OrderDirectCreateRequest) (*OrderDirectCreateResponse, error)
	// courier
	AddCourierInfo(context.Context, *AddCourierInfoRequest) (*CourierResponse, error)
	// cart
	CartCreate(context.Context, *CartCreateRequest) (*CartCreateResponse, error)
	CartGetByID(context.Context, *CartGetByIDRequest) (*CartGetByIDResponse, error)
	CartGetAll(context.Context, *CartGetAllRequest) (*CartGetAllResponse, error)
	CartUpdate(context.Context, *CartUpdateRequest) (*CartUpdateResponse, error)
	CartDeleteOne(context.Context, *CartDeleteOneRequest) (*CartDeleteOneResponse, error)
	CartDeleteAll(context.Context, *CartDeleteAllRequest) (*CartDeleteAllResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) OrderDirectCreate(context.Context, *OrderDirectCreateRequest) (*OrderDirectCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderDirectCreate not implemented")
}
func (UnimplementedOrderServiceServer) AddCourierInfo(context.Context, *AddCourierInfoRequest) (*CourierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCourierInfo not implemented")
}
func (UnimplementedOrderServiceServer) CartCreate(context.Context, *CartCreateRequest) (*CartCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartCreate not implemented")
}
func (UnimplementedOrderServiceServer) CartGetByID(context.Context, *CartGetByIDRequest) (*CartGetByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartGetByID not implemented")
}
func (UnimplementedOrderServiceServer) CartGetAll(context.Context, *CartGetAllRequest) (*CartGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartGetAll not implemented")
}
func (UnimplementedOrderServiceServer) CartUpdate(context.Context, *CartUpdateRequest) (*CartUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartUpdate not implemented")
}
func (UnimplementedOrderServiceServer) CartDeleteOne(context.Context, *CartDeleteOneRequest) (*CartDeleteOneResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartDeleteOne not implemented")
}
func (UnimplementedOrderServiceServer) CartDeleteAll(context.Context, *CartDeleteAllRequest) (*CartDeleteAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CartDeleteAll not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_OrderDirectCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderDirectCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).OrderDirectCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/OrderDirectCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).OrderDirectCreate(ctx, req.(*OrderDirectCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_AddCourierInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCourierInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).AddCourierInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/AddCourierInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).AddCourierInfo(ctx, req.(*AddCourierInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CartCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CartCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/CartCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CartCreate(ctx, req.(*CartCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CartGetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartGetByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CartGetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/CartGetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CartGetByID(ctx, req.(*CartGetByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CartGetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartGetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CartGetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/CartGetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CartGetAll(ctx, req.(*CartGetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CartUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CartUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/CartUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CartUpdate(ctx, req.(*CartUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CartDeleteOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartDeleteOneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CartDeleteOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/CartDeleteOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CartDeleteOne(ctx, req.(*CartDeleteOneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CartDeleteAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CartDeleteAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CartDeleteAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OrderService/CartDeleteAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CartDeleteAll(ctx, req.(*CartDeleteAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OrderDirectCreate",
			Handler:    _OrderService_OrderDirectCreate_Handler,
		},
		{
			MethodName: "AddCourierInfo",
			Handler:    _OrderService_AddCourierInfo_Handler,
		},
		{
			MethodName: "CartCreate",
			Handler:    _OrderService_CartCreate_Handler,
		},
		{
			MethodName: "CartGetByID",
			Handler:    _OrderService_CartGetByID_Handler,
		},
		{
			MethodName: "CartGetAll",
			Handler:    _OrderService_CartGetAll_Handler,
		},
		{
			MethodName: "CartUpdate",
			Handler:    _OrderService_CartUpdate_Handler,
		},
		{
			MethodName: "CartDeleteOne",
			Handler:    _OrderService_CartDeleteOne_Handler,
		},
		{
			MethodName: "CartDeleteAll",
			Handler:    _OrderService_CartDeleteAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
