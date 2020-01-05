package main

import (
	"flag"

	"go_api_template/pkg/api"
)

func main() {
	flag.Parse()

	if err := ParseEnvConfig(); err != nil {
		panic(err)
	}

	api.ServeAPI(&api.Config{
		BindAddr: cfg.BindAddr,
	})
}
