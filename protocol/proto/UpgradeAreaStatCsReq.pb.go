// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: UpgradeAreaStatCsReq.proto

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

type UpgradeAreaStatCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MPLMAAHLHGD StatType `protobuf:"varint,5,opt,name=MPLMAAHLHGD,proto3,enum=StatType" json:"MPLMAAHLHGD,omitempty"`
	AreaId      uint32   `protobuf:"varint,10,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	Level       uint32   `protobuf:"varint,14,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *UpgradeAreaStatCsReq) Reset() {
	*x = UpgradeAreaStatCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UpgradeAreaStatCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpgradeAreaStatCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpgradeAreaStatCsReq) ProtoMessage() {}

func (x *UpgradeAreaStatCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_UpgradeAreaStatCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpgradeAreaStatCsReq.ProtoReflect.Descriptor instead.
func (*UpgradeAreaStatCsReq) Descriptor() ([]byte, []int) {
	return file_UpgradeAreaStatCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *UpgradeAreaStatCsReq) GetMPLMAAHLHGD() StatType {
	if x != nil {
		return x.MPLMAAHLHGD
	}
	return StatType_STAT_TYPE_NONE
}

func (x *UpgradeAreaStatCsReq) GetAreaId() uint32 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

func (x *UpgradeAreaStatCsReq) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

var File_UpgradeAreaStatCsReq_proto protoreflect.FileDescriptor

var file_UpgradeAreaStatCsReq_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x41, 0x72, 0x65, 0x61, 0x53, 0x74, 0x61,
	0x74, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x53, 0x74,
	0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x72, 0x0a, 0x14,
	0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x41, 0x72, 0x65, 0x61, 0x53, 0x74, 0x61, 0x74, 0x43,
	0x73, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x0b, 0x4d, 0x50, 0x4c, 0x4d, 0x41, 0x41, 0x48, 0x4c,
	0x48, 0x47, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x4d, 0x50, 0x4c, 0x4d, 0x41, 0x41, 0x48, 0x4c, 0x48, 0x47,
	0x44, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_UpgradeAreaStatCsReq_proto_rawDescOnce sync.Once
	file_UpgradeAreaStatCsReq_proto_rawDescData = file_UpgradeAreaStatCsReq_proto_rawDesc
)

func file_UpgradeAreaStatCsReq_proto_rawDescGZIP() []byte {
	file_UpgradeAreaStatCsReq_proto_rawDescOnce.Do(func() {
		file_UpgradeAreaStatCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_UpgradeAreaStatCsReq_proto_rawDescData)
	})
	return file_UpgradeAreaStatCsReq_proto_rawDescData
}

var file_UpgradeAreaStatCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_UpgradeAreaStatCsReq_proto_goTypes = []interface{}{
	(*UpgradeAreaStatCsReq)(nil), // 0: UpgradeAreaStatCsReq
	(StatType)(0),                // 1: StatType
}
var file_UpgradeAreaStatCsReq_proto_depIdxs = []int32{
	1, // 0: UpgradeAreaStatCsReq.MPLMAAHLHGD:type_name -> StatType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_UpgradeAreaStatCsReq_proto_init() }
func file_UpgradeAreaStatCsReq_proto_init() {
	if File_UpgradeAreaStatCsReq_proto != nil {
		return
	}
	file_StatType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_UpgradeAreaStatCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpgradeAreaStatCsReq); i {
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
			RawDescriptor: file_UpgradeAreaStatCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UpgradeAreaStatCsReq_proto_goTypes,
		DependencyIndexes: file_UpgradeAreaStatCsReq_proto_depIdxs,
		MessageInfos:      file_UpgradeAreaStatCsReq_proto_msgTypes,
	}.Build()
	File_UpgradeAreaStatCsReq_proto = out.File
	file_UpgradeAreaStatCsReq_proto_rawDesc = nil
	file_UpgradeAreaStatCsReq_proto_goTypes = nil
	file_UpgradeAreaStatCsReq_proto_depIdxs = nil
}
