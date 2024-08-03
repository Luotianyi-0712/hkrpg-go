// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: AetherdivideSpiritLineupType.proto

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

type AetherdivideSpiritLineupType int32

const (
	AetherdivideSpiritLineupType_AETHERDIVIDE_SPIRIT_LINEUP_NONE   AetherdivideSpiritLineupType = 0
	AetherdivideSpiritLineupType_AETHERDIVIDE_SPIRIT_LINEUP_NORMAL AetherdivideSpiritLineupType = 1
	AetherdivideSpiritLineupType_AETHERDIVIDE_SPIRIT_LINEUP_TRIAL  AetherdivideSpiritLineupType = 2
)

// Enum value maps for AetherdivideSpiritLineupType.
var (
	AetherdivideSpiritLineupType_name = map[int32]string{
		0: "AETHERDIVIDE_SPIRIT_LINEUP_NONE",
		1: "AETHERDIVIDE_SPIRIT_LINEUP_NORMAL",
		2: "AETHERDIVIDE_SPIRIT_LINEUP_TRIAL",
	}
	AetherdivideSpiritLineupType_value = map[string]int32{
		"AETHERDIVIDE_SPIRIT_LINEUP_NONE":   0,
		"AETHERDIVIDE_SPIRIT_LINEUP_NORMAL": 1,
		"AETHERDIVIDE_SPIRIT_LINEUP_TRIAL":  2,
	}
)

func (x AetherdivideSpiritLineupType) Enum() *AetherdivideSpiritLineupType {
	p := new(AetherdivideSpiritLineupType)
	*p = x
	return p
}

func (x AetherdivideSpiritLineupType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AetherdivideSpiritLineupType) Descriptor() protoreflect.EnumDescriptor {
	return file_AetherdivideSpiritLineupType_proto_enumTypes[0].Descriptor()
}

func (AetherdivideSpiritLineupType) Type() protoreflect.EnumType {
	return &file_AetherdivideSpiritLineupType_proto_enumTypes[0]
}

func (x AetherdivideSpiritLineupType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AetherdivideSpiritLineupType.Descriptor instead.
func (AetherdivideSpiritLineupType) EnumDescriptor() ([]byte, []int) {
	return file_AetherdivideSpiritLineupType_proto_rawDescGZIP(), []int{0}
}

var File_AetherdivideSpiritLineupType_proto protoreflect.FileDescriptor

var file_AetherdivideSpiritLineupType_proto_rawDesc = []byte{
	0x0a, 0x22, 0x41, 0x65, 0x74, 0x68, 0x65, 0x72, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x53, 0x70,
	0x69, 0x72, 0x69, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x90, 0x01, 0x0a, 0x1c, 0x41, 0x65, 0x74, 0x68, 0x65, 0x72, 0x64,
	0x69, 0x76, 0x69, 0x64, 0x65, 0x53, 0x70, 0x69, 0x72, 0x69, 0x74, 0x4c, 0x69, 0x6e, 0x65, 0x75,
	0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x45, 0x54, 0x48, 0x45, 0x52, 0x44,
	0x49, 0x56, 0x49, 0x44, 0x45, 0x5f, 0x53, 0x50, 0x49, 0x52, 0x49, 0x54, 0x5f, 0x4c, 0x49, 0x4e,
	0x45, 0x55, 0x50, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x41, 0x45,
	0x54, 0x48, 0x45, 0x52, 0x44, 0x49, 0x56, 0x49, 0x44, 0x45, 0x5f, 0x53, 0x50, 0x49, 0x52, 0x49,
	0x54, 0x5f, 0x4c, 0x49, 0x4e, 0x45, 0x55, 0x50, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10,
	0x01, 0x12, 0x24, 0x0a, 0x20, 0x41, 0x45, 0x54, 0x48, 0x45, 0x52, 0x44, 0x49, 0x56, 0x49, 0x44,
	0x45, 0x5f, 0x53, 0x50, 0x49, 0x52, 0x49, 0x54, 0x5f, 0x4c, 0x49, 0x4e, 0x45, 0x55, 0x50, 0x5f,
	0x54, 0x52, 0x49, 0x41, 0x4c, 0x10, 0x02, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AetherdivideSpiritLineupType_proto_rawDescOnce sync.Once
	file_AetherdivideSpiritLineupType_proto_rawDescData = file_AetherdivideSpiritLineupType_proto_rawDesc
)

func file_AetherdivideSpiritLineupType_proto_rawDescGZIP() []byte {
	file_AetherdivideSpiritLineupType_proto_rawDescOnce.Do(func() {
		file_AetherdivideSpiritLineupType_proto_rawDescData = protoimpl.X.CompressGZIP(file_AetherdivideSpiritLineupType_proto_rawDescData)
	})
	return file_AetherdivideSpiritLineupType_proto_rawDescData
}

var file_AetherdivideSpiritLineupType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_AetherdivideSpiritLineupType_proto_goTypes = []interface{}{
	(AetherdivideSpiritLineupType)(0), // 0: AetherdivideSpiritLineupType
}
var file_AetherdivideSpiritLineupType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AetherdivideSpiritLineupType_proto_init() }
func file_AetherdivideSpiritLineupType_proto_init() {
	if File_AetherdivideSpiritLineupType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_AetherdivideSpiritLineupType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AetherdivideSpiritLineupType_proto_goTypes,
		DependencyIndexes: file_AetherdivideSpiritLineupType_proto_depIdxs,
		EnumInfos:         file_AetherdivideSpiritLineupType_proto_enumTypes,
	}.Build()
	File_AetherdivideSpiritLineupType_proto = out.File
	file_AetherdivideSpiritLineupType_proto_rawDesc = nil
	file_AetherdivideSpiritLineupType_proto_goTypes = nil
	file_AetherdivideSpiritLineupType_proto_depIdxs = nil
}