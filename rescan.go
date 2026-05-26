// NOTE: THIS API IS UNSTABLE RIGHT NOW.

package neutrino

import (
	"errors"
	"sync"
	"time"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/gcs"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/wire"
	"github.com/lightninglabs/neutrino/blockntfns"
	"github.com/lightninglabs/neutrino/headerfs"
)

var (
	// zeroOutPoint indicates that we should match on an output's script
	// when dispatching a spend notification.
	zeroOutPoint wire.OutPoint

	// ErrRescanExit is an error returned to the caller in case the ongoing
	// rescan exits.
	ErrRescanExit = errors.New("rescan exited")

	// errRetryBlock is an internal error used to signal to the rescan
	// should it should attempt to retry processing a block.
	errRetryBlock = errors.New("block must be retried")
)

// ChainSource is an interface that's in charge of retrieving information about
// the existing chain.
type ChainSource interface {
	// ChainParams returns the parameters of the current chain.
	ChainParams() chaincfg.Params

	// BestBlock retrieves the most recent block's height and hash where we
	// have both the header and filter header ready.
	BestBlock() (*headerfs.BlockStamp, error)

	// GetBlockHeaderByHeight returns the header of the block with the given
	// height.
	GetBlockHeaderByHeight(uint32) (*wire.BlockHeader, error)

	// GetBlockHeader returns the header of the block with the given hash.
	GetBlockHeader(*chainhash.Hash) (*wire.BlockHeader, uint32, error)

	// GetBlock returns the block with the given hash.
	GetBlock(chainhash.Hash, ...QueryOption) (*btcutil.Block, error)

	// GetFilterHeaderByHeight returns the filter header of the block with
	// the given height.
	GetFilterHeaderByHeight(uint32) (*chainhash.Hash, error)

	// GetCFilter returns the filter of the given type for the block with
	// the given hash.
	GetCFilter(chainhash.Hash,
		wire.FilterType, ...QueryOption) (*gcs.Filter, error)

	// Subscribe returns a block subscription that delivers block
	// notifications in order. The bestHeight parameter can be used to
	// signal that a backlog of notifications should be delivered from this
	// height. When providing a bestHeight of 0, a backlog will not be
	// delivered.
	//
	// TODO(wilmer): extend with best hash as well.
	Subscribe(bestHeight uint32) (*blockntfns.Subscription, error)

	// IsCurrent returns true if the backend chain thinks that its view of
	// the network is current.
	IsCurrent() bool
}

// ScanProgressHandler is used in rescanOptions to update the caller with the
// rescan progress.
type ScanProgressHandler func(lastProcessedBlock uint32)

// rescanOptions holds the set of functional parameters for Rescan.
type rescanOptions struct {
	queryOptions []QueryOption

	ntfn            rpcclient.NotificationHandlers
	progressHandler ScanProgressHandler

	startTime  time.Time
	startBlock *headerfs.BlockStamp

	endBlock *headerfs.BlockStamp

	watchAddrs  []btcutil.Address
	watchInputs []InputWithScript
	watchList   [][]byte
	txIdx       uint32

	update <-chan *updateOptions
	quit   <-chan struct{}
}

// RescanOption is a functional option argument to any of the rescan and
// notification subscription methods. These are always processed in order, with
// later options overriding earlier ones.
type RescanOption func(ro *rescanOptions)

func defaultRescanOptions() *rescanOptions { _ = "STUB: not implemented"; return nil }

// QueryOptions pass onto the underlying queries.
func QueryOptions(options ...QueryOption) RescanOption {
	_ = "STUB: not implemented"
	return *new(RescanOption)
}

// NotificationHandlers specifies notification handlers for the rescan. These
// will always run in the same goroutine as the caller.
func NotificationHandlers(ntfn rpcclient.NotificationHandlers) RescanOption {
	_ = "STUB: not implemented"
	return *new(RescanOption)
}

// ProgressHandler specifies a handler to be used when the utxo
// scanner reports its progress.
// The passed handler should be non-blocking for the rescan to continue
// normally.
func ProgressHandler(
	handler ScanProgressHandler) RescanOption {
	_ = "STUB: not implemented"
	return *new(RescanOption)
}

