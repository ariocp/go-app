package repository

import (
	"github.com/ariocp/go-app/internal/users/entities"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user entities.User) (int, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return 0, err
	}
	return int(user.ID), nil
}

func (r *AuthPostgres) GetUser(username, password string) (entities.User, error) {
	var user entities.User

	err := r.db.Where("username = ? AND password = ?", username, password).First(&user).Error

	return user, err
}
