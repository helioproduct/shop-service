package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"shop-service/internal/domain"
	"shop-service/pkg/logger"

	sq "github.com/Masterminds/squirrel"
)

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	caller := "UserRepository.GetUserByUsername"

	query, args, err := sq.Select("id", "username", "balance").
		From("users").
		Where(sq.Eq{"username": username}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		err = fmt.Errorf("failed to build GetUserByUsername query: %w", err)
		logger.Error(err, caller)
		return nil, err
	}

	var user domain.User
	trOrDB := r.txGetter.DefaultTrOrDB(ctx, r.db)
	if err := trOrDB.QueryRowContext(ctx, query, args...).
		Scan(&user.ID, &user.Username, &user.Balance); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		err = fmt.Errorf("failed to execute GetUserByUsername query: %w", err)
		logger.Error(err, caller)
		return nil, err
	}

	return &user, nil
}
