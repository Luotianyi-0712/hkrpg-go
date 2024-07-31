// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FinishEmotionDialoguePerformanceScRsp.proto

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

type FinishEmotionDialoguePerformanceScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode     uint32    `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
	JNEDEELEEKO uint32    `protobuf:"varint,3,opt,name=JNEDEELEEKO,proto3" json:"JNEDEELEEKO,omitempty"`
	RewardList  *ItemList `protobuf:"bytes,7,opt,name=reward_list,json=rewardList,proto3" json:"reward_list,omitempty"`
	ONKKHGBNJOI uint32    `protobuf:"varint,4,opt,name=ONKKHGBNJOI,proto3" json:"ONKKHGBNJOI,omitempty"`
}

func (x *FinishEmotionDialoguePerformanceScRsp) Reset() {
	*x = FinishEmotionDialoguePerformanceScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FinishEmotionDialoguePerformanceScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishEmotionDialoguePerformanceScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishEmotionDialoguePerformanceScRsp) ProtoMessage() {}

func (x *FinishEmotionDialoguePerformanceScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_FinishEmotionDialoguePerformanceScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishEmotionDialoguePerformanceScRsp.ProtoReflect.Descriptor instead.
func (*FinishEmotionDialoguePerformanceScRsp) Descriptor() ([]byte, []int) {
	return file_FinishEmotionDialoguePerformanceScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *FinishEmotionDialoguePerformanceScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *FinishEmotionDialoguePerformanceScRsp) GetJNEDEELEEKO() uint32 {
	if x != nil {
		return x.JNEDEELEEKO
	}
	return 0
}

func (x *FinishEmotionDialoguePerformanceScRsp) GetRewardList() *ItemList {
	if x != nil {
		return x.RewardList
	}
	return nil
}

func (x *FinishEmotionDialoguePerformanceScRsp) GetONKKHGBNJOI() uint32 {
	if x != nil {
		return x.ONKKHGBNJOI
	}
	return 0
}

var File_FinishEmotionDialoguePerformanceScRsp_proto protoreflect.FileDescriptor

var file_FinishEmotionDialoguePerformanceScRsp_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x45, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49,
	0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01,
	0x0a, 0x25, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x45, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x69, 0x61, 0x6c, 0x6f, 0x67, 0x75, 0x65, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x4e, 0x45, 0x44, 0x45, 0x45, 0x4c, 0x45, 0x45, 0x4b, 0x4f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x4e, 0x45, 0x44, 0x45, 0x45, 0x4c, 0x45,
	0x45, 0x4b, 0x4f, 0x12, 0x2a, 0x0a, 0x0b, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x0a, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x4f, 0x4e, 0x4b, 0x4b, 0x48, 0x47, 0x42, 0x4e, 0x4a, 0x4f, 0x49, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x4e, 0x4b, 0x4b, 0x48, 0x47, 0x42, 0x4e, 0x4a, 0x4f,
	0x49, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_FinishEmotionDialoguePerformanceScRsp_proto_rawDescOnce sync.Once
	file_FinishEmotionDialoguePerformanceScRsp_proto_rawDescData = file_FinishEmotionDialoguePerformanceScRsp_proto_rawDesc
)

func file_FinishEmotionDialoguePerformanceScRsp_proto_rawDescGZIP() []byte {
	file_FinishEmotionDialoguePerformanceScRsp_proto_rawDescOnce.Do(func() {
		file_FinishEmotionDialoguePerformanceScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_FinishEmotionDialoguePerformanceScRsp_proto_rawDescData)
	})
	return file_FinishEmotionDialoguePerformanceScRsp_proto_rawDescData
}

var file_FinishEmotionDialoguePerformanceScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FinishEmotionDialoguePerformanceScRsp_proto_goTypes = []interface{}{
	(*FinishEmotionDialoguePerformanceScRsp)(nil), // 0: FinishEmotionDialoguePerformanceScRsp
	(*ItemList)(nil), // 1: ItemList
}
var file_FinishEmotionDialoguePerformanceScRsp_proto_depIdxs = []int32{
	1, // 0: FinishEmotionDialoguePerformanceScRsp.reward_list:type_name -> ItemList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FinishEmotionDialoguePerformanceScRsp_proto_init() }
func file_FinishEmotionDialoguePerformanceScRsp_proto_init() {
	if File_FinishEmotionDialoguePerformanceScRsp_proto != nil {
		return
	}
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FinishEmotionDialoguePerformanceScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishEmotionDialoguePerformanceScRsp); i {
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
			RawDescriptor: file_FinishEmotionDialoguePerformanceScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FinishEmotionDialoguePerformanceScRsp_proto_goTypes,
		DependencyIndexes: file_FinishEmotionDialoguePerformanceScRsp_proto_depIdxs,
		MessageInfos:      file_FinishEmotionDialoguePerformanceScRsp_proto_msgTypes,
	}.Build()
	File_FinishEmotionDialoguePerformanceScRsp_proto = out.File
	file_FinishEmotionDialoguePerformanceScRsp_proto_rawDesc = nil
	file_FinishEmotionDialoguePerformanceScRsp_proto_goTypes = nil
	file_FinishEmotionDialoguePerformanceScRsp_proto_depIdxs = nil
}
