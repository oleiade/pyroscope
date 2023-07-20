// SPDX-License-Identifier: AGPL-3.0-only
// Provenance-includes-location: https://github.com/cortexproject/cortex/blob/master/pkg/scheduler/schedulerpb/scheduler.proto
// Provenance-includes-license: Apache-2.0
// Provenance-includes-copyright: The Cortex Authors.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: scheduler/schedulerpb/scheduler.proto

package schedulerpb

import (
	httpgrpc "github.com/grafana/pyroscope/pkg/util/httpgrpc"
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

type FrontendToSchedulerType int32

const (
	FrontendToSchedulerType_INIT    FrontendToSchedulerType = 0
	FrontendToSchedulerType_ENQUEUE FrontendToSchedulerType = 1
	FrontendToSchedulerType_CANCEL  FrontendToSchedulerType = 2
)

// Enum value maps for FrontendToSchedulerType.
var (
	FrontendToSchedulerType_name = map[int32]string{
		0: "INIT",
		1: "ENQUEUE",
		2: "CANCEL",
	}
	FrontendToSchedulerType_value = map[string]int32{
		"INIT":    0,
		"ENQUEUE": 1,
		"CANCEL":  2,
	}
)

func (x FrontendToSchedulerType) Enum() *FrontendToSchedulerType {
	p := new(FrontendToSchedulerType)
	*p = x
	return p
}

func (x FrontendToSchedulerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FrontendToSchedulerType) Descriptor() protoreflect.EnumDescriptor {
	return file_scheduler_schedulerpb_scheduler_proto_enumTypes[0].Descriptor()
}

func (FrontendToSchedulerType) Type() protoreflect.EnumType {
	return &file_scheduler_schedulerpb_scheduler_proto_enumTypes[0]
}

func (x FrontendToSchedulerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FrontendToSchedulerType.Descriptor instead.
func (FrontendToSchedulerType) EnumDescriptor() ([]byte, []int) {
	return file_scheduler_schedulerpb_scheduler_proto_rawDescGZIP(), []int{0}
}

type SchedulerToFrontendStatus int32

const (
	SchedulerToFrontendStatus_OK                           SchedulerToFrontendStatus = 0
	SchedulerToFrontendStatus_TOO_MANY_REQUESTS_PER_TENANT SchedulerToFrontendStatus = 1
	SchedulerToFrontendStatus_ERROR                        SchedulerToFrontendStatus = 2
	SchedulerToFrontendStatus_SHUTTING_DOWN                SchedulerToFrontendStatus = 3
)

// Enum value maps for SchedulerToFrontendStatus.
var (
	SchedulerToFrontendStatus_name = map[int32]string{
		0: "OK",
		1: "TOO_MANY_REQUESTS_PER_TENANT",
		2: "ERROR",
		3: "SHUTTING_DOWN",
	}
	SchedulerToFrontendStatus_value = map[string]int32{
		"OK":                           0,
		"TOO_MANY_REQUESTS_PER_TENANT": 1,
		"ERROR":                        2,
		"SHUTTING_DOWN":                3,
	}
)

func (x SchedulerToFrontendStatus) Enum() *SchedulerToFrontendStatus {
	p := new(SchedulerToFrontendStatus)
	*p = x
	return p
}

func (x SchedulerToFrontendStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SchedulerToFrontendStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_scheduler_schedulerpb_scheduler_proto_enumTypes[1].Descriptor()
}

func (SchedulerToFrontendStatus) Type() protoreflect.EnumType {
	return &file_scheduler_schedulerpb_scheduler_proto_enumTypes[1]
}

func (x SchedulerToFrontendStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SchedulerToFrontendStatus.Descriptor instead.
func (SchedulerToFrontendStatus) EnumDescriptor() ([]byte, []int) {
	return file_scheduler_schedulerpb_scheduler_proto_rawDescGZIP(), []int{1}
}

// Querier reports its own clientID when it connects, so that scheduler knows how many *different* queriers are connected.
// To signal that querier is ready to accept another request, querier sends empty message.
type QuerierToScheduler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuerierID string `protobuf:"bytes,1,opt,name=querierID,proto3" json:"querierID,omitempty"`
}

func (x *QuerierToScheduler) Reset() {
	*x = QuerierToScheduler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_schedulerpb_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuerierToScheduler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuerierToScheduler) ProtoMessage() {}

func (x *QuerierToScheduler) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_schedulerpb_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuerierToScheduler.ProtoReflect.Descriptor instead.
func (*QuerierToScheduler) Descriptor() ([]byte, []int) {
	return file_scheduler_schedulerpb_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *QuerierToScheduler) GetQuerierID() string {
	if x != nil {
		return x.QuerierID
	}
	return ""
}

type SchedulerToQuerier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Query ID as reported by frontend. When querier sends the response back to frontend (using frontendAddress),
	// it identifies the query by using this ID.
	QueryID     uint64                `protobuf:"varint,1,opt,name=queryID,proto3" json:"queryID,omitempty"`
	HttpRequest *httpgrpc.HTTPRequest `protobuf:"bytes,2,opt,name=httpRequest,proto3" json:"httpRequest,omitempty"`
	// Where should querier send HTTP Response to (using FrontendForQuerier interface).
	FrontendAddress string `protobuf:"bytes,3,opt,name=frontendAddress,proto3" json:"frontendAddress,omitempty"`
	// User who initiated the request. Needed to send reply back to frontend.
	UserID string `protobuf:"bytes,4,opt,name=userID,proto3" json:"userID,omitempty"`
	// Whether query statistics tracking should be enabled. The response will include
	// statistics only when this option is enabled.
	StatsEnabled bool `protobuf:"varint,5,opt,name=statsEnabled,proto3" json:"statsEnabled,omitempty"`
}

func (x *SchedulerToQuerier) Reset() {
	*x = SchedulerToQuerier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_schedulerpb_scheduler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerToQuerier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerToQuerier) ProtoMessage() {}

func (x *SchedulerToQuerier) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_schedulerpb_scheduler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerToQuerier.ProtoReflect.Descriptor instead.
func (*SchedulerToQuerier) Descriptor() ([]byte, []int) {
	return file_scheduler_schedulerpb_scheduler_proto_rawDescGZIP(), []int{1}
}

func (x *SchedulerToQuerier) GetQueryID() uint64 {
	if x != nil {
		return x.QueryID
	}
	return 0
}

func (x *SchedulerToQuerier) GetHttpRequest() *httpgrpc.HTTPRequest {
	if x != nil {
		return x.HttpRequest
	}
	return nil
}

func (x *SchedulerToQuerier) GetFrontendAddress() string {
	if x != nil {
		return x.FrontendAddress
	}
	return ""
}

func (x *SchedulerToQuerier) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *SchedulerToQuerier) GetStatsEnabled() bool {
	if x != nil {
		return x.StatsEnabled
	}
	return false
}

