package models

import "database/sql"

type Session struct {
	ID     int
	UserID int
	// Token is only set when creating a new session. When looking up a session
	// token will be left empty as we only store the hash of a session token
	// in our database and we cannot reverse it into a raw token.
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(userID int) (*Session, error) {
	// 1. Create session token
	return nil, nil
}

func (ss *SessionService) User(token string) (*User, error) {
	return nil, nil

}
