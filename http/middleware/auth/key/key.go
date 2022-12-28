package key

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/edobtc/cloudkit/config"
)

var (
	ErrApiKeyNotFound = errors.New("no API key present, please add a valid API key")
	ErrApiKeyInvalid  = errors.New("we found and API key but it was invalid")
)

// ApiKey looks up an API key, validates it's existence
// and associates related developer/application config
func ApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		value := r.Header.Get("X-EDO-ApiKey")

		if value == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("%s", ErrApiKeyNotFound)))
			return
		}

		if value != config.Read().APIKey {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("%s", ErrApiKeyInvalid)))
			return

		}
		next.ServeHTTP(w, r)
	})
}
