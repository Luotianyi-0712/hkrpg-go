// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueEnterCellScRsp.proto

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

type ChessRogueEnterCellScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode          uint32                   `protobuf:"varint,15,opt,name=retcode,proto3" json:"retcode,omitempty"`
	PlayerInfo       *ChessRoguePlayerInfo    `protobuf:"bytes,4,opt,name=player_info,json=playerInfo,proto3" json:"player_info,omitempty"`
	CellId           uint32                   `protobuf:"varint,11,opt,name=cell_id,json=cellId,proto3" json:"cell_id,omitempty"`
	RogueCurrentInfo *ChessRogueQueryGameInfo `protobuf:"bytes,5,opt,name=rogue_current_info,json=rogueCurrentInfo,proto3" json:"rogue_current_info,omitempty"`
	Info             *ChessRogueCurrentInfo   `protobuf:"bytes,7,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *ChessRogueEnterCellScRsp) Reset() {
	*x = ChessRogueEnterCellScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueEnterCellScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueEnterCellScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueEnterCellScRsp) ProtoMessage() {}

func (x *ChessRogueEnterCellScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueEnterCellScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueEnterCellScRsp.ProtoReflect.Descriptor instead.
func (*ChessRogueEnterCellScRsp) Descriptor() ([]byte, []int) {
	return file_ChessRogueEnterCellScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueEnterCellScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ChessRogueEnterCellScRsp) GetPlayerInfo() *ChessRoguePlayerInfo {
	if x != nil {
		return x.PlayerInfo
	}
	return nil
}

func (x *ChessRogueEnterCellScRsp) GetCellId() uint32 {
	if x != nil {
		return x.CellId
	}
	return 0
}

func (x *ChessRogueEnterCellScRsp) GetRogueCurrentInfo() *ChessRogueQueryGameInfo {
	if x != nil {
		return x.RogueCurrentInfo
	}
	return nil
}

func (x *ChessRogueEnterCellScRsp) GetInfo() *ChessRogueCurrentInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_ChessRogueEnterCellScRsp_proto protoreflect.FileDescriptor

var file_ChessRogueEnterCellScRsp_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x65,
	0x72, 0x43, 0x65, 0x6c, 0x6c, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x50, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x43, 0x68,
	0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x01, 0x0a, 0x18, 0x43, 0x68, 0x65,
	0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x65, 0x72, 0x43, 0x65, 0x6c, 0x6c,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x36, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x65, 0x6c, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63, 0x65, 0x6c, 0x6c, 0x49, 0x64,
	0x12, 0x46, 0x0a, 0x12, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x47, 0x61,
	0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2a, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x69, 0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueEnterCellScRsp_proto_rawDescOnce sync.Once
	file_ChessRogueEnterCellScRsp_proto_rawDescData = file_ChessRogueEnterCellScRsp_proto_rawDesc
)

func file_ChessRogueEnterCellScRsp_proto_rawDescGZIP() []byte {
	file_ChessRogueEnterCellScRsp_proto_rawDescOnce.Do(func() {
		file_ChessRogueEnterCellScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueEnterCellScRsp_proto_rawDescData)
	})
	return file_ChessRogueEnterCellScRsp_proto_rawDescData
}

var file_ChessRogueEnterCellScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueEnterCellScRsp_proto_goTypes = []interface{}{
	(*ChessRogueEnterCellScRsp)(nil), // 0: ChessRogueEnterCellScRsp
	(*ChessRoguePlayerInfo)(nil),     // 1: ChessRoguePlayerInfo
	(*ChessRogueQueryGameInfo)(nil),  // 2: ChessRogueQueryGameInfo
	(*ChessRogueCurrentInfo)(nil),    // 3: ChessRogueCurrentInfo
}
var file_ChessRogueEnterCellScRsp_proto_depIdxs = []int32{
	1, // 0: ChessRogueEnterCellScRsp.player_info:type_name -> ChessRoguePlayerInfo
	2, // 1: ChessRogueEnterCellScRsp.rogue_current_info:type_name -> ChessRogueQueryGameInfo
	3, // 2: ChessRogueEnterCellScRsp.info:type_name -> ChessRogueCurrentInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ChessRogueEnterCellScRsp_proto_init() }
func file_ChessRogueEnterCellScRsp_proto_init() {
	if File_ChessRogueEnterCellScRsp_proto != nil {
		return
	}
	file_ChessRoguePlayerInfo_proto_init()
	file_ChessRogueCurrentInfo_proto_init()
	file_ChessRogueQueryGameInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueEnterCellScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueEnterCellScRsp); i {
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
			RawDescriptor: file_ChessRogueEnterCellScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueEnterCellScRsp_proto_goTypes,
		DependencyIndexes: file_ChessRogueEnterCellScRsp_proto_depIdxs,
		MessageInfos:      file_ChessRogueEnterCellScRsp_proto_msgTypes,
	}.Build()
	File_ChessRogueEnterCellScRsp_proto = out.File
	file_ChessRogueEnterCellScRsp_proto_rawDesc = nil
	file_ChessRogueEnterCellScRsp_proto_goTypes = nil
	file_ChessRogueEnterCellScRsp_proto_depIdxs = nil
}
