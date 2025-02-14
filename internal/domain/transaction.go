package domain

import "time"

type TransactionID string

type Transaction struct {
	ID     TransactionID
	From   UserID
	To     UserID
	Amount uint64
	Time   time.Time
}
