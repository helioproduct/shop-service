package domain

import "time"

type Transaction struct {
	From   UserID
	To     UserID
	Amount uint64
	Time   time.Time
}
