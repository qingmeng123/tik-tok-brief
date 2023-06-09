// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: proto/video.proto

package pb

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

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	VideoId       int64  `protobuf:"varint,2,opt,name=videoId,proto3" json:"videoId,omitempty"`             // 视频唯一标识
	UserId        int64  `protobuf:"varint,3,opt,name=userId,proto3" json:"userId,omitempty"`               //视频所属用户id
	PlayUrl       string `protobuf:"bytes,4,opt,name=playUrl,proto3" json:"playUrl,omitempty"`              // 视频播放地址
	CoverUrl      string `protobuf:"bytes,5,opt,name=coverUrl,proto3" json:"coverUrl,omitempty"`            // 视频封面地址
	FavoriteCount int64  `protobuf:"varint,6,opt,name=favoriteCount,proto3" json:"favoriteCount,omitempty"` // 视频的点赞总数
	CommentCount  int64  `protobuf:"varint,7,opt,name=commentCount,proto3" json:"commentCount,omitempty"`   // 视频的评论总数
	IsFavorite    bool   `protobuf:"varint,8,opt,name=isFavorite,proto3" json:"isFavorite,omitempty"`       // true-已点赞，false-未点赞
	Title         string `protobuf:"bytes,9,opt,name=title,proto3" json:"title,omitempty"`                  // 视频标题
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{0}
}

func (x *Video) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *Video) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Video) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type SaveVideoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	UserId   int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	PlayUrl  string `protobuf:"bytes,3,opt,name=playUrl,proto3" json:"playUrl,omitempty"`
	CoverUrl string `protobuf:"bytes,4,opt,name=coverUrl,proto3" json:"coverUrl,omitempty"`
}

func (x *SaveVideoReq) Reset() {
	*x = SaveVideoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveVideoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveVideoReq) ProtoMessage() {}

func (x *SaveVideoReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveVideoReq.ProtoReflect.Descriptor instead.
func (*SaveVideoReq) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{1}
}

func (x *SaveVideoReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SaveVideoReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *SaveVideoReq) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *SaveVideoReq) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

type SaveVideoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId int64 `protobuf:"varint,1,opt,name=videoId,proto3" json:"videoId,omitempty"`
}

func (x *SaveVideoResp) Reset() {
	*x = SaveVideoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveVideoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveVideoResp) ProtoMessage() {}

func (x *SaveVideoResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveVideoResp.ProtoReflect.Descriptor instead.
func (*SaveVideoResp) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{2}
}

func (x *SaveVideoResp) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type PublishListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *PublishListReq) Reset() {
	*x = PublishListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishListReq) ProtoMessage() {}

func (x *PublishListReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishListReq.ProtoReflect.Descriptor instead.
func (*PublishListReq) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{3}
}

func (x *PublishListReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type PublishListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoList []*Video `protobuf:"bytes,1,rep,name=videoList,proto3" json:"videoList,omitempty"`
}

func (x *PublishListResp) Reset() {
	*x = PublishListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishListResp) ProtoMessage() {}

func (x *PublishListResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishListResp.ProtoReflect.Descriptor instead.
func (*PublishListResp) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{4}
}

func (x *PublishListResp) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

// 视频流请求体
type FeedReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   *int64 `protobuf:"varint,1,opt,name=userId,proto3,oneof" json:"userId,omitempty"`
	LastTime *int64 `protobuf:"varint,2,opt,name=lastTime,proto3,oneof" json:"lastTime,omitempty"`
}

func (x *FeedReq) Reset() {
	*x = FeedReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedReq) ProtoMessage() {}

func (x *FeedReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedReq.ProtoReflect.Descriptor instead.
func (*FeedReq) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{5}
}

func (x *FeedReq) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *FeedReq) GetLastTime() int64 {
	if x != nil && x.LastTime != nil {
		return *x.LastTime
	}
	return 0
}

// 视频流返回体
type FeedResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextTime  *int64   `protobuf:"varint,1,opt,name=nextTime,proto3,oneof" json:"nextTime,omitempty"`
	VideoList []*Video `protobuf:"bytes,2,rep,name=videoList,proto3" json:"videoList,omitempty"`
}

func (x *FeedResp) Reset() {
	*x = FeedResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedResp) ProtoMessage() {}

func (x *FeedResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedResp.ProtoReflect.Descriptor instead.
func (*FeedResp) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{6}
}

