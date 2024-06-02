package models

import "time"

type User struct {
	Id                 int64     `json:"-" db:"id"`
	Username           string    `json:"username" binding:"required"`
	Password           string    `json:"password" binding:"required"`
	Email              string    `json:"email" binding:"required"`
	ConfirmationCode   string    `json:"-" db:"confirmation_code"`
	IsConfirmed        bool      `json:"-" db:"is_confirmed"`
	ConfirmationExpiry time.Time `json:"-" db:"confirmation_expiry"`
}
