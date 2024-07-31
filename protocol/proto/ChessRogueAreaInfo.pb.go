// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueAreaInfo.proto

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

type ChessRogueAreaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LayerStatus             ChessRogueBoardCellStatus    `protobuf:"varint,13,opt,name=layer_status,json=layerStatus,proto3,enum=ChessRogueBoardCellStatus" json:"layer_status,omitempty"`
	OJNCMJDAABJ             *JDIPIHPMEKN                 `protobuf:"bytes,1,opt,name=OJNCMJDAABJ,proto3" json:"OJNCMJDAABJ,omitempty"`
	AllowedSelectCellIdList []uint32                     `protobuf:"varint,10,rep,packed,name=allowed_select_cell_id_list,json=allowedSelectCellIdList,proto3" json:"allowed_select_cell_id_list,omitempty"`
	Cell                    *CellInfo                    `protobuf:"bytes,2,opt,name=cell,proto3" json:"cell,omitempty"`
	HistoryCell             []*ChessRogueHistoryCellInfo `protobuf:"bytes,7,rep,name=history_cell,json=historyCell,proto3" json:"history_cell,omitempty"`
	CurId                   uint32                       `protobuf:"varint,12,opt,name=cur_id,json=curId,proto3" json:"cur_id,omitempty"`
	CurBoardId              uint32                       `protobuf:"varint,15,opt,name=cur_board_id,json=curBoardId,proto3" json:"cur_board_id,omitempty"`
}

func (x *ChessRogueAreaInfo) Reset() {
	*x = ChessRogueAreaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueAreaInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueAreaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueAreaInfo) ProtoMessage() {}

func (x *ChessRogueAreaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueAreaInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueAreaInfo.ProtoReflect.Descriptor instead.
func (*ChessRogueAreaInfo) Descriptor() ([]byte, []int) {
	return file_ChessRogueAreaInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueAreaInfo) GetLayerStatus() ChessRogueBoardCellStatus {
	if x != nil {
		return x.LayerStatus
	}
	return ChessRogueBoardCellStatus_IDLE
}

func (x *ChessRogueAreaInfo) GetOJNCMJDAABJ() *JDIPIHPMEKN {
	if x != nil {
		return x.OJNCMJDAABJ
	}
	return nil
}

func (x *ChessRogueAreaInfo) GetAllowedSelectCellIdList() []uint32 {
	if x != nil {
		return x.AllowedSelectCellIdList
	}
	return nil
}

func (x *ChessRogueAreaInfo) GetCell() *CellInfo {
	if x != nil {
		return x.Cell
	}
	return nil
}

func (x *ChessRogueAreaInfo) GetHistoryCell() []*ChessRogueHistoryCellInfo {
	if x != nil {
		return x.HistoryCell
	}
	return nil
}

func (x *ChessRogueAreaInfo) GetCurId() uint32 {
	if x != nil {
		return x.CurId
	}
	return 0
}

func (x *ChessRogueAreaInfo) GetCurBoardId() uint32 {
	if x != nil {
		return x.CurBoardId
	}
	return 0
}

var File_ChessRogueAreaInfo_proto protoreflect.FileDescriptor

var file_ChessRogueAreaInfo_proto_rawDesc = []byte{
	0x0a, 0x18, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x72, 0x65, 0x61,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x44, 0x49, 0x50,
	0x49, 0x48, 0x50, 0x4d, 0x45, 0x4b, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x43, 0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x43,
	0x65, 0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x0e, 0x43, 0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd8, 0x02, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x41, 0x72,
	0x65, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x43, 0x65,
	0x6c, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0b, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x0b, 0x4f, 0x4a, 0x4e, 0x43, 0x4d, 0x4a, 0x44,
	0x41, 0x41, 0x42, 0x4a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a, 0x44, 0x49,
	0x50, 0x49, 0x48, 0x50, 0x4d, 0x45, 0x4b, 0x4e, 0x52, 0x0b, 0x4f, 0x4a, 0x4e, 0x43, 0x4d, 0x4a,
	0x44, 0x41, 0x41, 0x42, 0x4a, 0x12, 0x3c, 0x0a, 0x1b, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64,
	0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x5f, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x17, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x65, 0x64, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x65, 0x6c, 0x6c, 0x49, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x63, 0x65, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x43, 0x65, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x63, 0x65,
	0x6c, 0x6c, 0x12, 0x3d, 0x0a, 0x0c, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x63, 0x65,
	0x6c, 0x6c, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x65, 0x6c, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x43, 0x65, 0x6c,
	0x6c, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x63, 0x75, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x5f,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x63, 0x75, 0x72, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x49, 0x64, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueAreaInfo_proto_rawDescOnce sync.Once
	file_ChessRogueAreaInfo_proto_rawDescData = file_ChessRogueAreaInfo_proto_rawDesc
)

func file_ChessRogueAreaInfo_proto_rawDescGZIP() []byte {
	file_ChessRogueAreaInfo_proto_rawDescOnce.Do(func() {
		file_ChessRogueAreaInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueAreaInfo_proto_rawDescData)
	})
	return file_ChessRogueAreaInfo_proto_rawDescData
}

var file_ChessRogueAreaInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueAreaInfo_proto_goTypes = []interface{}{
	(*ChessRogueAreaInfo)(nil),        // 0: ChessRogueAreaInfo
	(ChessRogueBoardCellStatus)(0),    // 1: ChessRogueBoardCellStatus
	(*JDIPIHPMEKN)(nil),               // 2: JDIPIHPMEKN
	(*CellInfo)(nil),                  // 3: CellInfo
	(*ChessRogueHistoryCellInfo)(nil), // 4: ChessRogueHistoryCellInfo
}
var file_ChessRogueAreaInfo_proto_depIdxs = []int32{
	1, // 0: ChessRogueAreaInfo.layer_status:type_name -> ChessRogueBoardCellStatus
	2, // 1: ChessRogueAreaInfo.OJNCMJDAABJ:type_name -> JDIPIHPMEKN
	3, // 2: ChessRogueAreaInfo.cell:type_name -> CellInfo
	4, // 3: ChessRogueAreaInfo.history_cell:type_name -> ChessRogueHistoryCellInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ChessRogueAreaInfo_proto_init() }
func file_ChessRogueAreaInfo_proto_init() {
	if File_ChessRogueAreaInfo_proto != nil {
		return
	}
	file_JDIPIHPMEKN_proto_init()
	file_ChessRogueHistoryCellInfo_proto_init()
	file_ChessRogueBoardCellStatus_proto_init()
	file_CellInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueAreaInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueAreaInfo); i {
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
			RawDescriptor: file_ChessRogueAreaInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueAreaInfo_proto_goTypes,
		DependencyIndexes: file_ChessRogueAreaInfo_proto_depIdxs,
		MessageInfos:      file_ChessRogueAreaInfo_proto_msgTypes,
	}.Build()
	File_ChessRogueAreaInfo_proto = out.File
	file_ChessRogueAreaInfo_proto_rawDesc = nil
	file_ChessRogueAreaInfo_proto_goTypes = nil
	file_ChessRogueAreaInfo_proto_depIdxs = nil
}
