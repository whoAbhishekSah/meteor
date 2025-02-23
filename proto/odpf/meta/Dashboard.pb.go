// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.11.4
// source: odpf/meta/Dashboard.proto

package meta

import (
	common "github.com/odpf/meteor/proto/odpf/meta/common"
	facets "github.com/odpf/meteor/proto/odpf/meta/facets"
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

type DashboardKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urn string `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
}

func (x *DashboardKey) Reset() {
	*x = DashboardKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_meta_Dashboard_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DashboardKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DashboardKey) ProtoMessage() {}

func (x *DashboardKey) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_meta_Dashboard_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DashboardKey.ProtoReflect.Descriptor instead.
func (*DashboardKey) Descriptor() ([]byte, []int) {
	return file_odpf_meta_Dashboard_proto_rawDescGZIP(), []int{0}
}

func (x *DashboardKey) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

type Dashboard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Urn         string            `protobuf:"bytes,1,opt,name=urn,proto3" json:"urn,omitempty"`
	Name        string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Source      string            `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Description string            `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Url         string            `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	Charts      []*Chart          `protobuf:"bytes,6,rep,name=charts,proto3" json:"charts,omitempty"`
	Ownership   *facets.Ownership `protobuf:"bytes,21,opt,name=ownership,proto3" json:"ownership,omitempty"`
	Tags        *facets.Tags      `protobuf:"bytes,22,opt,name=tags,proto3" json:"tags,omitempty"`
	Custom      *facets.Custom    `protobuf:"bytes,23,opt,name=custom,proto3" json:"custom,omitempty"`
	Timestamps  *common.Timestamp `protobuf:"bytes,24,opt,name=timestamps,proto3" json:"timestamps,omitempty"`
	Event       *common.Event     `protobuf:"bytes,100,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *Dashboard) Reset() {
	*x = Dashboard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_odpf_meta_Dashboard_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dashboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dashboard) ProtoMessage() {}

func (x *Dashboard) ProtoReflect() protoreflect.Message {
	mi := &file_odpf_meta_Dashboard_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dashboard.ProtoReflect.Descriptor instead.
func (*Dashboard) Descriptor() ([]byte, []int) {
	return file_odpf_meta_Dashboard_proto_rawDescGZIP(), []int{1}
}

func (x *Dashboard) GetUrn() string {
	if x != nil {
		return x.Urn
	}
	return ""
}

func (x *Dashboard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Dashboard) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Dashboard) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Dashboard) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Dashboard) GetCharts() []*Chart {
	if x != nil {
		return x.Charts
	}
	return nil
}

func (x *Dashboard) GetOwnership() *facets.Ownership {
	if x != nil {
		return x.Ownership
	}
	return nil
}

func (x *Dashboard) GetTags() *facets.Tags {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Dashboard) GetCustom() *facets.Custom {
	if x != nil {
		return x.Custom
	}
	return nil
}

func (x *Dashboard) GetTimestamps() *common.Timestamp {
	if x != nil {
		return x.Timestamps
	}
	return nil
}

func (x *Dashboard) GetEvent() *common.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

var File_odpf_meta_Dashboard_proto protoreflect.FileDescriptor

var file_odpf_meta_Dashboard_proto_rawDesc = []byte{
	0x0a, 0x19, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x44, 0x61, 0x73, 0x68,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6f, 0x64, 0x70,
	0x66, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x15, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x6d, 0x65, 0x74,
	0x61, 0x2f, 0x43, 0x68, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6f,
	0x64, 0x70, 0x66, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2f,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x6f, 0x64, 0x70, 0x66, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x74,
	0x73, 0x2f, 0x54, 0x61, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x6f, 0x64,
	0x70, 0x66, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2f, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6f, 0x64, 0x70,
	0x66, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x6f,
	0x64, 0x70, 0x66, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0c, 0x44,
	0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x22, 0xac, 0x03,
	0x0a, 0x09, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x28, 0x0a,
	0x06, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x6f, 0x64, 0x70, 0x66, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x74, 0x52,
	0x06, 0x63, 0x68, 0x61, 0x72, 0x74, 0x73, 0x12, 0x39, 0x0a, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x64, 0x70,
	0x66, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x74, 0x73, 0x2e, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x09, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x68,
	0x69, 0x70, 0x12, 0x2a, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x66, 0x61, 0x63,
	0x65, 0x74, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x30,
	0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x66, 0x61, 0x63, 0x65, 0x74,
	0x73, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x12, 0x3b, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x18, 0x18,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x6d, 0x65, 0x74, 0x61,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x73, 0x12, 0x2d, 0x0a,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6f,
	0x64, 0x70, 0x66, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x3d, 0x0a, 0x13,
	0x69, 0x6f, 0x2e, 0x6f, 0x64, 0x70, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x42, 0x09, 0x44, 0x61, 0x73, 0x68, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x5a, 0x1b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x64, 0x70, 0x66, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6e, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_odpf_meta_Dashboard_proto_rawDescOnce sync.Once
	file_odpf_meta_Dashboard_proto_rawDescData = file_odpf_meta_Dashboard_proto_rawDesc
)

func file_odpf_meta_Dashboard_proto_rawDescGZIP() []byte {
	file_odpf_meta_Dashboard_proto_rawDescOnce.Do(func() {
		file_odpf_meta_Dashboard_proto_rawDescData = protoimpl.X.CompressGZIP(file_odpf_meta_Dashboard_proto_rawDescData)
	})
	return file_odpf_meta_Dashboard_proto_rawDescData
}

var file_odpf_meta_Dashboard_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_odpf_meta_Dashboard_proto_goTypes = []interface{}{
	(*DashboardKey)(nil),     // 0: odpf.meta.DashboardKey
	(*Dashboard)(nil),        // 1: odpf.meta.Dashboard
	(*Chart)(nil),            // 2: odpf.meta.Chart
	(*facets.Ownership)(nil), // 3: odpf.meta.facets.Ownership
	(*facets.Tags)(nil),      // 4: odpf.meta.facets.Tags
	(*facets.Custom)(nil),    // 5: odpf.meta.facets.Custom
	(*common.Timestamp)(nil), // 6: odpf.meta.common.Timestamp
	(*common.Event)(nil),     // 7: odpf.meta.common.Event
}
var file_odpf_meta_Dashboard_proto_depIdxs = []int32{
	2, // 0: odpf.meta.Dashboard.charts:type_name -> odpf.meta.Chart
	3, // 1: odpf.meta.Dashboard.ownership:type_name -> odpf.meta.facets.Ownership
	4, // 2: odpf.meta.Dashboard.tags:type_name -> odpf.meta.facets.Tags
	5, // 3: odpf.meta.Dashboard.custom:type_name -> odpf.meta.facets.Custom
	6, // 4: odpf.meta.Dashboard.timestamps:type_name -> odpf.meta.common.Timestamp
	7, // 5: odpf.meta.Dashboard.event:type_name -> odpf.meta.common.Event
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_odpf_meta_Dashboard_proto_init() }
func file_odpf_meta_Dashboard_proto_init() {
	if File_odpf_meta_Dashboard_proto != nil {
		return
	}
	file_odpf_meta_Chart_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_odpf_meta_Dashboard_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DashboardKey); i {
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
		file_odpf_meta_Dashboard_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dashboard); i {
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
			RawDescriptor: file_odpf_meta_Dashboard_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_odpf_meta_Dashboard_proto_goTypes,
		DependencyIndexes: file_odpf_meta_Dashboard_proto_depIdxs,
		MessageInfos:      file_odpf_meta_Dashboard_proto_msgTypes,
	}.Build()
	File_odpf_meta_Dashboard_proto = out.File
	file_odpf_meta_Dashboard_proto_rawDesc = nil
	file_odpf_meta_Dashboard_proto_goTypes = nil
	file_odpf_meta_Dashboard_proto_depIdxs = nil
}
