// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: employee/v1/employee.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	v1 "northwind/api/typespb/v1"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeId *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_v1_employee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_v1_employee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_employee_v1_employee_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteRequest) GetEmployeeId() *wrapperspb.Int64Value {
	if x != nil {
		return x.EmployeeId
	}
	return nil
}

type EmployeeByEmployeeIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmployeeId *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
}

func (x *EmployeeByEmployeeIDRequest) Reset() {
	*x = EmployeeByEmployeeIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_v1_employee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeeByEmployeeIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeeByEmployeeIDRequest) ProtoMessage() {}

func (x *EmployeeByEmployeeIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_v1_employee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeeByEmployeeIDRequest.ProtoReflect.Descriptor instead.
func (*EmployeeByEmployeeIDRequest) Descriptor() ([]byte, []int) {
	return file_employee_v1_employee_proto_rawDescGZIP(), []int{1}
}

func (x *EmployeeByEmployeeIDRequest) GetEmployeeId() *wrapperspb.Int64Value {
	if x != nil {
		return x.EmployeeId
	}
	return nil
}

type InsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BirthDate  *timestamppb.Timestamp  `protobuf:"bytes,1,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	EmployeeId *wrapperspb.Int64Value  `protobuf:"bytes,2,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
	FirstName  *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName   *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Notes      *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=notes,proto3" json:"notes,omitempty"`
	Photo      *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=photo,proto3" json:"photo,omitempty"`
}

func (x *InsertRequest) Reset() {
	*x = InsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_v1_employee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertRequest) ProtoMessage() {}

func (x *InsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_v1_employee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertRequest.ProtoReflect.Descriptor instead.
func (*InsertRequest) Descriptor() ([]byte, []int) {
	return file_employee_v1_employee_proto_rawDescGZIP(), []int{2}
}

func (x *InsertRequest) GetBirthDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BirthDate
	}
	return nil
}

func (x *InsertRequest) GetEmployeeId() *wrapperspb.Int64Value {
	if x != nil {
		return x.EmployeeId
	}
	return nil
}

func (x *InsertRequest) GetFirstName() *wrapperspb.StringValue {
	if x != nil {
		return x.FirstName
	}
	return nil
}

func (x *InsertRequest) GetLastName() *wrapperspb.StringValue {
	if x != nil {
		return x.LastName
	}
	return nil
}

func (x *InsertRequest) GetNotes() *wrapperspb.StringValue {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *InsertRequest) GetPhoto() *wrapperspb.StringValue {
	if x != nil {
		return x.Photo
	}
	return nil
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BirthDate *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	// Output only.
	EmployeeId *wrapperspb.Int64Value  `protobuf:"bytes,2,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
	FirstName  *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName   *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Notes      *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=notes,proto3" json:"notes,omitempty"`
	Photo      *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=photo,proto3" json:"photo,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_v1_employee_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_v1_employee_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_employee_v1_employee_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateRequest) GetBirthDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BirthDate
	}
	return nil
}

func (x *UpdateRequest) GetEmployeeId() *wrapperspb.Int64Value {
	if x != nil {
		return x.EmployeeId
	}
	return nil
}

func (x *UpdateRequest) GetFirstName() *wrapperspb.StringValue {
	if x != nil {
		return x.FirstName
	}
	return nil
}

func (x *UpdateRequest) GetLastName() *wrapperspb.StringValue {
	if x != nil {
		return x.LastName
	}
	return nil
}

func (x *UpdateRequest) GetNotes() *wrapperspb.StringValue {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *UpdateRequest) GetPhoto() *wrapperspb.StringValue {
	if x != nil {
		return x.Photo
	}
	return nil
}

type UpsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BirthDate  *timestamppb.Timestamp  `protobuf:"bytes,1,opt,name=birth_date,json=birthDate,proto3" json:"birth_date,omitempty"`
	EmployeeId *wrapperspb.Int64Value  `protobuf:"bytes,2,opt,name=employee_id,json=employeeId,proto3" json:"employee_id,omitempty"`
	FirstName  *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName   *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Notes      *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=notes,proto3" json:"notes,omitempty"`
	Photo      *wrapperspb.StringValue `protobuf:"bytes,6,opt,name=photo,proto3" json:"photo,omitempty"`
}

func (x *UpsertRequest) Reset() {
	*x = UpsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_employee_v1_employee_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertRequest) ProtoMessage() {}

