package service

import (
	"github.com/ariocp/go-app/internal/users/entities"
	"github.com/ariocp/go-app/internal/users/repository"
)

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GenerateToken(name, password string) (string, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Authorization: NewAuthService(repos.Authorization)}
}
