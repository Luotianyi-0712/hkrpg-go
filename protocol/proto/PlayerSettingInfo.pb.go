// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PlayerSettingInfo.proto

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

type PlayerSettingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DAAAIHDPCFE bool `protobuf:"varint,2,opt,name=DAAAIHDPCFE,proto3" json:"DAAAIHDPCFE,omitempty"`
	LJDEMIHBKAJ bool `protobuf:"varint,1,opt,name=LJDEMIHBKAJ,proto3" json:"LJDEMIHBKAJ,omitempty"`
	MOKMEEDBECL bool `protobuf:"varint,3,opt,name=MOKMEEDBECL,proto3" json:"MOKMEEDBECL,omitempty"`
	LLNLDDKKGPH bool `protobuf:"varint,15,opt,name=LLNLDDKKGPH,proto3" json:"LLNLDDKKGPH,omitempty"`
	FLJHGDKONLC bool `protobuf:"varint,9,opt,name=FLJHGDKONLC,proto3" json:"FLJHGDKONLC,omitempty"`
	MAJIMDCHNDL bool `protobuf:"varint,11,opt,name=MAJIMDCHNDL,proto3" json:"MAJIMDCHNDL,omitempty"`
	BBJGEGEJJFB bool `protobuf:"varint,12,opt,name=BBJGEGEJJFB,proto3" json:"BBJGEGEJJFB,omitempty"`
	AOLKDAJONOH bool `protobuf:"varint,7,opt,name=AOLKDAJONOH,proto3" json:"AOLKDAJONOH,omitempty"`
	OJNELKIOAOK bool `protobuf:"varint,14,opt,name=OJNELKIOAOK,proto3" json:"OJNELKIOAOK,omitempty"`
}

func (x *PlayerSettingInfo) Reset() {
	*x = PlayerSettingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerSettingInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerSettingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerSettingInfo) ProtoMessage() {}

