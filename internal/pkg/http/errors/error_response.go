package errors

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	ErrCode string `json:"err_code"`
}

func RespondWithError(w http.ResponseWriter, errCode string, status int) {
	errJson, _ := json.Marshal(errorResponse{ErrCode: errCode})

	http.Error(w, string(errJson), status)
}
