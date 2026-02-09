package middleware

import (
	"net/http"

	"github.com/OkciD/whos_on_call/internal/pkg/logger"
	appContext "github.com/OkciD/whos_on_call/internal/server/pkg/context"
	"github.com/OkciD/whos_on_call/internal/server/pkg/http/handler"
	"github.com/OkciD/whos_on_call/internal/server/user"
)

const API_KEY_HEADER = "X-Api-Key"

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

			contextWithUser := appContext.StoreUser(r.Context(), user)
			requestWithUser := r.WithContext(contextWithUser)

			next.ServeHTTP(w, requestWithUser)
		})
	}
}
