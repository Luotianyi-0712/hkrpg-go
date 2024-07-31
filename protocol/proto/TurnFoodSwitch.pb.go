// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TurnFoodSwitch.proto

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

type TurnFoodSwitch int32

const (
	TurnFoodSwitch_TURN_FOOD_SWITCH_NONE   TurnFoodSwitch = 0
	TurnFoodSwitch_attack                  TurnFoodSwitch = 1
	TurnFoodSwitch_TURN_FOOD_SWITCH_DEFINE TurnFoodSwitch = 2
)

// Enum value maps for TurnFoodSwitch.
var (
	TurnFoodSwitch_name = map[int32]string{
		0: "TURN_FOOD_SWITCH_NONE",
		1: "attack",
		2: "TURN_FOOD_SWITCH_DEFINE",
	}
	TurnFoodSwitch_value = map[string]int32{
		"TURN_FOOD_SWITCH_NONE":   0,
		"attack":                  1,
		"TURN_FOOD_SWITCH_DEFINE": 2,
	}
)

func (x TurnFoodSwitch) Enum() *TurnFoodSwitch {
	p := new(TurnFoodSwitch)
	*p = x
	return p
}

func (x TurnFoodSwitch) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TurnFoodSwitch) Descriptor() protoreflect.EnumDescriptor {
	return file_TurnFoodSwitch_proto_enumTypes[0].Descriptor()
}

func (TurnFoodSwitch) Type() protoreflect.EnumType {
	return &file_TurnFoodSwitch_proto_enumTypes[0]
}

func (x TurnFoodSwitch) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TurnFoodSwitch.Descriptor instead.
func (TurnFoodSwitch) EnumDescriptor() ([]byte, []int) {
	return file_TurnFoodSwitch_proto_rawDescGZIP(), []int{0}
}

var File_TurnFoodSwitch_proto protoreflect.FileDescriptor

var file_TurnFoodSwitch_proto_rawDesc = []byte{
	0x0a, 0x14, 0x54, 0x75, 0x72, 0x6e, 0x46, 0x6f, 0x6f, 0x64, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x54, 0x0a, 0x0e, 0x54, 0x75, 0x72, 0x6e, 0x46, 0x6f,
	0x6f, 0x64, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x55, 0x52, 0x4e,
	0x5f, 0x46, 0x4f, 0x4f, 0x44, 0x5f, 0x53, 0x57, 0x49, 0x54, 0x43, 0x48, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x10, 0x01, 0x12,
	0x1b, 0x0a, 0x17, 0x54, 0x55, 0x52, 0x4e, 0x5f, 0x46, 0x4f, 0x4f, 0x44, 0x5f, 0x53, 0x57, 0x49,
	0x54, 0x43, 0x48, 0x5f, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x10, 0x02, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TurnFoodSwitch_proto_rawDescOnce sync.Once
	file_TurnFoodSwitch_proto_rawDescData = file_TurnFoodSwitch_proto_rawDesc
)

func file_TurnFoodSwitch_proto_rawDescGZIP() []byte {
	file_TurnFoodSwitch_proto_rawDescOnce.Do(func() {
		file_TurnFoodSwitch_proto_rawDescData = protoimpl.X.CompressGZIP(file_TurnFoodSwitch_proto_rawDescData)
	})
	return file_TurnFoodSwitch_proto_rawDescData
}

var file_TurnFoodSwitch_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_TurnFoodSwitch_proto_goTypes = []interface{}{
	(TurnFoodSwitch)(0), // 0: TurnFoodSwitch
}
var file_TurnFoodSwitch_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TurnFoodSwitch_proto_init() }
func file_TurnFoodSwitch_proto_init() {
	if File_TurnFoodSwitch_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_TurnFoodSwitch_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TurnFoodSwitch_proto_goTypes,
		DependencyIndexes: file_TurnFoodSwitch_proto_depIdxs,
		EnumInfos:         file_TurnFoodSwitch_proto_enumTypes,
	}.Build()
	File_TurnFoodSwitch_proto = out.File
	file_TurnFoodSwitch_proto_rawDesc = nil
	file_TurnFoodSwitch_proto_goTypes = nil
	file_TurnFoodSwitch_proto_depIdxs = nil
}
