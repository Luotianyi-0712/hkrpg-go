// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RelicList.proto

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

type RelicList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SetId uint32 `protobuf:"varint,4,opt,name=set_id,json=setId,proto3" json:"set_id,omitempty"`
	Type  uint32 `protobuf:"varint,10,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *RelicList) Reset() {
	*x = RelicList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RelicList_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelicList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelicList) ProtoMessage() {}

func (x *RelicList) ProtoReflect() protoreflect.Message {
	mi := &file_RelicList_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelicList.ProtoReflect.Descriptor instead.
func (*RelicList) Descriptor() ([]byte, []int) {
	return file_RelicList_proto_rawDescGZIP(), []int{0}
}

func (x *RelicList) GetSetId() uint32 {
	if x != nil {
		return x.SetId
	}
	return 0
}

func (x *RelicList) GetType() uint32 {
	if x != nil {
		return x.Type
	}
	return 0
}

var File_RelicList_proto protoreflect.FileDescriptor

var file_RelicList_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x36, 0x0a, 0x09, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15,
	0x0a, 0x06, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RelicList_proto_rawDescOnce sync.Once
	file_RelicList_proto_rawDescData = file_RelicList_proto_rawDesc
)

func file_RelicList_proto_rawDescGZIP() []byte {
	file_RelicList_proto_rawDescOnce.Do(func() {
		file_RelicList_proto_rawDescData = protoimpl.X.CompressGZIP(file_RelicList_proto_rawDescData)
	})
	return file_RelicList_proto_rawDescData
}

var file_RelicList_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RelicList_proto_goTypes = []interface{}{
	(*RelicList)(nil), // 0: RelicList
}
var file_RelicList_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_RelicList_proto_init() }
func file_RelicList_proto_init() {
	if File_RelicList_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_RelicList_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelicList); i {
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
			RawDescriptor: file_RelicList_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RelicList_proto_goTypes,
		DependencyIndexes: file_RelicList_proto_depIdxs,
		MessageInfos:      file_RelicList_proto_msgTypes,
	}.Build()
	File_RelicList_proto = out.File
	file_RelicList_proto_rawDesc = nil
	file_RelicList_proto_goTypes = nil
	file_RelicList_proto_depIdxs = nil
}
