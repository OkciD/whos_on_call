package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *HTTPClient) Do(ctx context.Context, req *http.Request, responseBody any) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		return nil, fmt.Errorf("failed to do request: %w", err)
	}
	defer resp.Body.Close()

	if c := resp.StatusCode; c < 200 || c > 299 {
		return nil,
			fmt.Errorf("%s %s returned invalid status code %d", resp.Request.Method, resp.Request.URL, resp.StatusCode)
	}

	respBodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body bytes: %w", err)
	}

	err = json.Unmarshal(respBodyBytes, responseBody)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body bytes: %w", err)
	}

	return resp, nil
}
