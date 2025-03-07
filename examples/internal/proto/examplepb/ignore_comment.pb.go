// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: examples/internal/proto/examplepb/ignore_comment.proto

package examplepb

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// This comment should be excluded from OpenAPI output
type FooRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This comment should be excluded from OpenAPI output
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	// This comment should be excluded from OpenAPI output
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *FooRequest) Reset() {
	*x = FooRequest{}
	mi := &file_examples_internal_proto_examplepb_ignore_comment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FooRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooRequest) ProtoMessage() {}

func (x *FooRequest) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_ignore_comment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooRequest.ProtoReflect.Descriptor instead.
func (*FooRequest) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_ignore_comment_proto_rawDescGZIP(), []int{0}
}

func (x *FooRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *FooRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// This comment should be excluded from OpenAPI output
type FooReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FooReply) Reset() {
	*x = FooReply{}
	mi := &file_examples_internal_proto_examplepb_ignore_comment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FooReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooReply) ProtoMessage() {}

func (x *FooReply) ProtoReflect() protoreflect.Message {
	mi := &file_examples_internal_proto_examplepb_ignore_comment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooReply.ProtoReflect.Descriptor instead.
func (*FooReply) Descriptor() ([]byte, []int) {
	return file_examples_internal_proto_examplepb_ignore_comment_proto_rawDescGZIP(), []int{1}
}

var File_examples_internal_proto_examplepb_ignore_comment_proto protoreflect.FileDescriptor

var file_examples_internal_proto_examplepb_ignore_comment_proto_rawDesc = []byte{
	0x0a, 0x36, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x70, 0x62, 0x2f, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x0a, 0x46, 0x6f, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x44, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0x92, 0x41, 0x25, 0x32, 0x23, 0x54, 0x68,
	0x69, 0x73, 0x20, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x73, 0x68,
	0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x20, 0x70, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x64, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x44, 0x0a, 0x08, 0x70,
	0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0x92,
	0x41, 0x25, 0x2a, 0x23, 0x54, 0x68, 0x69, 0x73, 0x20, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x20, 0x70, 0x72,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x0a, 0x0a, 0x08, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xa6, 0x01,
	0x0a, 0x0a, 0x46, 0x6f, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x97, 0x01, 0x0a,
	0x03, 0x46, 0x6f, 0x6f, 0x12, 0x3a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x70, 0x62, 0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x38, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x70,
	0x62, 0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x66, 0x6f, 0x6f, 0x42, 0x4d, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x76, 0x32, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_examples_internal_proto_examplepb_ignore_comment_proto_rawDescOnce sync.Once
	file_examples_internal_proto_examplepb_ignore_comment_proto_rawDescData = file_examples_internal_proto_examplepb_ignore_comment_proto_rawDesc
)

func file_examples_internal_proto_examplepb_ignore_comment_proto_rawDescGZIP() []byte {
	file_examples_internal_proto_examplepb_ignore_comment_proto_rawDescOnce.Do(func() {
		file_examples_internal_proto_examplepb_ignore_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_examples_internal_proto_examplepb_ignore_comment_proto_rawDescData)
	})
	return file_examples_internal_proto_examplepb_ignore_comment_proto_rawDescData
}

var file_examples_internal_proto_examplepb_ignore_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_examples_internal_proto_examplepb_ignore_comment_proto_goTypes = []any{
	(*FooRequest)(nil), // 0: grpc.gateway.examples.internal.proto.examplepb.FooRequest
	(*FooReply)(nil),   // 1: grpc.gateway.examples.internal.proto.examplepb.FooReply
}
var file_examples_internal_proto_examplepb_ignore_comment_proto_depIdxs = []int32{
	0, // 0: grpc.gateway.examples.internal.proto.examplepb.FooService.Foo:input_type -> grpc.gateway.examples.internal.proto.examplepb.FooRequest
	1, // 1: grpc.gateway.examples.internal.proto.examplepb.FooService.Foo:output_type -> grpc.gateway.examples.internal.proto.examplepb.FooReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_examples_internal_proto_examplepb_ignore_comment_proto_init() }
func file_examples_internal_proto_examplepb_ignore_comment_proto_init() {
	if File_examples_internal_proto_examplepb_ignore_comment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_examples_internal_proto_examplepb_ignore_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_examples_internal_proto_examplepb_ignore_comment_proto_goTypes,
		DependencyIndexes: file_examples_internal_proto_examplepb_ignore_comment_proto_depIdxs,
		MessageInfos:      file_examples_internal_proto_examplepb_ignore_comment_proto_msgTypes,
	}.Build()
	File_examples_internal_proto_examplepb_ignore_comment_proto = out.File
	file_examples_internal_proto_examplepb_ignore_comment_proto_rawDesc = nil
	file_examples_internal_proto_examplepb_ignore_comment_proto_goTypes = nil
	file_examples_internal_proto_examplepb_ignore_comment_proto_depIdxs = nil
}
