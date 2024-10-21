// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: CIAOIKEANEA.proto

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

type CIAOIKEANEA int32

const (
	CIAOIKEANEA_MONOPOLY_ACTION_RESULT_SOURCE_TYPE_NONE             CIAOIKEANEA = 0
	CIAOIKEANEA_MONOPOLY_ACTION_RESULT_SOURCE_TYPE_EFFECT           CIAOIKEANEA = 1
	CIAOIKEANEA_MONOPOLY_ACTION_RESULT_SOURCE_TYPE_ASSET_BONUS      CIAOIKEANEA = 2
	CIAOIKEANEA_MONOPOLY_ACTION_RESULT_SOURCE_TYPE_ASSET_TAX        CIAOIKEANEA = 3
	CIAOIKEANEA_MONOPOLY_ACTION_RESULT_SOURCE_TYPE_ASSET_UPGRADE    CIAOIKEANEA = 4
	CIAOIKEANEA_MONOPOLY_ACTION_RESULT_SOURCE_TYPE_GAME_SETTLE      CIAOIKEANEA = 5
	CIAOIKEANEA_MONOPOLY_ACTION_RESULT_SOURCE_TYPE_BUY_GOODS        CIAOIKEANEA = 6
	CIAOIKEANEA_MONOPOLY_ACTION_RESULT_SOURCE_TYPE_CLICK            CIAOIKEANEA = 7
	CIAOIKEANEA_MONOPOLY_ACTION_RESULT_SOURCE_TYPE_SOCIAL_EVENT     CIAOIKEANEA = 8
	CIAOIKEANEA_MONOPOLY_ACTION_RESULT_SOURCE_TYPE_LIKE             CIAOIKEANEA = 9
	CIAOIKEANEA_MONOPOLY_ACTION_RESULT_SOURCE_TYPE_QUIZ_GAME_SETTLE CIAOIKEANEA = 10
)

// Enum value maps for CIAOIKEANEA.
var (
	CIAOIKEANEA_name = map[int32]string{
		0:  "MONOPOLY_ACTION_RESULT_SOURCE_TYPE_NONE",
		1:  "MONOPOLY_ACTION_RESULT_SOURCE_TYPE_EFFECT",
		2:  "MONOPOLY_ACTION_RESULT_SOURCE_TYPE_ASSET_BONUS",
		3:  "MONOPOLY_ACTION_RESULT_SOURCE_TYPE_ASSET_TAX",
		4:  "MONOPOLY_ACTION_RESULT_SOURCE_TYPE_ASSET_UPGRADE",
		5:  "MONOPOLY_ACTION_RESULT_SOURCE_TYPE_GAME_SETTLE",
		6:  "MONOPOLY_ACTION_RESULT_SOURCE_TYPE_BUY_GOODS",
		7:  "MONOPOLY_ACTION_RESULT_SOURCE_TYPE_CLICK",
		8:  "MONOPOLY_ACTION_RESULT_SOURCE_TYPE_SOCIAL_EVENT",
		9:  "MONOPOLY_ACTION_RESULT_SOURCE_TYPE_LIKE",
		10: "MONOPOLY_ACTION_RESULT_SOURCE_TYPE_QUIZ_GAME_SETTLE",
	}
	CIAOIKEANEA_value = map[string]int32{
		"MONOPOLY_ACTION_RESULT_SOURCE_TYPE_NONE":             0,
		"MONOPOLY_ACTION_RESULT_SOURCE_TYPE_EFFECT":           1,
		"MONOPOLY_ACTION_RESULT_SOURCE_TYPE_ASSET_BONUS":      2,
		"MONOPOLY_ACTION_RESULT_SOURCE_TYPE_ASSET_TAX":        3,
		"MONOPOLY_ACTION_RESULT_SOURCE_TYPE_ASSET_UPGRADE":    4,
		"MONOPOLY_ACTION_RESULT_SOURCE_TYPE_GAME_SETTLE":      5,
		"MONOPOLY_ACTION_RESULT_SOURCE_TYPE_BUY_GOODS":        6,
		"MONOPOLY_ACTION_RESULT_SOURCE_TYPE_CLICK":            7,
		"MONOPOLY_ACTION_RESULT_SOURCE_TYPE_SOCIAL_EVENT":     8,
		"MONOPOLY_ACTION_RESULT_SOURCE_TYPE_LIKE":             9,
		"MONOPOLY_ACTION_RESULT_SOURCE_TYPE_QUIZ_GAME_SETTLE": 10,
	}
)

func (x CIAOIKEANEA) Enum() *CIAOIKEANEA {
	p := new(CIAOIKEANEA)
	*p = x
	return p
}

func (x CIAOIKEANEA) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CIAOIKEANEA) Descriptor() protoreflect.EnumDescriptor {
	return file_CIAOIKEANEA_proto_enumTypes[0].Descriptor()
}

func (CIAOIKEANEA) Type() protoreflect.EnumType {
	return &file_CIAOIKEANEA_proto_enumTypes[0]
}

