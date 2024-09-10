// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SetAssistAvatarScRsp.proto

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

type SetAssistAvatarScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode      uint32   `protobuf:"varint,3,opt,name=retcode,proto3" json:"retcode,omitempty"`
	AvatarId     uint32   `protobuf:"varint,9,opt,name=avatar_id,json=avatarId,proto3" json:"avatar_id,omitempty"`
	AvatarIdList []uint32 `protobuf:"varint,7,rep,packed,name=avatar_id_list,json=avatarIdList,proto3" json:"avatar_id_list,omitempty"`
}

func (x *SetAssistAvatarScRsp) Reset() {
	*x = SetAssistAvatarScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetAssistAvatarScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAssistAvatarScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAssistAvatarScRsp) ProtoMessage() {}

func (x *SetAssistAvatarScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SetAssistAvatarScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAssistAvatarScRsp.ProtoReflect.Descriptor instead.
func (*SetAssistAvatarScRsp) Descriptor() ([]byte, []int) {
	return file_SetAssistAvatarScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SetAssistAvatarScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *SetAssistAvatarScRsp) GetAvatarId() uint32 {
	if x != nil {
		return x.AvatarId
	}
	return 0
}

func (x *SetAssistAvatarScRsp) GetAvatarIdList() []uint32 {
	if x != nil {
		return x.AvatarIdList
	}
	return nil
}

var File_SetAssistAvatarScRsp_proto protoreflect.FileDescriptor

var file_SetAssistAvatarScRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x53, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x14,
	0x53, 0x65, 0x74, 0x41, 0x73, 0x73, 0x69, 0x73, 0x74, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x07, 0x20,
	0x03, 0x28, 0x0d, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x49, 0x64, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SetAssistAvatarScRsp_proto_rawDescOnce sync.Once
	file_SetAssistAvatarScRsp_proto_rawDescData = file_SetAssistAvatarScRsp_proto_rawDesc
)

func file_SetAssistAvatarScRsp_proto_rawDescGZIP() []byte {
	file_SetAssistAvatarScRsp_proto_rawDescOnce.Do(func() {
		file_SetAssistAvatarScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetAssistAvatarScRsp_proto_rawDescData)
	})
	return file_SetAssistAvatarScRsp_proto_rawDescData
}

var file_SetAssistAvatarScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetAssistAvatarScRsp_proto_goTypes = []interface{}{
	(*SetAssistAvatarScRsp)(nil), // 0: SetAssistAvatarScRsp
}
var file_SetAssistAvatarScRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SetAssistAvatarScRsp_proto_init() }
func file_SetAssistAvatarScRsp_proto_init() {
	if File_SetAssistAvatarScRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SetAssistAvatarScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAssistAvatarScRsp); i {
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
			RawDescriptor: file_SetAssistAvatarScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetAssistAvatarScRsp_proto_goTypes,
		DependencyIndexes: file_SetAssistAvatarScRsp_proto_depIdxs,
		MessageInfos:      file_SetAssistAvatarScRsp_proto_msgTypes,
	}.Build()
	File_SetAssistAvatarScRsp_proto = out.File
	file_SetAssistAvatarScRsp_proto_rawDesc = nil
	file_SetAssistAvatarScRsp_proto_goTypes = nil
	file_SetAssistAvatarScRsp_proto_depIdxs = nil
}
