// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        (unknown)
// source: customer/v1/customer.proto

package customer_v1alpha0

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

type CustomerBranch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   float64 `protobuf:"fixed64,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CustomerBranch) Reset() {
	*x = CustomerBranch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerBranch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerBranch) ProtoMessage() {}

func (x *CustomerBranch) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerBranch.ProtoReflect.Descriptor instead.
func (*CustomerBranch) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerBranch) GetId() float64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CustomerBranch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetOrdersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *GetOrdersRequest) Reset() {
	*x = GetOrdersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersRequest) ProtoMessage() {}

func (x *GetOrdersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersRequest.ProtoReflect.Descriptor instead.
func (*GetOrdersRequest) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{1}
}

func (x *GetOrdersRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type GetOrdersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Orders []*GetOrdersResponse_Order `protobuf:"bytes,1,rep,name=orders,proto3" json:"orders,omitempty"`
}

func (x *GetOrdersResponse) Reset() {
	*x = GetOrdersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersResponse) ProtoMessage() {}

func (x *GetOrdersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersResponse.ProtoReflect.Descriptor instead.
func (*GetOrdersResponse) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{2}
}

func (x *GetOrdersResponse) GetOrders() []*GetOrdersResponse_Order {
	if x != nil {
		return x.Orders
	}
	return nil
}

type GetCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CustomerId float64 `protobuf:"fixed64,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *GetCustomerRequest) Reset() {
	*x = GetCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerRequest) ProtoMessage() {}

func (x *GetCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerRequest.ProtoReflect.Descriptor instead.
func (*GetCustomerRequest) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{3}
}

func (x *GetCustomerRequest) GetCustomerId() float64 {
	if x != nil {
		return x.CustomerId
	}
	return 0
}

type GetCustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCustomerResponse) Reset() {
	*x = GetCustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCustomerResponse) ProtoMessage() {}

func (x *GetCustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCustomerResponse.ProtoReflect.Descriptor instead.
func (*GetCustomerResponse) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{4}
}

type GetContactByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetContactByIdRequest) Reset() {
	*x = GetContactByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContactByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactByIdRequest) ProtoMessage() {}

func (x *GetContactByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactByIdRequest.ProtoReflect.Descriptor instead.
func (*GetContactByIdRequest) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{5}
}

func (x *GetContactByIdRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetContactByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Contact  *GetContactByIdResponse_Contact `protobuf:"bytes,1,opt,name=contact,proto3" json:"contact,omitempty"`
	Branches []*CustomerBranch               `protobuf:"bytes,2,rep,name=branches,proto3" json:"branches,omitempty"`
}

func (x *GetContactByIdResponse) Reset() {
	*x = GetContactByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContactByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactByIdResponse) ProtoMessage() {}

func (x *GetContactByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactByIdResponse.ProtoReflect.Descriptor instead.
func (*GetContactByIdResponse) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{6}
}

func (x *GetContactByIdResponse) GetContact() *GetContactByIdResponse_Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

func (x *GetContactByIdResponse) GetBranches() []*CustomerBranch {
	if x != nil {
		return x.Branches
	}
	return nil
}

type GetOrdersResponse_Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetOrdersResponse_Order) Reset() {
	*x = GetOrdersResponse_Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersResponse_Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersResponse_Order) ProtoMessage() {}

func (x *GetOrdersResponse_Order) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersResponse_Order.ProtoReflect.Descriptor instead.
func (*GetOrdersResponse_Order) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GetOrdersResponse_Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetOrdersResponse_Order) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetContactByIdResponse_Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *GetContactByIdResponse_Contact) Reset() {
	*x = GetContactByIdResponse_Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_customer_v1_customer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetContactByIdResponse_Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetContactByIdResponse_Contact) ProtoMessage() {}

func (x *GetContactByIdResponse_Contact) ProtoReflect() protoreflect.Message {
	mi := &file_customer_v1_customer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetContactByIdResponse_Contact.ProtoReflect.Descriptor instead.
func (*GetContactByIdResponse_Contact) Descriptor() ([]byte, []int) {
	return file_customer_v1_customer_proto_rawDescGZIP(), []int{6, 0}
}

func (x *GetContactByIdResponse_Contact) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetContactByIdResponse_Contact) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetContactByIdResponse_Contact) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GetContactByIdResponse_Contact) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetContactByIdResponse_Contact) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_customer_v1_customer_proto protoreflect.FileDescriptor

