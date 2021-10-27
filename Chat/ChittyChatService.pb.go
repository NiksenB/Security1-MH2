// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: Chat/ChittyChatService.proto

package Chat

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

type FromClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Body    string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Lamport int32  `protobuf:"varint,3,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *FromClient) Reset() {
	*x = FromClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chat_ChittyChatService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FromClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromClient) ProtoMessage() {}

func (x *FromClient) ProtoReflect() protoreflect.Message {
	mi := &file_Chat_ChittyChatService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromClient.ProtoReflect.Descriptor instead.
func (*FromClient) Descriptor() ([]byte, []int) {
	return file_Chat_ChittyChatService_proto_rawDescGZIP(), []int{0}
}

func (x *FromClient) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FromClient) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *FromClient) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

type FromServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Body    string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	Lamport int32  `protobuf:"varint,3,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *FromServer) Reset() {
	*x = FromServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chat_ChittyChatService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FromServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FromServer) ProtoMessage() {}

func (x *FromServer) ProtoReflect() protoreflect.Message {
	mi := &file_Chat_ChittyChatService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FromServer.ProtoReflect.Descriptor instead.
func (*FromServer) Descriptor() ([]byte, []int) {
	return file_Chat_ChittyChatService_proto_rawDescGZIP(), []int{1}
}

func (x *FromServer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FromServer) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *FromServer) GetLamport() int32 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

type UserName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UserName) Reset() {
	*x = UserName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chat_ChittyChatService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserName) ProtoMessage() {}

func (x *UserName) ProtoReflect() protoreflect.Message {
	mi := &file_Chat_ChittyChatService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserName.ProtoReflect.Descriptor instead.
func (*UserName) Descriptor() ([]byte, []int) {
	return file_Chat_ChittyChatService_proto_rawDescGZIP(), []int{2}
}

func (x *UserName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chat_ChittyChatService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_Chat_ChittyChatService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_Chat_ChittyChatService_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Chat_ChittyChatService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_Chat_ChittyChatService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_Chat_ChittyChatService_proto_rawDescGZIP(), []int{4}
}

var File_Chat_ChittyChatService_proto protoreflect.FileDescriptor

var file_Chat_ChittyChatService_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x43, 0x68, 0x61, 0x74, 0x2f, 0x43, 0x68, 0x69, 0x74, 0x74, 0x79, 0x43, 0x68, 0x61,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x43, 0x68, 0x61, 0x74, 0x22, 0x4e, 0x0a, 0x0a, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x61, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x22, 0x4e, 0x0a, 0x0a, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6c, 0x61, 0x6d,
	0x70, 0x6f, 0x72, 0x74, 0x22, 0x1e, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0x76, 0x0a, 0x11, 0x43, 0x68, 0x69,
	0x74, 0x74, 0x79, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c,
	0x0a, 0x08, 0x4a, 0x6f, 0x69, 0x6e, 0x43, 0x68, 0x61, 0x74, 0x12, 0x10, 0x2e, 0x43, 0x68, 0x61,
	0x74, 0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x0a, 0x2e, 0x43,
	0x68, 0x61, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x28, 0x01, 0x12, 0x33, 0x0a, 0x07,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x12, 0x10, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x2e, 0x46,
	0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x10, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x2e, 0x46, 0x72, 0x6f, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x00, 0x28, 0x01, 0x30,
	0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x43, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_Chat_ChittyChatService_proto_rawDescOnce sync.Once
	file_Chat_ChittyChatService_proto_rawDescData = file_Chat_ChittyChatService_proto_rawDesc
)

func file_Chat_ChittyChatService_proto_rawDescGZIP() []byte {
	file_Chat_ChittyChatService_proto_rawDescOnce.Do(func() {
		file_Chat_ChittyChatService_proto_rawDescData = protoimpl.X.CompressGZIP(file_Chat_ChittyChatService_proto_rawDescData)
	})
	return file_Chat_ChittyChatService_proto_rawDescData
}

var file_Chat_ChittyChatService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_Chat_ChittyChatService_proto_goTypes = []interface{}{
	(*FromClient)(nil), // 0: Chat.FromClient
	(*FromServer)(nil), // 1: Chat.FromServer
	(*UserName)(nil),   // 2: Chat.UserName
	(*User)(nil),       // 3: Chat.User
	(*Empty)(nil),      // 4: Chat.Empty
}
var file_Chat_ChittyChatService_proto_depIdxs = []int32{
	0, // 0: Chat.ChittyChatService.JoinChat:input_type -> Chat.FromClient
	0, // 1: Chat.ChittyChatService.Publish:input_type -> Chat.FromClient
	3, // 2: Chat.ChittyChatService.JoinChat:output_type -> Chat.User
	1, // 3: Chat.ChittyChatService.Publish:output_type -> Chat.FromServer
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Chat_ChittyChatService_proto_init() }
func file_Chat_ChittyChatService_proto_init() {
	if File_Chat_ChittyChatService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Chat_ChittyChatService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FromClient); i {
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
		file_Chat_ChittyChatService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FromServer); i {
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
		file_Chat_ChittyChatService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserName); i {
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
		file_Chat_ChittyChatService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_Chat_ChittyChatService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_Chat_ChittyChatService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Chat_ChittyChatService_proto_goTypes,
		DependencyIndexes: file_Chat_ChittyChatService_proto_depIdxs,
		MessageInfos:      file_Chat_ChittyChatService_proto_msgTypes,
	}.Build()
	File_Chat_ChittyChatService_proto = out.File
	file_Chat_ChittyChatService_proto_rawDesc = nil
	file_Chat_ChittyChatService_proto_goTypes = nil
	file_Chat_ChittyChatService_proto_depIdxs = nil
}
