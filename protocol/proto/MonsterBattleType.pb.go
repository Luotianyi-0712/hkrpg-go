// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MonsterBattleType.proto

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

type MonsterBattleType int32

const (
	MonsterBattleType_MONSTER_BATTLE_TYPE_NONE                       MonsterBattleType = 0
	MonsterBattleType_MONSTER_BATTLE_TYPE_TRIGGER_BATTLE             MonsterBattleType = 1
	MonsterBattleType_MONSTER_BATTLE_TYPE_DIRECT_DIE_SIMULATE_BATTLE MonsterBattleType = 2
	MonsterBattleType_MONSTER_BATTLE_TYPE_DIRECT_DIE_SKIP_BATTLE     MonsterBattleType = 3
	MonsterBattleType_MONSTER_BATTLE_TYPE_NO_BATTLE                  MonsterBattleType = 4
)

// Enum value maps for MonsterBattleType.
var (
	MonsterBattleType_name = map[int32]string{
		0: "MONSTER_BATTLE_TYPE_NONE",
		1: "MONSTER_BATTLE_TYPE_TRIGGER_BATTLE",
		2: "MONSTER_BATTLE_TYPE_DIRECT_DIE_SIMULATE_BATTLE",
		3: "MONSTER_BATTLE_TYPE_DIRECT_DIE_SKIP_BATTLE",
		4: "MONSTER_BATTLE_TYPE_NO_BATTLE",
	}
	MonsterBattleType_value = map[string]int32{
		"MONSTER_BATTLE_TYPE_NONE":                       0,
		"MONSTER_BATTLE_TYPE_TRIGGER_BATTLE":             1,
		"MONSTER_BATTLE_TYPE_DIRECT_DIE_SIMULATE_BATTLE": 2,
		"MONSTER_BATTLE_TYPE_DIRECT_DIE_SKIP_BATTLE":     3,
		"MONSTER_BATTLE_TYPE_NO_BATTLE":                  4,
	}
)

func (x MonsterBattleType) Enum() *MonsterBattleType {
	p := new(MonsterBattleType)
	*p = x
	return p
}

func (x MonsterBattleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MonsterBattleType) Descriptor() protoreflect.EnumDescriptor {
	return file_MonsterBattleType_proto_enumTypes[0].Descriptor()
}

func (MonsterBattleType) Type() protoreflect.EnumType {
	return &file_MonsterBattleType_proto_enumTypes[0]
}

func (x MonsterBattleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MonsterBattleType.Descriptor instead.
func (MonsterBattleType) EnumDescriptor() ([]byte, []int) {
	return file_MonsterBattleType_proto_rawDescGZIP(), []int{0}
}

var File_MonsterBattleType_proto protoreflect.FileDescriptor

var file_MonsterBattleType_proto_rawDesc = []byte{
	0x0a, 0x17, 0x4d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xe0, 0x01, 0x0a, 0x11, 0x4d, 0x6f,
	0x6e, 0x73, 0x74, 0x65, 0x72, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1c, 0x0a, 0x18, 0x4d, 0x4f, 0x4e, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x26, 0x0a,
	0x22, 0x4d, 0x4f, 0x4e, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x52, 0x49, 0x47, 0x47, 0x45, 0x52, 0x5f, 0x42, 0x41, 0x54,
	0x54, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x32, 0x0a, 0x2e, 0x4d, 0x4f, 0x4e, 0x53, 0x54, 0x45, 0x52,
	0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x49, 0x52,
	0x45, 0x43, 0x54, 0x5f, 0x44, 0x49, 0x45, 0x5f, 0x53, 0x49, 0x4d, 0x55, 0x4c, 0x41, 0x54, 0x45,
	0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x2e, 0x0a, 0x2a, 0x4d, 0x4f, 0x4e,
	0x53, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x5f, 0x44, 0x49, 0x45, 0x5f, 0x53, 0x4b, 0x49, 0x50,
	0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x21, 0x0a, 0x1d, 0x4d, 0x4f, 0x4e,
	0x53, 0x54, 0x45, 0x52, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x4e, 0x4f, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x10, 0x04, 0x42, 0x2e, 0x5a, 0x0e,
	0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MonsterBattleType_proto_rawDescOnce sync.Once
	file_MonsterBattleType_proto_rawDescData = file_MonsterBattleType_proto_rawDesc
)

func file_MonsterBattleType_proto_rawDescGZIP() []byte {
	file_MonsterBattleType_proto_rawDescOnce.Do(func() {
		file_MonsterBattleType_proto_rawDescData = protoimpl.X.CompressGZIP(file_MonsterBattleType_proto_rawDescData)
	})
	return file_MonsterBattleType_proto_rawDescData
}

var file_MonsterBattleType_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MonsterBattleType_proto_goTypes = []interface{}{
	(MonsterBattleType)(0), // 0: MonsterBattleType
}
var file_MonsterBattleType_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MonsterBattleType_proto_init() }
func file_MonsterBattleType_proto_init() {
	if File_MonsterBattleType_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MonsterBattleType_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MonsterBattleType_proto_goTypes,
		DependencyIndexes: file_MonsterBattleType_proto_depIdxs,
		EnumInfos:         file_MonsterBattleType_proto_enumTypes,
	}.Build()
	File_MonsterBattleType_proto = out.File
	file_MonsterBattleType_proto_rawDesc = nil
	file_MonsterBattleType_proto_goTypes = nil
	file_MonsterBattleType_proto_depIdxs = nil
}
