package main

import (
	"github.com/ariocp/go-app/config"
	"github.com/ariocp/go-app/internal/app"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
