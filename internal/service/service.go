package service

import (
	"github.com/ariocp/go-app/internal/models"
	"github.com/ariocp/go-app/internal/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int64, error)
	GenerateToken(username, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
