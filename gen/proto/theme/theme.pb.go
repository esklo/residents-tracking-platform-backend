// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: theme/theme.proto

package theme

import (
	empty "github.com/esklo/residents-tracking-platform/gen/proto/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title        string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Priority     int64  `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	DepartmentId string `protobuf:"bytes,3,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_theme_theme_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_theme_theme_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_theme_theme_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateRequest) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *CreateRequest) GetDepartmentId() string {
	if x != nil {
		return x.DepartmentId
	}
	return ""
}

type ByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ByIdRequest) Reset() {
	*x = ByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_theme_theme_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByIdRequest) ProtoMessage() {}

func (x *ByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_theme_theme_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByIdRequest.ProtoReflect.Descriptor instead.
func (*ByIdRequest) Descriptor() ([]byte, []int) {
	return file_theme_theme_proto_rawDescGZIP(), []int{1}
}

func (x *ByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  *int64 `protobuf:"varint,1,opt,name=limit,proto3,oneof" json:"limit,omitempty"`
	Offset *int64 `protobuf:"varint,2,opt,name=offset,proto3,oneof" json:"offset,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_theme_theme_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_theme_theme_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_theme_theme_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetLimit() int64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *GetRequest) GetOffset() int64 {
	if x != nil && x.Offset != nil {
		return *x.Offset
	}
	return 0
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Themes []*Theme `protobuf:"bytes,1,rep,name=themes,proto3" json:"themes,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_theme_theme_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_theme_theme_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_theme_theme_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetThemes() []*Theme {
	if x != nil {
		return x.Themes
	}
	return nil
}

type Theme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        string                 `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Priority     int64                  `protobuf:"varint,3,opt,name=priority,proto3" json:"priority,omitempty"`
	DepartmentId string                 `protobuf:"bytes,4,opt,name=department_id,json=departmentId,proto3" json:"department_id,omitempty"`
	DeletedAt    *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=deleted_at,json=deletedAt,proto3,oneof" json:"deleted_at,omitempty"`
}

func (x *Theme) Reset() {
	*x = Theme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_theme_theme_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Theme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Theme) ProtoMessage() {}

func (x *Theme) ProtoReflect() protoreflect.Message {
	mi := &file_theme_theme_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Theme.ProtoReflect.Descriptor instead.
func (*Theme) Descriptor() ([]byte, []int) {
	return file_theme_theme_proto_rawDescGZIP(), []int{4}
}

func (x *Theme) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Theme) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Theme) GetPriority() int64 {
	if x != nil {
		return x.Priority
	}
	return 0
}

func (x *Theme) GetDepartmentId() string {
	if x != nil {
		return x.DepartmentId
	}
	return ""
}

func (x *Theme) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

var File_theme_theme_proto protoreflect.FileDescriptor

var file_theme_theme_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x2f, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x1a, 0x11, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x66,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x1d, 0x0a, 0x0b, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x59, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x22, 0x33, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x06, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0c, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x06, 0x74,
	0x68, 0x65, 0x6d, 0x65, 0x73, 0x22, 0xbd, 0x01, 0x0a, 0x05, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x3e, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x00, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x32, 0xe9, 0x01, 0x0a, 0x0c, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x14, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x54,
	0x68, 0x65, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x12, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x54, 0x68, 0x65, 0x6d,
	0x65, 0x12, 0x2c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x74, 0x68,
	0x65, 0x6d, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x74, 0x68, 0x65, 0x6d,
	0x65, 0x2e, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x1a, 0x0c, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x2e,
	0x54, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x12, 0x2e, 0x74, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x73, 0x6b, 0x6c, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x2d,
	0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x68, 0x65, 0x6d,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_theme_theme_proto_rawDescOnce sync.Once
	file_theme_theme_proto_rawDescData = file_theme_theme_proto_rawDesc
)

func file_theme_theme_proto_rawDescGZIP() []byte {
	file_theme_theme_proto_rawDescOnce.Do(func() {
		file_theme_theme_proto_rawDescData = protoimpl.X.CompressGZIP(file_theme_theme_proto_rawDescData)
	})
	return file_theme_theme_proto_rawDescData
}

var file_theme_theme_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_theme_theme_proto_goTypes = []interface{}{
	(*CreateRequest)(nil),         // 0: theme.CreateRequest
	(*ByIdRequest)(nil),           // 1: theme.ByIdRequest
	(*GetRequest)(nil),            // 2: theme.GetRequest
	(*GetResponse)(nil),           // 3: theme.GetResponse
	(*Theme)(nil),                 // 4: theme.Theme
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*empty.Empty)(nil),           // 6: empty.Empty
}
var file_theme_theme_proto_depIdxs = []int32{
	4, // 0: theme.GetResponse.themes:type_name -> theme.Theme
	5, // 1: theme.Theme.deleted_at:type_name -> google.protobuf.Timestamp
	0, // 2: theme.ThemeService.Create:input_type -> theme.CreateRequest
	1, // 3: theme.ThemeService.GetById:input_type -> theme.ByIdRequest
	2, // 4: theme.ThemeService.Get:input_type -> theme.GetRequest
	4, // 5: theme.ThemeService.Update:input_type -> theme.Theme
	1, // 6: theme.ThemeService.Delete:input_type -> theme.ByIdRequest
	4, // 7: theme.ThemeService.Create:output_type -> theme.Theme
	4, // 8: theme.ThemeService.GetById:output_type -> theme.Theme
	3, // 9: theme.ThemeService.Get:output_type -> theme.GetResponse
	4, // 10: theme.ThemeService.Update:output_type -> theme.Theme
	6, // 11: theme.ThemeService.Delete:output_type -> empty.Empty
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_theme_theme_proto_init() }
func file_theme_theme_proto_init() {
	if File_theme_theme_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_theme_theme_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_theme_theme_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByIdRequest); i {
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
		file_theme_theme_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_theme_theme_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_theme_theme_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Theme); i {
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
	file_theme_theme_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_theme_theme_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_theme_theme_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_theme_theme_proto_goTypes,
		DependencyIndexes: file_theme_theme_proto_depIdxs,
		MessageInfos:      file_theme_theme_proto_msgTypes,
	}.Build()
	File_theme_theme_proto = out.File
	file_theme_theme_proto_rawDesc = nil
	file_theme_theme_proto_goTypes = nil
	file_theme_theme_proto_depIdxs = nil
}
