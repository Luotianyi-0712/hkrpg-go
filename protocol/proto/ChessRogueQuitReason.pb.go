// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChessRogueQuitReason.proto

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

type ChessRogueQuitReason int32

const (
	ChessRogueQuitReason_CHESS_ROGUE_ACCOUNT_BY_NONE          ChessRogueQuitReason = 0
	ChessRogueQuitReason_CHESS_ROGUE_ACCOUNT_BY_NORMAL_FINISH ChessRogueQuitReason = 1
	ChessRogueQuitReason_CHESS_ROGUE_ACCOUNT_BY_NORMAL_QUIT   ChessRogueQuitReason = 2
	ChessRogueQuitReason_CHESS_ROGUE_ACCOUNT_BY_DIALOG        ChessRogueQuitReason = 3
	ChessRogueQuitReason_CHESS_ROGUE_ACCOUNT_BY_FAILED        ChessRogueQuitReason = 4
	ChessRogueQuitReason_CHESS_ROGUE_ACCOUNT_BY_CUSTOM_OP     ChessRogueQuitReason = 5
)

// Enum value maps for ChessRogueQuitReason.
var (
	ChessRogueQuitReason_name = map[int32]string{
		0: "CHESS_ROGUE_ACCOUNT_BY_NONE",
		1: "CHESS_ROGUE_ACCOUNT_BY_NORMAL_FINISH",
		2: "CHESS_ROGUE_ACCOUNT_BY_NORMAL_QUIT",
		3: "CHESS_ROGUE_ACCOUNT_BY_DIALOG",
		4: "CHESS_ROGUE_ACCOUNT_BY_FAILED",
		5: "CHESS_ROGUE_ACCOUNT_BY_CUSTOM_OP",
	}
	ChessRogueQuitReason_value = map[string]int32{
		"CHESS_ROGUE_ACCOUNT_BY_NONE":          0,
		"CHESS_ROGUE_ACCOUNT_BY_NORMAL_FINISH": 1,
		"CHESS_ROGUE_ACCOUNT_BY_NORMAL_QUIT":   2,
		"CHESS_ROGUE_ACCOUNT_BY_DIALOG":        3,
		"CHESS_ROGUE_ACCOUNT_BY_FAILED":        4,
		"CHESS_ROGUE_ACCOUNT_BY_CUSTOM_OP":     5,
	}
)

func (x ChessRogueQuitReason) Enum() *ChessRogueQuitReason {
	p := new(ChessRogueQuitReason)
	*p = x
	return p
}

func (x ChessRogueQuitReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChessRogueQuitReason) Descriptor() protoreflect.EnumDescriptor {
	return file_ChessRogueQuitReason_proto_enumTypes[0].Descriptor()
}

func (ChessRogueQuitReason) Type() protoreflect.EnumType {
	return &file_ChessRogueQuitReason_proto_enumTypes[0]
}

func (x ChessRogueQuitReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChessRogueQuitReason.Descriptor instead.
func (ChessRogueQuitReason) EnumDescriptor() ([]byte, []int) {
	return file_ChessRogueQuitReason_proto_rawDescGZIP(), []int{0}
}

var File_ChessRogueQuitReason_proto protoreflect.FileDescriptor

var file_ChessRogueQuitReason_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x69, 0x74,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xf5, 0x01, 0x0a,
	0x14, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x51, 0x75, 0x69, 0x74, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x48, 0x45, 0x53, 0x53, 0x5f, 0x52,
	0x4f, 0x47, 0x55, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x42, 0x59, 0x5f,
	0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x28, 0x0a, 0x24, 0x43, 0x48, 0x45, 0x53, 0x53, 0x5f,
	0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x42, 0x59,
	0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x10, 0x01,
	0x12, 0x26, 0x0a, 0x22, 0x43, 0x48, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f,
	0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41,
	0x4c, 0x5f, 0x51, 0x55, 0x49, 0x54, 0x10, 0x02, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x48, 0x45, 0x53,
	0x53, 0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f,
	0x42, 0x59, 0x5f, 0x44, 0x49, 0x41, 0x4c, 0x4f, 0x47, 0x10, 0x03, 0x12, 0x21, 0x0a, 0x1d, 0x43,
	0x48, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x24,
	0x0a, 0x20, 0x43, 0x48, 0x45, 0x53, 0x53, 0x5f, 0x52, 0x4f, 0x47, 0x55, 0x45, 0x5f, 0x41, 0x43,
	0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x42, 0x59, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x5f,
	0x4f, 0x50, 0x10, 0x05, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChessRogueQuitReason_proto_rawDescOnce sync.Once
	file_ChessRogueQuitReason_proto_rawDescData = file_ChessRogueQuitReason_proto_rawDesc
)

func file_ChessRogueQuitReason_proto_rawDescGZIP() []byte {
	file_ChessRogueQuitReason_proto_rawDescOnce.Do(func() {
		file_ChessRogueQuitReason_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChessRogueQuitReason_proto_rawDescData)
	})
	return file_ChessRogueQuitReason_proto_rawDescData
}

var file_ChessRogueQuitReason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_ChessRogueQuitReason_proto_goTypes = []interface{}{
	(ChessRogueQuitReason)(0), // 0: ChessRogueQuitReason
}
var file_ChessRogueQuitReason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ChessRogueQuitReason_proto_init() }
func file_ChessRogueQuitReason_proto_init() {
	if File_ChessRogueQuitReason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ChessRogueQuitReason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChessRogueQuitReason_proto_goTypes,
		DependencyIndexes: file_ChessRogueQuitReason_proto_depIdxs,
		EnumInfos:         file_ChessRogueQuitReason_proto_enumTypes,
	}.Build()
	File_ChessRogueQuitReason_proto = out.File
	file_ChessRogueQuitReason_proto_rawDesc = nil
	file_ChessRogueQuitReason_proto_goTypes = nil
	file_ChessRogueQuitReason_proto_depIdxs = nil
}
