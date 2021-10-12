// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: cvcio/classification/accounts.proto

package classification

import (
	common "github.com/cvcio/proto/cvcio/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type AccountClass int32

const (
	AccountClass_ROLE_UNKNOWN    AccountClass = 0
	AccountClass_ROLE_ACTIVE     AccountClass = 1
	AccountClass_ROLE_INFLUENCER AccountClass = 2
	AccountClass_ROLE_NEW        AccountClass = 3
	AccountClass_ROLE_AMPLIFIER  AccountClass = 4
)

// Enum value maps for AccountClass.
var (
	AccountClass_name = map[int32]string{
		0: "ROLE_UNKNOWN",
		1: "ROLE_ACTIVE",
		2: "ROLE_INFLUENCER",
		3: "ROLE_NEW",
		4: "ROLE_AMPLIFIER",
	}
	AccountClass_value = map[string]int32{
		"ROLE_UNKNOWN":    0,
		"ROLE_ACTIVE":     1,
		"ROLE_INFLUENCER": 2,
		"ROLE_NEW":        3,
		"ROLE_AMPLIFIER":  4,
	}
)

func (x AccountClass) Enum() *AccountClass {
	p := new(AccountClass)
	*p = x
	return p
}

func (x AccountClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AccountClass) Descriptor() protoreflect.EnumDescriptor {
	return file_cvcio_classification_accounts_proto_enumTypes[0].Descriptor()
}

func (AccountClass) Type() protoreflect.EnumType {
	return &file_cvcio_classification_accounts_proto_enumTypes[0]
}

func (x AccountClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AccountClass.Descriptor instead.
func (AccountClass) EnumDescriptor() ([]byte, []int) {
	return file_cvcio_classification_accounts_proto_rawDescGZIP(), []int{0}
}

type TwitterAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Followers int64        `protobuf:"varint,1,opt,name=followers,proto3" json:"followers,omitempty"`
	Friends   int64        `protobuf:"varint,2,opt,name=friends,proto3" json:"friends,omitempty"`
	Statuses  int64        `protobuf:"varint,3,opt,name=statuses,proto3" json:"statuses,omitempty"`
	Favorites int64        `protobuf:"varint,4,opt,name=favorites,proto3" json:"favorites,omitempty"`
	Lists     int64        `protobuf:"varint,5,opt,name=lists,proto3" json:"lists,omitempty"`
	Ffr       float64      `protobuf:"fixed64,6,opt,name=ffr,proto3" json:"ffr,omitempty"`
	Stfv      float64      `protobuf:"fixed64,7,opt,name=stfv,proto3" json:"stfv,omitempty"`
	Fstfv     float64      `protobuf:"fixed64,8,opt,name=fstfv,proto3" json:"fstfv,omitempty"`
	Dates     float64      `protobuf:"fixed64,9,opt,name=dates,proto3" json:"dates,omitempty"`
	Actions   float64      `protobuf:"fixed64,10,opt,name=actions,proto3" json:"actions,omitempty"`
	Meta      *common.Meta `protobuf:"bytes,11,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *TwitterAccount) Reset() {
	*x = TwitterAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cvcio_classification_accounts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TwitterAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TwitterAccount) ProtoMessage() {}

func (x *TwitterAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cvcio_classification_accounts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TwitterAccount.ProtoReflect.Descriptor instead.
func (*TwitterAccount) Descriptor() ([]byte, []int) {
	return file_cvcio_classification_accounts_proto_rawDescGZIP(), []int{0}
}

func (x *TwitterAccount) GetFollowers() int64 {
	if x != nil {
		return x.Followers
	}
	return 0
}

func (x *TwitterAccount) GetFriends() int64 {
	if x != nil {
		return x.Friends
	}
	return 0
}

func (x *TwitterAccount) GetStatuses() int64 {
	if x != nil {
		return x.Statuses
	}
	return 0
}

func (x *TwitterAccount) GetFavorites() int64 {
	if x != nil {
		return x.Favorites
	}
	return 0
}

func (x *TwitterAccount) GetLists() int64 {
	if x != nil {
		return x.Lists
	}
	return 0
}

func (x *TwitterAccount) GetFfr() float64 {
	if x != nil {
		return x.Ffr
	}
	return 0
}

func (x *TwitterAccount) GetStfv() float64 {
	if x != nil {
		return x.Stfv
	}
	return 0
}

func (x *TwitterAccount) GetFstfv() float64 {
	if x != nil {
		return x.Fstfv
	}
	return 0
}

func (x *TwitterAccount) GetDates() float64 {
	if x != nil {
		return x.Dates
	}
	return 0
}

func (x *TwitterAccount) GetActions() float64 {
	if x != nil {
		return x.Actions
	}
	return 0
}

func (x *TwitterAccount) GetMeta() *common.Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type ResponseAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Predictions []*common.Prediction `protobuf:"bytes,1,rep,name=predictions,proto3" json:"predictions,omitempty"`
}

func (x *ResponseAccount) Reset() {
	*x = ResponseAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cvcio_classification_accounts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseAccount) ProtoMessage() {}

func (x *ResponseAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cvcio_classification_accounts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseAccount.ProtoReflect.Descriptor instead.
func (*ResponseAccount) Descriptor() ([]byte, []int) {
	return file_cvcio_classification_accounts_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseAccount) GetPredictions() []*common.Prediction {
	if x != nil {
		return x.Predictions
	}
	return nil
}

var File_cvcio_classification_accounts_proto protoreflect.FileDescriptor

var file_cvcio_classification_accounts_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x76, 0x63, 0x69, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x76, 0x63, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x76,
	0x63, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x02, 0x0a,
	0x0e, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x66, 0x72, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x66, 0x66, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x66,
	0x76, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x73, 0x74, 0x66, 0x76, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x73, 0x74, 0x66, 0x76, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x66, 0x73,
	0x74, 0x66, 0x76, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x61, 0x74, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x05, 0x64, 0x61, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x47, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0x68,
	0x0a, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x10,
	0x0a, 0x0c, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10,
	0x01, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x49, 0x4e, 0x46, 0x4c, 0x55, 0x45,
	0x4e, 0x43, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x4e,
	0x45, 0x57, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x41, 0x4d, 0x50,
	0x4c, 0x49, 0x46, 0x49, 0x45, 0x52, 0x10, 0x04, 0x32, 0x8a, 0x01, 0x0a, 0x0e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x78, 0x0a, 0x14, 0x44,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x1a, 0x1f, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76,
	0x33, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x74, 0x77, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x3a, 0x01, 0x2a, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x76, 0x63, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x76, 0x63, 0x69, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cvcio_classification_accounts_proto_rawDescOnce sync.Once
	file_cvcio_classification_accounts_proto_rawDescData = file_cvcio_classification_accounts_proto_rawDesc
)

func file_cvcio_classification_accounts_proto_rawDescGZIP() []byte {
	file_cvcio_classification_accounts_proto_rawDescOnce.Do(func() {
		file_cvcio_classification_accounts_proto_rawDescData = protoimpl.X.CompressGZIP(file_cvcio_classification_accounts_proto_rawDescData)
	})
	return file_cvcio_classification_accounts_proto_rawDescData
}

var file_cvcio_classification_accounts_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cvcio_classification_accounts_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cvcio_classification_accounts_proto_goTypes = []interface{}{
	(AccountClass)(0),         // 0: classification.AccountClass
	(*TwitterAccount)(nil),    // 1: classification.TwitterAccount
	(*ResponseAccount)(nil),   // 2: classification.ResponseAccount
	(*common.Meta)(nil),       // 3: common.Meta
	(*common.Prediction)(nil), // 4: common.Prediction
}
var file_cvcio_classification_accounts_proto_depIdxs = []int32{
	3, // 0: classification.TwitterAccount.meta:type_name -> common.Meta
	4, // 1: classification.ResponseAccount.predictions:type_name -> common.Prediction
	1, // 2: classification.AccountService.DetectTwitterAccount:input_type -> classification.TwitterAccount
	2, // 3: classification.AccountService.DetectTwitterAccount:output_type -> classification.ResponseAccount
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cvcio_classification_accounts_proto_init() }
func file_cvcio_classification_accounts_proto_init() {
	if File_cvcio_classification_accounts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cvcio_classification_accounts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TwitterAccount); i {
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
		file_cvcio_classification_accounts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseAccount); i {
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
			RawDescriptor: file_cvcio_classification_accounts_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cvcio_classification_accounts_proto_goTypes,
		DependencyIndexes: file_cvcio_classification_accounts_proto_depIdxs,
		EnumInfos:         file_cvcio_classification_accounts_proto_enumTypes,
		MessageInfos:      file_cvcio_classification_accounts_proto_msgTypes,
	}.Build()
	File_cvcio_classification_accounts_proto = out.File
	file_cvcio_classification_accounts_proto_rawDesc = nil
	file_cvcio_classification_accounts_proto_goTypes = nil
	file_cvcio_classification_accounts_proto_depIdxs = nil
}
