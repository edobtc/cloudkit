package proxy

import (
	"net/http"
	"net/http/httputil"

	"github.com/edobtc/cloudkit/proxy/directors"
	"github.com/edobtc/cloudkit/proxy/modifiers"

	log "github.com/sirupsen/logrus"
)

// ResponseHandler is the pluggable func format for
// response handlers from the downstream service
type ResponseHandler func(*http.Response) error

// ErrorHandler defines a type of pluggable error handlers for
// responding to downstream requests that have failed
type ErrorHandler func(http.ResponseWriter, *http.Request, error)

// NewErrHandler returns a handler for processing downstream
// errors of the filehandler service
func NewErrHandler() func(w http.ResponseWriter, r *http.Request, err error) {
	return func(w http.ResponseWriter, r *http.Request, err error) {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

// NewAPIRequestForwarder returns a ReverseProxy that will ensure a request
// is handled by the downstream API service BEFORE any subsequent
// handling we want to do is done
func NewAPIRequestForwarder(
	resp ResponseHandler,
	onError ErrorHandler) (*httputil.ReverseProxy, error) {

	p := httputil.ReverseProxy{
		// Director handles the incoming request and
		// makes modifications. in it's simplest form
		// this will add configured Auth settings to the
		// outgoing request as configured for the app resource
		// that is being loaded
		Director: directors.ConfiguredAPIForwarder,

		// ErrorHandler customizes error handling
		// as requests go through the proxy which can
		// be oblique to observe otherwise
		ErrorHandler: onError,

		// Optionally define a response handler
		// that can alter the responses
		ModifyResponse: modifiers.ExampleProxyResponseHandler,
	}

	return &p, nil
}