// StartBlock specifies the start block. The hash is checked first; if there's
// no such hash (zero hash avoids lookup), the height is checked next. If the
// height is 0 or the start block isn't specified, starts from the genesis
// block. This block is assumed to already be known, and no notifications will
// be sent for this block. The rescan uses the latter of StartBlock and
// StartTime.
func StartBlock(startBlock *headerfs.BlockStamp) RescanOption {
	_ = "STUB: not implemented"
	return *new(RescanOption)
}

// StartTime specifies the start time. The time is compared to the timestamp of
// each block, and the rescan only begins once the first block crosses that
// timestamp. When using this, it is advisable to use a margin of error and
// start rescans slightly earlier than required. The rescan uses the latter of
// StartBlock and StartTime.
func StartTime(startTime time.Time) RescanOption {
	_ = "STUB: not implemented"
	return *new(RescanOption)
}

// EndBlock specifies the end block. The hash is checked first; if there's no
// such hash (zero hash avoids lookup), the height is checked next. If the
// height is 0 or in the future or the end block isn't specified, the quit
// channel MUST be specified as Rescan will sync to the tip of the blockchain
// and continue to stay in sync and pass notifications. This is enforced at
// runtime.
func EndBlock(endBlock *headerfs.BlockStamp) RescanOption {
	_ = "STUB: not implemented"
	return *new(RescanOption)
}

// WatchAddrs specifies the addresses to watch/filter for. Each call to this
// function adds to the list of addresses being watched rather than replacing
// the list. Each time a transaction spends to the specified address, the
// outpoint is added to the WatchOutPoints list.
func WatchAddrs(watchAddrs ...btcutil.Address) RescanOption {
	_ = "STUB: not implemented"
	return *new(RescanOption)
}

// InputWithScript couples an previous outpoint along with its input script.
// We'll use the prev script to match the filter itself, but then scan for the
// particular outpoint when we need to make a notification decision.
type InputWithScript struct {
	// OutPoint identifies the previous output to watch.
	OutPoint wire.OutPoint

	// PkScript is the script of the previous output.
	PkScript []byte
}

// WatchInputs specifies the outpoints to watch for on-chain spends. We also
// require the script as we'll match on the script, but then notify based on
// the outpoint. Each call to this function adds to the list of outpoints being
// watched rather than replacing the list.
func WatchInputs(watchInputs ...InputWithScript) RescanOption {
	_ = "STUB: not implemented"
	return *new(RescanOption)
}

// TxIdx specifies a hint transaction index into the block in which the UTXO is
// created (eg, coinbase is 0, next transaction is 1, etc.)
func TxIdx(txIdx uint32) RescanOption { _ = "STUB: not implemented"; return *new(RescanOption) }

// QuitChan specifies the quit channel. This can be used by the caller to let
// an indefinite rescan (one with no EndBlock set) know it should gracefully
// shut down. If this isn't specified, an end block MUST be specified as Rescan
// must know when to stop. This is enforced at runtime.
func QuitChan(quit <-chan struct{}) RescanOption {
	_ = "STUB: not implemented"
	return *new(RescanOption)
}

// blockRetryQueue is a helper struct that maintains a queue of blocks for which
// we need to fetch filters.
type blockRetryQueue struct {
	blocks []*blockntfns.Connected
}

// newBlockRetryQueue constructs a new, empty retry block queue.
func newBlockRetryQueue() *blockRetryQueue { _ = "STUB: not implemented"; return nil }

// push enqueues a block at the end of the queue.
func (q *blockRetryQueue) push(block *blockntfns.Connected) { _ = "STUB: not implemented"; return }

// peek returns the next block to be retried but doesn't consume it.
func (q *blockRetryQueue) peek() *blockntfns.Connected { _ = "STUB: not implemented"; return nil }

// pop returns and consumes the next block to be retried.
func (q *blockRetryQueue) pop() *blockntfns.Connected { _ = "STUB: not implemented"; return nil }

// remove removes the block from the queue and any others that follow it. If the
// block doesn't exit within the queue, then this acts as a NOP.
func (q *blockRetryQueue) remove(header wire.BlockHeader) { _ = "STUB: not implemented"; return }

// clear clears the queue.
func (q *blockRetryQueue) clear() {
	_ = "STUB: not implemented"

	// updateChan specifies an update channel. This is for internal use by the
	// Rescan.Update functionality.
	return
}

func updateChan(update <-chan *updateOptions) RescanOption {
	_ = "STUB: not implemented"
	return *new(RescanOption)
}

