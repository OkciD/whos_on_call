package inmemory

import (
	"database/sql"

	"github.com/OkciD/whos_on_call/internal/app/user"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
)

type Repository struct {
	logger logger.Logger

	db *sql.DB
}

func New(logger logger.Logger, db *sql.DB) user.Repository {
	return &Repository{
		logger: logger,

		db: db,
	}
}
