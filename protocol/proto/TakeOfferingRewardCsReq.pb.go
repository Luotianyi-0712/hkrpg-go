// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: TakeOfferingRewardCsReq.proto

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

type TakeOfferingRewardCsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BBOAEJKGEBO uint32   `protobuf:"varint,4,opt,name=BBOAEJKGEBO,proto3" json:"BBOAEJKGEBO,omitempty"`
	MonsterId   uint32   `protobuf:"varint,14,opt,name=monster_id,json=monsterId,proto3" json:"monster_id,omitempty"`
	KMHLLGLJNMN []uint32 `protobuf:"varint,3,rep,packed,name=KMHLLGLJNMN,proto3" json:"KMHLLGLJNMN,omitempty"`
}

func (x *TakeOfferingRewardCsReq) Reset() {
	*x = TakeOfferingRewardCsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_TakeOfferingRewardCsReq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TakeOfferingRewardCsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TakeOfferingRewardCsReq) ProtoMessage() {}

func (x *TakeOfferingRewardCsReq) ProtoReflect() protoreflect.Message {
	mi := &file_TakeOfferingRewardCsReq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TakeOfferingRewardCsReq.ProtoReflect.Descriptor instead.
func (*TakeOfferingRewardCsReq) Descriptor() ([]byte, []int) {
	return file_TakeOfferingRewardCsReq_proto_rawDescGZIP(), []int{0}
}

func (x *TakeOfferingRewardCsReq) GetBBOAEJKGEBO() uint32 {
	if x != nil {
		return x.BBOAEJKGEBO
	}
	return 0
}

func (x *TakeOfferingRewardCsReq) GetMonsterId() uint32 {
	if x != nil {
		return x.MonsterId
	}
	return 0
}

func (x *TakeOfferingRewardCsReq) GetKMHLLGLJNMN() []uint32 {
	if x != nil {
		return x.KMHLLGLJNMN
	}
	return nil
}

var File_TakeOfferingRewardCsReq_proto protoreflect.FileDescriptor

var file_TakeOfferingRewardCsReq_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x54, 0x61, 0x6b, 0x65, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x43, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7c, 0x0a, 0x17, 0x54, 0x61, 0x6b, 0x65, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x43, 0x73, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x42,
	0x4f, 0x41, 0x45, 0x4a, 0x4b, 0x47, 0x45, 0x42, 0x4f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0b, 0x42, 0x42, 0x4f, 0x41, 0x45, 0x4a, 0x4b, 0x47, 0x45, 0x42, 0x4f, 0x12, 0x1d, 0x0a, 0x0a,
	0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x6d, 0x6f, 0x6e, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x4b,
	0x4d, 0x48, 0x4c, 0x4c, 0x47, 0x4c, 0x4a, 0x4e, 0x4d, 0x4e, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x0b, 0x4b, 0x4d, 0x48, 0x4c, 0x4c, 0x47, 0x4c, 0x4a, 0x4e, 0x4d, 0x4e, 0x42, 0x2e, 0x5a,
	0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa,
	0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_TakeOfferingRewardCsReq_proto_rawDescOnce sync.Once
	file_TakeOfferingRewardCsReq_proto_rawDescData = file_TakeOfferingRewardCsReq_proto_rawDesc
)

func file_TakeOfferingRewardCsReq_proto_rawDescGZIP() []byte {
	file_TakeOfferingRewardCsReq_proto_rawDescOnce.Do(func() {
		file_TakeOfferingRewardCsReq_proto_rawDescData = protoimpl.X.CompressGZIP(file_TakeOfferingRewardCsReq_proto_rawDescData)
	})
	return file_TakeOfferingRewardCsReq_proto_rawDescData
}

var file_TakeOfferingRewardCsReq_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_TakeOfferingRewardCsReq_proto_goTypes = []interface{}{
	(*TakeOfferingRewardCsReq)(nil), // 0: TakeOfferingRewardCsReq
}
var file_TakeOfferingRewardCsReq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_TakeOfferingRewardCsReq_proto_init() }
func file_TakeOfferingRewardCsReq_proto_init() {
	if File_TakeOfferingRewardCsReq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_TakeOfferingRewardCsReq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TakeOfferingRewardCsReq); i {
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
			RawDescriptor: file_TakeOfferingRewardCsReq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_TakeOfferingRewardCsReq_proto_goTypes,
		DependencyIndexes: file_TakeOfferingRewardCsReq_proto_depIdxs,
		MessageInfos:      file_TakeOfferingRewardCsReq_proto_msgTypes,
	}.Build()
	File_TakeOfferingRewardCsReq_proto = out.File
	file_TakeOfferingRewardCsReq_proto_rawDesc = nil
	file_TakeOfferingRewardCsReq_proto_goTypes = nil
	file_TakeOfferingRewardCsReq_proto_depIdxs = nil
}
