// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: JFLDMDNBOKF.proto

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

type JFLDMDNBOKF int32

const (
	JFLDMDNBOKF_ROGUE_MODIFIER_CONTENT_DEFINITE JFLDMDNBOKF = 0
	JFLDMDNBOKF_ROGUE_MODIFIER_CONTENT_RANDOM   JFLDMDNBOKF = 1
)

// Enum value maps for JFLDMDNBOKF.
var (
	JFLDMDNBOKF_name = map[int32]string{
		0: "ROGUE_MODIFIER_CONTENT_DEFINITE",
		1: "ROGUE_MODIFIER_CONTENT_RANDOM",
	}
	JFLDMDNBOKF_value = map[string]int32{
		"ROGUE_MODIFIER_CONTENT_DEFINITE": 0,
		"ROGUE_MODIFIER_CONTENT_RANDOM":   1,
	}
)

func (x JFLDMDNBOKF) Enum() *JFLDMDNBOKF {
	p := new(JFLDMDNBOKF)
	*p = x
	return p
}

func (x JFLDMDNBOKF) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JFLDMDNBOKF) Descriptor() protoreflect.EnumDescriptor {
	return file_JFLDMDNBOKF_proto_enumTypes[0].Descriptor()
}

func (JFLDMDNBOKF) Type() protoreflect.EnumType {
	return &file_JFLDMDNBOKF_proto_enumTypes[0]
}

func (x JFLDMDNBOKF) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JFLDMDNBOKF.Descriptor instead.
func (JFLDMDNBOKF) EnumDescriptor() ([]byte, []int) {
	return file_JFLDMDNBOKF_proto_rawDescGZIP(), []int{0}
}

var File_JFLDMDNBOKF_proto protoreflect.FileDescriptor

var file_JFLDMDNBOKF_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4a, 0x46, 0x4c, 0x44, 0x4d, 0x44, 0x4e, 0x42, 0x4f, 0x4b, 0x46, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x55, 0x0a, 0x0b, 0x4a, 0x46, 0x4c, 0x44, 0x4d, 0x44, 0x4e, 0x42, 0x4f,
	0x4b, 0x46, 0x12, 0x23, 0x0a, 0x1f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x49,
	0x46, 0x49, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e, 0x54, 0x5f, 0x44, 0x45, 0x46,
	0x49, 0x4e, 0x49, 0x54, 0x45, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x52, 0x4f, 0x47, 0x55, 0x45,
	0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x49, 0x45, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x45, 0x4e,
	0x54, 0x5f, 0x52, 0x41, 0x4e, 0x44, 0x4f, 0x4d, 0x10, 0x01, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_JFLDMDNBOKF_proto_rawDescOnce sync.Once
	file_JFLDMDNBOKF_proto_rawDescData = file_JFLDMDNBOKF_proto_rawDesc
)

func file_JFLDMDNBOKF_proto_rawDescGZIP() []byte {
	file_JFLDMDNBOKF_proto_rawDescOnce.Do(func() {
		file_JFLDMDNBOKF_proto_rawDescData = protoimpl.X.CompressGZIP(file_JFLDMDNBOKF_proto_rawDescData)
	})
	return file_JFLDMDNBOKF_proto_rawDescData
}

var file_JFLDMDNBOKF_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_JFLDMDNBOKF_proto_goTypes = []interface{}{
	(JFLDMDNBOKF)(0), // 0: JFLDMDNBOKF
}
var file_JFLDMDNBOKF_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_JFLDMDNBOKF_proto_init() }
func file_JFLDMDNBOKF_proto_init() {
	if File_JFLDMDNBOKF_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_JFLDMDNBOKF_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_JFLDMDNBOKF_proto_goTypes,
		DependencyIndexes: file_JFLDMDNBOKF_proto_depIdxs,
		EnumInfos:         file_JFLDMDNBOKF_proto_enumTypes,
	}.Build()
	File_JFLDMDNBOKF_proto = out.File
	file_JFLDMDNBOKF_proto_rawDesc = nil
	file_JFLDMDNBOKF_proto_goTypes = nil
	file_JFLDMDNBOKF_proto_depIdxs = nil
}