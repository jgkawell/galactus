// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: core/aggregates/v1/registry.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/jgkawell/protoc-gen-gorm/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// deregister vs unregister reference: https://grammarhow.com/unregister-vs-deregister/
type ServiceStatus int32

const (
	ServiceStatus_SERVICE_STATUS_UNSPECIFIED  ServiceStatus = 0
	ServiceStatus_SERVICE_STATUS_REGISTERED   ServiceStatus = 1
	ServiceStatus_SERVICE_STATUS_DEREGISTERED ServiceStatus = 2
	ServiceStatus_SERVICE_STATUS_HEALTHY      ServiceStatus = 3
	ServiceStatus_SERVICE_STATUS_UNHEALTHY    ServiceStatus = 4
)

// Enum value maps for ServiceStatus.
var (
	ServiceStatus_name = map[int32]string{
		0: "SERVICE_STATUS_UNSPECIFIED",
		1: "SERVICE_STATUS_REGISTERED",
		2: "SERVICE_STATUS_DEREGISTERED",
		3: "SERVICE_STATUS_HEALTHY",
		4: "SERVICE_STATUS_UNHEALTHY",
	}
	ServiceStatus_value = map[string]int32{
		"SERVICE_STATUS_UNSPECIFIED":  0,
		"SERVICE_STATUS_REGISTERED":   1,
		"SERVICE_STATUS_DEREGISTERED": 2,
		"SERVICE_STATUS_HEALTHY":      3,
		"SERVICE_STATUS_UNHEALTHY":    4,
	}
)

func (x ServiceStatus) Enum() *ServiceStatus {
	p := new(ServiceStatus)
	*p = x
	return p
}

func (x ServiceStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServiceStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_core_aggregates_v1_registry_proto_enumTypes[0].Descriptor()
}

func (ServiceStatus) Type() protoreflect.EnumType {
	return &file_core_aggregates_v1_registry_proto_enumTypes[0]
}

