package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/OkciD/whos_on_call/internal/app/models"
	"github.com/OkciD/whos_on_call/internal/app/user"
	appErrors "github.com/OkciD/whos_on_call/internal/pkg/errors"
	httpErrors "github.com/OkciD/whos_on_call/internal/pkg/http/errors"
)

const API_KEY_HEADER = "X-Api-Key"
const USER_CTX_KEY = "user"

func NewAuthMiddleware(userUseCase user.UseCase) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get(API_KEY_HEADER)

			user, err := userUseCase.GetUserByApiKey(r.Context(), apiKey)
			if err != nil {
				if errors.Is(err, appErrors.ErrUserNotFound) {
					httpErrors.RespondWithError(w, "unauthorized", http.StatusUnauthorized)
					return
				} else {
					httpErrors.RespondWithError(w, "internal", http.StatusInternalServerError)
					return
				}
			}

			contextWithUser := context.WithValue(r.Context(), USER_CTX_KEY, user)
			requestWithUser := r.WithContext(contextWithUser)

			next.ServeHTTP(w, requestWithUser)
		})
	}
}

func GetUserFromRequest(r *http.Request) (*models.User, error) {
	user := r.Context().Value(USER_CTX_KEY).(*models.User)
	if (user == nil || *user == models.User{}) {
		return nil, appErrors.Unauthorized
	}

	return user, nil
}
