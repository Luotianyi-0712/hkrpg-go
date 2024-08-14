// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueTournCurSceneInfo.proto

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

type RogueTournCurSceneInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scene      *SceneInfo          `protobuf:"bytes,13,opt,name=scene,proto3" json:"scene,omitempty"`
	Lineup     *LineupInfo         `protobuf:"bytes,2,opt,name=lineup,proto3" json:"lineup,omitempty"`
	RotateInfo *RogueMapRotateInfo `protobuf:"bytes,8,opt,name=rotate_info,json=rotateInfo,proto3" json:"rotate_info,omitempty"`
}

func (x *RogueTournCurSceneInfo) Reset() {
	*x = RogueTournCurSceneInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueTournCurSceneInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueTournCurSceneInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueTournCurSceneInfo) ProtoMessage() {}

func (x *RogueTournCurSceneInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueTournCurSceneInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueTournCurSceneInfo.ProtoReflect.Descriptor instead.
func (*RogueTournCurSceneInfo) Descriptor() ([]byte, []int) {
	return file_RogueTournCurSceneInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueTournCurSceneInfo) GetScene() *SceneInfo {
	if x != nil {
		return x.Scene
	}
	return nil
}

func (x *RogueTournCurSceneInfo) GetLineup() *LineupInfo {
	if x != nil {
		return x.Lineup
	}
	return nil
}

func (x *RogueTournCurSceneInfo) GetRotateInfo() *RogueMapRotateInfo {
	if x != nil {
		return x.RotateInfo
	}
	return nil
}

var File_RogueTournCurSceneInfo_proto protoreflect.FileDescriptor

var file_RogueTournCurSceneInfo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75, 0x72, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10,
	0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x95, 0x01, 0x0a, 0x16,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x43, 0x75, 0x72, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x05, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x12, 0x23, 0x0a, 0x06, 0x6c, 0x69, 0x6e, 0x65,
	0x75, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x75,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x12, 0x34, 0x0a,
	0x0b, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x61, 0x70, 0x52, 0x6f, 0x74,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueTournCurSceneInfo_proto_rawDescOnce sync.Once
	file_RogueTournCurSceneInfo_proto_rawDescData = file_RogueTournCurSceneInfo_proto_rawDesc
)

func file_RogueTournCurSceneInfo_proto_rawDescGZIP() []byte {
	file_RogueTournCurSceneInfo_proto_rawDescOnce.Do(func() {
		file_RogueTournCurSceneInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueTournCurSceneInfo_proto_rawDescData)
	})
	return file_RogueTournCurSceneInfo_proto_rawDescData
}

var file_RogueTournCurSceneInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueTournCurSceneInfo_proto_goTypes = []interface{}{
	(*RogueTournCurSceneInfo)(nil), // 0: RogueTournCurSceneInfo
	(*SceneInfo)(nil),              // 1: SceneInfo
	(*LineupInfo)(nil),             // 2: LineupInfo
	(*RogueMapRotateInfo)(nil),     // 3: RogueMapRotateInfo
}
var file_RogueTournCurSceneInfo_proto_depIdxs = []int32{
	1, // 0: RogueTournCurSceneInfo.scene:type_name -> SceneInfo
	2, // 1: RogueTournCurSceneInfo.lineup:type_name -> LineupInfo
	3, // 2: RogueTournCurSceneInfo.rotate_info:type_name -> RogueMapRotateInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_RogueTournCurSceneInfo_proto_init() }
func file_RogueTournCurSceneInfo_proto_init() {
	if File_RogueTournCurSceneInfo_proto != nil {
		return
	}
	file_LineupInfo_proto_init()
	file_RogueMapRotateInfo_proto_init()
	file_SceneInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueTournCurSceneInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueTournCurSceneInfo); i {
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
			RawDescriptor: file_RogueTournCurSceneInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueTournCurSceneInfo_proto_goTypes,
		DependencyIndexes: file_RogueTournCurSceneInfo_proto_depIdxs,
		MessageInfos:      file_RogueTournCurSceneInfo_proto_msgTypes,
	}.Build()
	File_RogueTournCurSceneInfo_proto = out.File
	file_RogueTournCurSceneInfo_proto_rawDesc = nil
	file_RogueTournCurSceneInfo_proto_goTypes = nil
	file_RogueTournCurSceneInfo_proto_depIdxs = nil
}
