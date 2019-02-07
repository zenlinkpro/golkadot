// Code generated by MockGen. DO NOT EDIT.
// Source: ./client/types/interface.go

// Package mock_clienttypes is a generated GoMock package.
package mock_clienttypes

import (
	types "github.com/opennetsys/go-substrate/client/p2p/handler/types"
	types0 "github.com/opennetsys/go-substrate/client/p2p/peer/types"
	types1 "github.com/opennetsys/go-substrate/client/p2p/peers/types"
	types2 "github.com/opennetsys/go-substrate/client/p2p/sync/types"
	types3 "github.com/opennetsys/go-substrate/client/p2p/types"
	types4 "github.com/opennetsys/go-substrate/client/types"
	crypto "github.com/opennetsys/go-substrate/common/crypto"
	gomock "github.com/golang/mock/gomock"
	go_libp2p_net "github.com/libp2p/go-libp2p-net"
	go_libp2p_peer "github.com/libp2p/go-libp2p-peer"
	go_libp2p_peerstore "github.com/libp2p/go-libp2p-peerstore"
	big "math/big"
	reflect "reflect"
)

// MockInterfaceSync is a mock of InterfaceSync interface
type MockInterfaceSync struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceSyncMockRecorder
}

// MockInterfaceSyncMockRecorder is the mock recorder for MockInterfaceSync
type MockInterfaceSyncMockRecorder struct {
	mock *MockInterfaceSync
}

// NewMockInterfaceSync creates a new mock instance
func NewMockInterfaceSync(ctrl *gomock.Controller) *MockInterfaceSync {
	mock := &MockInterfaceSync{ctrl: ctrl}
	mock.recorder = &MockInterfaceSyncMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterfaceSync) EXPECT() *MockInterfaceSyncMockRecorder {
	return m.recorder
}

// On mocks base method
func (m *MockInterfaceSync) On(event types2.EventEnum, cb types4.EventCallback) {
	m.ctrl.Call(m, "On", event, cb)
}

// On indicates an expected call of On
func (mr *MockInterfaceSyncMockRecorder) On(event, cb interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "On", reflect.TypeOf((*MockInterfaceSync)(nil).On), event, cb)
}

