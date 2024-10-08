// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: vault/hcp_link/proto/meta/meta.proto

package meta

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	HCPLinkMeta_ListNamespaces_FullMethodName   = "/meta.HCPLinkMeta/ListNamespaces"
	HCPLinkMeta_ListMounts_FullMethodName       = "/meta.HCPLinkMeta/ListMounts"
	HCPLinkMeta_ListAuths_FullMethodName        = "/meta.HCPLinkMeta/ListAuths"
	HCPLinkMeta_GetClusterStatus_FullMethodName = "/meta.HCPLinkMeta/GetClusterStatus"
)

// HCPLinkMetaClient is the client API for HCPLinkMeta service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HCPLinkMetaClient interface {
	// ListNamespaces will be used to recursively list all namespaces
	ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error)
	// ListMounts will be used to recursively list all mounts in all namespaces
	ListMounts(ctx context.Context, in *ListMountsRequest, opts ...grpc.CallOption) (*ListMountsResponse, error)
	// ListAuths will be used to recursively list all auths in all namespaces
	ListAuths(ctx context.Context, in *ListAuthsRequest, opts ...grpc.CallOption) (*ListAuthResponse, error)
	// GetClusterStatus will provide various cluster-level information
	GetClusterStatus(ctx context.Context, in *GetClusterStatusRequest, opts ...grpc.CallOption) (*GetClusterStatusResponse, error)
}

type hCPLinkMetaClient struct {
	cc grpc.ClientConnInterface
}

func NewHCPLinkMetaClient(cc grpc.ClientConnInterface) HCPLinkMetaClient {
	return &hCPLinkMetaClient{cc}
}

func (c *hCPLinkMetaClient) ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNamespacesResponse)
	err := c.cc.Invoke(ctx, HCPLinkMeta_ListNamespaces_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hCPLinkMetaClient) ListMounts(ctx context.Context, in *ListMountsRequest, opts ...grpc.CallOption) (*ListMountsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListMountsResponse)
	err := c.cc.Invoke(ctx, HCPLinkMeta_ListMounts_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hCPLinkMetaClient) ListAuths(ctx context.Context, in *ListAuthsRequest, opts ...grpc.CallOption) (*ListAuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAuthResponse)
	err := c.cc.Invoke(ctx, HCPLinkMeta_ListAuths_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hCPLinkMetaClient) GetClusterStatus(ctx context.Context, in *GetClusterStatusRequest, opts ...grpc.CallOption) (*GetClusterStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetClusterStatusResponse)
	err := c.cc.Invoke(ctx, HCPLinkMeta_GetClusterStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HCPLinkMetaServer is the server API for HCPLinkMeta service.
// All implementations must embed UnimplementedHCPLinkMetaServer
// for forward compatibility.
type HCPLinkMetaServer interface {
	// ListNamespaces will be used to recursively list all namespaces
	ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error)
	// ListMounts will be used to recursively list all mounts in all namespaces
	ListMounts(context.Context, *ListMountsRequest) (*ListMountsResponse, error)
	// ListAuths will be used to recursively list all auths in all namespaces
	ListAuths(context.Context, *ListAuthsRequest) (*ListAuthResponse, error)
	// GetClusterStatus will provide various cluster-level information
	GetClusterStatus(context.Context, *GetClusterStatusRequest) (*GetClusterStatusResponse, error)
	mustEmbedUnimplementedHCPLinkMetaServer()
}

// UnimplementedHCPLinkMetaServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHCPLinkMetaServer struct{}

func (UnimplementedHCPLinkMetaServer) ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNamespaces not implemented")
}
func (UnimplementedHCPLinkMetaServer) ListMounts(context.Context, *ListMountsRequest) (*ListMountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMounts not implemented")
}
func (UnimplementedHCPLinkMetaServer) ListAuths(context.Context, *ListAuthsRequest) (*ListAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuths not implemented")
}
func (UnimplementedHCPLinkMetaServer) GetClusterStatus(context.Context, *GetClusterStatusRequest) (*GetClusterStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClusterStatus not implemented")
}
func (UnimplementedHCPLinkMetaServer) mustEmbedUnimplementedHCPLinkMetaServer() {}
func (UnimplementedHCPLinkMetaServer) testEmbeddedByValue()                     {}

// UnsafeHCPLinkMetaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HCPLinkMetaServer will
// result in compilation errors.
type UnsafeHCPLinkMetaServer interface {
	mustEmbedUnimplementedHCPLinkMetaServer()
}

func RegisterHCPLinkMetaServer(s grpc.ServiceRegistrar, srv HCPLinkMetaServer) {
	// If the following call pancis, it indicates UnimplementedHCPLinkMetaServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HCPLinkMeta_ServiceDesc, srv)
}

func _HCPLinkMeta_ListNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNamespacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HCPLinkMetaServer).ListNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HCPLinkMeta_ListNamespaces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HCPLinkMetaServer).ListNamespaces(ctx, req.(*ListNamespacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HCPLinkMeta_ListMounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HCPLinkMetaServer).ListMounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HCPLinkMeta_ListMounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HCPLinkMetaServer).ListMounts(ctx, req.(*ListMountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HCPLinkMeta_ListAuths_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuthsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HCPLinkMetaServer).ListAuths(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HCPLinkMeta_ListAuths_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HCPLinkMetaServer).ListAuths(ctx, req.(*ListAuthsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HCPLinkMeta_GetClusterStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HCPLinkMetaServer).GetClusterStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HCPLinkMeta_GetClusterStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HCPLinkMetaServer).GetClusterStatus(ctx, req.(*GetClusterStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HCPLinkMeta_ServiceDesc is the grpc.ServiceDesc for HCPLinkMeta service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HCPLinkMeta_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "meta.HCPLinkMeta",
	HandlerType: (*HCPLinkMetaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNamespaces",
			Handler:    _HCPLinkMeta_ListNamespaces_Handler,
		},
		{
			MethodName: "ListMounts",
			Handler:    _HCPLinkMeta_ListMounts_Handler,
		},
		{
			MethodName: "ListAuths",
			Handler:    _HCPLinkMeta_ListAuths_Handler,
		},
		{
			MethodName: "GetClusterStatus",
			Handler:    _HCPLinkMeta_GetClusterStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vault/hcp_link/proto/meta/meta.proto",
}
