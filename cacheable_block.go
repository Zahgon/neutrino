package neutrino

import "github.com/btcsuite/btcd/btcutil"

// CacheableBlock is a wrapper around the btcutil.Block type which provides a
// Size method used by the cache to target certain memory usage.
type CacheableBlock struct {
	*btcutil.Block
}

// Size returns size of this block in bytes.
func (c *CacheableBlock) Size() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }
