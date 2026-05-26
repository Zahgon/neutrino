package headerfs

import (
	"bytes"
	"io"
	"os"
	"sync"
	"time"

	"github.com/btcsuite/btcd/blockchain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcwallet/walletdb"
)

// BlockStamp represents a block, identified by its height and time stamp in
// the chain. We also lift the timestamp from the block header itself into this
// struct as well.
type BlockStamp struct {
	// Height is the height of the target block.
	Height int32

	// Hash is the hash that uniquely identifies this block.
	Hash chainhash.Hash

	// Timestamp is the timestamp of the block in the chain.
	Timestamp time.Time
}

// BlockHeaderStore is an interface that provides an abstraction for a generic
// store for block headers.
type BlockHeaderStore interface {
	// ChainTip returns the best known block header and height for the
	// BlockHeaderStore.
	ChainTip() (*wire.BlockHeader, uint32, error)

	// LatestBlockLocator returns the latest block locator object based on
	// the tip of the current main chain from the PoV of the
	// BlockHeaderStore.
	LatestBlockLocator() (blockchain.BlockLocator, error)

	// FetchHeaderByHeight attempts to retrieve a target block header based
	// on a block height.
	FetchHeaderByHeight(height uint32) (*wire.BlockHeader, error)

	// FetchHeaderAncestors fetches the numHeaders block headers that are
	// the ancestors of the target stop hash. A total of numHeaders+1
	// headers will be returned, as we'll walk back numHeaders distance to
	// collect each header, then return the final header specified by the
	// stop hash. We'll also return the starting height of the header range
	// as well so callers can compute the height of each header without
	// knowing the height of the stop hash.
	FetchHeaderAncestors(uint32, *chainhash.Hash) ([]wire.BlockHeader,
		uint32, error)

	// HeightFromHash returns the height of a particular block header given
	// its hash.
	HeightFromHash(*chainhash.Hash) (uint32, error)

	// FetchHeader attempts to retrieve a block header determined by the
	// passed block height.
	FetchHeader(*chainhash.Hash) (*wire.BlockHeader, uint32, error)

	// WriteHeaders adds a set of headers to the BlockHeaderStore in a
	// single atomic transaction.
	WriteHeaders(...BlockHeader) error

	// RollbackBlockHeaders rolls back a specified number of headers from
	// the tip of the chain. It removes the most recent 'numHeaders' from
	// the block headers file and updates the corresponding indices. This
	// method is used during blockchain reorganizations to remove headers
	// that are no longer part of the main chain. The function will return
	// an error if the rollback would reach or go before the genesis block
	// (height 0). The information about the new header tip after truncation
	// is returned.
	RollbackBlockHeaders(numHeaders uint32) (*BlockStamp, error)

	// RollbackLastBlock rolls back the BlockHeaderStore by a _single_
	// header. This method is meant to be used in the case of re-org which
	// disconnects the latest block header from the end of the main chain.
	// The information about the new header tip after truncation is
	// returned.
	//
	// NOTE: This function is maintained for backward compatibility since it
	// is a publicly exposed function. It now internally utilizes
	// RollbackBlockHeaders API.
	RollbackLastBlock() (*BlockStamp, error)
}

// headerBufPool is a pool of bytes.Buffer that will be re-used by the various
// headerStore implementations to batch their header writes to disk. By
// utilizing this variable we can minimize the total number of allocations when
// writing headers to disk.
var headerBufPool = sync.Pool{
	New: func() interface{} { return new(bytes.Buffer) },
}

// File defines the minimum file operations needed by headerStore.
type File interface {
	// Basic I/O operations.
	io.Reader
	io.Writer
	io.Closer

	// Extended I/O positioning.
	io.Seeker
	io.ReaderAt

	// File-specific operations.
	Stat() (os.FileInfo, error)
	Sync() error
	Truncate(size int64) error

	// Returns the name of the file.
	Name() string
}

