package db

import (
	"database/sql"
	"time"

	loggerPkg "github.com/OkciD/whos_on_call/internal/pkg/logger"
)

var ticker *time.Ticker

func startDBMonitoring(db *sql.DB, logger loggerPkg.Logger, duration time.Duration) {
	ticker = time.NewTicker(duration)

	logger.Debug("start db stats ticker")
	go func() {
		for range ticker.C {
			stats := db.Stats()
			logger.WithFields(loggerPkg.Fields{
				"openConnections":   stats.OpenConnections,
				"inUse":             stats.InUse,
				"idle":              stats.Idle,
				"waitCount":         stats.WaitCount,
				"waitDuration":      stats.WaitDuration,
				"maxIdleClosed":     stats.MaxIdleClosed,
				"maxIdleTimeClosed": stats.MaxIdleTimeClosed,
				"maxLifetimeClosed": stats.MaxLifetimeClosed,
			}).Debug("db stats")
		}
	}()
}

func stopDBMonitoring(logger loggerPkg.Logger) {
	ticker.Stop()
	logger.Debug("stop db stats ticker")
}