// rescanState hold the state used throughout a rescan.
type rescanState struct {
	// chain is the backend chain that the rescan has access to.
	chain ChainSource

	// opts holds the various rescan configuration options.
	opts *rescanOptions

	// curHeader is the block header of our current position in the chain.
	curHeader wire.BlockHeader

	// curStamp is the block stamp of our current position in the chain.
	curStamp headerfs.BlockStamp

	// scanning is true if the current block should be scanned for filter
	// matches.
	scanning bool
}

// newRescanState constructs a new rescanState.
func newRescanState(chain ChainSource, options ...RescanOption) (*rescanState,
	error) {
	_ = "STUB: not implemented"

	// First, we'll apply the set of default options, then serially apply
	// all the options that've been passed in.
	return nil, nil
}

// If we have something to watch, create a watch list. The watch list
// can be composed of a set of scripts, outpoints, and txids.

// Check that we have either an end block or a quit channel.

// If the end block hash is non-nil, then we'll query the
// database to find out the stop height.

// If the ending hash it nil, then check to see if the target
// height is non-nil. If not, then we'll use this to find the
// stopping hash.

// If we don't have a quit channel, and the end height is still
// unspecified, then we'll exit out here.

// If no start block is specified, start the scan from our current best
// block.

// To find our starting block, either the start hash should be set, or
// the start height should be set. If neither is, then we'll be
// starting from the genesis block.

// rescan is a single-threaded function that uses headers from the database and
// functional options as arguments.
func (rs *rescanState) rescan() error { _ = "STUB: not implemented"; return nil }

// We'll need to ensure that the backing chain has actually caught up to
// the rescan's starting height.

// To ensure that we batch as many filter queries as possible, we also
// wait for the header chain to either be current or for it to at least
// have caught up with the specified end block.

// If the header chain is current, then there is no need to
// wait.

// If an end height was specified then we wait until the
// notification corresponding to that block height.

// If a block hash was specified, check if the notification is
// for that block.

// Compare the start time to the start block. If the start time is
// later, cycle through blocks until we find a block timestamp later
// than the start time, and begin filter download at that block. Since
// time is non-monotonic between blocks, we look for the first block to
// trip the switch, and download filters from there, rather than
// checking timestamps at each block.

// Even though we'll have multiple subscriptions, they'll always be
// referred to by the same variable, so we only need to defer its
// cancellation once at the end. Any intermediate subscriptions should
// be properly canceled before registering a new one.

// blockRetryInterval is the interval in which we'll continually
// retry to fetch the latest filter from our peers.
//
// TODO(roasbeef): add exponential back-off

// We'll need to keep track of whether we are current with the chain in
// order to properly recover from a re-org. We'll start by assuming that
// we are not current in order to catch up from the starting point to
// the tip of the chain.

// Loop through blocks, one at a time. This relies on the underlying
// chain source to deliver notifications in the correct order.

// If we've reached the ending height or hash for this rescan,
// then we'll exit.

// If we're current, we wait for notifications that will be
// delivered each time a block is connecting, disconnecting, or
// we can an update to the filter we should be looking for.

// Wait for a signal that we have a newly connected
// header and cfheader, or a newly disconnected header;
// alternatively, forward ourselves to the next block
// if possible.

// An update message has just come across, if it points
// to a prior point in the chain, then we may need to
// rewind a bit in order to provide the client all its
// requested client.

// If we have to rewind our state, then we'll
// mark ourselves as not current so we can walk
// forward in the chain again until we are
// current. This is our way of doing a manual
// rescan.

// If we have any blocks to retry, we'll
// defer processing this notification
// until later.

// We'll need to retry the block again
// as we couldn't fetch its filter.

// Since we weren't able to successfully
// process the block, we'll set
// ourselves to not be current in order
// to attempt catching up with the chain
// ourselves.
//
// TODO(wilmer): determine if the error
// is fatal and return it?

// Check whether the block being
// disconnected is one for which we've
// queued up to retry. If it is, we'll
// remove it and any others that follow
// as they are now considered stale.

// Our retry signal has fired, so we'll attempt to
// refetch and notify the filter for our queued blocks.

// We'll go through all of our retry blocks in
// order unless we need to retry any of them.

// We successfully notified the block
// this time, so we can remove it from
// our queue and move on to the next.

