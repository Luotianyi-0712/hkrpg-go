// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetStageLineupCsReq.proto

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

type GetStageLineupCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStageLineupCsReq) Reset() {
	*x = GetStageLineupCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetStageLineupCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStageLineupCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStageLineupCsReq) ProtoMessage() {}

func (x *GetStageLineupCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_GetStageLineupCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStageLineupCsReq.ProtoReflect.Descriptor instead.
func (*GetStageLineupCsReq) Descriptor() ([]byte, []int) {
	return file_GetStageLineupCsReq_proto_rawDescGZIP(), []int{0}
}

var File_GetStageLineupCsReq_proto protoreflect.FileDescriptor

var file_GetStageLineupCsReq_proto_rawDesc = []byte{
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70,
	0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x15, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x6e, 0x65, 0x75, 0x70, 0x43, 0x73, 0x52,
	0x65, 0x71, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetStageLineupCsReq_proto_rawDescOnce sync.Once
	file_GetStageLineupCsReq_proto_rawDescData = file_GetStageLineupCsReq_proto_rawDesc
)

func file_GetStageLineupCsReq_proto_rawDescGZIP() []byte {
	file_GetStageLineupCsReq_proto_rawDescOnce.Do(func() {
		file_GetStageLineupCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetStageLineupCsReq_proto_rawDescData)
	})
	return file_GetStageLineupCsReq_proto_rawDescData
}

var file_GetStageLineupCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetStageLineupCsReq_proto_goTypes = []interface{}{
	(*GetStageLineupCsReq)(nil), // 0: GetStageLineupCsReq
}
var file_GetStageLineupCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetStageLineupCsReq_proto_init() }
func file_GetStageLineupCsReq_proto_init() {
	if File_GetStageLineupCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetStageLineupCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStageLineupCsReq); i {
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
			RawDescriptor: file_GetStageLineupCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetStageLineupCsReq_proto_goTypes,
		DependencyIndexes: file_GetStageLineupCsReq_proto_depIdxs,
		MessageInfos:      file_GetStageLineupCsReq_proto_msgTypes,
	}.Build()
	File_GetStageLineupCsReq_proto = out.File
	file_GetStageLineupCsReq_proto_rawDesc = nil
	file_GetStageLineupCsReq_proto_goTypes = nil
	file_GetStageLineupCsReq_proto_depIdxs = nil
}