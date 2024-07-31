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

	Language      LanguageType `protobuf:"varint,14,opt,name=language,proto3,enum=LanguageType" json:"language,omitempty"`
	AIONAIPAKJC   string       `protobuf:"bytes,1032,opt,name=AIONAIPAKJC,proto3" json:"AIONAIPAKJC,omitempty"`
	EJKNHJGAIBM   string       `protobuf:"bytes,1133,opt,name=EJKNHJGAIBM,proto3" json:"EJKNHJGAIBM,omitempty"`
	BDCEEMHGBHI   uint32       `protobuf:"varint,1208,opt,name=BDCEEMHGBHI,proto3" json:"BDCEEMHGBHI,omitempty"`
	Signature     string       `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	NGMENNDFODP   string       `protobuf:"bytes,8,opt,name=NGMENNDFODP,proto3" json:"NGMENNDFODP,omitempty"`
	HAIKFFCCIFL   string       `protobuf:"bytes,9,opt,name=HAIKFFCCIFL,proto3" json:"HAIKFFCCIFL,omitempty"`
	OKAFNHKOKFJ   string       `protobuf:"bytes,6,opt,name=OKAFNHKOKFJ,proto3" json:"OKAFNHKOKFJ,omitempty"`
	DHNMDCNBNLD   string       `protobuf:"bytes,13,opt,name=DHNMDCNBNLD,proto3" json:"DHNMDCNBNLD,omitempty"`
	HHLOJBCCNPL   string       `protobuf:"bytes,10,opt,name=HHLOJBCCNPL,proto3" json:"HHLOJBCCNPL,omitempty"`
	JAKEDMONFKM   uint32       `protobuf:"varint,794,opt,name=JAKEDMONFKM,proto3" json:"JAKEDMONFKM,omitempty"`
	JAPDDCEJDKN   *MKDGHOMDDNM `protobuf:"bytes,387,opt,name=JAPDDCEJDKN,proto3" json:"JAPDDCEJDKN,omitempty"`
	HLDGIGHGHCF   string       `protobuf:"bytes,2,opt,name=HLDGIGHGHCF,proto3" json:"HLDGIGHGHCF,omitempty"`
	Platform      PlatformType `protobuf:"varint,7,opt,name=platform,proto3,enum=PlatformType" json:"platform,omitempty"`
	MOLBBPGIOOJ   string       `protobuf:"bytes,1928,opt,name=MOLBBPGIOOJ,proto3" json:"MOLBBPGIOOJ,omitempty"`
	JBJFBOCCBBA   string       `protobuf:"bytes,1,opt,name=JBJFBOCCBBA,proto3" json:"JBJFBOCCBBA,omitempty"`
	PIMJDJPHEIF   string       `protobuf:"bytes,904,opt,name=PIMJDJPHEIF,proto3" json:"PIMJDJPHEIF,omitempty"`
	ResVersion    uint32       `protobuf:"varint,11,opt,name=res_version,json=resVersion,proto3" json:"res_version,omitempty"`
	CLGEKPFAPNG   bool         `protobuf:"varint,1104,opt,name=CLGEKPFAPNG,proto3" json:"CLGEKPFAPNG,omitempty"`
	HHHIHGDICGF   string       `protobuf:"bytes,3,opt,name=HHHIHGDICGF,proto3" json:"HHHIHGDICGF,omitempty"`
	ClientVersion string       `protobuf:"bytes,12,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	DKMGKNMJEPM   string       `protobuf:"bytes,1458,opt,name=DKMGKNMJEPM,proto3" json:"DKMGKNMJEPM,omitempty"`
	PHLGMEIBOIM   uint32       `protobuf:"varint,15,opt,name=PHLGMEIBOIM,proto3" json:"PHLGMEIBOIM,omitempty"`
	LoginRandom   uint64       `protobuf:"varint,4,opt,name=login_random,json=loginRandom,proto3" json:"login_random,omitempty"`
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

func (x *PlayerLoginCsReq) GetLanguage() LanguageType {
	if x != nil {
		return x.Language
	}
	return LanguageType_LANGUAGE_NONE
}

func (x *PlayerLoginCsReq) GetAIONAIPAKJC() string {
	if x != nil {
		return x.AIONAIPAKJC
	}
	return ""
}

func (x *PlayerLoginCsReq) GetEJKNHJGAIBM() string {
	if x != nil {
		return x.EJKNHJGAIBM
	}
	return ""
}

func (x *PlayerLoginCsReq) GetBDCEEMHGBHI() uint32 {
	if x != nil {
		return x.BDCEEMHGBHI
	}
	return 0
}

func (x *PlayerLoginCsReq) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *PlayerLoginCsReq) GetNGMENNDFODP() string {
	if x != nil {
		return x.NGMENNDFODP
	}
	return ""
}

func (x *PlayerLoginCsReq) GetHAIKFFCCIFL() string {
	if x != nil {
		return x.HAIKFFCCIFL
	}
	return ""
}

func (x *PlayerLoginCsReq) GetOKAFNHKOKFJ() string {
	if x != nil {
		return x.OKAFNHKOKFJ
	}
	return ""
}

