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
	0x6f, 0x74, 0x6f, 0x32, 0xcf, 0x04, 0x0a, 0x09, 0x4f, 0x41, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x23, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x05, 0x2e, 0x50, 0x6f, 0x6e, 0x67,
	0x1a, 0x05, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12,
	0x05, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x59, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69,
	0x4c, 0x6f, 0x67, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x69, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x4c, 0x6f, 0x67, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x11, 0x2f, 0x6f, 0x61,
	0x6d, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x6c, 0x6f, 0x67, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x3a, 0x01,
	0x2a, 0x12, 0x56, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x55, 0x70, 0x46, 0x6c, 0x61,
	0x67, 0x73, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x46, 0x6c, 0x61, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69,
	0x55, 0x70, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x6f, 0x61, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x2f, 0x75,
	0x70, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x64, 0x0a, 0x10, 0x41, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e,
	0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x22, 0x13, 0x2f, 0x6f, 0x61, 0x6d, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12,
	0x4f, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x10, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x6f, 0x61,
	0x6d, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x3a, 0x01, 0x2a,
	0x12, 0x47, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0a, 0x2e, 0x4f, 0x41, 0x4d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x6f, 0x61, 0x6d, 0x2f, 0x67, 0x65, 0x74, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x6a, 0x0a, 0x16, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x46, 0x6c,
	0x61, 0x67, 0x73, 0x12, 0x1e, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22,
	0x11, 0x2f, 0x6f, 0x61, 0x6d, 0x2f, 0x70, 0x75, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x66, 0x6c, 0x61,
	0x67, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_oam_proto_goTypes = []interface{}{
	(*Pong)(nil),                          // 0: Pong
	(*GetCliFlagsRequest)(nil),            // 1: GetCliFlagsRequest
	(*AcceptClientListRequest)(nil),       // 2: AcceptClientListRequest
	(*ListClientRequest)(nil),             // 3: ListClientRequest
	(*BatchUpdateClientFlagsRequest)(nil), // 4: BatchUpdateClientFlagsRequest
	(*GetCliLogFlagsReply)(nil),           // 5: GetCliLogFlagsReply
	(*GetCliUpFlagsReply)(nil),            // 6: GetCliUpFlagsReply
	(*AcceptClientListReply)(nil),         // 7: AcceptClientListReply
	(*ListClientReply)(nil),               // 8: ListClientReply
	(*OAMClient)(nil),                     // 9: OAMClient
	(*UpdateClientReply)(nil),             // 10: UpdateClientReply
}
var file_oam_proto_depIdxs = []int32{
	0,  // 0: OAMServer.Ping:input_type -> Pong
	1,  // 1: OAMServer.GetCliLogFlags:input_type -> GetCliFlagsRequest
	1,  // 2: OAMServer.GetCliUpFlags:input_type -> GetCliFlagsRequest
	2,  // 3: OAMServer.AcceptClientList:input_type -> AcceptClientListRequest
	3,  // 4: OAMServer.ListClient:input_type -> ListClientRequest
	3,  // 5: OAMServer.GetClient:input_type -> ListClientRequest
	4,  // 6: OAMServer.BatchUpdateClientFlags:input_type -> BatchUpdateClientFlagsRequest
	0,  // 7: OAMServer.Ping:output_type -> Pong
	5,  // 8: OAMServer.GetCliLogFlags:output_type -> GetCliLogFlagsReply
	6,  // 9: OAMServer.GetCliUpFlags:output_type -> GetCliUpFlagsReply
	7,  // 10: OAMServer.AcceptClientList:output_type -> AcceptClientListReply
	8,  // 11: OAMServer.ListClient:output_type -> ListClientReply
	9,  // 12: OAMServer.GetClient:output_type -> OAMClient
	10, // 13: OAMServer.BatchUpdateClientFlags:output_type -> UpdateClientReply
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
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
