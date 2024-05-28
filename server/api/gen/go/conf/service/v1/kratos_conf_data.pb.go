// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: conf/service/v1/kratos_conf_data.proto

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

// 数据
type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database *Data_Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"` // 数据库DSN
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_service_v1_kratos_conf_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_conf_service_v1_kratos_conf_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_conf_service_v1_kratos_conf_data_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetDatabase() *Data_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

// 数据库
type Data_Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver                string               `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`                                                               // 驱动名：mysql、postgresql、mongodb、sqlite……
	Source                string               `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`                                                               // 数据源（DSN字符串）
	Migrate               bool                 `protobuf:"varint,10,opt,name=migrate,proto3" json:"migrate,omitempty"`                                                           // 数据迁移开关
	Debug                 bool                 `protobuf:"varint,11,opt,name=debug,proto3" json:"debug,omitempty"`                                                               // 调试开关
	EnableTrace           bool                 `protobuf:"varint,12,opt,name=enable_trace,json=enableTrace,proto3" json:"enable_trace,omitempty"`                                // 链路追踪开关
	EnableMetrics         bool                 `protobuf:"varint,13,opt,name=enable_metrics,json=enableMetrics,proto3" json:"enable_metrics,omitempty"`                          // 性能分析开关
	MaxIdleConnections    int32                `protobuf:"varint,20,opt,name=max_idle_connections,json=maxIdleConnections,proto3" json:"max_idle_connections,omitempty"`         // 连接池最大空闲连接数
	MaxOpenConnections    int32                `protobuf:"varint,21,opt,name=max_open_connections,json=maxOpenConnections,proto3" json:"max_open_connections,omitempty"`         // 连接池最大打开连接数
	ConnectionMaxLifetime *durationpb.Duration `protobuf:"bytes,22,opt,name=connection_max_lifetime,json=connectionMaxLifetime,proto3" json:"connection_max_lifetime,omitempty"` // 连接可重用的最大时间长度
}

func (x *Data_Database) Reset() {
	*x = Data_Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_service_v1_kratos_conf_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Database) ProtoMessage() {}

func (x *Data_Database) ProtoReflect() protoreflect.Message {
	mi := &file_conf_service_v1_kratos_conf_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Database.ProtoReflect.Descriptor instead.
func (*Data_Database) Descriptor() ([]byte, []int) {
	return file_conf_service_v1_kratos_conf_data_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Data_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Data_Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Data_Database) GetMigrate() bool {
	if x != nil {
		return x.Migrate
	}
	return false
}

func (x *Data_Database) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *Data_Database) GetEnableTrace() bool {
	if x != nil {
		return x.EnableTrace
	}
	return false
}

func (x *Data_Database) GetEnableMetrics() bool {
	if x != nil {
		return x.EnableMetrics
	}
	return false
}

func (x *Data_Database) GetMaxIdleConnections() int32 {
	if x != nil {
		return x.MaxIdleConnections
	}
	return 0
}

func (x *Data_Database) GetMaxOpenConnections() int32 {
	if x != nil {
		return x.MaxOpenConnections
	}
	return 0
}

func (x *Data_Database) GetConnectionMaxLifetime() *durationpb.Duration {
	if x != nil {
		return x.ConnectionMaxLifetime
	}
	return nil
}

// redis
type Data_Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network       string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`                                   // 网络
	Addr          string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`                                         // 服务端地址
	Password      string               `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`                                 // 密码
	Db            int32                `protobuf:"varint,4,opt,name=db,proto3" json:"db,omitempty"`                                            // 数据库索引
	DialTimeout   *durationpb.Duration `protobuf:"bytes,5,opt,name=dial_timeout,json=dialTimeout,proto3" json:"dial_timeout,omitempty"`        // 连接超时时间
	ReadTimeout   *durationpb.Duration `protobuf:"bytes,6,opt,name=read_timeout,json=readTimeout,proto3" json:"read_timeout,omitempty"`        // 读取超时时间
	WriteTimeout  *durationpb.Duration `protobuf:"bytes,7,opt,name=write_timeout,json=writeTimeout,proto3" json:"write_timeout,omitempty"`     // 写入超时时间
	EnableTracing bool                 `protobuf:"varint,8,opt,name=enable_tracing,json=enableTracing,proto3" json:"enable_tracing,omitempty"` // 打开链路追踪
	EnableMetrics bool                 `protobuf:"varint,9,opt,name=enable_metrics,json=enableMetrics,proto3" json:"enable_metrics,omitempty"` // 打开性能度量
}

func (x *Data_Redis) Reset() {
	*x = Data_Redis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_service_v1_kratos_conf_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data_Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data_Redis) ProtoMessage() {}

func (x *Data_Redis) ProtoReflect() protoreflect.Message {
	mi := &file_conf_service_v1_kratos_conf_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data_Redis.ProtoReflect.Descriptor instead.
func (*Data_Redis) Descriptor() ([]byte, []int) {
	return file_conf_service_v1_kratos_conf_data_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Data_Redis) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Data_Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Data_Redis) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Data_Redis) GetDb() int32 {
	if x != nil {
		return x.Db
	}
	return 0
}

