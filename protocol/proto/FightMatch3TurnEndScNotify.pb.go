// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FightMatch3TurnEndScNotify.proto

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

type FightMatch3TurnEndScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NCAJNCEAJHB *HPCDLEMPBEK `protobuf:"bytes,4,opt,name=NCAJNCEAJHB,proto3" json:"NCAJNCEAJHB,omitempty"`
	DICPLJNDHHG *HPCDLEMPBEK `protobuf:"bytes,10,opt,name=DICPLJNDHHG,proto3" json:"DICPLJNDHHG,omitempty"`
}

func (x *FightMatch3TurnEndScNotify) Reset() {
	*x = FightMatch3TurnEndScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FightMatch3TurnEndScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FightMatch3TurnEndScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FightMatch3TurnEndScNotify) ProtoMessage() {}

func (x *FightMatch3TurnEndScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_FightMatch3TurnEndScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FightMatch3TurnEndScNotify.ProtoReflect.Descriptor instead.
func (*FightMatch3TurnEndScNotify) Descriptor() ([]byte, []int) {
	return file_FightMatch3TurnEndScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *FightMatch3TurnEndScNotify) GetNCAJNCEAJHB() *HPCDLEMPBEK {
	if x != nil {
		return x.NCAJNCEAJHB
	}
	return nil
}

func (x *FightMatch3TurnEndScNotify) GetDICPLJNDHHG() *HPCDLEMPBEK {
	if x != nil {
		return x.DICPLJNDHHG
	}
	return nil
}

var File_FightMatch3TurnEndScNotify_proto protoreflect.FileDescriptor

var file_FightMatch3TurnEndScNotify_proto_rawDesc = []byte{
	0x0a, 0x20, 0x46, 0x69, 0x67, 0x68, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x33, 0x54, 0x75, 0x72,
	0x6e, 0x45, 0x6e, 0x64, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x48, 0x50, 0x43, 0x44, 0x4c, 0x45, 0x4d, 0x50, 0x42, 0x45, 0x4b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x1a, 0x46, 0x69, 0x67, 0x68, 0x74, 0x4d, 0x61,
	0x74, 0x63, 0x68, 0x33, 0x54, 0x75, 0x72, 0x6e, 0x45, 0x6e, 0x64, 0x53, 0x63, 0x4e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x12, 0x2e, 0x0a, 0x0b, 0x4e, 0x43, 0x41, 0x4a, 0x4e, 0x43, 0x45, 0x41, 0x4a,
	0x48, 0x42, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x50, 0x43, 0x44, 0x4c,
	0x45, 0x4d, 0x50, 0x42, 0x45, 0x4b, 0x52, 0x0b, 0x4e, 0x43, 0x41, 0x4a, 0x4e, 0x43, 0x45, 0x41,
	0x4a, 0x48, 0x42, 0x12, 0x2e, 0x0a, 0x0b, 0x44, 0x49, 0x43, 0x50, 0x4c, 0x4a, 0x4e, 0x44, 0x48,
	0x48, 0x47, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x50, 0x43, 0x44, 0x4c,
	0x45, 0x4d, 0x50, 0x42, 0x45, 0x4b, 0x52, 0x0b, 0x44, 0x49, 0x43, 0x50, 0x4c, 0x4a, 0x4e, 0x44,
	0x48, 0x48, 0x47, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FightMatch3TurnEndScNotify_proto_rawDescOnce sync.Once
	file_FightMatch3TurnEndScNotify_proto_rawDescData = file_FightMatch3TurnEndScNotify_proto_rawDesc
)

func file_FightMatch3TurnEndScNotify_proto_rawDescGZIP() []byte {
	file_FightMatch3TurnEndScNotify_proto_rawDescOnce.Do(func() {
		file_FightMatch3TurnEndScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_FightMatch3TurnEndScNotify_proto_rawDescData)
	})
	return file_FightMatch3TurnEndScNotify_proto_rawDescData
}

var file_FightMatch3TurnEndScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FightMatch3TurnEndScNotify_proto_goTypes = []interface{}{
	(*FightMatch3TurnEndScNotify)(nil), // 0: FightMatch3TurnEndScNotify
	(*HPCDLEMPBEK)(nil),                // 1: HPCDLEMPBEK
}
var file_FightMatch3TurnEndScNotify_proto_depIdxs = []int32{
	1, // 0: FightMatch3TurnEndScNotify.NCAJNCEAJHB:type_name -> HPCDLEMPBEK
	1, // 1: FightMatch3TurnEndScNotify.DICPLJNDHHG:type_name -> HPCDLEMPBEK
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_FightMatch3TurnEndScNotify_proto_init() }
func file_FightMatch3TurnEndScNotify_proto_init() {
	if File_FightMatch3TurnEndScNotify_proto != nil {
		return
	}
	file_HPCDLEMPBEK_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FightMatch3TurnEndScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FightMatch3TurnEndScNotify); i {
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
			RawDescriptor: file_FightMatch3TurnEndScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FightMatch3TurnEndScNotify_proto_goTypes,
		DependencyIndexes: file_FightMatch3TurnEndScNotify_proto_depIdxs,
		MessageInfos:      file_FightMatch3TurnEndScNotify_proto_msgTypes,
	}.Build()
	File_FightMatch3TurnEndScNotify_proto = out.File
	file_FightMatch3TurnEndScNotify_proto_rawDesc = nil
	file_FightMatch3TurnEndScNotify_proto_goTypes = nil
	file_FightMatch3TurnEndScNotify_proto_depIdxs = nil
}
