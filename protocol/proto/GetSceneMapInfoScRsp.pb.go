// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: GetSceneMapInfoScRsp.proto

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

type GetSceneMapInfoScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MFDIBEECLPP  bool            `protobuf:"varint,10,opt,name=MFDIBEECLPP,proto3" json:"MFDIBEECLPP,omitempty"`
	SceneMapInfo []*SceneMapInfo `protobuf:"bytes,5,rep,name=scene_map_info,json=sceneMapInfo,proto3" json:"scene_map_info,omitempty"`
	Retcode      uint32          `protobuf:"varint,9,opt,name=retcode,proto3" json:"retcode,omitempty"`
}

func (x *GetSceneMapInfoScRsp) Reset() {
	*x = GetSceneMapInfoScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GetSceneMapInfoScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSceneMapInfoScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSceneMapInfoScRsp) ProtoMessage() {}

func (x *GetSceneMapInfoScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_GetSceneMapInfoScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSceneMapInfoScRsp.ProtoReflect.Descriptor instead.
func (*GetSceneMapInfoScRsp) Descriptor() ([]byte, []int) {
	return file_GetSceneMapInfoScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *GetSceneMapInfoScRsp) GetMFDIBEECLPP() bool {
	if x != nil {
		return x.MFDIBEECLPP
	}
	return false
}

func (x *GetSceneMapInfoScRsp) GetSceneMapInfo() []*SceneMapInfo {
	if x != nil {
		return x.SceneMapInfo
	}
	return nil
}

func (x *GetSceneMapInfoScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

var File_GetSceneMapInfoScRsp_proto protoreflect.FileDescriptor

var file_GetSceneMapInfoScRsp_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x53, 0x63,
	0x65, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x87, 0x01, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x61, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x46, 0x44,
	0x49, 0x42, 0x45, 0x45, 0x43, 0x4c, 0x50, 0x50, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x4d, 0x46, 0x44, 0x49, 0x42, 0x45, 0x45, 0x43, 0x4c, 0x50, 0x50, 0x12, 0x33, 0x0a, 0x0e, 0x73,
	0x63, 0x65, 0x6e, 0x65, 0x5f, 0x6d, 0x61, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0c, 0x73, 0x63, 0x65, 0x6e, 0x65, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45,
	0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_GetSceneMapInfoScRsp_proto_rawDescOnce sync.Once
	file_GetSceneMapInfoScRsp_proto_rawDescData = file_GetSceneMapInfoScRsp_proto_rawDesc
)

func file_GetSceneMapInfoScRsp_proto_rawDescGZIP() []byte {
	file_GetSceneMapInfoScRsp_proto_rawDescOnce.Do(func() {
		file_GetSceneMapInfoScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_GetSceneMapInfoScRsp_proto_rawDescData)
	})
	return file_GetSceneMapInfoScRsp_proto_rawDescData
}

var file_GetSceneMapInfoScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_GetSceneMapInfoScRsp_proto_goTypes = []interface{}{
	(*GetSceneMapInfoScRsp)(nil), // 0: GetSceneMapInfoScRsp
	(*SceneMapInfo)(nil),         // 1: SceneMapInfo
}
var file_GetSceneMapInfoScRsp_proto_depIdxs = []int32{
	1, // 0: GetSceneMapInfoScRsp.scene_map_info:type_name -> SceneMapInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_GetSceneMapInfoScRsp_proto_init() }
func file_GetSceneMapInfoScRsp_proto_init() {
	if File_GetSceneMapInfoScRsp_proto != nil {
		return
	}
	file_SceneMapInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GetSceneMapInfoScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSceneMapInfoScRsp); i {
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
			RawDescriptor: file_GetSceneMapInfoScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GetSceneMapInfoScRsp_proto_goTypes,
		DependencyIndexes: file_GetSceneMapInfoScRsp_proto_depIdxs,
		MessageInfos:      file_GetSceneMapInfoScRsp_proto_msgTypes,
	}.Build()
	File_GetSceneMapInfoScRsp_proto = out.File
	file_GetSceneMapInfoScRsp_proto_rawDesc = nil
	file_GetSceneMapInfoScRsp_proto_goTypes = nil
	file_GetSceneMapInfoScRsp_proto_depIdxs = nil
}
