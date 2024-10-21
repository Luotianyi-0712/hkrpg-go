// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TrackMainMissionUpdate.proto

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

type TrackMainMissionUpdate int32

const (
	TrackMainMissionUpdate_TRACK_MAIN_MISSION_UPDATE_NONE         TrackMainMissionUpdate = 0
	TrackMainMissionUpdate_TRACK_MAIN_MISSION_UPDATE_AUTO         TrackMainMissionUpdate = 1
	TrackMainMissionUpdate_TRACK_MAIN_MISSION_UPDATE_MANUAL       TrackMainMissionUpdate = 2
	TrackMainMissionUpdate_TRACK_MAIN_MISSION_UPDATE_LOGIN_REPORT TrackMainMissionUpdate = 3
)

// Enum value maps for TrackMainMissionUpdate.
var (
	TrackMainMissionUpdate_name = map[int32]string{
		0: "TRACK_MAIN_MISSION_UPDATE_NONE",
		1: "TRACK_MAIN_MISSION_UPDATE_AUTO",
		2: "TRACK_MAIN_MISSION_UPDATE_MANUAL",
		3: "TRACK_MAIN_MISSION_UPDATE_LOGIN_REPORT",
	}
	TrackMainMissionUpdate_value = map[string]int32{
		"TRACK_MAIN_MISSION_UPDATE_NONE":         0,
		"TRACK_MAIN_MISSION_UPDATE_AUTO":         1,
		"TRACK_MAIN_MISSION_UPDATE_MANUAL":       2,
		"TRACK_MAIN_MISSION_UPDATE_LOGIN_REPORT": 3,
	}
)

func (x TrackMainMissionUpdate) Enum() *TrackMainMissionUpdate {
	p := new(TrackMainMissionUpdate)
	*p = x
	return p
}

func (x TrackMainMissionUpdate) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TrackMainMissionUpdate) Descriptor() protoreflect.EnumDescriptor {
	return file_TrackMainMissionUpdate_proto_enumTypes[0].Descriptor()
}

func (TrackMainMissionUpdate) Type() protoreflect.EnumType {
	return &file_TrackMainMissionUpdate_proto_enumTypes[0]
}

func (x TrackMainMissionUpdate) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TrackMainMissionUpdate.Descriptor instead.
func (TrackMainMissionUpdate) EnumDescriptor() ([]byte, []int) {
	return file_TrackMainMissionUpdate_proto_rawDescGZIP(), []int{0}
}

var File_TrackMainMissionUpdate_proto protoreflect.FileDescriptor

var file_TrackMainMissionUpdate_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x4d, 0x61, 0x69, 0x6e, 0x4d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xb2,
	0x01, 0x0a, 0x16, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x4d, 0x61, 0x69, 0x6e, 0x4d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x1e, 0x54, 0x52, 0x41,
	0x43, 0x4b, 0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x22, 0x0a,
	0x1e, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x4d, 0x49, 0x53, 0x53,
	0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x55, 0x54, 0x4f, 0x10,
	0x01, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x52, 0x41, 0x43, 0x4b, 0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x5f,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x4d,
	0x41, 0x4e, 0x55, 0x41, 0x4c, 0x10, 0x02, 0x12, 0x2a, 0x0a, 0x26, 0x54, 0x52, 0x41, 0x43, 0x4b,
	0x5f, 0x4d, 0x41, 0x49, 0x4e, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52,
	0x54, 0x10, 0x03, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TrackMainMissionUpdate_proto_rawDescOnce sync.Once
	file_TrackMainMissionUpdate_proto_rawDescData = file_TrackMainMissionUpdate_proto_rawDesc
)

func file_TrackMainMissionUpdate_proto_rawDescGZIP() []byte {
	file_TrackMainMissionUpdate_proto_rawDescOnce.Do(func() {
		file_TrackMainMissionUpdate_proto_rawDescData = protoimpl.X.CompressGZIP(file_TrackMainMissionUpdate_proto_rawDescData)
	})
	return file_TrackMainMissionUpdate_proto_rawDescData
}

var file_TrackMainMissionUpdate_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_TrackMainMissionUpdate_proto_goTypes = []interface{}{
	(TrackMainMissionUpdate)(0), // 0: TrackMainMissionUpdate
}
var file_TrackMainMissionUpdate_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TrackMainMissionUpdate_proto_init() }
func file_TrackMainMissionUpdate_proto_init() {
	if File_TrackMainMissionUpdate_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_TrackMainMissionUpdate_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TrackMainMissionUpdate_proto_goTypes,
		DependencyIndexes: file_TrackMainMissionUpdate_proto_depIdxs,
		EnumInfos:         file_TrackMainMissionUpdate_proto_enumTypes,
	}.Build()
	File_TrackMainMissionUpdate_proto = out.File
	file_TrackMainMissionUpdate_proto_rawDesc = nil
	file_TrackMainMissionUpdate_proto_goTypes = nil
	file_TrackMainMissionUpdate_proto_depIdxs = nil
}
