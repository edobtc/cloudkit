// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: relay/config/v1/config.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/anypb"
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

type ListenerType int32

const (
	ListenerType_LISTENER_TYPE_UNSPECIFIED ListenerType = 0
	ListenerType_LISTENER_TYPE_HTTP        ListenerType = 1
	ListenerType_LISTENER_TYPE_WEBSOCKET   ListenerType = 2
	ListenerType_LISTENER_TYPE_SNS         ListenerType = 3
	ListenerType_LISTENER_TYPE_SQS         ListenerType = 4
	ListenerType_LISTENER_TYPE_KINESIS     ListenerType = 5
	ListenerType_LISTENER_TYPE_KAFKA       ListenerType = 6
	ListenerType_LISTENER_TYPE_UDP         ListenerType = 7
)

// Enum value maps for ListenerType.
var (
	ListenerType_name = map[int32]string{
		0: "LISTENER_TYPE_UNSPECIFIED",
		1: "LISTENER_TYPE_HTTP",
		2: "LISTENER_TYPE_WEBSOCKET",
		3: "LISTENER_TYPE_SNS",
		4: "LISTENER_TYPE_SQS",
		5: "LISTENER_TYPE_KINESIS",
		6: "LISTENER_TYPE_KAFKA",
		7: "LISTENER_TYPE_UDP",
	}
	ListenerType_value = map[string]int32{
		"LISTENER_TYPE_UNSPECIFIED": 0,
		"LISTENER_TYPE_HTTP":        1,
		"LISTENER_TYPE_WEBSOCKET":   2,
		"LISTENER_TYPE_SNS":         3,
		"LISTENER_TYPE_SQS":         4,
		"LISTENER_TYPE_KINESIS":     5,
		"LISTENER_TYPE_KAFKA":       6,
		"LISTENER_TYPE_UDP":         7,
	}
)

func (x ListenerType) Enum() *ListenerType {
	p := new(ListenerType)
	*p = x
	return p
}

func (x ListenerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListenerType) Descriptor() protoreflect.EnumDescriptor {
	return file_relay_config_v1_config_proto_enumTypes[0].Descriptor()
}

func (ListenerType) Type() protoreflect.EnumType {
	return &file_relay_config_v1_config_proto_enumTypes[0]
}

func (x ListenerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListenerType.Descriptor instead.
func (ListenerType) EnumDescriptor() ([]byte, []int) {
	return file_relay_config_v1_config_proto_rawDescGZIP(), []int{0}
}

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *ConfigRequest `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_config_v1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relay_config_v1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_relay_config_v1_config_proto_rawDescGZIP(), []int{0}
}

func (x *StartRequest) GetConfig() *ConfigRequest {
	if x != nil {
		return x.Config
	}
	return nil
}

type StartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *StatusResponse `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StartResponse) Reset() {
	*x = StartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_config_v1_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartResponse) ProtoMessage() {}

func (x *StartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relay_config_v1_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartResponse.ProtoReflect.Descriptor instead.
func (*StartResponse) Descriptor() ([]byte, []int) {
	return file_relay_config_v1_config_proto_rawDescGZIP(), []int{1}
}

func (x *StartResponse) GetStatus() *StatusResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

type StopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopRequest) Reset() {
	*x = StopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_config_v1_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRequest) ProtoMessage() {}

