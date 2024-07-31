// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SyncRogueRewardInfoScNotify.proto

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

type SyncRogueRewardInfoScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RogueScoreRewardInfo *RogueScoreRewardInfo `protobuf:"bytes,11,opt,name=rogue_score_reward_info,json=rogueScoreRewardInfo,proto3" json:"rogue_score_reward_info,omitempty"`
}

func (x *SyncRogueRewardInfoScNotify) Reset() {
	*x = SyncRogueRewardInfoScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SyncRogueRewardInfoScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRogueRewardInfoScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRogueRewardInfoScNotify) ProtoMessage() {}

func (x *SyncRogueRewardInfoScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SyncRogueRewardInfoScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRogueRewardInfoScNotify.ProtoReflect.Descriptor instead.
func (*SyncRogueRewardInfoScNotify) Descriptor() ([]byte, []int) {
	return file_SyncRogueRewardInfoScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SyncRogueRewardInfoScNotify) GetRogueScoreRewardInfo() *RogueScoreRewardInfo {
	if x != nil {
		return x.RogueScoreRewardInfo
	}
	return nil
}

var File_SyncRogueRewardInfoScNotify_proto protoreflect.FileDescriptor

var file_SyncRogueRewardInfoScNotify_proto_rawDesc = []byte{
	0x0a, 0x21, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6b, 0x0a, 0x1b, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x4c,
	0x0a, 0x17, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x63, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SyncRogueRewardInfoScNotify_proto_rawDescOnce sync.Once
	file_SyncRogueRewardInfoScNotify_proto_rawDescData = file_SyncRogueRewardInfoScNotify_proto_rawDesc
)

func file_SyncRogueRewardInfoScNotify_proto_rawDescGZIP() []byte {
	file_SyncRogueRewardInfoScNotify_proto_rawDescOnce.Do(func() {
		file_SyncRogueRewardInfoScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SyncRogueRewardInfoScNotify_proto_rawDescData)
	})
	return file_SyncRogueRewardInfoScNotify_proto_rawDescData
}

var file_SyncRogueRewardInfoScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SyncRogueRewardInfoScNotify_proto_goTypes = []interface{}{
	(*SyncRogueRewardInfoScNotify)(nil), // 0: SyncRogueRewardInfoScNotify
	(*RogueScoreRewardInfo)(nil),        // 1: RogueScoreRewardInfo
}
var file_SyncRogueRewardInfoScNotify_proto_depIdxs = []int32{
	1, // 0: SyncRogueRewardInfoScNotify.rogue_score_reward_info:type_name -> RogueScoreRewardInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SyncRogueRewardInfoScNotify_proto_init() }
func file_SyncRogueRewardInfoScNotify_proto_init() {
	if File_SyncRogueRewardInfoScNotify_proto != nil {
		return
	}
	file_RogueScoreRewardInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SyncRogueRewardInfoScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRogueRewardInfoScNotify); i {
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
			RawDescriptor: file_SyncRogueRewardInfoScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SyncRogueRewardInfoScNotify_proto_goTypes,
		DependencyIndexes: file_SyncRogueRewardInfoScNotify_proto_depIdxs,
		MessageInfos:      file_SyncRogueRewardInfoScNotify_proto_msgTypes,
	}.Build()
	File_SyncRogueRewardInfoScNotify_proto = out.File
	file_SyncRogueRewardInfoScNotify_proto_rawDesc = nil
	file_SyncRogueRewardInfoScNotify_proto_goTypes = nil
	file_SyncRogueRewardInfoScNotify_proto_depIdxs = nil
}
