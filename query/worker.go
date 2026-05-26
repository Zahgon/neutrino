package query

import (
	"errors"
	"time"

	"github.com/btcsuite/btcd/wire"
)

var (
	// ErrQueryTimeout is an error returned if the worker doesn't respond
	// with a valid response to the request within the timeout.
	ErrQueryTimeout = errors.New("did not get response before timeout")

	// ErrPeerDisconnected is returned if the worker's peer disconnect
	// before the query has been answered.
	ErrPeerDisconnected = errors.New("peer disconnected")

	// ErrJobCanceled is returned if the job is canceled before the query
	// has been answered.
	ErrJobCanceled = errors.New("job canceled")
)

// queryJob is the internal struct that wraps the Query to work on, in
// addition to some information about the query.
type queryJob struct {
	tries              uint8
	index              uint64
	timeout            time.Duration
	encoding           wire.MessageEncoding
	cancelChan         <-chan struct{}
	internalCancelChan <-chan struct{}
	*Request
}

// queryJob should satisfy the Task interface in order to be sorted by the
// workQueue.
var _ Task = (*queryJob)(nil)

// Index returns the queryJob's index within the work queue.
//
// NOTE: Part of the Task interface.
func (q *queryJob) Index() uint64 {
	_ = "STUB: not implemented"

	// jobResult is the final result of the worker's handling of the queryJob.
	return 0
}

type jobResult struct {
	job  *queryJob
	peer Peer
	err  error
}

// worker is responsible for polling work from its work queue, and handing it
// to the associated peer. It validates incoming responses with the current
// query's response handler, and polls more work for the peer when it has
// successfully received a response to the request.
type worker struct {
	peer Peer

	// nextJob is a channel of queries to be distributed, where the worker
	// will poll new work from.
	nextJob chan *queryJob
}

// A compile-time check to ensure worker satisfies the Worker interface.
var _ Worker = (*worker)(nil)

// NewWorker creates a new worker associated with the given peer.
func NewWorker(peer Peer) Worker { _ = "STUB: not implemented"; return *new(Worker) }

// Run starts the worker. The worker will supply its peer with queries, and
// handle responses from it. Results for any query handled by this worker will
// be delivered on the results channel. quit can be closed to immediately make
// the worker exit.
//
// The method is blocking, and should be started in a goroutine. It will run
// until the peer disconnects or the worker is told to quit.
//
// NOTE: Part of the Worker interface.
func (w *worker) Run(results chan<- *jobResult, quit <-chan struct{}) {
	_ = "STUB: not implemented"

	// Subscribe to messages from the peer.
	return
}

// Poll a new job from the nextJob channel.

// Ignore any message received while not working on anything.

// If the peer disconnected, we can exit immediately, as we
// weren't working on a query.

// There is no point in queueing the request if the job already
// is canceled, so we check this quickly.

// We break to the below loop, where we'll check the
// cancel channel again and the ErrJobCanceled
// result will be sent back.

// We break to the below loop, where we'll check the
// internal cancel channel again and the ErrJobCanceled
// result will be sent back.

// We received a non-canceled query job, send it to the peer.

// Wait for the correct response to be received from the peer,
// or an error happening.

// A message was received from the peer, use the
// response handler to check whether it was answering
// our request.

// If the response did not answer our query, we
// check whether it did progress it.

// If it did make progress we reset the
// timeout. This ensures that the
// queries with multiple responses
// expected won't timeout before all
// responses have been handled.
// TODO(halseth): separate progress
// timeout value.

// We did get a valid response, and can break
// the loop.

// If the timeout is reached before a valid response
// has been received, we exit with an error.

// The query did experience a timeout and will
// be given to someone else.

// If the peer disconnects before giving us a valid
// answer, we'll also exit with an error.

// If the job was canceled, we report this back to the
// work manager.

// Stop to allow garbage collection.

// We have a result ready for the query, hand it off before
// getting a new job.

// If the peer disconnected, we can exit immediately.

// NewJob returns a channel where work that is to be handled by the worker can
// be sent. If the worker reads a queryJob from this channel, it is guaranteed
// that a response will eventually be deliverd on the results channel (except
// when the quit channel has been closed).
//
// NOTE: Part of the Worker interface.
func (w *worker) NewJob() chan<- *queryJob { _ = "STUB: not implemented"; return nil }
