// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MatchThreeLevelEndCsReq.proto

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

type MatchThreeLevelEndCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DHFDCMJDOGL []uint32          `protobuf:"varint,2,rep,packed,name=DHFDCMJDOGL,proto3" json:"DHFDCMJDOGL,omitempty"`
	AJPBDPPHGHC string            `protobuf:"bytes,1,opt,name=AJPBDPPHGHC,proto3" json:"AJPBDPPHGHC,omitempty"`
	PCKBGKDCHAB map[uint32]uint32 `protobuf:"bytes,6,rep,name=PCKBGKDCHAB,proto3" json:"PCKBGKDCHAB,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	JIELNNCBKOD uint32            `protobuf:"varint,11,opt,name=JIELNNCBKOD,proto3" json:"JIELNNCBKOD,omitempty"`
	BGMAGGLPJNN uint32            `protobuf:"varint,8,opt,name=BGMAGGLPJNN,proto3" json:"BGMAGGLPJNN,omitempty"`
	CPKLPJGGEOM uint32            `protobuf:"varint,14,opt,name=CPKLPJGGEOM,proto3" json:"CPKLPJGGEOM,omitempty"`
	LAAFPDHGMBG uint32            `protobuf:"varint,4,opt,name=LAAFPDHGMBG,proto3" json:"LAAFPDHGMBG,omitempty"`
}

func (x *MatchThreeLevelEndCsReq) Reset() {
	*x = MatchThreeLevelEndCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MatchThreeLevelEndCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchThreeLevelEndCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchThreeLevelEndCsReq) ProtoMessage() {}

func (x *MatchThreeLevelEndCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_MatchThreeLevelEndCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchThreeLevelEndCsReq.ProtoReflect.Descriptor instead.
func (*MatchThreeLevelEndCsReq) Descriptor() ([]byte, []int) {
	return file_MatchThreeLevelEndCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *MatchThreeLevelEndCsReq) GetDHFDCMJDOGL() []uint32 {
	if x != nil {
		return x.DHFDCMJDOGL
	}
	return nil
}

func (x *MatchThreeLevelEndCsReq) GetAJPBDPPHGHC() string {
	if x != nil {
		return x.AJPBDPPHGHC
	}
	return ""
}

func (x *MatchThreeLevelEndCsReq) GetPCKBGKDCHAB() map[uint32]uint32 {
	if x != nil {
		return x.PCKBGKDCHAB
	}
	return nil
}

func (x *MatchThreeLevelEndCsReq) GetJIELNNCBKOD() uint32 {
	if x != nil {
		return x.JIELNNCBKOD
	}
	return 0
}

func (x *MatchThreeLevelEndCsReq) GetBGMAGGLPJNN() uint32 {
	if x != nil {
		return x.BGMAGGLPJNN
	}
	return 0
}

func (x *MatchThreeLevelEndCsReq) GetCPKLPJGGEOM() uint32 {
	if x != nil {
		return x.CPKLPJGGEOM
	}
	return 0
}

func (x *MatchThreeLevelEndCsReq) GetLAAFPDHGMBG() uint32 {
	if x != nil {
		return x.LAAFPDHGMBG
	}
	return 0
}

var File_MatchThreeLevelEndCsReq_proto protoreflect.FileDescriptor

var file_MatchThreeLevelEndCsReq_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x68, 0x72, 0x65, 0x65, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x45, 0x6e, 0x64, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xf2, 0x02, 0x0a, 0x17, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x68, 0x72, 0x65, 0x65, 0x4c, 0x65,
	0x76, 0x65, 0x6c, 0x45, 0x6e, 0x64, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x48, 0x46, 0x44, 0x43, 0x4d, 0x4a, 0x44, 0x4f, 0x47, 0x4c, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0b, 0x44, 0x48, 0x46, 0x44, 0x43, 0x4d, 0x4a, 0x44, 0x4f, 0x47, 0x4c, 0x12, 0x20, 0x0a,
	0x0b, 0x41, 0x4a, 0x50, 0x42, 0x44, 0x50, 0x50, 0x48, 0x47, 0x48, 0x43, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x41, 0x4a, 0x50, 0x42, 0x44, 0x50, 0x50, 0x48, 0x47, 0x48, 0x43, 0x12,
	0x4b, 0x0a, 0x0b, 0x50, 0x43, 0x4b, 0x42, 0x47, 0x4b, 0x44, 0x43, 0x48, 0x41, 0x42, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x54, 0x68, 0x72, 0x65,
	0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x45, 0x6e, 0x64, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x50,
	0x43, 0x4b, 0x42, 0x47, 0x4b, 0x44, 0x43, 0x48, 0x41, 0x42, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0b, 0x50, 0x43, 0x4b, 0x42, 0x47, 0x4b, 0x44, 0x43, 0x48, 0x41, 0x42, 0x12, 0x20, 0x0a, 0x0b,
	0x4a, 0x49, 0x45, 0x4c, 0x4e, 0x4e, 0x43, 0x42, 0x4b, 0x4f, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x4a, 0x49, 0x45, 0x4c, 0x4e, 0x4e, 0x43, 0x42, 0x4b, 0x4f, 0x44, 0x12, 0x20,
	0x0a, 0x0b, 0x42, 0x47, 0x4d, 0x41, 0x47, 0x47, 0x4c, 0x50, 0x4a, 0x4e, 0x4e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x42, 0x47, 0x4d, 0x41, 0x47, 0x47, 0x4c, 0x50, 0x4a, 0x4e, 0x4e,
	0x12, 0x20, 0x0a, 0x0b, 0x43, 0x50, 0x4b, 0x4c, 0x50, 0x4a, 0x47, 0x47, 0x45, 0x4f, 0x4d, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x50, 0x4b, 0x4c, 0x50, 0x4a, 0x47, 0x47, 0x45,
	0x4f, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x41, 0x41, 0x46, 0x50, 0x44, 0x48, 0x47, 0x4d, 0x42,
	0x47, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x41, 0x41, 0x46, 0x50, 0x44, 0x48,
	0x47, 0x4d, 0x42, 0x47, 0x1a, 0x3e, 0x0a, 0x10, 0x50, 0x43, 0x4b, 0x42, 0x47, 0x4b, 0x44, 0x43,
	0x48, 0x41, 0x42, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MatchThreeLevelEndCsReq_proto_rawDescOnce sync.Once
	file_MatchThreeLevelEndCsReq_proto_rawDescData = file_MatchThreeLevelEndCsReq_proto_rawDesc
)

func file_MatchThreeLevelEndCsReq_proto_rawDescGZIP() []byte {
	file_MatchThreeLevelEndCsReq_proto_rawDescOnce.Do(func() {
		file_MatchThreeLevelEndCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_MatchThreeLevelEndCsReq_proto_rawDescData)
	})
	return file_MatchThreeLevelEndCsReq_proto_rawDescData
}

var file_MatchThreeLevelEndCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_MatchThreeLevelEndCsReq_proto_goTypes = []interface{}{
	(*MatchThreeLevelEndCsReq)(nil), // 0: MatchThreeLevelEndCsReq
	nil,                             // 1: MatchThreeLevelEndCsReq.PCKBGKDCHABEntry
}
var file_MatchThreeLevelEndCsReq_proto_depIdxs = []int32{
	1, // 0: MatchThreeLevelEndCsReq.PCKBGKDCHAB:type_name -> MatchThreeLevelEndCsReq.PCKBGKDCHABEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_MatchThreeLevelEndCsReq_proto_init() }
func file_MatchThreeLevelEndCsReq_proto_init() {
	if File_MatchThreeLevelEndCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MatchThreeLevelEndCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchThreeLevelEndCsReq); i {
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
			RawDescriptor: file_MatchThreeLevelEndCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MatchThreeLevelEndCsReq_proto_goTypes,
		DependencyIndexes: file_MatchThreeLevelEndCsReq_proto_depIdxs,
		MessageInfos:      file_MatchThreeLevelEndCsReq_proto_msgTypes,
	}.Build()
	File_MatchThreeLevelEndCsReq_proto = out.File
	file_MatchThreeLevelEndCsReq_proto_rawDesc = nil
	file_MatchThreeLevelEndCsReq_proto_goTypes = nil
	file_MatchThreeLevelEndCsReq_proto_depIdxs = nil
}
