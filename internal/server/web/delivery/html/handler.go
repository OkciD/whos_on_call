package html

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/server/callstatus"
	"github.com/OkciD/whos_on_call/internal/server/web"
	"github.com/OkciD/whos_on_call/internal/server/web/assets"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type WebHandler struct {
	logger logger.Logger

	config web.Config

	callStatusUseCase callstatus.UseCase
}

func New(
	mux *http.ServeMux,
	logger logger.Logger,
	callStatusUseCase callstatus.UseCase,
	config web.Config,
) *WebHandler {
	h := &WebHandler{
		logger: logger,

		config: config,

		callStatusUseCase: callStatusUseCase,
	}

	mux.Handle("GET /", h.callStatusPage())
	mux.Handle("GET /partials/callstatus", h.callStatusPartial())
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(assets.StaticFiles))))

	return h
}
