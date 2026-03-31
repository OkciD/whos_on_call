package http

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/server/device"
	"github.com/OkciD/whos_on_call/internal/server/pkg/http/handler"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type UserHandler struct {
	logger logger.Logger

	deviceUseCase device.UseCase
}

func New(mux *http.ServeMux, logger logger.Logger, deviceUseCase device.UseCase) *UserHandler {
	h := &UserHandler{
		logger: logger,

		deviceUseCase: deviceUseCase,
	}

	mux.Handle("POST /api/v1/device", handler.GenericHandler(logger, h.CreateDevice))

	return h
}
