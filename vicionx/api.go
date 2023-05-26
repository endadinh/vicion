package vicionx

import (
	"context"
	"errors"
	"sync"
	"time"
)

const (
	LimitThresholdOrderNonceInQueue = 100
)

// List of errors
var (
	ErrNoTopics          = errors.New("missing topic(s)")
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicVicionXAPI provides the vicionX RPC service that can be
// use publicly without security implications.
type PublicVicionXAPI struct {
	t        *VicionX
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicVicionXAPI create a new RPC vicionX service.
func NewPublicVicionXAPI(t *VicionX) *PublicVicionXAPI {
	api := &PublicVicionXAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the VicionX sub-protocol version.
func (api *PublicVicionXAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}
