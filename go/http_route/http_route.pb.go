// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: http_route.proto

package http_route

import (
	http_types "github.com/linkerd/linkerd2-proxy-api/go/http_types"
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

// Describes how to match an `:authority` or `host` header.
type HostMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Match:
	//	*HostMatch_Exact
	//	*HostMatch_Suffix_
	Match isHostMatch_Match `protobuf_oneof:"match"`
}

func (x *HostMatch) Reset() {
	*x = HostMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_route_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostMatch) ProtoMessage() {}

func (x *HostMatch) ProtoReflect() protoreflect.Message {
	mi := &file_http_route_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostMatch.ProtoReflect.Descriptor instead.
func (*HostMatch) Descriptor() ([]byte, []int) {
	return file_http_route_proto_rawDescGZIP(), []int{0}
}

func (m *HostMatch) GetMatch() isHostMatch_Match {
	if m != nil {
		return m.Match
	}
	return nil
}

func (x *HostMatch) GetExact() string {
	if x, ok := x.GetMatch().(*HostMatch_Exact); ok {
		return x.Exact
	}
	return ""
}

func (x *HostMatch) GetSuffix() *HostMatch_Suffix {
	if x, ok := x.GetMatch().(*HostMatch_Suffix_); ok {
		return x.Suffix
	}
	return nil
}

type isHostMatch_Match interface {
	isHostMatch_Match()
}

type HostMatch_Exact struct {
	// Match an exact hostname, e.g. www.example.com.
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof"`
}

type HostMatch_Suffix_ struct {
	// Match a hostname as a wildcard suffix, e.g. *.example.com.
	Suffix *HostMatch_Suffix `protobuf:"bytes,2,opt,name=suffix,proto3,oneof"`
}

func (*HostMatch_Exact) isHostMatch_Match() {}

func (*HostMatch_Suffix_) isHostMatch_Match() {}

// Describes a set of matches, ALL of which must apply.
type RouteMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Matches requests by path.
	Path *PathMatch `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// A set of header value matches that must be satisified. This match is not
	// comprehensive, so requests may include headers that are not covered by this
	// match.
	Header []*HeaderMatch `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty"`
	// A set of query parmaeter value matches that must be satisified. This match
	// is not comprehensive, so requests may include query parameters that are not
	// covered by this match.
	QueryParam []*QueryParamMatch `protobuf:"bytes,3,rep,name=query_param,json=queryParam,proto3" json:"query_param,omitempty"`
	// If specified, restricts the match to a single HTTP method.
	Method *http_types.HttpMethod `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
}

func (x *RouteMatch) Reset() {
	*x = RouteMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_route_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RouteMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RouteMatch) ProtoMessage() {}

func (x *RouteMatch) ProtoReflect() protoreflect.Message {
	mi := &file_http_route_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RouteMatch.ProtoReflect.Descriptor instead.
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return file_http_route_proto_rawDescGZIP(), []int{1}
}

func (x *RouteMatch) GetPath() *PathMatch {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *RouteMatch) GetHeader() []*HeaderMatch {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *RouteMatch) GetQueryParam() []*QueryParamMatch {
	if x != nil {
		return x.QueryParam
	}
	return nil
}

func (x *RouteMatch) GetMethod() *http_types.HttpMethod {
	if x != nil {
		return x.Method
	}
	return nil
}

// Describes how to match a path.
type PathMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Kind:
	//	*PathMatch_Exact
	//	*PathMatch_Prefix
	//	*PathMatch_Regex
	Kind isPathMatch_Kind `protobuf_oneof:"kind"`
}

func (x *PathMatch) Reset() {
	*x = PathMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_route_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathMatch) ProtoMessage() {}

func (x *PathMatch) ProtoReflect() protoreflect.Message {
	mi := &file_http_route_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathMatch.ProtoReflect.Descriptor instead.
func (*PathMatch) Descriptor() ([]byte, []int) {
	return file_http_route_proto_rawDescGZIP(), []int{2}
}

func (m *PathMatch) GetKind() isPathMatch_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (x *PathMatch) GetExact() string {
	if x, ok := x.GetKind().(*PathMatch_Exact); ok {
		return x.Exact
	}
	return ""
}

func (x *PathMatch) GetPrefix() string {
	if x, ok := x.GetKind().(*PathMatch_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (x *PathMatch) GetRegex() string {
	if x, ok := x.GetKind().(*PathMatch_Regex); ok {
		return x.Regex
	}
	return ""
}

type isPathMatch_Kind interface {
	isPathMatch_Kind()
}

type PathMatch_Exact struct {
	Exact string `protobuf:"bytes,1,opt,name=exact,proto3,oneof"`
}

type PathMatch_Prefix struct {
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}

type PathMatch_Regex struct {
	Regex string `protobuf:"bytes,3,opt,name=regex,proto3,oneof"`
}

func (*PathMatch_Exact) isPathMatch_Kind() {}

func (*PathMatch_Prefix) isPathMatch_Kind() {}

func (*PathMatch_Regex) isPathMatch_Kind() {}

// Describes how to match a header by name and value.
type HeaderMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to Value:
	//	*HeaderMatch_Exact
	//	*HeaderMatch_Regex
	Value isHeaderMatch_Value `protobuf_oneof:"value"`
}

func (x *HeaderMatch) Reset() {
	*x = HeaderMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_route_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeaderMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderMatch) ProtoMessage() {}

func (x *HeaderMatch) ProtoReflect() protoreflect.Message {
	mi := &file_http_route_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderMatch.ProtoReflect.Descriptor instead.
func (*HeaderMatch) Descriptor() ([]byte, []int) {
	return file_http_route_proto_rawDescGZIP(), []int{3}
}

func (x *HeaderMatch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *HeaderMatch) GetValue() isHeaderMatch_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *HeaderMatch) GetExact() string {
	if x, ok := x.GetValue().(*HeaderMatch_Exact); ok {
		return x.Exact
	}
	return ""
}

func (x *HeaderMatch) GetRegex() string {
	if x, ok := x.GetValue().(*HeaderMatch_Regex); ok {
		return x.Regex
	}
	return ""
}

type isHeaderMatch_Value interface {
	isHeaderMatch_Value()
}

type HeaderMatch_Exact struct {
	Exact string `protobuf:"bytes,2,opt,name=exact,proto3,oneof"`
}

type HeaderMatch_Regex struct {
	Regex string `protobuf:"bytes,3,opt,name=regex,proto3,oneof"`
}

func (*HeaderMatch_Exact) isHeaderMatch_Value() {}

func (*HeaderMatch_Regex) isHeaderMatch_Value() {}

// Describes how to match a query parameter by name and value.
type QueryParamMatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to Value:
	//	*QueryParamMatch_Exact
	//	*QueryParamMatch_Regex
	Value isQueryParamMatch_Value `protobuf_oneof:"value"`
}

func (x *QueryParamMatch) Reset() {
	*x = QueryParamMatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_route_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryParamMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryParamMatch) ProtoMessage() {}

func (x *QueryParamMatch) ProtoReflect() protoreflect.Message {
	mi := &file_http_route_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryParamMatch.ProtoReflect.Descriptor instead.
func (*QueryParamMatch) Descriptor() ([]byte, []int) {
	return file_http_route_proto_rawDescGZIP(), []int{4}
}

func (x *QueryParamMatch) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *QueryParamMatch) GetValue() isQueryParamMatch_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *QueryParamMatch) GetExact() string {
	if x, ok := x.GetValue().(*QueryParamMatch_Exact); ok {
		return x.Exact
	}
	return ""
}

func (x *QueryParamMatch) GetRegex() string {
	if x, ok := x.GetValue().(*QueryParamMatch_Regex); ok {
		return x.Regex
	}
	return ""
}

type isQueryParamMatch_Value interface {
	isQueryParamMatch_Value()
}

type QueryParamMatch_Exact struct {
	Exact string `protobuf:"bytes,2,opt,name=exact,proto3,oneof"`
}

type QueryParamMatch_Regex struct {
	Regex string `protobuf:"bytes,3,opt,name=regex,proto3,oneof"`
}

func (*QueryParamMatch_Exact) isQueryParamMatch_Value() {}

func (*QueryParamMatch_Regex) isQueryParamMatch_Value() {}

// Configures a route to modify a request's headers.
//
// Modifications are to be applied in the order they are described here:
// additions apply first, then sets, and then removals.
type RequestHeaderModifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of headers name-value pairs to set on requests, augmenting any
	// existing values for the header.
	Add *http_types.Headers `protobuf:"bytes,1,opt,name=add,proto3" json:"add,omitempty"`
	// A list of headers name-value pairs to set on requests, replacing any
	// existing values for the header.
	Set *http_types.Headers `protobuf:"bytes,2,opt,name=set,proto3" json:"set,omitempty"`
	// A list of headers names to be removed from requests.
	Remove []string `protobuf:"bytes,3,rep,name=remove,proto3" json:"remove,omitempty"`
}

func (x *RequestHeaderModifier) Reset() {
	*x = RequestHeaderModifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_route_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHeaderModifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHeaderModifier) ProtoMessage() {}

func (x *RequestHeaderModifier) ProtoReflect() protoreflect.Message {
	mi := &file_http_route_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHeaderModifier.ProtoReflect.Descriptor instead.
func (*RequestHeaderModifier) Descriptor() ([]byte, []int) {
	return file_http_route_proto_rawDescGZIP(), []int{5}
}

func (x *RequestHeaderModifier) GetAdd() *http_types.Headers {
	if x != nil {
		return x.Add
	}
	return nil
}

func (x *RequestHeaderModifier) GetSet() *http_types.Headers {
	if x != nil {
		return x.Set
	}
	return nil
}

func (x *RequestHeaderModifier) GetRemove() []string {
	if x != nil {
		return x.Remove
	}
	return nil
}

// Configures a route to respond with a redirect response. The `location` header
// is set with the given URL parameters.
type RequestRedirect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The scheme value to be used in the `location` header. If not specified,
	// the original request's scheme is used.
	Scheme *http_types.Scheme `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty"`
	// The host value to be used in the `location` header. If not specified, the
	// original request's host is used.
	Host string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`
	// If set, configures how the request's path should be modified for use in
	// the `location` header. If not specified, the original request's path is
	// used.
	Path *PathModifier `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// If set, specifies the port to use in the `location` header.
	Port uint32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	// The status code to use in the HTTP response. If not specified, 301 is
	// used.
	Status uint32 `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *RequestRedirect) Reset() {
	*x = RequestRedirect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_route_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestRedirect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestRedirect) ProtoMessage() {}

func (x *RequestRedirect) ProtoReflect() protoreflect.Message {
	mi := &file_http_route_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestRedirect.ProtoReflect.Descriptor instead.
func (*RequestRedirect) Descriptor() ([]byte, []int) {
	return file_http_route_proto_rawDescGZIP(), []int{6}
}

func (x *RequestRedirect) GetScheme() *http_types.Scheme {
	if x != nil {
		return x.Scheme
	}
	return nil
}

func (x *RequestRedirect) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RequestRedirect) GetPath() *PathModifier {
	if x != nil {
		return x.Path
	}
	return nil
}

func (x *RequestRedirect) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *RequestRedirect) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

// Describes how a path value may be rewritten in a route.
type PathModifier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Replace:
	//	*PathModifier_Full
	//	*PathModifier_Prefix
	Replace isPathModifier_Replace `protobuf_oneof:"replace"`
}

func (x *PathModifier) Reset() {
	*x = PathModifier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_route_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PathModifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PathModifier) ProtoMessage() {}

func (x *PathModifier) ProtoReflect() protoreflect.Message {
	mi := &file_http_route_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PathModifier.ProtoReflect.Descriptor instead.
func (*PathModifier) Descriptor() ([]byte, []int) {
	return file_http_route_proto_rawDescGZIP(), []int{7}
}

func (m *PathModifier) GetReplace() isPathModifier_Replace {
	if m != nil {
		return m.Replace
	}
	return nil
}

func (x *PathModifier) GetFull() string {
	if x, ok := x.GetReplace().(*PathModifier_Full); ok {
		return x.Full
	}
	return ""
}

func (x *PathModifier) GetPrefix() string {
	if x, ok := x.GetReplace().(*PathModifier_Prefix); ok {
		return x.Prefix
	}
	return ""
}

type isPathModifier_Replace interface {
	isPathModifier_Replace()
}

type PathModifier_Full struct {
	// Indicates that the path should be replaced with the given value.
	Full string `protobuf:"bytes,1,opt,name=full,proto3,oneof"`
}

type PathModifier_Prefix struct {
	// Indicates that the path prefix should be replaced with the given
	// value. When used, the route MUST match the request with PathMatch
	// prefix match. Server implementations SHOULD prevent the useof prefix
	// modifiers on routes that do use a PathMatch prefix match. Proxyies
	// MUST not process requests for routes where this condition is not
	// satisfied.
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof"`
}

