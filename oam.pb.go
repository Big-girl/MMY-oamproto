// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: oam.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_oam_proto protoreflect.FileDescriptor

var file_oam_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6f, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x70, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x63, 0x6c, 0x69, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x6c, 0x69, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0xb8, 0x0e, 0x0a, 0x09, 0x4f, 0x41, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x26, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x05, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x1a,
	0x05, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x22, 0x05,
	0x2f, 0x70, 0x69, 0x6e, 0x67, 0x3a, 0x01, 0x2a, 0x12, 0x59, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x69, 0x4c, 0x6f, 0x67, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x6c, 0x69, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x4c, 0x6f, 0x67, 0x46, 0x6c, 0x61, 0x67, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f,
	0x6f, 0x61, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x61, 0x67, 0x73,
	0x3a, 0x01, 0x2a, 0x12, 0x56, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x55, 0x70, 0x46,
	0x6c, 0x61, 0x67, 0x73, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x46, 0x6c, 0x61,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6c, 0x69, 0x55, 0x70, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x6f, 0x61, 0x6d, 0x2f, 0x67, 0x65, 0x74,
	0x2f, 0x75, 0x70, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x64, 0x0a, 0x10, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x18, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x6f, 0x61, 0x6d, 0x2f,
	0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x01,
	0x2a, 0x12, 0x4f, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f,
	0x6f, 0x61, 0x6d, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x3a,
	0x01, 0x2a, 0x12, 0x47, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12,
	0x12, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x4f, 0x41, 0x4d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x6f, 0x61, 0x6d, 0x2f, 0x67, 0x65,
	0x74, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x6a, 0x0a, 0x16, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x1e, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x22, 0x11, 0x2f, 0x6f, 0x61, 0x6d, 0x2f, 0x70, 0x75, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x66,
	0x6c, 0x61, 0x67, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x58, 0x0a, 0x0d, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x15, 0x2e, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x6f, 0x61, 0x6d, 0x2f,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x01,
	0x2a, 0x12, 0x49, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x11,
	0x2e, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0d, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x6f, 0x61, 0x6d, 0x2f, 0x61,
	0x64, 0x64, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x4c, 0x0a, 0x09,
	0x50, 0x75, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x11, 0x2e, 0x50, 0x75, 0x74, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x6f, 0x61, 0x6d, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x47, 0x0a, 0x09, 0x44, 0x65,
	0x6c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0c, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0d, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x6f,
	0x61, 0x6d, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x3a, 0x01, 0x2a, 0x12, 0x49, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x0c, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a,
	0x10, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x6f, 0x61, 0x6d, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x46,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x0c, 0x2e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x22, 0x0f, 0x2f, 0x6f, 0x61, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x3a, 0x01, 0x2a, 0x12, 0x53, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x2e, 0x4d, 0x6f, 0x64,
	0x65, 0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x0d, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b,
	0x22, 0x16, 0x2f, 0x6f, 0x61, 0x6d, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x64, 0x0a, 0x10, 0x41,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x18, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x41, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x6f, 0x61, 0x6d, 0x2f,
	0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x73, 0x3a, 0x01,
	0x2a, 0x12, 0x48, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x70,
	0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x70, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x0a, 0x2f, 0x6f,
	0x61, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x69, 0x70, 0x3a, 0x01, 0x2a, 0x12, 0x47, 0x0a, 0x0e, 0x41,
	0x64, 0x64, 0x50, 0x6c, 0x61, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0f, 0x2e,
	0x41, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a,
	0x2e, 0x50, 0x6c, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x12, 0x22, 0x0d, 0x2f, 0x6f, 0x61, 0x6d, 0x2f, 0x61, 0x64, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x74, 0x3a, 0x01, 0x2a, 0x12, 0x55, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x74, 0x4c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x2e, 0x41, 0x64, 0x64, 0x50, 0x6c, 0x61, 0x74,
	0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a,
	0x2e, 0x50, 0x6c, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x19, 0x22, 0x14, 0x2f, 0x6f, 0x61, 0x6d, 0x2f, 0x61, 0x64, 0x64, 0x2f, 0x70, 0x6c, 0x61,
	0x74, 0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x4e, 0x0a, 0x0e, 0x53,
	0x65, 0x6e, 0x64, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x50, 0x6c, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x22, 0x0d, 0x2f, 0x6f, 0x61, 0x6d, 0x2f,
	0x73, 0x65, 0x6e, 0x64, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x5a, 0x0a, 0x0e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x50, 0x6c,
	0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x6f, 0x61, 0x6d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x53, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6c,
	0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x61, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x6f, 0x61, 0x6d, 0x2f, 0x67, 0x65, 0x74,
	0x2f, 0x70, 0x6c, 0x61, 0x74, 0x69, 0x6e, 0x66, 0x6f, 0x3a, 0x01, 0x2a, 0x12, 0x6f, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x4c, 0x69, 0x63, 0x65,
	0x6e, 0x73, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x4c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d,
	0x22, 0x18, 0x2f, 0x6f, 0x61, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x74, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x69, 0x6e, 0x66, 0x6f, 0x3a, 0x01, 0x2a, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_oam_proto_goTypes = []interface{}{
	(*Pong)(nil),                          // 0: Pong
	(*GetCliFlagsRequest)(nil),            // 1: GetCliFlagsRequest
	(*AcceptClientListRequest)(nil),       // 2: AcceptClientListRequest
	(*ListClientRequest)(nil),             // 3: ListClientRequest
	(*BatchUpdateClientFlagsRequest)(nil), // 4: BatchUpdateClientFlagsRequest
	(*ReleasePolicyRequest)(nil),          // 5: ReleasePolicyRequest
	(*AddPolicyRequest)(nil),              // 6: AddPolicyRequest
	(*PutPolicyRequest)(nil),              // 7: PutPolicyRequest
	(*PolicyQuery)(nil),                   // 8: PolicyQuery
	(*ModelVersion)(nil),                  // 9: ModelVersion
	(*AcceptTenantListRequest)(nil),       // 10: AcceptTenantListRequest
	(*GetPlatRequest)(nil),                // 11: GetPlatRequest
	(*AddPlatRequest)(nil),                // 12: AddPlatRequest
	(*AddPlatLicenseRequest)(nil),         // 13: AddPlatLicenseRequest
	(*ActivePlatInfoRequest)(nil),         // 14: ActivePlatInfoRequest
	(*GetPlatInfoRequest)(nil),            // 15: GetPlatInfoRequest
	(*GetPlatLicenseInfoRequest)(nil),     // 16: GetPlatLicenseInfoRequest
	(*GetCliLogFlagsReply)(nil),           // 17: GetCliLogFlagsReply
	(*GetCliUpFlagsReply)(nil),            // 18: GetCliUpFlagsReply
	(*AcceptClientListReply)(nil),         // 19: AcceptClientListReply
	(*ListClientReply)(nil),               // 20: ListClientReply
	(*OAMClient)(nil),                     // 21: OAMClient
	(*UpdateClientReply)(nil),             // 22: UpdateClientReply
	(*ListPolicyReply)(nil),               // 23: ListPolicyReply
	(*OperateReply)(nil),                  // 24: OperateReply
	(*GetPolicyReply)(nil),                // 25: GetPolicyReply
	(*AcceptTenantListReply)(nil),         // 26: AcceptTenantListReply
	(*GetServerIpReply)(nil),              // 27: GetServerIpReply
	(*PlatReply)(nil),                     // 28: PlatReply
	(*ActivePlatInfoReply)(nil),           // 29: ActivePlatInfoReply
	(*GetPlatInfoReply)(nil),              // 30: GetPlatInfoReply
	(*GetPlatLicenseInfoReply)(nil),       // 31: GetPlatLicenseInfoReply
}
var file_oam_proto_depIdxs = []int32{
	0,  // 0: OAMServer.Ping:input_type -> Pong
	1,  // 1: OAMServer.GetCliLogFlags:input_type -> GetCliFlagsRequest
	1,  // 2: OAMServer.GetCliUpFlags:input_type -> GetCliFlagsRequest
	2,  // 3: OAMServer.AcceptClientList:input_type -> AcceptClientListRequest
	3,  // 4: OAMServer.ListClient:input_type -> ListClientRequest
	3,  // 5: OAMServer.GetClient:input_type -> ListClientRequest
	4,  // 6: OAMServer.BatchUpdateClientFlags:input_type -> BatchUpdateClientFlagsRequest
	5,  // 7: OAMServer.ReleasePolicy:input_type -> ReleasePolicyRequest
	6,  // 8: OAMServer.AddPolicy:input_type -> AddPolicyRequest
	7,  // 9: OAMServer.PutPolicy:input_type -> PutPolicyRequest
	8,  // 10: OAMServer.DelPolicy:input_type -> PolicyQuery
	8,  // 11: OAMServer.ListPolicy:input_type -> PolicyQuery
	8,  // 12: OAMServer.GetPolicy:input_type -> PolicyQuery
	9,  // 13: OAMServer.ListModelVersion:input_type -> ModelVersion
	10, // 14: OAMServer.AcceptTenantList:input_type -> AcceptTenantListRequest
	11, // 15: OAMServer.GetServerIp:input_type -> GetPlatRequest
	12, // 16: OAMServer.AddPlatAccount:input_type -> AddPlatRequest
	13, // 17: OAMServer.AddPlatLicense:input_type -> AddPlatLicenseRequest
	14, // 18: OAMServer.SendActiveCode:input_type -> ActivePlatInfoRequest
	14, // 19: OAMServer.ActivePlatInfo:input_type -> ActivePlatInfoRequest
	15, // 20: OAMServer.GetPlatInfo:input_type -> GetPlatInfoRequest
	16, // 21: OAMServer.GetPlatLicenseInfo:input_type -> GetPlatLicenseInfoRequest
	0,  // 22: OAMServer.Ping:output_type -> Pong
	17, // 23: OAMServer.GetCliLogFlags:output_type -> GetCliLogFlagsReply
	18, // 24: OAMServer.GetCliUpFlags:output_type -> GetCliUpFlagsReply
	19, // 25: OAMServer.AcceptClientList:output_type -> AcceptClientListReply
	20, // 26: OAMServer.ListClient:output_type -> ListClientReply
	21, // 27: OAMServer.GetClient:output_type -> OAMClient
	22, // 28: OAMServer.BatchUpdateClientFlags:output_type -> UpdateClientReply
	23, // 29: OAMServer.ReleasePolicy:output_type -> ListPolicyReply
	24, // 30: OAMServer.AddPolicy:output_type -> OperateReply
	24, // 31: OAMServer.PutPolicy:output_type -> OperateReply
	24, // 32: OAMServer.DelPolicy:output_type -> OperateReply
	23, // 33: OAMServer.ListPolicy:output_type -> ListPolicyReply
	25, // 34: OAMServer.GetPolicy:output_type -> GetPolicyReply
	9,  // 35: OAMServer.ListModelVersion:output_type -> ModelVersion
	26, // 36: OAMServer.AcceptTenantList:output_type -> AcceptTenantListReply
	27, // 37: OAMServer.GetServerIp:output_type -> GetServerIpReply
	28, // 38: OAMServer.AddPlatAccount:output_type -> PlatReply
	28, // 39: OAMServer.AddPlatLicense:output_type -> PlatReply
	28, // 40: OAMServer.SendActiveCode:output_type -> PlatReply
	29, // 41: OAMServer.ActivePlatInfo:output_type -> ActivePlatInfoReply
	30, // 42: OAMServer.GetPlatInfo:output_type -> GetPlatInfoReply
	31, // 43: OAMServer.GetPlatLicenseInfo:output_type -> GetPlatLicenseInfoReply
	22, // [22:44] is the sub-list for method output_type
	0,  // [0:22] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_oam_proto_init() }
func file_oam_proto_init() {
	if File_oam_proto != nil {
		return
	}
	file_ping_proto_init()
	file_cliflags_proto_init()
	file_client_proto_init()
	file_policylib_proto_init()
	file_modelversion_proto_init()
	file_tenant_proto_init()
	file_plat_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_oam_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oam_proto_goTypes,
		DependencyIndexes: file_oam_proto_depIdxs,
	}.Build()
	File_oam_proto = out.File
	file_oam_proto_rawDesc = nil
	file_oam_proto_goTypes = nil
	file_oam_proto_depIdxs = nil
}
