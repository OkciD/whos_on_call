package apiclient

import (
	"context"
	"fmt"

	"github.com/OkciD/whos_on_call/internal/client/apiclient/gen"
	"github.com/OkciD/whos_on_call/internal/client/pkg/httpclient"
	"github.com/OkciD/whos_on_call/internal/shared/models"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type ApiClient interface {
	GetUser(ctx context.Context) (*models.User, error)
	CreateDevice(ctx context.Context, newDevice *models.Device) (*models.Device, error)
}

type apiClient struct {
	genClient *gen.ClientWithResponses
}

func NewWithHttpClient(logger logger.Logger, hc httpclient.HttpDoer, cfg Config) (ApiClient, error) {
	genClient, err := gen.NewClientWithResponses(
		cfg.BaseURL,
		gen.WithHTTPClient(hc),
		gen.WithRequestEditorFn(newAuthRequestEditor(cfg.ApiKey)),
	)
	if err != nil {
		return nil, fmt.Errorf("error initing generated api client: %w", err)
	}

	return &apiClient{
		genClient: genClient,
	}, nil
}

func New(logger logger.Logger, cfg Config) (ApiClient, error) {
	hc := httpclient.New(logger, cfg.HTTPClientConfig)

	return NewWithHttpClient(logger, hc, cfg)
}
