package apiclient

import (
	"context"
	"fmt"

	"github.com/OkciD/whos_on_call/internal/shared/errors/mapper"
	"github.com/OkciD/whos_on_call/internal/shared/models"
	"github.com/OkciD/whos_on_call/internal/shared/models/api"
)

func (c *apiClient) ListDevices(ctx context.Context, params *models.DeviceListParams) ([]models.Device, error) {
	apiListDeviceParams, err := api.FromDeviceListParamsAppModel(params)
	if err != nil {
		return nil, fmt.Errorf("failed to convert list devices params from model to api: %w", err)
	}

	resp, err := c.genClient.ListDevicesWithResponse(ctx, apiListDeviceParams)
	if err != nil {
		return nil, fmt.Errorf("failed to get devices list from api: %w", err)
	}

	statusCode := resp.StatusCode()

	if statusCode == 200 {
		appDevices := make([]models.Device, 0, len(*resp.JSON200))

		for _, apiDevice := range *resp.JSON200 {
			appDevice, err := apiDevice.ToAppModel()
			if err != nil {
				return nil, fmt.Errorf("failed to convert device from api to model: %w", err)
			}
			appDevices = append(appDevices, *appDevice)
		}

		return appDevices, nil
	} else {
		return nil, mapper.RespToError(statusCode, *resp.JSONDefault)
	}
}
