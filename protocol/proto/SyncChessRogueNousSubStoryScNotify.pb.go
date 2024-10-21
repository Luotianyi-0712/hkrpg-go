// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SyncChessRogueNousSubStoryScNotify.proto

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

type SyncChessRogueNousSubStoryScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChessRogueSubStoryId uint32 `protobuf:"varint,4,opt,name=chess_rogue_sub_story_id,json=chessRogueSubStoryId,proto3" json:"chess_rogue_sub_story_id,omitempty"`
}

func (x *SyncChessRogueNousSubStoryScNotify) Reset() {
	*x = SyncChessRogueNousSubStoryScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SyncChessRogueNousSubStoryScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncChessRogueNousSubStoryScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncChessRogueNousSubStoryScNotify) ProtoMessage() {}

func (x *SyncChessRogueNousSubStoryScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_SyncChessRogueNousSubStoryScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncChessRogueNousSubStoryScNotify.ProtoReflect.Descriptor instead.
func (*SyncChessRogueNousSubStoryScNotify) Descriptor() ([]byte, []int) {
	return file_SyncChessRogueNousSubStoryScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *SyncChessRogueNousSubStoryScNotify) GetChessRogueSubStoryId() uint32 {
	if x != nil {
		return x.ChessRogueSubStoryId
	}
	return 0
}

var File_SyncChessRogueNousSubStoryScNotify_proto protoreflect.FileDescriptor

var file_SyncChessRogueNousSubStoryScNotify_proto_rawDesc = []byte{
	0x0a, 0x28, 0x53, 0x79, 0x6e, 0x63, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65,
	0x4e, 0x6f, 0x75, 0x73, 0x53, 0x75, 0x62, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x63, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x22, 0x53, 0x79,
	0x6e, 0x63, 0x43, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x4e, 0x6f, 0x75, 0x73,
	0x53, 0x75, 0x62, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x12, 0x36, 0x0a, 0x18, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x5f,
	0x73, 0x75, 0x62, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x14, 0x63, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x53, 0x75,
	0x62, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SyncChessRogueNousSubStoryScNotify_proto_rawDescOnce sync.Once
	file_SyncChessRogueNousSubStoryScNotify_proto_rawDescData = file_SyncChessRogueNousSubStoryScNotify_proto_rawDesc
)

func file_SyncChessRogueNousSubStoryScNotify_proto_rawDescGZIP() []byte {
	file_SyncChessRogueNousSubStoryScNotify_proto_rawDescOnce.Do(func() {
		file_SyncChessRogueNousSubStoryScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_SyncChessRogueNousSubStoryScNotify_proto_rawDescData)
	})
	return file_SyncChessRogueNousSubStoryScNotify_proto_rawDescData
}

var file_SyncChessRogueNousSubStoryScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SyncChessRogueNousSubStoryScNotify_proto_goTypes = []interface{}{
	(*SyncChessRogueNousSubStoryScNotify)(nil), // 0: SyncChessRogueNousSubStoryScNotify
}
var file_SyncChessRogueNousSubStoryScNotify_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SyncChessRogueNousSubStoryScNotify_proto_init() }
func file_SyncChessRogueNousSubStoryScNotify_proto_init() {
	if File_SyncChessRogueNousSubStoryScNotify_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SyncChessRogueNousSubStoryScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncChessRogueNousSubStoryScNotify); i {
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
			RawDescriptor: file_SyncChessRogueNousSubStoryScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SyncChessRogueNousSubStoryScNotify_proto_goTypes,
		DependencyIndexes: file_SyncChessRogueNousSubStoryScNotify_proto_depIdxs,
		MessageInfos:      file_SyncChessRogueNousSubStoryScNotify_proto_msgTypes,
	}.Build()
	File_SyncChessRogueNousSubStoryScNotify_proto = out.File
	file_SyncChessRogueNousSubStoryScNotify_proto_rawDesc = nil
	file_SyncChessRogueNousSubStoryScNotify_proto_goTypes = nil
	file_SyncChessRogueNousSubStoryScNotify_proto_depIdxs = nil
}
