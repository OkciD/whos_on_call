package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/OkciD/whos_on_call/cmd/whos_on_call/config"

	"github.com/pressly/goose/v3"

	_ "github.com/mattn/go-sqlite3"

	userHttpDelivery "github.com/OkciD/whos_on_call/internal/app/user/delivery/http"
	userRepositorySqlite "github.com/OkciD/whos_on_call/internal/app/user/repository/sqlite"
	userUseCase "github.com/OkciD/whos_on_call/internal/app/user/usecase"
	configUtils "github.com/OkciD/whos_on_call/internal/pkg/config"
	"github.com/OkciD/whos_on_call/internal/pkg/db"
	"github.com/OkciD/whos_on_call/internal/pkg/db/migrations"
	"github.com/OkciD/whos_on_call/internal/pkg/http/middleware"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
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

	logger := logger.NewLogrusBasedLogger(&cfg.Logger)

	db, err := db.NewDBConnection(logger, &cfg.DB)
	if err != nil {
		logger.WithError(err).Fatal("db connection failed")
	}
	defer db.Close()

	goose.SetBaseFS(migrations.EmbedMigrations)

	if err := goose.SetDialect(cfg.DB.Driver); err != nil {
		logger.WithError(err).Fatal("error setting goose dialect")
	}

	if err := goose.Up(db, "."); err != nil {
		logger.WithError(err).Fatal("failed to apply migrations")
	}
	logger.Info("migrations applied")

	userRepo := userRepositorySqlite.New(logger.ForModule("user_repo"), db)

	userUseCase := userUseCase.New(logger.ForModule("user_usecase"), userRepo)

	mux := http.NewServeMux()

	userHttpDelivery.New(mux, userUseCase)

	wrappedMux := middleware.ApplyMiddlewares(
		mux,
		middleware.NewAuthMiddleware(userUseCase),
		middleware.NewAccessLogMiddleware(logger),
		middleware.NewRequestIdMiddleware(),
		middleware.NewRecoveryMiddleware(logger),
	)

	logger.WithField("addr", cfg.Server.ListenAddr).Info("http server starting")
	err = http.ListenAndServe(cfg.Server.ListenAddr, wrappedMux)
	if err != nil {
		logger.WithError(err).Fatal("http server error")
	}
}
