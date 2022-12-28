package directors

import (
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func ConfiguredAPIForwarder(req *http.Request) {
	// Need to fetch a specific ID here
	vars := mux.Vars(req)

	log.Info(vars)

	// TODO:

	// using relevant request ID
	// fetch resources
	// with resources name
	// get URL
	// add Auth config (IP, HostHeader, API Key)
	// make request, return response

	// Temporary Thing
	rand.Seed(time.Now().UnixNano())

	endpoints := []string{
		"http://127.0.0.1:8090/headers",
		"http://127.0.0.1:8090/hello",
		"https://www.google.com/",
		"https://edobtc.com/",
	}

	index := rand.Intn(len(endpoints))
	u, err := url.Parse(endpoints[index])
	if err != nil {
		log.Error(err)
		return
	}

	// Test injecting the thing
	req.Header.Set("AuthKey", "2323435435434")
	req.URL.RawQuery = "?apiKey=23434343234234"

	req.URL.Scheme = u.Scheme
	req.URL.Host = u.Host
	req.URL.Path = u.Path
	req.Host = u.Host
}
