package event

import (
	"net/http"

	"github.com/edobtc/cloudkit/http/response"
)

// Show will show the registered event
func Show(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResponder(w)
	resp.Data = nil

	resp.Data = response.Status{
		Message: "ok",
	}

	resp.Success()
}
