package request

import (
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	blacklist = []string{
		"health",
		"status",
		"metrics",
	}
)

// LogRequest logs each request at the start of the http request
// so we can get webserver style logs when log level is set to INFO
func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var blacklistTerm bool

		for _, term := range blacklist {
			blacklistTerm = strings.Contains(r.RequestURI, term)

			if blacklistTerm {
				break
			}
		}

		if !blacklistTerm {
			log.Infof("START: %s: %s, %s", r.Method, r.RequestURI, time.Now())
		}

		next.ServeHTTP(w, r)
	})
}
