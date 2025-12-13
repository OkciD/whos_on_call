package handler

import (
	"encoding/json"
	"net/http"

	appErrors "github.com/OkciD/whos_on_call/internal/pkg/errors"
	httpErrors "github.com/OkciD/whos_on_call/internal/pkg/http/errors"
)

type SuccessResponse struct {
	Status *int
	Body   any
}

type Handler func(r *http.Request) (SuccessResponse, error)

func GenericHandler(handlers map[string]Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		handler, ok := handlers[r.Method]
		if !ok {
			httpErrors.MapErrorToResponse(w, appErrors.ErrNotFound)
			return
		}

		response, err := handler(r)
		if err != nil {
			httpErrors.MapErrorToResponse(w, err)
			return
		}

		// todo: support redirects

		if response.Status != nil {
			w.WriteHeader(*response.Status)
		} else {
			w.WriteHeader(http.StatusOK)
		}

		if response.Body != nil {
			err = json.NewEncoder(w).Encode(response.Body)
			if err != nil {
				httpErrors.MapErrorToResponse(w, err)
				return
			}
		}
	})
}
