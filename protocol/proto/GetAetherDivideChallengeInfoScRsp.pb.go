// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetAetherDivideChallengeInfoScRsp.proto

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

type GetAetherDivideChallengeInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FIBDNGJJDFD         uint32   `protobuf:"varint,5,opt,name=FIBDNGJJDFD,proto3" json:"FIBDNGJJDFD,omitempty"`
	Retcode             uint32   `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
	EELDBMONEIM         []uint32 `protobuf:"varint,6,rep,packed,name=EELDBMONEIM,proto3" json:"EELDBMONEIM,omitempty"`
	FinishChallengeList []uint32 `protobuf:"varint,3,rep,packed,name=finish_challenge_list,json=finishChallengeList,proto3" json:"finish_challenge_list,omitempty"`
}

func (x *GetAetherDivideChallengeInfoScRsp) Reset() {
	*x = GetAetherDivideChallengeInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetAetherDivideChallengeInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAetherDivideChallengeInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAetherDivideChallengeInfoScRsp) ProtoMessage() {}

func (x *GetAetherDivideChallengeInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetAetherDivideChallengeInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAetherDivideChallengeInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetAetherDivideChallengeInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetAetherDivideChallengeInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetAetherDivideChallengeInfoScRsp) GetFIBDNGJJDFD() uint32 {
	if x != nil {
		return x.FIBDNGJJDFD
	}
	return 0
}

func (x *GetAetherDivideChallengeInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetAetherDivideChallengeInfoScRsp) GetEELDBMONEIM() []uint32 {
	if x != nil {
		return x.EELDBMONEIM
	}
	return nil
}

func (x *GetAetherDivideChallengeInfoScRsp) GetFinishChallengeList() []uint32 {
	if x != nil {
		return x.FinishChallengeList
	}
	return nil
}

var File_GetAetherDivideChallengeInfoScRsp_proto protoreflect.FileDescriptor

var file_GetAetherDivideChallengeInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x27, 0x47, 0x65, 0x74, 0x41, 0x65, 0x74, 0x68, 0x65, 0x72, 0x44, 0x69, 0x76, 0x69, 0x64,
	0x65, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63,
	0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb5, 0x01, 0x0a, 0x21, 0x47, 0x65,
	0x74, 0x41, 0x65, 0x74, 0x68, 0x65, 0x72, 0x44, 0x69, 0x76, 0x69, 0x64, 0x65, 0x43, 0x68, 0x61,
	0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12,
	0x20, 0x0a, 0x0b, 0x46, 0x49, 0x42, 0x44, 0x4e, 0x47, 0x4a, 0x4a, 0x44, 0x46, 0x44, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x46, 0x49, 0x42, 0x44, 0x4e, 0x47, 0x4a, 0x4a, 0x44, 0x46,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x45,
	0x45, 0x4c, 0x44, 0x42, 0x4d, 0x4f, 0x4e, 0x45, 0x49, 0x4d, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0b, 0x45, 0x45, 0x4c, 0x44, 0x42, 0x4d, 0x4f, 0x4e, 0x45, 0x49, 0x4d, 0x12, 0x32, 0x0a,
	0x15, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x13, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetAetherDivideChallengeInfoScRsp_proto_rawDescOnce sync.Once
	file_GetAetherDivideChallengeInfoScRsp_proto_rawDescData = file_GetAetherDivideChallengeInfoScRsp_proto_rawDesc
)

func file_GetAetherDivideChallengeInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetAetherDivideChallengeInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetAetherDivideChallengeInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetAetherDivideChallengeInfoScRsp_proto_rawDescData)
	})
	return file_GetAetherDivideChallengeInfoScRsp_proto_rawDescData
}

var file_GetAetherDivideChallengeInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetAetherDivideChallengeInfoScRsp_proto_goTypes = []interface{}{
	(*GetAetherDivideChallengeInfoScRsp)(nil), // 0: GetAetherDivideChallengeInfoScRsp
}
var file_GetAetherDivideChallengeInfoScRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GetAetherDivideChallengeInfoScRsp_proto_init() }
func file_GetAetherDivideChallengeInfoScRsp_proto_init() {
	if File_GetAetherDivideChallengeInfoScRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GetAetherDivideChallengeInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAetherDivideChallengeInfoScRsp); i {
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
			RawDescriptor: file_GetAetherDivideChallengeInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetAetherDivideChallengeInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetAetherDivideChallengeInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetAetherDivideChallengeInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetAetherDivideChallengeInfoScRsp_proto = out.File
	file_GetAetherDivideChallengeInfoScRsp_proto_rawDesc = nil
	file_GetAetherDivideChallengeInfoScRsp_proto_goTypes = nil
	file_GetAetherDivideChallengeInfoScRsp_proto_depIdxs = nil
}
