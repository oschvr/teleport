// Copyright 2022 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        (unknown)
// source: v1/startup_service.proto

package v1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type ResolveTshdEventsServerAddressRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ResolveTshdEventsServerAddressRequest) Reset() {
	*x = ResolveTshdEventsServerAddressRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_startup_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveTshdEventsServerAddressRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveTshdEventsServerAddressRequest) ProtoMessage() {}

func (x *ResolveTshdEventsServerAddressRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_startup_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveTshdEventsServerAddressRequest.ProtoReflect.Descriptor instead.
func (*ResolveTshdEventsServerAddressRequest) Descriptor() ([]byte, []int) {
	return file_v1_startup_service_proto_rawDescGZIP(), []int{0}
}

func (x *ResolveTshdEventsServerAddressRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ResolveTshdEventsServerAddressResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResolveTshdEventsServerAddressResponse) Reset() {
	*x = ResolveTshdEventsServerAddressResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_startup_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveTshdEventsServerAddressResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveTshdEventsServerAddressResponse) ProtoMessage() {}

func (x *ResolveTshdEventsServerAddressResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_startup_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveTshdEventsServerAddressResponse.ProtoReflect.Descriptor instead.
func (*ResolveTshdEventsServerAddressResponse) Descriptor() ([]byte, []int) {
	return file_v1_startup_service_proto_rawDescGZIP(), []int{1}
}

var File_v1_startup_service_proto protoreflect.FileDescriptor

var file_v1_startup_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x31,
	0x22, 0x41, 0x0a, 0x25, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x54, 0x73, 0x68, 0x64, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x28, 0x0a, 0x26, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x54, 0x73,
	0x68, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xae, 0x01,
	0x0a, 0x0e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x9b, 0x01, 0x0a, 0x1e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x54, 0x73, 0x68, 0x64,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x3b, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x74,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x54, 0x73, 0x68, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x3c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x74, 0x65, 0x72, 0x6d,
	0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x54,
	0x73, 0x68, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33,
	0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x72, 0x6d,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_startup_service_proto_rawDescOnce sync.Once
	file_v1_startup_service_proto_rawDescData = file_v1_startup_service_proto_rawDesc
)

func file_v1_startup_service_proto_rawDescGZIP() []byte {
	file_v1_startup_service_proto_rawDescOnce.Do(func() {
		file_v1_startup_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_startup_service_proto_rawDescData)
	})
	return file_v1_startup_service_proto_rawDescData
}

var file_v1_startup_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_startup_service_proto_goTypes = []interface{}{
	(*ResolveTshdEventsServerAddressRequest)(nil),  // 0: teleport.terminal.v1.ResolveTshdEventsServerAddressRequest
	(*ResolveTshdEventsServerAddressResponse)(nil), // 1: teleport.terminal.v1.ResolveTshdEventsServerAddressResponse
}
var file_v1_startup_service_proto_depIdxs = []int32{
	0, // 0: teleport.terminal.v1.StartupService.ResolveTshdEventsServerAddress:input_type -> teleport.terminal.v1.ResolveTshdEventsServerAddressRequest
	1, // 1: teleport.terminal.v1.StartupService.ResolveTshdEventsServerAddress:output_type -> teleport.terminal.v1.ResolveTshdEventsServerAddressResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_startup_service_proto_init() }
func file_v1_startup_service_proto_init() {
	if File_v1_startup_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_startup_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveTshdEventsServerAddressRequest); i {
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
		file_v1_startup_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveTshdEventsServerAddressResponse); i {
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
			RawDescriptor: file_v1_startup_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_startup_service_proto_goTypes,
		DependencyIndexes: file_v1_startup_service_proto_depIdxs,
		MessageInfos:      file_v1_startup_service_proto_msgTypes,
	}.Build()
	File_v1_startup_service_proto = out.File
	file_v1_startup_service_proto_rawDesc = nil
	file_v1_startup_service_proto_goTypes = nil
	file_v1_startup_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StartupServiceClient is the client API for StartupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StartupServiceClient interface {
	// ResolveTshdEventsServerAddress is called by the Electron app after the tshd events server has
	// started.
	ResolveTshdEventsServerAddress(ctx context.Context, in *ResolveTshdEventsServerAddressRequest, opts ...grpc.CallOption) (*ResolveTshdEventsServerAddressResponse, error)
}

type startupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStartupServiceClient(cc grpc.ClientConnInterface) StartupServiceClient {
	return &startupServiceClient{cc}
}

func (c *startupServiceClient) ResolveTshdEventsServerAddress(ctx context.Context, in *ResolveTshdEventsServerAddressRequest, opts ...grpc.CallOption) (*ResolveTshdEventsServerAddressResponse, error) {
	out := new(ResolveTshdEventsServerAddressResponse)
	err := c.cc.Invoke(ctx, "/teleport.terminal.v1.StartupService/ResolveTshdEventsServerAddress", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StartupServiceServer is the server API for StartupService service.
type StartupServiceServer interface {
	// ResolveTshdEventsServerAddress is called by the Electron app after the tshd events server has
	// started.
	ResolveTshdEventsServerAddress(context.Context, *ResolveTshdEventsServerAddressRequest) (*ResolveTshdEventsServerAddressResponse, error)
}

// UnimplementedStartupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStartupServiceServer struct {
}

func (*UnimplementedStartupServiceServer) ResolveTshdEventsServerAddress(context.Context, *ResolveTshdEventsServerAddressRequest) (*ResolveTshdEventsServerAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResolveTshdEventsServerAddress not implemented")
}

func RegisterStartupServiceServer(s *grpc.Server, srv StartupServiceServer) {
	s.RegisterService(&_StartupService_serviceDesc, srv)
}

func _StartupService_ResolveTshdEventsServerAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveTshdEventsServerAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartupServiceServer).ResolveTshdEventsServerAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teleport.terminal.v1.StartupService/ResolveTshdEventsServerAddress",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartupServiceServer).ResolveTshdEventsServerAddress(ctx, req.(*ResolveTshdEventsServerAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _StartupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "teleport.terminal.v1.StartupService",
	HandlerType: (*StartupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ResolveTshdEventsServerAddress",
			Handler:    _StartupService_ResolveTshdEventsServerAddress_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/startup_service.proto",
}
