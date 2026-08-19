package apiclient

import (
	"context"
	"fmt"

	"github.com/OkciD/whos_on_call/internal/shared/errors/mapper"
	"github.com/OkciD/whos_on_call/internal/shared/models"
	"github.com/OkciD/whos_on_call/internal/shared/models/api"
)

func (c *apiClient) UpdateDeviceFeature(ctx context.Context, newDeviceFeature *models.DeviceFeature) (*models.DeviceFeature, error) {
	apiDeviceFeature, err := api.FromDeviceFeatureAppModel(newDeviceFeature)
	if err != nil {
		return nil, fmt.Errorf("failed to convert device from model to api: %w", err)
	}

	apiDevice, err := api.FromDeviceAppModel(newDeviceFeature.Device)
	if err != nil {
		return nil, fmt.Errorf("failed to convert device from model to api: %w", err)
	}

	// todo: эта конвертация - дичь какая-то
	resp, err := c.genClient.UpdateDeviceFeatureWithResponse(ctx, apiDevice.Id, api.DeviceFeatureInput{
		Status: apiDeviceFeature.Status,
		Type:   apiDeviceFeature.Type,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to update device feature via api: %w", err)
	}

	statusCode := resp.StatusCode()

	if statusCode == 200 {
		appDeviceFeature, err := resp.JSON200.ToAppModel()
		if err != nil {
			return nil, fmt.Errorf("failed to convert device from api to model: %w", err)
		}
		return appDeviceFeature, nil
	} else {
		return nil, mapper.RespToError(statusCode, *resp.JSONDefault)
	}
}