// QueueBlocks mocks base method
func (m *MockInterfaceSync) QueueBlocks(pr types4.InterfacePeer, response *types4.BlockResponse) error {
	ret := m.ctrl.Call(m, "QueueBlocks", pr, response)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueueBlocks indicates an expected call of QueueBlocks
func (mr *MockInterfaceSyncMockRecorder) QueueBlocks(pr, response interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueueBlocks", reflect.TypeOf((*MockInterfaceSync)(nil).QueueBlocks), pr, response)
}

// RequestBlocks mocks base method
func (m *MockInterfaceSync) RequestBlocks(pr types4.InterfacePeer) error {
	ret := m.ctrl.Call(m, "RequestBlocks", pr)
	ret0, _ := ret[0].(error)
	return ret0
}

// RequestBlocks indicates an expected call of RequestBlocks
func (mr *MockInterfaceSyncMockRecorder) RequestBlocks(pr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestBlocks", reflect.TypeOf((*MockInterfaceSync)(nil).RequestBlocks), pr)
}

// ProvideBlocks mocks base method
func (m *MockInterfaceSync) ProvideBlocks(pr types4.InterfacePeer, request *types4.BlockRequest) error {
	ret := m.ctrl.Call(m, "ProvideBlocks", pr, request)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProvideBlocks indicates an expected call of ProvideBlocks
func (mr *MockInterfaceSyncMockRecorder) ProvideBlocks(pr, request interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProvideBlocks", reflect.TypeOf((*MockInterfaceSync)(nil).ProvideBlocks), pr, request)
}

// PeerRequests mocks base method
func (m *MockInterfaceSync) PeerRequests(pr types4.InterfacePeer) (types4.Requests, error) {
	ret := m.ctrl.Call(m, "PeerRequests", pr)
	ret0, _ := ret[0].(types4.Requests)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PeerRequests indicates an expected call of PeerRequests
func (mr *MockInterfaceSyncMockRecorder) PeerRequests(pr interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PeerRequests", reflect.TypeOf((*MockInterfaceSync)(nil).PeerRequests), pr)
}

// MockInterfaceChains is a mock of InterfaceChains interface
type MockInterfaceChains struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceChainsMockRecorder
}

// MockInterfaceChainsMockRecorder is the mock recorder for MockInterfaceChains
type MockInterfaceChainsMockRecorder struct {
	mock *MockInterfaceChains
}

// NewMockInterfaceChains creates a new mock instance
func NewMockInterfaceChains(ctrl *gomock.Controller) *MockInterfaceChains {
	mock := &MockInterfaceChains{ctrl: ctrl}
	mock.recorder = &MockInterfaceChainsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterfaceChains) EXPECT() *MockInterfaceChainsMockRecorder {
	return m.recorder
}

// GetBestBlocksNumber mocks base method
func (m *MockInterfaceChains) GetBestBlocksNumber() (*big.Int, error) {
	ret := m.ctrl.Call(m, "GetBestBlocksNumber")
	ret0, _ := ret[0].(*big.Int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBestBlocksNumber indicates an expected call of GetBestBlocksNumber
func (mr *MockInterfaceChainsMockRecorder) GetBestBlocksNumber() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBestBlocksNumber", reflect.TypeOf((*MockInterfaceChains)(nil).GetBestBlocksNumber))
}

// GetBestBlocksHash mocks base method
func (m *MockInterfaceChains) GetBestBlocksHash() (*crypto.Blake2b256Hash, error) {
	ret := m.ctrl.Call(m, "GetBestBlocksHash")
	ret0, _ := ret[0].(*crypto.Blake2b256Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBestBlocksHash indicates an expected call of GetBestBlocksHash
func (mr *MockInterfaceChainsMockRecorder) GetBestBlocksHash() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBestBlocksHash", reflect.TypeOf((*MockInterfaceChains)(nil).GetBestBlocksHash))
}

// GetGenesisHash mocks base method
func (m *MockInterfaceChains) GetGenesisHash() (*crypto.Blake2b256Hash, error) {
	ret := m.ctrl.Call(m, "GetGenesisHash")
	ret0, _ := ret[0].(*crypto.Blake2b256Hash)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGenesisHash indicates an expected call of GetGenesisHash
func (mr *MockInterfaceChainsMockRecorder) GetGenesisHash() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenesisHash", reflect.TypeOf((*MockInterfaceChains)(nil).GetGenesisHash))
}

// ImportBlock mocks base method
func (m *MockInterfaceChains) ImportBlock(block *types4.StateBlock) (bool, error) {
	ret := m.ctrl.Call(m, "ImportBlock", block)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ImportBlock indicates an expected call of ImportBlock
func (mr *MockInterfaceChainsMockRecorder) ImportBlock(block interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportBlock", reflect.TypeOf((*MockInterfaceChains)(nil).ImportBlock), block)
}

// GetBlockDataByHash mocks base method
func (m *MockInterfaceChains) GetBlockDataByHash(hash *crypto.Blake2b256Hash) (*types4.StateBlock, error) {
	ret := m.ctrl.Call(m, "GetBlockDataByHash", hash)
	ret0, _ := ret[0].(*types4.StateBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBlockDataByHash indicates an expected call of GetBlockDataByHash
func (mr *MockInterfaceChainsMockRecorder) GetBlockDataByHash(hash interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlockDataByHash", reflect.TypeOf((*MockInterfaceChains)(nil).GetBlockDataByHash), hash)
}

// MockInterfacePeer is a mock of InterfacePeer interface
type MockInterfacePeer struct {
	ctrl     *gomock.Controller
	recorder *MockInterfacePeerMockRecorder
}

// MockInterfacePeerMockRecorder is the mock recorder for MockInterfacePeer
type MockInterfacePeerMockRecorder struct {
	mock *MockInterfacePeer
}

// NewMockInterfacePeer creates a new mock instance
func NewMockInterfacePeer(ctrl *gomock.Controller) *MockInterfacePeer {
	mock := &MockInterfacePeer{ctrl: ctrl}
	mock.recorder = &MockInterfacePeerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterfacePeer) EXPECT() *MockInterfacePeerMockRecorder {
	return m.recorder
}

// AddConnection mocks base method
func (m *MockInterfacePeer) AddConnection(conn go_libp2p_net.Conn, isWritable bool) (uint, error) {
	ret := m.ctrl.Call(m, "AddConnection", conn, isWritable)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddConnection indicates an expected call of AddConnection
func (mr *MockInterfacePeerMockRecorder) AddConnection(conn, isWritable interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddConnection", reflect.TypeOf((*MockInterfacePeer)(nil).AddConnection), conn, isWritable)
}

// Disconnect mocks base method
func (m *MockInterfacePeer) Disconnect() error {
	ret := m.ctrl.Call(m, "Disconnect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Disconnect indicates an expected call of Disconnect
func (mr *MockInterfacePeerMockRecorder) Disconnect() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockInterfacePeer)(nil).Disconnect))
}

// IsActive mocks base method
func (m *MockInterfacePeer) IsActive() (bool, error) {
	ret := m.ctrl.Call(m, "IsActive")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsActive indicates an expected call of IsActive
func (mr *MockInterfacePeerMockRecorder) IsActive() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsActive", reflect.TypeOf((*MockInterfacePeer)(nil).IsActive))
}

// IsWritable mocks base method
func (m *MockInterfacePeer) IsWritable() (bool, error) {
	ret := m.ctrl.Call(m, "IsWritable")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsWritable indicates an expected call of IsWritable
func (mr *MockInterfacePeerMockRecorder) IsWritable() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsWritable", reflect.TypeOf((*MockInterfacePeer)(nil).IsWritable))
}

// On mocks base method
func (m *MockInterfacePeer) On(event types0.EventEnum, cb types4.PeerEventCallback) {
	m.ctrl.Call(m, "On", event, cb)
}

// On indicates an expected call of On
func (mr *MockInterfacePeerMockRecorder) On(event, cb interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "On", reflect.TypeOf((*MockInterfacePeer)(nil).On), event, cb)
}

// Send mocks base method
func (m *MockInterfacePeer) Send(msg types4.InterfaceMessage) (bool, error) {
	ret := m.ctrl.Call(m, "Send", msg)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send
func (mr *MockInterfacePeerMockRecorder) Send(msg interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockInterfacePeer)(nil).Send), msg)
}

// SetBest mocks base method
func (m *MockInterfacePeer) SetBest(blockNumber *big.Int, hash []byte) error {
	ret := m.ctrl.Call(m, "SetBest", blockNumber, hash)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBest indicates an expected call of SetBest
func (mr *MockInterfacePeerMockRecorder) SetBest(blockNumber, hash interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBest", reflect.TypeOf((*MockInterfacePeer)(nil).SetBest), blockNumber, hash)
}

// Cfg mocks base method
func (m *MockInterfacePeer) Cfg() types4.ConfigClient {
	ret := m.ctrl.Call(m, "Cfg")
	ret0, _ := ret[0].(types4.ConfigClient)
	return ret0
}

// Cfg indicates an expected call of Cfg
func (mr *MockInterfacePeerMockRecorder) Cfg() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cfg", reflect.TypeOf((*MockInterfacePeer)(nil).Cfg))
}

// GetChain mocks base method
func (m *MockInterfacePeer) GetChain() (types4.InterfaceChains, error) {
	ret := m.ctrl.Call(m, "GetChain")
	ret0, _ := ret[0].(types4.InterfaceChains)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChain indicates an expected call of GetChain
func (mr *MockInterfacePeerMockRecorder) GetChain() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChain", reflect.TypeOf((*MockInterfacePeer)(nil).GetChain))
}

// GetID mocks base method
func (m *MockInterfacePeer) GetID() string {
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetID indicates an expected call of GetID
func (mr *MockInterfacePeerMockRecorder) GetID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockInterfacePeer)(nil).GetID))
}

