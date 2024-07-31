// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MADMJLJELDI.proto

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

type MADMJLJELDI int32

const (
	MADMJLJELDI_MATCH3_STATE_IDLE     MADMJLJELDI = 0
	MADMJLJELDI_MATCH3_STATE_START    MADMJLJELDI = 1
	MADMJLJELDI_MATCH3_STATE_MATCH    MADMJLJELDI = 2
	MADMJLJELDI_MATCH3_STATE_GAME     MADMJLJELDI = 3
	MADMJLJELDI_MATCH3_STATE_HALFTIME MADMJLJELDI = 4
	MADMJLJELDI_MATCH3_STATE_OVER     MADMJLJELDI = 5
)

// Enum value maps for MADMJLJELDI.
var (
	MADMJLJELDI_name = map[int32]string{
		0: "MATCH3_STATE_IDLE",
		1: "MATCH3_STATE_START",
		2: "MATCH3_STATE_MATCH",
		3: "MATCH3_STATE_GAME",
		4: "MATCH3_STATE_HALFTIME",
		5: "MATCH3_STATE_OVER",
	}
	MADMJLJELDI_value = map[string]int32{
		"MATCH3_STATE_IDLE":     0,
		"MATCH3_STATE_START":    1,
		"MATCH3_STATE_MATCH":    2,
		"MATCH3_STATE_GAME":     3,
		"MATCH3_STATE_HALFTIME": 4,
		"MATCH3_STATE_OVER":     5,
	}
)

func (x MADMJLJELDI) Enum() *MADMJLJELDI {
	p := new(MADMJLJELDI)
	*p = x
	return p
}

func (x MADMJLJELDI) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MADMJLJELDI) Descriptor() protoreflect.EnumDescriptor {
	return file_MADMJLJELDI_proto_enumTypes[0].Descriptor()
}

func (MADMJLJELDI) Type() protoreflect.EnumType {
	return &file_MADMJLJELDI_proto_enumTypes[0]
}

func (x MADMJLJELDI) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MADMJLJELDI.Descriptor instead.
func (MADMJLJELDI) EnumDescriptor() ([]byte, []int) {
	return file_MADMJLJELDI_proto_rawDescGZIP(), []int{0}
}

var File_MADMJLJELDI_proto protoreflect.FileDescriptor

var file_MADMJLJELDI_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4d, 0x41, 0x44, 0x4d, 0x4a, 0x4c, 0x4a, 0x45, 0x4c, 0x44, 0x49, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0x9d, 0x01, 0x0a, 0x0b, 0x4d, 0x41, 0x44, 0x4d, 0x4a, 0x4c, 0x4a, 0x45,
	0x4c, 0x44, 0x49, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x33, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x49, 0x44, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x41,
	0x54, 0x43, 0x48, 0x33, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x33, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x4d, 0x41,
	0x54, 0x43, 0x48, 0x33, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x10,
	0x03, 0x12, 0x19, 0x0a, 0x15, 0x4d, 0x41, 0x54, 0x43, 0x48, 0x33, 0x5f, 0x53, 0x54, 0x41, 0x54,
	0x45, 0x5f, 0x48, 0x41, 0x4c, 0x46, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11,
	0x4d, 0x41, 0x54, 0x43, 0x48, 0x33, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x56, 0x45,
	0x52, 0x10, 0x05, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MADMJLJELDI_proto_rawDescOnce sync.Once
	file_MADMJLJELDI_proto_rawDescData = file_MADMJLJELDI_proto_rawDesc
)

func file_MADMJLJELDI_proto_rawDescGZIP() []byte {
	file_MADMJLJELDI_proto_rawDescOnce.Do(func() {
		file_MADMJLJELDI_proto_rawDescData = protoimpl.X.CompressGZIP(file_MADMJLJELDI_proto_rawDescData)
	})
	return file_MADMJLJELDI_proto_rawDescData
}

var file_MADMJLJELDI_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_MADMJLJELDI_proto_goTypes = []interface{}{
	(MADMJLJELDI)(0), // 0: MADMJLJELDI
}
var file_MADMJLJELDI_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MADMJLJELDI_proto_init() }
func file_MADMJLJELDI_proto_init() {
	if File_MADMJLJELDI_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MADMJLJELDI_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MADMJLJELDI_proto_goTypes,
		DependencyIndexes: file_MADMJLJELDI_proto_depIdxs,
		EnumInfos:         file_MADMJLJELDI_proto_enumTypes,
	}.Build()
	File_MADMJLJELDI_proto = out.File
	file_MADMJLJELDI_proto_rawDesc = nil
	file_MADMJLJELDI_proto_goTypes = nil
	file_MADMJLJELDI_proto_depIdxs = nil
}
