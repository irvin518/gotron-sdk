// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/Contract.proto

/*
Package core is a generated protocol buffer package.

It is generated from these files:
	core/Contract.proto
	core/Discover.proto
	core/Tron.proto
	core/TronInventoryItems.proto

It has these top-level messages:
	AccountCreateContract
	AccountUpdateContract
	TransferContract
	TransferAssetContract
	VoteAssetContract
	VoteWitnessContract
	WitnessCreateContract
	WitnessUpdateContract
	AssetIssueContract
	ParticipateAssetIssueContract
	DeployContract
	FreezeBalanceContract
	UnfreezeBalanceContract
	UnfreezeAssetContract
	WithdrawBalanceContract
	UpdateAssetContract
	Endpoint
	PingMessage
	PongMessage
	FindNeighbours
	Neighbours
	AccountId
	Vote
	Account
	Acuthrity
	Permision
	Witness
	Votes
	TXOutput
	TXInput
	TXOutputs
	Transaction
	Transactions
	BlockHeader
	Block
	ChainInventory
	BlockInventory
	Inventory
	Items
	DynamicProperties
	DisconnectMessage
	HelloMessage
	InventoryItems
*/
package core

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AccountCreateContract struct {
	OwnerAddress   []byte      `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	AccountAddress []byte      `protobuf:"bytes,2,opt,name=account_address,json=accountAddress,proto3" json:"account_address,omitempty"`
	Type           AccountType `protobuf:"varint,3,opt,name=type,enum=protocol.AccountType" json:"type,omitempty"`
}

func (m *AccountCreateContract) Reset()                    { *m = AccountCreateContract{} }
func (m *AccountCreateContract) String() string            { return proto.CompactTextString(m) }
func (*AccountCreateContract) ProtoMessage()               {}
func (*AccountCreateContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AccountCreateContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *AccountCreateContract) GetAccountAddress() []byte {
	if m != nil {
		return m.AccountAddress
	}
	return nil
}

func (m *AccountCreateContract) GetType() AccountType {
	if m != nil {
		return m.Type
	}
	return AccountType_Normal
}

// update account name if the account has no name.
type AccountUpdateContract struct {
	AccountName  []byte `protobuf:"bytes,1,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	OwnerAddress []byte `protobuf:"bytes,2,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
}

func (m *AccountUpdateContract) Reset()                    { *m = AccountUpdateContract{} }
func (m *AccountUpdateContract) String() string            { return proto.CompactTextString(m) }
func (*AccountUpdateContract) ProtoMessage()               {}
func (*AccountUpdateContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AccountUpdateContract) GetAccountName() []byte {
	if m != nil {
		return m.AccountName
	}
	return nil
}

