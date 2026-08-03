package apiclient

import (
	"context"
	"fmt"

	"github.com/OkciD/whos_on_call/internal/shared/errors/mapper"
	"github.com/OkciD/whos_on_call/internal/shared/models"
	"github.com/OkciD/whos_on_call/internal/shared/models/api"
)

func (c *apiClient) CreateDevice(ctx context.Context, newDevice *models.Device) (*models.Device, error) {
	apiDevice, err := api.FromDeviceAppModel(newDevice)
	if err != nil {
		return nil, fmt.Errorf("failed to convert device from model to api: %w", err)
	}

	// todo: эта конвертация - дичь какая-то
	resp, err := c.genClient.CreateDeviceWithResponse(ctx, api.DeviceInput{
		Name: apiDevice.Name,
		Type: apiDevice.Type,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get create device in api: %w", err)
	}

	statusCode := resp.StatusCode()

	if statusCode == 201 {
		appDevice, err := resp.JSON201.ToAppModel()
		if err != nil {
			return nil, fmt.Errorf("failed to convert device from api to model: %w", err)
		}
		return appDevice, nil
	} else {
		return nil, mapper.RespToError(statusCode, *resp.JSONDefault)
	}
}
