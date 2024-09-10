// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: HealPoolInfoNotify.proto

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

type HealPoolInfoNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HealPoolInfo *HealPoolInfo `protobuf:"bytes,5,opt,name=heal_pool_info,json=healPoolInfo,proto3" json:"heal_pool_info,omitempty"`
}

func (x *HealPoolInfoNotify) Reset() {
	*x = HealPoolInfoNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HealPoolInfoNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealPoolInfoNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealPoolInfoNotify) ProtoMessage() {}

func (x *HealPoolInfoNotify) ProtoReflect() protoreflect.Message {
	mi := &file_HealPoolInfoNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealPoolInfoNotify.ProtoReflect.Descriptor instead.
func (*HealPoolInfoNotify) Descriptor() ([]byte, []int) {
	return file_HealPoolInfoNotify_proto_rawDescGZIP(), []int{0}
}

func (x *HealPoolInfoNotify) GetHealPoolInfo() *HealPoolInfo {
	if x != nil {
		return x.HealPoolInfo
	}
	return nil
}

var File_HealPoolInfoNotify_proto protoreflect.FileDescriptor

var file_HealPoolInfoNotify_proto_rawDesc = []byte{
	0x0a, 0x18, 0x48, 0x65, 0x61, 0x6c, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x48, 0x65, 0x61, 0x6c,
	0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x49,
	0x0a, 0x12, 0x48, 0x65, 0x61, 0x6c, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x12, 0x33, 0x0a, 0x0e, 0x68, 0x65, 0x61, 0x6c, 0x5f, 0x70, 0x6f, 0x6f,
	0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x48,
	0x65, 0x61, 0x6c, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x68, 0x65, 0x61,
	0x6c, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_HealPoolInfoNotify_proto_rawDescOnce sync.Once
	file_HealPoolInfoNotify_proto_rawDescData = file_HealPoolInfoNotify_proto_rawDesc
)

func file_HealPoolInfoNotify_proto_rawDescGZIP() []byte {
	file_HealPoolInfoNotify_proto_rawDescOnce.Do(func() {
		file_HealPoolInfoNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_HealPoolInfoNotify_proto_rawDescData)
	})
	return file_HealPoolInfoNotify_proto_rawDescData
}

var file_HealPoolInfoNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HealPoolInfoNotify_proto_goTypes = []interface{}{
	(*HealPoolInfoNotify)(nil), // 0: HealPoolInfoNotify
	(*HealPoolInfo)(nil),       // 1: HealPoolInfo
}
var file_HealPoolInfoNotify_proto_depIdxs = []int32{
	1, // 0: HealPoolInfoNotify.heal_pool_info:type_name -> HealPoolInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HealPoolInfoNotify_proto_init() }
func file_HealPoolInfoNotify_proto_init() {
	if File_HealPoolInfoNotify_proto != nil {
		return
	}
	file_HealPoolInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HealPoolInfoNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealPoolInfoNotify); i {
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
			RawDescriptor: file_HealPoolInfoNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HealPoolInfoNotify_proto_goTypes,
		DependencyIndexes: file_HealPoolInfoNotify_proto_depIdxs,
		MessageInfos:      file_HealPoolInfoNotify_proto_msgTypes,
	}.Build()
	File_HealPoolInfoNotify_proto = out.File
	file_HealPoolInfoNotify_proto_rawDesc = nil
	file_HealPoolInfoNotify_proto_goTypes = nil
	file_HealPoolInfoNotify_proto_depIdxs = nil
}
