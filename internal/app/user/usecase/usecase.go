package usecase

import (
	"github.com/OkciD/whos_on_call/internal/app/user"
	"github.com/sirupsen/logrus"
)

type UseCase struct {
	logger *logrus.Entry

	userRepo user.Repository
}

func New(logger *logrus.Entry, userRepo user.Repository) user.UseCase {
	return &UseCase{
		logger: logger,

		userRepo: userRepo,
	}
}
