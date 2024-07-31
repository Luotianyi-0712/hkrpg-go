// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: KMEDPBBAOHC.proto

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

type KMEDPBBAOHC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemList    []uint32 `protobuf:"varint,2,rep,packed,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
	CPHKOCPKCOI bool     `protobuf:"varint,15,opt,name=CPHKOCPKCOI,proto3" json:"CPHKOCPKCOI,omitempty"`
	UniqueId    uint32   `protobuf:"varint,5,opt,name=unique_id,json=uniqueId,proto3" json:"unique_id,omitempty"`
	LLJJPOFJNKK []uint32 `protobuf:"varint,9,rep,packed,name=LLJJPOFJNKK,proto3" json:"LLJJPOFJNKK,omitempty"`
	EFICNPDCKPK uint32   `protobuf:"varint,4,opt,name=EFICNPDCKPK,proto3" json:"EFICNPDCKPK,omitempty"`
	GJIOOFGKPJC uint32   `protobuf:"varint,1,opt,name=GJIOOFGKPJC,proto3" json:"GJIOOFGKPJC,omitempty"`
}

func (x *KMEDPBBAOHC) Reset() {
	*x = KMEDPBBAOHC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KMEDPBBAOHC_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KMEDPBBAOHC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KMEDPBBAOHC) ProtoMessage() {}

func (x *KMEDPBBAOHC) ProtoReflect() protoreflect.Message {
	mi := &file_KMEDPBBAOHC_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KMEDPBBAOHC.ProtoReflect.Descriptor instead.
func (*KMEDPBBAOHC) Descriptor() ([]byte, []int) {
	return file_KMEDPBBAOHC_proto_rawDescGZIP(), []int{0}
}

func (x *KMEDPBBAOHC) GetItemList() []uint32 {
	if x != nil {
		return x.ItemList
	}
	return nil
}

func (x *KMEDPBBAOHC) GetCPHKOCPKCOI() bool {
	if x != nil {
		return x.CPHKOCPKCOI
	}
	return false
}

func (x *KMEDPBBAOHC) GetUniqueId() uint32 {
	if x != nil {
		return x.UniqueId
	}
	return 0
}

func (x *KMEDPBBAOHC) GetLLJJPOFJNKK() []uint32 {
	if x != nil {
		return x.LLJJPOFJNKK
	}
	return nil
}

func (x *KMEDPBBAOHC) GetEFICNPDCKPK() uint32 {
	if x != nil {
		return x.EFICNPDCKPK
	}
	return 0
}

func (x *KMEDPBBAOHC) GetGJIOOFGKPJC() uint32 {
	if x != nil {
		return x.GJIOOFGKPJC
	}
	return 0
}

var File_KMEDPBBAOHC_proto protoreflect.FileDescriptor

var file_KMEDPBBAOHC_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4b, 0x4d, 0x45, 0x44, 0x50, 0x42, 0x42, 0x41, 0x4f, 0x48, 0x43, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x0b, 0x4b, 0x4d, 0x45, 0x44, 0x50, 0x42, 0x42, 0x41,
	0x4f, 0x48, 0x43, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x43, 0x50, 0x48, 0x4b, 0x4f, 0x43, 0x50, 0x4b, 0x43, 0x4f, 0x49, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x43, 0x50, 0x48, 0x4b, 0x4f, 0x43, 0x50, 0x4b, 0x43,
	0x4f, 0x49, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x4c, 0x4c, 0x4a, 0x4a, 0x50, 0x4f, 0x46, 0x4a, 0x4e, 0x4b, 0x4b, 0x18, 0x09,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x4c, 0x4a, 0x4a, 0x50, 0x4f, 0x46, 0x4a, 0x4e, 0x4b,
	0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x45, 0x46, 0x49, 0x43, 0x4e, 0x50, 0x44, 0x43, 0x4b, 0x50, 0x4b,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x45, 0x46, 0x49, 0x43, 0x4e, 0x50, 0x44, 0x43,
	0x4b, 0x50, 0x4b, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x4a, 0x49, 0x4f, 0x4f, 0x46, 0x47, 0x4b, 0x50,
	0x4a, 0x43, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x47, 0x4a, 0x49, 0x4f, 0x4f, 0x46,
	0x47, 0x4b, 0x50, 0x4a, 0x43, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_KMEDPBBAOHC_proto_rawDescOnce sync.Once
	file_KMEDPBBAOHC_proto_rawDescData = file_KMEDPBBAOHC_proto_rawDesc
)

func file_KMEDPBBAOHC_proto_rawDescGZIP() []byte {
	file_KMEDPBBAOHC_proto_rawDescOnce.Do(func() {
		file_KMEDPBBAOHC_proto_rawDescData = protoimpl.X.CompressGZIP(file_KMEDPBBAOHC_proto_rawDescData)
	})
	return file_KMEDPBBAOHC_proto_rawDescData
}

var file_KMEDPBBAOHC_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_KMEDPBBAOHC_proto_goTypes = []interface{}{
	(*KMEDPBBAOHC)(nil), // 0: KMEDPBBAOHC
}
var file_KMEDPBBAOHC_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_KMEDPBBAOHC_proto_init() }
func file_KMEDPBBAOHC_proto_init() {
	if File_KMEDPBBAOHC_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_KMEDPBBAOHC_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KMEDPBBAOHC); i {
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
			RawDescriptor: file_KMEDPBBAOHC_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_KMEDPBBAOHC_proto_goTypes,
		DependencyIndexes: file_KMEDPBBAOHC_proto_depIdxs,
		MessageInfos:      file_KMEDPBBAOHC_proto_msgTypes,
	}.Build()
	File_KMEDPBBAOHC_proto = out.File
	file_KMEDPBBAOHC_proto_rawDesc = nil
	file_KMEDPBBAOHC_proto_goTypes = nil
	file_KMEDPBBAOHC_proto_depIdxs = nil
}
