package repository

import (
	"github.com/jmoiron/sqlx"
)

<<<<<<< HEAD
type UserRepository struct {
=======
type Repository struct {
>>>>>>> 39260269a65d547ef035ec84c6d4c737c0756251
	Authorization
}

func NewRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		Authorization: NewAuthPostgres(db),
	}
}
