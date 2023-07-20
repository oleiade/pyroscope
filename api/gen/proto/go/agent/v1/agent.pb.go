// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: agent/v1/agent.proto

package agentv1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Health int32

const (
	Health_HEALTH_UNSPECIFIED Health = 0
	Health_HEALTH_UP          Health = 1
	Health_HEALTH_DOWN        Health = 2
)

// Enum value maps for Health.
var (
	Health_name = map[int32]string{
		0: "HEALTH_UNSPECIFIED",
		1: "HEALTH_UP",
		2: "HEALTH_DOWN",
	}
	Health_value = map[string]int32{
		"HEALTH_UNSPECIFIED": 0,
		"HEALTH_UP":          1,
		"HEALTH_DOWN":        2,
	}
)

func (x Health) Enum() *Health {
	p := new(Health)
	*p = x
	return p
}

func (x Health) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Health) Descriptor() protoreflect.EnumDescriptor {
	return file_agent_v1_agent_proto_enumTypes[0].Descriptor()
}

func (Health) Type() protoreflect.EnumType {
	return &file_agent_v1_agent_proto_enumTypes[0]
}

func (x Health) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Health.Descriptor instead.
func (Health) EnumDescriptor() ([]byte, []int) {
	return file_agent_v1_agent_proto_rawDescGZIP(), []int{0}
}

type State int32

const (
	State_STATE_UNSPECIFIED State = 0
	State_STATE_ACTIVE      State = 1
	State_STATE_DROPPED     State = 2
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STATE_ACTIVE",
		2: "STATE_DROPPED",
	}
	State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"STATE_ACTIVE":      1,
		"STATE_DROPPED":     2,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_agent_v1_agent_proto_enumTypes[1].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_agent_v1_agent_proto_enumTypes[1]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_agent_v1_agent_proto_rawDescGZIP(), []int{1}
}

type GetTargetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State State `protobuf:"varint,1,opt,name=state,proto3,enum=agent.v1.State" json:"state,omitempty"`
}

func (x *GetTargetsRequest) Reset() {
	*x = GetTargetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_v1_agent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTargetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTargetsRequest) ProtoMessage() {}

func (x *GetTargetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_agent_v1_agent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTargetsRequest.ProtoReflect.Descriptor instead.
func (*GetTargetsRequest) Descriptor() ([]byte, []int) {
	return file_agent_v1_agent_proto_rawDescGZIP(), []int{0}
}

func (x *GetTargetsRequest) GetState() State {
	if x != nil {
		return x.State
	}
	return State_STATE_UNSPECIFIED
}

type GetTargetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActiveTargets  []*Target `protobuf:"bytes,1,rep,name=active_targets,json=activeTargets,proto3" json:"active_targets,omitempty"`
	DroppedTargets []*Target `protobuf:"bytes,2,rep,name=dropped_targets,json=droppedTargets,proto3" json:"dropped_targets,omitempty"`
}

func (x *GetTargetsResponse) Reset() {
	*x = GetTargetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_v1_agent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTargetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTargetsResponse) ProtoMessage() {}

