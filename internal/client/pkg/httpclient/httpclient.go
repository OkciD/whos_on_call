package httpclient

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/pkg/logger"
)

func New(logger logger.Logger, config Config) http.Client {
	return http.Client{
		// todo: fn to apply roundtrippers sequentially
		Transport: loggingRoundTripper{
			logger: logger,
			next: authRoundTripper{
				apiKey: config.ApiKey,
				next:   http.DefaultTransport,
			},
		},
		Timeout: config.Timeout.Duration,
	}
}
