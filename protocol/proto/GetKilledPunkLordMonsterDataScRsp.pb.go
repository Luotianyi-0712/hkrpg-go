// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetKilledPunkLordMonsterDataScRsp.proto

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

type GetKilledPunkLordMonsterDataScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode     uint32         `protobuf:"varint,5,opt,name=retcode,proto3" json:"retcode,omitempty"`
	MPHOCFEIEGH []*KJHICPDAFPM `protobuf:"bytes,2,rep,name=MPHOCFEIEGH,proto3" json:"MPHOCFEIEGH,omitempty"`
	PIOAHDCPDAC []*IBJHGNNGJCD `protobuf:"bytes,7,rep,name=PIOAHDCPDAC,proto3" json:"PIOAHDCPDAC,omitempty"`
}

func (x *GetKilledPunkLordMonsterDataScRsp) Reset() {
	*x = GetKilledPunkLordMonsterDataScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetKilledPunkLordMonsterDataScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetKilledPunkLordMonsterDataScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetKilledPunkLordMonsterDataScRsp) ProtoMessage() {}

func (x *GetKilledPunkLordMonsterDataScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetKilledPunkLordMonsterDataScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetKilledPunkLordMonsterDataScRsp.ProtoReflect.Descriptor instead.
func (*GetKilledPunkLordMonsterDataScRsp) Descriptor() ([]byte, []int) {
	return file_GetKilledPunkLordMonsterDataScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetKilledPunkLordMonsterDataScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetKilledPunkLordMonsterDataScRsp) GetMPHOCFEIEGH() []*KJHICPDAFPM {
	if x != nil {
		return x.MPHOCFEIEGH
	}
	return nil
}

func (x *GetKilledPunkLordMonsterDataScRsp) GetPIOAHDCPDAC() []*IBJHGNNGJCD {
	if x != nil {
		return x.PIOAHDCPDAC
	}
	return nil
}

var File_GetKilledPunkLordMonsterDataScRsp_proto protoreflect.FileDescriptor

var file_GetKilledPunkLordMonsterDataScRsp_proto_rawDesc = []byte{
	0x0a, 0x27, 0x47, 0x65, 0x74, 0x4b, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x50, 0x75, 0x6e, 0x6b, 0x4c,
	0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x4a, 0x48, 0x49, 0x43,
	0x50, 0x44, 0x41, 0x46, 0x50, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x42,
	0x4a, 0x48, 0x47, 0x4e, 0x4e, 0x47, 0x4a, 0x43, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x9d, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x4b, 0x69, 0x6c, 0x6c, 0x65, 0x64, 0x50, 0x75, 0x6e,
	0x6b, 0x4c, 0x6f, 0x72, 0x64, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x2e, 0x0a, 0x0b, 0x4d, 0x50, 0x48, 0x4f, 0x43, 0x46, 0x45, 0x49, 0x45, 0x47, 0x48, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x4a, 0x48, 0x49, 0x43, 0x50, 0x44, 0x41, 0x46,
	0x50, 0x4d, 0x52, 0x0b, 0x4d, 0x50, 0x48, 0x4f, 0x43, 0x46, 0x45, 0x49, 0x45, 0x47, 0x48, 0x12,
	0x2e, 0x0a, 0x0b, 0x50, 0x49, 0x4f, 0x41, 0x48, 0x44, 0x43, 0x50, 0x44, 0x41, 0x43, 0x18, 0x07,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x42, 0x4a, 0x48, 0x47, 0x4e, 0x4e, 0x47, 0x4a,
	0x43, 0x44, 0x52, 0x0b, 0x50, 0x49, 0x4f, 0x41, 0x48, 0x44, 0x43, 0x50, 0x44, 0x41, 0x43, 0x42,
	0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_GetKilledPunkLordMonsterDataScRsp_proto_rawDescOnce sync.Once
	file_GetKilledPunkLordMonsterDataScRsp_proto_rawDescData = file_GetKilledPunkLordMonsterDataScRsp_proto_rawDesc
)

func file_GetKilledPunkLordMonsterDataScRsp_proto_rawDescGZIP() []byte {
	file_GetKilledPunkLordMonsterDataScRsp_proto_rawDescOnce.Do(func() {
		file_GetKilledPunkLordMonsterDataScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetKilledPunkLordMonsterDataScRsp_proto_rawDescData)
	})
	return file_GetKilledPunkLordMonsterDataScRsp_proto_rawDescData
}

var file_GetKilledPunkLordMonsterDataScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetKilledPunkLordMonsterDataScRsp_proto_goTypes = []interface{}{
	(*GetKilledPunkLordMonsterDataScRsp)(nil), // 0: GetKilledPunkLordMonsterDataScRsp
	(*KJHICPDAFPM)(nil),                       // 1: KJHICPDAFPM
	(*IBJHGNNGJCD)(nil),                       // 2: IBJHGNNGJCD
}
var file_GetKilledPunkLordMonsterDataScRsp_proto_depIdxs = []int32{
	1, // 0: GetKilledPunkLordMonsterDataScRsp.MPHOCFEIEGH:type_name -> KJHICPDAFPM
	2, // 1: GetKilledPunkLordMonsterDataScRsp.PIOAHDCPDAC:type_name -> IBJHGNNGJCD
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GetKilledPunkLordMonsterDataScRsp_proto_init() }
func file_GetKilledPunkLordMonsterDataScRsp_proto_init() {
	if File_GetKilledPunkLordMonsterDataScRsp_proto != nil {
		return
	}
	file_KJHICPDAFPM_proto_init()
	file_IBJHGNNGJCD_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetKilledPunkLordMonsterDataScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetKilledPunkLordMonsterDataScRsp); i {
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
			RawDescriptor: file_GetKilledPunkLordMonsterDataScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetKilledPunkLordMonsterDataScRsp_proto_goTypes,
		DependencyIndexes: file_GetKilledPunkLordMonsterDataScRsp_proto_depIdxs,
		MessageInfos:      file_GetKilledPunkLordMonsterDataScRsp_proto_msgTypes,
	}.Build()
	File_GetKilledPunkLordMonsterDataScRsp_proto = out.File
	file_GetKilledPunkLordMonsterDataScRsp_proto_rawDesc = nil
	file_GetKilledPunkLordMonsterDataScRsp_proto_goTypes = nil
	file_GetKilledPunkLordMonsterDataScRsp_proto_depIdxs = nil
}
