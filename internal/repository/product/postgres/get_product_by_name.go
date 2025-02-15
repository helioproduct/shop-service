package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"shop-service/internal/domain"

	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog/log"
)

// GetProductByName получает товар по имени
func (repo *ProductRepository) GetProductByName(ctx context.Context, name string) (*domain.Product, error) {
	caller := "ProductRepository.GetProductByName"

	query, args, err := sq.Select("id", "name", "price").
		From("products").
		Where(sq.Eq{"name": name}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		err = fmt.Errorf("failed to build GetProductByName query: %w", err)
		log.Err(err).Str("caller", caller).Send()
		return nil, err
	}

	var product domain.Product
	txOrDB := repo.txGetter.DefaultTrOrDB(ctx, repo.DB)
	if err := txOrDB.QueryRowContext(ctx, query, args...).
		Scan(&product.ID, &product.Name, &product.Price); err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrProductNotFound
		}
		err = fmt.Errorf("failed to execute GetProductByName query: %w", err)
		log.Err(err).Str("caller", caller).Send()
		return nil, err
	}

	return &product, nil
}
