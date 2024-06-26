package main

import (
	"github.com/ariocp/go-app/config"
	"github.com/ariocp/go-app/internal/app"
)

// @title go app
// @version 1.0

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apiKey apiAuthKey
// @in header
// @name authorization

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	app.Run(cfg)
}
