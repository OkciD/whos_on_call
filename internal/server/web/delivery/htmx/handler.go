package htmx

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/server/web/assets"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type Handler struct {
	logger logger.Logger
}

func New(
	mux *http.ServeMux,
	logger logger.Logger,
) *Handler {
	h := &Handler{
		logger: logger,
	}

	mux.Handle("GET /", h.callStatusPage())
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(assets.StaticFiles))))

	return h
}
