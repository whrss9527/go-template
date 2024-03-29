// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.9
// source: template_proj/errors/error_reason.proto

package errors

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
	// 未知异常
	ErrorReason_CANCELLATION_IN_PROGRESS ErrorReason = 0
	// 用户未找到
	ErrorReason_USER_NOT_FOUND ErrorReason = 1
	// 请求繁忙
	ErrorReason_REQUEST_BUSY ErrorReason = 2
	// 客户端版本需要升级
	ErrorReason_CLIENT_VERSION_NEEDS_TO_BE_UPGRADED ErrorReason = 3
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "CANCELLATION_IN_PROGRESS",
		1: "USER_NOT_FOUND",
		2: "REQUEST_BUSY",
		3: "CLIENT_VERSION_NEEDS_TO_BE_UPGRADED",
	}
	ErrorReason_value = map[string]int32{
		"CANCELLATION_IN_PROGRESS":            0,
		"USER_NOT_FOUND":                      1,
		"REQUEST_BUSY":                        2,
		"CLIENT_VERSION_NEEDS_TO_BE_UPGRADED": 3,
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
	return file_template_proj_errors_error_reason_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_template_proj_errors_error_reason_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_template_proj_errors_error_reason_proto_rawDescGZIP(), []int{0}
}

var File_template_proj_errors_error_reason_proto protoreflect.FileDescriptor

var file_template_proj_errors_error_reason_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x72, 0x65, 0x61,
	0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a,
	0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x80, 0x01, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53,
	0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46,
	0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x5f, 0x42, 0x55, 0x53, 0x59, 0x10, 0x02, 0x12, 0x27, 0x0a, 0x23, 0x43, 0x4c, 0x49, 0x45,
	0x4e, 0x54, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x4e, 0x45, 0x45, 0x44, 0x53,
	0x5f, 0x54, 0x4f, 0x5f, 0x42, 0x45, 0x5f, 0x55, 0x50, 0x47, 0x52, 0x41, 0x44, 0x45, 0x44, 0x10,
	0x03, 0x1a, 0x04, 0xa0, 0x45, 0x90, 0x03, 0x42, 0x35, 0x0a, 0x14, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x50,
	0x01, 0x5a, 0x1b, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x6a,
	0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_template_proj_errors_error_reason_proto_rawDescOnce sync.Once
	file_template_proj_errors_error_reason_proto_rawDescData = file_template_proj_errors_error_reason_proto_rawDesc
)

func file_template_proj_errors_error_reason_proto_rawDescGZIP() []byte {
	file_template_proj_errors_error_reason_proto_rawDescOnce.Do(func() {
		file_template_proj_errors_error_reason_proto_rawDescData = protoimpl.X.CompressGZIP(file_template_proj_errors_error_reason_proto_rawDescData)
	})
	return file_template_proj_errors_error_reason_proto_rawDescData
}

var file_template_proj_errors_error_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_template_proj_errors_error_reason_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: template_proj.errors.ErrorReason
}
var file_template_proj_errors_error_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_template_proj_errors_error_reason_proto_init() }
func file_template_proj_errors_error_reason_proto_init() {
	if File_template_proj_errors_error_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_template_proj_errors_error_reason_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_template_proj_errors_error_reason_proto_goTypes,
		DependencyIndexes: file_template_proj_errors_error_reason_proto_depIdxs,
		EnumInfos:         file_template_proj_errors_error_reason_proto_enumTypes,
	}.Build()
	File_template_proj_errors_error_reason_proto = out.File
	file_template_proj_errors_error_reason_proto_rawDesc = nil
	file_template_proj_errors_error_reason_proto_goTypes = nil
	file_template_proj_errors_error_reason_proto_depIdxs = nil
}
