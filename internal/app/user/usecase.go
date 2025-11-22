package user

import (
	"context"

	"github.com/OkciD/whos_on_call/internal/app/models"
)

type UseCase interface {
	Authorize(ctx context.Context, apiKey string) (*models.User, error)
}
