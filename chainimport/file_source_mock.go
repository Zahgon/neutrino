package chainimport

import (
	"github.com/stretchr/testify/mock"
)

// mockHeaderImportSource mocks a header import source for testing import source
// interactions.
type mockHeaderImportSource struct {
	mock.Mock
	uri string
}

// Open opens the mock header import source.
func (m *mockHeaderImportSource) Open() error { _ = "STUB: not implemented"; return nil }

// Close closes the mock header import source.
func (m *mockHeaderImportSource) Close() error { _ = "STUB: not implemented"; return nil }

// GetHeaderMetadata gets header metadata from the mock header import source.
func (m *mockHeaderImportSource) GetHeaderMetadata() (*headerMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetHeader gets a header by index from the mock header import source.
func (m *mockHeaderImportSource) GetHeader(index uint32) (Header, error) {
	_ = "STUB: not implemented"
	return *new(Header), nil
}

// Iterator returns a header iterator from the mock header import source.
func (m *mockHeaderImportSource) Iterator(start, end uint32,
	batchSize uint32) HeaderIterator {
	_ = "STUB: not implemented"
	return *new(HeaderIterator)
}

// GetURI gets the URI from the mock header import source.
func (m *mockHeaderImportSource) GetURI() string { _ = "STUB: not implemented"; return "" }

// SetURI sets the URI for the mock header import source.
func (m *mockHeaderImportSource) SetURI(uri string) { _ = "STUB: not implemented"; return }
