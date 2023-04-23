// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: domain/common.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 请求的策略类型
type RequestType int32

const (
	// 0 is reserved for errors
	RequestType_UNKNOWN        RequestType = 0
	RequestType_UNIT           RequestType = 1
	RequestType_ROUTE          RequestType = 2
	RequestType_CIRCUITBREAKER RequestType = 3
	RequestType_AUTH           RequestType = 4
	RequestType_LOADBALANCE    RequestType = 5
	RequestType_RATELIMIT      RequestType = 6
)

// Enum value maps for RequestType.
var (
	RequestType_name = map[int32]string{
		0: "UNKNOWN",
		1: "UNIT",
		2: "ROUTE",
		3: "CIRCUITBREAKER",
		4: "AUTH",
		5: "LOADBALANCE",
		6: "RATELIMIT",
	}
	RequestType_value = map[string]int32{
		"UNKNOWN":        0,
		"UNIT":           1,
		"ROUTE":          2,
		"CIRCUITBREAKER": 3,
		"AUTH":           4,
		"LOADBALANCE":    5,
		"RATELIMIT":      6,
	}
)

func (x RequestType) Enum() *RequestType {
	p := new(RequestType)
	*p = x
	return p
}

func (x RequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_domain_common_proto_enumTypes[0].Descriptor()
}

func (RequestType) Type() protoreflect.EnumType {
	return &file_domain_common_proto_enumTypes[0]
}

func (x RequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RequestType.Descriptor instead.
func (RequestType) EnumDescriptor() ([]byte, []int) {
	return file_domain_common_proto_rawDescGZIP(), []int{0}
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestType RequestType       `protobuf:"varint,1,opt,name=requestType,proto3,enum=v1.RequestType" json:"requestType,omitempty"`
	Key         string            `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Headers     map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_domain_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_domain_common_proto_rawDescGZIP(), []int{0}
}

func (x *Metadata) GetRequestType() RequestType {
	if x != nil {
		return x.RequestType
	}
	return RequestType_UNKNOWN
}

func (x *Metadata) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Metadata) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *Metadata  `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Body     *anypb.Any `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_domain_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_domain_common_proto_rawDescGZIP(), []int{1}
}

func (x *Payload) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Payload) GetBody() *anypb.Any {
	if x != nil {
		return x.Body
	}
	return nil
}

var File_domain_common_proto protoreflect.FileDescriptor

var file_domain_common_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x01, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x31, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5d, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x28, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x2a, 0x6d, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x4e, 0x49, 0x54, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x52, 0x4f, 0x55, 0x54, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x49, 0x52, 0x43, 0x55,
	0x49, 0x54, 0x42, 0x52, 0x45, 0x41, 0x4b, 0x45, 0x52, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x41,
	0x55, 0x54, 0x48, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x4c, 0x4f, 0x41, 0x44, 0x42, 0x41, 0x4c,
	0x41, 0x4e, 0x43, 0x45, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x52, 0x41, 0x54, 0x45, 0x4c, 0x49,
	0x4d, 0x49, 0x54, 0x10, 0x06, 0x42, 0x23, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x6a, 0x64, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6a, 0x6d, 0x73, 0x66, 0x2e, 0x70, 0x62, 0x2e, 0x64, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x50, 0x01, 0x5a, 0x03, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_domain_common_proto_rawDescOnce sync.Once
	file_domain_common_proto_rawDescData = file_domain_common_proto_rawDesc
)

func file_domain_common_proto_rawDescGZIP() []byte {
	file_domain_common_proto_rawDescOnce.Do(func() {
		file_domain_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_common_proto_rawDescData)
	})
	return file_domain_common_proto_rawDescData
}

var file_domain_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_domain_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_domain_common_proto_goTypes = []interface{}{
	(RequestType)(0),  // 0: v1.RequestType
	(*Metadata)(nil),  // 1: v1.Metadata
	(*Payload)(nil),   // 2: v1.Payload
	nil,               // 3: v1.Metadata.HeadersEntry
	(*anypb.Any)(nil), // 4: google.protobuf.Any
}
var file_domain_common_proto_depIdxs = []int32{
	0, // 0: v1.Metadata.requestType:type_name -> v1.RequestType
	3, // 1: v1.Metadata.headers:type_name -> v1.Metadata.HeadersEntry
	1, // 2: v1.Payload.metadata:type_name -> v1.Metadata
	4, // 3: v1.Payload.body:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_domain_common_proto_init() }
func file_domain_common_proto_init() {
	if File_domain_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_domain_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_domain_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
			RawDescriptor: file_domain_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_common_proto_goTypes,
		DependencyIndexes: file_domain_common_proto_depIdxs,
		EnumInfos:         file_domain_common_proto_enumTypes,
		MessageInfos:      file_domain_common_proto_msgTypes,
	}.Build()
	File_domain_common_proto = out.File
	file_domain_common_proto_rawDesc = nil
	file_domain_common_proto_goTypes = nil
	file_domain_common_proto_depIdxs = nil
}
