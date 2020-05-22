package ljwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	// Create the JWT key used to create the signature
	jwtKey = []byte(ServerKey)
)

// IssueNewToken returns a new JWT for a given user
func IssueNewToken(userID, username string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)
	tokenUser := TokenUser{
		Username: username,
		ID:       userID,
	}
	claims := &Claims{
		User: tokenUser,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("Internal server error")
	}
	return tokenString, nil
}
