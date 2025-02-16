package transfer

import (
	"context"
	transferUsecase "shop-service/internal/usecase/transfer"
)

type (
	TransferUsecase interface {
		SendCoins(ctx context.Context, req transferUsecase.SendCoinsRequest) error
	}

	UserUsecase interface {
		GetBalance(ctx context.Context, username string) (uint64, error)
	}
)

type Handler struct {
	transferUsecase TransferUsecase
	userUsecase     UserUsecase
}

func NewTransferHandler(transferUsecase TransferUsecase, userUsecase UserUsecase) *Handler {
	return &Handler{
		transferUsecase: transferUsecase,
		userUsecase:     userUsecase,
	}
}
