package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"shop-service/internal/domain"

	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog/log"
)

func (repo *UserRepository) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	caller := "UserRepository.GetUserByUsername"

	query, args, err := sq.Select("id", "username", "balance").
		From("users").
		Where(sq.Eq{"username": username}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		err = fmt.Errorf("failed to build GetUserByUsername query: %w", err)
		log.Err(err).Str("caller", caller).Send()
		return nil, err
	}

	trOrDB := repo.txGetter.DefaultTrOrDB(ctx, repo.db)

	var user domain.User
	if err := trOrDB.QueryRowContext(ctx, query, args...).
		Scan(&user.ID, &user.Username, &user.Balance); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrUserNotFound
		}
		err = fmt.Errorf("failed to execute GetUserByUsername query: %w", err)
		log.Err(err).Str("caller", caller).Send()
		return nil, err
	}

	return &user, nil
}
