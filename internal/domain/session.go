package domain

import "time"

type Session struct {
	UserID    UserID
	Token     string
	IssuedAt  time.Time
	ExpiresAt time.Time
}
