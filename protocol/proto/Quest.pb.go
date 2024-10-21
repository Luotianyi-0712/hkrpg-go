// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: Quest.proto

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

type Quest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FinishTime  int64       `protobuf:"varint,13,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
	Id          uint32      `protobuf:"varint,10,opt,name=id,proto3" json:"id,omitempty"`
	Status      QuestStatus `protobuf:"varint,15,opt,name=status,proto3,enum=QuestStatus" json:"status,omitempty"`
	Progress    uint32      `protobuf:"varint,6,opt,name=progress,proto3" json:"progress,omitempty"`
	ANAKGMFIDBJ []uint32    `protobuf:"varint,5,rep,packed,name=ANAKGMFIDBJ,proto3" json:"ANAKGMFIDBJ,omitempty"`
}

func (x *Quest) Reset() {
	*x = Quest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Quest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Quest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Quest) ProtoMessage() {}

func (x *Quest) ProtoReflect() protoreflect.Message {
	mi := &file_Quest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Quest.ProtoReflect.Descriptor instead.
func (*Quest) Descriptor() ([]byte, []int) {
	return file_Quest_proto_rawDescGZIP(), []int{0}
}

func (x *Quest) GetFinishTime() int64 {
	if x != nil {
		return x.FinishTime
	}
	return 0
}

func (x *Quest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Quest) GetStatus() QuestStatus {
	if x != nil {
		return x.Status
	}
	return QuestStatus_QUEST_NONE
}

func (x *Quest) GetProgress() uint32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *Quest) GetANAKGMFIDBJ() []uint32 {
	if x != nil {
		return x.ANAKGMFIDBJ
	}
	return nil
}

var File_Quest_proto protoreflect.FileDescriptor

var file_Quest_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x51, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9c, 0x01, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0c, 0x2e, 0x51, 0x75,
	0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x41, 0x4e, 0x41, 0x4b, 0x47, 0x4d, 0x46, 0x49, 0x44, 0x42, 0x4a, 0x18, 0x05, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x0b, 0x41, 0x4e, 0x41, 0x4b, 0x47, 0x4d, 0x46, 0x49, 0x44, 0x42, 0x4a, 0x42,
	0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Quest_proto_rawDescOnce sync.Once
	file_Quest_proto_rawDescData = file_Quest_proto_rawDesc
)

func file_Quest_proto_rawDescGZIP() []byte {
	file_Quest_proto_rawDescOnce.Do(func() {
		file_Quest_proto_rawDescData = protoimpl.X.CompressGZIP(file_Quest_proto_rawDescData)
	})
	return file_Quest_proto_rawDescData
}

var file_Quest_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Quest_proto_goTypes = []interface{}{
	(*Quest)(nil),    // 0: Quest
	(QuestStatus)(0), // 1: QuestStatus
}
var file_Quest_proto_depIdxs = []int32{
	1, // 0: Quest.status:type_name -> QuestStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Quest_proto_init() }
func file_Quest_proto_init() {
	if File_Quest_proto != nil {
		return
	}
	file_QuestStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Quest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Quest); i {
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
			RawDescriptor: file_Quest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Quest_proto_goTypes,
		DependencyIndexes: file_Quest_proto_depIdxs,
		MessageInfos:      file_Quest_proto_msgTypes,
	}.Build()
	File_Quest_proto = out.File
	file_Quest_proto_rawDesc = nil
	file_Quest_proto_goTypes = nil
	file_Quest_proto_depIdxs = nil
}
