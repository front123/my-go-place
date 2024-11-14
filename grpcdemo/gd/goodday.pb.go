// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: grpcdemo/gd/goodday.proto

package gd

import (
	hw "go-example/grpcdemo/hw"
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

type DayOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32        `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Req   *hw.HelloReq `protobuf:"bytes,2,opt,name=req,proto3" json:"req,omitempty"`
}

func (x *DayOperation) Reset() {
	*x = DayOperation{}
	mi := &file_grpcdemo_gd_goodday_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DayOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DayOperation) ProtoMessage() {}

func (x *DayOperation) ProtoReflect() protoreflect.Message {
	mi := &file_grpcdemo_gd_goodday_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DayOperation.ProtoReflect.Descriptor instead.
func (*DayOperation) Descriptor() ([]byte, []int) {
	return file_grpcdemo_gd_goodday_proto_rawDescGZIP(), []int{0}
}

func (x *DayOperation) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *DayOperation) GetReq() *hw.HelloReq {
	if x != nil {
		return x.Req
	}
	return nil
}

type SimpleMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SimpleMsg) Reset() {
	*x = SimpleMsg{}
	mi := &file_grpcdemo_gd_goodday_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SimpleMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SimpleMsg) ProtoMessage() {}

func (x *SimpleMsg) ProtoReflect() protoreflect.Message {
	mi := &file_grpcdemo_gd_goodday_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SimpleMsg.ProtoReflect.Descriptor instead.
func (*SimpleMsg) Descriptor() ([]byte, []int) {
	return file_grpcdemo_gd_goodday_proto_rawDescGZIP(), []int{1}
}

func (x *SimpleMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_grpcdemo_gd_goodday_proto protoreflect.FileDescriptor

var file_grpcdemo_gd_goodday_proto_rawDesc = []byte{
	0x0a, 0x19, 0x67, 0x72, 0x70, 0x63, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x67, 0x64, 0x2f, 0x67, 0x6f,
	0x6f, 0x64, 0x64, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x64, 0x70,
	0x6b, 0x67, 0x1a, 0x1c, 0x67, 0x72, 0x70, 0x63, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x68, 0x77, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x47, 0x0a, 0x0c, 0x44, 0x61, 0x79, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x21, 0x0a, 0x03, 0x72, 0x65, 0x71, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x68, 0x77, 0x70, 0x6b, 0x67, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x52, 0x03, 0x72, 0x65, 0x71, 0x22, 0x1d, 0x0a, 0x09, 0x53, 0x69, 0x6d,
	0x70, 0x6c, 0x65, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x6f, 0x2d, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x64, 0x65, 0x6d, 0x6f, 0x2f,
	0x67, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcdemo_gd_goodday_proto_rawDescOnce sync.Once
	file_grpcdemo_gd_goodday_proto_rawDescData = file_grpcdemo_gd_goodday_proto_rawDesc
)

func file_grpcdemo_gd_goodday_proto_rawDescGZIP() []byte {
	file_grpcdemo_gd_goodday_proto_rawDescOnce.Do(func() {
		file_grpcdemo_gd_goodday_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcdemo_gd_goodday_proto_rawDescData)
	})
	return file_grpcdemo_gd_goodday_proto_rawDescData
}

var file_grpcdemo_gd_goodday_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpcdemo_gd_goodday_proto_goTypes = []any{
	(*DayOperation)(nil), // 0: gdpkg.DayOperation
	(*SimpleMsg)(nil),    // 1: gdpkg.SimpleMsg
	(*hw.HelloReq)(nil),  // 2: hwpkg.HelloReq
}
var file_grpcdemo_gd_goodday_proto_depIdxs = []int32{
	2, // 0: gdpkg.DayOperation.req:type_name -> hwpkg.HelloReq
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpcdemo_gd_goodday_proto_init() }
func file_grpcdemo_gd_goodday_proto_init() {
	if File_grpcdemo_gd_goodday_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpcdemo_gd_goodday_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpcdemo_gd_goodday_proto_goTypes,
		DependencyIndexes: file_grpcdemo_gd_goodday_proto_depIdxs,
		MessageInfos:      file_grpcdemo_gd_goodday_proto_msgTypes,
	}.Build()
	File_grpcdemo_gd_goodday_proto = out.File
	file_grpcdemo_gd_goodday_proto_rawDesc = nil
	file_grpcdemo_gd_goodday_proto_goTypes = nil
	file_grpcdemo_gd_goodday_proto_depIdxs = nil
}
