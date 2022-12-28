package probe

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/edobtc/cloudkit/httpclient"
)

var (
	// ErrURLNotDefined is returned if the probe config does not
	// have a url set in the Target field of Config
	ErrURLNotDefined = errors.New("URL not set in Config.Target")
)

// HTTPProbe is the HTTP probe
type HTTPProbe struct {
	URL string
	c   *http.Client
}

// NewHTTPProbe instantiates a HTTP probe with
// the default configured http client
func NewHTTPProbe() *HTTPProbe {
	return &HTTPProbe{
		c: httpclient.New(),
	}
}

// Check performs a status check of the probe
func (h *HTTPProbe) Check() (bool, error) {
	if h.URL == "" {
		return false, ErrURLNotDefined
	}

	req, err := http.NewRequest("GET", h.URL, nil)
	if err != nil {
		return false, err
	}

	resp, err := h.c.Do(req)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Sprintf("status not ok %v", resp.StatusCode)
		return false, errors.New(err)

	}

	return true, nil
}
