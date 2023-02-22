// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: spire/plugin/agent/workloadattestor/v1/workloadattestor.proto

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

type AttestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The process ID of the workload to attest.
	Pid int32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	// Optional. Variable of type map[string]string
	Meta map[string]string `protobuf:"bytes,2,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AttestRequest) Reset() {
	*x = AttestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestRequest) ProtoMessage() {}

func (x *AttestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestRequest.ProtoReflect.Descriptor instead.
func (*AttestRequest) Descriptor() ([]byte, []int) {
	return file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_rawDescGZIP(), []int{0}
}

func (x *AttestRequest) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *AttestRequest) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

type AttestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Selector values related to the attested workload. The type
	// of the selector is inferred from the plugin name.
	SelectorValues []string `protobuf:"bytes,1,rep,name=selector_values,json=selectorValues,proto3" json:"selector_values,omitempty"`
}

func (x *AttestResponse) Reset() {
	*x = AttestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttestResponse) ProtoMessage() {}

func (x *AttestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttestResponse.ProtoReflect.Descriptor instead.
func (*AttestResponse) Descriptor() ([]byte, []int) {
	return file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_rawDescGZIP(), []int{1}
}

func (x *AttestResponse) GetSelectorValues() []string {
	if x != nil {
		return x.SelectorValues
	}
	return nil
}

var File_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto protoreflect.FileDescriptor

var file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x26, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x22, 0xaf, 0x01, 0x0a, 0x0d, 0x41, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x53, 0x0a, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x73, 0x70, 0x69, 0x72,
	0x65, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x77,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x39, 0x0a, 0x0e, 0x41, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x32, 0x8b, 0x01, 0x0a, 0x10, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x12, 0x77, 0x0a, 0x06, 0x41, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x12, 0x35, 0x2e, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61,
	0x64, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x73, 0x70, 0x69,
	0x72, 0x65, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x74, 0x74, 0x65, 0x73, 0x74, 0x6f, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x53, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x63, 0x63, 0x75, 0x6b, 0x6e, 0x6f, 0x78, 0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2d,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2d, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x70, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x61, 0x74, 0x74, 0x65,
	0x73, 0x74, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_rawDescOnce sync.Once
	file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_rawDescData = file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_rawDesc
)

func file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_rawDescGZIP() []byte {
	file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_rawDescOnce.Do(func() {
		file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_rawDescData = protoimpl.X.CompressGZIP(file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_rawDescData)
	})
	return file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_rawDescData
}

var file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_goTypes = []interface{}{
	(*AttestRequest)(nil),  // 0: spire.plugin.agent.workloadattestor.v1.AttestRequest
	(*AttestResponse)(nil), // 1: spire.plugin.agent.workloadattestor.v1.AttestResponse
	nil,                    // 2: spire.plugin.agent.workloadattestor.v1.AttestRequest.MetaEntry
}
var file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_depIdxs = []int32{
	2, // 0: spire.plugin.agent.workloadattestor.v1.AttestRequest.meta:type_name -> spire.plugin.agent.workloadattestor.v1.AttestRequest.MetaEntry
	0, // 1: spire.plugin.agent.workloadattestor.v1.WorkloadAttestor.Attest:input_type -> spire.plugin.agent.workloadattestor.v1.AttestRequest
	1, // 2: spire.plugin.agent.workloadattestor.v1.WorkloadAttestor.Attest:output_type -> spire.plugin.agent.workloadattestor.v1.AttestResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_init() }
func file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_init() {
	if File_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestRequest); i {
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
		file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttestResponse); i {
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
			RawDescriptor: file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_goTypes,
		DependencyIndexes: file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_depIdxs,
		MessageInfos:      file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_msgTypes,
	}.Build()
	File_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto = out.File
	file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_rawDesc = nil
	file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_goTypes = nil
	file_spire_plugin_agent_workloadattestor_v1_workloadattestor_proto_depIdxs = nil
}
