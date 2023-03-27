// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: supervisor/v1alpha1/job.proto

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

type JobStatus int32

const (
	JobStatus_JOB_STATUS_UNSPECIFIED    JobStatus = 0
	JobStatus_JOB_STATUS_PENDING        JobStatus = 1
	JobStatus_JOB_STATUS_META_SCHEDULED JobStatus = 2
	JobStatus_JOB_STATUS_SCHEDULED      JobStatus = 3
	JobStatus_JOB_STATUS_RUNNING        JobStatus = 4
	JobStatus_JOB_STATUS_CANCELLING     JobStatus = 5
	JobStatus_JOB_STATUS_CANCELLED      JobStatus = 6
	JobStatus_JOB_STATUS_FINISHED       JobStatus = 7
	JobStatus_JOB_STATUS_FAILED         JobStatus = 8
	JobStatus_JOB_STATUS_OUT_OF_CREDITS JobStatus = 9
)

// Enum value maps for JobStatus.
var (
	JobStatus_name = map[int32]string{
		0: "JOB_STATUS_UNSPECIFIED",
		1: "JOB_STATUS_PENDING",
		2: "JOB_STATUS_META_SCHEDULED",
		3: "JOB_STATUS_SCHEDULED",
		4: "JOB_STATUS_RUNNING",
		5: "JOB_STATUS_CANCELLING",
		6: "JOB_STATUS_CANCELLED",
		7: "JOB_STATUS_FINISHED",
		8: "JOB_STATUS_FAILED",
		9: "JOB_STATUS_OUT_OF_CREDITS",
	}
	JobStatus_value = map[string]int32{
		"JOB_STATUS_UNSPECIFIED":    0,
		"JOB_STATUS_PENDING":        1,
		"JOB_STATUS_META_SCHEDULED": 2,
		"JOB_STATUS_SCHEDULED":      3,
		"JOB_STATUS_RUNNING":        4,
		"JOB_STATUS_CANCELLING":     5,
		"JOB_STATUS_CANCELLED":      6,
		"JOB_STATUS_FINISHED":       7,
		"JOB_STATUS_FAILED":         8,
		"JOB_STATUS_OUT_OF_CREDITS": 9,
	}
)

func (x JobStatus) Enum() *JobStatus {
	p := new(JobStatus)
	*p = x
	return p
}

func (x JobStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_supervisor_v1alpha1_job_proto_enumTypes[0].Descriptor()
}

func (JobStatus) Type() protoreflect.EnumType {
	return &file_supervisor_v1alpha1_job_proto_enumTypes[0]
}

