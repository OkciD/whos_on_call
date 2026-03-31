package http

import (
	"fmt"
	"net/http"

	appContext "github.com/OkciD/whos_on_call/internal/server/pkg/context"
	"github.com/OkciD/whos_on_call/internal/server/pkg/http/handler"
	"github.com/OkciD/whos_on_call/internal/shared/models/api"
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
