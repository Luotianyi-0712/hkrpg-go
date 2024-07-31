// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.0
// source: JPFECHLHHEN.proto

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

type JPFECHLHHEN struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OIIEELPECGL           uint32   `protobuf:"varint,12,opt,name=OIIEELPECGL,proto3" json:"OIIEELPECGL,omitempty"`
	ChessRogueMainStoryId uint32   `protobuf:"varint,7,opt,name=chess_rogue_main_story_id,json=chessRogueMainStoryId,proto3" json:"chess_rogue_main_story_id,omitempty"`
	ALNMBFMEIAC           bool     `protobuf:"varint,15,opt,name=ALNMBFMEIAC,proto3" json:"ALNMBFMEIAC,omitempty"`
	NJJKIOHDOMP           bool     `protobuf:"varint,8,opt,name=NJJKIOHDOMP,proto3" json:"NJJKIOHDOMP,omitempty"`
	NDJJNLKCNHM           []uint32 `protobuf:"varint,2,rep,packed,name=NDJJNLKCNHM,proto3" json:"NDJJNLKCNHM,omitempty"`
	ChessRogueSubStoryId  uint32   `protobuf:"varint,11,opt,name=chess_rogue_sub_story_id,json=chessRogueSubStoryId,proto3" json:"chess_rogue_sub_story_id,omitempty"`
	PCOMKOLMBBG           []uint32 `protobuf:"varint,3,rep,packed,name=PCOMKOLMBBG,proto3" json:"PCOMKOLMBBG,omitempty"`
	NIMEABELKEH           []uint32 `protobuf:"varint,6,rep,packed,name=NIMEABELKEH,proto3" json:"NIMEABELKEH,omitempty"`
}

func (x *JPFECHLHHEN) Reset() {
	*x = JPFECHLHHEN{}
	if protoimpl.UnsafeEnabled {
		mi := &file_JPFECHLHHEN_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JPFECHLHHEN) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JPFECHLHHEN) ProtoMessage() {}

