package service

import (
	"github.com/ariocp/go-app/internal/repository"
)

type UserService struct {
	Authorization
}

func NewService(repos *repository.UserRepository) *UserService {
	return &UserService{
		Authorization: NewAuthService(repos.Authorization),
	}
}