func (x ServiceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceStatus.Descriptor instead.
func (ServiceStatus) EnumDescriptor() ([]byte, []int) {
	return file_core_aggregates_v1_registry_proto_rawDescGZIP(), []int{0}
}

type ProtocolKind int32

const (
	ProtocolKind_PROTOCOL_KIND_UNSPECIFIED ProtocolKind = 0
	ProtocolKind_PROTOCOL_KIND_GRPC        ProtocolKind = 1
	ProtocolKind_PROTOCOL_KIND_HTTP        ProtocolKind = 2
)

// Enum value maps for ProtocolKind.
var (
	ProtocolKind_name = map[int32]string{
		0: "PROTOCOL_KIND_UNSPECIFIED",
		1: "PROTOCOL_KIND_GRPC",
		2: "PROTOCOL_KIND_HTTP",
	}
	ProtocolKind_value = map[string]int32{
		"PROTOCOL_KIND_UNSPECIFIED": 0,
		"PROTOCOL_KIND_GRPC":        1,
		"PROTOCOL_KIND_HTTP":        2,
	}
)

func (x ProtocolKind) Enum() *ProtocolKind {
	p := new(ProtocolKind)
	*p = x
	return p
}

func (x ProtocolKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProtocolKind) Descriptor() protoreflect.EnumDescriptor {
	return file_core_aggregates_v1_registry_proto_enumTypes[1].Descriptor()
}

func (ProtocolKind) Type() protoreflect.EnumType {
	return &file_core_aggregates_v1_registry_proto_enumTypes[1]
}

func (x ProtocolKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProtocolKind.Descriptor instead.
func (ProtocolKind) EnumDescriptor() ([]byte, []int) {
	return file_core_aggregates_v1_registry_proto_rawDescGZIP(), []int{1}
}

// consumer kind is whether the consumer is a queue (unicast - 1:N queue to consumer) or topic (multicast - 1:1 queue to consumer)
type ConsumerKind int32

const (
	ConsumerKind_CONSUMER_KIND_UNSPECIFIED ConsumerKind = 0
	ConsumerKind_CONSUMER_KIND_QUEUE       ConsumerKind = 1
	ConsumerKind_CONSUMER_KIND_TOPIC       ConsumerKind = 2
)

// Enum value maps for ConsumerKind.
var (
	ConsumerKind_name = map[int32]string{
		0: "CONSUMER_KIND_UNSPECIFIED",
		1: "CONSUMER_KIND_QUEUE",
		2: "CONSUMER_KIND_TOPIC",
	}
	ConsumerKind_value = map[string]int32{
		"CONSUMER_KIND_UNSPECIFIED": 0,
		"CONSUMER_KIND_QUEUE":       1,
		"CONSUMER_KIND_TOPIC":       2,
	}
)

func (x ConsumerKind) Enum() *ConsumerKind {
	p := new(ConsumerKind)
	*p = x
	return p
}

func (x ConsumerKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConsumerKind) Descriptor() protoreflect.EnumDescriptor {
	return file_core_aggregates_v1_registry_proto_enumTypes[2].Descriptor()
}

func (ConsumerKind) Type() protoreflect.EnumType {
	return &file_core_aggregates_v1_registry_proto_enumTypes[2]
}

func (x ConsumerKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConsumerKind.Descriptor instead.
func (ConsumerKind) EnumDescriptor() ([]byte, []int) {
	return file_core_aggregates_v1_registry_proto_rawDescGZIP(), []int{2}
}

type Registration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the table primary key
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// NOTE: name + version + domain must be unique
	// the service domain
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// the service name
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// the service version
	Version string `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	// plain text description of the service
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// the service' current status
	Status ServiceStatus `protobuf:"varint,6,opt,name=status,proto3,enum=core.aggregates.v1.ServiceStatus" json:"status,omitempty"`
	// the routes this service exposes
	Routes []*Route `protobuf:"bytes,16,rep,name=routes,proto3" json:"routes,omitempty"`
	// the consumer configuration of the service
	Consumers []*Consumer `protobuf:"bytes,17,rep,name=consumers,proto3" json:"consumers,omitempty"`
}

func (x *Registration) Reset() {
	*x = Registration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_aggregates_v1_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registration) ProtoMessage() {}

func (x *Registration) ProtoReflect() protoreflect.Message {
	mi := &file_core_aggregates_v1_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registration.ProtoReflect.Descriptor instead.
func (*Registration) Descriptor() ([]byte, []int) {
	return file_core_aggregates_v1_registry_proto_rawDescGZIP(), []int{0}
}

func (x *Registration) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Registration) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Registration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Registration) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Registration) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Registration) GetStatus() ServiceStatus {
	if x != nil {
		return x.Status
	}
	return ServiceStatus_SERVICE_STATUS_UNSPECIFIED
}

func (x *Registration) GetRoutes() []*Route {
	if x != nil {
		return x.Routes
	}
	return nil
}

func (x *Registration) GetConsumers() []*Consumer {
	if x != nil {
		return x.Consumers
	}
	return nil
}

type Route struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the table primary key
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// NOTE: the route path must be unique
	// the route path (e.g. '/core.commandhandler.v1.CommandHandler' (grpc) or '/core/eventstore/v1' (http))
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// the host of the route (e.g. 'localhost' or the name of the service in Istio)
	Host string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	// the route port (e.g. 8080 or 8090)
	Port int32        `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Kind ProtocolKind `protobuf:"varint,5,opt,name=kind,proto3,enum=core.aggregates.v1.ProtocolKind" json:"kind,omitempty"`
}

func (x *Route) Reset() {
	*x = Route{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_aggregates_v1_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Route) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Route) ProtoMessage() {}

func (x *Route) ProtoReflect() protoreflect.Message {
	mi := &file_core_aggregates_v1_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Route.ProtoReflect.Descriptor instead.
func (*Route) Descriptor() ([]byte, []int) {
	return file_core_aggregates_v1_registry_proto_rawDescGZIP(), []int{1}
}

func (x *Route) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Route) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Route) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Route) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Route) GetKind() ProtocolKind {
	if x != nil {
		return x.Kind
	}
	return ProtocolKind_PROTOCOL_KIND_UNSPECIFIED
}

