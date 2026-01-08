package html

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/app/callstatus"
	"github.com/OkciD/whos_on_call/internal/app/web/static"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
)

type WebHandler struct {
	logger logger.Logger

	callStatusUseCase callstatus.UseCase
}

func New(mux *http.ServeMux, logger logger.Logger, callStatusUseCase callstatus.UseCase) *WebHandler {
	h := &WebHandler{
		logger: logger,

		callStatusUseCase: callStatusUseCase,
	}

	mux.Handle("GET /", h.CallStatus())
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(static.StaticFS))))

	return h
}
