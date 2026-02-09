package sqlite

import (
	"database/sql"

	"github.com/OkciD/whos_on_call/internal/pkg/logger"
	"github.com/OkciD/whos_on_call/internal/server/devicefeature"
)

type Repository struct {
	logger logger.Logger

	db *sql.DB
}

func New(logger logger.Logger, db *sql.DB) devicefeature.Repository {
	return &Repository{
		logger: logger,

		db: db,
	}
}
