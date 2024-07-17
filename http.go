package bigcartel

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// doRequest is a helper method to handle HTTP requests and responses
func (c *Client) doRequest(method, endpoint string, body interface{}) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, c.BaseURL+endpoint, reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.api+json")
	req.Header.Set("User-Agent", c.UserAgent1)
	req.Header.Add("User-Agent", c.UserAgent2)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.SetBasicAuth(c.Username, c.Password)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("failed request: %s", resp.Status)
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}
