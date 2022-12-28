package namespace

import (
	"context"
	"net/http"

	"github.com/edobtc/cloudkit/config"
	"github.com/edobtc/cloudkit/namespace"
)

// DecorateNamespace inspects the incoming request header and sets a namespace into
// both the request header as middleware, but also sets up
func DecorateNamespace(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqNS := r.Header.Get(namespace.HTTPHeaderKey)
		if reqNS == "" {

			reqNS = config.Read().DefaultNamespace
			r.Header.Set(namespace.HTTPHeaderKey, reqNS)
		}

		// augment the request context to have a namespace
		ctx := context.WithValue(r.Context(), namespace.NamespaceContext("Namespace"), reqNS)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
