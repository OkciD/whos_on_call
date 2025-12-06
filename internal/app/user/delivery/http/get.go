package http

import (
	"encoding/json"
	"net/http"

	"github.com/OkciD/whos_on_call/internal/pkg/http/middleware"
)

type ApiUser struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	user, _ := middleware.GetUserFromRequest(r)

	// todo: respond ok tool
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ApiUser{
		ID:   user.ID,
		Name: user.Name,
	})
}
