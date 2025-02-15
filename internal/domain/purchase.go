package domain

import "time"

type PurchaseID string

type Purchase struct {
	ID        PurchaseID
	UserID    UserID
	ProductID ProductID
	Amount    uint64
	Time      time.Time
}
