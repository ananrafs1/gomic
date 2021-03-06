// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: orchestrator/proto/scrapper.proto

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

type ScrapRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Page     int64  `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
	Quantity int64  `protobuf:"varint,3,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
}

func (x *ScrapRequest) Reset() {
	*x = ScrapRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_scrapper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScrapRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScrapRequest) ProtoMessage() {}

func (x *ScrapRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_scrapper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScrapRequest.ProtoReflect.Descriptor instead.
func (*ScrapRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_scrapper_proto_rawDescGZIP(), []int{0}
}

func (x *ScrapRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ScrapRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ScrapRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ScrapPerChapterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=Title,proto3" json:"Title,omitempty"`
	Id    string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *ScrapPerChapterRequest) Reset() {
	*x = ScrapPerChapterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_scrapper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScrapPerChapterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScrapPerChapterRequest) ProtoMessage() {}

func (x *ScrapPerChapterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_scrapper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScrapPerChapterRequest.ProtoReflect.Descriptor instead.
func (*ScrapPerChapterRequest) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_scrapper_proto_rawDescGZIP(), []int{1}
}

func (x *ScrapPerChapterRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ScrapPerChapterRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ScrapResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comic *Comic `protobuf:"bytes,1,opt,name=Comic,proto3" json:"Comic,omitempty"`
}

func (x *ScrapResponse) Reset() {
	*x = ScrapResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_scrapper_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScrapResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScrapResponse) ProtoMessage() {}

func (x *ScrapResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_scrapper_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScrapResponse.ProtoReflect.Descriptor instead.
func (*ScrapResponse) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_scrapper_proto_rawDescGZIP(), []int{2}
}

func (x *ScrapResponse) GetComic() *Comic {
	if x != nil {
		return x.Comic
	}
	return nil
}

type Comic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComicFlat *ComicFlat `protobuf:"bytes,1,opt,name=ComicFlat,proto3" json:"ComicFlat,omitempty"`
	Chapters  []*Chapter `protobuf:"bytes,2,rep,name=Chapters,proto3" json:"Chapters,omitempty"`
}

func (x *Comic) Reset() {
	*x = Comic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_scrapper_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comic) ProtoMessage() {}

func (x *Comic) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_scrapper_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comic.ProtoReflect.Descriptor instead.
func (*Comic) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_scrapper_proto_rawDescGZIP(), []int{3}
}

func (x *Comic) GetComicFlat() *ComicFlat {
	if x != nil {
		return x.ComicFlat
	}
	return nil
}

func (x *Comic) GetChapters() []*Chapter {
	if x != nil {
		return x.Chapters
	}
	return nil
}

type ComicFlat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Host string `protobuf:"bytes,3,opt,name=Host,proto3" json:"Host,omitempty"`
}

func (x *ComicFlat) Reset() {
	*x = ComicFlat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_scrapper_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComicFlat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComicFlat) ProtoMessage() {}

func (x *ComicFlat) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_scrapper_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComicFlat.ProtoReflect.Descriptor instead.
func (*ComicFlat) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_scrapper_proto_rawDescGZIP(), []int{4}
}

func (x *ComicFlat) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ComicFlat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ComicFlat) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type Chapter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChapterFlat *ChapterFlat     `protobuf:"bytes,1,opt,name=ChapterFlat,proto3" json:"ChapterFlat,omitempty"`
	Images      []*ImageProvider `protobuf:"bytes,2,rep,name=Images,proto3" json:"Images,omitempty"`
}

func (x *Chapter) Reset() {
	*x = Chapter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_scrapper_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chapter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chapter) ProtoMessage() {}

func (x *Chapter) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_scrapper_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chapter.ProtoReflect.Descriptor instead.
func (*Chapter) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_scrapper_proto_rawDescGZIP(), []int{5}
}

func (x *Chapter) GetChapterFlat() *ChapterFlat {
	if x != nil {
		return x.ChapterFlat
	}
	return nil
}

func (x *Chapter) GetImages() []*ImageProvider {
	if x != nil {
		return x.Images
	}
	return nil
}

type ChapterFlat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *ChapterFlat) Reset() {
	*x = ChapterFlat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_scrapper_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChapterFlat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChapterFlat) ProtoMessage() {}

