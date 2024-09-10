// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TakeRogueEndlessActivityPointRewardScRsp.proto

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

type TakeRogueEndlessActivityPointRewardScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TakenRewardLevelList []uint32  `protobuf:"varint,12,rep,packed,name=taken_reward_level_list,json=takenRewardLevelList,proto3" json:"taken_reward_level_list,omitempty"`
	Reward               *ItemList `protobuf:"bytes,7,opt,name=reward,proto3" json:"reward,omitempty"`
	Level                uint32    `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	HFEDJFGPNOH          bool      `protobuf:"varint,11,opt,name=HFEDJFGPNOH,proto3" json:"HFEDJFGPNOH,omitempty"`
	IPOJGPBKKPD          uint32    `protobuf:"varint,8,opt,name=IPOJGPBKKPD,proto3" json:"IPOJGPBKKPD,omitempty"`
	Retcode              uint32    `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *TakeRogueEndlessActivityPointRewardScRsp) Reset() {
	*x = TakeRogueEndlessActivityPointRewardScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TakeRogueEndlessActivityPointRewardScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeRogueEndlessActivityPointRewardScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeRogueEndlessActivityPointRewardScRsp) ProtoMessage() {}

func (x *TakeRogueEndlessActivityPointRewardScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_TakeRogueEndlessActivityPointRewardScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeRogueEndlessActivityPointRewardScRsp.ProtoReflect.Descriptor instead.
func (*TakeRogueEndlessActivityPointRewardScRsp) Descriptor() ([]byte, []int) {
	return file_TakeRogueEndlessActivityPointRewardScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *TakeRogueEndlessActivityPointRewardScRsp) GetTakenRewardLevelList() []uint32 {
	if x != nil {
		return x.TakenRewardLevelList
	}
	return nil
}

func (x *TakeRogueEndlessActivityPointRewardScRsp) GetReward() *ItemList {
	if x != nil {
		return x.Reward
	}
	return nil
}

func (x *TakeRogueEndlessActivityPointRewardScRsp) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *TakeRogueEndlessActivityPointRewardScRsp) GetHFEDJFGPNOH() bool {
	if x != nil {
		return x.HFEDJFGPNOH
	}
	return false
}

func (x *TakeRogueEndlessActivityPointRewardScRsp) GetIPOJGPBKKPD() uint32 {
	if x != nil {
		return x.IPOJGPBKKPD
	}
	return 0
}

func (x *TakeRogueEndlessActivityPointRewardScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_TakeRogueEndlessActivityPointRewardScRsp_proto protoreflect.FileDescriptor

var file_TakeRogueEndlessActivityPointRewardScRsp_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x54, 0x61, 0x6b, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x6e, 0x64, 0x6c, 0x65,
	0x73, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf8, 0x01, 0x0a, 0x28, 0x54, 0x61, 0x6b, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x6e,
	0x64, 0x6c, 0x65, 0x73, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x35, 0x0a,
	0x17, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x14,
	0x74, 0x61, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a,
	0x0b, 0x48, 0x46, 0x45, 0x44, 0x4a, 0x46, 0x47, 0x50, 0x4e, 0x4f, 0x48, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x48, 0x46, 0x45, 0x44, 0x4a, 0x46, 0x47, 0x50, 0x4e, 0x4f, 0x48, 0x12,
	0x20, 0x0a, 0x0b, 0x49, 0x50, 0x4f, 0x4a, 0x47, 0x50, 0x42, 0x4b, 0x4b, 0x50, 0x44, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x50, 0x4f, 0x4a, 0x47, 0x50, 0x42, 0x4b, 0x4b, 0x50,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e,
	0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_TakeRogueEndlessActivityPointRewardScRsp_proto_rawDescOnce sync.Once
	file_TakeRogueEndlessActivityPointRewardScRsp_proto_rawDescData = file_TakeRogueEndlessActivityPointRewardScRsp_proto_rawDesc
)

func file_TakeRogueEndlessActivityPointRewardScRsp_proto_rawDescGZIP() []byte {
	file_TakeRogueEndlessActivityPointRewardScRsp_proto_rawDescOnce.Do(func() {
		file_TakeRogueEndlessActivityPointRewardScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_TakeRogueEndlessActivityPointRewardScRsp_proto_rawDescData)
	})
	return file_TakeRogueEndlessActivityPointRewardScRsp_proto_rawDescData
}

var file_TakeRogueEndlessActivityPointRewardScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TakeRogueEndlessActivityPointRewardScRsp_proto_goTypes = []interface{}{
	(*TakeRogueEndlessActivityPointRewardScRsp)(nil), // 0: TakeRogueEndlessActivityPointRewardScRsp
	(*ItemList)(nil), // 1: ItemList
}
var file_TakeRogueEndlessActivityPointRewardScRsp_proto_depIdxs = []int32{
	1, // 0: TakeRogueEndlessActivityPointRewardScRsp.reward:type_name -> ItemList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_TakeRogueEndlessActivityPointRewardScRsp_proto_init() }
func file_TakeRogueEndlessActivityPointRewardScRsp_proto_init() {
	if File_TakeRogueEndlessActivityPointRewardScRsp_proto != nil {
		return
	}
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TakeRogueEndlessActivityPointRewardScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeRogueEndlessActivityPointRewardScRsp); i {
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
			RawDescriptor: file_TakeRogueEndlessActivityPointRewardScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TakeRogueEndlessActivityPointRewardScRsp_proto_goTypes,
		DependencyIndexes: file_TakeRogueEndlessActivityPointRewardScRsp_proto_depIdxs,
		MessageInfos:      file_TakeRogueEndlessActivityPointRewardScRsp_proto_msgTypes,
	}.Build()
	File_TakeRogueEndlessActivityPointRewardScRsp_proto = out.File
	file_TakeRogueEndlessActivityPointRewardScRsp_proto_rawDesc = nil
	file_TakeRogueEndlessActivityPointRewardScRsp_proto_goTypes = nil
	file_TakeRogueEndlessActivityPointRewardScRsp_proto_depIdxs = nil
}
