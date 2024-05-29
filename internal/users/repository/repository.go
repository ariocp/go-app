package repository

import (
	"github.com/ariocp/go-app/internal/users/entities"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GetUser(username, password string) (entities.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{Authorization: NewAuthPostgres(db)}
}
