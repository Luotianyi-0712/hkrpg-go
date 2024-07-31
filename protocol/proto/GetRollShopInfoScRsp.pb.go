// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetRollShopInfoScRsp.proto

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

type GetRollShopInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopGroupIdList []uint32 `protobuf:"varint,15,rep,packed,name=shop_group_id_list,json=shopGroupIdList,proto3" json:"shop_group_id_list,omitempty"`
	Retcode         uint32   `protobuf:"varint,14,opt,name=retcode,proto3" json:"retcode,omitempty"`
	GachaRandom     uint32   `protobuf:"varint,9,opt,name=gacha_random,json=gachaRandom,proto3" json:"gacha_random,omitempty"`
	RollShopId      uint32   `protobuf:"varint,7,opt,name=roll_shop_id,json=rollShopId,proto3" json:"roll_shop_id,omitempty"`
}

func (x *GetRollShopInfoScRsp) Reset() {
	*x = GetRollShopInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetRollShopInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRollShopInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRollShopInfoScRsp) ProtoMessage() {}

func (x *GetRollShopInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetRollShopInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRollShopInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetRollShopInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetRollShopInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetRollShopInfoScRsp) GetShopGroupIdList() []uint32 {
	if x != nil {
		return x.ShopGroupIdList
	}
	return nil
}

func (x *GetRollShopInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetRollShopInfoScRsp) GetGachaRandom() uint32 {
	if x != nil {
		return x.GachaRandom
	}
	return 0
}

func (x *GetRollShopInfoScRsp) GetRollShopId() uint32 {
	if x != nil {
		return x.RollShopId
	}
	return 0
}

var File_GetRollShopInfoScRsp_proto protoreflect.FileDescriptor

var file_GetRollShopInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x6c, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x52, 0x6f, 0x6c, 0x6c, 0x53, 0x68, 0x6f, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x2b, 0x0a, 0x12, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0f, 0x20, 0x03, 0x28,
	0x0d, 0x52, 0x0f, 0x73, 0x68, 0x6f, 0x70, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x67, 0x61, 0x63, 0x68, 0x61, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x67, 0x61, 0x63, 0x68, 0x61, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x12,
	0x20, 0x0a, 0x0c, 0x72, 0x6f, 0x6c, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72, 0x6f, 0x6c, 0x6c, 0x53, 0x68, 0x6f, 0x70, 0x49,
	0x64, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_GetRollShopInfoScRsp_proto_rawDescOnce sync.Once
	file_GetRollShopInfoScRsp_proto_rawDescData = file_GetRollShopInfoScRsp_proto_rawDesc
)

func file_GetRollShopInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetRollShopInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetRollShopInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetRollShopInfoScRsp_proto_rawDescData)
	})
	return file_GetRollShopInfoScRsp_proto_rawDescData
}

var file_GetRollShopInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetRollShopInfoScRsp_proto_goTypes = []interface{}{
	(*GetRollShopInfoScRsp)(nil), // 0: GetRollShopInfoScRsp
}
var file_GetRollShopInfoScRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetRollShopInfoScRsp_proto_init() }
func file_GetRollShopInfoScRsp_proto_init() {
	if File_GetRollShopInfoScRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetRollShopInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRollShopInfoScRsp); i {
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
			RawDescriptor: file_GetRollShopInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetRollShopInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetRollShopInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetRollShopInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetRollShopInfoScRsp_proto = out.File
	file_GetRollShopInfoScRsp_proto_rawDesc = nil
	file_GetRollShopInfoScRsp_proto_goTypes = nil
	file_GetRollShopInfoScRsp_proto_depIdxs = nil
}
