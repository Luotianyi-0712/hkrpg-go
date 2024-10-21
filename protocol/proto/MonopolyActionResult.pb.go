// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MonopolyActionResult.proto

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

type MonopolyActionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClickCellIdFieldNumber   uint32       `protobuf:"varint,10,opt,name=ClickCellIdFieldNumber,proto3" json:"ClickCellIdFieldNumber,omitempty"`
	TriggerCellIdFieldNumber uint32       `protobuf:"varint,4,opt,name=TriggerCellIdFieldNumber,proto3" json:"TriggerCellIdFieldNumber,omitempty"`
	DetailFieldNumber        *BAKGCCKJDLK `protobuf:"bytes,11,opt,name=DetailFieldNumber,proto3" json:"DetailFieldNumber,omitempty"`
	SourceType               CIAOIKEANEA  `protobuf:"varint,12,opt,name=source_type,json=sourceType,proto3,enum=CIAOIKEANEA" json:"source_type,omitempty"`
	Effecttype               uint32       `protobuf:"varint,9,opt,name=Effecttype,proto3" json:"Effecttype,omitempty"`
	TriggerMapIdFieldNumber  uint32       `protobuf:"varint,5,opt,name=TriggerMapIdFieldNumber,proto3" json:"TriggerMapIdFieldNumber,omitempty"`
	ClickMapIdFieldNumber    uint32       `protobuf:"varint,6,opt,name=ClickMapIdFieldNumber,proto3" json:"ClickMapIdFieldNumber,omitempty"`
}

func (x *MonopolyActionResult) Reset() {
	*x = MonopolyActionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MonopolyActionResult_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonopolyActionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonopolyActionResult) ProtoMessage() {}

func (x *MonopolyActionResult) ProtoReflect() protoreflect.Message {
	mi := &file_MonopolyActionResult_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonopolyActionResult.ProtoReflect.Descriptor instead.
func (*MonopolyActionResult) Descriptor() ([]byte, []int) {
	return file_MonopolyActionResult_proto_rawDescGZIP(), []int{0}
}

func (x *MonopolyActionResult) GetClickCellIdFieldNumber() uint32 {
	if x != nil {
		return x.ClickCellIdFieldNumber
	}
	return 0
}

func (x *MonopolyActionResult) GetTriggerCellIdFieldNumber() uint32 {
	if x != nil {
		return x.TriggerCellIdFieldNumber
	}
	return 0
}

func (x *MonopolyActionResult) GetDetailFieldNumber() *BAKGCCKJDLK {
	if x != nil {
		return x.DetailFieldNumber
	}
	return nil
}

func (x *MonopolyActionResult) GetSourceType() CIAOIKEANEA {
	if x != nil {
		return x.SourceType
	}
	return CIAOIKEANEA_MONOPOLY_ACTION_RESULT_SOURCE_TYPE_NONE
}

func (x *MonopolyActionResult) GetEffecttype() uint32 {
	if x != nil {
		return x.Effecttype
	}
	return 0
}

func (x *MonopolyActionResult) GetTriggerMapIdFieldNumber() uint32 {
	if x != nil {
		return x.TriggerMapIdFieldNumber
	}
	return 0
}

func (x *MonopolyActionResult) GetClickMapIdFieldNumber() uint32 {
	if x != nil {
		return x.ClickMapIdFieldNumber
	}
	return 0
}

var File_MonopolyActionResult_proto protoreflect.FileDescriptor

var file_MonopolyActionResult_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x43, 0x49,
	0x41, 0x4f, 0x49, 0x4b, 0x45, 0x41, 0x4e, 0x45, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x42, 0x41, 0x4b, 0x47, 0x43, 0x43, 0x4b, 0x4a, 0x44, 0x4c, 0x4b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x85, 0x03, 0x0a, 0x14, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x36, 0x0a, 0x16, 0x43,
	0x6c, 0x69, 0x63, 0x6b, 0x43, 0x65, 0x6c, 0x6c, 0x49, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x43, 0x6c, 0x69,
	0x63, 0x6b, 0x43, 0x65, 0x6c, 0x6c, 0x49, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x18, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x43, 0x65,
	0x6c, 0x6c, 0x49, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x18, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x43, 0x65,
	0x6c, 0x6c, 0x49, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x3a, 0x0a, 0x11, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x41, 0x4b,
	0x47, 0x43, 0x43, 0x4b, 0x4a, 0x44, 0x4c, 0x4b, 0x52, 0x11, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x0b, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0c, 0x2e, 0x43, 0x49, 0x41, 0x4f, 0x49, 0x4b, 0x45, 0x41, 0x4e, 0x45, 0x41, 0x52, 0x0a,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x45, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x74, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x54, 0x72,
	0x69, 0x67, 0x67, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x49, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x17, 0x54, 0x72, 0x69,
	0x67, 0x67, 0x65, 0x72, 0x4d, 0x61, 0x70, 0x49, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x15, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x4d, 0x61, 0x70,
	0x49, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x15, 0x43, 0x6c, 0x69, 0x63, 0x6b, 0x4d, 0x61, 0x70, 0x49, 0x64, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_MonopolyActionResult_proto_rawDescOnce sync.Once
	file_MonopolyActionResult_proto_rawDescData = file_MonopolyActionResult_proto_rawDesc
)

func file_MonopolyActionResult_proto_rawDescGZIP() []byte {
	file_MonopolyActionResult_proto_rawDescOnce.Do(func() {
		file_MonopolyActionResult_proto_rawDescData = protoimpl.X.CompressGZIP(file_MonopolyActionResult_proto_rawDescData)
	})
	return file_MonopolyActionResult_proto_rawDescData
}

var file_MonopolyActionResult_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MonopolyActionResult_proto_goTypes = []interface{}{
	(*MonopolyActionResult)(nil), // 0: MonopolyActionResult
	(*BAKGCCKJDLK)(nil),          // 1: BAKGCCKJDLK
	(CIAOIKEANEA)(0),             // 2: CIAOIKEANEA
}
var file_MonopolyActionResult_proto_depIdxs = []int32{
	1, // 0: MonopolyActionResult.DetailFieldNumber:type_name -> BAKGCCKJDLK
	2, // 1: MonopolyActionResult.source_type:type_name -> CIAOIKEANEA
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_MonopolyActionResult_proto_init() }
func file_MonopolyActionResult_proto_init() {
	if File_MonopolyActionResult_proto != nil {
		return
	}
	file_CIAOIKEANEA_proto_init()
	file_BAKGCCKJDLK_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MonopolyActionResult_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonopolyActionResult); i {
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
			RawDescriptor: file_MonopolyActionResult_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MonopolyActionResult_proto_goTypes,
		DependencyIndexes: file_MonopolyActionResult_proto_depIdxs,
		MessageInfos:      file_MonopolyActionResult_proto_msgTypes,
	}.Build()
	File_MonopolyActionResult_proto = out.File
	file_MonopolyActionResult_proto_rawDesc = nil
	file_MonopolyActionResult_proto_goTypes = nil
	file_MonopolyActionResult_proto_depIdxs = nil
}
