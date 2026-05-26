// NOTE: THIS API IS UNSTABLE RIGHT NOW.
// TODO: Add functional options to ChainService instantiation.

package neutrino

import (
	"context"
	"net"
	"sync"
	"time"

	"github.com/btcsuite/btcd/addrmgr"
	"github.com/btcsuite/btcd/blockchain"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/connmgr"
	"github.com/btcsuite/btcd/peer"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcwallet/walletdb"
	"github.com/lightninglabs/neutrino/banman"
	"github.com/lightninglabs/neutrino/blockntfns"
	"github.com/lightninglabs/neutrino/cache/lru"
	"github.com/lightninglabs/neutrino/chanutils"
	"github.com/lightninglabs/neutrino/filterdb"
	"github.com/lightninglabs/neutrino/headerfs"
	"github.com/lightninglabs/neutrino/pushtx"
	"github.com/lightninglabs/neutrino/query"
)

// These are exported variables so they can be changed by users.
//
// TODO: Export functional options for these as much as possible so they can be
// changed call-to-call.
var (
	// ConnectionRetryInterval is the base amount of time to wait in
	// between retries when connecting to persistent peers.  It is adjusted
	// by the number of retries such that there is a retry backoff.
	ConnectionRetryInterval = time.Second * 5

	// UserAgentName is the user agent name and is used to help identify
	// ourselves to other bitcoin peers.
	UserAgentName = "neutrino"

	// UserAgentVersion is the user agent version and is used to help
	// identify ourselves to other bitcoin peers.
	UserAgentVersion = "0.17.1"

	// Services describes the services that are supported by the server.
	Services = wire.SFNodeWitness | wire.SFNodeCF

	// RequiredServices describes the services that are required to be
	// supported by outbound peers.
	RequiredServices = wire.SFNodeNetwork | wire.SFNodeWitness | wire.SFNodeCF

	// BanThreshold is the maximum ban score before a peer is banned.
	BanThreshold = uint32(100)

	// BanDuration is the duration of a ban.
	BanDuration = time.Hour * 24

	// TargetOutbound is the number of outbound peers to target.
	TargetOutbound = 8

	// MaxPeers is the maximum number of connections the client maintains.
	MaxPeers = 125

	// DisableDNSSeed disables getting initial addresses for Bitcoin nodes
	// from DNS.
	DisableDNSSeed = false

	// DefaultFilterCacheSize is the size (in bytes) of filters neutrino
	// will keep in memory if no size is specified in the neutrino.Config.
	// Since we utilize the cache during batch filter fetching, it is
	// beneficial if it is able to keep a whole batch. The current batch
	// size is 1000, so we default to 30 MB, which can fit about 1450 to
	// 2300 mainnet filters.
	DefaultFilterCacheSize uint64 = 3120 * 10 * 1000

	// DefaultBlockCacheSize is the size (in bytes) of blocks neutrino will
	// keep in memory if no size is specified in the neutrino.Config.
	DefaultBlockCacheSize uint64 = 4096 * 10 * 1000 // 40 MB
)

// isDevNetwork indicates if the chain is a private development network, namely
// simnet or regtest/regnet.
func isDevNetwork(net wire.BitcoinNet) bool { _ = "STUB: not implemented"; return false }

// updatePeerHeightsMsg is a message sent from the blockmanager to the server
// after a new block has been accepted. The purpose of the message is to update
// the heights of peers that were known to announce the block before we
// connected it to the main chain or recognized it as an orphan. With these
// updates, peer heights will be kept up to date, allowing for fresh data when
// selecting sync peer candidacy.
type updatePeerHeightsMsg struct {
	newHash    *chainhash.Hash
	newHeight  int32
	originPeer *ServerPeer
}

// peerState maintains state of inbound, persistent, outbound peers as well
// as banned peers and outbound groups.
type peerState struct {
	outboundPeers   map[int32]*ServerPeer
	persistentPeers map[int32]*ServerPeer
	outboundGroups  map[string]int
}

// Count returns the count of all known peers.
func (ps *peerState) Count() int { _ = "STUB: not implemented"; return 0 }

// forAllOutboundPeers is a helper function that runs closure on all outbound
// peers known to peerState.
func (ps *peerState) forAllOutboundPeers(closure func(sp *ServerPeer)) {
	_ = "STUB: not implemented"
	return
}

