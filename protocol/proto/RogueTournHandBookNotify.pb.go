// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueTournHandBookNotify.proto

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

type RogueTournHandBookNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MiracleIdList      []uint32           `protobuf:"varint,14,rep,packed,name=miracle_id_list,json=miracleIdList,proto3" json:"miracle_id_list,omitempty"`
	RogueTournHandbook RogueTournHandbook `protobuf:"varint,7,opt,name=rogue_tourn_handbook,json=rogueTournHandbook,proto3,enum=RogueTournHandbook" json:"rogue_tourn_handbook,omitempty"`
}

func (x *RogueTournHandBookNotify) Reset() {
	*x = RogueTournHandBookNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueTournHandBookNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueTournHandBookNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueTournHandBookNotify) ProtoMessage() {}

func (x *RogueTournHandBookNotify) ProtoReflect() protoreflect.Message {
	mi := &file_RogueTournHandBookNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueTournHandBookNotify.ProtoReflect.Descriptor instead.
func (*RogueTournHandBookNotify) Descriptor() ([]byte, []int) {
	return file_RogueTournHandBookNotify_proto_rawDescGZIP(), []int{0}
}

func (x *RogueTournHandBookNotify) GetMiracleIdList() []uint32 {
	if x != nil {
		return x.MiracleIdList
	}
	return nil
}

func (x *RogueTournHandBookNotify) GetRogueTournHandbook() RogueTournHandbook {
	if x != nil {
		return x.RogueTournHandbook
	}
	return RogueTournHandbook_ROGUE_TOURN_HANDBOOK_NONE
}

var File_RogueTournHandBookNotify_proto protoreflect.FileDescriptor

var file_RogueTournHandBookNotify_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x48, 0x61, 0x6e, 0x64,
	0x42, 0x6f, 0x6f, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x18, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x48, 0x61, 0x6e, 0x64,
	0x62, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x18, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x42, 0x6f, 0x6f,
	0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x26, 0x0a, 0x0f, 0x6d, 0x69, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0d, 0x6d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x45, 0x0a, 0x14, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x74, 0x6f, 0x75, 0x72, 0x6e, 0x5f, 0x68,
	0x61, 0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x48, 0x61, 0x6e, 0x64, 0x62, 0x6f,
	0x6f, 0x6b, 0x52, 0x12, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x6f, 0x75, 0x72, 0x6e, 0x48, 0x61,
	0x6e, 0x64, 0x62, 0x6f, 0x6f, 0x6b, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueTournHandBookNotify_proto_rawDescOnce sync.Once
	file_RogueTournHandBookNotify_proto_rawDescData = file_RogueTournHandBookNotify_proto_rawDesc
)

func file_RogueTournHandBookNotify_proto_rawDescGZIP() []byte {
	file_RogueTournHandBookNotify_proto_rawDescOnce.Do(func() {
		file_RogueTournHandBookNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueTournHandBookNotify_proto_rawDescData)
	})
	return file_RogueTournHandBookNotify_proto_rawDescData
}

var file_RogueTournHandBookNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueTournHandBookNotify_proto_goTypes = []interface{}{
	(*RogueTournHandBookNotify)(nil), // 0: RogueTournHandBookNotify
	(RogueTournHandbook)(0),          // 1: RogueTournHandbook
}
var file_RogueTournHandBookNotify_proto_depIdxs = []int32{
	1, // 0: RogueTournHandBookNotify.rogue_tourn_handbook:type_name -> RogueTournHandbook
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_RogueTournHandBookNotify_proto_init() }
func file_RogueTournHandBookNotify_proto_init() {
	if File_RogueTournHandBookNotify_proto != nil {
		return
	}
	file_RogueTournHandbook_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueTournHandBookNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueTournHandBookNotify); i {
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
			RawDescriptor: file_RogueTournHandBookNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueTournHandBookNotify_proto_goTypes,
		DependencyIndexes: file_RogueTournHandBookNotify_proto_depIdxs,
		MessageInfos:      file_RogueTournHandBookNotify_proto_msgTypes,
	}.Build()
	File_RogueTournHandBookNotify_proto = out.File
	file_RogueTournHandBookNotify_proto_rawDesc = nil
	file_RogueTournHandBookNotify_proto_goTypes = nil
	file_RogueTournHandBookNotify_proto_depIdxs = nil
}
