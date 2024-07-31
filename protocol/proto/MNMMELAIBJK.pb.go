// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MNMMELAIBJK.proto

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

type MNMMELAIBJK struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level       uint32         `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	DGPECHFOPHK uint32         `protobuf:"varint,13,opt,name=DGPECHFOPHK,proto3" json:"DGPECHFOPHK,omitempty"`
	GLGFCBDLOOI []*GGJDIKHDJOM `protobuf:"bytes,8,rep,name=GLGFCBDLOOI,proto3" json:"GLGFCBDLOOI,omitempty"`
	OIIBFPIBLBK uint32         `protobuf:"varint,9,opt,name=OIIBFPIBLBK,proto3" json:"OIIBFPIBLBK,omitempty"`
	HMBIHJFHHHI uint32         `protobuf:"varint,14,opt,name=HMBIHJFHHHI,proto3" json:"HMBIHJFHHHI,omitempty"`
	AreaId      uint32         `protobuf:"varint,2,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
}

func (x *MNMMELAIBJK) Reset() {
	*x = MNMMELAIBJK{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MNMMELAIBJK_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MNMMELAIBJK) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MNMMELAIBJK) ProtoMessage() {}

func (x *MNMMELAIBJK) ProtoReflect() protoreflect.Message {
	mi := &file_MNMMELAIBJK_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MNMMELAIBJK.ProtoReflect.Descriptor instead.
func (*MNMMELAIBJK) Descriptor() ([]byte, []int) {
	return file_MNMMELAIBJK_proto_rawDescGZIP(), []int{0}
}

func (x *MNMMELAIBJK) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *MNMMELAIBJK) GetDGPECHFOPHK() uint32 {
	if x != nil {
		return x.DGPECHFOPHK
	}
	return 0
}

func (x *MNMMELAIBJK) GetGLGFCBDLOOI() []*GGJDIKHDJOM {
	if x != nil {
		return x.GLGFCBDLOOI
	}
	return nil
}

func (x *MNMMELAIBJK) GetOIIBFPIBLBK() uint32 {
	if x != nil {
		return x.OIIBFPIBLBK
	}
	return 0
}

func (x *MNMMELAIBJK) GetHMBIHJFHHHI() uint32 {
	if x != nil {
		return x.HMBIHJFHHHI
	}
	return 0
}

func (x *MNMMELAIBJK) GetAreaId() uint32 {
	if x != nil {
		return x.AreaId
	}
	return 0
}

var File_MNMMELAIBJK_proto protoreflect.FileDescriptor

var file_MNMMELAIBJK_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4d, 0x4e, 0x4d, 0x4d, 0x45, 0x4c, 0x41, 0x49, 0x42, 0x4a, 0x4b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x47, 0x47, 0x4a, 0x44, 0x49, 0x4b, 0x48, 0x44, 0x4a, 0x4f, 0x4d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd2, 0x01, 0x0a, 0x0b, 0x4d, 0x4e, 0x4d, 0x4d, 0x45,
	0x4c, 0x41, 0x49, 0x42, 0x4a, 0x4b, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x47, 0x50, 0x45, 0x43, 0x48, 0x46, 0x4f, 0x50, 0x48, 0x4b, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x44, 0x47, 0x50, 0x45, 0x43, 0x48, 0x46, 0x4f, 0x50, 0x48, 0x4b, 0x12, 0x2e,
	0x0a, 0x0b, 0x47, 0x4c, 0x47, 0x46, 0x43, 0x42, 0x44, 0x4c, 0x4f, 0x4f, 0x49, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x47, 0x47, 0x4a, 0x44, 0x49, 0x4b, 0x48, 0x44, 0x4a, 0x4f,
	0x4d, 0x52, 0x0b, 0x47, 0x4c, 0x47, 0x46, 0x43, 0x42, 0x44, 0x4c, 0x4f, 0x4f, 0x49, 0x12, 0x20,
	0x0a, 0x0b, 0x4f, 0x49, 0x49, 0x42, 0x46, 0x50, 0x49, 0x42, 0x4c, 0x42, 0x4b, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x49, 0x49, 0x42, 0x46, 0x50, 0x49, 0x42, 0x4c, 0x42, 0x4b,
	0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4d, 0x42, 0x49, 0x48, 0x4a, 0x46, 0x48, 0x48, 0x48, 0x49, 0x18,
	0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x4d, 0x42, 0x49, 0x48, 0x4a, 0x46, 0x48, 0x48,
	0x48, 0x49, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x61, 0x72, 0x65, 0x61, 0x49, 0x64, 0x42, 0x28, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MNMMELAIBJK_proto_rawDescOnce sync.Once
	file_MNMMELAIBJK_proto_rawDescData = file_MNMMELAIBJK_proto_rawDesc
)

func file_MNMMELAIBJK_proto_rawDescGZIP() []byte {
	file_MNMMELAIBJK_proto_rawDescOnce.Do(func() {
		file_MNMMELAIBJK_proto_rawDescData = protoimpl.X.CompressGZIP(file_MNMMELAIBJK_proto_rawDescData)
	})
	return file_MNMMELAIBJK_proto_rawDescData
}

var file_MNMMELAIBJK_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MNMMELAIBJK_proto_goTypes = []interface{}{
	(*MNMMELAIBJK)(nil), // 0: MNMMELAIBJK
	(*GGJDIKHDJOM)(nil), // 1: GGJDIKHDJOM
}
var file_MNMMELAIBJK_proto_depIdxs = []int32{
	1, // 0: MNMMELAIBJK.GLGFCBDLOOI:type_name -> GGJDIKHDJOM
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_MNMMELAIBJK_proto_init() }
func file_MNMMELAIBJK_proto_init() {
	if File_MNMMELAIBJK_proto != nil {
		return
	}
	file_GGJDIKHDJOM_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MNMMELAIBJK_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MNMMELAIBJK); i {
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
			RawDescriptor: file_MNMMELAIBJK_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MNMMELAIBJK_proto_goTypes,
		DependencyIndexes: file_MNMMELAIBJK_proto_depIdxs,
		MessageInfos:      file_MNMMELAIBJK_proto_msgTypes,
	}.Build()
	File_MNMMELAIBJK_proto = out.File
	file_MNMMELAIBJK_proto_rawDesc = nil
	file_MNMMELAIBJK_proto_goTypes = nil
	file_MNMMELAIBJK_proto_depIdxs = nil
}