func (x *PlayerLoginCsReq) GetDHNMDCNBNLD() string {
	if x != nil {
		return x.DHNMDCNBNLD
	}
	return ""
}

func (x *PlayerLoginCsReq) GetHHLOJBCCNPL() string {
	if x != nil {
		return x.HHLOJBCCNPL
	}
	return ""
}

func (x *PlayerLoginCsReq) GetJAKEDMONFKM() uint32 {
	if x != nil {
		return x.JAKEDMONFKM
	}
	return 0
}

func (x *PlayerLoginCsReq) GetJAPDDCEJDKN() *MKDGHOMDDNM {
	if x != nil {
		return x.JAPDDCEJDKN
	}
	return nil
}

func (x *PlayerLoginCsReq) GetHLDGIGHGHCF() string {
	if x != nil {
		return x.HLDGIGHGHCF
	}
	return ""
}

func (x *PlayerLoginCsReq) GetPlatform() PlatformType {
	if x != nil {
		return x.Platform
	}
	return PlatformType_EDITOR
}

func (x *PlayerLoginCsReq) GetMOLBBPGIOOJ() string {
	if x != nil {
		return x.MOLBBPGIOOJ
	}
	return ""
}

func (x *PlayerLoginCsReq) GetJBJFBOCCBBA() string {
	if x != nil {
		return x.JBJFBOCCBBA
	}
	return ""
}

func (x *PlayerLoginCsReq) GetPIMJDJPHEIF() string {
	if x != nil {
		return x.PIMJDJPHEIF
	}
	return ""
}

func (x *PlayerLoginCsReq) GetResVersion() uint32 {
	if x != nil {
		return x.ResVersion
	}
	return 0
}

func (x *PlayerLoginCsReq) GetCLGEKPFAPNG() bool {
	if x != nil {
		return x.CLGEKPFAPNG
	}
	return false
}

func (x *PlayerLoginCsReq) GetHHHIHGDICGF() string {
	if x != nil {
		return x.HHHIHGDICGF
	}
	return ""
}

func (x *PlayerLoginCsReq) GetClientVersion() string {
	if x != nil {
		return x.ClientVersion
	}
	return ""
}

func (x *PlayerLoginCsReq) GetDKMGKNMJEPM() string {
	if x != nil {
		return x.DKMGKNMJEPM
	}
	return ""
}

func (x *PlayerLoginCsReq) GetPHLGMEIBOIM() uint32 {
	if x != nil {
		return x.PHLGMEIBOIM
	}
	return 0
}

func (x *PlayerLoginCsReq) GetLoginRandom() uint64 {
	if x != nil {
		return x.LoginRandom
	}
	return 0
}

var File_PlayerLoginCsReq_proto protoreflect.FileDescriptor

