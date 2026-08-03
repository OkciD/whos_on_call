package httpclient

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
)

func (c *HTTPClient) logRequest(req *http.Request) {
	reqCopy := req.Clone(context.Background())

	for _, headerName := range c.config.Logging.Request.HideHeaders {
		if reqCopy.Header.Get(headerName) != "" {
			reqCopy.Header.Set(headerName, "<hidden>")
		}
	}

	reqBytes, err := httputil.DumpRequestOut(reqCopy, true)
	if err != nil {
		c.logger.WithError(err).Error("failed to dump request")
		return
	}

	c.logger.WithField("request", string(reqBytes)).Debug("sending request")
}

func (c *HTTPClient) logResponse(resp *http.Response) {
	respBytes, err := httputil.DumpResponse(resp, true)
	if err != nil {
		c.logger.WithError(err).Error("failed to dump response")
		return
	}

	c.logger.WithField("response", string(respBytes)).Debug("got response")
}

func (c *HTTPClient) Do(req *http.Request) (*http.Response, error) {
	// todo: assign request id

	c.logRequest(req)

	resp, err := c.Client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("failed to do request: %w", err)
	}

	c.logResponse(resp)

	return resp, nil
}
