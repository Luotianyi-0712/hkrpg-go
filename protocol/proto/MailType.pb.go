// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MailType.proto

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

type MailType int32

const (
	MailType_MAIL_TYPE_NORMAL MailType = 0
	MailType_MAIL_TYPE_STAR   MailType = 1
)

// Enum value maps for MailType.
var (
	MailType_name = map[int32]string{
		0: "MAIL_TYPE_NORMAL",
		1: "MAIL_TYPE_STAR",
	}
	MailType_value = map[string]int32{
		"MAIL_TYPE_NORMAL": 0,
		"MAIL_TYPE_STAR":   1,
	}
)

func (x MailType) Enum() *MailType {
	p := new(MailType)
	*p = x
	return p
}

func (x MailType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MailType) Descriptor() protoreflect.EnumDescriptor {
	return file_MailType_proto_enumTypes[0].Descriptor()
}

func (MailType) Type() protoreflect.EnumType {
	return &file_MailType_proto_enumTypes[0]
}

func (x MailType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MailType.Descriptor instead.
func (MailType) EnumDescriptor() ([]byte, []int) {
	return file_MailType_proto_rawDescGZIP(), []int{0}
}

var File_MailType_proto protoreflect.FileDescriptor

var file_MailType_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0x34, 0x0a, 0x08, 0x4d, 0x61, 0x69, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10,
	0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x54, 0x41, 0x52, 0x10, 0x01, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MailType_proto_rawDescOnce sync.Once
	file_MailType_proto_rawDescData = file_MailType_proto_rawDesc
)

func file_MailType_proto_rawDescGZIP() []byte {
	file_MailType_proto_rawDescOnce.Do(func() {
		file_MailType_proto_rawDescData = protoimpl.X.CompressGZIP(file_MailType_proto_rawDescData)
	})
	return file_MailType_proto_rawDescData
}

var file_MailType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MailType_proto_goTypes = []interface{}{
	(MailType)(0), // 0: MailType
}
var file_MailType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MailType_proto_init() }
func file_MailType_proto_init() {
	if File_MailType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MailType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MailType_proto_goTypes,
		DependencyIndexes: file_MailType_proto_depIdxs,
		EnumInfos:         file_MailType_proto_enumTypes,
	}.Build()
	File_MailType_proto = out.File
	file_MailType_proto_rawDesc = nil
	file_MailType_proto_goTypes = nil
	file_MailType_proto_depIdxs = nil
}
