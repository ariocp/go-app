package repository

import "github.com/ariocp/go-app/internal/models"

type Authorization interface {
	CreateUser(user models.User) (int64, error)
	GetUser(username, password string) (models.User, error)
	ConfirmUser(username, code string) error
	GetUserByUsername(username string) (models.User, error)
}
