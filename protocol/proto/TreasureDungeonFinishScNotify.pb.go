// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TreasureDungeonFinishScNotify.proto

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

type TreasureDungeonFinishScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LOJIIDHFDGP uint32            `protobuf:"varint,4,opt,name=LOJIIDHFDGP,proto3" json:"LOJIIDHFDGP,omitempty"`
	IHPBBEGGKAI uint32            `protobuf:"varint,5,opt,name=IHPBBEGGKAI,proto3" json:"IHPBBEGGKAI,omitempty"`
	FCCLDGIMEGA uint32            `protobuf:"varint,2,opt,name=FCCLDGIMEGA,proto3" json:"FCCLDGIMEGA,omitempty"`
	BIIJLDFLMID uint32            `protobuf:"varint,7,opt,name=BIIJLDFLMID,proto3" json:"BIIJLDFLMID,omitempty"`
	IsWin       bool              `protobuf:"varint,14,opt,name=is_win,json=isWin,proto3" json:"is_win,omitempty"`
	OLCPBIGEPBK map[uint32]uint32 `protobuf:"bytes,15,rep,name=OLCPBIGEPBK,proto3" json:"OLCPBIGEPBK,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	CNCLNMDPDME map[uint32]uint32 `protobuf:"bytes,11,rep,name=CNCLNMDPDME,proto3" json:"CNCLNMDPDME,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *TreasureDungeonFinishScNotify) Reset() {
	*x = TreasureDungeonFinishScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TreasureDungeonFinishScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TreasureDungeonFinishScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TreasureDungeonFinishScNotify) ProtoMessage() {}

func (x *TreasureDungeonFinishScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_TreasureDungeonFinishScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TreasureDungeonFinishScNotify.ProtoReflect.Descriptor instead.
func (*TreasureDungeonFinishScNotify) Descriptor() ([]byte, []int) {
	return file_TreasureDungeonFinishScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *TreasureDungeonFinishScNotify) GetLOJIIDHFDGP() uint32 {
	if x != nil {
		return x.LOJIIDHFDGP
	}
	return 0
}

func (x *TreasureDungeonFinishScNotify) GetIHPBBEGGKAI() uint32 {
	if x != nil {
		return x.IHPBBEGGKAI
	}
	return 0
}

func (x *TreasureDungeonFinishScNotify) GetFCCLDGIMEGA() uint32 {
	if x != nil {
		return x.FCCLDGIMEGA
	}
	return 0
}

func (x *TreasureDungeonFinishScNotify) GetBIIJLDFLMID() uint32 {
	if x != nil {
		return x.BIIJLDFLMID
	}
	return 0
}

func (x *TreasureDungeonFinishScNotify) GetIsWin() bool {
	if x != nil {
		return x.IsWin
	}
	return false
}

func (x *TreasureDungeonFinishScNotify) GetOLCPBIGEPBK() map[uint32]uint32 {
	if x != nil {
		return x.OLCPBIGEPBK
	}
	return nil
}

func (x *TreasureDungeonFinishScNotify) GetCNCLNMDPDME() map[uint32]uint32 {
	if x != nil {
		return x.CNCLNMDPDME
	}
	return nil
}

var File_TreasureDungeonFinishScNotify_proto protoreflect.FileDescriptor

var file_TreasureDungeonFinishScNotify_proto_rawDesc = []byte{
	0x0a, 0x23, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f,
	0x6e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x03, 0x0a, 0x1d, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75,
	0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x53,
	0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x4f, 0x4a, 0x49, 0x49,
	0x44, 0x48, 0x46, 0x44, 0x47, 0x50, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x4f,
	0x4a, 0x49, 0x49, 0x44, 0x48, 0x46, 0x44, 0x47, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x48, 0x50,
	0x42, 0x42, 0x45, 0x47, 0x47, 0x4b, 0x41, 0x49, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x49, 0x48, 0x50, 0x42, 0x42, 0x45, 0x47, 0x47, 0x4b, 0x41, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x46,
	0x43, 0x43, 0x4c, 0x44, 0x47, 0x49, 0x4d, 0x45, 0x47, 0x41, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x46, 0x43, 0x43, 0x4c, 0x44, 0x47, 0x49, 0x4d, 0x45, 0x47, 0x41, 0x12, 0x20, 0x0a,
	0x0b, 0x42, 0x49, 0x49, 0x4a, 0x4c, 0x44, 0x46, 0x4c, 0x4d, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x42, 0x49, 0x49, 0x4a, 0x4c, 0x44, 0x46, 0x4c, 0x4d, 0x49, 0x44, 0x12,
	0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x77, 0x69, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x69, 0x73, 0x57, 0x69, 0x6e, 0x12, 0x51, 0x0a, 0x0b, 0x4f, 0x4c, 0x43, 0x50, 0x42, 0x49,
	0x47, 0x45, 0x50, 0x42, 0x4b, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x54, 0x72,
	0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x4f, 0x4c, 0x43, 0x50,
	0x42, 0x49, 0x47, 0x45, 0x50, 0x42, 0x4b, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x4f, 0x4c,
	0x43, 0x50, 0x42, 0x49, 0x47, 0x45, 0x50, 0x42, 0x4b, 0x12, 0x51, 0x0a, 0x0b, 0x43, 0x4e, 0x43,
	0x4c, 0x4e, 0x4d, 0x44, 0x50, 0x44, 0x4d, 0x45, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f,
	0x2e, 0x54, 0x72, 0x65, 0x61, 0x73, 0x75, 0x72, 0x65, 0x44, 0x75, 0x6e, 0x67, 0x65, 0x6f, 0x6e,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x43,
	0x4e, 0x43, 0x4c, 0x4e, 0x4d, 0x44, 0x50, 0x44, 0x4d, 0x45, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0b, 0x43, 0x4e, 0x43, 0x4c, 0x4e, 0x4d, 0x44, 0x50, 0x44, 0x4d, 0x45, 0x1a, 0x3e, 0x0a, 0x10,
	0x4f, 0x4c, 0x43, 0x50, 0x42, 0x49, 0x47, 0x45, 0x50, 0x42, 0x4b, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10,
	0x43, 0x4e, 0x43, 0x4c, 0x4e, 0x4d, 0x44, 0x50, 0x44, 0x4d, 0x45, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x2e, 0x5a, 0x0e,
	0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TreasureDungeonFinishScNotify_proto_rawDescOnce sync.Once
	file_TreasureDungeonFinishScNotify_proto_rawDescData = file_TreasureDungeonFinishScNotify_proto_rawDesc
)

func file_TreasureDungeonFinishScNotify_proto_rawDescGZIP() []byte {
	file_TreasureDungeonFinishScNotify_proto_rawDescOnce.Do(func() {
		file_TreasureDungeonFinishScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_TreasureDungeonFinishScNotify_proto_rawDescData)
	})
	return file_TreasureDungeonFinishScNotify_proto_rawDescData
}

