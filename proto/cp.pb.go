// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.1
// source: proto/cp.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type PasteOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Length int64 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *PasteOption) Reset() {
	*x = PasteOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PasteOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PasteOption) ProtoMessage() {}

func (x *PasteOption) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PasteOption.ProtoReflect.Descriptor instead.
func (*PasteOption) Descriptor() ([]byte, []int) {
	return file_proto_cp_proto_rawDescGZIP(), []int{0}
}

func (x *PasteOption) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

type Stream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*Stream_Info_
	//	*Stream_Chunk
	//	*Stream_Hash
	Data isStream_Data `protobuf_oneof:"data"`
}

func (x *Stream) Reset() {
	*x = Stream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stream) ProtoMessage() {}

func (x *Stream) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stream.ProtoReflect.Descriptor instead.
func (*Stream) Descriptor() ([]byte, []int) {
	return file_proto_cp_proto_rawDescGZIP(), []int{1}
}

func (m *Stream) GetData() isStream_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Stream) GetInfo() *Stream_Info {
	if x, ok := x.GetData().(*Stream_Info_); ok {
		return x.Info
	}
	return nil
}

func (x *Stream) GetChunk() []byte {
	if x, ok := x.GetData().(*Stream_Chunk); ok {
		return x.Chunk
	}
	return nil
}

func (x *Stream) GetHash() []byte {
	if x, ok := x.GetData().(*Stream_Hash); ok {
		return x.Hash
	}
	return nil
}

type isStream_Data interface {
	isStream_Data()
}

type Stream_Info_ struct {
	Info *Stream_Info `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type Stream_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

type Stream_Hash struct {
	Hash []byte `protobuf:"bytes,3,opt,name=hash,proto3,oneof"`
}

func (*Stream_Info_) isStream_Data() {}

func (*Stream_Chunk) isStream_Data() {}

func (*Stream_Hash) isStream_Data() {}

type Stream_Info struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ttl int64 `protobuf:"varint,1,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *Stream_Info) Reset() {
	*x = Stream_Info{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_cp_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stream_Info) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stream_Info) ProtoMessage() {}

func (x *Stream_Info) ProtoReflect() protoreflect.Message {
	mi := &file_proto_cp_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stream_Info.ProtoReflect.Descriptor instead.
func (*Stream_Info) Descriptor() ([]byte, []int) {
	return file_proto_cp_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Stream_Info) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

var File_proto_cp_proto protoreflect.FileDescriptor

var file_proto_cp_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0b, 0x50, 0x61, 0x73, 0x74, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x82, 0x01, 0x0a, 0x06,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x28, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f,
	0x12, 0x16, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x14, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x1a, 0x18,
	0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x32, 0x71, 0x0a, 0x0c, 0x57, 0x65, 0x62, 0x43, 0x6c, 0x69, 0x70, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x12, 0x31, 0x0a, 0x04, 0x43, 0x6f, 0x70, 0x79, 0x12, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x28, 0x01, 0x12, 0x2e, 0x0a, 0x05, 0x50, 0x61, 0x73, 0x74, 0x65, 0x12, 0x12, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x73, 0x74, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22,
	0x00, 0x30, 0x01, 0x42, 0x24, 0x5a, 0x22, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x69, 0x67, 0x65, 0x72, 0x77, 0x69, 0x6c, 0x6c, 0x39, 0x30, 0x2f, 0x77, 0x65,
	0x62, 0x63, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_cp_proto_rawDescOnce sync.Once
	file_proto_cp_proto_rawDescData = file_proto_cp_proto_rawDesc
)

func file_proto_cp_proto_rawDescGZIP() []byte {
	file_proto_cp_proto_rawDescOnce.Do(func() {
		file_proto_cp_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_cp_proto_rawDescData)
	})
	return file_proto_cp_proto_rawDescData
}

var file_proto_cp_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_cp_proto_goTypes = []interface{}{
	(*PasteOption)(nil),   // 0: proto.PasteOption
	(*Stream)(nil),        // 1: proto.Stream
	(*Stream_Info)(nil),   // 2: proto.Stream.Info
	(*emptypb.Empty)(nil), // 3: google.protobuf.Empty
}
var file_proto_cp_proto_depIdxs = []int32{
	2, // 0: proto.Stream.info:type_name -> proto.Stream.Info
	1, // 1: proto.WebClipboard.Copy:input_type -> proto.Stream
	0, // 2: proto.WebClipboard.Paste:input_type -> proto.PasteOption
	3, // 3: proto.WebClipboard.Copy:output_type -> google.protobuf.Empty
	1, // 4: proto.WebClipboard.Paste:output_type -> proto.Stream
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_cp_proto_init() }
func file_proto_cp_proto_init() {
	if File_proto_cp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_cp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PasteOption); i {
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
		file_proto_cp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stream); i {
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
		file_proto_cp_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stream_Info); i {
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
	file_proto_cp_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Stream_Info_)(nil),
		(*Stream_Chunk)(nil),
		(*Stream_Hash)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_cp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cp_proto_goTypes,
		DependencyIndexes: file_proto_cp_proto_depIdxs,
		MessageInfos:      file_proto_cp_proto_msgTypes,
	}.Build()
	File_proto_cp_proto = out.File
	file_proto_cp_proto_rawDesc = nil
	file_proto_cp_proto_goTypes = nil
	file_proto_cp_proto_depIdxs = nil
}
