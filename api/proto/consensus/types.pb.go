// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: proto/consensus/types.proto

package consensus

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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*Request_Echo
	Value isRequest_Value `protobuf_oneof:"value"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consensus_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consensus_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_proto_consensus_types_proto_rawDescGZIP(), []int{0}
}

func (m *Request) GetValue() isRequest_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Request) GetEcho() *RequestEcho {
	if x, ok := x.GetValue().(*Request_Echo); ok {
		return x.Echo
	}
	return nil
}

type isRequest_Value interface {
	isRequest_Value()
}

type Request_Echo struct {
	Echo *RequestEcho `protobuf:"bytes,1,opt,name=echo,proto3,oneof"`
}

func (*Request_Echo) isRequest_Value() {}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//
	//	*Response_Echo
	Value isResponse_Value `protobuf_oneof:"value"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consensus_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consensus_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_consensus_types_proto_rawDescGZIP(), []int{1}
}

func (m *Response) GetValue() isResponse_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Response) GetEcho() *ResponseEcho {
	if x, ok := x.GetValue().(*Response_Echo); ok {
		return x.Echo
	}
	return nil
}

type isResponse_Value interface {
	isResponse_Value()
}

type Response_Echo struct {
	Echo *ResponseEcho `protobuf:"bytes,1,opt,name=echo,proto3,oneof"`
}

func (*Response_Echo) isResponse_Value() {}

type RequestEcho struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RequestEcho) Reset() {
	*x = RequestEcho{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consensus_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestEcho) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestEcho) ProtoMessage() {}

func (x *RequestEcho) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consensus_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestEcho.ProtoReflect.Descriptor instead.
func (*RequestEcho) Descriptor() ([]byte, []int) {
	return file_proto_consensus_types_proto_rawDescGZIP(), []int{2}
}

func (x *RequestEcho) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ResponseEcho struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponseEcho) Reset() {
	*x = ResponseEcho{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consensus_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseEcho) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseEcho) ProtoMessage() {}

func (x *ResponseEcho) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consensus_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseEcho.ProtoReflect.Descriptor instead.
func (*ResponseEcho) Descriptor() ([]byte, []int) {
	return file_proto_consensus_types_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseEcho) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ExternalTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`            //Namespace used to classify services
	TxBytes   []byte `protobuf:"bytes,2,opt,name=tx_bytes,json=txBytes,proto3" json:"tx_bytes,omitempty"` //Transaction hash or some raw data
}

func (x *ExternalTransaction) Reset() {
	*x = ExternalTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consensus_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalTransaction) ProtoMessage() {}

func (x *ExternalTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consensus_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalTransaction.ProtoReflect.Descriptor instead.
func (*ExternalTransaction) Descriptor() ([]byte, []int) {
	return file_proto_consensus_types_proto_rawDescGZIP(), []int{4}
}

func (x *ExternalTransaction) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ExternalTransaction) GetTxBytes() []byte {
	if x != nil {
		return x.TxBytes
	}
	return nil
}

type CommitedTransactions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*ExternalTransaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *CommitedTransactions) Reset() {
	*x = CommitedTransactions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_consensus_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommitedTransactions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommitedTransactions) ProtoMessage() {}

func (x *CommitedTransactions) ProtoReflect() protoreflect.Message {
	mi := &file_proto_consensus_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommitedTransactions.ProtoReflect.Descriptor instead.
func (*CommitedTransactions) Descriptor() ([]byte, []int) {
	return file_proto_consensus_types_proto_rawDescGZIP(), []int{5}
}

func (x *CommitedTransactions) GetTransactions() []*ExternalTransaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

var File_proto_consensus_types_proto protoreflect.FileDescriptor

var file_proto_consensus_types_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x22, 0x46, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x65, 0x63, 0x68,
	0x6f, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03,
	0x22, 0x48, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x04,
	0x65, 0x63, 0x68, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45,
	0x63, 0x68, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x42, 0x07, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0x27, 0x0a, 0x0b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x28, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45,
	0x63, 0x68, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4e, 0x0a,
	0x13, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x78, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x78, 0x42, 0x79, 0x74, 0x65, 0x73, 0x22, 0x5a, 0x0a,
	0x14, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x42, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f,
	0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x6f, 0x72,
	0x67, 0x2f, 0x63, 0x6f, 0x6d, 0x65, 0x74, 0x62, 0x66, 0x74, 0x2f, 0x73, 0x63, 0x61, 0x6c, 0x61,
	0x72, 0x69, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_consensus_types_proto_rawDescOnce sync.Once
	file_proto_consensus_types_proto_rawDescData = file_proto_consensus_types_proto_rawDesc
)

func file_proto_consensus_types_proto_rawDescGZIP() []byte {
	file_proto_consensus_types_proto_rawDescOnce.Do(func() {
		file_proto_consensus_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_consensus_types_proto_rawDescData)
	})
	return file_proto_consensus_types_proto_rawDescData
}

var file_proto_consensus_types_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_consensus_types_proto_goTypes = []interface{}{
	(*Request)(nil),              // 0: consensus.Request
	(*Response)(nil),             // 1: consensus.Response
	(*RequestEcho)(nil),          // 2: consensus.RequestEcho
	(*ResponseEcho)(nil),         // 3: consensus.ResponseEcho
	(*ExternalTransaction)(nil),  // 4: consensus.ExternalTransaction
	(*CommitedTransactions)(nil), // 5: consensus.CommitedTransactions
}
var file_proto_consensus_types_proto_depIdxs = []int32{
	2, // 0: consensus.Request.echo:type_name -> consensus.RequestEcho
	3, // 1: consensus.Response.echo:type_name -> consensus.ResponseEcho
	4, // 2: consensus.CommitedTransactions.transactions:type_name -> consensus.ExternalTransaction
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_consensus_types_proto_init() }
func file_proto_consensus_types_proto_init() {
	if File_proto_consensus_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_consensus_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_proto_consensus_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_consensus_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestEcho); i {
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
		file_proto_consensus_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseEcho); i {
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
		file_proto_consensus_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalTransaction); i {
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
		file_proto_consensus_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommitedTransactions); i {
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
	file_proto_consensus_types_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Request_Echo)(nil),
	}
	file_proto_consensus_types_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Response_Echo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_consensus_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_consensus_types_proto_goTypes,
		DependencyIndexes: file_proto_consensus_types_proto_depIdxs,
		MessageInfos:      file_proto_consensus_types_proto_msgTypes,
	}.Build()
	File_proto_consensus_types_proto = out.File
	file_proto_consensus_types_proto_rawDesc = nil
	file_proto_consensus_types_proto_goTypes = nil
	file_proto_consensus_types_proto_depIdxs = nil
}
