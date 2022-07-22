// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: core/eventstore/v1/eventstore.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Wrapper for the Event model, this is to keep inline with our linters
type CreateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *CreateEventRequest) Reset() {
	*x = CreateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_eventstore_v1_eventstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventRequest) ProtoMessage() {}

func (x *CreateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_eventstore_v1_eventstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventRequest.ProtoReflect.Descriptor instead.
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return file_core_eventstore_v1_eventstore_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEventRequest) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type CreateEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// event id
	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IsPublished bool   `protobuf:"varint,2,opt,name=is_published,json=isPublished,proto3" json:"is_published,omitempty"`
}

func (x *CreateEventResponse) Reset() {
	*x = CreateEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_eventstore_v1_eventstore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventResponse) ProtoMessage() {}

func (x *CreateEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_eventstore_v1_eventstore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventResponse.ProtoReflect.Descriptor instead.
func (*CreateEventResponse) Descriptor() ([]byte, []int) {
	return file_core_eventstore_v1_eventstore_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEventResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateEventResponse) GetIsPublished() bool {
	if x != nil {
		return x.IsPublished
	}
	return false
}

// Event - Is the model of the event store
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// do not set this on Create() call, this is a uuid generated in the grpc handler
	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// the time the event create request was received
	ReceivedDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=received_date,json=receivedDate,proto3" json:"received_date,omitempty"`
	// the time the event was been published to it's respective topic/queue
	PublishedDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=published_date,json=publishedDate,proto3" json:"published_date,omitempty"`
	// uuid generated by the caller (e.g. commandhandler) and received by the client (e.g. react app)
	TransactionId string `protobuf:"bytes,4,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	// whether or not to publish the event to the messagebus
	Publish bool `protobuf:"varint,5,opt,name=publish,proto3" json:"publish,omitempty"`
	// the integer identifier (enum value) of the aggregate type this event belongs to
	// NOTE: this is simply an integer to keep the eventstore from depending on changing types
	AggregateType int64 `protobuf:"varint,17,opt,name=aggregate_type,json=aggregateType,proto3" json:"aggregate_type,omitempty"`
	// map of all event types
	// NOTE: this is simply an integer to keep the eventstore from depending on changing types
	EventType int64 `protobuf:"varint,18,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	// the id of the aggregate that this event belongs to
	AggregateId string `protobuf:"bytes,19,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	// data representing the state of the system that this event encapsulates
	// it is saved as a json string so that the eventstore can be completey agnostic to the
	// the data structure of the system (and thus be a static service)
	// this data MUST be able to be unmarshalled into a Proto message type
	EventData string `protobuf:"bytes,20,opt,name=event_data,json=eventData,proto3" json:"event_data,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_eventstore_v1_eventstore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_core_eventstore_v1_eventstore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_core_eventstore_v1_eventstore_proto_rawDescGZIP(), []int{2}
}

func (x *Event) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *Event) GetReceivedDate() *timestamppb.Timestamp {
	if x != nil {
		return x.ReceivedDate
	}
	return nil
}

func (x *Event) GetPublishedDate() *timestamppb.Timestamp {
	if x != nil {
		return x.PublishedDate
	}
	return nil
}

func (x *Event) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *Event) GetPublish() bool {
	if x != nil {
		return x.Publish
	}
	return false
}

func (x *Event) GetAggregateType() int64 {
	if x != nil {
		return x.AggregateType
	}
	return 0
}

func (x *Event) GetEventType() int64 {
	if x != nil {
		return x.EventType
	}
	return 0
}

func (x *Event) GetAggregateId() string {
	if x != nil {
		return x.AggregateId
	}
	return ""
}

func (x *Event) GetEventData() string {
	if x != nil {
		return x.EventData
	}
	return ""
}

var File_core_eventstore_v1_eventstore_proto protoreflect.FileDescriptor

var file_core_eventstore_v1_eventstore_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0x48, 0x0a, 0x13, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x22, 0x90, 0x03, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x26,
	0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0xb0, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x07, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0d, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76,
	0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x63, 0x65, 0x69,
	0x76, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x41, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x0e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x0c, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0x52, 0x0b, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x32, 0x69, 0x0a, 0x0a, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x5b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x26, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x69, 0x72, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x2d, 0x6f, 0x66, 0x66, 0x69,
	0x63, 0x69, 0x61, 0x6c, 0x2f, 0x67, 0x61, 0x6c, 0x61, 0x63, 0x74, 0x75, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_core_eventstore_v1_eventstore_proto_rawDescOnce sync.Once
	file_core_eventstore_v1_eventstore_proto_rawDescData = file_core_eventstore_v1_eventstore_proto_rawDesc
)

func file_core_eventstore_v1_eventstore_proto_rawDescGZIP() []byte {
	file_core_eventstore_v1_eventstore_proto_rawDescOnce.Do(func() {
		file_core_eventstore_v1_eventstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_eventstore_v1_eventstore_proto_rawDescData)
	})
	return file_core_eventstore_v1_eventstore_proto_rawDescData
}

var file_core_eventstore_v1_eventstore_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_core_eventstore_v1_eventstore_proto_goTypes = []interface{}{
	(*CreateEventRequest)(nil),    // 0: core.eventstore.v1.CreateEventRequest
	(*CreateEventResponse)(nil),   // 1: core.eventstore.v1.CreateEventResponse
	(*Event)(nil),                 // 2: core.eventstore.v1.Event
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_core_eventstore_v1_eventstore_proto_depIdxs = []int32{
	2, // 0: core.eventstore.v1.CreateEventRequest.event:type_name -> core.eventstore.v1.Event
	3, // 1: core.eventstore.v1.Event.received_date:type_name -> google.protobuf.Timestamp
	3, // 2: core.eventstore.v1.Event.published_date:type_name -> google.protobuf.Timestamp
	0, // 3: core.eventstore.v1.EventStore.Create:input_type -> core.eventstore.v1.CreateEventRequest
	1, // 4: core.eventstore.v1.EventStore.Create:output_type -> core.eventstore.v1.CreateEventResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_core_eventstore_v1_eventstore_proto_init() }
func file_core_eventstore_v1_eventstore_proto_init() {
	if File_core_eventstore_v1_eventstore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_eventstore_v1_eventstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventRequest); i {
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
		file_core_eventstore_v1_eventstore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventResponse); i {
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
		file_core_eventstore_v1_eventstore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
			RawDescriptor: file_core_eventstore_v1_eventstore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_eventstore_v1_eventstore_proto_goTypes,
		DependencyIndexes: file_core_eventstore_v1_eventstore_proto_depIdxs,
		MessageInfos:      file_core_eventstore_v1_eventstore_proto_msgTypes,
	}.Build()
	File_core_eventstore_v1_eventstore_proto = out.File
	file_core_eventstore_v1_eventstore_proto_rawDesc = nil
	file_core_eventstore_v1_eventstore_proto_goTypes = nil
	file_core_eventstore_v1_eventstore_proto_depIdxs = nil
}
