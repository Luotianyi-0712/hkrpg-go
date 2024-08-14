// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetMuseumInfoScRsp.proto

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

type GetMuseumInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MELNILMPOBA []*MNMMELAIBJK `protobuf:"bytes,4,rep,name=MELNILMPOBA,proto3" json:"MELNILMPOBA,omitempty"`
	MDHKFPDCFGP *KPNPOAAHDAC   `protobuf:"bytes,13,opt,name=MDHKFPDCFGP,proto3" json:"MDHKFPDCFGP,omitempty"`
	Retcode     uint32         `protobuf:"varint,2,opt,name=retcode,proto3" json:"retcode,omitempty"`
	CCKBBOEOKGO uint32         `protobuf:"varint,6,opt,name=CCKBBOEOKGO,proto3" json:"CCKBBOEOKGO,omitempty"`
	Exp         uint32         `protobuf:"varint,3,opt,name=exp,proto3" json:"exp,omitempty"`
	ADOBEODJBLJ *IPMIGBGLPNJ   `protobuf:"bytes,15,opt,name=ADOBEODJBLJ,proto3" json:"ADOBEODJBLJ,omitempty"`
	Level       uint32         `protobuf:"varint,11,opt,name=level,proto3" json:"level,omitempty"`
	EBNNBEEGJFN uint32         `protobuf:"varint,7,opt,name=EBNNBEEGJFN,proto3" json:"EBNNBEEGJFN,omitempty"`
	NBBKPIJEJPP uint32         `protobuf:"varint,9,opt,name=NBBKPIJEJPP,proto3" json:"NBBKPIJEJPP,omitempty"`
	PFGENNDADFA uint32         `protobuf:"varint,1,opt,name=PFGENNDADFA,proto3" json:"PFGENNDADFA,omitempty"`
	ODNOFEPBJAG []uint32       `protobuf:"varint,14,rep,packed,name=ODNOFEPBJAG,proto3" json:"ODNOFEPBJAG,omitempty"`
	GLGFCBDLOOI []*OPOHDHHDOAE `protobuf:"bytes,8,rep,name=GLGFCBDLOOI,proto3" json:"GLGFCBDLOOI,omitempty"`
	BHIHOPLDPGF []uint32       `protobuf:"varint,5,rep,packed,name=BHIHOPLDPGF,proto3" json:"BHIHOPLDPGF,omitempty"`
	JOANGDLCHKM uint32         `protobuf:"varint,12,opt,name=JOANGDLCHKM,proto3" json:"JOANGDLCHKM,omitempty"`
}

func (x *GetMuseumInfoScRsp) Reset() {
	*x = GetMuseumInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetMuseumInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMuseumInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMuseumInfoScRsp) ProtoMessage() {}

func (x *GetMuseumInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetMuseumInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMuseumInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetMuseumInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetMuseumInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetMuseumInfoScRsp) GetMELNILMPOBA() []*MNMMELAIBJK {
	if x != nil {
		return x.MELNILMPOBA
	}
	return nil
}

func (x *GetMuseumInfoScRsp) GetMDHKFPDCFGP() *KPNPOAAHDAC {
	if x != nil {
		return x.MDHKFPDCFGP
	}
	return nil
}

func (x *GetMuseumInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetMuseumInfoScRsp) GetCCKBBOEOKGO() uint32 {
	if x != nil {
		return x.CCKBBOEOKGO
	}
	return 0
}

func (x *GetMuseumInfoScRsp) GetExp() uint32 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *GetMuseumInfoScRsp) GetADOBEODJBLJ() *IPMIGBGLPNJ {
	if x != nil {
		return x.ADOBEODJBLJ
	}
	return nil
}

func (x *GetMuseumInfoScRsp) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *GetMuseumInfoScRsp) GetEBNNBEEGJFN() uint32 {
	if x != nil {
		return x.EBNNBEEGJFN
	}
	return 0
}

func (x *GetMuseumInfoScRsp) GetNBBKPIJEJPP() uint32 {
	if x != nil {
		return x.NBBKPIJEJPP
	}
	return 0
}

