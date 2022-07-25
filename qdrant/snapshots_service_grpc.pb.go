// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: snapshots_service.proto

package go_client

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

// SnapshotsClient is the client API for Snapshots service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SnapshotsClient interface {
	//
	//Create collection snapshot
	Create(ctx context.Context, in *CreateSnapshotRequest, opts ...grpc.CallOption) (*CreateSnapshotResponse, error)
	//
	//List collection snapshots
	List(ctx context.Context, in *ListSnapshotsRequest, opts ...grpc.CallOption) (*ListSnapshotsResponse, error)
	//
	//Create full storage snapshot
	CreateFull(ctx context.Context, in *CreateFullSnapshotRequest, opts ...grpc.CallOption) (*CreateSnapshotResponse, error)
	//
	//List full storage snapshots
	ListFull(ctx context.Context, in *ListFullSnapshotsRequest, opts ...grpc.CallOption) (*ListSnapshotsResponse, error)
}

type snapshotsClient struct {
	cc grpc.ClientConnInterface
}

func NewSnapshotsClient(cc grpc.ClientConnInterface) SnapshotsClient {
	return &snapshotsClient{cc}
}

func (c *snapshotsClient) Create(ctx context.Context, in *CreateSnapshotRequest, opts ...grpc.CallOption) (*CreateSnapshotResponse, error) {
	out := new(CreateSnapshotResponse)
	err := c.cc.Invoke(ctx, "/qdrant.Snapshots/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsClient) List(ctx context.Context, in *ListSnapshotsRequest, opts ...grpc.CallOption) (*ListSnapshotsResponse, error) {
	out := new(ListSnapshotsResponse)
	err := c.cc.Invoke(ctx, "/qdrant.Snapshots/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsClient) CreateFull(ctx context.Context, in *CreateFullSnapshotRequest, opts ...grpc.CallOption) (*CreateSnapshotResponse, error) {
	out := new(CreateSnapshotResponse)
	err := c.cc.Invoke(ctx, "/qdrant.Snapshots/CreateFull", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snapshotsClient) ListFull(ctx context.Context, in *ListFullSnapshotsRequest, opts ...grpc.CallOption) (*ListSnapshotsResponse, error) {
	out := new(ListSnapshotsResponse)
	err := c.cc.Invoke(ctx, "/qdrant.Snapshots/ListFull", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnapshotsServer is the server API for Snapshots service.
// All implementations must embed UnimplementedSnapshotsServer
// for forward compatibility
type SnapshotsServer interface {
	//
	//Create collection snapshot
	Create(context.Context, *CreateSnapshotRequest) (*CreateSnapshotResponse, error)
	//
	//List collection snapshots
	List(context.Context, *ListSnapshotsRequest) (*ListSnapshotsResponse, error)
	//
	//Create full storage snapshot
	CreateFull(context.Context, *CreateFullSnapshotRequest) (*CreateSnapshotResponse, error)
	//
	//List full storage snapshots
	ListFull(context.Context, *ListFullSnapshotsRequest) (*ListSnapshotsResponse, error)
	mustEmbedUnimplementedSnapshotsServer()
}

// UnimplementedSnapshotsServer must be embedded to have forward compatible implementations.
type UnimplementedSnapshotsServer struct {
}

func (UnimplementedSnapshotsServer) Create(context.Context, *CreateSnapshotRequest) (*CreateSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSnapshotsServer) List(context.Context, *ListSnapshotsRequest) (*ListSnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSnapshotsServer) CreateFull(context.Context, *CreateFullSnapshotRequest) (*CreateSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFull not implemented")
}
func (UnimplementedSnapshotsServer) ListFull(context.Context, *ListFullSnapshotsRequest) (*ListSnapshotsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFull not implemented")
}
func (UnimplementedSnapshotsServer) mustEmbedUnimplementedSnapshotsServer() {}

// UnsafeSnapshotsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SnapshotsServer will
// result in compilation errors.
type UnsafeSnapshotsServer interface {
	mustEmbedUnimplementedSnapshotsServer()
}

func RegisterSnapshotsServer(s grpc.ServiceRegistrar, srv SnapshotsServer) {
	s.RegisterService(&Snapshots_ServiceDesc, srv)
}

func _Snapshots_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qdrant.Snapshots/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).Create(ctx, req.(*CreateSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshots_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qdrant.Snapshots/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).List(ctx, req.(*ListSnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshots_CreateFull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFullSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).CreateFull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qdrant.Snapshots/CreateFull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).CreateFull(ctx, req.(*CreateFullSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Snapshots_ListFull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFullSnapshotsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnapshotsServer).ListFull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qdrant.Snapshots/ListFull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnapshotsServer).ListFull(ctx, req.(*ListFullSnapshotsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Snapshots_ServiceDesc is the grpc.ServiceDesc for Snapshots service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Snapshots_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "qdrant.Snapshots",
	HandlerType: (*SnapshotsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Snapshots_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Snapshots_List_Handler,
		},
		{
			MethodName: "CreateFull",
			Handler:    _Snapshots_CreateFull_Handler,
		},
		{
			MethodName: "ListFull",
			Handler:    _Snapshots_ListFull_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "snapshots_service.proto",
}
