package errors

import (
	"encoding/json"
	"errors"
	"net/http"

	appErrors "github.com/OkciD/whos_on_call/internal/pkg/errors"
)

type errorResponse struct {
	ErrCode string `json:"err_code"`
}

func respondWithError(w http.ResponseWriter, errCode string, status int) {
	respBytes, err := json.Marshal(errorResponse{ErrCode: errCode})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"err_code\":\"internal\"}"))
		return
	}

	w.WriteHeader(status)
	w.Write(respBytes)
}

func MapErrorToResponse(w http.ResponseWriter, err error) {
	if errors.Is(err, appErrors.ErrUnauthorized) || errors.Is(err, appErrors.ErrUserNotFound) {
		respondWithError(w, "unauthorized", http.StatusUnauthorized)
	} else if errors.Is(err, appErrors.ErrNotFound) {
		respondWithError(w, "not found", http.StatusNotFound)
	} else {
		respondWithError(w, "internal", http.StatusInternalServerError)
	}
}