func (x *GetMuseumInfoScRsp) GetPFGENNDADFA() uint32 {
	if x != nil {
		return x.PFGENNDADFA
	}
	return 0
}

func (x *GetMuseumInfoScRsp) GetODNOFEPBJAG() []uint32 {
	if x != nil {
		return x.ODNOFEPBJAG
	}
	return nil
}

func (x *GetMuseumInfoScRsp) GetGLGFCBDLOOI() []*OPOHDHHDOAE {
	if x != nil {
		return x.GLGFCBDLOOI
	}
	return nil
}

func (x *GetMuseumInfoScRsp) GetBHIHOPLDPGF() []uint32 {
	if x != nil {
		return x.BHIHOPLDPGF
	}
	return nil
}

func (x *GetMuseumInfoScRsp) GetJOANGDLCHKM() uint32 {
	if x != nil {
		return x.JOANGDLCHKM
	}
	return 0
}

var File_GetMuseumInfoScRsp_proto protoreflect.FileDescriptor

var file_GetMuseumInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x18, 0x47, 0x65, 0x74, 0x4d, 0x75, 0x73, 0x65, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4b, 0x50, 0x4e, 0x50,
	0x4f, 0x41, 0x41, 0x48, 0x44, 0x41, 0x43, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49,
	0x50, 0x4d, 0x49, 0x47, 0x42, 0x47, 0x4c, 0x50, 0x4e, 0x4a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x4d, 0x4e, 0x4d, 0x4d, 0x45, 0x4c, 0x41, 0x49, 0x42, 0x4a, 0x4b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4f, 0x50, 0x4f, 0x48, 0x44, 0x48, 0x48, 0x44, 0x4f, 0x41, 0x45,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x04, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x75,
	0x73, 0x65, 0x75, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x2e, 0x0a,
	0x0b, 0x4d, 0x45, 0x4c, 0x4e, 0x49, 0x4c, 0x4d, 0x50, 0x4f, 0x42, 0x41, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x4e, 0x4d, 0x4d, 0x45, 0x4c, 0x41, 0x49, 0x42, 0x4a, 0x4b,
	0x52, 0x0b, 0x4d, 0x45, 0x4c, 0x4e, 0x49, 0x4c, 0x4d, 0x50, 0x4f, 0x42, 0x41, 0x12, 0x2e, 0x0a,
	0x0b, 0x4d, 0x44, 0x48, 0x4b, 0x46, 0x50, 0x44, 0x43, 0x46, 0x47, 0x50, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x50, 0x4e, 0x50, 0x4f, 0x41, 0x41, 0x48, 0x44, 0x41, 0x43,
	0x52, 0x0b, 0x4d, 0x44, 0x48, 0x4b, 0x46, 0x50, 0x44, 0x43, 0x46, 0x47, 0x50, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x43, 0x4b, 0x42, 0x42,
	0x4f, 0x45, 0x4f, 0x4b, 0x47, 0x4f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x43,
	0x4b, 0x42, 0x42, 0x4f, 0x45, 0x4f, 0x4b, 0x47, 0x4f, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x65, 0x78, 0x70, 0x12, 0x2e, 0x0a, 0x0b, 0x41,
	0x44, 0x4f, 0x42, 0x45, 0x4f, 0x44, 0x4a, 0x42, 0x4c, 0x4a, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x49, 0x50, 0x4d, 0x49, 0x47, 0x42, 0x47, 0x4c, 0x50, 0x4e, 0x4a, 0x52, 0x0b,
	0x41, 0x44, 0x4f, 0x42, 0x45, 0x4f, 0x44, 0x4a, 0x42, 0x4c, 0x4a, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x42, 0x4e, 0x4e, 0x42, 0x45, 0x45, 0x47, 0x4a, 0x46, 0x4e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x42, 0x4e, 0x4e, 0x42, 0x45, 0x45, 0x47,
	0x4a, 0x46, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x42, 0x42, 0x4b, 0x50, 0x49, 0x4a, 0x45, 0x4a,
	0x50, 0x50, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x42, 0x42, 0x4b, 0x50, 0x49,
	0x4a, 0x45, 0x4a, 0x50, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x46, 0x47, 0x45, 0x4e, 0x4e, 0x44,
	0x41, 0x44, 0x46, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x50, 0x46, 0x47, 0x45,
	0x4e, 0x4e, 0x44, 0x41, 0x44, 0x46, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x44, 0x4e, 0x4f, 0x46,
	0x45, 0x50, 0x42, 0x4a, 0x41, 0x47, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x44,
	0x4e, 0x4f, 0x46, 0x45, 0x50, 0x42, 0x4a, 0x41, 0x47, 0x12, 0x2e, 0x0a, 0x0b, 0x47, 0x4c, 0x47,
	0x46, 0x43, 0x42, 0x44, 0x4c, 0x4f, 0x4f, 0x49, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x4f, 0x50, 0x4f, 0x48, 0x44, 0x48, 0x48, 0x44, 0x4f, 0x41, 0x45, 0x52, 0x0b, 0x47, 0x4c,
	0x47, 0x46, 0x43, 0x42, 0x44, 0x4c, 0x4f, 0x4f, 0x49, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x48, 0x49,
	0x48, 0x4f, 0x50, 0x4c, 0x44, 0x50, 0x47, 0x46, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b,
	0x42, 0x48, 0x49, 0x48, 0x4f, 0x50, 0x4c, 0x44, 0x50, 0x47, 0x46, 0x12, 0x20, 0x0a, 0x0b, 0x4a,
	0x4f, 0x41, 0x4e, 0x47, 0x44, 0x4c, 0x43, 0x48, 0x4b, 0x4d, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0b, 0x4a, 0x4f, 0x41, 0x4e, 0x47, 0x44, 0x4c, 0x43, 0x48, 0x4b, 0x4d, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetMuseumInfoScRsp_proto_rawDescOnce sync.Once
	file_GetMuseumInfoScRsp_proto_rawDescData = file_GetMuseumInfoScRsp_proto_rawDesc
)