// GetNextID mocks base method
func (m *MockInterfacePeer) GetNextID() uint {
	ret := m.ctrl.Call(m, "GetNextID")
	ret0, _ := ret[0].(uint)
	return ret0
}

// GetNextID indicates an expected call of GetNextID
func (mr *MockInterfacePeerMockRecorder) GetNextID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNextID", reflect.TypeOf((*MockInterfacePeer)(nil).GetNextID))
}

// GetPeerInfo mocks base method
func (m *MockInterfacePeer) GetPeerInfo() go_libp2p_peerstore.PeerInfo {
	ret := m.ctrl.Call(m, "GetPeerInfo")
	ret0, _ := ret[0].(go_libp2p_peerstore.PeerInfo)
	return ret0
}

// GetPeerInfo indicates an expected call of GetPeerInfo
func (mr *MockInterfacePeerMockRecorder) GetPeerInfo() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPeerInfo", reflect.TypeOf((*MockInterfacePeer)(nil).GetPeerInfo))
}

// GetShortID mocks base method
func (m *MockInterfacePeer) GetShortID() string {
	ret := m.ctrl.Call(m, "GetShortID")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetShortID indicates an expected call of GetShortID
func (mr *MockInterfacePeerMockRecorder) GetShortID() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShortID", reflect.TypeOf((*MockInterfacePeer)(nil).GetShortID))
}

