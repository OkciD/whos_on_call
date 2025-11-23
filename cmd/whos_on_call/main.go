package main

import (
	"flag"

	"github.com/OkciD/whos_on_call/cmd/whos_on_call/config"

	UserRepositoryInmemory "github.com/OkciD/whos_on_call/internal/app/user/repository/inmemory"
	UserUseCase "github.com/OkciD/whos_on_call/internal/app/user/usecase"
	ConfigUtils "github.com/OkciD/whos_on_call/internal/pkg/config"
)

func main() {
	configFilePathPtr := flag.String("config", "", "path to config file")

	flag.Parse()

	if configFilePathPtr == nil || *configFilePathPtr == "" {
		// TODO: no panic
		panic("failed to parse config path")
	}

	config, err := ConfigUtils.ReadConfig[config.Config](*configFilePathPtr)
	if err != nil {
		// TODO: no panic
		panic(err)
	}

	userRepo, err := UserRepositoryInmemory.New(&config.User.Repository)
	if err != nil {
		// TODO: no panic
		panic(err)
	}

	_ = UserUseCase.New(userRepo)
}
