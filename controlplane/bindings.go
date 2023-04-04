package controlplane

import (
	"github.com/edobtc/cloudkit/controlplane/http/handlers"
	"github.com/edobtc/cloudkit/controlplane/ws/handlers/subscribe"

	"github.com/gorilla/mux"
)

var (
	prefix = "/cp"
)

func BindRoutes(r *mux.Router) {
	r.HandleFunc("/ping", handlers.Ping)
	r.HandleFunc("/ws/subscribe", subscribe.SubscribeCommand)
	r.HandleFunc("/v1/events/ws/publish", handlers.PublishCommand).Methods("POST")
}

func BindRoutesWithPrefix(r *mux.Router) {
	pr := r.PathPrefix(prefix).Subrouter()
	BindRoutes(pr)
}