// We'll need to retry the block again
// as we couldn't fetch its filter.

// Since we weren't able to successfully
// process the block, we'll set
// ourselves to not be current in order
// to attempt catching up with the chain
// ourselves.
//
// TODO(wilmer): determine if the error
// is fatal and return it?

// If we're not yet current, then we'll walk down the chain
// until we reach the tip of the chain as we know it. At this
// point, we'll be "current" again.

// Apply all queued filter updates.

// Since we're not current, we try to manually advance
// the block. If the next height is above the best
// height known to the chain service, then we mark
// ourselves as current and follow notifications.

// Ensure we cancel the old subscription if
// we're going back to scan for missed blocks.

// Subscribe to block notifications.

// If the next height is known to the chain service,
// then we'll fetch the next block and send a
// notification, maybe also scanning the filters for
// the block.

// waitForBlocks is a helper function that can be used to wait on block
// notifications until the given predicate returns true.
func (rs *rescanState) waitForBlocks(predicate func(hash chainhash.Hash,
	height uint32) bool) error {
	_ = "STUB: not implemented"
	return nil
}

// Before subscribing to block notifications, first check if the
// predicate is not already satisfied by the current best block.

// We'll make sure to process any updates while we're syncing to
// prevent blocking the client.

// A new block notification for the tip of the chain has
// arrived. We'll determine we've caught up to the rescan's
// starting height by receiving a block connected notification
// for the same height.

// If any updates were queued while waiting to catch up to the
// start height of the rescan, apply them now.

// notifyBlock calls appropriate listeners based on the block filter.
func (rs *rescanState) notifyBlock() error { _ = "STUB: not implemented"; return nil }

// Find relevant transactions based on watch list. If scanning is
// false, we can safely assume this block has no relevant transactions.

// If we have a non-empty watch list, then we need to see if it
// matches the rescan's filters, so we get the basic filter
// from the DB or network.

// nolint:staticcheck
// nolint:staticcheck

// handleBlockConnected handles a new block connected notification.
func (rs *rescanState) handleBlockConnected(ntfn *blockntfns.Connected) error {
	_ = "STUB: not implemented"
	return nil
}

// If we've somehow missed a header in the range, then we'll mark
// ourselves as not current so we can walk down the chain and notify the
// callers of blocks we may have missed.

// Ensure the filter header still exists before attempting to fetch the
// filter. This should usually succeed since notifications are delivered
// once filter headers are synced.

// We're only scanning if the header is beyond the horizon of
// our start time.

// If we're not scanning or our watch list is empty, then we can just
// notify the block without fetching any filters/blocks.

// nolint:staticcheck
// nolint:staticcheck

// Otherwise, we'll attempt to fetch the filter to retrieve the relevant
// transactions and notify them.

// If the query failed, then this either means that we don't
// have any peers to fetch this filter from, or the peer(s) that
// we're trying to fetch from are in the progress of a re-org.

// With the block successfully notified, we'll advance our state to it.

// extractBlockMatches fetches the target block from the network, and filters
// out any relevant transactions found within the block.
func extractBlockMatches(chain ChainSource, ro *rescanOptions,
	curStamp *headerfs.BlockStamp, filter *gcs.Filter) ([]*btcutil.Tx,
	error) {
	_ = "STUB: not implemented"

	// We've matched. Now we actually get the block and cycle through the
	// transactions to see which ones are relevant.
	return nil, nil
}

// Before we go through the transactions, let's make sure the filter we
// got from our peer is valid and includes all spent previous output
// scripts. If there's a problem, the error returned here will be
// interpreted by the block manager to disconnect/ban said peer.

// nolint:staticcheck
// nolint:staticcheck

// Even though the transaction may already be known as relevant
// and there might not be a notification callback, we need to
// call paysWatchedAddr anyway as it updates the rescan
// options.

// nolint:staticcheck
// nolint:staticcheck

// handleBlockDisconnected handles a new block disconnected notification.
func (rs *rescanState) handleBlockDisconnected(ntfn *blockntfns.Disconnected) {
	_ = "STUB: not implemented"
	return
}

// Only deal with it if it's the current block we know about. Otherwise,
// it's in the future.

// Run through notifications. This is all single-threaded. We include
// deprecated calls as they're still used, for now.

// nolint:staticcheck
// nolint:staticcheck