func (x *StopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relay_config_v1_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRequest.ProtoReflect.Descriptor instead.
func (*StopRequest) Descriptor() ([]byte, []int) {
	return file_relay_config_v1_config_proto_rawDescGZIP(), []int{2}
}

type StopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *StatusResponse `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StopResponse) Reset() {
	*x = StopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_config_v1_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopResponse) ProtoMessage() {}

func (x *StopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relay_config_v1_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopResponse.ProtoReflect.Descriptor instead.
func (*StopResponse) Descriptor() ([]byte, []int) {
	return file_relay_config_v1_config_proto_rawDescGZIP(), []int{3}
}

func (x *StopResponse) GetStatus() *StatusResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

type ConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchInterval int32 `protobuf:"varint,1,opt,name=batchInterval,proto3" json:"batchInterval,omitempty"`
	BatchSize     int32 `protobuf:"varint,2,opt,name=batchSize,proto3" json:"batchSize,omitempty"`
	// Ability to disable listener type
	ListenerType []ListenerType `protobuf:"varint,3,rep,packed,name=listenerType,proto3,enum=relay.config.v1.ListenerType" json:"listenerType,omitempty"`
}

func (x *ConfigRequest) Reset() {
	*x = ConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_config_v1_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigRequest) ProtoMessage() {}

func (x *ConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relay_config_v1_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigRequest.ProtoReflect.Descriptor instead.
func (*ConfigRequest) Descriptor() ([]byte, []int) {
	return file_relay_config_v1_config_proto_rawDescGZIP(), []int{4}
}

func (x *ConfigRequest) GetBatchInterval() int32 {
	if x != nil {
		return x.BatchInterval
	}
	return 0
}

func (x *ConfigRequest) GetBatchSize() int32 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

func (x *ConfigRequest) GetListenerType() []ListenerType {
	if x != nil {
		return x.ListenerType
	}
	return nil
}

type ConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success   bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Identifer string `protobuf:"bytes,2,opt,name=identifer,proto3" json:"identifer,omitempty"`
}

func (x *ConfigResponse) Reset() {
	*x = ConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_config_v1_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigResponse) ProtoMessage() {}

func (x *ConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relay_config_v1_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigResponse.ProtoReflect.Descriptor instead.
func (*ConfigResponse) Descriptor() ([]byte, []int) {
	return file_relay_config_v1_config_proto_rawDescGZIP(), []int{5}
}

func (x *ConfigResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ConfigResponse) GetIdentifer() string {
	if x != nil {
		return x.Identifer
	}
	return ""
}

type StatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusRequest) Reset() {
	*x = StatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_config_v1_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusRequest) ProtoMessage() {}

func (x *StatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_relay_config_v1_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusRequest.ProtoReflect.Descriptor instead.
func (*StatusRequest) Descriptor() ([]byte, []int) {
	return file_relay_config_v1_config_proto_rawDescGZIP(), []int{6}
}

type StatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Processed    int32                  `protobuf:"varint,1,opt,name=processed,proto3" json:"processed,omitempty"`
	CurrentSize  int32                  `protobuf:"varint,2,opt,name=currentSize,proto3" json:"currentSize,omitempty"`
	Failed       int32                  `protobuf:"varint,3,opt,name=failed,proto3" json:"failed,omitempty"`
	ListenerType []*Listener            `protobuf:"bytes,4,rep,name=listenerType,proto3" json:"listenerType,omitempty"`
	Running      bool                   `protobuf:"varint,9,opt,name=running,proto3" json:"running,omitempty"`
	LastSent     *timestamppb.Timestamp `protobuf:"bytes,12,opt,name=lastSent,proto3" json:"lastSent,omitempty"`
	StartedAt    *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
}

func (x *StatusResponse) Reset() {
	*x = StatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_config_v1_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusResponse) ProtoMessage() {}

func (x *StatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_relay_config_v1_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusResponse.ProtoReflect.Descriptor instead.
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return file_relay_config_v1_config_proto_rawDescGZIP(), []int{7}
}

func (x *StatusResponse) GetProcessed() int32 {
	if x != nil {
		return x.Processed
	}
	return 0
}

func (x *StatusResponse) GetCurrentSize() int32 {
	if x != nil {
		return x.CurrentSize
	}
	return 0
}

func (x *StatusResponse) GetFailed() int32 {
	if x != nil {
		return x.Failed
	}
	return 0
}

func (x *StatusResponse) GetListenerType() []*Listener {
	if x != nil {
		return x.ListenerType
	}
	return nil
}

func (x *StatusResponse) GetRunning() bool {
	if x != nil {
		return x.Running
	}
	return false
}

func (x *StatusResponse) GetLastSent() *timestamppb.Timestamp {
	if x != nil {
		return x.LastSent
	}
	return nil
}

func (x *StatusResponse) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

type Listener struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListenerType ListenerType `protobuf:"varint,1,opt,name=listenerType,proto3,enum=relay.config.v1.ListenerType" json:"listenerType,omitempty"`
	Active       bool         `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty"`
}

