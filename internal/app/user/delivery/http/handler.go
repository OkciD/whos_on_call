package http

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/app/user"
)

type UserHandler struct {
	userUseCase user.UseCase
}

func New(mux *http.ServeMux, userUseCase user.UseCase) *UserHandler {
	h := &UserHandler{
		userUseCase: userUseCase,
	}

	mux.Handle("/api/v1/user", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.GetUser(w, r)
		}
	}))

	return h
}
