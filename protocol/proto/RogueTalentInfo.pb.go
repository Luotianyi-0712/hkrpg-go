// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: RogueTalentInfo.proto

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

type RogueTalentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TalentId                uint32                 `protobuf:"varint,12,opt,name=talent_id,json=talentId,proto3" json:"talent_id,omitempty"`
	Status                  RogueTalentStatus      `protobuf:"varint,7,opt,name=status,proto3,enum=RogueTalentStatus" json:"status,omitempty"`
	RogueUnlockProgressList []*RogueUnlockProgress `protobuf:"bytes,5,rep,name=rogue_unlock_progress_list,json=rogueUnlockProgressList,proto3" json:"rogue_unlock_progress_list,omitempty"`
}

func (x *RogueTalentInfo) Reset() {
	*x = RogueTalentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_RogueTalentInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RogueTalentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RogueTalentInfo) ProtoMessage() {}

func (x *RogueTalentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_RogueTalentInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RogueTalentInfo.ProtoReflect.Descriptor instead.
func (*RogueTalentInfo) Descriptor() ([]byte, []int) {
	return file_RogueTalentInfo_proto_rawDescGZIP(), []int{0}
}

func (x *RogueTalentInfo) GetTalentId() uint32 {
	if x != nil {
		return x.TalentId
	}
	return 0
}

func (x *RogueTalentInfo) GetStatus() RogueTalentStatus {
	if x != nil {
		return x.Status
	}
	return RogueTalentStatus_ROGUE_TALENT_STATUS_LOCK
}

func (x *RogueTalentInfo) GetRogueUnlockProgressList() []*RogueUnlockProgress {
	if x != nil {
		return x.RogueUnlockProgressList
	}
	return nil
}

var File_RogueTalentInfo_proto protoreflect.FileDescriptor

var file_RogueTalentInfo_proto_rawDesc = []byte{
	0x0a, 0x15, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x61,
	0x6c, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x52, 0x6f, 0x67, 0x75, 0x65, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x0f,
	0x52, 0x6f, 0x67, 0x75, 0x65, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x54, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x51, 0x0a, 0x1a, 0x72, 0x6f, 0x67, 0x75,
	0x65, 0x5f, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x17, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x28, 0x5a, 0x08, 0x2e,
	0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e,
	0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_RogueTalentInfo_proto_rawDescOnce sync.Once
	file_RogueTalentInfo_proto_rawDescData = file_RogueTalentInfo_proto_rawDesc
)

func file_RogueTalentInfo_proto_rawDescGZIP() []byte {
	file_RogueTalentInfo_proto_rawDescOnce.Do(func() {
		file_RogueTalentInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_RogueTalentInfo_proto_rawDescData)
	})
	return file_RogueTalentInfo_proto_rawDescData
}

var file_RogueTalentInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_RogueTalentInfo_proto_goTypes = []interface{}{
	(*RogueTalentInfo)(nil),     // 0: RogueTalentInfo
	(RogueTalentStatus)(0),      // 1: RogueTalentStatus
	(*RogueUnlockProgress)(nil), // 2: RogueUnlockProgress
}
var file_RogueTalentInfo_proto_depIdxs = []int32{
	1, // 0: RogueTalentInfo.status:type_name -> RogueTalentStatus
	2, // 1: RogueTalentInfo.rogue_unlock_progress_list:type_name -> RogueUnlockProgress
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_RogueTalentInfo_proto_init() }
func file_RogueTalentInfo_proto_init() {
	if File_RogueTalentInfo_proto != nil {
		return
	}
	file_RogueTalentStatus_proto_init()
	file_RogueUnlockProgress_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_RogueTalentInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RogueTalentInfo); i {
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
			RawDescriptor: file_RogueTalentInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_RogueTalentInfo_proto_goTypes,
		DependencyIndexes: file_RogueTalentInfo_proto_depIdxs,
		MessageInfos:      file_RogueTalentInfo_proto_msgTypes,
	}.Build()
	File_RogueTalentInfo_proto = out.File
	file_RogueTalentInfo_proto_rawDesc = nil
	file_RogueTalentInfo_proto_goTypes = nil
	file_RogueTalentInfo_proto_depIdxs = nil
}
