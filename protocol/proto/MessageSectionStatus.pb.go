// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MessageSectionStatus.proto

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

type MessageSectionStatus int32

const (
	MessageSectionStatus_MESSAGE_SECTION_NONE   MessageSectionStatus = 0
	MessageSectionStatus_MESSAGE_SECTION_DOING  MessageSectionStatus = 1
	MessageSectionStatus_MESSAGE_SECTION_FINISH MessageSectionStatus = 2
	MessageSectionStatus_MESSAGE_SECTION_FROZEN MessageSectionStatus = 3
)

// Enum value maps for MessageSectionStatus.
var (
	MessageSectionStatus_name = map[int32]string{
		0: "MESSAGE_SECTION_NONE",
		1: "MESSAGE_SECTION_DOING",
		2: "MESSAGE_SECTION_FINISH",
		3: "MESSAGE_SECTION_FROZEN",
	}
	MessageSectionStatus_value = map[string]int32{
		"MESSAGE_SECTION_NONE":   0,
		"MESSAGE_SECTION_DOING":  1,
		"MESSAGE_SECTION_FINISH": 2,
		"MESSAGE_SECTION_FROZEN": 3,
	}
)

func (x MessageSectionStatus) Enum() *MessageSectionStatus {
	p := new(MessageSectionStatus)
	*p = x
	return p
}

func (x MessageSectionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageSectionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_MessageSectionStatus_proto_enumTypes[0].Descriptor()
}

func (MessageSectionStatus) Type() protoreflect.EnumType {
	return &file_MessageSectionStatus_proto_enumTypes[0]
}

func (x MessageSectionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageSectionStatus.Descriptor instead.
func (MessageSectionStatus) EnumDescriptor() ([]byte, []int) {
	return file_MessageSectionStatus_proto_rawDescGZIP(), []int{0}
}

var File_MessageSectionStatus_proto protoreflect.FileDescriptor

var file_MessageSectionStatus_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x83, 0x01, 0x0a,
	0x14, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x14, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45,
	0x5f, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12,
	0x19, 0x0a, 0x15, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x45, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x44, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x49,
	0x4e, 0x49, 0x53, 0x48, 0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47,
	0x45, 0x5f, 0x53, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x46, 0x52, 0x4f, 0x5a, 0x45, 0x4e,
	0x10, 0x03, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MessageSectionStatus_proto_rawDescOnce sync.Once
	file_MessageSectionStatus_proto_rawDescData = file_MessageSectionStatus_proto_rawDesc
)

func file_MessageSectionStatus_proto_rawDescGZIP() []byte {
	file_MessageSectionStatus_proto_rawDescOnce.Do(func() {
		file_MessageSectionStatus_proto_rawDescData = protoimpl.X.CompressGZIP(file_MessageSectionStatus_proto_rawDescData)
	})
	return file_MessageSectionStatus_proto_rawDescData
}

var file_MessageSectionStatus_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MessageSectionStatus_proto_goTypes = []interface{}{
	(MessageSectionStatus)(0), // 0: MessageSectionStatus
}
var file_MessageSectionStatus_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MessageSectionStatus_proto_init() }
func file_MessageSectionStatus_proto_init() {
	if File_MessageSectionStatus_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MessageSectionStatus_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MessageSectionStatus_proto_goTypes,
		DependencyIndexes: file_MessageSectionStatus_proto_depIdxs,
		EnumInfos:         file_MessageSectionStatus_proto_enumTypes,
	}.Build()
	File_MessageSectionStatus_proto = out.File
	file_MessageSectionStatus_proto_rawDesc = nil
	file_MessageSectionStatus_proto_goTypes = nil
	file_MessageSectionStatus_proto_depIdxs = nil
}
