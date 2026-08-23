package http

import (
	"context"
	"fmt"

	"github.com/OkciD/whos_on_call/internal/server/pkg/apiserver/gen"
	appContext "github.com/OkciD/whos_on_call/internal/server/pkg/context"
	"github.com/OkciD/whos_on_call/internal/shared/models/api"
)

// GET /api/v1/devices?type=...&name=...

func (h DeviceHandler) ListDevices(
	ctx context.Context,
	request gen.ListDevicesRequestObject,
) (gen.ListDevicesResponseObject, error) {
	user, err := appContext.GetUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from request: %w", err)
	}

	appParams, err := request.Params.ToAppModel()
	if err != nil {
		return nil, fmt.Errorf("error converting request params: %w", err)
	}

	appDevices, err := h.deviceUseCase.List(ctx, user, appParams)
	if err != nil {
		return nil, fmt.Errorf("error listing devices: %w", err)
	}

	apiDevices := make([]api.Device, 0, len(appDevices))
	for _, appDevice := range appDevices {
		apiDevice, err := api.FromDeviceAppModel(&appDevice)
		if err != nil {
			return nil, fmt.Errorf("error converting app device model to api models: %w", err)
		}

		apiDevices = append(apiDevices, *apiDevice)
	}

	return gen.ListDevices200JSONResponse(apiDevices), nil
}
