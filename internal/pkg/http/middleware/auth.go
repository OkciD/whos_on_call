package middleware

import (
	"context"
	"net/http"

	"github.com/OkciD/whos_on_call/internal/app/models"
	"github.com/OkciD/whos_on_call/internal/app/user"
	appErrors "github.com/OkciD/whos_on_call/internal/pkg/errors"
	"github.com/OkciD/whos_on_call/internal/pkg/http/handler"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
)

const API_KEY_HEADER = "X-Api-Key"

type userCtxKey struct{}

func NewAuthMiddleware(logger logger.Logger, userUseCase user.UseCase) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get(API_KEY_HEADER)

			user, err := userUseCase.GetUserByApiKey(r.Context(), apiKey)
			if err != nil {
				logger.WithError(err).Error("failed to get user by api key")
				handler.RespondJSONError(w, err)
				return
			}

			contextWithUser := context.WithValue(r.Context(), userCtxKey{}, user)
			requestWithUser := r.WithContext(contextWithUser)

			next.ServeHTTP(w, requestWithUser)
		})
	}
}

func GetUserFromRequest(r *http.Request) (*models.User, error) {
	user, ok := r.Context().Value(userCtxKey{}).(*models.User)
	if !ok {
		return nil, appErrors.ErrUnauthorized
	}

	return user, nil
}
