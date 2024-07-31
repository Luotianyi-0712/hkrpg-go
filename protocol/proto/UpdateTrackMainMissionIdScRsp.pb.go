// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: UpdateTrackMainMissionIdScRsp.proto

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

type UpdateTrackMainMissionIdScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode            uint32 `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
	TrackMissionId     uint32 `protobuf:"varint,13,opt,name=track_mission_id,json=trackMissionId,proto3" json:"track_mission_id,omitempty"`
	PrevTrackMissionId uint32 `protobuf:"varint,4,opt,name=prev_track_mission_id,json=prevTrackMissionId,proto3" json:"prev_track_mission_id,omitempty"`
}

func (x *UpdateTrackMainMissionIdScRsp) Reset() {
	*x = UpdateTrackMainMissionIdScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UpdateTrackMainMissionIdScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTrackMainMissionIdScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTrackMainMissionIdScRsp) ProtoMessage() {}

func (x *UpdateTrackMainMissionIdScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_UpdateTrackMainMissionIdScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTrackMainMissionIdScRsp.ProtoReflect.Descriptor instead.
func (*UpdateTrackMainMissionIdScRsp) Descriptor() ([]byte, []int) {
	return file_UpdateTrackMainMissionIdScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateTrackMainMissionIdScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *UpdateTrackMainMissionIdScRsp) GetTrackMissionId() uint32 {
	if x != nil {
		return x.TrackMissionId
	}
	return 0
}

func (x *UpdateTrackMainMissionIdScRsp) GetPrevTrackMissionId() uint32 {
	if x != nil {
		return x.PrevTrackMissionId
	}
	return 0
}

var File_UpdateTrackMainMissionIdScRsp_proto protoreflect.FileDescriptor

var file_UpdateTrackMainMissionIdScRsp_proto_rawDesc = []byte{
	0x0a, 0x23, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x4d, 0x61, 0x69,
	0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x1d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x4d, 0x61, 0x69, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x31, 0x0a, 0x15, 0x70,
	0x72, 0x65, 0x76, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x12, 0x70, 0x72, 0x65, 0x76,
	0x54, 0x72, 0x61, 0x63, 0x6b, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x28,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_UpdateTrackMainMissionIdScRsp_proto_rawDescOnce sync.Once
	file_UpdateTrackMainMissionIdScRsp_proto_rawDescData = file_UpdateTrackMainMissionIdScRsp_proto_rawDesc
)

func file_UpdateTrackMainMissionIdScRsp_proto_rawDescGZIP() []byte {
	file_UpdateTrackMainMissionIdScRsp_proto_rawDescOnce.Do(func() {
		file_UpdateTrackMainMissionIdScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_UpdateTrackMainMissionIdScRsp_proto_rawDescData)
	})
	return file_UpdateTrackMainMissionIdScRsp_proto_rawDescData
}

var file_UpdateTrackMainMissionIdScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_UpdateTrackMainMissionIdScRsp_proto_goTypes = []interface{}{
	(*UpdateTrackMainMissionIdScRsp)(nil), // 0: UpdateTrackMainMissionIdScRsp
}
var file_UpdateTrackMainMissionIdScRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_UpdateTrackMainMissionIdScRsp_proto_init() }
func file_UpdateTrackMainMissionIdScRsp_proto_init() {
	if File_UpdateTrackMainMissionIdScRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_UpdateTrackMainMissionIdScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTrackMainMissionIdScRsp); i {
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
			RawDescriptor: file_UpdateTrackMainMissionIdScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UpdateTrackMainMissionIdScRsp_proto_goTypes,
		DependencyIndexes: file_UpdateTrackMainMissionIdScRsp_proto_depIdxs,
		MessageInfos:      file_UpdateTrackMainMissionIdScRsp_proto_msgTypes,
	}.Build()
	File_UpdateTrackMainMissionIdScRsp_proto = out.File
	file_UpdateTrackMainMissionIdScRsp_proto_rawDesc = nil
	file_UpdateTrackMainMissionIdScRsp_proto_goTypes = nil
	file_UpdateTrackMainMissionIdScRsp_proto_depIdxs = nil
}