func (x *PlayerSettingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerSettingInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerSettingInfo.ProtoReflect.Descriptor instead.
func (*PlayerSettingInfo) Descriptor() ([]byte, []int) {
	return file_PlayerSettingInfo_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerSettingInfo) GetDAAAIHDPCFE() bool {
	if x != nil {
		return x.DAAAIHDPCFE
	}
	return false
}

func (x *PlayerSettingInfo) GetLJDEMIHBKAJ() bool {
	if x != nil {
		return x.LJDEMIHBKAJ
	}
	return false
}

func (x *PlayerSettingInfo) GetMOKMEEDBECL() bool {
	if x != nil {
		return x.MOKMEEDBECL
	}
	return false
}

func (x *PlayerSettingInfo) GetLLNLDDKKGPH() bool {
	if x != nil {
		return x.LLNLDDKKGPH
	}
	return false
}

func (x *PlayerSettingInfo) GetFLJHGDKONLC() bool {
	if x != nil {
		return x.FLJHGDKONLC
	}
	return false
}

func (x *PlayerSettingInfo) GetMAJIMDCHNDL() bool {
	if x != nil {
		return x.MAJIMDCHNDL
	}
	return false
}

func (x *PlayerSettingInfo) GetBBJGEGEJJFB() bool {
	if x != nil {
		return x.BBJGEGEJJFB
	}
	return false
}

func (x *PlayerSettingInfo) GetAOLKDAJONOH() bool {
	if x != nil {
		return x.AOLKDAJONOH
	}
	return false
}

func (x *PlayerSettingInfo) GetOJNELKIOAOK() bool {
	if x != nil {
		return x.OJNELKIOAOK
	}
	return false
}

var File_PlayerSettingInfo_proto protoreflect.FileDescriptor

var file_PlayerSettingInfo_proto_rawDesc = []byte{
	0x0a, 0x17, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x02, 0x0a, 0x11, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x20, 0x0a, 0x0b, 0x44, 0x41, 0x41, 0x41, 0x49, 0x48, 0x44, 0x50, 0x43, 0x46, 0x45, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x44, 0x41, 0x41, 0x41, 0x49, 0x48, 0x44, 0x50, 0x43, 0x46,
	0x45, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4a, 0x44, 0x45, 0x4d, 0x49, 0x48, 0x42, 0x4b, 0x41, 0x4a,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4c, 0x4a, 0x44, 0x45, 0x4d, 0x49, 0x48, 0x42,
	0x4b, 0x41, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x4f, 0x4b, 0x4d, 0x45, 0x45, 0x44, 0x42, 0x45,
	0x43, 0x4c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4d, 0x4f, 0x4b, 0x4d, 0x45, 0x45,
	0x44, 0x42, 0x45, 0x43, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4c, 0x4e, 0x4c, 0x44, 0x44, 0x4b,
	0x4b, 0x47, 0x50, 0x48, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4c, 0x4c, 0x4e, 0x4c,
	0x44, 0x44, 0x4b, 0x4b, 0x47, 0x50, 0x48, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x4c, 0x4a, 0x48, 0x47,
	0x44, 0x4b, 0x4f, 0x4e, 0x4c, 0x43, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x46, 0x4c,
	0x4a, 0x48, 0x47, 0x44, 0x4b, 0x4f, 0x4e, 0x4c, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x41, 0x4a,
	0x49, 0x4d, 0x44, 0x43, 0x48, 0x4e, 0x44, 0x4c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x4d, 0x41, 0x4a, 0x49, 0x4d, 0x44, 0x43, 0x48, 0x4e, 0x44, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x42,
	0x42, 0x4a, 0x47, 0x45, 0x47, 0x45, 0x4a, 0x4a, 0x46, 0x42, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x42, 0x42, 0x4a, 0x47, 0x45, 0x47, 0x45, 0x4a, 0x4a, 0x46, 0x42, 0x12, 0x20, 0x0a,
	0x0b, 0x41, 0x4f, 0x4c, 0x4b, 0x44, 0x41, 0x4a, 0x4f, 0x4e, 0x4f, 0x48, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x41, 0x4f, 0x4c, 0x4b, 0x44, 0x41, 0x4a, 0x4f, 0x4e, 0x4f, 0x48, 0x12,
	0x20, 0x0a, 0x0b, 0x4f, 0x4a, 0x4e, 0x45, 0x4c, 0x4b, 0x49, 0x4f, 0x41, 0x4f, 0x4b, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4f, 0x4a, 0x4e, 0x45, 0x4c, 0x4b, 0x49, 0x4f, 0x41, 0x4f,
	0x4b, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PlayerSettingInfo_proto_rawDescOnce sync.Once
	file_PlayerSettingInfo_proto_rawDescData = file_PlayerSettingInfo_proto_rawDesc
)

func file_PlayerSettingInfo_proto_rawDescGZIP() []byte {
	file_PlayerSettingInfo_proto_rawDescOnce.Do(func() {
		file_PlayerSettingInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerSettingInfo_proto_rawDescData)
	})
	return file_PlayerSettingInfo_proto_rawDescData
}

var file_PlayerSettingInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerSettingInfo_proto_goTypes = []interface{}{
	(*PlayerSettingInfo)(nil), // 0: PlayerSettingInfo
}
var file_PlayerSettingInfo_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PlayerSettingInfo_proto_init() }
func file_PlayerSettingInfo_proto_init() {
	if File_PlayerSettingInfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PlayerSettingInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerSettingInfo); i {
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
			RawDescriptor: file_PlayerSettingInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerSettingInfo_proto_goTypes,
		DependencyIndexes: file_PlayerSettingInfo_proto_depIdxs,
		MessageInfos:      file_PlayerSettingInfo_proto_msgTypes,
	}.Build()
	File_PlayerSettingInfo_proto = out.File
	file_PlayerSettingInfo_proto_rawDesc = nil
	file_PlayerSettingInfo_proto_goTypes = nil
	file_PlayerSettingInfo_proto_depIdxs = nil
}
