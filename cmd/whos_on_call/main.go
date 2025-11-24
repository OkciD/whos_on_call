package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/OkciD/whos_on_call/cmd/whos_on_call/config"
	"github.com/sirupsen/logrus"

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

	logger := logrus.WithFields(logrus.Fields{
		"host": hostname,
	})

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

	userRepo, err := userRepositoryInmemory.New(&cfg.User.Repository)
	if err != nil {
		logrus.WithError(err).Fatalf("failed to create user repository")
	}

	userUseCase := userUseCase.New(userRepo)

	mux := http.NewServeMux()

	userHttpDelivery.New(mux, userUseCase)

	contentTypeMiddleware := middleware.NewContentTypeMiddleware("application/json")
	requestIdMiddleWare := middleware.NewRequestIdMiddleware()
	authMiddleware := middleware.NewAuthMiddleware(userUseCase)

	wrappedMux := contentTypeMiddleware(requestIdMiddleWare(authMiddleware(mux)))

	logrus.WithField("addr", cfg.Server.ListenAddr).Info("http server starting")
	err = http.ListenAndServe(cfg.Server.ListenAddr, wrappedMux)
	if err != nil {
		logrus.WithError(err).Fatalf("http server error")
	}
}
