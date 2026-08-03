package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/OkciD/whos_on_call/internal/client/pkg/apiclient"
	"github.com/OkciD/whos_on_call/internal/client/pkg/apiclient/gen"
	"github.com/OkciD/whos_on_call/internal/client/pkg/httpclient"
	configUtils "github.com/OkciD/whos_on_call/internal/shared/pkg/config"
	loggerPkg "github.com/OkciD/whos_on_call/internal/shared/pkg/logger"
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

	logger := loggerPkg.NewLogrusBasedLogger(&cfg.Logger)

	hc := httpclient.New(logger.ForModule("http_client"), cfg.ApiClient.HTTPClientConfig)
	c, err := gen.NewClientWithResponses(
		cfg.ApiClient.BaseURL,
		gen.WithHTTPClient(&hc),
		gen.WithRequestEditorFn(apiclient.NewAuthRequestEditor(cfg.ApiClient.ApiKey)),
	)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.GetUserWithResponse(context.TODO())
	if err != nil {
		logger.WithError(err).Fatal("failed to get user")
	}
	logger.WithFields(loggerPkg.Fields{
		"id":   resp.JSON200.Id,
		"name": resp.JSON200.Name,
	}).Info("got user")
}