func (x *FeedResp) GetNextTime() int64 {
	if x != nil && x.NextTime != nil {
		return *x.NextTime
	}
	return 0
}

func (x *FeedResp) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

// 更新视频点赞数
type UpdateFavoriteCountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId int64 `protobuf:"varint,1,opt,name=videoId,proto3" json:"videoId,omitempty"`
	Number  int64 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *UpdateFavoriteCountReq) Reset() {
	*x = UpdateFavoriteCountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFavoriteCountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFavoriteCountReq) ProtoMessage() {}

func (x *UpdateFavoriteCountReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFavoriteCountReq.ProtoReflect.Descriptor instead.
func (*UpdateFavoriteCountReq) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateFavoriteCountReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *UpdateFavoriteCountReq) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type UpdateFavoriteCountResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateFavoriteCountResp) Reset() {
	*x = UpdateFavoriteCountResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFavoriteCountResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFavoriteCountResp) ProtoMessage() {}

func (x *UpdateFavoriteCountResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFavoriteCountResp.ProtoReflect.Descriptor instead.
func (*UpdateFavoriteCountResp) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{8}
}

// 更新视频评论数
type UpdateCommentCountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId int64 `protobuf:"varint,1,opt,name=videoId,proto3" json:"videoId,omitempty"`
	Number  int64 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *UpdateCommentCountReq) Reset() {
	*x = UpdateCommentCountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentCountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentCountReq) ProtoMessage() {}

func (x *UpdateCommentCountReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentCountReq.ProtoReflect.Descriptor instead.
func (*UpdateCommentCountReq) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateCommentCountReq) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *UpdateCommentCountReq) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type UpdateCommentCountResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCommentCountResp) Reset() {
	*x = UpdateCommentCountResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentCountResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentCountResp) ProtoMessage() {}

func (x *UpdateCommentCountResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentCountResp.ProtoReflect.Descriptor instead.
func (*UpdateCommentCountResp) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{10}
}

// 获取视频列表
type GetVideoListByIdsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   *int64  `protobuf:"varint,1,opt,name=userId,proto3,oneof" json:"userId,omitempty"`
	VideoIds []int64 `protobuf:"varint,2,rep,packed,name=videoIds,proto3" json:"videoIds,omitempty"`
}

func (x *GetVideoListByIdsReq) Reset() {
	*x = GetVideoListByIdsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoListByIdsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoListByIdsReq) ProtoMessage() {}

func (x *GetVideoListByIdsReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoListByIdsReq.ProtoReflect.Descriptor instead.
func (*GetVideoListByIdsReq) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{11}
}

func (x *GetVideoListByIdsReq) GetUserId() int64 {
	if x != nil && x.UserId != nil {
		return *x.UserId
	}
	return 0
}

func (x *GetVideoListByIdsReq) GetVideoIds() []int64 {
	if x != nil {
		return x.VideoIds
	}
	return nil
}

type GetVideoListByIdsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoList []*Video `protobuf:"bytes,1,rep,name=videoList,proto3" json:"videoList,omitempty"`
}

func (x *GetVideoListByIdsResp) Reset() {
	*x = GetVideoListByIdsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_video_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoListByIdsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoListByIdsResp) ProtoMessage() {}

func (x *GetVideoListByIdsResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_video_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoListByIdsResp.ProtoReflect.Descriptor instead.
func (*GetVideoListByIdsResp) Descriptor() ([]byte, []int) {
	return file_proto_video_proto_rawDescGZIP(), []int{12}
}

func (x *GetVideoListByIdsResp) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

var File_proto_video_proto protoreflect.FileDescriptor

var file_proto_video_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0xff, 0x01, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72,
	0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x72, 0x0a, 0x0c, 0x53, 0x61, 0x76,
	0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x22, 0x29, 0x0a,
	0x0d, 0x53, 0x61, 0x76, 0x65, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x0e, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x3a, 0x0a, 0x0f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x27, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x5f,
	0x0a, 0x07, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22,
	0x61, 0x0a, 0x08, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x08, 0x6e,
	0x65, 0x78, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x08, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x09,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x09, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x19,
	0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x49, 0x0a, 0x15, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x5a,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x73, 0x42,
	0x09, 0x0a, 0x07, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x15, 0x47, 0x65,
	0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x27, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x32, 0xfc, 0x02, 0x0a,
	0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x30, 0x0a, 0x09, 0x53, 0x61, 0x76, 0x65, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x21, 0x0a, 0x04, 0x46, 0x65, 0x65, 0x64, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65,
	0x65, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x4e, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x4c, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x48, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x73, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_video_proto_rawDescOnce sync.Once
	file_proto_video_proto_rawDescData = file_proto_video_proto_rawDesc
)

