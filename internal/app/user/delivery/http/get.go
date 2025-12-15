package http

import (
	"fmt"
	"net/http"

	"github.com/OkciD/whos_on_call/internal/app/models/api"
	"github.com/OkciD/whos_on_call/internal/pkg/http/handler"
	"github.com/OkciD/whos_on_call/internal/pkg/http/middleware"
)

func (h *UserHandler) GetUser(r *http.Request) (handler.ResponseWriter, error) {
	user, err := middleware.GetUserFromRequest(r)
	if err != nil {
		return nil, fmt.Errorf("failed to get user from request: %w", err)
	}

	return handler.RespondJSON(api.FromUserAppModel(user)), nil
}
