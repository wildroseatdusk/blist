// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identity/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgCreateIdentity defines a message to create an identity
type MsgCreateIdentity struct {
	Id          string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PubKey      *PubKeyInfo `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pubkey" yaml:"pubkey"`
	Certificate string      `protobuf:"bytes,3,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Credentials string      `protobuf:"bytes,4,opt,name=credentials,proto3" json:"credentials,omitempty"`
	Owner       string      `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	Data        string      `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MsgCreateIdentity) Reset()         { *m = MsgCreateIdentity{} }
func (m *MsgCreateIdentity) String() string { return proto.CompactTextString(m) }
func (*MsgCreateIdentity) ProtoMessage()    {}
func (*MsgCreateIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a49ec0beed01e79, []int{0}
}
func (m *MsgCreateIdentity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateIdentity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateIdentity.Merge(m, src)
}
func (m *MsgCreateIdentity) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateIdentity proto.InternalMessageInfo

// MsgCreateIdentityResponse defines the Msg/Create response type.
type MsgCreateIdentityResponse struct {
}

func (m *MsgCreateIdentityResponse) Reset()         { *m = MsgCreateIdentityResponse{} }
func (m *MsgCreateIdentityResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateIdentityResponse) ProtoMessage()    {}
func (*MsgCreateIdentityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a49ec0beed01e79, []int{1}
}
func (m *MsgCreateIdentityResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateIdentityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateIdentityResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateIdentityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateIdentityResponse.Merge(m, src)
}
func (m *MsgCreateIdentityResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateIdentityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateIdentityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateIdentityResponse proto.InternalMessageInfo

// MsgUpdateIdentity defines a message to update an identity
type MsgUpdateIdentity struct {
	Id          string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PubKey      *PubKeyInfo `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pubkey" yaml:"pubkey"`
	Certificate string      `protobuf:"bytes,3,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Credentials string      `protobuf:"bytes,4,opt,name=credentials,proto3" json:"credentials,omitempty"`
	Owner       string      `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
	Data        string      `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *MsgUpdateIdentity) Reset()         { *m = MsgUpdateIdentity{} }
func (m *MsgUpdateIdentity) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateIdentity) ProtoMessage()    {}
func (*MsgUpdateIdentity) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a49ec0beed01e79, []int{2}
}
func (m *MsgUpdateIdentity) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateIdentity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateIdentity.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateIdentity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateIdentity.Merge(m, src)
}
func (m *MsgUpdateIdentity) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateIdentity) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateIdentity.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateIdentity proto.InternalMessageInfo

// MsgUpdateIdentityResponse defines the Msg/Update response type.
type MsgUpdateIdentityResponse struct {
}

func (m *MsgUpdateIdentityResponse) Reset()         { *m = MsgUpdateIdentityResponse{} }
func (m *MsgUpdateIdentityResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateIdentityResponse) ProtoMessage()    {}
func (*MsgUpdateIdentityResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a49ec0beed01e79, []int{3}
}
func (m *MsgUpdateIdentityResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateIdentityResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateIdentityResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateIdentityResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateIdentityResponse.Merge(m, src)
}
func (m *MsgUpdateIdentityResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateIdentityResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateIdentityResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateIdentityResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateIdentity)(nil), "iritamod.identity.MsgCreateIdentity")
	proto.RegisterType((*MsgCreateIdentityResponse)(nil), "iritamod.identity.MsgCreateIdentityResponse")
	proto.RegisterType((*MsgUpdateIdentity)(nil), "iritamod.identity.MsgUpdateIdentity")
	proto.RegisterType((*MsgUpdateIdentityResponse)(nil), "iritamod.identity.MsgUpdateIdentityResponse")
}

func init() { proto.RegisterFile("identity/tx.proto", fileDescriptor_4a49ec0beed01e79) }

var fileDescriptor_4a49ec0beed01e79 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe4, 0x53, 0xcd, 0xaa, 0xda, 0x40,
	0x18, 0xcd, 0xf8, 0x57, 0x3a, 0x52, 0xc1, 0x41, 0x68, 0xaa, 0x34, 0x4a, 0xe8, 0xc2, 0x45, 0x49,
	0xc0, 0x76, 0xe5, 0xd2, 0xae, 0x44, 0x84, 0x92, 0xd2, 0x4d, 0x37, 0x65, 0xe2, 0x7c, 0xa6, 0x53,
	0x4d, 0x26, 0x24, 0x13, 0xda, 0xbc, 0x45, 0x1f, 0xa1, 0x8f, 0xe3, 0xd2, 0xee, 0xba, 0x92, 0x56,
	0x37, 0x97, 0xcb, 0x5d, 0xdd, 0x27, 0xb8, 0x38, 0x31, 0xd7, 0xeb, 0x1f, 0xdc, 0xfd, 0xdd, 0x9d,
	0x39, 0xdf, 0xc9, 0x77, 0x38, 0x87, 0x7c, 0xb8, 0xce, 0x19, 0x04, 0x92, 0xcb, 0xd4, 0x96, 0x3f,
	0xad, 0x30, 0x12, 0x52, 0x90, 0x3a, 0x8f, 0xb8, 0xa4, 0xbe, 0x60, 0x56, 0x3e, 0x6b, 0xbe, 0xbc,
	0x57, 0xe5, 0x20, 0xd3, 0x36, 0x1b, 0x9e, 0xf0, 0x84, 0x82, 0xf6, 0x16, 0x65, 0xac, 0x79, 0x83,
	0x70, 0x7d, 0x1c, 0x7b, 0x1f, 0x22, 0xa0, 0x12, 0x86, 0xbb, 0x2f, 0x48, 0x0d, 0x17, 0x38, 0xd3,
	0x51, 0x07, 0x75, 0x9f, 0x3b, 0x05, 0xce, 0xc8, 0x27, 0xfc, 0x2c, 0x4c, 0xdc, 0xaf, 0x33, 0x48,
	0xf5, 0x42, 0x07, 0x75, 0xab, 0xbd, 0xd7, 0xd6, 0x89, 0xb3, 0xf5, 0x31, 0x71, 0x47, 0x90, 0x0e,
	0x83, 0xa9, 0x18, 0xb4, 0xae, 0x57, 0xed, 0x4a, 0x98, 0xb8, 0x33, 0x48, 0x6f, 0x57, 0xed, 0x17,
	0x29, 0xf5, 0xe7, 0x7d, 0x33, 0x7b, 0x9b, 0xce, 0x76, 0x30, 0x82, 0x94, 0x74, 0x70, 0x75, 0x02,
	0x91, 0xe4, 0x53, 0x3e, 0xa1, 0x12, 0xf4, 0xa2, 0x72, 0x7b, 0x48, 0x29, 0x45, 0x04, 0x6a, 0x3f,
	0x9d, 0xc7, 0x7a, 0x69, 0xa7, 0xd8, 0x53, 0xa4, 0x81, 0xcb, 0xe2, 0x47, 0x00, 0x91, 0x5e, 0x56,
	0xb3, 0xec, 0x41, 0x08, 0x2e, 0x31, 0x2a, 0xa9, 0x5e, 0x51, 0xa4, 0xc2, 0xfd, 0xd2, 0xd5, 0xef,
	0x36, 0x32, 0x5b, 0xf8, 0xd5, 0x49, 0x5a, 0x07, 0xe2, 0x50, 0x04, 0x31, 0xe4, 0x5d, 0x7c, 0x0e,
	0xd9, 0x13, 0xea, 0xe2, 0x30, 0x6d, 0xde, 0x45, 0xef, 0x0f, 0xc2, 0xc5, 0x71, 0xec, 0x11, 0x86,
	0x6b, 0x47, 0xff, 0xc6, 0x9b, 0x33, 0x71, 0x4f, 0x3a, 0x6d, 0xbe, 0x7d, 0x8c, 0x2a, 0x77, 0xdb,
	0xba, 0x1c, 0xb5, 0x7e, 0xc1, 0xe5, 0x50, 0x75, 0xc9, 0xe5, 0x7c, 0xa6, 0x81, 0xb3, 0xf8, 0x6f,
	0x68, 0x8b, 0xb5, 0x81, 0x96, 0x6b, 0x03, 0xfd, 0x5b, 0x1b, 0xe8, 0xd7, 0xc6, 0xd0, 0x96, 0x1b,
	0x43, 0xfb, 0xbb, 0x31, 0xb4, 0x2f, 0xef, 0x3d, 0x2e, 0xbf, 0x25, 0xae, 0x35, 0x11, 0xbe, 0xed,
	0x72, 0x1a, 0x7c, 0xe7, 0x40, 0xb9, 0x9d, 0xef, 0xb7, 0x7d, 0xc1, 0x92, 0x39, 0xc4, 0xf6, 0xfe,
	0x08, 0xd3, 0x10, 0x62, 0xb7, 0xa2, 0xce, 0xe8, 0xdd, 0x5d, 0x00, 0x00, 0x00, 0xff, 0xff, 0x37,
	0x28, 0xae, 0x2a, 0x9d, 0x03, 0x00, 0x00,
}

func (this *MsgCreateIdentity) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgCreateIdentity)
	if !ok {
		that2, ok := that.(MsgCreateIdentity)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if !this.PubKey.Equal(that1.PubKey) {
		return false
	}
	if this.Certificate != that1.Certificate {
		return false
	}
	if this.Credentials != that1.Credentials {
		return false
	}
	if this.Owner != that1.Owner {
		return false
	}
	if this.Data != that1.Data {
		return false
	}
	return true
}
func (this *MsgUpdateIdentity) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgUpdateIdentity)
	if !ok {
		that2, ok := that.(MsgUpdateIdentity)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if !this.PubKey.Equal(that1.PubKey) {
		return false
	}
	if this.Certificate != that1.Certificate {
		return false
	}
	if this.Credentials != that1.Credentials {
		return false
	}
	if this.Owner != that1.Owner {
		return false
	}
	if this.Data != that1.Data {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// CreateIdentity defines a method for creating a new identity.
	CreateIdentity(ctx context.Context, in *MsgCreateIdentity, opts ...grpc.CallOption) (*MsgCreateIdentityResponse, error)
	// UpdateIdentity defines a method for Updating a identity.
	UpdateIdentity(ctx context.Context, in *MsgUpdateIdentity, opts ...grpc.CallOption) (*MsgUpdateIdentityResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateIdentity(ctx context.Context, in *MsgCreateIdentity, opts ...grpc.CallOption) (*MsgCreateIdentityResponse, error) {
	out := new(MsgCreateIdentityResponse)
	err := c.cc.Invoke(ctx, "/iritamod.identity.Msg/CreateIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateIdentity(ctx context.Context, in *MsgUpdateIdentity, opts ...grpc.CallOption) (*MsgUpdateIdentityResponse, error) {
	out := new(MsgUpdateIdentityResponse)
	err := c.cc.Invoke(ctx, "/iritamod.identity.Msg/UpdateIdentity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateIdentity defines a method for creating a new identity.
	CreateIdentity(context.Context, *MsgCreateIdentity) (*MsgCreateIdentityResponse, error)
	// UpdateIdentity defines a method for Updating a identity.
	UpdateIdentity(context.Context, *MsgUpdateIdentity) (*MsgUpdateIdentityResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateIdentity(ctx context.Context, req *MsgCreateIdentity) (*MsgCreateIdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIdentity not implemented")
}
func (*UnimplementedMsgServer) UpdateIdentity(ctx context.Context, req *MsgUpdateIdentity) (*MsgUpdateIdentityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateIdentity not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iritamod.identity.Msg/CreateIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateIdentity(ctx, req.(*MsgCreateIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateIdentity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateIdentity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateIdentity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iritamod.identity.Msg/UpdateIdentity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateIdentity(ctx, req.(*MsgUpdateIdentity))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iritamod.identity.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIdentity",
			Handler:    _Msg_CreateIdentity_Handler,
		},
		{
			MethodName: "UpdateIdentity",
			Handler:    _Msg_UpdateIdentity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identity/tx.proto",
}

func (m *MsgCreateIdentity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateIdentity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateIdentity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Credentials) > 0 {
		i -= len(m.Credentials)
		copy(dAtA[i:], m.Credentials)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Credentials)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Certificate) > 0 {
		i -= len(m.Certificate)
		copy(dAtA[i:], m.Certificate)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Certificate)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PubKey != nil {
		{
			size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateIdentityResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateIdentityResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateIdentityResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateIdentity) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateIdentity) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateIdentity) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Credentials) > 0 {
		i -= len(m.Credentials)
		copy(dAtA[i:], m.Credentials)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Credentials)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Certificate) > 0 {
		i -= len(m.Certificate)
		copy(dAtA[i:], m.Certificate)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Certificate)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PubKey != nil {
		{
			size, err := m.PubKey.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateIdentityResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateIdentityResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateIdentityResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateIdentity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PubKey != nil {
		l = m.PubKey.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Certificate)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Credentials)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateIdentityResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateIdentity) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PubKey != nil {
		l = m.PubKey.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Certificate)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Credentials)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateIdentityResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateIdentity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateIdentity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateIdentity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PubKey == nil {
				m.PubKey = &PubKeyInfo{}
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certificate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credentials", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Credentials = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateIdentityResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateIdentityResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateIdentityResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateIdentity) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateIdentity: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateIdentity: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PubKey == nil {
				m.PubKey = &PubKeyInfo{}
			}
			if err := m.PubKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certificate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credentials", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Credentials = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateIdentityResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateIdentityResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateIdentityResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
