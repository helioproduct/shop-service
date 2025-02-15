package postgres

import (
	"database/sql"

	trmsql "github.com/avito-tech/go-transaction-manager/drivers/sql/v2"
)

type ProductRepository struct {
	DB       *sql.DB
	txGetter *trmsql.CtxGetter
}

func NewProductRepository(db *sql.DB, txGetter *trmsql.CtxGetter) *ProductRepository {
	return &ProductRepository{
		DB:       db,
		txGetter: txGetter,
	}
}
