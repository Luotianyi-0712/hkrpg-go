// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: BatchGetQuestDataScRsp.proto

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

type BatchGetQuestDataScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuestList []*Quest `protobuf:"bytes,6,rep,name=quest_list,json=questList,proto3" json:"quest_list,omitempty"`
	Retcode   uint32   `protobuf:"varint,7,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *BatchGetQuestDataScRsp) Reset() {
	*x = BatchGetQuestDataScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BatchGetQuestDataScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetQuestDataScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetQuestDataScRsp) ProtoMessage() {}

func (x *BatchGetQuestDataScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_BatchGetQuestDataScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetQuestDataScRsp.ProtoReflect.Descriptor instead.
func (*BatchGetQuestDataScRsp) Descriptor() ([]byte, []int) {
	return file_BatchGetQuestDataScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *BatchGetQuestDataScRsp) GetQuestList() []*Quest {
	if x != nil {
		return x.QuestList
	}
	return nil
}

func (x *BatchGetQuestDataScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_BatchGetQuestDataScRsp_proto protoreflect.FileDescriptor

var file_BatchGetQuestDataScRsp_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b,
	0x51, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x16, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x51, 0x75, 0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x51, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x09, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72,
	0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BatchGetQuestDataScRsp_proto_rawDescOnce sync.Once
	file_BatchGetQuestDataScRsp_proto_rawDescData = file_BatchGetQuestDataScRsp_proto_rawDesc
)

func file_BatchGetQuestDataScRsp_proto_rawDescGZIP() []byte {
	file_BatchGetQuestDataScRsp_proto_rawDescOnce.Do(func() {
		file_BatchGetQuestDataScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_BatchGetQuestDataScRsp_proto_rawDescData)
	})
	return file_BatchGetQuestDataScRsp_proto_rawDescData
}

var file_BatchGetQuestDataScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BatchGetQuestDataScRsp_proto_goTypes = []interface{}{
	(*BatchGetQuestDataScRsp)(nil), // 0: BatchGetQuestDataScRsp
	(*Quest)(nil),                  // 1: Quest
}
var file_BatchGetQuestDataScRsp_proto_depIdxs = []int32{
	1, // 0: BatchGetQuestDataScRsp.quest_list:type_name -> Quest
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BatchGetQuestDataScRsp_proto_init() }
func file_BatchGetQuestDataScRsp_proto_init() {
	if File_BatchGetQuestDataScRsp_proto != nil {
		return
	}
	file_Quest_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BatchGetQuestDataScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetQuestDataScRsp); i {
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
			RawDescriptor: file_BatchGetQuestDataScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BatchGetQuestDataScRsp_proto_goTypes,
		DependencyIndexes: file_BatchGetQuestDataScRsp_proto_depIdxs,
		MessageInfos:      file_BatchGetQuestDataScRsp_proto_msgTypes,
	}.Build()
	File_BatchGetQuestDataScRsp_proto = out.File
	file_BatchGetQuestDataScRsp_proto_rawDesc = nil
	file_BatchGetQuestDataScRsp_proto_goTypes = nil
	file_BatchGetQuestDataScRsp_proto_depIdxs = nil
}
