// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: conf/service/v1/kratos_conf_server.proto

package servicev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 服务器
type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RPC
	Rest *Server_REST `protobuf:"bytes,1,opt,name=rest,proto3" json:"rest,omitempty"` // REST服务
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_service_v1_kratos_conf_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_service_v1_kratos_conf_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_conf_service_v1_kratos_conf_server_proto_rawDescGZIP(), []int{0}
}

func (x *Server) GetRest() *Server_REST {
	if x != nil {
		return x.Rest
	}
	return nil
}

// REST
type Server_REST struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network       string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`                                   // 网络
	Addr          string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`                                         // 服务监听地址
	Timeout       *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`                                   // 超时时间
	Cors          *Server_REST_CORS    `protobuf:"bytes,4,opt,name=cors,proto3" json:"cors,omitempty"`                                         // 服务监听地址
	Middleware    *Middleware          `protobuf:"bytes,5,opt,name=middleware,proto3" json:"middleware,omitempty"`                             // 中间件
	EnableSwagger bool                 `protobuf:"varint,6,opt,name=enable_swagger,json=enableSwagger,proto3" json:"enable_swagger,omitempty"` // 启用SwaggerUI
	EnablePprof   bool                 `protobuf:"varint,7,opt,name=enable_pprof,json=enablePprof,proto3" json:"enable_pprof,omitempty"`       // 启用pprof
}

func (x *Server_REST) Reset() {
	*x = Server_REST{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_service_v1_kratos_conf_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_REST) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_REST) ProtoMessage() {}

func (x *Server_REST) ProtoReflect() protoreflect.Message {
	mi := &file_conf_service_v1_kratos_conf_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_REST.ProtoReflect.Descriptor instead.
func (*Server_REST) Descriptor() ([]byte, []int) {
	return file_conf_service_v1_kratos_conf_server_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Server_REST) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Server_REST) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server_REST) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Server_REST) GetCors() *Server_REST_CORS {
	if x != nil {
		return x.Cors
	}
	return nil
}

func (x *Server_REST) GetMiddleware() *Middleware {
	if x != nil {
		return x.Middleware
	}
	return nil
}

func (x *Server_REST) GetEnableSwagger() bool {
	if x != nil {
		return x.EnableSwagger
	}
	return false
}

func (x *Server_REST) GetEnablePprof() bool {
	if x != nil {
		return x.EnablePprof
	}
	return false
}

type Server_REST_CORS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers []string `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"` //
	Methods []string `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty"` //
	Origins []string `protobuf:"bytes,3,rep,name=origins,proto3" json:"origins,omitempty"` //
}

func (x *Server_REST_CORS) Reset() {
	*x = Server_REST_CORS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_service_v1_kratos_conf_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server_REST_CORS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server_REST_CORS) ProtoMessage() {}

func (x *Server_REST_CORS) ProtoReflect() protoreflect.Message {
	mi := &file_conf_service_v1_kratos_conf_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server_REST_CORS.ProtoReflect.Descriptor instead.
func (*Server_REST_CORS) Descriptor() ([]byte, []int) {
	return file_conf_service_v1_kratos_conf_server_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Server_REST_CORS) GetHeaders() []string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Server_REST_CORS) GetMethods() []string {
	if x != nil {
		return x.Methods
	}
	return nil
}

func (x *Server_REST_CORS) GetOrigins() []string {
	if x != nil {
		return x.Origins
	}
	return nil
}

var File_conf_service_v1_kratos_conf_server_proto protoreflect.FileDescriptor

