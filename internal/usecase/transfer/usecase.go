package transfer

import (
	"context"
	"shop-service/internal/domain"
	"shop-service/internal/repository/transfer"
)

type (
	TransferRepository interface {
		CreateTransfer(ctx context.Context, transfer *domain.Transfer) (*domain.Transfer, error)
		GetTransfers(ctx context.Context, filter *transfer.GetTransfersFilter) ([]domain.Transfer, error)
	}

	UserRepository interface {
		GetUserByID(ctx context.Context, userID domain.UserID) (*domain.User, error)
	}
)

type TransferUsecase struct {
	transferRepo TransferRepository
	userRepo     UserRepository
}

func NewTransferRepository(transferRepo TransferRepository, userRepo UserRepository) *TransferUsecase {
	return &TransferUsecase{
		transferRepo: transferRepo,
		userRepo:     userRepo,
	}
}