func (x *ChapterFlat) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_scrapper_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChapterFlat.ProtoReflect.Descriptor instead.
func (*ChapterFlat) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_scrapper_proto_rawDescGZIP(), []int{6}
}

func (x *ChapterFlat) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ImageProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Episode  int64  `protobuf:"varint,1,opt,name=Episode,proto3" json:"Episode,omitempty"`
	Provider string `protobuf:"bytes,2,opt,name=Provider,proto3" json:"Provider,omitempty"`
	Link     string `protobuf:"bytes,3,opt,name=Link,proto3" json:"Link,omitempty"`
}

func (x *ImageProvider) Reset() {
	*x = ImageProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_scrapper_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImageProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageProvider) ProtoMessage() {}

func (x *ImageProvider) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_scrapper_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageProvider.ProtoReflect.Descriptor instead.
func (*ImageProvider) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_scrapper_proto_rawDescGZIP(), []int{7}
}

func (x *ImageProvider) GetEpisode() int64 {
	if x != nil {
		return x.Episode
	}
	return 0
}

func (x *ImageProvider) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *ImageProvider) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type ScrapPerChapterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Chapter *Chapter `protobuf:"bytes,1,opt,name=Chapter,proto3" json:"Chapter,omitempty"`
}

func (x *ScrapPerChapterResponse) Reset() {
	*x = ScrapPerChapterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orchestrator_proto_scrapper_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScrapPerChapterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScrapPerChapterResponse) ProtoMessage() {}

func (x *ScrapPerChapterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_orchestrator_proto_scrapper_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScrapPerChapterResponse.ProtoReflect.Descriptor instead.
func (*ScrapPerChapterResponse) Descriptor() ([]byte, []int) {
	return file_orchestrator_proto_scrapper_proto_rawDescGZIP(), []int{8}
}

func (x *ScrapPerChapterResponse) GetChapter() *Chapter {
	if x != nil {
		return x.Chapter
	}
	return nil
}

var File_orchestrator_proto_scrapper_proto protoreflect.FileDescriptor

var file_orchestrator_proto_scrapper_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x0c, 0x53, 0x63,
	0x72, 0x61, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x50, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x22, 0x3e, 0x0a, 0x16, 0x53, 0x63, 0x72, 0x61, 0x70, 0x50, 0x65, 0x72, 0x43, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64,
	0x22, 0x33, 0x0a, 0x0d, 0x53, 0x63, 0x72, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x05, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x52, 0x05,
	0x43, 0x6f, 0x6d, 0x69, 0x63, 0x22, 0x63, 0x0a, 0x05, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x12, 0x2e,
	0x0a, 0x09, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x46, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x46,
	0x6c, 0x61, 0x74, 0x52, 0x09, 0x43, 0x6f, 0x6d, 0x69, 0x63, 0x46, 0x6c, 0x61, 0x74, 0x12, 0x2a,
	0x0a, 0x08, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x52, 0x08, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x73, 0x22, 0x43, 0x0a, 0x09, 0x43, 0x6f,
	0x6d, 0x69, 0x63, 0x46, 0x6c, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x48,
	0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x48, 0x6f, 0x73, 0x74, 0x22,
	0x6d, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x0b, 0x43, 0x68,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x46, 0x6c, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x46,
	0x6c, 0x61, 0x74, 0x52, 0x0b, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x46, 0x6c, 0x61, 0x74,
	0x12, 0x2c, 0x0a, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x1d,
	0x0a, 0x0b, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x46, 0x6c, 0x61, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x22, 0x59, 0x0a,
	0x0d, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x43, 0x0a, 0x17, 0x53, 0x63, 0x72, 0x61,
	0x70, 0x50, 0x65, 0x72, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x52, 0x07, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x32, 0x94, 0x01,
	0x0a, 0x08, 0x53, 0x63, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x12, 0x34, 0x0a, 0x05, 0x53, 0x63,
	0x72, 0x61, 0x70, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x72, 0x61,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x63, 0x72, 0x61, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x52, 0x0a, 0x0f, 0x53, 0x63, 0x72, 0x61, 0x70, 0x50, 0x65, 0x72, 0x43, 0x68, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x72, 0x61,
	0x70, 0x50, 0x65, 0x72, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x63, 0x72, 0x61, 0x70,
	0x50, 0x65, 0x72, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_orchestrator_proto_scrapper_proto_rawDescOnce sync.Once
	file_orchestrator_proto_scrapper_proto_rawDescData = file_orchestrator_proto_scrapper_proto_rawDesc
)