func (x *UpsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_v1_employee_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertRequest.ProtoReflect.Descriptor instead.
func (*UpsertRequest) Descriptor() ([]byte, []int) {
	return file_employee_v1_employee_proto_rawDescGZIP(), []int{4}
}

func (x *UpsertRequest) GetBirthDate() *timestamppb.Timestamp {
	if x != nil {
		return x.BirthDate
	}
	return nil
}

func (x *UpsertRequest) GetEmployeeId() *wrapperspb.Int64Value {
	if x != nil {
		return x.EmployeeId
	}
	return nil
}

func (x *UpsertRequest) GetFirstName() *wrapperspb.StringValue {
	if x != nil {
		return x.FirstName
	}
	return nil
}

func (x *UpsertRequest) GetLastName() *wrapperspb.StringValue {
	if x != nil {
		return x.LastName
	}
	return nil
}

func (x *UpsertRequest) GetNotes() *wrapperspb.StringValue {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *UpsertRequest) GetPhoto() *wrapperspb.StringValue {
	if x != nil {
		return x.Photo
	}
	return nil
}

var File_employee_v1_employee_proto protoreflect.FileDescriptor

var file_employee_v1_employee_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x74, 0x79, 0x70, 0x65, 0x73, 0x70, 0x62, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4d, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3c, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x22, 0x5b,
	0x0a, 0x1b, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x42, 0x79, 0x45, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a,
	0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x22, 0xe8, 0x02, 0x0a, 0x0d,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x62,
	0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32,
	0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x74,
	0x65, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x02, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49,
	0x64, 0x12, 0x3b, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39,
	0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x6e, 0x6f, 0x74,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x32, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x22, 0xe8, 0x02, 0x0a, 0x0d, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x62, 0x69, 0x72, 0x74, 0x68, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x62, 0x69, 0x72, 0x74, 0x68, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3c,
	0x0a, 0x0b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0a, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x64, 0x12, 0x3b, 0x0a, 0x0a,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x09, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x32, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x32, 0x89, 0x04, 0x0a,
	0x0f, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x60, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x2a, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70, 0x6c,
	0x6f, 0x79, 0x65, 0x65, 0x2f, 0x7b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x7a, 0x0a, 0x14, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x42, 0x79,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x44, 0x12, 0x28, 0x2e, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x42, 0x79, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x70, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65,
	0x2f, 0x7b, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x55,
	0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x1a, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x17, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70,
	0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x63, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x3a, 0x01, 0x2a, 0x1a, 0x1a,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2f, 0x7b, 0x65, 0x6d,
	0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x5c, 0x0a, 0x06, 0x55, 0x70,
	0x73, 0x65, 0x72, 0x74, 0x12, 0x1a, 0x2e, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18,
	0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65,
	0x65, 0x2f, 0x75, 0x70, 0x73, 0x65, 0x72, 0x74, 0x42, 0x1b, 0x5a, 0x19, 0x6e, 0x6f, 0x72, 0x74,
	0x68, 0x77, 0x69, 0x6e, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6d, 0x70, 0x6c, 0x6f, 0x79,
	0x65, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_employee_v1_employee_proto_rawDescOnce sync.Once
	file_employee_v1_employee_proto_rawDescData = file_employee_v1_employee_proto_rawDesc
)

func file_employee_v1_employee_proto_rawDescGZIP() []byte {
	file_employee_v1_employee_proto_rawDescOnce.Do(func() {
		file_employee_v1_employee_proto_rawDescData = protoimpl.X.CompressGZIP(file_employee_v1_employee_proto_rawDescData)
	})
	return file_employee_v1_employee_proto_rawDescData
}