// Receive mocks base method
func (m *MockInterfacePeer) Receive(stream go_libp2p_net.Stream) error {
	ret := m.ctrl.Call(m, "Receive", stream)
	ret0, _ := ret[0].(error)
	return ret0
}

// Receive indicates an expected call of Receive
func (mr *MockInterfacePeerMockRecorder) Receive(stream interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Receive", reflect.TypeOf((*MockInterfacePeer)(nil).Receive), stream)
}

// GetBestNumber mocks base method
func (m *MockInterfacePeer) GetBestNumber() *big.Int {
	ret := m.ctrl.Call(m, "GetBestNumber")
	ret0, _ := ret[0].(*big.Int)
	return ret0
}

// GetBestNumber indicates an expected call of GetBestNumber
func (mr *MockInterfacePeerMockRecorder) GetBestNumber() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBestNumber", reflect.TypeOf((*MockInterfacePeer)(nil).GetBestNumber))
}

// MockInterfacePeers is a mock of InterfacePeers interface
type MockInterfacePeers struct {
	ctrl     *gomock.Controller
	recorder *MockInterfacePeersMockRecorder
}

// MockInterfacePeersMockRecorder is the mock recorder for MockInterfacePeers
type MockInterfacePeersMockRecorder struct {
	mock *MockInterfacePeers
}

// NewMockInterfacePeers creates a new mock instance
func NewMockInterfacePeers(ctrl *gomock.Controller) *MockInterfacePeers {
	mock := &MockInterfacePeers{ctrl: ctrl}
	mock.recorder = &MockInterfacePeersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterfacePeers) EXPECT() *MockInterfacePeersMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockInterfacePeers) Add(pi go_libp2p_peerstore.PeerInfo) (*types4.KnownPeer, error) {
	ret := m.ctrl.Call(m, "Add", pi)
	ret0, _ := ret[0].(*types4.KnownPeer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add
func (mr *MockInterfacePeersMockRecorder) Add(pi interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockInterfacePeers)(nil).Add), pi)
}

// CountAll mocks base method
func (m *MockInterfacePeers) CountAll() (uint, error) {
	ret := m.ctrl.Call(m, "CountAll")
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountAll indicates an expected call of CountAll
func (mr *MockInterfacePeersMockRecorder) CountAll() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountAll", reflect.TypeOf((*MockInterfacePeers)(nil).CountAll))
}

// Count mocks base method
func (m *MockInterfacePeers) Count() (uint, error) {
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockInterfacePeersMockRecorder) Count() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockInterfacePeers)(nil).Count))
}

// Get mocks base method
func (m *MockInterfacePeers) Get(pi go_libp2p_peerstore.PeerInfo) (*types4.KnownPeer, error) {
	ret := m.ctrl.Call(m, "Get", pi)
	ret0, _ := ret[0].(*types4.KnownPeer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockInterfacePeersMockRecorder) Get(pi interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterfacePeers)(nil).Get), pi)
}

