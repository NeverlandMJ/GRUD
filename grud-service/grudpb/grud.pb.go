// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: grudpb/grud.proto

package grudpb

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

type GetUserPostsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int32 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
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

func (x *GetUserPostsRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId int32  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title  string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body   string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Data) Reset() {
	*x = Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grudpb_grud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_grudpb_grud_proto_rawDescGZIP(), []int{1}
}

func (x *Data) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Data) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Data) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Data) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type GetUserPostsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datas []*Data `protobuf:"bytes,1,rep,name=datas,proto3" json:"datas,omitempty"`
}

func (x *GetUserPostsResponse) Reset() {
	*x = GetUserPostsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grudpb_grud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserPostsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPostsResponse) ProtoMessage() {}

func (x *GetUserPostsResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetUserPostsResponse.ProtoReflect.Descriptor instead.
func (*GetUserPostsResponse) Descriptor() ([]byte, []int) {
	return file_grudpb_grud_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserPostsResponse) GetDatas() []*Data {
	if x != nil {
		return x.Datas
	}
	return nil
}

type GetPostByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId int32 `protobuf:"varint,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *GetPostByIDRequest) Reset() {
	*x = GetPostByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grudpb_grud_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostByIDRequest) ProtoMessage() {}

func (x *GetPostByIDRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetPostByIDRequest.ProtoReflect.Descriptor instead.
func (*GetPostByIDRequest) Descriptor() ([]byte, []int) {
	return file_grudpb_grud_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostByIDRequest) GetPostId() int32 {
	if x != nil {
		return x.PostId
	}
	return 0
}

var File_grudpb_grud_proto protoreflect.FileDescriptor

var file_grudpb_grud_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x72, 0x75, 0x64, 0x70, 0x62, 0x2f, 0x67, 0x72, 0x75, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x70, 0x62, 0x22, 0x2e,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x59,
	0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x3d, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x25, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x61, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x05, 0x64, 0x61, 0x74, 0x61, 0x73, 0x22, 0x2d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x32, 0xa1, 0x01, 0x0a, 0x0b, 0x47, 0x72, 0x75, 0x64,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1e, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1d, 0x2e, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x70, 0x62, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x42, 0x31, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x65, 0x76, 0x65, 0x72, 0x6c,
	0x61, 0x6e, 0x64, 0x4d, 0x4a, 0x2f, 0x47, 0x52, 0x55, 0x44, 0x2f, 0x67, 0x72, 0x75, 0x64, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x67, 0x72, 0x75, 0x64, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_grudpb_grud_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_grudpb_grud_proto_goTypes = []interface{}{
	(*GetUserPostsRequest)(nil),  // 0: collectpb.GetUserPostsRequest
	(*Data)(nil),                 // 1: collectpb.Data
	(*GetUserPostsResponse)(nil), // 2: collectpb.GetUserPostsResponse
	(*GetPostByIDRequest)(nil),   // 3: collectpb.GetPostByIDRequest
}
var file_grudpb_grud_proto_depIdxs = []int32{
	1, // 0: collectpb.GetUserPostsResponse.datas:type_name -> collectpb.Data
	0, // 1: collectpb.GrudService.GetPostsByUserID:input_type -> collectpb.GetUserPostsRequest
	3, // 2: collectpb.GrudService.GetPostByID:input_type -> collectpb.GetPostByIDRequest
	2, // 3: collectpb.GrudService.GetPostsByUserID:output_type -> collectpb.GetUserPostsResponse
	1, // 4: collectpb.GrudService.GetPostByID:output_type -> collectpb.Data
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
		file_grudpb_grud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserPostsResponse); i {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grudpb_grud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
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
