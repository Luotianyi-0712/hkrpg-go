// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueTournLevel.proto

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

type RogueTournLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TournRoomList []*RogueTournRoomList `protobuf:"bytes,4,rep,name=tourn_room_list,json=tournRoomList,proto3" json:"tourn_room_list,omitempty"`
	LevelIndex    uint32                `protobuf:"varint,9,opt,name=level_index,json=levelIndex,proto3" json:"level_index,omitempty"`
	CurRoomIndex  uint32                `protobuf:"varint,3,opt,name=cur_room_index,json=curRoomIndex,proto3" json:"cur_room_index,omitempty"`
	LayerId       uint32                `protobuf:"varint,7,opt,name=layer_id,json=layerId,proto3" json:"layer_id,omitempty"`
	Status        RogueTournLayerStatus `protobuf:"varint,11,opt,name=status,proto3,enum=RogueTournLayerStatus" json:"status,omitempty"`
}

func (x *RogueTournLevel) Reset() {
	*x = RogueTournLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueTournLevel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueTournLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueTournLevel) ProtoMessage() {}

func (x *RogueTournLevel) ProtoReflect() protoreflect.Message {
	mi := &file_RogueTournLevel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueTournLevel.ProtoReflect.Descriptor instead.
func (*RogueTournLevel) Descriptor() ([]byte, []int) {
	return file_RogueTournLevel_proto_rawDescGZIP(), []int{0}
}

func (x *RogueTournLevel) GetTournRoomList() []*RogueTournRoomList {
	if x != nil {
		return x.TournRoomList
	}
	return nil
}

func (x *RogueTournLevel) GetLevelIndex() uint32 {
	if x != nil {
		return x.LevelIndex
	}
	return 0
}

func (x *RogueTournLevel) GetCurRoomIndex() uint32 {
	if x != nil {
		return x.CurRoomIndex
	}
	return 0
}

func (x *RogueTournLevel) GetLayerId() uint32 {
	if x != nil {
		return x.LayerId
	}
	return 0
}

func (x *RogueTournLevel) GetStatus() RogueTournLayerStatus {
	if x != nil {
		return x.Status
	}
	return RogueTournLayerStatus_ROGUE_TOURN_LAYER_STATUS_NONE
}

var File_RogueTournLevel_proto protoreflect.FileDescriptor

var file_RogueTournLevel_proto_rawDesc = []byte{
	0x0a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f,
	0x75, 0x72, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x4c, 0x61, 0x79,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0,
	0x01, 0x0a, 0x0f, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x12, 0x3b, 0x0a, 0x0f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x5f, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x0d, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x12, 0x24, 0x0a, 0x0e, 0x63, 0x75, 0x72, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x52, 0x6f, 0x6f,
	0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x4c, 0x61,
	0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueTournLevel_proto_rawDescOnce sync.Once
	file_RogueTournLevel_proto_rawDescData = file_RogueTournLevel_proto_rawDesc
)

func file_RogueTournLevel_proto_rawDescGZIP() []byte {
	file_RogueTournLevel_proto_rawDescOnce.Do(func() {
		file_RogueTournLevel_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueTournLevel_proto_rawDescData)
	})
	return file_RogueTournLevel_proto_rawDescData
}

var file_RogueTournLevel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueTournLevel_proto_goTypes = []interface{}{
	(*RogueTournLevel)(nil),    // 0: RogueTournLevel
	(*RogueTournRoomList)(nil), // 1: RogueTournRoomList
	(RogueTournLayerStatus)(0), // 2: RogueTournLayerStatus
}
var file_RogueTournLevel_proto_depIdxs = []int32{
	1, // 0: RogueTournLevel.tourn_room_list:type_name -> RogueTournRoomList
	2, // 1: RogueTournLevel.status:type_name -> RogueTournLayerStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RogueTournLevel_proto_init() }
func file_RogueTournLevel_proto_init() {
	if File_RogueTournLevel_proto != nil {
		return
	}
	file_RogueTournRoomList_proto_init()
	file_RogueTournLayerStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueTournLevel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueTournLevel); i {
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
			RawDescriptor: file_RogueTournLevel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueTournLevel_proto_goTypes,
		DependencyIndexes: file_RogueTournLevel_proto_depIdxs,
		MessageInfos:      file_RogueTournLevel_proto_msgTypes,
	}.Build()
	File_RogueTournLevel_proto = out.File
	file_RogueTournLevel_proto_rawDesc = nil
	file_RogueTournLevel_proto_goTypes = nil
	file_RogueTournLevel_proto_depIdxs = nil
}
