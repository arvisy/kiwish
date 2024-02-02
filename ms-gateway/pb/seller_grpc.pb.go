// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: seller.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SellerServiceClient is the client API for SellerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SellerServiceClient interface {
	// product
	AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	GetProductsBySeller(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
	GetProductsByCategory(ctx context.Context, in *GetProductByCategoryRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
	GetProductByID(ctx context.Context, in *GetProductByIDRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
}

type sellerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSellerServiceClient(cc grpc.ClientConnInterface) SellerServiceClient {
	return &sellerServiceClient{cc}
}

func (c *sellerServiceClient) AddProduct(ctx context.Context, in *AddProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/seller.SellerService/AddProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerServiceClient) GetProductsBySeller(ctx context.Context, in *GetProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	out := new(GetProductsResponse)
	err := c.cc.Invoke(ctx, "/seller.SellerService/GetProductsBySeller", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerServiceClient) GetProductsByCategory(ctx context.Context, in *GetProductByCategoryRequest, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	out := new(GetProductsResponse)
	err := c.cc.Invoke(ctx, "/seller.SellerService/GetProductsByCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerServiceClient) GetProductByID(ctx context.Context, in *GetProductByIDRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/seller.SellerService/GetProductByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerServiceClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/seller.SellerService/DeleteProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sellerServiceClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/seller.SellerService/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SellerServiceServer is the server API for SellerService service.
// All implementations must embed UnimplementedSellerServiceServer
// for forward compatibility
type SellerServiceServer interface {
	// product
	AddProduct(context.Context, *AddProductRequest) (*ProductResponse, error)
	GetProductsBySeller(context.Context, *GetProductsRequest) (*GetProductsResponse, error)
	GetProductsByCategory(context.Context, *GetProductByCategoryRequest) (*GetProductsResponse, error)
	GetProductByID(context.Context, *GetProductByIDRequest) (*ProductResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*emptypb.Empty, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*ProductResponse, error)
	mustEmbedUnimplementedSellerServiceServer()
}

// UnimplementedSellerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSellerServiceServer struct {
}

func (UnimplementedSellerServiceServer) AddProduct(context.Context, *AddProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProduct not implemented")
}
func (UnimplementedSellerServiceServer) GetProductsBySeller(context.Context, *GetProductsRequest) (*GetProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsBySeller not implemented")
}
func (UnimplementedSellerServiceServer) GetProductsByCategory(context.Context, *GetProductByCategoryRequest) (*GetProductsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductsByCategory not implemented")
}
func (UnimplementedSellerServiceServer) GetProductByID(context.Context, *GetProductByIDRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductByID not implemented")
}
func (UnimplementedSellerServiceServer) DeleteProduct(context.Context, *DeleteProductRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedSellerServiceServer) UpdateProduct(context.Context, *UpdateProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedSellerServiceServer) mustEmbedUnimplementedSellerServiceServer() {}

// UnsafeSellerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SellerServiceServer will
// result in compilation errors.
type UnsafeSellerServiceServer interface {
	mustEmbedUnimplementedSellerServiceServer()
}

func RegisterSellerServiceServer(s grpc.ServiceRegistrar, srv SellerServiceServer) {
	s.RegisterService(&SellerService_ServiceDesc, srv)
}

func _SellerService_AddProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerServiceServer).AddProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seller.SellerService/AddProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerServiceServer).AddProduct(ctx, req.(*AddProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerService_GetProductsBySeller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerServiceServer).GetProductsBySeller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seller.SellerService/GetProductsBySeller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerServiceServer).GetProductsBySeller(ctx, req.(*GetProductsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerService_GetProductsByCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductByCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerServiceServer).GetProductsByCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seller.SellerService/GetProductsByCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerServiceServer).GetProductsByCategory(ctx, req.(*GetProductByCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerService_GetProductByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerServiceServer).GetProductByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seller.SellerService/GetProductByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerServiceServer).GetProductByID(ctx, req.(*GetProductByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerService_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerServiceServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seller.SellerService/DeleteProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerServiceServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SellerService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SellerServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seller.SellerService/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SellerServiceServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
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
			MethodName: "AddProduct",
			Handler:    _SellerService_AddProduct_Handler,
		},
		{
			MethodName: "GetProductsBySeller",
			Handler:    _SellerService_GetProductsBySeller_Handler,
		},
		{
			MethodName: "GetProductsByCategory",
			Handler:    _SellerService_GetProductsByCategory_Handler,
		},
		{
			MethodName: "GetProductByID",
			Handler:    _SellerService_GetProductByID_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _SellerService_DeleteProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _SellerService_UpdateProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "seller.proto",
}