// forAllPeers is a helper function that runs closure on all peers known to
// peerState.
func (ps *peerState) forAllPeers(closure func(sp *ServerPeer)) { _ = "STUB: not implemented"; return }

// spMsg represents a message over the wire from a specific peer.
type spMsg struct {
	sp  *ServerPeer
	msg wire.Message
}

// spMsgSubscription sends all messages from a peer over a channel, allowing
// pluggable filtering of the messages.
type spMsgSubscription struct {
	msgChan  chan<- spMsg
	quitChan <-chan struct{}
}

// msgSubscription sends all messages from a peer over a channel, allowing
// pluggable filtering of the messages.
type msgSubscription struct {
	msgChan  chan<- wire.Message
	quitChan <-chan struct{}
}

// ServerPeer extends the peer to maintain state shared by the server and the
// blockmanager.
type ServerPeer struct {
	// The following variables must only be used atomically
	feeFilter int64

	*peer.Peer

	connReq        *connmgr.ConnReq
	server         *ChainService
	persistent     bool
	knownAddresses *lru.Cache[string, *cachedAddr]
	quit           chan struct{}

	// The following map of subcribers is used to subscribe to messages
	// from the peer. This allows broadcast to multiple subscribers at
	// once, allowing for multiple queries to be going to multiple peers at
	// any one time. The mutex is for subscribe/unsubscribe functionality.
	// The sends on these channels WILL NOT block; any messages the channel
	// can't accept will be dropped silently.
	// TODO(halseth): remove one of the maps when all queries go through
	// work manager.
	recvSubscribers  map[spMsgSubscription]struct{}
	recvSubscribers2 map[msgSubscription]struct{}
	mtxSubscribers   sync.RWMutex
}

// NewServerPeer returns a new ServerPeer instance. The peer needs to be set by
// the caller.
func NewServerPeer(s *ChainService, isPersistent bool) *ServerPeer {
	_ = "STUB: not implemented"
	return nil
}

