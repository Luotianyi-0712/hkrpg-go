// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PlayerLoginCsReq.proto

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

type PlayerLoginCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ILEHKCJPGGE   string       `protobuf:"bytes,8,opt,name=ILEHKCJPGGE,proto3" json:"ILEHKCJPGGE,omitempty"`
	PLINCKJBBMN   string       `protobuf:"bytes,10,opt,name=PLINCKJBBMN,proto3" json:"PLINCKJBBMN,omitempty"`
	LGLEFCNIJKL   string       `protobuf:"bytes,13,opt,name=LGLEFCNIJKL,proto3" json:"LGLEFCNIJKL,omitempty"`
	LoginRandom   uint64       `protobuf:"varint,9,opt,name=login_random,json=loginRandom,proto3" json:"login_random,omitempty"`
	ClientVersion string       `protobuf:"bytes,4,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	LanguageType  LanguageType `protobuf:"varint,7,opt,name=language_type,json=languageType,proto3,enum=LanguageType" json:"language_type,omitempty"`
	DHICFKJEPBI   uint32       `protobuf:"varint,1687,opt,name=DHICFKJEPBI,proto3" json:"DHICFKJEPBI,omitempty"`
	Signature     string       `protobuf:"bytes,12,opt,name=signature,proto3" json:"signature,omitempty"`
	JJNINCMBIDN   string       `protobuf:"bytes,5,opt,name=JJNINCMBIDN,proto3" json:"JJNINCMBIDN,omitempty"`
	OPIACEKOANJ   string       `protobuf:"bytes,188,opt,name=OPIACEKOANJ,proto3" json:"OPIACEKOANJ,omitempty"`
	NJLPCOEBMFJ   string       `protobuf:"bytes,2,opt,name=NJLPCOEBMFJ,proto3" json:"NJLPCOEBMFJ,omitempty"`
	LGOMKMJMEDM   string       `protobuf:"bytes,155,opt,name=LGOMKMJMEDM,proto3" json:"LGOMKMJMEDM,omitempty"`
	BCIJKADEMIE   string       `protobuf:"bytes,940,opt,name=BCIJKADEMIE,proto3" json:"BCIJKADEMIE,omitempty"`
	GIMCNHMAPBP   string       `protobuf:"bytes,11,opt,name=GIMCNHMAPBP,proto3" json:"GIMCNHMAPBP,omitempty"`
	Platform      PlatformType `protobuf:"varint,3,opt,name=platform,proto3,enum=PlatformType" json:"platform,omitempty"`
	CIDHAHDCHDL   uint32       `protobuf:"varint,15,opt,name=CIDHAHDCHDL,proto3" json:"CIDHAHDCHDL,omitempty"`
	ResVersion    uint32       `protobuf:"varint,1,opt,name=res_version,json=resVersion,proto3" json:"res_version,omitempty"`
	MCKKBKPOMLI   string       `protobuf:"bytes,138,opt,name=MCKKBKPOMLI,proto3" json:"MCKKBKPOMLI,omitempty"`
	GANJMFCLNFL   *GBAMNAPGHFF `protobuf:"bytes,885,opt,name=GANJMFCLNFL,proto3" json:"GANJMFCLNFL,omitempty"`
	RogueGetInfo  string       `protobuf:"bytes,6,opt,name=rogue_get_info,json=rogueGetInfo,proto3" json:"rogue_get_info,omitempty"`
	OKPDCDEDOIL   uint32       `protobuf:"varint,1630,opt,name=OKPDCDEDOIL,proto3" json:"OKPDCDEDOIL,omitempty"`
	BOLHLHDEGGG   string       `protobuf:"bytes,2038,opt,name=BOLHLHDEGGG,proto3" json:"BOLHLHDEGGG,omitempty"`
	HKIGMBKIDLA   string       `protobuf:"bytes,14,opt,name=HKIGMBKIDLA,proto3" json:"HKIGMBKIDLA,omitempty"`
	BHHOMHKAMLO   bool         `protobuf:"varint,801,opt,name=BHHOMHKAMLO,proto3" json:"BHHOMHKAMLO,omitempty"`
}

func (x *PlayerLoginCsReq) Reset() {
	*x = PlayerLoginCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PlayerLoginCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerLoginCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerLoginCsReq) ProtoMessage() {}

func (x *PlayerLoginCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_PlayerLoginCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerLoginCsReq.ProtoReflect.Descriptor instead.
func (*PlayerLoginCsReq) Descriptor() ([]byte, []int) {
	return file_PlayerLoginCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerLoginCsReq) GetILEHKCJPGGE() string {
	if x != nil {
		return x.ILEHKCJPGGE
	}
	return ""
}

func (x *PlayerLoginCsReq) GetPLINCKJBBMN() string {
	if x != nil {
		return x.PLINCKJBBMN
	}
	return ""
}

func (x *PlayerLoginCsReq) GetLGLEFCNIJKL() string {
	if x != nil {
		return x.LGLEFCNIJKL
	}
	return ""
}

func (x *PlayerLoginCsReq) GetLoginRandom() uint64 {
	if x != nil {
		return x.LoginRandom
	}
	return 0
}

func (x *PlayerLoginCsReq) GetClientVersion() string {
	if x != nil {
		return x.ClientVersion
	}
	return ""
}

func (x *PlayerLoginCsReq) GetLanguageType() LanguageType {
	if x != nil {
		return x.LanguageType
	}
	return LanguageType_LANGUAGE_NONE
}

func (x *PlayerLoginCsReq) GetDHICFKJEPBI() uint32 {
	if x != nil {
		return x.DHICFKJEPBI
	}
	return 0
}

func (x *PlayerLoginCsReq) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *PlayerLoginCsReq) GetJJNINCMBIDN() string {
	if x != nil {
		return x.JJNINCMBIDN
	}
	return ""
}

func (x *PlayerLoginCsReq) GetOPIACEKOANJ() string {
	if x != nil {
		return x.OPIACEKOANJ
	}
	return ""
}

func (x *PlayerLoginCsReq) GetNJLPCOEBMFJ() string {
	if x != nil {
		return x.NJLPCOEBMFJ
	}
	return ""
}

func (x *PlayerLoginCsReq) GetLGOMKMJMEDM() string {
	if x != nil {
		return x.LGOMKMJMEDM
	}
	return ""
}

func (x *PlayerLoginCsReq) GetBCIJKADEMIE() string {
	if x != nil {
		return x.BCIJKADEMIE
	}
	return ""
}

func (x *PlayerLoginCsReq) GetGIMCNHMAPBP() string {
	if x != nil {
		return x.GIMCNHMAPBP
	}
	return ""
}

func (x *PlayerLoginCsReq) GetPlatform() PlatformType {
	if x != nil {
		return x.Platform
	}
	return PlatformType_EDITOR
}

func (x *PlayerLoginCsReq) GetCIDHAHDCHDL() uint32 {
	if x != nil {
		return x.CIDHAHDCHDL
	}
	return 0
}

func (x *PlayerLoginCsReq) GetResVersion() uint32 {
	if x != nil {
		return x.ResVersion
	}
	return 0
}

func (x *PlayerLoginCsReq) GetMCKKBKPOMLI() string {
	if x != nil {
		return x.MCKKBKPOMLI
	}
	return ""
}

func (x *PlayerLoginCsReq) GetGANJMFCLNFL() *GBAMNAPGHFF {
	if x != nil {
		return x.GANJMFCLNFL
	}
	return nil
}

func (x *PlayerLoginCsReq) GetRogueGetInfo() string {
	if x != nil {
		return x.RogueGetInfo
	}
	return ""
}

func (x *PlayerLoginCsReq) GetOKPDCDEDOIL() uint32 {
	if x != nil {
		return x.OKPDCDEDOIL
	}
	return 0
}

func (x *PlayerLoginCsReq) GetBOLHLHDEGGG() string {
	if x != nil {
		return x.BOLHLHDEGGG
	}
	return ""
}

func (x *PlayerLoginCsReq) GetHKIGMBKIDLA() string {
	if x != nil {
		return x.HKIGMBKIDLA
	}
	return ""
}

func (x *PlayerLoginCsReq) GetBHHOMHKAMLO() bool {
	if x != nil {
		return x.BHHOMHKAMLO
	}
	return false
}

var File_PlayerLoginCsReq_proto protoreflect.FileDescriptor

var file_PlayerLoginCsReq_proto_rawDesc = []byte{
	0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x73, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x42,
	0x41, 0x4d, 0x4e, 0x41, 0x50, 0x47, 0x48, 0x46, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xf9, 0x06, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x4c, 0x45, 0x48,
	0x4b, 0x43, 0x4a, 0x50, 0x47, 0x47, 0x45, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x49,
	0x4c, 0x45, 0x48, 0x4b, 0x43, 0x4a, 0x50, 0x47, 0x47, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x4c,
	0x49, 0x4e, 0x43, 0x4b, 0x4a, 0x42, 0x42, 0x4d, 0x4e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x50, 0x4c, 0x49, 0x4e, 0x43, 0x4b, 0x4a, 0x42, 0x42, 0x4d, 0x4e, 0x12, 0x20, 0x0a, 0x0b,
	0x4c, 0x47, 0x4c, 0x45, 0x46, 0x43, 0x4e, 0x49, 0x4a, 0x4b, 0x4c, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x4c, 0x47, 0x4c, 0x45, 0x46, 0x43, 0x4e, 0x49, 0x4a, 0x4b, 0x4c, 0x12, 0x21,
	0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x0d, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x0d, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0b,
	0x44, 0x48, 0x49, 0x43, 0x46, 0x4b, 0x4a, 0x45, 0x50, 0x42, 0x49, 0x18, 0x97, 0x0d, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x44, 0x48, 0x49, 0x43, 0x46, 0x4b, 0x4a, 0x45, 0x50, 0x42, 0x49, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x4a, 0x4a, 0x4e, 0x49, 0x4e, 0x43, 0x4d, 0x42, 0x49, 0x44, 0x4e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x4a, 0x4a, 0x4e, 0x49, 0x4e, 0x43, 0x4d, 0x42, 0x49, 0x44, 0x4e, 0x12,
	0x21, 0x0a, 0x0b, 0x4f, 0x50, 0x49, 0x41, 0x43, 0x45, 0x4b, 0x4f, 0x41, 0x4e, 0x4a, 0x18, 0xbc,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x50, 0x49, 0x41, 0x43, 0x45, 0x4b, 0x4f, 0x41,
	0x4e, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x4a, 0x4c, 0x50, 0x43, 0x4f, 0x45, 0x42, 0x4d, 0x46,
	0x4a, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4e, 0x4a, 0x4c, 0x50, 0x43, 0x4f, 0x45,
	0x42, 0x4d, 0x46, 0x4a, 0x12, 0x21, 0x0a, 0x0b, 0x4c, 0x47, 0x4f, 0x4d, 0x4b, 0x4d, 0x4a, 0x4d,
	0x45, 0x44, 0x4d, 0x18, 0x9b, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4c, 0x47, 0x4f, 0x4d,
	0x4b, 0x4d, 0x4a, 0x4d, 0x45, 0x44, 0x4d, 0x12, 0x21, 0x0a, 0x0b, 0x42, 0x43, 0x49, 0x4a, 0x4b,
	0x41, 0x44, 0x45, 0x4d, 0x49, 0x45, 0x18, 0xac, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x42,
	0x43, 0x49, 0x4a, 0x4b, 0x41, 0x44, 0x45, 0x4d, 0x49, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x49,
	0x4d, 0x43, 0x4e, 0x48, 0x4d, 0x41, 0x50, 0x42, 0x50, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x47, 0x49, 0x4d, 0x43, 0x4e, 0x48, 0x4d, 0x41, 0x50, 0x42, 0x50, 0x12, 0x29, 0x0a, 0x08,
	0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d,
	0x2e, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x49, 0x44, 0x48, 0x41,
	0x48, 0x44, 0x43, 0x48, 0x44, 0x4c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x49,
	0x44, 0x48, 0x41, 0x48, 0x44, 0x43, 0x48, 0x44, 0x4c, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0b, 0x4d, 0x43,
	0x4b, 0x4b, 0x42, 0x4b, 0x50, 0x4f, 0x4d, 0x4c, 0x49, 0x18, 0x8a, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x4d, 0x43, 0x4b, 0x4b, 0x42, 0x4b, 0x50, 0x4f, 0x4d, 0x4c, 0x49, 0x12, 0x2f, 0x0a,
	0x0b, 0x47, 0x41, 0x4e, 0x4a, 0x4d, 0x46, 0x43, 0x4c, 0x4e, 0x46, 0x4c, 0x18, 0xf5, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x42, 0x41, 0x4d, 0x4e, 0x41, 0x50, 0x47, 0x48, 0x46,
	0x46, 0x52, 0x0b, 0x47, 0x41, 0x4e, 0x4a, 0x4d, 0x46, 0x43, 0x4c, 0x4e, 0x46, 0x4c, 0x12, 0x24,
	0x0a, 0x0e, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0b, 0x4f, 0x4b, 0x50, 0x44, 0x43, 0x44, 0x45, 0x44,
	0x4f, 0x49, 0x4c, 0x18, 0xde, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x4b, 0x50, 0x44,
	0x43, 0x44, 0x45, 0x44, 0x4f, 0x49, 0x4c, 0x12, 0x21, 0x0a, 0x0b, 0x42, 0x4f, 0x4c, 0x48, 0x4c,
	0x48, 0x44, 0x45, 0x47, 0x47, 0x47, 0x18, 0xf6, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x42,
	0x4f, 0x4c, 0x48, 0x4c, 0x48, 0x44, 0x45, 0x47, 0x47, 0x47, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4b,
	0x49, 0x47, 0x4d, 0x42, 0x4b, 0x49, 0x44, 0x4c, 0x41, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x48, 0x4b, 0x49, 0x47, 0x4d, 0x42, 0x4b, 0x49, 0x44, 0x4c, 0x41, 0x12, 0x21, 0x0a, 0x0b,
	0x42, 0x48, 0x48, 0x4f, 0x4d, 0x48, 0x4b, 0x41, 0x4d, 0x4c, 0x4f, 0x18, 0xa1, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0b, 0x42, 0x48, 0x48, 0x4f, 0x4d, 0x48, 0x4b, 0x41, 0x4d, 0x4c, 0x4f, 0x42,
	0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67,
	0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_PlayerLoginCsReq_proto_rawDescOnce sync.Once
	file_PlayerLoginCsReq_proto_rawDescData = file_PlayerLoginCsReq_proto_rawDesc
)

func file_PlayerLoginCsReq_proto_rawDescGZIP() []byte {
	file_PlayerLoginCsReq_proto_rawDescOnce.Do(func() {
		file_PlayerLoginCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_PlayerLoginCsReq_proto_rawDescData)
	})
	return file_PlayerLoginCsReq_proto_rawDescData
}

var file_PlayerLoginCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PlayerLoginCsReq_proto_goTypes = []interface{}{
	(*PlayerLoginCsReq)(nil), // 0: PlayerLoginCsReq
	(LanguageType)(0),        // 1: LanguageType
	(PlatformType)(0),        // 2: PlatformType
	(*GBAMNAPGHFF)(nil),      // 3: GBAMNAPGHFF
}
var file_PlayerLoginCsReq_proto_depIdxs = []int32{
	1, // 0: PlayerLoginCsReq.language_type:type_name -> LanguageType
	2, // 1: PlayerLoginCsReq.platform:type_name -> PlatformType
	3, // 2: PlayerLoginCsReq.GANJMFCLNFL:type_name -> GBAMNAPGHFF
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_PlayerLoginCsReq_proto_init() }
func file_PlayerLoginCsReq_proto_init() {
	if File_PlayerLoginCsReq_proto != nil {
		return
	}
	file_LanguageType_proto_init()
	file_GBAMNAPGHFF_proto_init()
	file_PlatformType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_PlayerLoginCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerLoginCsReq); i {
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
			RawDescriptor: file_PlayerLoginCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PlayerLoginCsReq_proto_goTypes,
		DependencyIndexes: file_PlayerLoginCsReq_proto_depIdxs,
		MessageInfos:      file_PlayerLoginCsReq_proto_msgTypes,
	}.Build()
	File_PlayerLoginCsReq_proto = out.File
	file_PlayerLoginCsReq_proto_rawDesc = nil
	file_PlayerLoginCsReq_proto_goTypes = nil
	file_PlayerLoginCsReq_proto_depIdxs = nil
}