func (x CIAOIKEANEA) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CIAOIKEANEA.Descriptor instead.
func (CIAOIKEANEA) EnumDescriptor() ([]byte, []int) {
	return file_CIAOIKEANEA_proto_rawDescGZIP(), []int{0}
}

var File_CIAOIKEANEA_proto protoreflect.FileDescriptor

var file_CIAOIKEANEA_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x49, 0x41, 0x4f, 0x49, 0x4b, 0x45, 0x41, 0x4e, 0x45, 0x41, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0xb4, 0x04, 0x0a, 0x0b, 0x43, 0x49, 0x41, 0x4f, 0x49, 0x4b, 0x45, 0x41,
	0x4e, 0x45, 0x41, 0x12, 0x2b, 0x0a, 0x27, 0x4d, 0x4f, 0x4e, 0x4f, 0x50, 0x4f, 0x4c, 0x59, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x2d, 0x0a, 0x29, 0x4d, 0x4f, 0x4e, 0x4f, 0x50, 0x4f, 0x4c, 0x59, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x46, 0x46, 0x45, 0x43, 0x54, 0x10, 0x01, 0x12,
	0x32, 0x0a, 0x2e, 0x4d, 0x4f, 0x4e, 0x4f, 0x50, 0x4f, 0x4c, 0x59, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f, 0x42, 0x4f, 0x4e, 0x55,
	0x53, 0x10, 0x02, 0x12, 0x30, 0x0a, 0x2c, 0x4d, 0x4f, 0x4e, 0x4f, 0x50, 0x4f, 0x4c, 0x59, 0x5f,
	0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x4f,
	0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45, 0x54, 0x5f,
	0x54, 0x41, 0x58, 0x10, 0x03, 0x12, 0x34, 0x0a, 0x30, 0x4d, 0x4f, 0x4e, 0x4f, 0x50, 0x4f, 0x4c,
	0x59, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f,
	0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x53, 0x53, 0x45,
	0x54, 0x5f, 0x55, 0x50, 0x47, 0x52, 0x41, 0x44, 0x45, 0x10, 0x04, 0x12, 0x32, 0x0a, 0x2e, 0x4d,
	0x4f, 0x4e, 0x4f, 0x50, 0x4f, 0x4c, 0x59, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52,
	0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x10, 0x05, 0x12,
	0x30, 0x0a, 0x2c, 0x4d, 0x4f, 0x4e, 0x4f, 0x50, 0x4f, 0x4c, 0x59, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x55, 0x59, 0x5f, 0x47, 0x4f, 0x4f, 0x44, 0x53, 0x10,
	0x06, 0x12, 0x2c, 0x0a, 0x28, 0x4d, 0x4f, 0x4e, 0x4f, 0x50, 0x4f, 0x4c, 0x59, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x10, 0x07, 0x12,
	0x33, 0x0a, 0x2f, 0x4d, 0x4f, 0x4e, 0x4f, 0x50, 0x4f, 0x4c, 0x59, 0x5f, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4f, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x10, 0x08, 0x12, 0x2b, 0x0a, 0x27, 0x4d, 0x4f, 0x4e, 0x4f, 0x50, 0x4f, 0x4c, 0x59,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53,
	0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4c, 0x49, 0x4b, 0x45, 0x10,
	0x09, 0x12, 0x37, 0x0a, 0x33, 0x4d, 0x4f, 0x4e, 0x4f, 0x50, 0x4f, 0x4c, 0x59, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x51, 0x55, 0x49, 0x5a, 0x5f, 0x47, 0x41, 0x4d,
	0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x10, 0x0a, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_CIAOIKEANEA_proto_rawDescOnce sync.Once
	file_CIAOIKEANEA_proto_rawDescData = file_CIAOIKEANEA_proto_rawDesc
)

func file_CIAOIKEANEA_proto_rawDescGZIP() []byte {
	file_CIAOIKEANEA_proto_rawDescOnce.Do(func() {
		file_CIAOIKEANEA_proto_rawDescData = protoimpl.X.CompressGZIP(file_CIAOIKEANEA_proto_rawDescData)
	})
	return file_CIAOIKEANEA_proto_rawDescData
}

var file_CIAOIKEANEA_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_CIAOIKEANEA_proto_goTypes = []interface{}{
	(CIAOIKEANEA)(0), // 0: CIAOIKEANEA
}
var file_CIAOIKEANEA_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_CIAOIKEANEA_proto_init() }
func file_CIAOIKEANEA_proto_init() {
	if File_CIAOIKEANEA_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_CIAOIKEANEA_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CIAOIKEANEA_proto_goTypes,
		DependencyIndexes: file_CIAOIKEANEA_proto_depIdxs,
		EnumInfos:         file_CIAOIKEANEA_proto_enumTypes,
	}.Build()
	File_CIAOIKEANEA_proto = out.File
	file_CIAOIKEANEA_proto_rawDesc = nil
	file_CIAOIKEANEA_proto_goTypes = nil
	file_CIAOIKEANEA_proto_depIdxs = nil
}
