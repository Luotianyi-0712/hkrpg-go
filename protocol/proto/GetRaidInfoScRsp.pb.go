// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetRaidInfoScRsp.proto

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

type GetRaidInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeTakenRewardIdList []uint32            `protobuf:"varint,4,rep,packed,name=challenge_taken_reward_id_list,json=challengeTakenRewardIdList,proto3" json:"challenge_taken_reward_id_list,omitempty"`
	FinishedRaidInfoList       []*FinishedRaidInfo `protobuf:"bytes,12,rep,name=finished_raid_info_list,json=finishedRaidInfoList,proto3" json:"finished_raid_info_list,omitempty"`
	ChallengeRaidList          []*ChallengeRaid    `protobuf:"bytes,9,rep,name=challenge_raid_list,json=challengeRaidList,proto3" json:"challenge_raid_list,omitempty"`
	Retcode                    uint32              `protobuf:"varint,1,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GetRaidInfoScRsp) Reset() {
	*x = GetRaidInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetRaidInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRaidInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRaidInfoScRsp) ProtoMessage() {}

func (x *GetRaidInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetRaidInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRaidInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetRaidInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetRaidInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetRaidInfoScRsp) GetChallengeTakenRewardIdList() []uint32 {
	if x != nil {
		return x.ChallengeTakenRewardIdList
	}
	return nil
}

func (x *GetRaidInfoScRsp) GetFinishedRaidInfoList() []*FinishedRaidInfo {
	if x != nil {
		return x.FinishedRaidInfoList
	}
	return nil
}

func (x *GetRaidInfoScRsp) GetChallengeRaidList() []*ChallengeRaid {
	if x != nil {
		return x.ChallengeRaidList
	}
	return nil
}

func (x *GetRaidInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_GetRaidInfoScRsp_proto protoreflect.FileDescriptor

var file_GetRaidInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x52, 0x61, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52,
	0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x65, 0x64, 0x52, 0x61, 0x69, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x13, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x52, 0x61, 0x69, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x61, 0x69,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x42, 0x0a, 0x1e, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x5f, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x1a, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x54, 0x61, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x49, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x48,
	0x0a, 0x17, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x72, 0x61, 0x69, 0x64, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x61, 0x69, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x14, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x52, 0x61, 0x69, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x13, 0x63, 0x68, 0x61, 0x6c,
	0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x72, 0x61, 0x69, 0x64, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x52, 0x61, 0x69, 0x64, 0x52, 0x11, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x52, 0x61, 0x69, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f,
	0x64, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GetRaidInfoScRsp_proto_rawDescOnce sync.Once
	file_GetRaidInfoScRsp_proto_rawDescData = file_GetRaidInfoScRsp_proto_rawDesc
)

func file_GetRaidInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetRaidInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetRaidInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetRaidInfoScRsp_proto_rawDescData)
	})
	return file_GetRaidInfoScRsp_proto_rawDescData
}

var file_GetRaidInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetRaidInfoScRsp_proto_goTypes = []interface{}{
	(*GetRaidInfoScRsp)(nil), // 0: GetRaidInfoScRsp
	(*FinishedRaidInfo)(nil), // 1: FinishedRaidInfo
	(*ChallengeRaid)(nil),    // 2: ChallengeRaid
}
var file_GetRaidInfoScRsp_proto_depIdxs = []int32{
	1, // 0: GetRaidInfoScRsp.finished_raid_info_list:type_name -> FinishedRaidInfo
	2, // 1: GetRaidInfoScRsp.challenge_raid_list:type_name -> ChallengeRaid
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_GetRaidInfoScRsp_proto_init() }
func file_GetRaidInfoScRsp_proto_init() {
	if File_GetRaidInfoScRsp_proto != nil {
		return
	}
	file_FinishedRaidInfo_proto_init()
	file_ChallengeRaid_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetRaidInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRaidInfoScRsp); i {
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
			RawDescriptor: file_GetRaidInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetRaidInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetRaidInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetRaidInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetRaidInfoScRsp_proto = out.File
	file_GetRaidInfoScRsp_proto_rawDesc = nil
	file_GetRaidInfoScRsp_proto_goTypes = nil
	file_GetRaidInfoScRsp_proto_depIdxs = nil
}
