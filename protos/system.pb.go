// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: system.proto

package protos

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

// 设置系统时间
type ReqSetTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp uint64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ReqSetTime) Reset() {
	*x = ReqSetTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqSetTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqSetTime) ProtoMessage() {}

func (x *ReqSetTime) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqSetTime.ProtoReflect.Descriptor instead.
func (*ReqSetTime) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{0}
}

func (x *ReqSetTime) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// 设置系统时区
type ReqSetTimezone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timezone string `protobuf:"bytes,1,opt,name=timezone,proto3" json:"timezone,omitempty"`
}

func (x *ReqSetTimezone) Reset() {
	*x = ReqSetTimezone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqSetTimezone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqSetTimezone) ProtoMessage() {}

func (x *ReqSetTimezone) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqSetTimezone.ProtoReflect.Descriptor instead.
func (*ReqSetTimezone) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{1}
}

func (x *ReqSetTimezone) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

// 设置MTP模式
type ReqSetMtpMode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode int32 `protobuf:"varint,1,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *ReqSetMtpMode) Reset() {
	*x = ReqSetMtpMode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqSetMtpMode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqSetMtpMode) ProtoMessage() {}

func (x *ReqSetMtpMode) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqSetMtpMode.ProtoReflect.Descriptor instead.
func (*ReqSetMtpMode) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{2}
}

func (x *ReqSetMtpMode) GetMode() int32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

// 设置CPU模式
type ReqSetCpuMode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode int32 `protobuf:"varint,1,opt,name=mode,proto3" json:"mode,omitempty"`
}

func (x *ReqSetCpuMode) Reset() {
	*x = ReqSetCpuMode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqSetCpuMode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqSetCpuMode) ProtoMessage() {}

func (x *ReqSetCpuMode) ProtoReflect() protoreflect.Message {
	mi := &file_system_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqSetCpuMode.ProtoReflect.Descriptor instead.
func (*ReqSetCpuMode) Descriptor() ([]byte, []int) {
	return file_system_proto_rawDescGZIP(), []int{3}
}

func (x *ReqSetCpuMode) GetMode() int32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

var File_system_proto protoreflect.FileDescriptor

var file_system_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a,
	0x0a, 0x0a, 0x52, 0x65, 0x71, 0x53, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x2c, 0x0a, 0x0e, 0x52, 0x65,
	0x71, 0x53, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x53,
	0x65, 0x74, 0x4d, 0x74, 0x70, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x23, 0x0a,
	0x0d, 0x52, 0x65, 0x71, 0x53, 0x65, 0x74, 0x43, 0x70, 0x75, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6d, 0x6f,
	0x64, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x76, 0x70, 0x72, 0x65, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x77, 0x61, 0x72, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_proto_rawDescOnce sync.Once
	file_system_proto_rawDescData = file_system_proto_rawDesc
)

func file_system_proto_rawDescGZIP() []byte {
	file_system_proto_rawDescOnce.Do(func() {
		file_system_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_proto_rawDescData)
	})
	return file_system_proto_rawDescData
}

var file_system_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_system_proto_goTypes = []interface{}{
	(*ReqSetTime)(nil),     // 0: ReqSetTime
	(*ReqSetTimezone)(nil), // 1: ReqSetTimezone
	(*ReqSetMtpMode)(nil),  // 2: ReqSetMtpMode
	(*ReqSetCpuMode)(nil),  // 3: ReqSetCpuMode
}
var file_system_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_system_proto_init() }
func file_system_proto_init() {
	if File_system_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_system_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqSetTime); i {
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
		file_system_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqSetTimezone); i {
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
		file_system_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqSetMtpMode); i {
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
		file_system_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqSetCpuMode); i {
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
			RawDescriptor: file_system_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_system_proto_goTypes,
		DependencyIndexes: file_system_proto_depIdxs,
		MessageInfos:      file_system_proto_msgTypes,
	}.Build()
	File_system_proto = out.File
	file_system_proto_rawDesc = nil
	file_system_proto_goTypes = nil
	file_system_proto_depIdxs = nil
}