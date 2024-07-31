// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MonopolyBuyGoodsCsReq.proto

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

type MonopolyBuyGoodsCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId  uint32 `protobuf:"varint,5,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
	GoodsId uint32 `protobuf:"varint,15,opt,name=goods_id,json=goodsId,proto3" json:"goods_id,omitempty"`
}

func (x *MonopolyBuyGoodsCsReq) Reset() {
	*x = MonopolyBuyGoodsCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MonopolyBuyGoodsCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonopolyBuyGoodsCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonopolyBuyGoodsCsReq) ProtoMessage() {}

func (x *MonopolyBuyGoodsCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_MonopolyBuyGoodsCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonopolyBuyGoodsCsReq.ProtoReflect.Descriptor instead.
func (*MonopolyBuyGoodsCsReq) Descriptor() ([]byte, []int) {
	return file_MonopolyBuyGoodsCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *MonopolyBuyGoodsCsReq) GetShopId() uint32 {
	if x != nil {
		return x.ShopId
	}
	return 0
}

func (x *MonopolyBuyGoodsCsReq) GetGoodsId() uint32 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

var File_MonopolyBuyGoodsCsReq_proto protoreflect.FileDescriptor

var file_MonopolyBuyGoodsCsReq_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x42, 0x75, 0x79, 0x47, 0x6f, 0x6f,
	0x64, 0x73, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a,
	0x15, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x42, 0x75, 0x79, 0x47, 0x6f, 0x6f, 0x64,
	0x73, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MonopolyBuyGoodsCsReq_proto_rawDescOnce sync.Once
	file_MonopolyBuyGoodsCsReq_proto_rawDescData = file_MonopolyBuyGoodsCsReq_proto_rawDesc
)

func file_MonopolyBuyGoodsCsReq_proto_rawDescGZIP() []byte {
	file_MonopolyBuyGoodsCsReq_proto_rawDescOnce.Do(func() {
		file_MonopolyBuyGoodsCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_MonopolyBuyGoodsCsReq_proto_rawDescData)
	})
	return file_MonopolyBuyGoodsCsReq_proto_rawDescData
}

var file_MonopolyBuyGoodsCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MonopolyBuyGoodsCsReq_proto_goTypes = []interface{}{
	(*MonopolyBuyGoodsCsReq)(nil), // 0: MonopolyBuyGoodsCsReq
}
var file_MonopolyBuyGoodsCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MonopolyBuyGoodsCsReq_proto_init() }
func file_MonopolyBuyGoodsCsReq_proto_init() {
	if File_MonopolyBuyGoodsCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MonopolyBuyGoodsCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonopolyBuyGoodsCsReq); i {
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
			RawDescriptor: file_MonopolyBuyGoodsCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MonopolyBuyGoodsCsReq_proto_goTypes,
		DependencyIndexes: file_MonopolyBuyGoodsCsReq_proto_depIdxs,
		MessageInfos:      file_MonopolyBuyGoodsCsReq_proto_msgTypes,
	}.Build()
	File_MonopolyBuyGoodsCsReq_proto = out.File
	file_MonopolyBuyGoodsCsReq_proto_rawDesc = nil
	file_MonopolyBuyGoodsCsReq_proto_goTypes = nil
	file_MonopolyBuyGoodsCsReq_proto_depIdxs = nil
}
