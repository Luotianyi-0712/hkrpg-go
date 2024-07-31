// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueMiracleInfo.proto

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

type RogueMiracleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BDDANOBJMEL uint32          `protobuf:"varint,8,opt,name=BDDANOBJMEL,proto3" json:"BDDANOBJMEL,omitempty"`
	BAILOBNCIGP uint32          `protobuf:"varint,4,opt,name=BAILOBNCIGP,proto3" json:"BAILOBNCIGP,omitempty"`
	MiracleList []*RogueMiracle `protobuf:"bytes,6,rep,name=miracle_list,json=miracleList,proto3" json:"miracle_list,omitempty"`
}

func (x *RogueMiracleInfo) Reset() {
	*x = RogueMiracleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueMiracleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueMiracleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueMiracleInfo) ProtoMessage() {}

func (x *RogueMiracleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueMiracleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueMiracleInfo.ProtoReflect.Descriptor instead.
func (*RogueMiracleInfo) Descriptor() ([]byte, []int) {
	return file_RogueMiracleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueMiracleInfo) GetBDDANOBJMEL() uint32 {
	if x != nil {
		return x.BDDANOBJMEL
	}
	return 0
}

func (x *RogueMiracleInfo) GetBAILOBNCIGP() uint32 {
	if x != nil {
		return x.BAILOBNCIGP
	}
	return 0
}

func (x *RogueMiracleInfo) GetMiracleList() []*RogueMiracle {
	if x != nil {
		return x.MiracleList
	}
	return nil
}

var File_RogueMiracleInfo_proto protoreflect.FileDescriptor

var file_RogueMiracleInfo_proto_rawDesc = []byte{
	0x0a, 0x16, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d,
	0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a,
	0x10, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x44, 0x44, 0x41, 0x4e, 0x4f, 0x42, 0x4a, 0x4d, 0x45, 0x4c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x44, 0x44, 0x41, 0x4e, 0x4f, 0x42, 0x4a,
	0x4d, 0x45, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x41, 0x49, 0x4c, 0x4f, 0x42, 0x4e, 0x43, 0x49,
	0x47, 0x50, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x41, 0x49, 0x4c, 0x4f, 0x42,
	0x4e, 0x43, 0x49, 0x47, 0x50, 0x12, 0x30, 0x0a, 0x0c, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x0b, 0x6d, 0x69, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueMiracleInfo_proto_rawDescOnce sync.Once
	file_RogueMiracleInfo_proto_rawDescData = file_RogueMiracleInfo_proto_rawDesc
)

func file_RogueMiracleInfo_proto_rawDescGZIP() []byte {
	file_RogueMiracleInfo_proto_rawDescOnce.Do(func() {
		file_RogueMiracleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueMiracleInfo_proto_rawDescData)
	})
	return file_RogueMiracleInfo_proto_rawDescData
}

var file_RogueMiracleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueMiracleInfo_proto_goTypes = []interface{}{
	(*RogueMiracleInfo)(nil), // 0: RogueMiracleInfo
	(*RogueMiracle)(nil),     // 1: RogueMiracle
}
var file_RogueMiracleInfo_proto_depIdxs = []int32{
	1, // 0: RogueMiracleInfo.miracle_list:type_name -> RogueMiracle
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_RogueMiracleInfo_proto_init() }
func file_RogueMiracleInfo_proto_init() {
	if File_RogueMiracleInfo_proto != nil {
		return
	}
	file_RogueMiracle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueMiracleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueMiracleInfo); i {
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
			RawDescriptor: file_RogueMiracleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueMiracleInfo_proto_goTypes,
		DependencyIndexes: file_RogueMiracleInfo_proto_depIdxs,
		MessageInfos:      file_RogueMiracleInfo_proto_msgTypes,
	}.Build()
	File_RogueMiracleInfo_proto = out.File
	file_RogueMiracleInfo_proto_rawDesc = nil
	file_RogueMiracleInfo_proto_goTypes = nil
	file_RogueMiracleInfo_proto_depIdxs = nil
}
