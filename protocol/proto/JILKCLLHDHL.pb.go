// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: JILKCLLHDHL.proto

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

type JILKCLLHDHL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LDCNDJDFJCE uint32 `protobuf:"varint,3,opt,name=LDCNDJDFJCE,proto3" json:"LDCNDJDFJCE,omitempty"`
}

func (x *JILKCLLHDHL) Reset() {
	*x = JILKCLLHDHL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_JILKCLLHDHL_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JILKCLLHDHL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JILKCLLHDHL) ProtoMessage() {}

func (x *JILKCLLHDHL) ProtoReflect() protoreflect.Message {
	mi := &file_JILKCLLHDHL_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JILKCLLHDHL.ProtoReflect.Descriptor instead.
func (*JILKCLLHDHL) Descriptor() ([]byte, []int) {
	return file_JILKCLLHDHL_proto_rawDescGZIP(), []int{0}
}

func (x *JILKCLLHDHL) GetLDCNDJDFJCE() uint32 {
	if x != nil {
		return x.LDCNDJDFJCE
	}
	return 0
}

var File_JILKCLLHDHL_proto protoreflect.FileDescriptor

var file_JILKCLLHDHL_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4a, 0x49, 0x4c, 0x4b, 0x43, 0x4c, 0x4c, 0x48, 0x44, 0x48, 0x4c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2f, 0x0a, 0x0b, 0x4a, 0x49, 0x4c, 0x4b, 0x43, 0x4c, 0x4c, 0x48, 0x44,
	0x48, 0x4c, 0x12, 0x20, 0x0a, 0x0b, 0x4c, 0x44, 0x43, 0x4e, 0x44, 0x4a, 0x44, 0x46, 0x4a, 0x43,
	0x45, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4c, 0x44, 0x43, 0x4e, 0x44, 0x4a, 0x44,
	0x46, 0x4a, 0x43, 0x45, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_JILKCLLHDHL_proto_rawDescOnce sync.Once
	file_JILKCLLHDHL_proto_rawDescData = file_JILKCLLHDHL_proto_rawDesc
)

func file_JILKCLLHDHL_proto_rawDescGZIP() []byte {
	file_JILKCLLHDHL_proto_rawDescOnce.Do(func() {
		file_JILKCLLHDHL_proto_rawDescData = protoimpl.X.CompressGZIP(file_JILKCLLHDHL_proto_rawDescData)
	})
	return file_JILKCLLHDHL_proto_rawDescData
}

var file_JILKCLLHDHL_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_JILKCLLHDHL_proto_goTypes = []interface{}{
	(*JILKCLLHDHL)(nil), // 0: JILKCLLHDHL
}
var file_JILKCLLHDHL_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_JILKCLLHDHL_proto_init() }
func file_JILKCLLHDHL_proto_init() {
	if File_JILKCLLHDHL_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_JILKCLLHDHL_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JILKCLLHDHL); i {
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
			RawDescriptor: file_JILKCLLHDHL_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_JILKCLLHDHL_proto_goTypes,
		DependencyIndexes: file_JILKCLLHDHL_proto_depIdxs,
		MessageInfos:      file_JILKCLLHDHL_proto_msgTypes,
	}.Build()
	File_JILKCLLHDHL_proto = out.File
	file_JILKCLLHDHL_proto_rawDesc = nil
	file_JILKCLLHDHL_proto_goTypes = nil
	file_JILKCLLHDHL_proto_depIdxs = nil
}