// notifyBlockWithFilter calls appropriate listeners based on the block filter.
// This differs from notifyBlock in that is expects the caller to already have
// obtained the target filter.
func (rs *rescanState) notifyBlockWithFilter(header *wire.BlockHeader,
	stamp *headerfs.BlockStamp, filter *gcs.Filter) error {
	_ = "STUB: not implemented"
	return nil
}

// Based on what we find within the block or the filter, we'll be
// sending out a set of notifications with transactions that are
// relevant to the rescan.

// If we actually have a filter, then we'll go ahead an attempt to
// match the items within the filter to ensure we create any relevant
// notifications.

// nolint:staticcheck
// nolint:staticcheck

// matchBlockFilter returns whether the block filter matches the watched items.
// If this returns false, it means the block is certainly not interesting to
// us. This method differs from blockFilterMatches in that it expects the
// filter to already be obtained, rather than fetching the filter from the
// network.
func matchBlockFilter(ro *rescanOptions, filter *gcs.Filter,
	blockHash *chainhash.Hash) (bool, error) {
	_ = "STUB: not implemented"

	// Now that we have the filter as well as the block hash of the block
	// used to construct the filter, we'll check to see if the block
	// matches any items in our watch list.
	return false, nil
}

// blockFilterMatches returns whether the block filter matches the watched
// items. If this returns false, it means the block is certainly not interesting
// to us.
func blockFilterMatches(chain ChainSource, ro *rescanOptions,
	blockHash *chainhash.Hash) (bool, *gcs.Filter, error) {
	_ = "STUB: not implemented"

	// TODO(roasbeef): need to ENSURE always get filter
	return false, nil, nil
}

// Since this method is called when we are not current, and from the
// utxoscanner, we expect more calls to follow for the subsequent
// filters. To speed up the fetching, we make an optimistic batch
// query.

// Block has been reorged out from under us.

// If we found the filter, then we'll check the items in the watch list
// against it.

// updateFilter atomically updates the filter and rewinds to the specified
// height if not 0.
func (ro *rescanOptions) updateFilter(chain ChainSource, update *updateOptions,
	curStamp *headerfs.BlockStamp, curHeader *wire.BlockHeader) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// If we don't need to rewind, then we can exit early.

// If we need to rewind, then we'll walk backwards in the chain until
// we arrive at the block _just_ before the rewind.

// nolint:staticcheck

// nolint:staticcheck

// We just disconnected a block above, so we're now in rewind
// mode. We set this to true here so we properly send
// notifications even if it was just a 1 block rewind.

// Rewind and continue.

// spendsWatchedInput returns whether the transaction matches the filter by
// spending a watched input.
func (ro *rescanOptions) spendsWatchedInput(tx *btcutil.Tx) bool {
	_ = "STUB: not implemented"
	return false
}

// If we're watching for a zero outpoint, then we should
// match on the output script being spent instead.

// Otherwise, we'll match on the outpoint being spent.

// paysWatchedAddr returns whether the transaction matches the filter by having
// an output paying to a watched address. If that is the case, this also
// updates the filter to watch the newly created output going forward.
func (ro *rescanOptions) paysWatchedAddr(tx *btcutil.Tx) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// We'll convert the address into its matching pkScript
// to in order to check for a match.

// If the script doesn't match, we'll move onto the
// next one.

// At this state, we have a matching output so we'll
// mark this transaction as matching.

// Update the filter by also watching this created
// outpoint for the event in the future that it's
// spent.

// Rescan is an object that represents a long-running rescan/notification
// client with updateable filters. It's meant to be close to a drop-in
// replacement for the btcd rescan and notification functionality used in
// wallets. It only contains information about whether a goroutine is running.
type Rescan struct { // nolint:maligned
	started uint32 // To be used atomically.

	running    chan struct{}
	updateChan chan *updateOptions

	options []RescanOption

	chain ChainSource

	errMtx sync.Mutex
	err    error

	wg sync.WaitGroup
}

// NewRescan returns a rescan object that runs in another goroutine and has an
// updatable filter. It returns the long-running rescan object, and a channel
// which returns any error on termination of the rescan process.
func NewRescan(chain ChainSource, options ...RescanOption) *Rescan {
	_ = "STUB: not implemented"
	return nil
}

