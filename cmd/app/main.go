package main

import (
	"os"

	"github.com/ariocp/go-app/pkg/database"
	"github.com/ariocp/go-app/pkg/server"

	"github.com/ariocp/go-app/config"
	v1 "github.com/ariocp/go-app/internal/users/delivery/http/v1"
	"github.com/ariocp/go-app/internal/users/repository"
	"github.com/ariocp/go-app/internal/users/service"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("%s", err)
	}

	dbConfig := cfg.Database
	dbConfig.Password = os.Getenv("DB_PWD")

	db, err := database.NewDatabase(dbConfig)
	if err != nil {
		logrus.Fatalf("%s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handler := v1.NewHandler(services)
	router := handler.Routes()
	srv := server.NewServer(cfg, router)

	if err = srv.Run(); err != nil {
		logrus.Fatalf("%s", err)
	}
}
