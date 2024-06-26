// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: food.proto

package gen

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

type CreateFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Category string `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Quantity uint32 `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *CreateFoodRequest) Reset() {
	*x = CreateFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFoodRequest) ProtoMessage() {}

func (x *CreateFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFoodRequest.ProtoReflect.Descriptor instead.
func (*CreateFoodRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{0}
}

func (x *CreateFoodRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateFoodRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CreateFoodRequest) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type GetAllFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllFoodRequest) Reset() {
	*x = GetAllFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllFoodRequest) ProtoMessage() {}

func (x *GetAllFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllFoodRequest.ProtoReflect.Descriptor instead.
func (*GetAllFoodRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{1}
}

type GetFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFoodRequest) Reset() {
	*x = GetFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFoodRequest) ProtoMessage() {}

func (x *GetFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFoodRequest.ProtoReflect.Descriptor instead.
func (*GetFoodRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{2}
}

func (x *GetFoodRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Quantity uint32 `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *UpdateFoodRequest) Reset() {
	*x = UpdateFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFoodRequest) ProtoMessage() {}

func (x *UpdateFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFoodRequest.ProtoReflect.Descriptor instead.
func (*UpdateFoodRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateFoodRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateFoodRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateFoodRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *UpdateFoodRequest) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type DeleteFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteFoodRequest) Reset() {
	*x = DeleteFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFoodRequest) ProtoMessage() {}

func (x *DeleteFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFoodRequest.ProtoReflect.Descriptor instead.
func (*DeleteFoodRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteFoodRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Category string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Quantity uint32 `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *FoodResponse) Reset() {
	*x = FoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodResponse) ProtoMessage() {}

func (x *FoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodResponse.ProtoReflect.Descriptor instead.
func (*FoodResponse) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{5}
}

func (x *FoodResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FoodResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FoodResponse) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *FoodResponse) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type DeleteFoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteFoodResponse) Reset() {
	*x = DeleteFoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFoodResponse) ProtoMessage() {}

func (x *DeleteFoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFoodResponse.ProtoReflect.Descriptor instead.
func (*DeleteFoodResponse) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteFoodResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_food_proto protoreflect.FileDescriptor

var file_food_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x6f,
	0x6f, 0x64, 0x22, 0x5f, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x6f, 0x6f,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6f, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x23, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x6a, 0x0a, 0x0c, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x2e, 0x0a, 0x12,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xb4, 0x02, 0x0a,
	0x0b, 0x46, 0x6f, 0x6f, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x17, 0x2e, 0x66, 0x6f, 0x6f,
	0x64, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x17, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x14, 0x2e,
	0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x17, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64,
	0x12, 0x17, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f,
	0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x66, 0x6f, 0x6f, 0x64,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_food_proto_rawDescOnce sync.Once
	file_food_proto_rawDescData = file_food_proto_rawDesc
)

func file_food_proto_rawDescGZIP() []byte {
	file_food_proto_rawDescOnce.Do(func() {
		file_food_proto_rawDescData = protoimpl.X.CompressGZIP(file_food_proto_rawDescData)
	})
	return file_food_proto_rawDescData
}

var file_food_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_food_proto_goTypes = []interface{}{
	(*CreateFoodRequest)(nil),  // 0: food.CreateFoodRequest
	(*GetAllFoodRequest)(nil),  // 1: food.GetAllFoodRequest
	(*GetFoodRequest)(nil),     // 2: food.GetFoodRequest
	(*UpdateFoodRequest)(nil),  // 3: food.UpdateFoodRequest
	(*DeleteFoodRequest)(nil),  // 4: food.DeleteFoodRequest
	(*FoodResponse)(nil),       // 5: food.FoodResponse
	(*DeleteFoodResponse)(nil), // 6: food.DeleteFoodResponse
}
var file_food_proto_depIdxs = []int32{
	1, // 0: food.FoodService.GetAllFood:input_type -> food.GetAllFoodRequest
	0, // 1: food.FoodService.CreateFood:input_type -> food.CreateFoodRequest
	2, // 2: food.FoodService.GetFood:input_type -> food.GetFoodRequest
	3, // 3: food.FoodService.UpdateFood:input_type -> food.UpdateFoodRequest
	4, // 4: food.FoodService.DeleteFood:input_type -> food.DeleteFoodRequest
	5, // 5: food.FoodService.GetAllFood:output_type -> food.FoodResponse
	5, // 6: food.FoodService.CreateFood:output_type -> food.FoodResponse
	5, // 7: food.FoodService.GetFood:output_type -> food.FoodResponse
	5, // 8: food.FoodService.UpdateFood:output_type -> food.FoodResponse
	6, // 9: food.FoodService.DeleteFood:output_type -> food.DeleteFoodResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_food_proto_init() }
func file_food_proto_init() {
	if File_food_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_food_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFoodRequest); i {
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
		file_food_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllFoodRequest); i {
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
		file_food_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFoodRequest); i {
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
		file_food_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFoodRequest); i {
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
		file_food_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFoodRequest); i {
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
		file_food_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodResponse); i {
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
		file_food_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFoodResponse); i {
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
			RawDescriptor: file_food_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_food_proto_goTypes,
		DependencyIndexes: file_food_proto_depIdxs,
		MessageInfos:      file_food_proto_msgTypes,
	}.Build()
	File_food_proto = out.File
	file_food_proto_rawDesc = nil
	file_food_proto_goTypes = nil
	file_food_proto_depIdxs = nil
}
