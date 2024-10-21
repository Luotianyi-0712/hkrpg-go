// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: POKLPKKCAAD.proto

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

type POKLPKKCAAD struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarType AvatarType `protobuf:"varint,5,opt,name=avatar_type,json=avatarType,proto3,enum=AvatarType" json:"avatar_type,omitempty"`
	AvatarId   uint32     `protobuf:"varint,4,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
}

func (x *POKLPKKCAAD) Reset() {
	*x = POKLPKKCAAD{}
	if protoimpl.UnsafeEnabled {
		mi := &file_POKLPKKCAAD_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *POKLPKKCAAD) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*POKLPKKCAAD) ProtoMessage() {}

func (x *POKLPKKCAAD) ProtoReflect() protoreflect.Message {
	mi := &file_POKLPKKCAAD_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use POKLPKKCAAD.ProtoReflect.Descriptor instead.
func (*POKLPKKCAAD) Descriptor() ([]byte, []int) {
	return file_POKLPKKCAAD_proto_rawDescGZIP(), []int{0}
}

func (x *POKLPKKCAAD) GetAvatarType() AvatarType {
	if x != nil {
		return x.AvatarType
	}
	return AvatarType_AVATAR_TYPE_NONE
}

func (x *POKLPKKCAAD) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

var File_POKLPKKCAAD_proto protoreflect.FileDescriptor

var file_POKLPKKCAAD_proto_rawDesc = []byte{
	0x0a, 0x11, 0x50, 0x4f, 0x4b, 0x4c, 0x50, 0x4b, 0x4b, 0x43, 0x41, 0x41, 0x44, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a, 0x0b, 0x50, 0x4f, 0x4b, 0x4c, 0x50, 0x4b, 0x4b,
	0x43, 0x41, 0x41, 0x44, 0x12, 0x2c, 0x0a, 0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x41, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x42,
	0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_POKLPKKCAAD_proto_rawDescOnce sync.Once
	file_POKLPKKCAAD_proto_rawDescData = file_POKLPKKCAAD_proto_rawDesc
)

func file_POKLPKKCAAD_proto_rawDescGZIP() []byte {
	file_POKLPKKCAAD_proto_rawDescOnce.Do(func() {
		file_POKLPKKCAAD_proto_rawDescData = protoimpl.X.CompressGZIP(file_POKLPKKCAAD_proto_rawDescData)
	})
	return file_POKLPKKCAAD_proto_rawDescData
}

var file_POKLPKKCAAD_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_POKLPKKCAAD_proto_goTypes = []interface{}{
	(*POKLPKKCAAD)(nil), // 0: POKLPKKCAAD
	(AvatarType)(0),     // 1: AvatarType
}
var file_POKLPKKCAAD_proto_depIdxs = []int32{
	1, // 0: POKLPKKCAAD.avatar_type:type_name -> AvatarType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_POKLPKKCAAD_proto_init() }
func file_POKLPKKCAAD_proto_init() {
	if File_POKLPKKCAAD_proto != nil {
		return
	}
	file_AvatarType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_POKLPKKCAAD_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*POKLPKKCAAD); i {
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
			RawDescriptor: file_POKLPKKCAAD_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_POKLPKKCAAD_proto_goTypes,
		DependencyIndexes: file_POKLPKKCAAD_proto_depIdxs,
		MessageInfos:      file_POKLPKKCAAD_proto_msgTypes,
	}.Build()
	File_POKLPKKCAAD_proto = out.File
	file_POKLPKKCAAD_proto_rawDesc = nil
	file_POKLPKKCAAD_proto_goTypes = nil
	file_POKLPKKCAAD_proto_depIdxs = nil
}
