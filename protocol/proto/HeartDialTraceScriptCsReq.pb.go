// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: HeartDialTraceScriptCsReq.proto

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

type HeartDialTraceScriptCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LICAGFDKGKE *CCOOEOEDCFN `protobuf:"bytes,12,opt,name=LICAGFDKGKE,proto3" json:"LICAGFDKGKE,omitempty"`
}

func (x *HeartDialTraceScriptCsReq) Reset() {
	*x = HeartDialTraceScriptCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HeartDialTraceScriptCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartDialTraceScriptCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartDialTraceScriptCsReq) ProtoMessage() {}

func (x *HeartDialTraceScriptCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_HeartDialTraceScriptCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartDialTraceScriptCsReq.ProtoReflect.Descriptor instead.
func (*HeartDialTraceScriptCsReq) Descriptor() ([]byte, []int) {
	return file_HeartDialTraceScriptCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *HeartDialTraceScriptCsReq) GetLICAGFDKGKE() *CCOOEOEDCFN {
	if x != nil {
		return x.LICAGFDKGKE
	}
	return nil
}

var File_HeartDialTraceScriptCsReq_proto protoreflect.FileDescriptor

var file_HeartDialTraceScriptCsReq_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x48, 0x65, 0x61, 0x72, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x63, 0x65,
	0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x43, 0x43, 0x4f, 0x4f, 0x45, 0x4f, 0x45, 0x44, 0x43, 0x46, 0x4e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x19, 0x48, 0x65, 0x61, 0x72, 0x74, 0x44, 0x69, 0x61,
	0x6c, 0x54, 0x72, 0x61, 0x63, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x43, 0x73, 0x52, 0x65,
	0x71, 0x12, 0x2e, 0x0a, 0x0b, 0x4c, 0x49, 0x43, 0x41, 0x47, 0x46, 0x44, 0x4b, 0x47, 0x4b, 0x45,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x43, 0x43, 0x4f, 0x4f, 0x45, 0x4f, 0x45,
	0x44, 0x43, 0x46, 0x4e, 0x52, 0x0b, 0x4c, 0x49, 0x43, 0x41, 0x47, 0x46, 0x44, 0x4b, 0x47, 0x4b,
	0x45, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_HeartDialTraceScriptCsReq_proto_rawDescOnce sync.Once
	file_HeartDialTraceScriptCsReq_proto_rawDescData = file_HeartDialTraceScriptCsReq_proto_rawDesc
)

func file_HeartDialTraceScriptCsReq_proto_rawDescGZIP() []byte {
	file_HeartDialTraceScriptCsReq_proto_rawDescOnce.Do(func() {
		file_HeartDialTraceScriptCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_HeartDialTraceScriptCsReq_proto_rawDescData)
	})
	return file_HeartDialTraceScriptCsReq_proto_rawDescData
}

var file_HeartDialTraceScriptCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HeartDialTraceScriptCsReq_proto_goTypes = []interface{}{
	(*HeartDialTraceScriptCsReq)(nil), // 0: HeartDialTraceScriptCsReq
	(*CCOOEOEDCFN)(nil),               // 1: CCOOEOEDCFN
}
var file_HeartDialTraceScriptCsReq_proto_depIdxs = []int32{
	1, // 0: HeartDialTraceScriptCsReq.LICAGFDKGKE:type_name -> CCOOEOEDCFN
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_HeartDialTraceScriptCsReq_proto_init() }
func file_HeartDialTraceScriptCsReq_proto_init() {
	if File_HeartDialTraceScriptCsReq_proto != nil {
		return
	}
	file_CCOOEOEDCFN_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HeartDialTraceScriptCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartDialTraceScriptCsReq); i {
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
			RawDescriptor: file_HeartDialTraceScriptCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HeartDialTraceScriptCsReq_proto_goTypes,
		DependencyIndexes: file_HeartDialTraceScriptCsReq_proto_depIdxs,
		MessageInfos:      file_HeartDialTraceScriptCsReq_proto_msgTypes,
	}.Build()
	File_HeartDialTraceScriptCsReq_proto = out.File
	file_HeartDialTraceScriptCsReq_proto_rawDesc = nil
	file_HeartDialTraceScriptCsReq_proto_goTypes = nil
	file_HeartDialTraceScriptCsReq_proto_depIdxs = nil
}
