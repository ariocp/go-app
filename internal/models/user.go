package models

type User struct {
	Id       int64  `json:"id" db:"id"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
