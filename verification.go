package neutrino

import (
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/btcutil/gcs"
)

// VerifyBasicBlockFilter asserts that a given block filter was constructed
// correctly and according to the rules of BIP-0158 to contain both the output's
// pk scripts as well as the pk scripts the inputs are spending.
func VerifyBasicBlockFilter(filter *gcs.Filter, block *btcutil.Block) (int,
	error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Skip coinbase transaction.

// Check outputs first.

// If the script itself is blank, then we'll skip this
// as it doesn't contain any useful information.

// We'll also skip any OP_RETURN scripts as well since
// we don't index these in order to avoid a circular
// dependency.

// Previous versions of the filters did include
// OP_RETURNs. To be able disconnect bad peers
// still serving these old filters we attempt to
// check if there's an unexpected match. Since
// there might be false positives, an OP_RETURN
// can still match filters not including them.
// Therefore, we count the number of such
// unexpected matches for each peer, such that
// we can ban peers matching more than the rest.

// Mark peer bad if we cannot match on
// its filter.

// If it matches on the OP_RETURN output, we
// increase the op return counter.

// This is a "normal" script where we definitely expect
// a match.

// Now we can go through all inputs and check that the filter
// also included any pk scripts of the outputs being _spent_.
// We can do this for witness items since the witness always
// contains the full script as the last element on the stack.

// There are too many edge cases to cover for non-
// witness scripts. And in LN land we're interested in
// witness spends only anyway. Therefore let's skip any
// input that has no witness.
//
// TODO(guggero): Add all those edge cases to
// ComputePkScript?

// The only input type that has both set is a nested
// P2PKH (P2SH-P2WKH). We can verify that one because
// the script hash has to be HASH160(OP_PUSH32 <PKH>).

// Just skip any inputs that we can't derive the pk
// script from.

// Something else went wrong. We can't really say the
// filter is faulty though so we also just skip over
// this input.
