package bitfinex

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/edobtc/cloudkit/httpclient"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

const (
	APIHost = "https://api.bitfinex.com"
	Version = "v1"

	//
	DefaultRateLimitEnabled = true

	DefaultCacheEnabled = false

	// RateLimit interval
	AllowedInterval = time.Second * 10
)

type cached struct {
	LastFetched time.Time
	Data        interface{}
}

type BitfinexClient struct {
	http    *http.Client
	Host    string
	Version string

	CacheEnabled bool
	RateLimit    bool
	cacheConn    *redis.Client

	lastRequest map[string]cached
}

// The simpler New returns the NewV1Client
func New() *BitfinexClient {
	return NewV1Client()
}

func NewV1Client() *BitfinexClient {
	return &BitfinexClient{
		http:         httpclient.New(),
		Host:         APIHost,
		Version:      Version,
		RateLimit:    true,
		CacheEnabled: false,

		lastRequest: map[string]cached{},
	}
}

func (b *BitfinexClient) Ticker(t Ticker) (*TickerData, error) {
	ticker := &TickerData{}

	tickerURL := fmt.Sprintf("pubticker/%s", t)

	if cached, ok := b.lastRequest[tickerURL]; ok {
		if time.Since(cached.LastFetched) < AllowedInterval {
			return cached.Data.(*TickerData), nil
		}
	}

	if b.CacheEnabled {
		log.Debug("reading from cache")
		data, err := b.FetchCached()
		if err != nil {
			return nil, err
		}

		if data != nil {
			return data, nil
		}
	}

	req, err := b.Request("GET", tickerURL, nil)
	if err != nil {
		return ticker, err
	}

	resp, err := b.Do(req)
	if err != nil {
		return ticker, err
	}

	data, err := b.ParseResponseBody(resp)
	if err != nil {
		return ticker, err
	}

	err = json.Unmarshal(data, ticker)
	if err != nil {
		return ticker, err
	}

	b.lastRequest[tickerURL] = cached{
		LastFetched: time.Now(),
		Data:        ticker,
	}

	if b.CacheEnabled {
		log.Debug("writing to cache")
		err = b.WriteCached(ticker)
		if err != nil {
			log.Error(err)
		}
	}

	return ticker, nil
}

// apiURL builds a full url/path for making an API request
func (b *BitfinexClient) apiURL(endpoint string) string {
	return fmt.Sprintf(
		"%s/%s/%s",
		b.Host,
		b.Version,
		endpoint,
	)
}

func (b *BitfinexClient) Request(method string, url string, data io.Reader) (*http.Request, error) {
	log.Debug("making request")
	u := b.apiURL(url)

	req, err := http.NewRequest(method, u, nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Do executes and parses an HTTP request
func (b *BitfinexClient) Do(req *http.Request) (*http.Response, error) {
	resp, err := b.http.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ParseResponseBody takes an http response and
// turns the body to a []byte
func (b *BitfinexClient) ParseResponseBody(resp *http.Response) ([]byte, error) {
	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, ErrNotFound
		}

		return nil, ErrRequestFailed
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
