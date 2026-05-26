package blockntfns

import (
	"errors"
	"sync"

	"github.com/lightningnetwork/lnd/queue"
)

var (
	// ErrSubscriptionManagerStopped is an error returned when we attempt to
	// register a new block subscription but the manager has been stopped.
	ErrSubscriptionManagerStopped = errors.New("subscription manager was " +
		"stopped")
)

// newSubscription is an internal message used within the SubscriptionManager to
// denote a new client's intent to receive block notifications.
type newSubscription struct {
	id uint64 // To be used atomically.

	bestHeight uint32

	canceled sync.Once

	ntfnChan  chan BlockNtfn
	ntfnQueue *queue.ConcurrentQueue

	errChan chan error

	quit chan struct{}
	wg   sync.WaitGroup
}

func (s *newSubscription) cancel() { _ = "STUB: not implemented"; return }

// cancelSubscription is an internal message used within the SubscriptionManager
// to denote an existing client's intent to stop receiving block notifications.
type cancelSubscription struct {
	id uint64
}

// NotificationSource is an interface responsible for delivering block
// notifications of a chain.
type NotificationSource interface {
	// Notifications returns a channel through which the latest
	// notifications of the tip of the chain can be retrieved from.
	Notifications() <-chan BlockNtfn

	// NotificationsSinceHeight returns a backlog of block notifications
	// starting from the given height to the tip of the chain.
	//
	// TODO(wilmer): extend with best hash to track reorgs.
	NotificationsSinceHeight(uint32) ([]BlockNtfn, uint32, error)
}

// Subscription represents an intent to receive notifications about the latest
// block events in the chain. The notifications will be streamed through the
// Notifications channel. A Cancel closure is also included to indicate that the
// client no longer wishes to receive any notifications.
type Subscription struct {
	// Notifications is the channel through which block notifications will
	// be sent through.
	//
	// TODO(wilmer): make read-only chan once we remove
	// resetBlockReFetchTimer hack from rescan.
	Notifications chan BlockNtfn

	// Cancel is closure that can be invoked to cancel the client's desire
	// to receive notifications.
	Cancel func()
}

// SubscriptionManager is a system responsible for managing the delivery of
// block notifications for a chain at tip to multiple clients in an asynchronous
// manner.
type SubscriptionManager struct {
	subscriberCounter uint64 // to be used atomically

	started int32 // to be used atomically
	stopped int32 // to be used atomically

	subscribers map[uint64]*newSubscription

	newSubscriptions    chan *newSubscription
	cancelSubscriptions chan *cancelSubscription

	ntfnSource NotificationSource

	quit chan struct{}
	wg   sync.WaitGroup
}

// NewSubscriptionManager creates a subscription manager backed by a
// NotificationSource.
func NewSubscriptionManager(ntfnSource NotificationSource) *SubscriptionManager {
	_ = "STUB: not implemented"
	return nil
}

// Start starts all the goroutines required for the SubscriptionManager to carry
// out its duties.
func (m *SubscriptionManager) Start() { _ = "STUB: not implemented"; return }

// Stop stops all active goroutines required for the SubscriptionManager to
// carry out its duties.
func (m *SubscriptionManager) Stop() { _ = "STUB: not implemented"; return }

// subscriptionHandler is the main event handler of the SubscriptionManager.
// It's responsible for atomically handling notifications for new blocks and
// creating/removing client block subscriptions.
//
// NOTE: This must be run as a goroutine.
func (m *SubscriptionManager) subscriptionHandler() { _ = "STUB: not implemented"; return }

// A new subscription request has been received from a client.

// A request to cancel an existing subscription has been
// received from a client.

// A new block notification for the tip of the chain has been
// received from the backing NotificationSource.

// NewSubscription creates a new block notification subscription for a client.
// The bestHeight parameter can be used by the client to indicate its best known
// state. A backlog of notifications from said point until the tip of the chain
// will be delivered upon the client's successful registration. When providing a
// bestHeight of 0, no backlog will be delivered.
//
// These notifications, along with the latest notifications of the chain, will
// be delivered through the Notifications channel within the Subscription
// returned. A Cancel closure is also provided, in the event that the client
// wishes to no longer receive any notifications.
func (m *SubscriptionManager) NewSubscription(bestHeight uint32) (*Subscription,
	error) {
	_ = "STUB: not implemented"

	// We'll start by constructing the internal messages that the
	// subscription handler will use to register the new client.
	return nil, nil
}

// We'll start the notification queue now so that it is ready in the
// event that a backlog of notifications is to be delivered.

// We'll also start a goroutine that will attempt to consume
// notifications from this queue by delivering them to the client
// itself.

// Now, we can deliver the notification to the subscription handler.

// It's possible that the registration failed if we were unable to
// deliver the backlog of notifications, so we'll make sure to handle
// the error.

// Finally, we can return to the client with its new subscription
// successfully registered.

// handleNewSubscription handles a request to create a new block subscription.
func (m *SubscriptionManager) handleNewSubscription(sub *newSubscription) error {
	_ = "STUB: not implemented"
	return nil
}

// We'll start by retrieving a backlog of notifications from the
// client's best height.

// We'll then attempt to deliver these notifications.

// With the notifications delivered, we can keep track of the new client
// internally in order to deliver new block notifications about the
// chain.

// cancelSubscription sends a request to the subscription handler to cancel an
// existing subscription.
func (m *SubscriptionManager) cancelSubscription(sub *newSubscription) {
	_ = "STUB: not implemented"
	return
}

// handleCancelSubscription handles a request to cancel an existing
// subscription.
func (m *SubscriptionManager) handleCancelSubscription(msg *cancelSubscription) {
	_ = "STUB: not implemented"
	// First, we'll attempt to look up an existing susbcriber with the given
	// ID.
	return
}

// If there is one, we'll stop their internal queue to no longer deliver
// notifications to them.

// notifySubscribers notifies all currently active subscribers about the block.
func (m *SubscriptionManager) notifySubscribers(ntfn BlockNtfn) { _ = "STUB: not implemented"; return }

// notifySubscriber notifies a single subscriber about the block.
func (m *SubscriptionManager) notifySubscriber(sub *newSubscription,
	block BlockNtfn) {
	_ = "STUB: not implemented"
	return
}
