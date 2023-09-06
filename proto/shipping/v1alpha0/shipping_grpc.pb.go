// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: shipping/v1alpha0/shipping.proto

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
	ShippingService_GetPickTicket_FullMethodName = "/shipping.v1alpha0.ShippingService/GetPickTicket"
)

// ShippingServiceClient is the client API for ShippingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ShippingServiceClient interface {
	GetPickTicket(ctx context.Context, in *GetPickTicketRequest, opts ...grpc.CallOption) (*GetPickTicketResponse, error)
}

type shippingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewShippingServiceClient(cc grpc.ClientConnInterface) ShippingServiceClient {
	return &shippingServiceClient{cc}
}

func (c *shippingServiceClient) GetPickTicket(ctx context.Context, in *GetPickTicketRequest, opts ...grpc.CallOption) (*GetPickTicketResponse, error) {
	out := new(GetPickTicketResponse)
	err := c.cc.Invoke(ctx, ShippingService_GetPickTicket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShippingServiceServer is the server API for ShippingService service.
// All implementations should embed UnimplementedShippingServiceServer
// for forward compatibility
type ShippingServiceServer interface {
	GetPickTicket(context.Context, *GetPickTicketRequest) (*GetPickTicketResponse, error)
}

// UnimplementedShippingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedShippingServiceServer struct {
}

func (UnimplementedShippingServiceServer) GetPickTicket(context.Context, *GetPickTicketRequest) (*GetPickTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPickTicket not implemented")
}

// UnsafeShippingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ShippingServiceServer will
// result in compilation errors.
type UnsafeShippingServiceServer interface {
	mustEmbedUnimplementedShippingServiceServer()
}

func RegisterShippingServiceServer(s grpc.ServiceRegistrar, srv ShippingServiceServer) {
	s.RegisterService(&ShippingService_ServiceDesc, srv)
}

func _ShippingService_GetPickTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPickTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShippingServiceServer).GetPickTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ShippingService_GetPickTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShippingServiceServer).GetPickTicket(ctx, req.(*GetPickTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ShippingService_ServiceDesc is the grpc.ServiceDesc for ShippingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ShippingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "shipping.v1alpha0.ShippingService",
	HandlerType: (*ShippingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPickTicket",
			Handler:    _ShippingService_GetPickTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shipping/v1alpha0/shipping.proto",
}