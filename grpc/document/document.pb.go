// Copyright 2015 gRPC authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: document.proto

package document

import (
	context "context"
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

// The request message containings the user's name.
type DocumentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DocumentRequest) Reset() {
	*x = DocumentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_document_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentRequest) ProtoMessage() {}

func (x *DocumentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_document_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentRequest.ProtoReflect.Descriptor instead.
func (*DocumentRequest) Descriptor() ([]byte, []int) {
	return file_document_proto_rawDescGZIP(), []int{0}
}

func (x *DocumentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The response message containing the greetings
type DocumentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DocumentReply) Reset() {
	*x = DocumentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_document_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DocumentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DocumentReply) ProtoMessage() {}

func (x *DocumentReply) ProtoReflect() protoreflect.Message {
	mi := &file_document_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DocumentReply.ProtoReflect.Descriptor instead.
func (*DocumentReply) Descriptor() ([]byte, []int) {
	return file_document_proto_rawDescGZIP(), []int{1}
}

func (x *DocumentReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_document_proto protoreflect.FileDescriptor

var file_document_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x25, 0x0a, 0x0f, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x29, 0x0a, 0x0d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xcb, 0x01, 0x0a,
	0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x40, 0x0a, 0x08, 0x53, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x19, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x04, 0x53,
	0x65, 0x6e, 0x64, 0x12, 0x19, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x46, 0x0a, 0x1b, 0x69, 0x6f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x42, 0x0d, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x16, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_document_proto_rawDescOnce sync.Once
	file_document_proto_rawDescData = file_document_proto_rawDesc
)

func file_document_proto_rawDescGZIP() []byte {
	file_document_proto_rawDescOnce.Do(func() {
		file_document_proto_rawDescData = protoimpl.X.CompressGZIP(file_document_proto_rawDescData)
	})
	return file_document_proto_rawDescData
}

var file_document_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_document_proto_goTypes = []interface{}{
	(*DocumentRequest)(nil), // 0: document.DocumentRequest
	(*DocumentReply)(nil),   // 1: document.DocumentReply
}
var file_document_proto_depIdxs = []int32{
	0, // 0: document.Document.SayHello:input_type -> document.DocumentRequest
	0, // 1: document.Document.Send:input_type -> document.DocumentRequest
	0, // 2: document.Document.GetData:input_type -> document.DocumentRequest
	1, // 3: document.Document.SayHello:output_type -> document.DocumentReply
	1, // 4: document.Document.Send:output_type -> document.DocumentReply
	1, // 5: document.Document.GetData:output_type -> document.DocumentReply
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_document_proto_init() }
func file_document_proto_init() {
	if File_document_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_document_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentRequest); i {
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
		file_document_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DocumentReply); i {
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
			RawDescriptor: file_document_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_document_proto_goTypes,
		DependencyIndexes: file_document_proto_depIdxs,
		MessageInfos:      file_document_proto_msgTypes,
	}.Build()
	File_document_proto = out.File
	file_document_proto_rawDesc = nil
	file_document_proto_goTypes = nil
	file_document_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DocumentClient is the client API for Document service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DocumentClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *DocumentRequest, opts ...grpc.CallOption) (*DocumentReply, error)
	Send(ctx context.Context, in *DocumentRequest, opts ...grpc.CallOption) (*DocumentReply, error)
	GetData(ctx context.Context, in *DocumentRequest, opts ...grpc.CallOption) (*DocumentReply, error)
}

type documentClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentClient(cc grpc.ClientConnInterface) DocumentClient {
	return &documentClient{cc}
}

func (c *documentClient) SayHello(ctx context.Context, in *DocumentRequest, opts ...grpc.CallOption) (*DocumentReply, error) {
	out := new(DocumentReply)
	err := c.cc.Invoke(ctx, "/document.Document/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentClient) Send(ctx context.Context, in *DocumentRequest, opts ...grpc.CallOption) (*DocumentReply, error) {
	out := new(DocumentReply)
	err := c.cc.Invoke(ctx, "/document.Document/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentClient) GetData(ctx context.Context, in *DocumentRequest, opts ...grpc.CallOption) (*DocumentReply, error) {
	out := new(DocumentReply)
	err := c.cc.Invoke(ctx, "/document.Document/GetData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentServer is the server API for Document service.
type DocumentServer interface {
	// Sends a greeting
	SayHello(context.Context, *DocumentRequest) (*DocumentReply, error)
	Send(context.Context, *DocumentRequest) (*DocumentReply, error)
	GetData(context.Context, *DocumentRequest) (*DocumentReply, error)
}

// UnimplementedDocumentServer can be embedded to have forward compatible implementations.
type UnimplementedDocumentServer struct {
}

func (*UnimplementedDocumentServer) SayHello(context.Context, *DocumentRequest) (*DocumentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedDocumentServer) Send(context.Context, *DocumentRequest) (*DocumentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedDocumentServer) GetData(context.Context, *DocumentRequest) (*DocumentReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetData not implemented")
}

func RegisterDocumentServer(s *grpc.Server, srv DocumentServer) {
	s.RegisterService(&_Document_serviceDesc, srv)
}

func _Document_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document.Document/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServer).SayHello(ctx, req.(*DocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Document_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document.Document/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServer).Send(ctx, req.(*DocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Document_GetData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServer).GetData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/document.Document/GetData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServer).GetData(ctx, req.(*DocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Document_serviceDesc = grpc.ServiceDesc{
	ServiceName: "document.Document",
	HandlerType: (*DocumentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Document_SayHello_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _Document_Send_Handler,
		},
		{
			MethodName: "GetData",
			Handler:    _Document_GetData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "document.proto",
}
