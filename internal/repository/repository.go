package repository

import "github.com/jmoiron/sqlx"

type UserRepository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		Authorization: NewAuthPostgres(db),
	}
}
