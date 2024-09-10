// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueVirtualItem.proto

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

type RogueVirtualItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CKOJDJGLADK uint32 `protobuf:"varint,12,opt,name=CKOJDJGLADK,proto3" json:"CKOJDJGLADK,omitempty"`
	RogueMoney  uint32 `protobuf:"varint,7,opt,name=rogue_money,json=rogueMoney,proto3" json:"rogue_money,omitempty"`
	CFMBNFAAJCG uint32 `protobuf:"varint,5,opt,name=CFMBNFAAJCG,proto3" json:"CFMBNFAAJCG,omitempty"`
	DIHCFKEFJKL uint32 `protobuf:"varint,3,opt,name=DIHCFKEFJKL,proto3" json:"DIHCFKEFJKL,omitempty"`
}

func (x *RogueVirtualItem) Reset() {
	*x = RogueVirtualItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueVirtualItem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueVirtualItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueVirtualItem) ProtoMessage() {}

func (x *RogueVirtualItem) ProtoReflect() protoreflect.Message {
	mi := &file_RogueVirtualItem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueVirtualItem.ProtoReflect.Descriptor instead.
func (*RogueVirtualItem) Descriptor() ([]byte, []int) {
	return file_RogueVirtualItem_proto_rawDescGZIP(), []int{0}
}

func (x *RogueVirtualItem) GetCKOJDJGLADK() uint32 {
	if x != nil {
		return x.CKOJDJGLADK
	}
	return 0
}

func (x *RogueVirtualItem) GetRogueMoney() uint32 {
	if x != nil {
		return x.RogueMoney
	}
	return 0
}

func (x *RogueVirtualItem) GetCFMBNFAAJCG() uint32 {
	if x != nil {
		return x.CFMBNFAAJCG
	}
	return 0
}

func (x *RogueVirtualItem) GetDIHCFKEFJKL() uint32 {
	if x != nil {
		return x.DIHCFKEFJKL
	}
	return 0
}

var File_RogueVirtualItem_proto protoreflect.FileDescriptor

var file_RogueVirtualItem_proto_rawDesc = []byte{
	0x0a, 0x16, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74,
	0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x01, 0x0a, 0x10, 0x52, 0x6f, 0x67,
	0x75, 0x65, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x20, 0x0a,
	0x0b, 0x43, 0x4b, 0x4f, 0x4a, 0x44, 0x4a, 0x47, 0x4c, 0x41, 0x44, 0x4b, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x43, 0x4b, 0x4f, 0x4a, 0x44, 0x4a, 0x47, 0x4c, 0x41, 0x44, 0x4b, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x43, 0x46, 0x4d, 0x42, 0x4e, 0x46, 0x41, 0x41, 0x4a, 0x43, 0x47, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x46, 0x4d, 0x42, 0x4e, 0x46, 0x41, 0x41, 0x4a,
	0x43, 0x47, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x49, 0x48, 0x43, 0x46, 0x4b, 0x45, 0x46, 0x4a, 0x4b,
	0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x44, 0x49, 0x48, 0x43, 0x46, 0x4b, 0x45,
	0x46, 0x4a, 0x4b, 0x4c, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueVirtualItem_proto_rawDescOnce sync.Once
	file_RogueVirtualItem_proto_rawDescData = file_RogueVirtualItem_proto_rawDesc
)

func file_RogueVirtualItem_proto_rawDescGZIP() []byte {
	file_RogueVirtualItem_proto_rawDescOnce.Do(func() {
		file_RogueVirtualItem_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueVirtualItem_proto_rawDescData)
	})
	return file_RogueVirtualItem_proto_rawDescData
}

var file_RogueVirtualItem_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueVirtualItem_proto_goTypes = []interface{}{
	(*RogueVirtualItem)(nil), // 0: RogueVirtualItem
}
var file_RogueVirtualItem_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RogueVirtualItem_proto_init() }
func file_RogueVirtualItem_proto_init() {
	if File_RogueVirtualItem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RogueVirtualItem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueVirtualItem); i {
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
			RawDescriptor: file_RogueVirtualItem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueVirtualItem_proto_goTypes,
		DependencyIndexes: file_RogueVirtualItem_proto_depIdxs,
		MessageInfos:      file_RogueVirtualItem_proto_msgTypes,
	}.Build()
	File_RogueVirtualItem_proto = out.File
	file_RogueVirtualItem_proto_rawDesc = nil
	file_RogueVirtualItem_proto_goTypes = nil
	file_RogueVirtualItem_proto_depIdxs = nil
}
