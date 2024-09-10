// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueEndlessAvatar.proto

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

type RogueEndlessAvatar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarId   uint32     `protobuf:"varint,15,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	AvatarType AvatarType `protobuf:"varint,1,opt,name=avatar_type,json=avatarType,proto3,enum=AvatarType" json:"avatar_type,omitempty"`
}

func (x *RogueEndlessAvatar) Reset() {
	*x = RogueEndlessAvatar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueEndlessAvatar_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueEndlessAvatar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueEndlessAvatar) ProtoMessage() {}

func (x *RogueEndlessAvatar) ProtoReflect() protoreflect.Message {
	mi := &file_RogueEndlessAvatar_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueEndlessAvatar.ProtoReflect.Descriptor instead.
func (*RogueEndlessAvatar) Descriptor() ([]byte, []int) {
	return file_RogueEndlessAvatar_proto_rawDescGZIP(), []int{0}
}

func (x *RogueEndlessAvatar) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *RogueEndlessAvatar) GetAvatarType() AvatarType {
	if x != nil {
		return x.AvatarType
	}
	return AvatarType_AVATAR_TYPE_NONE
}

var File_RogueEndlessAvatar_proto protoreflect.FileDescriptor

var file_RogueEndlessAvatar_proto_rawDesc = []byte{
	0x0a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x12,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x45, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x73, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12,
	0x2c, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42, 0x2e, 0x5a,
	0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueEndlessAvatar_proto_rawDescOnce sync.Once
	file_RogueEndlessAvatar_proto_rawDescData = file_RogueEndlessAvatar_proto_rawDesc
)

func file_RogueEndlessAvatar_proto_rawDescGZIP() []byte {
	file_RogueEndlessAvatar_proto_rawDescOnce.Do(func() {
		file_RogueEndlessAvatar_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueEndlessAvatar_proto_rawDescData)
	})
	return file_RogueEndlessAvatar_proto_rawDescData
}

var file_RogueEndlessAvatar_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueEndlessAvatar_proto_goTypes = []interface{}{
	(*RogueEndlessAvatar)(nil), // 0: RogueEndlessAvatar
	(AvatarType)(0),            // 1: AvatarType
}
var file_RogueEndlessAvatar_proto_depIdxs = []int32{
	1, // 0: RogueEndlessAvatar.avatar_type:type_name -> AvatarType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_RogueEndlessAvatar_proto_init() }
func file_RogueEndlessAvatar_proto_init() {
	if File_RogueEndlessAvatar_proto != nil {
		return
	}
	file_AvatarType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueEndlessAvatar_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueEndlessAvatar); i {
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
			RawDescriptor: file_RogueEndlessAvatar_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueEndlessAvatar_proto_goTypes,
		DependencyIndexes: file_RogueEndlessAvatar_proto_depIdxs,
		MessageInfos:      file_RogueEndlessAvatar_proto_msgTypes,
	}.Build()
	File_RogueEndlessAvatar_proto = out.File
	file_RogueEndlessAvatar_proto_rawDesc = nil
	file_RogueEndlessAvatar_proto_goTypes = nil
	file_RogueEndlessAvatar_proto_depIdxs = nil
}
