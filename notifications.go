// Copyright (c) 2013-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

// NOTE: THIS API IS UNSTABLE RIGHT NOW.

package neutrino

import (
	"github.com/lightninglabs/neutrino/query"
)

type getConnCountMsg struct {
	reply chan int32
}

type subConnPeersReply struct {
	peerChan   chan query.Peer
	cancelChan chan struct{}
}
type subConnPeersMsg struct {
	reply chan subConnPeersReply
}

type getPeersMsg struct {
	reply chan []*ServerPeer
}

type getOutboundGroup struct {
	key   string
	reply chan int
}

type getAddedNodesMsg struct {
	reply chan []*ServerPeer
}

type disconnectNodeMsg struct {
	cmp   func(*ServerPeer) bool
	reply chan error
}

type connectNodeMsg struct {
	addr      string
	permanent bool
	reply     chan error
}

type removeNodeMsg struct {
	cmp   func(*ServerPeer) bool
	reply chan error
}

type forAllPeersMsg struct {
	closure func(*ServerPeer)
}

// TODO: General - abstract out more of blockmanager into queries. It'll make
// this way more maintainable and usable.

// handleQuery is the central handler for all queries and commands from other
// goroutines related to peer state.
func (s *ChainService) handleQuery(state *peerState, querymsg interface{}) {
	_ = "STUB: not implemented"
	return
}

// Subscription for connected peers requested.

// Create a channel and fill it with the current set of
// connected peers.

// TODO: duplicate oneshots?
// Limit max number of total peers.

// TODO: if too many, nuke a non-perm peer.

// Keep group counts ok since we remove from
// the list now.

// Request a list of the persistent (added) peers.

// Respond with a slice of the relevant peers.

// Check outbound peers.

// Keep group counts ok since we remove from
// the list now.

// If there are multiple outbound connections to the same
// ip:port, continue disconnecting them all until no such
// peers are found.

// TODO: Remove this when it's unnecessary due to wider use of
// queryPeers.
// Run the closure on all peers in the passed state.

// Even though this is a query, there's no reply channel as the
// forAllPeers method doesn't return anything. An error might be
// useful in the future.

// ConnectedCount returns the number of currently connected peers.
func (s *ChainService) ConnectedCount() int32 { _ = "STUB: not implemented"; return 0 }

// ConnectedPeers is a function that returns a channel where all connected
// peers will be sent. It is assumed that all current peers will be sent
// imemdiately, and new peers as they connect.
func (s *ChainService) ConnectedPeers() (<-chan query.Peer, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// OutboundGroupCount returns the number of peers connected to the given
// outbound group key.
func (s *ChainService) OutboundGroupCount(key string) int { _ = "STUB: not implemented"; return 0 }

// AddedNodeInfo returns an array of btcjson.GetAddedNodeInfoResult structures
// describing the persistent (added) nodes.
func (s *ChainService) AddedNodeInfo() []*ServerPeer { _ = "STUB: not implemented"; return nil }

// Peers returns an array of all connected peers.
func (s *ChainService) Peers() []*ServerPeer { _ = "STUB: not implemented"; return nil }

// DisconnectNodeByAddr disconnects a peer by target address. Both outbound and
// inbound nodes will be searched for the target node. An error message will
// be returned if the peer was not found.
func (s *ChainService) DisconnectNodeByAddr(addr string) error {
	_ = "STUB: not implemented"
	return nil
}

// DisconnectNodeByID disconnects a peer by target node id. Both outbound and
// inbound nodes will be searched for the target node. An error message will be
// returned if the peer was not found.
func (s *ChainService) DisconnectNodeByID(id int32) error { _ = "STUB: not implemented"; return nil }

// RemoveNodeByAddr removes a peer from the list of persistent peers if
// present. An error will be returned if the peer was not found.
func (s *ChainService) RemoveNodeByAddr(addr string) error { _ = "STUB: not implemented"; return nil }

// RemoveNodeByID removes a peer by node ID from the list of persistent peers
// if present. An error will be returned if the peer was not found.
func (s *ChainService) RemoveNodeByID(id int32) error { _ = "STUB: not implemented"; return nil }

// ConnectNode adds `addr' as a new outbound peer. If permanent is true then the
// peer will be persistent and reconnect if the connection is lost.
// It is an error to call this with an already existing peer.
func (s *ChainService) ConnectNode(addr string, permanent bool) error {
	_ = "STUB: not implemented"
	return nil
}

// ForAllPeers runs a closure over all peers (outbound and persistent) to which
// the ChainService is connected. Nothing is returned because the peerState's
// ForAllPeers method doesn't return anything as the closure passed to it
// doesn't return anything.
func (s *ChainService) ForAllPeers(closure func(sp *ServerPeer)) { _ = "STUB: not implemented"; return }
