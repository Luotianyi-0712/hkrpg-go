// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RebattleType.proto

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

type RebattleType int32

const (
	RebattleType_REBATTLE_TYPE_NONE                   RebattleType = 0
	RebattleType_REBATTLE_TYPE_REBATTLE_MIDWAY        RebattleType = 1
	RebattleType_REBATTLE_TYPE_REBATTLE_LOSE          RebattleType = 2
	RebattleType_REBATTLE_TYPE_REBATTLE_MIDWAY_LINEUP RebattleType = 3
	RebattleType_REBATTLE_TYPE_REBATTLE_LOSE_LINEUP   RebattleType = 4
	RebattleType_REBATTLE_TYPE_QUIT_MIDWAY            RebattleType = 5
	RebattleType_REBATTLE_TYPE_QUIT_LOSE              RebattleType = 6
)

// Enum value maps for RebattleType.
var (
	RebattleType_name = map[int32]string{
		0: "REBATTLE_TYPE_NONE",
		1: "REBATTLE_TYPE_REBATTLE_MIDWAY",
		2: "REBATTLE_TYPE_REBATTLE_LOSE",
		3: "REBATTLE_TYPE_REBATTLE_MIDWAY_LINEUP",
		4: "REBATTLE_TYPE_REBATTLE_LOSE_LINEUP",
		5: "REBATTLE_TYPE_QUIT_MIDWAY",
		6: "REBATTLE_TYPE_QUIT_LOSE",
	}
	RebattleType_value = map[string]int32{
		"REBATTLE_TYPE_NONE":                   0,
		"REBATTLE_TYPE_REBATTLE_MIDWAY":        1,
		"REBATTLE_TYPE_REBATTLE_LOSE":          2,
		"REBATTLE_TYPE_REBATTLE_MIDWAY_LINEUP": 3,
		"REBATTLE_TYPE_REBATTLE_LOSE_LINEUP":   4,
		"REBATTLE_TYPE_QUIT_MIDWAY":            5,
		"REBATTLE_TYPE_QUIT_LOSE":              6,
	}
)

func (x RebattleType) Enum() *RebattleType {
	p := new(RebattleType)
	*p = x
	return p
}

func (x RebattleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RebattleType) Descriptor() protoreflect.EnumDescriptor {
	return file_RebattleType_proto_enumTypes[0].Descriptor()
}

func (RebattleType) Type() protoreflect.EnumType {
	return &file_RebattleType_proto_enumTypes[0]
}

func (x RebattleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RebattleType.Descriptor instead.
func (RebattleType) EnumDescriptor() ([]byte, []int) {
	return file_RebattleType_proto_rawDescGZIP(), []int{0}
}

var File_RebattleType_proto protoreflect.FileDescriptor

var file_RebattleType_proto_rawDesc = []byte{
	0x0a, 0x12, 0x52, 0x65, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xf8, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x62, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x52, 0x45, 0x42, 0x41, 0x54, 0x54, 0x4c,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x21, 0x0a,
	0x1d, 0x52, 0x45, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52,
	0x45, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x4d, 0x49, 0x44, 0x57, 0x41, 0x59, 0x10, 0x01,
	0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x52, 0x45, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x4c, 0x4f, 0x53, 0x45, 0x10,
	0x02, 0x12, 0x28, 0x0a, 0x24, 0x52, 0x45, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x52, 0x45, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x4d, 0x49, 0x44, 0x57,
	0x41, 0x59, 0x5f, 0x4c, 0x49, 0x4e, 0x45, 0x55, 0x50, 0x10, 0x03, 0x12, 0x26, 0x0a, 0x22, 0x52,
	0x45, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x42,
	0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x4c, 0x4f, 0x53, 0x45, 0x5f, 0x4c, 0x49, 0x4e, 0x45, 0x55,
	0x50, 0x10, 0x04, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x51, 0x55, 0x49, 0x54, 0x5f, 0x4d, 0x49, 0x44, 0x57, 0x41, 0x59,
	0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x51, 0x55, 0x49, 0x54, 0x5f, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x06, 0x42,
	0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RebattleType_proto_rawDescOnce sync.Once
	file_RebattleType_proto_rawDescData = file_RebattleType_proto_rawDesc
)

func file_RebattleType_proto_rawDescGZIP() []byte {
	file_RebattleType_proto_rawDescOnce.Do(func() {
		file_RebattleType_proto_rawDescData = protoimpl.X.CompressGZIP(file_RebattleType_proto_rawDescData)
	})
	return file_RebattleType_proto_rawDescData
}

var file_RebattleType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_RebattleType_proto_goTypes = []interface{}{
	(RebattleType)(0), // 0: RebattleType
}
var file_RebattleType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RebattleType_proto_init() }
func file_RebattleType_proto_init() {
	if File_RebattleType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RebattleType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RebattleType_proto_goTypes,
		DependencyIndexes: file_RebattleType_proto_depIdxs,
		EnumInfos:         file_RebattleType_proto_enumTypes,
	}.Build()
	File_RebattleType_proto = out.File
	file_RebattleType_proto_rawDesc = nil
	file_RebattleType_proto_goTypes = nil
	file_RebattleType_proto_depIdxs = nil
}
