// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: NBBDAMGNJPN.proto

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

type NBBDAMGNJPN int32

const (
	NBBDAMGNJPN_MONOPOLY_SOCIAL_EVENT_STATUS_NONE                  NBBDAMGNJPN = 0
	NBBDAMGNJPN_MONOPOLY_SOCIAL_EVENT_STATUS_WAITING_SELECT_FRIEND NBBDAMGNJPN = 1
)

// Enum value maps for NBBDAMGNJPN.
var (
	NBBDAMGNJPN_name = map[int32]string{
		0: "MONOPOLY_SOCIAL_EVENT_STATUS_NONE",
		1: "MONOPOLY_SOCIAL_EVENT_STATUS_WAITING_SELECT_FRIEND",
	}
	NBBDAMGNJPN_value = map[string]int32{
		"MONOPOLY_SOCIAL_EVENT_STATUS_NONE":                  0,
		"MONOPOLY_SOCIAL_EVENT_STATUS_WAITING_SELECT_FRIEND": 1,
	}
)

func (x NBBDAMGNJPN) Enum() *NBBDAMGNJPN {
	p := new(NBBDAMGNJPN)
	*p = x
	return p
}

func (x NBBDAMGNJPN) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NBBDAMGNJPN) Descriptor() protoreflect.EnumDescriptor {
	return file_NBBDAMGNJPN_proto_enumTypes[0].Descriptor()
}

func (NBBDAMGNJPN) Type() protoreflect.EnumType {
	return &file_NBBDAMGNJPN_proto_enumTypes[0]
}

func (x NBBDAMGNJPN) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NBBDAMGNJPN.Descriptor instead.
func (NBBDAMGNJPN) EnumDescriptor() ([]byte, []int) {
	return file_NBBDAMGNJPN_proto_rawDescGZIP(), []int{0}
}

var File_NBBDAMGNJPN_proto protoreflect.FileDescriptor

var file_NBBDAMGNJPN_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4e, 0x42, 0x42, 0x44, 0x41, 0x4d, 0x47, 0x4e, 0x4a, 0x50, 0x4e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x6c, 0x0a, 0x0b, 0x4e, 0x42, 0x42, 0x44, 0x41, 0x4d, 0x47, 0x4e, 0x4a,
	0x50, 0x4e, 0x12, 0x25, 0x0a, 0x21, 0x4d, 0x4f, 0x4e, 0x4f, 0x50, 0x4f, 0x4c, 0x59, 0x5f, 0x53,
	0x4f, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x36, 0x0a, 0x32, 0x4d, 0x4f, 0x4e,
	0x4f, 0x50, 0x4f, 0x4c, 0x59, 0x5f, 0x53, 0x4f, 0x43, 0x49, 0x41, 0x4c, 0x5f, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x57, 0x41, 0x49, 0x54, 0x49, 0x4e,
	0x47, 0x5f, 0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x10,
	0x01, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_NBBDAMGNJPN_proto_rawDescOnce sync.Once
	file_NBBDAMGNJPN_proto_rawDescData = file_NBBDAMGNJPN_proto_rawDesc
)

func file_NBBDAMGNJPN_proto_rawDescGZIP() []byte {
	file_NBBDAMGNJPN_proto_rawDescOnce.Do(func() {
		file_NBBDAMGNJPN_proto_rawDescData = protoimpl.X.CompressGZIP(file_NBBDAMGNJPN_proto_rawDescData)
	})
	return file_NBBDAMGNJPN_proto_rawDescData
}

var file_NBBDAMGNJPN_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_NBBDAMGNJPN_proto_goTypes = []interface{}{
	(NBBDAMGNJPN)(0), // 0: NBBDAMGNJPN
}
var file_NBBDAMGNJPN_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_NBBDAMGNJPN_proto_init() }
func file_NBBDAMGNJPN_proto_init() {
	if File_NBBDAMGNJPN_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_NBBDAMGNJPN_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NBBDAMGNJPN_proto_goTypes,
		DependencyIndexes: file_NBBDAMGNJPN_proto_depIdxs,
		EnumInfos:         file_NBBDAMGNJPN_proto_enumTypes,
	}.Build()
	File_NBBDAMGNJPN_proto = out.File
	file_NBBDAMGNJPN_proto_rawDesc = nil
	file_NBBDAMGNJPN_proto_goTypes = nil
	file_NBBDAMGNJPN_proto_depIdxs = nil
}
