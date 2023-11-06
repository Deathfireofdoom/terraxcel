package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ReadExtensions() ([]string, error) {
	if c == nil {
		return nil, fmt.Errorf("client is nil")
	}

	req, err := c.NewRequest(http.MethodGet, "/extension", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		var extensions []string
		err := json.NewDecoder(resp.Body).Decode(&extensions)
		if err != nil {
			return nil, fmt.Errorf("error decoding response: %w", err)
		}
		return extensions, nil
	}
	return nil, fmt.Errorf("got non-200 status code: %d", resp.StatusCode)

}
