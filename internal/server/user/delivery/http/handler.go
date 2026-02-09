package http

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/pkg/logger"
	"github.com/OkciD/whos_on_call/internal/server/pkg/http/handler"
	"github.com/OkciD/whos_on_call/internal/server/user"
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

	mux.Handle("GET /api/v1/user", handler.GenericHandler(logger, h.GetUser))

	return h
}
