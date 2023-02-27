// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: checkout.proto

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

type CheckoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemsID []int32 `protobuf:"varint,1,rep,packed,name=itemsID,proto3" json:"itemsID,omitempty"`
}

func (x *CheckoutRequest) Reset() {
	*x = CheckoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutRequest) ProtoMessage() {}

func (x *CheckoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutRequest.ProtoReflect.Descriptor instead.
func (*CheckoutRequest) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{0}
}

func (x *CheckoutRequest) GetItemsID() []int32 {
	if x != nil {
		return x.ItemsID
	}
	return nil
}

type CheckoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalAmount int32 `protobuf:"varint,1,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"`
}

func (x *CheckoutResponse) Reset() {
	*x = CheckoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkout_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutResponse) ProtoMessage() {}

func (x *CheckoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checkout_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutResponse.ProtoReflect.Descriptor instead.
func (*CheckoutResponse) Descriptor() ([]byte, []int) {
	return file_checkout_proto_rawDescGZIP(), []int{1}
}

func (x *CheckoutResponse) GetTotalAmount() int32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

var File_checkout_proto protoreflect.FileDescriptor

var file_checkout_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x49, 0x44, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x07,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x49, 0x44, 0x22, 0x34, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x51, 0x0a,
	0x08, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x45, 0x0a, 0x0a, 0x44, 0x6f, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x12, 0x19, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x2e, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x6d, 0x6f, 0x72, 0x65, 0x6c, 0x6c, 0x69, 0x39, 0x32, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x2d, 0x74,
	0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_checkout_proto_rawDescOnce sync.Once
	file_checkout_proto_rawDescData = file_checkout_proto_rawDesc
)

func file_checkout_proto_rawDescGZIP() []byte {
	file_checkout_proto_rawDescOnce.Do(func() {
		file_checkout_proto_rawDescData = protoimpl.X.CompressGZIP(file_checkout_proto_rawDescData)
	})
	return file_checkout_proto_rawDescData
}

var file_checkout_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_checkout_proto_goTypes = []interface{}{
	(*CheckoutRequest)(nil),  // 0: checkout.CheckoutRequest
	(*CheckoutResponse)(nil), // 1: checkout.CheckoutResponse
}
var file_checkout_proto_depIdxs = []int32{
	0, // 0: checkout.Checkout.DoCheckout:input_type -> checkout.CheckoutRequest
	1, // 1: checkout.Checkout.DoCheckout:output_type -> checkout.CheckoutResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_checkout_proto_init() }
func file_checkout_proto_init() {
	if File_checkout_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_checkout_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutRequest); i {
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
		file_checkout_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutResponse); i {
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
			RawDescriptor: file_checkout_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_checkout_proto_goTypes,
		DependencyIndexes: file_checkout_proto_depIdxs,
		MessageInfos:      file_checkout_proto_msgTypes,
	}.Build()
	File_checkout_proto = out.File
	file_checkout_proto_rawDesc = nil
	file_checkout_proto_goTypes = nil
	file_checkout_proto_depIdxs = nil
}
