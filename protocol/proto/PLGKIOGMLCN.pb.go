// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: PLGKIOGMLCN.proto

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

type PLGKIOGMLCN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OLKLCKCPNMM uint32 `protobuf:"varint,8,opt,name=OLKLCKCPNMM,proto3" json:"OLKLCKCPNMM,omitempty"`
	HMDIKALOMIP uint32 `protobuf:"varint,9,opt,name=HMDIKALOMIP,proto3" json:"HMDIKALOMIP,omitempty"`
	LBJDPLONOLA uint32 `protobuf:"varint,11,opt,name=LBJDPLONOLA,proto3" json:"LBJDPLONOLA,omitempty"`
	MEMEMFKPPHF uint32 `protobuf:"varint,13,opt,name=MEMEMFKPPHF,proto3" json:"MEMEMFKPPHF,omitempty"`
	Uid         uint32 `protobuf:"varint,5,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *PLGKIOGMLCN) Reset() {
	*x = PLGKIOGMLCN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_PLGKIOGMLCN_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PLGKIOGMLCN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PLGKIOGMLCN) ProtoMessage() {}

func (x *PLGKIOGMLCN) ProtoReflect() protoreflect.Message {
	mi := &file_PLGKIOGMLCN_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PLGKIOGMLCN.ProtoReflect.Descriptor instead.
func (*PLGKIOGMLCN) Descriptor() ([]byte, []int) {
	return file_PLGKIOGMLCN_proto_rawDescGZIP(), []int{0}
}

func (x *PLGKIOGMLCN) GetOLKLCKCPNMM() uint32 {
	if x != nil {
		return x.OLKLCKCPNMM
	}
	return 0
}

func (x *PLGKIOGMLCN) GetHMDIKALOMIP() uint32 {
	if x != nil {
		return x.HMDIKALOMIP
	}
	return 0
}

func (x *PLGKIOGMLCN) GetLBJDPLONOLA() uint32 {
	if x != nil {
		return x.LBJDPLONOLA
	}
	return 0
}

func (x *PLGKIOGMLCN) GetMEMEMFKPPHF() uint32 {
	if x != nil {
		return x.MEMEMFKPPHF
	}
	return 0
}

func (x *PLGKIOGMLCN) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

var File_PLGKIOGMLCN_proto protoreflect.FileDescriptor

var file_PLGKIOGMLCN_proto_rawDesc = []byte{
	0x0a, 0x11, 0x50, 0x4c, 0x47, 0x4b, 0x49, 0x4f, 0x47, 0x4d, 0x4c, 0x43, 0x4e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x0b, 0x50, 0x4c, 0x47, 0x4b, 0x49, 0x4f, 0x47, 0x4d,
	0x4c, 0x43, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x4c, 0x4b, 0x4c, 0x43, 0x4b, 0x43, 0x50, 0x4e,
	0x4d, 0x4d, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x4c, 0x4b, 0x4c, 0x43, 0x4b,
	0x43, 0x50, 0x4e, 0x4d, 0x4d, 0x12, 0x20, 0x0a, 0x0b, 0x48, 0x4d, 0x44, 0x49, 0x4b, 0x41, 0x4c,
	0x4f, 0x4d, 0x49, 0x50, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x48, 0x4d, 0x44, 0x49,
	0x4b, 0x41, 0x4c, 0x4f, 0x4d, 0x49, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x42, 0x4a, 0x44, 0x50,
	0x4c, 0x4f, 0x4e, 0x4f, 0x4c, 0x41, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x42,
	0x4a, 0x44, 0x50, 0x4c, 0x4f, 0x4e, 0x4f, 0x4c, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x45, 0x4d,
	0x45, 0x4d, 0x46, 0x4b, 0x50, 0x50, 0x48, 0x46, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b,
	0x4d, 0x45, 0x4d, 0x45, 0x4d, 0x46, 0x4b, 0x50, 0x50, 0x48, 0x46, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x42, 0x28, 0x5a,
	0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_PLGKIOGMLCN_proto_rawDescOnce sync.Once
	file_PLGKIOGMLCN_proto_rawDescData = file_PLGKIOGMLCN_proto_rawDesc
)

func file_PLGKIOGMLCN_proto_rawDescGZIP() []byte {
	file_PLGKIOGMLCN_proto_rawDescOnce.Do(func() {
		file_PLGKIOGMLCN_proto_rawDescData = protoimpl.X.CompressGZIP(file_PLGKIOGMLCN_proto_rawDescData)
	})
	return file_PLGKIOGMLCN_proto_rawDescData
}

var file_PLGKIOGMLCN_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_PLGKIOGMLCN_proto_goTypes = []interface{}{
	(*PLGKIOGMLCN)(nil), // 0: PLGKIOGMLCN
}
var file_PLGKIOGMLCN_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_PLGKIOGMLCN_proto_init() }
func file_PLGKIOGMLCN_proto_init() {
	if File_PLGKIOGMLCN_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_PLGKIOGMLCN_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PLGKIOGMLCN); i {
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
			RawDescriptor: file_PLGKIOGMLCN_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_PLGKIOGMLCN_proto_goTypes,
		DependencyIndexes: file_PLGKIOGMLCN_proto_depIdxs,
		MessageInfos:      file_PLGKIOGMLCN_proto_msgTypes,
	}.Build()
	File_PLGKIOGMLCN_proto = out.File
	file_PLGKIOGMLCN_proto_rawDesc = nil
	file_PLGKIOGMLCN_proto_goTypes = nil
	file_PLGKIOGMLCN_proto_depIdxs = nil
}
