// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: DPNMFCCFJBE.proto

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

type DPNMFCCFJBE struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GEBIOEEKCEF uint32 `protobuf:"varint,2,opt,name=GEBIOEEKCEF,proto3" json:"GEBIOEEKCEF,omitempty"`
	JBLIKFLGJIA uint32 `protobuf:"varint,5,opt,name=JBLIKFLGJIA,proto3" json:"JBLIKFLGJIA,omitempty"`
	OKFIGEFNFAJ bool   `protobuf:"varint,1,opt,name=OKFIGEFNFAJ,proto3" json:"OKFIGEFNFAJ,omitempty"`
	Level       uint32 `protobuf:"varint,13,opt,name=level,proto3" json:"level,omitempty"`
}

func (x *DPNMFCCFJBE) Reset() {
	*x = DPNMFCCFJBE{}
	if protoimpl.UnsafeEnabled {
		mi := &file_DPNMFCCFJBE_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DPNMFCCFJBE) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DPNMFCCFJBE) ProtoMessage() {}

func (x *DPNMFCCFJBE) ProtoReflect() protoreflect.Message {
	mi := &file_DPNMFCCFJBE_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DPNMFCCFJBE.ProtoReflect.Descriptor instead.
func (*DPNMFCCFJBE) Descriptor() ([]byte, []int) {
	return file_DPNMFCCFJBE_proto_rawDescGZIP(), []int{0}
}

func (x *DPNMFCCFJBE) GetGEBIOEEKCEF() uint32 {
	if x != nil {
		return x.GEBIOEEKCEF
	}
	return 0
}

func (x *DPNMFCCFJBE) GetJBLIKFLGJIA() uint32 {
	if x != nil {
		return x.JBLIKFLGJIA
	}
	return 0
}

func (x *DPNMFCCFJBE) GetOKFIGEFNFAJ() bool {
	if x != nil {
		return x.OKFIGEFNFAJ
	}
	return false
}

func (x *DPNMFCCFJBE) GetLevel() uint32 {
	if x != nil {
		return x.Level
	}
	return 0
}

var File_DPNMFCCFJBE_proto protoreflect.FileDescriptor

var file_DPNMFCCFJBE_proto_rawDesc = []byte{
	0x0a, 0x11, 0x44, 0x50, 0x4e, 0x4d, 0x46, 0x43, 0x43, 0x46, 0x4a, 0x42, 0x45, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x89, 0x01, 0x0a, 0x0b, 0x44, 0x50, 0x4e, 0x4d, 0x46, 0x43, 0x43, 0x46,
	0x4a, 0x42, 0x45, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x45, 0x42, 0x49, 0x4f, 0x45, 0x45, 0x4b, 0x43,
	0x45, 0x46, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x45, 0x42, 0x49, 0x4f, 0x45,
	0x45, 0x4b, 0x43, 0x45, 0x46, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x42, 0x4c, 0x49, 0x4b, 0x46, 0x4c,
	0x47, 0x4a, 0x49, 0x41, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4a, 0x42, 0x4c, 0x49,
	0x4b, 0x46, 0x4c, 0x47, 0x4a, 0x49, 0x41, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x4b, 0x46, 0x49, 0x47,
	0x45, 0x46, 0x4e, 0x46, 0x41, 0x4a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4f, 0x4b,
	0x46, 0x49, 0x47, 0x45, 0x46, 0x4e, 0x46, 0x41, 0x4a, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x65, 0x76,
	0x65, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x42,
	0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_DPNMFCCFJBE_proto_rawDescOnce sync.Once
	file_DPNMFCCFJBE_proto_rawDescData = file_DPNMFCCFJBE_proto_rawDesc
)

func file_DPNMFCCFJBE_proto_rawDescGZIP() []byte {
	file_DPNMFCCFJBE_proto_rawDescOnce.Do(func() {
		file_DPNMFCCFJBE_proto_rawDescData = protoimpl.X.CompressGZIP(file_DPNMFCCFJBE_proto_rawDescData)
	})
	return file_DPNMFCCFJBE_proto_rawDescData
}

var file_DPNMFCCFJBE_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_DPNMFCCFJBE_proto_goTypes = []interface{}{
	(*DPNMFCCFJBE)(nil), // 0: DPNMFCCFJBE
}
var file_DPNMFCCFJBE_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_DPNMFCCFJBE_proto_init() }
func file_DPNMFCCFJBE_proto_init() {
	if File_DPNMFCCFJBE_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_DPNMFCCFJBE_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DPNMFCCFJBE); i {
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
			RawDescriptor: file_DPNMFCCFJBE_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_DPNMFCCFJBE_proto_goTypes,
		DependencyIndexes: file_DPNMFCCFJBE_proto_depIdxs,
		MessageInfos:      file_DPNMFCCFJBE_proto_msgTypes,
	}.Build()
	File_DPNMFCCFJBE_proto = out.File
	file_DPNMFCCFJBE_proto_rawDesc = nil
	file_DPNMFCCFJBE_proto_goTypes = nil
	file_DPNMFCCFJBE_proto_depIdxs = nil
}