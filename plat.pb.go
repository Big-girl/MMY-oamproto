// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: plat.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 获取平台域名的请求参数
type GetPlatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 每个接口都需要的校验参数
	Co *CheckObj `protobuf:"bytes,1,opt,name=co,proto3" json:"co,omitempty"`
}

func (x *GetPlatRequest) Reset() {
	*x = GetPlatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlatRequest) ProtoMessage() {}

func (x *GetPlatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlatRequest.ProtoReflect.Descriptor instead.
func (*GetPlatRequest) Descriptor() ([]byte, []int) {
	return file_plat_proto_rawDescGZIP(), []int{0}
}

func (x *GetPlatRequest) GetCo() *CheckObj {
	if x != nil {
		return x.Co
	}
	return nil
}

// 返回的数据结构响应参数
type IpAddr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 数据端口
	Dport string `protobuf:"bytes,1,opt,name=dport,proto3" json:"dport,omitempty"`
	// 文件端口
	Fport string `protobuf:"bytes,2,opt,name=fport,proto3" json:"fport,omitempty"`
	// ip或者域名
	Ip string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	// 是否本地化部署
	Islocal int32 `protobuf:"varint,4,opt,name=islocal,proto3" json:"islocal,omitempty"`
	// 是否OSS保存文件
	Isoss int32 `protobuf:"varint,5,opt,name=isoss,proto3" json:"isoss,omitempty"`
}

func (x *IpAddr) Reset() {
	*x = IpAddr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpAddr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpAddr) ProtoMessage() {}

func (x *IpAddr) ProtoReflect() protoreflect.Message {
	mi := &file_plat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpAddr.ProtoReflect.Descriptor instead.
func (*IpAddr) Descriptor() ([]byte, []int) {
	return file_plat_proto_rawDescGZIP(), []int{1}
}

func (x *IpAddr) GetDport() string {
	if x != nil {
		return x.Dport
	}
	return ""
}

func (x *IpAddr) GetFport() string {
	if x != nil {
		return x.Fport
	}
	return ""
}

func (x *IpAddr) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *IpAddr) GetIslocal() int32 {
	if x != nil {
		return x.Islocal
	}
	return 0
}

func (x *IpAddr) GetIsoss() int32 {
	if x != nil {
		return x.Isoss
	}
	return 0
}

// 获取平台服务器的ip或者域名的响应参数
type GetServerIpReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 错误码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// ip域名参数
	Ipaddr *IpAddr `protobuf:"bytes,2,opt,name=ipaddr,proto3" json:"ipaddr,omitempty"`
}

func (x *GetServerIpReply) Reset() {
	*x = GetServerIpReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServerIpReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServerIpReply) ProtoMessage() {}