type FrontendToScheduler struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type FrontendToSchedulerType `protobuf:"varint,1,opt,name=type,proto3,enum=schedulerpb.FrontendToSchedulerType" json:"type,omitempty"`
	// Used by INIT message. Will be put into all requests passed to querier.
	FrontendAddress string `protobuf:"bytes,2,opt,name=frontendAddress,proto3" json:"frontendAddress,omitempty"`
	// Used by ENQUEUE and CANCEL.
	// Each frontend manages its own queryIDs. Different frontends may use same set of query IDs.
	QueryID uint64 `protobuf:"varint,3,opt,name=queryID,proto3" json:"queryID,omitempty"`
	// Following are used by ENQUEUE only.
	UserID       string                `protobuf:"bytes,4,opt,name=userID,proto3" json:"userID,omitempty"`
	HttpRequest  *httpgrpc.HTTPRequest `protobuf:"bytes,5,opt,name=httpRequest,proto3" json:"httpRequest,omitempty"`
	StatsEnabled bool                  `protobuf:"varint,6,opt,name=statsEnabled,proto3" json:"statsEnabled,omitempty"`
}

func (x *FrontendToScheduler) Reset() {
	*x = FrontendToScheduler{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_schedulerpb_scheduler_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrontendToScheduler) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontendToScheduler) ProtoMessage() {}

