// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: coffeeservice/v1/coffeeservice.proto

package coffeeservicev1

import (
	_ "github.com/datacircus/sideshow/highwire/gen/go/buf/validate"
	v1 "github.com/datacircus/sideshow/highwire/gen/go/sideshow/coffeeco/v1"
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

type CoffeeOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *v1.Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *CoffeeOrderRequest) Reset() {
	*x = CoffeeOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coffeeservice_v1_coffeeservice_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoffeeOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoffeeOrderRequest) ProtoMessage() {}

func (x *CoffeeOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coffeeservice_v1_coffeeservice_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoffeeOrderRequest.ProtoReflect.Descriptor instead.
func (*CoffeeOrderRequest) Descriptor() ([]byte, []int) {
	return file_coffeeservice_v1_coffeeservice_proto_rawDescGZIP(), []int{0}
}

func (x *CoffeeOrderRequest) GetOrder() *v1.Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type CoffeeOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CoffeeOrderResponse) Reset() {
	*x = CoffeeOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coffeeservice_v1_coffeeservice_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoffeeOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoffeeOrderResponse) ProtoMessage() {}

func (x *CoffeeOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coffeeservice_v1_coffeeservice_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoffeeOrderResponse.ProtoReflect.Descriptor instead.
func (*CoffeeOrderResponse) Descriptor() ([]byte, []int) {
	return file_coffeeservice_v1_coffeeservice_proto_rawDescGZIP(), []int{1}
}

func (x *CoffeeOrderResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_coffeeservice_v1_coffeeservice_proto protoreflect.FileDescriptor

var file_coffeeservice_v1_coffeeservice_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x73, 0x69, 0x64, 0x65, 0x73, 0x68,
	0x6f, 0x77, 0x2f, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x12, 0x43, 0x6f, 0x66, 0x66, 0x65,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73,
	0x69, 0x64, 0x65, 0x73, 0x68, 0x6f, 0x77, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x13, 0x43, 0x6f, 0x66, 0x66,
	0x65, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x6d, 0x0a, 0x0d, 0x43,
	0x6f, 0x66, 0x66, 0x65, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x0b,
	0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x24, 0x2e, 0x63, 0x6f,
	0x66, 0x66, 0x65, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x66, 0x66, 0x65, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xdc, 0x01, 0x0a, 0x14, 0x63,
	0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x42, 0x12, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x69, 0x72, 0x63, 0x75, 0x73,
	0x2f, 0x73, 0x69, 0x64, 0x65, 0x73, 0x68, 0x6f, 0x77, 0x2f, 0x68, 0x69, 0x67, 0x68, 0x77, 0x69,
	0x72, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x66, 0x66, 0x65,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58, 0x58,
	0xaa, 0x02, 0x10, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x43, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_coffeeservice_v1_coffeeservice_proto_rawDescOnce sync.Once
	file_coffeeservice_v1_coffeeservice_proto_rawDescData = file_coffeeservice_v1_coffeeservice_proto_rawDesc
)

func file_coffeeservice_v1_coffeeservice_proto_rawDescGZIP() []byte {
	file_coffeeservice_v1_coffeeservice_proto_rawDescOnce.Do(func() {
		file_coffeeservice_v1_coffeeservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_coffeeservice_v1_coffeeservice_proto_rawDescData)
	})
	return file_coffeeservice_v1_coffeeservice_proto_rawDescData
}

var file_coffeeservice_v1_coffeeservice_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_coffeeservice_v1_coffeeservice_proto_goTypes = []interface{}{
	(*CoffeeOrderRequest)(nil),  // 0: coffeeservice.v1.CoffeeOrderRequest
	(*CoffeeOrderResponse)(nil), // 1: coffeeservice.v1.CoffeeOrderResponse
	(*v1.Order)(nil),            // 2: sideshow.coffeeco.v1.Order
}
var file_coffeeservice_v1_coffeeservice_proto_depIdxs = []int32{
	2, // 0: coffeeservice.v1.CoffeeOrderRequest.order:type_name -> sideshow.coffeeco.v1.Order
	0, // 1: coffeeservice.v1.CoffeeService.CoffeeOrder:input_type -> coffeeservice.v1.CoffeeOrderRequest
	1, // 2: coffeeservice.v1.CoffeeService.CoffeeOrder:output_type -> coffeeservice.v1.CoffeeOrderResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_coffeeservice_v1_coffeeservice_proto_init() }
func file_coffeeservice_v1_coffeeservice_proto_init() {
	if File_coffeeservice_v1_coffeeservice_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coffeeservice_v1_coffeeservice_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoffeeOrderRequest); i {
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
		file_coffeeservice_v1_coffeeservice_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoffeeOrderResponse); i {
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
			RawDescriptor: file_coffeeservice_v1_coffeeservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coffeeservice_v1_coffeeservice_proto_goTypes,
		DependencyIndexes: file_coffeeservice_v1_coffeeservice_proto_depIdxs,
		MessageInfos:      file_coffeeservice_v1_coffeeservice_proto_msgTypes,
	}.Build()
	File_coffeeservice_v1_coffeeservice_proto = out.File
	file_coffeeservice_v1_coffeeservice_proto_rawDesc = nil
	file_coffeeservice_v1_coffeeservice_proto_goTypes = nil
	file_coffeeservice_v1_coffeeservice_proto_depIdxs = nil
}
