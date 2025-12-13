package http

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/app/user"
	"github.com/OkciD/whos_on_call/internal/pkg/http/handler"
)

type UserHandler struct {
	userUseCase user.UseCase
}

func New(mux *http.ServeMux, userUseCase user.UseCase) *UserHandler {
	h := &UserHandler{
		userUseCase: userUseCase,
	}

	mux.Handle("/api/v1/user", handler.GenericHandler(map[string]handler.Handler{
		http.MethodGet: h.GetUser,
	}))

	return h
}
