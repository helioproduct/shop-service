package transfer

import (
	"context"
	"fmt"
	"shop-service/internal/domain"
	"shop-service/pkg/logger"

	sq "github.com/Masterminds/squirrel"
)

type TotalTransfer struct {
	FromUsername string
	ToUsername   string
	Amount       uint64
}

type GetTransfersFilter struct {
	FromUsername *string
	ToUsername   *string
	Limit        uint64
	Offset       uint64
}

func (r *TransferRepository) GetTransfers(ctx context.Context, filter GetTransfersFilter) ([]domain.Transfer, error) {
	caller := "TransferRepository.GetTransfers"

	queryBuilder := sq.Select(
		"t.id",
		"t.from_user_id",
		"t.to_user_id",
		"t.amount",
		"t.created_at",
		"fu.username AS from_username",
		"tu.username AS to_username").
		From("transfers t").
		LeftJoin("users fu ON t.from_user_id = fu.id").
		LeftJoin("users tu ON t.to_user_id = tu.id").
		PlaceholderFormat(sq.Dollar)

	if filter.FromUsername != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"fu.username": *filter.FromUsername})
	}

	if filter.ToUsername != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"tu.username": *filter.ToUsername})
	}

	if filter.Limit > 0 {
		queryBuilder = queryBuilder.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		queryBuilder = queryBuilder.Offset(filter.Offset)
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		err = fmt.Errorf("failed to build GetTransfers query: %w", err)
		logger.Error(err, caller)
		return nil, domain.ErrInternalError
	}

	trOrDB := r.txGetter.DefaultTrOrDB(ctx, r.db)
	rows, err := trOrDB.QueryContext(ctx, query, args...)
	if err != nil {
		err = fmt.Errorf("failed to execute GetTransfers query: %w", err)
		logger.Error(err, caller)
		return nil, domain.ErrInternalError
	}
	defer rows.Close()

	var transfers []domain.Transfer
	for rows.Next() {
		var t domain.Transfer
		var fromUsername, toUsername string
		if err := rows.Scan(
			&t.ID, &t.From, &t.To, &t.Amount, &t.Time,
			&fromUsername, &toUsername,
		); err != nil {
			err = fmt.Errorf("failed to scan GetTransfers result: %w", err)
			logger.Error(err, caller)
			return nil, domain.ErrInternalError
		}
		t.FromUsername = fromUsername
		t.ToUsername = toUsername
		transfers = append(transfers, t)
	}

	if err := rows.Err(); err != nil {
		logger.Error(err, caller)
		return nil, domain.ErrInternalError
	}

	return transfers, nil
}
