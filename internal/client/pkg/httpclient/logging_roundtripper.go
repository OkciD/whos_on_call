package httpclient

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/pkg/logger"
)

type loggingRoundTripper struct {
	next   http.RoundTripper
	logger logger.Logger
}

func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	l.logger.WithFields(logger.Fields{
		// todo: request id
		// todo: body if any
		"method": r.Method,
		"url":    r.URL.String(),
	}).Info("making request")

	resp, err := l.next.RoundTrip(r)

	l.logger.WithFields(logger.Fields{
		"statusCode": resp.StatusCode,
	}).Info("got response")

	return resp, err
}