func (x *Data_Redis) GetDialTimeout() *durationpb.Duration {
	if x != nil {
		return x.DialTimeout
	}
	return nil
}

func (x *Data_Redis) GetReadTimeout() *durationpb.Duration {
	if x != nil {
		return x.ReadTimeout
	}
	return nil
}

func (x *Data_Redis) GetWriteTimeout() *durationpb.Duration {
	if x != nil {
		return x.WriteTimeout
	}
	return nil
}

func (x *Data_Redis) GetEnableTracing() bool {
	if x != nil {
		return x.EnableTracing
	}
	return false
}

func (x *Data_Redis) GetEnableMetrics() bool {
	if x != nil {
		return x.EnableMetrics
	}
	return false
}

var File_conf_service_v1_kratos_conf_data_proto protoreflect.FileDescriptor

var file_conf_service_v1_kratos_conf_data_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x06, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x3a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x1a, 0xeb,
	0x02, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x69, 0x67, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x6d, 0x69,
	0x67, 0x72, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0b, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x63, 0x65, 0x12, 0x25,
	0x0a, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x6c,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x49, 0x64, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x6f,
	0x70, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x51, 0x0a, 0x17, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x61, 0x78, 0x5f, 0x6c, 0x69, 0x66, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x61, 0x78, 0x4c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0xeb, 0x02, 0x0a,
	0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x64, 0x62,
	0x12, 0x3c, 0x0a, 0x0c, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3c,
	0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x72, 0x65, 0x61, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x3e, 0x0a, 0x0d,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x72, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x42, 0xb5, 0x01, 0x0a, 0x13, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x42, 0x13, 0x4b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x44, 0x61,
	0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x6d, 0x6f, 0x73, 0x6f, 0x6e,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x53, 0x58, 0xaa, 0x02, 0x0f, 0x43,
	0x6f, 0x6e, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x0f, 0x43, 0x6f, 0x6e, 0x66, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x1b, 0x43, 0x6f, 0x6e, 0x66, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x11, 0x43, 0x6f, 0x6e, 0x66, 0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_service_v1_kratos_conf_data_proto_rawDescOnce sync.Once
	file_conf_service_v1_kratos_conf_data_proto_rawDescData = file_conf_service_v1_kratos_conf_data_proto_rawDesc
)

func file_conf_service_v1_kratos_conf_data_proto_rawDescGZIP() []byte {
	file_conf_service_v1_kratos_conf_data_proto_rawDescOnce.Do(func() {
		file_conf_service_v1_kratos_conf_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_service_v1_kratos_conf_data_proto_rawDescData)
	})
	return file_conf_service_v1_kratos_conf_data_proto_rawDescData
}

var file_conf_service_v1_kratos_conf_data_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_conf_service_v1_kratos_conf_data_proto_goTypes = []interface{}{
	(*Data)(nil),                // 0: conf.service.v1.Data
	(*Data_Database)(nil),       // 1: conf.service.v1.Data.Database
	(*Data_Redis)(nil),          // 2: conf.service.v1.Data.Redis
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
}
var file_conf_service_v1_kratos_conf_data_proto_depIdxs = []int32{
	1, // 0: conf.service.v1.Data.database:type_name -> conf.service.v1.Data.Database
	3, // 1: conf.service.v1.Data.Database.connection_max_lifetime:type_name -> google.protobuf.Duration
	3, // 2: conf.service.v1.Data.Redis.dial_timeout:type_name -> google.protobuf.Duration
	3, // 3: conf.service.v1.Data.Redis.read_timeout:type_name -> google.protobuf.Duration
	3, // 4: conf.service.v1.Data.Redis.write_timeout:type_name -> google.protobuf.Duration
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_conf_service_v1_kratos_conf_data_proto_init() }
func file_conf_service_v1_kratos_conf_data_proto_init() {
	if File_conf_service_v1_kratos_conf_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_service_v1_kratos_conf_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data); i {
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
		file_conf_service_v1_kratos_conf_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Database); i {
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
		file_conf_service_v1_kratos_conf_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Data_Redis); i {
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
			RawDescriptor: file_conf_service_v1_kratos_conf_data_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_service_v1_kratos_conf_data_proto_goTypes,
		DependencyIndexes: file_conf_service_v1_kratos_conf_data_proto_depIdxs,
		MessageInfos:      file_conf_service_v1_kratos_conf_data_proto_msgTypes,
	}.Build()
	File_conf_service_v1_kratos_conf_data_proto = out.File
	file_conf_service_v1_kratos_conf_data_proto_rawDesc = nil
	file_conf_service_v1_kratos_conf_data_proto_goTypes = nil
	file_conf_service_v1_kratos_conf_data_proto_depIdxs = nil
}
