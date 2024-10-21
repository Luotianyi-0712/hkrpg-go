// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: ClockParkUseBuffScRsp.proto

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

type ClockParkUseBuffScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	INMCJFNAGCI   uint32       `protobuf:"varint,3,opt,name=INMCJFNAGCI,proto3" json:"INMCJFNAGCI,omitempty"`
	Retcode       uint32       `protobuf:"varint,13,opt,name=retcode,proto3" json:"retcode,omitempty"`
	ScriptId      uint32       `protobuf:"varint,15,opt,name=script_id,json=scriptId,proto3" json:"script_id,omitempty"`
	RogueBuffInfo *LOGAFBJDPKC `protobuf:"bytes,12,opt,name=rogue_buff_info,json=rogueBuffInfo,proto3" json:"rogue_buff_info,omitempty"`
	AAINPIJDLBL   *KPEMELKKNJB `protobuf:"bytes,1363,opt,name=AAINPIJDLBL,proto3" json:"AAINPIJDLBL,omitempty"`
	MENPEMHEIEF   *BLMOJDPMNJF `protobuf:"bytes,602,opt,name=MENPEMHEIEF,proto3" json:"MENPEMHEIEF,omitempty"`
}

func (x *ClockParkUseBuffScRsp) Reset() {
	*x = ClockParkUseBuffScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ClockParkUseBuffScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClockParkUseBuffScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClockParkUseBuffScRsp) ProtoMessage() {}

func (x *ClockParkUseBuffScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ClockParkUseBuffScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClockParkUseBuffScRsp.ProtoReflect.Descriptor instead.
func (*ClockParkUseBuffScRsp) Descriptor() ([]byte, []int) {
	return file_ClockParkUseBuffScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *ClockParkUseBuffScRsp) GetINMCJFNAGCI() uint32 {
	if x != nil {
		return x.INMCJFNAGCI
	}
	return 0
}

func (x *ClockParkUseBuffScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *ClockParkUseBuffScRsp) GetScriptId() uint32 {
	if x != nil {
		return x.ScriptId
	}
	return 0
}

func (x *ClockParkUseBuffScRsp) GetRogueBuffInfo() *LOGAFBJDPKC {
	if x != nil {
		return x.RogueBuffInfo
	}
	return nil
}

func (x *ClockParkUseBuffScRsp) GetAAINPIJDLBL() *KPEMELKKNJB {
	if x != nil {
		return x.AAINPIJDLBL
	}
	return nil
}

func (x *ClockParkUseBuffScRsp) GetMENPEMHEIEF() *BLMOJDPMNJF {
	if x != nil {
		return x.MENPEMHEIEF
	}
	return nil
}

var File_ClockParkUseBuffScRsp_proto protoreflect.FileDescriptor

var file_ClockParkUseBuffScRsp_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x61, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x42, 0x75,
	0x66, 0x66, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x42,
	0x4c, 0x4d, 0x4f, 0x4a, 0x44, 0x50, 0x4d, 0x4e, 0x4a, 0x46, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x11, 0x4b, 0x50, 0x45, 0x4d, 0x45, 0x4c, 0x4b, 0x4b, 0x4e, 0x4a, 0x42, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x4c, 0x4f, 0x47, 0x41, 0x46, 0x42, 0x4a, 0x44, 0x50, 0x4b, 0x43,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x02, 0x0a, 0x15, 0x43, 0x6c, 0x6f, 0x63, 0x6b,
	0x50, 0x61, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x42, 0x75, 0x66, 0x66, 0x53, 0x63, 0x52, 0x73, 0x70,
	0x12, 0x20, 0x0a, 0x0b, 0x49, 0x4e, 0x4d, 0x43, 0x4a, 0x46, 0x4e, 0x41, 0x47, 0x43, 0x49, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x49, 0x4e, 0x4d, 0x43, 0x4a, 0x46, 0x4e, 0x41, 0x47,
	0x43, 0x49, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x0f, 0x72, 0x6f, 0x67,
	0x75, 0x65, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4c, 0x4f, 0x47, 0x41, 0x46, 0x42, 0x4a, 0x44, 0x50, 0x4b, 0x43,
	0x52, 0x0d, 0x72, 0x6f, 0x67, 0x75, 0x65, 0x42, 0x75, 0x66, 0x66, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x2f, 0x0a, 0x0b, 0x41, 0x41, 0x49, 0x4e, 0x50, 0x49, 0x4a, 0x44, 0x4c, 0x42, 0x4c, 0x18, 0xd3,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x4b, 0x50, 0x45, 0x4d, 0x45, 0x4c, 0x4b, 0x4b,
	0x4e, 0x4a, 0x42, 0x52, 0x0b, 0x41, 0x41, 0x49, 0x4e, 0x50, 0x49, 0x4a, 0x44, 0x4c, 0x42, 0x4c,
	0x12, 0x2f, 0x0a, 0x0b, 0x4d, 0x45, 0x4e, 0x50, 0x45, 0x4d, 0x48, 0x45, 0x49, 0x45, 0x46, 0x18,
	0xda, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x42, 0x4c, 0x4d, 0x4f, 0x4a, 0x44, 0x50,
	0x4d, 0x4e, 0x4a, 0x46, 0x52, 0x0b, 0x4d, 0x45, 0x4e, 0x50, 0x45, 0x4d, 0x48, 0x45, 0x49, 0x45,
	0x46, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61,
	0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ClockParkUseBuffScRsp_proto_rawDescOnce sync.Once
	file_ClockParkUseBuffScRsp_proto_rawDescData = file_ClockParkUseBuffScRsp_proto_rawDesc
)

