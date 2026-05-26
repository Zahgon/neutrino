package chanutils

import (
	"sync"

	"github.com/lightninglabs/neutrino/cache/lru"
)

const (
	// DefaultQueueSize is the default size to use for concurrent queues.
	DefaultQueueSize = 10
)

// ConcurrentQueue is a typed concurrent-safe FIFO queue with unbounded
// capacity. Clients interact with the queue by pushing items into the in
// channel and popping items from the out channel. There is a goroutine that
// manages moving items from the in channel to the out channel in the correct
// order that must be started by calling Start().
type ConcurrentQueue[T any] struct {
	started sync.Once
	stopped sync.Once

	chanIn   chan T
	chanOut  chan T
	overflow *lru.List[T]

	wg   sync.WaitGroup
	quit chan struct{}
}

// NewConcurrentQueue constructs a ConcurrentQueue. The bufferSize parameter is
// the capacity of the output channel. When the size of the queue is below this
// threshold, pushes do not incur the overhead of the less efficient overflow
// structure.
func NewConcurrentQueue[T any](bufferSize int) *ConcurrentQueue[T] {
	_ = "STUB: not implemented"
	return nil
}

// ChanIn returns a channel that can be used to push new items into the queue.
func (cq *ConcurrentQueue[T]) ChanIn() chan<- T {
	_ = "STUB: not implemented"

	// ChanOut returns a channel that can be used to pop items from the queue.
	return nil
}

func (cq *ConcurrentQueue[T]) ChanOut() <-chan T {
	_ = "STUB: not implemented"

	// Start begins a goroutine that manages moving items from the in channel to the
	// out channel. The queue tries to move items directly to the out channel
	// minimize overhead, but if the out channel is full it pushes items to an
	// overflow queue. This must be called before using the queue.
	return nil
}

func (cq *ConcurrentQueue[T]) Start() { _ = "STUB: not implemented"; return }

func (cq *ConcurrentQueue[T]) start() { _ = "STUB: not implemented"; return }

// Overflow queue is empty so incoming items can
// be pushed directly to the output channel. If
// output channel is full though, push to
// overflow.

// Optimistically push directly
// to chanOut.

// Overflow queue is not empty, so any new items
// get pushed to the back to preserve order.

// Incoming channel has been closed. Empty overflow queue into
// the outgoing channel.

// Close outgoing channel.

// Stop ends the goroutine that moves items from the in channel to the out
// channel. This does not clear the queue state, so the queue can be restarted
// without dropping items.
func (cq *ConcurrentQueue[T]) Stop() { _ = "STUB: not implemented"; return }
