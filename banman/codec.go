package banman

import (
	"io"
	"net"
)

// ipType represents the different types of IP addresses supported by the
// BanStore interface.
type ipType = byte

const (
	// ipv4 represents an IP address of type IPv4.
	ipv4 ipType = 0

	// ipv6 represents an IP address of type IPv6.
	ipv6 ipType = 1
)

// encodeIPNet serializes the IP network into the given reader.
func encodeIPNet(w io.Writer, ipNet *net.IPNet) error {
	_ = "STUB: not implemented"
	// Determine the appropriate IP type for the IP address contained in the
	// network.
	return nil
}

// Write the IP type first in order to properly identify it when
// deserializing it, followed by the IP itself and its mask.

// decodeIPNet deserialized an IP network from the given reader.
func decodeIPNet(r io.Reader) (*net.IPNet, error) {
	_ = "STUB: not implemented"
	// Read the IP address type and determine whether it is supported.
	return nil, nil
}

// Once we have the type and its corresponding length, attempt to read
// it and its mask.
