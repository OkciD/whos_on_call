package http

import (
	"fmt"
	"net/http"

	"github.com/OkciD/whos_on_call/internal/pkg/http/handler"
	"github.com/OkciD/whos_on_call/internal/pkg/http/middleware"
)

type ApiUser struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (h *UserHandler) GetUser(r *http.Request) (handler.SuccessResponse, error) {
	user, err := middleware.GetUserFromRequest(r)
	if err != nil {
		return handler.SuccessResponse{}, fmt.Errorf("failed to get user from request: %w", err)
	}

	return handler.SuccessResponse{
		Body: ApiUser{
			ID:   user.ID,
			Name: user.Name,
		},
	}, nil
}