type headerFile struct {
	file File
}

// truncateHeaders truncates one or more headers from the end of the header
// file. This can be used in the case of a re-org to remove headers from the end
// of the main chain.
//
// The numHeaders parameter specifies how many headers to truncate. If
// numHeaders is 1, this is equivalent to the old singleTruncate behavior.
//
// This function handles platform-specific differences in file truncation. On
// Windows, the file is closed, truncated, and reopened due to Windows
// limitations on truncating open files. On other platforms, the file is
// truncated directly without closing.
func (h *headerFile) truncateHeaders(numHeaders uint32,
	headerType HeaderType) error {
	_ = "STUB: not implemented"

	// If numHeaders is 0, treat it as a no-op and return no error.
	return nil
}

// In order to truncate the file, we'll need to grab the absolute size
// of the file as it stands currently.

// Calculate the total bytes to truncate based on number of headers.

// Finally, we'll use both of these values to calculate the new size of
// the file and truncate it accordingly.

// On Windows, we need to close, truncate, and reopen the file.

// headerStore combines a on-disk set of headers within a flat file in addition
// to a database which indexes that flat file. Together, these two abstractions
// can be used in order to build an indexed header store for any type of
// "header" as it deals only with raw bytes, and leaves it to a higher layer to
// interpret those raw bytes accordingly.
//
// TODO(roasbeef): quickcheck coverage.
type headerStore struct {
	mtx sync.RWMutex // nolint:structcheck // false positive because used as embedded struct only

	*headerFile

	*headerIndex
}

