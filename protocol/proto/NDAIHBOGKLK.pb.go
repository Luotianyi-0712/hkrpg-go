// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: NDAIHBOGKLK.proto

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

type NDAIHBOGKLK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarIdList []uint32 `protobuf:"varint,13,rep,packed,name=avatar_id_list,json=avatarIdList,proto3" json:"avatar_id_list,omitempty"`
	Id           uint32   `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
	MFGNACAOBNB  int64    `protobuf:"varint,10,opt,name=MFGNACAOBNB,proto3" json:"MFGNACAOBNB,omitempty"`
	FNPOIHMEPFA  uint32   `protobuf:"varint,5,opt,name=FNPOIHMEPFA,proto3" json:"FNPOIHMEPFA,omitempty"`
}

func (x *NDAIHBOGKLK) Reset() {
	*x = NDAIHBOGKLK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_NDAIHBOGKLK_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NDAIHBOGKLK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NDAIHBOGKLK) ProtoMessage() {}

func (x *NDAIHBOGKLK) ProtoReflect() protoreflect.Message {
	mi := &file_NDAIHBOGKLK_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NDAIHBOGKLK.ProtoReflect.Descriptor instead.
func (*NDAIHBOGKLK) Descriptor() ([]byte, []int) {
	return file_NDAIHBOGKLK_proto_rawDescGZIP(), []int{0}
}

func (x *NDAIHBOGKLK) GetAvatarIdList() []uint32 {
	if x != nil {
		return x.AvatarIdList
	}
	return nil
}

func (x *NDAIHBOGKLK) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NDAIHBOGKLK) GetMFGNACAOBNB() int64 {
	if x != nil {
		return x.MFGNACAOBNB
	}
	return 0
}

func (x *NDAIHBOGKLK) GetFNPOIHMEPFA() uint32 {
	if x != nil {
		return x.FNPOIHMEPFA
	}
	return 0
}

var File_NDAIHBOGKLK_proto protoreflect.FileDescriptor

var file_NDAIHBOGKLK_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4e, 0x44, 0x41, 0x49, 0x48, 0x42, 0x4f, 0x47, 0x4b, 0x4c, 0x4b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x0b, 0x4e, 0x44, 0x41, 0x49, 0x48, 0x42, 0x4f, 0x47,
	0x4b, 0x4c, 0x4b, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0c, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x46, 0x47,
	0x4e, 0x41, 0x43, 0x41, 0x4f, 0x42, 0x4e, 0x42, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x4d, 0x46, 0x47, 0x4e, 0x41, 0x43, 0x41, 0x4f, 0x42, 0x4e, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x46,
	0x4e, 0x50, 0x4f, 0x49, 0x48, 0x4d, 0x45, 0x50, 0x46, 0x41, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x46, 0x4e, 0x50, 0x4f, 0x49, 0x48, 0x4d, 0x45, 0x50, 0x46, 0x41, 0x42, 0x2e, 0x5a,
	0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_NDAIHBOGKLK_proto_rawDescOnce sync.Once
	file_NDAIHBOGKLK_proto_rawDescData = file_NDAIHBOGKLK_proto_rawDesc
)

func file_NDAIHBOGKLK_proto_rawDescGZIP() []byte {
	file_NDAIHBOGKLK_proto_rawDescOnce.Do(func() {
		file_NDAIHBOGKLK_proto_rawDescData = protoimpl.X.CompressGZIP(file_NDAIHBOGKLK_proto_rawDescData)
	})
	return file_NDAIHBOGKLK_proto_rawDescData
}

var file_NDAIHBOGKLK_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_NDAIHBOGKLK_proto_goTypes = []interface{}{
	(*NDAIHBOGKLK)(nil), // 0: NDAIHBOGKLK
}
var file_NDAIHBOGKLK_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_NDAIHBOGKLK_proto_init() }
func file_NDAIHBOGKLK_proto_init() {
	if File_NDAIHBOGKLK_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_NDAIHBOGKLK_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NDAIHBOGKLK); i {
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
			RawDescriptor: file_NDAIHBOGKLK_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_NDAIHBOGKLK_proto_goTypes,
		DependencyIndexes: file_NDAIHBOGKLK_proto_depIdxs,
		MessageInfos:      file_NDAIHBOGKLK_proto_msgTypes,
	}.Build()
	File_NDAIHBOGKLK_proto = out.File
	file_NDAIHBOGKLK_proto_rawDesc = nil
	file_NDAIHBOGKLK_proto_goTypes = nil
	file_NDAIHBOGKLK_proto_depIdxs = nil
}
