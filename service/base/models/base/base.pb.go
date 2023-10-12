// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.4
// source: volcengine/base/base.proto

package base

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResponseMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string         `protobuf:"bytes,1,opt,name=RequestId,proto3" json:"RequestId,omitempty"` // 请求ID
	Action    string         `protobuf:"bytes,2,opt,name=Action,proto3" json:"Action,omitempty"`       // 请求接口
	Version   string         `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version,omitempty"`     // 版本
	Service   string         `protobuf:"bytes,4,opt,name=Service,proto3" json:"Service,omitempty"`     // 服务
	Region    string         `protobuf:"bytes,5,opt,name=Region,proto3" json:"Region,omitempty"`       // 区域
	Error     *ResponseError `protobuf:"bytes,6,opt,name=Error,proto3" json:"Error,omitempty"`         // 异常信息
}

func (x *ResponseMetadata) Reset() {
	*x = ResponseMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volcengine_base_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMetadata) ProtoMessage() {}

func (x *ResponseMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_volcengine_base_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMetadata.ProtoReflect.Descriptor instead.
func (*ResponseMetadata) Descriptor() ([]byte, []int) {
	return file_volcengine_base_base_proto_rawDescGZIP(), []int{0}
}

func (x *ResponseMetadata) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ResponseMetadata) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *ResponseMetadata) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ResponseMetadata) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ResponseMetadata) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *ResponseMetadata) GetError() *ResponseError {
	if x != nil {
		return x.Error
	}
	return nil
}

type ResponseError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`       // 错误码
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"` // 详细错误信息
}

func (x *ResponseError) Reset() {
	*x = ResponseError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volcengine_base_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseError) ProtoMessage() {}

func (x *ResponseError) ProtoReflect() protoreflect.Message {
	mi := &file_volcengine_base_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseError.ProtoReflect.Descriptor instead.
func (*ResponseError) Descriptor() ([]byte, []int) {
	return file_volcengine_base_base_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseError) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ResponseError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_volcengine_base_base_proto protoreflect.FileDescriptor

var file_volcengine_base_base_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x56, 0x6f,
	0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x01, 0x0a, 0x10,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x3d, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0xc0, 0x01, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x6f, 0x6c, 0x63,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x42, 0x04,
	0x42, 0x61, 0x73, 0x65, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x6f,
	0x6c, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0xa0, 0x01, 0x01, 0xd8, 0x01, 0x01, 0xc2, 0x02, 0x00, 0xca,
	0x02, 0x1d, 0x56, 0x6f, 0x6c, 0x63, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x42,
	0x61, 0x73, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x42, 0x61, 0x73, 0x65, 0xe2,
	0x02, 0x24, 0x56, 0x6f, 0x6c, 0x63, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x42,
	0x61, 0x73, 0x65, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_volcengine_base_base_proto_rawDescOnce sync.Once
	file_volcengine_base_base_proto_rawDescData = file_volcengine_base_base_proto_rawDesc
)

func file_volcengine_base_base_proto_rawDescGZIP() []byte {
	file_volcengine_base_base_proto_rawDescOnce.Do(func() {
		file_volcengine_base_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_volcengine_base_base_proto_rawDescData)
	})
	return file_volcengine_base_base_proto_rawDescData
}

var file_volcengine_base_base_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_volcengine_base_base_proto_goTypes = []interface{}{
	(*ResponseMetadata)(nil), // 0: Volcengine.Base.Models.Base.ResponseMetadata
	(*ResponseError)(nil),    // 1: Volcengine.Base.Models.Base.ResponseError
}
var file_volcengine_base_base_proto_depIdxs = []int32{
	1, // 0: Volcengine.Base.Models.Base.ResponseMetadata.Error:type_name -> Volcengine.Base.Models.Base.ResponseError
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_volcengine_base_base_proto_init() }
func file_volcengine_base_base_proto_init() {
	if File_volcengine_base_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_volcengine_base_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMetadata); i {
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
		file_volcengine_base_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseError); i {
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
			RawDescriptor: file_volcengine_base_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_volcengine_base_base_proto_goTypes,
		DependencyIndexes: file_volcengine_base_base_proto_depIdxs,
		MessageInfos:      file_volcengine_base_base_proto_msgTypes,
	}.Build()
	File_volcengine_base_base_proto = out.File
	file_volcengine_base_base_proto_rawDesc = nil
	file_volcengine_base_base_proto_goTypes = nil
	file_volcengine_base_base_proto_depIdxs = nil
}
