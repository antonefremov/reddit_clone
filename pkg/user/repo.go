package user

import (
	"crypto/sha256"
	"database/sql"
	"strconv"
)

// Repo represent the Users Repository
type Repo struct {
	DB *sql.DB
}

// NewRepo creates a new repository for Users
func NewRepo(db *sql.DB) *Repo {
	return &Repo{DB: db}
}

// Authorize takes care about authorising a user
func (repo *Repo) Authorize(login, pass string) (*User, error) {
	u, err := repo.GetByUserName(login)
	if err != nil {
		return nil, err
	}

	// u.Token, _ = ljwt.IssueNewToken(u.ID, u.Username)
	calcHash := getSHA256(pass)
	if u.PasswordHash != calcHash {
		return nil, ErrBadPass
	}

	return u, nil
}

// Register creates a new User in the repository when they sign up
func (repo *Repo) Register(user *User) (*User, error) {
	u, _ := repo.GetByUserName(user.Username)
	if u != nil {
		return nil, ErrUserExists
	}

	user.PasswordHash = getSHA256(user.Password)
	user.Token = ""

	uID, err := repo.add(user)
	if err != nil {
		return nil, err
	}

	u, err = repo.GetByID(uID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetByUserName retrieves a User by their User Name (Login) from the database
func (repo *Repo) GetByUserName(login string) (*User, error) {
	user := &User{}
	// QueryRow сам закрывает коннект
	err := repo.DB.
		QueryRow("SELECT id, username, admin, passwordHash FROM users WHERE username = ?", login).
		Scan(&user.ID, &user.Username, &user.Admin, &user.PasswordHash)
	if err != nil {
		return nil, ErrNoUser
	}
	return user, nil
}

// GetByID retrieves a User by their ID from the database
func (repo *Repo) GetByID(ID string) (*User, error) {
	user := &User{}
	err := repo.DB.
		QueryRow("SELECT id, username, admin, passwordHash, token FROM users WHERE id = ?", ID).
		Scan(&user.ID, &user.Username, &user.Admin, &user.PasswordHash, &user.Token)
	if err != nil {
		return nil, ErrNoUser
	}
	return user, nil
}

// add saves a new user into the database
func (repo *Repo) add(user *User) (string, error) {

	result, err := repo.DB.Exec(
		"INSERT INTO users (`username`, `admin`, `passwordHash`, `token`) VALUES (?, ?, ?, ?)",
		user.Username,
		user.Admin,
		user.PasswordHash,
		user.Token,
	)
	if err != nil {
		return "", err
	}
	lastInsertIDAsInt64, err := result.LastInsertId()
	if err != nil {
		return "", err
	}
	retID := strconv.FormatInt(lastInsertIDAsInt64, 10)
	return retID, nil
}

func getSHA256(value string) string {
	sha256Instance := sha256.New()
	sha256Instance.Write([]byte(value))
	return string(sha256Instance.Sum(nil))
}
