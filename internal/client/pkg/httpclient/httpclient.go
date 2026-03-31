package httpclient

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type HTTPClient struct {
	client *http.Client

	baseURL *url.URL
	apiKey  string
}

func New(logger logger.Logger, config Config) (*HTTPClient, error) {
	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing baseURL: %w", err)
	}

	return &HTTPClient{
		client: &http.Client{
			Timeout: config.Timeout.Duration,
		},
		baseURL: baseURL,
		apiKey:  config.ApiKey,
	}, nil
}