func (x *JPFECHLHHEN) ProtoReflect() protoreflect.Message {
	mi := &file_JPFECHLHHEN_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JPFECHLHHEN.ProtoReflect.Descriptor instead.
func (*JPFECHLHHEN) Descriptor() ([]byte, []int) {
	return file_JPFECHLHHEN_proto_rawDescGZIP(), []int{0}
}

func (x *JPFECHLHHEN) GetOIIEELPECGL() uint32 {
	if x != nil {
		return x.OIIEELPECGL
	}
	return 0
}

func (x *JPFECHLHHEN) GetChessRogueMainStoryId() uint32 {
	if x != nil {
		return x.ChessRogueMainStoryId
	}
	return 0
}

func (x *JPFECHLHHEN) GetALNMBFMEIAC() bool {
	if x != nil {
		return x.ALNMBFMEIAC
	}
	return false
}

func (x *JPFECHLHHEN) GetNJJKIOHDOMP() bool {
	if x != nil {
		return x.NJJKIOHDOMP
	}
	return false
}

func (x *JPFECHLHHEN) GetNDJJNLKCNHM() []uint32 {
	if x != nil {
		return x.NDJJNLKCNHM
	}
	return nil
}

func (x *JPFECHLHHEN) GetChessRogueSubStoryId() uint32 {
	if x != nil {
		return x.ChessRogueSubStoryId
	}
	return 0
}

func (x *JPFECHLHHEN) GetPCOMKOLMBBG() []uint32 {
	if x != nil {
		return x.PCOMKOLMBBG
	}
	return nil
}

func (x *JPFECHLHHEN) GetNIMEABELKEH() []uint32 {
	if x != nil {
		return x.NIMEABELKEH
	}
	return nil
}

var File_JPFECHLHHEN_proto protoreflect.FileDescriptor

var file_JPFECHLHHEN_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4a, 0x50, 0x46, 0x45, 0x43, 0x48, 0x4c, 0x48, 0x48, 0x45, 0x4e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xcb, 0x02, 0x0a, 0x0b, 0x4a, 0x50, 0x46, 0x45, 0x43, 0x48, 0x4c, 0x48,
	0x48, 0x45, 0x4e, 0x12, 0x20, 0x0a, 0x0b, 0x4f, 0x49, 0x49, 0x45, 0x45, 0x4c, 0x50, 0x45, 0x43,
	0x47, 0x4c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x4f, 0x49, 0x49, 0x45, 0x45, 0x4c,
	0x50, 0x45, 0x43, 0x47, 0x4c, 0x12, 0x38, 0x0a, 0x19, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x72,
	0x6f, 0x67, 0x75, 0x65, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x63, 0x68, 0x65, 0x73, 0x73, 0x52,
	0x6f, 0x67, 0x75, 0x65, 0x4d, 0x61, 0x69, 0x6e, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x41, 0x4c, 0x4e, 0x4d, 0x42, 0x46, 0x4d, 0x45, 0x49, 0x41, 0x43, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x41, 0x4c, 0x4e, 0x4d, 0x42, 0x46, 0x4d, 0x45, 0x49, 0x41,
	0x43, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x4a, 0x4a, 0x4b, 0x49, 0x4f, 0x48, 0x44, 0x4f, 0x4d, 0x50,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x4e, 0x4a, 0x4a, 0x4b, 0x49, 0x4f, 0x48, 0x44,
	0x4f, 0x4d, 0x50, 0x12, 0x20, 0x0a, 0x0b, 0x4e, 0x44, 0x4a, 0x4a, 0x4e, 0x4c, 0x4b, 0x43, 0x4e,
	0x48, 0x4d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x44, 0x4a, 0x4a, 0x4e, 0x4c,
	0x4b, 0x43, 0x4e, 0x48, 0x4d, 0x12, 0x36, 0x0a, 0x18, 0x63, 0x68, 0x65, 0x73, 0x73, 0x5f, 0x72,
	0x6f, 0x67, 0x75, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x63, 0x68, 0x65, 0x73, 0x73, 0x52, 0x6f,
	0x67, 0x75, 0x65, 0x53, 0x75, 0x62, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x50, 0x43, 0x4f, 0x4d, 0x4b, 0x4f, 0x4c, 0x4d, 0x42, 0x42, 0x47, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x0b, 0x50, 0x43, 0x4f, 0x4d, 0x4b, 0x4f, 0x4c, 0x4d, 0x42, 0x42, 0x47, 0x12,
	0x20, 0x0a, 0x0b, 0x4e, 0x49, 0x4d, 0x45, 0x41, 0x42, 0x45, 0x4c, 0x4b, 0x45, 0x48, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0d, 0x52, 0x0b, 0x4e, 0x49, 0x4d, 0x45, 0x41, 0x42, 0x45, 0x4c, 0x4b, 0x45,
	0x48, 0x42, 0x28, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x1b,
	0x45, 0x67, 0x67, 0x4c, 0x69, 0x6e, 0x6b, 0x2e, 0x44, 0x61, 0x6e, 0x68, 0x65, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_JPFECHLHHEN_proto_rawDescOnce sync.Once
	file_JPFECHLHHEN_proto_rawDescData = file_JPFECHLHHEN_proto_rawDesc
)

func file_JPFECHLHHEN_proto_rawDescGZIP() []byte {
	file_JPFECHLHHEN_proto_rawDescOnce.Do(func() {
		file_JPFECHLHHEN_proto_rawDescData = protoimpl.X.CompressGZIP(file_JPFECHLHHEN_proto_rawDescData)
	})
	return file_JPFECHLHHEN_proto_rawDescData
}

var file_JPFECHLHHEN_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_JPFECHLHHEN_proto_goTypes = []interface{}{
	(*JPFECHLHHEN)(nil), // 0: JPFECHLHHEN
}
var file_JPFECHLHHEN_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_JPFECHLHHEN_proto_init() }
func file_JPFECHLHHEN_proto_init() {
	if File_JPFECHLHHEN_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_JPFECHLHHEN_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JPFECHLHHEN); i {
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
			RawDescriptor: file_JPFECHLHHEN_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_JPFECHLHHEN_proto_goTypes,
		DependencyIndexes: file_JPFECHLHHEN_proto_depIdxs,
		MessageInfos:      file_JPFECHLHHEN_proto_msgTypes,
	}.Build()
	File_JPFECHLHHEN_proto = out.File
	file_JPFECHLHHEN_proto_rawDesc = nil
	file_JPFECHLHHEN_proto_goTypes = nil
	file_JPFECHLHHEN_proto_depIdxs = nil
}
