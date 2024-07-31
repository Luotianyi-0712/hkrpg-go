// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueEnterScRsp.proto

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

type ChessRogueEnterScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerInfo       *ChessRoguePlayerInfo    `protobuf:"bytes,15,opt,name=player_info,json=playerInfo,proto3" json:"player_info,omitempty"`
	Retcode          uint32                   `protobuf:"varint,4,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Info             *ChessRogueCurrentInfo   `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	Id               uint32                   `protobuf:"varint,12,opt,name=id,proto3" json:"id,omitempty"`
	RogueCurrentInfo *ChessRogueQueryGameInfo `protobuf:"bytes,8,opt,name=rogue_current_info,json=rogueCurrentInfo,proto3" json:"rogue_current_info,omitempty"`
}

func (x *ChessRogueEnterScRsp) Reset() {
	*x = ChessRogueEnterScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueEnterScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueEnterScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueEnterScRsp) ProtoMessage() {}

func (x *ChessRogueEnterScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueEnterScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueEnterScRsp.ProtoReflect.Descriptor instead.
func (*ChessRogueEnterScRsp) Descriptor() ([]byte, []int) {
	return file_ChessRogueEnterScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueEnterScRsp) GetPlayerInfo() *ChessRoguePlayerInfo {
	if x != nil {
		return x.PlayerInfo
	}
	return nil
}

func (x *ChessRogueEnterScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ChessRogueEnterScRsp) GetInfo() *ChessRogueCurrentInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ChessRogueEnterScRsp) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChessRogueEnterScRsp) GetRogueCurrentInfo() *ChessRogueQueryGameInfo {
	if x != nil {
		return x.RogueCurrentInfo
	}
	return nil
}

var File_ChessRogueEnterScRsp_proto protoreflect.FileDescriptor

var file_ChessRogueEnterScRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x43, 0x68,
	0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a, 0x14, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x36, 0x0a,
	0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x2a, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x46, 0x0a, 0x12, 0x72,
	0x6f, 0x67, 0x75, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x10, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueEnterScRsp_proto_rawDescOnce sync.Once
	file_ChessRogueEnterScRsp_proto_rawDescData = file_ChessRogueEnterScRsp_proto_rawDesc
)

func file_ChessRogueEnterScRsp_proto_rawDescGZIP() []byte {
	file_ChessRogueEnterScRsp_proto_rawDescOnce.Do(func() {
		file_ChessRogueEnterScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueEnterScRsp_proto_rawDescData)
	})
	return file_ChessRogueEnterScRsp_proto_rawDescData
}

var file_ChessRogueEnterScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueEnterScRsp_proto_goTypes = []interface{}{
	(*ChessRogueEnterScRsp)(nil),    // 0: ChessRogueEnterScRsp
	(*ChessRoguePlayerInfo)(nil),    // 1: ChessRoguePlayerInfo
	(*ChessRogueCurrentInfo)(nil),   // 2: ChessRogueCurrentInfo
	(*ChessRogueQueryGameInfo)(nil), // 3: ChessRogueQueryGameInfo
}
var file_ChessRogueEnterScRsp_proto_depIdxs = []int32{
	1, // 0: ChessRogueEnterScRsp.player_info:type_name -> ChessRoguePlayerInfo
	2, // 1: ChessRogueEnterScRsp.info:type_name -> ChessRogueCurrentInfo
	3, // 2: ChessRogueEnterScRsp.rogue_current_info:type_name -> ChessRogueQueryGameInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ChessRogueEnterScRsp_proto_init() }
func file_ChessRogueEnterScRsp_proto_init() {
	if File_ChessRogueEnterScRsp_proto != nil {
		return
	}
	file_ChessRoguePlayerInfo_proto_init()
	file_ChessRogueCurrentInfo_proto_init()
	file_ChessRogueQueryGameInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueEnterScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueEnterScRsp); i {
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
			RawDescriptor: file_ChessRogueEnterScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueEnterScRsp_proto_goTypes,
		DependencyIndexes: file_ChessRogueEnterScRsp_proto_depIdxs,
		MessageInfos:      file_ChessRogueEnterScRsp_proto_msgTypes,
	}.Build()
	File_ChessRogueEnterScRsp_proto = out.File
	file_ChessRogueEnterScRsp_proto_rawDesc = nil
	file_ChessRogueEnterScRsp_proto_goTypes = nil
	file_ChessRogueEnterScRsp_proto_depIdxs = nil
}
