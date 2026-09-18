package apiserver

import (
	"context"
	"fmt"

	"github.com/getkin/kin-openapi/openapi3filter"

	appContext "github.com/OkciD/whos_on_call/internal/server/pkg/context"
	"github.com/OkciD/whos_on_call/internal/server/user"
	"github.com/OkciD/whos_on_call/internal/shared/errors"
	loggerPkg "github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

func NewAuthFunc(logger loggerPkg.Logger, userUseCase user.UseCase) openapi3filter.AuthenticationFunc {
	return func(ctx context.Context, input *openapi3filter.AuthenticationInput) error {
		if input.SecurityScheme.Type == "apiKey" && input.SecurityScheme.In == "header" {
			return apiKeyHeaderAuthFn(ctx, input, logger, userUseCase)
		}

		logger.WithField("name", input.SecuritySchemeName).Error("security scheme not supported")
		return errors.ErrUnauthorized
	}
}

func apiKeyHeaderAuthFn(
	ctx context.Context,
	input *openapi3filter.AuthenticationInput,
	logger loggerPkg.Logger,
	userUseCase user.UseCase,
) error {
	r := input.RequestValidationInput.Request
	apiKey := r.Header.Get(input.SecurityScheme.Name)

	user, err := userUseCase.GetUserByAPIKey(ctx, apiKey)
	if err != nil {
		return fmt.Errorf("%w, %w", errors.ErrUnauthorized, err)
	}
	logger.WithFields(loggerPkg.Fields{
		"id":   user.ID,
		"name": user.Name,
	}).Info("got user")

	contextWithUser := appContext.StoreUser(r.Context(), user)
	*r = *r.WithContext(contextWithUser)

	return nil
}
