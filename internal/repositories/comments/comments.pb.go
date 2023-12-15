// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: comments.proto

package comments

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

type CommentGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PostId uint32 `protobuf:"varint,3,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
	Text   string `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CommentGRPC) Reset() {
	*x = CommentGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentGRPC) ProtoMessage() {}

func (x *CommentGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentGRPC.ProtoReflect.Descriptor instead.
func (*CommentGRPC) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{0}
}

func (x *CommentGRPC) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CommentGRPC) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CommentGRPC) GetPostId() uint32 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *CommentGRPC) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type CommentsGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentsMap map[int32]*CommentGRPC `protobuf:"bytes,1,rep,name=comments_map,json=commentsMap,proto3" json:"comments_map,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CommentsGRPC) Reset() {
	*x = CommentsGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentsGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentsGRPC) ProtoMessage() {}

func (x *CommentsGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentsGRPC.ProtoReflect.Descriptor instead.
func (*CommentsGRPC) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{1}
}

func (x *CommentsGRPC) GetCommentsMap() map[int32]*CommentGRPC {
	if x != nil {
		return x.CommentsMap
	}
	return nil
}

type CommentsMapGRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments map[int32]*CommentsGRPC `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *CommentsMapGRPC) Reset() {
	*x = CommentsMapGRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentsMapGRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentsMapGRPC) ProtoMessage() {}

func (x *CommentsMapGRPC) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentsMapGRPC.ProtoReflect.Descriptor instead.
func (*CommentsMapGRPC) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{2}
}

func (x *CommentsMapGRPC) GetComments() map[int32]*CommentsGRPC {
	if x != nil {
		return x.Comments
	}
	return nil
}

type UInt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UInt) Reset() {
	*x = UInt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UInt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UInt) ProtoMessage() {}

func (x *UInt) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UInt.ProtoReflect.Descriptor instead.
func (*UInt) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{3}
}

func (x *UInt) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Int struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I int32 `protobuf:"varint,1,opt,name=i,proto3" json:"i,omitempty"`
}

func (x *Int) Reset() {
	*x = Int{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Int) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int) ProtoMessage() {}

func (x *Int) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Int.ProtoReflect.Descriptor instead.
func (*Int) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{4}
}

func (x *Int) GetI() int32 {
	if x != nil {
		return x.I
	}
	return 0
}

type Nothing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dummy bool `protobuf:"varint,1,opt,name=dummy,proto3" json:"dummy,omitempty"`
}

func (x *Nothing) Reset() {
	*x = Nothing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nothing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nothing) ProtoMessage() {}

func (x *Nothing) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nothing.ProtoReflect.Descriptor instead.
func (*Nothing) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{5}
}

func (x *Nothing) GetDummy() bool {
	if x != nil {
		return x.Dummy
	}
	return false
}

var File_comments_proto protoreflect.FileDescriptor

var file_comments_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x63, 0x0a, 0x0b, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x52, 0x50, 0x43, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22,
	0xb1, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x47, 0x52, 0x50, 0x43,
	0x12, 0x4a, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f, 0x6d, 0x61, 0x70,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x47, 0x52, 0x50, 0x43, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4d, 0x61, 0x70, 0x1a, 0x55, 0x0a, 0x10,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4d, 0x61, 0x70, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x47, 0x52, 0x50, 0x43, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xab, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x4d, 0x61, 0x70, 0x47, 0x52, 0x50, 0x43, 0x12, 0x43, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4d, 0x61, 0x70,
	0x47, 0x52, 0x50, 0x43, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0x53, 0x0a, 0x0d,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x47, 0x52, 0x50, 0x43, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x16, 0x0a, 0x04, 0x55, 0x49, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x13, 0x0a, 0x03, 0x49, 0x6e, 0x74,
	0x12, 0x0c, 0x0a, 0x01, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x69, 0x22, 0x1f,
	0x0a, 0x07, 0x4e, 0x6f, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x75, 0x6d,
	0x6d, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x64, 0x75, 0x6d, 0x6d, 0x79, 0x32,
	0xc5, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3a, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x43, 0x74, 0x78, 0x12, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x52, 0x50, 0x43, 0x1a, 0x0d, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x22, 0x00, 0x12, 0x37,
	0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x74, 0x78, 0x12, 0x0e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x55, 0x49,
	0x6e, 0x74, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4e, 0x6f,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x74, 0x78, 0x12, 0x15, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x52,
	0x50, 0x43, 0x1a, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4e, 0x6f,
	0x74, 0x68, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comments_proto_rawDescOnce sync.Once
	file_comments_proto_rawDescData = file_comments_proto_rawDesc
)

func file_comments_proto_rawDescGZIP() []byte {
	file_comments_proto_rawDescOnce.Do(func() {
		file_comments_proto_rawDescData = protoimpl.X.CompressGZIP(file_comments_proto_rawDescData)
	})
	return file_comments_proto_rawDescData
}

var file_comments_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_comments_proto_goTypes = []interface{}{
	(*CommentGRPC)(nil),     // 0: comments.CommentGRPC
	(*CommentsGRPC)(nil),    // 1: comments.CommentsGRPC
	(*CommentsMapGRPC)(nil), // 2: comments.CommentsMapGRPC
	(*UInt)(nil),            // 3: comments.UInt
	(*Int)(nil),             // 4: comments.Int
	(*Nothing)(nil),         // 5: comments.Nothing
	nil,                     // 6: comments.CommentsGRPC.CommentsMapEntry
	nil,                     // 7: comments.CommentsMapGRPC.CommentsEntry
}
var file_comments_proto_depIdxs = []int32{
	6, // 0: comments.CommentsGRPC.comments_map:type_name -> comments.CommentsGRPC.CommentsMapEntry
	7, // 1: comments.CommentsMapGRPC.comments:type_name -> comments.CommentsMapGRPC.CommentsEntry
	0, // 2: comments.CommentsGRPC.CommentsMapEntry.value:type_name -> comments.CommentGRPC
	1, // 3: comments.CommentsMapGRPC.CommentsEntry.value:type_name -> comments.CommentsGRPC
	0, // 4: comments.CommentService.CreateCommentCtx:input_type -> comments.CommentGRPC
	3, // 5: comments.CommentService.DeleteCommentCtx:input_type -> comments.UInt
	0, // 6: comments.CommentService.ChangeCommentCtx:input_type -> comments.CommentGRPC
	4, // 7: comments.CommentService.CreateCommentCtx:output_type -> comments.Int
	5, // 8: comments.CommentService.DeleteCommentCtx:output_type -> comments.Nothing
	5, // 9: comments.CommentService.ChangeCommentCtx:output_type -> comments.Nothing
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_comments_proto_init() }
func file_comments_proto_init() {
	if File_comments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentGRPC); i {
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
		file_comments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentsGRPC); i {
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
		file_comments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentsMapGRPC); i {
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
		file_comments_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UInt); i {
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
		file_comments_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Int); i {
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
		file_comments_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nothing); i {
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
			RawDescriptor: file_comments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comments_proto_goTypes,
		DependencyIndexes: file_comments_proto_depIdxs,
		MessageInfos:      file_comments_proto_msgTypes,
	}.Build()
	File_comments_proto = out.File
	file_comments_proto_rawDesc = nil
	file_comments_proto_goTypes = nil
	file_comments_proto_depIdxs = nil
}