func (x *GetTargetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_agent_v1_agent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTargetsResponse.ProtoReflect.Descriptor instead.
func (*GetTargetsResponse) Descriptor() ([]byte, []int) {
	return file_agent_v1_agent_proto_rawDescGZIP(), []int{1}
}

func (x *GetTargetsResponse) GetActiveTargets() []*Target {
	if x != nil {
		return x.ActiveTargets
	}
	return nil
}

func (x *GetTargetsResponse) GetDroppedTargets() []*Target {
	if x != nil {
		return x.DroppedTargets
	}
	return nil
}

type Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Contains labels before any processing.
	DiscoveredLabels map[string]string `protobuf:"bytes,1,rep,name=discovered_labels,json=discoveredLabels,proto3" json:"discovered_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Labels as they are attached to ingested profiles.
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Name of the scrape pool used
	ScrapePool string `protobuf:"bytes,3,opt,name=scrape_pool,json=scrapePool,proto3" json:"scrape_pool,omitempty"`
	// URL that is used for retrieving the profile.
	ScrapeUrl string `protobuf:"bytes,4,opt,name=scrape_url,json=scrapeUrl,proto3" json:"scrape_url,omitempty"`
	// Contains the error if the last attempted scrape has failed.
	LastError string `protobuf:"bytes,5,opt,name=last_error,json=lastError,proto3" json:"last_error,omitempty"`
	// Timestamp of last scrape
	LastScrape *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=last_scrape,json=lastScrape,proto3" json:"last_scrape,omitempty"`
	// Duation of last scrape.
	LastScrapeDuration *durationpb.Duration `protobuf:"bytes,7,opt,name=last_scrape_duration,json=lastScrapeDuration,proto3" json:"last_scrape_duration,omitempty"`
	// Health of last scrape.
	Health Health `protobuf:"varint,8,opt,name=health,proto3,enum=agent.v1.Health" json:"health,omitempty"`
	// Timeout duration for the scrape request.
	ScrapeTimeout *durationpb.Duration `protobuf:"bytes,9,opt,name=scrape_timeout,json=scrapeTimeout,proto3" json:"scrape_timeout,omitempty"`
	// Interval how often profiles are scraped.
	ScrapeInterval *durationpb.Duration `protobuf:"bytes,10,opt,name=scrape_interval,json=scrapeInterval,proto3" json:"scrape_interval,omitempty"`
}

func (x *Target) Reset() {
	*x = Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_agent_v1_agent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target) ProtoMessage() {}

func (x *Target) ProtoReflect() protoreflect.Message {
	mi := &file_agent_v1_agent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target.ProtoReflect.Descriptor instead.
func (*Target) Descriptor() ([]byte, []int) {
	return file_agent_v1_agent_proto_rawDescGZIP(), []int{2}
}

func (x *Target) GetDiscoveredLabels() map[string]string {
	if x != nil {
		return x.DiscoveredLabels
	}
	return nil
}

func (x *Target) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Target) GetScrapePool() string {
	if x != nil {
		return x.ScrapePool
	}
	return ""
}

func (x *Target) GetScrapeUrl() string {
	if x != nil {
		return x.ScrapeUrl
	}
	return ""
}

func (x *Target) GetLastError() string {
	if x != nil {
		return x.LastError
	}
	return ""
}

func (x *Target) GetLastScrape() *timestamppb.Timestamp {
	if x != nil {
		return x.LastScrape
	}
	return nil
}

func (x *Target) GetLastScrapeDuration() *durationpb.Duration {
	if x != nil {
		return x.LastScrapeDuration
	}
	return nil
}

func (x *Target) GetHealth() Health {
	if x != nil {
		return x.Health
	}
	return Health_HEALTH_UNSPECIFIED
}

func (x *Target) GetScrapeTimeout() *durationpb.Duration {
	if x != nil {
		return x.ScrapeTimeout
	}
	return nil
}

func (x *Target) GetScrapeInterval() *durationpb.Duration {
	if x != nil {
		return x.ScrapeInterval
	}
	return nil
}

var File_agent_v1_agent_proto protoreflect.FileDescriptor

var file_agent_v1_agent_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x40, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x04, 0xe2, 0x41, 0x01, 0x01, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0e, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x52, 0x0d, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x73, 0x12, 0x39, 0x0a, 0x0f, 0x64, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x52, 0x0e, 0x64,
	0x72, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x22, 0xac, 0x05,
	0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x53, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x34, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x5f, 0x70, 0x6f,
	0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65,
	0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65,
	0x55, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x63, 0x72, 0x61, 0x70,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x63, 0x72, 0x61, 0x70, 0x65, 0x12,
	0x4b, 0x0a, 0x14, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65, 0x5f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x63,
	0x72, 0x61, 0x70, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x52, 0x06,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x40, 0x0a, 0x0e, 0x73, 0x63, 0x72, 0x61, 0x70, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x73, 0x63, 0x72, 0x61, 0x70,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x42, 0x0a, 0x0f, 0x73, 0x63, 0x72, 0x61,
	0x70, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x73, 0x63,
	0x72, 0x61, 0x70, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x1a, 0x43, 0x0a, 0x15,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x40, 0x0a, 0x06,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x12, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d,
	0x0a, 0x09, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x55, 0x50, 0x10, 0x01, 0x12, 0x0f, 0x0a,
	0x0b, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x5f, 0x44, 0x4f, 0x57, 0x4e, 0x10, 0x02, 0x2a, 0x43,
	0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01,
	0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x44, 0x52, 0x4f, 0x50, 0x50, 0x45,
	0x44, 0x10, 0x02, 0x32, 0x70, 0x0a, 0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x60, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x73, 0x12, 0x1b, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x42, 0x9b, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x67, 0x72, 0x61, 0x66, 0x61, 0x6e, 0x61, 0x2f, 0x70, 0x79, 0x72, 0x6f, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x14, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_agent_v1_agent_proto_rawDescOnce sync.Once
	file_agent_v1_agent_proto_rawDescData = file_agent_v1_agent_proto_rawDesc
)

