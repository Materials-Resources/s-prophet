// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: inventory/v1/inventory.proto

package inventory_v1alpha0

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	InventoryService_GetProductStock_FullMethodName = "/inventory.v1.InventoryService/GetProductStock"
	InventoryService_GetReceiptByID_FullMethodName  = "/inventory.v1.InventoryService/GetReceiptByID"
)

// InventoryServiceClient is the client API for InventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryServiceClient interface {
	GetProductStock(ctx context.Context, in *GetProductStockRequest, opts ...grpc.CallOption) (*GetProductStockResponse, error)
	// GetReceiptByID returns details for a inventory receipt given an identifier.
	GetReceiptByID(ctx context.Context, in *GetReceiptByIDRequest, opts ...grpc.CallOption) (*GetReceiptByIDResponse, error)
}

type inventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryServiceClient(cc grpc.ClientConnInterface) InventoryServiceClient {
	return &inventoryServiceClient{cc}
}

func (c *inventoryServiceClient) GetProductStock(ctx context.Context, in *GetProductStockRequest, opts ...grpc.CallOption) (*GetProductStockResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetProductStockResponse)
	err := c.cc.Invoke(ctx, InventoryService_GetProductStock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) GetReceiptByID(ctx context.Context, in *GetReceiptByIDRequest, opts ...grpc.CallOption) (*GetReceiptByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetReceiptByIDResponse)
	err := c.cc.Invoke(ctx, InventoryService_GetReceiptByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServiceServer is the server API for InventoryService service.
// All implementations should embed UnimplementedInventoryServiceServer
// for forward compatibility
type InventoryServiceServer interface {
	GetProductStock(context.Context, *GetProductStockRequest) (*GetProductStockResponse, error)
	// GetReceiptByID returns details for a inventory receipt given an identifier.
	GetReceiptByID(context.Context, *GetReceiptByIDRequest) (*GetReceiptByIDResponse, error)
}

// UnimplementedInventoryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedInventoryServiceServer struct {
}

func (UnimplementedInventoryServiceServer) GetProductStock(context.Context, *GetProductStockRequest) (*GetProductStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductStock not implemented")
}
func (UnimplementedInventoryServiceServer) GetReceiptByID(context.Context, *GetReceiptByIDRequest) (*GetReceiptByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceiptByID not implemented")
}

// UnsafeInventoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InventoryServiceServer will
// result in compilation errors.
type UnsafeInventoryServiceServer interface {
	mustEmbedUnimplementedInventoryServiceServer()
}

func RegisterInventoryServiceServer(s grpc.ServiceRegistrar, srv InventoryServiceServer) {
	s.RegisterService(&InventoryService_ServiceDesc, srv)
}

func _InventoryService_GetProductStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetProductStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_GetProductStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetProductStock(ctx, req.(*GetProductStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_GetReceiptByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReceiptByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).GetReceiptByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_GetReceiptByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).GetReceiptByID(ctx, req.(*GetReceiptByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InventoryService_ServiceDesc is the grpc.ServiceDesc for InventoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.v1.InventoryService",
	HandlerType: (*InventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductStock",
			Handler:    _InventoryService_GetProductStock_Handler,
		},
		{
			MethodName: "GetReceiptByID",
			Handler:    _InventoryService_GetReceiptByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory/v1/inventory.proto",
}
