package product

import (
	"context"
	"database/sql"
	"fmt"
	"shop-service/internal/domain"
	"shop-service/pkg/logger"

	sq "github.com/Masterminds/squirrel"
)

func (repo *ProductRepository) GetProductByName(ctx context.Context, name string) (*domain.Product, error) {
	caller := "ProductRepository.GetProductByName"

	query, args, err := sq.Select("id", "name", "price").
		From("products").
		Where(sq.Eq{"name": name}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		err = fmt.Errorf("failed to build GetProductByName query: %w", err)
		logger.Error(err, caller)
		return nil, domain.ErrInternalError
	}

	var product domain.Product
	txOrDB := repo.txGetter.DefaultTrOrDB(ctx, repo.DB)
	if err := txOrDB.QueryRowContext(ctx, query, args...).
		Scan(&product.ID, &product.Name, &product.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrProductNotFound
		}
		err = fmt.Errorf("failed to execute GetProductByName query: %w", err)
		logger.Error(err, caller)
		return nil, domain.ErrInternalError
	}

	return &product, nil
}
