// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: mqtt_error.proto

package v1

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

type TransportMqttErrorReason int32

const (
	TransportMqttErrorReason_UNKNOWN_ERROR TransportMqttErrorReason = 0
)

// Enum value maps for TransportMqttErrorReason.
var (
	TransportMqttErrorReason_name = map[int32]string{
		0: "UNKNOWN_ERROR",
	}
	TransportMqttErrorReason_value = map[string]int32{
		"UNKNOWN_ERROR": 0,
	}
)

func (x TransportMqttErrorReason) Enum() *TransportMqttErrorReason {
	p := new(TransportMqttErrorReason)
	*p = x
	return p
}

func (x TransportMqttErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransportMqttErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_mqtt_error_proto_enumTypes[0].Descriptor()
}

func (TransportMqttErrorReason) Type() protoreflect.EnumType {
	return &file_mqtt_error_proto_enumTypes[0]
}

func (x TransportMqttErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransportMqttErrorReason.Descriptor instead.
func (TransportMqttErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_mqtt_error_proto_rawDescGZIP(), []int{0}
}

var File_mqtt_error_proto protoreflect.FileDescriptor

var file_mqtt_error_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6d, 0x71, 0x74, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x6d, 0x71,
	0x74, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x33, 0x0a, 0x18, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x71, 0x74, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x11, 0x0a, 0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42,
	0x1a, 0x5a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x6d, 0x71, 0x74, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_mqtt_error_proto_rawDescOnce sync.Once
	file_mqtt_error_proto_rawDescData = file_mqtt_error_proto_rawDesc
)

func file_mqtt_error_proto_rawDescGZIP() []byte {
	file_mqtt_error_proto_rawDescOnce.Do(func() {
		file_mqtt_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_mqtt_error_proto_rawDescData)
	})
	return file_mqtt_error_proto_rawDescData
}

var file_mqtt_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mqtt_error_proto_goTypes = []interface{}{
	(TransportMqttErrorReason)(0), // 0: transport.mqtt.v1.TransportMqttErrorReason
}
var file_mqtt_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mqtt_error_proto_init() }
func file_mqtt_error_proto_init() {
	if File_mqtt_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mqtt_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mqtt_error_proto_goTypes,
		DependencyIndexes: file_mqtt_error_proto_depIdxs,
		EnumInfos:         file_mqtt_error_proto_enumTypes,
	}.Build()
	File_mqtt_error_proto = out.File
	file_mqtt_error_proto_rawDesc = nil
	file_mqtt_error_proto_goTypes = nil
	file_mqtt_error_proto_depIdxs = nil
}