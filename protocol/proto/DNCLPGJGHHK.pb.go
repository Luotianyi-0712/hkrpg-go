// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: DNCLPGJGHHK.proto

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

type DNCLPGJGHHK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KEACGBKOFKF bool             `protobuf:"varint,13,opt,name=KEACGBKOFKF,proto3" json:"KEACGBKOFKF,omitempty"`
	DAFHJJEDMOF bool             `protobuf:"varint,8,opt,name=DAFHJJEDMOF,proto3" json:"DAFHJJEDMOF,omitempty"`
	KDAKDMCGFND bool             `protobuf:"varint,11,opt,name=KDAKDMCGFND,proto3" json:"KDAKDMCGFND,omitempty"`
	BPELFJGIJID bool             `protobuf:"varint,7,opt,name=BPELFJGIJID,proto3" json:"BPELFJGIJID,omitempty"`
	DKLJGCEHPJL bool             `protobuf:"varint,5,opt,name=DKLJGCEHPJL,proto3" json:"DKLJGCEHPJL,omitempty"`
	EPMCKMCDIGB BattleRecordType `protobuf:"varint,15,opt,name=EPMCKMCDIGB,proto3,enum=BattleRecordType" json:"EPMCKMCDIGB,omitempty"`
}

func (x *DNCLPGJGHHK) Reset() {
	*x = DNCLPGJGHHK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DNCLPGJGHHK_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DNCLPGJGHHK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DNCLPGJGHHK) ProtoMessage() {}

func (x *DNCLPGJGHHK) ProtoReflect() protoreflect.Message {
	mi := &file_DNCLPGJGHHK_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DNCLPGJGHHK.ProtoReflect.Descriptor instead.
func (*DNCLPGJGHHK) Descriptor() ([]byte, []int) {
	return file_DNCLPGJGHHK_proto_rawDescGZIP(), []int{0}
}

func (x *DNCLPGJGHHK) GetKEACGBKOFKF() bool {
	if x != nil {
		return x.KEACGBKOFKF
	}
	return false
}

func (x *DNCLPGJGHHK) GetDAFHJJEDMOF() bool {
	if x != nil {
		return x.DAFHJJEDMOF
	}
	return false
}

func (x *DNCLPGJGHHK) GetKDAKDMCGFND() bool {
	if x != nil {
		return x.KDAKDMCGFND
	}
	return false
}

func (x *DNCLPGJGHHK) GetBPELFJGIJID() bool {
	if x != nil {
		return x.BPELFJGIJID
	}
	return false
}

func (x *DNCLPGJGHHK) GetDKLJGCEHPJL() bool {
	if x != nil {
		return x.DKLJGCEHPJL
	}
	return false
}

func (x *DNCLPGJGHHK) GetEPMCKMCDIGB() BattleRecordType {
	if x != nil {
		return x.EPMCKMCDIGB
	}
	return BattleRecordType_BATTLE_RECORD_NONE
}

var File_DNCLPGJGHHK_proto protoreflect.FileDescriptor

var file_DNCLPGJGHHK_proto_rawDesc = []byte{
	0x0a, 0x11, 0x44, 0x4e, 0x43, 0x4c, 0x50, 0x47, 0x4a, 0x47, 0x48, 0x48, 0x4b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a, 0x0b,
	0x44, 0x4e, 0x43, 0x4c, 0x50, 0x47, 0x4a, 0x47, 0x48, 0x48, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x4b,
	0x45, 0x41, 0x43, 0x47, 0x42, 0x4b, 0x4f, 0x46, 0x4b, 0x46, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x4b, 0x45, 0x41, 0x43, 0x47, 0x42, 0x4b, 0x4f, 0x46, 0x4b, 0x46, 0x12, 0x20, 0x0a,
	0x0b, 0x44, 0x41, 0x46, 0x48, 0x4a, 0x4a, 0x45, 0x44, 0x4d, 0x4f, 0x46, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x44, 0x41, 0x46, 0x48, 0x4a, 0x4a, 0x45, 0x44, 0x4d, 0x4f, 0x46, 0x12,
	0x20, 0x0a, 0x0b, 0x4b, 0x44, 0x41, 0x4b, 0x44, 0x4d, 0x43, 0x47, 0x46, 0x4e, 0x44, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4b, 0x44, 0x41, 0x4b, 0x44, 0x4d, 0x43, 0x47, 0x46, 0x4e,
	0x44, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x50, 0x45, 0x4c, 0x46, 0x4a, 0x47, 0x49, 0x4a, 0x49, 0x44,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x42, 0x50, 0x45, 0x4c, 0x46, 0x4a, 0x47, 0x49,
	0x4a, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x4b, 0x4c, 0x4a, 0x47, 0x43, 0x45, 0x48, 0x50,
	0x4a, 0x4c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x44, 0x4b, 0x4c, 0x4a, 0x47, 0x43,
	0x45, 0x48, 0x50, 0x4a, 0x4c, 0x12, 0x33, 0x0a, 0x0b, 0x45, 0x50, 0x4d, 0x43, 0x4b, 0x4d, 0x43,
	0x44, 0x49, 0x47, 0x42, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x45,
	0x50, 0x4d, 0x43, 0x4b, 0x4d, 0x43, 0x44, 0x49, 0x47, 0x42, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DNCLPGJGHHK_proto_rawDescOnce sync.Once
	file_DNCLPGJGHHK_proto_rawDescData = file_DNCLPGJGHHK_proto_rawDesc
)

func file_DNCLPGJGHHK_proto_rawDescGZIP() []byte {
	file_DNCLPGJGHHK_proto_rawDescOnce.Do(func() {
		file_DNCLPGJGHHK_proto_rawDescData = protoimpl.X.CompressGZIP(file_DNCLPGJGHHK_proto_rawDescData)
	})
	return file_DNCLPGJGHHK_proto_rawDescData
}

var file_DNCLPGJGHHK_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DNCLPGJGHHK_proto_goTypes = []interface{}{
	(*DNCLPGJGHHK)(nil),   // 0: DNCLPGJGHHK
	(BattleRecordType)(0), // 1: BattleRecordType
}
var file_DNCLPGJGHHK_proto_depIdxs = []int32{
	1, // 0: DNCLPGJGHHK.EPMCKMCDIGB:type_name -> BattleRecordType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_DNCLPGJGHHK_proto_init() }
func file_DNCLPGJGHHK_proto_init() {
	if File_DNCLPGJGHHK_proto != nil {
		return
	}
	file_BattleRecordType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_DNCLPGJGHHK_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DNCLPGJGHHK); i {
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
			RawDescriptor: file_DNCLPGJGHHK_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DNCLPGJGHHK_proto_goTypes,
		DependencyIndexes: file_DNCLPGJGHHK_proto_depIdxs,
		MessageInfos:      file_DNCLPGJGHHK_proto_msgTypes,
	}.Build()
	File_DNCLPGJGHHK_proto = out.File
	file_DNCLPGJGHHK_proto_rawDesc = nil
	file_DNCLPGJGHHK_proto_goTypes = nil
	file_DNCLPGJGHHK_proto_depIdxs = nil
}
