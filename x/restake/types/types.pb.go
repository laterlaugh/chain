// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: restake/v1beta1/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Key message represents a key definition.
type Key struct {
	// Name is the name of the key.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Address is the address that holds rewards for this key.
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// IsActive is the status of the key
	IsActive bool `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// RewardPerShares is a list of reward_per_share of rewards.
	RewardPerShares github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,4,rep,name=reward_per_shares,json=rewardPerShares,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"reward_per_shares"`
	// TotalLock is the total locked power of the key.
	TotalLock github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=total_lock,json=totalLock,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"total_lock"`
	// Remainders is a list of the remainder amounts in the key.
	Remainders github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,6,rep,name=remainders,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"remainders"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_be4ee20adc0c7118, []int{0}
}
func (m *Key) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Key.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return m.Size()
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func (m *Key) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Key) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Key) GetIsActive() bool {
	if m != nil {
		return m.IsActive
	}
	return false
}

func (m *Key) GetRewardPerShares() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.RewardPerShares
	}
	return nil
}

func (m *Key) GetRemainders() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Remainders
	}
	return nil
}

// Stake message represents a stake detail.
type Stake struct {
	// Address is the owner's address of the stake.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Key is the key that this stake is locked to.
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	// Amount is the locked power amount.
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
	// PosRewardDebts is a list of reward debt for each reward (only the positive side).
	// Note: Coin and DecCoin can't have negative amounts. so, we split it into two numbers.
	PosRewardDebts github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=pos_reward_debts,json=posRewardDebts,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"pos_reward_debts"`
	// NegRewardDebts is a list of reward debt for each reward (only negative side).
	NegRewardDebts github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,5,rep,name=neg_reward_debts,json=negRewardDebts,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"neg_reward_debts"`
}

func (m *Stake) Reset()         { *m = Stake{} }
func (m *Stake) String() string { return proto.CompactTextString(m) }
func (*Stake) ProtoMessage()    {}
func (*Stake) Descriptor() ([]byte, []int) {
	return fileDescriptor_be4ee20adc0c7118, []int{1}
}
func (m *Stake) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Stake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Stake.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Stake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stake.Merge(m, src)
}
func (m *Stake) XXX_Size() int {
	return m.Size()
}
func (m *Stake) XXX_DiscardUnknown() {
	xxx_messageInfo_Stake.DiscardUnknown(m)
}

var xxx_messageInfo_Stake proto.InternalMessageInfo

func (m *Stake) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Stake) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Stake) GetPosRewardDebts() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.PosRewardDebts
	}
	return nil
}

func (m *Stake) GetNegRewardDebts() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.NegRewardDebts
	}
	return nil
}

// Reward message represents a reward detail.
type Reward struct {
	// Key is the key that this reward belongs to.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Rewards is a list of rewards.
	Rewards github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,2,rep,name=rewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"rewards"`
}

func (m *Reward) Reset()         { *m = Reward{} }
func (m *Reward) String() string { return proto.CompactTextString(m) }
func (*Reward) ProtoMessage()    {}
func (*Reward) Descriptor() ([]byte, []int) {
	return fileDescriptor_be4ee20adc0c7118, []int{2}
}
func (m *Reward) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Reward) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Reward.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Reward) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reward.Merge(m, src)
}
func (m *Reward) XXX_Size() int {
	return m.Size()
}
func (m *Reward) XXX_DiscardUnknown() {
	xxx_messageInfo_Reward.DiscardUnknown(m)
}

var xxx_messageInfo_Reward proto.InternalMessageInfo

func (m *Reward) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Reward) GetRewards() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.Rewards
	}
	return nil
}

// Lock message represents a lock definition.
type Lock struct {
	// Key is the key that this lock belongs to.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Amount is a the number of locked power.
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
}

func (m *Lock) Reset()         { *m = Lock{} }
func (m *Lock) String() string { return proto.CompactTextString(m) }
func (*Lock) ProtoMessage()    {}
func (*Lock) Descriptor() ([]byte, []int) {
	return fileDescriptor_be4ee20adc0c7118, []int{3}
}
func (m *Lock) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Lock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Lock.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Lock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Lock.Merge(m, src)
}
func (m *Lock) XXX_Size() int {
	return m.Size()
}
func (m *Lock) XXX_DiscardUnknown() {
	xxx_messageInfo_Lock.DiscardUnknown(m)
}

