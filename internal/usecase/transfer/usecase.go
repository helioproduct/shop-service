package transfer

import (
	"context"
	"shop-service/internal/domain"
	"shop-service/internal/repository/transfer"

	"github.com/avito-tech/go-transaction-manager/trm"
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
	trm          trm.Manager
	transferRepo TransferRepository
	userRepo     UserRepository
}

func NewTransferRepository(
	trm trm.Manager,
	transferRepo TransferRepository,
	userRepo UserRepository,
) *TransferUsecase {
	return &TransferUsecase{
		transferRepo: transferRepo,
		userRepo:     userRepo,
	}
}
