// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/state/types.proto

package state

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	types "github.com/tendermint/tendermint/abci/types"
	types1 "github.com/tendermint/tendermint/proto/types"
	version "github.com/tendermint/tendermint/proto/version"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ABCIResponses retains the responses
// of the various ABCI calls during block processing.
// It is persisted to disk for each height before calling Commit.
type ABCIResponses struct {
	DeliverTx            []*types.ResponseDeliverTx `protobuf:"bytes,1,rep,name=deliver_tx,json=deliverTx,proto3" json:"deliver_tx,omitempty"`
	EndBlock             *types.ResponseEndBlock    `protobuf:"bytes,2,opt,name=end_block,json=endBlock,proto3" json:"end_block,omitempty"`
	BeginBlock           *types.ResponseBeginBlock  `protobuf:"bytes,3,opt,name=begin_block,json=beginBlock,proto3" json:"begin_block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ABCIResponses) Reset()         { *m = ABCIResponses{} }
func (m *ABCIResponses) String() string { return proto.CompactTextString(m) }
func (*ABCIResponses) ProtoMessage()    {}
func (*ABCIResponses) Descriptor() ([]byte, []int) {
	return fileDescriptor_00e69fef8162ea9b, []int{0}
}
func (m *ABCIResponses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ABCIResponses.Unmarshal(m, b)
}
func (m *ABCIResponses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ABCIResponses.Marshal(b, m, deterministic)
}
func (m *ABCIResponses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ABCIResponses.Merge(m, src)
}
func (m *ABCIResponses) XXX_Size() int {
	return xxx_messageInfo_ABCIResponses.Size(m)
}
func (m *ABCIResponses) XXX_DiscardUnknown() {
	xxx_messageInfo_ABCIResponses.DiscardUnknown(m)
}

var xxx_messageInfo_ABCIResponses proto.InternalMessageInfo

func (m *ABCIResponses) GetDeliverTx() []*types.ResponseDeliverTx {
	if m != nil {
		return m.DeliverTx
	}
	return nil
}

func (m *ABCIResponses) GetEndBlock() *types.ResponseEndBlock {
	if m != nil {
		return m.EndBlock
	}
	return nil
}

func (m *ABCIResponses) GetBeginBlock() *types.ResponseBeginBlock {
	if m != nil {
		return m.BeginBlock
	}
	return nil
}

// ValidatorsInfo represents the latest validator set, or the last height it changed
type ValidatorsInfo struct {
	ValidatorSet         *types1.ValidatorSet `protobuf:"bytes,1,opt,name=validator_set,json=validatorSet,proto3" json:"validator_set,omitempty"`
	LastHeightChanged    int64                `protobuf:"varint,2,opt,name=last_height_changed,json=lastHeightChanged,proto3" json:"last_height_changed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ValidatorsInfo) Reset()         { *m = ValidatorsInfo{} }
func (m *ValidatorsInfo) String() string { return proto.CompactTextString(m) }
func (*ValidatorsInfo) ProtoMessage()    {}
func (*ValidatorsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_00e69fef8162ea9b, []int{1}
}
func (m *ValidatorsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValidatorsInfo.Unmarshal(m, b)
}
func (m *ValidatorsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValidatorsInfo.Marshal(b, m, deterministic)
}
func (m *ValidatorsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorsInfo.Merge(m, src)
}
func (m *ValidatorsInfo) XXX_Size() int {
	return xxx_messageInfo_ValidatorsInfo.Size(m)
}
func (m *ValidatorsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorsInfo proto.InternalMessageInfo

func (m *ValidatorsInfo) GetValidatorSet() *types1.ValidatorSet {
	if m != nil {
		return m.ValidatorSet
	}
	return nil
}

func (m *ValidatorsInfo) GetLastHeightChanged() int64 {
	if m != nil {
		return m.LastHeightChanged
	}
	return 0
}

// ConsensusParamsInfo represents the latest consensus params, or the last height it changed
type ConsensusParamsInfo struct {
	ConsensusParams      types1.ConsensusParams `protobuf:"bytes,1,opt,name=consensus_params,json=consensusParams,proto3" json:"consensus_params"`
	LastHeightChanged    int64                  `protobuf:"varint,2,opt,name=last_height_changed,json=lastHeightChanged,proto3" json:"last_height_changed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ConsensusParamsInfo) Reset()         { *m = ConsensusParamsInfo{} }
func (m *ConsensusParamsInfo) String() string { return proto.CompactTextString(m) }
func (*ConsensusParamsInfo) ProtoMessage()    {}
func (*ConsensusParamsInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_00e69fef8162ea9b, []int{2}
}
func (m *ConsensusParamsInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsensusParamsInfo.Unmarshal(m, b)
}
func (m *ConsensusParamsInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsensusParamsInfo.Marshal(b, m, deterministic)
}
func (m *ConsensusParamsInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusParamsInfo.Merge(m, src)
}
func (m *ConsensusParamsInfo) XXX_Size() int {
	return xxx_messageInfo_ConsensusParamsInfo.Size(m)
}
func (m *ConsensusParamsInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusParamsInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusParamsInfo proto.InternalMessageInfo

func (m *ConsensusParamsInfo) GetConsensusParams() types1.ConsensusParams {
	if m != nil {
		return m.ConsensusParams
	}
	return types1.ConsensusParams{}
}

func (m *ConsensusParamsInfo) GetLastHeightChanged() int64 {
	if m != nil {
		return m.LastHeightChanged
	}
	return 0
}

type Version struct {
	Consensus            version.Consensus `protobuf:"bytes,1,opt,name=consensus,proto3" json:"consensus"`
	Software             string            `protobuf:"bytes,2,opt,name=software,proto3" json:"software,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_00e69fef8162ea9b, []int{3}
}
func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (m *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(m, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetConsensus() version.Consensus {
	if m != nil {
		return m.Consensus
	}
	return version.Consensus{}
}

func (m *Version) GetSoftware() string {
	if m != nil {
		return m.Software
	}
	return ""
}

type State struct {
	Version Version `protobuf:"bytes,1,opt,name=version,proto3" json:"version"`
	// immutable
	ChainID string `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// LastBlockHeight=0 at genesis (ie. block(H=0) does not exist)
	LastBlockHeight  int64          `protobuf:"varint,3,opt,name=last_block_height,json=lastBlockHeight,proto3" json:"last_block_height,omitempty"`
	LastBlockTotalTx int64          `protobuf:"varint,4,opt,name=last_block_total_tx,json=lastBlockTotalTx,proto3" json:"last_block_total_tx,omitempty"`
	LastBlockID      types1.BlockID `protobuf:"bytes,5,opt,name=last_block_id,json=lastBlockId,proto3" json:"last_block_id"`
	LastBlockTime    time.Time      `protobuf:"bytes,6,opt,name=last_block_time,json=lastBlockTime,proto3,stdtime" json:"last_block_time"`
	// LastValidators is used to validate block.LastCommit.
	// Validators are persisted to the database separately every time they change,
	// so we can query for historical validator sets.
	// Note that if s.LastBlockHeight causes a valset change,
	// we set s.LastHeightValidatorsChanged = s.LastBlockHeight + 1 + 1
	// Extra +1 due to nextValSet delay.
	NextValidators              *types1.ValidatorSet `protobuf:"bytes,7,opt,name=next_validators,json=nextValidators,proto3" json:"next_validators,omitempty"`
	Validators                  *types1.ValidatorSet `protobuf:"bytes,8,opt,name=validators,proto3" json:"validators,omitempty"`
	LastValidators              *types1.ValidatorSet `protobuf:"bytes,9,opt,name=last_validators,json=lastValidators,proto3" json:"last_validators,omitempty"`
	LastHeightValidatorsChanged int64                `protobuf:"varint,20,opt,name=last_height_validators_changed,json=lastHeightValidatorsChanged,proto3" json:"last_height_validators_changed,omitempty"`
	// Consensus parameters used for validating blocks.
	// Changes returned by EndBlock and updated after Commit.
	ConsensusParams                  types1.ConsensusParams `protobuf:"bytes,12,opt,name=consensus_params,json=consensusParams,proto3" json:"consensus_params"`
	LastHeightConsensusParamsChanged int64                  `protobuf:"varint,13,opt,name=last_height_consensus_params_changed,json=lastHeightConsensusParamsChanged,proto3" json:"last_height_consensus_params_changed,omitempty"`
	// Merkle root of the results from executing prev block
	LastResultsHash []byte `protobuf:"bytes,14,opt,name=last_results_hash,json=lastResultsHash,proto3" json:"last_results_hash,omitempty"`
	// the latest AppHash we've received from calling abci.Commit()
	AppHash              []byte   `protobuf:"bytes,15,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_00e69fef8162ea9b, []int{4}
}
func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetVersion() Version {
	if m != nil {
		return m.Version
	}
	return Version{}
}

func (m *State) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *State) GetLastBlockHeight() int64 {
	if m != nil {
		return m.LastBlockHeight
	}
	return 0
}

func (m *State) GetLastBlockTotalTx() int64 {
	if m != nil {
		return m.LastBlockTotalTx
	}
	return 0
}

func (m *State) GetLastBlockID() types1.BlockID {
	if m != nil {
		return m.LastBlockID
	}
	return types1.BlockID{}
}

func (m *State) GetLastBlockTime() time.Time {
	if m != nil {
		return m.LastBlockTime
	}
	return time.Time{}
}

func (m *State) GetNextValidators() *types1.ValidatorSet {
	if m != nil {
		return m.NextValidators
	}
	return nil
}

func (m *State) GetValidators() *types1.ValidatorSet {
	if m != nil {
		return m.Validators
	}
	return nil
}

func (m *State) GetLastValidators() *types1.ValidatorSet {
	if m != nil {
		return m.LastValidators
	}
	return nil
}

func (m *State) GetLastHeightValidatorsChanged() int64 {
	if m != nil {
		return m.LastHeightValidatorsChanged
	}
	return 0
}

func (m *State) GetConsensusParams() types1.ConsensusParams {
	if m != nil {
		return m.ConsensusParams
	}
	return types1.ConsensusParams{}
}

func (m *State) GetLastHeightConsensusParamsChanged() int64 {
	if m != nil {
		return m.LastHeightConsensusParamsChanged
	}
	return 0
}

func (m *State) GetLastResultsHash() []byte {
	if m != nil {
		return m.LastResultsHash
	}
	return nil
}

func (m *State) GetAppHash() []byte {
	if m != nil {
		return m.AppHash
	}
	return nil
}

func init() {
	proto.RegisterType((*ABCIResponses)(nil), "tendermint.proto.state.ABCIResponses")
	proto.RegisterType((*ValidatorsInfo)(nil), "tendermint.proto.state.ValidatorsInfo")
	proto.RegisterType((*ConsensusParamsInfo)(nil), "tendermint.proto.state.ConsensusParamsInfo")
	proto.RegisterType((*Version)(nil), "tendermint.proto.state.Version")
	proto.RegisterType((*State)(nil), "tendermint.proto.state.State")
}

func init() { proto.RegisterFile("proto/state/types.proto", fileDescriptor_00e69fef8162ea9b) }

var fileDescriptor_00e69fef8162ea9b = []byte{
	// 751 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x6d, 0x6a, 0xdb, 0x4a,
	0x14, 0x7d, 0x7a, 0x4e, 0x62, 0x7b, 0x1c, 0xc7, 0x79, 0x93, 0x47, 0xaa, 0x3a, 0x50, 0x1b, 0x37,
	0x24, 0x6e, 0x69, 0x65, 0x48, 0x17, 0x50, 0x2a, 0xbb, 0x24, 0x2a, 0x69, 0x29, 0x8a, 0x09, 0xa1,
	0x7f, 0xc4, 0xd8, 0x9a, 0x48, 0x43, 0x6d, 0x49, 0x68, 0xc6, 0xae, 0xb3, 0x86, 0x52, 0xe8, 0x0e,
	0xba, 0x9d, 0xae, 0x22, 0x85, 0xfc, 0xec, 0x2a, 0xca, 0x7c, 0x48, 0x9a, 0x7c, 0x11, 0x0c, 0xfd,
	0xe5, 0xd1, 0xdc, 0x7b, 0xce, 0x3d, 0xf7, 0xce, 0x3d, 0x18, 0x3c, 0x4a, 0xd2, 0x98, 0xc5, 0x3d,
	0xca, 0x10, 0xc3, 0x3d, 0x76, 0x91, 0x60, 0x6a, 0x89, 0x1b, 0xb8, 0xcd, 0x70, 0xe4, 0xe3, 0x74,
	0x4a, 0x22, 0x26, 0x6f, 0x2c, 0x91, 0xd3, 0xdc, 0x63, 0x21, 0x49, 0x7d, 0x2f, 0x41, 0x29, 0xbb,
	0xe8, 0x49, 0x70, 0x10, 0x07, 0x71, 0x71, 0x92, 0xd9, 0xcd, 0x6d, 0x34, 0x1a, 0x13, 0xc9, 0xa8,
	0xf3, 0x36, 0x55, 0xc1, 0xdb, 0x81, 0x1d, 0x3d, 0x30, 0x47, 0x13, 0xe2, 0x23, 0x16, 0xa7, 0x2a,
	0x68, 0xea, 0xc1, 0x04, 0xa5, 0x68, 0x7a, 0x03, 0x36, 0xc7, 0x29, 0x25, 0x71, 0x94, 0xfd, 0xaa,
	0x60, 0x2b, 0x88, 0xe3, 0x60, 0x82, 0xa5, 0xce, 0xd1, 0xec, 0xbc, 0xc7, 0xc8, 0x14, 0x53, 0x86,
	0xa6, 0x89, 0x4c, 0xe8, 0xfc, 0x36, 0x40, 0xfd, 0x8d, 0xdd, 0x77, 0x5c, 0x4c, 0x93, 0x38, 0xa2,
	0x98, 0xc2, 0x43, 0x00, 0x7c, 0x3c, 0x21, 0x73, 0x9c, 0x7a, 0x6c, 0x61, 0x1a, 0xed, 0x52, 0xb7,
	0x76, 0xd0, 0xb5, 0xb4, 0x61, 0xf0, 0xbe, 0x2c, 0x29, 0x3c, 0x43, 0x0d, 0x24, 0x60, 0xb8, 0x70,
	0xab, 0x7e, 0x76, 0x84, 0x03, 0x50, 0xc5, 0x91, 0xef, 0x8d, 0x26, 0xf1, 0xf8, 0xb3, 0xf9, 0x6f,
	0xdb, 0xe8, 0xd6, 0x0e, 0xf6, 0x1f, 0xe0, 0x79, 0x1b, 0xf9, 0x36, 0x4f, 0x77, 0x2b, 0x58, 0x9d,
	0xe0, 0x3b, 0x50, 0x1b, 0xe1, 0x80, 0x44, 0x8a, 0xa7, 0x24, 0x78, 0x9e, 0x3d, 0xc0, 0x63, 0x73,
	0x84, 0x64, 0x02, 0xa3, 0xfc, 0xdc, 0xf9, 0x6a, 0x80, 0x8d, 0xd3, 0x6c, 0xb0, 0xd4, 0x89, 0xce,
	0x63, 0xe8, 0x80, 0x7a, 0x3e, 0x6a, 0x8f, 0x62, 0x66, 0x1a, 0xa2, 0xc0, 0xae, 0x75, 0xeb, 0xf5,
	0x65, 0x85, 0x1c, 0x7e, 0x82, 0x99, 0xbb, 0x3e, 0xd7, 0xbe, 0xa0, 0x05, 0xb6, 0x26, 0x88, 0x32,
	0x2f, 0xc4, 0x24, 0x08, 0x99, 0x37, 0x0e, 0x51, 0x14, 0x60, 0x5f, 0x74, 0x5e, 0x72, 0xff, 0xe3,
	0xa1, 0x23, 0x11, 0xe9, 0xcb, 0x40, 0xe7, 0x87, 0x01, 0xb6, 0xfa, 0x5c, 0x6d, 0x44, 0x67, 0xf4,
	0xa3, 0x78, 0x52, 0x21, 0xe9, 0x0c, 0x6c, 0x8e, 0xb3, 0x6b, 0x4f, 0x3e, 0xb5, 0x52, 0xb5, 0x7f,
	0x9f, 0xaa, 0x1b, 0x34, 0xf6, 0xca, 0xcf, 0xcb, 0xd6, 0x3f, 0x6e, 0x63, 0x7c, 0xfd, 0x7a, 0x69,
	0x85, 0x11, 0x28, 0x9f, 0xca, 0x75, 0x82, 0x87, 0xa0, 0x9a, 0xb3, 0x29, 0x35, 0x4f, 0x6f, 0xab,
	0xc9, 0x96, 0x2f, 0xd7, 0xa3, 0x94, 0x14, 0x58, 0xd8, 0x04, 0x15, 0x1a, 0x9f, 0xb3, 0x2f, 0x28,
	0xc5, 0xa2, 0x70, 0xd5, 0xcd, 0xbf, 0x3b, 0xdf, 0xca, 0x60, 0xf5, 0x84, 0x9b, 0x0c, 0xbe, 0x06,
	0x65, 0xc5, 0xa5, 0x8a, 0xb5, 0xac, 0xbb, 0xed, 0x68, 0x29, 0x81, 0xaa, 0x50, 0x86, 0x82, 0x7b,
	0xa0, 0x32, 0x0e, 0x11, 0x89, 0x3c, 0x22, 0xfb, 0xab, 0xda, 0xb5, 0xab, 0xcb, 0x56, 0xb9, 0xcf,
	0xef, 0x9c, 0x81, 0x5b, 0x16, 0x41, 0xc7, 0x87, 0xcf, 0x81, 0xe8, 0x5b, 0x6e, 0x97, 0x1a, 0x8c,
	0x58, 0xb2, 0x92, 0xdb, 0xe0, 0x01, 0xb1, 0x38, 0x72, 0x2a, 0xf0, 0xa5, 0x1a, 0x9f, 0xcc, 0x65,
	0x31, 0x43, 0x13, 0x6e, 0x91, 0x15, 0x91, 0xbd, 0x99, 0x67, 0x0f, 0x79, 0x60, 0xb8, 0x80, 0x67,
	0xa0, 0xae, 0xa5, 0x13, 0xdf, 0x5c, 0xbd, 0xaf, 0x13, 0xf9, 0x88, 0x02, 0xec, 0x0c, 0xec, 0x2d,
	0xde, 0xc9, 0xd5, 0x65, 0xab, 0x76, 0x9c, 0x31, 0x3a, 0x03, 0xb7, 0x96, 0xd3, 0x3b, 0x3e, 0x3c,
	0x06, 0x0d, 0x5d, 0x08, 0x99, 0x62, 0x73, 0x4d, 0x70, 0x37, 0x2d, 0xe9, 0x77, 0x2b, 0xf3, 0xbb,
	0x35, 0xcc, 0xfc, 0x6e, 0x57, 0x38, 0xed, 0xf7, 0x5f, 0x2d, 0xc3, 0xad, 0x17, 0x52, 0xc9, 0x14,
	0xc3, 0xf7, 0xa0, 0x11, 0xe1, 0x05, 0xf3, 0xf2, 0x65, 0xa6, 0x66, 0x79, 0x09, 0x13, 0x6c, 0x70,
	0x70, 0xe1, 0x2a, 0x38, 0x00, 0x40, 0x63, 0xaa, 0x2c, 0xc1, 0xa4, 0xe1, 0xb8, 0x28, 0xd1, 0xa2,
	0x46, 0x55, 0x5d, 0x46, 0x14, 0x07, 0x6b, 0xa2, 0xfa, 0xe0, 0x89, 0xbe, 0xf9, 0x05, 0x6b, 0x6e,
	0x82, 0xff, 0xc5, 0x2b, 0xee, 0x14, 0x26, 0x28, 0xd0, 0xca, 0x0e, 0x77, 0x1a, 0x73, 0xfd, 0xaf,
	0x18, 0xf3, 0x03, 0xd8, 0xbd, 0x66, 0xcc, 0x1b, 0x55, 0x72, 0x91, 0x75, 0x21, 0xb2, 0xad, 0x39,
	0xf5, 0x3a, 0x51, 0xa6, 0x34, 0xdb, 0xea, 0x14, 0xd3, 0xd9, 0x84, 0x51, 0x2f, 0x44, 0x34, 0x34,
	0x37, 0xda, 0x46, 0x77, 0x5d, 0x6e, 0xb5, 0x2b, 0xef, 0x8f, 0x10, 0x0d, 0xe1, 0x63, 0x50, 0x41,
	0x49, 0x22, 0x53, 0x1a, 0x22, 0xa5, 0x8c, 0x92, 0x84, 0x87, 0x6c, 0xeb, 0xd3, 0x8b, 0x80, 0xb0,
	0x70, 0x36, 0xb2, 0xc6, 0xf1, 0xb4, 0x57, 0xb4, 0xa8, 0x1f, 0xb5, 0xbf, 0xcf, 0xd1, 0x9a, 0xf8,
	0x78, 0xf5, 0x27, 0x00, 0x00, 0xff, 0xff, 0x90, 0xac, 0x9a, 0x4f, 0x54, 0x07, 0x00, 0x00,
}