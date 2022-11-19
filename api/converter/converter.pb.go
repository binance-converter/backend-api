// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: proto/converter.proto

package converter

import (
	currencies "/api/currencies"
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

type AdditionalErrorCode int32

const (
	AdditionalErrorCode_OK                     AdditionalErrorCode = 0
	AdditionalErrorCode_INVALID_CONVERTER_PAIR AdditionalErrorCode = 100
	AdditionalErrorCode_INVALID_FORMAT         AdditionalErrorCode = 101
	AdditionalErrorCode_INVALID_CURRENCY_CODE  AdditionalErrorCode = 102
	AdditionalErrorCode_INVALID_BANK_CODE      AdditionalErrorCode = 103
)

// Enum value maps for AdditionalErrorCode.
var (
	AdditionalErrorCode_name = map[int32]string{
		0:   "OK",
		100: "INVALID_CONVERTER_PAIR",
		101: "INVALID_FORMAT",
		102: "INVALID_CURRENCY_CODE",
		103: "INVALID_BANK_CODE",
	}
	AdditionalErrorCode_value = map[string]int32{
		"OK":                     0,
		"INVALID_CONVERTER_PAIR": 100,
		"INVALID_FORMAT":         101,
		"INVALID_CURRENCY_CODE":  102,
		"INVALID_BANK_CODE":      103,
	}
)

func (x AdditionalErrorCode) Enum() *AdditionalErrorCode {
	p := new(AdditionalErrorCode)
	*p = x
	return p
}

func (x AdditionalErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdditionalErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_converter_proto_enumTypes[0].Descriptor()
}

func (AdditionalErrorCode) Type() protoreflect.EnumType {
	return &file_proto_converter_proto_enumTypes[0]
}

func (x AdditionalErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdditionalErrorCode.Descriptor instead.
func (AdditionalErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_proto_converter_proto_rawDescGZIP(), []int{0}
}

type ConverterPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConverterPair []*currencies.FullCurrency `protobuf:"bytes,1,rep,name=converterPair,proto3" json:"converterPair,omitempty"`
}

func (x *ConverterPair) Reset() {
	*x = ConverterPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_converter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConverterPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConverterPair) ProtoMessage() {}

func (x *ConverterPair) ProtoReflect() protoreflect.Message {
	mi := &file_proto_converter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConverterPair.ProtoReflect.Descriptor instead.
func (*ConverterPair) Descriptor() ([]byte, []int) {
	return file_proto_converter_proto_rawDescGZIP(), []int{0}
}

func (x *ConverterPair) GetConverterPair() []*currencies.FullCurrency {
	if x != nil {
		return x.ConverterPair
	}
	return nil
}

type ConverterPairs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConverterPairs []*ConverterPair `protobuf:"bytes,1,rep,name=converterPairs,proto3" json:"converterPairs,omitempty"`
}

func (x *ConverterPairs) Reset() {
	*x = ConverterPairs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_converter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConverterPairs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConverterPairs) ProtoMessage() {}

func (x *ConverterPairs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_converter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConverterPairs.ProtoReflect.Descriptor instead.
func (*ConverterPairs) Descriptor() ([]byte, []int) {
	return file_proto_converter_proto_rawDescGZIP(), []int{1}
}

func (x *ConverterPairs) GetConverterPairs() []*ConverterPair {
	if x != nil {
		return x.ConverterPairs
	}
	return nil
}

type Exchange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exchange float32 `protobuf:"fixed32,1,opt,name=exchange,proto3" json:"exchange,omitempty"`
}

func (x *Exchange) Reset() {
	*x = Exchange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_converter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Exchange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Exchange) ProtoMessage() {}

func (x *Exchange) ProtoReflect() protoreflect.Message {
	mi := &file_proto_converter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Exchange.ProtoReflect.Descriptor instead.
func (*Exchange) Descriptor() ([]byte, []int) {
	return file_proto_converter_proto_rawDescGZIP(), []int{2}
}

func (x *Exchange) GetExchange() float32 {
	if x != nil {
		return x.Exchange
	}
	return 0
}

type ThresholdConvertPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConverterPair *currencies.FullCurrency `protobuf:"bytes,1,opt,name=converterPair,proto3" json:"converterPair,omitempty"`
	Exchange      *Exchange                `protobuf:"bytes,2,opt,name=exchange,proto3" json:"exchange,omitempty"`
}

func (x *ThresholdConvertPair) Reset() {
	*x = ThresholdConvertPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_converter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThresholdConvertPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThresholdConvertPair) ProtoMessage() {}

func (x *ThresholdConvertPair) ProtoReflect() protoreflect.Message {
	mi := &file_proto_converter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThresholdConvertPair.ProtoReflect.Descriptor instead.
func (*ThresholdConvertPair) Descriptor() ([]byte, []int) {
	return file_proto_converter_proto_rawDescGZIP(), []int{3}
}

func (x *ThresholdConvertPair) GetConverterPair() *currencies.FullCurrency {
	if x != nil {
		return x.ConverterPair
	}
	return nil
}

func (x *ThresholdConvertPair) GetExchange() *Exchange {
	if x != nil {
		return x.Exchange
	}
	return nil
}

type ThresholdConvertPairs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConverterPairs []*ThresholdConvertPair `protobuf:"bytes,1,rep,name=converterPairs,proto3" json:"converterPairs,omitempty"`
}

func (x *ThresholdConvertPairs) Reset() {
	*x = ThresholdConvertPairs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_converter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThresholdConvertPairs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThresholdConvertPairs) ProtoMessage() {}

func (x *ThresholdConvertPairs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_converter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThresholdConvertPairs.ProtoReflect.Descriptor instead.
func (*ThresholdConvertPairs) Descriptor() ([]byte, []int) {
	return file_proto_converter_proto_rawDescGZIP(), []int{4}
}

func (x *ThresholdConvertPairs) GetConverterPairs() []*ThresholdConvertPair {
	if x != nil {
		return x.ConverterPairs
	}
	return nil
}

var File_proto_converter_proto protoreflect.FileDescriptor

var file_proto_converter_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x27, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x12, 0x5c, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e,
	0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65,
	0x72, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x66, 0x75, 0x6c, 0x6c, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72,
	0x50, 0x61, 0x69, 0x72, 0x22, 0x70, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65,
	0x72, 0x50, 0x61, 0x69, 0x72, 0x73, 0x12, 0x5e, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65,
	0x72, 0x50, 0x61, 0x69, 0x72, 0x73, 0x22, 0x26, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0xc3,
	0x01, 0x0a, 0x14, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x50, 0x61, 0x69, 0x72, 0x12, 0x5c, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x74, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x65, 0x72, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x66, 0x75, 0x6c, 0x6c, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65,
	0x72, 0x50, 0x61, 0x69, 0x72, 0x12, 0x4d, 0x0a, 0x08, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65,
	0x72, 0x2e, 0x65, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x08, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x22, 0x7e, 0x0a, 0x15, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c,
	0x64, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x50, 0x61, 0x69, 0x72, 0x73, 0x12, 0x65, 0x0a,
	0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x50, 0x61, 0x69, 0x72, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x50,
	0x61, 0x69, 0x72, 0x73, 0x2a, 0x7f, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x06, 0x0a, 0x02, 0x4f,
	0x4b, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43,
	0x4f, 0x4e, 0x56, 0x45, 0x52, 0x54, 0x45, 0x52, 0x5f, 0x50, 0x41, 0x49, 0x52, 0x10, 0x64, 0x12,
	0x12, 0x0a, 0x0e, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41,
	0x54, 0x10, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43,
	0x55, 0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x10, 0x66, 0x12, 0x15,
	0x0a, 0x11, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x42, 0x41, 0x4e, 0x4b, 0x5f, 0x43,
	0x4f, 0x44, 0x45, 0x10, 0x67, 0x32, 0xac, 0x05, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x65, 0x72, 0x12, 0x6d, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x37, 0x2e, 0x62, 0x69, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x50, 0x61, 0x69,
	0x72, 0x73, 0x12, 0x60, 0x0a, 0x0e, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74,
	0x50, 0x61, 0x69, 0x72, 0x12, 0x36, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x64, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x43, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x74, 0x50, 0x61, 0x69, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x37, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x65, 0x72, 0x50, 0x61, 0x69, 0x72, 0x73, 0x12, 0x71, 0x0a, 0x18, 0x53, 0x65,
	0x74, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x50, 0x61, 0x69, 0x72, 0x73, 0x12, 0x3d, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65,
	0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72,
	0x2e, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x50, 0x61, 0x69, 0x72, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x74, 0x0a,
	0x1a, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x54, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x43,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x50, 0x61, 0x69, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x1a, 0x3e, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x74, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x50, 0x61,
	0x69, 0x72, 0x73, 0x12, 0x7f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x45, 0x78, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x36, 0x2e, 0x62, 0x69, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x50, 0x61, 0x69,
	0x72, 0x1a, 0x31, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x65, 0x78, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x42, 0x10, 0x5a, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e,
	0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_converter_proto_rawDescOnce sync.Once
	file_proto_converter_proto_rawDescData = file_proto_converter_proto_rawDesc
)

func file_proto_converter_proto_rawDescGZIP() []byte {
	file_proto_converter_proto_rawDescOnce.Do(func() {
		file_proto_converter_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_converter_proto_rawDescData)
	})
	return file_proto_converter_proto_rawDescData
}

