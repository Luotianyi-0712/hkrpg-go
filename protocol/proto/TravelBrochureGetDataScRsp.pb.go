// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TravelBrochureGetDataScRsp.proto

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

type TravelBrochureGetDataScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomValue uint32                  `protobuf:"varint,3,opt,name=custom_value,json=customValue,proto3" json:"custom_value,omitempty"`
	Retcode     uint32                  `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
	OLDHAJNFNHA map[uint32]*OHBCINICBHP `protobuf:"bytes,4,rep,name=OLDHAJNFNHA,proto3" json:"OLDHAJNFNHA,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BPBFCBCLMKL map[uint32]uint32       `protobuf:"bytes,6,rep,name=BPBFCBCLMKL,proto3" json:"BPBFCBCLMKL,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *TravelBrochureGetDataScRsp) Reset() {
	*x = TravelBrochureGetDataScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TravelBrochureGetDataScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TravelBrochureGetDataScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TravelBrochureGetDataScRsp) ProtoMessage() {}

func (x *TravelBrochureGetDataScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_TravelBrochureGetDataScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TravelBrochureGetDataScRsp.ProtoReflect.Descriptor instead.
func (*TravelBrochureGetDataScRsp) Descriptor() ([]byte, []int) {
	return file_TravelBrochureGetDataScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *TravelBrochureGetDataScRsp) GetCustomValue() uint32 {
	if x != nil {
		return x.CustomValue
	}
	return 0
}

func (x *TravelBrochureGetDataScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *TravelBrochureGetDataScRsp) GetOLDHAJNFNHA() map[uint32]*OHBCINICBHP {
	if x != nil {
		return x.OLDHAJNFNHA
	}
	return nil
}

func (x *TravelBrochureGetDataScRsp) GetBPBFCBCLMKL() map[uint32]uint32 {
	if x != nil {
		return x.BPBFCBCLMKL
	}
	return nil
}

var File_TravelBrochureGetDataScRsp_proto protoreflect.FileDescriptor

var file_TravelBrochureGetDataScRsp_proto_rawDesc = []byte{
	0x0a, 0x20, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x42, 0x72, 0x6f, 0x63, 0x68, 0x75, 0x72, 0x65,
	0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x48, 0x42, 0x43, 0x49, 0x4e, 0x49, 0x43, 0x42, 0x48, 0x50, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x03, 0x0a, 0x1a, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c,
	0x42, 0x72, 0x6f, 0x63, 0x68, 0x75, 0x72, 0x65, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x4f, 0x4c, 0x44, 0x48, 0x41, 0x4a, 0x4e, 0x46, 0x4e, 0x48, 0x41,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x42,
	0x72, 0x6f, 0x63, 0x68, 0x75, 0x72, 0x65, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x2e, 0x4f, 0x4c, 0x44, 0x48, 0x41, 0x4a, 0x4e, 0x46, 0x4e, 0x48, 0x41, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x4f, 0x4c, 0x44, 0x48, 0x41, 0x4a, 0x4e, 0x46, 0x4e, 0x48,
	0x41, 0x12, 0x4e, 0x0a, 0x0b, 0x42, 0x50, 0x42, 0x46, 0x43, 0x42, 0x43, 0x4c, 0x4d, 0x4b, 0x4c,
	0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x42,
	0x72, 0x6f, 0x63, 0x68, 0x75, 0x72, 0x65, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x2e, 0x42, 0x50, 0x42, 0x46, 0x43, 0x42, 0x43, 0x4c, 0x4d, 0x4b, 0x4c, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x42, 0x50, 0x42, 0x46, 0x43, 0x42, 0x43, 0x4c, 0x4d, 0x4b,
	0x4c, 0x1a, 0x4c, 0x0a, 0x10, 0x4f, 0x4c, 0x44, 0x48, 0x41, 0x4a, 0x4e, 0x46, 0x4e, 0x48, 0x41,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x48, 0x42, 0x43, 0x49, 0x4e, 0x49,
	0x43, 0x42, 0x48, 0x50, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x3e, 0x0a, 0x10, 0x42, 0x50, 0x42, 0x46, 0x43, 0x42, 0x43, 0x4c, 0x4d, 0x4b, 0x4c, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_TravelBrochureGetDataScRsp_proto_rawDescOnce sync.Once
	file_TravelBrochureGetDataScRsp_proto_rawDescData = file_TravelBrochureGetDataScRsp_proto_rawDesc
)

func file_TravelBrochureGetDataScRsp_proto_rawDescGZIP() []byte {
	file_TravelBrochureGetDataScRsp_proto_rawDescOnce.Do(func() {
		file_TravelBrochureGetDataScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_TravelBrochureGetDataScRsp_proto_rawDescData)
	})
	return file_TravelBrochureGetDataScRsp_proto_rawDescData
}

var file_TravelBrochureGetDataScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_TravelBrochureGetDataScRsp_proto_goTypes = []interface{}{
	(*TravelBrochureGetDataScRsp)(nil), // 0: TravelBrochureGetDataScRsp
	nil,                                // 1: TravelBrochureGetDataScRsp.OLDHAJNFNHAEntry
	nil,                                // 2: TravelBrochureGetDataScRsp.BPBFCBCLMKLEntry
	(*OHBCINICBHP)(nil),                // 3: OHBCINICBHP
}
var file_TravelBrochureGetDataScRsp_proto_depIdxs = []int32{
	1, // 0: TravelBrochureGetDataScRsp.OLDHAJNFNHA:type_name -> TravelBrochureGetDataScRsp.OLDHAJNFNHAEntry
	2, // 1: TravelBrochureGetDataScRsp.BPBFCBCLMKL:type_name -> TravelBrochureGetDataScRsp.BPBFCBCLMKLEntry
	3, // 2: TravelBrochureGetDataScRsp.OLDHAJNFNHAEntry.value:type_name -> OHBCINICBHP
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_TravelBrochureGetDataScRsp_proto_init() }
func file_TravelBrochureGetDataScRsp_proto_init() {
	if File_TravelBrochureGetDataScRsp_proto != nil {
		return
	}
	file_OHBCINICBHP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TravelBrochureGetDataScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TravelBrochureGetDataScRsp); i {
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
			RawDescriptor: file_TravelBrochureGetDataScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TravelBrochureGetDataScRsp_proto_goTypes,
		DependencyIndexes: file_TravelBrochureGetDataScRsp_proto_depIdxs,
		MessageInfos:      file_TravelBrochureGetDataScRsp_proto_msgTypes,
	}.Build()
	File_TravelBrochureGetDataScRsp_proto = out.File
	file_TravelBrochureGetDataScRsp_proto_rawDesc = nil
	file_TravelBrochureGetDataScRsp_proto_goTypes = nil
	file_TravelBrochureGetDataScRsp_proto_depIdxs = nil
}
