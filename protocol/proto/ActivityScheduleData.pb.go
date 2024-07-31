// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ActivityScheduleData.proto

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

type ActivityScheduleData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActivityId uint32 `protobuf:"varint,6,opt,name=activity_id,json=activityId,proto3" json:"activity_id,omitempty"`
	EndTime    int64  `protobuf:"varint,7,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	BeginTime  int64  `protobuf:"varint,13,opt,name=begin_time,json=beginTime,proto3" json:"begin_time,omitempty"`
	PanelId    uint32 `protobuf:"varint,9,opt,name=panel_id,json=panelId,proto3" json:"panel_id,omitempty"`
}

func (x *ActivityScheduleData) Reset() {
	*x = ActivityScheduleData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ActivityScheduleData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityScheduleData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityScheduleData) ProtoMessage() {}

func (x *ActivityScheduleData) ProtoReflect() protoreflect.Message {
	mi := &file_ActivityScheduleData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityScheduleData.ProtoReflect.Descriptor instead.
func (*ActivityScheduleData) Descriptor() ([]byte, []int) {
	return file_ActivityScheduleData_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityScheduleData) GetActivityId() uint32 {
	if x != nil {
		return x.ActivityId
	}
	return 0
}

func (x *ActivityScheduleData) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *ActivityScheduleData) GetBeginTime() int64 {
	if x != nil {
		return x.BeginTime
	}
	return 0
}

func (x *ActivityScheduleData) GetPanelId() uint32 {
	if x != nil {
		return x.PanelId
	}
	return 0
}

var File_ActivityScheduleData_proto protoreflect.FileDescriptor

var file_ActivityScheduleData_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a,
	0x14, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x65, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x70, 0x61, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x42, 0x28, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ActivityScheduleData_proto_rawDescOnce sync.Once
	file_ActivityScheduleData_proto_rawDescData = file_ActivityScheduleData_proto_rawDesc
)

func file_ActivityScheduleData_proto_rawDescGZIP() []byte {
	file_ActivityScheduleData_proto_rawDescOnce.Do(func() {
		file_ActivityScheduleData_proto_rawDescData = protoimpl.X.CompressGZIP(file_ActivityScheduleData_proto_rawDescData)
	})
	return file_ActivityScheduleData_proto_rawDescData
}

var file_ActivityScheduleData_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ActivityScheduleData_proto_goTypes = []interface{}{
	(*ActivityScheduleData)(nil), // 0: ActivityScheduleData
}
var file_ActivityScheduleData_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ActivityScheduleData_proto_init() }
func file_ActivityScheduleData_proto_init() {
	if File_ActivityScheduleData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ActivityScheduleData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityScheduleData); i {
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
			RawDescriptor: file_ActivityScheduleData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ActivityScheduleData_proto_goTypes,
		DependencyIndexes: file_ActivityScheduleData_proto_depIdxs,
		MessageInfos:      file_ActivityScheduleData_proto_msgTypes,
	}.Build()
	File_ActivityScheduleData_proto = out.File
	file_ActivityScheduleData_proto_rawDesc = nil
	file_ActivityScheduleData_proto_goTypes = nil
	file_ActivityScheduleData_proto_depIdxs = nil
}