var file_conf_service_v1_kratos_conf_server_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x6e, 0x66,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x63, 0x6f, 0x6e,
	0x66, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x6d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77,
	0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x03, 0x0a, 0x06, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x45, 0x53, 0x54,
	0x52, 0x04, 0x72, 0x65, 0x73, 0x74, 0x1a, 0xfd, 0x02, 0x0a, 0x04, 0x52, 0x45, 0x53, 0x54, 0x12,
	0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x33, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x63, 0x6f, 0x72, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x52, 0x45, 0x53, 0x54, 0x2e, 0x43,
	0x4f, 0x52, 0x53, 0x52, 0x04, 0x63, 0x6f, 0x72, 0x73, 0x12, 0x3b, 0x0a, 0x0a, 0x6d, 0x69, 0x64,
	0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x69, 0x64, 0x64, 0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x52, 0x0a, 0x6d, 0x69, 0x64, 0x64,
	0x6c, 0x65, 0x77, 0x61, 0x72, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x5f, 0x73, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x70, 0x72, 0x6f, 0x66, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x70, 0x72, 0x6f, 0x66,
	0x1a, 0x54, 0x0a, 0x04, 0x43, 0x4f, 0x52, 0x53, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x42, 0xb7, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x15,
	0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x6d, 0x6f, 0x73, 0x6f, 0x6e, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x43, 0x6f, 0x6e,
	0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0f, 0x43,
	0x6f, 0x6e, 0x66, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1b, 0x43, 0x6f, 0x6e, 0x66, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x43,
	0x6f, 0x6e, 0x66, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_service_v1_kratos_conf_server_proto_rawDescOnce sync.Once
	file_conf_service_v1_kratos_conf_server_proto_rawDescData = file_conf_service_v1_kratos_conf_server_proto_rawDesc
)

func file_conf_service_v1_kratos_conf_server_proto_rawDescGZIP() []byte {
	file_conf_service_v1_kratos_conf_server_proto_rawDescOnce.Do(func() {
		file_conf_service_v1_kratos_conf_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_service_v1_kratos_conf_server_proto_rawDescData)
	})
	return file_conf_service_v1_kratos_conf_server_proto_rawDescData
}

var file_conf_service_v1_kratos_conf_server_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_conf_service_v1_kratos_conf_server_proto_goTypes = []interface{}{
	(*Server)(nil),              // 0: conf.service.v1.Server
	(*Server_REST)(nil),         // 1: conf.service.v1.Server.REST
	(*Server_REST_CORS)(nil),    // 2: conf.service.v1.Server.REST.CORS
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
	(*Middleware)(nil),          // 4: conf.service.v1.Middleware
}
var file_conf_service_v1_kratos_conf_server_proto_depIdxs = []int32{
	1, // 0: conf.service.v1.Server.rest:type_name -> conf.service.v1.Server.REST
	3, // 1: conf.service.v1.Server.REST.timeout:type_name -> google.protobuf.Duration
	2, // 2: conf.service.v1.Server.REST.cors:type_name -> conf.service.v1.Server.REST.CORS
	4, // 3: conf.service.v1.Server.REST.middleware:type_name -> conf.service.v1.Middleware
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_conf_service_v1_kratos_conf_server_proto_init() }
func file_conf_service_v1_kratos_conf_server_proto_init() {
	if File_conf_service_v1_kratos_conf_server_proto != nil {
		return
	}
	file_conf_service_v1_kratos_conf_middleware_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_conf_service_v1_kratos_conf_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_conf_service_v1_kratos_conf_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_REST); i {
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
		file_conf_service_v1_kratos_conf_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server_REST_CORS); i {
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
			RawDescriptor: file_conf_service_v1_kratos_conf_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_service_v1_kratos_conf_server_proto_goTypes,
		DependencyIndexes: file_conf_service_v1_kratos_conf_server_proto_depIdxs,
		MessageInfos:      file_conf_service_v1_kratos_conf_server_proto_msgTypes,
	}.Build()
	File_conf_service_v1_kratos_conf_server_proto = out.File
	file_conf_service_v1_kratos_conf_server_proto_rawDesc = nil
	file_conf_service_v1_kratos_conf_server_proto_goTypes = nil
	file_conf_service_v1_kratos_conf_server_proto_depIdxs = nil
}
