// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MakeDrinkCsReq.proto

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

type MakeDrinkCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ONMCCOPHDEM uint32       `protobuf:"varint,13,opt,name=ONMCCOPHDEM,proto3" json:"ONMCCOPHDEM,omitempty"`
	LIKBNKOPFIF *EICPBAEMNEC `protobuf:"bytes,12,opt,name=LIKBNKOPFIF,proto3" json:"LIKBNKOPFIF,omitempty"`
}

func (x *MakeDrinkCsReq) Reset() {
	*x = MakeDrinkCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MakeDrinkCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MakeDrinkCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MakeDrinkCsReq) ProtoMessage() {}

func (x *MakeDrinkCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_MakeDrinkCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MakeDrinkCsReq.ProtoReflect.Descriptor instead.
func (*MakeDrinkCsReq) Descriptor() ([]byte, []int) {
	return file_MakeDrinkCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *MakeDrinkCsReq) GetONMCCOPHDEM() uint32 {
	if x != nil {
		return x.ONMCCOPHDEM
	}
	return 0
}

func (x *MakeDrinkCsReq) GetLIKBNKOPFIF() *EICPBAEMNEC {
	if x != nil {
		return x.LIKBNKOPFIF
	}
	return nil
}

var File_MakeDrinkCsReq_proto protoreflect.FileDescriptor

var file_MakeDrinkCsReq_proto_rawDesc = []byte{
	0x0a, 0x14, 0x4d, 0x61, 0x6b, 0x65, 0x44, 0x72, 0x69, 0x6e, 0x6b, 0x43, 0x73, 0x52, 0x65, 0x71,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x49, 0x43, 0x50, 0x42, 0x41, 0x45, 0x4d,
	0x4e, 0x45, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x0e, 0x4d, 0x61, 0x6b,
	0x65, 0x44, 0x72, 0x69, 0x6e, 0x6b, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x4f,
	0x4e, 0x4d, 0x43, 0x43, 0x4f, 0x50, 0x48, 0x44, 0x45, 0x4d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4f, 0x4e, 0x4d, 0x43, 0x43, 0x4f, 0x50, 0x48, 0x44, 0x45, 0x4d, 0x12, 0x2e, 0x0a,
	0x0b, 0x4c, 0x49, 0x4b, 0x42, 0x4e, 0x4b, 0x4f, 0x50, 0x46, 0x49, 0x46, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x49, 0x43, 0x50, 0x42, 0x41, 0x45, 0x4d, 0x4e, 0x45, 0x43,
	0x52, 0x0b, 0x4c, 0x49, 0x4b, 0x42, 0x4e, 0x4b, 0x4f, 0x50, 0x46, 0x49, 0x46, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MakeDrinkCsReq_proto_rawDescOnce sync.Once
	file_MakeDrinkCsReq_proto_rawDescData = file_MakeDrinkCsReq_proto_rawDesc
)

func file_MakeDrinkCsReq_proto_rawDescGZIP() []byte {
	file_MakeDrinkCsReq_proto_rawDescOnce.Do(func() {
		file_MakeDrinkCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_MakeDrinkCsReq_proto_rawDescData)
	})
	return file_MakeDrinkCsReq_proto_rawDescData
}

var file_MakeDrinkCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MakeDrinkCsReq_proto_goTypes = []interface{}{
	(*MakeDrinkCsReq)(nil), // 0: MakeDrinkCsReq
	(*EICPBAEMNEC)(nil),    // 1: EICPBAEMNEC
}
var file_MakeDrinkCsReq_proto_depIdxs = []int32{
	1, // 0: MakeDrinkCsReq.LIKBNKOPFIF:type_name -> EICPBAEMNEC
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_MakeDrinkCsReq_proto_init() }
func file_MakeDrinkCsReq_proto_init() {
	if File_MakeDrinkCsReq_proto != nil {
		return
	}
	file_EICPBAEMNEC_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MakeDrinkCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MakeDrinkCsReq); i {
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
			RawDescriptor: file_MakeDrinkCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MakeDrinkCsReq_proto_goTypes,
		DependencyIndexes: file_MakeDrinkCsReq_proto_depIdxs,
		MessageInfos:      file_MakeDrinkCsReq_proto_msgTypes,
	}.Build()
	File_MakeDrinkCsReq_proto = out.File
	file_MakeDrinkCsReq_proto_rawDesc = nil
	file_MakeDrinkCsReq_proto_goTypes = nil
	file_MakeDrinkCsReq_proto_depIdxs = nil
}
