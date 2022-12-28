package event

import (
	"net/http"

	"github.com/edobtc/cloudkit/http/response"
)

// Delete will soft delete a registered event
func Status(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResponder(w)
	resp.Data = nil

	resp.Data = response.Status{
		Message: "ok",
	}

	resp.Success()
}
