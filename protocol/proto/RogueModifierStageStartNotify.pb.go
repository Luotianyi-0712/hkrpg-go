// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueModifierStageStartNotify.proto

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

type RogueModifierStageStartNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModifierSourceType RogueModifierSourceType `protobuf:"varint,10,opt,name=modifier_source_type,json=modifierSourceType,proto3,enum=RogueModifierSourceType" json:"modifier_source_type,omitempty"`
}

func (x *RogueModifierStageStartNotify) Reset() {
	*x = RogueModifierStageStartNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueModifierStageStartNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueModifierStageStartNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueModifierStageStartNotify) ProtoMessage() {}

func (x *RogueModifierStageStartNotify) ProtoReflect() protoreflect.Message {
	mi := &file_RogueModifierStageStartNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueModifierStageStartNotify.ProtoReflect.Descriptor instead.
func (*RogueModifierStageStartNotify) Descriptor() ([]byte, []int) {
	return file_RogueModifierStageStartNotify_proto_rawDescGZIP(), []int{0}
}

func (x *RogueModifierStageStartNotify) GetModifierSourceType() RogueModifierSourceType {
	if x != nil {
		return x.ModifierSourceType
	}
	return RogueModifierSourceType_ROGUE_MODIFIER_SOURCE_NONE
}

var File_RogueModifierStageStartNotify_proto protoreflect.FileDescriptor

var file_RogueModifierStageStartNotify_proto_rawDesc = []byte{
	0x0a, 0x23, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x1d, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x74, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x4a, 0x0a, 0x14, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x12, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_RogueModifierStageStartNotify_proto_rawDescOnce sync.Once
	file_RogueModifierStageStartNotify_proto_rawDescData = file_RogueModifierStageStartNotify_proto_rawDesc
)

func file_RogueModifierStageStartNotify_proto_rawDescGZIP() []byte {
	file_RogueModifierStageStartNotify_proto_rawDescOnce.Do(func() {
		file_RogueModifierStageStartNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueModifierStageStartNotify_proto_rawDescData)
	})
	return file_RogueModifierStageStartNotify_proto_rawDescData
}

var file_RogueModifierStageStartNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueModifierStageStartNotify_proto_goTypes = []interface{}{
	(*RogueModifierStageStartNotify)(nil), // 0: RogueModifierStageStartNotify
	(RogueModifierSourceType)(0),          // 1: RogueModifierSourceType
}
var file_RogueModifierStageStartNotify_proto_depIdxs = []int32{
	1, // 0: RogueModifierStageStartNotify.modifier_source_type:type_name -> RogueModifierSourceType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_RogueModifierStageStartNotify_proto_init() }
func file_RogueModifierStageStartNotify_proto_init() {
	if File_RogueModifierStageStartNotify_proto != nil {
		return
	}
	file_RogueModifierSourceType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueModifierStageStartNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueModifierStageStartNotify); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_RogueModifierStageStartNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueModifierStageStartNotify_proto_goTypes,
		DependencyIndexes: file_RogueModifierStageStartNotify_proto_depIdxs,
		MessageInfos:      file_RogueModifierStageStartNotify_proto_msgTypes,
	}.Build()
	File_RogueModifierStageStartNotify_proto = out.File
	file_RogueModifierStageStartNotify_proto_rawDesc = nil
	file_RogueModifierStageStartNotify_proto_goTypes = nil
	file_RogueModifierStageStartNotify_proto_depIdxs = nil
}
