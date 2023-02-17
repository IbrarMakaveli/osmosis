// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/gamm/pool-models/balancer/tx/tx.proto

package balancer

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// ===================== MsgCreatePool
type MsgCreateBalancerPool struct {
	Sender             string      `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	PoolParams         *PoolParams `protobuf:"bytes,2,opt,name=pool_params,json=poolParams,proto3" json:"pool_params,omitempty" yaml:"pool_params"`
	PoolAssets         []PoolAsset `protobuf:"bytes,3,rep,name=pool_assets,json=poolAssets,proto3" json:"pool_assets"`
	FuturePoolGovernor string      `protobuf:"bytes,4,opt,name=future_pool_governor,json=futurePoolGovernor,proto3" json:"future_pool_governor,omitempty" yaml:"future_pool_governor"`
}

func (m *MsgCreateBalancerPool) Reset()         { *m = MsgCreateBalancerPool{} }
func (m *MsgCreateBalancerPool) String() string { return proto.CompactTextString(m) }
func (*MsgCreateBalancerPool) ProtoMessage()    {}
func (*MsgCreateBalancerPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_0647ee155de97433, []int{0}
}
func (m *MsgCreateBalancerPool) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateBalancerPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateBalancerPool.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateBalancerPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateBalancerPool.Merge(m, src)
}
func (m *MsgCreateBalancerPool) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateBalancerPool) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateBalancerPool.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateBalancerPool proto.InternalMessageInfo

func (m *MsgCreateBalancerPool) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgCreateBalancerPool) GetPoolParams() *PoolParams {
	if m != nil {
		return m.PoolParams
	}
	return nil
}

func (m *MsgCreateBalancerPool) GetPoolAssets() []PoolAsset {
	if m != nil {
		return m.PoolAssets
	}
	return nil
}

func (m *MsgCreateBalancerPool) GetFuturePoolGovernor() string {
	if m != nil {
		return m.FuturePoolGovernor
	}
	return ""
}

// Returns the poolID
type MsgCreateBalancerPoolResponse struct {
	PoolID uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
}

func (m *MsgCreateBalancerPoolResponse) Reset()         { *m = MsgCreateBalancerPoolResponse{} }
func (m *MsgCreateBalancerPoolResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateBalancerPoolResponse) ProtoMessage()    {}
func (*MsgCreateBalancerPoolResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0647ee155de97433, []int{1}
}
func (m *MsgCreateBalancerPoolResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateBalancerPoolResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateBalancerPoolResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateBalancerPoolResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateBalancerPoolResponse.Merge(m, src)
}
func (m *MsgCreateBalancerPoolResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateBalancerPoolResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateBalancerPoolResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateBalancerPoolResponse proto.InternalMessageInfo

func (m *MsgCreateBalancerPoolResponse) GetPoolID() uint64 {
	if m != nil {
		return m.PoolID
	}
	return 0
}

// ===================== MsgMigrateSharesToFullRangeConcentratedPosition
type MsgMigrateSharesToFullRangeConcentratedPosition struct {
	Sender          string     `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty" yaml:"sender"`
	SharesToMigrate types.Coin `protobuf:"bytes,2,opt,name=shares_to_migrate,json=sharesToMigrate,proto3" json:"shares_to_migrate" yaml:"shares_to_migrate"`
}

func (m *MsgMigrateSharesToFullRangeConcentratedPosition) Reset() {
	*m = MsgMigrateSharesToFullRangeConcentratedPosition{}
}
func (m *MsgMigrateSharesToFullRangeConcentratedPosition) String() string {
	return proto.CompactTextString(m)
}
func (*MsgMigrateSharesToFullRangeConcentratedPosition) ProtoMessage() {}
func (*MsgMigrateSharesToFullRangeConcentratedPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_0647ee155de97433, []int{2}
}
func (m *MsgMigrateSharesToFullRangeConcentratedPosition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMigrateSharesToFullRangeConcentratedPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMigrateSharesToFullRangeConcentratedPosition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMigrateSharesToFullRangeConcentratedPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMigrateSharesToFullRangeConcentratedPosition.Merge(m, src)
}
func (m *MsgMigrateSharesToFullRangeConcentratedPosition) XXX_Size() int {
	return m.Size()
}
func (m *MsgMigrateSharesToFullRangeConcentratedPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMigrateSharesToFullRangeConcentratedPosition.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMigrateSharesToFullRangeConcentratedPosition proto.InternalMessageInfo

func (m *MsgMigrateSharesToFullRangeConcentratedPosition) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgMigrateSharesToFullRangeConcentratedPosition) GetSharesToMigrate() types.Coin {
	if m != nil {
		return m.SharesToMigrate
	}
	return types.Coin{}
}