func (*PathModifier_Full) isPathModifier_Replace() {}

func (*PathModifier_Prefix) isPathModifier_Replace() {}

// Configures a route to respond with a fixed response.
type Responder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status code to use in the HTTP response. Must be specified.
	Status  uint32              `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Headers *http_types.Headers `protobuf:"bytes,2,opt,name=headers,proto3" json:"headers,omitempty"`
	Body    []byte              `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Responder) Reset() {
	*x = Responder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_route_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Responder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Responder) ProtoMessage() {}

func (x *Responder) ProtoReflect() protoreflect.Message {
	mi := &file_http_route_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Responder.ProtoReflect.Descriptor instead.
func (*Responder) Descriptor() ([]byte, []int) {
	return file_http_route_proto_rawDescGZIP(), []int{8}
}

func (x *Responder) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Responder) GetHeaders() *http_types.Headers {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *Responder) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

// A match like `*.example.com` is encoded as [com, example].
type HostMatch_Suffix struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReverseLabels []string `protobuf:"bytes,1,rep,name=reverse_labels,json=reverseLabels,proto3" json:"reverse_labels,omitempty"`
}

func (x *HostMatch_Suffix) Reset() {
	*x = HostMatch_Suffix{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_route_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostMatch_Suffix) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostMatch_Suffix) ProtoMessage() {}

