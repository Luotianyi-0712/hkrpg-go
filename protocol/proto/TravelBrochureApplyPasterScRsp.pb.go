// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TravelBrochureApplyPasterScRsp.proto

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

type TravelBrochureApplyPasterScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode     uint32       `protobuf:"varint,12,opt,name=retcode,proto3" json:"retcode,omitempty"`
	CJNINFNKNML *OHBCINICBHP `protobuf:"bytes,10,opt,name=CJNINFNKNML,proto3" json:"CJNINFNKNML,omitempty"`
}

func (x *TravelBrochureApplyPasterScRsp) Reset() {
	*x = TravelBrochureApplyPasterScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TravelBrochureApplyPasterScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TravelBrochureApplyPasterScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TravelBrochureApplyPasterScRsp) ProtoMessage() {}

func (x *TravelBrochureApplyPasterScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_TravelBrochureApplyPasterScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TravelBrochureApplyPasterScRsp.ProtoReflect.Descriptor instead.
func (*TravelBrochureApplyPasterScRsp) Descriptor() ([]byte, []int) {
	return file_TravelBrochureApplyPasterScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *TravelBrochureApplyPasterScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *TravelBrochureApplyPasterScRsp) GetCJNINFNKNML() *OHBCINICBHP {
	if x != nil {
		return x.CJNINFNKNML
	}
	return nil
}

var File_TravelBrochureApplyPasterScRsp_proto protoreflect.FileDescriptor

var file_TravelBrochureApplyPasterScRsp_proto_rawDesc = []byte{
	0x0a, 0x24, 0x54, 0x72, 0x61, 0x76, 0x65, 0x6c, 0x42, 0x72, 0x6f, 0x63, 0x68, 0x75, 0x72, 0x65,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x50, 0x61, 0x73, 0x74, 0x65, 0x72, 0x53, 0x63, 0x52, 0x73, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x48, 0x42, 0x43, 0x49, 0x4e, 0x49, 0x43,
	0x42, 0x48, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x1e, 0x54, 0x72, 0x61,
	0x76, 0x65, 0x6c, 0x42, 0x72, 0x6f, 0x63, 0x68, 0x75, 0x72, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x50, 0x61, 0x73, 0x74, 0x65, 0x72, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65,
	0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x43, 0x4a, 0x4e, 0x49, 0x4e, 0x46, 0x4e,
	0x4b, 0x4e, 0x4d, 0x4c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4f, 0x48, 0x42,
	0x43, 0x49, 0x4e, 0x49, 0x43, 0x42, 0x48, 0x50, 0x52, 0x0b, 0x43, 0x4a, 0x4e, 0x49, 0x4e, 0x46,
	0x4e, 0x4b, 0x4e, 0x4d, 0x4c, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TravelBrochureApplyPasterScRsp_proto_rawDescOnce sync.Once
	file_TravelBrochureApplyPasterScRsp_proto_rawDescData = file_TravelBrochureApplyPasterScRsp_proto_rawDesc
)

func file_TravelBrochureApplyPasterScRsp_proto_rawDescGZIP() []byte {
	file_TravelBrochureApplyPasterScRsp_proto_rawDescOnce.Do(func() {
		file_TravelBrochureApplyPasterScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_TravelBrochureApplyPasterScRsp_proto_rawDescData)
	})
	return file_TravelBrochureApplyPasterScRsp_proto_rawDescData
}

var file_TravelBrochureApplyPasterScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TravelBrochureApplyPasterScRsp_proto_goTypes = []interface{}{
	(*TravelBrochureApplyPasterScRsp)(nil), // 0: TravelBrochureApplyPasterScRsp
	(*OHBCINICBHP)(nil),                    // 1: OHBCINICBHP
}
var file_TravelBrochureApplyPasterScRsp_proto_depIdxs = []int32{
	1, // 0: TravelBrochureApplyPasterScRsp.CJNINFNKNML:type_name -> OHBCINICBHP
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_TravelBrochureApplyPasterScRsp_proto_init() }
func file_TravelBrochureApplyPasterScRsp_proto_init() {
	if File_TravelBrochureApplyPasterScRsp_proto != nil {
		return
	}
	file_OHBCINICBHP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_TravelBrochureApplyPasterScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TravelBrochureApplyPasterScRsp); i {
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
			RawDescriptor: file_TravelBrochureApplyPasterScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TravelBrochureApplyPasterScRsp_proto_goTypes,
		DependencyIndexes: file_TravelBrochureApplyPasterScRsp_proto_depIdxs,
		MessageInfos:      file_TravelBrochureApplyPasterScRsp_proto_msgTypes,
	}.Build()
	File_TravelBrochureApplyPasterScRsp_proto = out.File
	file_TravelBrochureApplyPasterScRsp_proto_rawDesc = nil
	file_TravelBrochureApplyPasterScRsp_proto_goTypes = nil
	file_TravelBrochureApplyPasterScRsp_proto_depIdxs = nil
}