func file_GetMuseumInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetMuseumInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetMuseumInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetMuseumInfoScRsp_proto_rawDescData)
	})
	return file_GetMuseumInfoScRsp_proto_rawDescData
}

var file_GetMuseumInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetMuseumInfoScRsp_proto_goTypes = []interface{}{
	(*GetMuseumInfoScRsp)(nil), // 0: GetMuseumInfoScRsp
	(*MNMMELAIBJK)(nil),        // 1: MNMMELAIBJK
	(*KPNPOAAHDAC)(nil),        // 2: KPNPOAAHDAC
	(*IPMIGBGLPNJ)(nil),        // 3: IPMIGBGLPNJ
	(*OPOHDHHDOAE)(nil),        // 4: OPOHDHHDOAE
}
var file_GetMuseumInfoScRsp_proto_depIdxs = []int32{
	1, // 0: GetMuseumInfoScRsp.MELNILMPOBA:type_name -> MNMMELAIBJK
	2, // 1: GetMuseumInfoScRsp.MDHKFPDCFGP:type_name -> KPNPOAAHDAC
	3, // 2: GetMuseumInfoScRsp.ADOBEODJBLJ:type_name -> IPMIGBGLPNJ
	4, // 3: GetMuseumInfoScRsp.GLGFCBDLOOI:type_name -> OPOHDHHDOAE
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_GetMuseumInfoScRsp_proto_init() }
func file_GetMuseumInfoScRsp_proto_init() {
	if File_GetMuseumInfoScRsp_proto != nil {
		return
	}
	file_KPNPOAAHDAC_proto_init()
	file_IPMIGBGLPNJ_proto_init()
	file_MNMMELAIBJK_proto_init()
	file_OPOHDHHDOAE_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetMuseumInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMuseumInfoScRsp); i {
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
			RawDescriptor: file_GetMuseumInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetMuseumInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetMuseumInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetMuseumInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetMuseumInfoScRsp_proto = out.File
	file_GetMuseumInfoScRsp_proto_rawDesc = nil
	file_GetMuseumInfoScRsp_proto_goTypes = nil
	file_GetMuseumInfoScRsp_proto_depIdxs = nil
}
