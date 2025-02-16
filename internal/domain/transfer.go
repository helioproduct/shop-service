package domain

import "time"

type TransferID string

type Transfer struct {
	ID           TransferID
	From         UserID
	To           UserID
	Amount       uint64
	Time         time.Time
	FromUsername string
	ToUsername   string
}
