// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: track.proto

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

// 开始跟踪
type ReqStartTrack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X int32 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y int32 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	W int32 `protobuf:"varint,3,opt,name=w,proto3" json:"w,omitempty"`
	H int32 `protobuf:"varint,4,opt,name=h,proto3" json:"h,omitempty"`
}

func (x *ReqStartTrack) Reset() {
	*x = ReqStartTrack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_track_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqStartTrack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqStartTrack) ProtoMessage() {}

func (x *ReqStartTrack) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqStartTrack.ProtoReflect.Descriptor instead.
func (*ReqStartTrack) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{0}
}

func (x *ReqStartTrack) GetX() int32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *ReqStartTrack) GetY() int32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *ReqStartTrack) GetW() int32 {
	if x != nil {
		return x.W
	}
	return 0
}

func (x *ReqStartTrack) GetH() int32 {
	if x != nil {
		return x.H
	}
	return 0
}

// 停止跟踪
type ReqStopTrack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqStopTrack) Reset() {
	*x = ReqStopTrack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_track_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqStopTrack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqStopTrack) ProtoMessage() {}

func (x *ReqStopTrack) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqStopTrack.ProtoReflect.Descriptor instead.
func (*ReqStopTrack) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{1}
}

// 暂停跟踪
type ReqPauseTrack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqPauseTrack) Reset() {
	*x = ReqPauseTrack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_track_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqPauseTrack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqPauseTrack) ProtoMessage() {}

func (x *ReqPauseTrack) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqPauseTrack.ProtoReflect.Descriptor instead.
func (*ReqPauseTrack) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{2}
}

// 继续跟踪
type ReqContinueTrack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqContinueTrack) Reset() {
	*x = ReqContinueTrack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_track_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqContinueTrack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqContinueTrack) ProtoMessage() {}

func (x *ReqContinueTrack) ProtoReflect() protoreflect.Message {
	mi := &file_track_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqContinueTrack.ProtoReflect.Descriptor instead.
func (*ReqContinueTrack) Descriptor() ([]byte, []int) {
	return file_track_proto_rawDescGZIP(), []int{3}
}

var File_track_proto protoreflect.FileDescriptor

var file_track_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a,
	0x0d, 0x52, 0x65, 0x71, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x12, 0x0c,
	0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x77, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x77, 0x12, 0x0c, 0x0a, 0x01, 0x68, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x01, 0x68, 0x22, 0x0e, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x53, 0x74, 0x6f,
	0x70, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x50, 0x61, 0x75,
	0x73, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x43, 0x6f,
	0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x42, 0x1b, 0x5a, 0x19, 0x76,
	0x70, 0x72, 0x65, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x77, 0x61, 0x72,
	0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_track_proto_rawDescOnce sync.Once
	file_track_proto_rawDescData = file_track_proto_rawDesc
)

func file_track_proto_rawDescGZIP() []byte {
	file_track_proto_rawDescOnce.Do(func() {
		file_track_proto_rawDescData = protoimpl.X.CompressGZIP(file_track_proto_rawDescData)
	})
	return file_track_proto_rawDescData
}

var file_track_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_track_proto_goTypes = []interface{}{
	(*ReqStartTrack)(nil),    // 0: ReqStartTrack
	(*ReqStopTrack)(nil),     // 1: ReqStopTrack
	(*ReqPauseTrack)(nil),    // 2: ReqPauseTrack
	(*ReqContinueTrack)(nil), // 3: ReqContinueTrack
}
var file_track_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_track_proto_init() }
func file_track_proto_init() {
	if File_track_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_track_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqStartTrack); i {
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
		file_track_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqStopTrack); i {
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
		file_track_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqPauseTrack); i {
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
		file_track_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqContinueTrack); i {
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
			RawDescriptor: file_track_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_track_proto_goTypes,
		DependencyIndexes: file_track_proto_depIdxs,
		MessageInfos:      file_track_proto_msgTypes,
	}.Build()
	File_track_proto = out.File
	file_track_proto_rawDesc = nil
	file_track_proto_goTypes = nil
	file_track_proto_depIdxs = nil
}
