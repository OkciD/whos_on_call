package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/OkciD/whos_on_call/internal/errors"
	"github.com/OkciD/whos_on_call/internal/models/api"
	appContext "github.com/OkciD/whos_on_call/internal/server/pkg/context"
	"github.com/OkciD/whos_on_call/internal/server/pkg/http/handler"
)

func (h *DeviceFeatureHandler) Update(r *http.Request) (handler.ResponseWriter, error) {
	deviceIdStr := r.PathValue("deviceid")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrInvalidUrlParam, err)
	}

	decoder := json.NewDecoder(r.Body)
	var newDeviceFeatureInput api.DeviceFeature
	if err := decoder.Decode(&newDeviceFeatureInput); err != nil {
		return nil, fmt.Errorf("%w: %w", errors.ErrBadJSON, err)
	}

	newDeviceFeatureApp, err := newDeviceFeatureInput.ToAppModel()
	if err != nil {
		return nil, fmt.Errorf("failed to convert to app device feature model: %w", err)
	}

	user, err := appContext.GetUser(r.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get user from request: %w", err)
	}

	newDeviceFeatureApp, err = h.deviceFeatureUseCase.Upsert(r.Context(), deviceId, user, newDeviceFeatureApp)
	if err != nil {
		return nil, fmt.Errorf("failed to upsert device feature: %w", err)
	}

	newDeviceFeatureApi, err := api.FromDeviceFeatureAppModel(newDeviceFeatureApp)
	if err != nil {
		return nil, fmt.Errorf("error converting device feature to api model: %w", err)
	}

	return handler.RespondJSON(newDeviceFeatureApi), nil
}
