// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GLICMBJEKHB.proto

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

type GLICMBJEKHB struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AeonIdList             []uint32 `protobuf:"varint,14,rep,packed,name=aeon_id_list,json=aeonIdList,proto3" json:"aeon_id_list,omitempty"`
	UnlockedAeonEnhanceNum uint32   `protobuf:"varint,15,opt,name=unlocked_aeon_enhance_num,json=unlockedAeonEnhanceNum,proto3" json:"unlocked_aeon_enhance_num,omitempty"`
	UnlockedAeonNum        uint32   `protobuf:"varint,11,opt,name=unlocked_aeon_num,json=unlockedAeonNum,proto3" json:"unlocked_aeon_num,omitempty"`
	AeonId                 uint32   `protobuf:"varint,6,opt,name=aeon_id,json=aeonId,proto3" json:"aeon_id,omitempty"`
	IsUnlocked             bool     `protobuf:"varint,1,opt,name=is_unlocked,json=isUnlocked,proto3" json:"is_unlocked,omitempty"`
}

func (x *GLICMBJEKHB) Reset() {
	*x = GLICMBJEKHB{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GLICMBJEKHB_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GLICMBJEKHB) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GLICMBJEKHB) ProtoMessage() {}

func (x *GLICMBJEKHB) ProtoReflect() protoreflect.Message {
	mi := &file_GLICMBJEKHB_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GLICMBJEKHB.ProtoReflect.Descriptor instead.
func (*GLICMBJEKHB) Descriptor() ([]byte, []int) {
	return file_GLICMBJEKHB_proto_rawDescGZIP(), []int{0}
}

func (x *GLICMBJEKHB) GetAeonIdList() []uint32 {
	if x != nil {
		return x.AeonIdList
	}
	return nil
}

func (x *GLICMBJEKHB) GetUnlockedAeonEnhanceNum() uint32 {
	if x != nil {
		return x.UnlockedAeonEnhanceNum
	}
	return 0
}

func (x *GLICMBJEKHB) GetUnlockedAeonNum() uint32 {
	if x != nil {
		return x.UnlockedAeonNum
	}
	return 0
}

func (x *GLICMBJEKHB) GetAeonId() uint32 {
	if x != nil {
		return x.AeonId
	}
	return 0
}

func (x *GLICMBJEKHB) GetIsUnlocked() bool {
	if x != nil {
		return x.IsUnlocked
	}
	return false
}

var File_GLICMBJEKHB_proto protoreflect.FileDescriptor

var file_GLICMBJEKHB_proto_rawDesc = []byte{
	0x0a, 0x11, 0x47, 0x4c, 0x49, 0x43, 0x4d, 0x42, 0x4a, 0x45, 0x4b, 0x48, 0x42, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x0b, 0x47, 0x4c, 0x49, 0x43, 0x4d, 0x42, 0x4a, 0x45,
	0x4b, 0x48, 0x42, 0x12, 0x20, 0x0a, 0x0c, 0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x61, 0x65, 0x6f, 0x6e, 0x49,
	0x64, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x19, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65,
	0x64, 0x5f, 0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x65, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x16, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b,
	0x65, 0x64, 0x41, 0x65, 0x6f, 0x6e, 0x45, 0x6e, 0x68, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d,
	0x12, 0x2a, 0x0a, 0x11, 0x75, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x61, 0x65, 0x6f,
	0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x75, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x65, 0x64, 0x41, 0x65, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x12, 0x17, 0x0a, 0x07,
	0x61, 0x65, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x61,
	0x65, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x75, 0x6e, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x55, 0x6e,
	0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e,
	0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GLICMBJEKHB_proto_rawDescOnce sync.Once
	file_GLICMBJEKHB_proto_rawDescData = file_GLICMBJEKHB_proto_rawDesc
)

func file_GLICMBJEKHB_proto_rawDescGZIP() []byte {
	file_GLICMBJEKHB_proto_rawDescOnce.Do(func() {
		file_GLICMBJEKHB_proto_rawDescData = protoimpl.X.CompressGZIP(file_GLICMBJEKHB_proto_rawDescData)
	})
	return file_GLICMBJEKHB_proto_rawDescData
}

var file_GLICMBJEKHB_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GLICMBJEKHB_proto_goTypes = []interface{}{
	(*GLICMBJEKHB)(nil), // 0: GLICMBJEKHB
}
var file_GLICMBJEKHB_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GLICMBJEKHB_proto_init() }
func file_GLICMBJEKHB_proto_init() {
	if File_GLICMBJEKHB_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GLICMBJEKHB_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GLICMBJEKHB); i {
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
			RawDescriptor: file_GLICMBJEKHB_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GLICMBJEKHB_proto_goTypes,
		DependencyIndexes: file_GLICMBJEKHB_proto_depIdxs,
		MessageInfos:      file_GLICMBJEKHB_proto_msgTypes,
	}.Build()
	File_GLICMBJEKHB_proto = out.File
	file_GLICMBJEKHB_proto_rawDesc = nil
	file_GLICMBJEKHB_proto_goTypes = nil
	file_GLICMBJEKHB_proto_depIdxs = nil
}
