// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: JKMJAIADMLI.proto

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

type JKMJAIADMLI int32

const (
	JKMJAIADMLI_SWORD_TRAINING_GAME_SETTLE_NONE          JKMJAIADMLI = 0
	JKMJAIADMLI_SWORD_TRAINING_GAME_SETTLE_FINISH        JKMJAIADMLI = 1
	JKMJAIADMLI_SWORD_TRAINING_GAME_SETTLE_GIVE_UP       JKMJAIADMLI = 2
	JKMJAIADMLI_SWORD_TRAINING_GAME_SETTLE_BATTLE_FAILED JKMJAIADMLI = 3
	JKMJAIADMLI_SWORD_TRAINING_GAME_SETTLE_FORCE         JKMJAIADMLI = 4
	JKMJAIADMLI_SWORD_TRAINING_GAME_SETTLE_BY_RESTORE    JKMJAIADMLI = 5
)

// Enum value maps for JKMJAIADMLI.
var (
	JKMJAIADMLI_name = map[int32]string{
		0: "SWORD_TRAINING_GAME_SETTLE_NONE",
		1: "SWORD_TRAINING_GAME_SETTLE_FINISH",
		2: "SWORD_TRAINING_GAME_SETTLE_GIVE_UP",
		3: "SWORD_TRAINING_GAME_SETTLE_BATTLE_FAILED",
		4: "SWORD_TRAINING_GAME_SETTLE_FORCE",
		5: "SWORD_TRAINING_GAME_SETTLE_BY_RESTORE",
	}
	JKMJAIADMLI_value = map[string]int32{
		"SWORD_TRAINING_GAME_SETTLE_NONE":          0,
		"SWORD_TRAINING_GAME_SETTLE_FINISH":        1,
		"SWORD_TRAINING_GAME_SETTLE_GIVE_UP":       2,
		"SWORD_TRAINING_GAME_SETTLE_BATTLE_FAILED": 3,
		"SWORD_TRAINING_GAME_SETTLE_FORCE":         4,
		"SWORD_TRAINING_GAME_SETTLE_BY_RESTORE":    5,
	}
)

func (x JKMJAIADMLI) Enum() *JKMJAIADMLI {
	p := new(JKMJAIADMLI)
	*p = x
	return p
}

func (x JKMJAIADMLI) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JKMJAIADMLI) Descriptor() protoreflect.EnumDescriptor {
	return file_JKMJAIADMLI_proto_enumTypes[0].Descriptor()
}

func (JKMJAIADMLI) Type() protoreflect.EnumType {
	return &file_JKMJAIADMLI_proto_enumTypes[0]
}

func (x JKMJAIADMLI) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JKMJAIADMLI.Descriptor instead.
func (JKMJAIADMLI) EnumDescriptor() ([]byte, []int) {
	return file_JKMJAIADMLI_proto_rawDescGZIP(), []int{0}
}

var File_JKMJAIADMLI_proto protoreflect.FileDescriptor

var file_JKMJAIADMLI_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4a, 0x4b, 0x4d, 0x4a, 0x41, 0x49, 0x41, 0x44, 0x4d, 0x4c, 0x49, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x80, 0x02, 0x0a, 0x0b, 0x4a, 0x4b, 0x4d, 0x4a, 0x41, 0x49, 0x41, 0x44,
	0x4d, 0x4c, 0x49, 0x12, 0x23, 0x0a, 0x1f, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x52, 0x41,
	0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x4c,
	0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x53, 0x57, 0x4f, 0x52,
	0x44, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f,
	0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x10, 0x01, 0x12,
	0x26, 0x0a, 0x22, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e,
	0x47, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x47, 0x49,
	0x56, 0x45, 0x5f, 0x55, 0x50, 0x10, 0x02, 0x12, 0x2c, 0x0a, 0x28, 0x53, 0x57, 0x4f, 0x52, 0x44,
	0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53,
	0x45, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x46, 0x41, 0x49,
	0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x24, 0x0a, 0x20, 0x53, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x54,
	0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x53, 0x45, 0x54,
	0x54, 0x4c, 0x45, 0x5f, 0x46, 0x4f, 0x52, 0x43, 0x45, 0x10, 0x04, 0x12, 0x29, 0x0a, 0x25, 0x53,
	0x57, 0x4f, 0x52, 0x44, 0x5f, 0x54, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x47, 0x41,
	0x4d, 0x45, 0x5f, 0x53, 0x45, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x42, 0x59, 0x5f, 0x52, 0x45, 0x53,
	0x54, 0x4f, 0x52, 0x45, 0x10, 0x05, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_JKMJAIADMLI_proto_rawDescOnce sync.Once
	file_JKMJAIADMLI_proto_rawDescData = file_JKMJAIADMLI_proto_rawDesc
)

func file_JKMJAIADMLI_proto_rawDescGZIP() []byte {
	file_JKMJAIADMLI_proto_rawDescOnce.Do(func() {
		file_JKMJAIADMLI_proto_rawDescData = protoimpl.X.CompressGZIP(file_JKMJAIADMLI_proto_rawDescData)
	})
	return file_JKMJAIADMLI_proto_rawDescData
}

var file_JKMJAIADMLI_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_JKMJAIADMLI_proto_goTypes = []interface{}{
	(JKMJAIADMLI)(0), // 0: JKMJAIADMLI
}
var file_JKMJAIADMLI_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_JKMJAIADMLI_proto_init() }
func file_JKMJAIADMLI_proto_init() {
	if File_JKMJAIADMLI_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_JKMJAIADMLI_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_JKMJAIADMLI_proto_goTypes,
		DependencyIndexes: file_JKMJAIADMLI_proto_depIdxs,
		EnumInfos:         file_JKMJAIADMLI_proto_enumTypes,
	}.Build()
	File_JKMJAIADMLI_proto = out.File
	file_JKMJAIADMLI_proto_rawDesc = nil
	file_JKMJAIADMLI_proto_goTypes = nil
	file_JKMJAIADMLI_proto_depIdxs = nil
}