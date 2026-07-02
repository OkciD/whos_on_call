package http

import (
	"context"
	"fmt"

	"github.com/OkciD/whos_on_call/cmd/server/apiserver/gen"
	appContext "github.com/OkciD/whos_on_call/internal/server/pkg/context"
	"github.com/OkciD/whos_on_call/internal/shared/models/api"
)

func (h *Handler) GetStatus(ctx context.Context, request gen.GetStatusRequestObject) (gen.GetStatusResponseObject, error) {
	_, err := appContext.GetUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from request: %w", err)
	}

	appCallStatus, err := h.callStatusUseCase.Calculate(ctx)
	if err != nil {
		return nil, fmt.Errorf("error calculating status: %w", err)
	}

	apiCallStatus, err := api.FromAppCallStatus(appCallStatus)
	if err != nil {
		return nil, fmt.Errorf("error transforming call status from app to api: %w", err)
	}

	return gen.GetStatus200JSONResponse(apiCallStatus), nil
}
