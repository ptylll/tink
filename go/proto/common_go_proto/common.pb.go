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

// Definitions for common cryptographic enum types.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0-devel
// 	protoc        v3.11.4
// source: third_party/tink/proto/common.proto

package common_go_proto

import (
	proto "github.com/golang/protobuf/proto"
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

type EllipticCurveType int32

const (
	EllipticCurveType_UNKNOWN_CURVE EllipticCurveType = 0
	EllipticCurveType_NIST_P256     EllipticCurveType = 2
	EllipticCurveType_NIST_P384     EllipticCurveType = 3
	EllipticCurveType_NIST_P521     EllipticCurveType = 4
	EllipticCurveType_CURVE25519    EllipticCurveType = 5
)

// Enum value maps for EllipticCurveType.
var (
	EllipticCurveType_name = map[int32]string{
		0: "UNKNOWN_CURVE",
		2: "NIST_P256",
		3: "NIST_P384",
		4: "NIST_P521",
		5: "CURVE25519",
	}
	EllipticCurveType_value = map[string]int32{
		"UNKNOWN_CURVE": 0,
		"NIST_P256":     2,
		"NIST_P384":     3,
		"NIST_P521":     4,
		"CURVE25519":    5,
	}
)

func (x EllipticCurveType) Enum() *EllipticCurveType {
	p := new(EllipticCurveType)
	*p = x
	return p
}

func (x EllipticCurveType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EllipticCurveType) Descriptor() protoreflect.EnumDescriptor {
	return file_third_party_tink_proto_common_proto_enumTypes[0].Descriptor()
}

func (EllipticCurveType) Type() protoreflect.EnumType {
	return &file_third_party_tink_proto_common_proto_enumTypes[0]
}

func (x EllipticCurveType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EllipticCurveType.Descriptor instead.
func (EllipticCurveType) EnumDescriptor() ([]byte, []int) {
	return file_third_party_tink_proto_common_proto_rawDescGZIP(), []int{0}
}

type EcPointFormat int32

const (
	EcPointFormat_UNKNOWN_FORMAT EcPointFormat = 0
	EcPointFormat_UNCOMPRESSED   EcPointFormat = 1
	EcPointFormat_COMPRESSED     EcPointFormat = 2
	// Like UNCOMPRESSED but without the \x04 prefix. Crunchy uses this format.
	// DO NOT USE unless you are a Crunchy user moving to Tink.
	EcPointFormat_DO_NOT_USE_CRUNCHY_UNCOMPRESSED EcPointFormat = 3
)

// Enum value maps for EcPointFormat.
var (
	EcPointFormat_name = map[int32]string{
		0: "UNKNOWN_FORMAT",
		1: "UNCOMPRESSED",
		2: "COMPRESSED",
		3: "DO_NOT_USE_CRUNCHY_UNCOMPRESSED",
	}
	EcPointFormat_value = map[string]int32{
		"UNKNOWN_FORMAT":                  0,
		"UNCOMPRESSED":                    1,
		"COMPRESSED":                      2,
		"DO_NOT_USE_CRUNCHY_UNCOMPRESSED": 3,
	}
)

func (x EcPointFormat) Enum() *EcPointFormat {
	p := new(EcPointFormat)
	*p = x
	return p
}

func (x EcPointFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EcPointFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_third_party_tink_proto_common_proto_enumTypes[1].Descriptor()
}

func (EcPointFormat) Type() protoreflect.EnumType {
	return &file_third_party_tink_proto_common_proto_enumTypes[1]
}

func (x EcPointFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EcPointFormat.Descriptor instead.
func (EcPointFormat) EnumDescriptor() ([]byte, []int) {
	return file_third_party_tink_proto_common_proto_rawDescGZIP(), []int{1}
}

type HashType int32

const (
	HashType_UNKNOWN_HASH HashType = 0
	HashType_SHA1         HashType = 1 // Using SHA1 for digital signature is deprecated but HMAC-SHA1 is
	// fine.
	HashType_SHA384 HashType = 2
	HashType_SHA256 HashType = 3
	HashType_SHA512 HashType = 4
)

// Enum value maps for HashType.
var (
	HashType_name = map[int32]string{
		0: "UNKNOWN_HASH",
		1: "SHA1",
		2: "SHA384",
		3: "SHA256",
		4: "SHA512",
	}
	HashType_value = map[string]int32{
		"UNKNOWN_HASH": 0,
		"SHA1":         1,
		"SHA384":       2,
		"SHA256":       3,
		"SHA512":       4,
	}
)

func (x HashType) Enum() *HashType {
	p := new(HashType)
	*p = x
	return p
}

func (x HashType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HashType) Descriptor() protoreflect.EnumDescriptor {
	return file_third_party_tink_proto_common_proto_enumTypes[2].Descriptor()
}

func (HashType) Type() protoreflect.EnumType {
	return &file_third_party_tink_proto_common_proto_enumTypes[2]
}

func (x HashType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HashType.Descriptor instead.
func (HashType) EnumDescriptor() ([]byte, []int) {
	return file_third_party_tink_proto_common_proto_rawDescGZIP(), []int{2}
}

var File_third_party_tink_proto_common_proto protoreflect.FileDescriptor

var file_third_party_tink_proto_common_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x69,
	0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2a, 0x63, 0x0a, 0x11, 0x45, 0x6c, 0x6c,
	0x69, 0x70, 0x74, 0x69, 0x63, 0x43, 0x75, 0x72, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x11,
	0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x43, 0x55, 0x52, 0x56, 0x45, 0x10,
	0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x49, 0x53, 0x54, 0x5f, 0x50, 0x32, 0x35, 0x36, 0x10, 0x02,
	0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x49, 0x53, 0x54, 0x5f, 0x50, 0x33, 0x38, 0x34, 0x10, 0x03, 0x12,
	0x0d, 0x0a, 0x09, 0x4e, 0x49, 0x53, 0x54, 0x5f, 0x50, 0x35, 0x32, 0x31, 0x10, 0x04, 0x12, 0x0e,
	0x0a, 0x0a, 0x43, 0x55, 0x52, 0x56, 0x45, 0x32, 0x35, 0x35, 0x31, 0x39, 0x10, 0x05, 0x2a, 0x6a,
	0x0a, 0x0d, 0x45, 0x63, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12,
	0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x46, 0x4f, 0x52, 0x4d, 0x41,
	0x54, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53,
	0x53, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4f, 0x4d, 0x50, 0x52, 0x45, 0x53,
	0x53, 0x45, 0x44, 0x10, 0x02, 0x12, 0x23, 0x0a, 0x1f, 0x44, 0x4f, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x55, 0x53, 0x45, 0x5f, 0x43, 0x52, 0x55, 0x4e, 0x43, 0x48, 0x59, 0x5f, 0x55, 0x4e, 0x43, 0x4f,
	0x4d, 0x50, 0x52, 0x45, 0x53, 0x53, 0x45, 0x44, 0x10, 0x03, 0x2a, 0x4a, 0x0a, 0x08, 0x48, 0x61,
	0x73, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x48, 0x41, 0x53, 0x48, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x48, 0x41, 0x31,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48, 0x41, 0x33, 0x38, 0x34, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x53, 0x48, 0x41, 0x32, 0x35, 0x36, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48,
	0x41, 0x35, 0x31, 0x32, 0x10, 0x04, 0x42, 0x57, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74, 0x69, 0x6e, 0x6b,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x67, 0x6f,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xa2, 0x02, 0x06, 0x54, 0x49, 0x4e, 0x4b, 0x50, 0x42, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_third_party_tink_proto_common_proto_rawDescOnce sync.Once
	file_third_party_tink_proto_common_proto_rawDescData = file_third_party_tink_proto_common_proto_rawDesc
)

func file_third_party_tink_proto_common_proto_rawDescGZIP() []byte {
	file_third_party_tink_proto_common_proto_rawDescOnce.Do(func() {
		file_third_party_tink_proto_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_third_party_tink_proto_common_proto_rawDescData)
	})
	return file_third_party_tink_proto_common_proto_rawDescData
}

var file_third_party_tink_proto_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_third_party_tink_proto_common_proto_goTypes = []interface{}{
	(EllipticCurveType)(0), // 0: google.crypto.tink.EllipticCurveType
	(EcPointFormat)(0),     // 1: google.crypto.tink.EcPointFormat
	(HashType)(0),          // 2: google.crypto.tink.HashType
}
var file_third_party_tink_proto_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_third_party_tink_proto_common_proto_init() }
func file_third_party_tink_proto_common_proto_init() {
	if File_third_party_tink_proto_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_third_party_tink_proto_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_third_party_tink_proto_common_proto_goTypes,
		DependencyIndexes: file_third_party_tink_proto_common_proto_depIdxs,
		EnumInfos:         file_third_party_tink_proto_common_proto_enumTypes,
	}.Build()
	File_third_party_tink_proto_common_proto = out.File
	file_third_party_tink_proto_common_proto_rawDesc = nil
	file_third_party_tink_proto_common_proto_goTypes = nil
	file_third_party_tink_proto_common_proto_depIdxs = nil
}
