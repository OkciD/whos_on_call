package sqlite

import (
	"database/sql"

	"github.com/OkciD/whos_on_call/internal/server/devicefeature"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
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
