// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: branch_office.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type BranchOffice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Country string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	State   string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Cp      uint32 `protobuf:"varint,4,opt,name=cp,proto3" json:"cp,omitempty"`
	Phone   string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"` //TODO: check if its viable to use products agenda services and employees
}

func (x *BranchOffice) Reset() {
	*x = BranchOffice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_office_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchOffice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchOffice) ProtoMessage() {}

func (x *BranchOffice) ProtoReflect() protoreflect.Message {
	mi := &file_branch_office_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchOffice.ProtoReflect.Descriptor instead.
func (*BranchOffice) Descriptor() ([]byte, []int) {
	return file_branch_office_proto_rawDescGZIP(), []int{0}
}

func (x *BranchOffice) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BranchOffice) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *BranchOffice) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *BranchOffice) GetCp() uint32 {
	if x != nil {
		return x.Cp
	}
	return 0
}

func (x *BranchOffice) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type BranchOfficeRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succes    bool                 `protobuf:"varint,1,opt,name=succes,proto3" json:"succes,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *BranchOfficeRes) Reset() {
	*x = BranchOfficeRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_branch_office_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BranchOfficeRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BranchOfficeRes) ProtoMessage() {}

func (x *BranchOfficeRes) ProtoReflect() protoreflect.Message {
	mi := &file_branch_office_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BranchOfficeRes.ProtoReflect.Descriptor instead.
func (*BranchOfficeRes) Descriptor() ([]byte, []int) {
	return file_branch_office_proto_rawDescGZIP(), []int{1}
}

func (x *BranchOfficeRes) GetSucces() bool {
	if x != nil {
		return x.Succes
	}
	return false
}

func (x *BranchOfficeRes) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_branch_office_proto protoreflect.FileDescriptor

var file_branch_office_proto_rawDesc = []byte{
	0x0a, 0x13, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x5f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x78, 0x0a, 0x0c, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x63, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x63, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x22, 0x64, 0x0a, 0x0f, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x4f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x54, 0x0a, 0x13, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x72, 0x61,
	0x6e, 0x63, 0x68, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x42,
	0x72, 0x61, 0x6e, 0x63, 0x68, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d,
	0x61, 0x75, 0x2d, 0x4d, 0x52, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_branch_office_proto_rawDescOnce sync.Once
	file_branch_office_proto_rawDescData = file_branch_office_proto_rawDesc
)

func file_branch_office_proto_rawDescGZIP() []byte {
	file_branch_office_proto_rawDescOnce.Do(func() {
		file_branch_office_proto_rawDescData = protoimpl.X.CompressGZIP(file_branch_office_proto_rawDescData)
	})
	return file_branch_office_proto_rawDescData
}

var file_branch_office_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_branch_office_proto_goTypes = []interface{}{
	(*BranchOffice)(nil),        // 0: pb.BranchOffice
	(*BranchOfficeRes)(nil),     // 1: pb.BranchOfficeRes
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_branch_office_proto_depIdxs = []int32{
	2, // 0: pb.BranchOfficeRes.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: pb.BranchOfficeService.CreateBranchOffice:input_type -> pb.BranchOffice
	1, // 2: pb.BranchOfficeService.CreateBranchOffice:output_type -> pb.BranchOfficeRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_branch_office_proto_init() }
func file_branch_office_proto_init() {
	if File_branch_office_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_branch_office_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BranchOffice); i {
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
		file_branch_office_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BranchOfficeRes); i {
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
			RawDescriptor: file_branch_office_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_branch_office_proto_goTypes,
		DependencyIndexes: file_branch_office_proto_depIdxs,
		MessageInfos:      file_branch_office_proto_msgTypes,
	}.Build()
	File_branch_office_proto = out.File
	file_branch_office_proto_rawDesc = nil
	file_branch_office_proto_goTypes = nil
	file_branch_office_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BranchOfficeServiceClient is the client API for BranchOfficeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BranchOfficeServiceClient interface {
	CreateBranchOffice(ctx context.Context, in *BranchOffice, opts ...grpc.CallOption) (*BranchOfficeRes, error)
}

type branchOfficeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBranchOfficeServiceClient(cc grpc.ClientConnInterface) BranchOfficeServiceClient {
	return &branchOfficeServiceClient{cc}
}

func (c *branchOfficeServiceClient) CreateBranchOffice(ctx context.Context, in *BranchOffice, opts ...grpc.CallOption) (*BranchOfficeRes, error) {
	out := new(BranchOfficeRes)
	err := c.cc.Invoke(ctx, "/pb.BranchOfficeService/CreateBranchOffice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BranchOfficeServiceServer is the server API for BranchOfficeService service.
type BranchOfficeServiceServer interface {
	CreateBranchOffice(context.Context, *BranchOffice) (*BranchOfficeRes, error)
}

// UnimplementedBranchOfficeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBranchOfficeServiceServer struct {
}

func (*UnimplementedBranchOfficeServiceServer) CreateBranchOffice(context.Context, *BranchOffice) (*BranchOfficeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBranchOffice not implemented")
}

func RegisterBranchOfficeServiceServer(s *grpc.Server, srv BranchOfficeServiceServer) {
	s.RegisterService(&_BranchOfficeService_serviceDesc, srv)
}

func _BranchOfficeService_CreateBranchOffice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BranchOffice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BranchOfficeServiceServer).CreateBranchOffice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.BranchOfficeService/CreateBranchOffice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BranchOfficeServiceServer).CreateBranchOffice(ctx, req.(*BranchOffice))
	}
	return interceptor(ctx, in, info, handler)
}

var _BranchOfficeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BranchOfficeService",
	HandlerType: (*BranchOfficeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBranchOffice",
			Handler:    _BranchOfficeService_CreateBranchOffice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "branch_office.proto",
}
