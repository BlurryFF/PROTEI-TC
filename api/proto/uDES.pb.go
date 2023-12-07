// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: api/proto/uDES.proto

package proto

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

type Employee struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Email       string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	WorkPhone   string `protobuf:"bytes,4,opt,name=work_phone,json=workPhone,proto3" json:"work_phone,omitempty"`
}

func (x *Employee) Reset() {
	*x = Employee{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_uDES_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Employee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Employee) ProtoMessage() {}

func (x *Employee) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_uDES_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Employee.ProtoReflect.Descriptor instead.
func (*Employee) Descriptor() ([]byte, []int) {
	return file_api_proto_uDES_proto_rawDescGZIP(), []int{0}
}

func (x *Employee) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Employee) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Employee) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Employee) GetWorkPhone() string {
	if x != nil {
		return x.WorkPhone
	}
	return ""
}

type Absence struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PersonId    int32  `protobuf:"varint,2,opt,name=personId,proto3" json:"personId,omitempty"`
	CreatedDate string `protobuf:"bytes,3,opt,name=createdDate,proto3" json:"createdDate,omitempty"`
	DateFrom    string `protobuf:"bytes,4,opt,name=dateFrom,proto3" json:"dateFrom,omitempty"`
	DateTo      string `protobuf:"bytes,5,opt,name=dateTo,proto3" json:"dateTo,omitempty"`
	ReasonId    int32  `protobuf:"varint,6,opt,name=reasonId,proto3" json:"reasonId,omitempty"`
	DisplayName string `protobuf:"bytes,7,opt,name=displayName,proto3" json:"displayName,omitempty"`
}

func (x *Absence) Reset() {
	*x = Absence{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_uDES_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Absence) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Absence) ProtoMessage() {}

func (x *Absence) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_uDES_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Absence.ProtoReflect.Descriptor instead.
func (*Absence) Descriptor() ([]byte, []int) {
	return file_api_proto_uDES_proto_rawDescGZIP(), []int{1}
}

func (x *Absence) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Absence) GetPersonId() int32 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *Absence) GetCreatedDate() string {
	if x != nil {
		return x.CreatedDate
	}
	return ""
}

func (x *Absence) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *Absence) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

func (x *Absence) GetReasonId() int32 {
	if x != nil {
		return x.ReasonId
	}
	return 0
}

func (x *Absence) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type AbsenceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonId    int32  `protobuf:"varint,1,opt,name=personId,proto3" json:"personId,omitempty"`
	DisplayName string `protobuf:"bytes,2,opt,name=displayName,proto3" json:"displayName,omitempty"`
}

func (x *AbsenceData) Reset() {
	*x = AbsenceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_uDES_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbsenceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbsenceData) ProtoMessage() {}

func (x *AbsenceData) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_uDES_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbsenceData.ProtoReflect.Descriptor instead.
func (*AbsenceData) Descriptor() ([]byte, []int) {
	return file_api_proto_uDES_proto_rawDescGZIP(), []int{2}
}

func (x *AbsenceData) GetPersonId() int32 {
	if x != nil {
		return x.PersonId
	}
	return 0
}

