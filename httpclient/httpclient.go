package httpclient

import (
	"net"
	"net/http"
	"runtime"
	"time"
)

const (
	// DefaultTimeout sets a sensible timeout
	DefaultTimeout = 10 * time.Second

	// DefaultKeepAlive set a sensible keepalive value
	DefaultKeepAlive = 10 * time.Second

	// IdleTimeout sets a sensible IdleTimeout
	IdleTimeout = 90 * time.Second

	// MaxIdleConnections sets sensible max idle connetions
	MaxIdleConnections = 100
)

// New Returns a new http client with sensible
// defaults to the stdlib but which does not share global client
// configuration
func New() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   DefaultTimeout,
				KeepAlive: DefaultKeepAlive,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          MaxIdleConnections,
			IdleConnTimeout:       IdleTimeout,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		},
	}
}
