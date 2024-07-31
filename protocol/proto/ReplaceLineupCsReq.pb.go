// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ReplaceLineupCsReq.proto

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

type ReplaceLineupCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameStoryLineId uint32            `protobuf:"varint,12,opt,name=game_story_line_id,json=gameStoryLineId,proto3" json:"game_story_line_id,omitempty"`
	PlaneId         uint32            `protobuf:"varint,3,opt,name=plane_id,json=planeId,proto3" json:"plane_id,omitempty"`
	LeaderSlot      uint32            `protobuf:"varint,7,opt,name=leader_slot,json=leaderSlot,proto3" json:"leader_slot,omitempty"`
	IsVirtual       bool              `protobuf:"varint,4,opt,name=is_virtual,json=isVirtual,proto3" json:"is_virtual,omitempty"`
	LineupSlotList  []*LineupSlotData `protobuf:"bytes,14,rep,name=lineup_slot_list,json=lineupSlotList,proto3" json:"lineup_slot_list,omitempty"`
	ExtraLineupType ExtraLineupType   `protobuf:"varint,10,opt,name=extra_lineup_type,json=extraLineupType,proto3,enum=ExtraLineupType" json:"extra_lineup_type,omitempty"`
	Index           uint32            `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *ReplaceLineupCsReq) Reset() {
	*x = ReplaceLineupCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ReplaceLineupCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplaceLineupCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplaceLineupCsReq) ProtoMessage() {}

func (x *ReplaceLineupCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_ReplaceLineupCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplaceLineupCsReq.ProtoReflect.Descriptor instead.
func (*ReplaceLineupCsReq) Descriptor() ([]byte, []int) {
	return file_ReplaceLineupCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *ReplaceLineupCsReq) GetGameStoryLineId() uint32 {
	if x != nil {
		return x.GameStoryLineId
	}
	return 0
}

func (x *ReplaceLineupCsReq) GetPlaneId() uint32 {
	if x != nil {
		return x.PlaneId
	}
	return 0
}

func (x *ReplaceLineupCsReq) GetLeaderSlot() uint32 {
	if x != nil {
		return x.LeaderSlot
	}
	return 0
}

func (x *ReplaceLineupCsReq) GetIsVirtual() bool {
	if x != nil {
		return x.IsVirtual
	}
	return false
}

func (x *ReplaceLineupCsReq) GetLineupSlotList() []*LineupSlotData {
	if x != nil {
		return x.LineupSlotList
	}
	return nil
}

func (x *ReplaceLineupCsReq) GetExtraLineupType() ExtraLineupType {
	if x != nil {
		return x.ExtraLineupType
	}
	return ExtraLineupType_LINEUP_NONE
}

func (x *ReplaceLineupCsReq) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

var File_ReplaceLineupCsReq_proto protoreflect.FileDescriptor

var file_ReplaceLineupCsReq_proto_rawDesc = []byte{
	0x0a, 0x18, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x43,
	0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x4c, 0x69, 0x6e, 0x65,
	0x75, 0x70, 0x53, 0x6c, 0x6f, 0x74, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x15, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x02, 0x0a, 0x12, 0x52, 0x65, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x2b,
	0x0a, 0x12, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x67, 0x61, 0x6d, 0x65,
	0x53, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70,
	0x6c, 0x61, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x76, 0x69,
	0x72, 0x74, 0x75, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x56,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x12, 0x39, 0x0a, 0x10, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70,
	0x5f, 0x73, 0x6c, 0x6f, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x53, 0x6c, 0x6f, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x0e, 0x6c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x53, 0x6c, 0x6f, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x3c, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x75,
	0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ReplaceLineupCsReq_proto_rawDescOnce sync.Once
	file_ReplaceLineupCsReq_proto_rawDescData = file_ReplaceLineupCsReq_proto_rawDesc
)

func file_ReplaceLineupCsReq_proto_rawDescGZIP() []byte {
	file_ReplaceLineupCsReq_proto_rawDescOnce.Do(func() {
		file_ReplaceLineupCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_ReplaceLineupCsReq_proto_rawDescData)
	})
	return file_ReplaceLineupCsReq_proto_rawDescData
}

var file_ReplaceLineupCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ReplaceLineupCsReq_proto_goTypes = []interface{}{
	(*ReplaceLineupCsReq)(nil), // 0: ReplaceLineupCsReq
	(*LineupSlotData)(nil),     // 1: LineupSlotData
	(ExtraLineupType)(0),       // 2: ExtraLineupType
}
var file_ReplaceLineupCsReq_proto_depIdxs = []int32{
	1, // 0: ReplaceLineupCsReq.lineup_slot_list:type_name -> LineupSlotData
	2, // 1: ReplaceLineupCsReq.extra_lineup_type:type_name -> ExtraLineupType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ReplaceLineupCsReq_proto_init() }
func file_ReplaceLineupCsReq_proto_init() {
	if File_ReplaceLineupCsReq_proto != nil {
		return
	}
	file_LineupSlotData_proto_init()
	file_ExtraLineupType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ReplaceLineupCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplaceLineupCsReq); i {
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
			RawDescriptor: file_ReplaceLineupCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ReplaceLineupCsReq_proto_goTypes,
		DependencyIndexes: file_ReplaceLineupCsReq_proto_depIdxs,
		MessageInfos:      file_ReplaceLineupCsReq_proto_msgTypes,
	}.Build()
	File_ReplaceLineupCsReq_proto = out.File
	file_ReplaceLineupCsReq_proto_rawDesc = nil
	file_ReplaceLineupCsReq_proto_goTypes = nil
	file_ReplaceLineupCsReq_proto_depIdxs = nil
}
