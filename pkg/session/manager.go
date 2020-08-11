package session

import (
	"database/sql"
	"time"
)

// SessionsManager manages sessions
type SessionsManager struct {
	// data map[string]*Session
	DB *sql.DB
}

// NewSessionsManager constructs a new Sess Man
func NewSessionsManager(db *sql.DB) *SessionsManager {
	return &SessionsManager{
		DB: db,
	}
}

// Check validates a session by its id
func (sm *SessionsManager) Check(sessionID string) (*Session, error) {
	sess := &Session{}
	var expiresUnix int64
	err := sm.DB.
		QueryRow("SELECT id, userId, expires FROM sessions WHERE id = ?", sessionID).
		Scan(&sess.ID, &sess.UserID, &expiresUnix)
	if err != nil {
		return nil, ErrNoAuth
	}

	sess.Expires = time.Unix(expiresUnix, 0)
	if sess.Expires.Unix() < time.Now().Unix() {
		return nil, ErrNoAuth
	}

	return sess, nil
}

// Create creates a new session for the passed userID
func (sm *SessionsManager) Create(userID string) (*Session, error) {
	sess, err := NewSession(userID)
	if err != nil {
		return nil, err
	}
	_, err = sm.DB.Exec(
		"INSERT INTO sessions (`id`, `userId`, `expires`) VALUES (?, ?, ?)",
		sess.ID,
		sess.UserID,
		sess.Expires.Unix(),
	)
	if err != nil {
		return nil, err
	}

	return sess, nil
}
