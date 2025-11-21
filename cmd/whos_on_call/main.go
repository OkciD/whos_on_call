package main

import (
	"flag"
	"fmt"

	"github.com/OkciD/whos_on_call/cmd/whos_on_call/config"
)

func main() {
	configFilePathPtr := flag.String("config", "", "path to config file")

	flag.Parse()

	if configFilePathPtr == nil || *configFilePathPtr == "" {
		// TODO: no panic
		panic("failed to parse config path")
	}

	config, err := config.ReadConfig(*configFilePathPtr)
	if err != nil {
		// TODO: no panic
		panic(err)
	}

	fmt.Printf("About to listen on addr %s", config.ListenAddr)
}
