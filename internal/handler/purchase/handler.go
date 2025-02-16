package purchase

import (
	"context"
	"shop-service/internal/usecase/purchase"
)

type (
	PurchaseUsecase interface {
		BuyItemByName(ctx context.Context, req purchase.BuyItemRequest) error
	}

	UserUsecase interface {
		GetBalance(ctx context.Context, username string) (uint64, error)
	}
)

type Handler struct {
	purchaseUsecase PurchaseUsecase
	userUsecase     UserUsecase
}

func NewPurchaseHandler(purchaseUsecase PurchaseUsecase, userUsecase UserUsecase) *Handler {
	return &Handler{
		purchaseUsecase: purchaseUsecase,
		userUsecase:     userUsecase,
	}
}
