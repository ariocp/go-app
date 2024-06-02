package repository

import (
	"fmt"
	"github.com/ariocp/go-app/internal/models"
	"github.com/jmoiron/sqlx"
)

const usersTable = "users"

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int64, error) {
	var id int64

	query := fmt.Sprintf("INSERT INTO %s (username, password, email, confirmation_code, is_confirmed, confirmation_expiry) values ($1, $2, $3, $4, $5, $6) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Username, user.Password, user.Email, user.ConfirmationCode, user.IsConfirmed, user.ConfirmationExpiry)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) ConfirmUser(username, code string) error {
	query := fmt.Sprintf("UPDATE %s SET is_confirmed = TRUE WHERE username=$1 AND confirmation_code=$2 AND confirmation_expiry > NOW()", usersTable)
	_, err := r.db.Exec(query, username, code)
	return err
}

func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password=$2", usersTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}

func (r *AuthPostgres) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE username=$1", usersTable)
	err := r.db.Get(&user, query, username)
	return user, err
}
