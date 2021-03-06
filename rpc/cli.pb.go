// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cli.proto

package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TransInfo struct {
	Seckey               string   `protobuf:"bytes,1,opt,name=seckey" json:"seckey,omitempty"`
	Nonce                int64    `protobuf:"varint,2,opt,name=nonce" json:"nonce,omitempty"`
	Contract             string   `protobuf:"bytes,3,opt,name=contract" json:"contract,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransInfo) Reset()         { *m = TransInfo{} }
func (m *TransInfo) String() string { return proto.CompactTextString(m) }
func (*TransInfo) ProtoMessage()    {}
func (*TransInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_441c81ee1c4e0f3c, []int{0}
}
func (m *TransInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransInfo.Unmarshal(m, b)
}
func (m *TransInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransInfo.Marshal(b, m, deterministic)
}
func (dst *TransInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransInfo.Merge(dst, src)
}
func (m *TransInfo) XXX_Size() int {
	return xxx_messageInfo_TransInfo.Size(m)
}
func (m *TransInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TransInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TransInfo proto.InternalMessageInfo

func (m *TransInfo) GetSeckey() string {
	if m != nil {
		return m.Seckey
	}
	return ""
}

func (m *TransInfo) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *TransInfo) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

type Transaction struct {
	Tx                   []byte   `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_441c81ee1c4e0f3c, []int{1}
}
func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (dst *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(dst, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetTx() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

type PublishRet struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Hash                 []byte   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishRet) Reset()         { *m = PublishRet{} }
func (m *PublishRet) String() string { return proto.CompactTextString(m) }
func (*PublishRet) ProtoMessage()    {}
func (*PublishRet) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_441c81ee1c4e0f3c, []int{2}
}
func (m *PublishRet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRet.Unmarshal(m, b)
}
func (m *PublishRet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRet.Marshal(b, m, deterministic)
}
func (dst *PublishRet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRet.Merge(dst, src)
}
func (m *PublishRet) XXX_Size() int {
	return xxx_messageInfo_PublishRet.Size(m)
}
func (m *PublishRet) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRet.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRet proto.InternalMessageInfo

func (m *PublishRet) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *PublishRet) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Response struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_441c81ee1c4e0f3c, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

