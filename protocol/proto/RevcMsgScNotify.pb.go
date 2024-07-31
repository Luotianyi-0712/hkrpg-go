// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RevcMsgScNotify.proto

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

type RevcMsgScNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageType MsgType      `protobuf:"varint,11,opt,name=message_type,json=messageType,proto3,enum=MsgType" json:"message_type,omitempty"`
	ENLMBCCJFBG *JDKPHOFLFEN `protobuf:"bytes,14,opt,name=ENLMBCCJFBG,proto3" json:"ENLMBCCJFBG,omitempty"`
	SourceUid   uint32       `protobuf:"varint,4,opt,name=source_uid,json=sourceUid,proto3" json:"source_uid,omitempty"`
	ChatType    ChatType     `protobuf:"varint,6,opt,name=chat_type,json=chatType,proto3,enum=ChatType" json:"chat_type,omitempty"`
	ExtraId     uint32       `protobuf:"varint,8,opt,name=extra_id,json=extraId,proto3" json:"extra_id,omitempty"`
	MessageText string       `protobuf:"bytes,10,opt,name=message_text,json=messageText,proto3" json:"message_text,omitempty"`
	TargetUid   uint32       `protobuf:"varint,7,opt,name=target_uid,json=targetUid,proto3" json:"target_uid,omitempty"`
}

func (x *RevcMsgScNotify) Reset() {
	*x = RevcMsgScNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RevcMsgScNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevcMsgScNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevcMsgScNotify) ProtoMessage() {}

func (x *RevcMsgScNotify) ProtoReflect() protoreflect.Message {
	mi := &file_RevcMsgScNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevcMsgScNotify.ProtoReflect.Descriptor instead.
func (*RevcMsgScNotify) Descriptor() ([]byte, []int) {
	return file_RevcMsgScNotify_proto_rawDescGZIP(), []int{0}
}

func (x *RevcMsgScNotify) GetMessageType() MsgType {
	if x != nil {
		return x.MessageType
	}
	return MsgType_MSG_TYPE_NONE
}

func (x *RevcMsgScNotify) GetENLMBCCJFBG() *JDKPHOFLFEN {
	if x != nil {
		return x.ENLMBCCJFBG
	}
	return nil
}

func (x *RevcMsgScNotify) GetSourceUid() uint32 {
	if x != nil {
		return x.SourceUid
	}
	return 0
}

func (x *RevcMsgScNotify) GetChatType() ChatType {
	if x != nil {
		return x.ChatType
	}
	return ChatType_CHAT_TYPE_NONE
}

func (x *RevcMsgScNotify) GetExtraId() uint32 {
	if x != nil {
		return x.ExtraId
	}
	return 0
}

func (x *RevcMsgScNotify) GetMessageText() string {
	if x != nil {
		return x.MessageText
	}
	return ""
}

func (x *RevcMsgScNotify) GetTargetUid() uint32 {
	if x != nil {
		return x.TargetUid
	}
	return 0
}

var File_RevcMsgScNotify_proto protoreflect.FileDescriptor

var file_RevcMsgScNotify_proto_rawDesc = []byte{
	0x0a, 0x15, 0x52, 0x65, 0x76, 0x63, 0x4d, 0x73, 0x67, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4a, 0x44, 0x4b, 0x50, 0x48, 0x4f, 0x46,
	0x4c, 0x46, 0x45, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x43, 0x68, 0x61, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x4d, 0x73, 0x67, 0x54,
	0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x02, 0x0a, 0x0f, 0x52, 0x65,
	0x76, 0x63, 0x4d, 0x73, 0x67, 0x53, 0x63, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2b, 0x0a,
	0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x08, 0x2e, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2e, 0x0a, 0x0b, 0x45, 0x4e,
	0x4c, 0x4d, 0x42, 0x43, 0x43, 0x4a, 0x46, 0x42, 0x47, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x4a, 0x44, 0x4b, 0x50, 0x48, 0x4f, 0x46, 0x4c, 0x46, 0x45, 0x4e, 0x52, 0x0b, 0x45,
	0x4e, 0x4c, 0x4d, 0x42, 0x43, 0x43, 0x4a, 0x46, 0x42, 0x47, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x09, 0x63, 0x68, 0x61,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x65, 0x78, 0x74, 0x72, 0x61, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x55, 0x69, 0x64, 0x42, 0x28,
	0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RevcMsgScNotify_proto_rawDescOnce sync.Once
	file_RevcMsgScNotify_proto_rawDescData = file_RevcMsgScNotify_proto_rawDesc
)

func file_RevcMsgScNotify_proto_rawDescGZIP() []byte {
	file_RevcMsgScNotify_proto_rawDescOnce.Do(func() {
		file_RevcMsgScNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_RevcMsgScNotify_proto_rawDescData)
	})
	return file_RevcMsgScNotify_proto_rawDescData
}

var file_RevcMsgScNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RevcMsgScNotify_proto_goTypes = []interface{}{
	(*RevcMsgScNotify)(nil), // 0: RevcMsgScNotify
	(MsgType)(0),            // 1: MsgType
	(*JDKPHOFLFEN)(nil),     // 2: JDKPHOFLFEN
	(ChatType)(0),           // 3: ChatType
}
var file_RevcMsgScNotify_proto_depIdxs = []int32{
	1, // 0: RevcMsgScNotify.message_type:type_name -> MsgType
	2, // 1: RevcMsgScNotify.ENLMBCCJFBG:type_name -> JDKPHOFLFEN
	3, // 2: RevcMsgScNotify.chat_type:type_name -> ChatType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_RevcMsgScNotify_proto_init() }
func file_RevcMsgScNotify_proto_init() {
	if File_RevcMsgScNotify_proto != nil {
		return
	}
	file_JDKPHOFLFEN_proto_init()
	file_ChatType_proto_init()
	file_MsgType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RevcMsgScNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevcMsgScNotify); i {
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
			RawDescriptor: file_RevcMsgScNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RevcMsgScNotify_proto_goTypes,
		DependencyIndexes: file_RevcMsgScNotify_proto_depIdxs,
		MessageInfos:      file_RevcMsgScNotify_proto_msgTypes,
	}.Build()
	File_RevcMsgScNotify_proto = out.File
	file_RevcMsgScNotify_proto_rawDesc = nil
	file_RevcMsgScNotify_proto_goTypes = nil
	file_RevcMsgScNotify_proto_depIdxs = nil
}
