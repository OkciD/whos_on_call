package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mattn/go-sqlite3"
	sqldblogger "github.com/simukti/sqldb-logger"

	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
	sqlDBLoggerAdapter "github.com/OkciD/whos_on_call/internal/shared/pkg/logger/sqldbloggeradapter"
)

func NewDBConnection(logger logger.Logger, cfg *Config) (*sql.DB, error) {
	loggerAdapter := sqlDBLoggerAdapter.New(logger)
	db := sqldblogger.OpenDriver(
		cfg.DSN,
		&sqlite3.SQLiteDriver{},
		loggerAdapter,
		sqldblogger.WithTimeFormat(sqldblogger.TimeFormatRFC3339),
		sqldblogger.WithExecerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithQueryerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithPreparerLevel(sqldblogger.LevelDebug),
	)

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime.Duration)
	db.SetConnMaxIdleTime(cfg.ConnMaxIdleTime.Duration)

	if cfg.Stats.Enabled {
		startDBMonitoring(db, logger, cfg.Stats.TickerDuration.Duration)
	}

	pingCtx, cancel := context.WithTimeout(context.Background(), cfg.PingTimeout.Duration)
	defer cancel()
	if err := db.PingContext(pingCtx); err != nil {
		return nil, fmt.Errorf("failed to ping db by dsn %s: %w", cfg.DSN, err)
	}
	logger.Info("open db successfully")

	return db, nil
}

func Close(db *sql.DB, logger logger.Logger) {
	stopDBMonitoring(logger)

	err := db.Close()
	if err != nil {
		logger.WithError(err).Error("error closing db")
	}
}