// WaitForShutdown waits until all goroutines associated with the rescan have
// exited. This method is to be called once the passed quitchan (if any) has
// been closed.
func (r *Rescan) WaitForShutdown() {
	_ = "STUB: not implemented"

	// Start kicks off the rescan goroutine, which will begin to scan the chain
	// according to the specified rescan options.
	return
}

func (r *Rescan) Start() <-chan error { _ = "STUB: not implemented"; return nil }

// nolint

// updateOptions are a set of functional parameters for Update.
type updateOptions struct {
	addrs                    []btcutil.Address
	inputs                   []InputWithScript
	txIDs                    []chainhash.Hash
	rewind                   uint32
	disableDisconnectedNtfns bool
}

// UpdateOption is a functional option argument for the Rescan.Update method.
type UpdateOption func(uo *updateOptions)

func defaultUpdateOptions() *updateOptions { _ = "STUB: not implemented"; return nil }

// AddAddrs adds addresses to the filter.
func AddAddrs(addrs ...btcutil.Address) UpdateOption {
	_ = "STUB: not implemented"
	return *new(UpdateOption)
}

// AddInputs adds inputs to watch to the filter.
func AddInputs(inputs ...InputWithScript) UpdateOption {
	_ = "STUB: not implemented"
	return *new(UpdateOption)
}

// Rewind rewinds the rescan to the specified height (meaning, disconnects down
// to the block immediately after the specified height) and restarts it from
// that point with the (possibly) newly expanded filter. Especially useful when
// called in the same Update() as one of the previous three options.
func Rewind(height uint32) UpdateOption { _ = "STUB: not implemented"; return *new(UpdateOption) }

// DisableDisconnectedNtfns tells the rescan not to send `OnBlockDisconnected`
// and `OnFilteredBlockDisconnected` notifications when rewinding.
func DisableDisconnectedNtfns(disabled bool) UpdateOption {
	_ = "STUB: not implemented"
	return *new(UpdateOption)
}

// Update sends an update to a long-running rescan/notification goroutine.
func (r *Rescan) Update(options ...UpdateOption) error { _ = "STUB: not implemented"; return nil }

// SpendReport is a struct which describes the current spentness state of a
// particular output. In the case that an output is spent, then the spending
// transaction and related details will be populated. Otherwise, only the
// target unspent output in the chain will be returned.
type SpendReport struct {
	// SpendingTx is the transaction that spent the output that a spend
	// report was requested for.
	//
	// NOTE: This field will only be populated if the target output has
	// been spent.
	SpendingTx *wire.MsgTx

	// SpendingTxIndex is the input index of the transaction above which
	// spends the target output.
	//
	// NOTE: This field will only be populated if the target output has
	// been spent.
	SpendingInputIndex uint32

	// SpendingTxHeight is the height of the block that included the
	// transaction  above which spent the target output.
	//
	// NOTE: This field will only be populated if the target output has
	// been spent.
	SpendingTxHeight uint32

	// Output is the raw output of the target outpoint.
	//
	// NOTE: This field will only be populated if the target is still
	// unspent.
	Output *wire.TxOut

	// BlockHash is the block hash of the block that includes the unspent
	// output.
	//
	// NOTE: This field will only be populated if the target is still
	// unspent.
	BlockHash *chainhash.Hash

	// BlockHeight is the height of the block that includes the unspent output.
	//
	// NOTE: This field will only be populated if the target is still
	// unspent.
	BlockHeight uint32

	// BlockIndex is the index of the output's transaction in its block.
	//
	// NOTE: This field will only be populated if the target is still
	// unspent.
	BlockIndex uint32
}

// GetUtxo gets the appropriate TxOut or errors if it's spent. The option
// WatchOutPoints (with a single outpoint) is required. StartBlock can be used
// to give a hint about which block the transaction is in, and TxIdx can be
// used to give a hint of which transaction in the block matches it (coinbase
// is 0, first normal transaction is 1, etc.).
//
// TODO(roasbeef): WTB utxo-commitments.
func (s *ChainService) GetUtxo(options ...RescanOption) (*SpendReport, error) {
	_ = "STUB: not implemented"
	// Before we start we'll fetch the set of default options, and apply
	// any user specified options in a functional manner.
	return nil, nil
}

// As this is meant to fetch UTXO's, the options MUST specify exactly
// one outpoint.

// Wait for the result to be delivered by the rescan or until a shutdown
// is signaled.
