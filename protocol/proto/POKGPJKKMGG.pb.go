// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: POKGPJKKMGG.proto

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

type POKGPJKKMGG struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CLMPJAOHKCD uint32         `protobuf:"varint,4,opt,name=CLMPJAOHKCD,proto3" json:"CLMPJAOHKCD,omitempty"`
	MEGDOBEGIPN uint64         `protobuf:"varint,11,opt,name=MEGDOBEGIPN,proto3" json:"MEGDOBEGIPN,omitempty"`
	LECMBIOIENL string         `protobuf:"bytes,257,opt,name=LECMBIOIENL,proto3" json:"LECMBIOIENL,omitempty"`
	PGFLHACOLBG []*NGPNIEAKNED `protobuf:"bytes,7,rep,name=PGFLHACOLBG,proto3" json:"PGFLHACOLBG,omitempty"`
	PMALMJDGCIB uint64         `protobuf:"varint,14,opt,name=PMALMJDGCIB,proto3" json:"PMALMJDGCIB,omitempty"`
	IENKLJIEAED []*NGPNIEAKNED `protobuf:"bytes,9,rep,name=IENKLJIEAED,proto3" json:"IENKLJIEAED,omitempty"`
	IDAEDOCPDID uint32         `protobuf:"varint,5,opt,name=IDAEDOCPDID,proto3" json:"IDAEDOCPDID,omitempty"`
	JKCCELHBDNN uint32         `protobuf:"varint,2,opt,name=JKCCELHBDNN,proto3" json:"JKCCELHBDNN,omitempty"`
	FANPKIMGLLG string         `protobuf:"bytes,242,opt,name=FANPKIMGLLG,proto3" json:"FANPKIMGLLG,omitempty"`
	KAJKPOHOPDO []*NGPNIEAKNED `protobuf:"bytes,1,rep,name=KAJKPOHOPDO,proto3" json:"KAJKPOHOPDO,omitempty"`
	HJFLFDPOGLB []*LDNACICHLEA `protobuf:"bytes,13,rep,name=HJFLFDPOGLB,proto3" json:"HJFLFDPOGLB,omitempty"`
	HINCOFBLIPO uint64         `protobuf:"varint,6,opt,name=HINCOFBLIPO,proto3" json:"HINCOFBLIPO,omitempty"`
	BBHJPICGJJE bool           `protobuf:"varint,10,opt,name=BBHJPICGJJE,proto3" json:"BBHJPICGJJE,omitempty"`
	PGLPGNCJEEC []*NGPNIEAKNED `protobuf:"bytes,12,rep,name=PGLPGNCJEEC,proto3" json:"PGLPGNCJEEC,omitempty"`
	MHMCKHGJJPA uint32         `protobuf:"varint,8,opt,name=MHMCKHGJJPA,proto3" json:"MHMCKHGJJPA,omitempty"`
	LECHMCAPHCF uint32         `protobuf:"varint,15,opt,name=LECHMCAPHCF,proto3" json:"LECHMCAPHCF,omitempty"`
	PoolId      uint32         `protobuf:"varint,3,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	HJLBKPMAHFA string         `protobuf:"bytes,850,opt,name=HJLBKPMAHFA,proto3" json:"HJLBKPMAHFA,omitempty"`
}

func (x *POKGPJKKMGG) Reset() {
	*x = POKGPJKKMGG{}
	if protoimpl.UnsafeEnabled {
		mi := &file_POKGPJKKMGG_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *POKGPJKKMGG) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*POKGPJKKMGG) ProtoMessage() {}

func (x *POKGPJKKMGG) ProtoReflect() protoreflect.Message {
	mi := &file_POKGPJKKMGG_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use POKGPJKKMGG.ProtoReflect.Descriptor instead.
func (*POKGPJKKMGG) Descriptor() ([]byte, []int) {
	return file_POKGPJKKMGG_proto_rawDescGZIP(), []int{0}
}

func (x *POKGPJKKMGG) GetCLMPJAOHKCD() uint32 {
	if x != nil {
		return x.CLMPJAOHKCD
	}
	return 0
}

func (x *POKGPJKKMGG) GetMEGDOBEGIPN() uint64 {
	if x != nil {
		return x.MEGDOBEGIPN
	}
	return 0
}

func (x *POKGPJKKMGG) GetLECMBIOIENL() string {
	if x != nil {
		return x.LECMBIOIENL
	}
	return ""
}

func (x *POKGPJKKMGG) GetPGFLHACOLBG() []*NGPNIEAKNED {
	if x != nil {
		return x.PGFLHACOLBG
	}
	return nil
}

func (x *POKGPJKKMGG) GetPMALMJDGCIB() uint64 {
	if x != nil {
		return x.PMALMJDGCIB
	}
	return 0
}

func (x *POKGPJKKMGG) GetIENKLJIEAED() []*NGPNIEAKNED {
	if x != nil {
		return x.IENKLJIEAED
	}
	return nil
}

func (x *POKGPJKKMGG) GetIDAEDOCPDID() uint32 {
	if x != nil {
		return x.IDAEDOCPDID
	}
	return 0
}

func (x *POKGPJKKMGG) GetJKCCELHBDNN() uint32 {
	if x != nil {
		return x.JKCCELHBDNN
	}
	return 0
}

func (x *POKGPJKKMGG) GetFANPKIMGLLG() string {
	if x != nil {
		return x.FANPKIMGLLG
	}
	return ""
}

func (x *POKGPJKKMGG) GetKAJKPOHOPDO() []*NGPNIEAKNED {
	if x != nil {
		return x.KAJKPOHOPDO
	}
	return nil
}

func (x *POKGPJKKMGG) GetHJFLFDPOGLB() []*LDNACICHLEA {
	if x != nil {
		return x.HJFLFDPOGLB
	}
	return nil
}

func (x *POKGPJKKMGG) GetHINCOFBLIPO() uint64 {
	if x != nil {
		return x.HINCOFBLIPO
	}
	return 0
}

func (x *POKGPJKKMGG) GetBBHJPICGJJE() bool {
	if x != nil {
		return x.BBHJPICGJJE
	}
	return false
}

func (x *POKGPJKKMGG) GetPGLPGNCJEEC() []*NGPNIEAKNED {
	if x != nil {
		return x.PGLPGNCJEEC
	}
	return nil
}

func (x *POKGPJKKMGG) GetMHMCKHGJJPA() uint32 {
	if x != nil {
		return x.MHMCKHGJJPA
	}
	return 0
}

func (x *POKGPJKKMGG) GetLECHMCAPHCF() uint32 {
	if x != nil {
		return x.LECHMCAPHCF
	}
	return 0
}

func (x *POKGPJKKMGG) GetPoolId() uint32 {
	if x != nil {
		return x.PoolId
	}
	return 0
}

func (x *POKGPJKKMGG) GetHJLBKPMAHFA() string {
	if x != nil {
		return x.HJLBKPMAHFA
	}
	return ""
}

var File_POKGPJKKMGG_proto protoreflect.FileDescriptor

var file_POKGPJKKMGG_proto_rawDesc = []byte{
	0x0a, 0x11, 0x50, 0x4f, 0x4b, 0x47, 0x50, 0x4a, 0x4b, 0x4b, 0x4d, 0x47, 0x47, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x47, 0x50, 0x4e, 0x49, 0x45, 0x41, 0x4b, 0x4e, 0x45, 0x44,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x44, 0x4e, 0x41, 0x43, 0x49, 0x43, 0x48,
	0x4c, 0x45, 0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x05, 0x0a, 0x0b, 0x50, 0x4f,
	0x4b, 0x47, 0x50, 0x4a, 0x4b, 0x4b, 0x4d, 0x47, 0x47, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x4c, 0x4d,
	0x50, 0x4a, 0x41, 0x4f, 0x48, 0x4b, 0x43, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x43, 0x4c, 0x4d, 0x50, 0x4a, 0x41, 0x4f, 0x48, 0x4b, 0x43, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4d,
	0x45, 0x47, 0x44, 0x4f, 0x42, 0x45, 0x47, 0x49, 0x50, 0x4e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x4d, 0x45, 0x47, 0x44, 0x4f, 0x42, 0x45, 0x47, 0x49, 0x50, 0x4e, 0x12, 0x21, 0x0a,
	0x0b, 0x4c, 0x45, 0x43, 0x4d, 0x42, 0x49, 0x4f, 0x49, 0x45, 0x4e, 0x4c, 0x18, 0x81, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x4c, 0x45, 0x43, 0x4d, 0x42, 0x49, 0x4f, 0x49, 0x45, 0x4e, 0x4c,
	0x12, 0x2e, 0x0a, 0x0b, 0x50, 0x47, 0x46, 0x4c, 0x48, 0x41, 0x43, 0x4f, 0x4c, 0x42, 0x47, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x47, 0x50, 0x4e, 0x49, 0x45, 0x41, 0x4b,
	0x4e, 0x45, 0x44, 0x52, 0x0b, 0x50, 0x47, 0x46, 0x4c, 0x48, 0x41, 0x43, 0x4f, 0x4c, 0x42, 0x47,
	0x12, 0x20, 0x0a, 0x0b, 0x50, 0x4d, 0x41, 0x4c, 0x4d, 0x4a, 0x44, 0x47, 0x43, 0x49, 0x42, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x50, 0x4d, 0x41, 0x4c, 0x4d, 0x4a, 0x44, 0x47, 0x43,
	0x49, 0x42, 0x12, 0x2e, 0x0a, 0x0b, 0x49, 0x45, 0x4e, 0x4b, 0x4c, 0x4a, 0x49, 0x45, 0x41, 0x45,
	0x44, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x47, 0x50, 0x4e, 0x49, 0x45,
	0x41, 0x4b, 0x4e, 0x45, 0x44, 0x52, 0x0b, 0x49, 0x45, 0x4e, 0x4b, 0x4c, 0x4a, 0x49, 0x45, 0x41,
	0x45, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x44, 0x41, 0x45, 0x44, 0x4f, 0x43, 0x50, 0x44, 0x49,
	0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x44, 0x41, 0x45, 0x44, 0x4f, 0x43,
	0x50, 0x44, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x4b, 0x43, 0x43, 0x45, 0x4c, 0x48, 0x42,
	0x44, 0x4e, 0x4e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x4b, 0x43, 0x43, 0x45,
	0x4c, 0x48, 0x42, 0x44, 0x4e, 0x4e, 0x12, 0x21, 0x0a, 0x0b, 0x46, 0x41, 0x4e, 0x50, 0x4b, 0x49,
	0x4d, 0x47, 0x4c, 0x4c, 0x47, 0x18, 0xf2, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x46, 0x41,
	0x4e, 0x50, 0x4b, 0x49, 0x4d, 0x47, 0x4c, 0x4c, 0x47, 0x12, 0x2e, 0x0a, 0x0b, 0x4b, 0x41, 0x4a,
	0x4b, 0x50, 0x4f, 0x48, 0x4f, 0x50, 0x44, 0x4f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x4e, 0x47, 0x50, 0x4e, 0x49, 0x45, 0x41, 0x4b, 0x4e, 0x45, 0x44, 0x52, 0x0b, 0x4b, 0x41,
	0x4a, 0x4b, 0x50, 0x4f, 0x48, 0x4f, 0x50, 0x44, 0x4f, 0x12, 0x2e, 0x0a, 0x0b, 0x48, 0x4a, 0x46,
	0x4c, 0x46, 0x44, 0x50, 0x4f, 0x47, 0x4c, 0x42, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x4c, 0x44, 0x4e, 0x41, 0x43, 0x49, 0x43, 0x48, 0x4c, 0x45, 0x41, 0x52, 0x0b, 0x48, 0x4a,
	0x46, 0x4c, 0x46, 0x44, 0x50, 0x4f, 0x47, 0x4c, 0x42, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x49, 0x4e,
	0x43, 0x4f, 0x46, 0x42, 0x4c, 0x49, 0x50, 0x4f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b,
	0x48, 0x49, 0x4e, 0x43, 0x4f, 0x46, 0x42, 0x4c, 0x49, 0x50, 0x4f, 0x12, 0x20, 0x0a, 0x0b, 0x42,
	0x42, 0x48, 0x4a, 0x50, 0x49, 0x43, 0x47, 0x4a, 0x4a, 0x45, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x42, 0x42, 0x48, 0x4a, 0x50, 0x49, 0x43, 0x47, 0x4a, 0x4a, 0x45, 0x12, 0x2e, 0x0a,
	0x0b, 0x50, 0x47, 0x4c, 0x50, 0x47, 0x4e, 0x43, 0x4a, 0x45, 0x45, 0x43, 0x18, 0x0c, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x47, 0x50, 0x4e, 0x49, 0x45, 0x41, 0x4b, 0x4e, 0x45, 0x44,
	0x52, 0x0b, 0x50, 0x47, 0x4c, 0x50, 0x47, 0x4e, 0x43, 0x4a, 0x45, 0x45, 0x43, 0x12, 0x20, 0x0a,
	0x0b, 0x4d, 0x48, 0x4d, 0x43, 0x4b, 0x48, 0x47, 0x4a, 0x4a, 0x50, 0x41, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x4d, 0x48, 0x4d, 0x43, 0x4b, 0x48, 0x47, 0x4a, 0x4a, 0x50, 0x41, 0x12,
	0x20, 0x0a, 0x0b, 0x4c, 0x45, 0x43, 0x48, 0x4d, 0x43, 0x41, 0x50, 0x48, 0x43, 0x46, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x45, 0x43, 0x48, 0x4d, 0x43, 0x41, 0x50, 0x48, 0x43,
	0x46, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x06, 0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x48, 0x4a,
	0x4c, 0x42, 0x4b, 0x50, 0x4d, 0x41, 0x48, 0x46, 0x41, 0x18, 0xd2, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x48, 0x4a, 0x4c, 0x42, 0x4b, 0x50, 0x4d, 0x41, 0x48, 0x46, 0x41, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_POKGPJKKMGG_proto_rawDescOnce sync.Once
	file_POKGPJKKMGG_proto_rawDescData = file_POKGPJKKMGG_proto_rawDesc
)

func file_POKGPJKKMGG_proto_rawDescGZIP() []byte {
	file_POKGPJKKMGG_proto_rawDescOnce.Do(func() {
		file_POKGPJKKMGG_proto_rawDescData = protoimpl.X.CompressGZIP(file_POKGPJKKMGG_proto_rawDescData)
	})
	return file_POKGPJKKMGG_proto_rawDescData
}

var file_POKGPJKKMGG_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_POKGPJKKMGG_proto_goTypes = []interface{}{
	(*POKGPJKKMGG)(nil), // 0: POKGPJKKMGG
	(*NGPNIEAKNED)(nil), // 1: NGPNIEAKNED
	(*LDNACICHLEA)(nil), // 2: LDNACICHLEA
}
var file_POKGPJKKMGG_proto_depIdxs = []int32{
	1, // 0: POKGPJKKMGG.PGFLHACOLBG:type_name -> NGPNIEAKNED
	1, // 1: POKGPJKKMGG.IENKLJIEAED:type_name -> NGPNIEAKNED
	1, // 2: POKGPJKKMGG.KAJKPOHOPDO:type_name -> NGPNIEAKNED
	2, // 3: POKGPJKKMGG.HJFLFDPOGLB:type_name -> LDNACICHLEA
	1, // 4: POKGPJKKMGG.PGLPGNCJEEC:type_name -> NGPNIEAKNED
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_POKGPJKKMGG_proto_init() }
func file_POKGPJKKMGG_proto_init() {
	if File_POKGPJKKMGG_proto != nil {
		return
	}
	file_NGPNIEAKNED_proto_init()
	file_LDNACICHLEA_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_POKGPJKKMGG_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*POKGPJKKMGG); i {
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
			RawDescriptor: file_POKGPJKKMGG_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_POKGPJKKMGG_proto_goTypes,
		DependencyIndexes: file_POKGPJKKMGG_proto_depIdxs,
		MessageInfos:      file_POKGPJKKMGG_proto_msgTypes,
	}.Build()
	File_POKGPJKKMGG_proto = out.File
	file_POKGPJKKMGG_proto_rawDesc = nil
	file_POKGPJKKMGG_proto_goTypes = nil
	file_POKGPJKKMGG_proto_depIdxs = nil
}
