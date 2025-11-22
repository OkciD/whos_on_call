package user

import "github.com/OkciD/whos_on_call/internal/app/models"

type UseCase interface {
	Authorize(apiKey string) (*models.User, error)
}
