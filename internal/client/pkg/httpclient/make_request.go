package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *HTTPClient) MakeRequest(method string, path string, body any) (*http.Request, error) {
	fullUrl, err := c.baseURL.Parse(path)
	if err != nil {
		return nil, fmt.Errorf("failed to build request url: %w", err)
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, fmt.Errorf("error marshalling request body to json: %w", err)
		}
	}

	req, err := http.NewRequest(method, fullUrl.String(), buf)
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	req.Header.Set("X-Api-Key", c.apiKey)

	return req, nil
}
