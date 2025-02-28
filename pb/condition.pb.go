// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: condition.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Condition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elements  []string `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements,omitempty"`
	IsEqualTo string   `protobuf:"bytes,2,opt,name=is_equal_to,json=isEqualTo,proto3" json:"is_equal_to,omitempty"`
}

func (x *Condition) Reset() {
	*x = Condition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_condition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Condition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Condition) ProtoMessage() {}

func (x *Condition) ProtoReflect() protoreflect.Message {
	mi := &file_condition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Condition.ProtoReflect.Descriptor instead.
func (*Condition) Descriptor() ([]byte, []int) {
	return file_condition_proto_rawDescGZIP(), []int{0}
}

func (x *Condition) GetElements() []string {
	if x != nil {
		return x.Elements
	}
	return nil
}

func (x *Condition) GetIsEqualTo() string {
	if x != nil {
		return x.IsEqualTo
	}
	return ""
}

var File_condition_proto protoreflect.FileDescriptor

var file_condition_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x47, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1e,
	0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x65, 0x71, 0x75, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x73, 0x45, 0x71, 0x75, 0x61, 0x6c, 0x54, 0x6f, 0x42, 0x21,
	0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x75,
	0x2d, 0x4d, 0x52, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_condition_proto_rawDescOnce sync.Once
	file_condition_proto_rawDescData = file_condition_proto_rawDesc
)

func file_condition_proto_rawDescGZIP() []byte {
	file_condition_proto_rawDescOnce.Do(func() {
		file_condition_proto_rawDescData = protoimpl.X.CompressGZIP(file_condition_proto_rawDescData)
	})
	return file_condition_proto_rawDescData
}

var file_condition_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_condition_proto_goTypes = []interface{}{
	(*Condition)(nil), // 0: pb.Condition
}
var file_condition_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_condition_proto_init() }
func file_condition_proto_init() {
	if File_condition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_condition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Condition); i {
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
			RawDescriptor: file_condition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_condition_proto_goTypes,
		DependencyIndexes: file_condition_proto_depIdxs,
		MessageInfos:      file_condition_proto_msgTypes,
	}.Build()
	File_condition_proto = out.File
	file_condition_proto_rawDesc = nil
	file_condition_proto_goTypes = nil
	file_condition_proto_depIdxs = nil
}
