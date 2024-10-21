// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: BattleEventBattleInfo.proto

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

type BattleEventBattleInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BattleEventId uint32               `protobuf:"varint,1,opt,name=battle_event_id,json=battleEventId,proto3" json:"battle_event_id,omitempty"`
	Status        *BattleEventProperty `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	SkillInfo     []*LLNHNHHLCIP       `protobuf:"bytes,3,rep,name=skill_info,json=skillInfo,proto3" json:"skill_info,omitempty"`
}

func (x *BattleEventBattleInfo) Reset() {
	*x = BattleEventBattleInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BattleEventBattleInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattleEventBattleInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattleEventBattleInfo) ProtoMessage() {}

func (x *BattleEventBattleInfo) ProtoReflect() protoreflect.Message {
	mi := &file_BattleEventBattleInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattleEventBattleInfo.ProtoReflect.Descriptor instead.
func (*BattleEventBattleInfo) Descriptor() ([]byte, []int) {
	return file_BattleEventBattleInfo_proto_rawDescGZIP(), []int{0}
}

func (x *BattleEventBattleInfo) GetBattleEventId() uint32 {
	if x != nil {
		return x.BattleEventId
	}
	return 0
}

func (x *BattleEventBattleInfo) GetStatus() *BattleEventProperty {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *BattleEventBattleInfo) GetSkillInfo() []*LLNHNHHLCIP {
	if x != nil {
		return x.SkillInfo
	}
	return nil
}

var File_BattleEventBattleInfo_proto protoreflect.FileDescriptor

var file_BattleEventBattleInfo_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x74,
	0x74, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x42,
	0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x4c, 0x4e, 0x48, 0x4e, 0x48,
	0x48, 0x4c, 0x43, 0x49, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x15,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x61, 0x74, 0x74, 0x6c,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x0f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65,
	0x72, 0x74, 0x79, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x0a, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x4c, 0x4c, 0x4e, 0x48, 0x4e, 0x48, 0x48, 0x4c, 0x43, 0x49, 0x50, 0x52, 0x09, 0x73,
	0x6b, 0x69, 0x6c, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67,
	0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BattleEventBattleInfo_proto_rawDescOnce sync.Once
	file_BattleEventBattleInfo_proto_rawDescData = file_BattleEventBattleInfo_proto_rawDesc
)

func file_BattleEventBattleInfo_proto_rawDescGZIP() []byte {
	file_BattleEventBattleInfo_proto_rawDescOnce.Do(func() {
		file_BattleEventBattleInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_BattleEventBattleInfo_proto_rawDescData)
	})
	return file_BattleEventBattleInfo_proto_rawDescData
}

var file_BattleEventBattleInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_BattleEventBattleInfo_proto_goTypes = []interface{}{
	(*BattleEventBattleInfo)(nil), // 0: BattleEventBattleInfo
	(*BattleEventProperty)(nil),   // 1: BattleEventProperty
	(*LLNHNHHLCIP)(nil),           // 2: LLNHNHHLCIP
}
var file_BattleEventBattleInfo_proto_depIdxs = []int32{
	1, // 0: BattleEventBattleInfo.status:type_name -> BattleEventProperty
	2, // 1: BattleEventBattleInfo.skill_info:type_name -> LLNHNHHLCIP
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_BattleEventBattleInfo_proto_init() }
func file_BattleEventBattleInfo_proto_init() {
	if File_BattleEventBattleInfo_proto != nil {
		return
	}
	file_BattleEventProperty_proto_init()
	file_LLNHNHHLCIP_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_BattleEventBattleInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattleEventBattleInfo); i {
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
			RawDescriptor: file_BattleEventBattleInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_BattleEventBattleInfo_proto_goTypes,
		DependencyIndexes: file_BattleEventBattleInfo_proto_depIdxs,
		MessageInfos:      file_BattleEventBattleInfo_proto_msgTypes,
	}.Build()
	File_BattleEventBattleInfo_proto = out.File
	file_BattleEventBattleInfo_proto_rawDesc = nil
	file_BattleEventBattleInfo_proto_goTypes = nil
	file_BattleEventBattleInfo_proto_depIdxs = nil
}
