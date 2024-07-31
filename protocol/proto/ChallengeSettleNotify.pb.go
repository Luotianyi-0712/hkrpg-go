// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ChallengeSettleNotify.proto

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

type ChallengeSettleNotify struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChallengeScore uint32       `protobuf:"varint,10,opt,name=challenge_score,json=challengeScore,proto3" json:"challenge_score,omitempty"`
	OKGCOBHLIIM    uint32       `protobuf:"varint,3,opt,name=OKGCOBHLIIM,proto3" json:"OKGCOBHLIIM,omitempty"`
	ScoreTwo       uint32       `protobuf:"varint,1,opt,name=score_two,json=scoreTwo,proto3" json:"score_two,omitempty"`
	IsWin          bool         `protobuf:"varint,5,opt,name=is_win,json=isWin,proto3" json:"is_win,omitempty"`
	JCDOIJEJKDH    []uint32     `protobuf:"varint,6,rep,packed,name=JCDOIJEJKDH,proto3" json:"JCDOIJEJKDH,omitempty"`
	LHDFJGBLFNH    *NMHNANJAINM `protobuf:"bytes,13,opt,name=LHDFJGBLFNH,proto3" json:"LHDFJGBLFNH,omitempty"`
	ChallengeId    uint32       `protobuf:"varint,14,opt,name=challenge_id,json=challengeId,proto3" json:"challenge_id,omitempty"`
	Star           uint32       `protobuf:"varint,8,opt,name=star,proto3" json:"star,omitempty"`
	Reward         *ItemList    `protobuf:"bytes,9,opt,name=reward,proto3" json:"reward,omitempty"`
}

func (x *ChallengeSettleNotify) Reset() {
	*x = ChallengeSettleNotify{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ChallengeSettleNotify_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChallengeSettleNotify) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChallengeSettleNotify) ProtoMessage() {}

func (x *ChallengeSettleNotify) ProtoReflect() protoreflect.Message {
	mi := &file_ChallengeSettleNotify_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChallengeSettleNotify.ProtoReflect.Descriptor instead.
func (*ChallengeSettleNotify) Descriptor() ([]byte, []int) {
	return file_ChallengeSettleNotify_proto_rawDescGZIP(), []int{0}
}

func (x *ChallengeSettleNotify) GetChallengeScore() uint32 {
	if x != nil {
		return x.ChallengeScore
	}
	return 0
}

func (x *ChallengeSettleNotify) GetOKGCOBHLIIM() uint32 {
	if x != nil {
		return x.OKGCOBHLIIM
	}
	return 0
}

func (x *ChallengeSettleNotify) GetScoreTwo() uint32 {
	if x != nil {
		return x.ScoreTwo
	}
	return 0
}

func (x *ChallengeSettleNotify) GetIsWin() bool {
	if x != nil {
		return x.IsWin
	}
	return false
}

func (x *ChallengeSettleNotify) GetJCDOIJEJKDH() []uint32 {
	if x != nil {
		return x.JCDOIJEJKDH
	}
	return nil
}

func (x *ChallengeSettleNotify) GetLHDFJGBLFNH() *NMHNANJAINM {
	if x != nil {
		return x.LHDFJGBLFNH
	}
	return nil
}

func (x *ChallengeSettleNotify) GetChallengeId() uint32 {
	if x != nil {
		return x.ChallengeId
	}
	return 0
}

func (x *ChallengeSettleNotify) GetStar() uint32 {
	if x != nil {
		return x.Star
	}
	return 0
}

func (x *ChallengeSettleNotify) GetReward() *ItemList {
	if x != nil {
		return x.Reward
	}
	return nil
}

var File_ChallengeSettleNotify_proto protoreflect.FileDescriptor

var file_ChallengeSettleNotify_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74, 0x6c,
	0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4e,
	0x4d, 0x48, 0x4e, 0x41, 0x4e, 0x4a, 0x41, 0x49, 0x4e, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x0e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc2, 0x02, 0x0a, 0x15, 0x43, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x65,
	0x74, 0x74, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x68,
	0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0e, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x53, 0x63,
	0x6f, 0x72, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x4b, 0x47, 0x43, 0x4f, 0x42, 0x48, 0x4c, 0x49,
	0x49, 0x4d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x4b, 0x47, 0x43, 0x4f, 0x42,
	0x48, 0x4c, 0x49, 0x49, 0x4d, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x74,
	0x77, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x54,
	0x77, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x77, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x57, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x43, 0x44,
	0x4f, 0x49, 0x4a, 0x45, 0x4a, 0x4b, 0x44, 0x48, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b,
	0x4a, 0x43, 0x44, 0x4f, 0x49, 0x4a, 0x45, 0x4a, 0x4b, 0x44, 0x48, 0x12, 0x2e, 0x0a, 0x0b, 0x4c,
	0x48, 0x44, 0x46, 0x4a, 0x47, 0x42, 0x4c, 0x46, 0x4e, 0x48, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x4e, 0x4d, 0x48, 0x4e, 0x41, 0x4e, 0x4a, 0x41, 0x49, 0x4e, 0x4d, 0x52, 0x0b,
	0x4c, 0x48, 0x44, 0x46, 0x4a, 0x47, 0x42, 0x4c, 0x46, 0x4e, 0x48, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x74, 0x61, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x74,
	0x61, 0x72, 0x12, 0x21, 0x0a, 0x06, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x09, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x06, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ChallengeSettleNotify_proto_rawDescOnce sync.Once
	file_ChallengeSettleNotify_proto_rawDescData = file_ChallengeSettleNotify_proto_rawDesc
)

func file_ChallengeSettleNotify_proto_rawDescGZIP() []byte {
	file_ChallengeSettleNotify_proto_rawDescOnce.Do(func() {
		file_ChallengeSettleNotify_proto_rawDescData = protoimpl.X.CompressGZIP(file_ChallengeSettleNotify_proto_rawDescData)
	})
	return file_ChallengeSettleNotify_proto_rawDescData
}

var file_ChallengeSettleNotify_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ChallengeSettleNotify_proto_goTypes = []interface{}{
	(*ChallengeSettleNotify)(nil), // 0: ChallengeSettleNotify
	(*NMHNANJAINM)(nil),           // 1: NMHNANJAINM
	(*ItemList)(nil),              // 2: ItemList
}
var file_ChallengeSettleNotify_proto_depIdxs = []int32{
	1, // 0: ChallengeSettleNotify.LHDFJGBLFNH:type_name -> NMHNANJAINM
	2, // 1: ChallengeSettleNotify.reward:type_name -> ItemList
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ChallengeSettleNotify_proto_init() }
func file_ChallengeSettleNotify_proto_init() {
	if File_ChallengeSettleNotify_proto != nil {
		return
	}
	file_NMHNANJAINM_proto_init()
	file_ItemList_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ChallengeSettleNotify_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChallengeSettleNotify); i {
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
			RawDescriptor: file_ChallengeSettleNotify_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ChallengeSettleNotify_proto_goTypes,
		DependencyIndexes: file_ChallengeSettleNotify_proto_depIdxs,
		MessageInfos:      file_ChallengeSettleNotify_proto_msgTypes,
	}.Build()
	File_ChallengeSettleNotify_proto = out.File
	file_ChallengeSettleNotify_proto_rawDesc = nil
	file_ChallengeSettleNotify_proto_goTypes = nil
	file_ChallengeSettleNotify_proto_depIdxs = nil
}
