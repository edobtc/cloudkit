package handlers

import (
	"net/http"

	"github.com/edobtc/cloudkit/http/response"
	log "github.com/sirupsen/logrus"
)

// Status handles health checking the http server
// and reporting if the service is operating, without
// validating anything else (db connections etc)
func Create(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResponder(w)
	resp.Data = nil

	log.Info("did it")
	log.Info("going ti")

	resp.Data = response.Status{
		Message: "ok",
	}

	resp.Success()
}
