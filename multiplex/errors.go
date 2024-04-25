package multiplex

import "errors"

var (
	// ErrMarshallingMessage is returned when an error occurs marshalling a message to JSON
	ErrMarshallingMessage = errors.New("error marshalling message to JSON")

	// ErrSendingMessage is returned when an error occurs sending a message to the destination
	ErrSendingMessage = errors.New("error sending message to destination")
)
