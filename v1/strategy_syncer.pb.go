// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: strategy_syncer.proto

package v1

import (
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

var File_strategy_syncer_proto protoreflect.FileDescriptor

var file_strategy_syncer_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x5f, 0x73, 0x79, 0x6e, 0x63, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x13, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x6d, 0x0a, 0x0e, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x53, 0x79, 0x6e, 0x63, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x0b, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x28, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e,
	0x6a, 0x64, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6a, 0x6d, 0x73, 0x66, 0x2e, 0x70, 0x62, 0x2e,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x50, 0x01, 0x5a, 0x03, 0x2f, 0x76, 0x31, 0x88,
	0x01, 0x01, 0x50, 0x00, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_strategy_syncer_proto_goTypes = []interface{}{
	(*Payload)(nil), // 0: v1.Payload
}
var file_strategy_syncer_proto_depIdxs = []int32{
	0, // 0: v1.StrategySyncer.GetStrategy:input_type -> v1.Payload
	0, // 1: v1.StrategySyncer.StreamStrategy:input_type -> v1.Payload
	0, // 2: v1.StrategySyncer.GetStrategy:output_type -> v1.Payload
	0, // 3: v1.StrategySyncer.StreamStrategy:output_type -> v1.Payload
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_strategy_syncer_proto_init() }
func file_strategy_syncer_proto_init() {
	if File_strategy_syncer_proto != nil {
		return
	}
	file_domain_common_proto_init()
	file_domain_transcoder_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_strategy_syncer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_strategy_syncer_proto_goTypes,
		DependencyIndexes: file_strategy_syncer_proto_depIdxs,
	}.Build()
	File_strategy_syncer_proto = out.File
	file_strategy_syncer_proto_rawDesc = nil
	file_strategy_syncer_proto_goTypes = nil
	file_strategy_syncer_proto_depIdxs = nil
}
