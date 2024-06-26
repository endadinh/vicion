package vicionxlending

import (
	"context"
	"errors"
	"sync"
	"time"
)

// List of errors
var (
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicVicionXLendingAPI provides the vicionX RPC service that can be
// use publicly without security implications.
type PublicVicionXLendingAPI struct {
	t        *Lending
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicVicionXLendingAPI create a new RPC vicionX service.
func NewPublicVicionXLendingAPI(t *Lending) *PublicVicionXLendingAPI {
	api := &PublicVicionXLendingAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the Lending sub-protocol version.
func (api *PublicVicionXLendingAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}
