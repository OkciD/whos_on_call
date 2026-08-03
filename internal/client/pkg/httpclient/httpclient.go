package httpclient

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type HttpDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type HTTPClient struct {
	http.Client
	config Config
	logger logger.Logger
}

func New(logger logger.Logger, config Config) HttpDoer {
	return &HTTPClient{
		Client: http.Client{
			Timeout: config.Timeout.Duration,
		},
		config: config,
		logger: logger,
	}
}
