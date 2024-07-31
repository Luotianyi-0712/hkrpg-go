// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EvolveBuildStartLevelCsReq.proto

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

type EvolveBuildStartLevelCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AvatarList  []*EvolveBuildAvatar `protobuf:"bytes,1,rep,name=avatar_list,json=avatarList,proto3" json:"avatar_list,omitempty"`
	LAAFPDHGMBG uint32               `protobuf:"varint,15,opt,name=LAAFPDHGMBG,proto3" json:"LAAFPDHGMBG,omitempty"`
	EALPPJEDKMH *AAOEPMKPNOK         `protobuf:"bytes,12,opt,name=EALPPJEDKMH,proto3" json:"EALPPJEDKMH,omitempty"`
}

func (x *EvolveBuildStartLevelCsReq) Reset() {
	*x = EvolveBuildStartLevelCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EvolveBuildStartLevelCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvolveBuildStartLevelCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvolveBuildStartLevelCsReq) ProtoMessage() {}

func (x *EvolveBuildStartLevelCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_EvolveBuildStartLevelCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvolveBuildStartLevelCsReq.ProtoReflect.Descriptor instead.
func (*EvolveBuildStartLevelCsReq) Descriptor() ([]byte, []int) {
	return file_EvolveBuildStartLevelCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *EvolveBuildStartLevelCsReq) GetAvatarList() []*EvolveBuildAvatar {
	if x != nil {
		return x.AvatarList
	}
	return nil
}

func (x *EvolveBuildStartLevelCsReq) GetLAAFPDHGMBG() uint32 {
	if x != nil {
		return x.LAAFPDHGMBG
	}
	return 0
}

func (x *EvolveBuildStartLevelCsReq) GetEALPPJEDKMH() *AAOEPMKPNOK {
	if x != nil {
		return x.EALPPJEDKMH
	}
	return nil
}

var File_EvolveBuildStartLevelCsReq_proto protoreflect.FileDescriptor

var file_EvolveBuildStartLevelCsReq_proto_rawDesc = []byte{
	0x0a, 0x20, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x41,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x41, 0x4f,
	0x45, 0x50, 0x4d, 0x4b, 0x50, 0x4e, 0x4f, 0x4b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3,
	0x01, 0x0a, 0x1a, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x33, 0x0a,
	0x0b, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x45, 0x76, 0x6f, 0x6c, 0x76, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x52, 0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x41, 0x41, 0x46, 0x50, 0x44, 0x48, 0x47, 0x4d, 0x42,
	0x47, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x41, 0x41, 0x46, 0x50, 0x44, 0x48,
	0x47, 0x4d, 0x42, 0x47, 0x12, 0x2e, 0x0a, 0x0b, 0x45, 0x41, 0x4c, 0x50, 0x50, 0x4a, 0x45, 0x44,
	0x4b, 0x4d, 0x48, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x41, 0x41, 0x4f, 0x45,
	0x50, 0x4d, 0x4b, 0x50, 0x4e, 0x4f, 0x4b, 0x52, 0x0b, 0x45, 0x41, 0x4c, 0x50, 0x50, 0x4a, 0x45,
	0x44, 0x4b, 0x4d, 0x48, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EvolveBuildStartLevelCsReq_proto_rawDescOnce sync.Once
	file_EvolveBuildStartLevelCsReq_proto_rawDescData = file_EvolveBuildStartLevelCsReq_proto_rawDesc
)

func file_EvolveBuildStartLevelCsReq_proto_rawDescGZIP() []byte {
	file_EvolveBuildStartLevelCsReq_proto_rawDescOnce.Do(func() {
		file_EvolveBuildStartLevelCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_EvolveBuildStartLevelCsReq_proto_rawDescData)
	})
	return file_EvolveBuildStartLevelCsReq_proto_rawDescData
}

var file_EvolveBuildStartLevelCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EvolveBuildStartLevelCsReq_proto_goTypes = []interface{}{
	(*EvolveBuildStartLevelCsReq)(nil), // 0: EvolveBuildStartLevelCsReq
	(*EvolveBuildAvatar)(nil),          // 1: EvolveBuildAvatar
	(*AAOEPMKPNOK)(nil),                // 2: AAOEPMKPNOK
}
var file_EvolveBuildStartLevelCsReq_proto_depIdxs = []int32{
	1, // 0: EvolveBuildStartLevelCsReq.avatar_list:type_name -> EvolveBuildAvatar
	2, // 1: EvolveBuildStartLevelCsReq.EALPPJEDKMH:type_name -> AAOEPMKPNOK
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_EvolveBuildStartLevelCsReq_proto_init() }
func file_EvolveBuildStartLevelCsReq_proto_init() {
	if File_EvolveBuildStartLevelCsReq_proto != nil {
		return
	}
	file_EvolveBuildAvatar_proto_init()
	file_AAOEPMKPNOK_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EvolveBuildStartLevelCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EvolveBuildStartLevelCsReq); i {
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
			RawDescriptor: file_EvolveBuildStartLevelCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EvolveBuildStartLevelCsReq_proto_goTypes,
		DependencyIndexes: file_EvolveBuildStartLevelCsReq_proto_depIdxs,
		MessageInfos:      file_EvolveBuildStartLevelCsReq_proto_msgTypes,
	}.Build()
	File_EvolveBuildStartLevelCsReq_proto = out.File
	file_EvolveBuildStartLevelCsReq_proto_rawDesc = nil
	file_EvolveBuildStartLevelCsReq_proto_goTypes = nil
	file_EvolveBuildStartLevelCsReq_proto_depIdxs = nil
}
