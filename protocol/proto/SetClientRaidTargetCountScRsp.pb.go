// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: SetClientRaidTargetCountScRsp.proto

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

type SetClientRaidTargetCountScRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Progress    uint32 `protobuf:"varint,11,opt,name=progress,proto3" json:"progress,omitempty"`
	Retcode     uint32 `protobuf:"varint,13,opt,name=retcode,proto3" json:"retcode,omitempty"`
	BFDNDBPPEOL uint32 `protobuf:"varint,1,opt,name=BFDNDBPPEOL,proto3" json:"BFDNDBPPEOL,omitempty"`
}

func (x *SetClientRaidTargetCountScRsp) Reset() {
	*x = SetClientRaidTargetCountScRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_SetClientRaidTargetCountScRsp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetClientRaidTargetCountScRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetClientRaidTargetCountScRsp) ProtoMessage() {}

func (x *SetClientRaidTargetCountScRsp) ProtoReflect() protoreflect.Message {
	mi := &file_SetClientRaidTargetCountScRsp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetClientRaidTargetCountScRsp.ProtoReflect.Descriptor instead.
func (*SetClientRaidTargetCountScRsp) Descriptor() ([]byte, []int) {
	return file_SetClientRaidTargetCountScRsp_proto_rawDescGZIP(), []int{0}
}

func (x *SetClientRaidTargetCountScRsp) GetProgress() uint32 {
	if x != nil {
		return x.Progress
	}
	return 0
}

func (x *SetClientRaidTargetCountScRsp) GetRetcode() uint32 {
	if x != nil {
		return x.Retcode
	}
	return 0
}

func (x *SetClientRaidTargetCountScRsp) GetBFDNDBPPEOL() uint32 {
	if x != nil {
		return x.BFDNDBPPEOL
	}
	return 0
}

var File_SetClientRaidTargetCountScRsp_proto protoreflect.FileDescriptor

var file_SetClientRaidTargetCountScRsp_proto_rawDesc = []byte{
	0x0a, 0x23, 0x53, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x69, 0x64, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x63, 0x52, 0x73, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x1d, 0x53, 0x65, 0x74, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x52, 0x61, 0x69, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x63, 0x52, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x42, 0x46, 0x44, 0x4e, 0x44, 0x42, 0x50, 0x50, 0x45, 0x4f, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0b, 0x42, 0x46, 0x44, 0x4e, 0x44, 0x42, 0x50, 0x50, 0x45, 0x4f, 0x4c, 0x42, 0x2e,
	0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65,
	0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_SetClientRaidTargetCountScRsp_proto_rawDescOnce sync.Once
	file_SetClientRaidTargetCountScRsp_proto_rawDescData = file_SetClientRaidTargetCountScRsp_proto_rawDesc
)

func file_SetClientRaidTargetCountScRsp_proto_rawDescGZIP() []byte {
	file_SetClientRaidTargetCountScRsp_proto_rawDescOnce.Do(func() {
		file_SetClientRaidTargetCountScRsp_proto_rawDescData = protoimpl.X.CompressGZIP(file_SetClientRaidTargetCountScRsp_proto_rawDescData)
	})
	return file_SetClientRaidTargetCountScRsp_proto_rawDescData
}

var file_SetClientRaidTargetCountScRsp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_SetClientRaidTargetCountScRsp_proto_goTypes = []interface{}{
	(*SetClientRaidTargetCountScRsp)(nil), // 0: SetClientRaidTargetCountScRsp
}
var file_SetClientRaidTargetCountScRsp_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_SetClientRaidTargetCountScRsp_proto_init() }
func file_SetClientRaidTargetCountScRsp_proto_init() {
	if File_SetClientRaidTargetCountScRsp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_SetClientRaidTargetCountScRsp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetClientRaidTargetCountScRsp); i {
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
			RawDescriptor: file_SetClientRaidTargetCountScRsp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_SetClientRaidTargetCountScRsp_proto_goTypes,
		DependencyIndexes: file_SetClientRaidTargetCountScRsp_proto_depIdxs,
		MessageInfos:      file_SetClientRaidTargetCountScRsp_proto_msgTypes,
	}.Build()
	File_SetClientRaidTargetCountScRsp_proto = out.File
	file_SetClientRaidTargetCountScRsp_proto_rawDesc = nil
	file_SetClientRaidTargetCountScRsp_proto_goTypes = nil
	file_SetClientRaidTargetCountScRsp_proto_depIdxs = nil
}
