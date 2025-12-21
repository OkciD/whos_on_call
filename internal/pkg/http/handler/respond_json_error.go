package handler

import (
	"errors"
	"net/http"

	appErrors "github.com/OkciD/whos_on_call/internal/pkg/errors"
)

type errorResponse struct {
	ErrCode string `json:"err_code"`
}

func respondJSONError(w http.ResponseWriter, errCode string, status int) {
	err := RespondJSONWithStatus(status, errorResponse{ErrCode: errCode})(w)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"err_code\":\"internal\"}"))
		return
	}
}

func RespondJSONError(w http.ResponseWriter, err error) {
	if errors.Is(err, appErrors.ErrUnauthorized) || errors.Is(err, appErrors.ErrUserNotFound) {
		respondJSONError(w, "unauthorized", http.StatusUnauthorized)
	} else if errors.Is(err, appErrors.ErrNotFound) || errors.Is(err, appErrors.ErrDeviceNotFound) {
		respondJSONError(w, "not found", http.StatusNotFound)
	} else if errors.Is(err, appErrors.ErrNotImplemented) {
		respondJSONError(w, "not impl", http.StatusNotImplemented)
	} else if errors.Is(err, appErrors.ErrInvalidJSON) {
		respondJSONError(w, "invalid json", http.StatusBadRequest)
	} else if errors.Is(err, appErrors.ErrDeviceExists) {
		respondJSONError(w, "duplicate", http.StatusBadRequest)
	} else {
		respondJSONError(w, "internal", http.StatusInternalServerError)
	}
}
