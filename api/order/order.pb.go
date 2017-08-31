// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order/order.proto

/*
Package order is a generated protocol buffer package.

It is generated from these files:
	order/order.proto

It has these top-level messages:
	Order
	CreateOrderRequest
	GetOrderRequest
	GetOrderDetailRequest
	GetOrderDetailResponse
	ListOrdersRequest
	ListOrdersResponse
	DeleteOrderRequest
*/
package order

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import item "github.com/bakins/kubernetes-envoy-example/api/item"
import _ "github.com/mwitkow/go-proto-validators"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Order struct {
	Id    string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	User  string   `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Items []string `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Order) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

type CreateOrderRequest struct {
	User  string   `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Items []string `protobuf:"bytes,2,rep,name=items" json:"items,omitempty"`
}

func (m *CreateOrderRequest) Reset()                    { *m = CreateOrderRequest{} }
func (m *CreateOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateOrderRequest) ProtoMessage()               {}
func (*CreateOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateOrderRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *CreateOrderRequest) GetItems() []string {
	if m != nil {
		return m.Items
	}
	return nil
}

type GetOrderRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetOrderRequest) Reset()                    { *m = GetOrderRequest{} }
func (m *GetOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOrderRequest) ProtoMessage()               {}
func (*GetOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetOrderRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetOrderDetailRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetOrderDetailRequest) Reset()                    { *m = GetOrderDetailRequest{} }
func (m *GetOrderDetailRequest) String() string            { return proto.CompactTextString(m) }
func (*GetOrderDetailRequest) ProtoMessage()               {}
func (*GetOrderDetailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GetOrderDetailRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetOrderDetailResponse struct {
	Id    string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	User  string       `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Items []*item.Item `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
}

func (m *GetOrderDetailResponse) Reset()                    { *m = GetOrderDetailResponse{} }
func (m *GetOrderDetailResponse) String() string            { return proto.CompactTextString(m) }
func (*GetOrderDetailResponse) ProtoMessage()               {}
func (*GetOrderDetailResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetOrderDetailResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetOrderDetailResponse) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *GetOrderDetailResponse) GetItems() []*item.Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type ListOrdersRequest struct {
	// user is optional
	User string `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *ListOrdersRequest) Reset()                    { *m = ListOrdersRequest{} }
func (m *ListOrdersRequest) String() string            { return proto.CompactTextString(m) }
func (*ListOrdersRequest) ProtoMessage()               {}
func (*ListOrdersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ListOrdersRequest) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

type ListOrdersResponse struct {
	Orders []*Order `protobuf:"bytes,1,rep,name=orders" json:"orders,omitempty"`
}

func (m *ListOrdersResponse) Reset()                    { *m = ListOrdersResponse{} }
func (m *ListOrdersResponse) String() string            { return proto.CompactTextString(m) }
func (*ListOrdersResponse) ProtoMessage()               {}
func (*ListOrdersResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListOrdersResponse) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type DeleteOrderRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteOrderRequest) Reset()                    { *m = DeleteOrderRequest{} }
func (m *DeleteOrderRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteOrderRequest) ProtoMessage()               {}
func (*DeleteOrderRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DeleteOrderRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Order)(nil), "order.Order")
	proto.RegisterType((*CreateOrderRequest)(nil), "order.CreateOrderRequest")
	proto.RegisterType((*GetOrderRequest)(nil), "order.GetOrderRequest")
	proto.RegisterType((*GetOrderDetailRequest)(nil), "order.GetOrderDetailRequest")
	proto.RegisterType((*GetOrderDetailResponse)(nil), "order.GetOrderDetailResponse")
	proto.RegisterType((*ListOrdersRequest)(nil), "order.ListOrdersRequest")
	proto.RegisterType((*ListOrdersResponse)(nil), "order.ListOrdersResponse")
	proto.RegisterType((*DeleteOrderRequest)(nil), "order.DeleteOrderRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for OrderService service

type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error)
	GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*Order, error)
	GetOrderDetail(ctx context.Context, in *GetOrderDetailRequest, opts ...grpc.CallOption) (*GetOrderDetailResponse, error)
	ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error)
	DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*Order, error)
	UpdateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
}

type orderServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrderServiceClient(cc *grpc.ClientConn) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/order.OrderService/CreateOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrder(ctx context.Context, in *GetOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/order.OrderService/GetOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrderDetail(ctx context.Context, in *GetOrderDetailRequest, opts ...grpc.CallOption) (*GetOrderDetailResponse, error) {
	out := new(GetOrderDetailResponse)
	err := grpc.Invoke(ctx, "/order.OrderService/GetOrderDetail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ListOrders(ctx context.Context, in *ListOrdersRequest, opts ...grpc.CallOption) (*ListOrdersResponse, error) {
	out := new(ListOrdersResponse)
	err := grpc.Invoke(ctx, "/order.OrderService/ListOrders", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/order.OrderService/DeleteOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/order.OrderService/UpdateOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OrderService service

type OrderServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*Order, error)
	GetOrder(context.Context, *GetOrderRequest) (*Order, error)
	GetOrderDetail(context.Context, *GetOrderDetailRequest) (*GetOrderDetailResponse, error)
	ListOrders(context.Context, *ListOrdersRequest) (*ListOrdersResponse, error)
	DeleteOrder(context.Context, *DeleteOrderRequest) (*Order, error)
	UpdateOrder(context.Context, *Order) (*Order, error)
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrder(ctx, req.(*GetOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrderDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/GetOrderDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderDetail(ctx, req.(*GetOrderDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ListOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOrdersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ListOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/ListOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ListOrders(ctx, req.(*ListOrdersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/DeleteOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).DeleteOrder(ctx, req.(*DeleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/UpdateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrder",
			Handler:    _OrderService_GetOrder_Handler,
		},
		{
			MethodName: "GetOrderDetail",
			Handler:    _OrderService_GetOrderDetail_Handler,
		},
		{
			MethodName: "ListOrders",
			Handler:    _OrderService_ListOrders_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _OrderService_DeleteOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _OrderService_UpdateOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order/order.proto",
}

func init() { proto.RegisterFile("order/order.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0x65, 0xb7, 0x4d, 0xb1, 0x37, 0xa1, 0x21, 0x17, 0xdb, 0x6e, 0x96, 0x8a, 0x71, 0x28, 0xb4,
	0x14, 0x92, 0xc1, 0x0a, 0x3e, 0xf4, 0x41, 0x14, 0x0b, 0x52, 0x50, 0x0a, 0x91, 0xbe, 0x0a, 0xdb,
	0xee, 0x25, 0x0e, 0x66, 0x77, 0xd6, 0x9d, 0x49, 0xfa, 0x20, 0xbe, 0xf4, 0x2f, 0xf8, 0xd3, 0xfc,
	0x0b, 0xfe, 0x10, 0xc9, 0xcc, 0x64, 0xb3, 0x5f, 0x15, 0x5f, 0x86, 0x99, 0xb9, 0xf7, 0x9c, 0x33,
	0xf7, 0x9c, 0x5d, 0x18, 0xc8, 0x3c, 0xa6, 0x9c, 0x9b, 0x75, 0x92, 0xe5, 0x52, 0x4b, 0xec, 0x98,
	0x43, 0x78, 0x34, 0x93, 0x72, 0x36, 0x27, 0x1e, 0x65, 0x82, 0x47, 0x69, 0x2a, 0x75, 0xa4, 0x85,
	0x4c, 0x95, 0x6d, 0x0a, 0xfb, 0x42, 0x53, 0xc2, 0x57, 0x8b, 0xbb, 0x78, 0x3d, 0x13, 0xfa, 0xeb,
	0xe2, 0x76, 0x72, 0x27, 0x13, 0x9e, 0xdc, 0x0b, 0xfd, 0x4d, 0xde, 0xf3, 0x99, 0x1c, 0x9b, 0xe2,
	0x78, 0x19, 0xcd, 0x45, 0x1c, 0x69, 0x99, 0x2b, 0x5e, 0x6c, 0x2d, 0x8e, 0xbd, 0x83, 0xce, 0xf5,
	0x4a, 0x0f, 0xf7, 0xc0, 0x17, 0x71, 0xe0, 0x8d, 0xbc, 0xd3, 0xdd, 0xa9, 0x2f, 0x62, 0x44, 0xd8,
	0x5e, 0x28, 0xca, 0x03, 0xdf, 0xdc, 0x98, 0x3d, 0x3e, 0x85, 0xce, 0x4a, 0x52, 0x05, 0x5b, 0xa3,
	0xad, 0xd3, 0xdd, 0xa9, 0x3d, 0xb0, 0x37, 0x80, 0xef, 0x73, 0x8a, 0x34, 0x19, 0xa2, 0x29, 0x7d,
	0x5f, 0x90, 0xd2, 0x05, 0xde, 0x6b, 0xc3, 0xfb, 0x65, 0xfc, 0x0b, 0xe8, 0x7f, 0x20, 0x5d, 0x01,
	0xd7, 0x1e, 0xc3, 0x4e, 0x60, 0x7f, 0xdd, 0x72, 0x49, 0x3a, 0x12, 0xf3, 0xc7, 0x1a, 0xbf, 0xc0,
	0x41, 0xbd, 0x51, 0x65, 0x32, 0x55, 0xf4, 0x5f, 0xf3, 0x8d, 0xca, 0xf3, 0x75, 0xcf, 0x61, 0x62,
	0x0c, 0xbe, 0xd2, 0x94, 0xac, 0xdf, 0x7a, 0x02, 0x83, 0x8f, 0x42, 0x59, 0x01, 0xf5, 0x8f, 0x51,
	0xd9, 0x05, 0x60, 0xb9, 0xd1, 0x3d, 0xe2, 0x18, 0x76, 0x4c, 0xba, 0x2a, 0xf0, 0x8c, 0x42, 0x6f,
	0x62, 0x93, 0xb7, 0xc3, 0xbb, 0x1a, 0x3b, 0x06, 0xbc, 0xa4, 0x39, 0xd5, 0x0c, 0xad, 0x0d, 0x70,
	0xfe, 0xb0, 0x0d, 0x3d, 0xd3, 0xf0, 0x99, 0xf2, 0xa5, 0xb8, 0x23, 0xfc, 0x04, 0xdd, 0x52, 0x0e,
	0x38, 0x74, 0xdc, 0xcd, 0x6c, 0xc2, 0x8a, 0x2c, 0xdb, 0x7f, 0xf8, 0xfd, 0xe7, 0x97, 0xdf, 0x67,
	0xc0, 0x97, 0x2f, 0xed, 0x97, 0xa8, 0x2e, 0xbc, 0x33, 0xbc, 0x82, 0x27, 0x6b, 0x2b, 0xf1, 0xc0,
	0x01, 0x6a, 0x39, 0xd5, 0x88, 0x0e, 0x0d, 0xd1, 0x00, 0xfb, 0x1b, 0x22, 0xfe, 0x43, 0xc4, 0x3f,
	0x31, 0x85, 0xbd, 0x6a, 0x2a, 0x78, 0x54, 0x23, 0xac, 0xa4, 0x1a, 0x3e, 0x7b, 0xa4, 0x6a, 0x5d,
	0x64, 0xcf, 0x8d, 0xce, 0x10, 0x0f, 0x37, 0x3a, 0xe3, 0xd8, 0xb4, 0x38, 0xbd, 0x1b, 0x80, 0x8d,
	0xf9, 0x18, 0x38, 0xb6, 0x46, 0x70, 0xe1, 0xb0, 0xa5, 0xe2, 0x34, 0xd0, 0x68, 0xf4, 0xb0, 0x64,
	0x0a, 0x5e, 0x43, 0xb7, 0x94, 0x4b, 0x61, 0x70, 0x33, 0xab, 0x76, 0x5f, 0xce, 0x1a, 0xbe, 0xbc,
	0x85, 0xee, 0x4d, 0x16, 0x17, 0x89, 0x55, 0x50, 0xed, 0x21, 0x85, 0xd5, 0x90, 0x6e, 0x77, 0xcc,
	0x5f, 0xfc, 0xea, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xec, 0x82, 0x18, 0xa6, 0x48, 0x04, 0x00,
	0x00,
}
