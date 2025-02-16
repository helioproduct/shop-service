package transfer

import (
	"context"
	"fmt"
	"shop-service/internal/domain"
	"shop-service/pkg/logger"

	sq "github.com/Masterminds/squirrel"
)

func (r *TransferRepository) CreateTransfer(ctx context.Context, transfer domain.Transfer) (*domain.Transfer, error) {
	caller := "TransferRepository.CreateTransfer"

	query, args, err := sq.Insert("transfers").
		Columns("from_user_id", "to_user_id", "amount").
		Values(transfer.From, transfer.To, transfer.Amount).
		Suffix("RETURNING id, created_at").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		err = fmt.Errorf("failed to build CreateTransfer query: %w", err)
		logger.Error(err, caller)
		return nil, err
	}

	var createdTransfer domain.Transfer
	trOrDB := r.txGetter.DefaultTrOrDB(ctx, r.db)
	if err := trOrDB.QueryRowContext(ctx, query, args...).
		Scan(&createdTransfer.ID, &createdTransfer.Time); err != nil {
		err = fmt.Errorf("failed to execute CreateTransfer query: %w", err)
		logger.Error(err, caller)
		return nil, err
	}

	createdTransfer.From = transfer.From
	createdTransfer.To = transfer.To
	createdTransfer.Amount = transfer.Amount

	return &createdTransfer, nil
}
