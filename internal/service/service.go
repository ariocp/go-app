package service

import (
	"github.com/ariocp/go-app/internal/repository"
)

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
