// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/connection_app/connection_app.proto

package pb

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConnectionAppBodyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnectionName string `protobuf:"bytes,1,opt,name=connection_name,json=connectionName,proto3" json:"connection_name,omitempty"`
	ConnectionType string `protobuf:"bytes,2,opt,name=connection_type,json=connectionType,proto3" json:"connection_type,omitempty"`
	AppId          string `protobuf:"bytes,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	QueueId        string `protobuf:"bytes,4,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	Status         string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ConnectionAppBodyRequest) Reset() {
	*x = ConnectionAppBodyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_connection_app_connection_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionAppBodyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionAppBodyRequest) ProtoMessage() {}

func (x *ConnectionAppBodyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_connection_app_connection_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionAppBodyRequest.ProtoReflect.Descriptor instead.
func (*ConnectionAppBodyRequest) Descriptor() ([]byte, []int) {
	return file_proto_connection_app_connection_app_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectionAppBodyRequest) GetConnectionName() string {
	if x != nil {
		return x.ConnectionName
	}
	return ""
}

func (x *ConnectionAppBodyRequest) GetConnectionType() string {
	if x != nil {
		return x.ConnectionType
	}
	return ""
}

func (x *ConnectionAppBodyRequest) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *ConnectionAppBodyRequest) GetQueueId() string {
	if x != nil {
		return x.QueueId
	}
	return ""
}

func (x *ConnectionAppBodyRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ConnectionAppRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit          int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset         int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	ConnectionName string `protobuf:"bytes,3,opt,name=connection_name,json=connectionName,proto3" json:"connection_name,omitempty"`
	ConnectionType string `protobuf:"bytes,4,opt,name=connection_type,json=connectionType,proto3" json:"connection_type,omitempty"`
	Status         string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ConnectionAppRequest) Reset() {
	*x = ConnectionAppRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_connection_app_connection_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionAppRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionAppRequest) ProtoMessage() {}

func (x *ConnectionAppRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_connection_app_connection_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionAppRequest.ProtoReflect.Descriptor instead.
func (*ConnectionAppRequest) Descriptor() ([]byte, []int) {
	return file_proto_connection_app_connection_app_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectionAppRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ConnectionAppRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ConnectionAppRequest) GetConnectionName() string {
	if x != nil {
		return x.ConnectionName
	}
	return ""
}

func (x *ConnectionAppRequest) GetConnectionType() string {
	if x != nil {
		return x.ConnectionType
	}
	return ""
}

func (x *ConnectionAppRequest) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type ConnectionAppResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Id      string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ConnectionAppResponse) Reset() {
	*x = ConnectionAppResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_connection_app_connection_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionAppResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionAppResponse) ProtoMessage() {}

func (x *ConnectionAppResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_connection_app_connection_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionAppResponse.ProtoReflect.Descriptor instead.
func (*ConnectionAppResponse) Descriptor() ([]byte, []int) {
	return file_proto_connection_app_connection_app_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectionAppResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ConnectionAppResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ConnectionAppResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ConnectionAppStructResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string             `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string             `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*structpb.Struct `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty"`
	Total   int32              `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *ConnectionAppStructResponse) Reset() {
	*x = ConnectionAppStructResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_connection_app_connection_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionAppStructResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionAppStructResponse) ProtoMessage() {}

func (x *ConnectionAppStructResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_connection_app_connection_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionAppStructResponse.ProtoReflect.Descriptor instead.
func (*ConnectionAppStructResponse) Descriptor() ([]byte, []int) {
	return file_proto_connection_app_connection_app_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectionAppStructResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ConnectionAppStructResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ConnectionAppStructResponse) GetData() []*structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ConnectionAppStructResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_proto_connection_app_connection_app_proto protoreflect.FileDescriptor

var file_proto_connection_app_connection_app_proto_rawDesc = []byte{
	0x0a, 0x29, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70,
	0x70, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01, 0x0a, 0x18,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70, 0x42, 0x6f, 0x64,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70,
	0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0xae, 0x01, 0x0a, 0x14, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x55, 0x0a, 0x15, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x8e, 0x01, 0x0a,
	0x1b, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xc3, 0x02,
	0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70, 0x12,
	0x98, 0x01, 0x0a, 0x11, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x70, 0x70, 0x12, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f,
	0x62, 0x73, 0x73, 0x2d, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x70, 0x12, 0x96, 0x01, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70, 0x12,
	0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x70, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61,
	0x70, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x70, 0x70,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x62, 0x73, 0x73, 0x2d, 0x63, 0x68, 0x61,
	0x74, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2d,
	0x61, 0x70, 0x70, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x74, 0x65, 0x6c, 0x34, 0x76, 0x6e, 0x2f, 0x66, 0x69, 0x6e, 0x73, 0x2d, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_connection_app_connection_app_proto_rawDescOnce sync.Once
	file_proto_connection_app_connection_app_proto_rawDescData = file_proto_connection_app_connection_app_proto_rawDesc
)

func file_proto_connection_app_connection_app_proto_rawDescGZIP() []byte {
	file_proto_connection_app_connection_app_proto_rawDescOnce.Do(func() {
		file_proto_connection_app_connection_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_connection_app_connection_app_proto_rawDescData)
	})
	return file_proto_connection_app_connection_app_proto_rawDescData
}

var file_proto_connection_app_connection_app_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_connection_app_connection_app_proto_goTypes = []interface{}{
	(*ConnectionAppBodyRequest)(nil),    // 0: proto.connection_app.ConnectionAppBodyRequest
	(*ConnectionAppRequest)(nil),        // 1: proto.connection_app.ConnectionAppRequest
	(*ConnectionAppResponse)(nil),       // 2: proto.connection_app.ConnectionAppResponse
	(*ConnectionAppStructResponse)(nil), // 3: proto.connection_app.ConnectionAppStructResponse
	(*structpb.Struct)(nil),             // 4: google.protobuf.Struct
}
var file_proto_connection_app_connection_app_proto_depIdxs = []int32{
	4, // 0: proto.connection_app.ConnectionAppStructResponse.data:type_name -> google.protobuf.Struct
	0, // 1: proto.connection_app.ConnectionApp.PostConnectionApp:input_type -> proto.connection_app.ConnectionAppBodyRequest
	1, // 2: proto.connection_app.ConnectionApp.GetConnectionApp:input_type -> proto.connection_app.ConnectionAppRequest
	2, // 3: proto.connection_app.ConnectionApp.PostConnectionApp:output_type -> proto.connection_app.ConnectionAppResponse
	3, // 4: proto.connection_app.ConnectionApp.GetConnectionApp:output_type -> proto.connection_app.ConnectionAppStructResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_connection_app_connection_app_proto_init() }
func file_proto_connection_app_connection_app_proto_init() {
	if File_proto_connection_app_connection_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_connection_app_connection_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionAppBodyRequest); i {
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
		file_proto_connection_app_connection_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionAppRequest); i {
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
		file_proto_connection_app_connection_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionAppResponse); i {
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
		file_proto_connection_app_connection_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionAppStructResponse); i {
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
			RawDescriptor: file_proto_connection_app_connection_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_connection_app_connection_app_proto_goTypes,
		DependencyIndexes: file_proto_connection_app_connection_app_proto_depIdxs,
		MessageInfos:      file_proto_connection_app_connection_app_proto_msgTypes,
	}.Build()
	File_proto_connection_app_connection_app_proto = out.File
	file_proto_connection_app_connection_app_proto_rawDesc = nil
	file_proto_connection_app_connection_app_proto_goTypes = nil
	file_proto_connection_app_connection_app_proto_depIdxs = nil
}
