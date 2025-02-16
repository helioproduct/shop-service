package purchase

import (
	"database/sql"

	trmsql "github.com/avito-tech/go-transaction-manager/drivers/sql/v2"
)

type PurchaseRepository struct {
	db       *sql.DB
	txGetter *trmsql.CtxGetter
}

func NewPurchaseRepository(db *sql.DB, txGetter *trmsql.CtxGetter) *PurchaseRepository {
	return &PurchaseRepository{
		db:       db,
		txGetter: txGetter,
	}
}