func (x *Listener) Reset() {
	*x = Listener{}
	if protoimpl.UnsafeEnabled {
		mi := &file_relay_config_v1_config_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Listener) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Listener) ProtoMessage() {}

func (x *Listener) ProtoReflect() protoreflect.Message {
	mi := &file_relay_config_v1_config_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Listener.ProtoReflect.Descriptor instead.
func (*Listener) Descriptor() ([]byte, []int) {
	return file_relay_config_v1_config_proto_rawDescGZIP(), []int{8}
}

func (x *Listener) GetListenerType() ListenerType {
	if x != nil {
		return x.ListenerType
	}
	return ListenerType_LISTENER_TYPE_UNSPECIFIED
}

func (x *Listener) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

var File_relay_config_v1_config_proto protoreflect.FileDescriptor

var file_relay_config_v1_config_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x0c, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x06, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x72, 0x65, 0x6c, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x22, 0x48, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x0d, 0x0a, 0x0b, 0x53,
	0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x47, 0x0a, 0x0c, 0x53, 0x74,
	0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x96, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x41, 0x0a, 0x0c, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x22, 0x48, 0x0a, 0x0e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x66, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x65, 0x72, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xb3, 0x02, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61,
	0x69, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x65, 0x72, 0x52, 0x0c, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x36, 0x0a, 0x08, 0x6c,
	0x61, 0x73, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x53,
	0x65, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x65, 0x0a,
	0x08, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0c, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x65, 0x2a, 0xdb, 0x01, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x4c, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x45,
	0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4c, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x45, 0x52,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x48, 0x54, 0x54, 0x50, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17,
	0x4c, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57, 0x45,
	0x42, 0x53, 0x4f, 0x43, 0x4b, 0x45, 0x54, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x49, 0x53,
	0x54, 0x45, 0x4e, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x4e, 0x53, 0x10, 0x03,
	0x12, 0x15, 0x0a, 0x11, 0x4c, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x53, 0x51, 0x53, 0x10, 0x04, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x49, 0x53, 0x54, 0x45,
	0x4e, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x45, 0x53, 0x49, 0x53,
	0x10, 0x05, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x45, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x4b, 0x41, 0x46, 0x4b, 0x41, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x4c,
	0x49, 0x53, 0x54, 0x45, 0x4e, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x44, 0x50,
	0x10, 0x07, 0x32, 0xaa, 0x02, 0x0a, 0x05, 0x52, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x46, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x1c, 0x2e, 0x72,
	0x65, 0x6c, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x72, 0x65, 0x6c,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x12, 0x1e, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x64,
	0x6f, 0x62, 0x74, 0x63, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x6b, 0x69, 0x74, 0x2f, 0x72, 0x70,
	0x63, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_relay_config_v1_config_proto_rawDescOnce sync.Once
	file_relay_config_v1_config_proto_rawDescData = file_relay_config_v1_config_proto_rawDesc
)

func file_relay_config_v1_config_proto_rawDescGZIP() []byte {
	file_relay_config_v1_config_proto_rawDescOnce.Do(func() {
		file_relay_config_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_relay_config_v1_config_proto_rawDescData)
	})
	return file_relay_config_v1_config_proto_rawDescData
}

