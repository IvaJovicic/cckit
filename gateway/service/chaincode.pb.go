// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: service/chaincode.proto

package service

import (
	peer "github.com/hyperledger/fabric-protos-go/peer"
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

type InvocationType int32

const (
	InvocationType_QUERY  InvocationType = 0
	InvocationType_INVOKE InvocationType = 1
)

// Enum value maps for InvocationType.
var (
	InvocationType_name = map[int32]string{
		0: "QUERY",
		1: "INVOKE",
	}
	InvocationType_value = map[string]int32{
		"QUERY":  0,
		"INVOKE": 1,
	}
)

func (x InvocationType) Enum() *InvocationType {
	p := new(InvocationType)
	*p = x
	return p
}

func (x InvocationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InvocationType) Descriptor() protoreflect.EnumDescriptor {
	return file_service_chaincode_proto_enumTypes[0].Descriptor()
}

func (InvocationType) Type() protoreflect.EnumType {
	return &file_service_chaincode_proto_enumTypes[0]
}

func (x InvocationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InvocationType.Descriptor instead.
func (InvocationType) EnumDescriptor() ([]byte, []int) {
	return file_service_chaincode_proto_rawDescGZIP(), []int{0}
}

type ChaincodeInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chaincode name
	Chaincode string `protobuf:"bytes,1,opt,name=chaincode,proto3" json:"chaincode,omitempty"`
	// Channel name
	Channel string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	// Input contains the arguments for invocation.
	Args [][]byte `protobuf:"bytes,3,rep,name=args,proto3" json:"args,omitempty"`
	// TransientMap contains data (e.g. cryptographic material) that might be used
	// to implement some form of application-level confidentiality. The contents
	// of this field are supposed to always be omitted from the transaction and
	// excluded from the ledger.
	Transient map[string][]byte `protobuf:"bytes,4,rep,name=transient,proto3" json:"transient,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ChaincodeInput) Reset() {
	*x = ChaincodeInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_chaincode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeInput) ProtoMessage() {}

func (x *ChaincodeInput) ProtoReflect() protoreflect.Message {
	mi := &file_service_chaincode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeInput.ProtoReflect.Descriptor instead.
func (*ChaincodeInput) Descriptor() ([]byte, []int) {
	return file_service_chaincode_proto_rawDescGZIP(), []int{0}
}

func (x *ChaincodeInput) GetChaincode() string {
	if x != nil {
		return x.Chaincode
	}
	return ""
}

func (x *ChaincodeInput) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *ChaincodeInput) GetArgs() [][]byte {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *ChaincodeInput) GetTransient() map[string][]byte {
	if x != nil {
		return x.Transient
	}
	return nil
}

type ChaincodeLocator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Chaincode name
	Chaincode string `protobuf:"bytes,1,opt,name=chaincode,proto3" json:"chaincode,omitempty"`
	// Channel name
	Channel string `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *ChaincodeLocator) Reset() {
	*x = ChaincodeLocator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_chaincode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeLocator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeLocator) ProtoMessage() {}

func (x *ChaincodeLocator) ProtoReflect() protoreflect.Message {
	mi := &file_service_chaincode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeLocator.ProtoReflect.Descriptor instead.
func (*ChaincodeLocator) Descriptor() ([]byte, []int) {
	return file_service_chaincode_proto_rawDescGZIP(), []int{1}
}

func (x *ChaincodeLocator) GetChaincode() string {
	if x != nil {
		return x.Chaincode
	}
	return ""
}

func (x *ChaincodeLocator) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

type ChaincodeExec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type  InvocationType  `protobuf:"varint,1,opt,name=type,proto3,enum=s7techlab.gateway.service.InvocationType" json:"type,omitempty"`
	Input *ChaincodeInput `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *ChaincodeExec) Reset() {
	*x = ChaincodeExec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_chaincode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeExec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeExec) ProtoMessage() {}

func (x *ChaincodeExec) ProtoReflect() protoreflect.Message {
	mi := &file_service_chaincode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeExec.ProtoReflect.Descriptor instead.
func (*ChaincodeExec) Descriptor() ([]byte, []int) {
	return file_service_chaincode_proto_rawDescGZIP(), []int{2}
}

func (x *ChaincodeExec) GetType() InvocationType {
	if x != nil {
		return x.Type
	}
	return InvocationType_QUERY
}

func (x *ChaincodeExec) GetInput() *ChaincodeInput {
	if x != nil {
		return x.Input
	}
	return nil
}

var File_service_chaincode_proto protoreflect.FileDescriptor

var file_service_chaincode_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x73, 0x37, 0x74, 0x65, 0x63,
	0x68, 0x6c, 0x61, 0x62, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2,
	0x01, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x56, 0x0a,
	0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x38, 0x2e, 0x73, 0x37, 0x74, 0x65, 0x63, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x3c, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x69, 0x65,
	0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x4a, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22,
	0x8f, 0x01, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x45, 0x78, 0x65,
	0x63, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x29, 0x2e, 0x73, 0x37, 0x74, 0x65, 0x63, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x12, 0x3f, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x73, 0x37, 0x74, 0x65, 0x63, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x2a, 0x27, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x51, 0x55, 0x45, 0x52, 0x59, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x10, 0x01, 0x32, 0xc5, 0x02, 0x0a, 0x09, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x4a, 0x0a, 0x04, 0x45, 0x78, 0x65, 0x63,
	0x12, 0x28, 0x2e, 0x73, 0x37, 0x74, 0x65, 0x63, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x45, 0x78, 0x65, 0x63, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x29, 0x2e,
	0x73, 0x37, 0x74, 0x65, 0x63, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x06, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x12, 0x29, 0x2e, 0x73,
	0x37, 0x74, 0x65, 0x63, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4f, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2b, 0x2e, 0x73, 0x37,
	0x74, 0x65, 0x63, 0x68, 0x6c, 0x61, 0x62, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x30, 0x01, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x37, 0x74, 0x65, 0x63, 0x68, 0x6c, 0x61, 0x62, 0x2f, 0x63, 0x63, 0x6b, 0x69, 0x74,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_chaincode_proto_rawDescOnce sync.Once
	file_service_chaincode_proto_rawDescData = file_service_chaincode_proto_rawDesc
)

func file_service_chaincode_proto_rawDescGZIP() []byte {
	file_service_chaincode_proto_rawDescOnce.Do(func() {
		file_service_chaincode_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_chaincode_proto_rawDescData)
	})
	return file_service_chaincode_proto_rawDescData
}

var file_service_chaincode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_service_chaincode_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_service_chaincode_proto_goTypes = []interface{}{
	(InvocationType)(0),           // 0: s7techlab.gateway.service.InvocationType
	(*ChaincodeInput)(nil),        // 1: s7techlab.gateway.service.ChaincodeInput
	(*ChaincodeLocator)(nil),      // 2: s7techlab.gateway.service.ChaincodeLocator
	(*ChaincodeExec)(nil),         // 3: s7techlab.gateway.service.ChaincodeExec
	nil,                           // 4: s7techlab.gateway.service.ChaincodeInput.TransientEntry
	(*peer.ProposalResponse)(nil), // 5: protos.ProposalResponse
	(*peer.ChaincodeEvent)(nil),   // 6: protos.ChaincodeEvent
}
var file_service_chaincode_proto_depIdxs = []int32{
	4, // 0: s7techlab.gateway.service.ChaincodeInput.transient:type_name -> s7techlab.gateway.service.ChaincodeInput.TransientEntry
	0, // 1: s7techlab.gateway.service.ChaincodeExec.type:type_name -> s7techlab.gateway.service.InvocationType
	1, // 2: s7techlab.gateway.service.ChaincodeExec.input:type_name -> s7techlab.gateway.service.ChaincodeInput
	3, // 3: s7techlab.gateway.service.Chaincode.Exec:input_type -> s7techlab.gateway.service.ChaincodeExec
	1, // 4: s7techlab.gateway.service.Chaincode.Query:input_type -> s7techlab.gateway.service.ChaincodeInput
	1, // 5: s7techlab.gateway.service.Chaincode.Invoke:input_type -> s7techlab.gateway.service.ChaincodeInput
	2, // 6: s7techlab.gateway.service.Chaincode.Events:input_type -> s7techlab.gateway.service.ChaincodeLocator
	5, // 7: s7techlab.gateway.service.Chaincode.Exec:output_type -> protos.ProposalResponse
	5, // 8: s7techlab.gateway.service.Chaincode.Query:output_type -> protos.ProposalResponse
	5, // 9: s7techlab.gateway.service.Chaincode.Invoke:output_type -> protos.ProposalResponse
	6, // 10: s7techlab.gateway.service.Chaincode.Events:output_type -> protos.ChaincodeEvent
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_service_chaincode_proto_init() }
func file_service_chaincode_proto_init() {
	if File_service_chaincode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_chaincode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeInput); i {
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
		file_service_chaincode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeLocator); i {
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
		file_service_chaincode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeExec); i {
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
			RawDescriptor: file_service_chaincode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_chaincode_proto_goTypes,
		DependencyIndexes: file_service_chaincode_proto_depIdxs,
		EnumInfos:         file_service_chaincode_proto_enumTypes,
		MessageInfos:      file_service_chaincode_proto_msgTypes,
	}.Build()
	File_service_chaincode_proto = out.File
	file_service_chaincode_proto_rawDesc = nil
	file_service_chaincode_proto_goTypes = nil
	file_service_chaincode_proto_depIdxs = nil
}
