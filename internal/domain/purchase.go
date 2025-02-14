package domain

import "time"

type Purchase struct {
	UserID    UserID
	ProductID ProductID
	Time      time.Time
}