var file_relay_config_v1_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_relay_config_v1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_relay_config_v1_config_proto_goTypes = []interface{}{
	(ListenerType)(0),             // 0: relay.config.v1.ListenerType
	(*StartRequest)(nil),          // 1: relay.config.v1.StartRequest
	(*StartResponse)(nil),         // 2: relay.config.v1.StartResponse
	(*StopRequest)(nil),           // 3: relay.config.v1.StopRequest
	(*StopResponse)(nil),          // 4: relay.config.v1.StopResponse
	(*ConfigRequest)(nil),         // 5: relay.config.v1.ConfigRequest
	(*ConfigResponse)(nil),        // 6: relay.config.v1.ConfigResponse
	(*StatusRequest)(nil),         // 7: relay.config.v1.StatusRequest
	(*StatusResponse)(nil),        // 8: relay.config.v1.StatusResponse
	(*Listener)(nil),              // 9: relay.config.v1.Listener
	(*timestamppb.Timestamp)(nil), // 10: google.protobuf.Timestamp
}
var file_relay_config_v1_config_proto_depIdxs = []int32{
	5,  // 0: relay.config.v1.StartRequest.config:type_name -> relay.config.v1.ConfigRequest
	8,  // 1: relay.config.v1.StartResponse.status:type_name -> relay.config.v1.StatusResponse
	8,  // 2: relay.config.v1.StopResponse.status:type_name -> relay.config.v1.StatusResponse
	0,  // 3: relay.config.v1.ConfigRequest.listenerType:type_name -> relay.config.v1.ListenerType
	9,  // 4: relay.config.v1.StatusResponse.listenerType:type_name -> relay.config.v1.Listener
	10, // 5: relay.config.v1.StatusResponse.lastSent:type_name -> google.protobuf.Timestamp
	10, // 6: relay.config.v1.StatusResponse.startedAt:type_name -> google.protobuf.Timestamp
	0,  // 7: relay.config.v1.Listener.listenerType:type_name -> relay.config.v1.ListenerType
	1,  // 8: relay.config.v1.Relay.Start:input_type -> relay.config.v1.StartRequest
	3,  // 9: relay.config.v1.Relay.Stop:input_type -> relay.config.v1.StopRequest
	5,  // 10: relay.config.v1.Relay.Config:input_type -> relay.config.v1.ConfigRequest
	7,  // 11: relay.config.v1.Relay.Status:input_type -> relay.config.v1.StatusRequest
	2,  // 12: relay.config.v1.Relay.Start:output_type -> relay.config.v1.StartResponse
	4,  // 13: relay.config.v1.Relay.Stop:output_type -> relay.config.v1.StopResponse
	6,  // 14: relay.config.v1.Relay.Config:output_type -> relay.config.v1.ConfigResponse
	8,  // 15: relay.config.v1.Relay.Status:output_type -> relay.config.v1.StatusResponse
	12, // [12:16] is the sub-list for method output_type
	8,  // [8:12] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_relay_config_v1_config_proto_init() }
func file_relay_config_v1_config_proto_init() {
	if File_relay_config_v1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_relay_config_v1_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRequest); i {
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
		file_relay_config_v1_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartResponse); i {
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
		file_relay_config_v1_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopRequest); i {
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
		file_relay_config_v1_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopResponse); i {
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
		file_relay_config_v1_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigRequest); i {
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
		file_relay_config_v1_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigResponse); i {
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
		file_relay_config_v1_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusRequest); i {
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
		file_relay_config_v1_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusResponse); i {
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
		file_relay_config_v1_config_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Listener); i {
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
			RawDescriptor: file_relay_config_v1_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_relay_config_v1_config_proto_goTypes,
		DependencyIndexes: file_relay_config_v1_config_proto_depIdxs,
		EnumInfos:         file_relay_config_v1_config_proto_enumTypes,
		MessageInfos:      file_relay_config_v1_config_proto_msgTypes,
	}.Build()
	File_relay_config_v1_config_proto = out.File
	file_relay_config_v1_config_proto_rawDesc = nil
	file_relay_config_v1_config_proto_goTypes = nil
	file_relay_config_v1_config_proto_depIdxs = nil
}
