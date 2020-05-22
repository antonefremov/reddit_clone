package users

import (
	"crypto/sha256"
	"errors"
	"golang-stepik-2020q1/5/99_hw/redditclone/pkg/ids"
	"golang-stepik-2020q1/5/99_hw/redditclone/pkg/ljwt"
)

// UsersRepo represent the Users Repository
type UsersRepo struct {
	data map[string]*User
}

// NewUsersRepo creates a repository of users on app start
func NewUsersRepo() *UsersRepo {
	hash := getSHA256("testtest")
	return &UsersRepo{
		data: map[string]*User{
			"test123456": &User{
				Username:     "test123456",
				Admin:        true,
				ID:           ids.GenerateID(),
				PasswordHash: hash,
			},
		},
	}
}

var (
	ErrNoUser        = errors.New("No user found")
	ErrBadPass       = errors.New("Invald password")
	ErrInternalError = errors.New("Internal server error")
	ErrUserExists    = errors.New("User name already exists")
)

// Authorize takes care about authorising a user
func (repo *UsersRepo) Authorize(login, pass string) (*User, error) {
	u, ok := repo.data[login]
	if !ok {
		return nil, ErrNoUser
	}

	u.Token, _ = ljwt.IssueNewToken(u.ID, u.Username)
	calcHash := getSHA256(pass)
	if u.PasswordHash != calcHash {
		return nil, ErrBadPass
	}

	return u, nil
}

// Register creates a new User in the repository when they sign up
func (repo *UsersRepo) Register(user *User) (*User, error) {
	passwordHash := getSHA256(user.Password)
	user.PasswordHash = passwordHash

	if _, userExists := repo.data[user.Username]; userExists {
		return nil, ErrUserExists
	}

	repo.data[user.Username] = user

	// clear the password for the returned User entity
	user.Password = ""
	return user, nil
}

func getSHA256(value string) string {
	sha256Instance := sha256.New()
	sha256Instance.Write([]byte(value))
	return string(sha256Instance.Sum(nil))
}
