package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pressly/goose/v3"

	_ "github.com/mattn/go-sqlite3"

	configUtils "github.com/OkciD/whos_on_call/internal/pkg/config"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
	callStatusDelivery "github.com/OkciD/whos_on_call/internal/server/callstatus/delivery/http"
	callStatusUseCase "github.com/OkciD/whos_on_call/internal/server/callstatus/usecase"
	deviceHttpDelivery "github.com/OkciD/whos_on_call/internal/server/device/delivery/http"
	deviceRepositorySqlite "github.com/OkciD/whos_on_call/internal/server/device/repository/sqlite"
	deviceUseCase "github.com/OkciD/whos_on_call/internal/server/device/usecase"
	deviceFeatureDelivery "github.com/OkciD/whos_on_call/internal/server/devicefeature/delivery/http"
	deviceFeatureRepositorySqlite "github.com/OkciD/whos_on_call/internal/server/devicefeature/repository/sqlite"
	deviceFeatureUseCase "github.com/OkciD/whos_on_call/internal/server/devicefeature/usecase"
	dbPkg "github.com/OkciD/whos_on_call/internal/server/pkg/db"
	"github.com/OkciD/whos_on_call/internal/server/pkg/db/migrations"
	"github.com/OkciD/whos_on_call/internal/server/pkg/http/middleware"
	"github.com/OkciD/whos_on_call/internal/server/pkg/http/server"
	userHttpDelivery "github.com/OkciD/whos_on_call/internal/server/user/delivery/http"
	userRepositorySqlite "github.com/OkciD/whos_on_call/internal/server/user/repository/sqlite"
	userUseCase "github.com/OkciD/whos_on_call/internal/server/user/usecase"
	webDelivery "github.com/OkciD/whos_on_call/internal/server/web/delivery/html"
)

func main() {
	configFilePathPtr := flag.String("config", "", "path to config file")

	flag.Parse()

	if configFilePathPtr == nil || *configFilePathPtr == "" {
		log.Fatal("-config option required")
	}

	cfg, err := configUtils.ReadConfig[config](*configFilePathPtr)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to read config: %w", err))
	}

	logger := logger.NewLogrusBasedLogger(&cfg.Logger)

	db, err := dbPkg.NewDBConnection(logger, &cfg.DB)
	if err != nil {
		logger.WithError(err).Fatal("db connection failed")
	}
	defer func() {
		logger.Info("closing db connection")
		dbPkg.Close(db, logger)
	}()

	goose.SetBaseFS(migrations.EmbedMigrations)

	if err := goose.SetDialect(cfg.DB.Driver); err != nil {
		logger.WithError(err).Fatal("error setting goose dialect")
	}

	if err := goose.Up(db, "."); err != nil {
		logger.WithError(err).Fatal("failed to apply migrations")
	}
	logger.Info("migrations applied")

	userRepo := userRepositorySqlite.New(logger.ForModule("user_repo"), db)
	deviceRepo := deviceRepositorySqlite.New(logger.ForModule("device_repo"), db)
	deviceFeatureRepo := deviceFeatureRepositorySqlite.New(logger.ForModule("devicefeature_repo"), db)

	userUseCase := userUseCase.New(logger.ForModule("user_usecase"), userRepo)
	deviceUseCase := deviceUseCase.New(logger.ForModule("device_usecase"), deviceRepo)
	deviceFeatureUseCase := deviceFeatureUseCase.New(logger.ForModule("devicefeature_usecase"), deviceRepo, deviceFeatureRepo)
	callStatusUseCase := callStatusUseCase.New(logger.ForModule("callstatus_usecase"), cfg.CallStatus.UseCase, userRepo, deviceRepo, deviceFeatureRepo)

	apiMux := http.NewServeMux()

	userHttpDelivery.New(apiMux, logger.ForModule("user_handler"), userUseCase)
	deviceHttpDelivery.New(apiMux, logger.ForModule("device_handler"), deviceUseCase)
	deviceFeatureDelivery.New(apiMux, logger.ForModule("devicefeature_handler"), deviceFeatureUseCase)
	callStatusDelivery.New(apiMux, logger.ForModule("callstatus_delivery"), callStatusUseCase)

	wrappedApiMux := middleware.ApplyMiddlewares(
		apiMux,
		middleware.NewAuthMiddleware(logger.ForModule("auth_middleware"), userUseCase),
		middleware.NewAccessLogMiddleware(logger),
		middleware.NewRequestIdMiddleware(),
		middleware.NewRecoveryMiddleware(logger),
	)

	apiServer := server.New("api", cfg.ApiServer, logger, wrappedApiMux)
	go func() {
		if err := apiServer.Start(); err != nil {
			logger.WithError(err).Fatal("error starting api server")
		}
	}()

	webMux := http.NewServeMux()

	webDelivery.New(webMux, logger.ForModule("web_delivery"), callStatusUseCase)

	wrappedWebMux := middleware.ApplyMiddlewares(
		webMux,
		middleware.NewAccessLogMiddleware(logger),
		middleware.NewRequestIdMiddleware(),
		middleware.NewRecoveryMiddleware(logger),
	)

	webServer := server.New("web", cfg.WebServer, logger, wrappedWebMux)
	go func() {
		if err := webServer.Start(); err != nil {
			logger.WithError(err).Fatal("error starting web server")
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan

	logger.WithField("signal", sig).Info("received signal")

	if err = apiServer.Stop(); err != nil {
		logger.WithError(err).Fatal("error stopping api server")
	}
	if err = webServer.Stop(); err != nil {
		logger.WithError(err).Fatal("error stopping web server")
	}

	logger.Info("graceful shutdown complete")
}
