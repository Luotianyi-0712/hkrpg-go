// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TakeRogueMiracleHandbookRewardCsReq.proto

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

type TakeRogueMiracleHandbookRewardCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HDANFHLJFEJ []uint32 `protobuf:"varint,12,rep,packed,name=HDANFHLJFEJ,proto3" json:"HDANFHLJFEJ,omitempty"`
}

func (x *TakeRogueMiracleHandbookRewardCsReq) Reset() {
	*x = TakeRogueMiracleHandbookRewardCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TakeRogueMiracleHandbookRewardCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeRogueMiracleHandbookRewardCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeRogueMiracleHandbookRewardCsReq) ProtoMessage() {}

func (x *TakeRogueMiracleHandbookRewardCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_TakeRogueMiracleHandbookRewardCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeRogueMiracleHandbookRewardCsReq.ProtoReflect.Descriptor instead.
func (*TakeRogueMiracleHandbookRewardCsReq) Descriptor() ([]byte, []int) {
	return file_TakeRogueMiracleHandbookRewardCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *TakeRogueMiracleHandbookRewardCsReq) GetHDANFHLJFEJ() []uint32 {
	if x != nil {
		return x.HDANFHLJFEJ
	}
	return nil
}

var File_TakeRogueMiracleHandbookRewardCsReq_proto protoreflect.FileDescriptor

var file_TakeRogueMiracleHandbookRewardCsReq_proto_rawDesc = []byte{
	0x0a, 0x29, 0x54, 0x61, 0x6b, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x23, 0x54,
	0x61, 0x6b, 0x65, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x48,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x43, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x44, 0x41, 0x4e, 0x46, 0x48, 0x4c, 0x4a, 0x46, 0x45,
	0x4a, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x44, 0x41, 0x4e, 0x46, 0x48, 0x4c,
	0x4a, 0x46, 0x45, 0x4a, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TakeRogueMiracleHandbookRewardCsReq_proto_rawDescOnce sync.Once
	file_TakeRogueMiracleHandbookRewardCsReq_proto_rawDescData = file_TakeRogueMiracleHandbookRewardCsReq_proto_rawDesc
)

func file_TakeRogueMiracleHandbookRewardCsReq_proto_rawDescGZIP() []byte {
	file_TakeRogueMiracleHandbookRewardCsReq_proto_rawDescOnce.Do(func() {
		file_TakeRogueMiracleHandbookRewardCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_TakeRogueMiracleHandbookRewardCsReq_proto_rawDescData)
	})
	return file_TakeRogueMiracleHandbookRewardCsReq_proto_rawDescData
}

var file_TakeRogueMiracleHandbookRewardCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TakeRogueMiracleHandbookRewardCsReq_proto_goTypes = []interface{}{
	(*TakeRogueMiracleHandbookRewardCsReq)(nil), // 0: TakeRogueMiracleHandbookRewardCsReq
}
var file_TakeRogueMiracleHandbookRewardCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TakeRogueMiracleHandbookRewardCsReq_proto_init() }
func file_TakeRogueMiracleHandbookRewardCsReq_proto_init() {
	if File_TakeRogueMiracleHandbookRewardCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TakeRogueMiracleHandbookRewardCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeRogueMiracleHandbookRewardCsReq); i {
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
			RawDescriptor: file_TakeRogueMiracleHandbookRewardCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TakeRogueMiracleHandbookRewardCsReq_proto_goTypes,
		DependencyIndexes: file_TakeRogueMiracleHandbookRewardCsReq_proto_depIdxs,
		MessageInfos:      file_TakeRogueMiracleHandbookRewardCsReq_proto_msgTypes,
	}.Build()
	File_TakeRogueMiracleHandbookRewardCsReq_proto = out.File
	file_TakeRogueMiracleHandbookRewardCsReq_proto_rawDesc = nil
	file_TakeRogueMiracleHandbookRewardCsReq_proto_goTypes = nil
	file_TakeRogueMiracleHandbookRewardCsReq_proto_depIdxs = nil
}
