package chainimport

import (
	"net/http"

	"github.com/stretchr/testify/mock"
)

// mockHTTPClient mocks an HTTP client for testing HTTP header import source
// interactions.
type mockHTTPClient struct {
	mock.Mock
}

// Get returns a response from the mock HTTP client.
func (m *mockHTTPClient) Get(url string) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
