package user

import (
	"context"
	"errors"
	"fmt"
	"shop-service/internal/domain"
	"shop-service/pkg/constant"
	"shop-service/pkg/logger"

	"github.com/lib/pq"

	sq "github.com/Masterminds/squirrel"
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
		return nil, domain.ErrInternalError
	}

	var user domain.User
	trOrDB := r.txGetter.DefaultTrOrDB(ctx, r.db)
	if err := trOrDB.QueryRowContext(ctx, query, args...).
		Scan(&user.ID, &user.Username, &user.Balance); err != nil {

		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == constant.PostgresUniqueViolationErr {
				logger.Error(domain.ErrUserExists, caller)
				return nil, domain.ErrUserExists
			}
		}
		logger.Error(err, caller)
		return nil, domain.ErrInternalError
	}

	return &user, nil
}
