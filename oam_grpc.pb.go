// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: oam.proto

package proto

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

// OAMServerClient is the client API for OAMServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OAMServerClient interface {
	// 服务测试接口
	Ping(ctx context.Context, in *Pong, opts ...grpc.CallOption) (*Pong, error)
	// 获取客户端日志标记
	GetCliLogFlags(ctx context.Context, in *GetCliFlagsRequest, opts ...grpc.CallOption) (*GetCliLogFlagsReply, error)
	// 获取客户端允许升级标记
	GetCliUpFlags(ctx context.Context, in *GetCliFlagsRequest, opts ...grpc.CallOption) (*GetCliUpFlagsReply, error)
	// 接受客户端信息从租户
	AcceptClientList(ctx context.Context, in *AcceptClientListRequest, opts ...grpc.CallOption) (*AcceptClientListReply, error)
	// 获取客户端列表
	ListClient(ctx context.Context, in *ListClientRequest, opts ...grpc.CallOption) (*ListClientReply, error)
	// 获取单个客户端
	GetClient(ctx context.Context, in *ListClientRequest, opts ...grpc.CallOption) (*OAMClient, error)
	// 批量更新客户端
	BatchUpdateClientFlags(ctx context.Context, in *BatchUpdateClientFlagsRequest, opts ...grpc.CallOption) (*UpdateClientReply, error)
}

type oAMServerClient struct {
	cc grpc.ClientConnInterface
}

func NewOAMServerClient(cc grpc.ClientConnInterface) OAMServerClient {
	return &oAMServerClient{cc}
}

