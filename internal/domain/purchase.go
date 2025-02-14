package domain

import "time"

type Purchase struct {
	UserID    UserID
	ProductID ProductID
	Amount    uint64
	Time      time.Time
}
