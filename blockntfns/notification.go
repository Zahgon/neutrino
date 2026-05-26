package blockntfns

import (
	"github.com/btcsuite/btcd/wire"
)

// BlockNtfn is an interface that coalesces all the different types of block
// notifications.
type BlockNtfn interface {
	// Header returns the header of the block for which this notification is
	// for.
	Header() wire.BlockHeader

	// Height returns the height of the block for which this notification is
	// for.
	Height() uint32

	// ChainTip returns the header of the new tip of the chain after
	// processing the block being connected/disconnected.
	ChainTip() wire.BlockHeader
}

// Connected is a block notification that gets dispatched to clients when the
// filter header of a new block has been found that extends the current chain.
type Connected struct {
	header wire.BlockHeader
	height uint32
}

// A compile-time check to ensure Connected satisfies the BlockNtfn interface.
var _ BlockNtfn = (*Connected)(nil)

// NewBlockConnected creates a new Connected notification for the given block.
func NewBlockConnected(header wire.BlockHeader, height uint32) *Connected {
	_ = "STUB: not implemented"
	return nil
}

// Header returns the header of the block extending the chain.
func (n *Connected) Header() wire.BlockHeader {
	_ = "STUB: not implemented"

	// Height returns the height of the block extending the chain.
	return *new(wire.BlockHeader)
}

func (n *Connected) Height() uint32 {
	_ = "STUB: not implemented"

	// ChainTip returns the header of the new tip of the chain after processing the
	// block being connected.
	return 0
}

func (n *Connected) ChainTip() wire.BlockHeader {
	_ = "STUB: not implemented"

	// String returns the string representation of a Connected notification.
	return *new(wire.BlockHeader)
}

func (n *Connected) String() string { _ = "STUB: not implemented"; return "" }

// Disconnected if a notification that gets dispatched to clients when a reorg
// has been detected at the tip of the chain.
type Disconnected struct {
	headerDisconnected wire.BlockHeader
	heightDisconnected uint32
	chainTip           wire.BlockHeader
}

// A compile-time check to ensure Disconnected satisfies the BlockNtfn
// interface.
var _ BlockNtfn = (*Disconnected)(nil)

// NewBlockDisconnected creates a Disconnected notification for the given block.
func NewBlockDisconnected(headerDisconnected wire.BlockHeader,
	heightDisconnected uint32, chainTip wire.BlockHeader) *Disconnected {
	_ = "STUB: not implemented"
	return nil
}

// Header returns the header of the block being disconnected.
func (n *Disconnected) Header() wire.BlockHeader {
	_ = "STUB: not implemented"
	return *new(wire.BlockHeader)
}

// Height returns the height of the block being disconnected.
func (n *Disconnected) Height() uint32 { _ = "STUB: not implemented"; return 0 }

// ChainTip returns the header of the new tip of the chain after processing the
// block being disconnected.
func (n *Disconnected) ChainTip() wire.BlockHeader {
	_ = "STUB: not implemented"

	// String returns the string representation of a Disconnected notification.
	return *new(wire.BlockHeader)
}

func (n *Disconnected) String() string { _ = "STUB: not implemented"; return "" }
