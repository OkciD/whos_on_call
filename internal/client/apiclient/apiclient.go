package apiclient

import (
	"context"
	"fmt"

	"github.com/OkciD/whos_on_call/internal/client/apiclient/gen"
	"github.com/OkciD/whos_on_call/internal/client/pkg/httpclient"
	"github.com/OkciD/whos_on_call/internal/shared/models"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type APIClient interface {
	GetUser(ctx context.Context) (*models.User, error)
	CreateDevice(ctx context.Context, newDevice *models.Device) (*models.Device, error)
	ListDevices(ctx context.Context, params *models.DeviceListParams) ([]models.Device, error)
	UpdateDeviceFeature(ctx context.Context, newDeviceFeature *models.DeviceFeature) (*models.DeviceFeature, error)
}

type apiClient struct {
	logger    logger.Logger
	genClient *gen.ClientWithResponses
}

func NewWithHTTPClient(logger logger.Logger, hc httpclient.HTTPDoer, cfg Config) (APIClient, error) {
	genClient, err := gen.NewClientWithResponses(
		cfg.BaseURL,
		gen.WithHTTPClient(hc),
		gen.WithRequestEditorFn(newAuthRequestEditor(cfg.APIKey)),
	)
	if err != nil {
		return nil, fmt.Errorf("error initing generated api client: %w", err)
	}

	return &apiClient{
		logger:    logger,
		genClient: genClient,
	}, nil
}

func New(logger logger.Logger, cfg Config) (APIClient, error) {
	hc := httpclient.New(logger, cfg.HTTPClientConfig)

	return NewWithHTTPClient(logger, hc, cfg)
}
