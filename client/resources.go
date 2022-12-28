package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/edobtc/cloudkit/plan"
)

func (c *Client) CreateResource(p plan.Definition) (plan.Definition, error) {
	data, err := json.Marshal(p)

	if err != nil {
		return p, err
	}

	req, err := c.Request("POST", ResourcesCreateEndpoint, bytes.NewBuffer(data))

	if err != nil {
		return p, err
	}

	resp, err := c.Do(req)

	if err != nil {
		return p, err
	}

	if resp.StatusCode != http.StatusOK {
		return p, ErrRequestFailed
	}

	defer resp.Body.Close()

	return p, nil
}
