package postgres

import (
	"database/sql"

	trmsql "github.com/avito-tech/go-transaction-manager/drivers/sql/v2"
)

type UserRepository struct {
	db       *sql.DB
	txGetter *trmsql.CtxGetter
}

func NewUserRepository(db *sql.DB, txGetter *trmsql.CtxGetter) *UserRepository {
	return &UserRepository{
		db:       db,
		txGetter: txGetter,
	}
}
