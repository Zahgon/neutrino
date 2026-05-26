package chainimport

import (
	"context"

	"github.com/btcsuite/btcd/wire"
	"github.com/lightninglabs/neutrino/headerfs"
)

// AddHeadersImportMetadata prepares a header file for import by adding the
// necessary metadata to the file. This function takes an existing header file
// and prepends metadata required by Neutrino's header import feature.
func AddHeadersImportMetadata(sourceFilePath string,
	networkMagic wire.BitcoinNet, version uint8,
	headerType headerfs.HeaderType, startHeight uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// setLastFilterHeaderHash updates the HeaderHash of the last filter header to
// match the block hash of the corresponding block header. This maintains chain
// tip consistency for the regular tip.
func setLastFilterHeaderHash(filterHeaders []headerfs.FilterHeader,
	chainTipBlockHeader headerfs.BlockHeader) {
	_ = "STUB: not implemented"

	// We only need to set the block header hash of the last filter header
	// to maintain chain tip consistency for regular tip.
	return
}

// targetHeightToImportSourceIndex converts the absolute blockchain target
// height to the equivalent import source height based on the start height
// input.
func targetHeightToImportSourceIndex(targetH, importStartH uint32) uint32 {
	_ = "STUB: not implemented"
	return 0
}

// ctxCancelled checks if context is cancelled and returns the error if so. It
// is used during long-running headers import operations to provide responsive
// cancellation.
func ctxCancelled(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// assertBlockHeader type asserts header to *blockHeader or returns error.
func assertBlockHeader(header Header) (*blockHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// assertFilterHeader type asserts header to *filterHeader or returns error.
func assertFilterHeader(header Header) (*filterHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
