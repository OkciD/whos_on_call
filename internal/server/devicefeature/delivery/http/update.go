package http

import (
	"context"
	"fmt"

	"github.com/OkciD/whos_on_call/internal/server/pkg/apiserver/gen"
	appContext "github.com/OkciD/whos_on_call/internal/server/pkg/context"
	"github.com/OkciD/whos_on_call/internal/shared/models/api"
)

// PUT /api/v1/device/{deviceid}/feature

func (h DeviceFeatureHandler) UpdateDeviceFeature(
	ctx context.Context,
	request gen.UpdateDeviceFeatureRequestObject,
) (gen.UpdateDeviceFeatureResponseObject, error) {
	deviceID := int(request.Deviceid)

	newDeviceFeatureInput := api.DeviceFeature{
		Status: request.Body.Status,
		Type:   request.Body.Type,
	}

	newDeviceFeatureApp, err := newDeviceFeatureInput.ToAppModel()
	if err != nil {
		return nil, fmt.Errorf("failed to convert to app device feature model: %w", err)
	}

	user, err := appContext.GetUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from request: %w", err)
	}

	newDeviceFeatureApp, err = h.deviceFeatureUseCase.Update(ctx, deviceID, user, newDeviceFeatureApp)
	if err != nil {
		return nil, fmt.Errorf("failed to upsert device feature: %w", err)
	}

	newDeviceFeatureAPI, err := api.FromDeviceFeatureAppModel(newDeviceFeatureApp)
	if err != nil {
		return nil, fmt.Errorf("error converting device feature to api model: %w", err)
	}

	return gen.UpdateDeviceFeature200JSONResponse(*newDeviceFeatureAPI), nil
}
