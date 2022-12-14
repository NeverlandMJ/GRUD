// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: grudpb/grud.proto

package grudpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetUserPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID int32 `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *GetUserPostsRequest) Reset() {
	*x = GetUserPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grudpb_grud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPostsRequest) ProtoMessage() {}

func (x *GetUserPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grudpb_grud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserPostsRequest.ProtoReflect.Descriptor instead.
func (*GetUserPostsRequest) Descriptor() ([]byte, []int) {
	return file_grudpb_grud_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserPostsRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int32  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title  string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body   string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grudpb_grud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_grudpb_grud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_grudpb_grud_proto_rawDescGZIP(), []int{1}
}

func (x *Post) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Post) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type GetPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *GetPostsResponse) Reset() {
	*x = GetPostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grudpb_grud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostsResponse) ProtoMessage() {}

func (x *GetPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grudpb_grud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostsResponse.ProtoReflect.Descriptor instead.
func (*GetPostsResponse) Descriptor() ([]byte, []int) {
	return file_grudpb_grud_proto_rawDescGZIP(), []int{2}
}

func (x *GetPostsResponse) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

type ListPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *GetPostsResponse `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	All  int32             `protobuf:"varint,2,opt,name=all,proto3" json:"all,omitempty"`
}

func (x *ListPostsResponse) Reset() {
	*x = ListPostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grudpb_grud_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsResponse) ProtoMessage() {}

func (x *ListPostsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grudpb_grud_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsResponse.ProtoReflect.Descriptor instead.
func (*ListPostsResponse) Descriptor() ([]byte, []int) {
	return file_grudpb_grud_proto_rawDescGZIP(), []int{3}
}

func (x *ListPostsResponse) GetData() *GetPostsResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ListPostsResponse) GetAll() int32 {
	if x != nil {
		return x.All
	}
	return 0
}

type GetPostByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostID int32 `protobuf:"varint,1,opt,name=postID,proto3" json:"postID,omitempty"`
}

func (x *GetPostByIDRequest) Reset() {
	*x = GetPostByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grudpb_grud_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostByIDRequest) ProtoMessage() {}

func (x *GetPostByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grudpb_grud_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostByIDRequest.ProtoReflect.Descriptor instead.
func (*GetPostByIDRequest) Descriptor() ([]byte, []int) {
	return file_grudpb_grud_proto_rawDescGZIP(), []int{4}
}

func (x *GetPostByIDRequest) GetPostID() int32 {
	if x != nil {
		return x.PostID
	}
	return 0
}

type DeletePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostID int32 `protobuf:"varint,1,opt,name=postID,proto3" json:"postID,omitempty"`
}

func (x *DeletePostRequest) Reset() {
	*x = DeletePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grudpb_grud_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostRequest) ProtoMessage() {}

func (x *DeletePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grudpb_grud_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostRequest.ProtoReflect.Descriptor instead.
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return file_grudpb_grud_proto_rawDescGZIP(), []int{5}
}

func (x *DeletePostRequest) GetPostID() int32 {
	if x != nil {
		return x.PostID
	}
	return 0
}

type UpdateTitleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostID   int32  `protobuf:"varint,1,opt,name=postID,proto3" json:"postID,omitempty"`
	NewTitle string `protobuf:"bytes,2,opt,name=newTitle,proto3" json:"newTitle,omitempty"`
}

func (x *UpdateTitleRequest) Reset() {
	*x = UpdateTitleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grudpb_grud_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTitleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTitleRequest) ProtoMessage() {}

func (x *UpdateTitleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grudpb_grud_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTitleRequest.ProtoReflect.Descriptor instead.
func (*UpdateTitleRequest) Descriptor() ([]byte, []int) {
	return file_grudpb_grud_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateTitleRequest) GetPostID() int32 {
	if x != nil {
		return x.PostID
	}
	return 0
}

func (x *UpdateTitleRequest) GetNewTitle() string {
	if x != nil {
		return x.NewTitle
	}
	return ""
}

type UpdateBodyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostID  int32  `protobuf:"varint,1,opt,name=postID,proto3" json:"postID,omitempty"`
	NewBody string `protobuf:"bytes,2,opt,name=newBody,proto3" json:"newBody,omitempty"`
}

func (x *UpdateBodyRequest) Reset() {
	*x = UpdateBodyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grudpb_grud_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBodyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBodyRequest) ProtoMessage() {}

func (x *UpdateBodyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grudpb_grud_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBodyRequest.ProtoReflect.Descriptor instead.
func (*UpdateBodyRequest) Descriptor() ([]byte, []int) {
	return file_grudpb_grud_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateBodyRequest) GetPostID() int32 {
	if x != nil {
		return x.PostID
	}
	return 0
}

func (x *UpdateBodyRequest) GetNewBody() string {
	if x != nil {
		return x.NewBody
	}
	return ""
}

type ListPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListPostsRequest) Reset() {
	*x = ListPostsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grudpb_grud_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostsRequest) ProtoMessage() {}

