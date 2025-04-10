//
//Copyright (c) 2024 Diagrid Inc.
//Licensed under the MIT License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.3
// source: proto/api/job.proto

package api

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

// Job defines a scheduled rhythmic job stored in the database.
// Job holds the desired spec of the job, not the current trigger state, held
// by Counter.
type Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// schedule is an optional schedule at which the job is to be run.
	// Accepts both systemd timer style cron expressions, as well as human
	// readable '@' prefixed period strings as defined below.
	//
	// Systemd timer style cron accepts 6 fields:
	// seconds | minutes | hours | day of month | month        | day of week
	// 0-59    | 0-59    | 0-23  | 1-31         | 1-12/jan-dec | 0-6/sun-sat
	//
	// "0 30 * * * *" - every hour on the half hour
	// "0 15 3 * * *" - every day at 03:15
	//
	// Period string expressions:
	// Entry                  | Description                                | Equivalent To
	// -----                  | -----------                                | -------------
	// @every <duration>      | Run every <duration> (e.g. '@every 1h30m') | N/A
	// @yearly (or @annually) | Run once a year, midnight, Jan. 1st        | 0 0 0 1 1 *
	// @monthly               | Run once a month, midnight, first of month | 0 0 0 1 * *
	// @weekly                | Run once a week, midnight on Sunday        | 0 0 0 * * 0
	// @daily (or @midnight)  | Run once a day, midnight                   | 0 0 0 * * *
	// @hourly                | Run once an hour, beginning of hour        | 0 0 * * * *
	Schedule *string `protobuf:"bytes,1,opt,name=schedule,proto3,oneof" json:"schedule,omitempty"`
	// due_time is the optional time at which the job should be active, or the
	// "one shot" time if other scheduling type fields are not provided.
	// Accepts a "point in time" string in the format of RFC3339, Go duration
	// string (therefore calculated from now), or non-repeating ISO8601.
	DueTime *string `protobuf:"bytes,2,opt,name=due_time,json=dueTime,proto3,oneof" json:"due_time,omitempty"`
	// ttl is the optional time to live or expiration of the job.
	// Accepts a "point in time" string in the format of RFC3339, Go duration
	// string (therefore calculated from now), or non-repeating ISO8601.
	Ttl *string `protobuf:"bytes,3,opt,name=ttl,proto3,oneof" json:"ttl,omitempty"`
	// repeats is the optional number of times in which the job should be
	// triggered. If not set, the job will run indefinitely or until expiration.
	Repeats *uint32 `protobuf:"varint,4,opt,name=repeats,proto3,oneof" json:"repeats,omitempty"`
	// metadata is a arbitrary metadata asociated with the job.
	Metadata *anypb.Any `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// payload is the serialized job payload that will be sent to the recipient
	// when the job is triggered.
	Payload *anypb.Any `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
	// failure_policy is the optional policy to apply when a job fails to
	// trigger.
	// By default, the failure policy is FailurePolicyConstant with a 1s interval
	// and 3 maximum retries.
	// See `failurepolicy.proto` for more information.
	FailurePolicy *FailurePolicy `protobuf:"bytes,7,opt,name=failure_policy,json=failurePolicy,proto3,oneof" json:"failure_policy,omitempty"`
}

func (x *Job) Reset() {
	*x = Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_job_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Job) ProtoMessage() {}

func (x *Job) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_job_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Job.ProtoReflect.Descriptor instead.
func (*Job) Descriptor() ([]byte, []int) {
	return file_proto_api_job_proto_rawDescGZIP(), []int{0}
}

func (x *Job) GetSchedule() string {
	if x != nil && x.Schedule != nil {
		return *x.Schedule
	}
	return ""
}

func (x *Job) GetDueTime() string {
	if x != nil && x.DueTime != nil {
		return *x.DueTime
	}
	return ""
}

func (x *Job) GetTtl() string {
	if x != nil && x.Ttl != nil {
		return *x.Ttl
	}
	return ""
}

func (x *Job) GetRepeats() uint32 {
	if x != nil && x.Repeats != nil {
		return *x.Repeats
	}
	return 0
}

func (x *Job) GetMetadata() *anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Job) GetPayload() *anypb.Any {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Job) GetFailurePolicy() *FailurePolicy {
	if x != nil {
		return x.FailurePolicy
	}
	return nil
}

var File_proto_api_job_proto protoreflect.FileDescriptor

var file_proto_api_job_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6a, 0x6f, 0x62, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x02, 0x0a, 0x03, 0x4a, 0x6f, 0x62, 0x12, 0x1f, 0x0a, 0x08,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x12, 0x1e, 0x0a,
	0x08, 0x64, 0x75, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x07, 0x64, 0x75, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a,
	0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x03, 0x74, 0x74,
	0x6c, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x03, 0x52, 0x07, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x73,
	0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2e, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x70, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x3e, 0x0a, 0x0e, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x48, 0x04, 0x52, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x64, 0x75, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42,
	0x06, 0x0a, 0x04, 0x5f, 0x74, 0x74, 0x6c, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x73, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x69, 0x61, 0x67, 0x72, 0x69, 0x64, 0x69, 0x6f, 0x2f, 0x67,
	0x6f, 0x2d, 0x65, 0x74, 0x63, 0x64, 0x2d, 0x63, 0x72, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_job_proto_rawDescOnce sync.Once
	file_proto_api_job_proto_rawDescData = file_proto_api_job_proto_rawDesc
)

func file_proto_api_job_proto_rawDescGZIP() []byte {
	file_proto_api_job_proto_rawDescOnce.Do(func() {
		file_proto_api_job_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_job_proto_rawDescData)
	})
	return file_proto_api_job_proto_rawDescData
}

var file_proto_api_job_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_api_job_proto_goTypes = []interface{}{
	(*Job)(nil),           // 0: api.Job
	(*anypb.Any)(nil),     // 1: google.protobuf.Any
	(*FailurePolicy)(nil), // 2: api.FailurePolicy
}
var file_proto_api_job_proto_depIdxs = []int32{
	1, // 0: api.Job.metadata:type_name -> google.protobuf.Any
	1, // 1: api.Job.payload:type_name -> google.protobuf.Any
	2, // 2: api.Job.failure_policy:type_name -> api.FailurePolicy
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_api_job_proto_init() }
func file_proto_api_job_proto_init() {
	if File_proto_api_job_proto != nil {
		return
	}
	file_proto_api_failurepolicy_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_api_job_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Job); i {
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
	file_proto_api_job_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_api_job_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_api_job_proto_goTypes,
		DependencyIndexes: file_proto_api_job_proto_depIdxs,
		MessageInfos:      file_proto_api_job_proto_msgTypes,
	}.Build()
	File_proto_api_job_proto = out.File
	file_proto_api_job_proto_rawDesc = nil
	file_proto_api_job_proto_goTypes = nil
	file_proto_api_job_proto_depIdxs = nil
}
