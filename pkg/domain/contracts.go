package domain

import (
	"time"
)

type Store interface {
	CreateUser(usr *User) (*User, error)
	DeleteUserByID(id int64) error
	DeleteAllUsers() error
	FindUserByEmail(email string) (*User, error)
	FindUserByID(id int64) (*User, error)
}

type JWT interface {
	CreateToken(user User, duration time.Duration) (*JWTPayload, error)
	VerifyToken(tokn string) (*JWTPayload, error)
}
