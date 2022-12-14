// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: list_shop.proto

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

// ListShopServiceClient is the client API for ListShopService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ListShopServiceClient interface {
	ListShop(ctx context.Context, in *ListShopReq, opts ...grpc.CallOption) (ListShopService_ListShopClient, error)
}

type listShopServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewListShopServiceClient(cc grpc.ClientConnInterface) ListShopServiceClient {
	return &listShopServiceClient{cc}
}

func (c *listShopServiceClient) ListShop(ctx context.Context, in *ListShopReq, opts ...grpc.CallOption) (ListShopService_ListShopClient, error) {
	stream, err := c.cc.NewStream(ctx, &ListShopService_ServiceDesc.Streams[0], "/pb.ListShopService/ListShop", opts...)
	if err != nil {
		return nil, err
	}
	x := &listShopServiceListShopClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ListShopService_ListShopClient interface {
	Recv() (*ListShopRes, error)
	grpc.ClientStream
}

type listShopServiceListShopClient struct {
	grpc.ClientStream
}

func (x *listShopServiceListShopClient) Recv() (*ListShopRes, error) {
	m := new(ListShopRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ListShopServiceServer is the server API for ListShopService service.
// All implementations must embed UnimplementedListShopServiceServer
// for forward compatibility
type ListShopServiceServer interface {
	ListShop(*ListShopReq, ListShopService_ListShopServer) error
	mustEmbedUnimplementedListShopServiceServer()
}

// UnimplementedListShopServiceServer must be embedded to have forward compatible implementations.
type UnimplementedListShopServiceServer struct {
}

func (UnimplementedListShopServiceServer) ListShop(*ListShopReq, ListShopService_ListShopServer) error {
	return status.Errorf(codes.Unimplemented, "method ListShop not implemented")
}
func (UnimplementedListShopServiceServer) mustEmbedUnimplementedListShopServiceServer() {}

// UnsafeListShopServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ListShopServiceServer will
// result in compilation errors.
type UnsafeListShopServiceServer interface {
	mustEmbedUnimplementedListShopServiceServer()
}

func RegisterListShopServiceServer(s grpc.ServiceRegistrar, srv ListShopServiceServer) {
	s.RegisterService(&ListShopService_ServiceDesc, srv)
}

func _ListShopService_ListShop_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListShopReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ListShopServiceServer).ListShop(m, &listShopServiceListShopServer{stream})
}

type ListShopService_ListShopServer interface {
	Send(*ListShopRes) error
	grpc.ServerStream
}

type listShopServiceListShopServer struct {
	grpc.ServerStream
}

func (x *listShopServiceListShopServer) Send(m *ListShopRes) error {
	return x.ServerStream.SendMsg(m)
}

// ListShopService_ServiceDesc is the grpc.ServiceDesc for ListShopService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ListShopService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ListShopService",
	HandlerType: (*ListShopServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListShop",
			Handler:       _ListShopService_ListShop_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "list_shop.proto",
}
