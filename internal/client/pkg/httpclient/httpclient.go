package httpclient

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/pkg/logger"
)

func New(logger logger.Logger, config Config) http.Client {
	return http.Client{
		Transport: loggingRoundTripper{
			logger: logger,
			next:   http.DefaultTransport,
		},
		Timeout: config.Timeout.Duration,
	}
}
