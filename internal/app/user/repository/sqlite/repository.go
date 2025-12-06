package inmemory

import (
	"database/sql"

	"github.com/OkciD/whos_on_call/internal/app/user"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	logger *logrus.Entry

	db *sql.DB
}

func New(logger *logrus.Entry, db *sql.DB) (user.Repository, error) {
	return &Repository{
		logger: logger,

		db: db,
	}, nil
}
