// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1-devel
// 	protoc        v3.19.4
// source: volcengine/vod/business/vod_smart_strategy.proto

package business

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VodSmartStrategyResponseStreamType int32

const (
	VodSmartStrategyResponseStreamType_UndefinedVodSmartStrategyResponseStreamType            VodSmartStrategyResponseStreamType = 0 // 未定义的（保留置空）
	VodSmartStrategyResponseStreamType_OriginalStreamVodSmartStrategyResponseStreamType       VodSmartStrategyResponseStreamType = 1 // 返回原片流
	VodSmartStrategyResponseStreamType_StrategyEncodeStreamVodSmartStrategyResponseStreamType VodSmartStrategyResponseStreamType = 2 // 返回匹配的策略转码流
	VodSmartStrategyResponseStreamType_FallbackEncodeStreamVodSmartStrategyResponseStreamType VodSmartStrategyResponseStreamType = 3 // 返回降级的转码流
)

// Enum value maps for VodSmartStrategyResponseStreamType.
var (
	VodSmartStrategyResponseStreamType_name = map[int32]string{
		0: "UndefinedVodSmartStrategyResponseStreamType",
		1: "OriginalStreamVodSmartStrategyResponseStreamType",
		2: "StrategyEncodeStreamVodSmartStrategyResponseStreamType",
		3: "FallbackEncodeStreamVodSmartStrategyResponseStreamType",
	}
	VodSmartStrategyResponseStreamType_value = map[string]int32{
		"UndefinedVodSmartStrategyResponseStreamType":            0,
		"OriginalStreamVodSmartStrategyResponseStreamType":       1,
		"StrategyEncodeStreamVodSmartStrategyResponseStreamType": 2,
		"FallbackEncodeStreamVodSmartStrategyResponseStreamType": 3,
	}
)

func (x VodSmartStrategyResponseStreamType) Enum() *VodSmartStrategyResponseStreamType {
	p := new(VodSmartStrategyResponseStreamType)
	*p = x
	return p
}

func (x VodSmartStrategyResponseStreamType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VodSmartStrategyResponseStreamType) Descriptor() protoreflect.EnumDescriptor {
	return file_volcengine_vod_business_vod_smart_strategy_proto_enumTypes[0].Descriptor()
}

func (VodSmartStrategyResponseStreamType) Type() protoreflect.EnumType {
	return &file_volcengine_vod_business_vod_smart_strategy_proto_enumTypes[0]
}