func (m *AccountUpdateContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

type TransferContract struct {
	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ToAddress    []byte `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount       int64  `protobuf:"varint,3,opt,name=amount" json:"amount,omitempty"`
}

func (m *TransferContract) Reset()                    { *m = TransferContract{} }
func (m *TransferContract) String() string            { return proto.CompactTextString(m) }
func (*TransferContract) ProtoMessage()               {}
func (*TransferContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TransferContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *TransferContract) GetToAddress() []byte {
	if m != nil {
		return m.ToAddress
	}
	return nil
}

func (m *TransferContract) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type TransferAssetContract struct {
	AssetName    []byte `protobuf:"bytes,1,opt,name=asset_name,json=assetName,proto3" json:"asset_name,omitempty"`
	OwnerAddress []byte `protobuf:"bytes,2,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ToAddress    []byte `protobuf:"bytes,3,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	Amount       int64  `protobuf:"varint,4,opt,name=amount" json:"amount,omitempty"`
}

func (m *TransferAssetContract) Reset()                    { *m = TransferAssetContract{} }
func (m *TransferAssetContract) String() string            { return proto.CompactTextString(m) }
func (*TransferAssetContract) ProtoMessage()               {}
func (*TransferAssetContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TransferAssetContract) GetAssetName() []byte {
	if m != nil {
		return m.AssetName
	}
	return nil
}

func (m *TransferAssetContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *TransferAssetContract) GetToAddress() []byte {
	if m != nil {
		return m.ToAddress
	}
	return nil
}

func (m *TransferAssetContract) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type VoteAssetContract struct {
	OwnerAddress []byte   `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	VoteAddress  [][]byte `protobuf:"bytes,2,rep,name=vote_address,json=voteAddress,proto3" json:"vote_address,omitempty"`
	Support      bool     `protobuf:"varint,3,opt,name=support" json:"support,omitempty"`
	Count        int32    `protobuf:"varint,5,opt,name=count" json:"count,omitempty"`
}

func (m *VoteAssetContract) Reset()                    { *m = VoteAssetContract{} }
func (m *VoteAssetContract) String() string            { return proto.CompactTextString(m) }
func (*VoteAssetContract) ProtoMessage()               {}
func (*VoteAssetContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *VoteAssetContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *VoteAssetContract) GetVoteAddress() [][]byte {
	if m != nil {
		return m.VoteAddress
	}
	return nil
}

func (m *VoteAssetContract) GetSupport() bool {
	if m != nil {
		return m.Support
	}
	return false
}

func (m *VoteAssetContract) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type VoteWitnessContract struct {
	OwnerAddress []byte                      `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Votes        []*VoteWitnessContract_Vote `protobuf:"bytes,2,rep,name=votes" json:"votes,omitempty"`
	Support      bool                        `protobuf:"varint,3,opt,name=support" json:"support,omitempty"`
}

func (m *VoteWitnessContract) Reset()                    { *m = VoteWitnessContract{} }
func (m *VoteWitnessContract) String() string            { return proto.CompactTextString(m) }
func (*VoteWitnessContract) ProtoMessage()               {}
func (*VoteWitnessContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *VoteWitnessContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *VoteWitnessContract) GetVotes() []*VoteWitnessContract_Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *VoteWitnessContract) GetSupport() bool {
	if m != nil {
		return m.Support
	}
	return false
}

type VoteWitnessContract_Vote struct {
	VoteAddress []byte `protobuf:"bytes,1,opt,name=vote_address,json=voteAddress,proto3" json:"vote_address,omitempty"`
	VoteCount   int64  `protobuf:"varint,2,opt,name=vote_count,json=voteCount" json:"vote_count,omitempty"`
}

func (m *VoteWitnessContract_Vote) Reset()                    { *m = VoteWitnessContract_Vote{} }
func (m *VoteWitnessContract_Vote) String() string            { return proto.CompactTextString(m) }
func (*VoteWitnessContract_Vote) ProtoMessage()               {}
func (*VoteWitnessContract_Vote) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

func (m *VoteWitnessContract_Vote) GetVoteAddress() []byte {
	if m != nil {
		return m.VoteAddress
	}
	return nil
}

func (m *VoteWitnessContract_Vote) GetVoteCount() int64 {
	if m != nil {
		return m.VoteCount
	}
	return 0
}