func (x JobStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobStatus.Descriptor instead.
func (JobStatus) EnumDescriptor() ([]byte, []int) {
	return file_supervisor_v1alpha1_job_proto_rawDescGZIP(), []int{0}
}

type SetJobStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id       uint64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Duration uint64    `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Status   JobStatus `protobuf:"varint,4,opt,name=status,proto3,enum=supervisor.v1alpha1.JobStatus" json:"status,omitempty"`
}

func (x *SetJobStatusRequest) Reset() {
	*x = SetJobStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supervisor_v1alpha1_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetJobStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetJobStatusRequest) ProtoMessage() {}

func (x *SetJobStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_supervisor_v1alpha1_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetJobStatusRequest.ProtoReflect.Descriptor instead.
func (*SetJobStatusRequest) Descriptor() ([]byte, []int) {
	return file_supervisor_v1alpha1_job_proto_rawDescGZIP(), []int{0}
}

func (x *SetJobStatusRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SetJobStatusRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SetJobStatusRequest) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *SetJobStatusRequest) GetStatus() JobStatus {
	if x != nil {
		return x.Status
	}
	return JobStatus_JOB_STATUS_UNSPECIFIED
}

type SetJobStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetJobStatusResponse) Reset() {
	*x = SetJobStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supervisor_v1alpha1_job_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetJobStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetJobStatusResponse) ProtoMessage() {}

func (x *SetJobStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_supervisor_v1alpha1_job_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetJobStatusResponse.ProtoReflect.Descriptor instead.
func (*SetJobStatusResponse) Descriptor() ([]byte, []int) {
	return file_supervisor_v1alpha1_job_proto_rawDescGZIP(), []int{1}
}

var File_supervisor_v1alpha1_job_proto protoreflect.FileDescriptor

var file_supervisor_v1alpha1_job_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6a, 0x6f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x22, 0x8d, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x73,
	0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x94, 0x02, 0x0a,
	0x09, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x16, 0x4a, 0x4f,
	0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1d,
	0x0a, 0x19, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4d, 0x45, 0x54,
	0x41, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x18, 0x0a,
	0x14, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x43, 0x48, 0x45,
	0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x16, 0x0a, 0x12, 0x4a, 0x4f, 0x42, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12,
	0x19, 0x0a, 0x15, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x41,
	0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x18, 0x0a, 0x14, 0x4a, 0x4f,
	0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c,
	0x45, 0x44, 0x10, 0x06, 0x12, 0x17, 0x0a, 0x13, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x07, 0x12, 0x15, 0x0a,
	0x11, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x08, 0x12, 0x1d, 0x0a, 0x19, 0x4a, 0x4f, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x49, 0x54,
	0x53, 0x10, 0x09, 0x32, 0x6f, 0x0a, 0x06, 0x4a, 0x6f, 0x62, 0x41, 0x50, 0x49, 0x12, 0x65, 0x0a,
	0x0c, 0x53, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x2e,
	0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76,
	0x69, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65,
	0x74, 0x4a, 0x6f, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0xfa, 0x01, 0x0a, 0x17, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x75, 0x70,
	0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x42, 0x08, 0x4a, 0x6f, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x65, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x65, 0x65, 0x70, 0x73, 0x71, 0x75,
	0x61, 0x72, 0x65, 0x2d, 0x69, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x2d, 0x67, 0x72, 0x69, 0x64, 0x2f,
	0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x75, 0x70,
	0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3b, 0x73, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xf8, 0x01, 0x00, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x13, 0x53,
	0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xca, 0x02, 0x13, 0x53, 0x75, 0x70, 0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x1f, 0x53, 0x75, 0x70, 0x65, 0x72,
	0x76, 0x69, 0x73, 0x6f, 0x72, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x53, 0x75, 0x70,
	0x65, 0x72, 0x76, 0x69, 0x73, 0x6f, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_supervisor_v1alpha1_job_proto_rawDescOnce sync.Once
	file_supervisor_v1alpha1_job_proto_rawDescData = file_supervisor_v1alpha1_job_proto_rawDesc
)

func file_supervisor_v1alpha1_job_proto_rawDescGZIP() []byte {
	file_supervisor_v1alpha1_job_proto_rawDescOnce.Do(func() {
		file_supervisor_v1alpha1_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_supervisor_v1alpha1_job_proto_rawDescData)
	})
	return file_supervisor_v1alpha1_job_proto_rawDescData
}

var file_supervisor_v1alpha1_job_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_supervisor_v1alpha1_job_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_supervisor_v1alpha1_job_proto_goTypes = []interface{}{
	(JobStatus)(0),               // 0: supervisor.v1alpha1.JobStatus
	(*SetJobStatusRequest)(nil),  // 1: supervisor.v1alpha1.SetJobStatusRequest
	(*SetJobStatusResponse)(nil), // 2: supervisor.v1alpha1.SetJobStatusResponse
}
var file_supervisor_v1alpha1_job_proto_depIdxs = []int32{
	0, // 0: supervisor.v1alpha1.SetJobStatusRequest.status:type_name -> supervisor.v1alpha1.JobStatus
	1, // 1: supervisor.v1alpha1.JobAPI.SetJobStatus:input_type -> supervisor.v1alpha1.SetJobStatusRequest
	2, // 2: supervisor.v1alpha1.JobAPI.SetJobStatus:output_type -> supervisor.v1alpha1.SetJobStatusResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_supervisor_v1alpha1_job_proto_init() }
func file_supervisor_v1alpha1_job_proto_init() {
	if File_supervisor_v1alpha1_job_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_supervisor_v1alpha1_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetJobStatusRequest); i {
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
		file_supervisor_v1alpha1_job_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetJobStatusResponse); i {
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
			RawDescriptor: file_supervisor_v1alpha1_job_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_supervisor_v1alpha1_job_proto_goTypes,
		DependencyIndexes: file_supervisor_v1alpha1_job_proto_depIdxs,
		EnumInfos:         file_supervisor_v1alpha1_job_proto_enumTypes,
		MessageInfos:      file_supervisor_v1alpha1_job_proto_msgTypes,
	}.Build()
	File_supervisor_v1alpha1_job_proto = out.File
	file_supervisor_v1alpha1_job_proto_rawDesc = nil
	file_supervisor_v1alpha1_job_proto_goTypes = nil
	file_supervisor_v1alpha1_job_proto_depIdxs = nil
}
