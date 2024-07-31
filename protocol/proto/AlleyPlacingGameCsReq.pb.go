// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: AlleyPlacingGameCsReq.proto

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

type AlleyPlacingGameCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JMAIKKKNCLA uint32            `protobuf:"varint,4,opt,name=JMAIKKKNCLA,proto3" json:"JMAIKKKNCLA,omitempty"`
	IEPKPGBMBLA uint32            `protobuf:"varint,5,opt,name=IEPKPGBMBLA,proto3" json:"IEPKPGBMBLA,omitempty"`
	LBLLCLNHCND uint32            `protobuf:"varint,7,opt,name=LBLLCLNHCND,proto3" json:"LBLLCLNHCND,omitempty"`
	NFKNELJBBML uint32            `protobuf:"varint,15,opt,name=NFKNELJBBML,proto3" json:"NFKNELJBBML,omitempty"`
	AHCCGGOECLA uint32            `protobuf:"varint,12,opt,name=AHCCGGOECLA,proto3" json:"AHCCGGOECLA,omitempty"`
	DFCIJPEFOKP uint32            `protobuf:"varint,9,opt,name=DFCIJPEFOKP,proto3" json:"DFCIJPEFOKP,omitempty"`
	CostTime    uint32            `protobuf:"varint,8,opt,name=cost_time,json=costTime,proto3" json:"cost_time,omitempty"`
	JDHMPJKBHIL uint32            `protobuf:"varint,1,opt,name=JDHMPJKBHIL,proto3" json:"JDHMPJKBHIL,omitempty"`
	DFNHEIDEEHM *AlleyPlacingShip `protobuf:"bytes,13,opt,name=DFNHEIDEEHM,proto3" json:"DFNHEIDEEHM,omitempty"`
}

func (x *AlleyPlacingGameCsReq) Reset() {
	*x = AlleyPlacingGameCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AlleyPlacingGameCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlleyPlacingGameCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlleyPlacingGameCsReq) ProtoMessage() {}

