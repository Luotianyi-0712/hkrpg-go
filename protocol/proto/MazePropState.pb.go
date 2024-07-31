// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: MazePropState.proto

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

type MazePropState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigId uint32 `protobuf:"varint,14,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	GroupId  uint32 `protobuf:"varint,15,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	State    uint32 `protobuf:"varint,8,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *MazePropState) Reset() {
	*x = MazePropState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MazePropState_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MazePropState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MazePropState) ProtoMessage() {}

func (x *MazePropState) ProtoReflect() protoreflect.Message {
	mi := &file_MazePropState_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MazePropState.ProtoReflect.Descriptor instead.
func (*MazePropState) Descriptor() ([]byte, []int) {
	return file_MazePropState_proto_rawDescGZIP(), []int{0}
}

func (x *MazePropState) GetConfigId() uint32 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *MazePropState) GetGroupId() uint32 {
	if x != nil {
		return x.GroupId
	}
	return 0
}

func (x *MazePropState) GetState() uint32 {
	if x != nil {
		return x.State
	}
	return 0
}

var File_MazePropState_proto protoreflect.FileDescriptor

var file_MazePropState_proto_rawDesc = []byte{
	0x0a, 0x13, 0x4d, 0x61, 0x7a, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5d, 0x0a, 0x0d, 0x4d, 0x61, 0x7a, 0x65, 0x50, 0x72, 0x6f,
	0x70, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MazePropState_proto_rawDescOnce sync.Once
	file_MazePropState_proto_rawDescData = file_MazePropState_proto_rawDesc
)

func file_MazePropState_proto_rawDescGZIP() []byte {
	file_MazePropState_proto_rawDescOnce.Do(func() {
		file_MazePropState_proto_rawDescData = protoimpl.X.CompressGZIP(file_MazePropState_proto_rawDescData)
	})
	return file_MazePropState_proto_rawDescData
}

var file_MazePropState_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MazePropState_proto_goTypes = []interface{}{
	(*MazePropState)(nil), // 0: MazePropState
}
var file_MazePropState_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MazePropState_proto_init() }
func file_MazePropState_proto_init() {
	if File_MazePropState_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MazePropState_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MazePropState); i {
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
			RawDescriptor: file_MazePropState_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MazePropState_proto_goTypes,
		DependencyIndexes: file_MazePropState_proto_depIdxs,
		MessageInfos:      file_MazePropState_proto_msgTypes,
	}.Build()
	File_MazePropState_proto = out.File
	file_MazePropState_proto_rawDesc = nil
	file_MazePropState_proto_goTypes = nil
	file_MazePropState_proto_depIdxs = nil
}
