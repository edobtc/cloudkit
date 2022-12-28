package http

import (
	"github.com/edobtc/cloudkit/relay/http/handlers"
	"github.com/edobtc/cloudkit/relay/ws/handlers/subscribe"

	"github.com/gorilla/mux"
)

var (
	prefix = "/relay"
)

func BindRoutes(r *mux.Router) {
	r.HandleFunc("/status", handlers.Status)
	r.HandleFunc("/broadcast", subscribe.SubscribeCommand)
}

func BindRoutesWithPrefix(r *mux.Router) {
	pr := r.PathPrefix(prefix).Subrouter()
	BindRoutes(pr)
}