func file_orchestrator_proto_scrapper_proto_rawDescGZIP() []byte {
	file_orchestrator_proto_scrapper_proto_rawDescOnce.Do(func() {
		file_orchestrator_proto_scrapper_proto_rawDescData = protoimpl.X.CompressGZIP(file_orchestrator_proto_scrapper_proto_rawDescData)
	})
	return file_orchestrator_proto_scrapper_proto_rawDescData
}

var file_orchestrator_proto_scrapper_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_orchestrator_proto_scrapper_proto_goTypes = []interface{}{
	(*ScrapRequest)(nil),            // 0: proto.ScrapRequest
	(*ScrapPerChapterRequest)(nil),  // 1: proto.ScrapPerChapterRequest
	(*ScrapResponse)(nil),           // 2: proto.ScrapResponse
	(*Comic)(nil),                   // 3: proto.Comic
	(*ComicFlat)(nil),               // 4: proto.ComicFlat
	(*Chapter)(nil),                 // 5: proto.Chapter
	(*ChapterFlat)(nil),             // 6: proto.ChapterFlat
	(*ImageProvider)(nil),           // 7: proto.ImageProvider
	(*ScrapPerChapterResponse)(nil), // 8: proto.ScrapPerChapterResponse
}
var file_orchestrator_proto_scrapper_proto_depIdxs = []int32{
	3, // 0: proto.ScrapResponse.Comic:type_name -> proto.Comic
	4, // 1: proto.Comic.ComicFlat:type_name -> proto.ComicFlat
	5, // 2: proto.Comic.Chapters:type_name -> proto.Chapter
	6, // 3: proto.Chapter.ChapterFlat:type_name -> proto.ChapterFlat
	7, // 4: proto.Chapter.Images:type_name -> proto.ImageProvider
	5, // 5: proto.ScrapPerChapterResponse.Chapter:type_name -> proto.Chapter
	0, // 6: proto.Scrapper.Scrap:input_type -> proto.ScrapRequest
	1, // 7: proto.Scrapper.ScrapPerChapter:input_type -> proto.ScrapPerChapterRequest
	2, // 8: proto.Scrapper.Scrap:output_type -> proto.ScrapResponse
	8, // 9: proto.Scrapper.ScrapPerChapter:output_type -> proto.ScrapPerChapterResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_orchestrator_proto_scrapper_proto_init() }
func file_orchestrator_proto_scrapper_proto_init() {
	if File_orchestrator_proto_scrapper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orchestrator_proto_scrapper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScrapRequest); i {
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
		file_orchestrator_proto_scrapper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScrapPerChapterRequest); i {
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
		file_orchestrator_proto_scrapper_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScrapResponse); i {
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
		file_orchestrator_proto_scrapper_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comic); i {
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
		file_orchestrator_proto_scrapper_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComicFlat); i {
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
		file_orchestrator_proto_scrapper_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chapter); i {
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
		file_orchestrator_proto_scrapper_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChapterFlat); i {
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
		file_orchestrator_proto_scrapper_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImageProvider); i {
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
		file_orchestrator_proto_scrapper_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScrapPerChapterResponse); i {
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
			RawDescriptor: file_orchestrator_proto_scrapper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_orchestrator_proto_scrapper_proto_goTypes,
		DependencyIndexes: file_orchestrator_proto_scrapper_proto_depIdxs,
		MessageInfos:      file_orchestrator_proto_scrapper_proto_msgTypes,
	}.Build()
	File_orchestrator_proto_scrapper_proto = out.File
	file_orchestrator_proto_scrapper_proto_rawDesc = nil
	file_orchestrator_proto_scrapper_proto_goTypes = nil
	file_orchestrator_proto_scrapper_proto_depIdxs = nil
}