func (x *AbsenceData) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type GetEmployeeParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids       []int32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
	Email     string  `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name      string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	WorkPhone string  `protobuf:"bytes,4,opt,name=workPhone,proto3" json:"workPhone,omitempty"`
	DateFrom  string  `protobuf:"bytes,5,opt,name=dateFrom,proto3" json:"dateFrom,omitempty"`
	DateTo    string  `protobuf:"bytes,6,opt,name=dateTo,proto3" json:"dateTo,omitempty"`
}

func (x *GetEmployeeParams) Reset() {
	*x = GetEmployeeParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_uDES_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmployeeParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmployeeParams) ProtoMessage() {}

func (x *GetEmployeeParams) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_uDES_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmployeeParams.ProtoReflect.Descriptor instead.
func (*GetEmployeeParams) Descriptor() ([]byte, []int) {
	return file_api_proto_uDES_proto_rawDescGZIP(), []int{3}
}

func (x *GetEmployeeParams) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *GetEmployeeParams) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetEmployeeParams) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetEmployeeParams) GetWorkPhone() string {
	if x != nil {
		return x.WorkPhone
	}
	return ""
}

func (x *GetEmployeeParams) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *GetEmployeeParams) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

type GetAbsencesParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonIds []int32 `protobuf:"varint,1,rep,packed,name=personIds,proto3" json:"personIds,omitempty"`
	DateFrom  string  `protobuf:"bytes,2,opt,name=dateFrom,proto3" json:"dateFrom,omitempty"`
	DateTo    string  `protobuf:"bytes,3,opt,name=dateTo,proto3" json:"dateTo,omitempty"`
}

func (x *GetAbsencesParams) Reset() {
	*x = GetAbsencesParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_uDES_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAbsencesParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAbsencesParams) ProtoMessage() {}

func (x *GetAbsencesParams) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_uDES_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAbsencesParams.ProtoReflect.Descriptor instead.
func (*GetAbsencesParams) Descriptor() ([]byte, []int) {
	return file_api_proto_uDES_proto_rawDescGZIP(), []int{4}
}

func (x *GetAbsencesParams) GetPersonIds() []int32 {
	if x != nil {
		return x.PersonIds
	}
	return nil
}

func (x *GetAbsencesParams) GetDateFrom() string {
	if x != nil {
		return x.DateFrom
	}
	return ""
}

func (x *GetAbsencesParams) GetDateTo() string {
	if x != nil {
		return x.DateTo
	}
	return ""
}

type AbsencesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string     `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data   []*Absence `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *AbsencesList) Reset() {
	*x = AbsencesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_uDES_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AbsencesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AbsencesList) ProtoMessage() {}

func (x *AbsencesList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_uDES_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AbsencesList.ProtoReflect.Descriptor instead.
func (*AbsencesList) Descriptor() ([]byte, []int) {
	return file_api_proto_uDES_proto_rawDescGZIP(), []int{5}
}

func (x *AbsencesList) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *AbsencesList) GetData() []*Absence {
	if x != nil {
		return x.Data
	}
	return nil
}

type EmployeesList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string      `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data   []*Employee `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *EmployeesList) Reset() {
	*x = EmployeesList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_uDES_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmployeesList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmployeesList) ProtoMessage() {}

func (x *EmployeesList) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_uDES_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmployeesList.ProtoReflect.Descriptor instead.
func (*EmployeesList) Descriptor() ([]byte, []int) {
	return file_api_proto_uDES_proto_rawDescGZIP(), []int{6}
}

func (x *EmployeesList) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *EmployeesList) GetData() []*Employee {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_api_proto_uDES_proto protoreflect.FileDescriptor

var file_api_proto_uDES_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x44, 0x45, 0x53,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x72, 0x0a, 0x08, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x22,
	0xc9, 0x01, 0x0a, 0x07, 0x41, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a, 0x0b, 0x41,
	0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f,
	0x72, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77,
	0x6f, 0x72, 0x6b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x22, 0x65, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x41, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x22, 0x48, 0x0a, 0x0c, 0x41, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4a, 0x0a,
	0x0d, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x89, 0x01, 0x0a, 0x0e, 0x55, 0x73,
	0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x41, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x62, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x73, 0x1a, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x62, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x45,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a,
	0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_uDES_proto_rawDescOnce sync.Once
	file_api_proto_uDES_proto_rawDescData = file_api_proto_uDES_proto_rawDesc
)

func file_api_proto_uDES_proto_rawDescGZIP() []byte {
	file_api_proto_uDES_proto_rawDescOnce.Do(func() {
		file_api_proto_uDES_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_uDES_proto_rawDescData)
	})
	return file_api_proto_uDES_proto_rawDescData
}

var file_api_proto_uDES_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_uDES_proto_goTypes = []interface{}{
	(*Employee)(nil),          // 0: api.Employee
	(*Absence)(nil),           // 1: api.Absence
	(*AbsenceData)(nil),       // 2: api.AbsenceData
	(*GetEmployeeParams)(nil), // 3: api.GetEmployeeParams
	(*GetAbsencesParams)(nil), // 4: api.GetAbsencesParams
	(*AbsencesList)(nil),      // 5: api.AbsencesList
	(*EmployeesList)(nil),     // 6: api.EmployeesList
}
var file_api_proto_uDES_proto_depIdxs = []int32{
	1, // 0: api.AbsencesList.data:type_name -> api.Absence
	0, // 1: api.EmployeesList.data:type_name -> api.Employee
	4, // 2: api.UserManagement.GetAbsences:input_type -> api.GetAbsencesParams
	3, // 3: api.UserManagement.GetEmployee:input_type -> api.GetEmployeeParams
	5, // 4: api.UserManagement.GetAbsences:output_type -> api.AbsencesList
	6, // 5: api.UserManagement.GetEmployee:output_type -> api.EmployeesList
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_uDES_proto_init() }
func file_api_proto_uDES_proto_init() {
	if File_api_proto_uDES_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_uDES_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Employee); i {
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
		file_api_proto_uDES_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Absence); i {
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
		file_api_proto_uDES_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbsenceData); i {
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
		file_api_proto_uDES_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmployeeParams); i {
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
		file_api_proto_uDES_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAbsencesParams); i {
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
		file_api_proto_uDES_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AbsencesList); i {
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
		file_api_proto_uDES_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmployeesList); i {
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
			RawDescriptor: file_api_proto_uDES_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_uDES_proto_goTypes,
		DependencyIndexes: file_api_proto_uDES_proto_depIdxs,
		MessageInfos:      file_api_proto_uDES_proto_msgTypes,
	}.Build()
	File_api_proto_uDES_proto = out.File
	file_api_proto_uDES_proto_rawDesc = nil
	file_api_proto_uDES_proto_goTypes = nil
	file_api_proto_uDES_proto_depIdxs = nil
}