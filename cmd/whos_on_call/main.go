package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/OkciD/whos_on_call/cmd/whos_on_call/config"
	"github.com/sirupsen/logrus"

	"github.com/pressly/goose/v3"

	_ "github.com/mattn/go-sqlite3"

	userHttpDelivery "github.com/OkciD/whos_on_call/internal/app/user/delivery/http"
	userRepositorySqlite "github.com/OkciD/whos_on_call/internal/app/user/repository/sqlite"
	userUseCase "github.com/OkciD/whos_on_call/internal/app/user/usecase"
	configUtils "github.com/OkciD/whos_on_call/internal/pkg/config"
	"github.com/OkciD/whos_on_call/internal/pkg/db/migrations"
	"github.com/OkciD/whos_on_call/internal/pkg/http/middleware"

	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/logrusadapter"
)

func main() {
	configFilePathPtr := flag.String("config", "", "path to config file")

	flag.Parse()

	if configFilePathPtr == nil || *configFilePathPtr == "" {
		log.Fatal("-config option required")
	}

	cfg, err := configUtils.ReadConfig[config.Config](*configFilePathPtr)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to read config: %w", err))
	}

	logger := logrus.New()

	if cfg.GetLogFormat() == config.LogFormatText {
		logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
		})
	} else {
		logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	}

	logLevel, err := logrus.ParseLevel(string(cfg.Logger.Level))
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logger.SetLevel(logLevel)

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unk"
	}
	logEntry := logger.WithField("host", hostname)

	db, err := sql.Open("sqlite3", cfg.DB.DSN)
	if err != nil {
		logger.WithError(err).WithField("dsn", cfg.DB.DSN).Fatal("failed to open db")
	}
	defer db.Close()

	logger.Info("open db successfully")

	loggerAdapter := logrusadapter.New(logger)
	db = sqldblogger.OpenDriver(
		cfg.DB.DSN,
		db.Driver(),
		loggerAdapter,
		sqldblogger.WithTimeFormat(sqldblogger.TimeFormatRFC3339),
		sqldblogger.WithExecerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithQueryerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithPreparerLevel(sqldblogger.LevelDebug),
	)

	db.SetMaxOpenConns(cfg.DB.MaxOpenConns)
	db.SetMaxIdleConns(cfg.DB.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.DB.ConnMaxLifetime.Duration)
	db.SetConnMaxIdleTime(cfg.DB.ConnMaxIdleTime.Duration)

	pingCtx, cancel := context.WithTimeout(context.Background(), cfg.DB.PingTimeout.Duration)
	if err := db.PingContext(pingCtx); err != nil {
		logger.WithError(err).Fatal("unable to connect to database")
	}
	cancel()

	logger.Info("ping db successfully")

	goose.SetBaseFS(migrations.EmbedMigrations)

	if err := goose.SetDialect("sqlite3"); err != nil {
		logger.WithError(err).Fatal("error setting goose dialect")
	}

	if err := goose.Up(db, "."); err != nil {
		logger.WithError(err).Fatal("failed to apply migrations")
	}
	logger.Info("migrations applied")

	userRepo := userRepositorySqlite.New(logEntry.WithField("module", "user_repo"), db)

	userUseCase := userUseCase.New(logEntry.WithField("module", "user_usecase"), userRepo)

	mux := http.NewServeMux()

	userHttpDelivery.New(mux, userUseCase)

	contentTypeMiddleware := middleware.NewContentTypeMiddleware("application/json")
	requestIdMiddleWare := middleware.NewRequestIdMiddleware()
	accessLogMiddleware := middleware.NewAccessLogMiddleware(logEntry)
	authMiddleware := middleware.NewAuthMiddleware(userUseCase)

	wrappedMux := contentTypeMiddleware(requestIdMiddleWare(accessLogMiddleware(authMiddleware(mux))))

	logger.WithField("addr", cfg.Server.ListenAddr).Info("http server starting")
	err = http.ListenAndServe(cfg.Server.ListenAddr, wrappedMux)
	if err != nil {
		logger.WithError(err).Fatalf("http server error")
	}
}
