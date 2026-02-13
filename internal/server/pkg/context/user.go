package context

import (
	"context"

	appErrors "github.com/OkciD/whos_on_call/internal/errors"
	"github.com/OkciD/whos_on_call/internal/models"
)

type userCtxKey struct{}

func StoreUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userCtxKey{}, user)
}

func GetUser(ctx context.Context) (*models.User, error) {
	user, ok := ctx.Value(userCtxKey{}).(*models.User)
	if !ok {
		return nil, appErrors.ErrUnauthorized
	}

	return user, nil
}
