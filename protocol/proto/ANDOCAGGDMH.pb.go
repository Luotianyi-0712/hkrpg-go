// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ANDOCAGGDMH.proto

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

type ANDOCAGGDMH struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VaryValue uint32 `protobuf:"varint,2,opt,name=vary_value,json=varyValue,proto3" json:"vary_value,omitempty"`
	CurValue  uint32 `protobuf:"varint,6,opt,name=cur_value,json=curValue,proto3" json:"cur_value,omitempty"`
}

func (x *ANDOCAGGDMH) Reset() {
	*x = ANDOCAGGDMH{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ANDOCAGGDMH_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ANDOCAGGDMH) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ANDOCAGGDMH) ProtoMessage() {}

func (x *ANDOCAGGDMH) ProtoReflect() protoreflect.Message {
	mi := &file_ANDOCAGGDMH_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ANDOCAGGDMH.ProtoReflect.Descriptor instead.
func (*ANDOCAGGDMH) Descriptor() ([]byte, []int) {
	return file_ANDOCAGGDMH_proto_rawDescGZIP(), []int{0}
}

func (x *ANDOCAGGDMH) GetVaryValue() uint32 {
	if x != nil {
		return x.VaryValue
	}
	return 0
}

func (x *ANDOCAGGDMH) GetCurValue() uint32 {
	if x != nil {
		return x.CurValue
	}
	return 0
}

var File_ANDOCAGGDMH_proto protoreflect.FileDescriptor

var file_ANDOCAGGDMH_proto_rawDesc = []byte{
	0x0a, 0x11, 0x41, 0x4e, 0x44, 0x4f, 0x43, 0x41, 0x47, 0x47, 0x44, 0x4d, 0x48, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x49, 0x0a, 0x0b, 0x41, 0x4e, 0x44, 0x4f, 0x43, 0x41, 0x47, 0x47, 0x44,
	0x4d, 0x48, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x72, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x76, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x75, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2e,
	0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ANDOCAGGDMH_proto_rawDescOnce sync.Once
	file_ANDOCAGGDMH_proto_rawDescData = file_ANDOCAGGDMH_proto_rawDesc
)

func file_ANDOCAGGDMH_proto_rawDescGZIP() []byte {
	file_ANDOCAGGDMH_proto_rawDescOnce.Do(func() {
		file_ANDOCAGGDMH_proto_rawDescData = protoimpl.X.CompressGZIP(file_ANDOCAGGDMH_proto_rawDescData)
	})
	return file_ANDOCAGGDMH_proto_rawDescData
}

var file_ANDOCAGGDMH_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ANDOCAGGDMH_proto_goTypes = []interface{}{
	(*ANDOCAGGDMH)(nil), // 0: ANDOCAGGDMH
}
var file_ANDOCAGGDMH_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ANDOCAGGDMH_proto_init() }
func file_ANDOCAGGDMH_proto_init() {
	if File_ANDOCAGGDMH_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ANDOCAGGDMH_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ANDOCAGGDMH); i {
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
			RawDescriptor: file_ANDOCAGGDMH_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ANDOCAGGDMH_proto_goTypes,
		DependencyIndexes: file_ANDOCAGGDMH_proto_depIdxs,
		MessageInfos:      file_ANDOCAGGDMH_proto_msgTypes,
	}.Build()
	File_ANDOCAGGDMH_proto = out.File
	file_ANDOCAGGDMH_proto_rawDesc = nil
	file_ANDOCAGGDMH_proto_goTypes = nil
	file_ANDOCAGGDMH_proto_depIdxs = nil
}
