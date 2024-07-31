// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: UpdateMovieRacingDataScRsp.proto

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

type UpdateMovieRacingDataScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Retcode    uint32      `protobuf:"varint,5,opt,name=retcode,proto3" json:"retcode,omitempty"`
	RacingData *RacingData `protobuf:"bytes,12,opt,name=racing_data,json=racingData,proto3" json:"racing_data,omitempty"`
}

func (x *UpdateMovieRacingDataScRsp) Reset() {
	*x = UpdateMovieRacingDataScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_UpdateMovieRacingDataScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMovieRacingDataScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMovieRacingDataScRsp) ProtoMessage() {}

func (x *UpdateMovieRacingDataScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_UpdateMovieRacingDataScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMovieRacingDataScRsp.ProtoReflect.Descriptor instead.
func (*UpdateMovieRacingDataScRsp) Descriptor() ([]byte, []int) {
	return file_UpdateMovieRacingDataScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateMovieRacingDataScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *UpdateMovieRacingDataScRsp) GetRacingData() *RacingData {
	if x != nil {
		return x.RacingData
	}
	return nil
}

var File_UpdateMovieRacingDataScRsp_proto protoreflect.FileDescriptor

var file_UpdateMovieRacingDataScRsp_proto_rawDesc = []byte{
	0x0a, 0x20, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x10, 0x52, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x6f,
	0x76, 0x69, 0x65, 0x52, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x53, 0x63, 0x52,
	0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x0b,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x52, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a,
	0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b,
	0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_UpdateMovieRacingDataScRsp_proto_rawDescOnce sync.Once
	file_UpdateMovieRacingDataScRsp_proto_rawDescData = file_UpdateMovieRacingDataScRsp_proto_rawDesc
)

func file_UpdateMovieRacingDataScRsp_proto_rawDescGZIP() []byte {
	file_UpdateMovieRacingDataScRsp_proto_rawDescOnce.Do(func() {
		file_UpdateMovieRacingDataScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_UpdateMovieRacingDataScRsp_proto_rawDescData)
	})
	return file_UpdateMovieRacingDataScRsp_proto_rawDescData
}

var file_UpdateMovieRacingDataScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_UpdateMovieRacingDataScRsp_proto_goTypes = []interface{}{
	(*UpdateMovieRacingDataScRsp)(nil), // 0: UpdateMovieRacingDataScRsp
	(*RacingData)(nil),                 // 1: RacingData
}
var file_UpdateMovieRacingDataScRsp_proto_depIdxs = []int32{
	1, // 0: UpdateMovieRacingDataScRsp.racing_data:type_name -> RacingData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_UpdateMovieRacingDataScRsp_proto_init() }
func file_UpdateMovieRacingDataScRsp_proto_init() {
	if File_UpdateMovieRacingDataScRsp_proto != nil {
		return
	}
	file_RacingData_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_UpdateMovieRacingDataScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMovieRacingDataScRsp); i {
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
			RawDescriptor: file_UpdateMovieRacingDataScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_UpdateMovieRacingDataScRsp_proto_goTypes,
		DependencyIndexes: file_UpdateMovieRacingDataScRsp_proto_depIdxs,
		MessageInfos:      file_UpdateMovieRacingDataScRsp_proto_msgTypes,
	}.Build()
	File_UpdateMovieRacingDataScRsp_proto = out.File
	file_UpdateMovieRacingDataScRsp_proto_rawDesc = nil
	file_UpdateMovieRacingDataScRsp_proto_goTypes = nil
	file_UpdateMovieRacingDataScRsp_proto_depIdxs = nil
}