func (x *AlleyPlacingGameCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_AlleyPlacingGameCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlleyPlacingGameCsReq.ProtoReflect.Descriptor instead.
func (*AlleyPlacingGameCsReq) Descriptor() ([]byte, []int) {
	return file_AlleyPlacingGameCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *AlleyPlacingGameCsReq) GetJMAIKKKNCLA() uint32 {
	if x != nil {
		return x.JMAIKKKNCLA
	}
	return 0
}

func (x *AlleyPlacingGameCsReq) GetIEPKPGBMBLA() uint32 {
	if x != nil {
		return x.IEPKPGBMBLA
	}
	return 0
}

func (x *AlleyPlacingGameCsReq) GetLBLLCLNHCND() uint32 {
	if x != nil {
		return x.LBLLCLNHCND
	}
	return 0
}

func (x *AlleyPlacingGameCsReq) GetNFKNELJBBML() uint32 {
	if x != nil {
		return x.NFKNELJBBML
	}
	return 0
}

func (x *AlleyPlacingGameCsReq) GetAHCCGGOECLA() uint32 {
	if x != nil {
		return x.AHCCGGOECLA
	}
	return 0
}

func (x *AlleyPlacingGameCsReq) GetDFCIJPEFOKP() uint32 {
	if x != nil {
		return x.DFCIJPEFOKP
	}
	return 0
}

func (x *AlleyPlacingGameCsReq) GetCostTime() uint32 {
	if x != nil {
		return x.CostTime
	}
	return 0
}

func (x *AlleyPlacingGameCsReq) GetJDHMPJKBHIL() uint32 {
	if x != nil {
		return x.JDHMPJKBHIL
	}
	return 0
}

func (x *AlleyPlacingGameCsReq) GetDFNHEIDEEHM() *AlleyPlacingShip {
	if x != nil {
		return x.DFNHEIDEEHM
	}
	return nil
}

var File_AlleyPlacingGameCsReq_proto protoreflect.FileDescriptor

var file_AlleyPlacingGameCsReq_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x41, 0x6c, 0x6c, 0x65, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x47, 0x61,
	0x6d, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x41,
	0x6c, 0x6c, 0x65, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x53, 0x68, 0x69, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x02, 0x0a, 0x15, 0x41, 0x6c, 0x6c, 0x65, 0x79, 0x50,
	0x6c, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12,
	0x20, 0x0a, 0x0b, 0x4a, 0x4d, 0x41, 0x49, 0x4b, 0x4b, 0x4b, 0x4e, 0x43, 0x4c, 0x41, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x4d, 0x41, 0x49, 0x4b, 0x4b, 0x4b, 0x4e, 0x43, 0x4c,
	0x41, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x45, 0x50, 0x4b, 0x50, 0x47, 0x42, 0x4d, 0x42, 0x4c, 0x41,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x45, 0x50, 0x4b, 0x50, 0x47, 0x42, 0x4d,
	0x42, 0x4c, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x42, 0x4c, 0x4c, 0x43, 0x4c, 0x4e, 0x48, 0x43,
	0x4e, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x42, 0x4c, 0x4c, 0x43, 0x4c,
	0x4e, 0x48, 0x43, 0x4e, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x46, 0x4b, 0x4e, 0x45, 0x4c, 0x4a,
	0x42, 0x42, 0x4d, 0x4c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x46, 0x4b, 0x4e,
	0x45, 0x4c, 0x4a, 0x42, 0x42, 0x4d, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x48, 0x43, 0x43, 0x47,
	0x47, 0x4f, 0x45, 0x43, 0x4c, 0x41, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x48,
	0x43, 0x43, 0x47, 0x47, 0x4f, 0x45, 0x43, 0x4c, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x46, 0x43,
	0x49, 0x4a, 0x50, 0x45, 0x46, 0x4f, 0x4b, 0x50, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x44, 0x46, 0x43, 0x49, 0x4a, 0x50, 0x45, 0x46, 0x4f, 0x4b, 0x50, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6f, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x63, 0x6f, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x44, 0x48, 0x4d,
	0x50, 0x4a, 0x4b, 0x42, 0x48, 0x49, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a,
	0x44, 0x48, 0x4d, 0x50, 0x4a, 0x4b, 0x42, 0x48, 0x49, 0x4c, 0x12, 0x33, 0x0a, 0x0b, 0x44, 0x46,
	0x4e, 0x48, 0x45, 0x49, 0x44, 0x45, 0x45, 0x48, 0x4d, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x41, 0x6c, 0x6c, 0x65, 0x79, 0x50, 0x6c, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x53, 0x68,
	0x69, 0x70, 0x52, 0x0b, 0x44, 0x46, 0x4e, 0x48, 0x45, 0x49, 0x44, 0x45, 0x45, 0x48, 0x4d, 0x42,
	0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_AlleyPlacingGameCsReq_proto_rawDescOnce sync.Once
	file_AlleyPlacingGameCsReq_proto_rawDescData = file_AlleyPlacingGameCsReq_proto_rawDesc
)

func file_AlleyPlacingGameCsReq_proto_rawDescGZIP() []byte {
	file_AlleyPlacingGameCsReq_proto_rawDescOnce.Do(func() {
		file_AlleyPlacingGameCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_AlleyPlacingGameCsReq_proto_rawDescData)
	})
	return file_AlleyPlacingGameCsReq_proto_rawDescData
}

var file_AlleyPlacingGameCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_AlleyPlacingGameCsReq_proto_goTypes = []interface{}{
	(*AlleyPlacingGameCsReq)(nil), // 0: AlleyPlacingGameCsReq
	(*AlleyPlacingShip)(nil),      // 1: AlleyPlacingShip
}
var file_AlleyPlacingGameCsReq_proto_depIdxs = []int32{
	1, // 0: AlleyPlacingGameCsReq.DFNHEIDEEHM:type_name -> AlleyPlacingShip
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_AlleyPlacingGameCsReq_proto_init() }
func file_AlleyPlacingGameCsReq_proto_init() {
	if File_AlleyPlacingGameCsReq_proto != nil {
		return
	}
	file_AlleyPlacingShip_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AlleyPlacingGameCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlleyPlacingGameCsReq); i {
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
			RawDescriptor: file_AlleyPlacingGameCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AlleyPlacingGameCsReq_proto_goTypes,
		DependencyIndexes: file_AlleyPlacingGameCsReq_proto_depIdxs,
		MessageInfos:      file_AlleyPlacingGameCsReq_proto_msgTypes,
	}.Build()
	File_AlleyPlacingGameCsReq_proto = out.File
	file_AlleyPlacingGameCsReq_proto_rawDesc = nil
	file_AlleyPlacingGameCsReq_proto_goTypes = nil
	file_AlleyPlacingGameCsReq_proto_depIdxs = nil
}
