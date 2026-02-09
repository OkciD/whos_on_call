package http

import (
	"fmt"
	"net/http"

	"github.com/OkciD/whos_on_call/internal/models/api"
	appContext "github.com/OkciD/whos_on_call/internal/server/pkg/context"
	"github.com/OkciD/whos_on_call/internal/server/pkg/http/handler"
)

func (h *CallStatusHandler) GetStatus(r *http.Request) (handler.ResponseWriter, error) {
	_, err := appContext.GetUser(r.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get user from request: %w", err)
	}

	appCallStatus, err := h.callStatusUseCase.Calculate(r.Context())
	if err != nil {
		return nil, fmt.Errorf("error calculating status: %w", err)
	}

	apiCallStatus, err := api.FromAppCallStatus(appCallStatus)
	if err != nil {
		return nil, fmt.Errorf("error transforming call status from app to api: %w", err)
	}

	return handler.RespondJSON(apiCallStatus), nil
}
