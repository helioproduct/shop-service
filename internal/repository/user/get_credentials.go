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

func (repo *UserRepository) GetUserHashedPassword(ctx context.Context, username string) (string, error) {
	caller := "UserRepository.GetUserHashedPassword"

	query, args, err := sq.Select("hashed_password").
		From("users").
		Where(sq.Eq{"username": username}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		err = fmt.Errorf("failed to build GetUserHashedPassword query: %w", err)
		logger.Error(err, caller)
		return "", err
	}

	var hashedPassword string
	if err := repo.db.QueryRowContext(ctx, query, args...).Scan(&hashedPassword); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", domain.ErrUserNotFound
		}
		err = fmt.Errorf("failed to execute GetUserHashedPassword query: %w", err)
		logger.Error(err, caller)
		return "", err
	}

	return hashedPassword, nil
}
