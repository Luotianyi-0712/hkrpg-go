// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GachaItem.proto

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

type GachaItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GachaItem        *Item     `protobuf:"bytes,3,opt,name=gacha_item,json=gachaItem,proto3" json:"gacha_item,omitempty"`
	IsNew            bool      `protobuf:"varint,9,opt,name=is_new,json=isNew,proto3" json:"is_new,omitempty"`
	TokenItem        *ItemList `protobuf:"bytes,12,opt,name=token_item,json=tokenItem,proto3" json:"token_item,omitempty"`
	TransferItemList *ItemList `protobuf:"bytes,13,opt,name=transfer_item_list,json=transferItemList,proto3" json:"transfer_item_list,omitempty"`
}

func (x *GachaItem) Reset() {
	*x = GachaItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GachaItem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GachaItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GachaItem) ProtoMessage() {}

func (x *GachaItem) ProtoReflect() protoreflect.Message {
	mi := &file_GachaItem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GachaItem.ProtoReflect.Descriptor instead.
func (*GachaItem) Descriptor() ([]byte, []int) {
	return file_GachaItem_proto_rawDescGZIP(), []int{0}
}

func (x *GachaItem) GetGachaItem() *Item {
	if x != nil {
		return x.GachaItem
	}
	return nil
}

func (x *GachaItem) GetIsNew() bool {
	if x != nil {
		return x.IsNew
	}
	return false
}

func (x *GachaItem) GetTokenItem() *ItemList {
	if x != nil {
		return x.TokenItem
	}
	return nil
}

func (x *GachaItem) GetTransferItemList() *ItemList {
	if x != nil {
		return x.TransferItemList
	}
	return nil
}

var File_GachaItem_proto protoreflect.FileDescriptor

var file_GachaItem_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x47, 0x61, 0x63, 0x68, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0a, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x01,
	0x0a, 0x09, 0x47, 0x61, 0x63, 0x68, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x24, 0x0a, 0x0a, 0x67,
	0x61, 0x63, 0x68, 0x61, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x05, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x67, 0x61, 0x63, 0x68, 0x61, 0x49, 0x74, 0x65,
	0x6d, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x77, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x69, 0x73, 0x4e, 0x65, 0x77, 0x12, 0x28, 0x0a, 0x0a, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49,
	0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x37, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x28, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GachaItem_proto_rawDescOnce sync.Once
	file_GachaItem_proto_rawDescData = file_GachaItem_proto_rawDesc
)

func file_GachaItem_proto_rawDescGZIP() []byte {
	file_GachaItem_proto_rawDescOnce.Do(func() {
		file_GachaItem_proto_rawDescData = protoimpl.X.CompressGZIP(file_GachaItem_proto_rawDescData)
	})
	return file_GachaItem_proto_rawDescData
}

var file_GachaItem_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GachaItem_proto_goTypes = []interface{}{
	(*GachaItem)(nil), // 0: GachaItem
	(*Item)(nil),      // 1: Item
	(*ItemList)(nil),  // 2: ItemList
}
var file_GachaItem_proto_depIdxs = []int32{
	1, // 0: GachaItem.gacha_item:type_name -> Item
	2, // 1: GachaItem.token_item:type_name -> ItemList
	2, // 2: GachaItem.transfer_item_list:type_name -> ItemList
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_GachaItem_proto_init() }
func file_GachaItem_proto_init() {
	if File_GachaItem_proto != nil {
		return
	}
	file_ItemList_proto_init()
	file_Item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GachaItem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GachaItem); i {
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
			RawDescriptor: file_GachaItem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GachaItem_proto_goTypes,
		DependencyIndexes: file_GachaItem_proto_depIdxs,
		MessageInfos:      file_GachaItem_proto_msgTypes,
	}.Build()
	File_GachaItem_proto = out.File
	file_GachaItem_proto_rawDesc = nil
	file_GachaItem_proto_goTypes = nil
	file_GachaItem_proto_depIdxs = nil
}
