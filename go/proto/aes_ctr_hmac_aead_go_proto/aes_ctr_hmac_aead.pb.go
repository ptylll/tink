// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: third_party/tink/proto/aes_ctr_hmac_aead.proto

package aes_ctr_hmac_aead_go_proto

import (
	proto "github.com/golang/protobuf/proto"
	aes_ctr_go_proto "github.com/google/tink/proto/aes_ctr_go_proto"
	hmac_go_proto "github.com/google/tink/proto/hmac_go_proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type AesCtrHmacAeadKeyFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AesCtrKeyFormat *aes_ctr_go_proto.AesCtrKeyFormat `protobuf:"bytes,1,opt,name=aes_ctr_key_format,json=aesCtrKeyFormat,proto3" json:"aes_ctr_key_format,omitempty"`
	HmacKeyFormat   *hmac_go_proto.HmacKeyFormat      `protobuf:"bytes,2,opt,name=hmac_key_format,json=hmacKeyFormat,proto3" json:"hmac_key_format,omitempty"`
}

func (x *AesCtrHmacAeadKeyFormat) Reset() {
	*x = AesCtrHmacAeadKeyFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_aes_ctr_hmac_aead_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AesCtrHmacAeadKeyFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AesCtrHmacAeadKeyFormat) ProtoMessage() {}

func (x *AesCtrHmacAeadKeyFormat) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_aes_ctr_hmac_aead_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AesCtrHmacAeadKeyFormat.ProtoReflect.Descriptor instead.
func (*AesCtrHmacAeadKeyFormat) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_aes_ctr_hmac_aead_proto_rawDescGZIP(), []int{0}
}

func (x *AesCtrHmacAeadKeyFormat) GetAesCtrKeyFormat() *aes_ctr_go_proto.AesCtrKeyFormat {
	if x != nil {
		return x.AesCtrKeyFormat
	}
	return nil
}

func (x *AesCtrHmacAeadKeyFormat) GetHmacKeyFormat() *hmac_go_proto.HmacKeyFormat {
	if x != nil {
		return x.HmacKeyFormat
	}
	return nil
}

