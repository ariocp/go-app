package database

import (
	"fmt"

	"github.com/ariocp/go-app/config"
	"github.com/ariocp/go-app/internal/users/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&entities.User{}); err != nil {
		return nil, err
	}

	return db, nil
}
