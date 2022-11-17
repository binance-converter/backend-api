// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: proto/currencies.proto

package currencies

import (
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

type CurrencyCode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyCode string `protobuf:"bytes,1,opt,name=currencyCode,proto3" json:"currencyCode,omitempty"`
}

func (x *CurrencyCode) Reset() {
	*x = CurrencyCode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currencies_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyCode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyCode) ProtoMessage() {}

func (x *CurrencyCode) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencies_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyCode.ProtoReflect.Descriptor instead.
func (*CurrencyCode) Descriptor() ([]byte, []int) {
	return file_proto_currencies_proto_rawDescGZIP(), []int{0}
}

func (x *CurrencyCode) GetCurrencyCode() string {
	if x != nil {
		return x.CurrencyCode
	}
	return ""
}

type CurrencyCodes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyCodes []*CurrencyCode `protobuf:"bytes,1,rep,name=currencyCodes,proto3" json:"currencyCodes,omitempty"`
}

func (x *CurrencyCodes) Reset() {
	*x = CurrencyCodes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currencies_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CurrencyCodes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CurrencyCodes) ProtoMessage() {}

func (x *CurrencyCodes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencies_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CurrencyCodes.ProtoReflect.Descriptor instead.
func (*CurrencyCodes) Descriptor() ([]byte, []int) {
	return file_proto_currencies_proto_rawDescGZIP(), []int{1}
}

func (x *CurrencyCodes) GetCurrencyCodes() []*CurrencyCode {
	if x != nil {
		return x.CurrencyCodes
	}
	return nil
}

type BankName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BankName string `protobuf:"bytes,1,opt,name=bankName,proto3" json:"bankName,omitempty"`
}

func (x *BankName) Reset() {
	*x = BankName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currencies_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankName) ProtoMessage() {}

func (x *BankName) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencies_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankName.ProtoReflect.Descriptor instead.
func (*BankName) Descriptor() ([]byte, []int) {
	return file_proto_currencies_proto_rawDescGZIP(), []int{2}
}

func (x *BankName) GetBankName() string {
	if x != nil {
		return x.BankName
	}
	return ""
}

type BankNames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BankNames []*BankName `protobuf:"bytes,1,rep,name=bankNames,proto3" json:"bankNames,omitempty"`
}

func (x *BankNames) Reset() {
	*x = BankNames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currencies_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BankNames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BankNames) ProtoMessage() {}

func (x *BankNames) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencies_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BankNames.ProtoReflect.Descriptor instead.
func (*BankNames) Descriptor() ([]byte, []int) {
	return file_proto_currencies_proto_rawDescGZIP(), []int{3}
}

func (x *BankNames) GetBankNames() []*BankName {
	if x != nil {
		return x.BankNames
	}
	return nil
}

type FullCurrency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrencyCode *CurrencyCode `protobuf:"bytes,1,opt,name=currencyCode,proto3" json:"currencyCode,omitempty"`
	BankName     *BankName     `protobuf:"bytes,2,opt,name=bankName,proto3" json:"bankName,omitempty"`
}

func (x *FullCurrency) Reset() {
	*x = FullCurrency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currencies_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FullCurrency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FullCurrency) ProtoMessage() {}

func (x *FullCurrency) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencies_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FullCurrency.ProtoReflect.Descriptor instead.
func (*FullCurrency) Descriptor() ([]byte, []int) {
	return file_proto_currencies_proto_rawDescGZIP(), []int{4}
}

func (x *FullCurrency) GetCurrencyCode() *CurrencyCode {
	if x != nil {
		return x.CurrencyCode
	}
	return nil
}

func (x *FullCurrency) GetBankName() *BankName {
	if x != nil {
		return x.BankName
	}
	return nil
}

type FullCurrencies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FullCurrencies []*FullCurrency `protobuf:"bytes,1,rep,name=fullCurrencies,proto3" json:"fullCurrencies,omitempty"`
}

func (x *FullCurrencies) Reset() {
	*x = FullCurrencies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_currencies_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FullCurrencies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FullCurrencies) ProtoMessage() {}

