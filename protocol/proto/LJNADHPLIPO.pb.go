// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: LJNADHPLIPO.proto

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

type LJNADHPLIPO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tid          uint32        `protobuf:"varint,13,opt,name=tid,proto3" json:"tid,omitempty"`
	Level        uint32        `protobuf:"varint,11,opt,name=level,proto3" json:"level,omitempty"`
	MainAffixId  uint32        `protobuf:"varint,4,opt,name=main_affix_id,json=mainAffixId,proto3" json:"main_affix_id,omitempty"`
	SubAffixList []*RelicAffix `protobuf:"bytes,8,rep,name=sub_affix_list,json=subAffixList,proto3" json:"sub_affix_list,omitempty"`
	Exp          uint32        `protobuf:"varint,6,opt,name=exp,proto3" json:"exp,omitempty"`
}

func (x *LJNADHPLIPO) Reset() {
	*x = LJNADHPLIPO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LJNADHPLIPO_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LJNADHPLIPO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LJNADHPLIPO) ProtoMessage() {}

func (x *LJNADHPLIPO) ProtoReflect() protoreflect.Message {
	mi := &file_LJNADHPLIPO_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LJNADHPLIPO.ProtoReflect.Descriptor instead.
func (*LJNADHPLIPO) Descriptor() ([]byte, []int) {
	return file_LJNADHPLIPO_proto_rawDescGZIP(), []int{0}
}

func (x *LJNADHPLIPO) GetTid() uint32 {
	if x != nil {
		return x.Tid
	}
	return 0
}

func (x *LJNADHPLIPO) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *LJNADHPLIPO) GetMainAffixId() uint32 {
	if x != nil {
		return x.MainAffixId
	}
	return 0
}

func (x *LJNADHPLIPO) GetSubAffixList() []*RelicAffix {
	if x != nil {
		return x.SubAffixList
	}
	return nil
}

func (x *LJNADHPLIPO) GetExp() uint32 {
	if x != nil {
		return x.Exp
	}
	return 0
}

var File_LJNADHPLIPO_proto protoreflect.FileDescriptor

var file_LJNADHPLIPO_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4c, 0x4a, 0x4e, 0x41, 0x44, 0x48, 0x50, 0x4c, 0x49, 0x50, 0x4f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x52, 0x65, 0x6c, 0x69, 0x63, 0x41, 0x66, 0x66, 0x69, 0x78, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x0b, 0x4c, 0x4a, 0x4e, 0x41, 0x44, 0x48,
	0x50, 0x4c, 0x49, 0x50, 0x4f, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x03, 0x74, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x22, 0x0a,
	0x0d, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x61, 0x66, 0x66, 0x69, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x61, 0x69, 0x6e, 0x41, 0x66, 0x66, 0x69, 0x78, 0x49,
	0x64, 0x12, 0x31, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x5f, 0x61, 0x66, 0x66, 0x69, 0x78, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x52, 0x65, 0x6c, 0x69,
	0x63, 0x41, 0x66, 0x66, 0x69, 0x78, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x41, 0x66, 0x66, 0x69, 0x78,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x03, 0x65, 0x78, 0x70, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_LJNADHPLIPO_proto_rawDescOnce sync.Once
	file_LJNADHPLIPO_proto_rawDescData = file_LJNADHPLIPO_proto_rawDesc
)

func file_LJNADHPLIPO_proto_rawDescGZIP() []byte {
	file_LJNADHPLIPO_proto_rawDescOnce.Do(func() {
		file_LJNADHPLIPO_proto_rawDescData = protoimpl.X.CompressGZIP(file_LJNADHPLIPO_proto_rawDescData)
	})
	return file_LJNADHPLIPO_proto_rawDescData
}

var file_LJNADHPLIPO_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_LJNADHPLIPO_proto_goTypes = []interface{}{
	(*LJNADHPLIPO)(nil), // 0: LJNADHPLIPO
	(*RelicAffix)(nil),  // 1: RelicAffix
}
var file_LJNADHPLIPO_proto_depIdxs = []int32{
	1, // 0: LJNADHPLIPO.sub_affix_list:type_name -> RelicAffix
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_LJNADHPLIPO_proto_init() }
func file_LJNADHPLIPO_proto_init() {
	if File_LJNADHPLIPO_proto != nil {
		return
	}
	file_RelicAffix_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LJNADHPLIPO_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LJNADHPLIPO); i {
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
			RawDescriptor: file_LJNADHPLIPO_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_LJNADHPLIPO_proto_goTypes,
		DependencyIndexes: file_LJNADHPLIPO_proto_depIdxs,
		MessageInfos:      file_LJNADHPLIPO_proto_msgTypes,
	}.Build()
	File_LJNADHPLIPO_proto = out.File
	file_LJNADHPLIPO_proto_rawDesc = nil
	file_LJNADHPLIPO_proto_goTypes = nil
	file_LJNADHPLIPO_proto_depIdxs = nil
}
