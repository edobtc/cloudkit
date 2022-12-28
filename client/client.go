package client

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/edobtc/cloudkit/httpclient"
	"github.com/edobtc/cloudkit/settings"
)

// Client implements a api client
type Client struct {
	HTTP    *http.Client
	APIHost string
	Token   string
}

// NewClient returns a new client for access
// the ai-platform API client
func NewClient() *Client {
	return &Client{
		HTTP: httpclient.New(),
	}
}

// NewClientWithToken returns a new client for access
// the ai-platform API client and reads the JWT token from
// settings
func NewClientWithToken() (*Client, error) {
	s, err := settings.Read()

	if err != nil {
		return nil, err
	}

	return &Client{
		HTTP:    httpclient.New(),
		APIHost: s.Data.URL,
		Token:   s.Data.Token,
	}, nil
}

// apiURL builds a full url/path for making an API request
func (c *Client) apiURL(endpoint string) string {
	return fmt.Sprintf(
		"%s/%s",
		c.APIHost,
		endpoint,
	)
}

// Request builds a request to execute, this will also
// add the JWT token if one is set up in the client.
//
// Strong Argument: that this should return and error an never make a request
// IF there is no token set? I'm open
func (c *Client) Request(method string, url string, data io.Reader) (*http.Request, error) {
	u := c.apiURL(url)

	req, err := http.NewRequest(method, u, data)
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "application/json")

	if c.Token != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}

	return req, nil
}

// Do executes and parses an HTTP request
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.HTTP.Do(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ParseResponseBody takes an http response and
// turns the body to a *[]byte pointer
func (c *Client) ParseResponseBody(resp *http.Response) ([]byte, error) {
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
