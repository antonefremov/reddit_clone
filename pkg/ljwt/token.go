package ljwt

import (
	"github.com/dgrijalva/jwt-go"
)

// TokenUser is used in the Claims struct below
type TokenUser struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

// Claims is the data structure inside of the JWT token
type Claims struct {
	User      TokenUser `json:"user"`
	SessionID string    `json:"sessionId"`
	jwt.StandardClaims
}

// ServerKey is used to form the server side signature
// TODO: use env variable instead
var ServerKey = "my_super_secret_key"
