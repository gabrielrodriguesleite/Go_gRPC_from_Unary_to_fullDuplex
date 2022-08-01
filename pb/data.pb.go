// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: pb/data.proto

package heroes

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

type Calling struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hero string `protobuf:"bytes,1,opt,name=hero,proto3" json:"hero,omitempty"`
}

func (x *Calling) Reset() {
	*x = Calling{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Calling) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Calling) ProtoMessage() {}

func (x *Calling) ProtoReflect() protoreflect.Message {
	mi := &file_pb_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Calling.ProtoReflect.Descriptor instead.
func (*Calling) Descriptor() ([]byte, []int) {
	return file_pb_data_proto_rawDescGZIP(), []int{0}
}

func (x *Calling) GetHero() string {
	if x != nil {
		return x.Hero
	}
	return ""
}

type CallRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calling *Calling `protobuf:"bytes,1,opt,name=calling,proto3" json:"calling,omitempty"`
}

func (x *CallRequest) Reset() {
	*x = CallRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallRequest) ProtoMessage() {}

func (x *CallRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallRequest.ProtoReflect.Descriptor instead.
func (*CallRequest) Descriptor() ([]byte, []int) {
	return file_pb_data_proto_rawDescGZIP(), []int{1}
}

func (x *CallRequest) GetCalling() *Calling {
	if x != nil {
		return x.Calling
	}
	return nil
}

type CallResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CallResponse) Reset() {
	*x = CallResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallResponse) ProtoMessage() {}

