package postgres

import (
	"database/sql"

	trmsql "github.com/avito-tech/go-transaction-manager/drivers/sql/v2"
)

type TransferRepository struct {
	db       *sql.DB
	txGetter *trmsql.CtxGetter
}

func NewTransferRepository(db *sql.DB, txGetter *trmsql.CtxGetter) *TransferRepository {
	return &TransferRepository{
		db:       db,
		txGetter: txGetter,
	}
}
