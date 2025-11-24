package usecase

import (
	"context"

	"github.com/OkciD/whos_on_call/internal/app/models"
)

func (u *UseCase) GetUserByApiKey(ctx context.Context, apiKey string) (*models.User, error) {
	return u.userRepo.GetUserByApiKey(ctx, apiKey)
}
