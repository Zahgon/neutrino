package query

const (
	// bestScore is the best score a peer can get after multiple rewards.
	bestScore = 0

	// defaultScore is the score given to a peer when it hasn't been
	// rewarded or punished.
	defaultScore = 4

	// worstScore is the worst score a peer can get after multiple
	// punishments.
	worstScore = 8
)

// peerRanking is a struct that keeps history of peer's previous query success
// rate, and uses that to prioritise which peers to give the next queries to.
type peerRanking struct {
	// rank keeps track of the current set of peers and their score. A
	// lower score is better.
	rank map[string]uint64
}

// A compile time check to ensure peerRanking satisfies the PeerRanking
// interface.
var _ PeerRanking = (*peerRanking)(nil)

// NewPeerRanking returns a new, empty ranking.
func NewPeerRanking() PeerRanking { _ = "STUB: not implemented"; return *new(PeerRanking) }

// Order sorts the given slice of peers based on their current score. If a
// peer has no current score given, the default will be used.
func (p *peerRanking) Order(peers []string) { _ = "STUB: not implemented"; return }

// AddPeer adds a new peer to the ranking, starting out with the default score.
func (p *peerRanking) AddPeer(peer string) { _ = "STUB: not implemented"; return }

// Punish increases the score of the given peer.
func (p *peerRanking) Punish(peer string) { _ = "STUB: not implemented"; return }

// Cannot punish more.

// Reward decreases the score of the given peer.
// TODO(halseth): use actual response time when ranking peers.
func (p *peerRanking) Reward(peer string) { _ = "STUB: not implemented"; return }

// Cannot reward more.

// ResetRanking sets the score of the passed peer to the defaultScore.
func (p *peerRanking) ResetRanking(peer string) { _ = "STUB: not implemented"; return }
