// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChallengeBossPhaseSettleNotify.proto

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

type ChallengeBossPhaseSettleNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeId       uint32          `protobuf:"varint,10,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	Phase             uint32          `protobuf:"varint,3,opt,name=phase,proto3" json:"phase,omitempty"`
	ScoreTwo          uint32          `protobuf:"varint,14,opt,name=score_two,json=scoreTwo,proto3" json:"score_two,omitempty"`
	IsRemainingAction bool            `protobuf:"varint,2,opt,name=is_remaining_action,json=isRemainingAction,proto3" json:"is_remaining_action,omitempty"`
	IsWin             bool            `protobuf:"varint,1,opt,name=is_win,json=isWin,proto3" json:"is_win,omitempty"`
	BattleTargetList  []*BattleTarget `protobuf:"bytes,7,rep,name=battle_target_list,json=battleTargetList,proto3" json:"battle_target_list,omitempty"`
	Star              uint32          `protobuf:"varint,9,opt,name=star,proto3" json:"star,omitempty"`
	ChallengeScore    uint32          `protobuf:"varint,6,opt,name=challenge_score,json=challengeScore,proto3" json:"challenge_score,omitempty"`
	CKDHFNAFNBN       uint32          `protobuf:"varint,8,opt,name=CKDHFNAFNBN,proto3" json:"CKDHFNAFNBN,omitempty"`
	IsReward          bool            `protobuf:"varint,15,opt,name=is_reward,json=isReward,proto3" json:"is_reward,omitempty"`
}

func (x *ChallengeBossPhaseSettleNotify) Reset() {
	*x = ChallengeBossPhaseSettleNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChallengeBossPhaseSettleNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeBossPhaseSettleNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeBossPhaseSettleNotify) ProtoMessage() {}

func (x *ChallengeBossPhaseSettleNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ChallengeBossPhaseSettleNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeBossPhaseSettleNotify.ProtoReflect.Descriptor instead.
func (*ChallengeBossPhaseSettleNotify) Descriptor() ([]byte, []int) {
	return file_ChallengeBossPhaseSettleNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ChallengeBossPhaseSettleNotify) GetChallengeId() uint32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

func (x *ChallengeBossPhaseSettleNotify) GetPhase() uint32 {
	if x != nil {
		return x.Phase
	}
	return 0
}

func (x *ChallengeBossPhaseSettleNotify) GetScoreTwo() uint32 {
	if x != nil {
		return x.ScoreTwo
	}
	return 0
}

func (x *ChallengeBossPhaseSettleNotify) GetIsRemainingAction() bool {
	if x != nil {
		return x.IsRemainingAction
	}
	return false
}

func (x *ChallengeBossPhaseSettleNotify) GetIsWin() bool {
	if x != nil {
		return x.IsWin
	}
	return false
}

func (x *ChallengeBossPhaseSettleNotify) GetBattleTargetList() []*BattleTarget {
	if x != nil {
		return x.BattleTargetList
	}
	return nil
}

func (x *ChallengeBossPhaseSettleNotify) GetStar() uint32 {
	if x != nil {
		return x.Star
	}
	return 0
}

func (x *ChallengeBossPhaseSettleNotify) GetChallengeScore() uint32 {
	if x != nil {
		return x.ChallengeScore
	}
	return 0
}

func (x *ChallengeBossPhaseSettleNotify) GetCKDHFNAFNBN() uint32 {
	if x != nil {
		return x.CKDHFNAFNBN
	}
	return 0
}

func (x *ChallengeBossPhaseSettleNotify) GetIsReward() bool {
	if x != nil {
		return x.IsReward
	}
	return false
}

var File_ChallengeBossPhaseSettleNotify_proto protoreflect.FileDescriptor

var file_ChallengeBossPhaseSettleNotify_proto_rawDesc = []byte{
	0x0a, 0x24, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x6f, 0x73, 0x73, 0x50,
	0x68, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf6, 0x02, 0x0a, 0x1e, 0x43,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x42, 0x6f, 0x73, 0x73, 0x50, 0x68, 0x61, 0x73,
	0x65, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x70, 0x68, 0x61, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x74, 0x77, 0x6f, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x63, 0x6f, 0x72, 0x65,
	0x54, 0x77, 0x6f, 0x12, 0x2e, 0x0a, 0x13, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x6d, 0x61, 0x69, 0x6e,
	0x69, 0x6e, 0x67, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x69, 0x73, 0x52, 0x65, 0x6d, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x77, 0x69, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x57, 0x69, 0x6e, 0x12, 0x3b, 0x0a, 0x12, 0x62, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x10, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x72, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x74, 0x61, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53,
	0x63, 0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x4b, 0x44, 0x48, 0x46, 0x4e, 0x41, 0x46,
	0x4e, 0x42, 0x4e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x43, 0x4b, 0x44, 0x48, 0x46,
	0x4e, 0x41, 0x46, 0x4e, 0x42, 0x4e, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e,
	0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChallengeBossPhaseSettleNotify_proto_rawDescOnce sync.Once
	file_ChallengeBossPhaseSettleNotify_proto_rawDescData = file_ChallengeBossPhaseSettleNotify_proto_rawDesc
)

func file_ChallengeBossPhaseSettleNotify_proto_rawDescGZIP() []byte {
	file_ChallengeBossPhaseSettleNotify_proto_rawDescOnce.Do(func() {
		file_ChallengeBossPhaseSettleNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChallengeBossPhaseSettleNotify_proto_rawDescData)
	})
	return file_ChallengeBossPhaseSettleNotify_proto_rawDescData
}

var file_ChallengeBossPhaseSettleNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChallengeBossPhaseSettleNotify_proto_goTypes = []interface{}{
	(*ChallengeBossPhaseSettleNotify)(nil), // 0: ChallengeBossPhaseSettleNotify
	(*BattleTarget)(nil),                   // 1: BattleTarget
}
var file_ChallengeBossPhaseSettleNotify_proto_depIdxs = []int32{
	1, // 0: ChallengeBossPhaseSettleNotify.battle_target_list:type_name -> BattleTarget
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ChallengeBossPhaseSettleNotify_proto_init() }
func file_ChallengeBossPhaseSettleNotify_proto_init() {
	if File_ChallengeBossPhaseSettleNotify_proto != nil {
		return
	}
	file_BattleTarget_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChallengeBossPhaseSettleNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeBossPhaseSettleNotify); i {
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
			RawDescriptor: file_ChallengeBossPhaseSettleNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChallengeBossPhaseSettleNotify_proto_goTypes,
		DependencyIndexes: file_ChallengeBossPhaseSettleNotify_proto_depIdxs,
		MessageInfos:      file_ChallengeBossPhaseSettleNotify_proto_msgTypes,
	}.Build()
	File_ChallengeBossPhaseSettleNotify_proto = out.File
	file_ChallengeBossPhaseSettleNotify_proto_rawDesc = nil
	file_ChallengeBossPhaseSettleNotify_proto_goTypes = nil
	file_ChallengeBossPhaseSettleNotify_proto_depIdxs = nil
}