var file_proto_converter_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_converter_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_converter_proto_goTypes = []interface{}{
	(AdditionalErrorCode)(0),        // 0: binance_converter.backend_api.converter.AdditionalErrorCode
	(*ConverterPair)(nil),           // 1: binance_converter.backend_api.converter.converterPair
	(*ConverterPairs)(nil),          // 2: binance_converter.backend_api.converter.converterPairs
	(*Exchange)(nil),                // 3: binance_converter.backend_api.converter.exchange
	(*ThresholdConvertPair)(nil),    // 4: binance_converter.backend_api.converter.thresholdConvertPair
	(*ThresholdConvertPairs)(nil),   // 5: binance_converter.backend_api.converter.thresholdConvertPairs
	(*currencies.FullCurrency)(nil), // 6: binance_converter.backend_api.currencies.fullCurrency
	(*emptypb.Empty)(nil),           // 7: google.protobuf.Empty
}
var file_proto_converter_proto_depIdxs = []int32{
	6,  // 0: binance_converter.backend_api.converter.converterPair.converterPair:type_name -> binance_converter.backend_api.currencies.fullCurrency
	1,  // 1: binance_converter.backend_api.converter.converterPairs.converterPairs:type_name -> binance_converter.backend_api.converter.converterPair
	6,  // 2: binance_converter.backend_api.converter.thresholdConvertPair.converterPair:type_name -> binance_converter.backend_api.currencies.fullCurrency
	3,  // 3: binance_converter.backend_api.converter.thresholdConvertPair.exchange:type_name -> binance_converter.backend_api.converter.exchange
	4,  // 4: binance_converter.backend_api.converter.thresholdConvertPairs.converterPairs:type_name -> binance_converter.backend_api.converter.thresholdConvertPair
	7,  // 5: binance_converter.backend_api.converter.converter.GetAvailableConverterPairs:input_type -> google.protobuf.Empty
	1,  // 6: binance_converter.backend_api.converter.converter.SetConvertPair:input_type -> binance_converter.backend_api.converter.converterPair
	7,  // 7: binance_converter.backend_api.converter.converter.GetMyConvertPairs:input_type -> google.protobuf.Empty
	4,  // 8: binance_converter.backend_api.converter.converter.SetThresholdConvertPairs:input_type -> binance_converter.backend_api.converter.thresholdConvertPair
	7,  // 9: binance_converter.backend_api.converter.converter.GetMyThresholdConvertPairs:input_type -> google.protobuf.Empty
	1,  // 10: binance_converter.backend_api.converter.converter.GetCurrentExchange:input_type -> binance_converter.backend_api.converter.converterPair
	2,  // 11: binance_converter.backend_api.converter.converter.GetAvailableConverterPairs:output_type -> binance_converter.backend_api.converter.converterPairs
	7,  // 12: binance_converter.backend_api.converter.converter.SetConvertPair:output_type -> google.protobuf.Empty
	2,  // 13: binance_converter.backend_api.converter.converter.GetMyConvertPairs:output_type -> binance_converter.backend_api.converter.converterPairs
	7,  // 14: binance_converter.backend_api.converter.converter.SetThresholdConvertPairs:output_type -> google.protobuf.Empty
	5,  // 15: binance_converter.backend_api.converter.converter.GetMyThresholdConvertPairs:output_type -> binance_converter.backend_api.converter.thresholdConvertPairs
	3,  // 16: binance_converter.backend_api.converter.converter.GetCurrentExchange:output_type -> binance_converter.backend_api.converter.exchange
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_converter_proto_init() }
func file_proto_converter_proto_init() {
	if File_proto_converter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_converter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConverterPair); i {
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
		file_proto_converter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConverterPairs); i {
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
		file_proto_converter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Exchange); i {
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
		file_proto_converter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThresholdConvertPair); i {
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
		file_proto_converter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThresholdConvertPairs); i {
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
			RawDescriptor: file_proto_converter_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_converter_proto_goTypes,
		DependencyIndexes: file_proto_converter_proto_depIdxs,
		EnumInfos:         file_proto_converter_proto_enumTypes,
		MessageInfos:      file_proto_converter_proto_msgTypes,
	}.Build()
	File_proto_converter_proto = out.File
	file_proto_converter_proto_rawDesc = nil
	file_proto_converter_proto_goTypes = nil
	file_proto_converter_proto_depIdxs = nil
}
