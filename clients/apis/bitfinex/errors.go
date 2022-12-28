package bitfinex

import (
	"errors"
)

var (
	// ErrRequestFailed is returned in the instance a
	// http request does not return a 200
	ErrRequestFailed = errors.New("failed request not successful")

	// ErrNotFound is returned in the instance a
	// http request returns a 404
	ErrNotFound = errors.New("resource not found")

	// ErrCacheEntryExpired is returned if
	// the cached entry that is returned has expired.
	// This SHOULD be taken care of by key expiry, but there is a chance
	// a stale result is returned and by the time of presentation
	// has exceeded the expiry timeline, therefore should be discarded
	ErrCacheEntryExpired = errors.New("cache entry expired, please refresh")
)

// Error is the structure of a API
// error response
type Error struct {
	Message string

	Info []string
}
