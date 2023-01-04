package handlers

import (
	"net/http"

	"github.com/edobtc/cloudkit/controlplane/ws"
	"github.com/edobtc/cloudkit/http/response"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResponder(w)
	resp.Data = nil

	ws.Pool.Publish([]byte("ping"))

	resp.Data = response.Status{
		Message: "ok",
	}

	resp.Success()
}
