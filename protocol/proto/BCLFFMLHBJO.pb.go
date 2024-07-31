// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: BCLFFMLHBJO.proto

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

type BCLFFMLHBJO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CostData    *ItemCostData `protobuf:"bytes,12,opt,name=cost_data,json=costData,proto3" json:"cost_data,omitempty"`
	MiracleId   uint32        `protobuf:"varint,6,opt,name=miracle_id,json=miracleId,proto3" json:"miracle_id,omitempty"`
	BCLEGGBLDIL bool          `protobuf:"varint,13,opt,name=BCLEGGBLDIL,proto3" json:"BCLEGGBLDIL,omitempty"`
	MMLFGHDAMLE *ItemCostData `protobuf:"bytes,1,opt,name=MMLFGHDAMLE,proto3" json:"MMLFGHDAMLE,omitempty"`
	MMKACGNCMEK bool          `protobuf:"varint,4,opt,name=MMKACGNCMEK,proto3" json:"MMKACGNCMEK,omitempty"`
}

func (x *BCLFFMLHBJO) Reset() {
	*x = BCLFFMLHBJO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BCLFFMLHBJO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BCLFFMLHBJO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BCLFFMLHBJO) ProtoMessage() {}

func (x *BCLFFMLHBJO) ProtoReflect() protoreflect.Message {
	mi := &file_BCLFFMLHBJO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BCLFFMLHBJO.ProtoReflect.Descriptor instead.
func (*BCLFFMLHBJO) Descriptor() ([]byte, []int) {
	return file_BCLFFMLHBJO_proto_rawDescGZIP(), []int{0}
}

func (x *BCLFFMLHBJO) GetCostData() *ItemCostData {
	if x != nil {
		return x.CostData
	}
	return nil
}

func (x *BCLFFMLHBJO) GetMiracleId() uint32 {
	if x != nil {
		return x.MiracleId
	}
	return 0
}

func (x *BCLFFMLHBJO) GetBCLEGGBLDIL() bool {
	if x != nil {
		return x.BCLEGGBLDIL
	}
	return false
}

func (x *BCLFFMLHBJO) GetMMLFGHDAMLE() *ItemCostData {
	if x != nil {
		return x.MMLFGHDAMLE
	}
	return nil
}

func (x *BCLFFMLHBJO) GetMMKACGNCMEK() bool {
	if x != nil {
		return x.MMKACGNCMEK
	}
	return false
}

var File_BCLFFMLHBJO_proto protoreflect.FileDescriptor

var file_BCLFFMLHBJO_proto_rawDesc = []byte{
	0x0a, 0x11, 0x42, 0x43, 0x4c, 0x46, 0x46, 0x4d, 0x4c, 0x48, 0x42, 0x4a, 0x4f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x49, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x0b, 0x42, 0x43, 0x4c, 0x46,
	0x46, 0x4d, 0x4c, 0x48, 0x42, 0x4a, 0x4f, 0x12, 0x2a, 0x0a, 0x09, 0x63, 0x6f, 0x73, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x43, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x63, 0x6f, 0x73, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x43, 0x4c, 0x45, 0x47, 0x47, 0x42, 0x4c, 0x44, 0x49,
	0x4c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x42, 0x43, 0x4c, 0x45, 0x47, 0x47, 0x42,
	0x4c, 0x44, 0x49, 0x4c, 0x12, 0x2f, 0x0a, 0x0b, 0x4d, 0x4d, 0x4c, 0x46, 0x47, 0x48, 0x44, 0x41,
	0x4d, 0x4c, 0x45, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x49, 0x74, 0x65, 0x6d,
	0x43, 0x6f, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x4d, 0x4d, 0x4c, 0x46, 0x47, 0x48,
	0x44, 0x41, 0x4d, 0x4c, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x4d, 0x4b, 0x41, 0x43, 0x47, 0x4e,
	0x43, 0x4d, 0x45, 0x4b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4d, 0x4d, 0x4b, 0x41,
	0x43, 0x47, 0x4e, 0x43, 0x4d, 0x45, 0x4b, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BCLFFMLHBJO_proto_rawDescOnce sync.Once
	file_BCLFFMLHBJO_proto_rawDescData = file_BCLFFMLHBJO_proto_rawDesc
)

func file_BCLFFMLHBJO_proto_rawDescGZIP() []byte {
	file_BCLFFMLHBJO_proto_rawDescOnce.Do(func() {
		file_BCLFFMLHBJO_proto_rawDescData = protoimpl.X.CompressGZIP(file_BCLFFMLHBJO_proto_rawDescData)
	})
	return file_BCLFFMLHBJO_proto_rawDescData
}

var file_BCLFFMLHBJO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BCLFFMLHBJO_proto_goTypes = []interface{}{
	(*BCLFFMLHBJO)(nil),  // 0: BCLFFMLHBJO
	(*ItemCostData)(nil), // 1: ItemCostData
}
var file_BCLFFMLHBJO_proto_depIdxs = []int32{
	1, // 0: BCLFFMLHBJO.cost_data:type_name -> ItemCostData
	1, // 1: BCLFFMLHBJO.MMLFGHDAMLE:type_name -> ItemCostData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_BCLFFMLHBJO_proto_init() }
func file_BCLFFMLHBJO_proto_init() {
	if File_BCLFFMLHBJO_proto != nil {
		return
	}
	file_ItemCostData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BCLFFMLHBJO_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BCLFFMLHBJO); i {
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
			RawDescriptor: file_BCLFFMLHBJO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BCLFFMLHBJO_proto_goTypes,
		DependencyIndexes: file_BCLFFMLHBJO_proto_depIdxs,
		MessageInfos:      file_BCLFFMLHBJO_proto_msgTypes,
	}.Build()
	File_BCLFFMLHBJO_proto = out.File
	file_BCLFFMLHBJO_proto_rawDesc = nil
	file_BCLFFMLHBJO_proto_goTypes = nil
	file_BCLFFMLHBJO_proto_depIdxs = nil
}
