package main

import (
	"flag"
	"net/http"

	"github.com/OkciD/whos_on_call/cmd/whos_on_call/config"

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
		// TODO: no panic
		panic("failed to parse config path")
	}

	config, err := configUtils.ReadConfig[config.Config](*configFilePathPtr)
	if err != nil {
		// TODO: no panic
		panic(err)
	}

	userRepo, err := userRepositoryInmemory.New(&config.User.Repository)
	if err != nil {
		// TODO: no panic
		panic(err)
	}

	userUseCase := userUseCase.New(userRepo)

	mux := http.NewServeMux()

	userHttpDelivery.New(mux, userUseCase)

	contentTypeMiddleware := middleware.NewContentTypeMiddleware("application/json")
	authMiddleware := middleware.NewAuthMiddleware(userUseCase)

	wrappedMux := contentTypeMiddleware(authMiddleware(mux))

	err = http.ListenAndServe(config.Server.ListenAddr, wrappedMux)
	if err != nil {
		// TODO: no panic
		panic(err)
	}
}
