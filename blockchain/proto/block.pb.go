// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: block.proto

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

type GetLatestBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetLatestBlockRequest) Reset() {
	*x = GetLatestBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestBlockRequest) ProtoMessage() {}

func (x *GetLatestBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestBlockRequest.ProtoReflect.Descriptor instead.
func (*GetLatestBlockRequest) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{0}
}

type GetLatestBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId string `protobuf:"bytes,1,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Block   string `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *GetLatestBlockResponse) Reset() {
	*x = GetLatestBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_block_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestBlockResponse) ProtoMessage() {}

func (x *GetLatestBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_block_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestBlockResponse.ProtoReflect.Descriptor instead.
func (*GetLatestBlockResponse) Descriptor() ([]byte, []int) {
	return file_block_proto_rawDescGZIP(), []int{1}
}

func (x *GetLatestBlockResponse) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *GetLatestBlockResponse) GetBlock() string {
	if x != nil {
		return x.Block
	}
	return ""
}

var File_block_proto protoreflect.FileDescriptor

var file_block_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74,
	0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x49, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x32, 0x69, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x21, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74,
	0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x66, 0x69, 0x75, 0x6c, 0x68, 0x63, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2d, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2d,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_block_proto_rawDescOnce sync.Once
	file_block_proto_rawDescData = file_block_proto_rawDesc
)

func file_block_proto_rawDescGZIP() []byte {
	file_block_proto_rawDescOnce.Do(func() {
		file_block_proto_rawDescData = protoimpl.X.CompressGZIP(file_block_proto_rawDescData)
	})
	return file_block_proto_rawDescData
}

var file_block_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_block_proto_goTypes = []interface{}{
	(*GetLatestBlockRequest)(nil),  // 0: blockchain.GetLatestBlockRequest
	(*GetLatestBlockResponse)(nil), // 1: blockchain.GetLatestBlockResponse
}
var file_block_proto_depIdxs = []int32{
	0, // 0: blockchain.GetLatestBlockService.Block:input_type -> blockchain.GetLatestBlockRequest
	1, // 1: blockchain.GetLatestBlockService.Block:output_type -> blockchain.GetLatestBlockResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_block_proto_init() }
func file_block_proto_init() {
	if File_block_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_block_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestBlockRequest); i {
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
		file_block_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestBlockResponse); i {
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
			RawDescriptor: file_block_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_block_proto_goTypes,
		DependencyIndexes: file_block_proto_depIdxs,
		MessageInfos:      file_block_proto_msgTypes,
	}.Build()
	File_block_proto = out.File
	file_block_proto_rawDesc = nil
	file_block_proto_goTypes = nil
	file_block_proto_depIdxs = nil
}
