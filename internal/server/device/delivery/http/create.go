package http

import (
	"context"
	"fmt"

	"github.com/OkciD/whos_on_call/cmd/server/apiserver/gen"
	appContext "github.com/OkciD/whos_on_call/internal/server/pkg/context"
	"github.com/OkciD/whos_on_call/internal/shared/models/api"
)

func (h *Handler) CreateDevice(ctx context.Context, request gen.CreateDeviceRequestObject) (gen.CreateDeviceResponseObject, error) {
	newDeviceInput := api.Device{
		Name: request.Body.Name,
		Type: request.Body.Type,
	}

	user, err := appContext.GetUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from request: %w", err)
	}

	newDeviceApp, err := newDeviceInput.ToAppModel()
	if err != nil {
		return nil, fmt.Errorf("failed to convert to app device model: %w", err)
	}
	newDeviceApp.User = user

	newDeviceApp, err = h.deviceUseCase.Create(ctx, newDeviceApp)
	if err != nil {
		return nil, fmt.Errorf("failed to create device: %w", err)
	}

	newDeviceOutput, err := api.FromDeviceAppModel(newDeviceApp)
	if err != nil {
		return nil, fmt.Errorf("failed to convert back to api device model: %w", err)
	}

	return gen.CreateDevice201JSONResponse(*newDeviceOutput), nil
}
