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

	_ "github.com/mattn/go-sqlite3"

	userHttpDelivery "github.com/OkciD/whos_on_call/internal/app/user/delivery/http"
	userRepositoryInmemory "github.com/OkciD/whos_on_call/internal/app/user/repository/inmemory"
	userUseCase "github.com/OkciD/whos_on_call/internal/app/user/usecase"
	configUtils "github.com/OkciD/whos_on_call/internal/pkg/config"
	"github.com/OkciD/whos_on_call/internal/pkg/http/middleware"
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

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unk"
	}

	logger := logrus.WithField("host", hostname)

	if cfg.GetLogFormat() == config.LogFormatText {
		logger.Logger.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: time.RFC3339,
		})
	} else {
		logger.Logger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		})
	}

	logLevel, err := logrus.ParseLevel(string(cfg.Logger.Level))
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	logger.Logger.SetLevel(logLevel)

	db, err := sql.Open("sqlite3", cfg.DB.DSN)
	if err != nil {
		logger.WithError(err).WithField("dsn", cfg.DB.DSN).Fatal("failed to open db")
	}
	defer db.Close()

	logger.Info("open db successfully")

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

	userRepo, err := userRepositoryInmemory.New(logger.WithField("module", "user_repo_inmem"), &cfg.User.Repository)
	if err != nil {
		logger.WithError(err).Fatalf("failed to create user repository")
	}

	userUseCase := userUseCase.New(logger.WithField("module", "user_usecase"), userRepo)

	mux := http.NewServeMux()

	userHttpDelivery.New(mux, userUseCase)

	contentTypeMiddleware := middleware.NewContentTypeMiddleware("application/json")
	requestIdMiddleWare := middleware.NewRequestIdMiddleware()
	accessLogMiddleware := middleware.NewAccessLogMiddleware(logger)
	authMiddleware := middleware.NewAuthMiddleware(userUseCase)

	wrappedMux := contentTypeMiddleware(requestIdMiddleWare(accessLogMiddleware(authMiddleware(mux))))

	logger.WithField("addr", cfg.Server.ListenAddr).Info("http server starting")
	err = http.ListenAndServe(cfg.Server.ListenAddr, wrappedMux)
	if err != nil {
		logger.WithError(err).Fatalf("http server error")
	}
}
