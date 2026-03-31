package main

import (
	"flag"
	"fmt"
	"log"

	deviceRepository "github.com/OkciD/whos_on_call/internal/client/device/repository/api"
	"github.com/OkciD/whos_on_call/internal/client/pkg/httpclient"
	configUtils "github.com/OkciD/whos_on_call/internal/shared/pkg/config"
	"github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
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

	httpClient, err := httpclient.New(logger.ForModule("http_client"), cfg.HttpClient)
	if err != nil {
		log.Fatal(fmt.Errorf("failed to init http client: %w", err))
	}

	_ = deviceRepository.New(logger.ForModule("device_repo"), httpClient)

	logger.Info("hello world")
}
