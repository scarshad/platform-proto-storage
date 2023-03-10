// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.21.12
// source: contract.proto

package tracer

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

type Album struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title  string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Artist string  `protobuf:"bytes,3,opt,name=artist,proto3" json:"artist,omitempty"`
	Price  float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Album) Reset() {
	*x = Album{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Album) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Album) ProtoMessage() {}

func (x *Album) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Album.ProtoReflect.Descriptor instead.
func (*Album) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{0}
}

func (x *Album) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Album) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Album) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

func (x *Album) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type CreateAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Album *Album `protobuf:"bytes,1,opt,name=album,proto3" json:"album,omitempty"`
}

func (x *CreateAlbumRequest) Reset() {
	*x = CreateAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlbumRequest) ProtoMessage() {}

func (x *CreateAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlbumRequest.ProtoReflect.Descriptor instead.
func (*CreateAlbumRequest) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAlbumRequest) GetAlbum() *Album {
	if x != nil {
		return x.Album
	}
	return nil
}

type CreateAlbumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Album *Album `protobuf:"bytes,1,opt,name=Album,proto3" json:"Album,omitempty"`
}

func (x *CreateAlbumResponse) Reset() {
	*x = CreateAlbumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAlbumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAlbumResponse) ProtoMessage() {}

func (x *CreateAlbumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAlbumResponse.ProtoReflect.Descriptor instead.
func (*CreateAlbumResponse) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAlbumResponse) GetAlbum() *Album {
	if x != nil {
		return x.Album
	}
	return nil
}

type GetAlbumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAlbumRequest) Reset() {
	*x = GetAlbumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlbumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlbumRequest) ProtoMessage() {}

func (x *GetAlbumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlbumRequest.ProtoReflect.Descriptor instead.
func (*GetAlbumRequest) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{3}
}

type GetAlbumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Album []*Album `protobuf:"bytes,1,rep,name=album,proto3" json:"album,omitempty"`
}

func (x *GetAlbumResponse) Reset() {
	*x = GetAlbumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlbumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlbumResponse) ProtoMessage() {}

func (x *GetAlbumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlbumResponse.ProtoReflect.Descriptor instead.
func (*GetAlbumResponse) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{4}
}

func (x *GetAlbumResponse) GetAlbum() []*Album {
	if x != nil {
		return x.Album
	}
	return nil
}

var File_contract_proto protoreflect.FileDescriptor

var file_contract_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x22, 0x5b, 0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x39, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x05, 0x61,
	0x6c, 0x62, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d,
	0x22, 0x3a, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x2e,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x22, 0x11, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x2e, 0x41, 0x6c, 0x62, 0x75,
	0x6d, 0x52, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x32, 0x97, 0x01, 0x0a, 0x0a, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x48, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x1a, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x17, 0x2e,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x6f, 0x3b, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_proto_rawDescOnce sync.Once
	file_contract_proto_rawDescData = file_contract_proto_rawDesc
)

func file_contract_proto_rawDescGZIP() []byte {
	file_contract_proto_rawDescOnce.Do(func() {
		file_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_proto_rawDescData)
	})
	return file_contract_proto_rawDescData
}

var file_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_contract_proto_goTypes = []interface{}{
	(*Album)(nil),               // 0: tracer.Album
	(*CreateAlbumRequest)(nil),  // 1: tracer.CreateAlbumRequest
	(*CreateAlbumResponse)(nil), // 2: tracer.CreateAlbumResponse
	(*GetAlbumRequest)(nil),     // 3: tracer.GetAlbumRequest
	(*GetAlbumResponse)(nil),    // 4: tracer.GetAlbumResponse
}
var file_contract_proto_depIdxs = []int32{
	0, // 0: tracer.CreateAlbumRequest.album:type_name -> tracer.Album
	0, // 1: tracer.CreateAlbumResponse.Album:type_name -> tracer.Album
	0, // 2: tracer.GetAlbumResponse.album:type_name -> tracer.Album
	1, // 3: tracer.HelloWorld.CreateAlbum:input_type -> tracer.CreateAlbumRequest
	3, // 4: tracer.HelloWorld.GetAlbum:input_type -> tracer.GetAlbumRequest
	2, // 5: tracer.HelloWorld.CreateAlbum:output_type -> tracer.CreateAlbumResponse
	4, // 6: tracer.HelloWorld.GetAlbum:output_type -> tracer.GetAlbumResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_contract_proto_init() }
func file_contract_proto_init() {
	if File_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Album); i {
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
		file_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAlbumRequest); i {
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
		file_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAlbumResponse); i {
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
		file_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlbumRequest); i {
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
		file_contract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAlbumResponse); i {
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
			RawDescriptor: file_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contract_proto_goTypes,
		DependencyIndexes: file_contract_proto_depIdxs,
		MessageInfos:      file_contract_proto_msgTypes,
	}.Build()
	File_contract_proto = out.File
	file_contract_proto_rawDesc = nil
	file_contract_proto_goTypes = nil
	file_contract_proto_depIdxs = nil
}
