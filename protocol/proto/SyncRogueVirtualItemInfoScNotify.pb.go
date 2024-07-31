// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SyncRogueVirtualItemInfoScNotify.proto

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

type SyncRogueVirtualItemInfoScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RogueVirtualItemInfo *RogueVirtualItemInfo `protobuf:"bytes,13,opt,name=rogue_virtual_item_info,json=rogueVirtualItemInfo,proto3" json:"rogue_virtual_item_info,omitempty"`
}

func (x *SyncRogueVirtualItemInfoScNotify) Reset() {
	*x = SyncRogueVirtualItemInfoScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SyncRogueVirtualItemInfoScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncRogueVirtualItemInfoScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncRogueVirtualItemInfoScNotify) ProtoMessage() {}

func (x *SyncRogueVirtualItemInfoScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SyncRogueVirtualItemInfoScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncRogueVirtualItemInfoScNotify.ProtoReflect.Descriptor instead.
func (*SyncRogueVirtualItemInfoScNotify) Descriptor() ([]byte, []int) {
	return file_SyncRogueVirtualItemInfoScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SyncRogueVirtualItemInfoScNotify) GetRogueVirtualItemInfo() *RogueVirtualItemInfo {
	if x != nil {
		return x.RogueVirtualItemInfo
	}
	return nil
}

var File_SyncRogueVirtualItemInfoScNotify_proto protoreflect.FileDescriptor

var file_SyncRogueVirtualItemInfoScNotify_proto_rawDesc = []byte{
	0x0a, 0x26, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x56, 0x69, 0x72, 0x74, 0x75,
	0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x56,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x20, 0x53, 0x79, 0x6e, 0x63, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f,
	0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x4c, 0x0a, 0x17, 0x72, 0x6f, 0x67, 0x75,
	0x65, 0x5f, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x14, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74,
	0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SyncRogueVirtualItemInfoScNotify_proto_rawDescOnce sync.Once
	file_SyncRogueVirtualItemInfoScNotify_proto_rawDescData = file_SyncRogueVirtualItemInfoScNotify_proto_rawDesc
)

func file_SyncRogueVirtualItemInfoScNotify_proto_rawDescGZIP() []byte {
	file_SyncRogueVirtualItemInfoScNotify_proto_rawDescOnce.Do(func() {
		file_SyncRogueVirtualItemInfoScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SyncRogueVirtualItemInfoScNotify_proto_rawDescData)
	})
	return file_SyncRogueVirtualItemInfoScNotify_proto_rawDescData
}

var file_SyncRogueVirtualItemInfoScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SyncRogueVirtualItemInfoScNotify_proto_goTypes = []interface{}{
	(*SyncRogueVirtualItemInfoScNotify)(nil), // 0: SyncRogueVirtualItemInfoScNotify
	(*RogueVirtualItemInfo)(nil),             // 1: RogueVirtualItemInfo
}
var file_SyncRogueVirtualItemInfoScNotify_proto_depIdxs = []int32{
	1, // 0: SyncRogueVirtualItemInfoScNotify.rogue_virtual_item_info:type_name -> RogueVirtualItemInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_SyncRogueVirtualItemInfoScNotify_proto_init() }
func file_SyncRogueVirtualItemInfoScNotify_proto_init() {
	if File_SyncRogueVirtualItemInfoScNotify_proto != nil {
		return
	}
	file_RogueVirtualItemInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SyncRogueVirtualItemInfoScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncRogueVirtualItemInfoScNotify); i {
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
			RawDescriptor: file_SyncRogueVirtualItemInfoScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SyncRogueVirtualItemInfoScNotify_proto_goTypes,
		DependencyIndexes: file_SyncRogueVirtualItemInfoScNotify_proto_depIdxs,
		MessageInfos:      file_SyncRogueVirtualItemInfoScNotify_proto_msgTypes,
	}.Build()
	File_SyncRogueVirtualItemInfoScNotify_proto = out.File
	file_SyncRogueVirtualItemInfoScNotify_proto_rawDesc = nil
	file_SyncRogueVirtualItemInfoScNotify_proto_goTypes = nil
	file_SyncRogueVirtualItemInfoScNotify_proto_depIdxs = nil
}
