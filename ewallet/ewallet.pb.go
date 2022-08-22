// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: ewallet/ewallet.proto

package ewallet

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

type Wallet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string  `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Balance float32 `protobuf:"fixed32,4,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *Wallet) Reset() {
	*x = Wallet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ewallet_ewallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Wallet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Wallet) ProtoMessage() {}

func (x *Wallet) ProtoReflect() protoreflect.Message {
	mi := &file_ewallet_ewallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Wallet.ProtoReflect.Descriptor instead.
func (*Wallet) Descriptor() ([]byte, []int) {
	return file_ewallet_ewallet_proto_rawDescGZIP(), []int{0}
}

func (x *Wallet) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Wallet) GetBalance() float32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datetime string  `protobuf:"bytes,1,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Address  string  `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Amount   float32 `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ewallet_ewallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_ewallet_ewallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_ewallet_ewallet_proto_rawDescGZIP(), []int{1}
}

func (x *Transaction) GetDatetime() string {
	if x != nil {
		return x.Datetime
	}
	return ""
}

func (x *Transaction) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Transaction) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CreateWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallet *Wallet `protobuf:"bytes,1,opt,name=wallet,proto3" json:"wallet,omitempty"`
}

func (x *CreateWalletRequest) Reset() {
	*x = CreateWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ewallet_ewallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletRequest) ProtoMessage() {}

func (x *CreateWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ewallet_ewallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletRequest.ProtoReflect.Descriptor instead.
func (*CreateWalletRequest) Descriptor() ([]byte, []int) {
	return file_ewallet_ewallet_proto_rawDescGZIP(), []int{2}
}

func (x *CreateWalletRequest) GetWallet() *Wallet {
	if x != nil {
		return x.Wallet
	}
	return nil
}

type CreateWalletReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wallet *Wallet `protobuf:"bytes,1,opt,name=wallet,proto3" json:"wallet,omitempty"`
}

func (x *CreateWalletReply) Reset() {
	*x = CreateWalletReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ewallet_ewallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWalletReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletReply) ProtoMessage() {}

func (x *CreateWalletReply) ProtoReflect() protoreflect.Message {
	mi := &file_ewallet_ewallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletReply.ProtoReflect.Descriptor instead.
func (*CreateWalletReply) Descriptor() ([]byte, []int) {
	return file_ewallet_ewallet_proto_rawDescGZIP(), []int{3}
}

func (x *CreateWalletReply) GetWallet() *Wallet {
	if x != nil {
		return x.Wallet
	}
	return nil
}

type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From   string  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To     string  `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Amount float32 `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ewallet_ewallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ewallet_ewallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_ewallet_ewallet_proto_rawDescGZIP(), []int{4}
}

func (x *SendRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SendRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SendRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type SendReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendReply) Reset() {
	*x = SendReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ewallet_ewallet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendReply) ProtoMessage() {}

func (x *SendReply) ProtoReflect() protoreflect.Message {
	mi := &file_ewallet_ewallet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendReply.ProtoReflect.Descriptor instead.
func (*SendReply) Descriptor() ([]byte, []int) {
	return file_ewallet_ewallet_proto_rawDescGZIP(), []int{5}
}

func (x *SendReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetLastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLastRequest) Reset() {
	*x = GetLastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ewallet_ewallet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastRequest) ProtoMessage() {}

func (x *GetLastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ewallet_ewallet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastRequest.ProtoReflect.Descriptor instead.
func (*GetLastRequest) Descriptor() ([]byte, []int) {
	return file_ewallet_ewallet_proto_rawDescGZIP(), []int{6}
}

type GetLastReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transaction *Transaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
}

func (x *GetLastReply) Reset() {
	*x = GetLastReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ewallet_ewallet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLastReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLastReply) ProtoMessage() {}

func (x *GetLastReply) ProtoReflect() protoreflect.Message {
	mi := &file_ewallet_ewallet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLastReply.ProtoReflect.Descriptor instead.
func (*GetLastReply) Descriptor() ([]byte, []int) {
	return file_ewallet_ewallet_proto_rawDescGZIP(), []int{7}
}

func (x *GetLastReply) GetTransaction() *Transaction {
	if x != nil {
		return x.Transaction
	}
	return nil
}

var File_ewallet_ewallet_proto protoreflect.FileDescriptor

var file_ewallet_ewallet_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x22, 0x3c, 0x0a, 0x06, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x5b,
	0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3e, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x27, 0x0a, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x52, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x22, 0x3c, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x27, 0x0a, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x52, 0x06, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x22, 0x49, 0x0a, 0x0b, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x10, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x46, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x36, 0x0a,
	0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xc8, 0x01, 0x0a, 0x07, 0x45, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x12, 0x4a, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x12, 0x1c, 0x2e, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x32, 0x0a,
	0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x14, 0x2e, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x65, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x3d, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x65,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c,
	0x61, 0x73, 0x74, 0x42, 0x69, 0x74, 0x39, 0x37, 0x2f, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2f, 0x65, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ewallet_ewallet_proto_rawDescOnce sync.Once
	file_ewallet_ewallet_proto_rawDescData = file_ewallet_ewallet_proto_rawDesc
)

func file_ewallet_ewallet_proto_rawDescGZIP() []byte {
	file_ewallet_ewallet_proto_rawDescOnce.Do(func() {
		file_ewallet_ewallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_ewallet_ewallet_proto_rawDescData)
	})
	return file_ewallet_ewallet_proto_rawDescData
}

var file_ewallet_ewallet_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_ewallet_ewallet_proto_goTypes = []interface{}{
	(*Wallet)(nil),              // 0: ewallet.Wallet
	(*Transaction)(nil),         // 1: ewallet.Transaction
	(*CreateWalletRequest)(nil), // 2: ewallet.CreateWalletRequest
	(*CreateWalletReply)(nil),   // 3: ewallet.CreateWalletReply
	(*SendRequest)(nil),         // 4: ewallet.SendRequest
	(*SendReply)(nil),           // 5: ewallet.SendReply
	(*GetLastRequest)(nil),      // 6: ewallet.GetLastRequest
	(*GetLastReply)(nil),        // 7: ewallet.GetLastReply
}
var file_ewallet_ewallet_proto_depIdxs = []int32{
	0, // 0: ewallet.CreateWalletRequest.wallet:type_name -> ewallet.Wallet
	0, // 1: ewallet.CreateWalletReply.wallet:type_name -> ewallet.Wallet
	1, // 2: ewallet.GetLastReply.transaction:type_name -> ewallet.Transaction
	2, // 3: ewallet.EWallet.CreateWallet:input_type -> ewallet.CreateWalletRequest
	4, // 4: ewallet.EWallet.Send:input_type -> ewallet.SendRequest
	6, // 5: ewallet.EWallet.GetLast:input_type -> ewallet.GetLastRequest
	3, // 6: ewallet.EWallet.CreateWallet:output_type -> ewallet.CreateWalletReply
	5, // 7: ewallet.EWallet.Send:output_type -> ewallet.SendReply
	7, // 8: ewallet.EWallet.GetLast:output_type -> ewallet.GetLastReply
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ewallet_ewallet_proto_init() }
func file_ewallet_ewallet_proto_init() {
	if File_ewallet_ewallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ewallet_ewallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Wallet); i {
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
		file_ewallet_ewallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_ewallet_ewallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWalletRequest); i {
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
		file_ewallet_ewallet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWalletReply); i {
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
		file_ewallet_ewallet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
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
		file_ewallet_ewallet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendReply); i {
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
		file_ewallet_ewallet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastRequest); i {
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
		file_ewallet_ewallet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLastReply); i {
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
			RawDescriptor: file_ewallet_ewallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ewallet_ewallet_proto_goTypes,
		DependencyIndexes: file_ewallet_ewallet_proto_depIdxs,
		MessageInfos:      file_ewallet_ewallet_proto_msgTypes,
	}.Build()
	File_ewallet_ewallet_proto = out.File
	file_ewallet_ewallet_proto_rawDesc = nil
	file_ewallet_ewallet_proto_goTypes = nil
	file_ewallet_ewallet_proto_depIdxs = nil
}
