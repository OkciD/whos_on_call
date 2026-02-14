package api

import (
	"github.com/OkciD/whos_on_call/internal/client/device"
	"github.com/OkciD/whos_on_call/internal/client/pkg/httpclient"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
)

type Repository struct {
	logger logger.Logger

	httpClient *httpclient.HTTPClient
}

func New(logger logger.Logger, httpClient *httpclient.HTTPClient) device.Repository {
	return &Repository{
		logger: logger,

		httpClient: httpClient,
	}
}
