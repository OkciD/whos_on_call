package usecase

import (
	"github.com/OkciD/whos_on_call/internal/app/user"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
)

type UseCase struct {
	logger logger.Logger

	userRepo user.Repository
}

func New(logger logger.Logger, userRepo user.Repository) user.UseCase {
	return &UseCase{
		logger: logger,

		userRepo: userRepo,
	}
}