func (x *CallResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallResponse.ProtoReflect.Descriptor instead.
func (*CallResponse) Descriptor() ([]byte, []int) {
	return file_pb_data_proto_rawDescGZIP(), []int{2}
}

func (x *CallResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type CallTeamRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calling *Calling `protobuf:"bytes,1,opt,name=calling,proto3" json:"calling,omitempty"`
}

func (x *CallTeamRequest) Reset() {
	*x = CallTeamRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallTeamRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallTeamRequest) ProtoMessage() {}

func (x *CallTeamRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallTeamRequest.ProtoReflect.Descriptor instead.
func (*CallTeamRequest) Descriptor() ([]byte, []int) {
	return file_pb_data_proto_rawDescGZIP(), []int{3}
}

func (x *CallTeamRequest) GetCalling() *Calling {
	if x != nil {
		return x.Calling
	}
	return nil
}

type CallTeamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CallTeamResponse) Reset() {
	*x = CallTeamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallTeamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallTeamResponse) ProtoMessage() {}

func (x *CallTeamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallTeamResponse.ProtoReflect.Descriptor instead.
func (*CallTeamResponse) Descriptor() ([]byte, []int) {
	return file_pb_data_proto_rawDescGZIP(), []int{4}
}

func (x *CallTeamResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type CallManyHeroesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calling *Calling `protobuf:"bytes,1,opt,name=calling,proto3" json:"calling,omitempty"`
}

func (x *CallManyHeroesRequest) Reset() {
	*x = CallManyHeroesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallManyHeroesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallManyHeroesRequest) ProtoMessage() {}

func (x *CallManyHeroesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallManyHeroesRequest.ProtoReflect.Descriptor instead.
func (*CallManyHeroesRequest) Descriptor() ([]byte, []int) {
	return file_pb_data_proto_rawDescGZIP(), []int{5}
}

func (x *CallManyHeroesRequest) GetCalling() *Calling {
	if x != nil {
		return x.Calling
	}
	return nil
}

type CallManyHeroesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CallManyHeroesResponse) Reset() {
	*x = CallManyHeroesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_data_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallManyHeroesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallManyHeroesResponse) ProtoMessage() {}

func (x *CallManyHeroesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_data_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallManyHeroesResponse.ProtoReflect.Descriptor instead.
func (*CallManyHeroesResponse) Descriptor() ([]byte, []int) {
	return file_pb_data_proto_rawDescGZIP(), []int{6}
}

func (x *CallManyHeroesResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type CallEveryoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Calling *Calling `protobuf:"bytes,1,opt,name=calling,proto3" json:"calling,omitempty"`
}

func (x *CallEveryoneRequest) Reset() {
	*x = CallEveryoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_data_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallEveryoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallEveryoneRequest) ProtoMessage() {}

func (x *CallEveryoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_data_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallEveryoneRequest.ProtoReflect.Descriptor instead.
func (*CallEveryoneRequest) Descriptor() ([]byte, []int) {
	return file_pb_data_proto_rawDescGZIP(), []int{7}
}

func (x *CallEveryoneRequest) GetCalling() *Calling {
	if x != nil {
		return x.Calling
	}
	return nil
}

type CallEveryoneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CallEveryoneResponse) Reset() {
	*x = CallEveryoneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_data_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallEveryoneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallEveryoneResponse) ProtoMessage() {}

func (x *CallEveryoneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_data_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallEveryoneResponse.ProtoReflect.Descriptor instead.
func (*CallEveryoneResponse) Descriptor() ([]byte, []int) {
	return file_pb_data_proto_rawDescGZIP(), []int{8}
}

func (x *CallEveryoneResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_pb_data_proto protoreflect.FileDescriptor

var file_pb_data_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x68, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x22, 0x1d, 0x0a, 0x07, 0x43, 0x61, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x65, 0x72, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x65, 0x72, 0x6f, 0x22, 0x38, 0x0a, 0x0b, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x68, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x22, 0x26, 0x0a, 0x0c, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x3c, 0x0a, 0x0f, 0x43, 0x61, 0x6c, 0x6c,
	0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x63,
	0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x68,
	0x65, 0x72, 0x6f, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x63,
	0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x2a, 0x0a, 0x10, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x65,
	0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x42, 0x0a, 0x15, 0x43, 0x61, 0x6c, 0x6c, 0x4d, 0x61, 0x6e, 0x79, 0x48, 0x65,
	0x72, 0x6f, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x63,
	0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x68,
	0x65, 0x72, 0x6f, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x52, 0x07, 0x63,
	0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x30, 0x0a, 0x16, 0x43, 0x61, 0x6c, 0x6c, 0x4d, 0x61,
	0x6e, 0x79, 0x48, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x40, 0x0a, 0x13, 0x43, 0x61, 0x6c, 0x6c,
	0x45, 0x76, 0x65, 0x72, 0x79, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x29, 0x0a, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x68, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x52, 0x07, 0x63, 0x61, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x22, 0x2e, 0x0a, 0x14, 0x43, 0x61,
	0x6c, 0x6c, 0x45, 0x76, 0x65, 0x72, 0x79, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xad, 0x02, 0x0a, 0x0d, 0x48,
	0x65, 0x72, 0x6f, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x04,
	0x43, 0x61, 0x6c, 0x6c, 0x12, 0x13, 0x2e, 0x68, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x2e, 0x43, 0x61,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x68, 0x65, 0x72, 0x6f,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x41, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x17, 0x2e,
	0x68, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x68, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x53, 0x0a, 0x0e, 0x43, 0x61, 0x6c, 0x6c, 0x4d, 0x61, 0x6e, 0x79,
	0x48, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x68, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x6c, 0x6c, 0x4d, 0x61, 0x6e, 0x79, 0x48, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x68, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x2e, 0x43,
	0x61, 0x6c, 0x6c, 0x4d, 0x61, 0x6e, 0x79, 0x48, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x4f, 0x0a, 0x0c, 0x43, 0x61, 0x6c,
	0x6c, 0x45, 0x76, 0x65, 0x72, 0x79, 0x6f, 0x6e, 0x65, 0x12, 0x1b, 0x2e, 0x68, 0x65, 0x72, 0x6f,
	0x65, 0x73, 0x2e, 0x43, 0x61, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x72, 0x79, 0x6f, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x2e,
	0x43, 0x61, 0x6c, 0x6c, 0x45, 0x76, 0x65, 0x72, 0x79, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x62,
	0x2f, 0x68, 0x65, 0x72, 0x6f, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_data_proto_rawDescOnce sync.Once
	file_pb_data_proto_rawDescData = file_pb_data_proto_rawDesc
)

func file_pb_data_proto_rawDescGZIP() []byte {
	file_pb_data_proto_rawDescOnce.Do(func() {
		file_pb_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_data_proto_rawDescData)
	})
	return file_pb_data_proto_rawDescData
}

var file_pb_data_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pb_data_proto_goTypes = []interface{}{
	(*Calling)(nil),                // 0: heroes.Calling
	(*CallRequest)(nil),            // 1: heroes.CallRequest
	(*CallResponse)(nil),           // 2: heroes.CallResponse
	(*CallTeamRequest)(nil),        // 3: heroes.CallTeamRequest
	(*CallTeamResponse)(nil),       // 4: heroes.CallTeamResponse
	(*CallManyHeroesRequest)(nil),  // 5: heroes.CallManyHeroesRequest
	(*CallManyHeroesResponse)(nil), // 6: heroes.CallManyHeroesResponse
	(*CallEveryoneRequest)(nil),    // 7: heroes.CallEveryoneRequest
	(*CallEveryoneResponse)(nil),   // 8: heroes.CallEveryoneResponse
}
var file_pb_data_proto_depIdxs = []int32{
	0, // 0: heroes.CallRequest.calling:type_name -> heroes.Calling
	0, // 1: heroes.CallTeamRequest.calling:type_name -> heroes.Calling
	0, // 2: heroes.CallManyHeroesRequest.calling:type_name -> heroes.Calling
	0, // 3: heroes.CallEveryoneRequest.calling:type_name -> heroes.Calling
	1, // 4: heroes.HeroesService.Call:input_type -> heroes.CallRequest
	3, // 5: heroes.HeroesService.CallTeam:input_type -> heroes.CallTeamRequest
	5, // 6: heroes.HeroesService.CallManyHeroes:input_type -> heroes.CallManyHeroesRequest
	7, // 7: heroes.HeroesService.CallEveryone:input_type -> heroes.CallEveryoneRequest
	2, // 8: heroes.HeroesService.Call:output_type -> heroes.CallResponse
	4, // 9: heroes.HeroesService.CallTeam:output_type -> heroes.CallTeamResponse
	6, // 10: heroes.HeroesService.CallManyHeroes:output_type -> heroes.CallManyHeroesResponse
	8, // 11: heroes.HeroesService.CallEveryone:output_type -> heroes.CallEveryoneResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pb_data_proto_init() }
func file_pb_data_proto_init() {
	if File_pb_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Calling); i {
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
		file_pb_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallRequest); i {
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
		file_pb_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallResponse); i {
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
		file_pb_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallTeamRequest); i {
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
		file_pb_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallTeamResponse); i {
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
		file_pb_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallManyHeroesRequest); i {
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
		file_pb_data_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallManyHeroesResponse); i {
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
		file_pb_data_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallEveryoneRequest); i {
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
		file_pb_data_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallEveryoneResponse); i {
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
			RawDescriptor: file_pb_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_data_proto_goTypes,
		DependencyIndexes: file_pb_data_proto_depIdxs,
		MessageInfos:      file_pb_data_proto_msgTypes,
	}.Build()
	File_pb_data_proto = out.File
	file_pb_data_proto_rawDesc = nil
	file_pb_data_proto_goTypes = nil
	file_pb_data_proto_depIdxs = nil
}
