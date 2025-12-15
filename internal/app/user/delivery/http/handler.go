package http

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/app/user"
	"github.com/OkciD/whos_on_call/internal/pkg/http/handler"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
)

type UserHandler struct {
	logger logger.Logger

	userUseCase user.UseCase
}

func New(mux *http.ServeMux, logger logger.Logger, userUseCase user.UseCase) *UserHandler {
	h := &UserHandler{
		logger: logger,

		userUseCase: userUseCase,
	}

	mux.Handle("/api/v1/user", handler.GenericHandler(logger, map[string]handler.Handler{
		http.MethodGet: h.GetUser,
	}))

	return h
}
