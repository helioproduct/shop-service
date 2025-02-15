package postgres

import (
	"context"
	"fmt"
	"shop-service/internal/domain"

	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog/log"
)

type CreateUserRequest struct {
	Username       string
	HashedPassword string
	Balance        int
}

func (repo *UserRepository) CreateUser(ctx context.Context, req *CreateUserRequest) (*domain.User, error) {
	caller := "UserRepository.CreateUser"

	query, args, err := sq.Insert("users").
		Columns("username", "hashed_password", "balance").
		Values(req.Username, req.HashedPassword, req.Balance).
		Suffix("RETURNING id, username, balance").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		err = fmt.Errorf("failed to build CreateUser query: %w", err)
		log.Err(err).Str("caller", caller).Send()
		return nil, err
	}

	trOrDB := repo.txGetter.DefaultTrOrDB(ctx, repo.db)

	var user domain.User
	if err := trOrDB.QueryRowContext(ctx, query, args...).
		Scan(&user.ID, &user.Username, &user.Balance); err != nil {
		err = fmt.Errorf("failed to insert user: %w", err)
		log.Err(err).Str("caller", caller).Send()
		return nil, err
	}

	return &user, nil
}
