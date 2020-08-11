package user

import "errors"

// User entity representation
type User struct {
	Username     string `json:"username"`
	Admin        bool   `json:"amdin"`
	ID           string `json:"id"`
	PasswordHash string `json:"passwordHash"`
	Password     string `json:"password"`
	Token        string `json:"token"`
}

var (
	ErrNoUser        = errors.New("No user found")
	ErrBadPass       = errors.New("Invald password")
	ErrInternalError = errors.New("Internal server error")
	ErrUserExists    = errors.New("User name already exists")
)
