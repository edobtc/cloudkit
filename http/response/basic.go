package response

type ResponseData map[string]interface{}

// Basic represents a general response type, which can
// be empty, hold just a message, or a message and a list of message
// strings containing more information
type Basic struct {
	Message string       `json:"message,omitempty"`
	Data    ResponseData `json:"data,omitempty"`
}
