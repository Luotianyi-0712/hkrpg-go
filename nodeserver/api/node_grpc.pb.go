// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.0
// - protoc             v5.26.0
// source: node.proto

package proto

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
	NodeDiscovery_Test_FullMethodName               = "/proto.NodeDiscovery/Test"
	NodeDiscovery_GetAllGateServerMq_FullMethodName = "/proto.NodeDiscovery/GetAllGateServerMq"
)

// NodeDiscoveryClient is the client API for NodeDiscovery service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 节点服务器注册发现服务
type NodeDiscoveryClient interface {
	// 测试
	Test(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestRsp, error)
	// 获取全部gate mq
	GetAllGateServerMq(ctx context.Context, in *GetAllGateServerMqReq, opts ...grpc.CallOption) (*GetAllGateServerMqRsp, error)
}

type nodeDiscoveryClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeDiscoveryClient(cc grpc.ClientConnInterface) NodeDiscoveryClient {
	return &nodeDiscoveryClient{cc}
}

func (c *nodeDiscoveryClient) Test(ctx context.Context, in *TestReq, opts ...grpc.CallOption) (*TestRsp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TestRsp)
	err := c.cc.Invoke(ctx, NodeDiscovery_Test_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeDiscoveryClient) GetAllGateServerMq(ctx context.Context, in *GetAllGateServerMqReq, opts ...grpc.CallOption) (*GetAllGateServerMqRsp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllGateServerMqRsp)
	err := c.cc.Invoke(ctx, NodeDiscovery_GetAllGateServerMq_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeDiscoveryServer is the server API for NodeDiscovery service.
// All implementations must embed UnimplementedNodeDiscoveryServer
// for forward compatibility.
//
// 节点服务器注册发现服务
type NodeDiscoveryServer interface {
	// 测试
	Test(context.Context, *TestReq) (*TestRsp, error)
	// 获取全部gate mq
	GetAllGateServerMq(context.Context, *GetAllGateServerMqReq) (*GetAllGateServerMqRsp, error)
	mustEmbedUnimplementedNodeDiscoveryServer()
}

// UnimplementedNodeDiscoveryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNodeDiscoveryServer struct{}

func (UnimplementedNodeDiscoveryServer) Test(context.Context, *TestReq) (*TestRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Test not implemented")
}
func (UnimplementedNodeDiscoveryServer) GetAllGateServerMq(context.Context, *GetAllGateServerMqReq) (*GetAllGateServerMqRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllGateServerMq not implemented")
}
func (UnimplementedNodeDiscoveryServer) mustEmbedUnimplementedNodeDiscoveryServer() {}
func (UnimplementedNodeDiscoveryServer) testEmbeddedByValue()                       {}

// UnsafeNodeDiscoveryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NodeDiscoveryServer will
// result in compilation errors.
type UnsafeNodeDiscoveryServer interface {
	mustEmbedUnimplementedNodeDiscoveryServer()
}

func RegisterNodeDiscoveryServer(s grpc.ServiceRegistrar, srv NodeDiscoveryServer) {
	// If the following call pancis, it indicates UnimplementedNodeDiscoveryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NodeDiscovery_ServiceDesc, srv)
}

func _NodeDiscovery_Test_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeDiscoveryServer).Test(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeDiscovery_Test_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeDiscoveryServer).Test(ctx, req.(*TestReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeDiscovery_GetAllGateServerMq_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllGateServerMqReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeDiscoveryServer).GetAllGateServerMq(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NodeDiscovery_GetAllGateServerMq_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeDiscoveryServer).GetAllGateServerMq(ctx, req.(*GetAllGateServerMqReq))
	}
	return interceptor(ctx, in, info, handler)
}

// NodeDiscovery_ServiceDesc is the grpc.ServiceDesc for NodeDiscovery service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NodeDiscovery_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.NodeDiscovery",
	HandlerType: (*NodeDiscoveryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Test",
			Handler:    _NodeDiscovery_Test_Handler,
		},
		{
			MethodName: "GetAllGateServerMq",
			Handler:    _NodeDiscovery_GetAllGateServerMq_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}