func (x *FullCurrencies) ProtoReflect() protoreflect.Message {
	mi := &file_proto_currencies_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FullCurrencies.ProtoReflect.Descriptor instead.
func (*FullCurrencies) Descriptor() ([]byte, []int) {
	return file_proto_currencies_proto_rawDescGZIP(), []int{5}
}

func (x *FullCurrencies) GetFullCurrencies() []*FullCurrency {
	if x != nil {
		return x.FullCurrencies
	}
	return nil
}

var File_proto_currencies_proto protoreflect.FileDescriptor

var file_proto_currencies_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x32, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x6d, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x73, 0x12, 0x5c, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x62, 0x69,
	0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x73, 0x22, 0x26, 0x0a, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x5d, 0x0a, 0x09, 0x62, 0x61,
	0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x62, 0x69, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x09,
	0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0xba, 0x01, 0x0a, 0x0c, 0x66, 0x75,
	0x6c, 0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x5a, 0x0a, 0x0c, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x4e, 0x0a, 0x08, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x08, 0x62, 0x61,
	0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x70, 0x0a, 0x0e, 0x66, 0x75, 0x6c, 0x6c, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x5e, 0x0a, 0x0e, 0x66, 0x75, 0x6c, 0x6c,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65,
	0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x66, 0x75, 0x6c, 0x6c,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x0e, 0x66, 0x75, 0x6c, 0x6c, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x32, 0xc7, 0x03, 0x0a, 0x0a, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x69, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x76,
	0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65,
	0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x37, 0x2e, 0x62, 0x69, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64,
	0x65, 0x73, 0x12, 0x89, 0x01, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x42, 0x79, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x36, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x1a, 0x33, 0x2e, 0x62, 0x69, 0x6e, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x2e, 0x62, 0x61, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x5d,
	0x0a, 0x0b, 0x53, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x36, 0x2e,
	0x62, 0x69, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65,
	0x72, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x2e, 0x66, 0x75, 0x6c, 0x6c, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x63, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x63, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x38, 0x2e, 0x62, 0x69, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x65, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x69, 0x65, 0x73, 0x2e, 0x66, 0x75, 0x6c, 0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x69,
	0x65, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_currencies_proto_rawDescOnce sync.Once
	file_proto_currencies_proto_rawDescData = file_proto_currencies_proto_rawDesc
)

func file_proto_currencies_proto_rawDescGZIP() []byte {
	file_proto_currencies_proto_rawDescOnce.Do(func() {
		file_proto_currencies_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_currencies_proto_rawDescData)
	})
	return file_proto_currencies_proto_rawDescData
}

var file_proto_currencies_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_currencies_proto_goTypes = []interface{}{
	(*CurrencyCode)(nil),   // 0: binance_converter.backend_api.currencies.currencyCode
	(*CurrencyCodes)(nil),  // 1: binance_converter.backend_api.currencies.currencyCodes
	(*BankName)(nil),       // 2: binance_converter.backend_api.currencies.bankName
	(*BankNames)(nil),      // 3: binance_converter.backend_api.currencies.bankNames
	(*FullCurrency)(nil),   // 4: binance_converter.backend_api.currencies.fullCurrency
	(*FullCurrencies)(nil), // 5: binance_converter.backend_api.currencies.fullCurrencies
	(*emptypb.Empty)(nil),  // 6: google.protobuf.Empty
}
var file_proto_currencies_proto_depIdxs = []int32{
	0, // 0: binance_converter.backend_api.currencies.currencyCodes.currencyCodes:type_name -> binance_converter.backend_api.currencies.currencyCode
	2, // 1: binance_converter.backend_api.currencies.bankNames.bankNames:type_name -> binance_converter.backend_api.currencies.bankName
	0, // 2: binance_converter.backend_api.currencies.fullCurrency.currencyCode:type_name -> binance_converter.backend_api.currencies.currencyCode
	2, // 3: binance_converter.backend_api.currencies.fullCurrency.bankName:type_name -> binance_converter.backend_api.currencies.bankName
	4, // 4: binance_converter.backend_api.currencies.fullCurrencies.fullCurrencies:type_name -> binance_converter.backend_api.currencies.fullCurrency
	6, // 5: binance_converter.backend_api.currencies.currencies.GetAvailableCurrencies:input_type -> google.protobuf.Empty
	0, // 6: binance_converter.backend_api.currencies.currencies.GetAvailableBankByCurrency:input_type -> binance_converter.backend_api.currencies.currencyCode
	4, // 7: binance_converter.backend_api.currencies.currencies.SetCurrency:input_type -> binance_converter.backend_api.currencies.fullCurrency
	6, // 8: binance_converter.backend_api.currencies.currencies.GetMyCurrencies:input_type -> google.protobuf.Empty
	1, // 9: binance_converter.backend_api.currencies.currencies.GetAvailableCurrencies:output_type -> binance_converter.backend_api.currencies.currencyCodes
	3, // 10: binance_converter.backend_api.currencies.currencies.GetAvailableBankByCurrency:output_type -> binance_converter.backend_api.currencies.bankNames
	6, // 11: binance_converter.backend_api.currencies.currencies.SetCurrency:output_type -> google.protobuf.Empty
	5, // 12: binance_converter.backend_api.currencies.currencies.GetMyCurrencies:output_type -> binance_converter.backend_api.currencies.fullCurrencies
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_currencies_proto_init() }
func file_proto_currencies_proto_init() {
	if File_proto_currencies_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_currencies_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyCode); i {
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
		file_proto_currencies_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CurrencyCodes); i {
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
		file_proto_currencies_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankName); i {
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
		file_proto_currencies_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BankNames); i {
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
		file_proto_currencies_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FullCurrency); i {
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
		file_proto_currencies_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FullCurrencies); i {
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
			RawDescriptor: file_proto_currencies_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_currencies_proto_goTypes,
		DependencyIndexes: file_proto_currencies_proto_depIdxs,
		MessageInfos:      file_proto_currencies_proto_msgTypes,
	}.Build()
	File_proto_currencies_proto = out.File
	file_proto_currencies_proto_rawDesc = nil
	file_proto_currencies_proto_goTypes = nil
	file_proto_currencies_proto_depIdxs = nil
}
