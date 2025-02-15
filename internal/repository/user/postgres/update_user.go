package postgres

import (
	"context"
	"fmt"
	"shop-service/internal/domain"

	sq "github.com/Masterminds/squirrel"

	"github.com/rs/zerolog/log"
)

type UpdateUserRequest struct {
	UserID         domain.UserID
	Username       *string
	HashedPassword *string
	Balance        *uint64
}

func (repo *UserRepository) UpdateUser(ctx context.Context, req *UpdateUserRequest) error {
	caller := "UserRepository.UpdateUser"

	queryBuilder := sq.Update("users").
		Set("updated_at", sq.Expr("NOW()")).
		Where(sq.Eq{"id": req.UserID}).
		PlaceholderFormat(sq.Dollar)

	if req.Username != nil {
		queryBuilder = queryBuilder.Set("username", *req.Username)
	}
	if req.HashedPassword != nil {
		queryBuilder = queryBuilder.Set("hashed_password", *req.HashedPassword)
	}
	if req.Balance != nil {
		queryBuilder = queryBuilder.Set("balance", *req.Balance)
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		err = fmt.Errorf("failed to build UpdateUser query: %w", err)
		log.Err(err).Str("caller", caller).Send()
		return err
	}

	trOrDB := repo.txGetter.DefaultTrOrDB(ctx, repo.db)

	result, err := trOrDB.ExecContext(ctx, query, args...)
	if err != nil {
		err = fmt.Errorf("failed to execute UpdateUser query: %w", err)
		log.Err(err).Str("caller", caller).Send()
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		err = fmt.Errorf("failed to get affected rows: %w", err)
		log.Err(err).Str("caller", caller).Send()
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrUserNotFound
	}

	return nil
}
