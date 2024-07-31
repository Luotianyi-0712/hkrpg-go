// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ShareScRsp.proto

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

type ShareScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reward      *ItemList  `protobuf:"bytes,1,opt,name=reward,proto3" json:"reward,omitempty"`
	NMGHCBDOLKI *ShareData `protobuf:"bytes,10,opt,name=NMGHCBDOLKI,proto3" json:"NMGHCBDOLKI,omitempty"`
	Retcode     uint32     `protobuf:"varint,8,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *ShareScRsp) Reset() {
	*x = ShareScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ShareScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShareScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShareScRsp) ProtoMessage() {}

func (x *ShareScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ShareScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShareScRsp.ProtoReflect.Descriptor instead.
func (*ShareScRsp) Descriptor() ([]byte, []int) {
	return file_ShareScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ShareScRsp) GetReward() *ItemList {
	if x != nil {
		return x.Reward
	}
	return nil
}

func (x *ShareScRsp) GetNMGHCBDOLKI() *ShareData {
	if x != nil {
		return x.NMGHCBDOLKI
	}
	return nil
}

func (x *ShareScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_ShareScRsp_proto protoreflect.FileDescriptor

var file_ShareScRsp_proto_rawDesc = []byte{
	0x0a, 0x10, 0x53, 0x68, 0x61, 0x72, 0x65, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0f, 0x53, 0x68, 0x61, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x0a, 0x53, 0x68, 0x61, 0x72, 0x65, 0x53, 0x63, 0x52, 0x73,
	0x70, 0x12, 0x21, 0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x06, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x12, 0x2c, 0x0a, 0x0b, 0x4e, 0x4d, 0x47, 0x48, 0x43, 0x42, 0x44, 0x4f,
	0x4c, 0x4b, 0x49, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x68, 0x61, 0x72,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x4e, 0x4d, 0x47, 0x48, 0x43, 0x42, 0x44, 0x4f, 0x4c,
	0x4b, 0x49, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x28, 0x5a, 0x08,
	0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69,
	0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ShareScRsp_proto_rawDescOnce sync.Once
	file_ShareScRsp_proto_rawDescData = file_ShareScRsp_proto_rawDesc
)

func file_ShareScRsp_proto_rawDescGZIP() []byte {
	file_ShareScRsp_proto_rawDescOnce.Do(func() {
		file_ShareScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ShareScRsp_proto_rawDescData)
	})
	return file_ShareScRsp_proto_rawDescData
}

var file_ShareScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ShareScRsp_proto_goTypes = []interface{}{
	(*ShareScRsp)(nil), // 0: ShareScRsp
	(*ItemList)(nil),   // 1: ItemList
	(*ShareData)(nil),  // 2: ShareData
}
var file_ShareScRsp_proto_depIdxs = []int32{
	1, // 0: ShareScRsp.reward:type_name -> ItemList
	2, // 1: ShareScRsp.NMGHCBDOLKI:type_name -> ShareData
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ShareScRsp_proto_init() }
func file_ShareScRsp_proto_init() {
	if File_ShareScRsp_proto != nil {
		return
	}
	file_ShareData_proto_init()
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ShareScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShareScRsp); i {
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
			RawDescriptor: file_ShareScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ShareScRsp_proto_goTypes,
		DependencyIndexes: file_ShareScRsp_proto_depIdxs,
		MessageInfos:      file_ShareScRsp_proto_msgTypes,
	}.Build()
	File_ShareScRsp_proto = out.File
	file_ShareScRsp_proto_rawDesc = nil
	file_ShareScRsp_proto_goTypes = nil
	file_ShareScRsp_proto_depIdxs = nil
}