type WitnessCreateContract struct {
	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Url          []byte `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *WitnessCreateContract) Reset()                    { *m = WitnessCreateContract{} }
func (m *WitnessCreateContract) String() string            { return proto.CompactTextString(m) }
func (*WitnessCreateContract) ProtoMessage()               {}
func (*WitnessCreateContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *WitnessCreateContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *WitnessCreateContract) GetUrl() []byte {
	if m != nil {
		return m.Url
	}
	return nil
}

type WitnessUpdateContract struct {
	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	UpdateUrl    []byte `protobuf:"bytes,12,opt,name=update_url,json=updateUrl,proto3" json:"update_url,omitempty"`
}

func (m *WitnessUpdateContract) Reset()                    { *m = WitnessUpdateContract{} }
func (m *WitnessUpdateContract) String() string            { return proto.CompactTextString(m) }
func (*WitnessUpdateContract) ProtoMessage()               {}
func (*WitnessUpdateContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *WitnessUpdateContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *WitnessUpdateContract) GetUpdateUrl() []byte {
	if m != nil {
		return m.UpdateUrl
	}
	return nil
}

type AssetIssueContract struct {
	OwnerAddress            []byte                             `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Name                    []byte                             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Abbr                    []byte                             `protobuf:"bytes,3,opt,name=abbr,proto3" json:"abbr,omitempty"`
	TotalSupply             int64                              `protobuf:"varint,4,opt,name=total_supply,json=totalSupply" json:"total_supply,omitempty"`
	FrozenSupply            []*AssetIssueContract_FrozenSupply `protobuf:"bytes,5,rep,name=frozen_supply,json=frozenSupply" json:"frozen_supply,omitempty"`
	TrxNum                  int32                              `protobuf:"varint,6,opt,name=trx_num,json=trxNum" json:"trx_num,omitempty"`
	Num                     int32                              `protobuf:"varint,8,opt,name=num" json:"num,omitempty"`
	StartTime               int64                              `protobuf:"varint,9,opt,name=start_time,json=startTime" json:"start_time,omitempty"`
	EndTime                 int64                              `protobuf:"varint,10,opt,name=end_time,json=endTime" json:"end_time,omitempty"`
	VoteScore               int32                              `protobuf:"varint,16,opt,name=vote_score,json=voteScore" json:"vote_score,omitempty"`
	Description             []byte                             `protobuf:"bytes,20,opt,name=description,proto3" json:"description,omitempty"`
	Url                     []byte                             `protobuf:"bytes,21,opt,name=url,proto3" json:"url,omitempty"`
	FreeAssetNetLimit       int64                              `protobuf:"varint,22,opt,name=free_asset_net_limit,json=freeAssetNetLimit" json:"free_asset_net_limit,omitempty"`
	PublicFreeAssetNetLimit int64                              `protobuf:"varint,23,opt,name=public_free_asset_net_limit,json=publicFreeAssetNetLimit" json:"public_free_asset_net_limit,omitempty"`
	PublicFreeAssetNetUsage int64                              `protobuf:"varint,24,opt,name=public_free_asset_net_usage,json=publicFreeAssetNetUsage" json:"public_free_asset_net_usage,omitempty"`
	PublicLatestFreeNetTime int64                              `protobuf:"varint,25,opt,name=public_latest_free_net_time,json=publicLatestFreeNetTime" json:"public_latest_free_net_time,omitempty"`
}

func (m *AssetIssueContract) Reset()                    { *m = AssetIssueContract{} }
func (m *AssetIssueContract) String() string            { return proto.CompactTextString(m) }
func (*AssetIssueContract) ProtoMessage()               {}
func (*AssetIssueContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *AssetIssueContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *AssetIssueContract) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *AssetIssueContract) GetAbbr() []byte {
	if m != nil {
		return m.Abbr
	}
	return nil
}

func (m *AssetIssueContract) GetTotalSupply() int64 {
	if m != nil {
		return m.TotalSupply
	}
	return 0
}

func (m *AssetIssueContract) GetFrozenSupply() []*AssetIssueContract_FrozenSupply {
	if m != nil {
		return m.FrozenSupply
	}
	return nil
}

func (m *AssetIssueContract) GetTrxNum() int32 {
	if m != nil {
		return m.TrxNum
	}
	return 0
}

func (m *AssetIssueContract) GetNum() int32 {
	if m != nil {
		return m.Num
	}
	return 0
}

func (m *AssetIssueContract) GetStartTime() int64 {
	if m != nil {
		return m.StartTime
	}
	return 0
}

func (m *AssetIssueContract) GetEndTime() int64 {
	if m != nil {
		return m.EndTime
	}
	return 0
}

func (m *AssetIssueContract) GetVoteScore() int32 {
	if m != nil {
		return m.VoteScore
	}
	return 0
}

func (m *AssetIssueContract) GetDescription() []byte {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *AssetIssueContract) GetUrl() []byte {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *AssetIssueContract) GetFreeAssetNetLimit() int64 {
	if m != nil {
		return m.FreeAssetNetLimit
	}
	return 0
}

func (m *AssetIssueContract) GetPublicFreeAssetNetLimit() int64 {
	if m != nil {
		return m.PublicFreeAssetNetLimit
	}
	return 0
}

func (m *AssetIssueContract) GetPublicFreeAssetNetUsage() int64 {
	if m != nil {
		return m.PublicFreeAssetNetUsage
	}
	return 0
}

func (m *AssetIssueContract) GetPublicLatestFreeNetTime() int64 {
	if m != nil {
		return m.PublicLatestFreeNetTime
	}
	return 0
}

