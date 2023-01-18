package handlers

import (
	"io"
	"net/http"

	"github.com/edobtc/cloudkit/controlplane/ws"
	"github.com/edobtc/cloudkit/http/response"
)

func PublishCommand(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResponder(w)
	resp.Data = map[string]string{
		"message": "received",
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		resp.UnexpectedWithError(err)
	}

	ws.Pool.Publish(body)
	resp.Success()
}