type MsgMigrateSharesToFullRangeConcentratedPositionResponse struct {
	Amount0          github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=amount0,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount0" yaml:"amount0"`
	Amount1          github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount1,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount1" yaml:"amount1"`
	LiquidityCreated github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=liquidity_created,json=liquidityCreated,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"liquidity_created" yaml:"liquidity_created"`
}

func (m *MsgMigrateSharesToFullRangeConcentratedPositionResponse) Reset() {
	*m = MsgMigrateSharesToFullRangeConcentratedPositionResponse{}
}
func (m *MsgMigrateSharesToFullRangeConcentratedPositionResponse) String() string {
	return proto.CompactTextString(m)
}
func (*MsgMigrateSharesToFullRangeConcentratedPositionResponse) ProtoMessage() {}
func (*MsgMigrateSharesToFullRangeConcentratedPositionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0647ee155de97433, []int{3}
}
func (m *MsgMigrateSharesToFullRangeConcentratedPositionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMigrateSharesToFullRangeConcentratedPositionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMigrateSharesToFullRangeConcentratedPositionResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMigrateSharesToFullRangeConcentratedPositionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMigrateSharesToFullRangeConcentratedPositionResponse.Merge(m, src)
}
func (m *MsgMigrateSharesToFullRangeConcentratedPositionResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgMigrateSharesToFullRangeConcentratedPositionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMigrateSharesToFullRangeConcentratedPositionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMigrateSharesToFullRangeConcentratedPositionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateBalancerPool)(nil), "osmosis.gamm.poolmodels.balancer.v1beta1.MsgCreateBalancerPool")
	proto.RegisterType((*MsgCreateBalancerPoolResponse)(nil), "osmosis.gamm.poolmodels.balancer.v1beta1.MsgCreateBalancerPoolResponse")
	proto.RegisterType((*MsgMigrateSharesToFullRangeConcentratedPosition)(nil), "osmosis.gamm.poolmodels.balancer.v1beta1.MsgMigrateSharesToFullRangeConcentratedPosition")
	proto.RegisterType((*MsgMigrateSharesToFullRangeConcentratedPositionResponse)(nil), "osmosis.gamm.poolmodels.balancer.v1beta1.MsgMigrateSharesToFullRangeConcentratedPositionResponse")
}

func init() {
	proto.RegisterFile("osmosis/gamm/pool-models/balancer/tx/tx.proto", fileDescriptor_0647ee155de97433)
}

var fileDescriptor_0647ee155de97433 = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x4f, 0x6b, 0x13, 0x4f,
	0x18, 0xc7, 0xb3, 0x4d, 0x49, 0xf9, 0x4d, 0xf9, 0xa9, 0x5d, 0xaa, 0xc4, 0x88, 0xbb, 0x61, 0x05,
	0x89, 0x60, 0x66, 0x4c, 0x2d, 0x08, 0x5e, 0xaa, 0xdb, 0xd2, 0x52, 0x21, 0x50, 0x57, 0x2f, 0xed,
	0x25, 0x4c, 0x76, 0xc7, 0xed, 0xe2, 0xee, 0x4c, 0xdc, 0x99, 0xd4, 0xf6, 0x5d, 0x78, 0xf1, 0xea,
	0xeb, 0xf0, 0x0d, 0x08, 0x3d, 0xf6, 0x28, 0x1e, 0x16, 0x49, 0x6f, 0xe2, 0x29, 0xaf, 0x40, 0xe6,
	0x4f, 0xd2, 0xaa, 0x29, 0xba, 0x14, 0x4f, 0xdd, 0x3e, 0xf3, 0x7d, 0x3e, 0xdf, 0xe7, 0x79, 0xe6,
	0xc9, 0x80, 0x36, 0xe3, 0x19, 0xe3, 0x09, 0x47, 0x31, 0xce, 0x32, 0x34, 0x60, 0x2c, 0x6d, 0x67,
	0x2c, 0x22, 0x29, 0x47, 0x7d, 0x9c, 0x62, 0x1a, 0x92, 0x1c, 0x89, 0x43, 0x24, 0x0e, 0xe1, 0x20,
	0x67, 0x82, 0xd9, 0x2d, 0x23, 0x87, 0x52, 0x0e, 0xa5, 0x5c, 0xab, 0xe1, 0x44, 0x0d, 0x0f, 0x3a,
	0x7d, 0x22, 0x70, 0xa7, 0xb1, 0x1c, 0xb3, 0x98, 0xa9, 0x24, 0x24, 0xbf, 0x74, 0x7e, 0x63, 0xf5,
	0xcf, 0x76, 0x93, 0x8f, 0x1d, 0xc6, 0x52, 0x93, 0xe5, 0x84, 0x2a, 0x0d, 0xf5, 0x31, 0x27, 0xc8,
	0x18, 0xa0, 0x90, 0x25, 0x54, 0x9f, 0x7b, 0x1f, 0xe7, 0xc0, 0xf5, 0x2e, 0x8f, 0xd7, 0x73, 0x82,
	0x05, 0xf1, 0xcf, 0xe5, 0xdb, 0xf7, 0x40, 0x8d, 0x13, 0x1a, 0x91, 0xbc, 0x6e, 0x35, 0xad, 0xd6,
	0x7f, 0xfe, 0xd2, 0xb8, 0x70, 0xff, 0x3f, 0xc2, 0x59, 0xfa, 0xd8, 0xd3, 0x71, 0x2f, 0x30, 0x02,
	0x7b, 0x17, 0x2c, 0xca, 0x7a, 0x7a, 0x03, 0x9c, 0xe3, 0x8c, 0xd7, 0xe7, 0x9a, 0x56, 0x6b, 0x71,
	0xa5, 0x09, 0x7f, 0x6a, 0xd8, 0x78, 0x43, 0xc9, 0xde, 0x51, 0x3a, 0xff, 0xc6, 0xb8, 0x70, 0x6d,
	0x4d, 0x3c, 0x97, 0xee, 0x05, 0x60, 0x30, 0xd5, 0xd8, 0x9b, 0x06, 0x8d, 0x39, 0x27, 0x82, 0xd7,
	0xab, 0xcd, 0x6a, 0x6b, 0x71, 0xc5, 0xbd, 0x18, 0xfd, 0x54, 0xea, 0xfc, 0xf9, 0xe3, 0xc2, 0xad,
	0x68, 0x8e, 0x0a, 0x70, 0xfb, 0x39, 0x58, 0x7e, 0x35, 0x14, 0xc3, 0x9c, 0xf4, 0x14, 0x2e, 0x66,
	0x07, 0x24, 0xa7, 0x2c, 0xaf, 0xcf, 0xab, 0xde, 0xdc, 0x71, 0xe1, 0xde, 0xd2, 0x95, 0xcc, 0x52,
	0x79, 0x81, 0xad, 0xc3, 0xd2, 0x61, 0x6b, 0x12, 0xdc, 0x00, 0xb7, 0x67, 0x4e, 0x2e, 0x20, 0x7c,
	0xc0, 0x28, 0x27, 0xf6, 0x1d, 0xb0, 0xa0, 0x30, 0x49, 0xa4, 0x46, 0x38, 0xef, 0x83, 0x51, 0xe1,
	0xd6, 0xa4, 0x64, 0x7b, 0x23, 0xa8, 0xc9, 0xa3, 0xed, 0xc8, 0xfb, 0x64, 0x01, 0xd4, 0xe5, 0x71,
	0x37, 0x89, 0x73, 0x2c, 0xc8, 0x8b, 0x7d, 0x9c, 0x13, 0xfe, 0x92, 0x6d, 0x0e, 0xd3, 0x34, 0xc0,
	0x34, 0x26, 0xeb, 0x8c, 0x86, 0x84, 0x0a, 0x79, 0x16, 0xed, 0x30, 0x9e, 0x88, 0x84, 0xd1, 0x32,
	0x57, 0x13, 0x83, 0x25, 0xae, 0x98, 0x3d, 0xc1, 0x7a, 0x99, 0x36, 0x31, 0x17, 0x74, 0x13, 0xea,
	0xdd, 0x80, 0x72, 0x37, 0xa6, 0x43, 0x5c, 0x67, 0x09, 0xf5, 0x9b, 0x72, 0x7e, 0xe3, 0xc2, 0xad,
	0x1b, 0xe8, 0xaf, 0x04, 0x2f, 0xb8, 0xca, 0x4d, 0xa5, 0xa6, 0x70, 0xef, 0xdb, 0x1c, 0x78, 0x54,
	0xb2, 0x8f, 0xe9, 0xa0, 0xf6, 0xc0, 0x02, 0xce, 0xd8, 0x90, 0x8a, 0x07, 0xa6, 0xa1, 0x27, 0xd2,
	0xff, 0x4b, 0xe1, 0xde, 0x8d, 0x13, 0xb1, 0x3f, 0xec, 0xc3, 0x90, 0x65, 0xc8, 0x2c, 0xb2, 0xfe,
	0xd3, 0xe6, 0xd1, 0x6b, 0x24, 0x8e, 0x06, 0x84, 0xc3, 0x6d, 0x2a, 0xc6, 0x85, 0x7b, 0x45, 0x57,
	0x6a, 0x30, 0x5e, 0x30, 0x01, 0x9e, 0xb1, 0x3b, 0xaa, 0xed, 0x4b, 0xb3, 0x3b, 0x53, 0x76, 0xc7,
	0x7e, 0x0b, 0x96, 0xd2, 0xe4, 0xcd, 0x30, 0x89, 0x12, 0x71, 0xd4, 0x0b, 0xd5, 0x22, 0x44, 0xf5,
	0xaa, 0x72, 0x79, 0x56, 0xc2, 0x65, 0x83, 0x84, 0x67, 0xb3, 0xfe, 0x0d, 0xe8, 0x05, 0xd7, 0xa6,
	0x31, 0xbd, 0x6c, 0xd1, 0xca, 0xfb, 0x2a, 0xa8, 0x76, 0x79, 0x6c, 0x7f, 0xb0, 0x80, 0x3d, 0xe3,
	0xa7, 0xbb, 0x06, 0xff, 0xf6, 0xad, 0x81, 0x33, 0x37, 0xb8, 0xb1, 0x75, 0x49, 0xc0, 0xf4, 0x66,
	0xbf, 0x5b, 0xe0, 0x7e, 0xa9, 0xd5, 0xde, 0x2d, 0xe5, 0x5c, 0x06, 0xdd, 0xc0, 0xff, 0x0c, 0x3d,
	0x69, 0xd7, 0xdf, 0x3d, 0x1e, 0x39, 0xd6, 0xc9, 0xc8, 0xb1, 0xbe, 0x8e, 0x1c, 0xeb, 0xdd, 0xa9,
	0x53, 0x39, 0x39, 0x75, 0x2a, 0x9f, 0x4f, 0x9d, 0xca, 0xde, 0xda, 0xb9, 0x3d, 0x30, 0x65, 0xb4,
	0x53, 0xdc, 0xe7, 0x93, 0x7f, 0xd0, 0x41, 0x67, 0x15, 0x1d, 0x5e, 0xfc, 0xb6, 0xf7, 0x6b, 0xea,
	0xbd, 0x7e, 0xf8, 0x23, 0x00, 0x00, 0xff, 0xff, 0x8c, 0x6c, 0xd3, 0x3f, 0x76, 0x06, 0x00, 0x00,
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
	CreateBalancerPool(ctx context.Context, in *MsgCreateBalancerPool, opts ...grpc.CallOption) (*MsgCreateBalancerPoolResponse, error)
	MigrateSharesToFullRangeConcentratedPosition(ctx context.Context, in *MsgMigrateSharesToFullRangeConcentratedPosition, opts ...grpc.CallOption) (*MsgMigrateSharesToFullRangeConcentratedPositionResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateBalancerPool(ctx context.Context, in *MsgCreateBalancerPool, opts ...grpc.CallOption) (*MsgCreateBalancerPoolResponse, error) {
	out := new(MsgCreateBalancerPoolResponse)
	err := c.cc.Invoke(ctx, "/osmosis.gamm.poolmodels.balancer.v1beta1.Msg/CreateBalancerPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MigrateSharesToFullRangeConcentratedPosition(ctx context.Context, in *MsgMigrateSharesToFullRangeConcentratedPosition, opts ...grpc.CallOption) (*MsgMigrateSharesToFullRangeConcentratedPositionResponse, error) {
	out := new(MsgMigrateSharesToFullRangeConcentratedPositionResponse)
	err := c.cc.Invoke(ctx, "/osmosis.gamm.poolmodels.balancer.v1beta1.Msg/MigrateSharesToFullRangeConcentratedPosition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateBalancerPool(context.Context, *MsgCreateBalancerPool) (*MsgCreateBalancerPoolResponse, error)
	MigrateSharesToFullRangeConcentratedPosition(context.Context, *MsgMigrateSharesToFullRangeConcentratedPosition) (*MsgMigrateSharesToFullRangeConcentratedPositionResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateBalancerPool(ctx context.Context, req *MsgCreateBalancerPool) (*MsgCreateBalancerPoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBalancerPool not implemented")
}
func (*UnimplementedMsgServer) MigrateSharesToFullRangeConcentratedPosition(ctx context.Context, req *MsgMigrateSharesToFullRangeConcentratedPosition) (*MsgMigrateSharesToFullRangeConcentratedPositionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MigrateSharesToFullRangeConcentratedPosition not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateBalancerPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateBalancerPool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateBalancerPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.gamm.poolmodels.balancer.v1beta1.Msg/CreateBalancerPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateBalancerPool(ctx, req.(*MsgCreateBalancerPool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MigrateSharesToFullRangeConcentratedPosition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMigrateSharesToFullRangeConcentratedPosition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MigrateSharesToFullRangeConcentratedPosition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.gamm.poolmodels.balancer.v1beta1.Msg/MigrateSharesToFullRangeConcentratedPosition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MigrateSharesToFullRangeConcentratedPosition(ctx, req.(*MsgMigrateSharesToFullRangeConcentratedPosition))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.gamm.poolmodels.balancer.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBalancerPool",
			Handler:    _Msg_CreateBalancerPool_Handler,
		},
		{
			MethodName: "MigrateSharesToFullRangeConcentratedPosition",
			Handler:    _Msg_MigrateSharesToFullRangeConcentratedPosition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/gamm/pool-models/balancer/tx/tx.proto",
}

func (m *MsgCreateBalancerPool) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateBalancerPool) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateBalancerPool) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FuturePoolGovernor) > 0 {
		i -= len(m.FuturePoolGovernor)
		copy(dAtA[i:], m.FuturePoolGovernor)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FuturePoolGovernor)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PoolAssets) > 0 {
		for iNdEx := len(m.PoolAssets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolAssets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.PoolParams != nil {
		{
			size, err := m.PoolParams.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateBalancerPoolResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateBalancerPoolResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateBalancerPoolResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PoolID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.PoolID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgMigrateSharesToFullRangeConcentratedPosition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMigrateSharesToFullRangeConcentratedPosition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMigrateSharesToFullRangeConcentratedPosition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.SharesToMigrate.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgMigrateSharesToFullRangeConcentratedPositionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMigrateSharesToFullRangeConcentratedPositionResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMigrateSharesToFullRangeConcentratedPositionResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.LiquidityCreated.Size()
		i -= size
		if _, err := m.LiquidityCreated.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Amount1.Size()
		i -= size
		if _, err := m.Amount1.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Amount0.Size()
		i -= size
		if _, err := m.Amount0.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *MsgCreateBalancerPool) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.PoolParams != nil {
		l = m.PoolParams.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.PoolAssets) > 0 {
		for _, e := range m.PoolAssets {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.FuturePoolGovernor)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateBalancerPoolResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolID != 0 {
		n += 1 + sovTx(uint64(m.PoolID))
	}
	return n
}

func (m *MsgMigrateSharesToFullRangeConcentratedPosition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.SharesToMigrate.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgMigrateSharesToFullRangeConcentratedPositionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Amount0.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.Amount1.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.LiquidityCreated.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateBalancerPool) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateBalancerPool: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateBalancerPool: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolParams", wireType)
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
			if m.PoolParams == nil {
				m.PoolParams = &PoolParams{}
			}
			if err := m.PoolParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolAssets", wireType)
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
			m.PoolAssets = append(m.PoolAssets, PoolAsset{})
			if err := m.PoolAssets[len(m.PoolAssets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FuturePoolGovernor", wireType)
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
			m.FuturePoolGovernor = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *MsgCreateBalancerPoolResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateBalancerPoolResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateBalancerPoolResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolID", wireType)
			}
			m.PoolID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *MsgMigrateSharesToFullRangeConcentratedPosition) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgMigrateSharesToFullRangeConcentratedPosition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMigrateSharesToFullRangeConcentratedPosition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
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
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharesToMigrate", wireType)
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
			if err := m.SharesToMigrate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *MsgMigrateSharesToFullRangeConcentratedPositionResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgMigrateSharesToFullRangeConcentratedPositionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMigrateSharesToFullRangeConcentratedPositionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount0", wireType)
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
			if err := m.Amount0.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount1", wireType)
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
			if err := m.Amount1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LiquidityCreated", wireType)
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
			if err := m.LiquidityCreated.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
