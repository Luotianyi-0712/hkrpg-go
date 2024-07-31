// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetRogueEndlessActivityDataScRsp.proto

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

type GetRogueEndlessActivityDataScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IPOJGPBKKPD          uint32                  `protobuf:"varint,13,opt,name=IPOJGPBKKPD,proto3" json:"IPOJGPBKKPD,omitempty"`
	LJNCNGJCLCM          uint32                  `protobuf:"varint,8,opt,name=LJNCNGJCLCM,proto3" json:"LJNCNGJCLCM,omitempty"`
	Retcode              uint32                  `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Data                 []*RogueEndlessAreaData `protobuf:"bytes,14,rep,name=data,proto3" json:"data,omitempty"`
	TakenRewardLevelList []uint32                `protobuf:"varint,12,rep,packed,name=taken_reward_level_list,json=takenRewardLevelList,proto3" json:"taken_reward_level_list,omitempty"`
}

func (x *GetRogueEndlessActivityDataScRsp) Reset() {
	*x = GetRogueEndlessActivityDataScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetRogueEndlessActivityDataScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRogueEndlessActivityDataScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRogueEndlessActivityDataScRsp) ProtoMessage() {}

func (x *GetRogueEndlessActivityDataScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetRogueEndlessActivityDataScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRogueEndlessActivityDataScRsp.ProtoReflect.Descriptor instead.
func (*GetRogueEndlessActivityDataScRsp) Descriptor() ([]byte, []int) {
	return file_GetRogueEndlessActivityDataScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetRogueEndlessActivityDataScRsp) GetIPOJGPBKKPD() uint32 {
	if x != nil {
		return x.IPOJGPBKKPD
	}
	return 0
}

func (x *GetRogueEndlessActivityDataScRsp) GetLJNCNGJCLCM() uint32 {
	if x != nil {
		return x.LJNCNGJCLCM
	}
	return 0
}

func (x *GetRogueEndlessActivityDataScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetRogueEndlessActivityDataScRsp) GetData() []*RogueEndlessAreaData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetRogueEndlessActivityDataScRsp) GetTakenRewardLevelList() []uint32 {
	if x != nil {
		return x.TakenRewardLevelList
	}
	return nil
}

var File_GetRogueEndlessActivityDataScRsp_proto protoreflect.FileDescriptor

var file_GetRogueEndlessActivityDataScRsp_proto_rawDesc = []byte{
	0x0a, 0x26, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x6e, 0x64, 0x6c, 0x65, 0x73,
	0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x52,
	0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x45,
	0x6e, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x41, 0x72, 0x65, 0x61, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe2, 0x01, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x45, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x50, 0x4f,
	0x4a, 0x47, 0x50, 0x42, 0x4b, 0x4b, 0x50, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x49, 0x50, 0x4f, 0x4a, 0x47, 0x50, 0x42, 0x4b, 0x4b, 0x50, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4c,
	0x4a, 0x4e, 0x43, 0x4e, 0x47, 0x4a, 0x43, 0x4c, 0x43, 0x4d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4c, 0x4a, 0x4e, 0x43, 0x4e, 0x47, 0x4a, 0x43, 0x4c, 0x43, 0x4d, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x6e, 0x64,
	0x6c, 0x65, 0x73, 0x73, 0x41, 0x72, 0x65, 0x61, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x35, 0x0a, 0x17, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x14, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetRogueEndlessActivityDataScRsp_proto_rawDescOnce sync.Once
	file_GetRogueEndlessActivityDataScRsp_proto_rawDescData = file_GetRogueEndlessActivityDataScRsp_proto_rawDesc
)

func file_GetRogueEndlessActivityDataScRsp_proto_rawDescGZIP() []byte {
	file_GetRogueEndlessActivityDataScRsp_proto_rawDescOnce.Do(func() {
		file_GetRogueEndlessActivityDataScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetRogueEndlessActivityDataScRsp_proto_rawDescData)
	})
	return file_GetRogueEndlessActivityDataScRsp_proto_rawDescData
}

var file_GetRogueEndlessActivityDataScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetRogueEndlessActivityDataScRsp_proto_goTypes = []interface{}{
	(*GetRogueEndlessActivityDataScRsp)(nil), // 0: GetRogueEndlessActivityDataScRsp
	(*RogueEndlessAreaData)(nil),             // 1: RogueEndlessAreaData
}
var file_GetRogueEndlessActivityDataScRsp_proto_depIdxs = []int32{
	1, // 0: GetRogueEndlessActivityDataScRsp.data:type_name -> RogueEndlessAreaData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetRogueEndlessActivityDataScRsp_proto_init() }
func file_GetRogueEndlessActivityDataScRsp_proto_init() {
	if File_GetRogueEndlessActivityDataScRsp_proto != nil {
		return
	}
	file_RogueEndlessAreaData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetRogueEndlessActivityDataScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRogueEndlessActivityDataScRsp); i {
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
			RawDescriptor: file_GetRogueEndlessActivityDataScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetRogueEndlessActivityDataScRsp_proto_goTypes,
		DependencyIndexes: file_GetRogueEndlessActivityDataScRsp_proto_depIdxs,
		MessageInfos:      file_GetRogueEndlessActivityDataScRsp_proto_msgTypes,
	}.Build()
	File_GetRogueEndlessActivityDataScRsp_proto = out.File
	file_GetRogueEndlessActivityDataScRsp_proto_rawDesc = nil
	file_GetRogueEndlessActivityDataScRsp_proto_goTypes = nil
	file_GetRogueEndlessActivityDataScRsp_proto_depIdxs = nil
}
