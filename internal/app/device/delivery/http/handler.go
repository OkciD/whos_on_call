package http

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/app/device"
	"github.com/OkciD/whos_on_call/internal/pkg/http/handler"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
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

	mux.Handle("/api/v1/device", handler.GenericHandler(logger, map[string]handler.Handler{
		http.MethodPost: h.CreateDevice,
	}))

	return h
}
