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
	// 策略下发:先放在这里
	ReleasePolicy(ctx context.Context, in *ReleasePolicyRequest, opts ...grpc.CallOption) (*ListPolicyReply, error)
	// 新增策略
	AddPolicy(ctx context.Context, in *AddPolicyRequest, opts ...grpc.CallOption) (*OperateReply, error)
	// 修改
	PutPolicy(ctx context.Context, in *PutPolicyRequest, opts ...grpc.CallOption) (*OperateReply, error)
	// 删除
	DelPolicy(ctx context.Context, in *PolicyQuery, opts ...grpc.CallOption) (*OperateReply, error)
	// 获取列表
	ListPolicy(ctx context.Context, in *PolicyQuery, opts ...grpc.CallOption) (*ListPolicyReply, error)
	// 获取单个
	GetPolicy(ctx context.Context, in *PolicyQuery, opts ...grpc.CallOption) (*GetPolicyReply, error)
	// 获取modelVersion
	ListModelVersion(ctx context.Context, in *ModelVersion, opts ...grpc.CallOption) (*ModelVersion, error)
	// 接受客户端信息从租户
	AcceptTenantList(ctx context.Context, in *AcceptTenantListRequest, opts ...grpc.CallOption) (*AcceptTenantListReply, error)
	// 获取租户所属平台的域名
	GetServerIp(ctx context.Context, in *GetPlatRequest, opts ...grpc.CallOption) (*GetServerIpReply, error)
	// 创建平台账号
	AddPlatAccount(ctx context.Context, in *AddPlatRequest, opts ...grpc.CallOption) (*PlatReply, error)
	// 添加平台授权
	AddPlatLicense(ctx context.Context, in *AddPlatLicenseRequest, opts ...grpc.CallOption) (*PlatReply, error)
	// 购买平台授权点数
	BuyPlatLicenseCount(ctx context.Context, in *BuyPlatLicenseCountRequest, opts ...grpc.CallOption) (*PlatReply, error)
	// 发送激活验证码
	SendActiveCode(ctx context.Context, in *ActivePlatInfoRequest, opts ...grpc.CallOption) (*PlatReply, error)
	// 激活平台
	ActivePlatInfo(ctx context.Context, in *ActivePlatInfoRequest, opts ...grpc.CallOption) (*ActivePlatInfoReply, error)
	// 获取平台信息
	GetPlatInfo(ctx context.Context, in *GetPlatInfoRequest, opts ...grpc.CallOption) (*GetPlatInfoReply, error)
	// 获取平台授权信息
	GetPlatLicenseInfo(ctx context.Context, in *GetPlatLicenseInfoRequest, opts ...grpc.CallOption) (*GetPlatLicenseInfoReply, error)
	// 更新平台信息
	UpdatePlatInfo(ctx context.Context, in *UpdatePlatInfoRequest, opts ...grpc.CallOption) (*UpdatePlatInfoReply, error)
	// 更新平台授权信息
	UpdatePlatLicenseInfo(ctx context.Context, in *UpdatePlatLicenseInfoRequest, opts ...grpc.CallOption) (*UpdatePlatLicenseInfoReply, error)
	// 获取平台域名
	GetPlatDomain(ctx context.Context, in *GetPlatDomainRequest, opts ...grpc.CallOption) (*GetPlatDomainReply, error)
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

