// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GMFILGICFFC.proto

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

type GMFILGICFFC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OLCPIFJCBFP uint32 `protobuf:"varint,9,opt,name=OLCPIFJCBFP,proto3" json:"OLCPIFJCBFP,omitempty"`
	KNCCCCGBPED bool   `protobuf:"varint,2,opt,name=KNCCCCGBPED,proto3" json:"KNCCCCGBPED,omitempty"`
	GroupId     uint32 `protobuf:"varint,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	INEMBEBIPAN uint32 `protobuf:"varint,14,opt,name=INEMBEBIPAN,proto3" json:"INEMBEBIPAN,omitempty"`
	AHLOGIGDNLF uint32 `protobuf:"varint,1,opt,name=AHLOGIGDNLF,proto3" json:"AHLOGIGDNLF,omitempty"`
}

func (x *GMFILGICFFC) Reset() {
	*x = GMFILGICFFC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GMFILGICFFC_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GMFILGICFFC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GMFILGICFFC) ProtoMessage() {}

func (x *GMFILGICFFC) ProtoReflect() protoreflect.Message {
	mi := &file_GMFILGICFFC_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GMFILGICFFC.ProtoReflect.Descriptor instead.
func (*GMFILGICFFC) Descriptor() ([]byte, []int) {
	return file_GMFILGICFFC_proto_rawDescGZIP(), []int{0}
}

func (x *GMFILGICFFC) GetOLCPIFJCBFP() uint32 {
	if x != nil {
		return x.OLCPIFJCBFP
	}
	return 0
}

func (x *GMFILGICFFC) GetKNCCCCGBPED() bool {
	if x != nil {
		return x.KNCCCCGBPED
	}
	return false
}

func (x *GMFILGICFFC) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *GMFILGICFFC) GetINEMBEBIPAN() uint32 {
	if x != nil {
		return x.INEMBEBIPAN
	}
	return 0
}

func (x *GMFILGICFFC) GetAHLOGIGDNLF() uint32 {
	if x != nil {
		return x.AHLOGIGDNLF
	}
	return 0
}

var File_GMFILGICFFC_proto protoreflect.FileDescriptor

var file_GMFILGICFFC_proto_rawDesc = []byte{
	0x0a, 0x11, 0x47, 0x4d, 0x46, 0x49, 0x4c, 0x47, 0x49, 0x43, 0x46, 0x46, 0x43, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x0b, 0x47, 0x4d, 0x46, 0x49, 0x4c, 0x47, 0x49, 0x43,
	0x46, 0x46, 0x43, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x4c, 0x43, 0x50, 0x49, 0x46, 0x4a, 0x43, 0x42,
	0x46, 0x50, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x4c, 0x43, 0x50, 0x49, 0x46,
	0x4a, 0x43, 0x42, 0x46, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x4b, 0x4e, 0x43, 0x43, 0x43, 0x43, 0x47,
	0x42, 0x50, 0x45, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4b, 0x4e, 0x43, 0x43,
	0x43, 0x43, 0x47, 0x42, 0x50, 0x45, 0x44, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x4e, 0x45, 0x4d, 0x42, 0x45, 0x42, 0x49, 0x50, 0x41,
	0x4e, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x4e, 0x45, 0x4d, 0x42, 0x45, 0x42,
	0x49, 0x50, 0x41, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x48, 0x4c, 0x4f, 0x47, 0x49, 0x47, 0x44,
	0x4e, 0x4c, 0x46, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x41, 0x48, 0x4c, 0x4f, 0x47,
	0x49, 0x47, 0x44, 0x4e, 0x4c, 0x46, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GMFILGICFFC_proto_rawDescOnce sync.Once
	file_GMFILGICFFC_proto_rawDescData = file_GMFILGICFFC_proto_rawDesc
)

func file_GMFILGICFFC_proto_rawDescGZIP() []byte {
	file_GMFILGICFFC_proto_rawDescOnce.Do(func() {
		file_GMFILGICFFC_proto_rawDescData = protoimpl.X.CompressGZIP(file_GMFILGICFFC_proto_rawDescData)
	})
	return file_GMFILGICFFC_proto_rawDescData
}

var file_GMFILGICFFC_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GMFILGICFFC_proto_goTypes = []interface{}{
	(*GMFILGICFFC)(nil), // 0: GMFILGICFFC
}
var file_GMFILGICFFC_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GMFILGICFFC_proto_init() }
func file_GMFILGICFFC_proto_init() {
	if File_GMFILGICFFC_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GMFILGICFFC_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GMFILGICFFC); i {
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
			RawDescriptor: file_GMFILGICFFC_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GMFILGICFFC_proto_goTypes,
		DependencyIndexes: file_GMFILGICFFC_proto_depIdxs,
		MessageInfos:      file_GMFILGICFFC_proto_msgTypes,
	}.Build()
	File_GMFILGICFFC_proto = out.File
	file_GMFILGICFFC_proto_rawDesc = nil
	file_GMFILGICFFC_proto_goTypes = nil
	file_GMFILGICFFC_proto_depIdxs = nil
}
