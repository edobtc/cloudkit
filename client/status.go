package client

// Status fetches and returns the raw json
// status payload from the API server. The stauts
// is a simple healthcheck primarily used for healthcheck
// and availability of the API, but tells us nothing about
// working functionality or connections to dbs
func (c *Client) Status() ([]byte, error) {
	req, err := c.Request("GET", "/status", nil)

	if err != nil {
		return []byte{}, err
	}

	resp, err := c.Do(req)

	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	data, err := c.ParseResponseBody(resp)

	if err != nil {
		if err == ErrNotFound {
			return []byte{}, nil
		}

		return []byte{}, err
	}

	return data, nil
}