func file_proto_video_proto_rawDescGZIP() []byte {
	file_proto_video_proto_rawDescOnce.Do(func() {
		file_proto_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_video_proto_rawDescData)
	})
	return file_proto_video_proto_rawDescData
}

var file_proto_video_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_video_proto_goTypes = []interface{}{
	(*Video)(nil),                   // 0: pb.Video
	(*SaveVideoReq)(nil),            // 1: pb.SaveVideoReq
	(*SaveVideoResp)(nil),           // 2: pb.SaveVideoResp
	(*PublishListReq)(nil),          // 3: pb.PublishListReq
	(*PublishListResp)(nil),         // 4: pb.PublishListResp
	(*FeedReq)(nil),                 // 5: pb.FeedReq
	(*FeedResp)(nil),                // 6: pb.FeedResp
	(*UpdateFavoriteCountReq)(nil),  // 7: pb.UpdateFavoriteCountReq
	(*UpdateFavoriteCountResp)(nil), // 8: pb.UpdateFavoriteCountResp
	(*UpdateCommentCountReq)(nil),   // 9: pb.UpdateCommentCountReq
	(*UpdateCommentCountResp)(nil),  // 10: pb.UpdateCommentCountResp
	(*GetVideoListByIdsReq)(nil),    // 11: pb.GetVideoListByIdsReq
	(*GetVideoListByIdsResp)(nil),   // 12: pb.GetVideoListByIdsResp
}
var file_proto_video_proto_depIdxs = []int32{
	0,  // 0: pb.PublishListResp.videoList:type_name -> pb.Video
	0,  // 1: pb.FeedResp.videoList:type_name -> pb.Video
	0,  // 2: pb.GetVideoListByIdsResp.videoList:type_name -> pb.Video
	1,  // 3: pb.video.SaveVideo:input_type -> pb.SaveVideoReq
	3,  // 4: pb.video.PublishList:input_type -> pb.PublishListReq
	5,  // 5: pb.video.Feed:input_type -> pb.FeedReq
	7,  // 6: pb.video.UpdateFavoriteCount:input_type -> pb.UpdateFavoriteCountReq
	9,  // 7: pb.video.UpdateCommentCount:input_type -> pb.UpdateCommentCountReq
	11, // 8: pb.video.GetVideoListByIds:input_type -> pb.GetVideoListByIdsReq
	2,  // 9: pb.video.SaveVideo:output_type -> pb.SaveVideoResp
	4,  // 10: pb.video.PublishList:output_type -> pb.PublishListResp
	6,  // 11: pb.video.Feed:output_type -> pb.FeedResp
	8,  // 12: pb.video.UpdateFavoriteCount:output_type -> pb.UpdateFavoriteCountResp
	8,  // 13: pb.video.UpdateCommentCount:output_type -> pb.UpdateFavoriteCountResp
	12, // 14: pb.video.GetVideoListByIds:output_type -> pb.GetVideoListByIdsResp
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_proto_video_proto_init() }
func file_proto_video_proto_init() {
	if File_proto_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
		file_proto_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveVideoReq); i {
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
		file_proto_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveVideoResp); i {
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
		file_proto_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishListReq); i {
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
		file_proto_video_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishListResp); i {
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
		file_proto_video_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedReq); i {
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
		file_proto_video_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedResp); i {
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
		file_proto_video_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFavoriteCountReq); i {
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
		file_proto_video_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFavoriteCountResp); i {
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
		file_proto_video_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentCountReq); i {
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
		file_proto_video_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentCountResp); i {
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
		file_proto_video_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoListByIdsReq); i {
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
		file_proto_video_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoListByIdsResp); i {
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
	file_proto_video_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_proto_video_proto_msgTypes[6].OneofWrappers = []interface{}{}
	file_proto_video_proto_msgTypes[11].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_video_proto_goTypes,
		DependencyIndexes: file_proto_video_proto_depIdxs,
		MessageInfos:      file_proto_video_proto_msgTypes,
	}.Build()
	File_proto_video_proto = out.File
	file_proto_video_proto_rawDesc = nil
	file_proto_video_proto_goTypes = nil
	file_proto_video_proto_depIdxs = nil
}