func (c *oAMServerClient) Ping(ctx context.Context, in *Pong, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := c.cc.Invoke(ctx, "/OAMServer/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) GetCliLogFlags(ctx context.Context, in *GetCliFlagsRequest, opts ...grpc.CallOption) (*GetCliLogFlagsReply, error) {
	out := new(GetCliLogFlagsReply)
	err := c.cc.Invoke(ctx, "/OAMServer/GetCliLogFlags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) GetCliUpFlags(ctx context.Context, in *GetCliFlagsRequest, opts ...grpc.CallOption) (*GetCliUpFlagsReply, error) {
	out := new(GetCliUpFlagsReply)
	err := c.cc.Invoke(ctx, "/OAMServer/GetCliUpFlags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) AcceptClientList(ctx context.Context, in *AcceptClientListRequest, opts ...grpc.CallOption) (*AcceptClientListReply, error) {
	out := new(AcceptClientListReply)
	err := c.cc.Invoke(ctx, "/OAMServer/AcceptClientList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) ListClient(ctx context.Context, in *ListClientRequest, opts ...grpc.CallOption) (*ListClientReply, error) {
	out := new(ListClientReply)
	err := c.cc.Invoke(ctx, "/OAMServer/ListClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) GetClient(ctx context.Context, in *ListClientRequest, opts ...grpc.CallOption) (*OAMClient, error) {
	out := new(OAMClient)
	err := c.cc.Invoke(ctx, "/OAMServer/GetClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) BatchUpdateClientFlags(ctx context.Context, in *BatchUpdateClientFlagsRequest, opts ...grpc.CallOption) (*UpdateClientReply, error) {
	out := new(UpdateClientReply)
	err := c.cc.Invoke(ctx, "/OAMServer/BatchUpdateClientFlags", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OAMServerServer is the server API for OAMServer service.
// All implementations should embed UnimplementedOAMServerServer
// for forward compatibility
type OAMServerServer interface {
	// 服务测试接口
	Ping(context.Context, *Pong) (*Pong, error)
	// 获取客户端日志标记
	GetCliLogFlags(context.Context, *GetCliFlagsRequest) (*GetCliLogFlagsReply, error)
	// 获取客户端允许升级标记
	GetCliUpFlags(context.Context, *GetCliFlagsRequest) (*GetCliUpFlagsReply, error)
	// 接受客户端信息从租户
	AcceptClientList(context.Context, *AcceptClientListRequest) (*AcceptClientListReply, error)
	// 获取客户端列表
	ListClient(context.Context, *ListClientRequest) (*ListClientReply, error)
	// 获取单个客户端
	GetClient(context.Context, *ListClientRequest) (*OAMClient, error)
	// 批量更新客户端
	BatchUpdateClientFlags(context.Context, *BatchUpdateClientFlagsRequest) (*UpdateClientReply, error)
}

// UnimplementedOAMServerServer should be embedded to have forward compatible implementations.
type UnimplementedOAMServerServer struct {
}

func (UnimplementedOAMServerServer) Ping(context.Context, *Pong) (*Pong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedOAMServerServer) GetCliLogFlags(context.Context, *GetCliFlagsRequest) (*GetCliLogFlagsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCliLogFlags not implemented")
}
func (UnimplementedOAMServerServer) GetCliUpFlags(context.Context, *GetCliFlagsRequest) (*GetCliUpFlagsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCliUpFlags not implemented")
}
func (UnimplementedOAMServerServer) AcceptClientList(context.Context, *AcceptClientListRequest) (*AcceptClientListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptClientList not implemented")
}
func (UnimplementedOAMServerServer) ListClient(context.Context, *ListClientRequest) (*ListClientReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClient not implemented")
}
func (UnimplementedOAMServerServer) GetClient(context.Context, *ListClientRequest) (*OAMClient, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClient not implemented")
}
func (UnimplementedOAMServerServer) BatchUpdateClientFlags(context.Context, *BatchUpdateClientFlagsRequest) (*UpdateClientReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchUpdateClientFlags not implemented")
}

// UnsafeOAMServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OAMServerServer will
// result in compilation errors.
type UnsafeOAMServerServer interface {
	mustEmbedUnimplementedOAMServerServer()
}

func RegisterOAMServerServer(s grpc.ServiceRegistrar, srv OAMServerServer) {
	s.RegisterService(&OAMServer_ServiceDesc, srv)
}

func _OAMServer_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pong)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).Ping(ctx, req.(*Pong))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_GetCliLogFlags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCliFlagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).GetCliLogFlags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/GetCliLogFlags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).GetCliLogFlags(ctx, req.(*GetCliFlagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_GetCliUpFlags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCliFlagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).GetCliUpFlags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/GetCliUpFlags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).GetCliUpFlags(ctx, req.(*GetCliFlagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_AcceptClientList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptClientListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).AcceptClientList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/AcceptClientList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).AcceptClientList(ctx, req.(*AcceptClientListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_ListClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).ListClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/ListClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).ListClient(ctx, req.(*ListClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_GetClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).GetClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/GetClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).GetClient(ctx, req.(*ListClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_BatchUpdateClientFlags_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchUpdateClientFlagsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).BatchUpdateClientFlags(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/BatchUpdateClientFlags",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).BatchUpdateClientFlags(ctx, req.(*BatchUpdateClientFlagsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OAMServer_ServiceDesc is the grpc.ServiceDesc for OAMServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OAMServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OAMServer",
	HandlerType: (*OAMServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _OAMServer_Ping_Handler,
		},
		{
			MethodName: "GetCliLogFlags",
			Handler:    _OAMServer_GetCliLogFlags_Handler,
		},
		{
			MethodName: "GetCliUpFlags",
			Handler:    _OAMServer_GetCliUpFlags_Handler,
		},
		{
			MethodName: "AcceptClientList",
			Handler:    _OAMServer_AcceptClientList_Handler,
		},
		{
			MethodName: "ListClient",
			Handler:    _OAMServer_ListClient_Handler,
		},
		{
			MethodName: "GetClient",
			Handler:    _OAMServer_GetClient_Handler,
		},
		{
			MethodName: "BatchUpdateClientFlags",
			Handler:    _OAMServer_BatchUpdateClientFlags_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oam.proto",
}
