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
// 	protoc        v3.19.4
// source: pipeline.proto

package Astounding

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PipeOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the Options needed for constructing the Pipeline :-)
	//
	// If not blank, indicates that a topic should get created.
	// TO DO - while breaking changes still tentatively allowed,
	//   see if can use bool, and get Topic Name from Message Name
	//   (when true).
	GenPubsubTopic bool `protobuf:"varint,1,opt,name=gen_pubsub_topic,json=genPubsubTopic,proto3" json:"gen_pubsub_topic,omitempty"`
}

func (x *PipeOptions) Reset() {
	*x = PipeOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pipeline_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PipeOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PipeOptions) ProtoMessage() {}

func (x *PipeOptions) ProtoReflect() protoreflect.Message {
	mi := &file_pipeline_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PipeOptions.ProtoReflect.Descriptor instead.
func (*PipeOptions) Descriptor() ([]byte, []int) {
	return file_pipeline_proto_rawDescGZIP(), []int{0}
}

func (x *PipeOptions) GetGenPubsubTopic() bool {
	if x != nil {
		return x.GenPubsubTopic
	}
	return false
}

var file_pipeline_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*PipeOptions)(nil),
		Field:         1090,
		Name:          "gen_pipes.pipe_opts",
		Tag:           "bytes,1090,opt,name=pipe_opts",
		Filename:      "pipeline.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// Pipeline message schema generation options.
	//
	// The field number is a globally unique id for this option, assigned by
	// protobuf-global-extension-registry@google.com
	//
	// optional gen_pipes.PipeOptions pipe_opts = 1090;
	E_PipeOpts = &file_pipeline_proto_extTypes[0]
)

var File_pipeline_proto protoreflect.FileDescriptor

var file_pipeline_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x09, 0x67, 0x65, 0x6e, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x73, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a,
	0x0b, 0x50, 0x69, 0x70, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x28, 0x0a, 0x10,
	0x67, 0x65, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x73, 0x75, 0x62, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x67, 0x65, 0x6e, 0x50, 0x75, 0x62, 0x73, 0x75,
	0x62, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x3a, 0x55, 0x0a, 0x09, 0x70, 0x69, 0x70, 0x65, 0x5f, 0x6f,
	0x70, 0x74, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xc2, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x67, 0x65,
	0x6e, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x73, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x08, 0x70, 0x69, 0x70, 0x65, 0x4f, 0x70, 0x74, 0x73, 0x42, 0x1f, 0x5a,
	0x1d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x62, 0x50, 0x69,
	0x70, 0x65, 0x73, 0x2f, 0x41, 0x73, 0x74, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pipeline_proto_rawDescOnce sync.Once
	file_pipeline_proto_rawDescData = file_pipeline_proto_rawDesc
)

func file_pipeline_proto_rawDescGZIP() []byte {
	file_pipeline_proto_rawDescOnce.Do(func() {
		file_pipeline_proto_rawDescData = protoimpl.X.CompressGZIP(file_pipeline_proto_rawDescData)
	})
	return file_pipeline_proto_rawDescData
}

var file_pipeline_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pipeline_proto_goTypes = []interface{}{
	(*PipeOptions)(nil),                 // 0: gen_pipes.PipeOptions
	(*descriptorpb.MessageOptions)(nil), // 1: google.protobuf.MessageOptions
}
var file_pipeline_proto_depIdxs = []int32{
	1, // 0: gen_pipes.pipe_opts:extendee -> google.protobuf.MessageOptions
	0, // 1: gen_pipes.pipe_opts:type_name -> gen_pipes.PipeOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pipeline_proto_init() }
func file_pipeline_proto_init() {
	if File_pipeline_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pipeline_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PipeOptions); i {
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
			RawDescriptor: file_pipeline_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_pipeline_proto_goTypes,
		DependencyIndexes: file_pipeline_proto_depIdxs,
		MessageInfos:      file_pipeline_proto_msgTypes,
		ExtensionInfos:    file_pipeline_proto_extTypes,
	}.Build()
	File_pipeline_proto = out.File
	file_pipeline_proto_rawDesc = nil
	file_pipeline_proto_goTypes = nil
	file_pipeline_proto_depIdxs = nil
}
