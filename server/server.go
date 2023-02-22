package server

import (
	"context"
	"net/http"
	"os"

	"github.com/edobtc/cloudkit/config"

	// General Handlers
	"github.com/edobtc/cloudkit/http/handlers/meta"

	// websocket handlers

	// Request formatters/handlers
	"github.com/edobtc/cloudkit/http/middleware/request"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Server is a structured wrapper for
// running the API server, encapsulating a
// route and error handling
type Server struct {
	http http.Server

	// add this to generic server to different
	// implementations of server units
	// cmd/**/* can bind subsets of routes
	// specific to their output binary
	r *mux.Router

	Close chan os.Signal
	Wait  chan bool
	Error chan error
}

// NewServer returns a new Server with
// defaults set, routes bootstrapped, and middleware
// set up
func NewServer() Server {
	r := mux.NewRouter()

	// Global middleware
	r.Use(request.LogRequest)

	// basic http status
	r.HandleFunc("/status", meta.Status)
	r.HandleFunc("/healthz", meta.Status) // assimilate to borg

	return Server{
		http: http.Server{
			Addr:    config.Read().Listen,
			Handler: r,
		},
		r:     r,
		Close: make(chan os.Signal),
		Wait:  make(chan bool),
		Error: make(chan error),
	}
}

func (s *Server) Router() *mux.Router {
	return s.r
}

// Start starts the http server
func (s *Server) Start() chan bool {
	s.r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		tpl, err := route.GetPathTemplate()
		if err != nil {
			log.Errorf("route: %s", err)
		}

		log.Info(tpl)
		return nil
	})

	go func() {
		log.Infof("server starting and listening on %s", s.http.Addr)
		err := s.http.ListenAndServe()
		if err != nil {
			if err != http.ErrServerClosed {
				log.Error(err)
				os.Exit(1)
			}
		}
	}()

	return s.Wait
}

func (s *Server) close() {
	close(s.Error)
	close(s.Wait)
	close(s.Close)
}

// GracefulShutdown safely shuts down the server
func (s *Server) GracefulShutdown() {
	if err := s.http.Shutdown(context.Background()); err != nil {
		log.Info("failed to shutdown properly")
		log.Error(err)
	}

	s.close()
}
