// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChallengeAvatarInfo.proto

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

type ChallengeAvatarInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarType AvatarType `protobuf:"varint,15,opt,name=avatar_type,json=avatarType,proto3,enum=AvatarType" json:"avatar_type,omitempty"`
	Level      uint32     `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	Id         uint32     `protobuf:"varint,7,opt,name=id,proto3" json:"id,omitempty"`
	Index      uint32     `protobuf:"varint,5,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *ChallengeAvatarInfo) Reset() {
	*x = ChallengeAvatarInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChallengeAvatarInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeAvatarInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeAvatarInfo) ProtoMessage() {}

func (x *ChallengeAvatarInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChallengeAvatarInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeAvatarInfo.ProtoReflect.Descriptor instead.
func (*ChallengeAvatarInfo) Descriptor() ([]byte, []int) {
	return file_ChallengeAvatarInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChallengeAvatarInfo) GetAvatarType() AvatarType {
	if x != nil {
		return x.AvatarType
	}
	return AvatarType_AVATAR_TYPE_NONE
}

func (x *ChallengeAvatarInfo) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *ChallengeAvatarInfo) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChallengeAvatarInfo) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

var File_ChallengeAvatarInfo_proto protoreflect.FileDescriptor

var file_ChallengeAvatarInfo_proto_rawDesc = []byte{
	0x0a, 0x19, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x41, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f, 0x0a,
	0x13, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2c, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x42, 0x2e,
	0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChallengeAvatarInfo_proto_rawDescOnce sync.Once
	file_ChallengeAvatarInfo_proto_rawDescData = file_ChallengeAvatarInfo_proto_rawDesc
)

func file_ChallengeAvatarInfo_proto_rawDescGZIP() []byte {
	file_ChallengeAvatarInfo_proto_rawDescOnce.Do(func() {
		file_ChallengeAvatarInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChallengeAvatarInfo_proto_rawDescData)
	})
	return file_ChallengeAvatarInfo_proto_rawDescData
}

var file_ChallengeAvatarInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChallengeAvatarInfo_proto_goTypes = []interface{}{
	(*ChallengeAvatarInfo)(nil), // 0: ChallengeAvatarInfo
	(AvatarType)(0),             // 1: AvatarType
}
var file_ChallengeAvatarInfo_proto_depIdxs = []int32{
	1, // 0: ChallengeAvatarInfo.avatar_type:type_name -> AvatarType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ChallengeAvatarInfo_proto_init() }
func file_ChallengeAvatarInfo_proto_init() {
	if File_ChallengeAvatarInfo_proto != nil {
		return
	}
	file_AvatarType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChallengeAvatarInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeAvatarInfo); i {
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
			RawDescriptor: file_ChallengeAvatarInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChallengeAvatarInfo_proto_goTypes,
		DependencyIndexes: file_ChallengeAvatarInfo_proto_depIdxs,
		MessageInfos:      file_ChallengeAvatarInfo_proto_msgTypes,
	}.Build()
	File_ChallengeAvatarInfo_proto = out.File
	file_ChallengeAvatarInfo_proto_rawDesc = nil
	file_ChallengeAvatarInfo_proto_goTypes = nil
	file_ChallengeAvatarInfo_proto_depIdxs = nil
}
