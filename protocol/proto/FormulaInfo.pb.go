// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: FormulaInfo.proto

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

type FormulaInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FormulaId           uint32                 `protobuf:"varint,1,opt,name=formula_id,json=formulaId,proto3" json:"formula_id,omitempty"`
	IsExpand            bool                   `protobuf:"varint,15,opt,name=is_expand,json=isExpand,proto3" json:"is_expand,omitempty"`
	FormulaBuffTypeList []*FormulaBuffTypeInfo `protobuf:"bytes,8,rep,name=formula_buff_type_list,json=formulaBuffTypeList,proto3" json:"formula_buff_type_list,omitempty"`
}

func (x *FormulaInfo) Reset() {
	*x = FormulaInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FormulaInfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FormulaInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FormulaInfo) ProtoMessage() {}

func (x *FormulaInfo) ProtoReflect() protoreflect.Message {
	mi := &file_FormulaInfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FormulaInfo.ProtoReflect.Descriptor instead.
func (*FormulaInfo) Descriptor() ([]byte, []int) {
	return file_FormulaInfo_proto_rawDescGZIP(), []int{0}
}

func (x *FormulaInfo) GetFormulaId() uint32 {
	if x != nil {
		return x.FormulaId
	}
	return 0
}

func (x *FormulaInfo) GetIsExpand() bool {
	if x != nil {
		return x.IsExpand
	}
	return false
}

func (x *FormulaInfo) GetFormulaBuffTypeList() []*FormulaBuffTypeInfo {
	if x != nil {
		return x.FormulaBuffTypeList
	}
	return nil
}

var File_FormulaInfo_proto protoreflect.FileDescriptor

var file_FormulaInfo_proto_rawDesc = []byte{
	0x0a, 0x11, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x42, 0x75, 0x66, 0x66,
	0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94,
	0x01, 0x0a, 0x0b, 0x46, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x09, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x73, 0x45, 0x78, 0x70, 0x61, 0x6e, 0x64, 0x12, 0x49, 0x0a, 0x16, 0x66, 0x6f,
	0x72, 0x6d, 0x75, 0x6c, 0x61, 0x5f, 0x62, 0x75, 0x66, 0x66, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x46, 0x6f, 0x72,
	0x6d, 0x75, 0x6c, 0x61, 0x42, 0x75, 0x66, 0x66, 0x54, 0x79, 0x70, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x13, 0x66, 0x6f, 0x72, 0x6d, 0x75, 0x6c, 0x61, 0x42, 0x75, 0x66, 0x66, 0x54, 0x79, 0x70,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0xaa, 0x02, 0x1b, 0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68,
	0x65, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FormulaInfo_proto_rawDescOnce sync.Once
	file_FormulaInfo_proto_rawDescData = file_FormulaInfo_proto_rawDesc
)

func file_FormulaInfo_proto_rawDescGZIP() []byte {
	file_FormulaInfo_proto_rawDescOnce.Do(func() {
		file_FormulaInfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_FormulaInfo_proto_rawDescData)
	})
	return file_FormulaInfo_proto_rawDescData
}

var file_FormulaInfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_FormulaInfo_proto_goTypes = []interface{}{
	(*FormulaInfo)(nil),         // 0: FormulaInfo
	(*FormulaBuffTypeInfo)(nil), // 1: FormulaBuffTypeInfo
}
var file_FormulaInfo_proto_depIdxs = []int32{
	1, // 0: FormulaInfo.formula_buff_type_list:type_name -> FormulaBuffTypeInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FormulaInfo_proto_init() }
func file_FormulaInfo_proto_init() {
	if File_FormulaInfo_proto != nil {
		return
	}
	file_FormulaBuffTypeInfo_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_FormulaInfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FormulaInfo); i {
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
			RawDescriptor: file_FormulaInfo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_FormulaInfo_proto_goTypes,
		DependencyIndexes: file_FormulaInfo_proto_depIdxs,
		MessageInfos:      file_FormulaInfo_proto_msgTypes,
	}.Build()
	File_FormulaInfo_proto = out.File
	file_FormulaInfo_proto_rawDesc = nil
	file_FormulaInfo_proto_goTypes = nil
	file_FormulaInfo_proto_depIdxs = nil
}
