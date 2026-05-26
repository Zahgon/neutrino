package neutrino

import (
	"github.com/btcsuite/btcd/wire"
)

// batchSpendReporter orchestrates the delivery of spend reports to
// GetUtxoRequests processed by the UtxoScanner. The reporter expects a sequence
// of blocks consisting of those containing a UTXO to watch, or any whose
// filter generates a match using current filterEntries. This instance supports
// multiple requests for the same outpoint.
type batchSpendReporter struct {
	// requests maps an outpoint to list of GetUtxoRequests waiting for that
	// UTXO's spend report.
	requests map[wire.OutPoint][]*GetUtxoRequest

	// initialTxns contains a map from an outpoint to the "unspent" version
	// of it's spend report. This value is populated by fetching the output
	// from the block in the request's start height. This spend report will
	// be returned in the case that the output remained unspent for the
	// duration of the scan.
	initialTxns map[wire.OutPoint]*SpendReport

	// outpoints caches the filter entry for each outpoint, conserving
	// allocations when reconstructing the current filterEntries.
	outpoints map[wire.OutPoint][]byte

	// filterEntries holds the current set of watched outpoint, and is
	// applied to cfilters to gauge whether we should download the block.
	//
	// NOTE: This watchlist is updated during each call to ProcessBlock.
	filterEntries [][]byte
}

// newBatchSpendReporter instantiates a fresh batchSpendReporter.
func newBatchSpendReporter() *batchSpendReporter { _ = "STUB: not implemented"; return nil }

// NotifyProgress notifies all requests with the last processed height.
func (b *batchSpendReporter) NotifyProgress(blockHeight uint32) { _ = "STUB: not implemented"; return }

// NotifyUnspentAndUnfound iterates through any requests for which no spends
// were detected. If we were able to find the initial output, this will be
// delivered signaling that no spend was detected. If the original output could
// not be found, a nil spend report is returned.
func (b *batchSpendReporter) NotifyUnspentAndUnfound() { _ = "STUB: not implemented"; return }

// A nil SpendReport indicates the output was not found.

// FailRemaining will return an error to all remaining requests in the event we
// experience a critical rescan error. The error is threaded through to allow
// the syntax:
//
//	return reporter.FailRemaining(err)
func (b *batchSpendReporter) FailRemaining(err error) error { _ = "STUB: not implemented"; return nil }

// notifyRequests delivers the same final response to the given requests, and
// cleans up any remaining state for the outpoint.
//
// NOTE: AT MOST ONE of `report` or `err` may be non-nil.
func (b *batchSpendReporter) notifyRequests(
	outpoint *wire.OutPoint,
	requests []*GetUtxoRequest,
	report *SpendReport,
	err error) {
	_ = "STUB: not implemented"
	return
}

// ProcessBlock accepts a block, block height, and any new requests whose start
// height matches the provided height. If a non-zero number of new requests are
// presented, the block will first be checked for the initial outputs from which
// spends may occur. Afterwards, any spends detected in the block are
// immediately dispatched, and the watchlist updated in preparation of filtering
// the next block.
func (b *batchSpendReporter) ProcessBlock(blk *wire.MsgBlock,
	newReqs []*GetUtxoRequest, height uint32) {
	_ = "STUB: not implemented"

	// If any requests want the UTXOs at this height, scan the block to find
	// the original outputs that might be spent from.
	return
}

// Next, filter the block for any spends using the current set of
// watched outpoints. This will include any new requests added above.

// Finally, rebuild filter entries from cached entries remaining in
// outpoints map. This will provide an updated watchlist used to scan
// the subsequent filters.

// addNewRequests adds a set of new GetUtxoRequests to the spend reporter's
// state. This method immediately adds the request's outpoints to the reporter's
// watchlist.
func (b *batchSpendReporter) addNewRequests(reqs []*GetUtxoRequest) {
	_ = "STUB: not implemented"
	return
}

// Build the filter entry only if it is the first time seeing
// the outpoint.

// findInitialTransactions searches the given block for the creation of the
// UTXOs that are supposed to be birthed in this block. If any are found, a
// spend report containing the initial outpoint will be saved in case the
// outpoint is not spent later on. Requests corresponding to outpoints that are
// not found in the block will return a nil spend report to indicate that the
// UTXO was not found.
func (b *batchSpendReporter) findInitialTransactions(block *wire.MsgBlock,
	newReqs []*GetUtxoRequest, height uint32) map[wire.OutPoint]*SpendReport {
	_ = "STUB: not implemented"

	// First, construct  a reverse index from txid to all a list of requests
	// whose outputs share the same txid.
	return nil
}

// Iterate over the transactions in this block, hashing each and
// querying our reverse index to see if any requests depend on the txn.

// If our reverse index has been cleared, we are done.

// For all requests that are watching this txid, use the output
// index of each to grab the initial output.

// Ensure that the outpoint's index references an actual
// output on the transaction. If not, we will be unable
// to find the initial output.

// Finally, we must reconcile any requests for which the txid did not
// exist in this block. A nil spend report is saved for every initial
// txn that could not be found, otherwise the result is copied from scan
// above. The copied values can include valid initial txns, as well as
// nil spend report if the output index was invalid.

// notifySpends finds any transactions in the block that spend from our watched
// outpoints. If a spend is detected, it is immediately delivered and cleaned up
// from the reporter's internal state.
func (b *batchSpendReporter) notifySpends(block *wire.MsgBlock,
	height uint32) map[wire.OutPoint]*SpendReport {
	_ = "STUB: not implemented"
	return nil
}

// Check each input to see if this transaction spends one of our
// watched outpoints.

// Find the requests this spend relates to.

// With the requests located, we remove this outpoint
// from both the requests, outpoints, and initial txns
// map. This will ensures we don't continue watching
// this outpoint.
