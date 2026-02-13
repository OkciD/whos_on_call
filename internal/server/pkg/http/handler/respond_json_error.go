package handler

import (
	"net/http"

	appErrors "github.com/OkciD/whos_on_call/internal/errors"
)

type errorResponseJSON struct {
	ErrCode string `json:"err_code"`
}

func RespondJSONError(w http.ResponseWriter, err error) {
	resp := appErrors.ErrorToResp(err)
	respondError := RespondJSONWithStatus(resp.StatusCode, errorResponseJSON{ErrCode: resp.ErrCode})(w)

	if respondError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"err_code\":\"internal\"}"))
		return
	}
}
