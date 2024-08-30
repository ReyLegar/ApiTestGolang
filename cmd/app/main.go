package main

import (
	"ApiTest/config"
	"ApiTest/internal/app"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
