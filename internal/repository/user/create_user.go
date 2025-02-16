package user

import (
	"context"
	"errors"
	"fmt"
	"shop-service/internal/domain"
	"shop-service/pkg/logger"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
)

type CreateUserRequest struct {
	Username       string
	HashedPassword string
	Balance        int
}

func (r *UserRepository) CreateUser(ctx context.Context, req CreateUserRequest) (*domain.User, error) {
	caller := "UserRepository.CreateUser"

	query, args, err := sq.Insert("users").
		Columns("username", "hashed_password", "balance").
		Values(req.Username, req.HashedPassword, req.Balance).
		Suffix("RETURNING id, username, balance").
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		err = fmt.Errorf("failed to build CreateUser query: %w", err)
		logger.Error(err, caller)
		return nil, err
	}

	var user domain.User
	trOrDB := r.txGetter.DefaultTrOrDB(ctx, r.db)
	if err := trOrDB.QueryRowContext(ctx, query, args...).
		Scan(&user.ID, &user.Username, &user.Balance); err != nil {

		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			logger.Error(domain.ErrUserExists, caller)
			logger.Print(caller, "here", "here2")
			return nil, domain.ErrUserExists
		}

		logger.Print(caller, "here", "here3")

		err = fmt.Errorf("failed to insert user: %w", err)
		logger.Error(err, caller)
		return nil, err
	}

	return &user, nil
}
