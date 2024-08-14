// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueGetInfo.proto

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

type ChessRogueGetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RogueDifficultyInfo *ChessRogueQueryDiffcultyInfo `protobuf:"bytes,8,opt,name=rogue_difficulty_info,json=rogueDifficultyInfo,proto3" json:"rogue_difficulty_info,omitempty"`
	AreaIdList          []uint32                      `protobuf:"varint,13,rep,packed,name=area_id_list,json=areaIdList,proto3" json:"area_id_list,omitempty"`
	TalentInfoList      *ChessRogueTalentInfo         `protobuf:"bytes,9,opt,name=talent_info_list,json=talentInfoList,proto3" json:"talent_info_list,omitempty"`
	QueryDiceInfo       *ChessRogueQueryDiceInfo      `protobuf:"bytes,3,opt,name=query_dice_info,json=queryDiceInfo,proto3" json:"query_dice_info,omitempty"`
	ChessAeonInfo       *ChessRogueQueryAeonInfo      `protobuf:"bytes,1,opt,name=chess_aeon_info,json=chessAeonInfo,proto3" json:"chess_aeon_info,omitempty"`
	ExploredAreaIdList  []uint32                      `protobuf:"varint,14,rep,packed,name=explored_area_id_list,json=exploredAreaIdList,proto3" json:"explored_area_id_list,omitempty"`
}

func (x *ChessRogueGetInfo) Reset() {
	*x = ChessRogueGetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueGetInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueGetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueGetInfo) ProtoMessage() {}

func (x *ChessRogueGetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueGetInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueGetInfo.ProtoReflect.Descriptor instead.
func (*ChessRogueGetInfo) Descriptor() ([]byte, []int) {
	return file_ChessRogueGetInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueGetInfo) GetRogueDifficultyInfo() *ChessRogueQueryDiffcultyInfo {
	if x != nil {
		return x.RogueDifficultyInfo
	}
	return nil
}

func (x *ChessRogueGetInfo) GetAreaIdList() []uint32 {
	if x != nil {
		return x.AreaIdList
	}
	return nil
}

func (x *ChessRogueGetInfo) GetTalentInfoList() *ChessRogueTalentInfo {
	if x != nil {
		return x.TalentInfoList
	}
	return nil
}

func (x *ChessRogueGetInfo) GetQueryDiceInfo() *ChessRogueQueryDiceInfo {
	if x != nil {
		return x.QueryDiceInfo
	}
	return nil
}

func (x *ChessRogueGetInfo) GetChessAeonInfo() *ChessRogueQueryAeonInfo {
	if x != nil {
		return x.ChessAeonInfo
	}
	return nil
}

func (x *ChessRogueGetInfo) GetExploredAreaIdList() []uint32 {
	if x != nil {
		return x.ExploredAreaIdList
	}
	return nil
}

var File_ChessRogueGetInfo_proto protoreflect.FileDescriptor

var file_ChessRogueGetInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x69, 0x66, 0x66, 0x63, 0x75,
	0x6c, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x65,
	0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x43, 0x68,
	0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x03, 0x0a, 0x11, 0x43, 0x68, 0x65, 0x73,
	0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x51, 0x0a,
	0x15, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x64, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74,
	0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x69,
	0x66, 0x66, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x13, 0x72, 0x6f, 0x67,
	0x75, 0x65, 0x44, 0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x20, 0x0a, 0x0c, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x3f, 0x0a, 0x10, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x64, 0x69, 0x63,
	0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x44, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x71, 0x75, 0x65, 0x72, 0x79, 0x44, 0x69, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x40, 0x0a, 0x0f, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x61,
	0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x41, 0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x63, 0x68, 0x65, 0x73, 0x73, 0x41,
	0x65, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x15, 0x65, 0x78, 0x70, 0x6c, 0x6f,
	0x72, 0x65, 0x64, 0x5f, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x12, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x64,
	0x41, 0x72, 0x65, 0x61, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueGetInfo_proto_rawDescOnce sync.Once
	file_ChessRogueGetInfo_proto_rawDescData = file_ChessRogueGetInfo_proto_rawDesc
)

func file_ChessRogueGetInfo_proto_rawDescGZIP() []byte {
	file_ChessRogueGetInfo_proto_rawDescOnce.Do(func() {
		file_ChessRogueGetInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueGetInfo_proto_rawDescData)
	})
	return file_ChessRogueGetInfo_proto_rawDescData
}

var file_ChessRogueGetInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueGetInfo_proto_goTypes = []interface{}{
	(*ChessRogueGetInfo)(nil),            // 0: ChessRogueGetInfo
	(*ChessRogueQueryDiffcultyInfo)(nil), // 1: ChessRogueQueryDiffcultyInfo
	(*ChessRogueTalentInfo)(nil),         // 2: ChessRogueTalentInfo
	(*ChessRogueQueryDiceInfo)(nil),      // 3: ChessRogueQueryDiceInfo
	(*ChessRogueQueryAeonInfo)(nil),      // 4: ChessRogueQueryAeonInfo
}
var file_ChessRogueGetInfo_proto_depIdxs = []int32{
	1, // 0: ChessRogueGetInfo.rogue_difficulty_info:type_name -> ChessRogueQueryDiffcultyInfo
	2, // 1: ChessRogueGetInfo.talent_info_list:type_name -> ChessRogueTalentInfo
	3, // 2: ChessRogueGetInfo.query_dice_info:type_name -> ChessRogueQueryDiceInfo
	4, // 3: ChessRogueGetInfo.chess_aeon_info:type_name -> ChessRogueQueryAeonInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ChessRogueGetInfo_proto_init() }
func file_ChessRogueGetInfo_proto_init() {
	if File_ChessRogueGetInfo_proto != nil {
		return
	}
	file_ChessRogueQueryDiffcultyInfo_proto_init()
	file_ChessRogueQueryAeonInfo_proto_init()
	file_ChessRogueTalentInfo_proto_init()
	file_ChessRogueQueryDiceInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueGetInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueGetInfo); i {
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
			RawDescriptor: file_ChessRogueGetInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueGetInfo_proto_goTypes,
		DependencyIndexes: file_ChessRogueGetInfo_proto_depIdxs,
		MessageInfos:      file_ChessRogueGetInfo_proto_msgTypes,
	}.Build()
	File_ChessRogueGetInfo_proto = out.File
	file_ChessRogueGetInfo_proto_rawDesc = nil
	file_ChessRogueGetInfo_proto_goTypes = nil
	file_ChessRogueGetInfo_proto_depIdxs = nil
}
