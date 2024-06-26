// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.21.12
// source: systemuser.proto

package user_service

import (
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

type SystemUserEmpty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SystemUserEmpty) Reset() {
	*x = SystemUserEmpty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_systemuser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUserEmpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUserEmpty) ProtoMessage() {}

func (x *SystemUserEmpty) ProtoReflect() protoreflect.Message {
	mi := &file_systemuser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUserEmpty.ProtoReflect.Descriptor instead.
func (*SystemUserEmpty) Descriptor() ([]byte, []int) {
	return file_systemuser_proto_rawDescGZIP(), []int{0}
}

type SystemUserPrimaryKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SystemUserPrimaryKey) Reset() {
	*x = SystemUserPrimaryKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_systemuser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUserPrimaryKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUserPrimaryKey) ProtoMessage() {}

func (x *SystemUserPrimaryKey) ProtoReflect() protoreflect.Message {
	mi := &file_systemuser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUserPrimaryKey.ProtoReflect.Descriptor instead.
func (*SystemUserPrimaryKey) Descriptor() ([]byte, []int) {
	return file_systemuser_proto_rawDescGZIP(), []int{1}
}

func (x *SystemUserPrimaryKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SystemUserGmail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gmail string `protobuf:"bytes,1,opt,name=gmail,proto3" json:"gmail,omitempty"`
}

func (x *SystemUserGmail) Reset() {
	*x = SystemUserGmail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_systemuser_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUserGmail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUserGmail) ProtoMessage() {}

func (x *SystemUserGmail) ProtoReflect() protoreflect.Message {
	mi := &file_systemuser_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUserGmail.ProtoReflect.Descriptor instead.
func (*SystemUserGmail) Descriptor() ([]byte, []int) {
	return file_systemuser_proto_rawDescGZIP(), []int{2}
}

func (x *SystemUserGmail) GetGmail() string {
	if x != nil {
		return x.Gmail
	}
	return ""
}

type CreateSystemUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Gmail string `protobuf:"bytes,3,opt,name=gmail,proto3" json:"gmail,omitempty"`
	Name  string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Role  string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *CreateSystemUser) Reset() {
	*x = CreateSystemUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_systemuser_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSystemUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSystemUser) ProtoMessage() {}

func (x *CreateSystemUser) ProtoReflect() protoreflect.Message {
	mi := &file_systemuser_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSystemUser.ProtoReflect.Descriptor instead.
func (*CreateSystemUser) Descriptor() ([]byte, []int) {
	return file_systemuser_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSystemUser) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateSystemUser) GetGmail() string {
	if x != nil {
		return x.Gmail
	}
	return ""
}

func (x *CreateSystemUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSystemUser) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type SystemUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Phone     string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Gmail     string `protobuf:"bytes,3,opt,name=gmail,proto3" json:"gmail,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Role      string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
	CreatedAt string `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *SystemUser) Reset() {
	*x = SystemUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_systemuser_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemUser) ProtoMessage() {}

func (x *SystemUser) ProtoReflect() protoreflect.Message {
	mi := &file_systemuser_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemUser.ProtoReflect.Descriptor instead.
func (*SystemUser) Descriptor() ([]byte, []int) {
	return file_systemuser_proto_rawDescGZIP(), []int{4}
}

func (x *SystemUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SystemUser) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *SystemUser) GetGmail() string {
	if x != nil {
		return x.Gmail
	}
	return ""
}

func (x *SystemUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SystemUser) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *SystemUser) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *SystemUser) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UpdateSystemUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Phone string `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Gmail string `protobuf:"bytes,3,opt,name=gmail,proto3" json:"gmail,omitempty"`
	Name  string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Role  string `protobuf:"bytes,5,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *UpdateSystemUserRequest) Reset() {
	*x = UpdateSystemUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_systemuser_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSystemUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSystemUserRequest) ProtoMessage() {}

func (x *UpdateSystemUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_systemuser_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSystemUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateSystemUserRequest) Descriptor() ([]byte, []int) {
	return file_systemuser_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSystemUserRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateSystemUserRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UpdateSystemUserRequest) GetGmail() string {
	if x != nil {
		return x.Gmail
	}
	return ""
}

func (x *UpdateSystemUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateSystemUserRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

type UpdateSystemUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateSystemUserResponse) Reset() {
	*x = UpdateSystemUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_systemuser_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSystemUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSystemUserResponse) ProtoMessage() {}

