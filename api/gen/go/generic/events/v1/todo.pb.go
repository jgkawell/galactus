// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: generic/events/v1/todo.proto

package v1

import (
	v1 "github.com/jgkawell/galactus/api/gen/go/todo/aggregates/v1"
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

type TodoEventCode int32

const (
	TodoEventCode_TODO_EVENT_CODE_INVALID_UNSPECIFIED TodoEventCode = 0
	TodoEventCode_TODO_EVENT_CODE_CREATED             TodoEventCode = 1
	TodoEventCode_TODO_EVENT_CODE_DELETED             TodoEventCode = 2
)

// Enum value maps for TodoEventCode.
var (
	TodoEventCode_name = map[int32]string{
		0: "TODO_EVENT_CODE_INVALID_UNSPECIFIED",
		1: "TODO_EVENT_CODE_CREATED",
		2: "TODO_EVENT_CODE_DELETED",
	}
	TodoEventCode_value = map[string]int32{
		"TODO_EVENT_CODE_INVALID_UNSPECIFIED": 0,
		"TODO_EVENT_CODE_CREATED":             1,
		"TODO_EVENT_CODE_DELETED":             2,
	}
)

func (x TodoEventCode) Enum() *TodoEventCode {
	p := new(TodoEventCode)
	*p = x
	return p
}

func (x TodoEventCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TodoEventCode) Descriptor() protoreflect.EnumDescriptor {
	return file_generic_events_v1_todo_proto_enumTypes[0].Descriptor()
}

func (TodoEventCode) Type() protoreflect.EnumType {
	return &file_generic_events_v1_todo_proto_enumTypes[0]
}

func (x TodoEventCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TodoEventCode.Descriptor instead.
func (TodoEventCode) EnumDescriptor() ([]byte, []int) {
	return file_generic_events_v1_todo_proto_rawDescGZIP(), []int{0}
}

type TodoCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo *v1.Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
}

func (x *TodoCreated) Reset() {
	*x = TodoCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generic_events_v1_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoCreated) ProtoMessage() {}

func (x *TodoCreated) ProtoReflect() protoreflect.Message {
	mi := &file_generic_events_v1_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoCreated.ProtoReflect.Descriptor instead.
func (*TodoCreated) Descriptor() ([]byte, []int) {
	return file_generic_events_v1_todo_proto_rawDescGZIP(), []int{0}
}

func (x *TodoCreated) GetTodo() *v1.Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

type TodoCreationFailed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo  *v1.Todo `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	Error string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *TodoCreationFailed) Reset() {
	*x = TodoCreationFailed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_generic_events_v1_todo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TodoCreationFailed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TodoCreationFailed) ProtoMessage() {}

func (x *TodoCreationFailed) ProtoReflect() protoreflect.Message {
	mi := &file_generic_events_v1_todo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TodoCreationFailed.ProtoReflect.Descriptor instead.
func (*TodoCreationFailed) Descriptor() ([]byte, []int) {
	return file_generic_events_v1_todo_proto_rawDescGZIP(), []int{1}
}

func (x *TodoCreationFailed) GetTodo() *v1.Todo {
	if x != nil {
		return x.Todo
	}
	return nil
}

func (x *TodoCreationFailed) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_generic_events_v1_todo_proto protoreflect.FileDescriptor

var file_generic_events_v1_todo_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x1d, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3b, 0x0a, 0x0b, 0x54, 0x6f, 0x64, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x2c, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x22, 0x58, 0x0a,
	0x12, 0x54, 0x6f, 0x64, 0x6f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x61, 0x69,
	0x6c, 0x65, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x52, 0x04, 0x74, 0x6f, 0x64,
	0x6f, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x72, 0x0a, 0x0d, 0x54, 0x6f, 0x64, 0x6f, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x27, 0x0a, 0x23, 0x54, 0x4f, 0x44, 0x4f,
	0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x4f, 0x44, 0x4f, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x43, 0x4f, 0x44, 0x45, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1b,
	0x0a, 0x17, 0x54, 0x4f, 0x44, 0x4f, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4f, 0x44,
	0x45, 0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x42, 0x3b, 0x5a, 0x39, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x67, 0x6b, 0x61, 0x77, 0x65,
	0x6c, 0x6c, 0x2f, 0x67, 0x61, 0x6c, 0x61, 0x63, 0x74, 0x75, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_generic_events_v1_todo_proto_rawDescOnce sync.Once
	file_generic_events_v1_todo_proto_rawDescData = file_generic_events_v1_todo_proto_rawDesc
)

func file_generic_events_v1_todo_proto_rawDescGZIP() []byte {
	file_generic_events_v1_todo_proto_rawDescOnce.Do(func() {
		file_generic_events_v1_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_generic_events_v1_todo_proto_rawDescData)
	})
	return file_generic_events_v1_todo_proto_rawDescData
}

var file_generic_events_v1_todo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_generic_events_v1_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_generic_events_v1_todo_proto_goTypes = []interface{}{
	(TodoEventCode)(0),         // 0: generic.events.v1.TodoEventCode
	(*TodoCreated)(nil),        // 1: generic.events.v1.TodoCreated
	(*TodoCreationFailed)(nil), // 2: generic.events.v1.TodoCreationFailed
	(*v1.Todo)(nil),            // 3: todo.aggregates.v1.Todo
}
var file_generic_events_v1_todo_proto_depIdxs = []int32{
	3, // 0: generic.events.v1.TodoCreated.todo:type_name -> todo.aggregates.v1.Todo
	3, // 1: generic.events.v1.TodoCreationFailed.todo:type_name -> todo.aggregates.v1.Todo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_generic_events_v1_todo_proto_init() }
func file_generic_events_v1_todo_proto_init() {
	if File_generic_events_v1_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_generic_events_v1_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoCreated); i {
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
		file_generic_events_v1_todo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TodoCreationFailed); i {
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
			RawDescriptor: file_generic_events_v1_todo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_generic_events_v1_todo_proto_goTypes,
		DependencyIndexes: file_generic_events_v1_todo_proto_depIdxs,
		EnumInfos:         file_generic_events_v1_todo_proto_enumTypes,
		MessageInfos:      file_generic_events_v1_todo_proto_msgTypes,
	}.Build()
	File_generic_events_v1_todo_proto = out.File
	file_generic_events_v1_todo_proto_rawDesc = nil
	file_generic_events_v1_todo_proto_goTypes = nil
	file_generic_events_v1_todo_proto_depIdxs = nil
}