type Consumer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the table primary key
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// NOTE: the aggregate_type + event_type + event_code must be unique
	AggregateType string       `protobuf:"bytes,2,opt,name=aggregate_type,json=aggregateType,proto3" json:"aggregate_type,omitempty"`
	EventType     string       `protobuf:"bytes,3,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	EventCode     string       `protobuf:"bytes,4,opt,name=event_code,json=eventCode,proto3" json:"event_code,omitempty"`
	Kind          ConsumerKind `protobuf:"varint,5,opt,name=kind,proto3,enum=core.aggregates.v1.ConsumerKind" json:"kind,omitempty"`
}

func (x *Consumer) Reset() {
	*x = Consumer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_aggregates_v1_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consumer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consumer) ProtoMessage() {}

func (x *Consumer) ProtoReflect() protoreflect.Message {
	mi := &file_core_aggregates_v1_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consumer.ProtoReflect.Descriptor instead.
func (*Consumer) Descriptor() ([]byte, []int) {
	return file_core_aggregates_v1_registry_proto_rawDescGZIP(), []int{2}
}

func (x *Consumer) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Consumer) GetAggregateType() string {
	if x != nil {
		return x.AggregateType
	}
	return ""
}

func (x *Consumer) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *Consumer) GetEventCode() string {
	if x != nil {
		return x.EventCode
	}
	return ""
}

func (x *Consumer) GetKind() ConsumerKind {
	if x != nil {
		return x.Kind
	}
	return ConsumerKind_CONSUMER_KIND_UNSPECIFIED
}

var File_core_aggregates_v1_registry_proto protoreflect.FileDescriptor

var file_core_aggregates_v1_registry_proto_rawDesc = []byte{
	0x0a, 0x21, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x03, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x16, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0xba, 0xb9, 0x19, 0x0a,
	0x0a, 0x08, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x28, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15,
	0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0xba, 0xb9, 0x19, 0x07, 0x0a, 0x05,
	0x5a, 0x03, 0x69, 0x64, 0x78, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x29, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xfa, 0x42, 0x07,
	0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0xba, 0xb9, 0x19, 0x07, 0x0a, 0x05, 0x5a, 0x03, 0x69,
	0x64, 0x78, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xfa, 0x42, 0x07, 0x72, 0x05,
	0x10, 0x01, 0x18, 0xff, 0x01, 0xba, 0xb9, 0x19, 0x07, 0x0a, 0x05, 0x5a, 0x03, 0x69, 0x64, 0x78,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82,
	0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x06,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x2a, 0x00, 0x52,
	0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x73, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x42, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x32, 0x00,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x73, 0x3a, 0x06, 0xba, 0xb9, 0x19,
	0x02, 0x08, 0x01, 0x22, 0xe3, 0x01, 0x0a, 0x05, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x26, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xfa, 0x42, 0x05, 0x72, 0x03,
	0xb0, 0x01, 0x01, 0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x28,
	0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x15, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0xba,
	0xb9, 0x19, 0x07, 0x0a, 0x05, 0x5a, 0x03, 0x69, 0x64, 0x78, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x1e, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a,
	0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x1f, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0b,
	0xfa, 0x42, 0x08, 0x1a, 0x06, 0x10, 0xff, 0xff, 0x03, 0x20, 0x01, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x3e, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x20, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4b, 0x69, 0x6e,
	0x64, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01, 0x22, 0xa4, 0x02, 0x0a, 0x08, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x16, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01, 0xba, 0xb9, 0x19, 0x0a,
	0x0a, 0x08, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x28, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3c,
	0x0a, 0x0e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18,
	0xff, 0x01, 0xba, 0xb9, 0x19, 0x07, 0x0a, 0x05, 0x5a, 0x03, 0x69, 0x64, 0x78, 0x52, 0x0d, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x0a,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x15, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0xba, 0xb9, 0x19, 0x07,
	0x0a, 0x05, 0x5a, 0x03, 0x69, 0x64, 0x78, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x34, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18,
	0xff, 0x01, 0xba, 0xb9, 0x19, 0x07, 0x0a, 0x05, 0x5a, 0x03, 0x69, 0x64, 0x78, 0x52, 0x09, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x3e, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x67,
	0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x3a, 0x06, 0xba, 0xb9, 0x19, 0x02, 0x08, 0x01,
	0x2a, 0xa9, 0x01, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x45, 0x44,
	0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x03, 0x12, 0x1c,
	0x0a, 0x18, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x04, 0x2a, 0x5d, 0x0a, 0x0c,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x19,
	0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x50,
	0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x47, 0x52, 0x50,
	0x43, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f,
	0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x48, 0x54, 0x54, 0x50, 0x10, 0x02, 0x2a, 0x5f, 0x0a, 0x0c, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x19, 0x43,
	0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x52, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f,
	0x4e, 0x53, 0x55, 0x4d, 0x45, 0x52, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x51, 0x55, 0x45, 0x55,
	0x45, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x52, 0x5f,
	0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x54, 0x4f, 0x50, 0x49, 0x43, 0x10, 0x02, 0x42, 0x3c, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x67, 0x6b, 0x61, 0x77,
	0x65, 0x6c, 0x6c, 0x2f, 0x67, 0x61, 0x6c, 0x61, 0x63, 0x74, 0x75, 0x73, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_core_aggregates_v1_registry_proto_rawDescOnce sync.Once
	file_core_aggregates_v1_registry_proto_rawDescData = file_core_aggregates_v1_registry_proto_rawDesc
)

func file_core_aggregates_v1_registry_proto_rawDescGZIP() []byte {
	file_core_aggregates_v1_registry_proto_rawDescOnce.Do(func() {
		file_core_aggregates_v1_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_aggregates_v1_registry_proto_rawDescData)
	})
	return file_core_aggregates_v1_registry_proto_rawDescData
}

var file_core_aggregates_v1_registry_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_core_aggregates_v1_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_core_aggregates_v1_registry_proto_goTypes = []interface{}{
	(ServiceStatus)(0),   // 0: core.aggregates.v1.ServiceStatus
	(ProtocolKind)(0),    // 1: core.aggregates.v1.ProtocolKind
	(ConsumerKind)(0),    // 2: core.aggregates.v1.ConsumerKind
	(*Registration)(nil), // 3: core.aggregates.v1.Registration
	(*Route)(nil),        // 4: core.aggregates.v1.Route
	(*Consumer)(nil),     // 5: core.aggregates.v1.Consumer
}
var file_core_aggregates_v1_registry_proto_depIdxs = []int32{
	0, // 0: core.aggregates.v1.Registration.status:type_name -> core.aggregates.v1.ServiceStatus
	4, // 1: core.aggregates.v1.Registration.routes:type_name -> core.aggregates.v1.Route
	5, // 2: core.aggregates.v1.Registration.consumers:type_name -> core.aggregates.v1.Consumer
	1, // 3: core.aggregates.v1.Route.kind:type_name -> core.aggregates.v1.ProtocolKind
	2, // 4: core.aggregates.v1.Consumer.kind:type_name -> core.aggregates.v1.ConsumerKind
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_core_aggregates_v1_registry_proto_init() }
func file_core_aggregates_v1_registry_proto_init() {
	if File_core_aggregates_v1_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_aggregates_v1_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registration); i {
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
		file_core_aggregates_v1_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Route); i {
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
		file_core_aggregates_v1_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Consumer); i {
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
			RawDescriptor: file_core_aggregates_v1_registry_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_core_aggregates_v1_registry_proto_goTypes,
		DependencyIndexes: file_core_aggregates_v1_registry_proto_depIdxs,
		EnumInfos:         file_core_aggregates_v1_registry_proto_enumTypes,
		MessageInfos:      file_core_aggregates_v1_registry_proto_msgTypes,
	}.Build()
	File_core_aggregates_v1_registry_proto = out.File
	file_core_aggregates_v1_registry_proto_rawDesc = nil
	file_core_aggregates_v1_registry_proto_goTypes = nil
	file_core_aggregates_v1_registry_proto_depIdxs = nil
}
