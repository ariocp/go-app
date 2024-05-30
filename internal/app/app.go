package app

import (
	"github.com/ariocp/go-app/config"
	"github.com/ariocp/go-app/internal/controller/http/v1"
	"github.com/ariocp/go-app/internal/repository"
	"github.com/ariocp/go-app/internal/service"
	"github.com/ariocp/go-app/pkg/database"
	"github.com/ariocp/go-app/pkg/server"
	_ "github.com/lib/pq"
	"os"
)

func Run(cfg config.Config) {
	dbConfig := cfg.Database
	dbConfig.Password = os.Getenv("DB_PASSWORD")

	db, err := database.NewPostgresDB(dbConfig)
	if err != nil {
		panic(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := v1.NewHandler(services)
	router := handler.Routes()
	srv := server.NewServer(cfg, router)

	if err = srv.Start(); err != nil {
		panic(err)
	}
}
