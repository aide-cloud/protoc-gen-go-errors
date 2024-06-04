// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.19.4
// source: err.proto

package api

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	ErrorReason_SYSTEM_ERROR        ErrorReason = 0
	ErrorReason_USER_NOT_FOUND      ErrorReason = 1
	ErrorReason_USER_ALREADY_EXISTS ErrorReason = 2
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "SYSTEM_ERROR",
		1: "USER_NOT_FOUND",
		2: "USER_ALREADY_EXISTS",
	}
	ErrorReason_value = map[string]int32{
		"SYSTEM_ERROR":        0,
		"USER_NOT_FOUND":      1,
		"USER_ALREADY_EXISTS": 2,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_err_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_err_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_err_proto_rawDescGZIP(), []int{0}
}

var File_err_proto protoreflect.FileDescriptor

var file_err_proto_rawDesc = []byte{
	0x0a, 0x09, 0x65, 0x72, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xcd, 0x01,
	0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x34, 0x0a,
	0x0c, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x1a,
	0x22, 0xa8, 0x45, 0xf4, 0x03, 0xb2, 0x45, 0x0c, 0xe7, 0xb3, 0xbb, 0xe7, 0xbb, 0x9f, 0xe9, 0x94,
	0x99, 0xe8, 0xaf, 0xaf, 0xba, 0x45, 0x0c, 0x53, 0x59, 0x53, 0x54, 0x45, 0x4d, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x12, 0x3b, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f,
	0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x1a, 0x27, 0xa8, 0x45, 0x94, 0x03, 0xb2, 0x45, 0x0f,
	0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe4, 0xb8, 0x8d, 0xe5, 0xad, 0x98, 0xe5, 0x9c, 0xa8, 0xba,
	0x45, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44,
	0x12, 0x45, 0x0a, 0x13, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59,
	0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x02, 0x1a, 0x2c, 0xa8, 0x45, 0x90, 0x03, 0xb2,
	0x45, 0x0f, 0xe7, 0x94, 0xa8, 0xe6, 0x88, 0xb7, 0xe5, 0xb7, 0xb2, 0xe5, 0xad, 0x98, 0xe5, 0x9c,
	0xa8, 0xba, 0x45, 0x13, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59,
	0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x4b, 0x0a,
	0x0b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x50, 0x01, 0x5a, 0x3a,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x69, 0x64, 0x65, 0x2d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_err_proto_rawDescOnce sync.Once
	file_err_proto_rawDescData = file_err_proto_rawDesc
)

func file_err_proto_rawDescGZIP() []byte {
	file_err_proto_rawDescOnce.Do(func() {
		file_err_proto_rawDescData = protoimpl.X.CompressGZIP(file_err_proto_rawDescData)
	})
	return file_err_proto_rawDescData
}

var file_err_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_err_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: example.api.ErrorReason
}
var file_err_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_err_proto_init() }
func file_err_proto_init() {
	if File_err_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_err_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_err_proto_goTypes,
		DependencyIndexes: file_err_proto_depIdxs,
		EnumInfos:         file_err_proto_enumTypes,
	}.Build()
	File_err_proto = out.File
	file_err_proto_rawDesc = nil
	file_err_proto_goTypes = nil
	file_err_proto_depIdxs = nil
}
