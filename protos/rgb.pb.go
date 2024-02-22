// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: rgb.proto

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

// 打开RGB
type ReqOpenRgb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqOpenRgb) Reset() {
	*x = ReqOpenRgb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqOpenRgb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqOpenRgb) ProtoMessage() {}

func (x *ReqOpenRgb) ProtoReflect() protoreflect.Message {
	mi := &file_rgb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqOpenRgb.ProtoReflect.Descriptor instead.
func (*ReqOpenRgb) Descriptor() ([]byte, []int) {
	return file_rgb_proto_rawDescGZIP(), []int{0}
}

// 关闭RGB
type ReqCloseRgb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqCloseRgb) Reset() {
	*x = ReqCloseRgb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqCloseRgb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqCloseRgb) ProtoMessage() {}

func (x *ReqCloseRgb) ProtoReflect() protoreflect.Message {
	mi := &file_rgb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqCloseRgb.ProtoReflect.Descriptor instead.
func (*ReqCloseRgb) Descriptor() ([]byte, []int) {
	return file_rgb_proto_rawDescGZIP(), []int{1}
}

// 关机
type ReqPowerDown struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqPowerDown) Reset() {
	*x = ReqPowerDown{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqPowerDown) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqPowerDown) ProtoMessage() {}

func (x *ReqPowerDown) ProtoReflect() protoreflect.Message {
	mi := &file_rgb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqPowerDown.ProtoReflect.Descriptor instead.
func (*ReqPowerDown) Descriptor() ([]byte, []int) {
	return file_rgb_proto_rawDescGZIP(), []int{2}
}

// 打开电量指示灯
type ReqOpenPowerInd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqOpenPowerInd) Reset() {
	*x = ReqOpenPowerInd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqOpenPowerInd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqOpenPowerInd) ProtoMessage() {}

func (x *ReqOpenPowerInd) ProtoReflect() protoreflect.Message {
	mi := &file_rgb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqOpenPowerInd.ProtoReflect.Descriptor instead.
func (*ReqOpenPowerInd) Descriptor() ([]byte, []int) {
	return file_rgb_proto_rawDescGZIP(), []int{3}
}

// 关闭电量指示灯
type ReqClosePowerInd struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqClosePowerInd) Reset() {
	*x = ReqClosePowerInd{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqClosePowerInd) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqClosePowerInd) ProtoMessage() {}

func (x *ReqClosePowerInd) ProtoReflect() protoreflect.Message {
	mi := &file_rgb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqClosePowerInd.ProtoReflect.Descriptor instead.
func (*ReqClosePowerInd) Descriptor() ([]byte, []int) {
	return file_rgb_proto_rawDescGZIP(), []int{4}
}

// 重启
type ReqReboot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqReboot) Reset() {
	*x = ReqReboot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rgb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqReboot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqReboot) ProtoMessage() {}

func (x *ReqReboot) ProtoReflect() protoreflect.Message {
	mi := &file_rgb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqReboot.ProtoReflect.Descriptor instead.
func (*ReqReboot) Descriptor() ([]byte, []int) {
	return file_rgb_proto_rawDescGZIP(), []int{5}
}

var File_rgb_proto protoreflect.FileDescriptor

var file_rgb_proto_rawDesc = []byte{
	0x0a, 0x09, 0x72, 0x67, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0c, 0x0a, 0x0a, 0x52,
	0x65, 0x71, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x67, 0x62, 0x22, 0x0d, 0x0a, 0x0b, 0x52, 0x65, 0x71,
	0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x67, 0x62, 0x22, 0x0e, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x50,
	0x6f, 0x77, 0x65, 0x72, 0x44, 0x6f, 0x77, 0x6e, 0x22, 0x11, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x4f,
	0x70, 0x65, 0x6e, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x52,
	0x65, 0x71, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x49, 0x6e, 0x64, 0x22,
	0x0b, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x52, 0x65, 0x62, 0x6f, 0x6f, 0x74, 0x42, 0x1b, 0x5a, 0x19,
	0x76, 0x70, 0x72, 0x65, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x77, 0x61,
	0x72, 0x66, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rgb_proto_rawDescOnce sync.Once
	file_rgb_proto_rawDescData = file_rgb_proto_rawDesc
)

func file_rgb_proto_rawDescGZIP() []byte {
	file_rgb_proto_rawDescOnce.Do(func() {
		file_rgb_proto_rawDescData = protoimpl.X.CompressGZIP(file_rgb_proto_rawDescData)
	})
	return file_rgb_proto_rawDescData
}

var file_rgb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rgb_proto_goTypes = []interface{}{
	(*ReqOpenRgb)(nil),       // 0: ReqOpenRgb
	(*ReqCloseRgb)(nil),      // 1: ReqCloseRgb
	(*ReqPowerDown)(nil),     // 2: ReqPowerDown
	(*ReqOpenPowerInd)(nil),  // 3: ReqOpenPowerInd
	(*ReqClosePowerInd)(nil), // 4: ReqClosePowerInd
	(*ReqReboot)(nil),        // 5: ReqReboot
}
var file_rgb_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rgb_proto_init() }
func file_rgb_proto_init() {
	if File_rgb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rgb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqOpenRgb); i {
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
		file_rgb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqCloseRgb); i {
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
		file_rgb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqPowerDown); i {
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
		file_rgb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqOpenPowerInd); i {
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
		file_rgb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqClosePowerInd); i {
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
		file_rgb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqReboot); i {
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
			RawDescriptor: file_rgb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rgb_proto_goTypes,
		DependencyIndexes: file_rgb_proto_depIdxs,
		MessageInfos:      file_rgb_proto_msgTypes,
	}.Build()
	File_rgb_proto = out.File
	file_rgb_proto_rawDesc = nil
	file_rgb_proto_goTypes = nil
	file_rgb_proto_depIdxs = nil
}
