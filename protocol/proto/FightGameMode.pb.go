// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FightGameMode.proto

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

type FightGameMode int32

const (
	FightGameMode_FIGHT_GAME_MODE_NONE   FightGameMode = 0
	FightGameMode_FIGHT_GAME_MODE_MATCH3 FightGameMode = 1
)

// Enum value maps for FightGameMode.
var (
	FightGameMode_name = map[int32]string{
		0: "FIGHT_GAME_MODE_NONE",
		1: "FIGHT_GAME_MODE_MATCH3",
	}
	FightGameMode_value = map[string]int32{
		"FIGHT_GAME_MODE_NONE":   0,
		"FIGHT_GAME_MODE_MATCH3": 1,
	}
)

func (x FightGameMode) Enum() *FightGameMode {
	p := new(FightGameMode)
	*p = x
	return p
}

func (x FightGameMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FightGameMode) Descriptor() protoreflect.EnumDescriptor {
	return file_FightGameMode_proto_enumTypes[0].Descriptor()
}

func (FightGameMode) Type() protoreflect.EnumType {
	return &file_FightGameMode_proto_enumTypes[0]
}

func (x FightGameMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FightGameMode.Descriptor instead.
func (FightGameMode) EnumDescriptor() ([]byte, []int) {
	return file_FightGameMode_proto_rawDescGZIP(), []int{0}
}

var File_FightGameMode_proto protoreflect.FileDescriptor

var file_FightGameMode_proto_rawDesc = []byte{
	0x0a, 0x13, 0x46, 0x69, 0x67, 0x68, 0x74, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x45, 0x0a, 0x0d, 0x46, 0x69, 0x67, 0x68, 0x74, 0x47, 0x61,
	0x6d, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x49, 0x47, 0x48, 0x54, 0x5f,
	0x47, 0x41, 0x4d, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00,
	0x12, 0x1a, 0x0a, 0x16, 0x46, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x4d,
	0x4f, 0x44, 0x45, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x33, 0x10, 0x01, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FightGameMode_proto_rawDescOnce sync.Once
	file_FightGameMode_proto_rawDescData = file_FightGameMode_proto_rawDesc
)

func file_FightGameMode_proto_rawDescGZIP() []byte {
	file_FightGameMode_proto_rawDescOnce.Do(func() {
		file_FightGameMode_proto_rawDescData = protoimpl.X.CompressGZIP(file_FightGameMode_proto_rawDescData)
	})
	return file_FightGameMode_proto_rawDescData
}

var file_FightGameMode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_FightGameMode_proto_goTypes = []interface{}{
	(FightGameMode)(0), // 0: FightGameMode
}
var file_FightGameMode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_FightGameMode_proto_init() }
func file_FightGameMode_proto_init() {
	if File_FightGameMode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FightGameMode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FightGameMode_proto_goTypes,
		DependencyIndexes: file_FightGameMode_proto_depIdxs,
		EnumInfos:         file_FightGameMode_proto_enumTypes,
	}.Build()
	File_FightGameMode_proto = out.File
	file_FightGameMode_proto_rawDesc = nil
	file_FightGameMode_proto_goTypes = nil
	file_FightGameMode_proto_depIdxs = nil
}
