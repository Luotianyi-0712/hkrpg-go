// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MonopolyAcceptQuizScRsp.proto

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

type MonopolyAcceptQuizScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EBPMCIHJLNP *IAGKOOPHPBM `protobuf:"bytes,15,opt,name=EBPMCIHJLNP,proto3" json:"EBPMCIHJLNP,omitempty"`
	Retcode     uint32       `protobuf:"varint,10,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *MonopolyAcceptQuizScRsp) Reset() {
	*x = MonopolyAcceptQuizScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MonopolyAcceptQuizScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MonopolyAcceptQuizScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MonopolyAcceptQuizScRsp) ProtoMessage() {}

func (x *MonopolyAcceptQuizScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_MonopolyAcceptQuizScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MonopolyAcceptQuizScRsp.ProtoReflect.Descriptor instead.
func (*MonopolyAcceptQuizScRsp) Descriptor() ([]byte, []int) {
	return file_MonopolyAcceptQuizScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *MonopolyAcceptQuizScRsp) GetEBPMCIHJLNP() *IAGKOOPHPBM {
	if x != nil {
		return x.EBPMCIHJLNP
	}
	return nil
}

func (x *MonopolyAcceptQuizScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_MonopolyAcceptQuizScRsp_proto protoreflect.FileDescriptor

var file_MonopolyAcceptQuizScRsp_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x51, 0x75, 0x69, 0x7a, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x49, 0x41, 0x47, 0x4b, 0x4f, 0x4f, 0x50, 0x48, 0x50, 0x42, 0x4d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x63, 0x0a, 0x17, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x41, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x51, 0x75, 0x69, 0x7a, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x2e, 0x0a,
	0x0b, 0x45, 0x42, 0x50, 0x4d, 0x43, 0x49, 0x48, 0x4a, 0x4c, 0x4e, 0x50, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x49, 0x41, 0x47, 0x4b, 0x4f, 0x4f, 0x50, 0x48, 0x50, 0x42, 0x4d,
	0x52, 0x0b, 0x45, 0x42, 0x50, 0x4d, 0x43, 0x49, 0x48, 0x4a, 0x4c, 0x4e, 0x50, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MonopolyAcceptQuizScRsp_proto_rawDescOnce sync.Once
	file_MonopolyAcceptQuizScRsp_proto_rawDescData = file_MonopolyAcceptQuizScRsp_proto_rawDesc
)

func file_MonopolyAcceptQuizScRsp_proto_rawDescGZIP() []byte {
	file_MonopolyAcceptQuizScRsp_proto_rawDescOnce.Do(func() {
		file_MonopolyAcceptQuizScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_MonopolyAcceptQuizScRsp_proto_rawDescData)
	})
	return file_MonopolyAcceptQuizScRsp_proto_rawDescData
}

var file_MonopolyAcceptQuizScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MonopolyAcceptQuizScRsp_proto_goTypes = []interface{}{
	(*MonopolyAcceptQuizScRsp)(nil), // 0: MonopolyAcceptQuizScRsp
	(*IAGKOOPHPBM)(nil),             // 1: IAGKOOPHPBM
}
var file_MonopolyAcceptQuizScRsp_proto_depIdxs = []int32{
	1, // 0: MonopolyAcceptQuizScRsp.EBPMCIHJLNP:type_name -> IAGKOOPHPBM
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_MonopolyAcceptQuizScRsp_proto_init() }
func file_MonopolyAcceptQuizScRsp_proto_init() {
	if File_MonopolyAcceptQuizScRsp_proto != nil {
		return
	}
	file_IAGKOOPHPBM_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MonopolyAcceptQuizScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MonopolyAcceptQuizScRsp); i {
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
			RawDescriptor: file_MonopolyAcceptQuizScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MonopolyAcceptQuizScRsp_proto_goTypes,
		DependencyIndexes: file_MonopolyAcceptQuizScRsp_proto_depIdxs,
		MessageInfos:      file_MonopolyAcceptQuizScRsp_proto_msgTypes,
	}.Build()
	File_MonopolyAcceptQuizScRsp_proto = out.File
	file_MonopolyAcceptQuizScRsp_proto_rawDesc = nil
	file_MonopolyAcceptQuizScRsp_proto_goTypes = nil
	file_MonopolyAcceptQuizScRsp_proto_depIdxs = nil
}