type TransactionKey struct {
	Publisher            []byte   `protobuf:"bytes,1,opt,name=publisher,proto3" json:"publisher,omitempty"`
	Nonce                int64    `protobuf:"varint,2,opt,name=nonce" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionKey) Reset()         { *m = TransactionKey{} }
func (m *TransactionKey) String() string { return proto.CompactTextString(m) }
func (*TransactionKey) ProtoMessage()    {}
func (*TransactionKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_441c81ee1c4e0f3c, []int{4}
}
func (m *TransactionKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionKey.Unmarshal(m, b)
}
func (m *TransactionKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionKey.Marshal(b, m, deterministic)
}
func (dst *TransactionKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionKey.Merge(dst, src)
}
func (m *TransactionKey) XXX_Size() int {
	return xxx_messageInfo_TransactionKey.Size(m)
}
func (m *TransactionKey) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionKey.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionKey proto.InternalMessageInfo

func (m *TransactionKey) GetPublisher() []byte {
	if m != nil {
		return m.Publisher
	}
	return nil
}

func (m *TransactionKey) GetNonce() int64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type TransactionHash struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionHash) Reset()         { *m = TransactionHash{} }
func (m *TransactionHash) String() string { return proto.CompactTextString(m) }
func (*TransactionHash) ProtoMessage()    {}
func (*TransactionHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_441c81ee1c4e0f3c, []int{5}
}
func (m *TransactionHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionHash.Unmarshal(m, b)
}
func (m *TransactionHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionHash.Marshal(b, m, deterministic)
}
func (dst *TransactionHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionHash.Merge(dst, src)
}
func (m *TransactionHash) XXX_Size() int {
	return xxx_messageInfo_TransactionHash.Size(m)
}
func (m *TransactionHash) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionHash.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionHash proto.InternalMessageInfo

func (m *TransactionHash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type Key struct {
	S                    string   `protobuf:"bytes,1,opt,name=s" json:"s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_441c81ee1c4e0f3c, []int{6}
}
func (m *Key) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Key.Unmarshal(m, b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Key.Marshal(b, m, deterministic)
}
func (dst *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(dst, src)
}
func (m *Key) XXX_Size() int {
	return xxx_messageInfo_Key.Size(m)
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetS() string {
	if m != nil {
		return m.S
	}
	return ""
}

type Value struct {
	Sv                   string   `protobuf:"bytes,2,opt,name=sv" json:"sv,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_441c81ee1c4e0f3c, []int{7}
}
func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (dst *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(dst, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetSv() string {
	if m != nil {
		return m.Sv
	}
	return ""
}

type BlockKey struct {
	Layer                int64    `protobuf:"varint,1,opt,name=layer" json:"layer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockKey) Reset()         { *m = BlockKey{} }
func (m *BlockKey) String() string { return proto.CompactTextString(m) }
func (*BlockKey) ProtoMessage()    {}
func (*BlockKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_441c81ee1c4e0f3c, []int{8}
}
func (m *BlockKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockKey.Unmarshal(m, b)
}
func (m *BlockKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockKey.Marshal(b, m, deterministic)
}
func (dst *BlockKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockKey.Merge(dst, src)
}
func (m *BlockKey) XXX_Size() int {
	return xxx_messageInfo_BlockKey.Size(m)
}
func (m *BlockKey) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockKey.DiscardUnknown(m)
}

var xxx_messageInfo_BlockKey proto.InternalMessageInfo

func (m *BlockKey) GetLayer() int64 {
	if m != nil {
		return m.Layer
	}
	return 0
}

type Head struct {
	Version              int64    `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	ParentHash           []byte   `protobuf:"bytes,2,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	TreeHash             []byte   `protobuf:"bytes,3,opt,name=treeHash,proto3" json:"treeHash,omitempty"`
	BlockHash            []byte   `protobuf:"bytes,4,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	Info                 []byte   `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	Number               int64    `protobuf:"varint,6,opt,name=number" json:"number,omitempty"`
	Witness              string   `protobuf:"bytes,7,opt,name=witness" json:"witness,omitempty"`
	Signature            []byte   `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
	Time                 int64    `protobuf:"varint,9,opt,name=time" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Head) Reset()         { *m = Head{} }
func (m *Head) String() string { return proto.CompactTextString(m) }
func (*Head) ProtoMessage()    {}
func (*Head) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_441c81ee1c4e0f3c, []int{9}
}
func (m *Head) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Head.Unmarshal(m, b)
}
func (m *Head) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Head.Marshal(b, m, deterministic)
}
func (dst *Head) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Head.Merge(dst, src)
}
func (m *Head) XXX_Size() int {
	return xxx_messageInfo_Head.Size(m)
}
func (m *Head) XXX_DiscardUnknown() {
	xxx_messageInfo_Head.DiscardUnknown(m)
}

var xxx_messageInfo_Head proto.InternalMessageInfo

func (m *Head) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Head) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *Head) GetTreeHash() []byte {
	if m != nil {
		return m.TreeHash
	}
	return nil
}

func (m *Head) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *Head) GetInfo() []byte {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *Head) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

func (m *Head) GetWitness() string {
	if m != nil {
		return m.Witness
	}
	return ""
}