var xxx_messageInfo_Lock proto.InternalMessageInfo

func (m *Lock) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*Key)(nil), "restake.v1beta1.Key")
	proto.RegisterType((*Stake)(nil), "restake.v1beta1.Stake")
	proto.RegisterType((*Reward)(nil), "restake.v1beta1.Reward")
	proto.RegisterType((*Lock)(nil), "restake.v1beta1.Lock")
}

func init() { proto.RegisterFile("restake/v1beta1/types.proto", fileDescriptor_be4ee20adc0c7118) }

var fileDescriptor_be4ee20adc0c7118 = []byte{
	// 560 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x94, 0x31, 0x6f, 0x13, 0x31,
	0x14, 0xc7, 0xe3, 0x5e, 0x9a, 0x36, 0x46, 0xa2, 0xed, 0xa9, 0xc3, 0xb5, 0x45, 0x97, 0xa8, 0x03,
	0x8a, 0x40, 0xc9, 0xd1, 0x20, 0x24, 0x84, 0x58, 0x1a, 0xba, 0x14, 0x18, 0xd0, 0x85, 0x09, 0x86,
	0x93, 0xef, 0xce, 0xba, 0x58, 0xc9, 0xd9, 0x27, 0xdb, 0x0d, 0x84, 0x91, 0x4f, 0x80, 0xf8, 0x04,
	0xb0, 0x21, 0x26, 0x86, 0x7e, 0x88, 0x8a, 0xa9, 0xea, 0x84, 0x18, 0x0a, 0x4a, 0x06, 0xf8, 0x18,
	0xe8, 0x6c, 0x9f, 0x72, 0x08, 0x86, 0x48, 0x90, 0x25, 0xb1, 0xfd, 0x5e, 0xde, 0xef, 0xf9, 0xff,
	0xfe, 0x31, 0xdc, 0xe3, 0x58, 0x48, 0x34, 0xc4, 0xde, 0xf8, 0x20, 0xc4, 0x12, 0x1d, 0x78, 0x72,
	0x92, 0x61, 0xd1, 0xc9, 0x38, 0x93, 0xcc, 0xde, 0x30, 0xc1, 0x8e, 0x09, 0xee, 0x6e, 0x27, 0x2c,
	0x61, 0x2a, 0xe6, 0xe5, 0x2b, 0x9d, 0xb6, 0xbb, 0x85, 0x52, 0x42, 0x99, 0xa7, 0x3e, 0xcd, 0xd1,
	0x4e, 0xc4, 0x44, 0xca, 0x44, 0xa0, 0x73, 0xf5, 0xc6, 0x84, 0x5c, 0xbd, 0xf3, 0x42, 0x24, 0xe6,
	0xd4, 0x88, 0x11, 0xaa, 0xe3, 0xfb, 0x9f, 0x2d, 0x68, 0x3d, 0xc2, 0x13, 0xdb, 0x86, 0x55, 0x8a,
	0x52, 0xec, 0x80, 0x26, 0x68, 0xd5, 0x7d, 0xb5, 0xb6, 0xbb, 0x70, 0x0d, 0xc5, 0x31, 0xc7, 0x42,
	0x38, 0x2b, 0xf9, 0x71, 0xcf, 0xb9, 0x38, 0x6d, 0x6f, 0x9b, 0xf2, 0x87, 0x3a, 0xd2, 0x97, 0x9c,
	0xd0, 0xc4, 0x2f, 0x12, 0xed, 0x3d, 0x58, 0x27, 0x22, 0x40, 0x91, 0x24, 0x63, 0xec, 0x58, 0x4d,
	0xd0, 0x5a, 0xf7, 0xd7, 0x89, 0x38, 0x54, 0x7b, 0xfb, 0x35, 0x80, 0x5b, 0x1c, 0xbf, 0x40, 0x3c,
	0x0e, 0x32, 0xcc, 0x03, 0x31, 0x40, 0x1c, 0x0b, 0xa7, 0xda, 0xb4, 0x5a, 0x57, 0xba, 0xd7, 0x3a,
	0xa6, 0x70, 0xde, 0x69, 0x21, 0x41, 0xe7, 0x08, 0x47, 0x0f, 0x18, 0xa1, 0xbd, 0xbb, 0x67, 0x97,
	0x8d, 0xca, 0xc7, 0x6f, 0x8d, 0x9b, 0x09, 0x91, 0x83, 0x93, 0xb0, 0x13, 0xb1, 0xd4, 0xdc, 0xd3,
	0x7c, 0xb5, 0x45, 0x3c, 0x34, 0x6a, 0x9a, 0xdf, 0x88, 0x0f, 0x3f, 0x3e, 0xdd, 0x00, 0xfe, 0x86,
	0x06, 0x3e, 0xc1, 0xbc, 0xaf, 0x70, 0xf6, 0x73, 0x08, 0x25, 0x93, 0x68, 0x14, 0x8c, 0x58, 0x34,
	0x74, 0x56, 0xd5, 0xc5, 0xee, 0xe7, 0xe5, 0xbf, 0x5e, 0x36, 0xae, 0x2f, 0x50, 0xfe, 0x98, 0xca,
	0x8b, 0xd3, 0x36, 0x34, 0xdd, 0x1e, 0x53, 0xe9, 0xd7, 0x55, 0xbd, 0xc7, 0x2c, 0x1a, 0xda, 0x63,
	0x08, 0x39, 0x4e, 0x11, 0xa1, 0x31, 0xe6, 0xc2, 0xa9, 0x2d, 0xf5, 0x66, 0x25, 0xd2, 0xbd, 0xea,
	0xcf, 0x77, 0x0d, 0xb0, 0xff, 0xde, 0x82, 0xab, 0xfd, 0xdc, 0x42, 0xe5, 0xd1, 0x81, 0x45, 0x47,
	0xb7, 0x09, 0xad, 0x21, 0x9e, 0xe8, 0x51, 0xfb, 0xf9, 0xd2, 0x7e, 0x0a, 0x6b, 0x28, 0x65, 0x27,
	0x54, 0xaa, 0x49, 0xfe, 0xab, 0x4c, 0xa6, 0x96, 0xfd, 0x0a, 0x6e, 0x66, 0x4c, 0x04, 0xc6, 0x08,
	0x31, 0x0e, 0x65, 0xe1, 0x81, 0x9d, 0xbf, 0x2a, 0xa5, 0x64, 0xba, 0x63, 0x64, 0x6a, 0x2d, 0x80,
	0x2e, 0x69, 0x74, 0x35, 0x63, 0xc2, 0x57, 0xa0, 0xa3, 0x9c, 0x93, 0xb3, 0x29, 0x4e, 0x7e, 0x67,
	0xaf, 0x2e, 0x8b, 0x4d, 0x71, 0x52, 0x62, 0x9b, 0x19, 0xbd, 0x05, 0xb0, 0xa6, 0x4f, 0x0b, 0xc1,
	0xc1, 0x5c, 0xf0, 0x0c, 0xae, 0xe9, 0xd6, 0xf2, 0x7f, 0xdc, 0x32, 0xbd, 0x53, 0x60, 0x4c, 0x53,
	0x12, 0x56, 0x95, 0x7d, 0xff, 0xec, 0x68, 0x6e, 0x81, 0x95, 0xff, 0x67, 0x01, 0x4d, 0xed, 0x3d,
	0x3c, 0x9b, 0xba, 0xe0, 0x7c, 0xea, 0x82, 0xef, 0x53, 0x17, 0xbc, 0x99, 0xb9, 0x95, 0xf3, 0x99,
	0x5b, 0xf9, 0x32, 0x73, 0x2b, 0xcf, 0x6e, 0x95, 0xaa, 0x87, 0x88, 0xc6, 0xea, 0xad, 0x8a, 0xd8,
	0xc8, 0x8b, 0x06, 0x88, 0x50, 0x6f, 0xdc, 0xf5, 0x5e, 0x7a, 0xc5, 0x53, 0xaa, 0x58, 0x61, 0x4d,
	0xa5, 0xdc, 0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0xa7, 0xf5, 0xfc, 0xba, 0x62, 0x05, 0x00, 0x00,
}

