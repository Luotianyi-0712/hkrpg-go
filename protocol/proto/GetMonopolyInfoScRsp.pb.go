// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetMonopolyInfoScRsp.proto

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

type GetMonopolyInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JJCJFNALGAD   []uint32     `protobuf:"varint,15,rep,packed,name=JJCJFNALGAD,proto3" json:"JJCJFNALGAD,omitempty"`
	Retcode       uint32       `protobuf:"varint,2,opt,name=retcode,proto3" json:"retcode,omitempty"`
	LACNOCDMDEK   *NFKOEBLKDBA `protobuf:"bytes,10,opt,name=LACNOCDMDEK,proto3" json:"LACNOCDMDEK,omitempty"`
	PBCMIIENBFA   *NILNKCMEFKO `protobuf:"bytes,5,opt,name=PBCMIIENBFA,proto3" json:"PBCMIIENBFA,omitempty"`
	Stt           *IHBPFNOBGEM `protobuf:"bytes,4,opt,name=stt,proto3" json:"stt,omitempty"`
	DMNEPHPOLEN   *PANJFGJGLOL `protobuf:"bytes,12,opt,name=DMNEPHPOLEN,proto3" json:"DMNEPHPOLEN,omitempty"`
	RoomMap       *MKKFCGGHEPH `protobuf:"bytes,1,opt,name=room_map,json=roomMap,proto3" json:"room_map,omitempty"`
	RogueBuffInfo *DLGMNAAKMKD `protobuf:"bytes,14,opt,name=rogue_buff_info,json=rogueBuffInfo,proto3" json:"rogue_buff_info,omitempty"`
	NJHIOBCOLHM   *HPMNAAJDCGB `protobuf:"bytes,7,opt,name=NJHIOBCOLHM,proto3" json:"NJHIOBCOLHM,omitempty"`
	IIKIHCFKLJD   *BJEHMMDMEGL `protobuf:"bytes,13,opt,name=IIKIHCFKLJD,proto3" json:"IIKIHCFKLJD,omitempty"`
	PLDNMECDKIN   *JIONFPDOKEH `protobuf:"bytes,6,opt,name=PLDNMECDKIN,proto3" json:"PLDNMECDKIN,omitempty"`
	CJGJGEHNLFL   *HDJLPNBNFPB `protobuf:"bytes,11,opt,name=CJGJGEHNLFL,proto3" json:"CJGJGEHNLFL,omitempty"`
	PCABNBHKFHP   *DNGGOODOCFE `protobuf:"bytes,9,opt,name=PCABNBHKFHP,proto3" json:"PCABNBHKFHP,omitempty"`
}

func (x *GetMonopolyInfoScRsp) Reset() {
	*x = GetMonopolyInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetMonopolyInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMonopolyInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMonopolyInfoScRsp) ProtoMessage() {}

