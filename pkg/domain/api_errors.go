package domain

import "errors"

var (
	ErrUserNotFound = errors.New("No user found.")
	ErrExpiredToken = errors.New("jwt token expired")
	ErrInvalidToken = errors.New("invalid token")
)
