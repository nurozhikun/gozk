// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: http.proto

package zpbf

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

// unit message
type Point2F struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
}

func (x *Point2F) Reset() {
	*x = Point2F{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point2F) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point2F) ProtoMessage() {}

func (x *Point2F) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point2F.ProtoReflect.Descriptor instead.
func (*Point2F) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{0}
}

func (x *Point2F) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Point2F) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

type Point3F struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Point3F) Reset() {
	*x = Point3F{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Point3F) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Point3F) ProtoMessage() {}

func (x *Point3F) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Point3F.ProtoReflect.Descriptor instead.
func (*Point3F) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{1}
}

func (x *Point3F) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Point3F) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Point3F) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

// for http request and response
type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd       int64  `protobuf:"varint,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"` //ms
	Jwt       string `protobuf:"bytes,3,opt,name=jwt,proto3" json:"jwt,omitempty"`
	Code      int64  `protobuf:"varint,4,opt,name=code,proto3" json:"code,omitempty"`
	Error     string `protobuf:"bytes,5,opt,name=error,proto3" json:"error,omitempty"`
	BodyTotal int64  `protobuf:"varint,6,opt,name=body_total,json=bodyTotal,proto3" json:"body_total,omitempty"`
	BodyIdx   int64  `protobuf:"varint,7,opt,name=body_idx,json=bodyIdx,proto3" json:"body_idx,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_http_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_http_proto_rawDescGZIP(), []int{2}
}

func (x *Header) GetCmd() int64 {
	if x != nil {
		return x.Cmd
	}
	return 0
}

func (x *Header) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Header) GetJwt() string {
	if x != nil {
		return x.Jwt
	}
	return ""
}

func (x *Header) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Header) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Header) GetBodyTotal() int64 {
	if x != nil {
		return x.BodyTotal
	}
	return 0
}

func (x *Header) GetBodyIdx() int64 {
	if x != nil {
		return x.BodyIdx
	}
	return 0
}

var File_http_proto protoreflect.FileDescriptor

var file_http_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x7a, 0x70,
	0x62, 0x66, 0x22, 0x25, 0x0a, 0x07, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x32, 0x46, 0x12, 0x0c, 0x0a,
	0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x22, 0x33, 0x0a, 0x07, 0x50, 0x6f, 0x69,
	0x6e, 0x74, 0x33, 0x46, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79,
	0x12, 0x0c, 0x0a, 0x01, 0x7a, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x7a, 0x22, 0xae,
	0x01, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6d, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6a, 0x77, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6a, 0x77, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x6f, 0x64, 0x79, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x64, 0x79, 0x5f, 0x69, 0x64, 0x78,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x6f, 0x64, 0x79, 0x49, 0x64, 0x78, 0x42,
	0x1a, 0x0a, 0x09, 0x67, 0x6f, 0x7a, 0x6b, 0x2e, 0x7a, 0x70, 0x62, 0x66, 0x48, 0x03, 0x5a, 0x0b,
	0x7a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x7a, 0x70, 0x62, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_http_proto_rawDescOnce sync.Once
	file_http_proto_rawDescData = file_http_proto_rawDesc
)

func file_http_proto_rawDescGZIP() []byte {
	file_http_proto_rawDescOnce.Do(func() {
		file_http_proto_rawDescData = protoimpl.X.CompressGZIP(file_http_proto_rawDescData)
	})
	return file_http_proto_rawDescData
}

var file_http_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_http_proto_goTypes = []interface{}{
	(*Point2F)(nil), // 0: zpbf.Point2F
	(*Point3F)(nil), // 1: zpbf.Point3F
	(*Header)(nil),  // 2: zpbf.Header
}
var file_http_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_http_proto_init() }
func file_http_proto_init() {
	if File_http_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_http_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point2F); i {
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
		file_http_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Point3F); i {
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
		file_http_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
			RawDescriptor: file_http_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_http_proto_goTypes,
		DependencyIndexes: file_http_proto_depIdxs,
		MessageInfos:      file_http_proto_msgTypes,
	}.Build()
	File_http_proto = out.File
	file_http_proto_rawDesc = nil
	file_http_proto_goTypes = nil
	file_http_proto_depIdxs = nil
}