// GetFromID mocks base method
func (m *MockInterfacePeers) GetFromID(id go_libp2p_peer.ID) (*types4.KnownPeer, error) {
	ret := m.ctrl.Call(m, "GetFromID", id)
	ret0, _ := ret[0].(*types4.KnownPeer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFromID indicates an expected call of GetFromID
func (mr *MockInterfacePeersMockRecorder) GetFromID(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFromID", reflect.TypeOf((*MockInterfacePeers)(nil).GetFromID), id)
}

// Log mocks base method
func (m *MockInterfacePeers) Log(event types1.EventEnum, iface interface{}) error {
	ret := m.ctrl.Call(m, "Log", event, iface)
	ret0, _ := ret[0].(error)
	return ret0
}

// Log indicates an expected call of Log
func (mr *MockInterfacePeersMockRecorder) Log(event, iface interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Log", reflect.TypeOf((*MockInterfacePeers)(nil).Log), event, iface)
}

// On mocks base method
func (m *MockInterfacePeers) On(event types1.EventEnum, cb types4.PeersEventCallback) {
	m.ctrl.Call(m, "On", event, cb)
}

// On indicates an expected call of On
func (mr *MockInterfacePeersMockRecorder) On(event, cb interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "On", reflect.TypeOf((*MockInterfacePeers)(nil).On), event, cb)
}

// KnownPeers mocks base method
func (m *MockInterfacePeers) KnownPeers() ([]*types4.KnownPeer, error) {
	ret := m.ctrl.Call(m, "KnownPeers")
	ret0, _ := ret[0].([]*types4.KnownPeer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// KnownPeers indicates an expected call of KnownPeers
func (mr *MockInterfacePeersMockRecorder) KnownPeers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KnownPeers", reflect.TypeOf((*MockInterfacePeers)(nil).KnownPeers))
}

// MockInterfaceMessage is a mock of InterfaceMessage interface
type MockInterfaceMessage struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMessageMockRecorder
}

// MockInterfaceMessageMockRecorder is the mock recorder for MockInterfaceMessage
type MockInterfaceMessageMockRecorder struct {
	mock *MockInterfaceMessage
}

// NewMockInterfaceMessage creates a new mock instance
func NewMockInterfaceMessage(ctrl *gomock.Controller) *MockInterfaceMessage {
	mock := &MockInterfaceMessage{ctrl: ctrl}
	mock.recorder = &MockInterfaceMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterfaceMessage) EXPECT() *MockInterfaceMessageMockRecorder {
	return m.recorder
}

// Kind mocks base method
func (m *MockInterfaceMessage) Kind() types.FuncEnum {
	ret := m.ctrl.Call(m, "Kind")
	ret0, _ := ret[0].(types.FuncEnum)
	return ret0
}

// Kind indicates an expected call of Kind
func (mr *MockInterfaceMessageMockRecorder) Kind() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Kind", reflect.TypeOf((*MockInterfaceMessage)(nil).Kind))
}

// Encode mocks base method
func (m *MockInterfaceMessage) Encode() ([]byte, error) {
	ret := m.ctrl.Call(m, "Encode")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encode indicates an expected call of Encode
func (mr *MockInterfaceMessageMockRecorder) Encode() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockInterfaceMessage)(nil).Encode))
}

// Decode mocks base method
func (m *MockInterfaceMessage) Decode(bytes []byte) error {
	ret := m.ctrl.Call(m, "Decode", bytes)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decode indicates an expected call of Decode
func (mr *MockInterfaceMessageMockRecorder) Decode(bytes interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockInterfaceMessage)(nil).Decode), bytes)
}

// MarshalJSON mocks base method
func (m *MockInterfaceMessage) MarshalJSON() ([]byte, error) {
	ret := m.ctrl.Call(m, "MarshalJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalJSON indicates an expected call of MarshalJSON
func (mr *MockInterfaceMessageMockRecorder) MarshalJSON() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalJSON", reflect.TypeOf((*MockInterfaceMessage)(nil).MarshalJSON))
}

// UnmarshalJSON mocks base method
func (m *MockInterfaceMessage) UnmarshalJSON(bytes []byte) error {
	ret := m.ctrl.Call(m, "UnmarshalJSON", bytes)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnmarshalJSON indicates an expected call of UnmarshalJSON
func (mr *MockInterfaceMessageMockRecorder) UnmarshalJSON(bytes interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalJSON", reflect.TypeOf((*MockInterfaceMessage)(nil).UnmarshalJSON), bytes)
}

// Header mocks base method
func (m *MockInterfaceMessage) Header() *types4.Header {
	ret := m.ctrl.Call(m, "Header")
	ret0, _ := ret[0].(*types4.Header)
	return ret0
}

// Header indicates an expected call of Header
func (mr *MockInterfaceMessageMockRecorder) Header() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Header", reflect.TypeOf((*MockInterfaceMessage)(nil).Header))
}

// MockInterfaceTelemetry is a mock of InterfaceTelemetry interface
type MockInterfaceTelemetry struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceTelemetryMockRecorder
}

// MockInterfaceTelemetryMockRecorder is the mock recorder for MockInterfaceTelemetry
type MockInterfaceTelemetryMockRecorder struct {
	mock *MockInterfaceTelemetry
}

