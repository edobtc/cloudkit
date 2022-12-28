package client

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
)

// Error is the structure of a API
// error response
type Error struct {
	Message string

	Info []string
}
