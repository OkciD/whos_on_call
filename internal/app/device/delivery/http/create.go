package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/OkciD/whos_on_call/internal/app/models/api"
	appContext "github.com/OkciD/whos_on_call/internal/pkg/context"
	"github.com/OkciD/whos_on_call/internal/pkg/errors"
	"github.com/OkciD/whos_on_call/internal/pkg/http/handler"
)

func (h *UserHandler) CreateDevice(r *http.Request) (handler.ResponseWriter, error) {
	decoder := json.NewDecoder(r.Body)
	var newDeviceInput api.Device
	if err := decoder.Decode(&newDeviceInput); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidJSON, err)
	}

	user, err := appContext.GetUser(r.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get user from request: %w", err)
	}

	newDeviceApp, err := newDeviceInput.ToAppModel()
	if err != nil {
		return nil, fmt.Errorf("failed to convert to app device model: %w", err)
	}
	newDeviceApp.User = user

	newDeviceApp, err = h.deviceUseCase.Create(r.Context(), newDeviceApp)
	if err != nil {
		return nil, fmt.Errorf("failed to create device: %w", err)
	}

	newDeviceOutput, err := api.FromDeviceAppModel(newDeviceApp)
	if err != nil {
		return nil, fmt.Errorf("failed to convert back to api device model: %w", err)
	}

	return handler.RespondJSONWithStatus(http.StatusCreated, newDeviceOutput), nil
}