func (c *oAMServerClient) ReleasePolicy(ctx context.Context, in *ReleasePolicyRequest, opts ...grpc.CallOption) (*ListPolicyReply, error) {
	out := new(ListPolicyReply)
	err := c.cc.Invoke(ctx, "/OAMServer/ReleasePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) AddPolicy(ctx context.Context, in *AddPolicyRequest, opts ...grpc.CallOption) (*OperateReply, error) {
	out := new(OperateReply)
	err := c.cc.Invoke(ctx, "/OAMServer/AddPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) PutPolicy(ctx context.Context, in *PutPolicyRequest, opts ...grpc.CallOption) (*OperateReply, error) {
	out := new(OperateReply)
	err := c.cc.Invoke(ctx, "/OAMServer/PutPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) DelPolicy(ctx context.Context, in *PolicyQuery, opts ...grpc.CallOption) (*OperateReply, error) {
	out := new(OperateReply)
	err := c.cc.Invoke(ctx, "/OAMServer/DelPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) ListPolicy(ctx context.Context, in *PolicyQuery, opts ...grpc.CallOption) (*ListPolicyReply, error) {
	out := new(ListPolicyReply)
	err := c.cc.Invoke(ctx, "/OAMServer/ListPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) GetPolicy(ctx context.Context, in *PolicyQuery, opts ...grpc.CallOption) (*GetPolicyReply, error) {
	out := new(GetPolicyReply)
	err := c.cc.Invoke(ctx, "/OAMServer/GetPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) ListModelVersion(ctx context.Context, in *ModelVersion, opts ...grpc.CallOption) (*ModelVersion, error) {
	out := new(ModelVersion)
	err := c.cc.Invoke(ctx, "/OAMServer/ListModelVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) AcceptTenantList(ctx context.Context, in *AcceptTenantListRequest, opts ...grpc.CallOption) (*AcceptTenantListReply, error) {
	out := new(AcceptTenantListReply)
	err := c.cc.Invoke(ctx, "/OAMServer/AcceptTenantList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) GetServerIp(ctx context.Context, in *GetPlatRequest, opts ...grpc.CallOption) (*GetServerIpReply, error) {
	out := new(GetServerIpReply)
	err := c.cc.Invoke(ctx, "/OAMServer/GetServerIp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) AddPlatAccount(ctx context.Context, in *AddPlatRequest, opts ...grpc.CallOption) (*PlatReply, error) {
	out := new(PlatReply)
	err := c.cc.Invoke(ctx, "/OAMServer/AddPlatAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) AddPlatLicense(ctx context.Context, in *AddPlatLicenseRequest, opts ...grpc.CallOption) (*PlatReply, error) {
	out := new(PlatReply)
	err := c.cc.Invoke(ctx, "/OAMServer/AddPlatLicense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) BuyPlatLicenseCount(ctx context.Context, in *BuyPlatLicenseCountRequest, opts ...grpc.CallOption) (*PlatReply, error) {
	out := new(PlatReply)
	err := c.cc.Invoke(ctx, "/OAMServer/BuyPlatLicenseCount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) SendActiveCode(ctx context.Context, in *ActivePlatInfoRequest, opts ...grpc.CallOption) (*PlatReply, error) {
	out := new(PlatReply)
	err := c.cc.Invoke(ctx, "/OAMServer/SendActiveCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) ActivePlatInfo(ctx context.Context, in *ActivePlatInfoRequest, opts ...grpc.CallOption) (*ActivePlatInfoReply, error) {
	out := new(ActivePlatInfoReply)
	err := c.cc.Invoke(ctx, "/OAMServer/ActivePlatInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) GetPlatInfo(ctx context.Context, in *GetPlatInfoRequest, opts ...grpc.CallOption) (*GetPlatInfoReply, error) {
	out := new(GetPlatInfoReply)
	err := c.cc.Invoke(ctx, "/OAMServer/GetPlatInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) GetPlatLicenseInfo(ctx context.Context, in *GetPlatLicenseInfoRequest, opts ...grpc.CallOption) (*GetPlatLicenseInfoReply, error) {
	out := new(GetPlatLicenseInfoReply)
	err := c.cc.Invoke(ctx, "/OAMServer/GetPlatLicenseInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) UpdatePlatInfo(ctx context.Context, in *UpdatePlatInfoRequest, opts ...grpc.CallOption) (*UpdatePlatInfoReply, error) {
	out := new(UpdatePlatInfoReply)
	err := c.cc.Invoke(ctx, "/OAMServer/UpdatePlatInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) UpdatePlatLicenseInfo(ctx context.Context, in *UpdatePlatLicenseInfoRequest, opts ...grpc.CallOption) (*UpdatePlatLicenseInfoReply, error) {
	out := new(UpdatePlatLicenseInfoReply)
	err := c.cc.Invoke(ctx, "/OAMServer/UpdatePlatLicenseInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oAMServerClient) GetPlatDomain(ctx context.Context, in *GetPlatDomainRequest, opts ...grpc.CallOption) (*GetPlatDomainReply, error) {
	out := new(GetPlatDomainReply)
	err := c.cc.Invoke(ctx, "/OAMServer/GetPlatDomain", in, out, opts...)
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
	// 策略下发:先放在这里
	ReleasePolicy(context.Context, *ReleasePolicyRequest) (*ListPolicyReply, error)
	// 新增策略
	AddPolicy(context.Context, *AddPolicyRequest) (*OperateReply, error)
	// 修改
	PutPolicy(context.Context, *PutPolicyRequest) (*OperateReply, error)
	// 删除
	DelPolicy(context.Context, *PolicyQuery) (*OperateReply, error)
	// 获取列表
	ListPolicy(context.Context, *PolicyQuery) (*ListPolicyReply, error)
	// 获取单个
	GetPolicy(context.Context, *PolicyQuery) (*GetPolicyReply, error)
	// 获取modelVersion
	ListModelVersion(context.Context, *ModelVersion) (*ModelVersion, error)
	// 接受客户端信息从租户
	AcceptTenantList(context.Context, *AcceptTenantListRequest) (*AcceptTenantListReply, error)
	// 获取租户所属平台的域名
	GetServerIp(context.Context, *GetPlatRequest) (*GetServerIpReply, error)
	// 创建平台账号
	AddPlatAccount(context.Context, *AddPlatRequest) (*PlatReply, error)
	// 添加平台授权
	AddPlatLicense(context.Context, *AddPlatLicenseRequest) (*PlatReply, error)
	// 购买平台授权点数
	BuyPlatLicenseCount(context.Context, *BuyPlatLicenseCountRequest) (*PlatReply, error)
	// 发送激活验证码
	SendActiveCode(context.Context, *ActivePlatInfoRequest) (*PlatReply, error)
	// 激活平台
	ActivePlatInfo(context.Context, *ActivePlatInfoRequest) (*ActivePlatInfoReply, error)
	// 获取平台信息
	GetPlatInfo(context.Context, *GetPlatInfoRequest) (*GetPlatInfoReply, error)
	// 获取平台授权信息
	GetPlatLicenseInfo(context.Context, *GetPlatLicenseInfoRequest) (*GetPlatLicenseInfoReply, error)
	// 更新平台信息
	UpdatePlatInfo(context.Context, *UpdatePlatInfoRequest) (*UpdatePlatInfoReply, error)
	// 更新平台授权信息
	UpdatePlatLicenseInfo(context.Context, *UpdatePlatLicenseInfoRequest) (*UpdatePlatLicenseInfoReply, error)
	// 获取平台域名
	GetPlatDomain(context.Context, *GetPlatDomainRequest) (*GetPlatDomainReply, error)
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
func (UnimplementedOAMServerServer) ReleasePolicy(context.Context, *ReleasePolicyRequest) (*ListPolicyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleasePolicy not implemented")
}
func (UnimplementedOAMServerServer) AddPolicy(context.Context, *AddPolicyRequest) (*OperateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPolicy not implemented")
}
func (UnimplementedOAMServerServer) PutPolicy(context.Context, *PutPolicyRequest) (*OperateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutPolicy not implemented")
}
func (UnimplementedOAMServerServer) DelPolicy(context.Context, *PolicyQuery) (*OperateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelPolicy not implemented")
}
func (UnimplementedOAMServerServer) ListPolicy(context.Context, *PolicyQuery) (*ListPolicyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPolicy not implemented")
}
func (UnimplementedOAMServerServer) GetPolicy(context.Context, *PolicyQuery) (*GetPolicyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPolicy not implemented")
}
func (UnimplementedOAMServerServer) ListModelVersion(context.Context, *ModelVersion) (*ModelVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListModelVersion not implemented")
}
func (UnimplementedOAMServerServer) AcceptTenantList(context.Context, *AcceptTenantListRequest) (*AcceptTenantListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptTenantList not implemented")
}
func (UnimplementedOAMServerServer) GetServerIp(context.Context, *GetPlatRequest) (*GetServerIpReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerIp not implemented")
}
func (UnimplementedOAMServerServer) AddPlatAccount(context.Context, *AddPlatRequest) (*PlatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPlatAccount not implemented")
}
func (UnimplementedOAMServerServer) AddPlatLicense(context.Context, *AddPlatLicenseRequest) (*PlatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPlatLicense not implemented")
}
func (UnimplementedOAMServerServer) BuyPlatLicenseCount(context.Context, *BuyPlatLicenseCountRequest) (*PlatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyPlatLicenseCount not implemented")
}
func (UnimplementedOAMServerServer) SendActiveCode(context.Context, *ActivePlatInfoRequest) (*PlatReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendActiveCode not implemented")
}
func (UnimplementedOAMServerServer) ActivePlatInfo(context.Context, *ActivePlatInfoRequest) (*ActivePlatInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivePlatInfo not implemented")
}
func (UnimplementedOAMServerServer) GetPlatInfo(context.Context, *GetPlatInfoRequest) (*GetPlatInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlatInfo not implemented")
}
func (UnimplementedOAMServerServer) GetPlatLicenseInfo(context.Context, *GetPlatLicenseInfoRequest) (*GetPlatLicenseInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlatLicenseInfo not implemented")
}
func (UnimplementedOAMServerServer) UpdatePlatInfo(context.Context, *UpdatePlatInfoRequest) (*UpdatePlatInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlatInfo not implemented")
}
func (UnimplementedOAMServerServer) UpdatePlatLicenseInfo(context.Context, *UpdatePlatLicenseInfoRequest) (*UpdatePlatLicenseInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlatLicenseInfo not implemented")
}
func (UnimplementedOAMServerServer) GetPlatDomain(context.Context, *GetPlatDomainRequest) (*GetPlatDomainReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlatDomain not implemented")
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

func _OAMServer_ReleasePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleasePolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).ReleasePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/ReleasePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).ReleasePolicy(ctx, req.(*ReleasePolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_AddPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).AddPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/AddPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).AddPolicy(ctx, req.(*AddPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_PutPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutPolicyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).PutPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/PutPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).PutPolicy(ctx, req.(*PutPolicyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_DelPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).DelPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/DelPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).DelPolicy(ctx, req.(*PolicyQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_ListPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).ListPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/ListPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).ListPolicy(ctx, req.(*PolicyQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_GetPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PolicyQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).GetPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/GetPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).GetPolicy(ctx, req.(*PolicyQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_ListModelVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelVersion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).ListModelVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/ListModelVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).ListModelVersion(ctx, req.(*ModelVersion))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_AcceptTenantList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AcceptTenantListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).AcceptTenantList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/AcceptTenantList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).AcceptTenantList(ctx, req.(*AcceptTenantListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_GetServerIp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).GetServerIp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/GetServerIp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).GetServerIp(ctx, req.(*GetPlatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_AddPlatAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPlatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).AddPlatAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/AddPlatAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).AddPlatAccount(ctx, req.(*AddPlatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_AddPlatLicense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPlatLicenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).AddPlatLicense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/AddPlatLicense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).AddPlatLicense(ctx, req.(*AddPlatLicenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_BuyPlatLicenseCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyPlatLicenseCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).BuyPlatLicenseCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/BuyPlatLicenseCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).BuyPlatLicenseCount(ctx, req.(*BuyPlatLicenseCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_SendActiveCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivePlatInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).SendActiveCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/SendActiveCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).SendActiveCode(ctx, req.(*ActivePlatInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_ActivePlatInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivePlatInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).ActivePlatInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/ActivePlatInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).ActivePlatInfo(ctx, req.(*ActivePlatInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_GetPlatInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlatInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).GetPlatInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/GetPlatInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).GetPlatInfo(ctx, req.(*GetPlatInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_GetPlatLicenseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlatLicenseInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).GetPlatLicenseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/GetPlatLicenseInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).GetPlatLicenseInfo(ctx, req.(*GetPlatLicenseInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_UpdatePlatInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlatInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).UpdatePlatInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/UpdatePlatInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).UpdatePlatInfo(ctx, req.(*UpdatePlatInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_UpdatePlatLicenseInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlatLicenseInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).UpdatePlatLicenseInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/UpdatePlatLicenseInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).UpdatePlatLicenseInfo(ctx, req.(*UpdatePlatLicenseInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OAMServer_GetPlatDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlatDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OAMServerServer).GetPlatDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/OAMServer/GetPlatDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OAMServerServer).GetPlatDomain(ctx, req.(*GetPlatDomainRequest))
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
		{
			MethodName: "ReleasePolicy",
			Handler:    _OAMServer_ReleasePolicy_Handler,
		},
		{
			MethodName: "AddPolicy",
			Handler:    _OAMServer_AddPolicy_Handler,
		},
		{
			MethodName: "PutPolicy",
			Handler:    _OAMServer_PutPolicy_Handler,
		},
		{
			MethodName: "DelPolicy",
			Handler:    _OAMServer_DelPolicy_Handler,
		},
		{
			MethodName: "ListPolicy",
			Handler:    _OAMServer_ListPolicy_Handler,
		},
		{
			MethodName: "GetPolicy",
			Handler:    _OAMServer_GetPolicy_Handler,
		},
		{
			MethodName: "ListModelVersion",
			Handler:    _OAMServer_ListModelVersion_Handler,
		},
		{
			MethodName: "AcceptTenantList",
			Handler:    _OAMServer_AcceptTenantList_Handler,
		},
		{
			MethodName: "GetServerIp",
			Handler:    _OAMServer_GetServerIp_Handler,
		},
		{
			MethodName: "AddPlatAccount",
			Handler:    _OAMServer_AddPlatAccount_Handler,
		},
		{
			MethodName: "AddPlatLicense",
			Handler:    _OAMServer_AddPlatLicense_Handler,
		},
		{
			MethodName: "BuyPlatLicenseCount",
			Handler:    _OAMServer_BuyPlatLicenseCount_Handler,
		},
		{
			MethodName: "SendActiveCode",
			Handler:    _OAMServer_SendActiveCode_Handler,
		},
		{
			MethodName: "ActivePlatInfo",
			Handler:    _OAMServer_ActivePlatInfo_Handler,
		},
		{
			MethodName: "GetPlatInfo",
			Handler:    _OAMServer_GetPlatInfo_Handler,
		},
		{
			MethodName: "GetPlatLicenseInfo",
			Handler:    _OAMServer_GetPlatLicenseInfo_Handler,
		},
		{
			MethodName: "UpdatePlatInfo",
			Handler:    _OAMServer_UpdatePlatInfo_Handler,
		},
		{
			MethodName: "UpdatePlatLicenseInfo",
			Handler:    _OAMServer_UpdatePlatLicenseInfo_Handler,
		},
		{
			MethodName: "GetPlatDomain",
			Handler:    _OAMServer_GetPlatDomain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "oam.proto",
}
