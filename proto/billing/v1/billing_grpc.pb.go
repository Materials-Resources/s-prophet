// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: billing/v1/billing.proto

package billing_v1

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
	BillingService_GetInvoicesByOrder_FullMethodName = "/billing.v1.BillingService/GetInvoicesByOrder"
)

// BillingServiceClient is the client API for BillingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BillingServiceClient interface {
	// GetInvoicesByOrder returns all invoices for a given order
	GetInvoicesByOrder(ctx context.Context, in *GetInvoicesByOrderRequest, opts ...grpc.CallOption) (*GetInvoicesByOrderResponse, error)
}

type billingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBillingServiceClient(cc grpc.ClientConnInterface) BillingServiceClient {
	return &billingServiceClient{cc}
}

func (c *billingServiceClient) GetInvoicesByOrder(ctx context.Context, in *GetInvoicesByOrderRequest, opts ...grpc.CallOption) (*GetInvoicesByOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetInvoicesByOrderResponse)
	err := c.cc.Invoke(ctx, BillingService_GetInvoicesByOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillingServiceServer is the server API for BillingService service.
// All implementations should embed UnimplementedBillingServiceServer
// for forward compatibility
type BillingServiceServer interface {
	// GetInvoicesByOrder returns all invoices for a given order
	GetInvoicesByOrder(context.Context, *GetInvoicesByOrderRequest) (*GetInvoicesByOrderResponse, error)
}

// UnimplementedBillingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBillingServiceServer struct {
}

func (UnimplementedBillingServiceServer) GetInvoicesByOrder(context.Context, *GetInvoicesByOrderRequest) (*GetInvoicesByOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInvoicesByOrder not implemented")
}

// UnsafeBillingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BillingServiceServer will
// result in compilation errors.
type UnsafeBillingServiceServer interface {
	mustEmbedUnimplementedBillingServiceServer()
}

func RegisterBillingServiceServer(s grpc.ServiceRegistrar, srv BillingServiceServer) {
	s.RegisterService(&BillingService_ServiceDesc, srv)
}

func _BillingService_GetInvoicesByOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInvoicesByOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServiceServer).GetInvoicesByOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillingService_GetInvoicesByOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServiceServer).GetInvoicesByOrder(ctx, req.(*GetInvoicesByOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BillingService_ServiceDesc is the grpc.ServiceDesc for BillingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BillingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "billing.v1.BillingService",
	HandlerType: (*BillingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInvoicesByOrder",
			Handler:    _BillingService_GetInvoicesByOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "billing/v1/billing.proto",
}