var file_employee_v1_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_employee_v1_employee_proto_goTypes = []interface{}{
	(*DeleteRequest)(nil),               // 0: employee.v1.DeleteRequest
	(*EmployeeByEmployeeIDRequest)(nil), // 1: employee.v1.EmployeeByEmployeeIDRequest
	(*InsertRequest)(nil),               // 2: employee.v1.InsertRequest
	(*UpdateRequest)(nil),               // 3: employee.v1.UpdateRequest
	(*UpsertRequest)(nil),               // 4: employee.v1.UpsertRequest
	(*wrapperspb.Int64Value)(nil),       // 5: google.protobuf.Int64Value
	(*timestamppb.Timestamp)(nil),       // 6: google.protobuf.Timestamp
	(*wrapperspb.StringValue)(nil),      // 7: google.protobuf.StringValue
	(*emptypb.Empty)(nil),               // 8: google.protobuf.Empty
	(*v1.Employee)(nil),                 // 9: typespb.v1.Employee
}
var file_employee_v1_employee_proto_depIdxs = []int32{
	5,  // 0: employee.v1.DeleteRequest.employee_id:type_name -> google.protobuf.Int64Value
	5,  // 1: employee.v1.EmployeeByEmployeeIDRequest.employee_id:type_name -> google.protobuf.Int64Value
	6,  // 2: employee.v1.InsertRequest.birth_date:type_name -> google.protobuf.Timestamp
	5,  // 3: employee.v1.InsertRequest.employee_id:type_name -> google.protobuf.Int64Value
	7,  // 4: employee.v1.InsertRequest.first_name:type_name -> google.protobuf.StringValue
	7,  // 5: employee.v1.InsertRequest.last_name:type_name -> google.protobuf.StringValue
	7,  // 6: employee.v1.InsertRequest.notes:type_name -> google.protobuf.StringValue
	7,  // 7: employee.v1.InsertRequest.photo:type_name -> google.protobuf.StringValue
	6,  // 8: employee.v1.UpdateRequest.birth_date:type_name -> google.protobuf.Timestamp
	5,  // 9: employee.v1.UpdateRequest.employee_id:type_name -> google.protobuf.Int64Value
	7,  // 10: employee.v1.UpdateRequest.first_name:type_name -> google.protobuf.StringValue
	7,  // 11: employee.v1.UpdateRequest.last_name:type_name -> google.protobuf.StringValue
	7,  // 12: employee.v1.UpdateRequest.notes:type_name -> google.protobuf.StringValue
	7,  // 13: employee.v1.UpdateRequest.photo:type_name -> google.protobuf.StringValue
	6,  // 14: employee.v1.UpsertRequest.birth_date:type_name -> google.protobuf.Timestamp
	5,  // 15: employee.v1.UpsertRequest.employee_id:type_name -> google.protobuf.Int64Value
	7,  // 16: employee.v1.UpsertRequest.first_name:type_name -> google.protobuf.StringValue
	7,  // 17: employee.v1.UpsertRequest.last_name:type_name -> google.protobuf.StringValue
	7,  // 18: employee.v1.UpsertRequest.notes:type_name -> google.protobuf.StringValue
	7,  // 19: employee.v1.UpsertRequest.photo:type_name -> google.protobuf.StringValue
	0,  // 20: employee.v1.EmployeeService.Delete:input_type -> employee.v1.DeleteRequest
	1,  // 21: employee.v1.EmployeeService.EmployeeByEmployeeID:input_type -> employee.v1.EmployeeByEmployeeIDRequest
	2,  // 22: employee.v1.EmployeeService.Insert:input_type -> employee.v1.InsertRequest
	3,  // 23: employee.v1.EmployeeService.Update:input_type -> employee.v1.UpdateRequest
	4,  // 24: employee.v1.EmployeeService.Upsert:input_type -> employee.v1.UpsertRequest
	8,  // 25: employee.v1.EmployeeService.Delete:output_type -> google.protobuf.Empty
	9,  // 26: employee.v1.EmployeeService.EmployeeByEmployeeID:output_type -> typespb.v1.Employee
	8,  // 27: employee.v1.EmployeeService.Insert:output_type -> google.protobuf.Empty
	8,  // 28: employee.v1.EmployeeService.Update:output_type -> google.protobuf.Empty
	8,  // 29: employee.v1.EmployeeService.Upsert:output_type -> google.protobuf.Empty
	25, // [25:30] is the sub-list for method output_type
	20, // [20:25] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_employee_v1_employee_proto_init() }
func file_employee_v1_employee_proto_init() {
	if File_employee_v1_employee_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_employee_v1_employee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_employee_v1_employee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeeByEmployeeIDRequest); i {
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
		file_employee_v1_employee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertRequest); i {
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
		file_employee_v1_employee_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
		file_employee_v1_employee_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertRequest); i {
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
			RawDescriptor: file_employee_v1_employee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_employee_v1_employee_proto_goTypes,
		DependencyIndexes: file_employee_v1_employee_proto_depIdxs,
		MessageInfos:      file_employee_v1_employee_proto_msgTypes,
	}.Build()
	File_employee_v1_employee_proto = out.File
	file_employee_v1_employee_proto_rawDesc = nil
	file_employee_v1_employee_proto_goTypes = nil
	file_employee_v1_employee_proto_depIdxs = nil
}
