// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.17.3
// source: internal/pkg/entity/user/member.proto

package user

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

type Kind int32

const (
	Kind_UNKNOWN Kind = 0
	Kind_OWNER   Kind = 1
	Kind_MEMBER  Kind = 2
)

// Enum value maps for Kind.
var (
	Kind_name = map[int32]string{
		0: "UNKNOWN",
		1: "OWNER",
		2: "MEMBER",
	}
	Kind_value = map[string]int32{
		"UNKNOWN": 0,
		"OWNER":   1,
		"MEMBER":  2,
	}
)

func (x Kind) Enum() *Kind {
	p := new(Kind)
	*p = x
	return p
}

func (x Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_pkg_entity_user_member_proto_enumTypes[0].Descriptor()
}

func (Kind) Type() protoreflect.EnumType {
	return &file_internal_pkg_entity_user_member_proto_enumTypes[0]
}

func (x Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Kind.Descriptor instead.
func (Kind) EnumDescriptor() ([]byte, []int) {
	return file_internal_pkg_entity_user_member_proto_rawDescGZIP(), []int{0}
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Kind    Kind   `protobuf:"varint,2,opt,name=kind,proto3,enum=user.Kind" json:"kind,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Name    string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Value   int64  `protobuf:"varint,5,opt,name=value,proto3" json:"value,omitempty"`
	Balance int64  `protobuf:"varint,6,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pkg_entity_user_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pkg_entity_user_member_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_internal_pkg_entity_user_member_proto_rawDescGZIP(), []int{0}
}

func (x *Member) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Member) GetKind() Kind {
	if x != nil {
		return x.Kind
	}
	return Kind_UNKNOWN
}

func (x *Member) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Member) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Member) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Member) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

var File_internal_pkg_entity_user_member_proto protoreflect.FileDescriptor

var file_internal_pkg_entity_user_member_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x92, 0x01,
	0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4b, 0x69,
	0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x2a, 0x2a, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x57, 0x4e, 0x45, 0x52,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x10, 0x02, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pkg_entity_user_member_proto_rawDescOnce sync.Once
	file_internal_pkg_entity_user_member_proto_rawDescData = file_internal_pkg_entity_user_member_proto_rawDesc
)

func file_internal_pkg_entity_user_member_proto_rawDescGZIP() []byte {
	file_internal_pkg_entity_user_member_proto_rawDescOnce.Do(func() {
		file_internal_pkg_entity_user_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pkg_entity_user_member_proto_rawDescData)
	})
	return file_internal_pkg_entity_user_member_proto_rawDescData
}

var file_internal_pkg_entity_user_member_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_pkg_entity_user_member_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_pkg_entity_user_member_proto_goTypes = []interface{}{
	(Kind)(0),      // 0: user.Kind
	(*Member)(nil), // 1: user.Member
}
var file_internal_pkg_entity_user_member_proto_depIdxs = []int32{
	0, // 0: user.Member.kind:type_name -> user.Kind
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_pkg_entity_user_member_proto_init() }
func file_internal_pkg_entity_user_member_proto_init() {
	if File_internal_pkg_entity_user_member_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pkg_entity_user_member_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
			RawDescriptor: file_internal_pkg_entity_user_member_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pkg_entity_user_member_proto_goTypes,
		DependencyIndexes: file_internal_pkg_entity_user_member_proto_depIdxs,
		EnumInfos:         file_internal_pkg_entity_user_member_proto_enumTypes,
		MessageInfos:      file_internal_pkg_entity_user_member_proto_msgTypes,
	}.Build()
	File_internal_pkg_entity_user_member_proto = out.File
	file_internal_pkg_entity_user_member_proto_rawDesc = nil
	file_internal_pkg_entity_user_member_proto_goTypes = nil
	file_internal_pkg_entity_user_member_proto_depIdxs = nil
}