// newHeaderStore creates a new headerStore given an already open database, a
// target file path for the flat-file and a particular header type. The target
// file will be created as necessary.
func newHeaderStore(db walletdb.DB, filePath string,
	hType HeaderType) (*headerStore, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We'll open the file, creating it if necessary and ensuring that all
// writes are actually appends to the end of the file.

// With the file open, we'll then create the header index so we can
// have random access into the flat files.

// blockHeaderStore is an implementation of the BlockHeaderStore interface, a
// fully fledged database for Bitcoin block headers. The blockHeaderStore
// combines a flat file to store the block headers with a database instance for
// managing the index into the set of flat files.
type blockHeaderStore struct {
	*headerStore
}

// A compile-time check to ensure the blockHeaderStore adheres to the
// BlockHeaderStore interface.
var _ BlockHeaderStore = (*blockHeaderStore)(nil)

// NewBlockHeaderStore creates a new instance of the blockHeaderStore based on
// a target file path, an open database instance, and finally a set of
// parameters for the target chain. These parameters are required as if this is
// the initial start up of the blockHeaderStore, then the initial genesis
// header will need to be inserted.
func NewBlockHeaderStore(filePath string, db walletdb.DB,
	netParams *chaincfg.Params) (BlockHeaderStore, error) {
	_ = "STUB: not implemented"
	return *new(BlockHeaderStore), nil
}

// With the header store created, we'll fetch the file size to see if
// we need to initialize it with the first header or not.

// If the size of the file is zero, then this means that we haven't yet
// written the initial genesis header to disk, so we'll do so now.

// As a final initialization step (if this isn't the first time), we'll
// ensure that the header tip within the flat files, is in sync with
// out database index.

// First, we'll compute the size of the current file so we can
// calculate the latest header written to disk.

// Using the file's current height, fetch the latest on-disk header.

// If the index's tip hash, and the file on-disk match, then we're
// done here.

// TODO(roasbeef): below assumes index can never get ahead?
//  * we always update files _then_ indexes
//  * need to dual pointer walk back for max safety

// Otherwise, we'll need to truncate the file until it matches the
// current index tip.

// FetchHeader attempts to retrieve a block header determined by the passed
// block height.
//
// NOTE: Part of the BlockHeaderStore interface.
func (h *blockHeaderStore) FetchHeader(hash *chainhash.Hash) (*wire.BlockHeader, uint32, error) {
	_ = "STUB: not implemented"
	// Lock store for read.
	return nil, 0, nil
}

// First, we'll query the index to obtain the block height of the
// passed block hash.

// With the height known, we can now read the header from disk.

// FetchHeaderByHeight attempts to retrieve a target block header based on a
// block height.
//
// NOTE: Part of the BlockHeaderStore interface.
func (h *blockHeaderStore) FetchHeaderByHeight(height uint32) (*wire.BlockHeader, error) {
	_ = "STUB: not implemented"
	// Lock store for read.
	return nil, nil
}

// For this query, we don't need to consult the index, and can instead
// just seek into the flat file based on the target height and return
// the full header.

// FetchHeaderAncestors fetches the numHeaders block headers that are the
// ancestors of the target stop hash. A total of numHeaders+1 headers will be
// returned, as we'll walk back numHeaders distance to collect each header,
// then return the final header specified by the stop hash. We'll also return
// the starting height of the header range as well so callers can compute the
// height of each header without knowing the height of the stop hash.
//
// NOTE: Part of the BlockHeaderStore interface.
func (h *blockHeaderStore) FetchHeaderAncestors(numHeaders uint32,
	stopHash *chainhash.Hash) ([]wire.BlockHeader, uint32, error) {
	_ = "STUB: not implemented"

	// First, we'll find the final header in the range, this will be the
	// ending height of our scan.
	return nil, 0, nil
}

// HeightFromHash returns the height of a particular block header given its
// hash.
//
// NOTE: Part of the BlockHeaderStore interface.
func (h *blockHeaderStore) HeightFromHash(hash *chainhash.Hash) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// RollbackBlockHeaders removes the specified number of block headers from the
// end of the chain. It returns a BlockStamp representing the new chain tip. If
// numHeaders is 0, it returns an empty BlockStamp without performing any
// operations.
//
// The function ensures rollback doesn't remove or go beyond the genesis block
// (height 0). It determines the current chain tip height, reads the header
// range to be removed along with the new tip header, truncates the headers file
// to remove the specified number of headers, and updates the header indices to
// reflect the new chain tip.
func (h *blockHeaderStore) RollbackBlockHeaders(n uint32) (*BlockStamp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Lock store for rollback.

// First, we'll obtain the latest height that the index knows of.

// Ensure the rollback doesn't remove or go beyond the genesis block.

// With this height obtained, we'll use it to read the previous header
// from disk, so we can populate our return value which requires the
// prev header hash.

// Transform to blockhashes for downstream operations, starting at
// headers + 1 skipping the previous header.

// RollbackLastBlock rollsback both the index, and on-disk header file by a
// _single_ header. This method is meant to be used in the case of re-org which
// disconnects the latest block header from the end of the main chain. The
// information about the new header tip after truncation is returned.
//
// NOTE: Part of the BlockHeaderStore interface.
func (h *blockHeaderStore) RollbackLastBlock() (*BlockStamp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// BlockHeader is a Bitcoin block header that also has its height included.
type BlockHeader struct {
	*wire.BlockHeader

	// Height is the height of this block header within the current main
	// chain.
	Height uint32
}

// toIndexEntry converts the BlockHeader into a matching headerEntry. This
// method is used when a header is to be written to disk.
func (b *BlockHeader) toIndexEntry() headerEntry {
	_ = "STUB: not implemented"
	return *new(headerEntry)
}

// WriteHeaders writes a set of headers to disk and updates the index in a
// single atomic transaction.
//
// NOTE: Part of the BlockHeaderStore interface.
func (h *blockHeaderStore) WriteHeaders(hdrs ...BlockHeader) error {
	_ = "STUB: not implemented"
	// Lock store for write.
	return nil
}

// First, we'll grab a buffer from the write buffer pool so we an
// reduce our total number of allocations, and also write the headers
// in a single swoop.

// Next, we'll write out all the passed headers in series into the
// buffer we just extracted from the pool.

// With all the headers written to the buffer, we'll now write out the
// entire batch in a single write call.

// Once those are written, we'll then collate all the headers into
// headerEntry instances so we can write them all into the index in a
// single atomic batch.

// Attempt to add the headers to the database. If this fails, we'll need
// to roll back any changes to the header file to maintain consistency.
// The rollback process bases on the number of header serialized. If
// both the initial operation and the rollback fail, we return
// a detailed error explaining both failures to aid in debugging.

// Since the probability of failing to write to the database is
// very low, it is mostly worth the cost of file sync operation
// to make sure truncate headers does it correctly.

// blockLocatorFromHash takes a given block hash and then creates a block
// locator using it as the root of the locator. We'll start by taking a single
// step backwards, then keep doubling the distance until genesis after we get
// 10 locators.
//
// TODO(roasbeef): make into single transaction.
func (h *blockHeaderStore) blockLocatorFromHash(hash *chainhash.Hash) (
	blockchain.BlockLocator, error) {
	_ = "STUB: not implemented"
	return *new(blockchain.BlockLocator), nil
}

// Append the initial hash

// If hash isn't found in DB or this is the genesis block, return the
// locator as is.

// Decrement by 1 for the first 10 blocks, then double the jump
// until we get to the genesis hash

// LatestBlockLocator returns the latest block locator object based on the tip
// of the current main chain from the PoV of the database and flat files.
//
// NOTE: Part of the BlockHeaderStore interface.
func (h *blockHeaderStore) LatestBlockLocator() (blockchain.BlockLocator, error) {
	_ = "STUB: not implemented"
	// Lock store for read.
	return *new(blockchain.BlockLocator), nil
}

// BlockLocatorFromHash computes a block locator given a particular hash. The
// standard Bitcoin algorithm to compute block locators are employed.
func (h *blockHeaderStore) BlockLocatorFromHash(hash *chainhash.Hash) (
	blockchain.BlockLocator, error) {
	_ = "STUB: not implemented"

	// Lock store for read.
	return *new(blockchain.BlockLocator), nil
}

// CheckConnectivity cycles through all of the block headers on disk, from last
// to first, and makes sure they all connect to each other. Additionally, at
// each block header, we also ensure that the index entry for that height and
// hash also match up properly.
func (h *blockHeaderStore) CheckConnectivity() error {
	_ = "STUB: not implemented"
	// Lock store for read.
	return nil
}

// First, we'll fetch the chain tip so we can start our
// backwards scan.

// With the height extracted, we'll now read the _last_ block
// header within the file before we kick off our connectivity
// loop.

// We'll now cycle backwards, seeking backwards along the
// header file to ensure each header connects properly and the
// index entries are also accurate. To do this, we start from a
// height of one before our current tip.

// First, read the block header for this block height,
// and also compute the block hash for it.

// With the header retrieved, we'll now fetch the
// height for this current header hash to ensure the
// on-disk state and the index matches up properly.

// With the index entry retrieved, we'll now assert
// that the height matches up with our current height
// in this backwards walk.

// Finally, we'll assert that this new header is
// actually the prev header of the target header from
// the last loop. This ensures connectivity.

// As all the checks have passed, we'll now reset our
// header pointer to this current location, and
// continue our backwards walk.

// ChainTip returns the best known block header and height for the
// blockHeaderStore.
//
// NOTE: Part of the BlockHeaderStore interface.
func (h *blockHeaderStore) ChainTip() (*wire.BlockHeader, uint32, error) {
	_ = "STUB: not implemented"
	// Lock store for read.
	return nil, 0, nil
}

// FilterHeaderStore defines the interface for storing and retrieving filter
// headers.
type FilterHeaderStore interface {
	// ChainTip returns the hash and height of the latest filter header.
	ChainTip() (*chainhash.Hash, uint32, error)

	// FetchHeader fetches the filter header for a specific block hash.
	FetchHeader(hash *chainhash.Hash) (*chainhash.Hash, error)

	// FetchHeaderAncestors fetches the given number of headers starting
	// from the specified stop hash and working backwards.
	FetchHeaderAncestors(numHeaders uint32,
		stopHash *chainhash.Hash) ([]chainhash.Hash, uint32, error)

	// FetchHeaderByHeight fetches the filter header for a specific block
	// height.
	FetchHeaderByHeight(height uint32) (*chainhash.Hash, error)

	// WriteHeaders writes a set of filter headers to the store.
	WriteHeaders(hdrs ...FilterHeader) error

	// RollbackLastBlock rolls back the last block, returning the new tip
	// after rollback.
	RollbackLastBlock(newTip *chainhash.Hash) (*BlockStamp, error)
}

// filterHeaderStore is an implementation of a fully fledged database for any
// variant of filter headers. The filterHeaderStore combines a flat file to
// store the block headers with a database instance for managing the index into
// the set of flat files.
type filterHeaderStore struct {
	*headerStore
}

// Compile-time assertion to ensure filterHeaderStore implements
// FilterHeaderStore interface.
var _ FilterHeaderStore = (*filterHeaderStore)(nil)

// NewFilterHeaderStore returns a new instance of the FilterHeaderStore based
// on a target file path, filter type, and target net parameters. These
// parameters are required as if this is the initial start up of the
// FilterHeaderStore, then the initial genesis filter header will need to be
// inserted.
func NewFilterHeaderStore(filePath string, db walletdb.DB,
	filterType HeaderType, netParams *chaincfg.Params,
	headerStateAssertion *FilterHeader) (FilterHeaderStore, error) {
	_ = "STUB: not implemented"
	return *new(FilterHeaderStore), nil
}

// With the header store created, we'll fetch the fiie size to see if
// we need to initialize it with the first header or not.

// TODO(roasbeef): also reconsile with block header state due to way
// roll back works atm

// If the size of the file is zero, then this means that we haven't yet
// written the initial genesis header to disk, so we'll do so now.

// If we have a state assertion then we'll check it now to see if we
// need to modify our filter header files before we proceed.

// If the filter header store was reset, we'll re-initialize it
// to recreate our on-disk state.

// As a final initialization step, we'll ensure that the header tip
// within the flat files, is in sync with out database index.

// First, we'll compute the size of the current file so we can
// calculate the latest header written to disk.

// Using the file's current height, fetch the latest on-disk header.

// If the index's tip hash, and the file on-disk match, then we're
// doing here.

// Otherwise, we'll need to truncate the file until it matches the
// current index tip.

// maybeResetHeaderState will reset the header state if the header assertion
// fails, but only if the target height is found. The boolean returned indicates
// that header state was reset.
func (f *filterHeaderStore) maybeResetHeaderState(
	headerStateAssertion *FilterHeader) (bool, error) {
	_ = "STUB: not implemented"

	// First, we'll attempt to locate the header at this height. If no such
	// header is found, then we'll exit early.
	return false, nil
}

// If our on disk state and the provided header assertion don't match,
// then we'll purge this state so we can sync it anew once we fully
// start up.

// Close the file before removing it. This is required by some
// OS, e.g., Windows.

// FetchHeader returns the filter header that corresponds to the passed block
// height.
//
// NOTE: Part of the FilterHeaderStore interface.
func (f *filterHeaderStore) FetchHeader(
	hash *chainhash.Hash) (*chainhash.Hash, error) {
	_ = "STUB: not implemented"

	// Lock store for read.
	return nil, nil
}

// FetchHeaderByHeight returns the filter header for a particular block height.
//
// NOTE: Part of the FilterHeaderStore interface.
func (f *filterHeaderStore) FetchHeaderByHeight(
	height uint32) (*chainhash.Hash, error) {
	_ = "STUB: not implemented"

	// Lock store for read.
	return nil, nil
}

// FetchHeaderAncestors fetches the numHeaders filter headers that are the
// ancestors of the target stop block hash. A total of numHeaders+1 headers will be
// returned, as we'll walk back numHeaders distance to collect each header,
// then return the final header specified by the stop hash. We'll also return
// the starting height of the header range as well so callers can compute the
// height of each header without knowing the height of the stop hash.
//
// NOTE: Part of the FilterHeaderStore interface.
func (f *filterHeaderStore) FetchHeaderAncestors(numHeaders uint32,
	stopHash *chainhash.Hash) ([]chainhash.Hash, uint32, error) {
	_ = "STUB: not implemented"

	// First, we'll find the final header in the range, this will be the
	// ending height of our scan.
	return nil, 0, nil
}

// FilterHeader represents a filter header (basic or extended). The filter
// header itself is coupled with the block height and hash of the filter's
// block.
type FilterHeader struct {
	// HeaderHash is the hash of the block header that this filter header
	// corresponds to.
	HeaderHash chainhash.Hash

	// FilterHash is the filter header itself.
	FilterHash chainhash.Hash

	// Height is the block height of the filter header in the main chain.
	Height uint32
}

// toIndexEntry converts the filter header into a index entry to be stored
// within the database.
func (f *FilterHeader) toIndexEntry() headerEntry {
	_ = "STUB: not implemented"
	return *new(headerEntry)
}

// WriteHeaders writes a batch of filter headers to persistent storage. The
// headers themselves are appended to the flat file, and then the index updated
// to reflect the new entries.
//
// NOTE: Part of the FilterHeaderStore interface.
func (f *filterHeaderStore) WriteHeaders(hdrs ...FilterHeader) error {
	_ = "STUB: not implemented"
	// Lock store for write.
	return nil
}

// If there are 0 headers to be written, return immediately. This
// prevents the newTip assignment from panicking because of an index
// of -1.

// First, we'll grab a buffer from the write buffer pool so we an
// reduce our total number of allocations, and also write the headers
// in a single swoop.

// Next, we'll write out all the passed headers in series into the
// buffer we just extracted from the pool.

// With all the headers written to the buffer, we'll now write out the
// entire batch in a single write call.

// As the block headers should already be written, we only need to
// update the tip pointer for this particular header type.

// Attempt to add the headers to the database. If this fails, we'll need
// to roll back any changes to the header file to maintain consistency.
// The rollback process bases on the number of header serialized. If
// both the initial operation and the rollback fail, we return
// a detailed error explaining both failures to aid in debugging.

// Since the probability of failing to write to the database is
// very low, it is mostly worth the cost of file sync operation
// to make sure truncate headers does it correctly.

// ChainTip returns the latest filter header and height known to the
// FilterHeaderStore.
//
// NOTE: Part of the FilterHeaderStore interface.
func (f *filterHeaderStore) ChainTip() (*chainhash.Hash, uint32, error) {
	_ = "STUB: not implemented"
	// Lock store for read.
	return nil, 0, nil
}

// RollbackLastBlock rollsback both the index, and on-disk header file by a
// _single_ filter header. This method is meant to be used in the case of
// re-org which disconnects the latest filter header from the end of the main
// chain. The information about the latest header tip after truncation is
// returned.
//
// NOTE: Part of the FilterHeaderStore interface.
func (f *filterHeaderStore) RollbackLastBlock(
	newTip *chainhash.Hash) (*BlockStamp, error) {
	_ = "STUB: not implemented"

	// Lock store for write.
	return nil, nil
}

// First, we'll obtain the latest height that the index knows of.

// With this height obtained, we'll use it to read what will be the new
// chain tip from disk.

// Now that we have the information we need to return from this
// function, we can now truncate both the header file and the index.

// TODO(roasbeef): return chain hash also?