func (this *Key) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Key)
	if !ok {
		that2, ok := that.(Key)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if this.IsActive != that1.IsActive {
		return false
	}
	if len(this.RewardPerShares) != len(that1.RewardPerShares) {
		return false
	}
	for i := range this.RewardPerShares {
		if !this.RewardPerShares[i].Equal(&that1.RewardPerShares[i]) {
			return false
		}
	}
	if !this.TotalLock.Equal(that1.TotalLock) {
		return false
	}
	if len(this.Remainders) != len(that1.Remainders) {
		return false
	}
	for i := range this.Remainders {
		if !this.Remainders[i].Equal(&that1.Remainders[i]) {
			return false
		}
	}
	return true
}
func (this *Stake) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Stake)
	if !ok {
		that2, ok := that.(Stake)
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
	if this.Address != that1.Address {
		return false
	}
	if this.Key != that1.Key {
		return false
	}
	if !this.Amount.Equal(that1.Amount) {
		return false
	}
	if len(this.PosRewardDebts) != len(that1.PosRewardDebts) {
		return false
	}
	for i := range this.PosRewardDebts {
		if !this.PosRewardDebts[i].Equal(&that1.PosRewardDebts[i]) {
			return false
		}
	}
	if len(this.NegRewardDebts) != len(that1.NegRewardDebts) {
		return false
	}
	for i := range this.NegRewardDebts {
		if !this.NegRewardDebts[i].Equal(&that1.NegRewardDebts[i]) {
			return false
		}
	}
	return true
}
func (this *Reward) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Reward)
	if !ok {
		that2, ok := that.(Reward)
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
	if this.Key != that1.Key {
		return false
	}
	if len(this.Rewards) != len(that1.Rewards) {
		return false
	}
	for i := range this.Rewards {
		if !this.Rewards[i].Equal(&that1.Rewards[i]) {
			return false
		}
	}
	return true
}
func (this *Lock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Lock)
	if !ok {
		that2, ok := that.(Lock)
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
	if this.Key != that1.Key {
		return false
	}
	if !this.Amount.Equal(that1.Amount) {
		return false
	}
	return true
}
func (m *Key) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Key) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Key) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Remainders) > 0 {
		for iNdEx := len(m.Remainders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Remainders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	{
		size := m.TotalLock.Size()
		i -= size
		if _, err := m.TotalLock.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.RewardPerShares) > 0 {
		for iNdEx := len(m.RewardPerShares) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RewardPerShares[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.IsActive {
		i--
		if m.IsActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Stake) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Stake) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Stake) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NegRewardDebts) > 0 {
		for iNdEx := len(m.NegRewardDebts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NegRewardDebts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.PosRewardDebts) > 0 {
		for iNdEx := len(m.PosRewardDebts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PosRewardDebts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Reward) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Reward) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Reward) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Lock) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Lock) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Lock) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Key) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.IsActive {
		n += 2
	}
	if len(m.RewardPerShares) > 0 {
		for _, e := range m.RewardPerShares {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = m.TotalLock.Size()
	n += 1 + l + sovTypes(uint64(l))
	if len(m.Remainders) > 0 {
		for _, e := range m.Remainders {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *Stake) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTypes(uint64(l))
	if len(m.PosRewardDebts) > 0 {
		for _, e := range m.PosRewardDebts {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.NegRewardDebts) > 0 {
		for _, e := range m.NegRewardDebts {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *Reward) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *Lock) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Key) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Key: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Key: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.IsActive = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RewardPerShares", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RewardPerShares = append(m.RewardPerShares, types.DecCoin{})
			if err := m.RewardPerShares[len(m.RewardPerShares)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalLock", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalLock.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Remainders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Remainders = append(m.Remainders, types.DecCoin{})
			if err := m.Remainders[len(m.Remainders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Stake) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Stake: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Stake: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PosRewardDebts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PosRewardDebts = append(m.PosRewardDebts, types.Coin{})
			if err := m.PosRewardDebts[len(m.PosRewardDebts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NegRewardDebts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NegRewardDebts = append(m.NegRewardDebts, types.Coin{})
			if err := m.NegRewardDebts[len(m.NegRewardDebts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Reward) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Reward: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Reward: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, types.DecCoin{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Lock) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Lock: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Lock: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
