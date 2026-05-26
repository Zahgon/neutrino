package headerfs

import (
	"bytes"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

// ErrHeaderNotFound is returned when a target header on disk (flat file) can't
// be found.
type ErrHeaderNotFound struct {
	error
}

// appendRaw appends a new raw header to the end of the flat file.
func (h *headerStore) appendRaw(header []byte) error {
	_ = "STUB: not implemented"
	// Get current file position before writing. We'll use this position to
	// revert to if the write fails partially.
	return nil
}

// If we wrote some bytes but not all (partial write),
// truncate the file back to its original size to maintain
// consistency. This removes the partial/corrupt header.

// readRaw reads a raw header from disk from a particular seek distance. The
// amount of bytes read past the seek distance is determined by the specified
// header type.
func (h *headerStore) readRaw(seekDist uint64) ([]byte, error) {
	_ = "STUB: not implemented"
	// Based on the defined header type, we'll determine the number of bytes
	// that we need to read past the sync point.
	return nil, nil
}

// TODO(roasbeef): add buffer pool

// With the number of bytes to read determined, we'll create a slice
// for that number of bytes, and read directly from the file into the
// buffer.

// readHeaderRange will attempt to fetch a series of block headers within the
// target height range. This method batches a set of reads into a single system
// call thereby increasing performance when reading a set of contiguous
// headers.
//
// NOTE: The end height is _inclusive_ so we'll fetch all headers from the
// startHeight up to the end height, including the final header.
func (h *blockHeaderStore) readHeaderRange(startHeight uint32,
	endHeight uint32) ([]wire.BlockHeader, error) {
	_ = "STUB: not implemented"

	// Based on the defined header type, we'll determine the number of
	// bytes that we need to read from the file.
	return nil, nil
}

// We'll now incrementally parse out the set of individual headers from
// our set of serialized contiguous raw headers.

// readHeader reads a full block header from the flat-file. The header read is
// determined by the height value.
func (h *blockHeaderStore) readHeader(height uint32) (wire.BlockHeader, error) {
	_ = "STUB: not implemented"
	return *new(wire.BlockHeader), nil
}

// Each header is 80 bytes, so using this information, we'll seek a
// distance to cover that height based on the size of block headers.

// With the distance calculated, we'll raw a raw header start from that
// offset.

// Finally, decode the raw bytes into a proper bitcoin header.

// readHeader reads a single filter header at the specified height from the
// flat files on disk.
func (f *filterHeaderStore) readHeader(height uint32) (*chainhash.Hash, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// readHeaderRange will attempt to fetch a series of filter headers within the
// target height range. This method batches a set of reads into a single system
// call thereby increasing performance when reading a set of contiguous
// headers.
//
// NOTE: The end height is _inclusive_ so we'll fetch all headers from the
// startHeight up to the end height, including the final header.
func (f *filterHeaderStore) readHeaderRange(startHeight uint32,
	endHeight uint32) ([]chainhash.Hash, error) {
	_ = "STUB: not implemented"

	// Based on the defined header type, we'll determine the number of
	// bytes that we need to read from the file.
	return nil, nil
}

// We'll now incrementally parse out the set of individual headers from
// our set of serialized contiguous raw headers.

// readHeadersFromFile reads a chunk of headers, each of size headerSize, from
// the given file, from startHeight to endHeight.
func readHeadersFromFile(f File, headerSize, startHeight,
	endHeight uint32) (*bytes.Reader, error) {
	_ = "STUB: not implemented"

	// Each header is headerSize bytes, so using this information, we'll
	// seek a distance to cover that height based on the size the headers.
	return nil, nil
}

// Based on the number of headers in the range, we'll allocate a single
// slice that's able to hold the entire range of headers.

// Now that we have our slice allocated, we'll read out the entire
// range of headers with a single system call.
