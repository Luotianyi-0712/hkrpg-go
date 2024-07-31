// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FinishSectionIdScRsp.proto

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

type FinishSectionIdScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SectionId uint32    `protobuf:"varint,8,opt,name=section_id,json=sectionId,proto3" json:"section_id,omitempty"`
	Retcode   uint32    `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
	Reward    *ItemList `protobuf:"bytes,9,opt,name=reward,proto3" json:"reward,omitempty"`
}

func (x *FinishSectionIdScRsp) Reset() {
	*x = FinishSectionIdScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FinishSectionIdScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishSectionIdScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishSectionIdScRsp) ProtoMessage() {}

func (x *FinishSectionIdScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_FinishSectionIdScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishSectionIdScRsp.ProtoReflect.Descriptor instead.
func (*FinishSectionIdScRsp) Descriptor() ([]byte, []int) {
	return file_FinishSectionIdScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *FinishSectionIdScRsp) GetSectionId() uint32 {
	if x != nil {
		return x.SectionId
	}
	return 0
}

func (x *FinishSectionIdScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *FinishSectionIdScRsp) GetReward() *ItemList {
	if x != nil {
		return x.Reward
	}
	return nil
}

var File_FinishSectionIdScRsp_proto protoreflect.FileDescriptor

var file_FinishSectionIdScRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74,
	0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x14,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a,
	0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e,
	0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_FinishSectionIdScRsp_proto_rawDescOnce sync.Once
	file_FinishSectionIdScRsp_proto_rawDescData = file_FinishSectionIdScRsp_proto_rawDesc
)

func file_FinishSectionIdScRsp_proto_rawDescGZIP() []byte {
	file_FinishSectionIdScRsp_proto_rawDescOnce.Do(func() {
		file_FinishSectionIdScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_FinishSectionIdScRsp_proto_rawDescData)
	})
	return file_FinishSectionIdScRsp_proto_rawDescData
}

var file_FinishSectionIdScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FinishSectionIdScRsp_proto_goTypes = []interface{}{
	(*FinishSectionIdScRsp)(nil), // 0: FinishSectionIdScRsp
	(*ItemList)(nil),             // 1: ItemList
}
var file_FinishSectionIdScRsp_proto_depIdxs = []int32{
	1, // 0: FinishSectionIdScRsp.reward:type_name -> ItemList
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FinishSectionIdScRsp_proto_init() }
func file_FinishSectionIdScRsp_proto_init() {
	if File_FinishSectionIdScRsp_proto != nil {
		return
	}
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FinishSectionIdScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishSectionIdScRsp); i {
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
			RawDescriptor: file_FinishSectionIdScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FinishSectionIdScRsp_proto_goTypes,
		DependencyIndexes: file_FinishSectionIdScRsp_proto_depIdxs,
		MessageInfos:      file_FinishSectionIdScRsp_proto_msgTypes,
	}.Build()
	File_FinishSectionIdScRsp_proto = out.File
	file_FinishSectionIdScRsp_proto_rawDesc = nil
	file_FinishSectionIdScRsp_proto_goTypes = nil
	file_FinishSectionIdScRsp_proto_depIdxs = nil
}