func (x VodSmartStrategyResponseStreamType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VodSmartStrategyResponseStreamType.Descriptor instead.
func (VodSmartStrategyResponseStreamType) EnumDescriptor() ([]byte, []int) {
	return file_volcengine_vod_business_vod_smart_strategy_proto_rawDescGZIP(), []int{0}
}

type VodSmartStrategyDataStoreStatus int32

const (
	VodSmartStrategyDataStoreStatus_UndefinedVodSmartStrategyDataStoreStatus         VodSmartStrategyDataStoreStatus = 0 // 未定义的
	VodSmartStrategyDataStoreStatus_NotExistVodSmartStrategyDataStoreStatus          VodSmartStrategyDataStoreStatus = 1 // VDA中不存在
	VodSmartStrategyDataStoreStatus_HasOriginalStreamVodSmartStrategyDataStoreStatus VodSmartStrategyDataStoreStatus = 2 // VDA中存在源片
	VodSmartStrategyDataStoreStatus_HasEncodeStreamVodSmartStrategyDataStoreStatus   VodSmartStrategyDataStoreStatus = 3 // VDA中存在转码流
)

// Enum value maps for VodSmartStrategyDataStoreStatus.
var (
	VodSmartStrategyDataStoreStatus_name = map[int32]string{
		0: "UndefinedVodSmartStrategyDataStoreStatus",
		1: "NotExistVodSmartStrategyDataStoreStatus",
		2: "HasOriginalStreamVodSmartStrategyDataStoreStatus",
		3: "HasEncodeStreamVodSmartStrategyDataStoreStatus",
	}
	VodSmartStrategyDataStoreStatus_value = map[string]int32{
		"UndefinedVodSmartStrategyDataStoreStatus":         0,
		"NotExistVodSmartStrategyDataStoreStatus":          1,
		"HasOriginalStreamVodSmartStrategyDataStoreStatus": 2,
		"HasEncodeStreamVodSmartStrategyDataStoreStatus":   3,
	}
)

func (x VodSmartStrategyDataStoreStatus) Enum() *VodSmartStrategyDataStoreStatus {
	p := new(VodSmartStrategyDataStoreStatus)
	*p = x
	return p
}

func (x VodSmartStrategyDataStoreStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VodSmartStrategyDataStoreStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_volcengine_vod_business_vod_smart_strategy_proto_enumTypes[1].Descriptor()
}

func (VodSmartStrategyDataStoreStatus) Type() protoreflect.EnumType {
	return &file_volcengine_vod_business_vod_smart_strategy_proto_enumTypes[1]
}

func (x VodSmartStrategyDataStoreStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VodSmartStrategyDataStoreStatus.Descriptor instead.
func (VodSmartStrategyDataStoreStatus) EnumDescriptor() ([]byte, []int) {
	return file_volcengine_vod_business_vod_smart_strategy_proto_rawDescGZIP(), []int{1}
}

type VodGetSmartStrategyLitePlayInfoResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StreamType      VodSmartStrategyResponseStreamType `protobuf:"varint,1,opt,name=StreamType,proto3,enum=Volcengine.Vod.Models.Business.VodSmartStrategyResponseStreamType" json:"StreamType,omitempty"`        // 当前返回的播放流类型
	DataStoreStatus VodSmartStrategyDataStoreStatus    `protobuf:"varint,2,opt,name=DataStoreStatus,proto3,enum=Volcengine.Vod.Models.Business.VodSmartStrategyDataStoreStatus" json:"DataStoreStatus,omitempty"` // 该视频在点播中的元信息状态
	PlayInfoModel   *VodPlayInfoModel                  `protobuf:"bytes,3,opt,name=PlayInfoModel,proto3" json:"PlayInfoModel,omitempty"`                                                                          // 匹配的VideoModel
	OriginalPlayUrl string                             `protobuf:"bytes,4,opt,name=OriginalPlayUrl,proto3" json:"OriginalPlayUrl,omitempty"`                                                                      // 原始URL
}

func (x *VodGetSmartStrategyLitePlayInfoResult) Reset() {
	*x = VodGetSmartStrategyLitePlayInfoResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_volcengine_vod_business_vod_smart_strategy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodGetSmartStrategyLitePlayInfoResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodGetSmartStrategyLitePlayInfoResult) ProtoMessage() {}

func (x *VodGetSmartStrategyLitePlayInfoResult) ProtoReflect() protoreflect.Message {
	mi := &file_volcengine_vod_business_vod_smart_strategy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodGetSmartStrategyLitePlayInfoResult.ProtoReflect.Descriptor instead.
func (*VodGetSmartStrategyLitePlayInfoResult) Descriptor() ([]byte, []int) {
	return file_volcengine_vod_business_vod_smart_strategy_proto_rawDescGZIP(), []int{0}
}

func (x *VodGetSmartStrategyLitePlayInfoResult) GetStreamType() VodSmartStrategyResponseStreamType {
	if x != nil {
		return x.StreamType
	}
	return VodSmartStrategyResponseStreamType_UndefinedVodSmartStrategyResponseStreamType
}

func (x *VodGetSmartStrategyLitePlayInfoResult) GetDataStoreStatus() VodSmartStrategyDataStoreStatus {
	if x != nil {
		return x.DataStoreStatus
	}
	return VodSmartStrategyDataStoreStatus_UndefinedVodSmartStrategyDataStoreStatus
}

func (x *VodGetSmartStrategyLitePlayInfoResult) GetPlayInfoModel() *VodPlayInfoModel {
	if x != nil {
		return x.PlayInfoModel
	}
	return nil
}

func (x *VodGetSmartStrategyLitePlayInfoResult) GetOriginalPlayUrl() string {
	if x != nil {
		return x.OriginalPlayUrl
	}
	return ""
}

var File_volcengine_vod_business_vod_smart_strategy_proto protoreflect.FileDescriptor

var file_volcengine_vod_business_vod_smart_strategy_proto_rawDesc = []byte{
	0x0a, 0x30, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x6f, 0x64,
	0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x6f, 0x64, 0x5f, 0x73, 0x6d,
	0x61, 0x72, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56,
	0x6f, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x1a, 0x28, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x6f, 0x64, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x6f, 0x64, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x02, 0x0a,
	0x25, 0x56, 0x6f, 0x64, 0x47, 0x65, 0x74, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x61,
	0x74, 0x65, 0x67, 0x79, 0x4c, 0x69, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x62, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x42, 0x2e, 0x56, 0x6f, 0x6c,
	0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x53,
	0x6d, 0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x69, 0x0a, 0x0f, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x3f, 0x2e, 0x56, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x56, 0x0a, 0x0d, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x56,
	0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x4d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2e, 0x56, 0x6f,
	0x64, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x0d,
	0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x28, 0x0a,
	0x0f, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x2a, 0x83, 0x02, 0x0a, 0x22, 0x56, 0x6f, 0x64, 0x53,
	0x6d, 0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2f,
	0x0a, 0x2b, 0x55, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x56, 0x6f, 0x64, 0x53, 0x6d,
	0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12,
	0x34, 0x0a, 0x30, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x56, 0x6f, 0x64, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54,
	0x79, 0x70, 0x65, 0x10, 0x01, 0x12, 0x3a, 0x0a, 0x36, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x6f, 0x64,
	0x53, 0x6d, 0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x10,
	0x02, 0x12, 0x3a, 0x0a, 0x36, 0x46, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x45, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x6f, 0x64, 0x53, 0x6d, 0x61, 0x72,
	0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x10, 0x03, 0x2a, 0xe6, 0x01,
	0x0a, 0x1f, 0x56, 0x6f, 0x64, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x67, 0x79, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x2c, 0x0a, 0x28, 0x55, 0x6e, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x56, 0x6f,
	0x64, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x10, 0x00, 0x12,
	0x2b, 0x0a, 0x27, 0x4e, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x56, 0x6f, 0x64, 0x53, 0x6d,
	0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x44, 0x61, 0x74, 0x61, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x10, 0x01, 0x12, 0x34, 0x0a, 0x30,
	0x48, 0x61, 0x73, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x56, 0x6f, 0x64, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x10, 0x02, 0x12, 0x32, 0x0a, 0x2e, 0x48, 0x61, 0x73, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x56, 0x6f, 0x64, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x53, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x67, 0x79, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x10, 0x03, 0x42, 0xd4, 0x01, 0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x76,
	0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x6f, 0x64, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x42, 0x10, 0x56, 0x6f, 0x64, 0x53, 0x6d, 0x61, 0x72, 0x74, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x6f, 0x6c, 0x63, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f,
	0x76, 0x6f, 0x6c, 0x63, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x6f, 0x64, 0x2f, 0x6d, 0x6f, 0x64, 0x65,
	0x6c, 0x73, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xa0, 0x01, 0x01, 0xd8, 0x01,
	0x01, 0xc2, 0x02, 0x00, 0xca, 0x02, 0x20, 0x56, 0x6f, 0x6c, 0x63, 0x5c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5c, 0x56, 0x6f, 0x64, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xe2, 0x02, 0x23, 0x56, 0x6f, 0x6c, 0x63, 0x5c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x6f, 0x64, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_volcengine_vod_business_vod_smart_strategy_proto_rawDescOnce sync.Once
	file_volcengine_vod_business_vod_smart_strategy_proto_rawDescData = file_volcengine_vod_business_vod_smart_strategy_proto_rawDesc
)

func file_volcengine_vod_business_vod_smart_strategy_proto_rawDescGZIP() []byte {
	file_volcengine_vod_business_vod_smart_strategy_proto_rawDescOnce.Do(func() {
		file_volcengine_vod_business_vod_smart_strategy_proto_rawDescData = protoimpl.X.CompressGZIP(file_volcengine_vod_business_vod_smart_strategy_proto_rawDescData)
	})
	return file_volcengine_vod_business_vod_smart_strategy_proto_rawDescData
}

var file_volcengine_vod_business_vod_smart_strategy_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_volcengine_vod_business_vod_smart_strategy_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_volcengine_vod_business_vod_smart_strategy_proto_goTypes = []interface{}{
	(VodSmartStrategyResponseStreamType)(0),       // 0: Volcengine.Vod.Models.Business.VodSmartStrategyResponseStreamType
	(VodSmartStrategyDataStoreStatus)(0),          // 1: Volcengine.Vod.Models.Business.VodSmartStrategyDataStoreStatus
	(*VodGetSmartStrategyLitePlayInfoResult)(nil), // 2: Volcengine.Vod.Models.Business.VodGetSmartStrategyLitePlayInfoResult
	(*VodPlayInfoModel)(nil),                      // 3: Volcengine.Vod.Models.Business.VodPlayInfoModel
}
var file_volcengine_vod_business_vod_smart_strategy_proto_depIdxs = []int32{
	0, // 0: Volcengine.Vod.Models.Business.VodGetSmartStrategyLitePlayInfoResult.StreamType:type_name -> Volcengine.Vod.Models.Business.VodSmartStrategyResponseStreamType
	1, // 1: Volcengine.Vod.Models.Business.VodGetSmartStrategyLitePlayInfoResult.DataStoreStatus:type_name -> Volcengine.Vod.Models.Business.VodSmartStrategyDataStoreStatus
	3, // 2: Volcengine.Vod.Models.Business.VodGetSmartStrategyLitePlayInfoResult.PlayInfoModel:type_name -> Volcengine.Vod.Models.Business.VodPlayInfoModel
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_volcengine_vod_business_vod_smart_strategy_proto_init() }
func file_volcengine_vod_business_vod_smart_strategy_proto_init() {
	if File_volcengine_vod_business_vod_smart_strategy_proto != nil {
		return
	}
	file_volcengine_vod_business_vod_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_volcengine_vod_business_vod_smart_strategy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodGetSmartStrategyLitePlayInfoResult); i {
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
			RawDescriptor: file_volcengine_vod_business_vod_smart_strategy_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_volcengine_vod_business_vod_smart_strategy_proto_goTypes,
		DependencyIndexes: file_volcengine_vod_business_vod_smart_strategy_proto_depIdxs,
		EnumInfos:         file_volcengine_vod_business_vod_smart_strategy_proto_enumTypes,
		MessageInfos:      file_volcengine_vod_business_vod_smart_strategy_proto_msgTypes,
	}.Build()
	File_volcengine_vod_business_vod_smart_strategy_proto = out.File
	file_volcengine_vod_business_vod_smart_strategy_proto_rawDesc = nil
	file_volcengine_vod_business_vod_smart_strategy_proto_goTypes = nil
	file_volcengine_vod_business_vod_smart_strategy_proto_depIdxs = nil
}