func file_agent_v1_agent_proto_rawDescGZIP() []byte {
	file_agent_v1_agent_proto_rawDescOnce.Do(func() {
		file_agent_v1_agent_proto_rawDescData = protoimpl.X.CompressGZIP(file_agent_v1_agent_proto_rawDescData)
	})
	return file_agent_v1_agent_proto_rawDescData
}

var file_agent_v1_agent_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_agent_v1_agent_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_agent_v1_agent_proto_goTypes = []interface{}{
	(Health)(0),                   // 0: agent.v1.Health
	(State)(0),                    // 1: agent.v1.State
	(*GetTargetsRequest)(nil),     // 2: agent.v1.GetTargetsRequest
	(*GetTargetsResponse)(nil),    // 3: agent.v1.GetTargetsResponse
	(*Target)(nil),                // 4: agent.v1.Target
	nil,                           // 5: agent.v1.Target.DiscoveredLabelsEntry
	nil,                           // 6: agent.v1.Target.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 8: google.protobuf.Duration
}
var file_agent_v1_agent_proto_depIdxs = []int32{
	1,  // 0: agent.v1.GetTargetsRequest.state:type_name -> agent.v1.State
	4,  // 1: agent.v1.GetTargetsResponse.active_targets:type_name -> agent.v1.Target
	4,  // 2: agent.v1.GetTargetsResponse.dropped_targets:type_name -> agent.v1.Target
	5,  // 3: agent.v1.Target.discovered_labels:type_name -> agent.v1.Target.DiscoveredLabelsEntry
	6,  // 4: agent.v1.Target.labels:type_name -> agent.v1.Target.LabelsEntry
	7,  // 5: agent.v1.Target.last_scrape:type_name -> google.protobuf.Timestamp
	8,  // 6: agent.v1.Target.last_scrape_duration:type_name -> google.protobuf.Duration
	0,  // 7: agent.v1.Target.health:type_name -> agent.v1.Health
	8,  // 8: agent.v1.Target.scrape_timeout:type_name -> google.protobuf.Duration
	8,  // 9: agent.v1.Target.scrape_interval:type_name -> google.protobuf.Duration
	2,  // 10: agent.v1.AgentService.GetTargets:input_type -> agent.v1.GetTargetsRequest
	3,  // 11: agent.v1.AgentService.GetTargets:output_type -> agent.v1.GetTargetsResponse
	11, // [11:12] is the sub-list for method output_type
	10, // [10:11] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_agent_v1_agent_proto_init() }
func file_agent_v1_agent_proto_init() {
	if File_agent_v1_agent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_agent_v1_agent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTargetsRequest); i {
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
		file_agent_v1_agent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTargetsResponse); i {
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
		file_agent_v1_agent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Target); i {
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
			RawDescriptor: file_agent_v1_agent_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_agent_v1_agent_proto_goTypes,
		DependencyIndexes: file_agent_v1_agent_proto_depIdxs,
		EnumInfos:         file_agent_v1_agent_proto_enumTypes,
		MessageInfos:      file_agent_v1_agent_proto_msgTypes,
	}.Build()
	File_agent_v1_agent_proto = out.File
	file_agent_v1_agent_proto_rawDesc = nil
	file_agent_v1_agent_proto_goTypes = nil
	file_agent_v1_agent_proto_depIdxs = nil
}
