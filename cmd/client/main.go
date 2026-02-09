package main

import (
	"flag"
	"fmt"
	"log"

	configUtils "github.com/OkciD/whos_on_call/internal/pkg/config"
	"github.com/OkciD/whos_on_call/internal/pkg/logger"
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

	logger.Info("hello world")
}
