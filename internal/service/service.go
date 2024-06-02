package service

import (
	"github.com/ariocp/go-app/internal/repository"
)

<<<<<<< HEAD
type UserService struct {
=======
type Service struct {
>>>>>>> 39260269a65d547ef035ec84c6d4c737c0756251
	Authorization
}

func NewService(repos *repository.UserRepository) *UserService {
	return &UserService{
		Authorization: NewAuthService(repos.Authorization),
	}
}
