// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: LJMJPODBCAE.proto

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

type LJMJPODBCAE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KAGIFENOLLC []*LIDHMBFBHAO `protobuf:"bytes,12,rep,name=KAGIFENOLLC,proto3" json:"KAGIFENOLLC,omitempty"`
}

func (x *LJMJPODBCAE) Reset() {
	*x = LJMJPODBCAE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LJMJPODBCAE_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LJMJPODBCAE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LJMJPODBCAE) ProtoMessage() {}

func (x *LJMJPODBCAE) ProtoReflect() protoreflect.Message {
	mi := &file_LJMJPODBCAE_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LJMJPODBCAE.ProtoReflect.Descriptor instead.
func (*LJMJPODBCAE) Descriptor() ([]byte, []int) {
	return file_LJMJPODBCAE_proto_rawDescGZIP(), []int{0}
}

func (x *LJMJPODBCAE) GetKAGIFENOLLC() []*LIDHMBFBHAO {
	if x != nil {
		return x.KAGIFENOLLC
	}
	return nil
}

var File_LJMJPODBCAE_proto protoreflect.FileDescriptor

var file_LJMJPODBCAE_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4c, 0x4a, 0x4d, 0x4a, 0x50, 0x4f, 0x44, 0x42, 0x43, 0x41, 0x45, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x49, 0x44, 0x48, 0x4d, 0x42, 0x46, 0x42, 0x48, 0x41, 0x4f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x0b, 0x4c, 0x4a, 0x4d, 0x4a, 0x50, 0x4f,
	0x44, 0x42, 0x43, 0x41, 0x45, 0x12, 0x2e, 0x0a, 0x0b, 0x4b, 0x41, 0x47, 0x49, 0x46, 0x45, 0x4e,
	0x4f, 0x4c, 0x4c, 0x43, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x49, 0x44,
	0x48, 0x4d, 0x42, 0x46, 0x42, 0x48, 0x41, 0x4f, 0x52, 0x0b, 0x4b, 0x41, 0x47, 0x49, 0x46, 0x45,
	0x4e, 0x4f, 0x4c, 0x4c, 0x43, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LJMJPODBCAE_proto_rawDescOnce sync.Once
	file_LJMJPODBCAE_proto_rawDescData = file_LJMJPODBCAE_proto_rawDesc
)

func file_LJMJPODBCAE_proto_rawDescGZIP() []byte {
	file_LJMJPODBCAE_proto_rawDescOnce.Do(func() {
		file_LJMJPODBCAE_proto_rawDescData = protoimpl.X.CompressGZIP(file_LJMJPODBCAE_proto_rawDescData)
	})
	return file_LJMJPODBCAE_proto_rawDescData
}

var file_LJMJPODBCAE_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LJMJPODBCAE_proto_goTypes = []interface{}{
	(*LJMJPODBCAE)(nil), // 0: LJMJPODBCAE
	(*LIDHMBFBHAO)(nil), // 1: LIDHMBFBHAO
}
var file_LJMJPODBCAE_proto_depIdxs = []int32{
	1, // 0: LJMJPODBCAE.KAGIFENOLLC:type_name -> LIDHMBFBHAO
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_LJMJPODBCAE_proto_init() }
func file_LJMJPODBCAE_proto_init() {
	if File_LJMJPODBCAE_proto != nil {
		return
	}
	file_LIDHMBFBHAO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LJMJPODBCAE_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LJMJPODBCAE); i {
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
			RawDescriptor: file_LJMJPODBCAE_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LJMJPODBCAE_proto_goTypes,
		DependencyIndexes: file_LJMJPODBCAE_proto_depIdxs,
		MessageInfos:      file_LJMJPODBCAE_proto_msgTypes,
	}.Build()
	File_LJMJPODBCAE_proto = out.File
	file_LJMJPODBCAE_proto_rawDesc = nil
	file_LJMJPODBCAE_proto_goTypes = nil
	file_LJMJPODBCAE_proto_depIdxs = nil
}
