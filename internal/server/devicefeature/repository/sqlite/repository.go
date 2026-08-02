package sqlite

import (
	"github.com/OkciD/whos_on_call/internal/server/devicefeature"
	"github.com/OkciD/whos_on_call/internal/server/pkg/db"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
)

type Repository struct {
	db.WithExecutor

	logger logger.Logger
}

func New(logger logger.Logger, e db.Executor) devicefeature.Repository {
	return &Repository{
		WithExecutor: db.NewWithExecutor(e),

		logger: logger,
	}
}
