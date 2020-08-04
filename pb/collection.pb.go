// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: collection.proto

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

type Collection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	//ralated to the qty of different products
	Qty int32 `protobuf:"varint,3,opt,name=qty,proto3" json:"qty,omitempty"`
}

func (x *Collection) Reset() {
	*x = Collection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Collection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collection) ProtoMessage() {}

func (x *Collection) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collection.ProtoReflect.Descriptor instead.
func (*Collection) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{0}
}

func (x *Collection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Collection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Collection) GetQty() int32 {
	if x != nil {
		return x.Qty
	}
	return 0
}

type CreateCollectionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to Type:
	//	*CreateCollectionReq_Product
	//	*CreateCollectionReq_Services
	Type isCreateCollectionReq_Type `protobuf_oneof:"type"`
	// Types that are assignable to Hierarchy:
	//	*CreateCollectionReq_Master
	//	*CreateCollectionReq_ParentId
	Hierarchy isCreateCollectionReq_Hierarchy `protobuf_oneof:"hierarchy"`
}

func (x *CreateCollectionReq) Reset() {
	*x = CreateCollectionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCollectionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCollectionReq) ProtoMessage() {}

func (x *CreateCollectionReq) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCollectionReq.ProtoReflect.Descriptor instead.
func (*CreateCollectionReq) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCollectionReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *CreateCollectionReq) GetType() isCreateCollectionReq_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *CreateCollectionReq) GetProduct() bool {
	if x, ok := x.GetType().(*CreateCollectionReq_Product); ok {
		return x.Product
	}
	return false
}

func (x *CreateCollectionReq) GetServices() bool {
	if x, ok := x.GetType().(*CreateCollectionReq_Services); ok {
		return x.Services
	}
	return false
}

func (m *CreateCollectionReq) GetHierarchy() isCreateCollectionReq_Hierarchy {
	if m != nil {
		return m.Hierarchy
	}
	return nil
}

func (x *CreateCollectionReq) GetMaster() bool {
	if x, ok := x.GetHierarchy().(*CreateCollectionReq_Master); ok {
		return x.Master
	}
	return false
}

func (x *CreateCollectionReq) GetParentId() string {
	if x, ok := x.GetHierarchy().(*CreateCollectionReq_ParentId); ok {
		return x.ParentId
	}
	return ""
}

type isCreateCollectionReq_Type interface {
	isCreateCollectionReq_Type()
}

type CreateCollectionReq_Product struct {
	Product bool `protobuf:"varint,2,opt,name=product,proto3,oneof"`
}

type CreateCollectionReq_Services struct {
	Services bool `protobuf:"varint,3,opt,name=services,proto3,oneof"`
}

func (*CreateCollectionReq_Product) isCreateCollectionReq_Type() {}

func (*CreateCollectionReq_Services) isCreateCollectionReq_Type() {}

type isCreateCollectionReq_Hierarchy interface {
	isCreateCollectionReq_Hierarchy()
}

type CreateCollectionReq_Master struct {
	Master bool `protobuf:"varint,4,opt,name=master,proto3,oneof"`
}

type CreateCollectionReq_ParentId struct {
	ParentId string `protobuf:"bytes,5,opt,name=parentId,proto3,oneof"`
}

func (*CreateCollectionReq_Master) isCreateCollectionReq_Hierarchy() {}

func (*CreateCollectionReq_ParentId) isCreateCollectionReq_Hierarchy() {}

type CollectionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Status:
	//	*CollectionRes_CreatedAt
	Status isCollectionRes_Status `protobuf_oneof:"status"`
}

func (x *CollectionRes) Reset() {
	*x = CollectionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionRes) ProtoMessage() {}

