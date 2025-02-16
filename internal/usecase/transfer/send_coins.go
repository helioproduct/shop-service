package transfer

import (
	"context"
	"shop-service/internal/domain"
)

type SendCoinsRequest struct {
	From   domain.UserID
	To     domain.UserID
	Amount uint64
}

func (uc *TransferUsecase) SendCoins(ctx context.Context) error {
	return nil
}
