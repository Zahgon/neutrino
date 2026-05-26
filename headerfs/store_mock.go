package headerfs

import (
	"io"

	"github.com/btcsuite/btcd/blockchain"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcwallet/walletdb"
	"github.com/stretchr/testify/mock"
)

// MockBlockHeaderStore is a mock implementation of the BlockHeaderStore.
type MockBlockHeaderStore struct {
	mock.Mock
}

// ChainTip returns the current chain tip for the mock block header store.
func (m *MockBlockHeaderStore) ChainTip() (*wire.BlockHeader, uint32, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// LatestBlockLocator returns the latest block locator for the mock block header
// store.
//
//nolint:lll
func (m *MockBlockHeaderStore) LatestBlockLocator() (blockchain.BlockLocator, error) {
	_ = "STUB: not implemented"
	return *new(blockchain.BlockLocator), nil
}

// FetchHeaderByHeight fetches a block header by height for the mock block
// header store.
func (m *MockBlockHeaderStore) FetchHeaderByHeight(
	height uint32) (*wire.BlockHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FetchHeaderAncestors fetches block header ancestors for the mock block header
// store.
func (m *MockBlockHeaderStore) FetchHeaderAncestors(numHeaders uint32,
	stopHash *chainhash.Hash) ([]wire.BlockHeader, uint32, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// HeightFromHash returns the height from a hash for the mock block header
// store.
func (m *MockBlockHeaderStore) HeightFromHash(
	hash *chainhash.Hash) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// FetchHeader fetches a block header by hash for the mock block header store.
func (m *MockBlockHeaderStore) FetchHeader(
	hash *chainhash.Hash) (*wire.BlockHeader, uint32, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// WriteHeaders writes block headers to the mock block header store.
func (m *MockBlockHeaderStore) WriteHeaders(hdrs ...BlockHeader) error {
	_ = "STUB: not implemented"
	return nil
}

// RollbackBlockHeaders rolls back block headers in the mock block header store.
func (m *MockBlockHeaderStore) RollbackBlockHeaders(
	numHeaders uint32) (*BlockStamp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RollbackLastBlock rolls back the last block in the mock block header store.
func (m *MockBlockHeaderStore) RollbackLastBlock() (*BlockStamp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MockFilterHeaderStore is a mock implementation of the FilterHeaderStore.
type MockFilterHeaderStore struct {
	mock.Mock
}

// ChainTip returns the current chain tip for the mock filter header store.
func (m *MockFilterHeaderStore) ChainTip() (*chainhash.Hash, uint32, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// FetchHeader fetches a filter header by hash for the mock filter header store.
func (m *MockFilterHeaderStore) FetchHeader(
	hash *chainhash.Hash) (*chainhash.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FetchHeaderAncestors fetches filter header ancestors for the mock filter
// header store.
func (m *MockFilterHeaderStore) FetchHeaderAncestors(numHeaders uint32,
	stopHash *chainhash.Hash) ([]chainhash.Hash, uint32, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// FetchHeaderByHeight fetches a filter header by height for the mock filter
// header store.
func (m *MockFilterHeaderStore) FetchHeaderByHeight(
	height uint32) (*chainhash.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WriteHeaders writes filter headers to the mock filter header store.
func (m *MockFilterHeaderStore) WriteHeaders(headers ...FilterHeader) error {
	_ = "STUB: not implemented"
	return nil
}

// RollbackLastBlock rolls back the last block in the mock filter header store.
func (m *MockFilterHeaderStore) RollbackLastBlock(
	newTip *chainhash.Hash) (*BlockStamp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MockWalletDB is a mock implementation of walletdb.DB for testing.
type MockWalletDB struct {
	mock.Mock
}

// Update implements the walletdb.DB interface.
func (m *MockWalletDB) Update(fn func(tx walletdb.ReadWriteTx) error) error {
	_ = "STUB: not implemented"
	return nil
}

// View implements the walletdb.DB interface.
func (m *MockWalletDB) View(fn func(tx walletdb.ReadTx) error) error {
	_ = "STUB: not implemented"
	return nil
}

// BeginReadWriteTx implements the walletdb.DB interface.
func (m *MockWalletDB) BeginReadWriteTx() (walletdb.ReadWriteTx, error) {
	_ = "STUB: not implemented"
	return *new(walletdb.ReadWriteTx), nil
}

// BeginReadTx implements the walletdb.DB interface.
func (m *MockWalletDB) BeginReadTx() (walletdb.ReadTx, error) {
	_ = "STUB: not implemented"
	return *new(walletdb.ReadTx), nil
}

// Copy implements the walletdb.DB interface.
func (m *MockWalletDB) Copy(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Close implements the walletdb.DB interface.
func (m *MockWalletDB) Close() error { _ = "STUB: not implemented"; return nil }
