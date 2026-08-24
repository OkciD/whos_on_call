package middleware

import (
	"errors"
	"net/http"

	appContext "github.com/OkciD/whos_on_call/internal/server/pkg/context"
	"github.com/OkciD/whos_on_call/internal/server/user"
	appErrors "github.com/OkciD/whos_on_call/internal/shared/errors"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

//nolint:gosec // фолзит
const APIKeyHeader = "X-Api-Key"

func NewAuthMiddleware(logger logger.Logger, userUseCase user.UseCase) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get(APIKeyHeader)

			user, err := userUseCase.GetUserByAPIKey(r.Context(), apiKey)
			if err != nil {
				logger.WithError(err).Error("failed to get user by api key")

				// todo: не писать ответ "руками"
				w.Header().Add("Content-Type", "application/json")

				if errors.Is(err, appErrors.ErrEntityNotFound) {
					w.WriteHeader(http.StatusUnauthorized)
					_, err := w.Write([]byte("{\"code\":\"unauthorized\"}"))
					if err != nil {
						logger.WithError(err).Error("error while writing 401 error from auth middleware")
					}
				}

				w.WriteHeader(http.StatusInternalServerError)
				_, err := w.Write([]byte("{\"code\":\"internal\"}"))
				if err != nil {
					logger.WithError(err).Error("error while writing 500 error from auth middleware")
				}
				return
			}

			contextWithUser := appContext.StoreUser(r.Context(), user)
			requestWithUser := r.WithContext(contextWithUser)

			next.ServeHTTP(w, requestWithUser)
		})
	}
}
