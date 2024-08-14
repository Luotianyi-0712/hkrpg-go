// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EvolveBuildQueryInfoScRsp.proto

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

type EvolveBuildQueryInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LevelInfo        *EvolveBuildLevelInfo `protobuf:"bytes,2,opt,name=level_info,json=levelInfo,proto3" json:"level_info,omitempty"`
	Retcode          uint32                `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
	RogueCurrentInfo *KEPAMJFOKDN          `protobuf:"bytes,7,opt,name=rogue_current_info,json=rogueCurrentInfo,proto3" json:"rogue_current_info,omitempty"`
}

func (x *EvolveBuildQueryInfoScRsp) Reset() {
	*x = EvolveBuildQueryInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvolveBuildQueryInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvolveBuildQueryInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvolveBuildQueryInfoScRsp) ProtoMessage() {}

func (x *EvolveBuildQueryInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_EvolveBuildQueryInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvolveBuildQueryInfoScRsp.ProtoReflect.Descriptor instead.
func (*EvolveBuildQueryInfoScRsp) Descriptor() ([]byte, []int) {
	return file_EvolveBuildQueryInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *EvolveBuildQueryInfoScRsp) GetLevelInfo() *EvolveBuildLevelInfo {
	if x != nil {
		return x.LevelInfo
	}
	return nil
}

func (x *EvolveBuildQueryInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *EvolveBuildQueryInfoScRsp) GetRogueCurrentInfo() *KEPAMJFOKDN {
	if x != nil {
		return x.RogueCurrentInfo
	}
	return nil
}

var File_EvolveBuildQueryInfoScRsp_proto protoreflect.FileDescriptor

var file_EvolveBuildQueryInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x4b, 0x45, 0x50, 0x41, 0x4d, 0x4a, 0x46, 0x4f, 0x4b, 0x44, 0x4e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa7, 0x01, 0x0a, 0x19, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x34,
	0x0a, 0x0a, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3a,
	0x0a, 0x12, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x45, 0x50,
	0x41, 0x4d, 0x4a, 0x46, 0x4f, 0x4b, 0x44, 0x4e, 0x52, 0x10, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvolveBuildQueryInfoScRsp_proto_rawDescOnce sync.Once
	file_EvolveBuildQueryInfoScRsp_proto_rawDescData = file_EvolveBuildQueryInfoScRsp_proto_rawDesc
)

func file_EvolveBuildQueryInfoScRsp_proto_rawDescGZIP() []byte {
	file_EvolveBuildQueryInfoScRsp_proto_rawDescOnce.Do(func() {
		file_EvolveBuildQueryInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvolveBuildQueryInfoScRsp_proto_rawDescData)
	})
	return file_EvolveBuildQueryInfoScRsp_proto_rawDescData
}

var file_EvolveBuildQueryInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvolveBuildQueryInfoScRsp_proto_goTypes = []interface{}{
	(*EvolveBuildQueryInfoScRsp)(nil), // 0: EvolveBuildQueryInfoScRsp
	(*EvolveBuildLevelInfo)(nil),      // 1: EvolveBuildLevelInfo
	(*KEPAMJFOKDN)(nil),               // 2: KEPAMJFOKDN
}
var file_EvolveBuildQueryInfoScRsp_proto_depIdxs = []int32{
	1, // 0: EvolveBuildQueryInfoScRsp.level_info:type_name -> EvolveBuildLevelInfo
	2, // 1: EvolveBuildQueryInfoScRsp.rogue_current_info:type_name -> KEPAMJFOKDN
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_EvolveBuildQueryInfoScRsp_proto_init() }
func file_EvolveBuildQueryInfoScRsp_proto_init() {
	if File_EvolveBuildQueryInfoScRsp_proto != nil {
		return
	}
	file_KEPAMJFOKDN_proto_init()
	file_EvolveBuildLevelInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvolveBuildQueryInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvolveBuildQueryInfoScRsp); i {
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
			RawDescriptor: file_EvolveBuildQueryInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvolveBuildQueryInfoScRsp_proto_goTypes,
		DependencyIndexes: file_EvolveBuildQueryInfoScRsp_proto_depIdxs,
		MessageInfos:      file_EvolveBuildQueryInfoScRsp_proto_msgTypes,
	}.Build()
	File_EvolveBuildQueryInfoScRsp_proto = out.File
	file_EvolveBuildQueryInfoScRsp_proto_rawDesc = nil
	file_EvolveBuildQueryInfoScRsp_proto_goTypes = nil
	file_EvolveBuildQueryInfoScRsp_proto_depIdxs = nil
}
