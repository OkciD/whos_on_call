package api

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/client/device"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
)

type Repository struct {
	logger logger.Logger

	httpClient *http.Client
}

func New(logger logger.Logger, httpClient *http.Client) device.Repository {
	return &Repository{
		logger: logger,

		httpClient: httpClient,
	}
}
