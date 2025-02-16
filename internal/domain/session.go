package domain

import "time"

type Session struct {
	UserID    UserID
	Username  string
	Token     string
	IssuedAt  time.Time
	ExpiresAt time.Time
}
