// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: Challenge.proto

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

type Challenge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RecordId    uint32       `protobuf:"varint,4,opt,name=record_id,json=recordId,proto3" json:"record_id,omitempty"`
	PlayerInfo  *LMJLNMPCJJA `protobuf:"bytes,3,opt,name=player_info,json=playerInfo,proto3" json:"player_info,omitempty"`
	ScoreTwo    uint32       `protobuf:"varint,6,opt,name=score_two,json=scoreTwo,proto3" json:"score_two,omitempty"`
	Star        uint32       `protobuf:"varint,14,opt,name=star,proto3" json:"star,omitempty"`
	AHNJDLJONFO bool         `protobuf:"varint,13,opt,name=AHNJDLJONFO,proto3" json:"AHNJDLJONFO,omitempty"`
	ChallengeId uint32       `protobuf:"varint,5,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	ScoreId     uint32       `protobuf:"varint,1,opt,name=score_id,json=scoreId,proto3" json:"score_id,omitempty"`
	TakenReward uint32       `protobuf:"varint,9,opt,name=taken_reward,json=takenReward,proto3" json:"taken_reward,omitempty"`
}

func (x *Challenge) Reset() {
	*x = Challenge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Challenge_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Challenge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Challenge) ProtoMessage() {}

func (x *Challenge) ProtoReflect() protoreflect.Message {
	mi := &file_Challenge_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Challenge.ProtoReflect.Descriptor instead.
func (*Challenge) Descriptor() ([]byte, []int) {
	return file_Challenge_proto_rawDescGZIP(), []int{0}
}

func (x *Challenge) GetRecordId() uint32 {
	if x != nil {
		return x.RecordId
	}
	return 0
}

func (x *Challenge) GetPlayerInfo() *LMJLNMPCJJA {
	if x != nil {
		return x.PlayerInfo
	}
	return nil
}

func (x *Challenge) GetScoreTwo() uint32 {
	if x != nil {
		return x.ScoreTwo
	}
	return 0
}

func (x *Challenge) GetStar() uint32 {
	if x != nil {
		return x.Star
	}
	return 0
}

func (x *Challenge) GetAHNJDLJONFO() bool {
	if x != nil {
		return x.AHNJDLJONFO
	}
	return false
}

func (x *Challenge) GetChallengeId() uint32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

func (x *Challenge) GetScoreId() uint32 {
	if x != nil {
		return x.ScoreId
	}
	return 0
}

func (x *Challenge) GetTakenReward() uint32 {
	if x != nil {
		return x.TakenReward
	}
	return 0
}

var File_Challenge_proto protoreflect.FileDescriptor

var file_Challenge_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x4c, 0x4d, 0x4a, 0x4c, 0x4e, 0x4d, 0x50, 0x43, 0x4a, 0x4a, 0x41, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8b, 0x02, 0x0a, 0x09, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12,
	0x2d, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x4d, 0x4a, 0x4c, 0x4e, 0x4d, 0x50, 0x43, 0x4a,
	0x4a, 0x41, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x77, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x54, 0x77, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x74, 0x61, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x74, 0x61, 0x72, 0x12,
	0x20, 0x0a, 0x0b, 0x41, 0x48, 0x4e, 0x4a, 0x44, 0x4c, 0x4a, 0x4f, 0x4e, 0x46, 0x4f, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x41, 0x48, 0x4e, 0x4a, 0x44, 0x4c, 0x4a, 0x4f, 0x4e, 0x46,
	0x4f, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e,
	0x67, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x61, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44,
	0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Challenge_proto_rawDescOnce sync.Once
	file_Challenge_proto_rawDescData = file_Challenge_proto_rawDesc
)

func file_Challenge_proto_rawDescGZIP() []byte {
	file_Challenge_proto_rawDescOnce.Do(func() {
		file_Challenge_proto_rawDescData = protoimpl.X.CompressGZIP(file_Challenge_proto_rawDescData)
	})
	return file_Challenge_proto_rawDescData
}

var file_Challenge_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_Challenge_proto_goTypes = []interface{}{
	(*Challenge)(nil),   // 0: Challenge
	(*LMJLNMPCJJA)(nil), // 1: LMJLNMPCJJA
}
var file_Challenge_proto_depIdxs = []int32{
	1, // 0: Challenge.player_info:type_name -> LMJLNMPCJJA
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_Challenge_proto_init() }
func file_Challenge_proto_init() {
	if File_Challenge_proto != nil {
		return
	}
	file_LMJLNMPCJJA_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Challenge_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Challenge); i {
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
			RawDescriptor: file_Challenge_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Challenge_proto_goTypes,
		DependencyIndexes: file_Challenge_proto_depIdxs,
		MessageInfos:      file_Challenge_proto_msgTypes,
	}.Build()
	File_Challenge_proto = out.File
	file_Challenge_proto_rawDesc = nil
	file_Challenge_proto_goTypes = nil
	file_Challenge_proto_depIdxs = nil
}