func (x *GetMonopolyInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetMonopolyInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMonopolyInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetMonopolyInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetMonopolyInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetMonopolyInfoScRsp) GetJJCJFNALGAD() []uint32 {
	if x != nil {
		return x.JJCJFNALGAD
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *GetMonopolyInfoScRsp) GetLACNOCDMDEK() *NFKOEBLKDBA {
	if x != nil {
		return x.LACNOCDMDEK
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetPBCMIIENBFA() *NILNKCMEFKO {
	if x != nil {
		return x.PBCMIIENBFA
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetStt() *IHBPFNOBGEM {
	if x != nil {
		return x.Stt
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetDMNEPHPOLEN() *PANJFGJGLOL {
	if x != nil {
		return x.DMNEPHPOLEN
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetRoomMap() *MKKFCGGHEPH {
	if x != nil {
		return x.RoomMap
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetRogueBuffInfo() *DLGMNAAKMKD {
	if x != nil {
		return x.RogueBuffInfo
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetNJHIOBCOLHM() *HPMNAAJDCGB {
	if x != nil {
		return x.NJHIOBCOLHM
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetIIKIHCFKLJD() *BJEHMMDMEGL {
	if x != nil {
		return x.IIKIHCFKLJD
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetPLDNMECDKIN() *JIONFPDOKEH {
	if x != nil {
		return x.PLDNMECDKIN
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetCJGJGEHNLFL() *HDJLPNBNFPB {
	if x != nil {
		return x.CJGJGEHNLFL
	}
	return nil
}

func (x *GetMonopolyInfoScRsp) GetPCABNBHKFHP() *DNGGOODOCFE {
	if x != nil {
		return x.PCABNBHKFHP
	}
	return nil
}

var File_GetMonopolyInfoScRsp_proto protoreflect.FileDescriptor

var file_GetMonopolyInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x50,
	0x4d, 0x4e, 0x41, 0x41, 0x4a, 0x44, 0x43, 0x47, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x11, 0x42, 0x4a, 0x45, 0x48, 0x4d, 0x4d, 0x44, 0x4d, 0x45, 0x47, 0x4c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x11, 0x44, 0x4c, 0x47, 0x4d, 0x4e, 0x41, 0x41, 0x4b, 0x4d, 0x4b, 0x44, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x50, 0x41, 0x4e, 0x4a, 0x46, 0x47, 0x4a, 0x47, 0x4c,
	0x4f, 0x4c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4d, 0x4b, 0x4b, 0x46, 0x43, 0x47,
	0x47, 0x48, 0x45, 0x50, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x49, 0x4f,
	0x4e, 0x46, 0x50, 0x44, 0x4f, 0x4b, 0x45, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x44, 0x4e, 0x47, 0x47, 0x4f, 0x4f, 0x44, 0x4f, 0x43, 0x46, 0x45, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x4e, 0x49, 0x4c, 0x4e, 0x4b, 0x43, 0x4d, 0x45, 0x46, 0x4b, 0x4f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e, 0x46, 0x4b, 0x4f, 0x45, 0x42, 0x4c, 0x4b, 0x44, 0x42,
	0x41, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x49, 0x48, 0x42, 0x50, 0x46, 0x4e, 0x4f,
	0x42, 0x47, 0x45, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x48, 0x44, 0x4a, 0x4c,
	0x50, 0x4e, 0x42, 0x4e, 0x46, 0x50, 0x42, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x04,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x6e, 0x6f, 0x70, 0x6f, 0x6c, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x4a, 0x43, 0x4a, 0x46, 0x4e,
	0x41, 0x4c, 0x47, 0x41, 0x44, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x4a, 0x43,
	0x4a, 0x46, 0x4e, 0x41, 0x4c, 0x47, 0x41, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x4c, 0x41, 0x43, 0x4e, 0x4f, 0x43, 0x44, 0x4d, 0x44, 0x45,
	0x4b, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x46, 0x4b, 0x4f, 0x45, 0x42,
	0x4c, 0x4b, 0x44, 0x42, 0x41, 0x52, 0x0b, 0x4c, 0x41, 0x43, 0x4e, 0x4f, 0x43, 0x44, 0x4d, 0x44,
	0x45, 0x4b, 0x12, 0x2e, 0x0a, 0x0b, 0x50, 0x42, 0x43, 0x4d, 0x49, 0x49, 0x45, 0x4e, 0x42, 0x46,
	0x41, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4e, 0x49, 0x4c, 0x4e, 0x4b, 0x43,
	0x4d, 0x45, 0x46, 0x4b, 0x4f, 0x52, 0x0b, 0x50, 0x42, 0x43, 0x4d, 0x49, 0x49, 0x45, 0x4e, 0x42,
	0x46, 0x41, 0x12, 0x1e, 0x0a, 0x03, 0x73, 0x74, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x49, 0x48, 0x42, 0x50, 0x46, 0x4e, 0x4f, 0x42, 0x47, 0x45, 0x4d, 0x52, 0x03, 0x73,
	0x74, 0x74, 0x12, 0x2e, 0x0a, 0x0b, 0x44, 0x4d, 0x4e, 0x45, 0x50, 0x48, 0x50, 0x4f, 0x4c, 0x45,
	0x4e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x50, 0x41, 0x4e, 0x4a, 0x46, 0x47,
	0x4a, 0x47, 0x4c, 0x4f, 0x4c, 0x52, 0x0b, 0x44, 0x4d, 0x4e, 0x45, 0x50, 0x48, 0x50, 0x4f, 0x4c,
	0x45, 0x4e, 0x12, 0x27, 0x0a, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x6d, 0x61, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4d, 0x4b, 0x4b, 0x46, 0x43, 0x47, 0x47, 0x48, 0x45,
	0x50, 0x48, 0x52, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x4d, 0x61, 0x70, 0x12, 0x34, 0x0a, 0x0f, 0x72,
	0x6f, 0x67, 0x75, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x4c, 0x47, 0x4d, 0x4e, 0x41, 0x41, 0x4b, 0x4d,
	0x4b, 0x44, 0x52, 0x0d, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x2e, 0x0a, 0x0b, 0x4e, 0x4a, 0x48, 0x49, 0x4f, 0x42, 0x43, 0x4f, 0x4c, 0x48, 0x4d,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x50, 0x4d, 0x4e, 0x41, 0x41, 0x4a,
	0x44, 0x43, 0x47, 0x42, 0x52, 0x0b, 0x4e, 0x4a, 0x48, 0x49, 0x4f, 0x42, 0x43, 0x4f, 0x4c, 0x48,
	0x4d, 0x12, 0x2e, 0x0a, 0x0b, 0x49, 0x49, 0x4b, 0x49, 0x48, 0x43, 0x46, 0x4b, 0x4c, 0x4a, 0x44,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x4a, 0x45, 0x48, 0x4d, 0x4d, 0x44,
	0x4d, 0x45, 0x47, 0x4c, 0x52, 0x0b, 0x49, 0x49, 0x4b, 0x49, 0x48, 0x43, 0x46, 0x4b, 0x4c, 0x4a,
	0x44, 0x12, 0x2e, 0x0a, 0x0b, 0x50, 0x4c, 0x44, 0x4e, 0x4d, 0x45, 0x43, 0x44, 0x4b, 0x49, 0x4e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4a, 0x49, 0x4f, 0x4e, 0x46, 0x50, 0x44,
	0x4f, 0x4b, 0x45, 0x48, 0x52, 0x0b, 0x50, 0x4c, 0x44, 0x4e, 0x4d, 0x45, 0x43, 0x44, 0x4b, 0x49,
	0x4e, 0x12, 0x2e, 0x0a, 0x0b, 0x43, 0x4a, 0x47, 0x4a, 0x47, 0x45, 0x48, 0x4e, 0x4c, 0x46, 0x4c,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x48, 0x44, 0x4a, 0x4c, 0x50, 0x4e, 0x42,
	0x4e, 0x46, 0x50, 0x42, 0x52, 0x0b, 0x43, 0x4a, 0x47, 0x4a, 0x47, 0x45, 0x48, 0x4e, 0x4c, 0x46,
	0x4c, 0x12, 0x2e, 0x0a, 0x0b, 0x50, 0x43, 0x41, 0x42, 0x4e, 0x42, 0x48, 0x4b, 0x46, 0x48, 0x50,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x44, 0x4e, 0x47, 0x47, 0x4f, 0x4f, 0x44,
	0x4f, 0x43, 0x46, 0x45, 0x52, 0x0b, 0x50, 0x43, 0x41, 0x42, 0x4e, 0x42, 0x48, 0x4b, 0x46, 0x48,
	0x50, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_GetMonopolyInfoScRsp_proto_rawDescOnce sync.Once
	file_GetMonopolyInfoScRsp_proto_rawDescData = file_GetMonopolyInfoScRsp_proto_rawDesc
)

func file_GetMonopolyInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetMonopolyInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetMonopolyInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetMonopolyInfoScRsp_proto_rawDescData)
	})
	return file_GetMonopolyInfoScRsp_proto_rawDescData
}

var file_GetMonopolyInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetMonopolyInfoScRsp_proto_goTypes = []interface{}{
	(*GetMonopolyInfoScRsp)(nil), // 0: GetMonopolyInfoScRsp
	(*NFKOEBLKDBA)(nil),          // 1: NFKOEBLKDBA
	(*NILNKCMEFKO)(nil),          // 2: NILNKCMEFKO
	(*IHBPFNOBGEM)(nil),          // 3: IHBPFNOBGEM
	(*PANJFGJGLOL)(nil),          // 4: PANJFGJGLOL
	(*MKKFCGGHEPH)(nil),          // 5: MKKFCGGHEPH
	(*DLGMNAAKMKD)(nil),          // 6: DLGMNAAKMKD
	(*HPMNAAJDCGB)(nil),          // 7: HPMNAAJDCGB
	(*BJEHMMDMEGL)(nil),          // 8: BJEHMMDMEGL
	(*JIONFPDOKEH)(nil),          // 9: JIONFPDOKEH
	(*HDJLPNBNFPB)(nil),          // 10: HDJLPNBNFPB
	(*DNGGOODOCFE)(nil),          // 11: DNGGOODOCFE
}
var file_GetMonopolyInfoScRsp_proto_depIdxs = []int32{
	1,  // 0: GetMonopolyInfoScRsp.LACNOCDMDEK:type_name -> NFKOEBLKDBA
	2,  // 1: GetMonopolyInfoScRsp.PBCMIIENBFA:type_name -> NILNKCMEFKO
	3,  // 2: GetMonopolyInfoScRsp.stt:type_name -> IHBPFNOBGEM
	4,  // 3: GetMonopolyInfoScRsp.DMNEPHPOLEN:type_name -> PANJFGJGLOL
	5,  // 4: GetMonopolyInfoScRsp.room_map:type_name -> MKKFCGGHEPH
	6,  // 5: GetMonopolyInfoScRsp.rogue_buff_info:type_name -> DLGMNAAKMKD
	7,  // 6: GetMonopolyInfoScRsp.NJHIOBCOLHM:type_name -> HPMNAAJDCGB
	8,  // 7: GetMonopolyInfoScRsp.IIKIHCFKLJD:type_name -> BJEHMMDMEGL
	9,  // 8: GetMonopolyInfoScRsp.PLDNMECDKIN:type_name -> JIONFPDOKEH
	10, // 9: GetMonopolyInfoScRsp.CJGJGEHNLFL:type_name -> HDJLPNBNFPB
	11, // 10: GetMonopolyInfoScRsp.PCABNBHKFHP:type_name -> DNGGOODOCFE
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_GetMonopolyInfoScRsp_proto_init() }
func file_GetMonopolyInfoScRsp_proto_init() {
	if File_GetMonopolyInfoScRsp_proto != nil {
		return
	}
	file_HPMNAAJDCGB_proto_init()
	file_BJEHMMDMEGL_proto_init()
	file_DLGMNAAKMKD_proto_init()
	file_PANJFGJGLOL_proto_init()
	file_MKKFCGGHEPH_proto_init()
	file_JIONFPDOKEH_proto_init()
	file_DNGGOODOCFE_proto_init()
	file_NILNKCMEFKO_proto_init()
	file_NFKOEBLKDBA_proto_init()
	file_IHBPFNOBGEM_proto_init()
	file_HDJLPNBNFPB_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetMonopolyInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMonopolyInfoScRsp); i {
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
			RawDescriptor: file_GetMonopolyInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetMonopolyInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetMonopolyInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetMonopolyInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetMonopolyInfoScRsp_proto = out.File
	file_GetMonopolyInfoScRsp_proto_rawDesc = nil
	file_GetMonopolyInfoScRsp_proto_goTypes = nil
	file_GetMonopolyInfoScRsp_proto_depIdxs = nil
}
