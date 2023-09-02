// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: supervisor/v1alpha1/ssh.proto

package supervisorv1alpha1

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

type FetchAuthorizedKeysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FetchAuthorizedKeysRequest) Reset() {
	*x = FetchAuthorizedKeysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supervisor_v1alpha1_ssh_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchAuthorizedKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAuthorizedKeysRequest) ProtoMessage() {}

func (x *FetchAuthorizedKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supervisor_v1alpha1_ssh_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAuthorizedKeysRequest.ProtoReflect.Descriptor instead.
func (*FetchAuthorizedKeysRequest) Descriptor() ([]byte, []int) {
	return file_supervisor_v1alpha1_ssh_proto_rawDescGZIP(), []int{0}
}

type FetchAuthorizedKeysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorizedKeys string `protobuf:"bytes,1,opt,name=authorized_keys,json=authorizedKeys,proto3" json:"authorized_keys,omitempty"`
}

func (x *FetchAuthorizedKeysResponse) Reset() {
	*x = FetchAuthorizedKeysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supervisor_v1alpha1_ssh_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchAuthorizedKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchAuthorizedKeysResponse) ProtoMessage() {}

func (x *FetchAuthorizedKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supervisor_v1alpha1_ssh_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchAuthorizedKeysResponse.ProtoReflect.Descriptor instead.
func (*FetchAuthorizedKeysResponse) Descriptor() ([]byte, []int) {
	return file_supervisor_v1alpha1_ssh_proto_rawDescGZIP(), []int{1}
}

func (x *FetchAuthorizedKeysResponse) GetAuthorizedKeys() string {
	if x != nil {
		return x.AuthorizedKeys
	}
	return ""
}

var File_supervisor_v1alpha1_ssh_proto protoreflect.FileDescriptor

var file_supervisor_v1alpha1_ssh_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x22, 0x1c, 0x0a, 0x1a, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x46, 0x0a, 0x1b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x5f,
	0x6b, 0x65, 0x79, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x32, 0x84, 0x01, 0x0a, 0x06, 0x53,
	0x73, 0x68, 0x41, 0x50, 0x49, 0x12, 0x7a, 0x0a, 0x13, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x2f, 0x2e, 0x73,
	0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e,
	0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0xfa, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76,
	0x69, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x08, 0x53,
	0x73, 0x68, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x65, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x65, 0x70, 0x73, 0x71, 0x75, 0x61, 0x72, 0x65,
	0x2d, 0x69, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x2d, 0x67, 0x72, 0x69, 0x64, 0x2f, 0x73, 0x75, 0x70,
	0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76,
	0x69, 0x73, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x73, 0x75,
	0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xf8, 0x01, 0x00, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x13, 0x53, 0x75, 0x70, 0x65,
	0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca,
	0x02, 0x13, 0x53, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x1f, 0x53, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73,
	0x6f, 0x72, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x53, 0x75, 0x70, 0x65, 0x72, 0x76,
	0x69, 0x73, 0x6f, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_supervisor_v1alpha1_ssh_proto_rawDescOnce sync.Once
	file_supervisor_v1alpha1_ssh_proto_rawDescData = file_supervisor_v1alpha1_ssh_proto_rawDesc
)

func file_supervisor_v1alpha1_ssh_proto_rawDescGZIP() []byte {
	file_supervisor_v1alpha1_ssh_proto_rawDescOnce.Do(func() {
		file_supervisor_v1alpha1_ssh_proto_rawDescData = protoimpl.X.CompressGZIP(file_supervisor_v1alpha1_ssh_proto_rawDescData)
	})
	return file_supervisor_v1alpha1_ssh_proto_rawDescData
}

var file_supervisor_v1alpha1_ssh_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_supervisor_v1alpha1_ssh_proto_goTypes = []interface{}{
	(*FetchAuthorizedKeysRequest)(nil),  // 0: supervisor.v1alpha1.FetchAuthorizedKeysRequest
	(*FetchAuthorizedKeysResponse)(nil), // 1: supervisor.v1alpha1.FetchAuthorizedKeysResponse
}
var file_supervisor_v1alpha1_ssh_proto_depIdxs = []int32{
	0, // 0: supervisor.v1alpha1.SshAPI.FetchAuthorizedKeys:input_type -> supervisor.v1alpha1.FetchAuthorizedKeysRequest
	1, // 1: supervisor.v1alpha1.SshAPI.FetchAuthorizedKeys:output_type -> supervisor.v1alpha1.FetchAuthorizedKeysResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_supervisor_v1alpha1_ssh_proto_init() }
func file_supervisor_v1alpha1_ssh_proto_init() {
	if File_supervisor_v1alpha1_ssh_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_supervisor_v1alpha1_ssh_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchAuthorizedKeysRequest); i {
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
		file_supervisor_v1alpha1_ssh_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchAuthorizedKeysResponse); i {
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
			RawDescriptor: file_supervisor_v1alpha1_ssh_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_supervisor_v1alpha1_ssh_proto_goTypes,
		DependencyIndexes: file_supervisor_v1alpha1_ssh_proto_depIdxs,
		MessageInfos:      file_supervisor_v1alpha1_ssh_proto_msgTypes,
	}.Build()
	File_supervisor_v1alpha1_ssh_proto = out.File
	file_supervisor_v1alpha1_ssh_proto_rawDesc = nil
	file_supervisor_v1alpha1_ssh_proto_goTypes = nil
	file_supervisor_v1alpha1_ssh_proto_depIdxs = nil
}
