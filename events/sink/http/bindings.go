package http

import (
	"github.com/edobtc/cloudkit/events/sink/http/handlers"
	"github.com/edobtc/cloudkit/events/sink/http/handlers/broadcast"
	"github.com/edobtc/cloudkit/events/sink/http/handlers/event"
	"github.com/edobtc/cloudkit/events/sink/ws/handlers/echo"

	"github.com/gorilla/mux"
)

var (
	prefix = "/events"
)

func BindRoutes(r *mux.Router) {
	r.HandleFunc("/status", handlers.Status)

	// Event registration
	r.HandleFunc("/register", event.Create).Methods("POST", "PUT")
	r.HandleFunc("/register/{id}", event.Create).Methods("DELETE")
	r.HandleFunc("/register/{id}", event.Show).Methods("GET")

	// Listener
	r.HandleFunc("/broadcast", broadcast.HTTP).Methods("POST")

	r.HandleFunc("/ws/echo", echo.EchoCommand)
}

func BindRoutesWithPrefix(r *mux.Router) {
	pr := r.PathPrefix(prefix).Subrouter()
	BindRoutes(pr)
}