func (x *GetServerIpReply) ProtoReflect() protoreflect.Message {
	mi := &file_plat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServerIpReply.ProtoReflect.Descriptor instead.
func (*GetServerIpReply) Descriptor() ([]byte, []int) {
	return file_plat_proto_rawDescGZIP(), []int{2}
}

func (x *GetServerIpReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetServerIpReply) GetIpaddr() *IpAddr {
	if x != nil {
		return x.Ipaddr
	}
	return nil
}

// 创建平台账号的请求参数
type AddPlatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 每个接口都需要的校验参数
	Co *CheckObj `protobuf:"bytes,1,opt,name=co,proto3" json:"co,omitempty"`
	// 平台名称
	Platname string `protobuf:"bytes,2,opt,name=platname,proto3" json:"platname,omitempty"`
	// 联系人
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// 手机号
	Phone string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	// 平台域名
	Ip string `protobuf:"bytes,5,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *AddPlatRequest) Reset() {
	*x = AddPlatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPlatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPlatRequest) ProtoMessage() {}

func (x *AddPlatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPlatRequest.ProtoReflect.Descriptor instead.
func (*AddPlatRequest) Descriptor() ([]byte, []int) {
	return file_plat_proto_rawDescGZIP(), []int{3}
}

func (x *AddPlatRequest) GetCo() *CheckObj {
	if x != nil {
		return x.Co
	}
	return nil
}

func (x *AddPlatRequest) GetPlatname() string {
	if x != nil {
		return x.Platname
	}
	return ""
}

func (x *AddPlatRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddPlatRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *AddPlatRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

// 平台账号的响应
type PlatReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 错误码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 提示信息
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// 影响行数
	Affectedrows int64 `protobuf:"varint,3,opt,name=affectedrows,proto3" json:"affectedrows,omitempty"`
}

func (x *PlatReply) Reset() {
	*x = PlatReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlatReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlatReply) ProtoMessage() {}

func (x *PlatReply) ProtoReflect() protoreflect.Message {
	mi := &file_plat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlatReply.ProtoReflect.Descriptor instead.
func (*PlatReply) Descriptor() ([]byte, []int) {
	return file_plat_proto_rawDescGZIP(), []int{4}
}

func (x *PlatReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *PlatReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *PlatReply) GetAffectedrows() int64 {
	if x != nil {
		return x.Affectedrows
	}
	return 0
}

// 创建平台授权的请求参数
type AddPlatLicenseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 每个接口都需要的校验参数
	Co *CheckObj `protobuf:"bytes,1,opt,name=co,proto3" json:"co,omitempty"`
	// 平台id
	Platid int64 `protobuf:"varint,2,opt,name=platid,proto3" json:"platid,omitempty"`
	// 授权类型
	Type int32 `protobuf:"varint,3,opt,name=type,proto3" json:"type,omitempty"`
	// 授权数量
	Count int64 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *AddPlatLicenseRequest) Reset() {
	*x = AddPlatLicenseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddPlatLicenseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddPlatLicenseRequest) ProtoMessage() {}

func (x *AddPlatLicenseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddPlatLicenseRequest.ProtoReflect.Descriptor instead.
func (*AddPlatLicenseRequest) Descriptor() ([]byte, []int) {
	return file_plat_proto_rawDescGZIP(), []int{5}
}

func (x *AddPlatLicenseRequest) GetCo() *CheckObj {
	if x != nil {
		return x.Co
	}
	return nil
}

func (x *AddPlatLicenseRequest) GetPlatid() int64 {
	if x != nil {
		return x.Platid
	}
	return 0
}

func (x *AddPlatLicenseRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *AddPlatLicenseRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// 发送激活码和激活平台信息请求
type ActivePlatInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 每个接口都需要的校验参数
	Co *CheckObj `protobuf:"bytes,1,opt,name=co,proto3" json:"co,omitempty"`
	// 传输的激活数据
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ActivePlatInfoRequest) Reset() {
	*x = ActivePlatInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivePlatInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivePlatInfoRequest) ProtoMessage() {}

func (x *ActivePlatInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_plat_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivePlatInfoRequest.ProtoReflect.Descriptor instead.
func (*ActivePlatInfoRequest) Descriptor() ([]byte, []int) {
	return file_plat_proto_rawDescGZIP(), []int{6}
}

func (x *ActivePlatInfoRequest) GetCo() *CheckObj {
	if x != nil {
		return x.Co
	}
	return nil
}

func (x *ActivePlatInfoRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// 激活平台返回的数据
type ActivePlatInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 错误码
	Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 返回的激活数据
	Data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ActivePlatInfoReply) Reset() {
	*x = ActivePlatInfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plat_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivePlatInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivePlatInfoReply) ProtoMessage() {}

func (x *ActivePlatInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_plat_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivePlatInfoReply.ProtoReflect.Descriptor instead.
func (*ActivePlatInfoReply) Descriptor() ([]byte, []int) {
	return file_plat_proto_rawDescGZIP(), []int{7}
}

func (x *ActivePlatInfoReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ActivePlatInfoReply) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_plat_proto protoreflect.FileDescriptor

var file_plat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x02,
	0x63, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x4f, 0x62, 0x6a, 0x52, 0x02, 0x63, 0x6f, 0x22, 0x74, 0x0a, 0x06, 0x49, 0x70, 0x41, 0x64, 0x64,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x64, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x73, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x69, 0x73, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x6f, 0x73, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x73, 0x6f, 0x73, 0x73, 0x22, 0x47, 0x0a,
	0x10, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x70, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x69, 0x70, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x52, 0x06,
	0x69, 0x70, 0x61, 0x64, 0x64, 0x72, 0x22, 0x81, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x50, 0x6c,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x02, 0x63, 0x6f, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x62, 0x6a,
	0x52, 0x02, 0x63, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x5d, 0x0a, 0x09, 0x50, 0x6c,
	0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x66, 0x66, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x61, 0x66, 0x66,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x72, 0x6f, 0x77, 0x73, 0x22, 0x74, 0x0a, 0x15, 0x41, 0x64, 0x64,
	0x50, 0x6c, 0x61, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x02, 0x63, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x62, 0x6a, 0x52, 0x02, 0x63, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x70, 0x6c, 0x61, 0x74, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70,
	0x6c, 0x61, 0x74, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x46, 0x0a, 0x15, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x02, 0x63, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4f, 0x62, 0x6a, 0x52,
	0x02, 0x63, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3d, 0x0a, 0x13, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x50, 0x6c, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plat_proto_rawDescOnce sync.Once
	file_plat_proto_rawDescData = file_plat_proto_rawDesc
)

func file_plat_proto_rawDescGZIP() []byte {
	file_plat_proto_rawDescOnce.Do(func() {
		file_plat_proto_rawDescData = protoimpl.X.CompressGZIP(file_plat_proto_rawDescData)
	})
	return file_plat_proto_rawDescData
}

var file_plat_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_plat_proto_goTypes = []interface{}{
	(*GetPlatRequest)(nil),        // 0: GetPlatRequest
	(*IpAddr)(nil),                // 1: IpAddr
	(*GetServerIpReply)(nil),      // 2: GetServerIpReply
	(*AddPlatRequest)(nil),        // 3: AddPlatRequest
	(*PlatReply)(nil),             // 4: PlatReply
	(*AddPlatLicenseRequest)(nil), // 5: AddPlatLicenseRequest
	(*ActivePlatInfoRequest)(nil), // 6: ActivePlatInfoRequest
	(*ActivePlatInfoReply)(nil),   // 7: ActivePlatInfoReply
	(*CheckObj)(nil),              // 8: CheckObj
}
var file_plat_proto_depIdxs = []int32{
	8, // 0: GetPlatRequest.co:type_name -> CheckObj
	1, // 1: GetServerIpReply.ipaddr:type_name -> IpAddr
	8, // 2: AddPlatRequest.co:type_name -> CheckObj
	8, // 3: AddPlatLicenseRequest.co:type_name -> CheckObj
	8, // 4: ActivePlatInfoRequest.co:type_name -> CheckObj
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_plat_proto_init() }
func file_plat_proto_init() {
	if File_plat_proto != nil {
		return
	}
	file_client_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_plat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlatRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_plat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IpAddr); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_plat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServerIpReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_plat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPlatRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_plat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlatReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_plat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddPlatLicenseRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_plat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivePlatInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_plat_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivePlatInfoReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_plat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plat_proto_goTypes,
		DependencyIndexes: file_plat_proto_depIdxs,
		MessageInfos:      file_plat_proto_msgTypes,
	}.Build()
	File_plat_proto = out.File
	file_plat_proto_rawDesc = nil
	file_plat_proto_goTypes = nil
	file_plat_proto_depIdxs = nil
}
