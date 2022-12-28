package broadcast

import (
	"net/http"

	"github.com/edobtc/cloudkit/http/response"
)

func HTTP(w http.ResponseWriter, r *http.Request) {
	resp := response.NewResponder(w)
	resp.Data = nil

	resp.Data = response.Status{
		Message: "ok",
	}

	resp.Success()
}
