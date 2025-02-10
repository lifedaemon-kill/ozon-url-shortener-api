// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.2
// 	protoc        v5.29.1
// source: url_shortener/url_shortener.proto

package url_shortener_api

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

type SaveURLRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Source_URL    string                 `protobuf:"bytes,1,opt,name=source_URL,json=sourceURL,proto3" json:"source_URL,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SaveURLRequest) Reset() {
	*x = SaveURLRequest{}
	mi := &file_url_shortener_url_shortener_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveURLRequest) ProtoMessage() {}

func (x *SaveURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_shortener_url_shortener_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveURLRequest.ProtoReflect.Descriptor instead.
func (*SaveURLRequest) Descriptor() ([]byte, []int) {
	return file_url_shortener_url_shortener_proto_rawDescGZIP(), []int{0}
}

func (x *SaveURLRequest) GetSource_URL() string {
	if x != nil {
		return x.Source_URL
	}
	return ""
}

type SaveURLResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Alias_URL     string                 `protobuf:"bytes,2,opt,name=alias_URL,json=aliasURL,proto3" json:"alias_URL,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SaveURLResponse) Reset() {
	*x = SaveURLResponse{}
	mi := &file_url_shortener_url_shortener_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SaveURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveURLResponse) ProtoMessage() {}

func (x *SaveURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_url_shortener_url_shortener_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveURLResponse.ProtoReflect.Descriptor instead.
func (*SaveURLResponse) Descriptor() ([]byte, []int) {
	return file_url_shortener_url_shortener_proto_rawDescGZIP(), []int{1}
}

func (x *SaveURLResponse) GetAlias_URL() string {
	if x != nil {
		return x.Alias_URL
	}
	return ""
}

type FetchURLRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Alias_URL     string                 `protobuf:"bytes,1,opt,name=alias_URL,json=aliasURL,proto3" json:"alias_URL,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchURLRequest) Reset() {
	*x = FetchURLRequest{}
	mi := &file_url_shortener_url_shortener_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchURLRequest) ProtoMessage() {}

func (x *FetchURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_url_shortener_url_shortener_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchURLRequest.ProtoReflect.Descriptor instead.
func (*FetchURLRequest) Descriptor() ([]byte, []int) {
	return file_url_shortener_url_shortener_proto_rawDescGZIP(), []int{2}
}

func (x *FetchURLRequest) GetAlias_URL() string {
	if x != nil {
		return x.Alias_URL
	}
	return ""
}

type FetchURLResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Source_URL    string                 `protobuf:"bytes,1,opt,name=source_URL,json=sourceURL,proto3" json:"source_URL,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FetchURLResponse) Reset() {
	*x = FetchURLResponse{}
	mi := &file_url_shortener_url_shortener_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FetchURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchURLResponse) ProtoMessage() {}

func (x *FetchURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_url_shortener_url_shortener_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchURLResponse.ProtoReflect.Descriptor instead.
func (*FetchURLResponse) Descriptor() ([]byte, []int) {
	return file_url_shortener_url_shortener_proto_rawDescGZIP(), []int{3}
}

func (x *FetchURLResponse) GetSource_URL() string {
	if x != nil {
		return x.Source_URL
	}
	return ""
}

var File_url_shortener_url_shortener_proto protoreflect.FileDescriptor

var file_url_shortener_url_shortener_proto_rawDesc = []byte{
	0x0a, 0x21, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2f,
	0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x22, 0x2f, 0x0a, 0x0e, 0x53, 0x61, 0x76, 0x65, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x55,
	0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x55, 0x52, 0x4c, 0x22, 0x2e, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x5f,
	0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x55, 0x52, 0x4c, 0x22, 0x2e, 0x0a, 0x0f, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x5f,
	0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x6c, 0x69, 0x61, 0x73,
	0x55, 0x52, 0x4c, 0x22, 0x31, 0x0a, 0x10, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x55, 0x52, 0x4c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x32, 0xa3, 0x01, 0x0a, 0x0a, 0x55, 0x52, 0x4c, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x53, 0x61, 0x76, 0x65, 0x55, 0x52, 0x4c,
	0x12, 0x1d, 0x2e, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x2e, 0x53, 0x61, 0x76, 0x65, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e,
	0x53, 0x61, 0x76, 0x65, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4b, 0x0a, 0x08, 0x46, 0x65, 0x74, 0x63, 0x68, 0x55, 0x52, 0x4c, 0x12, 0x1e, 0x2e, 0x75, 0x72,
	0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x75, 0x72,
	0x6c, 0x5f, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x46, 0x65, 0x74, 0x63,
	0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b,
	0x6f, 0x7a, 0x6f, 0x6e, 0x2d, 0x6c, 0x66, 0x64, 0x2d, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x3b, 0x75, 0x72, 0x6c, 0x5f, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_url_shortener_url_shortener_proto_rawDescOnce sync.Once
	file_url_shortener_url_shortener_proto_rawDescData = file_url_shortener_url_shortener_proto_rawDesc
)

func file_url_shortener_url_shortener_proto_rawDescGZIP() []byte {
	file_url_shortener_url_shortener_proto_rawDescOnce.Do(func() {
		file_url_shortener_url_shortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_url_shortener_url_shortener_proto_rawDescData)
	})
	return file_url_shortener_url_shortener_proto_rawDescData
}

var file_url_shortener_url_shortener_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_url_shortener_url_shortener_proto_goTypes = []any{
	(*SaveURLRequest)(nil),   // 0: url_shortener.SaveURLRequest
	(*SaveURLResponse)(nil),  // 1: url_shortener.SaveURLResponse
	(*FetchURLRequest)(nil),  // 2: url_shortener.FetchURLRequest
	(*FetchURLResponse)(nil), // 3: url_shortener.FetchURLResponse
}
var file_url_shortener_url_shortener_proto_depIdxs = []int32{
	0, // 0: url_shortener.URLService.SaveURL:input_type -> url_shortener.SaveURLRequest
	2, // 1: url_shortener.URLService.FetchURL:input_type -> url_shortener.FetchURLRequest
	1, // 2: url_shortener.URLService.SaveURL:output_type -> url_shortener.SaveURLResponse
	3, // 3: url_shortener.URLService.FetchURL:output_type -> url_shortener.FetchURLResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_url_shortener_url_shortener_proto_init() }
func file_url_shortener_url_shortener_proto_init() {
	if File_url_shortener_url_shortener_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_url_shortener_url_shortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_url_shortener_url_shortener_proto_goTypes,
		DependencyIndexes: file_url_shortener_url_shortener_proto_depIdxs,
		MessageInfos:      file_url_shortener_url_shortener_proto_msgTypes,
	}.Build()
	File_url_shortener_url_shortener_proto = out.File
	file_url_shortener_url_shortener_proto_rawDesc = nil
	file_url_shortener_url_shortener_proto_goTypes = nil
	file_url_shortener_url_shortener_proto_depIdxs = nil
}
