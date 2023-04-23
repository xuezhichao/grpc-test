// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: domain/transcoder.proto

package v1

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

type TranscoderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata    *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3,oneof" json:"metadata,omitempty"`
	ServiceName string    `protobuf:"bytes,2,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	TransType   string    `protobuf:"bytes,3,opt,name=transType,proto3" json:"transType,omitempty"`
}

func (x *TranscoderRequest) Reset() {
	*x = TranscoderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_transcoder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranscoderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranscoderRequest) ProtoMessage() {}

func (x *TranscoderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_domain_transcoder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranscoderRequest.ProtoReflect.Descriptor instead.
func (*TranscoderRequest) Descriptor() ([]byte, []int) {
	return file_domain_transcoder_proto_rawDescGZIP(), []int{0}
}

func (x *TranscoderRequest) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *TranscoderRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *TranscoderRequest) GetTransType() string {
	if x != nil {
		return x.TransType
	}
	return ""
}

type Matcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MatcherType string            `protobuf:"bytes,1,opt,name=matcherType,proto3" json:"matcherType,omitempty"`
	MeshId      string            `protobuf:"bytes,2,opt,name=meshId,proto3" json:"meshId,omitempty"`
	ServiceName string            `protobuf:"bytes,3,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	TransSide   string            `protobuf:"bytes,4,opt,name=transSide,proto3" json:"transSide,omitempty"`
	Config      map[string]string `protobuf:"bytes,5,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Matcher) Reset() {
	*x = Matcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_transcoder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Matcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Matcher) ProtoMessage() {}

func (x *Matcher) ProtoReflect() protoreflect.Message {
	mi := &file_domain_transcoder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Matcher.ProtoReflect.Descriptor instead.
func (*Matcher) Descriptor() ([]byte, []int) {
	return file_domain_transcoder_proto_rawDescGZIP(), []int{1}
}

func (x *Matcher) GetMatcherType() string {
	if x != nil {
		return x.MatcherType
	}
	return ""
}

func (x *Matcher) GetMeshId() string {
	if x != nil {
		return x.MeshId
	}
	return ""
}

func (x *Matcher) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Matcher) GetTransSide() string {
	if x != nil {
		return x.TransSide
	}
	return ""
}

func (x *Matcher) GetConfig() map[string]string {
	if x != nil {
		return x.Config
	}
	return nil
}

type RuleConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enable      bool   `protobuf:"varint,1,opt,name=enable,proto3" json:"enable,omitempty"`
	UniqueId    string `protobuf:"bytes,2,opt,name=uniqueId,json=unique_id,proto3" json:"uniqueId,omitempty"`
	TargetApp   string `protobuf:"bytes,3,opt,name=targetApp,json=target_app,proto3" json:"targetApp,omitempty"`
	ReqMapping  string `protobuf:"bytes,4,opt,name=reqMapping,json=req_mapping,proto3" json:"reqMapping,omitempty"`
	RespMapping string `protobuf:"bytes,5,opt,name=respMapping,json=resp_mapping,proto3" json:"respMapping,omitempty"`
	Url         string `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	Method      string `protobuf:"bytes,7,opt,name=method,proto3" json:"method,omitempty"`
	ClusterName string `protobuf:"bytes,8,opt,name=clusterName,json=cluster_name,proto3" json:"clusterName,omitempty"`
}

func (x *RuleConfig) Reset() {
	*x = RuleConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_transcoder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RuleConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RuleConfig) ProtoMessage() {}

func (x *RuleConfig) ProtoReflect() protoreflect.Message {
	mi := &file_domain_transcoder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RuleConfig.ProtoReflect.Descriptor instead.
func (*RuleConfig) Descriptor() ([]byte, []int) {
	return file_domain_transcoder_proto_rawDescGZIP(), []int{2}
}

func (x *RuleConfig) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *RuleConfig) GetUniqueId() string {
	if x != nil {
		return x.UniqueId
	}
	return ""
}

func (x *RuleConfig) GetTargetApp() string {
	if x != nil {
		return x.TargetApp
	}
	return ""
}

func (x *RuleConfig) GetReqMapping() string {
	if x != nil {
		return x.ReqMapping
	}
	return ""
}

func (x *RuleConfig) GetRespMapping() string {
	if x != nil {
		return x.RespMapping
	}
	return ""
}

func (x *RuleConfig) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *RuleConfig) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *RuleConfig) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

type TranscoderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata    *Metadata     `protobuf:"bytes,1,opt,name=metadata,proto3,oneof" json:"metadata,omitempty"`
	ServiceName string        `protobuf:"bytes,2,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	TransType   string        `protobuf:"bytes,3,opt,name=transType,proto3" json:"transType,omitempty"`
	Matcher     *Matcher      `protobuf:"bytes,4,opt,name=matcher,proto3,oneof" json:"matcher,omitempty"`
	Rules       []*RuleConfig `protobuf:"bytes,5,rep,name=rules,proto3" json:"rules,omitempty"`
	Enable      bool          `protobuf:"varint,6,opt,name=enable,proto3" json:"enable,omitempty"`
	Version     string        `protobuf:"bytes,7,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *TranscoderResponse) Reset() {
	*x = TranscoderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_domain_transcoder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranscoderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranscoderResponse) ProtoMessage() {}

func (x *TranscoderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_domain_transcoder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranscoderResponse.ProtoReflect.Descriptor instead.
func (*TranscoderResponse) Descriptor() ([]byte, []int) {
	return file_domain_transcoder_proto_rawDescGZIP(), []int{3}
}

func (x *TranscoderResponse) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *TranscoderResponse) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *TranscoderResponse) GetTransType() string {
	if x != nil {
		return x.TransType
	}
	return ""
}

func (x *TranscoderResponse) GetMatcher() *Matcher {
	if x != nil {
		return x.Matcher
	}
	return nil
}

func (x *TranscoderResponse) GetRules() []*RuleConfig {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *TranscoderResponse) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

func (x *TranscoderResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_domain_transcoder_proto protoreflect.FileDescriptor

var file_domain_transcoder_proto_rawDesc = []byte{
	0x0a, 0x17, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f,
	0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x13, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xef, 0x01, 0x0a, 0x07, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x73, 0x68, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x73, 0x68, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x53, 0x69, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x53, 0x69, 0x64, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x39, 0x0a, 0x0b, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf1, 0x01, 0x0a, 0x0a, 0x52, 0x75, 0x6c, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1b, 0x0a,
	0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x09, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x41, 0x70, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x70, 0x70, 0x12, 0x1f, 0x0a, 0x0a, 0x72, 0x65, 0x71,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72,
	0x65, 0x71, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x70, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x72, 0x65, 0x73, 0x70, 0x5f, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12,
	0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xa0, 0x02, 0x0a, 0x12, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01,
	0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x48, 0x01,
	0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x05,
	0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x75, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x23, 0x0a,
	0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x6a, 0x64, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6a, 0x6d, 0x73,
	0x66, 0x2e, 0x70, 0x62, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x50, 0x01, 0x5a, 0x03, 0x2f,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_domain_transcoder_proto_rawDescOnce sync.Once
	file_domain_transcoder_proto_rawDescData = file_domain_transcoder_proto_rawDesc
)

func file_domain_transcoder_proto_rawDescGZIP() []byte {
	file_domain_transcoder_proto_rawDescOnce.Do(func() {
		file_domain_transcoder_proto_rawDescData = protoimpl.X.CompressGZIP(file_domain_transcoder_proto_rawDescData)
	})
	return file_domain_transcoder_proto_rawDescData
}

var file_domain_transcoder_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_domain_transcoder_proto_goTypes = []interface{}{
	(*TranscoderRequest)(nil),  // 0: v1.TranscoderRequest
	(*Matcher)(nil),            // 1: v1.Matcher
	(*RuleConfig)(nil),         // 2: v1.RuleConfig
	(*TranscoderResponse)(nil), // 3: v1.TranscoderResponse
	nil,                        // 4: v1.Matcher.ConfigEntry
	(*Metadata)(nil),           // 5: v1.Metadata
}
var file_domain_transcoder_proto_depIdxs = []int32{
	5, // 0: v1.TranscoderRequest.metadata:type_name -> v1.Metadata
	4, // 1: v1.Matcher.config:type_name -> v1.Matcher.ConfigEntry
	5, // 2: v1.TranscoderResponse.metadata:type_name -> v1.Metadata
	1, // 3: v1.TranscoderResponse.matcher:type_name -> v1.Matcher
	2, // 4: v1.TranscoderResponse.rules:type_name -> v1.RuleConfig
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_domain_transcoder_proto_init() }
func file_domain_transcoder_proto_init() {
	if File_domain_transcoder_proto != nil {
		return
	}
	file_domain_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_domain_transcoder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranscoderRequest); i {
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
		file_domain_transcoder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Matcher); i {
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
		file_domain_transcoder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RuleConfig); i {
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
		file_domain_transcoder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranscoderResponse); i {
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
	file_domain_transcoder_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_domain_transcoder_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_domain_transcoder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_domain_transcoder_proto_goTypes,
		DependencyIndexes: file_domain_transcoder_proto_depIdxs,
		MessageInfos:      file_domain_transcoder_proto_msgTypes,
	}.Build()
	File_domain_transcoder_proto = out.File
	file_domain_transcoder_proto_rawDesc = nil
	file_domain_transcoder_proto_goTypes = nil
	file_domain_transcoder_proto_depIdxs = nil
}
