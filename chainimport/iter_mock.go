package chainimport

import (
	"iter"

	"github.com/stretchr/testify/mock"
)

// mockHeaderIterator mocks a header iterator for testing header iteration
// logic.
type mockHeaderIterator struct {
	mock.Mock
}

// Iterator returns an iter.Seq2[Header, error] for the specified range.
func (m *mockHeaderIterator) Iterator(start,
	end uint32) iter.Seq2[Header, error] {
	_ = "STUB: not implemented"
	return nil
}

// BatchIterator returns an iterator that yields batches of headers.
func (m *mockHeaderIterator) BatchIterator(start, end,
	batchSize uint32) iter.Seq2[[]Header, error] {
	_ = "STUB: not implemented"
	return nil
}

// ReadBatch collects all headers from the given range into a slice.
func (m *mockHeaderIterator) ReadBatch(start, end,
	batchSize uint32) ([]Header, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetStartIndex returns the configured start index for this iterator.
func (m *mockHeaderIterator) GetStartIndex() uint32 { _ = "STUB: not implemented"; return 0 }

// GetEndIndex returns the configured end index for this iterator.
func (m *mockHeaderIterator) GetEndIndex() uint32 { _ = "STUB: not implemented"; return 0 }

// GetBatchSize returns the configured batch size for this iterator.
func (m *mockHeaderIterator) GetBatchSize() uint32 { _ = "STUB: not implemented"; return 0 }
