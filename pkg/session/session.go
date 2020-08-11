package session

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"
)

type Session struct {
	ID      string
	UserID  string
	Expires time.Time
}

const IDLength int = 32

func NewSession(userID string) (*Session, error) {
	id, err := generateRandomString(IDLength)
	if err != nil {
		return nil, err
	}

	return &Session{
		ID:      id,
		UserID:  userID,
		Expires: time.Now().Add(time.Hour),
	}, nil
}

var (
	ErrNoAuth = errors.New("No session found")
)

type sessKey string

var SessionKey sessKey = "sessionKey"

func SessionFromContext(ctx context.Context) (*Session, error) {
	sess, ok := ctx.Value(SessionKey).(*Session)
	if !ok || sess == nil {
		return nil, ErrNoAuth
	}
	return sess, nil
}

func generateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	return base64.URLEncoding.EncodeToString(bytes), err
}
