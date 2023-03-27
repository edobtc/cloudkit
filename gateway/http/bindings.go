package http

import (
	"github.com/edobtc/cloudkit/relay/http/handlers"

	"github.com/gorilla/mux"
)

var (
	// gateway prefix, simpler than "gateway"
	prefix = "/gw"
)

func BindRoutes(r *mux.Router) {
	r.HandleFunc("/status", handlers.Status)
}

// BindRoutesWithPrefix binds routes with an optionally
// defined prefix, this is useful for adding routes to a
// monolithic server where we might be adding multiple bindings
// and want to avoid collisions
//
// eg, if both:
//
// package-a/handler.go /cool/thing
// package-b/handler.go /cool/thing
//
// are defined, then we want
//
// /a/cool/thing
// /b/cool/thing
func BindRoutesWithPrefix(r *mux.Router) {
	pr := r.PathPrefix(prefix).Subrouter()
	BindRoutes(pr)
}