// newestBlock returns the current best block hash and height using the format
// required by the configuration for the peer package.
func (sp *ServerPeer) newestBlock() (*chainhash.Hash, int32, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// addKnownAddresses adds the given addresses to the set of known addresses to
// the peer to prevent sending duplicate addresses.
func (sp *ServerPeer) addKnownAddresses(addresses []*wire.NetAddressV2) {
	_ = "STUB: not implemented"
	return
}

// OnVerAck is invoked when a peer receives a verack bitcoin message and is used
// to kick start communication with them.
func (sp *ServerPeer) OnVerAck(_ *peer.Peer, msg *wire.MsgVerAck) {
	_ = "STUB: not implemented"
	return

	// OnVersion is invoked when a peer receives a version bitcoin message
	// and is used to negotiate the protocol version details as well as kick start
	// the communications.
}

func (sp *ServerPeer) OnVersion(_ *peer.Peer, msg *wire.MsgVersion) *wire.MsgReject {
	_ = "STUB: not implemented"
	// Add the remote peer time as a sample for creating an offset against
	// the local clock to keep the network time in sync.
	return nil
}

// Check to see if the peer supports the latest protocol version and
// service bits required to service us. If not, then we'll disconnect
// so we can find compatible peers.

// Disconnect the peer even though BanPeer attempts to do so
// because it has yet to be added.

// Update the address manager with the advertised services for outbound
// connections in case they have changed. This is not done for inbound
// connections to help prevent malicious behavior and is skipped when
// running on the simulation test network since it is only intended to
// connect to specified peers and actively avoids advertising and
// connecting to discovered peers.

// OnInv is invoked when a peer receives an inv bitcoin message and is
// used to examine the inventory being advertised by the remote peer and react
// accordingly.  We pass the message down to blockmanager which will call
// QueueMessage with any appropriate responses.
func (sp *ServerPeer) OnInv(p *peer.Peer, msg *wire.MsgInv) { _ = "STUB: not implemented"; return }

// OnHeaders is invoked when a peer receives a headers bitcoin
// message.  The message is passed down to the block manager.
func (sp *ServerPeer) OnHeaders(p *peer.Peer, msg *wire.MsgHeaders) {
	_ = "STUB: not implemented"
	return
}

// OnFeeFilter is invoked when a peer receives a feefilter bitcoin message and
// is used by remote peers to request that no transactions which have a fee rate
// lower than provided value are inventoried to them.  The peer will be
// disconnected if an invalid fee filter value is provided.
func (sp *ServerPeer) OnFeeFilter(_ *peer.Peer, msg *wire.MsgFeeFilter) {
	_ = "STUB: not implemented"
	// Check that the passed minimum fee is a valid amount.
	return
}

// OnReject is invoked when a peer receives a reject bitcoin message and is
// used to notify the server about a rejected transaction.
func (sp *ServerPeer) OnReject(_ *peer.Peer, msg *wire.MsgReject) {
	_ = "STUB: not implemented"
	// TODO(roaseef): log?
	return

	// OnAddr is invoked when a peer receives an addr bitcoin message and is
	// used to notify the server about advertised addresses.
}

func (sp *ServerPeer) OnAddr(_ *peer.Peer, msg *wire.MsgAddr) {
	_ = "STUB: not implemented"
	// Ignore addresses when running on a private development network.  This
	// helps prevent the network from becoming another public test network
	// since it will not be able to learn about other peers that have not
	// specifically been provided.
	return
}

// Ignore old style addresses which don't include a timestamp.

// A message that has no addresses is invalid.

// Don't add more address if we're disconnecting.

// Skip any that don't advertise our required services.

// Set the timestamp to 5 days ago if it's more than 10 minutes
// in the future so this address is one of the first to be
// removed when space is needed.

// Convert the wire.NetAddress to wire.NetAddressV2 since that
// is what is used by the addrmgr.

// Ignore any addr messages if none of them contained our required
// services.

// Add addresses to the set of known addresses for this peer.

// Add addresses to server address manager.  The address manager handles
// the details of things such as preventing duplicate addresses, max
// addresses, and last seen updates.
// XXX bitcoind gives a 2 hour time penalty here, do we want to do the
// same?

// OnAddrV2 is called when a peer receives an AddrV2 message from its peer.
func (sp *ServerPeer) OnAddrV2(_ *peer.Peer, msg *wire.MsgAddrV2) {
	_ = "STUB: not implemented"
	// Ignore addresses when running on a private development network for
	// the same reason that OnAddr does.
	return
}

// An empty AddrV2 message is invalid.

// Don't add more addresses if we're disconnecting.

// Skip any that don't advertise our required services.

// Set the timestamp to 5 days ago if it's more than 10 minutes
// in the future so this address is one of the first to be
// removed when space is needed.

// Ignore addrv2 message if no addresses contained our required
// services.

// Add the addresses to the set of known addresses for this peer.

// Add addresses to the address manager.

// OnRead is invoked when a peer receives a message and it is used to update
// the bytes received by the server.
func (sp *ServerPeer) OnRead(_ *peer.Peer, bytesRead int, msg wire.Message,
	err error) {
	_ = "STUB: not implemented"
	return
}

// Send a message to each subscriber. Each message gets its own
// goroutine to prevent blocking on the mutex lock.
// TODO: Flood control.

// Quickly determine if this subscription has been canceled, if
// so delete it.

// subscribeRecvMsg handles adding OnRead subscriptions to the server peer.
func (sp *ServerPeer) subscribeRecvMsg(subscription spMsgSubscription) {
	_ = "STUB: not implemented"
	return
}

// unsubscribeRecvMsgs handles removing OnRead subscriptions from the server
// peer.
func (sp *ServerPeer) unsubscribeRecvMsgs(subscription spMsgSubscription) {
	_ = "STUB: not implemented"
	return
}

// A compile-time check to ensure that ServerPeer implements the query.Peer
// interface.
var _ query.Peer = (*ServerPeer)(nil)

// SubscribeRecvMsg adds a OnRead subscription to the peer. All bitcoin
// messages received from this peer will be sent on the returned channel. A
// closure is also returned, that should be called to cancel the subscription.
//
// NOTE: Part of the query.Peer interface.
func (sp *ServerPeer) SubscribeRecvMsg() (<-chan wire.Message, func()) {
	_ = "STUB: not implemented"
	// We won't have to buffer this channel, since we'll always send on it
	// from a new goroutine.
	return nil, nil
}

// OnDisconnect returns a channel that will be closed when this peer is
// disconnected.
//
// NOTE: Part of the query.Peer interface.
func (sp *ServerPeer) OnDisconnect() <-chan struct{} {
	_ = "STUB: not implemented"

	// OnWrite is invoked when a peer sends a message and it is used to update
	// the bytes sent by the server.
	return nil
}

func (sp *ServerPeer) OnWrite(_ *peer.Peer, bytesWritten int, msg wire.Message, err error) {
	_ = "STUB: not implemented"
	return
}

// Config is a struct detailing the configuration of the chain service.
type Config struct {
	// DataDir is the directory that neutrino will store all header
	// information within.
	DataDir string

	// Database is an *open* database instance that we'll use to storm
	// indexes of the chain.
	Database walletdb.DB

	// ChainParams is the chain that we're running on.
	ChainParams chaincfg.Params

	// ConnectPeers is a slice of hosts that should be connected to on
	// startup, and be established as persistent peers.
	//
	// NOTE: If specified, we'll *only* connect to this set of peers and
	// won't attempt to automatically seek outbound peers.
	ConnectPeers []string

	// AddPeers is a slice of hosts that should be connected to on startup,
	// and be maintained as persistent peers.
	AddPeers []string

	// Dialer is an optional function closure that will be used to
	// establish outbound TCP connections. If specified, then the
	// connection manager will use this in place of net.Dial for all
	// outbound connection attempts.
	Dialer func(addr net.Addr) (net.Conn, error)

	// NameResolver is an optional function closure that will be used to
	// lookup the IP of any host. If specified, then the address manager,
	// along with regular outbound connection attempts will use this
	// instead.
	NameResolver func(host string) ([]net.IP, error)

	// FilterCacheSize indicates the size (in bytes) of filters the cache will
	// hold in memory at most.
	FilterCacheSize uint64

	// BlockCache is an LRU block cache. If none is provided then the a new
	// one will be instantiated.
	BlockCache *lru.Cache[wire.InvVect, *CacheableBlock]

	// BlockCacheSize indicates the size (in bytes) of blocks the block
	// cache will hold in memory at most. If a BlockCache is provided then
	// BlockCacheSize is ignored.
	BlockCacheSize uint64

	// persistToDisk indicates whether the filter should also be written
	// to disk in addition to the memory cache. For "normal" wallets, they'll
	// almost never need to re-match a filter once it's been fetched unless
	// they're doing something like a key import.
	PersistToDisk bool

	// HeadersImport contains configuration options for importing headers
	// from external sources. When these options are set, neutrino will
	// attempt to import headers from file before falling back to P2P
	// synchronization.
	HeadersImport *HeadersImportConfig

	// AssertFilterHeader is an optional field that allows the creator of
	// the ChainService to ensure that if any chain data exists, it's
	// compliant with the expected filter header state. If neutrino starts
	// up and this filter header state has diverged, then it'll remove the
	// current on disk filter headers to sync them anew.
	AssertFilterHeader *headerfs.FilterHeader

	// BroadcastTimeout is the amount of time we'll wait before giving up on
	// a transaction broadcast attempt. Broadcasting transactions consists
	// of three steps:
	//
	// 1. Neutrino sends an inv for the transaction.
	// 2. The recipient node determines if the inv is known, and if it's
	//    not, replies with a getdata message.
	// 3. Neutrino sends the raw transaction.
	BroadcastTimeout time.Duration
}

// HeadersImportConfig contains configuration options for importing headers
// from external sources.
type HeadersImportConfig struct {
	// BlockHeadersSource specifies where to obtain block headers from.
	// This could be a file path, URL, or other source identifier.
	BlockHeadersSource string

	// FilterHeadersSource specifies where to obtain filter headers from.
	// This could be a file path, URL, or other source identifier.
	FilterHeadersSource string

	// WriteBatchSizePerRegion defines the number of headers to write in a
	// single batch per processing region during import. The import process
	// divides header ranges into distinct regions
	// (overlap, divergence, and new headers), each with different
	// processing requirements. Each region is processed in batches of this
	// size to optimize database performance while maintaining data
	// integrity boundaries. Larger values improve speed but increase memory
	// usage.
	//
	// Default value: 16,384 (2^14) entries.
	// This results in 16,384 block headers and 16,384 filter headers per
	// batch. The total upper bound size per batch is:
	//   - Block headers: (16,384 * 80 bytes) / (2^10) = 1.28 MB
	//   - Filter headers: (16,384 * 32 bytes) / (2^10) = 0.5 MB
	// Peak memory usage observed during benchmarking was ≈ 66 MB.
	// Actual memory usage can vary depending on factors such as database
	// state, Go garbage collector semantics and activity.
	WriteBatchSizePerRegion int

	// ValidationFlags specifies the behavior flags used during header
	// validation. It defaults to BFNone.
	ValidationFlags blockchain.BehaviorFlags
}

// peerSubscription holds a peer subscription which we'll notify about any
// connected peers.
type peerSubscription struct {
	peers  chan<- query.Peer
	cancel <-chan struct{}
}

// ChainService is instantiated with functional options.
type ChainService struct { // nolint:maligned
	// The following variables must only be used atomically.
	// Putting the uint64s first makes them 64-bit aligned for 32-bit systems.
	bytesReceived uint64 // Total bytes received from all peers since start.
	bytesSent     uint64 // Total bytes sent by all peers since start.
	started       int32
	shutdown      int32

	FilterDB         filterdb.FilterDatabase
	BlockHeaders     headerfs.BlockHeaderStore
	RegFilterHeaders headerfs.FilterHeaderStore
	persistToDisk    bool

	headersImport *HeadersImportConfig

	FilterCache *lru.Cache[FilterCacheKey, *CacheableFilter]
	BlockCache  *lru.Cache[wire.InvVect, *CacheableBlock]

	chainParams          chaincfg.Params
	addrManager          *addrmgr.AddrManager
	connManager          *connmgr.ConnManager
	blockManager         *blockManager
	blockSubscriptionMgr *blockntfns.SubscriptionManager
	newPeers             chan *ServerPeer
	donePeers            chan *ServerPeer
	query                chan interface{}
	firstPeerConnect     chan struct{}
	peerHeightsUpdate    chan updatePeerHeightsMsg
	wg                   sync.WaitGroup
	quit                 chan struct{}
	timeSource           blockchain.MedianTimeSource
	services             wire.ServiceFlag
	utxoScanner          *UtxoScanner
	broadcaster          *pushtx.Broadcaster
	banStore             banman.Store
	workManager          query.WorkManager
	filterBatchWriter    *chanutils.BatchWriter[*filterdb.FilterData]

	// peerSubscribers is a slice of active peer subscriptions, that we
	// will notify each time a new peer is connected.
	peerSubscribers []*peerSubscription

	// TODO: Add a map for more granular exclusion?
	mtxCFilter sync.Mutex

	userAgentName    string
	userAgentVersion string

	nameResolver func(string) ([]net.IP, error)
	dialer       func(net.Addr) (net.Conn, error)

	broadcastTimeout time.Duration
}

// NewChainService returns a new chain service configured to connect to the
// bitcoin network type specified by chainParams.  Use start to begin syncing
// with peers.
func NewChainService(cfg Config) (*ChainService, error) {
	_ = "STUB: not implemented"
	// Use the default broadcast timeout if one isn't provided.
	return nil, nil
}

// First, we'll sort out the methods that we'll use to established
// outbound TCP connections, as well as perform any DNS queries.
//
// If the dialler was specified, then we'll use that in place of the
// default net.Dial function.

// Similarly, if the user specified as function to use for name
// resolution, then we'll use that everywhere as well.

// When creating the addr manager, we'll check to see if the user has
// provided their own resolution function. If so, then we'll use that
// instead as this may be proxying requests over an anonymizing
// network.

// Only setup a function to return new addresses to connect to when not
// running in connect-only mode.  Private development networks are always in
// connect-only mode since it is only intended to connect to specified peers
// and actively avoid advertising and connecting to discovered peers in
// order to prevent it from becoming a public test network.

// Gather our set of currently connected peers to avoid
// connecting to them again.

// Ignore peers that we've already banned.

// Skip any addresses that correspond to our set
// of currently connected peers.

// The peer behind this address should support
// all of our required services.

// Address will not be invalid, local or unroutable
// because addrmanager rejects those on addition.
// Just check that we don't already have an address
// in the same group so that we are not connecting
// to the same network segment at the expense of
// others.

// only allow recent nodes (10mins) after we failed 30
// times

// allow nondefault ports after 50 failed tries.

// Mark an attempt for the valid address.

// Create a connection manager.

// Start up persistent peers.

// Since netwok access might not be established yet, we
// loop until we are able to look up the permanent
// peer.

// Try again in 5 seconds.

// BestBlock retrieves the most recent block's height and hash where we
// have both the header and filter header ready.
func (s *ChainService) BestBlock() (*headerfs.BlockStamp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Filter headers might lag behind block headers, so we can fetch a
// previous block header if the filter headers are not caught up.

// GetBlockHash returns the block hash at the given height.
func (s *ChainService) GetBlockHash(height int64) (*chainhash.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBlockHeader returns the block header for the given block hash, or an
// error if the hash doesn't exist or is unknown.
func (s *ChainService) GetBlockHeader(
	blockHash *chainhash.Hash) (*wire.BlockHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBlockHeight gets the height of a block by its hash. An error is returned
// if the given block hash is unknown.
func (s *ChainService) GetBlockHeight(hash *chainhash.Hash) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// BanPeer disconnects and bans a peer due to a specific reason for a duration
// of BanDuration.
func (s *ChainService) BanPeer(addr string, reason banman.Reason) error {
	_ = "STUB: not implemented"
	return nil
}

// We'll want to disconnect the peer after we return regardless of
// whether we ban the peer or not. We do this to prevent a possible race
// condition where we end up reconnecting with the peer slightly
// before the ban succeeds.

// We do so in a goroutine to prevent blocking if the server is
// handling a query or a new/stale peer.

// UnbanPeer connects and unbans a previously banned peer.
func (s *ChainService) UnbanPeer(addr string, parmanent bool) error {
	_ = "STUB: not implemented"
	return nil
}

// IsBanned returns true if the peer is banned, and false otherwise.
func (s *ChainService) IsBanned(addr string) bool { _ = "STUB: not implemented"; return false }

// Log how much time left the peer will remain banned for, if any.

// AddPeer adds a new peer that has already been connected to the server.
func (s *ChainService) AddPeer(sp *ServerPeer) { _ = "STUB: not implemented"; return }

// AddBytesSent adds the passed number of bytes to the total bytes sent counter
// for the server.  It is safe for concurrent access.
func (s *ChainService) AddBytesSent(bytesSent uint64) { _ = "STUB: not implemented"; return }

// AddBytesReceived adds the passed number of bytes to the total bytes received
// counter for the server.  It is safe for concurrent access.
func (s *ChainService) AddBytesReceived(bytesReceived uint64) { _ = "STUB: not implemented"; return }

// NetTotals returns the sum of all bytes received and sent across the network
// for all peers.  It is safe for concurrent access.
func (s *ChainService) NetTotals() (uint64, uint64) { _ = "STUB: not implemented"; return 0, 0 }

// peerHandler is used to handle peer operations such as adding and removing
// peers to and from the server, banning peers, and broadcasting messages to
// peers.  It must be run in a goroutine.
func (s *ChainService) peerHandler() { _ = "STUB: not implemented"; return }

// Add peers discovered through DNS to the address manager.

// Bitcoind uses a lookup of the dns seeder
// here. This is rather strange since the
// values looked up by the DNS seed lookups
// will vary quite a lot.  to replicate this
// behaviour we put all addresses as having
// come from the first one.

// New peers connected to the server.

// Disconnected peers.

// Block accepted in mainchain or orphan, update peer height.

// Disconnect all peers on server shutdown.

// Drain channels before exiting so nothing is left waiting around
// to send.

// addrStringToNetAddr takes an address in the form of 'host:port' or 'host'
// and returns a net.Addr which maps to the original address with any host
// names resolved to IP addresses and a default port added, if not specified,
// from the ChainService's network parameters.
func (s *ChainService) addrStringToNetAddr(addr string) (net.Addr, error) {
	_ = "STUB: not implemented"
	return *new(net.Addr), nil
}

// Tor addresses cannot be resolved to an IP, so just return onionAddr
// instead.

// Attempt to look up an IP address associated with the parsed host.

// handleUpdatePeerHeight updates the heights of all peers who were known to
// announce a block we recently accepted.
func (s *ChainService) handleUpdatePeerHeights(state *peerState, umsg updatePeerHeightsMsg) {
	_ = "STUB: not implemented"
	return
}

// The origin peer should already have the updated height.

// This is a pointer to the underlying memory which doesn't
// change.

// Skip this peer if it hasn't recently announced any new blocks.

// If the peer has recently announced a block, and this block
// matches our newly accepted block, then update their block
// height.

// handleAddPeerMsg deals with adding new peers.  It is invoked from the
// peerHandler goroutine.
func (s *ChainService) handleAddPeerMsg(state *peerState, sp *ServerPeer) bool {
	_ = "STUB: not implemented"
	return false
}

// Ignore new peers if we're shutting down.

// Disconnect banned peers.

// TODO: Check for max peers from a single IP.

// Limit max number of total peers.

// TODO: how to handle permanent peers here?
// they should be rescheduled.

// Add the new peer and start it.

// Close firstPeerConnect channel so blockManager will be notified.

// Update the address' last seen time if the peer has acknowledged our
// version and has sent us its version as well.

// Signal the block manager this peer is a new sync candidate.

// Update the address manager and request known addresses from the
// remote peer for outbound connections. This is skipped when running on
// a development network since it is only intended to connect to
// specified peers and actively avoids advertising and connecting to
// discovered peers.

// Request known addresses if the server address manager needs
// more and the peer has a protocol version new enough to
// include a timestamp with addresses.

// Add the address to the addr manager anew, and also mark it as
// a good address.

// We'll go through each peer subscriber and notify it about the added
// peer.

// Quickly check whether this subscription has been canceled.

// Avoid GC leak.

// Keep non-canceled subscribers around.

// Send a notification in a goroutine to avoid blocking the
// peerHandler.

// Re-align the slice to only active subscribers.

// notifyConnectedPeer sends the given peer to the peerSubsription.
//
// NOTE: MUST be run as a goroutine.
func (s *ChainService) notifyConnectedPeer(
	sub *peerSubscription, sp *ServerPeer) {
	_ = "STUB: not implemented"
	return
}

// handleDonePeerMsg deals with peers that have signalled they are done.  It is
// invoked from the peerHandler goroutine.
func (s *ChainService) handleDonePeerMsg(state *peerState, sp *ServerPeer) {
	_ = "STUB: not implemented"
	// If the peer is being tracked internally, i.e., we received their
	// VerAck, we'll need to remove them.
	return
}

// Only request a new connection if the peer being disconnected is not
// persistent and we still need more peer connections. There's no need
// to do so if the peer is persistent since the connection manager will
// attempt to reconnect.

// disconnectPeer attempts to drop the connection of a targeted peer in the
// passed peer list. Targets are identified via usage of the passed
// `compareFunc`, which should return `true` if the passed peer is the target
// peer. This function returns true on success and false if the peer is unable
// to be located. If the peer is found, and the passed callback: `whenFound'
// isn't nil, we call it with the peer as the argument before it is removed
// from the peerList, and is disconnected from the server.
func disconnectPeer(peerList map[int32]*ServerPeer,
	compareFunc func(*ServerPeer) bool, whenFound func(*ServerPeer)) bool {
	_ = "STUB: not implemented"
	return false
}

// This is ok because we are not continuing
// to iterate so won't corrupt the loop.

// SendTransaction broadcasts the transaction to all currently active peers so
// it can be propagated to other nodes and eventually mined. An error won't be
// returned if the transaction already exists within the mempool. Any
// transaction broadcast through this method will be rebroadcast upon every
// change of the tip of the chain.
func (s *ChainService) SendTransaction(tx *wire.MsgTx) error {
	_ = "STUB: not implemented"
	// TODO(roasbeef): pipe through querying interface
	return nil
}

// NewPeerConfig returns the configuration for the given ServerPeer.
func NewPeerConfig(sp *ServerPeer) *peer.Config { _ = "STUB: not implemented"; return nil }

// outboundPeerConnected is invoked by the connection manager when a new
// outbound connection is established.  It initializes a new outbound server
// peer instance, associates it with the relevant state such as the connection
// request instance and the connection itself, and finally notifies the address
// manager of the attempt.
func (s *ChainService) outboundPeerConnected(c *connmgr.ConnReq, conn net.Conn) {
	_ = "STUB: not implemented"
	// In the event that we have to disconnect the peer, we'll choose the
	// appropriate method to do so based on whether the connection request
	// is for a persistent peer or not.
	return
}

// Since we're completely removing the request for this
// peer, we'll need to request a new one.

// If the peer is banned, then we'll disconnect them.

// If we're already connected to this peer, then we'll close out the new
// connection and keep the old.

// peerDoneHandler handles peer disconnects by notifying the server that it's
// done along with other performing other desirable cleanup.
func (s *ChainService) peerDoneHandler(sp *ServerPeer) { _ = "STUB: not implemented"; return }

// Only tell block manager we are gone if we ever told it we existed.

// UpdatePeerHeights updates the heights of all peers who have announced the
// latest connected main chain block, or a recognized orphan. These height
// updates allow us to dynamically refresh peer heights, ensuring sync peer
// selection has access to the latest block heights for each peer.
func (s *ChainService) UpdatePeerHeights(latestBlkHash *chainhash.Hash,
	latestHeight int32, updateSource *ServerPeer) {
	_ = "STUB: not implemented"
	return
}

// ChainParams returns a copy of the ChainService's chaincfg.Params.
func (s *ChainService) ChainParams() chaincfg.Params {
	_ = "STUB: not implemented"
	return *

	// Start begins connecting to peers and syncing the blockchain.
	new(chaincfg.Params)
}

func (s *ChainService) Start(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Already started?
	return nil
}

// Import headers if configured.
//nolint:lll

// The block manager was constructed before the import ran,
// so its internal header tracking state (headerList,
// headerTip, filterHeaderTip, etc.) reflects the
// pre-import chain tips. Re-read the now-updated stores
// so the block manager starts syncing from the correct
// height rather than from genesis.

// Start the address manager and block manager, both of which are
// needed by peers.

// Start the peer handler which in turn starts the address and block
// managers.

// Stop gracefully shuts down the server by stopping and disconnecting all
// peers and the main listener.
func (s *ChainService) Stop() error {
	_ = "STUB: not implemented"
	// Make sure this only happens once.
	return nil
}

// Signal the remaining goroutines to quit.

// IsCurrent lets the caller know whether the chain service's block manager
// thinks its view of the network is current.
func (s *ChainService) IsCurrent() bool { _ = "STUB: not implemented"; return false }

// PeerByAddr lets the caller look up a peer address in the service's peer
// table, if connected to that peer address.
func (s *ChainService) PeerByAddr(addr string) *ServerPeer { _ = "STUB: not implemented"; return nil }

// RescanChainSource is a wrapper type around the ChainService struct that will
// be used to satisfy the rescan.ChainSource interface.
type RescanChainSource struct {
	*ChainService
}

// A compile-time check to ensure that RescanChainSource implements the
// rescan.ChainSource interface.
var _ ChainSource = (*RescanChainSource)(nil)

// GetBlockHeaderByHeight returns the header of the block with the given height.
func (s *RescanChainSource) GetBlockHeaderByHeight(
	height uint32) (*wire.BlockHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetBlockHeader returns the header of the block with the given hash.
func (s *RescanChainSource) GetBlockHeader(
	hash *chainhash.Hash) (*wire.BlockHeader, uint32, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// GetFilterHeaderByHeight returns the filter header of the block with the given
// height.
func (s *RescanChainSource) GetFilterHeaderByHeight(
	height uint32) (*chainhash.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Subscribe returns a block subscription that delivers block notifications in
// order. The bestHeight parameter can be used to signal that a backlog of
// notifications should be delivered from this height. When providing a height
// of 0, a backlog will not be delivered.
func (s *RescanChainSource) Subscribe(
	bestHeight uint32) (*blockntfns.Subscription, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// cachedAddr is an empty struct used to satisfy the cache.Value interface.
type cachedAddr struct{}

// Size returns the size of cachedAddr, which is 1.
func (c *cachedAddr) Size() (uint64, error) {
	_ = "STUB: not implemented"

	// onionAddr implements the net.Addr interface and represents a tor address.
	// This code is identical to btcd's unexported onionAddr. It is used so that
	// neutrino can connect to v2 addresses without relying on the OnionCat
	// encoding. It also enables connecting to v3 addresses.
	return 0, nil
}

type onionAddr struct {
	addr string
}

// String returns the onion address.
func (o *onionAddr) String() string {
	_ = "STUB: not implemented"

	// Network returns "onion".
	return ""
}

func (o *onionAddr) Network() string {
	_ = "STUB: not implemented"

	// Ensure onionAddr implements the net.Addr interface.
	return ""
}

var _ net.Addr = (*onionAddr)(nil)
