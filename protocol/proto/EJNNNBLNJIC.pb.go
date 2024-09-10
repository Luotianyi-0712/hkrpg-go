// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: EJNNNBLNJIC.proto

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

type EJNNNBLNJIC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status    SwordTrainingStatus `protobuf:"varint,15,opt,name=Status,proto3,enum=SwordTrainingStatus" json:"Status,omitempty"`
	VaryValue uint32              `protobuf:"varint,11,opt,name=vary_value,json=varyValue,proto3" json:"vary_value,omitempty"`
	CurValue  uint32              `protobuf:"varint,1,opt,name=cur_value,json=curValue,proto3" json:"cur_value,omitempty"`
}

func (x *EJNNNBLNJIC) Reset() {
	*x = EJNNNBLNJIC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_EJNNNBLNJIC_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EJNNNBLNJIC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EJNNNBLNJIC) ProtoMessage() {}

func (x *EJNNNBLNJIC) ProtoReflect() protoreflect.Message {
	mi := &file_EJNNNBLNJIC_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EJNNNBLNJIC.ProtoReflect.Descriptor instead.
func (*EJNNNBLNJIC) Descriptor() ([]byte, []int) {
	return file_EJNNNBLNJIC_proto_rawDescGZIP(), []int{0}
}

func (x *EJNNNBLNJIC) GetStatus() SwordTrainingStatus {
	if x != nil {
		return x.Status
	}
	return SwordTrainingStatus_SWORD_TRAINING_STATUS_TYPE_NONE
}

func (x *EJNNNBLNJIC) GetVaryValue() uint32 {
	if x != nil {
		return x.VaryValue
	}
	return 0
}

func (x *EJNNNBLNJIC) GetCurValue() uint32 {
	if x != nil {
		return x.CurValue
	}
	return 0
}

var File_EJNNNBLNJIC_proto protoreflect.FileDescriptor

var file_EJNNNBLNJIC_proto_rawDesc = []byte{
	0x0a, 0x11, 0x45, 0x4a, 0x4e, 0x4e, 0x4e, 0x42, 0x4c, 0x4e, 0x4a, 0x49, 0x43, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77,
	0x0a, 0x0b, 0x45, 0x4a, 0x4e, 0x4e, 0x4e, 0x42, 0x4c, 0x4e, 0x4a, 0x49, 0x43, 0x12, 0x2c, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x53, 0x77, 0x6f, 0x72, 0x64, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x76,
	0x61, 0x72, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x76, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75,
	0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x63,
	0x75, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x2e, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c,
	0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EJNNNBLNJIC_proto_rawDescOnce sync.Once
	file_EJNNNBLNJIC_proto_rawDescData = file_EJNNNBLNJIC_proto_rawDesc
)

func file_EJNNNBLNJIC_proto_rawDescGZIP() []byte {
	file_EJNNNBLNJIC_proto_rawDescOnce.Do(func() {
		file_EJNNNBLNJIC_proto_rawDescData = protoimpl.X.CompressGZIP(file_EJNNNBLNJIC_proto_rawDescData)
	})
	return file_EJNNNBLNJIC_proto_rawDescData
}

var file_EJNNNBLNJIC_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_EJNNNBLNJIC_proto_goTypes = []interface{}{
	(*EJNNNBLNJIC)(nil),      // 0: EJNNNBLNJIC
	(SwordTrainingStatus)(0), // 1: SwordTrainingStatus
}
var file_EJNNNBLNJIC_proto_depIdxs = []int32{
	1, // 0: EJNNNBLNJIC.Status:type_name -> SwordTrainingStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_EJNNNBLNJIC_proto_init() }
func file_EJNNNBLNJIC_proto_init() {
	if File_EJNNNBLNJIC_proto != nil {
		return
	}
	file_SwordTrainingStatus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_EJNNNBLNJIC_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EJNNNBLNJIC); i {
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
			RawDescriptor: file_EJNNNBLNJIC_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EJNNNBLNJIC_proto_goTypes,
		DependencyIndexes: file_EJNNNBLNJIC_proto_depIdxs,
		MessageInfos:      file_EJNNNBLNJIC_proto_msgTypes,
	}.Build()
	File_EJNNNBLNJIC_proto = out.File
	file_EJNNNBLNJIC_proto_rawDesc = nil
	file_EJNNNBLNJIC_proto_goTypes = nil
	file_EJNNNBLNJIC_proto_depIdxs = nil
}