var file_PlayerLoginCsReq_proto_rawDesc = []byte{
	0x0a, 0x16, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x73, 0x52,
	0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x50, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x4d, 0x4b, 0x44, 0x47, 0x48, 0x4f, 0x4d, 0x44, 0x44, 0x4e, 0x4d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xec, 0x06, 0x0a, 0x10, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x29, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x41, 0x49, 0x4f, 0x4e, 0x41, 0x49, 0x50, 0x41, 0x4b,
	0x4a, 0x43, 0x18, 0x88, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x49, 0x4f, 0x4e, 0x41,
	0x49, 0x50, 0x41, 0x4b, 0x4a, 0x43, 0x12, 0x21, 0x0a, 0x0b, 0x45, 0x4a, 0x4b, 0x4e, 0x48, 0x4a,
	0x47, 0x41, 0x49, 0x42, 0x4d, 0x18, 0xed, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x45, 0x4a,
	0x4b, 0x4e, 0x48, 0x4a, 0x47, 0x41, 0x49, 0x42, 0x4d, 0x12, 0x21, 0x0a, 0x0b, 0x42, 0x44, 0x43,
	0x45, 0x45, 0x4d, 0x48, 0x47, 0x42, 0x48, 0x49, 0x18, 0xb8, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x42, 0x44, 0x43, 0x45, 0x45, 0x4d, 0x48, 0x47, 0x42, 0x48, 0x49, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x47,
	0x4d, 0x45, 0x4e, 0x4e, 0x44, 0x46, 0x4f, 0x44, 0x50, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x4e, 0x47, 0x4d, 0x45, 0x4e, 0x4e, 0x44, 0x46, 0x4f, 0x44, 0x50, 0x12, 0x20, 0x0a, 0x0b,
	0x48, 0x41, 0x49, 0x4b, 0x46, 0x46, 0x43, 0x43, 0x49, 0x46, 0x4c, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x48, 0x41, 0x49, 0x4b, 0x46, 0x46, 0x43, 0x43, 0x49, 0x46, 0x4c, 0x12, 0x20,
	0x0a, 0x0b, 0x4f, 0x4b, 0x41, 0x46, 0x4e, 0x48, 0x4b, 0x4f, 0x4b, 0x46, 0x4a, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x4f, 0x4b, 0x41, 0x46, 0x4e, 0x48, 0x4b, 0x4f, 0x4b, 0x46, 0x4a,
	0x12, 0x20, 0x0a, 0x0b, 0x44, 0x48, 0x4e, 0x4d, 0x44, 0x43, 0x4e, 0x42, 0x4e, 0x4c, 0x44, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x48, 0x4e, 0x4d, 0x44, 0x43, 0x4e, 0x42, 0x4e,
	0x4c, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x48, 0x4c, 0x4f, 0x4a, 0x42, 0x43, 0x43, 0x4e, 0x50,
	0x4c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x48, 0x48, 0x4c, 0x4f, 0x4a, 0x42, 0x43,
	0x43, 0x4e, 0x50, 0x4c, 0x12, 0x21, 0x0a, 0x0b, 0x4a, 0x41, 0x4b, 0x45, 0x44, 0x4d, 0x4f, 0x4e,
	0x46, 0x4b, 0x4d, 0x18, 0x9a, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x41, 0x4b, 0x45,
	0x44, 0x4d, 0x4f, 0x4e, 0x46, 0x4b, 0x4d, 0x12, 0x2f, 0x0a, 0x0b, 0x4a, 0x41, 0x50, 0x44, 0x44,
	0x43, 0x45, 0x4a, 0x44, 0x4b, 0x4e, 0x18, 0x83, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x4d, 0x4b, 0x44, 0x47, 0x48, 0x4f, 0x4d, 0x44, 0x44, 0x4e, 0x4d, 0x52, 0x0b, 0x4a, 0x41, 0x50,
	0x44, 0x44, 0x43, 0x45, 0x4a, 0x44, 0x4b, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4c, 0x44, 0x47,
	0x49, 0x47, 0x48, 0x47, 0x48, 0x43, 0x46, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x48,
	0x4c, 0x44, 0x47, 0x49, 0x47, 0x48, 0x47, 0x48, 0x43, 0x46, 0x12, 0x29, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x50,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x21, 0x0a, 0x0b, 0x4d, 0x4f, 0x4c, 0x42, 0x42, 0x50, 0x47,
	0x49, 0x4f, 0x4f, 0x4a, 0x18, 0x88, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d, 0x4f, 0x4c,
	0x42, 0x42, 0x50, 0x47, 0x49, 0x4f, 0x4f, 0x4a, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x42, 0x4a, 0x46,
	0x42, 0x4f, 0x43, 0x43, 0x42, 0x42, 0x41, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4a,
	0x42, 0x4a, 0x46, 0x42, 0x4f, 0x43, 0x43, 0x42, 0x42, 0x41, 0x12, 0x21, 0x0a, 0x0b, 0x50, 0x49,
	0x4d, 0x4a, 0x44, 0x4a, 0x50, 0x48, 0x45, 0x49, 0x46, 0x18, 0x88, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x50, 0x49, 0x4d, 0x4a, 0x44, 0x4a, 0x50, 0x48, 0x45, 0x49, 0x46, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21,
	0x0a, 0x0b, 0x43, 0x4c, 0x47, 0x45, 0x4b, 0x50, 0x46, 0x41, 0x50, 0x4e, 0x47, 0x18, 0xd0, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x43, 0x4c, 0x47, 0x45, 0x4b, 0x50, 0x46, 0x41, 0x50, 0x4e,
	0x47, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x48, 0x48, 0x49, 0x48, 0x47, 0x44, 0x49, 0x43, 0x47, 0x46,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x48, 0x48, 0x48, 0x49, 0x48, 0x47, 0x44, 0x49,
	0x43, 0x47, 0x46, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0b, 0x44, 0x4b,
	0x4d, 0x47, 0x4b, 0x4e, 0x4d, 0x4a, 0x45, 0x50, 0x4d, 0x18, 0xb2, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x4b, 0x4d, 0x47, 0x4b, 0x4e, 0x4d, 0x4a, 0x45, 0x50, 0x4d, 0x12, 0x20, 0x0a,
	0x0b, 0x50, 0x48, 0x4c, 0x47, 0x4d, 0x45, 0x49, 0x42, 0x4f, 0x49, 0x4d, 0x18, 0x0f, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x50, 0x48, 0x4c, 0x47, 0x4d, 0x45, 0x49, 0x42, 0x4f, 0x49, 0x4d, 0x12,
	0x21, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x61, 0x6e, 0x64,
	0x6f, 0x6d, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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
	(*MKDGHOMDDNM)(nil),      // 2: MKDGHOMDDNM
	(PlatformType)(0),        // 3: PlatformType
}
var file_PlayerLoginCsReq_proto_depIdxs = []int32{
	1, // 0: PlayerLoginCsReq.language:type_name -> LanguageType
	2, // 1: PlayerLoginCsReq.JAPDDCEJDKN:type_name -> MKDGHOMDDNM
	3, // 2: PlayerLoginCsReq.platform:type_name -> PlatformType
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
	file_PlatformType_proto_init()
	file_MKDGHOMDDNM_proto_init()
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
