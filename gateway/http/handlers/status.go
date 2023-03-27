package handlers

import (
	"net/http"

	"github.com/edobtc/cloudkit/http/response"
)

// Status handles health checking the http server
// and reporting if the service is operating, without
// validating anything else (db connections etc)
func Status(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResponder(w)
	resp.Data = nil

	resp.Data = response.Status{
		Message: "gateway status: ok",
	}

	resp.Success()
}