// NewMockInterfaceTelemetry creates a new mock instance
func NewMockInterfaceTelemetry(ctrl *gomock.Controller) *MockInterfaceTelemetry {
	mock := &MockInterfaceTelemetry{ctrl: ctrl}
	mock.recorder = &MockInterfaceTelemetryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterfaceTelemetry) EXPECT() *MockInterfaceTelemetryMockRecorder {
	return m.recorder
}

// MockInterfaceRPC is a mock of InterfaceRPC interface
type MockInterfaceRPC struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceRPCMockRecorder
}

// MockInterfaceRPCMockRecorder is the mock recorder for MockInterfaceRPC
type MockInterfaceRPCMockRecorder struct {
	mock *MockInterfaceRPC
}

// NewMockInterfaceRPC creates a new mock instance
func NewMockInterfaceRPC(ctrl *gomock.Controller) *MockInterfaceRPC {
	mock := &MockInterfaceRPC{ctrl: ctrl}
	mock.recorder = &MockInterfaceRPCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterfaceRPC) EXPECT() *MockInterfaceRPCMockRecorder {
	return m.recorder
}

// MockInterfaceP2P is a mock of InterfaceP2P interface
type MockInterfaceP2P struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceP2PMockRecorder
}

// MockInterfaceP2PMockRecorder is the mock recorder for MockInterfaceP2P
type MockInterfaceP2PMockRecorder struct {
	mock *MockInterfaceP2P
}

// NewMockInterfaceP2P creates a new mock instance
func NewMockInterfaceP2P(ctrl *gomock.Controller) *MockInterfaceP2P {
	mock := &MockInterfaceP2P{ctrl: ctrl}
	mock.recorder = &MockInterfaceP2PMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInterfaceP2P) EXPECT() *MockInterfaceP2PMockRecorder {
	return m.recorder
}

// IsStarted mocks base method
func (m *MockInterfaceP2P) IsStarted() bool {
	ret := m.ctrl.Call(m, "IsStarted")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsStarted indicates an expected call of IsStarted
func (mr *MockInterfaceP2PMockRecorder) IsStarted() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsStarted", reflect.TypeOf((*MockInterfaceP2P)(nil).IsStarted))
}

// GetNumPeers mocks base method
func (m *MockInterfaceP2P) GetNumPeers() (uint, error) {
	ret := m.ctrl.Call(m, "GetNumPeers")
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNumPeers indicates an expected call of GetNumPeers
func (mr *MockInterfaceP2PMockRecorder) GetNumPeers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumPeers", reflect.TypeOf((*MockInterfaceP2P)(nil).GetNumPeers))
}

// On mocks base method
func (m *MockInterfaceP2P) On(event types3.EventEnum, cb types4.EventCallback) {
	m.ctrl.Call(m, "On", event, cb)
}

// On indicates an expected call of On
func (mr *MockInterfaceP2PMockRecorder) On(event, cb interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "On", reflect.TypeOf((*MockInterfaceP2P)(nil).On), event, cb)
}

// Start mocks base method
func (m *MockInterfaceP2P) Start() error {
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (mr *MockInterfaceP2PMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockInterfaceP2P)(nil).Start))
}

// Stop mocks base method
func (m *MockInterfaceP2P) Stop() error {
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockInterfaceP2PMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockInterfaceP2P)(nil).Stop))
}

// Cfg mocks base method
func (m *MockInterfaceP2P) Cfg() types4.ConfigClient {
	ret := m.ctrl.Call(m, "Cfg")
	ret0, _ := ret[0].(types4.ConfigClient)
	return ret0
}

// Cfg indicates an expected call of Cfg
func (mr *MockInterfaceP2PMockRecorder) Cfg() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cfg", reflect.TypeOf((*MockInterfaceP2P)(nil).Cfg))
}

// GetSyncer mocks base method
func (m *MockInterfaceP2P) GetSyncer() (types4.InterfaceSync, error) {
	ret := m.ctrl.Call(m, "GetSyncer")
	ret0, _ := ret[0].(types4.InterfaceSync)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSyncer indicates an expected call of GetSyncer
func (mr *MockInterfaceP2PMockRecorder) GetSyncer() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSyncer", reflect.TypeOf((*MockInterfaceP2P)(nil).GetSyncer))
}