func (x *UpdateSystemUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_systemuser_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSystemUserResponse.ProtoReflect.Descriptor instead.
func (*UpdateSystemUserResponse) Descriptor() ([]byte, []int) {
	return file_systemuser_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateSystemUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetListSystemUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetListSystemUserRequest) Reset() {
	*x = GetListSystemUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_systemuser_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListSystemUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListSystemUserRequest) ProtoMessage() {}

func (x *GetListSystemUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_systemuser_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListSystemUserRequest.ProtoReflect.Descriptor instead.
func (*GetListSystemUserRequest) Descriptor() ([]byte, []int) {
	return file_systemuser_proto_rawDescGZIP(), []int{7}
}

func (x *GetListSystemUserRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListSystemUserRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListSystemUserRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetListSystemUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count      int64         `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	SystemUser []*SystemUser `protobuf:"bytes,2,rep,name=SystemUser,proto3" json:"SystemUser,omitempty"`
}

func (x *GetListSystemUserResponse) Reset() {
	*x = GetListSystemUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_systemuser_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListSystemUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListSystemUserResponse) ProtoMessage() {}

func (x *GetListSystemUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_systemuser_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListSystemUserResponse.ProtoReflect.Descriptor instead.
func (*GetListSystemUserResponse) Descriptor() ([]byte, []int) {
	return file_systemuser_proto_rawDescGZIP(), []int{8}
}

func (x *GetListSystemUserResponse) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetListSystemUserResponse) GetSystemUser() []*SystemUser {
	if x != nil {
		return x.SystemUser
	}
	return nil
}

var File_systemuser_proto protoreflect.FileDescriptor

var file_systemuser_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x11, 0x0a, 0x0f, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x26, 0x0a, 0x14, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x27, 0x0a, 0x0f, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x66, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0xae, 0x01, 0x0a,
	0x0a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x6f, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x7d, 0x0a,
	0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x34, 0x0a, 0x18,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x60, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x22, 0x6b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x38, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x55, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x0a, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65,
	0x72, 0x32, 0x89, 0x04, 0x0a, 0x11, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65,
	0x72, 0x1a, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72,
	0x22, 0x00, 0x12, 0x5c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x26, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x59, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x25, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x26, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73,
	0x65, 0x72, 0x47, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x22, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x22, 0x00, 0x42, 0x17, 0x5a,
	0x15, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_systemuser_proto_rawDescOnce sync.Once
	file_systemuser_proto_rawDescData = file_systemuser_proto_rawDesc
)

func file_systemuser_proto_rawDescGZIP() []byte {
	file_systemuser_proto_rawDescOnce.Do(func() {
		file_systemuser_proto_rawDescData = protoimpl.X.CompressGZIP(file_systemuser_proto_rawDescData)
	})
	return file_systemuser_proto_rawDescData
}

var file_systemuser_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_systemuser_proto_goTypes = []interface{}{
	(*SystemUserEmpty)(nil),           // 0: user_service.SystemUserEmpty
	(*SystemUserPrimaryKey)(nil),      // 1: user_service.SystemUserPrimaryKey
	(*SystemUserGmail)(nil),           // 2: user_service.SystemUserGmail
	(*CreateSystemUser)(nil),          // 3: user_service.CreateSystemUser
	(*SystemUser)(nil),                // 4: user_service.SystemUser
	(*UpdateSystemUserRequest)(nil),   // 5: user_service.UpdateSystemUserRequest
	(*UpdateSystemUserResponse)(nil),  // 6: user_service.UpdateSystemUserResponse
	(*GetListSystemUserRequest)(nil),  // 7: user_service.GetListSystemUserRequest
	(*GetListSystemUserResponse)(nil), // 8: user_service.GetListSystemUserResponse
}
var file_systemuser_proto_depIdxs = []int32{
	4, // 0: user_service.GetListSystemUserResponse.SystemUser:type_name -> user_service.SystemUser
	3, // 1: user_service.SystemUserService.Create:input_type -> user_service.CreateSystemUser
	1, // 2: user_service.SystemUserService.GetByID:input_type -> user_service.SystemUserPrimaryKey
	7, // 3: user_service.SystemUserService.GetList:input_type -> user_service.GetListSystemUserRequest
	5, // 4: user_service.SystemUserService.Update:input_type -> user_service.UpdateSystemUserRequest
	1, // 5: user_service.SystemUserService.Delete:input_type -> user_service.SystemUserPrimaryKey
	2, // 6: user_service.SystemUserService.GetByGmail:input_type -> user_service.SystemUserGmail
	1, // 7: user_service.SystemUserService.Create:output_type -> user_service.SystemUserPrimaryKey
	4, // 8: user_service.SystemUserService.GetByID:output_type -> user_service.SystemUser
	8, // 9: user_service.SystemUserService.GetList:output_type -> user_service.GetListSystemUserResponse
	6, // 10: user_service.SystemUserService.Update:output_type -> user_service.UpdateSystemUserResponse
	0, // 11: user_service.SystemUserService.Delete:output_type -> user_service.SystemUserEmpty
	1, // 12: user_service.SystemUserService.GetByGmail:output_type -> user_service.SystemUserPrimaryKey
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_systemuser_proto_init() }
func file_systemuser_proto_init() {
	if File_systemuser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_systemuser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemUserEmpty); i {
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
		file_systemuser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemUserPrimaryKey); i {
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
		file_systemuser_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemUserGmail); i {
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
		file_systemuser_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSystemUser); i {
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
		file_systemuser_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemUser); i {
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
		file_systemuser_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSystemUserRequest); i {
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
		file_systemuser_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSystemUserResponse); i {
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
		file_systemuser_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListSystemUserRequest); i {
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
		file_systemuser_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListSystemUserResponse); i {
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
			RawDescriptor: file_systemuser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_systemuser_proto_goTypes,
		DependencyIndexes: file_systemuser_proto_depIdxs,
		MessageInfos:      file_systemuser_proto_msgTypes,
	}.Build()
	File_systemuser_proto = out.File
	file_systemuser_proto_rawDesc = nil
	file_systemuser_proto_goTypes = nil
	file_systemuser_proto_depIdxs = nil
}
