// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GHBIABDNMFA.proto

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

type GHBIABDNMFA int32

const (
	GHBIABDNMFA_UNKNOWN          GHBIABDNMFA = 0
	GHBIABDNMFA_KILLED_BY_OTHERS GHBIABDNMFA = 1
	GHBIABDNMFA_KILLED_BY_SELF   GHBIABDNMFA = 2
	GHBIABDNMFA_ESCAPE           GHBIABDNMFA = 3
)

// Enum value maps for GHBIABDNMFA.
var (
	GHBIABDNMFA_name = map[int32]string{
		0: "UNKNOWN",
		1: "KILLED_BY_OTHERS",
		2: "KILLED_BY_SELF",
		3: "ESCAPE",
	}
	GHBIABDNMFA_value = map[string]int32{
		"UNKNOWN":          0,
		"KILLED_BY_OTHERS": 1,
		"KILLED_BY_SELF":   2,
		"ESCAPE":           3,
	}
)

func (x GHBIABDNMFA) Enum() *GHBIABDNMFA {
	p := new(GHBIABDNMFA)
	*p = x
	return p
}

func (x GHBIABDNMFA) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GHBIABDNMFA) Descriptor() protoreflect.EnumDescriptor {
	return file_GHBIABDNMFA_proto_enumTypes[0].Descriptor()
}

func (GHBIABDNMFA) Type() protoreflect.EnumType {
	return &file_GHBIABDNMFA_proto_enumTypes[0]
}

func (x GHBIABDNMFA) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GHBIABDNMFA.Descriptor instead.
func (GHBIABDNMFA) EnumDescriptor() ([]byte, []int) {
	return file_GHBIABDNMFA_proto_rawDescGZIP(), []int{0}
}

var File_GHBIABDNMFA_proto protoreflect.FileDescriptor

var file_GHBIABDNMFA_proto_rawDesc = []byte{
	0x0a, 0x11, 0x47, 0x48, 0x42, 0x49, 0x41, 0x42, 0x44, 0x4e, 0x4d, 0x46, 0x41, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x50, 0x0a, 0x0b, 0x47, 0x48, 0x42, 0x49, 0x41, 0x42, 0x44, 0x4e, 0x4d,
	0x46, 0x41, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12,
	0x14, 0x0a, 0x10, 0x4b, 0x49, 0x4c, 0x4c, 0x45, 0x44, 0x5f, 0x42, 0x59, 0x5f, 0x4f, 0x54, 0x48,
	0x45, 0x52, 0x53, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4b, 0x49, 0x4c, 0x4c, 0x45, 0x44, 0x5f,
	0x42, 0x59, 0x5f, 0x53, 0x45, 0x4c, 0x46, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x53, 0x43,
	0x41, 0x50, 0x45, 0x10, 0x03, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GHBIABDNMFA_proto_rawDescOnce sync.Once
	file_GHBIABDNMFA_proto_rawDescData = file_GHBIABDNMFA_proto_rawDesc
)

func file_GHBIABDNMFA_proto_rawDescGZIP() []byte {
	file_GHBIABDNMFA_proto_rawDescOnce.Do(func() {
		file_GHBIABDNMFA_proto_rawDescData = protoimpl.X.CompressGZIP(file_GHBIABDNMFA_proto_rawDescData)
	})
	return file_GHBIABDNMFA_proto_rawDescData
}

var file_GHBIABDNMFA_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GHBIABDNMFA_proto_goTypes = []interface{}{
	(GHBIABDNMFA)(0), // 0: GHBIABDNMFA
}
var file_GHBIABDNMFA_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GHBIABDNMFA_proto_init() }
func file_GHBIABDNMFA_proto_init() {
	if File_GHBIABDNMFA_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GHBIABDNMFA_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GHBIABDNMFA_proto_goTypes,
		DependencyIndexes: file_GHBIABDNMFA_proto_depIdxs,
		EnumInfos:         file_GHBIABDNMFA_proto_enumTypes,
	}.Build()
	File_GHBIABDNMFA_proto = out.File
	file_GHBIABDNMFA_proto_rawDesc = nil
	file_GHBIABDNMFA_proto_goTypes = nil
	file_GHBIABDNMFA_proto_depIdxs = nil
}