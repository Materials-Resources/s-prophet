// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: inventory/v1alpha0/inventory.proto

package v1alpha0

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

const (
	InventoryService_AddSupplier_FullMethodName    = "/inventory.v1alpha0.InventoryService/AddSupplier"
	InventoryService_DeleteSupplier_FullMethodName = "/inventory.v1alpha0.InventoryService/DeleteSupplier"
)

// InventoryServiceClient is the client API for InventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InventoryServiceClient interface {
	AddSupplier(ctx context.Context, in *AddSupplierRequest, opts ...grpc.CallOption) (*AddSupplierResponse, error)
	DeleteSupplier(ctx context.Context, in *DeleteSupplierRequest, opts ...grpc.CallOption) (*DeleteSupplierResponse, error)
}

type inventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInventoryServiceClient(cc grpc.ClientConnInterface) InventoryServiceClient {
	return &inventoryServiceClient{cc}
}

func (c *inventoryServiceClient) AddSupplier(ctx context.Context, in *AddSupplierRequest, opts ...grpc.CallOption) (*AddSupplierResponse, error) {
	out := new(AddSupplierResponse)
	err := c.cc.Invoke(ctx, InventoryService_AddSupplier_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *inventoryServiceClient) DeleteSupplier(ctx context.Context, in *DeleteSupplierRequest, opts ...grpc.CallOption) (*DeleteSupplierResponse, error) {
	out := new(DeleteSupplierResponse)
	err := c.cc.Invoke(ctx, InventoryService_DeleteSupplier_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InventoryServiceServer is the server API for InventoryService service.
// All implementations should embed UnimplementedInventoryServiceServer
// for forward compatibility
type InventoryServiceServer interface {
	AddSupplier(context.Context, *AddSupplierRequest) (*AddSupplierResponse, error)
	DeleteSupplier(context.Context, *DeleteSupplierRequest) (*DeleteSupplierResponse, error)
}

// UnimplementedInventoryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedInventoryServiceServer struct {
}

func (UnimplementedInventoryServiceServer) AddSupplier(context.Context, *AddSupplierRequest) (*AddSupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSupplier not implemented")
}
func (UnimplementedInventoryServiceServer) DeleteSupplier(context.Context, *DeleteSupplierRequest) (*DeleteSupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSupplier not implemented")
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

func _InventoryService_AddSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).AddSupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_AddSupplier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).AddSupplier(ctx, req.(*AddSupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InventoryService_DeleteSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSupplierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InventoryServiceServer).DeleteSupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: InventoryService_DeleteSupplier_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InventoryServiceServer).DeleteSupplier(ctx, req.(*DeleteSupplierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InventoryService_ServiceDesc is the grpc.ServiceDesc for InventoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InventoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.v1alpha0.InventoryService",
	HandlerType: (*InventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSupplier",
			Handler:    _InventoryService_AddSupplier_Handler,
		},
		{
			MethodName: "DeleteSupplier",
			Handler:    _InventoryService_DeleteSupplier_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory/v1alpha0/inventory.proto",
}