//*
// Copyright 2022 buexplain@qq.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: limitResp.proto

package protocol

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

// business向worker请求，返回网关中的限流配置的真实情况
type LimitResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// worker会原样回传给business
	CtxData []byte `protobuf:"bytes,1,opt,name=ctxData,proto3" json:"ctxData,omitempty"`
	// 每个worker的配置情况
	Items []*LimitCountItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *LimitResp) Reset() {
	*x = LimitResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limitResp_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitResp) ProtoMessage() {}

func (x *LimitResp) ProtoReflect() protoreflect.Message {
	mi := &file_limitResp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitResp.ProtoReflect.Descriptor instead.
func (*LimitResp) Descriptor() ([]byte, []int) {
	return file_limitResp_proto_rawDescGZIP(), []int{0}
}

func (x *LimitResp) GetCtxData() []byte {
	if x != nil {
		return x.CtxData
	}
	return nil
}

func (x *LimitResp) GetItems() []*LimitCountItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type LimitCountItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// workerId集合
	WorkerIds []int32 `protobuf:"varint,1,rep,packed,name=workerIds,proto3" json:"workerIds,omitempty"`
	// 限流大小
	Num int32 `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
}

func (x *LimitCountItem) Reset() {
	*x = LimitCountItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_limitResp_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LimitCountItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LimitCountItem) ProtoMessage() {}

func (x *LimitCountItem) ProtoReflect() protoreflect.Message {
	mi := &file_limitResp_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LimitCountItem.ProtoReflect.Descriptor instead.
func (*LimitCountItem) Descriptor() ([]byte, []int) {
	return file_limitResp_proto_rawDescGZIP(), []int{1}
}

func (x *LimitCountItem) GetWorkerIds() []int32 {
	if x != nil {
		return x.WorkerIds
	}
	return nil
}

func (x *LimitCountItem) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

var File_limitResp_proto protoreflect.FileDescriptor

var file_limitResp_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x6e, 0x65, 0x74, 0x73, 0x76, 0x72, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x5d, 0x0a, 0x09, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x74, 0x78, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x07, 0x63, 0x74, 0x78, 0x44, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x65, 0x74, 0x73,
	0x76, 0x72, 0x2e, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x22, 0x40, 0x0a, 0x0e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49,
	0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6e, 0x75, 0x6d, 0x42, 0x17, 0x5a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_limitResp_proto_rawDescOnce sync.Once
	file_limitResp_proto_rawDescData = file_limitResp_proto_rawDesc
)

func file_limitResp_proto_rawDescGZIP() []byte {
	file_limitResp_proto_rawDescOnce.Do(func() {
		file_limitResp_proto_rawDescData = protoimpl.X.CompressGZIP(file_limitResp_proto_rawDescData)
	})
	return file_limitResp_proto_rawDescData
}

var file_limitResp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_limitResp_proto_goTypes = []interface{}{
	(*LimitResp)(nil),      // 0: netsvr.limitResp.LimitResp
	(*LimitCountItem)(nil), // 1: netsvr.limitResp.LimitCountItem
}
var file_limitResp_proto_depIdxs = []int32{
	1, // 0: netsvr.limitResp.LimitResp.items:type_name -> netsvr.limitResp.LimitCountItem
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_limitResp_proto_init() }
func file_limitResp_proto_init() {
	if File_limitResp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_limitResp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitResp); i {
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
		file_limitResp_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LimitCountItem); i {
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
			RawDescriptor: file_limitResp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_limitResp_proto_goTypes,
		DependencyIndexes: file_limitResp_proto_depIdxs,
		MessageInfos:      file_limitResp_proto_msgTypes,
	}.Build()
	File_limitResp_proto = out.File
	file_limitResp_proto_rawDesc = nil
	file_limitResp_proto_goTypes = nil
	file_limitResp_proto_depIdxs = nil
}