func (m *Head) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Head) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type BlockInfo struct {
	Head                 *Head             `protobuf:"bytes,1,opt,name=head" json:"head,omitempty"`
	Txcnt                int64             `protobuf:"varint,2,opt,name=Txcnt" json:"Txcnt,omitempty"`
	TxList               []*TransactionKey `protobuf:"bytes,3,rep,name=txList" json:"txList,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *BlockInfo) Reset()         { *m = BlockInfo{} }
func (m *BlockInfo) String() string { return proto.CompactTextString(m) }
func (*BlockInfo) ProtoMessage()    {}
func (*BlockInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cli_441c81ee1c4e0f3c, []int{10}
}
func (m *BlockInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockInfo.Unmarshal(m, b)
}
func (m *BlockInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockInfo.Marshal(b, m, deterministic)
}
func (dst *BlockInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockInfo.Merge(dst, src)
}
func (m *BlockInfo) XXX_Size() int {
	return xxx_messageInfo_BlockInfo.Size(m)
}
func (m *BlockInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BlockInfo proto.InternalMessageInfo

func (m *BlockInfo) GetHead() *Head {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *BlockInfo) GetTxcnt() int64 {
	if m != nil {
		return m.Txcnt
	}
	return 0
}

func (m *BlockInfo) GetTxList() []*TransactionKey {
	if m != nil {
		return m.TxList
	}
	return nil
}

func init() {
	proto.RegisterType((*TransInfo)(nil), "rpc.TransInfo")
	proto.RegisterType((*Transaction)(nil), "rpc.Transaction")
	proto.RegisterType((*PublishRet)(nil), "rpc.PublishRet")
	proto.RegisterType((*Response)(nil), "rpc.Response")
	proto.RegisterType((*TransactionKey)(nil), "rpc.TransactionKey")
	proto.RegisterType((*TransactionHash)(nil), "rpc.TransactionHash")
	proto.RegisterType((*Key)(nil), "rpc.Key")
	proto.RegisterType((*Value)(nil), "rpc.Value")
	proto.RegisterType((*BlockKey)(nil), "rpc.BlockKey")
	proto.RegisterType((*Head)(nil), "rpc.Head")
	proto.RegisterType((*BlockInfo)(nil), "rpc.BlockInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CliClient is the client API for Cli service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CliClient interface {
	PublishTx(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*PublishRet, error)
	GetTransaction(ctx context.Context, in *TransactionKey, opts ...grpc.CallOption) (*Transaction, error)
	GetTransactionByHash(ctx context.Context, in *TransactionHash, opts ...grpc.CallOption) (*Transaction, error)
	GetBalance(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error)
	GetState(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error)
	GetBlock(ctx context.Context, in *BlockKey, opts ...grpc.CallOption) (*BlockInfo, error)
	GetBlockByHeight(ctx context.Context, in *BlockKey, opts ...grpc.CallOption) (*BlockInfo, error)
	Transfer(ctx context.Context, in *TransInfo, opts ...grpc.CallOption) (*PublishRet, error)
}

type cliClient struct {
	cc *grpc.ClientConn
}

func NewCliClient(cc *grpc.ClientConn) CliClient {
	return &cliClient{cc}
}

func (c *cliClient) PublishTx(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*PublishRet, error) {
	out := new(PublishRet)
	err := c.cc.Invoke(ctx, "/rpc.Cli/PublishTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliClient) GetTransaction(ctx context.Context, in *TransactionKey, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/rpc.Cli/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliClient) GetTransactionByHash(ctx context.Context, in *TransactionHash, opts ...grpc.CallOption) (*Transaction, error) {
	out := new(Transaction)
	err := c.cc.Invoke(ctx, "/rpc.Cli/GetTransactionByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliClient) GetBalance(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, "/rpc.Cli/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliClient) GetState(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Value, error) {
	out := new(Value)
	err := c.cc.Invoke(ctx, "/rpc.Cli/GetState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliClient) GetBlock(ctx context.Context, in *BlockKey, opts ...grpc.CallOption) (*BlockInfo, error) {
	out := new(BlockInfo)
	err := c.cc.Invoke(ctx, "/rpc.Cli/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliClient) GetBlockByHeight(ctx context.Context, in *BlockKey, opts ...grpc.CallOption) (*BlockInfo, error) {
	out := new(BlockInfo)
	err := c.cc.Invoke(ctx, "/rpc.Cli/GetBlockByHeight", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliClient) Transfer(ctx context.Context, in *TransInfo, opts ...grpc.CallOption) (*PublishRet, error) {
	out := new(PublishRet)
	err := c.cc.Invoke(ctx, "/rpc.Cli/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cli service

type CliServer interface {
	PublishTx(context.Context, *Transaction) (*PublishRet, error)
	GetTransaction(context.Context, *TransactionKey) (*Transaction, error)
	GetTransactionByHash(context.Context, *TransactionHash) (*Transaction, error)
	GetBalance(context.Context, *Key) (*Value, error)
	GetState(context.Context, *Key) (*Value, error)
	GetBlock(context.Context, *BlockKey) (*BlockInfo, error)
	GetBlockByHeight(context.Context, *BlockKey) (*BlockInfo, error)
	Transfer(context.Context, *TransInfo) (*PublishRet, error)
}

func RegisterCliServer(s *grpc.Server, srv CliServer) {
	s.RegisterService(&_Cli_serviceDesc, srv)
}

func _Cli_PublishTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServer).PublishTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Cli/PublishTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServer).PublishTx(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cli_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Cli/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServer).GetTransaction(ctx, req.(*TransactionKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cli_GetTransactionByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServer).GetTransactionByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Cli/GetTransactionByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServer).GetTransactionByHash(ctx, req.(*TransactionHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cli_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Cli/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServer).GetBalance(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cli_GetState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServer).GetState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Cli/GetState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServer).GetState(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cli_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Cli/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServer).GetBlock(ctx, req.(*BlockKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cli_GetBlockByHeight_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServer).GetBlockByHeight(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Cli/GetBlockByHeight",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServer).GetBlockByHeight(ctx, req.(*BlockKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cli_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Cli/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliServer).Transfer(ctx, req.(*TransInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cli_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Cli",
	HandlerType: (*CliServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishTx",
			Handler:    _Cli_PublishTx_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _Cli_GetTransaction_Handler,
		},
		{
			MethodName: "GetTransactionByHash",
			Handler:    _Cli_GetTransactionByHash_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _Cli_GetBalance_Handler,
		},
		{
			MethodName: "GetState",
			Handler:    _Cli_GetState_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _Cli_GetBlock_Handler,
		},
		{
			MethodName: "GetBlockByHeight",
			Handler:    _Cli_GetBlockByHeight_Handler,
		},
		{
			MethodName: "Transfer",
			Handler:    _Cli_Transfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cli.proto",
}

func init() { proto.RegisterFile("cli.proto", fileDescriptor_cli_441c81ee1c4e0f3c) }

var fileDescriptor_cli_441c81ee1c4e0f3c = []byte{
	// 572 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x6f, 0x12, 0x41,
	0x14, 0x65, 0x59, 0xa0, 0xbb, 0xb7, 0x95, 0x36, 0xd3, 0x46, 0x37, 0xc4, 0x36, 0xcd, 0x44, 0x13,
	0x92, 0x46, 0x62, 0xa8, 0x2f, 0xbe, 0x19, 0x34, 0x01, 0x53, 0x1f, 0xcc, 0x8a, 0xbe, 0x0f, 0xcb,
	0xa5, 0x3b, 0xe9, 0x32, 0xbb, 0x99, 0x19, 0x10, 0x7e, 0x85, 0xff, 0xd5, 0x5f, 0x60, 0xe6, 0xee,
	0xf2, 0xd1, 0x8a, 0xf1, 0x6d, 0xce, 0xfd, 0x3a, 0xf7, 0x9e, 0x3d, 0x00, 0x61, 0x92, 0xc9, 0x5e,
	0xa1, 0x73, 0x9b, 0x33, 0x5f, 0x17, 0x09, 0xff, 0x0e, 0xe1, 0x58, 0x0b, 0x65, 0x3e, 0xab, 0x59,
	0xce, 0x9e, 0x43, 0xcb, 0x60, 0xf2, 0x80, 0xeb, 0xc8, 0xbb, 0xf6, 0xba, 0x61, 0x5c, 0x21, 0x76,
	0x01, 0x4d, 0x95, 0xab, 0x04, 0xa3, 0xfa, 0xb5, 0xd7, 0xf5, 0xe3, 0x12, 0xb0, 0x0e, 0x04, 0x49,
	0xae, 0xac, 0x16, 0x89, 0x8d, 0x7c, 0xaa, 0xdf, 0x62, 0x7e, 0x09, 0xc7, 0x34, 0x56, 0x24, 0x56,
	0xe6, 0x8a, 0xb5, 0xa1, 0x6e, 0x57, 0x34, 0xf4, 0x24, 0xae, 0xdb, 0x15, 0x7f, 0x07, 0xf0, 0x75,
	0x31, 0xc9, 0xa4, 0x49, 0x63, 0xb4, 0x8c, 0x41, 0x23, 0xc9, 0xa7, 0x48, 0xf9, 0x66, 0x4c, 0x6f,
	0x17, 0x4b, 0x85, 0x49, 0x89, 0xf1, 0x24, 0xa6, 0x37, 0xbf, 0x82, 0x20, 0x46, 0x53, 0xe4, 0xca,
	0xe0, 0xa1, 0x1e, 0xfe, 0x09, 0xda, 0x7b, 0xa4, 0x77, 0xb8, 0x66, 0x2f, 0x21, 0x2c, 0x4a, 0x1e,
	0xd4, 0x15, 0xfd, 0x2e, 0x70, 0xf8, 0x2c, 0xfe, 0x1a, 0x4e, 0xf7, 0xa6, 0x8c, 0x84, 0x49, 0xb7,
	0xcb, 0x78, 0x7b, 0xcb, 0x9c, 0x83, 0xef, 0x18, 0x4e, 0xc0, 0x33, 0x95, 0x5a, 0x9e, 0xe1, 0x2f,
	0xa0, 0xf9, 0x43, 0x64, 0x0b, 0x74, 0x07, 0x9b, 0x25, 0xcd, 0x0d, 0xe3, 0xba, 0x59, 0xf2, 0x6b,
	0x08, 0x06, 0x59, 0x9e, 0x3c, 0xdc, 0x95, 0x6a, 0x66, 0x62, 0x5d, 0x2d, 0xe4, 0xc7, 0x25, 0xe0,
	0xbf, 0x3d, 0x68, 0x8c, 0x50, 0x4c, 0x59, 0x04, 0x47, 0x4b, 0xd4, 0x46, 0xe6, 0xaa, 0x2a, 0xd8,
	0x40, 0x76, 0x05, 0x50, 0x08, 0x8d, 0xca, 0x8e, 0x76, 0xca, 0xec, 0x45, 0xdc, 0x07, 0xb1, 0x1a,
	0x91, 0xb2, 0x3e, 0x65, 0xb7, 0xd8, 0x29, 0x31, 0x71, 0x0b, 0x50, 0xb2, 0x51, 0x2a, 0xb1, 0x0d,
	0xb8, 0x03, 0xa5, 0x9a, 0xe5, 0x51, 0xb3, 0x3c, 0x50, 0x56, 0x66, 0x50, 0x8b, 0xf9, 0x04, 0x75,
	0xd4, 0xa2, 0x35, 0x2a, 0xe4, 0xf6, 0xfb, 0x29, 0xad, 0x42, 0x63, 0xa2, 0x23, 0xba, 0x6f, 0x03,
	0x1d, 0x87, 0x91, 0xf7, 0x4a, 0xd8, 0x85, 0xc6, 0x28, 0x28, 0x39, 0xb6, 0x01, 0xc7, 0x61, 0xe5,
	0x1c, 0xa3, 0x90, 0xa6, 0xd1, 0x9b, 0xcf, 0x21, 0x24, 0x59, 0xc8, 0x7d, 0x97, 0xd0, 0x48, 0x51,
	0x4c, 0xe9, 0xea, 0xe3, 0x7e, 0xd8, 0xd3, 0x45, 0xd2, 0x73, 0x8a, 0xc4, 0x14, 0x76, 0xb2, 0x8d,
	0x57, 0x89, 0xb2, 0x9b, 0xaf, 0x45, 0x80, 0xdd, 0x40, 0xcb, 0xae, 0xbe, 0x48, 0xe3, 0x2c, 0xe8,
	0x77, 0x8f, 0xfb, 0xe7, 0xd4, 0xf6, 0xd8, 0x06, 0x71, 0x55, 0xd2, 0xff, 0xe5, 0x83, 0xff, 0x31,
	0x93, 0xec, 0x2d, 0x84, 0x95, 0xfd, 0xc6, 0x2b, 0x76, 0xf6, 0xb4, 0xa3, 0x73, 0x4a, 0x91, 0x9d,
	0x41, 0x79, 0x8d, 0xbd, 0x87, 0xf6, 0x10, 0xed, 0xbe, 0xa5, 0x0f, 0x11, 0x75, 0xfe, 0x9a, 0xc5,
	0x6b, 0xec, 0x03, 0x5c, 0x3c, 0x6e, 0x1d, 0xac, 0x49, 0xf3, 0x8b, 0xa7, 0xb5, 0x2e, 0x7a, 0x70,
	0xc2, 0x2b, 0x80, 0x21, 0xda, 0x81, 0xc8, 0x84, 0xfb, 0xd9, 0x05, 0x54, 0xe1, 0xd8, 0x80, 0x5e,
	0x64, 0x38, 0x5e, 0x63, 0x1c, 0x82, 0x21, 0xda, 0x6f, 0x56, 0xd8, 0x7f, 0xd7, 0xdc, 0x50, 0x0d,
	0x49, 0xce, 0x9e, 0x51, 0x66, 0xe3, 0xca, 0x4e, 0x7b, 0x07, 0xdd, 0xd7, 0xe0, 0x35, 0x76, 0x0b,
	0x67, 0x9b, 0xe2, 0xc1, 0x7a, 0x84, 0xf2, 0x3e, 0xb5, 0xff, 0x6f, 0x7a, 0x03, 0x01, 0x2d, 0x3f,
	0x43, 0xcd, 0xda, 0xbb, 0x5b, 0x5c, 0xf6, 0x80, 0xae, 0x93, 0x16, 0xfd, 0x15, 0xdd, 0xfe, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x7f, 0xc1, 0xab, 0x86, 0x97, 0x04, 0x00, 0x00,
}
