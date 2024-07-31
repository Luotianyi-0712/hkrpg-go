// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SceneEntityBuffChange.proto

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

type SceneEntityBuffChange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason   SceneEntityBuffChangeType `protobuf:"varint,4,opt,name=reason,proto3,enum=SceneEntityBuffChangeType" json:"reason,omitempty"`
	EntityId uint32                    `protobuf:"varint,12,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	// Types that are assignable to KEMFJDJHJPK:
	//
	//	*SceneEntityBuffChange_BuffChangeInfo
	//	*SceneEntityBuffChange_RemoveBuffId
	KEMFJDJHJPK isSceneEntityBuffChange_KEMFJDJHJPK `protobuf_oneof:"KEMFJDJHJPK"`
}

func (x *SceneEntityBuffChange) Reset() {
	*x = SceneEntityBuffChange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SceneEntityBuffChange_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SceneEntityBuffChange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SceneEntityBuffChange) ProtoMessage() {}

func (x *SceneEntityBuffChange) ProtoReflect() protoreflect.Message {
	mi := &file_SceneEntityBuffChange_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SceneEntityBuffChange.ProtoReflect.Descriptor instead.
func (*SceneEntityBuffChange) Descriptor() ([]byte, []int) {
	return file_SceneEntityBuffChange_proto_rawDescGZIP(), []int{0}
}

func (x *SceneEntityBuffChange) GetReason() SceneEntityBuffChangeType {
	if x != nil {
		return x.Reason
	}
	return SceneEntityBuffChangeType_SCENE_ENTITY_BUFF_CHANGE_TYPE_DEFAULT
}

func (x *SceneEntityBuffChange) GetEntityId() uint32 {
	if x != nil {
		return x.EntityId
	}
	return 0
}

func (m *SceneEntityBuffChange) GetKEMFJDJHJPK() isSceneEntityBuffChange_KEMFJDJHJPK {
	if m != nil {
		return m.KEMFJDJHJPK
	}
	return nil
}

func (x *SceneEntityBuffChange) GetBuffChangeInfo() *BuffInfo {
	if x, ok := x.GetKEMFJDJHJPK().(*SceneEntityBuffChange_BuffChangeInfo); ok {
		return x.BuffChangeInfo
	}
	return nil
}

func (x *SceneEntityBuffChange) GetRemoveBuffId() uint32 {
	if x, ok := x.GetKEMFJDJHJPK().(*SceneEntityBuffChange_RemoveBuffId); ok {
		return x.RemoveBuffId
	}
	return 0
}

type isSceneEntityBuffChange_KEMFJDJHJPK interface {
	isSceneEntityBuffChange_KEMFJDJHJPK()
}

type SceneEntityBuffChange_BuffChangeInfo struct {
	BuffChangeInfo *BuffInfo `protobuf:"bytes,5,opt,name=buff_change_info,json=buffChangeInfo,proto3,oneof"`
}

type SceneEntityBuffChange_RemoveBuffId struct {
	RemoveBuffId uint32 `protobuf:"varint,15,opt,name=remove_buff_id,json=removeBuffId,proto3,oneof"`
}

func (*SceneEntityBuffChange_BuffChangeInfo) isSceneEntityBuffChange_KEMFJDJHJPK() {}

func (*SceneEntityBuffChange_RemoveBuffId) isSceneEntityBuffChange_KEMFJDJHJPK() {}

var File_SceneEntityBuffChange_proto protoreflect.FileDescriptor

var file_SceneEntityBuffChange_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x75, 0x66,
	0x66, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x53,
	0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x75, 0x66, 0x66, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e,
	0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6,
	0x01, 0x0a, 0x15, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x75,
	0x66, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x75, 0x66, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x10, 0x62, 0x75, 0x66,
	0x66, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00,
	0x52, 0x0e, 0x62, 0x75, 0x66, 0x66, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f,
	0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0c, 0x72, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x42, 0x75, 0x66, 0x66, 0x49, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x4b, 0x45, 0x4d, 0x46,
	0x4a, 0x44, 0x4a, 0x48, 0x4a, 0x50, 0x4b, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SceneEntityBuffChange_proto_rawDescOnce sync.Once
	file_SceneEntityBuffChange_proto_rawDescData = file_SceneEntityBuffChange_proto_rawDesc
)

func file_SceneEntityBuffChange_proto_rawDescGZIP() []byte {
	file_SceneEntityBuffChange_proto_rawDescOnce.Do(func() {
		file_SceneEntityBuffChange_proto_rawDescData = protoimpl.X.CompressGZIP(file_SceneEntityBuffChange_proto_rawDescData)
	})
	return file_SceneEntityBuffChange_proto_rawDescData
}

var file_SceneEntityBuffChange_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SceneEntityBuffChange_proto_goTypes = []interface{}{
	(*SceneEntityBuffChange)(nil),  // 0: SceneEntityBuffChange
	(SceneEntityBuffChangeType)(0), // 1: SceneEntityBuffChangeType
	(*BuffInfo)(nil),               // 2: BuffInfo
}
var file_SceneEntityBuffChange_proto_depIdxs = []int32{
	1, // 0: SceneEntityBuffChange.reason:type_name -> SceneEntityBuffChangeType
	2, // 1: SceneEntityBuffChange.buff_change_info:type_name -> BuffInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_SceneEntityBuffChange_proto_init() }
func file_SceneEntityBuffChange_proto_init() {
	if File_SceneEntityBuffChange_proto != nil {
		return
	}
	file_SceneEntityBuffChangeType_proto_init()
	file_BuffInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_SceneEntityBuffChange_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SceneEntityBuffChange); i {
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
	file_SceneEntityBuffChange_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SceneEntityBuffChange_BuffChangeInfo)(nil),
		(*SceneEntityBuffChange_RemoveBuffId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_SceneEntityBuffChange_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SceneEntityBuffChange_proto_goTypes,
		DependencyIndexes: file_SceneEntityBuffChange_proto_depIdxs,
		MessageInfos:      file_SceneEntityBuffChange_proto_msgTypes,
	}.Build()
	File_SceneEntityBuffChange_proto = out.File
	file_SceneEntityBuffChange_proto_rawDesc = nil
	file_SceneEntityBuffChange_proto_goTypes = nil
	file_SceneEntityBuffChange_proto_depIdxs = nil
}
