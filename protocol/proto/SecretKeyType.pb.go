// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SecretKeyType.proto

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

type SecretKeyType int32

const (
	SecretKeyType_SECRET_KEY_NONE         SecretKeyType = 0
	SecretKeyType_SECRET_KEY_SERVER_CHECK SecretKeyType = 1
	SecretKeyType_SECRET_KEY_VIDEO        SecretKeyType = 2
	SecretKeyType_SECRET_KEY_BATTLE_TIME  SecretKeyType = 3
)

// Enum value maps for SecretKeyType.
var (
	SecretKeyType_name = map[int32]string{
		0: "SECRET_KEY_NONE",
		1: "SECRET_KEY_SERVER_CHECK",
		2: "SECRET_KEY_VIDEO",
		3: "SECRET_KEY_BATTLE_TIME",
	}
	SecretKeyType_value = map[string]int32{
		"SECRET_KEY_NONE":         0,
		"SECRET_KEY_SERVER_CHECK": 1,
		"SECRET_KEY_VIDEO":        2,
		"SECRET_KEY_BATTLE_TIME":  3,
	}
)

func (x SecretKeyType) Enum() *SecretKeyType {
	p := new(SecretKeyType)
	*p = x
	return p
}

func (x SecretKeyType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SecretKeyType) Descriptor() protoreflect.EnumDescriptor {
	return file_SecretKeyType_proto_enumTypes[0].Descriptor()
}

func (SecretKeyType) Type() protoreflect.EnumType {
	return &file_SecretKeyType_proto_enumTypes[0]
}

func (x SecretKeyType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SecretKeyType.Descriptor instead.
func (SecretKeyType) EnumDescriptor() ([]byte, []int) {
	return file_SecretKeyType_proto_rawDescGZIP(), []int{0}
}

var File_SecretKeyType_proto protoreflect.FileDescriptor

var file_SecretKeyType_proto_rawDesc = []byte{
	0x0a, 0x13, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x73, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b,
	0x65, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54,
	0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x53,
	0x45, 0x43, 0x52, 0x45, 0x54, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52,
	0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x45, 0x43, 0x52,
	0x45, 0x54, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x02, 0x12, 0x1a,
	0x0a, 0x16, 0x53, 0x45, 0x43, 0x52, 0x45, 0x54, 0x5f, 0x4b, 0x45, 0x59, 0x5f, 0x42, 0x41, 0x54,
	0x54, 0x4c, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x03, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_SecretKeyType_proto_rawDescOnce sync.Once
	file_SecretKeyType_proto_rawDescData = file_SecretKeyType_proto_rawDesc
)

func file_SecretKeyType_proto_rawDescGZIP() []byte {
	file_SecretKeyType_proto_rawDescOnce.Do(func() {
		file_SecretKeyType_proto_rawDescData = protoimpl.X.CompressGZIP(file_SecretKeyType_proto_rawDescData)
	})
	return file_SecretKeyType_proto_rawDescData
}

var file_SecretKeyType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_SecretKeyType_proto_goTypes = []interface{}{
	(SecretKeyType)(0), // 0: SecretKeyType
}
var file_SecretKeyType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SecretKeyType_proto_init() }
func file_SecretKeyType_proto_init() {
	if File_SecretKeyType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SecretKeyType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SecretKeyType_proto_goTypes,
		DependencyIndexes: file_SecretKeyType_proto_depIdxs,
		EnumInfos:         file_SecretKeyType_proto_enumTypes,
	}.Build()
	File_SecretKeyType_proto = out.File
	file_SecretKeyType_proto_rawDesc = nil
	file_SecretKeyType_proto_goTypes = nil
	file_SecretKeyType_proto_depIdxs = nil
}
