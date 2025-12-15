package handler

import (
	"encoding/json"
	"net/http"
)

func RespondJSONWithStatus(status int, body any) func(w http.ResponseWriter) error {
	return func(w http.ResponseWriter) error {
		w.Header().Add("Content-Type", "application/json")

		w.WriteHeader(status)

		if body != nil {
			err := json.NewEncoder(w).Encode(body)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func RespondJSON(body any) func(w http.ResponseWriter) error {
	return RespondJSONWithStatus(http.StatusOK, body)
}
