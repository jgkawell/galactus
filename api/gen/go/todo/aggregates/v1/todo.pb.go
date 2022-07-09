// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: todo/aggregates/v1/todo.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/infobloxopen/protoc-gen-gorm/options"
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

type TodoStatus int32

const (
	TodoStatus_TODO_STATUS_INVALID TodoStatus = 0
	TodoStatus_COMPLETE            TodoStatus = 1
	TodoStatus_INCOMPLETE          TodoStatus = 2
)

// Enum value maps for TodoStatus.
var (
	TodoStatus_name = map[int32]string{
		0: "TODO_STATUS_INVALID",
		1: "COMPLETE",
		2: "INCOMPLETE",
	}
	TodoStatus_value = map[string]int32{
		"TODO_STATUS_INVALID": 0,
		"COMPLETE":            1,
		"INCOMPLETE":          2,
	}
)

func (x TodoStatus) Enum() *TodoStatus {
	p := new(TodoStatus)
	*p = x
	return p
}

func (x TodoStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TodoStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_todo_aggregates_v1_todo_proto_enumTypes[0].Descriptor()
}

func (TodoStatus) Type() protoreflect.EnumType {
	return &file_todo_aggregates_v1_todo_proto_enumTypes[0]
}

func (x TodoStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TodoStatus.Descriptor instead.
func (TodoStatus) EnumDescriptor() ([]byte, []int) {
	return file_todo_aggregates_v1_todo_proto_rawDescGZIP(), []int{0}
}

type Todo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//values
	Id          string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title       string     `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string     `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status      TodoStatus `protobuf:"varint,4,opt,name=status,proto3,enum=todo.aggregates.v1.TodoStatus" json:"status,omitempty"`
	// TODO: the proto will throw an error if you don't have a google.protobuf.Timestamp somewhere in the model. Need to figure out a way around this.
	ScheduledTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=scheduled_time,json=scheduledTime,proto3" json:"scheduled_time,omitempty"`
}

func (x *Todo) Reset() {
	*x = Todo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_todo_aggregates_v1_todo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Todo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Todo) ProtoMessage() {}

func (x *Todo) ProtoReflect() protoreflect.Message {
	mi := &file_todo_aggregates_v1_todo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Todo.ProtoReflect.Descriptor instead.
func (*Todo) Descriptor() ([]byte, []int) {
	return file_todo_aggregates_v1_todo_proto_rawDescGZIP(), []int{0}
}

func (x *Todo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Todo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Todo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Todo) GetStatus() TodoStatus {
	if x != nil {
		return x.Status
	}
	return TodoStatus_TODO_STATUS_INVALID
}

func (x *Todo) GetScheduledTime() *timestamppb.Timestamp {
	if x != nil {
		return x.ScheduledTime
	}
	return nil
}

var File_todo_aggregates_v1_todo_proto protoreflect.FileDescriptor

var file_todo_aggregates_v1_todo_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x74, 0x6f, 0x64, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x74, 0x6f, 0x64, 0x6f, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x67, 0x6f, 0x72, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xbf, 0x02, 0x0a, 0x04, 0x54, 0x6f, 0x64, 0x6f, 0x12, 0x26, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xfa, 0x42, 0x05, 0x72, 0x03, 0xb0, 0x01, 0x01,
	0xba, 0xb9, 0x19, 0x0a, 0x0a, 0x08, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x28, 0x01, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x2e, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x18, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xff, 0x01, 0xba, 0xb9, 0x19,
	0x0a, 0x0a, 0x08, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x40, 0x01, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x38, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01,
	0x18, 0xff, 0x01, 0xba, 0xb9, 0x19, 0x08, 0x0a, 0x06, 0x12, 0x04, 0x75, 0x75, 0x69, 0x64, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x74,
	0x6f, 0x64, 0x6f, 0x2e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0f, 0xba, 0xb9,
	0x19, 0x0b, 0x0a, 0x09, 0x12, 0x05, 0x65, 0x6e, 0x75, 0x6d, 0x3f, 0x40, 0x01, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x54, 0x0a, 0x0e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x11, 0xba, 0xb9, 0x19, 0x0d, 0x0a,
	0x0b, 0x12, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x3a, 0x06, 0xba, 0xb9, 0x19,
	0x02, 0x08, 0x01, 0x2a, 0x43, 0x0a, 0x0a, 0x54, 0x6f, 0x64, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x17, 0x0a, 0x13, 0x54, 0x4f, 0x44, 0x4f, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x42, 0x47, 0x5a, 0x45, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x72, 0x63, 0x61, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x2d, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x69, 0x61, 0x6c, 0x2f, 0x67, 0x61, 0x6c, 0x61, 0x63,
	0x74, 0x75, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74,
	0x6f, 0x64, 0x6f, 0x2f, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x73, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_todo_aggregates_v1_todo_proto_rawDescOnce sync.Once
	file_todo_aggregates_v1_todo_proto_rawDescData = file_todo_aggregates_v1_todo_proto_rawDesc
)

func file_todo_aggregates_v1_todo_proto_rawDescGZIP() []byte {
	file_todo_aggregates_v1_todo_proto_rawDescOnce.Do(func() {
		file_todo_aggregates_v1_todo_proto_rawDescData = protoimpl.X.CompressGZIP(file_todo_aggregates_v1_todo_proto_rawDescData)
	})
	return file_todo_aggregates_v1_todo_proto_rawDescData
}

var file_todo_aggregates_v1_todo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_todo_aggregates_v1_todo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_todo_aggregates_v1_todo_proto_goTypes = []interface{}{
	(TodoStatus)(0),               // 0: todo.aggregates.v1.TodoStatus
	(*Todo)(nil),                  // 1: todo.aggregates.v1.Todo
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_todo_aggregates_v1_todo_proto_depIdxs = []int32{
	0, // 0: todo.aggregates.v1.Todo.status:type_name -> todo.aggregates.v1.TodoStatus
	2, // 1: todo.aggregates.v1.Todo.scheduled_time:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_todo_aggregates_v1_todo_proto_init() }
func file_todo_aggregates_v1_todo_proto_init() {
	if File_todo_aggregates_v1_todo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_todo_aggregates_v1_todo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Todo); i {
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
			RawDescriptor: file_todo_aggregates_v1_todo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_todo_aggregates_v1_todo_proto_goTypes,
		DependencyIndexes: file_todo_aggregates_v1_todo_proto_depIdxs,
		EnumInfos:         file_todo_aggregates_v1_todo_proto_enumTypes,
		MessageInfos:      file_todo_aggregates_v1_todo_proto_msgTypes,
	}.Build()
	File_todo_aggregates_v1_todo_proto = out.File
	file_todo_aggregates_v1_todo_proto_rawDesc = nil
	file_todo_aggregates_v1_todo_proto_goTypes = nil
	file_todo_aggregates_v1_todo_proto_depIdxs = nil
}
