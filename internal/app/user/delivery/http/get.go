package http

import (
	"encoding/json"
	"errors"
	"net/http"

	appErrors "github.com/OkciD/whos_on_call/internal/pkg/errors"
)

const API_KEY_HEADER = "X-Api-Key"

type errorResponse struct {
	Err string `json:"err"`
}

type ApiUser struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	apiKey := r.Header.Get(API_KEY_HEADER)

	user, err := h.userUseCase.GetUserByApiKey(r.Context(), apiKey)
	if err != nil {
		if errors.Is(err, appErrors.ErrUserNotFound) {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(errorResponse{Err: "unauthorized"})
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errorResponse{Err: "internal"})
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ApiUser{
		UID:  user.UID.String(),
		Name: user.Name,
	})
}
