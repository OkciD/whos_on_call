package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/OkciD/whos_on_call/internal/models"
	"github.com/OkciD/whos_on_call/internal/models/api"
)

func (r *Repository) Create(ctx context.Context, newDevice *models.Device) (*models.Device, error) {
	reqDevice, err := api.FromDeviceAppModel(newDevice)
	if err != nil {
		return nil, fmt.Errorf("failed to convert device to api entity: %w", err)
	}

	request, err := r.httpClient.MakeRequest(http.MethodPost, "api/v1/device", reqDevice)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	var respDevice api.Device
	_, err = r.httpClient.Do(ctx, request, reqDevice)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}

	newDeviceOut, err := respDevice.ToAppModel()
	if err != nil {
		return nil, fmt.Errorf("error transforming response api model to app: %w", err)
	}

	return newDeviceOut, nil
}
