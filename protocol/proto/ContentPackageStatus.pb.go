// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ContentPackageStatus.proto

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

type ContentPackageStatus int32

const (
	ContentPackageStatus_ContentPackageStatus_None     ContentPackageStatus = 0
	ContentPackageStatus_ContentPackageStatus_Init     ContentPackageStatus = 1
	ContentPackageStatus_ContentPackageStatus_Doing    ContentPackageStatus = 2
	ContentPackageStatus_ContentPackageStatus_Finished ContentPackageStatus = 3
	ContentPackageStatus_ContentPackageStatus_Release  ContentPackageStatus = 4
)

// Enum value maps for ContentPackageStatus.
var (
	ContentPackageStatus_name = map[int32]string{
		0: "ContentPackageStatus_None",
		1: "ContentPackageStatus_Init",
		2: "ContentPackageStatus_Doing",
		3: "ContentPackageStatus_Finished",
		4: "ContentPackageStatus_Release",
	}
	ContentPackageStatus_value = map[string]int32{
		"ContentPackageStatus_None":     0,
		"ContentPackageStatus_Init":     1,
		"ContentPackageStatus_Doing":    2,
		"ContentPackageStatus_Finished": 3,
		"ContentPackageStatus_Release":  4,
	}
)

func (x ContentPackageStatus) Enum() *ContentPackageStatus {
	p := new(ContentPackageStatus)
	*p = x
	return p
}

func (x ContentPackageStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContentPackageStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_ContentPackageStatus_proto_enumTypes[0].Descriptor()
}

func (ContentPackageStatus) Type() protoreflect.EnumType {
	return &file_ContentPackageStatus_proto_enumTypes[0]
}

func (x ContentPackageStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContentPackageStatus.Descriptor instead.
func (ContentPackageStatus) EnumDescriptor() ([]byte, []int) {
	return file_ContentPackageStatus_proto_rawDescGZIP(), []int{0}
}

var File_ContentPackageStatus_proto protoreflect.FileDescriptor

var file_ContentPackageStatus_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xb9, 0x01, 0x0a,
	0x14, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x4e, 0x6f,
	0x6e, 0x65, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x49, 0x6e, 0x69,
	0x74, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x44, 0x6f, 0x69, 0x6e,
	0x67, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x50, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x46, 0x69, 0x6e, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x10, 0x03, 0x12, 0x20, 0x0a, 0x1c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x52,
	0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x10, 0x04, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ContentPackageStatus_proto_rawDescOnce sync.Once
	file_ContentPackageStatus_proto_rawDescData = file_ContentPackageStatus_proto_rawDesc
)

func file_ContentPackageStatus_proto_rawDescGZIP() []byte {
	file_ContentPackageStatus_proto_rawDescOnce.Do(func() {
		file_ContentPackageStatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_ContentPackageStatus_proto_rawDescData)
	})
	return file_ContentPackageStatus_proto_rawDescData
}

var file_ContentPackageStatus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ContentPackageStatus_proto_goTypes = []interface{}{
	(ContentPackageStatus)(0), // 0: ContentPackageStatus
}
var file_ContentPackageStatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ContentPackageStatus_proto_init() }
func file_ContentPackageStatus_proto_init() {
	if File_ContentPackageStatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ContentPackageStatus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ContentPackageStatus_proto_goTypes,
		DependencyIndexes: file_ContentPackageStatus_proto_depIdxs,
		EnumInfos:         file_ContentPackageStatus_proto_enumTypes,
	}.Build()
	File_ContentPackageStatus_proto = out.File
	file_ContentPackageStatus_proto_rawDesc = nil
	file_ContentPackageStatus_proto_goTypes = nil
	file_ContentPackageStatus_proto_depIdxs = nil
}