type AssetIssueContract_FrozenSupply struct {
	FrozenAmount int64 `protobuf:"varint,1,opt,name=frozen_amount,json=frozenAmount" json:"frozen_amount,omitempty"`
	FrozenDays   int64 `protobuf:"varint,2,opt,name=frozen_days,json=frozenDays" json:"frozen_days,omitempty"`
}

func (m *AssetIssueContract_FrozenSupply) Reset()         { *m = AssetIssueContract_FrozenSupply{} }
func (m *AssetIssueContract_FrozenSupply) String() string { return proto.CompactTextString(m) }
func (*AssetIssueContract_FrozenSupply) ProtoMessage()    {}
func (*AssetIssueContract_FrozenSupply) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{8, 0}
}

func (m *AssetIssueContract_FrozenSupply) GetFrozenAmount() int64 {
	if m != nil {
		return m.FrozenAmount
	}
	return 0
}

func (m *AssetIssueContract_FrozenSupply) GetFrozenDays() int64 {
	if m != nil {
		return m.FrozenDays
	}
	return 0
}

type ParticipateAssetIssueContract struct {
	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	ToAddress    []byte `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty"`
	AssetName    []byte `protobuf:"bytes,3,opt,name=asset_name,json=assetName,proto3" json:"asset_name,omitempty"`
	Amount       int64  `protobuf:"varint,4,opt,name=amount" json:"amount,omitempty"`
}

func (m *ParticipateAssetIssueContract) Reset()                    { *m = ParticipateAssetIssueContract{} }
func (m *ParticipateAssetIssueContract) String() string            { return proto.CompactTextString(m) }
func (*ParticipateAssetIssueContract) ProtoMessage()               {}
func (*ParticipateAssetIssueContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ParticipateAssetIssueContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *ParticipateAssetIssueContract) GetToAddress() []byte {
	if m != nil {
		return m.ToAddress
	}
	return nil
}

func (m *ParticipateAssetIssueContract) GetAssetName() []byte {
	if m != nil {
		return m.AssetName
	}
	return nil
}

func (m *ParticipateAssetIssueContract) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type DeployContract struct {
	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Script       []byte `protobuf:"bytes,2,opt,name=script,proto3" json:"script,omitempty"`
}

func (m *DeployContract) Reset()                    { *m = DeployContract{} }
func (m *DeployContract) String() string            { return proto.CompactTextString(m) }
func (*DeployContract) ProtoMessage()               {}
func (*DeployContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeployContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *DeployContract) GetScript() []byte {
	if m != nil {
		return m.Script
	}
	return nil
}

type FreezeBalanceContract struct {
	OwnerAddress   []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	FrozenBalance  int64  `protobuf:"varint,2,opt,name=frozen_balance,json=frozenBalance" json:"frozen_balance,omitempty"`
	FrozenDuration int64  `protobuf:"varint,3,opt,name=frozen_duration,json=frozenDuration" json:"frozen_duration,omitempty"`
}

func (m *FreezeBalanceContract) Reset()                    { *m = FreezeBalanceContract{} }
func (m *FreezeBalanceContract) String() string            { return proto.CompactTextString(m) }
func (*FreezeBalanceContract) ProtoMessage()               {}
func (*FreezeBalanceContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *FreezeBalanceContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *FreezeBalanceContract) GetFrozenBalance() int64 {
	if m != nil {
		return m.FrozenBalance
	}
	return 0
}

func (m *FreezeBalanceContract) GetFrozenDuration() int64 {
	if m != nil {
		return m.FrozenDuration
	}
	return 0
}

type UnfreezeBalanceContract struct {
	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
}

