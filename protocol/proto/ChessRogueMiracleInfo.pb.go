// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueMiracleInfo.proto

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

type ChessRogueMiracleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChessRogueMiracleInfo *ChessRogueMiracle `protobuf:"bytes,5,opt,name=chess_rogue_miracle_info,json=chessRogueMiracleInfo,proto3" json:"chess_rogue_miracle_info,omitempty"`
}

func (x *ChessRogueMiracleInfo) Reset() {
	*x = ChessRogueMiracleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChessRogueMiracleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChessRogueMiracleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChessRogueMiracleInfo) ProtoMessage() {}

func (x *ChessRogueMiracleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ChessRogueMiracleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChessRogueMiracleInfo.ProtoReflect.Descriptor instead.
func (*ChessRogueMiracleInfo) Descriptor() ([]byte, []int) {
	return file_ChessRogueMiracleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *ChessRogueMiracleInfo) GetChessRogueMiracleInfo() *ChessRogueMiracle {
	if x != nil {
		return x.ChessRogueMiracleInfo
	}
	return nil
}

var File_ChessRogueMiracleInfo_proto protoreflect.FileDescriptor

var file_ChessRogueMiracleInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61,
	0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x43,
	0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x4b, 0x0a, 0x18, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x6d,
	0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4d, 0x69,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x15, 0x63, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75,
	0x65, 0x4d, 0x69, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueMiracleInfo_proto_rawDescOnce sync.Once
	file_ChessRogueMiracleInfo_proto_rawDescData = file_ChessRogueMiracleInfo_proto_rawDesc
)

func file_ChessRogueMiracleInfo_proto_rawDescGZIP() []byte {
	file_ChessRogueMiracleInfo_proto_rawDescOnce.Do(func() {
		file_ChessRogueMiracleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueMiracleInfo_proto_rawDescData)
	})
	return file_ChessRogueMiracleInfo_proto_rawDescData
}

var file_ChessRogueMiracleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChessRogueMiracleInfo_proto_goTypes = []interface{}{
	(*ChessRogueMiracleInfo)(nil), // 0: ChessRogueMiracleInfo
	(*ChessRogueMiracle)(nil),     // 1: ChessRogueMiracle
}
var file_ChessRogueMiracleInfo_proto_depIdxs = []int32{
	1, // 0: ChessRogueMiracleInfo.chess_rogue_miracle_info:type_name -> ChessRogueMiracle
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ChessRogueMiracleInfo_proto_init() }
func file_ChessRogueMiracleInfo_proto_init() {
	if File_ChessRogueMiracleInfo_proto != nil {
		return
	}
	file_ChessRogueMiracle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChessRogueMiracleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChessRogueMiracleInfo); i {
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
			RawDescriptor: file_ChessRogueMiracleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueMiracleInfo_proto_goTypes,
		DependencyIndexes: file_ChessRogueMiracleInfo_proto_depIdxs,
		MessageInfos:      file_ChessRogueMiracleInfo_proto_msgTypes,
	}.Build()
	File_ChessRogueMiracleInfo_proto = out.File
	file_ChessRogueMiracleInfo_proto_rawDesc = nil
	file_ChessRogueMiracleInfo_proto_goTypes = nil
	file_ChessRogueMiracleInfo_proto_depIdxs = nil
}
