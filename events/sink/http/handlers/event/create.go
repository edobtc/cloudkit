package event

import (
	"net/http"

	"github.com/edobtc/cloudkit/http/response"
)

// Create registers an event in the event registry so
// payment events can be broadcast to it
func Create(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResponder(w)
	resp.Data = nil

	resp.Data = response.Status{
		Message: "ok",
	}

	resp.Success()
}
