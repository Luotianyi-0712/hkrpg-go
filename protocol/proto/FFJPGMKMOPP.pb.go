// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FFJPGMKMOPP.proto

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

type FFJPGMKMOPP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemList []*ECHCJDNIHKO `protobuf:"bytes,4,rep,name=item_list,json=itemList,proto3" json:"item_list,omitempty"`
}

func (x *FFJPGMKMOPP) Reset() {
	*x = FFJPGMKMOPP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FFJPGMKMOPP_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FFJPGMKMOPP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FFJPGMKMOPP) ProtoMessage() {}

func (x *FFJPGMKMOPP) ProtoReflect() protoreflect.Message {
	mi := &file_FFJPGMKMOPP_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FFJPGMKMOPP.ProtoReflect.Descriptor instead.
func (*FFJPGMKMOPP) Descriptor() ([]byte, []int) {
	return file_FFJPGMKMOPP_proto_rawDescGZIP(), []int{0}
}

func (x *FFJPGMKMOPP) GetItemList() []*ECHCJDNIHKO {
	if x != nil {
		return x.ItemList
	}
	return nil
}

var File_FFJPGMKMOPP_proto protoreflect.FileDescriptor

var file_FFJPGMKMOPP_proto_rawDesc = []byte{
	0x0a, 0x11, 0x46, 0x46, 0x4a, 0x50, 0x47, 0x4d, 0x4b, 0x4d, 0x4f, 0x50, 0x50, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x45, 0x43, 0x48, 0x43, 0x4a, 0x44, 0x4e, 0x49, 0x48, 0x4b, 0x4f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x0b, 0x46, 0x46, 0x4a, 0x50, 0x47, 0x4d,
	0x4b, 0x4d, 0x4f, 0x50, 0x50, 0x12, 0x29, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x43, 0x48, 0x43, 0x4a,
	0x44, 0x4e, 0x49, 0x48, 0x4b, 0x4f, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FFJPGMKMOPP_proto_rawDescOnce sync.Once
	file_FFJPGMKMOPP_proto_rawDescData = file_FFJPGMKMOPP_proto_rawDesc
)

func file_FFJPGMKMOPP_proto_rawDescGZIP() []byte {
	file_FFJPGMKMOPP_proto_rawDescOnce.Do(func() {
		file_FFJPGMKMOPP_proto_rawDescData = protoimpl.X.CompressGZIP(file_FFJPGMKMOPP_proto_rawDescData)
	})
	return file_FFJPGMKMOPP_proto_rawDescData
}

var file_FFJPGMKMOPP_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FFJPGMKMOPP_proto_goTypes = []interface{}{
	(*FFJPGMKMOPP)(nil), // 0: FFJPGMKMOPP
	(*ECHCJDNIHKO)(nil), // 1: ECHCJDNIHKO
}
var file_FFJPGMKMOPP_proto_depIdxs = []int32{
	1, // 0: FFJPGMKMOPP.item_list:type_name -> ECHCJDNIHKO
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FFJPGMKMOPP_proto_init() }
func file_FFJPGMKMOPP_proto_init() {
	if File_FFJPGMKMOPP_proto != nil {
		return
	}
	file_ECHCJDNIHKO_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FFJPGMKMOPP_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FFJPGMKMOPP); i {
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
			RawDescriptor: file_FFJPGMKMOPP_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FFJPGMKMOPP_proto_goTypes,
		DependencyIndexes: file_FFJPGMKMOPP_proto_depIdxs,
		MessageInfos:      file_FFJPGMKMOPP_proto_msgTypes,
	}.Build()
	File_FFJPGMKMOPP_proto = out.File
	file_FFJPGMKMOPP_proto_rawDesc = nil
	file_FFJPGMKMOPP_proto_goTypes = nil
	file_FFJPGMKMOPP_proto_depIdxs = nil
}
