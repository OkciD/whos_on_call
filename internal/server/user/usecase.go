package user

import (
	"context"

	"github.com/OkciD/whos_on_call/internal/shared/models"
)

type UseCase interface {
	GetUserByApiKey(ctx context.Context, apiKey string) (*models.User, error)
}
