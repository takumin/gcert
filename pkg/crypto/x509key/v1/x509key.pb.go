// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: crypto/x509key/v1/x509key.proto

package x509key

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_crypto_x509key_v1_x509key_proto protoreflect.FileDescriptor

var file_crypto_x509key_v1_x509key_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x78, 0x35, 0x30, 0x39, 0x6b, 0x65, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x78, 0x35, 0x30, 0x39, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x11, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x78, 0x35, 0x30, 0x39, 0x6b, 0x65,
	0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x18, 0x5a, 0x16, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f,
	0x78, 0x35, 0x30, 0x39, 0x6b, 0x65, 0x79, 0x3b, 0x78, 0x35, 0x30, 0x39, 0x6b, 0x65, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_crypto_x509key_v1_x509key_proto_goTypes = []interface{}{}
var file_crypto_x509key_v1_x509key_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_crypto_x509key_v1_x509key_proto_init() }
func file_crypto_x509key_v1_x509key_proto_init() {
	if File_crypto_x509key_v1_x509key_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_crypto_x509key_v1_x509key_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_crypto_x509key_v1_x509key_proto_goTypes,
		DependencyIndexes: file_crypto_x509key_v1_x509key_proto_depIdxs,
	}.Build()
	File_crypto_x509key_v1_x509key_proto = out.File
	file_crypto_x509key_v1_x509key_proto_rawDesc = nil
	file_crypto_x509key_v1_x509key_proto_goTypes = nil
	file_crypto_x509key_v1_x509key_proto_depIdxs = nil
}
