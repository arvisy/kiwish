// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/cart.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId int64 `protobuf:"varint,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int64 `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_proto_cart_proto_rawDescGZIP(), []int{0}
}

func (x *Cart) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Cart) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Cart) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *Cart) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CreateCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ProductId int64 `protobuf:"varint,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity  int64 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *CreateCartRequest) Reset() {
	*x = CreateCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCartRequest) ProtoMessage() {}

func (x *CreateCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCartRequest.ProtoReflect.Descriptor instead.
func (*CreateCartRequest) Descriptor() ([]byte, []int) {
	return file_proto_cart_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCartRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateCartRequest) GetProductId() int64 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *CreateCartRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type CreateCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cart    *Cart  `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateCartResponse) Reset() {
	*x = CreateCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCartResponse) ProtoMessage() {}

func (x *CreateCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCartResponse.ProtoReflect.Descriptor instead.
func (*CreateCartResponse) Descriptor() ([]byte, []int) {
	return file_proto_cart_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCartResponse) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

func (x *CreateCartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetAllCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetAllCartRequest) Reset() {
	*x = GetAllCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCartRequest) ProtoMessage() {}

func (x *GetAllCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCartRequest.ProtoReflect.Descriptor instead.
func (*GetAllCartRequest) Descriptor() ([]byte, []int) {
	return file_proto_cart_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllCartRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetAllCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Carts []*Cart `protobuf:"bytes,1,rep,name=carts,proto3" json:"carts,omitempty"`
}

func (x *GetAllCartResponse) Reset() {
	*x = GetAllCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCartResponse) ProtoMessage() {}

func (x *GetAllCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCartResponse.ProtoReflect.Descriptor instead.
func (*GetAllCartResponse) Descriptor() ([]byte, []int) {
	return file_proto_cart_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllCartResponse) GetCarts() []*Cart {
	if x != nil {
		return x.Carts
	}
	return nil
}

type UpadateCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId   int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Quantity int64 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *UpadateCartRequest) Reset() {
	*x = UpadateCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpadateCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpadateCartRequest) ProtoMessage() {}

func (x *UpadateCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpadateCartRequest.ProtoReflect.Descriptor instead.
func (*UpadateCartRequest) Descriptor() ([]byte, []int) {
	return file_proto_cart_proto_rawDescGZIP(), []int{5}
}

func (x *UpadateCartRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpadateCartRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpadateCartRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type UpdateCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cart    *Cart  `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateCartResponse) Reset() {
	*x = UpdateCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCartResponse) ProtoMessage() {}

func (x *UpdateCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCartResponse.ProtoReflect.Descriptor instead.
func (*UpdateCartResponse) Descriptor() ([]byte, []int) {
	return file_proto_cart_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCartResponse) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

func (x *UpdateCartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteOneCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteOneCartRequest) Reset() {
	*x = DeleteOneCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOneCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOneCartRequest) ProtoMessage() {}

func (x *DeleteOneCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOneCartRequest.ProtoReflect.Descriptor instead.
func (*DeleteOneCartRequest) Descriptor() ([]byte, []int) {
	return file_proto_cart_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteOneCartRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *DeleteOneCartRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DeleteOneCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cart    *Cart  `protobuf:"bytes,1,opt,name=cart,proto3" json:"cart,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteOneCartResponse) Reset() {
	*x = DeleteOneCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOneCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOneCartResponse) ProtoMessage() {}

func (x *DeleteOneCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOneCartResponse.ProtoReflect.Descriptor instead.
func (*DeleteOneCartResponse) Descriptor() ([]byte, []int) {
	return file_proto_cart_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteOneCartResponse) GetCart() *Cart {
	if x != nil {
		return x.Cart
	}
	return nil
}

func (x *DeleteOneCartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteAllCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteAllCartRequest) Reset() {
	*x = DeleteAllCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAllCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllCartRequest) ProtoMessage() {}

func (x *DeleteAllCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllCartRequest.ProtoReflect.Descriptor instead.
func (*DeleteAllCartRequest) Descriptor() ([]byte, []int) {
	return file_proto_cart_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteAllCartRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type DeleteAllCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      []int64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
	Message string  `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteAllCartResponse) Reset() {
	*x = DeleteAllCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cart_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAllCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllCartResponse) ProtoMessage() {}

func (x *DeleteAllCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cart_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllCartResponse.ProtoReflect.Descriptor instead.
func (*DeleteAllCartResponse) Descriptor() ([]byte, []int) {
	return file_proto_cart_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteAllCartResponse) GetId() []int64 {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *DeleteAllCartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_proto_cart_proto protoreflect.FileDescriptor

var file_proto_cart_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x67,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x49, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x2c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x31, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x63, 0x61, 0x72, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x05, 0x63, 0x61,
	0x72, 0x74, 0x73, 0x22, 0x59, 0x0a, 0x12, 0x55, 0x70, 0x61, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x49,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3f, 0x0a, 0x14, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4f, 0x6e, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x6e, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x05, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x52, 0x04, 0x63, 0x61, 0x72, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2f, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x15, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xa9, 0x02, 0x0a,
	0x0b, 0x43, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x33, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x12, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x13, 0x2e, 0x55, 0x70, 0x61, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x6e, 0x65, 0x12, 0x15, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4f, 0x6e, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x6e, 0x65, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x09, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x6d, 0x73, 0x2d, 0x63,
	0x61, 0x72, 0x74, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_cart_proto_rawDescOnce sync.Once
	file_proto_cart_proto_rawDescData = file_proto_cart_proto_rawDesc
)

func file_proto_cart_proto_rawDescGZIP() []byte {
	file_proto_cart_proto_rawDescOnce.Do(func() {
		file_proto_cart_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cart_proto_rawDescData)
	})
	return file_proto_cart_proto_rawDescData
}

var file_proto_cart_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_cart_proto_goTypes = []interface{}{
	(*Cart)(nil),                  // 0: Cart
	(*CreateCartRequest)(nil),     // 1: CreateCartRequest
	(*CreateCartResponse)(nil),    // 2: CreateCartResponse
	(*GetAllCartRequest)(nil),     // 3: GetAllCartRequest
	(*GetAllCartResponse)(nil),    // 4: GetAllCartResponse
	(*UpadateCartRequest)(nil),    // 5: UpadateCartRequest
	(*UpdateCartResponse)(nil),    // 6: UpdateCartResponse
	(*DeleteOneCartRequest)(nil),  // 7: DeleteOneCartRequest
	(*DeleteOneCartResponse)(nil), // 8: DeleteOneCartResponse
	(*DeleteAllCartRequest)(nil),  // 9: DeleteAllCartRequest
	(*DeleteAllCartResponse)(nil), // 10: DeleteAllCartResponse
}
var file_proto_cart_proto_depIdxs = []int32{
	0,  // 0: CreateCartResponse.cart:type_name -> Cart
	0,  // 1: GetAllCartResponse.carts:type_name -> Cart
	0,  // 2: UpdateCartResponse.cart:type_name -> Cart
	0,  // 3: DeleteOneCartResponse.cart:type_name -> Cart
	1,  // 4: CartService.Create:input_type -> CreateCartRequest
	3,  // 5: CartService.GetAll:input_type -> GetAllCartRequest
	5,  // 6: CartService.Update:input_type -> UpadateCartRequest
	7,  // 7: CartService.DeleteOne:input_type -> DeleteOneCartRequest
	9,  // 8: CartService.DeleteAll:input_type -> DeleteAllCartRequest
	2,  // 9: CartService.Create:output_type -> CreateCartResponse
	4,  // 10: CartService.GetAll:output_type -> GetAllCartResponse
	6,  // 11: CartService.Update:output_type -> UpdateCartResponse
	8,  // 12: CartService.DeleteOne:output_type -> DeleteOneCartResponse
	10, // 13: CartService.DeleteAll:output_type -> DeleteAllCartResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_cart_proto_init() }
func file_proto_cart_proto_init() {
	if File_proto_cart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cart_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cart_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cart_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cart_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpadateCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cart_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cart_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOneCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cart_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOneCartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cart_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAllCartRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_cart_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAllCartResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_cart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cart_proto_goTypes,
		DependencyIndexes: file_proto_cart_proto_depIdxs,
		MessageInfos:      file_proto_cart_proto_msgTypes,
	}.Build()
	File_proto_cart_proto = out.File
	file_proto_cart_proto_rawDesc = nil
	file_proto_cart_proto_goTypes = nil
	file_proto_cart_proto_depIdxs = nil
}
