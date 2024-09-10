// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MonopolySelectOptionScRsp.proto

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

type MonopolySelectOptionScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OptionId    uint32         `protobuf:"varint,13,opt,name=option_id,json=optionId,proto3" json:"option_id,omitempty"`
	Retcode     uint32         `protobuf:"varint,2,opt,name=retcode,proto3" json:"retcode,omitempty"`
	PCABNBHKFHP *OKANJDMIODN   `protobuf:"bytes,5,opt,name=PCABNBHKFHP,proto3" json:"PCABNBHKFHP,omitempty"`
	EventId     uint32         `protobuf:"varint,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	EMKNBMDJALN []*JMLKGHLGIBI `protobuf:"bytes,3,rep,name=EMKNBMDJALN,proto3" json:"EMKNBMDJALN,omitempty"`
}

func (x *MonopolySelectOptionScRsp) Reset() {
	*x = MonopolySelectOptionScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MonopolySelectOptionScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonopolySelectOptionScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonopolySelectOptionScRsp) ProtoMessage() {}

func (x *MonopolySelectOptionScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_MonopolySelectOptionScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonopolySelectOptionScRsp.ProtoReflect.Descriptor instead.
func (*MonopolySelectOptionScRsp) Descriptor() ([]byte, []int) {
	return file_MonopolySelectOptionScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *MonopolySelectOptionScRsp) GetOptionId() uint32 {
	if x != nil {
		return x.OptionId
	}
	return 0
}

func (x *MonopolySelectOptionScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *MonopolySelectOptionScRsp) GetPCABNBHKFHP() *OKANJDMIODN {
	if x != nil {
		return x.PCABNBHKFHP
	}
	return nil
}

func (x *MonopolySelectOptionScRsp) GetEventId() uint32 {
	if x != nil {
		return x.EventId
	}
	return 0
}

func (x *MonopolySelectOptionScRsp) GetEMKNBMDJALN() []*JMLKGHLGIBI {
	if x != nil {
		return x.EMKNBMDJALN
	}
	return nil
}

var File_MonopolySelectOptionScRsp_proto protoreflect.FileDescriptor

var file_MonopolySelectOptionScRsp_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x4f, 0x4b, 0x41, 0x4e, 0x4a, 0x44, 0x4d, 0x49, 0x4f, 0x44, 0x4e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x4d, 0x4c, 0x4b, 0x47, 0x48, 0x4c, 0x47, 0x49, 0x42,
	0x49, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x19, 0x4d, 0x6f, 0x6e, 0x6f,
	0x70, 0x6f, 0x6c, 0x79, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x1b, 0x0a, 0x09, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x0b,
	0x50, 0x43, 0x41, 0x42, 0x4e, 0x42, 0x48, 0x4b, 0x46, 0x48, 0x50, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x4b, 0x41, 0x4e, 0x4a, 0x44, 0x4d, 0x49, 0x4f, 0x44, 0x4e, 0x52,
	0x0b, 0x50, 0x43, 0x41, 0x42, 0x4e, 0x42, 0x48, 0x4b, 0x46, 0x48, 0x50, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0b, 0x45, 0x4d, 0x4b, 0x4e, 0x42,
	0x4d, 0x44, 0x4a, 0x41, 0x4c, 0x4e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a,
	0x4d, 0x4c, 0x4b, 0x47, 0x48, 0x4c, 0x47, 0x49, 0x42, 0x49, 0x52, 0x0b, 0x45, 0x4d, 0x4b, 0x4e,
	0x42, 0x4d, 0x44, 0x4a, 0x41, 0x4c, 0x4e, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MonopolySelectOptionScRsp_proto_rawDescOnce sync.Once
	file_MonopolySelectOptionScRsp_proto_rawDescData = file_MonopolySelectOptionScRsp_proto_rawDesc
)

func file_MonopolySelectOptionScRsp_proto_rawDescGZIP() []byte {
	file_MonopolySelectOptionScRsp_proto_rawDescOnce.Do(func() {
		file_MonopolySelectOptionScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_MonopolySelectOptionScRsp_proto_rawDescData)
	})
	return file_MonopolySelectOptionScRsp_proto_rawDescData
}

var file_MonopolySelectOptionScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MonopolySelectOptionScRsp_proto_goTypes = []interface{}{
	(*MonopolySelectOptionScRsp)(nil), // 0: MonopolySelectOptionScRsp
	(*OKANJDMIODN)(nil),               // 1: OKANJDMIODN
	(*JMLKGHLGIBI)(nil),               // 2: JMLKGHLGIBI
}
var file_MonopolySelectOptionScRsp_proto_depIdxs = []int32{
	1, // 0: MonopolySelectOptionScRsp.PCABNBHKFHP:type_name -> OKANJDMIODN
	2, // 1: MonopolySelectOptionScRsp.EMKNBMDJALN:type_name -> JMLKGHLGIBI
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_MonopolySelectOptionScRsp_proto_init() }
func file_MonopolySelectOptionScRsp_proto_init() {
	if File_MonopolySelectOptionScRsp_proto != nil {
		return
	}
	file_OKANJDMIODN_proto_init()
	file_JMLKGHLGIBI_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MonopolySelectOptionScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonopolySelectOptionScRsp); i {
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
			RawDescriptor: file_MonopolySelectOptionScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MonopolySelectOptionScRsp_proto_goTypes,
		DependencyIndexes: file_MonopolySelectOptionScRsp_proto_depIdxs,
		MessageInfos:      file_MonopolySelectOptionScRsp_proto_msgTypes,
	}.Build()
	File_MonopolySelectOptionScRsp_proto = out.File
	file_MonopolySelectOptionScRsp_proto_rawDesc = nil
	file_MonopolySelectOptionScRsp_proto_goTypes = nil
	file_MonopolySelectOptionScRsp_proto_depIdxs = nil
}
