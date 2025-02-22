package purchase

import (
	"context"
	"fmt"
	"shop-service/internal/domain"
	"shop-service/pkg/logger"

	sq "github.com/Masterminds/squirrel"
)

type CreatePurchaseRequest struct {
	UserID    domain.UserID
	ProductID domain.ProductID
	Quantity  uint64
}

func (r *PurchaseRepository) CreatePurchase(ctx context.Context, req CreatePurchaseRequest) (*domain.Purchase, error) {
	caller := "PurchaseRepository.CreatePurchase"

	query, args, err := sq.Insert("purchases").
		Columns("user_id", "product_id", "quantity").
		Values(req.UserID, req.ProductID, req.Quantity).
		Suffix("RETURNING id, user_id, product_id, quantity, created_at").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		err = fmt.Errorf("failed to build CreatePurchase query: %w", err)
		logger.Error(err, caller)
		return nil, domain.ErrInternalError
	}

	var purchase domain.Purchase
	trOrDB := r.txGetter.DefaultTrOrDB(ctx, r.db)
	if err := trOrDB.QueryRowContext(ctx, query, args...).Scan(
		&purchase.ID,
		&purchase.UserID,
		&purchase.ProductID,
		&purchase.Amount,
		&purchase.Time,
	); err != nil {
		err = fmt.Errorf("failed to execute CreatePurchase query: %w", err)
		logger.Error(err, caller)
		return nil, domain.ErrInternalError
	}

	return &purchase, nil
}
