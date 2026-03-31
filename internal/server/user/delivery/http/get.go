package http

import (
	"fmt"
	"net/http"

	appContext "github.com/OkciD/whos_on_call/internal/server/pkg/context"
	"github.com/OkciD/whos_on_call/internal/server/pkg/http/handler"
	"github.com/OkciD/whos_on_call/internal/shared/models/api"
)

func (h *UserHandler) GetUser(r *http.Request) (handler.ResponseWriter, error) {
	user, err := appContext.GetUser(r.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get user from request: %w", err)
	}

	return handler.RespondJSON(api.FromUserAppModel(user)), nil
}
