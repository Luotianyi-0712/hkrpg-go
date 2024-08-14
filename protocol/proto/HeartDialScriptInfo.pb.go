// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: HeartDialScriptInfo.proto

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

type HeartDialScriptInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JAOCGOEMJPI    bool                 `protobuf:"varint,11,opt,name=JAOCGOEMJPI,proto3" json:"JAOCGOEMJPI,omitempty"`
	CurEmotionType HeartDialEmotionType `protobuf:"varint,10,opt,name=cur_emotion_type,json=curEmotionType,proto3,enum=HeartDialEmotionType" json:"cur_emotion_type,omitempty"`
	NNEEAINEBPM    bool                 `protobuf:"varint,1,opt,name=NNEEAINEBPM,proto3" json:"NNEEAINEBPM,omitempty"`
	ScriptId       uint32               `protobuf:"varint,5,opt,name=script_id,json=scriptId,proto3" json:"script_id,omitempty"`
	Step           HeartDialStepType    `protobuf:"varint,9,opt,name=step,proto3,enum=HeartDialStepType" json:"step,omitempty"`
}

func (x *HeartDialScriptInfo) Reset() {
	*x = HeartDialScriptInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HeartDialScriptInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartDialScriptInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartDialScriptInfo) ProtoMessage() {}

func (x *HeartDialScriptInfo) ProtoReflect() protoreflect.Message {
	mi := &file_HeartDialScriptInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartDialScriptInfo.ProtoReflect.Descriptor instead.
func (*HeartDialScriptInfo) Descriptor() ([]byte, []int) {
	return file_HeartDialScriptInfo_proto_rawDescGZIP(), []int{0}
}

func (x *HeartDialScriptInfo) GetJAOCGOEMJPI() bool {
	if x != nil {
		return x.JAOCGOEMJPI
	}
	return false
}

func (x *HeartDialScriptInfo) GetCurEmotionType() HeartDialEmotionType {
	if x != nil {
		return x.CurEmotionType
	}
	return HeartDialEmotionType_HEART_DIAL_EMOTION_TYPE_PEACE
}

func (x *HeartDialScriptInfo) GetNNEEAINEBPM() bool {
	if x != nil {
		return x.NNEEAINEBPM
	}
	return false
}

func (x *HeartDialScriptInfo) GetScriptId() uint32 {
	if x != nil {
		return x.ScriptId
	}
	return 0
}

func (x *HeartDialScriptInfo) GetStep() HeartDialStepType {
	if x != nil {
		return x.Step
	}
	return HeartDialStepType_HEART_DIAL_STEP_TYPE_MISSING
}

var File_HeartDialScriptInfo_proto protoreflect.FileDescriptor

var file_HeartDialScriptInfo_proto_rawDesc = []byte{
	0x0a, 0x19, 0x48, 0x65, 0x61, 0x72, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x53, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x45, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x48, 0x65, 0x61, 0x72, 0x74, 0x44, 0x69,
	0x61, 0x6c, 0x53, 0x74, 0x65, 0x70, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xdf, 0x01, 0x0a, 0x13, 0x48, 0x65, 0x61, 0x72, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x53, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x4a, 0x41, 0x4f, 0x43,
	0x47, 0x4f, 0x45, 0x4d, 0x4a, 0x50, 0x49, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4a,
	0x41, 0x4f, 0x43, 0x47, 0x4f, 0x45, 0x4d, 0x4a, 0x50, 0x49, 0x12, 0x3f, 0x0a, 0x10, 0x63, 0x75,
	0x72, 0x5f, 0x65, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x44, 0x69, 0x61, 0x6c,
	0x45, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0e, 0x63, 0x75, 0x72,
	0x45, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x4e,
	0x4e, 0x45, 0x45, 0x41, 0x49, 0x4e, 0x45, 0x42, 0x50, 0x4d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x4e, 0x4e, 0x45, 0x45, 0x41, 0x49, 0x4e, 0x45, 0x42, 0x50, 0x4d, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x08, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x73, 0x74,
	0x65, 0x70, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x44, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x65, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x73, 0x74,
	0x65, 0x70, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02,
	0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_HeartDialScriptInfo_proto_rawDescOnce sync.Once
	file_HeartDialScriptInfo_proto_rawDescData = file_HeartDialScriptInfo_proto_rawDesc
)

func file_HeartDialScriptInfo_proto_rawDescGZIP() []byte {
	file_HeartDialScriptInfo_proto_rawDescOnce.Do(func() {
		file_HeartDialScriptInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_HeartDialScriptInfo_proto_rawDescData)
	})
	return file_HeartDialScriptInfo_proto_rawDescData
}

var file_HeartDialScriptInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HeartDialScriptInfo_proto_goTypes = []interface{}{
	(*HeartDialScriptInfo)(nil), // 0: HeartDialScriptInfo
	(HeartDialEmotionType)(0),   // 1: HeartDialEmotionType
	(HeartDialStepType)(0),      // 2: HeartDialStepType
}
var file_HeartDialScriptInfo_proto_depIdxs = []int32{
	1, // 0: HeartDialScriptInfo.cur_emotion_type:type_name -> HeartDialEmotionType
	2, // 1: HeartDialScriptInfo.step:type_name -> HeartDialStepType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_HeartDialScriptInfo_proto_init() }
func file_HeartDialScriptInfo_proto_init() {
	if File_HeartDialScriptInfo_proto != nil {
		return
	}
	file_HeartDialEmotionType_proto_init()
	file_HeartDialStepType_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_HeartDialScriptInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartDialScriptInfo); i {
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
			RawDescriptor: file_HeartDialScriptInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HeartDialScriptInfo_proto_goTypes,
		DependencyIndexes: file_HeartDialScriptInfo_proto_depIdxs,
		MessageInfos:      file_HeartDialScriptInfo_proto_msgTypes,
	}.Build()
	File_HeartDialScriptInfo_proto = out.File
	file_HeartDialScriptInfo_proto_rawDesc = nil
	file_HeartDialScriptInfo_proto_goTypes = nil
	file_HeartDialScriptInfo_proto_depIdxs = nil
}
