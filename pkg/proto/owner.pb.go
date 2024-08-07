// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.26.1
// source: pkg/proto/owner.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Owner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64  `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Owner) Reset() {
	*x = Owner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_owner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Owner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Owner) ProtoMessage() {}

func (x *Owner) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_owner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Owner.ProtoReflect.Descriptor instead.
func (*Owner) Descriptor() ([]byte, []int) {
	return file_pkg_proto_owner_proto_rawDescGZIP(), []int{0}
}

func (x *Owner) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Owner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Owner) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Owner) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type ListOwnersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page    int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage int64 `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
}

func (x *ListOwnersRequest) Reset() {
	*x = ListOwnersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_owner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOwnersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOwnersRequest) ProtoMessage() {}

func (x *ListOwnersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_owner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOwnersRequest.ProtoReflect.Descriptor instead.
func (*ListOwnersRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_owner_proto_rawDescGZIP(), []int{1}
}

func (x *ListOwnersRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListOwnersRequest) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type ListOwnersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows       []*Owner `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
	Page       int64    `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
	PerPage    int64    `protobuf:"varint,3,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	Total      int64    `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty"`
	TotalPages int64    `protobuf:"varint,5,opt,name=total_pages,json=totalPages,proto3" json:"total_pages,omitempty"`
}

func (x *ListOwnersResponse) Reset() {
	*x = ListOwnersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_owner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOwnersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOwnersResponse) ProtoMessage() {}

func (x *ListOwnersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_owner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOwnersResponse.ProtoReflect.Descriptor instead.
func (*ListOwnersResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_owner_proto_rawDescGZIP(), []int{2}
}

func (x *ListOwnersResponse) GetRows() []*Owner {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *ListOwnersResponse) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListOwnersResponse) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *ListOwnersResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListOwnersResponse) GetTotalPages() int64 {
	if x != nil {
		return x.TotalPages
	}
	return 0
}

type GetOwnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetOwnerRequest) Reset() {
	*x = GetOwnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_owner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOwnerRequest) ProtoMessage() {}

func (x *GetOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_owner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOwnerRequest.ProtoReflect.Descriptor instead.
func (*GetOwnerRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_owner_proto_rawDescGZIP(), []int{3}
}

func (x *GetOwnerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetOwnerRequest) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

type GetOwnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt int64  `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64  `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *GetOwnerResponse) Reset() {
	*x = GetOwnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_owner_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOwnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOwnerResponse) ProtoMessage() {}

func (x *GetOwnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_owner_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOwnerResponse.ProtoReflect.Descriptor instead.
func (*GetOwnerResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_owner_proto_rawDescGZIP(), []int{4}
}

func (x *GetOwnerResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetOwnerResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetOwnerResponse) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *GetOwnerResponse) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreateOwnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateOwnerRequest) Reset() {
	*x = CreateOwnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_owner_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOwnerRequest) ProtoMessage() {}

func (x *CreateOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_owner_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOwnerRequest.ProtoReflect.Descriptor instead.
func (*CreateOwnerRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_owner_proto_rawDescGZIP(), []int{5}
}

func (x *CreateOwnerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateOwnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateOwnerResponse) Reset() {
	*x = CreateOwnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_owner_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOwnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOwnerResponse) ProtoMessage() {}

func (x *CreateOwnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_owner_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOwnerResponse.ProtoReflect.Descriptor instead.
func (*CreateOwnerResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_owner_proto_rawDescGZIP(), []int{6}
}

func (x *CreateOwnerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateOwnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateOwnerRequest) Reset() {
	*x = UpdateOwnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_owner_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOwnerRequest) ProtoMessage() {}

func (x *UpdateOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_owner_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOwnerRequest.ProtoReflect.Descriptor instead.
func (*UpdateOwnerRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_owner_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateOwnerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateOwnerRequest) GetName() *wrapperspb.StringValue {
	if x != nil {
		return x.Name
	}
	return nil
}

type UpdateOwnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateOwnerResponse) Reset() {
	*x = UpdateOwnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_owner_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOwnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOwnerResponse) ProtoMessage() {}

func (x *UpdateOwnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_owner_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOwnerResponse.ProtoReflect.Descriptor instead.
func (*UpdateOwnerResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_owner_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateOwnerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteOwnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteOwnerRequest) Reset() {
	*x = DeleteOwnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_owner_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOwnerRequest) ProtoMessage() {}

func (x *DeleteOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_owner_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOwnerRequest.ProtoReflect.Descriptor instead.
func (*DeleteOwnerRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_owner_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteOwnerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteOwnerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteOwnerResponse) Reset() {
	*x = DeleteOwnerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_owner_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOwnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOwnerResponse) ProtoMessage() {}

func (x *DeleteOwnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_owner_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOwnerResponse.ProtoReflect.Descriptor instead.
func (*DeleteOwnerResponse) Descriptor() ([]byte, []int) {
	return file_pkg_proto_owner_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteOwnerResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pkg_proto_owner_proto protoreflect.FileDescriptor

var file_pkg_proto_owner_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x05, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x42, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70,
	0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70,
	0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x96, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x1f,
	0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x73, 0x22,
	0x53, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x74, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x28, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x56, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a,
	0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x24,
	0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x2f, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xae, 0x02, 0x0a, 0x0c, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x73, 0x12, 0x12, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x31, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65,
	0x72, 0x12, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f,
	0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a,
	0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x13, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0b, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x13, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_proto_owner_proto_rawDescOnce sync.Once
	file_pkg_proto_owner_proto_rawDescData = file_pkg_proto_owner_proto_rawDesc
)

func file_pkg_proto_owner_proto_rawDescGZIP() []byte {
	file_pkg_proto_owner_proto_rawDescOnce.Do(func() {
		file_pkg_proto_owner_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_owner_proto_rawDescData)
	})
	return file_pkg_proto_owner_proto_rawDescData
}

var file_pkg_proto_owner_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pkg_proto_owner_proto_goTypes = []interface{}{
	(*Owner)(nil),                  // 0: Owner
	(*ListOwnersRequest)(nil),      // 1: ListOwnersRequest
	(*ListOwnersResponse)(nil),     // 2: ListOwnersResponse
	(*GetOwnerRequest)(nil),        // 3: GetOwnerRequest
	(*GetOwnerResponse)(nil),       // 4: GetOwnerResponse
	(*CreateOwnerRequest)(nil),     // 5: CreateOwnerRequest
	(*CreateOwnerResponse)(nil),    // 6: CreateOwnerResponse
	(*UpdateOwnerRequest)(nil),     // 7: UpdateOwnerRequest
	(*UpdateOwnerResponse)(nil),    // 8: UpdateOwnerResponse
	(*DeleteOwnerRequest)(nil),     // 9: DeleteOwnerRequest
	(*DeleteOwnerResponse)(nil),    // 10: DeleteOwnerResponse
	(*wrapperspb.StringValue)(nil), // 11: google.protobuf.StringValue
}
var file_pkg_proto_owner_proto_depIdxs = []int32{
	0,  // 0: ListOwnersResponse.rows:type_name -> Owner
	11, // 1: GetOwnerRequest.name:type_name -> google.protobuf.StringValue
	11, // 2: UpdateOwnerRequest.name:type_name -> google.protobuf.StringValue
	1,  // 3: OwnerService.ListOwners:input_type -> ListOwnersRequest
	3,  // 4: OwnerService.GetOwner:input_type -> GetOwnerRequest
	5,  // 5: OwnerService.CreateOwner:input_type -> CreateOwnerRequest
	7,  // 6: OwnerService.UpdateOwner:input_type -> UpdateOwnerRequest
	9,  // 7: OwnerService.DeleteOwner:input_type -> DeleteOwnerRequest
	2,  // 8: OwnerService.ListOwners:output_type -> ListOwnersResponse
	4,  // 9: OwnerService.GetOwner:output_type -> GetOwnerResponse
	6,  // 10: OwnerService.CreateOwner:output_type -> CreateOwnerResponse
	8,  // 11: OwnerService.UpdateOwner:output_type -> UpdateOwnerResponse
	10, // 12: OwnerService.DeleteOwner:output_type -> DeleteOwnerResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_proto_owner_proto_init() }
func file_pkg_proto_owner_proto_init() {
	if File_pkg_proto_owner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_owner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Owner); i {
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
		file_pkg_proto_owner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOwnersRequest); i {
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
		file_pkg_proto_owner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOwnersResponse); i {
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
		file_pkg_proto_owner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOwnerRequest); i {
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
		file_pkg_proto_owner_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOwnerResponse); i {
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
		file_pkg_proto_owner_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOwnerRequest); i {
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
		file_pkg_proto_owner_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOwnerResponse); i {
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
		file_pkg_proto_owner_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOwnerRequest); i {
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
		file_pkg_proto_owner_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOwnerResponse); i {
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
		file_pkg_proto_owner_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOwnerRequest); i {
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
		file_pkg_proto_owner_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOwnerResponse); i {
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
			RawDescriptor: file_pkg_proto_owner_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_owner_proto_goTypes,
		DependencyIndexes: file_pkg_proto_owner_proto_depIdxs,
		MessageInfos:      file_pkg_proto_owner_proto_msgTypes,
	}.Build()
	File_pkg_proto_owner_proto = out.File
	file_pkg_proto_owner_proto_rawDesc = nil
	file_pkg_proto_owner_proto_goTypes = nil
	file_pkg_proto_owner_proto_depIdxs = nil
}
