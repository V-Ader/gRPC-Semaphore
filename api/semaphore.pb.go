// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: semaphore.proto

package __

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

type IncreaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *IncreaseRequest) Reset() {
	*x = IncreaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_semaphore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncreaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncreaseRequest) ProtoMessage() {}

func (x *IncreaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_semaphore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncreaseRequest.ProtoReflect.Descriptor instead.
func (*IncreaseRequest) Descriptor() ([]byte, []int) {
	return file_semaphore_proto_rawDescGZIP(), []int{0}
}

func (x *IncreaseRequest) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type DecreaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *DecreaseRequest) Reset() {
	*x = DecreaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_semaphore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecreaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecreaseRequest) ProtoMessage() {}

func (x *DecreaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_semaphore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecreaseRequest.ProtoReflect.Descriptor instead.
func (*DecreaseRequest) Descriptor() ([]byte, []int) {
	return file_semaphore_proto_rawDescGZIP(), []int{1}
}

func (x *DecreaseRequest) GetValue() int32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type SemaphoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SemaphoreResponse) Reset() {
	*x = SemaphoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_semaphore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SemaphoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SemaphoreResponse) ProtoMessage() {}

func (x *SemaphoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_semaphore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SemaphoreResponse.ProtoReflect.Descriptor instead.
func (*SemaphoreResponse) Descriptor() ([]byte, []int) {
	return file_semaphore_proto_rawDescGZIP(), []int{2}
}

func (x *SemaphoreResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SemaphoreResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_semaphore_proto protoreflect.FileDescriptor

var file_semaphore_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x65, 0x6d, 0x61, 0x70, 0x68, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x65, 0x6d, 0x61, 0x70, 0x68, 0x6f, 0x72, 0x65, 0x22, 0x27, 0x0a, 0x0f,
	0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x27, 0x0a, 0x0f, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x47,
	0x0a, 0x11, 0x53, 0x65, 0x6d, 0x61, 0x70, 0x68, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x9e, 0x01, 0x0a, 0x10, 0x53, 0x65, 0x6d, 0x61,
	0x70, 0x68, 0x6f, 0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x08,
	0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x12, 0x1a, 0x2e, 0x73, 0x65, 0x6d, 0x61, 0x70,
	0x68, 0x6f, 0x72, 0x65, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x6d, 0x61, 0x70, 0x68, 0x6f, 0x72, 0x65,
	0x2e, 0x53, 0x65, 0x6d, 0x61, 0x70, 0x68, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x44, 0x0a, 0x08, 0x44, 0x65, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x12, 0x1a,
	0x2e, 0x73, 0x65, 0x6d, 0x61, 0x70, 0x68, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x63, 0x72, 0x65,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x6d,
	0x61, 0x70, 0x68, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x65, 0x6d, 0x61, 0x70, 0x68, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_semaphore_proto_rawDescOnce sync.Once
	file_semaphore_proto_rawDescData = file_semaphore_proto_rawDesc
)

func file_semaphore_proto_rawDescGZIP() []byte {
	file_semaphore_proto_rawDescOnce.Do(func() {
		file_semaphore_proto_rawDescData = protoimpl.X.CompressGZIP(file_semaphore_proto_rawDescData)
	})
	return file_semaphore_proto_rawDescData
}

var file_semaphore_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_semaphore_proto_goTypes = []any{
	(*IncreaseRequest)(nil),   // 0: semaphore.IncreaseRequest
	(*DecreaseRequest)(nil),   // 1: semaphore.DecreaseRequest
	(*SemaphoreResponse)(nil), // 2: semaphore.SemaphoreResponse
}
var file_semaphore_proto_depIdxs = []int32{
	0, // 0: semaphore.SemaphoreService.Increase:input_type -> semaphore.IncreaseRequest
	1, // 1: semaphore.SemaphoreService.Decrease:input_type -> semaphore.DecreaseRequest
	2, // 2: semaphore.SemaphoreService.Increase:output_type -> semaphore.SemaphoreResponse
	2, // 3: semaphore.SemaphoreService.Decrease:output_type -> semaphore.SemaphoreResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_semaphore_proto_init() }
func file_semaphore_proto_init() {
	if File_semaphore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_semaphore_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*IncreaseRequest); i {
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
		file_semaphore_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DecreaseRequest); i {
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
		file_semaphore_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SemaphoreResponse); i {
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
			RawDescriptor: file_semaphore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_semaphore_proto_goTypes,
		DependencyIndexes: file_semaphore_proto_depIdxs,
		MessageInfos:      file_semaphore_proto_msgTypes,
	}.Build()
	File_semaphore_proto = out.File
	file_semaphore_proto_rawDesc = nil
	file_semaphore_proto_goTypes = nil
	file_semaphore_proto_depIdxs = nil
}
