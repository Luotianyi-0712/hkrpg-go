// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EvolveBuildLevelInfo.proto

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

type EvolveBuildLevelInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BattleInfo       *EvolveBuildBattleInfo `protobuf:"bytes,9,opt,name=battle_info,json=battleInfo,proto3" json:"battle_info,omitempty"`
	CurGameExp       uint32                 `protobuf:"varint,7,opt,name=cur_game_exp,json=curGameExp,proto3" json:"cur_game_exp,omitempty"`
	BattleTargetList []*BattleTarget        `protobuf:"bytes,3,rep,name=battle_target_list,json=battleTargetList,proto3" json:"battle_target_list,omitempty"`
	RoundCnt         uint32                 `protobuf:"varint,10,opt,name=round_cnt,json=roundCnt,proto3" json:"round_cnt,omitempty"`
	PeriodIdList     []uint32               `protobuf:"varint,2,rep,packed,name=period_id_list,json=periodIdList,proto3" json:"period_id_list,omitempty"`
	AvatarList       []*EvolveBuildAvatar   `protobuf:"bytes,1,rep,name=avatar_list,json=avatarList,proto3" json:"avatar_list,omitempty"`
}

func (x *EvolveBuildLevelInfo) Reset() {
	*x = EvolveBuildLevelInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvolveBuildLevelInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvolveBuildLevelInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvolveBuildLevelInfo) ProtoMessage() {}

func (x *EvolveBuildLevelInfo) ProtoReflect() protoreflect.Message {
	mi := &file_EvolveBuildLevelInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvolveBuildLevelInfo.ProtoReflect.Descriptor instead.
func (*EvolveBuildLevelInfo) Descriptor() ([]byte, []int) {
	return file_EvolveBuildLevelInfo_proto_rawDescGZIP(), []int{0}
}

func (x *EvolveBuildLevelInfo) GetBattleInfo() *EvolveBuildBattleInfo {
	if x != nil {
		return x.BattleInfo
	}
	return nil
}

func (x *EvolveBuildLevelInfo) GetCurGameExp() uint32 {
	if x != nil {
		return x.CurGameExp
	}
	return 0
}

func (x *EvolveBuildLevelInfo) GetBattleTargetList() []*BattleTarget {
	if x != nil {
		return x.BattleTargetList
	}
	return nil
}

func (x *EvolveBuildLevelInfo) GetRoundCnt() uint32 {
	if x != nil {
		return x.RoundCnt
	}
	return 0
}

func (x *EvolveBuildLevelInfo) GetPeriodIdList() []uint32 {
	if x != nil {
		return x.PeriodIdList
	}
	return nil
}

func (x *EvolveBuildLevelInfo) GetAvatarList() []*EvolveBuildAvatar {
	if x != nil {
		return x.AvatarList
	}
	return nil
}

var File_EvolveBuildLevelInfo_proto protoreflect.FileDescriptor

var file_EvolveBuildLevelInfo_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x45, 0x76,
	0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x45, 0x76, 0x6f, 0x6c, 0x76,
	0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x02, 0x0a, 0x14, 0x45, 0x76, 0x6f, 0x6c, 0x76,
	0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x37, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x62, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x5f,
	0x67, 0x61, 0x6d, 0x65, 0x5f, 0x65, 0x78, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x63, 0x75, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x45, 0x78, 0x70, 0x12, 0x3b, 0x0a, 0x12, 0x62, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x10, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x75, 0x6e, 0x64,
	0x5f, 0x63, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x43, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x5f, 0x69,
	0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0b, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_EvolveBuildLevelInfo_proto_rawDescOnce sync.Once
	file_EvolveBuildLevelInfo_proto_rawDescData = file_EvolveBuildLevelInfo_proto_rawDesc
)

func file_EvolveBuildLevelInfo_proto_rawDescGZIP() []byte {
	file_EvolveBuildLevelInfo_proto_rawDescOnce.Do(func() {
		file_EvolveBuildLevelInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvolveBuildLevelInfo_proto_rawDescData)
	})
	return file_EvolveBuildLevelInfo_proto_rawDescData
}

var file_EvolveBuildLevelInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvolveBuildLevelInfo_proto_goTypes = []interface{}{
	(*EvolveBuildLevelInfo)(nil),  // 0: EvolveBuildLevelInfo
	(*EvolveBuildBattleInfo)(nil), // 1: EvolveBuildBattleInfo
	(*BattleTarget)(nil),          // 2: BattleTarget
	(*EvolveBuildAvatar)(nil),     // 3: EvolveBuildAvatar
}
var file_EvolveBuildLevelInfo_proto_depIdxs = []int32{
	1, // 0: EvolveBuildLevelInfo.battle_info:type_name -> EvolveBuildBattleInfo
	2, // 1: EvolveBuildLevelInfo.battle_target_list:type_name -> BattleTarget
	3, // 2: EvolveBuildLevelInfo.avatar_list:type_name -> EvolveBuildAvatar
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_EvolveBuildLevelInfo_proto_init() }
func file_EvolveBuildLevelInfo_proto_init() {
	if File_EvolveBuildLevelInfo_proto != nil {
		return
	}
	file_EvolveBuildAvatar_proto_init()
	file_BattleTarget_proto_init()
	file_EvolveBuildBattleInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvolveBuildLevelInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvolveBuildLevelInfo); i {
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
			RawDescriptor: file_EvolveBuildLevelInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvolveBuildLevelInfo_proto_goTypes,
		DependencyIndexes: file_EvolveBuildLevelInfo_proto_depIdxs,
		MessageInfos:      file_EvolveBuildLevelInfo_proto_msgTypes,
	}.Build()
	File_EvolveBuildLevelInfo_proto = out.File
	file_EvolveBuildLevelInfo_proto_rawDesc = nil
	file_EvolveBuildLevelInfo_proto_goTypes = nil
	file_EvolveBuildLevelInfo_proto_depIdxs = nil
}
