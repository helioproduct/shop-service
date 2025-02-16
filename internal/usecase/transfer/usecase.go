package transfer

import (
	"context"
	"shop-service/internal/domain"
	"shop-service/internal/repository/transfer"
	"shop-service/internal/repository/user"

	"github.com/avito-tech/go-transaction-manager/trm/v2"
)

type (
	TransferRepository interface {
		CreateTransfer(ctx context.Context, transfer domain.Transfer) (*domain.Transfer, error)
		GetSentCoinsSummary(ctx context.Context, fromUsername string) ([]transfer.SentCoinsSummary, error)
		GetReceivedCoinsSummary(ctx context.Context, toUsername string) ([]transfer.ReceivedCoinsSummary, error)
	}

	UserRepository interface {
		GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
		UpdateUser(ctx context.Context, req user.UpdateUserRequest) error
	}
)

type TransferUsecase struct {
	trm          trm.Manager
	transferRepo TransferRepository
	userRepo     UserRepository
}

func NewTransferUsecase(
	trm trm.Manager,
	transferRepo TransferRepository,
	userRepo UserRepository,
) *TransferUsecase {
	return &TransferUsecase{
		trm:          trm,
		transferRepo: transferRepo,
		userRepo:     userRepo,
	}
}
