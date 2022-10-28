// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: proto/math.proto

package proto

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

// MathRequest will be converted into Go variable that we will used in Go apps.
type MathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstParam  int64 `protobuf:"varint,1,opt,name=firstParam,proto3" json:"firstParam,omitempty"`
	SecondParam int64 `protobuf:"varint,2,opt,name=secondParam,proto3" json:"secondParam,omitempty"`
}

func (x *MathRequest) Reset() {
	*x = MathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_math_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MathRequest) ProtoMessage() {}

func (x *MathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_math_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MathRequest.ProtoReflect.Descriptor instead.
func (*MathRequest) Descriptor() ([]byte, []int) {
	return file_proto_math_proto_rawDescGZIP(), []int{0}
}

func (x *MathRequest) GetFirstParam() int64 {
	if x != nil {
		return x.FirstParam
	}
	return 0
}

func (x *MathRequest) GetSecondParam() int64 {
	if x != nil {
		return x.SecondParam
	}
	return 0
}

// MathResponse will be converted into Go variable that we will used in Go apps.
type MathResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *MathResponse) Reset() {
	*x = MathResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_math_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MathResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MathResponse) ProtoMessage() {}

func (x *MathResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_math_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MathResponse.ProtoReflect.Descriptor instead.
func (*MathResponse) Descriptor() ([]byte, []int) {
	return file_proto_math_proto_rawDescGZIP(), []int{1}
}

func (x *MathResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_proto_math_proto protoreflect.FileDescriptor

var file_proto_math_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0b, 0x4d, 0x61, 0x74,
	0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x26, 0x0a, 0x0c, 0x4d, 0x61,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x32, 0x6b, 0x0a, 0x04, 0x4d, 0x61, 0x74, 0x68, 0x12, 0x2e, 0x0a, 0x03, 0x41, 0x64,
	0x64, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x61,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x08, 0x4d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4d, 0x61, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x08, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_math_proto_rawDescOnce sync.Once
	file_proto_math_proto_rawDescData = file_proto_math_proto_rawDesc
)

func file_proto_math_proto_rawDescGZIP() []byte {
	file_proto_math_proto_rawDescOnce.Do(func() {
		file_proto_math_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_math_proto_rawDescData)
	})
	return file_proto_math_proto_rawDescData
}

var file_proto_math_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_math_proto_goTypes = []interface{}{
	(*MathRequest)(nil),  // 0: proto.MathRequest
	(*MathResponse)(nil), // 1: proto.MathResponse
}
var file_proto_math_proto_depIdxs = []int32{
	0, // 0: proto.Math.Add:input_type -> proto.MathRequest
	0, // 1: proto.Math.Multiply:input_type -> proto.MathRequest
	1, // 2: proto.Math.Add:output_type -> proto.MathResponse
	1, // 3: proto.Math.Multiply:output_type -> proto.MathResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_math_proto_init() }
func file_proto_math_proto_init() {
	if File_proto_math_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_math_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MathRequest); i {
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
		file_proto_math_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MathResponse); i {
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
			RawDescriptor: file_proto_math_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_math_proto_goTypes,
		DependencyIndexes: file_proto_math_proto_depIdxs,
		MessageInfos:      file_proto_math_proto_msgTypes,
	}.Build()
	File_proto_math_proto = out.File
	file_proto_math_proto_rawDesc = nil
	file_proto_math_proto_goTypes = nil
	file_proto_math_proto_depIdxs = nil
}