func (x *CollectionRes) ProtoReflect() protoreflect.Message {
	mi := &file_collection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionRes.ProtoReflect.Descriptor instead.
func (*CollectionRes) Descriptor() ([]byte, []int) {
	return file_collection_proto_rawDescGZIP(), []int{2}
}

func (x *CollectionRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CollectionRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *CollectionRes) GetStatus() isCollectionRes_Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (x *CollectionRes) GetCreatedAt() *timestamp.Timestamp {
	if x, ok := x.GetStatus().(*CollectionRes_CreatedAt); ok {
		return x.CreatedAt
	}
	return nil
}

type isCollectionRes_Status interface {
	isCollectionRes_Status()
}

type CollectionRes_CreatedAt struct {
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3,oneof"`
}

func (*CollectionRes_CreatedAt) isCollectionRes_Status() {}

var File_collection_proto protoreflect.FileDescriptor

var file_collection_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x71, 0x74, 0x79, 0x22, 0xb0, 0x01, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x01, 0x52, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x08, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x42, 0x0b, 0x0a, 0x09, 0x68, 0x69, 0x65, 0x72, 0x61, 0x72, 0x63, 0x68, 0x79, 0x22, 0x80,
	0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0x53, 0x0a, 0x11, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4d, 0x61, 0x75, 0x2d, 0x4d, 0x52, 0x2f, 0x72, 0x70, 0x63, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_collection_proto_rawDescOnce sync.Once
	file_collection_proto_rawDescData = file_collection_proto_rawDesc
)

func file_collection_proto_rawDescGZIP() []byte {
	file_collection_proto_rawDescOnce.Do(func() {
		file_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_collection_proto_rawDescData)
	})
	return file_collection_proto_rawDescData
}

var file_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_collection_proto_goTypes = []interface{}{
	(*Collection)(nil),          // 0: pb.Collection
	(*CreateCollectionReq)(nil), // 1: pb.CreateCollectionReq
	(*CollectionRes)(nil),       // 2: pb.CollectionRes
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_collection_proto_depIdxs = []int32{
	3, // 0: pb.CollectionRes.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: pb.CollectionService.CreateCollection:input_type -> pb.CreateCollectionReq
	2, // 2: pb.CollectionService.CreateCollection:output_type -> pb.CollectionRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_collection_proto_init() }
func file_collection_proto_init() {
	if File_collection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Collection); i {
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
		file_collection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCollectionReq); i {
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
		file_collection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionRes); i {
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
	file_collection_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CreateCollectionReq_Product)(nil),
		(*CreateCollectionReq_Services)(nil),
		(*CreateCollectionReq_Master)(nil),
		(*CreateCollectionReq_ParentId)(nil),
	}
	file_collection_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CollectionRes_CreatedAt)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_collection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_collection_proto_goTypes,
		DependencyIndexes: file_collection_proto_depIdxs,
		MessageInfos:      file_collection_proto_msgTypes,
	}.Build()
	File_collection_proto = out.File
	file_collection_proto_rawDesc = nil
	file_collection_proto_goTypes = nil
	file_collection_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CollectionServiceClient is the client API for CollectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CollectionServiceClient interface {
	CreateCollection(ctx context.Context, in *CreateCollectionReq, opts ...grpc.CallOption) (*CollectionRes, error)
}

type collectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectionServiceClient(cc grpc.ClientConnInterface) CollectionServiceClient {
	return &collectionServiceClient{cc}
}

func (c *collectionServiceClient) CreateCollection(ctx context.Context, in *CreateCollectionReq, opts ...grpc.CallOption) (*CollectionRes, error) {
	out := new(CollectionRes)
	err := c.cc.Invoke(ctx, "/pb.CollectionService/CreateCollection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectionServiceServer is the server API for CollectionService service.
type CollectionServiceServer interface {
	CreateCollection(context.Context, *CreateCollectionReq) (*CollectionRes, error)
}

// UnimplementedCollectionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCollectionServiceServer struct {
}

func (*UnimplementedCollectionServiceServer) CreateCollection(context.Context, *CreateCollectionReq) (*CollectionRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollection not implemented")
}

func RegisterCollectionServiceServer(s *grpc.Server, srv CollectionServiceServer) {
	s.RegisterService(&_CollectionService_serviceDesc, srv)
}

func _CollectionService_CreateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCollectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).CreateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CollectionService/CreateCollection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).CreateCollection(ctx, req.(*CreateCollectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _CollectionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CollectionService",
	HandlerType: (*CollectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCollection",
			Handler:    _CollectionService_CreateCollection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "collection.proto",
}