func (x *HostMatch_Suffix) ProtoReflect() protoreflect.Message {
	mi := &file_http_route_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostMatch_Suffix.ProtoReflect.Descriptor instead.
func (*HostMatch_Suffix) Descriptor() ([]byte, []int) {
	return file_http_route_proto_rawDescGZIP(), []int{0, 0}
}

func (x *HostMatch_Suffix) GetReverseLabels() []string {
	if x != nil {
		return x.ReverseLabels
	}
	return nil
}

var File_http_route_proto protoreflect.FileDescriptor

var file_http_route_proto_rawDesc = []byte{
	0x0a, 0x10, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1b, 0x69, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x1a,
	0x10, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x09, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12,
	0x16, 0x0a, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x12, 0x47, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x69, 0x6f, 0x2e, 0x6c, 0x69, 0x6e,
	0x6b, 0x65, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x2e,
	0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x48, 0x00, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78,
	0x1a, 0x2f, 0x0a, 0x06, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65,
	0x76, 0x65, 0x72, 0x73, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x42, 0x07, 0x0a, 0x05, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x22, 0x9a, 0x02, 0x0a, 0x0a, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x3a, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x69, 0x6f, 0x2e, 0x6c, 0x69, 0x6e,
	0x6b, 0x65, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x40, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x69, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65,
	0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x4d, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x69,
	0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x0a, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x3f, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x69, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b,
	0x65, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52,
	0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x5d, 0x0a, 0x09, 0x50, 0x61, 0x74, 0x68, 0x4d,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x12, 0x18, 0x0a, 0x06,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06,
	0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x16, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x42, 0x06,
	0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x22, 0x5a, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x78, 0x61,
	0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x78, 0x61, 0x63,
	0x74, 0x12, 0x16, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x5e, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x78, 0x61,
	0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x78, 0x61, 0x63,
	0x74, 0x12, 0x16, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x9f, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x03,
	0x61, 0x64, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x6f, 0x2e, 0x6c,
	0x69, 0x6e, 0x6b, 0x65, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x03, 0x61, 0x64, 0x64, 0x12, 0x36, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x69, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x78, 0x79, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x03, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x22, 0xcd, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x69, 0x6f, 0x2e, 0x6c, 0x69,
	0x6e, 0x6b, 0x65, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x3d, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x69, 0x6f, 0x2e, 0x6c, 0x69, 0x6e,
	0x6b, 0x65, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e, 0x50, 0x61, 0x74, 0x68, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x49, 0x0a, 0x0c, 0x50, 0x61, 0x74, 0x68, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x04, 0x66, 0x75, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x66, 0x75, 0x6c, 0x6c, 0x12, 0x18, 0x0a, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x42, 0x09, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x22,
	0x77, 0x0a, 0x09, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x3e, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x69, 0x6f, 0x2e, 0x6c, 0x69, 0x6e, 0x6b, 0x65,
	0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x52, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x65, 0x72, 0x64, 0x2f, 0x6c,
	0x69, 0x6e, 0x6b, 0x65, 0x72, 0x64, 0x32, 0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_http_route_proto_rawDescOnce sync.Once
	file_http_route_proto_rawDescData = file_http_route_proto_rawDesc
)

func file_http_route_proto_rawDescGZIP() []byte {
	file_http_route_proto_rawDescOnce.Do(func() {
		file_http_route_proto_rawDescData = protoimpl.X.CompressGZIP(file_http_route_proto_rawDescData)
	})
	return file_http_route_proto_rawDescData
}

var file_http_route_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_http_route_proto_goTypes = []interface{}{
	(*HostMatch)(nil),             // 0: io.linkerd.proxy.http_route.HostMatch
	(*RouteMatch)(nil),            // 1: io.linkerd.proxy.http_route.RouteMatch
	(*PathMatch)(nil),             // 2: io.linkerd.proxy.http_route.PathMatch
	(*HeaderMatch)(nil),           // 3: io.linkerd.proxy.http_route.HeaderMatch
	(*QueryParamMatch)(nil),       // 4: io.linkerd.proxy.http_route.QueryParamMatch
	(*RequestHeaderModifier)(nil), // 5: io.linkerd.proxy.http_route.RequestHeaderModifier
	(*RequestRedirect)(nil),       // 6: io.linkerd.proxy.http_route.RequestRedirect
	(*PathModifier)(nil),          // 7: io.linkerd.proxy.http_route.PathModifier
	(*Responder)(nil),             // 8: io.linkerd.proxy.http_route.Responder
	(*HostMatch_Suffix)(nil),      // 9: io.linkerd.proxy.http_route.HostMatch.Suffix
	(*http_types.HttpMethod)(nil), // 10: io.linkerd.proxy.http_types.HttpMethod
	(*http_types.Headers)(nil),    // 11: io.linkerd.proxy.http_types.Headers
	(*http_types.Scheme)(nil),     // 12: io.linkerd.proxy.http_types.Scheme
}
var file_http_route_proto_depIdxs = []int32{
	9,  // 0: io.linkerd.proxy.http_route.HostMatch.suffix:type_name -> io.linkerd.proxy.http_route.HostMatch.Suffix
	2,  // 1: io.linkerd.proxy.http_route.RouteMatch.path:type_name -> io.linkerd.proxy.http_route.PathMatch
	3,  // 2: io.linkerd.proxy.http_route.RouteMatch.header:type_name -> io.linkerd.proxy.http_route.HeaderMatch
	4,  // 3: io.linkerd.proxy.http_route.RouteMatch.query_param:type_name -> io.linkerd.proxy.http_route.QueryParamMatch
	10, // 4: io.linkerd.proxy.http_route.RouteMatch.method:type_name -> io.linkerd.proxy.http_types.HttpMethod
	11, // 5: io.linkerd.proxy.http_route.RequestHeaderModifier.add:type_name -> io.linkerd.proxy.http_types.Headers
	11, // 6: io.linkerd.proxy.http_route.RequestHeaderModifier.set:type_name -> io.linkerd.proxy.http_types.Headers
	12, // 7: io.linkerd.proxy.http_route.RequestRedirect.scheme:type_name -> io.linkerd.proxy.http_types.Scheme
	7,  // 8: io.linkerd.proxy.http_route.RequestRedirect.path:type_name -> io.linkerd.proxy.http_route.PathModifier
	11, // 9: io.linkerd.proxy.http_route.Responder.headers:type_name -> io.linkerd.proxy.http_types.Headers
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_http_route_proto_init() }
func file_http_route_proto_init() {
	if File_http_route_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_http_route_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostMatch); i {
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
		file_http_route_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RouteMatch); i {
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
		file_http_route_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathMatch); i {
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
		file_http_route_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeaderMatch); i {
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
		file_http_route_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryParamMatch); i {
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
		file_http_route_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHeaderModifier); i {
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
		file_http_route_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestRedirect); i {
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
		file_http_route_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PathModifier); i {
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
		file_http_route_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Responder); i {
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
		file_http_route_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostMatch_Suffix); i {
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
	file_http_route_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*HostMatch_Exact)(nil),
		(*HostMatch_Suffix_)(nil),
	}
	file_http_route_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*PathMatch_Exact)(nil),
		(*PathMatch_Prefix)(nil),
		(*PathMatch_Regex)(nil),
	}
	file_http_route_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*HeaderMatch_Exact)(nil),
		(*HeaderMatch_Regex)(nil),
	}
	file_http_route_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*QueryParamMatch_Exact)(nil),
		(*QueryParamMatch_Regex)(nil),
	}
	file_http_route_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*PathModifier_Full)(nil),
		(*PathModifier_Prefix)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_http_route_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_http_route_proto_goTypes,
		DependencyIndexes: file_http_route_proto_depIdxs,
		MessageInfos:      file_http_route_proto_msgTypes,
	}.Build()
	File_http_route_proto = out.File
	file_http_route_proto_rawDesc = nil
	file_http_route_proto_goTypes = nil
	file_http_route_proto_depIdxs = nil
}
