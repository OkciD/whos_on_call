package http

import (
	"fmt"
	"net/http"

	"github.com/OkciD/whos_on_call/internal/app/models/api"
	appContext "github.com/OkciD/whos_on_call/internal/pkg/context"
	"github.com/OkciD/whos_on_call/internal/pkg/http/handler"
)

func (h *UserHandler) GetUser(r *http.Request) (handler.ResponseWriter, error) {
	user, err := appContext.GetUser(r.Context())
	if err != nil {
		return nil, fmt.Errorf("failed to get user from request: %w", err)
	}

	return handler.RespondJSON(api.FromUserAppModel(user)), nil
}