var file_customer_v1_customer_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x22, 0x34, 0x0a, 0x0e, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x33, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x06, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x06, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x2f, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x35, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x9c, 0x02, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x42, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x52, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x65, 0x73, 0x1a, 0x81, 0x01, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x32,
	0x8a, 0x02, 0x0a, 0x0f, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x1d, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x50, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1f,
	0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x59, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x22, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x50, 0x5a, 0x4e,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x74, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x73, 0x2d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73,
	0x5f, 0x70, 0x72, 0x6f, 0x70, 0x68, 0x65, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x30, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_customer_v1_customer_proto_rawDescOnce sync.Once
	file_customer_v1_customer_proto_rawDescData = file_customer_v1_customer_proto_rawDesc
)

func file_customer_v1_customer_proto_rawDescGZIP() []byte {
	file_customer_v1_customer_proto_rawDescOnce.Do(func() {
		file_customer_v1_customer_proto_rawDescData = protoimpl.X.CompressGZIP(file_customer_v1_customer_proto_rawDescData)
	})
	return file_customer_v1_customer_proto_rawDescData
}

var file_customer_v1_customer_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_customer_v1_customer_proto_goTypes = []interface{}{
	(*CustomerBranch)(nil),                 // 0: customer.v1.CustomerBranch
	(*GetOrdersRequest)(nil),               // 1: customer.v1.GetOrdersRequest
	(*GetOrdersResponse)(nil),              // 2: customer.v1.GetOrdersResponse
	(*GetCustomerRequest)(nil),             // 3: customer.v1.GetCustomerRequest
	(*GetCustomerResponse)(nil),            // 4: customer.v1.GetCustomerResponse
	(*GetContactByIdRequest)(nil),          // 5: customer.v1.GetContactByIdRequest
	(*GetContactByIdResponse)(nil),         // 6: customer.v1.GetContactByIdResponse
	(*GetOrdersResponse_Order)(nil),        // 7: customer.v1.GetOrdersResponse.Order
	(*GetContactByIdResponse_Contact)(nil), // 8: customer.v1.GetContactByIdResponse.Contact
}
var file_customer_v1_customer_proto_depIdxs = []int32{
	7, // 0: customer.v1.GetOrdersResponse.orders:type_name -> customer.v1.GetOrdersResponse.Order
	8, // 1: customer.v1.GetContactByIdResponse.contact:type_name -> customer.v1.GetContactByIdResponse.Contact
	0, // 2: customer.v1.GetContactByIdResponse.branches:type_name -> customer.v1.CustomerBranch
	1, // 3: customer.v1.CustomerService.GetOrders:input_type -> customer.v1.GetOrdersRequest
	3, // 4: customer.v1.CustomerService.GetCustomer:input_type -> customer.v1.GetCustomerRequest
	5, // 5: customer.v1.CustomerService.GetContactById:input_type -> customer.v1.GetContactByIdRequest
	2, // 6: customer.v1.CustomerService.GetOrders:output_type -> customer.v1.GetOrdersResponse
	4, // 7: customer.v1.CustomerService.GetCustomer:output_type -> customer.v1.GetCustomerResponse
	6, // 8: customer.v1.CustomerService.GetContactById:output_type -> customer.v1.GetContactByIdResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_customer_v1_customer_proto_init() }
func file_customer_v1_customer_proto_init() {
	if File_customer_v1_customer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_customer_v1_customer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerBranch); i {
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
		file_customer_v1_customer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrdersRequest); i {
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
		file_customer_v1_customer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrdersResponse); i {
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
		file_customer_v1_customer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerRequest); i {
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
		file_customer_v1_customer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCustomerResponse); i {
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
		file_customer_v1_customer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContactByIdRequest); i {
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
		file_customer_v1_customer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContactByIdResponse); i {
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
		file_customer_v1_customer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrdersResponse_Order); i {
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
		file_customer_v1_customer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetContactByIdResponse_Contact); i {
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
			RawDescriptor: file_customer_v1_customer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_customer_v1_customer_proto_goTypes,
		DependencyIndexes: file_customer_v1_customer_proto_depIdxs,
		MessageInfos:      file_customer_v1_customer_proto_msgTypes,
	}.Build()
	File_customer_v1_customer_proto = out.File
	file_customer_v1_customer_proto_rawDesc = nil
	file_customer_v1_customer_proto_goTypes = nil
	file_customer_v1_customer_proto_depIdxs = nil
}
