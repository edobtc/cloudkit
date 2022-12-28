package controlplane

import (
	"github.com/edobtc/cloudkit/controlplane/http/handlers"
	"github.com/edobtc/cloudkit/controlplane/ws/handlers/subscribe"

	"github.com/gorilla/mux"
)

var (
	prefix = "/control"
)

func BindRoutes(r *mux.Router) {
	r.HandleFunc("/create", handlers.Create)
	r.HandleFunc("/ws/subscribe", subscribe.SubscribeCommand)
}

func BindRoutesWithPrefix(r *mux.Router) {
	pr := r.PathPrefix(prefix).Subrouter()
	BindRoutes(pr)
}