func file_ClockParkUseBuffScRsp_proto_rawDescGZIP() []byte {
	file_ClockParkUseBuffScRsp_proto_rawDescOnce.Do(func() {
		file_ClockParkUseBuffScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_ClockParkUseBuffScRsp_proto_rawDescData)
	})
	return file_ClockParkUseBuffScRsp_proto_rawDescData
}

var file_ClockParkUseBuffScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ClockParkUseBuffScRsp_proto_goTypes = []interface{}{
	(*ClockParkUseBuffScRsp)(nil), // 0: ClockParkUseBuffScRsp
	(*LOGAFBJDPKC)(nil),           // 1: LOGAFBJDPKC
	(*KPEMELKKNJB)(nil),           // 2: KPEMELKKNJB
	(*BLMOJDPMNJF)(nil),           // 3: BLMOJDPMNJF
}
var file_ClockParkUseBuffScRsp_proto_depIdxs = []int32{
	1, // 0: ClockParkUseBuffScRsp.rogue_buff_info:type_name -> LOGAFBJDPKC
	2, // 1: ClockParkUseBuffScRsp.AAINPIJDLBL:type_name -> KPEMELKKNJB
	3, // 2: ClockParkUseBuffScRsp.MENPEMHEIEF:type_name -> BLMOJDPMNJF
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ClockParkUseBuffScRsp_proto_init() }
func file_ClockParkUseBuffScRsp_proto_init() {
	if File_ClockParkUseBuffScRsp_proto != nil {
		return
	}
	file_BLMOJDPMNJF_proto_init()
	file_KPEMELKKNJB_proto_init()
	file_LOGAFBJDPKC_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ClockParkUseBuffScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClockParkUseBuffScRsp); i {
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
			RawDescriptor: file_ClockParkUseBuffScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ClockParkUseBuffScRsp_proto_goTypes,
		DependencyIndexes: file_ClockParkUseBuffScRsp_proto_depIdxs,
		MessageInfos:      file_ClockParkUseBuffScRsp_proto_msgTypes,
	}.Build()
	File_ClockParkUseBuffScRsp_proto = out.File
	file_ClockParkUseBuffScRsp_proto_rawDesc = nil
	file_ClockParkUseBuffScRsp_proto_goTypes = nil
	file_ClockParkUseBuffScRsp_proto_depIdxs = nil
}
