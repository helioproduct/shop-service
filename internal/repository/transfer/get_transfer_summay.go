package transfer

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)

type SentCoinsSummary struct {
	ToUsername string
	TotalSent  uint64
}

type ReceivedCoinsSummary struct {
	FromUsername  string
	TotalReceived uint64
}

func (r *TransferRepository) GetSentCoinsSummary(ctx context.Context, fromUsername string) ([]SentCoinsSummary, error) {
	// caller := "TransferRepository.GetSentCoinsSummary"

	queryBuilder := sq.Select("tu.username AS to_username", "SUM(t.amount) AS total_sent").
		From("transfers t").
		Join("users fu ON t.from_user_id = fu.id").
		Join("users tu ON t.to_user_id = tu.id").
		Where(sq.Eq{"fu.username": fromUsername}).
		GroupBy("tu.username").
		PlaceholderFormat(sq.Dollar)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build GetSentCoinsSummary query: %w", err)
	}

	trOrDB := r.txGetter.DefaultTrOrDB(ctx, r.db)
	rows, err := trOrDB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute GetSentCoinsSummary query: %w", err)
	}
	defer rows.Close()

	var summaries []SentCoinsSummary
	for rows.Next() {
		var summary SentCoinsSummary
		if err := rows.Scan(&summary.ToUsername, &summary.TotalSent); err != nil {
			return nil, fmt.Errorf("failed to scan GetSentCoinsSummary result: %w", err)
		}
		summaries = append(summaries, summary)
	}

	return summaries, nil
}

func (r *TransferRepository) GetReceivedCoinsSummary(ctx context.Context, toUsername string) ([]ReceivedCoinsSummary, error) {
	// caller := "TransferRepository.GetReceivedCoinsSummary"

	queryBuilder := sq.Select("fu.username AS from_username", "SUM(t.amount) AS total_received").
		From("transfers t").
		Join("users fu ON t.from_user_id = fu.id").
		Join("users tu ON t.to_user_id = tu.id").
		Where(sq.Eq{"tu.username": toUsername}).
		GroupBy("fu.username").
		PlaceholderFormat(sq.Dollar)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build GetReceivedCoinsSummary query: %w", err)
	}

	trOrDB := r.txGetter.DefaultTrOrDB(ctx, r.db)
	rows, err := trOrDB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute GetReceivedCoinsSummary query: %w", err)
	}
	defer rows.Close()

	var summaries []ReceivedCoinsSummary
	for rows.Next() {
		var summary ReceivedCoinsSummary
		if err := rows.Scan(&summary.FromUsername, &summary.TotalReceived); err != nil {
			return nil, fmt.Errorf("failed to scan GetReceivedCoinsSummary result: %w", err)
		}
		summaries = append(summaries, summary)
	}

	return summaries, nil
}
