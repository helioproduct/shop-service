package purchase

import (
	"context"
	"fmt"
	"shop-service/internal/domain"

	sq "github.com/Masterminds/squirrel"
)

type PurchaseSummary struct {
	Product domain.Product
	Amount  uint64
}

type PurchaseSummaryRequest struct {
	UserID domain.UserID
	Limit  uint64
	Offset uint64
}

func (r *PurchaseRepository) GetPurchaseSummary(ctx context.Context, req PurchaseSummaryRequest) ([]*PurchaseSummary, error) {
	// caller := "PurchaseRepository.GetPurchaseSummary"

	queryBuilder := sq.Select(
		"pr.id AS product_id",
		"pr.name AS product_name",
		"pr.price",
		"SUM(p.quantity) AS amount",
	).
		From("purchases p").
		Join("products pr ON p.product_id = pr.id").
		Where(sq.Eq{"p.user_id": req.UserID}).
		GroupBy("pr.id", "pr.name", "pr.price").
		OrderBy("amount DESC").
		Limit(req.Limit).
		Offset(req.Offset).
		PlaceholderFormat(sq.Dollar)

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build GetPurchaseSummary query: %w", err)
	}

	trOrDB := r.txGetter.DefaultTrOrDB(ctx, r.db)
	rows, err := trOrDB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute GetPurchaseSummary query: %w", err)
	}
	defer rows.Close()

	var summaries []*PurchaseSummary
	for rows.Next() {
		summary := new(PurchaseSummary)
		// var summary *PurchaseSummary
		if err := rows.Scan(
			&summary.Product.ID,
			&summary.Product.Name,
			&summary.Product.Price,
			&summary.Amount,
		); err != nil {
			return nil, fmt.Errorf("failed to scan GetPurchaseSummary result: %w", err)
		}
		summaries = append(summaries, summary)
	}

	return summaries, nil
}
