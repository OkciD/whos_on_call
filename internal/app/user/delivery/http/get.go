package http

import (
	"encoding/json"
	"net/http"

	"github.com/OkciD/whos_on_call/internal/pkg/http/middleware"
)

type ApiUser struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	user, _ := middleware.GetUserFromRequest(r)

	// todo: respond ok tool
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ApiUser{
		UID:  user.UID.String(),
		Name: user.Name,
	})
}
