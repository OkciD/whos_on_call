package http

import (
	"context"
	"fmt"

	"github.com/OkciD/whos_on_call/cmd/server/apiserver/gen"
	appContext "github.com/OkciD/whos_on_call/internal/server/pkg/context"
	"github.com/OkciD/whos_on_call/internal/shared/models/api"
)

// GET /api/v1/user
func (h UserHandler) GetUser(ctx context.Context, request gen.GetUserRequestObject) (gen.GetUserResponseObject, error) {
	user, err := appContext.GetUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from request: %w", err)
	}

	return gen.GetUser200JSONResponse(*api.FromUserAppModel(user)), nil
}
