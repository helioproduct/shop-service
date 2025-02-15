package transfer

import (
	"context"
	"fmt"
	"shop-service/internal/domain"

	sq "github.com/Masterminds/squirrel"

	"github.com/rs/zerolog/log"
)

type GetTransfersFilter struct {
	FromUserID *domain.UserID
	ToUserID   *domain.UserID
	Limit      uint64
	Offset     uint64
}

func (repo *TransferRepository) GetTransfers(ctx context.Context, filter *GetTransfersFilter) ([]domain.Transfer, error) {
	caller := "TransferRepository.GetTransfers"

	queryBuilder := sq.Select("id", "from_user_id", "to_user_id", "amount", "created_at").
		From("transfers").
		PlaceholderFormat(sq.Dollar)

	if filter.FromUserID != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"from_user_id": *filter.FromUserID})
	}
	if filter.ToUserID != nil {
		queryBuilder = queryBuilder.Where(sq.Eq{"to_user_id": *filter.ToUserID})
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
		log.Err(err).Str("caller", caller).Send()
		return nil, err
	}

	trOrDB := repo.txGetter.DefaultTrOrDB(ctx, repo.db)
	rows, err := trOrDB.QueryContext(ctx, query, args...)
	if err != nil {
		err = fmt.Errorf("failed to execute GetTransfers query: %w", err)
		log.Err(err).Str("caller", caller).Send()
		return nil, err
	}
	defer rows.Close()

	var transfers []domain.Transfer
	for rows.Next() {
		var t domain.Transfer
		if err := rows.Scan(&t.ID, &t.From, &t.To, &t.Amount, &t.Time); err != nil {
			err = fmt.Errorf("failed to scan GetTransfers result: %w", err)
			log.Err(err).Str("caller", caller).Send()
			return nil, err
		}
		transfers = append(transfers, t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error in GetTransfers: %w", err)
	}

	return transfers, nil
}