func (m *UnfreezeBalanceContract) Reset()                    { *m = UnfreezeBalanceContract{} }
func (m *UnfreezeBalanceContract) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeBalanceContract) ProtoMessage()               {}
func (*UnfreezeBalanceContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *UnfreezeBalanceContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

type UnfreezeAssetContract struct {
	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
}

func (m *UnfreezeAssetContract) Reset()                    { *m = UnfreezeAssetContract{} }
func (m *UnfreezeAssetContract) String() string            { return proto.CompactTextString(m) }
func (*UnfreezeAssetContract) ProtoMessage()               {}
func (*UnfreezeAssetContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *UnfreezeAssetContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

type WithdrawBalanceContract struct {
	OwnerAddress []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
}

func (m *WithdrawBalanceContract) Reset()                    { *m = WithdrawBalanceContract{} }
func (m *WithdrawBalanceContract) String() string            { return proto.CompactTextString(m) }
func (*WithdrawBalanceContract) ProtoMessage()               {}
func (*WithdrawBalanceContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *WithdrawBalanceContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

type UpdateAssetContract struct {
	OwnerAddress   []byte `protobuf:"bytes,1,opt,name=owner_address,json=ownerAddress,proto3" json:"owner_address,omitempty"`
	Description    []byte `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Url            []byte `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	NewLimit       int64  `protobuf:"varint,4,opt,name=new_limit,json=newLimit" json:"new_limit,omitempty"`
	NewPublicLimit int64  `protobuf:"varint,5,opt,name=new_public_limit,json=newPublicLimit" json:"new_public_limit,omitempty"`
}

func (m *UpdateAssetContract) Reset()                    { *m = UpdateAssetContract{} }
func (m *UpdateAssetContract) String() string            { return proto.CompactTextString(m) }
func (*UpdateAssetContract) ProtoMessage()               {}
func (*UpdateAssetContract) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *UpdateAssetContract) GetOwnerAddress() []byte {
	if m != nil {
		return m.OwnerAddress
	}
	return nil
}

func (m *UpdateAssetContract) GetDescription() []byte {
	if m != nil {
		return m.Description
	}
	return nil
}

func (m *UpdateAssetContract) GetUrl() []byte {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *UpdateAssetContract) GetNewLimit() int64 {
	if m != nil {
		return m.NewLimit
	}
	return 0
}

func (m *UpdateAssetContract) GetNewPublicLimit() int64 {
	if m != nil {
		return m.NewPublicLimit
	}
	return 0
}

func init() {
	proto.RegisterType((*AccountCreateContract)(nil), "protocol.AccountCreateContract")
	proto.RegisterType((*AccountUpdateContract)(nil), "protocol.AccountUpdateContract")
	proto.RegisterType((*TransferContract)(nil), "protocol.TransferContract")
	proto.RegisterType((*TransferAssetContract)(nil), "protocol.TransferAssetContract")
	proto.RegisterType((*VoteAssetContract)(nil), "protocol.VoteAssetContract")
	proto.RegisterType((*VoteWitnessContract)(nil), "protocol.VoteWitnessContract")
	proto.RegisterType((*VoteWitnessContract_Vote)(nil), "protocol.VoteWitnessContract.Vote")
	proto.RegisterType((*WitnessCreateContract)(nil), "protocol.WitnessCreateContract")
	proto.RegisterType((*WitnessUpdateContract)(nil), "protocol.WitnessUpdateContract")
	proto.RegisterType((*AssetIssueContract)(nil), "protocol.AssetIssueContract")
	proto.RegisterType((*AssetIssueContract_FrozenSupply)(nil), "protocol.AssetIssueContract.FrozenSupply")
	proto.RegisterType((*ParticipateAssetIssueContract)(nil), "protocol.ParticipateAssetIssueContract")
	proto.RegisterType((*DeployContract)(nil), "protocol.DeployContract")
	proto.RegisterType((*FreezeBalanceContract)(nil), "protocol.FreezeBalanceContract")
	proto.RegisterType((*UnfreezeBalanceContract)(nil), "protocol.UnfreezeBalanceContract")
	proto.RegisterType((*UnfreezeAssetContract)(nil), "protocol.UnfreezeAssetContract")
	proto.RegisterType((*WithdrawBalanceContract)(nil), "protocol.WithdrawBalanceContract")
	proto.RegisterType((*UpdateAssetContract)(nil), "protocol.UpdateAssetContract")
}

func init() { proto.RegisterFile("core/Contract.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 928 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x10, 0x06, 0x2d, 0x4b, 0x96, 0x46, 0x8a, 0xed, 0xd0, 0x96, 0xcd, 0x38, 0x30, 0xea, 0xb0, 0x28,
	0xea, 0x14, 0x88, 0x04, 0xa4, 0x97, 0x1e, 0x82, 0x02, 0xfe, 0x41, 0xda, 0x02, 0xa9, 0x10, 0x28,
	0x72, 0x03, 0xb4, 0x07, 0x62, 0x45, 0xae, 0x1c, 0x02, 0xe4, 0x2e, 0xb1, 0xbb, 0xac, 0xa2, 0xbc,
	0x42, 0x7b, 0x28, 0x7a, 0xec, 0x7b, 0xf4, 0x31, 0x7a, 0xef, 0xe3, 0x14, 0x3b, 0xbb, 0xab, 0xd0,
	0x72, 0x84, 0x4a, 0x3a, 0x89, 0xfc, 0xe6, 0x67, 0xbf, 0x9d, 0xf9, 0x66, 0x28, 0x38, 0x88, 0xb9,
	0xa0, 0xfd, 0x2b, 0xce, 0x94, 0x20, 0xb1, 0xea, 0x15, 0x82, 0x2b, 0xee, 0x37, 0xf1, 0x27, 0xe6,
	0xd9, 0xc9, 0x1e, 0x9a, 0x47, 0x82, 0x33, 0x63, 0x0a, 0xff, 0xf0, 0xa0, 0x7b, 0x11, 0xc7, 0xbc,
	0x64, 0xea, 0x4a, 0x50, 0xa2, 0xa8, 0x0b, 0xf5, 0x3f, 0x87, 0x07, 0x7c, 0xca, 0xa8, 0x88, 0x48,
	0x92, 0x08, 0x2a, 0x65, 0xe0, 0x9d, 0x79, 0xe7, 0x9d, 0x61, 0x07, 0xc1, 0x0b, 0x83, 0xf9, 0x5f,
	0xc2, 0x1e, 0x31, 0xd1, 0x73, 0xb7, 0x2d, 0x74, 0xdb, 0xb5, 0xb0, 0x73, 0x7c, 0x0a, 0xdb, 0x6a,
	0x56, 0xd0, 0xa0, 0x76, 0xe6, 0x9d, 0xef, 0x3e, 0xef, 0xf6, 0x1c, 0xa3, 0x9e, 0x3d, 0x7c, 0x34,
	0x2b, 0xe8, 0x10, 0x5d, 0xc2, 0x68, 0xce, 0xe8, 0xa6, 0x48, 0xaa, 0x8c, 0x9e, 0x40, 0xc7, 0x1d,
	0xc6, 0x48, 0x4e, 0x2d, 0xa1, 0xb6, 0xc5, 0x06, 0x24, 0xa7, 0xf7, 0x49, 0x6f, 0xdd, 0x27, 0x1d,
	0x32, 0xd8, 0x1f, 0x09, 0xc2, 0xe4, 0x84, 0x8a, 0xf5, 0x6e, 0x7b, 0x0a, 0xa0, 0xf8, 0x42, 0xea,
	0x96, 0xe2, 0xce, 0x7c, 0x04, 0x0d, 0x92, 0x6b, 0x2a, 0x78, 0xcb, 0xda, 0xd0, 0xbe, 0x85, 0x7f,
	0x7a, 0xd0, 0x75, 0x07, 0x5e, 0x48, 0x49, 0xd5, 0xfc, 0xd4, 0x53, 0x00, 0xa2, 0x81, 0xea, 0x7d,
	0x5a, 0x88, 0xac, 0x7c, 0x9b, 0x05, 0x52, 0xb5, 0xe5, 0xa4, 0xb6, 0xef, 0x90, 0xfa, 0xcd, 0x83,
	0x87, 0x3f, 0x71, 0x45, 0xef, 0x12, 0x5a, 0xa9, 0x0c, 0x4f, 0xa0, 0xf3, 0x2b, 0x57, 0xb4, 0xc2,
	0xaa, 0xa6, 0xfb, 0xa0, 0x31, 0xe7, 0x12, 0xc0, 0x8e, 0x2c, 0x8b, 0x82, 0x0b, 0x53, 0x8b, 0xe6,
	0xd0, 0xbd, 0xfa, 0x87, 0x50, 0xc7, 0x76, 0x05, 0xf5, 0x33, 0xef, 0xbc, 0x3e, 0x34, 0x2f, 0xe1,
	0xbf, 0x1e, 0x1c, 0x68, 0x36, 0x6f, 0x53, 0xc5, 0xa8, 0x94, 0xeb, 0xf1, 0xf9, 0x06, 0xea, 0xfa,
	0x6c, 0x43, 0xa4, 0xfd, 0x3c, 0xfc, 0x28, 0xae, 0x4f, 0xa4, 0x44, 0x6c, 0x68, 0x02, 0x96, 0xd3,
	0x3c, 0xf9, 0x1e, 0xb6, 0xb5, 0xe3, 0xbd, 0xbb, 0x5a, 0xcd, 0x55, 0xef, 0x7a, 0x0a, 0x80, 0x2e,
	0xe6, 0x5a, 0x5b, 0x58, 0xe5, 0x96, 0x46, 0xae, 0xf0, 0x6a, 0x03, 0xe8, 0x3a, 0x0a, 0x1b, 0x0c,
	0xd8, 0x3e, 0xd4, 0x4a, 0x91, 0xd9, 0xc6, 0xeb, 0xc7, 0xf0, 0x97, 0x79, 0xbe, 0x85, 0xf1, 0x58,
	0x55, 0xc2, 0x25, 0x86, 0x45, 0x3a, 0x6d, 0xc7, 0xa8, 0xc5, 0x20, 0x37, 0x22, 0x0b, 0xff, 0xa9,
	0x83, 0x8f, 0x8a, 0xf8, 0x41, 0xca, 0x72, 0xcd, 0xd4, 0x3e, 0x6c, 0xa3, 0x8c, 0x0d, 0x57, 0x7c,
	0xd6, 0x18, 0x19, 0x8f, 0x85, 0x95, 0x25, 0x3e, 0xeb, 0x92, 0x2a, 0xae, 0x48, 0x16, 0xe9, 0x5a,
	0x67, 0x33, 0xab, 0xcb, 0x36, 0x62, 0x6f, 0x10, 0xf2, 0x07, 0xf0, 0x60, 0x22, 0xf8, 0x07, 0xca,
	0x9c, 0x4f, 0x1d, 0x3b, 0xfb, 0xb4, 0xb2, 0x36, 0xee, 0x91, 0xec, 0xbd, 0xc4, 0x08, 0x93, 0x61,
	0xd8, 0x99, 0x54, 0xde, 0xfc, 0x63, 0xd8, 0x51, 0xe2, 0x7d, 0xc4, 0xca, 0x3c, 0x68, 0xa0, 0xec,
	0x1a, 0x4a, 0xbc, 0x1f, 0x94, 0xb9, 0x2e, 0xaf, 0x06, 0x9b, 0x08, 0xea, 0x47, 0x5d, 0x20, 0xa9,
	0x88, 0x50, 0x91, 0x4a, 0x73, 0x1a, 0xb4, 0x4c, 0x37, 0x11, 0x19, 0xa5, 0x39, 0xf5, 0x1f, 0x41,
	0x93, 0xb2, 0xc4, 0x18, 0x01, 0x8d, 0x3b, 0x94, 0x25, 0x68, 0x72, 0x3a, 0x90, 0x7a, 0xc7, 0x06,
	0xfb, 0x98, 0x12, 0x75, 0xf0, 0x46, 0x03, 0xfe, 0x19, 0xb4, 0x13, 0x2a, 0x63, 0x91, 0x16, 0x2a,
	0xe5, 0x2c, 0x38, 0x34, 0x42, 0xaa, 0x40, 0xae, 0xd7, 0xdd, 0x79, 0xaf, 0xfd, 0x3e, 0x1c, 0x4e,
	0x04, 0xa5, 0x91, 0x5d, 0x12, 0x54, 0x45, 0x59, 0x9a, 0xa7, 0x2a, 0x38, 0xc2, 0x93, 0x1f, 0x6a,
	0x1b, 0x16, 0x62, 0x40, 0xd5, 0x2b, 0x6d, 0xf0, 0x5f, 0xc0, 0xe3, 0xa2, 0x1c, 0x67, 0x69, 0x1c,
	0x7d, 0x32, 0xee, 0x18, 0xe3, 0x8e, 0x8d, 0xcb, 0xcb, 0xd5, 0xa3, 0x4b, 0x49, 0x6e, 0x69, 0x10,
	0x2c, 0x8b, 0xbe, 0xd1, 0xe6, 0x4a, 0x74, 0x46, 0x14, 0x95, 0xca, 0x24, 0xd1, 0xe1, 0x58, 0xad,
	0x47, 0xd5, 0xe8, 0x57, 0xe8, 0xa1, 0x73, 0x0c, 0x28, 0x16, 0xf6, 0x64, 0x04, 0x9d, 0x6a, 0x03,
	0xb5, 0xe4, 0xac, 0x04, 0xec, 0xfa, 0xf2, 0x30, 0xde, 0xf6, 0xf5, 0x02, 0x31, 0xff, 0x33, 0x68,
	0x5b, 0xa7, 0x84, 0xcc, 0xa4, 0x9d, 0x3d, 0x30, 0xd0, 0x35, 0x99, 0xc9, 0xf0, 0x2f, 0x0f, 0x4e,
	0x5f, 0x13, 0xa1, 0xd2, 0x38, 0x2d, 0x88, 0x5d, 0x76, 0x1b, 0x48, 0xfb, 0x7f, 0x16, 0xff, 0xdd,
	0x35, 0x5e, 0x5b, 0x5c, 0xe3, 0xcb, 0x56, 0xf0, 0x8f, 0xb0, 0x7b, 0x4d, 0x8b, 0x8c, 0xcf, 0xd6,
	0x23, 0x73, 0x04, 0x0d, 0xa3, 0x19, 0x4b, 0xc4, 0xbe, 0x85, 0xbf, 0x7b, 0xd0, 0xd5, 0x15, 0xfd,
	0x40, 0x2f, 0x49, 0x46, 0x58, 0xbc, 0xe6, 0x1d, 0xbf, 0x80, 0x5d, 0x5b, 0xcb, 0xb1, 0x09, 0xb7,
	0xe5, 0xb4, 0x6d, 0xb0, 0x39, 0xf5, 0x17, 0xdf, 0x95, 0xbc, 0x14, 0x04, 0xa5, 0x6c, 0xbe, 0x76,
	0x36, 0xfa, 0xda, 0xa2, 0xe1, 0xb7, 0x70, 0x7c, 0xc3, 0x26, 0x1b, 0xf3, 0x09, 0x5f, 0x40, 0xd7,
	0xc5, 0xaf, 0xff, 0x8d, 0xd2, 0xa7, 0xbf, 0x4d, 0xd5, 0xbb, 0x44, 0x90, 0xe9, 0x46, 0xa7, 0xff,
	0xed, 0xc1, 0x81, 0xd9, 0xaf, 0x1b, 0x7c, 0x20, 0x17, 0x46, 0x7d, 0x6b, 0xe9, 0xa8, 0xd7, 0x3e,
	0x8e, 0xfa, 0x63, 0x68, 0x31, 0x3a, 0xb5, 0x73, 0x6a, 0x74, 0xd2, 0x64, 0x74, 0x6a, 0x06, 0xf3,
	0x1c, 0xf6, 0xb5, 0xd1, 0x8d, 0x17, 0xfa, 0xd4, 0x4d, 0xd5, 0x19, 0x9d, 0xbe, 0x36, 0x23, 0xa5,
	0xd1, 0xcb, 0xef, 0x60, 0x8f, 0x8b, 0xdb, 0x9e, 0x9a, 0xff, 0xc3, 0x93, 0x97, 0x4d, 0x47, 0xfe,
	0xe7, 0xaf, 0x6e, 0x53, 0xf5, 0xae, 0x1c, 0xf7, 0x62, 0x9e, 0xf7, 0xb5, 0x87, 0xdb, 0xa6, 0xfd,
	0x5b, 0xfe, 0x2c, 0xce, 0x52, 0xca, 0xd4, 0x33, 0x52, 0xa4, 0x7d, 0xbd, 0xae, 0xc6, 0x0d, 0x34,
	0x7e, 0xfd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x34, 0x8f, 0xc8, 0x5d, 0x51, 0x0a, 0x00, 0x00,
}
