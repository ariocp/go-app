package app

import (
	"github.com/ariocp/go-app/config"
	"github.com/ariocp/go-app/internal/database"
	"github.com/ariocp/go-app/internal/delivery/http/v1"
	"github.com/ariocp/go-app/internal/repository"
	"github.com/ariocp/go-app/internal/service"
	_ "github.com/lib/pq"
	"os"
)

func Run(cfg config.Config) {
	dbConfig := cfg.Database
	dbConfig.Password = os.Getenv("DB_PWD")

	db, err := database.NewPostgresDB(dbConfig)
	if err != nil {
		panic(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := v1.NewHandler(services)
	router := handler.Routes()
	srv := v1.NewServer(cfg, router)

	if err = srv.Start(); err != nil {
		panic(err)
	}
}
