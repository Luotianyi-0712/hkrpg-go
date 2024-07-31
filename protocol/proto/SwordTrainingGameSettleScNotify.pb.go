// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SwordTrainingGameSettleScNotify.proto

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

type SwordTrainingGameSettleScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason          JKMJAIADMLI `protobuf:"varint,2,opt,name=reason,proto3,enum=JKMJAIADMLI" json:"reason,omitempty"`
	PCMAAKHAEBC     uint32      `protobuf:"varint,1,opt,name=PCMAAKHAEBC,proto3" json:"PCMAAKHAEBC,omitempty"`
	GameStoryLineId uint32      `protobuf:"varint,7,opt,name=game_story_line_id,json=gameStoryLineId,proto3" json:"game_story_line_id,omitempty"`
	DKFLLECGCCK     []uint32    `protobuf:"varint,11,rep,packed,name=DKFLLECGCCK,proto3" json:"DKFLLECGCCK,omitempty"`
	CKEDFNKALOI     uint32      `protobuf:"varint,15,opt,name=CKEDFNKALOI,proto3" json:"CKEDFNKALOI,omitempty"`
	CIDICFBPHAF     []uint32    `protobuf:"varint,3,rep,packed,name=CIDICFBPHAF,proto3" json:"CIDICFBPHAF,omitempty"`
	OOJPFKLPHFH     uint32      `protobuf:"varint,6,opt,name=OOJPFKLPHFH,proto3" json:"OOJPFKLPHFH,omitempty"`
	Reward          *ItemList   `protobuf:"bytes,12,opt,name=reward,proto3" json:"reward,omitempty"`
}

func (x *SwordTrainingGameSettleScNotify) Reset() {
	*x = SwordTrainingGameSettleScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SwordTrainingGameSettleScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwordTrainingGameSettleScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwordTrainingGameSettleScNotify) ProtoMessage() {}

func (x *SwordTrainingGameSettleScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SwordTrainingGameSettleScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwordTrainingGameSettleScNotify.ProtoReflect.Descriptor instead.
func (*SwordTrainingGameSettleScNotify) Descriptor() ([]byte, []int) {
	return file_SwordTrainingGameSettleScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SwordTrainingGameSettleScNotify) GetReason() JKMJAIADMLI {
	if x != nil {
		return x.Reason
	}
	return JKMJAIADMLI_SWORD_TRAINING_GAME_SETTLE_NONE
}

func (x *SwordTrainingGameSettleScNotify) GetPCMAAKHAEBC() uint32 {
	if x != nil {
		return x.PCMAAKHAEBC
	}
	return 0
}

func (x *SwordTrainingGameSettleScNotify) GetGameStoryLineId() uint32 {
	if x != nil {
		return x.GameStoryLineId
	}
	return 0
}

func (x *SwordTrainingGameSettleScNotify) GetDKFLLECGCCK() []uint32 {
	if x != nil {
		return x.DKFLLECGCCK
	}
	return nil
}

func (x *SwordTrainingGameSettleScNotify) GetCKEDFNKALOI() uint32 {
	if x != nil {
		return x.CKEDFNKALOI
	}
	return 0
}

func (x *SwordTrainingGameSettleScNotify) GetCIDICFBPHAF() []uint32 {
	if x != nil {
		return x.CIDICFBPHAF
	}
	return nil
}

func (x *SwordTrainingGameSettleScNotify) GetOOJPFKLPHFH() uint32 {
	if x != nil {
		return x.OOJPFKLPHFH
	}
	return 0
}

func (x *SwordTrainingGameSettleScNotify) GetReward() *ItemList {
	if x != nil {
		return x.Reward
	}
	return nil
}

var File_SwordTrainingGameSettleScNotify_proto protoreflect.FileDescriptor

var file_SwordTrainingGameSettleScNotify_proto_rawDesc = []byte{
	0x0a, 0x25, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x4b, 0x4d, 0x4a, 0x41, 0x49, 0x41,
	0x44, 0x4d, 0x4c, 0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x02, 0x0a, 0x1f, 0x53,
	0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x47, 0x61, 0x6d, 0x65,
	0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x24,
	0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c,
	0x2e, 0x4a, 0x4b, 0x4d, 0x4a, 0x41, 0x49, 0x41, 0x44, 0x4d, 0x4c, 0x49, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x43, 0x4d, 0x41, 0x41, 0x4b, 0x48, 0x41,
	0x45, 0x42, 0x43, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x43, 0x4d, 0x41, 0x41,
	0x4b, 0x48, 0x41, 0x45, 0x42, 0x43, 0x12, 0x2b, 0x0a, 0x12, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0f, 0x67, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x6e,
	0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4b, 0x46, 0x4c, 0x4c, 0x45, 0x43, 0x47, 0x43,
	0x43, 0x4b, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x4b, 0x46, 0x4c, 0x4c, 0x45,
	0x43, 0x47, 0x43, 0x43, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x4b, 0x45, 0x44, 0x46, 0x4e, 0x4b,
	0x41, 0x4c, 0x4f, 0x49, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x4b, 0x45, 0x44,
	0x46, 0x4e, 0x4b, 0x41, 0x4c, 0x4f, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x49, 0x44, 0x49, 0x43,
	0x46, 0x42, 0x50, 0x48, 0x41, 0x46, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x49,
	0x44, 0x49, 0x43, 0x46, 0x42, 0x50, 0x48, 0x41, 0x46, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x4f, 0x4a,
	0x50, 0x46, 0x4b, 0x4c, 0x50, 0x48, 0x46, 0x48, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4f, 0x4f, 0x4a, 0x50, 0x46, 0x4b, 0x4c, 0x50, 0x48, 0x46, 0x48, 0x12, 0x21, 0x0a, 0x06, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x42, 0x28,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SwordTrainingGameSettleScNotify_proto_rawDescOnce sync.Once
	file_SwordTrainingGameSettleScNotify_proto_rawDescData = file_SwordTrainingGameSettleScNotify_proto_rawDesc
)

func file_SwordTrainingGameSettleScNotify_proto_rawDescGZIP() []byte {
	file_SwordTrainingGameSettleScNotify_proto_rawDescOnce.Do(func() {
		file_SwordTrainingGameSettleScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SwordTrainingGameSettleScNotify_proto_rawDescData)
	})
	return file_SwordTrainingGameSettleScNotify_proto_rawDescData
}

var file_SwordTrainingGameSettleScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SwordTrainingGameSettleScNotify_proto_goTypes = []interface{}{
	(*SwordTrainingGameSettleScNotify)(nil), // 0: SwordTrainingGameSettleScNotify
	(JKMJAIADMLI)(0),                        // 1: JKMJAIADMLI
	(*ItemList)(nil),                        // 2: ItemList
}
var file_SwordTrainingGameSettleScNotify_proto_depIdxs = []int32{
	1, // 0: SwordTrainingGameSettleScNotify.reason:type_name -> JKMJAIADMLI
	2, // 1: SwordTrainingGameSettleScNotify.reward:type_name -> ItemList
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_SwordTrainingGameSettleScNotify_proto_init() }
func file_SwordTrainingGameSettleScNotify_proto_init() {
	if File_SwordTrainingGameSettleScNotify_proto != nil {
		return
	}
	file_JKMJAIADMLI_proto_init()
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SwordTrainingGameSettleScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwordTrainingGameSettleScNotify); i {
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
			RawDescriptor: file_SwordTrainingGameSettleScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SwordTrainingGameSettleScNotify_proto_goTypes,
		DependencyIndexes: file_SwordTrainingGameSettleScNotify_proto_depIdxs,
		MessageInfos:      file_SwordTrainingGameSettleScNotify_proto_msgTypes,
	}.Build()
	File_SwordTrainingGameSettleScNotify_proto = out.File
	file_SwordTrainingGameSettleScNotify_proto_rawDesc = nil
	file_SwordTrainingGameSettleScNotify_proto_goTypes = nil
	file_SwordTrainingGameSettleScNotify_proto_depIdxs = nil
}