func (x *FrontendToScheduler) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_schedulerpb_scheduler_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontendToScheduler.ProtoReflect.Descriptor instead.
func (*FrontendToScheduler) Descriptor() ([]byte, []int) {
	return file_scheduler_schedulerpb_scheduler_proto_rawDescGZIP(), []int{2}
}

func (x *FrontendToScheduler) GetType() FrontendToSchedulerType {
	if x != nil {
		return x.Type
	}
	return FrontendToSchedulerType_INIT
}

func (x *FrontendToScheduler) GetFrontendAddress() string {
	if x != nil {
		return x.FrontendAddress
	}
	return ""
}

func (x *FrontendToScheduler) GetQueryID() uint64 {
	if x != nil {
		return x.QueryID
	}
	return 0
}

func (x *FrontendToScheduler) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *FrontendToScheduler) GetHttpRequest() *httpgrpc.HTTPRequest {
	if x != nil {
		return x.HttpRequest
	}
	return nil
}

func (x *FrontendToScheduler) GetStatsEnabled() bool {
	if x != nil {
		return x.StatsEnabled
	}
	return false
}

type SchedulerToFrontend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status SchedulerToFrontendStatus `protobuf:"varint,1,opt,name=status,proto3,enum=schedulerpb.SchedulerToFrontendStatus" json:"status,omitempty"`
	Error  string                    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *SchedulerToFrontend) Reset() {
	*x = SchedulerToFrontend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_schedulerpb_scheduler_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerToFrontend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerToFrontend) ProtoMessage() {}

func (x *SchedulerToFrontend) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_schedulerpb_scheduler_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerToFrontend.ProtoReflect.Descriptor instead.
func (*SchedulerToFrontend) Descriptor() ([]byte, []int) {
	return file_scheduler_schedulerpb_scheduler_proto_rawDescGZIP(), []int{3}
}

func (x *SchedulerToFrontend) GetStatus() SchedulerToFrontendStatus {
	if x != nil {
		return x.Status
	}
	return SchedulerToFrontendStatus_OK
}

func (x *SchedulerToFrontend) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type NotifyQuerierShutdownRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QuerierID string `protobuf:"bytes,1,opt,name=querierID,proto3" json:"querierID,omitempty"`
}

func (x *NotifyQuerierShutdownRequest) Reset() {
	*x = NotifyQuerierShutdownRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_schedulerpb_scheduler_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyQuerierShutdownRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyQuerierShutdownRequest) ProtoMessage() {}

func (x *NotifyQuerierShutdownRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_schedulerpb_scheduler_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyQuerierShutdownRequest.ProtoReflect.Descriptor instead.
func (*NotifyQuerierShutdownRequest) Descriptor() ([]byte, []int) {
	return file_scheduler_schedulerpb_scheduler_proto_rawDescGZIP(), []int{4}
}

func (x *NotifyQuerierShutdownRequest) GetQuerierID() string {
	if x != nil {
		return x.QuerierID
	}
	return ""
}

type NotifyQuerierShutdownResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NotifyQuerierShutdownResponse) Reset() {
	*x = NotifyQuerierShutdownResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scheduler_schedulerpb_scheduler_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NotifyQuerierShutdownResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotifyQuerierShutdownResponse) ProtoMessage() {}

func (x *NotifyQuerierShutdownResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scheduler_schedulerpb_scheduler_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotifyQuerierShutdownResponse.ProtoReflect.Descriptor instead.
func (*NotifyQuerierShutdownResponse) Descriptor() ([]byte, []int) {
	return file_scheduler_schedulerpb_scheduler_proto_rawDescGZIP(), []int{5}
}

var File_scheduler_schedulerpb_scheduler_proto protoreflect.FileDescriptor

var file_scheduler_schedulerpb_scheduler_proto_rawDesc = []byte{
	0x0a, 0x25, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2f, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x70, 0x62, 0x1a, 0x1c, 0x75, 0x74, 0x69, 0x6c, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x32, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x54, 0x6f, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x72,
	0x69, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x65,
	0x72, 0x69, 0x65, 0x72, 0x49, 0x44, 0x22, 0xcd, 0x01, 0x0a, 0x12, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x54, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x12, 0x37, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x28, 0x0a, 0x0f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x66, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x73, 0x45,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x88, 0x02, 0x0a, 0x13, 0x46, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x38,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x46, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x66, 0x72, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x71, 0x75, 0x65, 0x72, 0x79, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x37, 0x0a, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x0b, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x73, 0x74, 0x61, 0x74, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x73, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x22, 0x6b, 0x0a, 0x13, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x54, 0x6f,
	0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x12, 0x3e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x54, 0x6f, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3c,
	0x0a, 0x1c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x53,
	0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x49, 0x44, 0x22, 0x1f, 0x0a, 0x1d,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x53, 0x68, 0x75,
	0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x3c, 0x0a,
	0x17, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x54, 0x6f, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x49, 0x54,
	0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x51, 0x55, 0x45, 0x55, 0x45, 0x10, 0x01, 0x12,
	0x0a, 0x0a, 0x06, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x10, 0x02, 0x2a, 0x63, 0x0a, 0x19, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x54, 0x6f, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00,
	0x12, 0x20, 0x0a, 0x1c, 0x54, 0x4f, 0x4f, 0x5f, 0x4d, 0x41, 0x4e, 0x59, 0x5f, 0x52, 0x45, 0x51,
	0x55, 0x45, 0x53, 0x54, 0x53, 0x5f, 0x50, 0x45, 0x52, 0x5f, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54,
	0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x11, 0x0a,
	0x0d, 0x53, 0x48, 0x55, 0x54, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x03,
	0x32, 0xdc, 0x01, 0x0a, 0x13, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x46, 0x6f,
	0x72, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x12, 0x55, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72,
	0x69, 0x65, 0x72, 0x4c, 0x6f, 0x6f, 0x70, 0x12, 0x1f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x54, 0x6f, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x1a, 0x1f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x54, 0x6f, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12,
	0x6e, 0x0a, 0x15, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72,
	0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x12, 0x29, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x51, 0x75, 0x65,
	0x72, 0x69, 0x65, 0x72, 0x53, 0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70,
	0x62, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x51, 0x75, 0x65, 0x72, 0x69, 0x65, 0x72, 0x53,
	0x68, 0x75, 0x74, 0x64, 0x6f, 0x77, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x70, 0x0a, 0x14, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x46,
	0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x12, 0x58, 0x0a, 0x0c, 0x46, 0x72, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x64, 0x4c, 0x6f, 0x6f, 0x70, 0x12, 0x20, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x54, 0x6f,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x1a, 0x20, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x54, 0x6f, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x42, 0xa5, 0x01, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x70, 0x62, 0x42, 0x0e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x79, 0x72, 0x6f,
	0x73, 0x63, 0x6f, 0x70, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x72, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0xa2,
	0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x70, 0x62, 0xca, 0x02, 0x0b, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70,
	0x62, 0xe2, 0x02, 0x17, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_scheduler_schedulerpb_scheduler_proto_rawDescOnce sync.Once
	file_scheduler_schedulerpb_scheduler_proto_rawDescData = file_scheduler_schedulerpb_scheduler_proto_rawDesc
)

func file_scheduler_schedulerpb_scheduler_proto_rawDescGZIP() []byte {
	file_scheduler_schedulerpb_scheduler_proto_rawDescOnce.Do(func() {
		file_scheduler_schedulerpb_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_scheduler_schedulerpb_scheduler_proto_rawDescData)
	})
	return file_scheduler_schedulerpb_scheduler_proto_rawDescData
}

