package service

import "github.com/ariocp/go-app/internal/models"

type Authorization interface {
	CreateUser(user models.User) (int64, error)
	GenerateToken(username, password string) (string, error)
	ConfirmUser(username, code string) error
}