// key_type: type.googleapis.com/google.crypto.tink.AesCtrHmacAeadKey
type AesCtrHmacAeadKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version   uint32                      `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	AesCtrKey *aes_ctr_go_proto.AesCtrKey `protobuf:"bytes,2,opt,name=aes_ctr_key,json=aesCtrKey,proto3" json:"aes_ctr_key,omitempty"`
	HmacKey   *hmac_go_proto.HmacKey      `protobuf:"bytes,3,opt,name=hmac_key,json=hmacKey,proto3" json:"hmac_key,omitempty"`
}

func (x *AesCtrHmacAeadKey) Reset() {
	*x = AesCtrHmacAeadKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_aes_ctr_hmac_aead_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AesCtrHmacAeadKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AesCtrHmacAeadKey) ProtoMessage() {}

func (x *AesCtrHmacAeadKey) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_aes_ctr_hmac_aead_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AesCtrHmacAeadKey.ProtoReflect.Descriptor instead.
func (*AesCtrHmacAeadKey) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_aes_ctr_hmac_aead_proto_rawDescGZIP(), []int{1}
}

func (x *AesCtrHmacAeadKey) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *AesCtrHmacAeadKey) GetAesCtrKey() *aes_ctr_go_proto.AesCtrKey {
	if x != nil {
		return x.AesCtrKey
	}
	return nil
}

func (x *AesCtrHmacAeadKey) GetHmacKey() *hmac_go_proto.HmacKey {
	if x != nil {
		return x.HmacKey
	}
	return nil
}

var File_third_party_tink_proto_aes_ctr_hmac_aead_proto protoreflect.FileDescriptor

var file_third_party_tink_proto_aes_ctr_hmac_aead_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x69,
	0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x65, 0x73, 0x5f, 0x63, 0x74, 0x72,
	0x5f, 0x68, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x65, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e,
	0x74, 0x69, 0x6e, 0x6b, 0x1a, 0x24, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x65, 0x73,
	0x5f, 0x63, 0x74, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x68, 0x6d, 0x61, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x01,
	0x0a, 0x17, 0x41, 0x65, 0x73, 0x43, 0x74, 0x72, 0x48, 0x6d, 0x61, 0x63, 0x41, 0x65, 0x61, 0x64,
	0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x50, 0x0a, 0x12, 0x61, 0x65, 0x73,
	0x5f, 0x63, 0x74, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x41, 0x65, 0x73, 0x43, 0x74,
	0x72, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x0f, 0x61, 0x65, 0x73, 0x43,
	0x74, 0x72, 0x4b, 0x65, 0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x49, 0x0a, 0x0f, 0x68,
	0x6d, 0x61, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x48, 0x6d, 0x61, 0x63, 0x4b, 0x65,
	0x79, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x0d, 0x68, 0x6d, 0x61, 0x63, 0x4b, 0x65, 0x79,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x22, 0xa4, 0x01, 0x0a, 0x11, 0x41, 0x65, 0x73, 0x43, 0x74,
	0x72, 0x48, 0x6d, 0x61, 0x63, 0x41, 0x65, 0x61, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3d, 0x0a, 0x0b, 0x61, 0x65, 0x73, 0x5f, 0x63, 0x74,
	0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b,
	0x2e, 0x41, 0x65, 0x73, 0x43, 0x74, 0x72, 0x4b, 0x65, 0x79, 0x52, 0x09, 0x61, 0x65, 0x73, 0x43,
	0x74, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x08, 0x68, 0x6d, 0x61, 0x63, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x48, 0x6d, 0x61,
	0x63, 0x4b, 0x65, 0x79, 0x52, 0x07, 0x68, 0x6d, 0x61, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x62, 0x0a,
	0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x74, 0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x65,
	0x73, 0x5f, 0x63, 0x74, 0x72, 0x5f, 0x68, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x65, 0x61, 0x64, 0x5f,
	0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x06, 0x54, 0x49, 0x4e, 0x4b, 0x50,
	0x42, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_third_party_tink_proto_aes_ctr_hmac_aead_proto_rawDescOnce sync.Once
	file_third_party_tink_proto_aes_ctr_hmac_aead_proto_rawDescData = file_third_party_tink_proto_aes_ctr_hmac_aead_proto_rawDesc
)

func file_third_party_tink_proto_aes_ctr_hmac_aead_proto_rawDescGZIP() []byte {
	file_third_party_tink_proto_aes_ctr_hmac_aead_proto_rawDescOnce.Do(func() {
		file_third_party_tink_proto_aes_ctr_hmac_aead_proto_rawDescData = protoimpl.X.CompressGZIP(file_third_party_tink_proto_aes_ctr_hmac_aead_proto_rawDescData)
	})
	return file_third_party_tink_proto_aes_ctr_hmac_aead_proto_rawDescData
}

var file_third_party_tink_proto_aes_ctr_hmac_aead_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_third_party_tink_proto_aes_ctr_hmac_aead_proto_goTypes = []interface{}{
	(*AesCtrHmacAeadKeyFormat)(nil),          // 0: google.crypto.tink.AesCtrHmacAeadKeyFormat
	(*AesCtrHmacAeadKey)(nil),                // 1: google.crypto.tink.AesCtrHmacAeadKey
	(*aes_ctr_go_proto.AesCtrKeyFormat)(nil), // 2: google.crypto.tink.AesCtrKeyFormat
	(*hmac_go_proto.HmacKeyFormat)(nil),      // 3: google.crypto.tink.HmacKeyFormat
	(*aes_ctr_go_proto.AesCtrKey)(nil),       // 4: google.crypto.tink.AesCtrKey
	(*hmac_go_proto.HmacKey)(nil),            // 5: google.crypto.tink.HmacKey
}
var file_third_party_tink_proto_aes_ctr_hmac_aead_proto_depIdxs = []int32{
	2, // 0: google.crypto.tink.AesCtrHmacAeadKeyFormat.aes_ctr_key_format:type_name -> google.crypto.tink.AesCtrKeyFormat
	3, // 1: google.crypto.tink.AesCtrHmacAeadKeyFormat.hmac_key_format:type_name -> google.crypto.tink.HmacKeyFormat
	4, // 2: google.crypto.tink.AesCtrHmacAeadKey.aes_ctr_key:type_name -> google.crypto.tink.AesCtrKey
	5, // 3: google.crypto.tink.AesCtrHmacAeadKey.hmac_key:type_name -> google.crypto.tink.HmacKey
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_third_party_tink_proto_aes_ctr_hmac_aead_proto_init() }
func file_third_party_tink_proto_aes_ctr_hmac_aead_proto_init() {
	if File_third_party_tink_proto_aes_ctr_hmac_aead_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_third_party_tink_proto_aes_ctr_hmac_aead_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AesCtrHmacAeadKeyFormat); i {
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
		file_third_party_tink_proto_aes_ctr_hmac_aead_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AesCtrHmacAeadKey); i {
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
			RawDescriptor: file_third_party_tink_proto_aes_ctr_hmac_aead_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_third_party_tink_proto_aes_ctr_hmac_aead_proto_goTypes,
		DependencyIndexes: file_third_party_tink_proto_aes_ctr_hmac_aead_proto_depIdxs,
		MessageInfos:      file_third_party_tink_proto_aes_ctr_hmac_aead_proto_msgTypes,
	}.Build()
	File_third_party_tink_proto_aes_ctr_hmac_aead_proto = out.File
	file_third_party_tink_proto_aes_ctr_hmac_aead_proto_rawDesc = nil
	file_third_party_tink_proto_aes_ctr_hmac_aead_proto_goTypes = nil
	file_third_party_tink_proto_aes_ctr_hmac_aead_proto_depIdxs = nil
}
