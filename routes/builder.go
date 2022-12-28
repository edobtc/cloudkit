package routes

import "net/http"

type RouteConfig struct {
	Rule    string
	Handler func(w http.ResponseWriter, r *http.Request)
}
