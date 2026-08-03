package httpclient

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type HTTPClient struct {
	http.Client
	config Config
	logger logger.Logger
}

func New(logger logger.Logger, config Config) HTTPClient {
	return HTTPClient{
		Client: http.Client{
			Timeout: config.Timeout.Duration,
		},
		config: config,
		logger: logger,
	}
}