var file_TreasureDungeonFinishScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_TreasureDungeonFinishScNotify_proto_goTypes = []interface{}{
	(*TreasureDungeonFinishScNotify)(nil), // 0: TreasureDungeonFinishScNotify
	nil,                                   // 1: TreasureDungeonFinishScNotify.OLCPBIGEPBKEntry
	nil,                                   // 2: TreasureDungeonFinishScNotify.CNCLNMDPDMEEntry
}
var file_TreasureDungeonFinishScNotify_proto_depIdxs = []int32{
	1, // 0: TreasureDungeonFinishScNotify.OLCPBIGEPBK:type_name -> TreasureDungeonFinishScNotify.OLCPBIGEPBKEntry
	2, // 1: TreasureDungeonFinishScNotify.CNCLNMDPDME:type_name -> TreasureDungeonFinishScNotify.CNCLNMDPDMEEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_TreasureDungeonFinishScNotify_proto_init() }
func file_TreasureDungeonFinishScNotify_proto_init() {
	if File_TreasureDungeonFinishScNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TreasureDungeonFinishScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TreasureDungeonFinishScNotify); i {
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
			RawDescriptor: file_TreasureDungeonFinishScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TreasureDungeonFinishScNotify_proto_goTypes,
		DependencyIndexes: file_TreasureDungeonFinishScNotify_proto_depIdxs,
		MessageInfos:      file_TreasureDungeonFinishScNotify_proto_msgTypes,
	}.Build()
	File_TreasureDungeonFinishScNotify_proto = out.File
	file_TreasureDungeonFinishScNotify_proto_rawDesc = nil
	file_TreasureDungeonFinishScNotify_proto_goTypes = nil
	file_TreasureDungeonFinishScNotify_proto_depIdxs = nil
}