var file_scheduler_schedulerpb_scheduler_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_scheduler_schedulerpb_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_scheduler_schedulerpb_scheduler_proto_goTypes = []interface{}{
	(FrontendToSchedulerType)(0),          // 0: schedulerpb.FrontendToSchedulerType
	(SchedulerToFrontendStatus)(0),        // 1: schedulerpb.SchedulerToFrontendStatus
	(*QuerierToScheduler)(nil),            // 2: schedulerpb.QuerierToScheduler
	(*SchedulerToQuerier)(nil),            // 3: schedulerpb.SchedulerToQuerier
	(*FrontendToScheduler)(nil),           // 4: schedulerpb.FrontendToScheduler
	(*SchedulerToFrontend)(nil),           // 5: schedulerpb.SchedulerToFrontend
	(*NotifyQuerierShutdownRequest)(nil),  // 6: schedulerpb.NotifyQuerierShutdownRequest
	(*NotifyQuerierShutdownResponse)(nil), // 7: schedulerpb.NotifyQuerierShutdownResponse
	(*httpgrpc.HTTPRequest)(nil),          // 8: httpgrpc.HTTPRequest
}
var file_scheduler_schedulerpb_scheduler_proto_depIdxs = []int32{
	8, // 0: schedulerpb.SchedulerToQuerier.httpRequest:type_name -> httpgrpc.HTTPRequest
	0, // 1: schedulerpb.FrontendToScheduler.type:type_name -> schedulerpb.FrontendToSchedulerType
	8, // 2: schedulerpb.FrontendToScheduler.httpRequest:type_name -> httpgrpc.HTTPRequest
	1, // 3: schedulerpb.SchedulerToFrontend.status:type_name -> schedulerpb.SchedulerToFrontendStatus
	2, // 4: schedulerpb.SchedulerForQuerier.QuerierLoop:input_type -> schedulerpb.QuerierToScheduler
	6, // 5: schedulerpb.SchedulerForQuerier.NotifyQuerierShutdown:input_type -> schedulerpb.NotifyQuerierShutdownRequest
	4, // 6: schedulerpb.SchedulerForFrontend.FrontendLoop:input_type -> schedulerpb.FrontendToScheduler
	3, // 7: schedulerpb.SchedulerForQuerier.QuerierLoop:output_type -> schedulerpb.SchedulerToQuerier
	7, // 8: schedulerpb.SchedulerForQuerier.NotifyQuerierShutdown:output_type -> schedulerpb.NotifyQuerierShutdownResponse
	5, // 9: schedulerpb.SchedulerForFrontend.FrontendLoop:output_type -> schedulerpb.SchedulerToFrontend
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_scheduler_schedulerpb_scheduler_proto_init() }
func file_scheduler_schedulerpb_scheduler_proto_init() {
	if File_scheduler_schedulerpb_scheduler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scheduler_schedulerpb_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuerierToScheduler); i {
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
		file_scheduler_schedulerpb_scheduler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerToQuerier); i {
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
		file_scheduler_schedulerpb_scheduler_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrontendToScheduler); i {
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
		file_scheduler_schedulerpb_scheduler_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerToFrontend); i {
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
		file_scheduler_schedulerpb_scheduler_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyQuerierShutdownRequest); i {
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
		file_scheduler_schedulerpb_scheduler_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NotifyQuerierShutdownResponse); i {
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
			RawDescriptor: file_scheduler_schedulerpb_scheduler_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_scheduler_schedulerpb_scheduler_proto_goTypes,
		DependencyIndexes: file_scheduler_schedulerpb_scheduler_proto_depIdxs,
		EnumInfos:         file_scheduler_schedulerpb_scheduler_proto_enumTypes,
		MessageInfos:      file_scheduler_schedulerpb_scheduler_proto_msgTypes,
	}.Build()
	File_scheduler_schedulerpb_scheduler_proto = out.File
	file_scheduler_schedulerpb_scheduler_proto_rawDesc = nil
	file_scheduler_schedulerpb_scheduler_proto_goTypes = nil
	file_scheduler_schedulerpb_scheduler_proto_depIdxs = nil
}
