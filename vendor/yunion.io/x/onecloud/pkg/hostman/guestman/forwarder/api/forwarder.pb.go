// Copyright 2019 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: forwarder.proto

package api

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

type OpenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetId      string `protobuf:"bytes,1,opt,name=netId,proto3" json:"netId,omitempty"`
	Proto      string `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
	BindAddr   string `protobuf:"bytes,3,opt,name=bindAddr,proto3" json:"bindAddr,omitempty"`
	BindPort   uint32 `protobuf:"varint,4,opt,name=bindPort,proto3" json:"bindPort,omitempty"`
	RemoteAddr string `protobuf:"bytes,5,opt,name=remoteAddr,proto3" json:"remoteAddr,omitempty"`
	RemotePort uint32 `protobuf:"varint,6,opt,name=remotePort,proto3" json:"remotePort,omitempty"`
}

func (x *OpenRequest) Reset() {
	*x = OpenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forwarder_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenRequest) ProtoMessage() {}

func (x *OpenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_forwarder_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenRequest.ProtoReflect.Descriptor instead.
func (*OpenRequest) Descriptor() ([]byte, []int) {
	return file_forwarder_proto_rawDescGZIP(), []int{0}
}

func (x *OpenRequest) GetNetId() string {
	if x != nil {
		return x.NetId
	}
	return ""
}

func (x *OpenRequest) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *OpenRequest) GetBindAddr() string {
	if x != nil {
		return x.BindAddr
	}
	return ""
}

func (x *OpenRequest) GetBindPort() uint32 {
	if x != nil {
		return x.BindPort
	}
	return 0
}

func (x *OpenRequest) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *OpenRequest) GetRemotePort() uint32 {
	if x != nil {
		return x.RemotePort
	}
	return 0
}

type OpenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetId      string `protobuf:"bytes,1,opt,name=netId,proto3" json:"netId,omitempty"`
	Proto      string `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
	BindAddr   string `protobuf:"bytes,3,opt,name=bindAddr,proto3" json:"bindAddr,omitempty"`
	BindPort   uint32 `protobuf:"varint,4,opt,name=bindPort,proto3" json:"bindPort,omitempty"`
	RemoteAddr string `protobuf:"bytes,5,opt,name=remoteAddr,proto3" json:"remoteAddr,omitempty"`
	RemotePort uint32 `protobuf:"varint,6,opt,name=remotePort,proto3" json:"remotePort,omitempty"`
}

func (x *OpenResponse) Reset() {
	*x = OpenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forwarder_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OpenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenResponse) ProtoMessage() {}

func (x *OpenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_forwarder_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenResponse.ProtoReflect.Descriptor instead.
func (*OpenResponse) Descriptor() ([]byte, []int) {
	return file_forwarder_proto_rawDescGZIP(), []int{1}
}

func (x *OpenResponse) GetNetId() string {
	if x != nil {
		return x.NetId
	}
	return ""
}

func (x *OpenResponse) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *OpenResponse) GetBindAddr() string {
	if x != nil {
		return x.BindAddr
	}
	return ""
}

func (x *OpenResponse) GetBindPort() uint32 {
	if x != nil {
		return x.BindPort
	}
	return 0
}

func (x *OpenResponse) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *OpenResponse) GetRemotePort() uint32 {
	if x != nil {
		return x.RemotePort
	}
	return 0
}

type CloseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetId    string `protobuf:"bytes,1,opt,name=netId,proto3" json:"netId,omitempty"`
	Proto    string `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
	BindAddr string `protobuf:"bytes,3,opt,name=bindAddr,proto3" json:"bindAddr,omitempty"`
	BindPort uint32 `protobuf:"varint,4,opt,name=bindPort,proto3" json:"bindPort,omitempty"`
}

func (x *CloseRequest) Reset() {
	*x = CloseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forwarder_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseRequest) ProtoMessage() {}

func (x *CloseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_forwarder_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseRequest.ProtoReflect.Descriptor instead.
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return file_forwarder_proto_rawDescGZIP(), []int{2}
}

func (x *CloseRequest) GetNetId() string {
	if x != nil {
		return x.NetId
	}
	return ""
}

func (x *CloseRequest) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *CloseRequest) GetBindAddr() string {
	if x != nil {
		return x.BindAddr
	}
	return ""
}

func (x *CloseRequest) GetBindPort() uint32 {
	if x != nil {
		return x.BindPort
	}
	return 0
}

type CloseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetId    string `protobuf:"bytes,1,opt,name=netId,proto3" json:"netId,omitempty"`
	Proto    string `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
	BindAddr string `protobuf:"bytes,3,opt,name=bindAddr,proto3" json:"bindAddr,omitempty"`
	BindPort uint32 `protobuf:"varint,4,opt,name=bindPort,proto3" json:"bindPort,omitempty"`
}

func (x *CloseResponse) Reset() {
	*x = CloseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forwarder_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseResponse) ProtoMessage() {}

func (x *CloseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_forwarder_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseResponse.ProtoReflect.Descriptor instead.
func (*CloseResponse) Descriptor() ([]byte, []int) {
	return file_forwarder_proto_rawDescGZIP(), []int{3}
}

func (x *CloseResponse) GetNetId() string {
	if x != nil {
		return x.NetId
	}
	return ""
}

func (x *CloseResponse) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *CloseResponse) GetBindAddr() string {
	if x != nil {
		return x.BindAddr
	}
	return ""
}

func (x *CloseResponse) GetBindPort() uint32 {
	if x != nil {
		return x.BindPort
	}
	return 0
}

type ListByRemoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NetId      string `protobuf:"bytes,1,opt,name=netId,proto3" json:"netId,omitempty"`
	Proto      string `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
	RemoteAddr string `protobuf:"bytes,3,opt,name=remoteAddr,proto3" json:"remoteAddr,omitempty"`
	RemotePort uint32 `protobuf:"varint,4,opt,name=remotePort,proto3" json:"remotePort,omitempty"`
}

func (x *ListByRemoteRequest) Reset() {
	*x = ListByRemoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forwarder_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListByRemoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListByRemoteRequest) ProtoMessage() {}

func (x *ListByRemoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_forwarder_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListByRemoteRequest.ProtoReflect.Descriptor instead.
func (*ListByRemoteRequest) Descriptor() ([]byte, []int) {
	return file_forwarder_proto_rawDescGZIP(), []int{4}
}

func (x *ListByRemoteRequest) GetNetId() string {
	if x != nil {
		return x.NetId
	}
	return ""
}

func (x *ListByRemoteRequest) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *ListByRemoteRequest) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *ListByRemoteRequest) GetRemotePort() uint32 {
	if x != nil {
		return x.RemotePort
	}
	return 0
}

type ListByRemoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Forwards []*OpenResponse `protobuf:"bytes,1,rep,name=Forwards,proto3" json:"Forwards,omitempty"`
}

func (x *ListByRemoteResponse) Reset() {
	*x = ListByRemoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_forwarder_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListByRemoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListByRemoteResponse) ProtoMessage() {}

func (x *ListByRemoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_forwarder_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListByRemoteResponse.ProtoReflect.Descriptor instead.
func (*ListByRemoteResponse) Descriptor() ([]byte, []int) {
	return file_forwarder_proto_rawDescGZIP(), []int{5}
}

func (x *ListByRemoteResponse) GetForwards() []*OpenResponse {
	if x != nil {
		return x.Forwards
	}
	return nil
}

var File_forwarder_proto protoreflect.FileDescriptor

var file_forwarder_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xb1, 0x01, 0x0a, 0x0b, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x62, 0x69, 0x6e, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x08, 0x62, 0x69, 0x6e, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x22, 0xb2, 0x01, 0x0a, 0x0c, 0x4f, 0x70,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x65,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x41, 0x64,
	0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x41, 0x64,
	0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x72,
	0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e,
	0x65, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69,
	0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x69,
	0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x50, 0x6f,
	0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x50, 0x6f,
	0x72, 0x74, 0x22, 0x73, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x0a, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x62,
	0x69, 0x6e, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x62,
	0x69, 0x6e, 0x64, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x81, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x79, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x22, 0x44, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x73, 0x32, 0xa7, 0x01, 0x0a, 0x09, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x29, 0x0a, 0x04, 0x4f, 0x70, 0x65, 0x6e, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70, 0x65,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4f, 0x70,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x05, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x79, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x52, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x52, 0x65, 0x6d,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x39, 0x5a, 0x37, 0x79,
	0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x2e, 0x69, 0x6f, 0x2f, 0x78, 0x2f, 0x6f, 0x6e, 0x65, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x68, 0x6f, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2f,
	0x67, 0x75, 0x65, 0x73, 0x74, 0x6d, 0x61, 0x6e, 0x2f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_forwarder_proto_rawDescOnce sync.Once
	file_forwarder_proto_rawDescData = file_forwarder_proto_rawDesc
)

func file_forwarder_proto_rawDescGZIP() []byte {
	file_forwarder_proto_rawDescOnce.Do(func() {
		file_forwarder_proto_rawDescData = protoimpl.X.CompressGZIP(file_forwarder_proto_rawDescData)
	})
	return file_forwarder_proto_rawDescData
}

var file_forwarder_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_forwarder_proto_goTypes = []interface{}{
	(*OpenRequest)(nil),          // 0: pb.OpenRequest
	(*OpenResponse)(nil),         // 1: pb.OpenResponse
	(*CloseRequest)(nil),         // 2: pb.CloseRequest
	(*CloseResponse)(nil),        // 3: pb.CloseResponse
	(*ListByRemoteRequest)(nil),  // 4: pb.ListByRemoteRequest
	(*ListByRemoteResponse)(nil), // 5: pb.ListByRemoteResponse
}
var file_forwarder_proto_depIdxs = []int32{
	1, // 0: pb.ListByRemoteResponse.Forwards:type_name -> pb.OpenResponse
	0, // 1: pb.Forwarder.Open:input_type -> pb.OpenRequest
	2, // 2: pb.Forwarder.Close:input_type -> pb.CloseRequest
	4, // 3: pb.Forwarder.ListByRemote:input_type -> pb.ListByRemoteRequest
	1, // 4: pb.Forwarder.Open:output_type -> pb.OpenResponse
	3, // 5: pb.Forwarder.Close:output_type -> pb.CloseResponse
	5, // 6: pb.Forwarder.ListByRemote:output_type -> pb.ListByRemoteResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_forwarder_proto_init() }
func file_forwarder_proto_init() {
	if File_forwarder_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_forwarder_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenRequest); i {
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
		file_forwarder_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OpenResponse); i {
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
		file_forwarder_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseRequest); i {
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
		file_forwarder_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseResponse); i {
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
		file_forwarder_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListByRemoteRequest); i {
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
		file_forwarder_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListByRemoteResponse); i {
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
			RawDescriptor: file_forwarder_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_forwarder_proto_goTypes,
		DependencyIndexes: file_forwarder_proto_depIdxs,
		MessageInfos:      file_forwarder_proto_msgTypes,
	}.Build()
	File_forwarder_proto = out.File
	file_forwarder_proto_rawDesc = nil
	file_forwarder_proto_goTypes = nil
	file_forwarder_proto_depIdxs = nil
}
