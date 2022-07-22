// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: core/registry/v1/registry.proto

package v1

import (
	v1 "github.com/circadence-official/galactus/api/gen/go/core/aggregates/v1"
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

type ServiceStatus int32

const (
	ServiceStatus_SERVICE_STATUS_INVALID   ServiceStatus = 0
	ServiceStatus_SERVICE_STATUS_HEALTHY   ServiceStatus = 1
	ServiceStatus_SERVICE_STATUS_UNHEALTHY ServiceStatus = 2
)

// Enum value maps for ServiceStatus.
var (
	ServiceStatus_name = map[int32]string{
		0: "SERVICE_STATUS_INVALID",
		1: "SERVICE_STATUS_HEALTHY",
		2: "SERVICE_STATUS_UNHEALTHY",
	}
	ServiceStatus_value = map[string]int32{
		"SERVICE_STATUS_INVALID":   0,
		"SERVICE_STATUS_HEALTHY":   1,
		"SERVICE_STATUS_UNHEALTHY": 2,
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
	return file_core_registry_v1_registry_proto_enumTypes[0].Descriptor()
}

func (ServiceStatus) Type() protoreflect.EnumType {
	return &file_core_registry_v1_registry_proto_enumTypes[0]
}

func (x ServiceStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServiceStatus.Descriptor instead.
func (ServiceStatus) EnumDescriptor() ([]byte, []int) {
	return file_core_registry_v1_registry_proto_rawDescGZIP(), []int{0}
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"` // this needs to match the k8s namespace for remote deployments
	// address = domain.name:port
	Version     string         `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"` // should be semver
	Description string         `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Protocols   []*Protocol    `protobuf:"bytes,5,rep,name=protocols,proto3" json:"protocols,omitempty"`
	Producers   []*v1.Producer `protobuf:"bytes,6,rep,name=producers,proto3" json:"producers,omitempty"`
	Consumers   []*v1.Consumer `protobuf:"bytes,7,rep,name=consumers,proto3" json:"consumers,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_registry_v1_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_registry_v1_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_core_registry_v1_registry_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *RegisterRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *RegisterRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RegisterRequest) GetProtocols() []*Protocol {
	if x != nil {
		return x.Protocols
	}
	return nil
}

func (x *RegisterRequest) GetProducers() []*v1.Producer {
	if x != nil {
		return x.Producers
	}
	return nil
}

func (x *RegisterRequest) GetConsumers() []*v1.Consumer {
	if x != nil {
		return x.Consumers
	}
	return nil
}

type Protocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kind v1.ProtocolKind `protobuf:"varint,1,opt,name=kind,proto3,enum=core.aggregates.v1.ProtocolKind" json:"kind,omitempty"`
}

func (x *Protocol) Reset() {
	*x = Protocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_registry_v1_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Protocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Protocol) ProtoMessage() {}

func (x *Protocol) ProtoReflect() protoreflect.Message {
	mi := &file_core_registry_v1_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Protocol.ProtoReflect.Descriptor instead.
func (*Protocol) Descriptor() ([]byte, []int) {
	return file_core_registry_v1_registry_proto_rawDescGZIP(), []int{1}
}

func (x *Protocol) GetKind() v1.ProtocolKind {
	if x != nil {
		return x.Kind
	}
	return v1.ProtocolKind(0)
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Registration *v1.Registration `protobuf:"bytes,1,opt,name=registration,proto3" json:"registration,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_registry_v1_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_registry_v1_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_core_registry_v1_registry_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterResponse) GetRegistration() *v1.Registration {
	if x != nil {
		return x.Registration
	}
	return nil
}

type ConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// this is the major version of the service to connect to (v1, v2, etc.)
	Version string          `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Type    v1.ProtocolKind `protobuf:"varint,3,opt,name=type,proto3,enum=core.aggregates.v1.ProtocolKind" json:"type,omitempty"`
}

func (x *ConnectionRequest) Reset() {
	*x = ConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_registry_v1_registry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionRequest) ProtoMessage() {}

func (x *ConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_core_registry_v1_registry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionRequest.ProtoReflect.Descriptor instead.
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return file_core_registry_v1_registry_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectionRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConnectionRequest) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ConnectionRequest) GetType() v1.ProtocolKind {
	if x != nil {
		return x.Type
	}
	return v1.ProtocolKind(0)
}

type ConnectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port    int32         `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Status  ServiceStatus `protobuf:"varint,3,opt,name=status,proto3,enum=core.registry.v1.ServiceStatus" json:"status,omitempty"`
}

func (x *ConnectionResponse) Reset() {
	*x = ConnectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_core_registry_v1_registry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionResponse) ProtoMessage() {}

func (x *ConnectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_core_registry_v1_registry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionResponse.ProtoReflect.Descriptor instead.
func (*ConnectionResponse) Descriptor() ([]byte, []int) {
	return file_core_registry_v1_registry_proto_rawDescGZIP(), []int{4}
}

func (x *ConnectionResponse) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ConnectionResponse) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *ConnectionResponse) GetStatus() ServiceStatus {
	if x != nil {
		return x.Status
	}
	return ServiceStatus_SERVICE_STATUS_INVALID
}

var File_core_registry_v1_registry_proto protoreflect.FileDescriptor

var file_core_registry_v1_registry_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f,
	0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x1a, 0x21, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67,
	0x61, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x12, 0x3a, 0x0a, 0x09,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x73, 0x12, 0x3a, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x73, 0x22, 0x40, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x12, 0x34, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4b, 0x69, 0x6e, 0x64,
	0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x58, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x77, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4b,
	0x69, 0x6e, 0x64, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x7b, 0x0a, 0x12, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x37, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x65, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x52, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x01, 0x12,
	0x1c, 0x0a, 0x18, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x59, 0x10, 0x02, 0x32, 0xba, 0x01,
	0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x53, 0x0a, 0x08, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x59, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x24, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x45, 0x5a, 0x43, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x72, 0x63, 0x61, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x2d, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x2f, 0x67, 0x61, 0x6c,
	0x61, 0x63, 0x74, 0x75, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f,
	0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_core_registry_v1_registry_proto_rawDescOnce sync.Once
	file_core_registry_v1_registry_proto_rawDescData = file_core_registry_v1_registry_proto_rawDesc
)

func file_core_registry_v1_registry_proto_rawDescGZIP() []byte {
	file_core_registry_v1_registry_proto_rawDescOnce.Do(func() {
		file_core_registry_v1_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_core_registry_v1_registry_proto_rawDescData)
	})
	return file_core_registry_v1_registry_proto_rawDescData
}

var file_core_registry_v1_registry_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_core_registry_v1_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_core_registry_v1_registry_proto_goTypes = []interface{}{
	(ServiceStatus)(0),         // 0: core.registry.v1.ServiceStatus
	(*RegisterRequest)(nil),    // 1: core.registry.v1.RegisterRequest
	(*Protocol)(nil),           // 2: core.registry.v1.Protocol
	(*RegisterResponse)(nil),   // 3: core.registry.v1.RegisterResponse
	(*ConnectionRequest)(nil),  // 4: core.registry.v1.ConnectionRequest
	(*ConnectionResponse)(nil), // 5: core.registry.v1.ConnectionResponse
	(*v1.Producer)(nil),        // 6: core.aggregates.v1.Producer
	(*v1.Consumer)(nil),        // 7: core.aggregates.v1.Consumer
	(v1.ProtocolKind)(0),       // 8: core.aggregates.v1.ProtocolKind
	(*v1.Registration)(nil),    // 9: core.aggregates.v1.Registration
}
var file_core_registry_v1_registry_proto_depIdxs = []int32{
	2, // 0: core.registry.v1.RegisterRequest.protocols:type_name -> core.registry.v1.Protocol
	6, // 1: core.registry.v1.RegisterRequest.producers:type_name -> core.aggregates.v1.Producer
	7, // 2: core.registry.v1.RegisterRequest.consumers:type_name -> core.aggregates.v1.Consumer
	8, // 3: core.registry.v1.Protocol.kind:type_name -> core.aggregates.v1.ProtocolKind
	9, // 4: core.registry.v1.RegisterResponse.registration:type_name -> core.aggregates.v1.Registration
	8, // 5: core.registry.v1.ConnectionRequest.type:type_name -> core.aggregates.v1.ProtocolKind
	0, // 6: core.registry.v1.ConnectionResponse.status:type_name -> core.registry.v1.ServiceStatus
	1, // 7: core.registry.v1.Registry.Register:input_type -> core.registry.v1.RegisterRequest
	4, // 8: core.registry.v1.Registry.Connection:input_type -> core.registry.v1.ConnectionRequest
	3, // 9: core.registry.v1.Registry.Register:output_type -> core.registry.v1.RegisterResponse
	5, // 10: core.registry.v1.Registry.Connection:output_type -> core.registry.v1.ConnectionResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_core_registry_v1_registry_proto_init() }
func file_core_registry_v1_registry_proto_init() {
	if File_core_registry_v1_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_core_registry_v1_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_core_registry_v1_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Protocol); i {
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
		file_core_registry_v1_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_core_registry_v1_registry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionRequest); i {
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
		file_core_registry_v1_registry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionResponse); i {
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
			RawDescriptor: file_core_registry_v1_registry_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_registry_v1_registry_proto_goTypes,
		DependencyIndexes: file_core_registry_v1_registry_proto_depIdxs,
		EnumInfos:         file_core_registry_v1_registry_proto_enumTypes,
		MessageInfos:      file_core_registry_v1_registry_proto_msgTypes,
	}.Build()
	File_core_registry_v1_registry_proto = out.File
	file_core_registry_v1_registry_proto_rawDesc = nil
	file_core_registry_v1_registry_proto_goTypes = nil
	file_core_registry_v1_registry_proto_depIdxs = nil
}