func (x *ListPostsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grudpb_grud_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostsRequest.ProtoReflect.Descriptor instead.
func (*ListPostsRequest) Descriptor() ([]byte, []int) {
	return file_grudpb_grud_proto_rawDescGZIP(), []int{8}
}

func (x *ListPostsRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListPostsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

var File_grudpb_grud_proto protoreflect.FileDescriptor

var file_grudpb_grud_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x72, 0x75, 0x64, 0x70, 0x62, 0x2f, 0x67, 0x72, 0x75, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x06, 0x67, 0x72, 0x75, 0x64, 0x70, 0x62, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x59, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f,
	0x64, 0x79, 0x22, 0x36, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x67, 0x72, 0x75, 0x64, 0x70, 0x62, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x53, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x67, 0x72, 0x75, 0x64, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x0a,
	0x03, 0x61, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x6c, 0x6c, 0x22,
	0x2c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x22, 0x2b, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x22, 0x48, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x65, 0x77, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x77, 0x54,
	0x69, 0x74, 0x6c, 0x65, 0x22, 0x45, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f,
	0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x42, 0x6f, 0x64, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x3c, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x32, 0x84, 0x03, 0x0a, 0x0b, 0x47, 0x72,
	0x75, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1b, 0x2e,
	0x67, 0x72, 0x75, 0x64, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x75,
	0x64, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x75, 0x64, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0c, 0x2e, 0x67, 0x72, 0x75, 0x64, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x3f, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x67, 0x72,
	0x75, 0x64, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x37,
	0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x2e,
	0x67, 0x72, 0x75, 0x64, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x75, 0x64,
	0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x19, 0x2e, 0x67, 0x72, 0x75, 0x64, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0c, 0x2e, 0x67, 0x72, 0x75, 0x64, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x40,
	0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x18, 0x2e, 0x67, 0x72,
	0x75, 0x64, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x75, 0x64, 0x70, 0x62, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e,
	0x65, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x6e, 0x64, 0x4d, 0x4a, 0x2f, 0x47, 0x52, 0x55, 0x44, 0x2f,
	0x67, 0x72, 0x75, 0x64, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x75,
	0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grudpb_grud_proto_rawDescOnce sync.Once
	file_grudpb_grud_proto_rawDescData = file_grudpb_grud_proto_rawDesc
)

func file_grudpb_grud_proto_rawDescGZIP() []byte {
	file_grudpb_grud_proto_rawDescOnce.Do(func() {
		file_grudpb_grud_proto_rawDescData = protoimpl.X.CompressGZIP(file_grudpb_grud_proto_rawDescData)
	})
	return file_grudpb_grud_proto_rawDescData
}

var file_grudpb_grud_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_grudpb_grud_proto_goTypes = []interface{}{
	(*GetUserPostsRequest)(nil), // 0: grudpb.GetUserPostsRequest
	(*Post)(nil),                // 1: grudpb.Post
	(*GetPostsResponse)(nil),    // 2: grudpb.GetPostsResponse
	(*ListPostsResponse)(nil),   // 3: grudpb.ListPostsResponse
	(*GetPostByIDRequest)(nil),  // 4: grudpb.GetPostByIDRequest
	(*DeletePostRequest)(nil),   // 5: grudpb.DeletePostRequest
	(*UpdateTitleRequest)(nil),  // 6: grudpb.UpdateTitleRequest
	(*UpdateBodyRequest)(nil),   // 7: grudpb.UpdateBodyRequest
	(*ListPostsRequest)(nil),    // 8: grudpb.ListPostsRequest
	(*emptypb.Empty)(nil),       // 9: google.protobuf.Empty
}
var file_grudpb_grud_proto_depIdxs = []int32{
	1, // 0: grudpb.GetPostsResponse.posts:type_name -> grudpb.Post
	2, // 1: grudpb.ListPostsResponse.data:type_name -> grudpb.GetPostsResponse
	0, // 2: grudpb.GrudService.GetPostsByUserID:input_type -> grudpb.GetUserPostsRequest
	4, // 3: grudpb.GrudService.GetPostByID:input_type -> grudpb.GetPostByIDRequest
	5, // 4: grudpb.GrudService.DeletePost:input_type -> grudpb.DeletePostRequest
	6, // 5: grudpb.GrudService.UpdateTitle:input_type -> grudpb.UpdateTitleRequest
	7, // 6: grudpb.GrudService.UpdateBody:input_type -> grudpb.UpdateBodyRequest
	8, // 7: grudpb.GrudService.ListPosts:input_type -> grudpb.ListPostsRequest
	2, // 8: grudpb.GrudService.GetPostsByUserID:output_type -> grudpb.GetPostsResponse
	1, // 9: grudpb.GrudService.GetPostByID:output_type -> grudpb.Post
	9, // 10: grudpb.GrudService.DeletePost:output_type -> google.protobuf.Empty
	1, // 11: grudpb.GrudService.UpdateTitle:output_type -> grudpb.Post
	1, // 12: grudpb.GrudService.UpdateBody:output_type -> grudpb.Post
	3, // 13: grudpb.GrudService.ListPosts:output_type -> grudpb.ListPostsResponse
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grudpb_grud_proto_init() }
func file_grudpb_grud_proto_init() {
	if File_grudpb_grud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grudpb_grud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserPostsRequest); i {
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
		file_grudpb_grud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_grudpb_grud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostsResponse); i {
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
		file_grudpb_grud_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostsResponse); i {
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
		file_grudpb_grud_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostByIDRequest); i {
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
		file_grudpb_grud_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostRequest); i {
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
		file_grudpb_grud_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTitleRequest); i {
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
		file_grudpb_grud_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBodyRequest); i {
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
		file_grudpb_grud_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostsRequest); i {
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
			RawDescriptor: file_grudpb_grud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grudpb_grud_proto_goTypes,
		DependencyIndexes: file_grudpb_grud_proto_depIdxs,
		MessageInfos:      file_grudpb_grud_proto_msgTypes,
	}.Build()
	File_grudpb_grud_proto = out.File
	file_grudpb_grud_proto_rawDesc = nil
	file_grudpb_grud_proto_goTypes = nil
	file_grudpb_grud_proto_depIdxs = nil
}
