// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SceneNpcMonsterInfo.proto

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

type SceneNpcMonsterInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AHNAIBGMCDB bool                 `protobuf:"varint,11,opt,name=AHNAIBGMCDB,proto3" json:"AHNAIBGMCDB,omitempty"`
	MonsterId   uint32               `protobuf:"varint,12,opt,name=monster_id,json=monsterId,proto3" json:"monster_id,omitempty"`
	EventId     uint32               `protobuf:"varint,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	MNGIHHFBPLO bool                 `protobuf:"varint,3,opt,name=MNGIHHFBPLO,proto3" json:"MNGIHHFBPLO,omitempty"`
	WorldLevel  uint32               `protobuf:"varint,10,opt,name=world_level,json=worldLevel,proto3" json:"world_level,omitempty"`
	ExtraInfo   *NpcMonsterExtraInfo `protobuf:"bytes,2,opt,name=extra_info,json=extraInfo,proto3" json:"extra_info,omitempty"`
}

func (x *SceneNpcMonsterInfo) Reset() {
	*x = SceneNpcMonsterInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneNpcMonsterInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneNpcMonsterInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneNpcMonsterInfo) ProtoMessage() {}

func (x *SceneNpcMonsterInfo) ProtoReflect() protoreflect.Message {
	mi := &file_SceneNpcMonsterInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneNpcMonsterInfo.ProtoReflect.Descriptor instead.
func (*SceneNpcMonsterInfo) Descriptor() ([]byte, []int) {
	return file_SceneNpcMonsterInfo_proto_rawDescGZIP(), []int{0}
}

func (x *SceneNpcMonsterInfo) GetAHNAIBGMCDB() bool {
	if x != nil {
		return x.AHNAIBGMCDB
	}
	return false
}

func (x *SceneNpcMonsterInfo) GetMonsterId() uint32 {
	if x != nil {
		return x.MonsterId
	}
	return 0
}

func (x *SceneNpcMonsterInfo) GetEventId() uint32 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *SceneNpcMonsterInfo) GetMNGIHHFBPLO() bool {
	if x != nil {
		return x.MNGIHHFBPLO
	}
	return false
}

func (x *SceneNpcMonsterInfo) GetWorldLevel() uint32 {
	if x != nil {
		return x.WorldLevel
	}
	return 0
}

func (x *SceneNpcMonsterInfo) GetExtraInfo() *NpcMonsterExtraInfo {
	if x != nil {
		return x.ExtraInfo
	}
	return nil
}

var File_SceneNpcMonsterInfo_proto protoreflect.FileDescriptor

var file_SceneNpcMonsterInfo_proto_rawDesc = []byte{
	0x0a, 0x19, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4e, 0x70, 0x63, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x4e, 0x70, 0x63,
	0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x45, 0x78, 0x74, 0x72, 0x61, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe9, 0x01, 0x0a, 0x13, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x4e, 0x70, 0x63, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20,
	0x0a, 0x0b, 0x41, 0x48, 0x4e, 0x41, 0x49, 0x42, 0x47, 0x4d, 0x43, 0x44, 0x42, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x41, 0x48, 0x4e, 0x41, 0x49, 0x42, 0x47, 0x4d, 0x43, 0x44, 0x42,
	0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x4e,
	0x47, 0x49, 0x48, 0x48, 0x46, 0x42, 0x50, 0x4c, 0x4f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x4d, 0x4e, 0x47, 0x49, 0x48, 0x48, 0x46, 0x42, 0x50, 0x4c, 0x4f, 0x12, 0x1f, 0x0a, 0x0b,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x33, 0x0a,
	0x0a, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x4e, 0x70, 0x63, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x65, 0x78, 0x74, 0x72, 0x61, 0x49, 0x6e,
	0x66, 0x6f, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneNpcMonsterInfo_proto_rawDescOnce sync.Once
	file_SceneNpcMonsterInfo_proto_rawDescData = file_SceneNpcMonsterInfo_proto_rawDesc
)

func file_SceneNpcMonsterInfo_proto_rawDescGZIP() []byte {
	file_SceneNpcMonsterInfo_proto_rawDescOnce.Do(func() {
		file_SceneNpcMonsterInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneNpcMonsterInfo_proto_rawDescData)
	})
	return file_SceneNpcMonsterInfo_proto_rawDescData
}

var file_SceneNpcMonsterInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneNpcMonsterInfo_proto_goTypes = []interface{}{
	(*SceneNpcMonsterInfo)(nil), // 0: SceneNpcMonsterInfo
	(*NpcMonsterExtraInfo)(nil), // 1: NpcMonsterExtraInfo
}
var file_SceneNpcMonsterInfo_proto_depIdxs = []int32{
	1, // 0: SceneNpcMonsterInfo.extra_info:type_name -> NpcMonsterExtraInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SceneNpcMonsterInfo_proto_init() }
func file_SceneNpcMonsterInfo_proto_init() {
	if File_SceneNpcMonsterInfo_proto != nil {
		return
	}
	file_NpcMonsterExtraInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneNpcMonsterInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneNpcMonsterInfo); i {
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
			RawDescriptor: file_SceneNpcMonsterInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneNpcMonsterInfo_proto_goTypes,
		DependencyIndexes: file_SceneNpcMonsterInfo_proto_depIdxs,
		MessageInfos:      file_SceneNpcMonsterInfo_proto_msgTypes,
	}.Build()
	File_SceneNpcMonsterInfo_proto = out.File
	file_SceneNpcMonsterInfo_proto_rawDesc = nil
	file_SceneNpcMonsterInfo_proto_goTypes = nil
	file_SceneNpcMonsterInfo_proto_depIdxs = nil
}
