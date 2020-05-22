package users

// User entity representation
type User struct {
	Username     string `json:"username"`
	Admin        bool   `json:"amdin"`
	ID           string `json:"id"`
	PasswordHash string `json:"passwordHash"`
	Password     string `json:"password"`
	Token        string `json:"token"`
}
