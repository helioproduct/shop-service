package info

import (
	"context"
	purchaseRepository "shop-service/internal/repository/purchase"
	"shop-service/internal/repository/transfer"
	purhcaseUsecase "shop-service/internal/usecase/purchase"
	transferUsecase "shop-service/internal/usecase/transfer"
)

type (
	PurchaseUsecase interface {
		GetSummary(ctx context.Context, req purchaseRepository.PurchaseSummaryRequest) ([]*purchaseRepository.PurchaseSummary, error)
		BuyItemByName(ctx context.Context, req purhcaseUsecase.BuyItemRequest) error
	}

	TransferUsecsae interface {
		GetSentCoinsSummary(ctx context.Context, username string) ([]*transfer.SentCoinsSummary, error)
		GetReceivedCoinsSummary(ctx context.Context, username string) ([]*transfer.ReceivedCoinsSummary, error)
		SendCoins(ctx context.Context, req transferUsecase.SendCoinsRequest) error
	}

	UserUsecase interface {
		GetBalance(ctx context.Context, username string) (uint64, error)
	}
)

type Hanlder struct {
	purchaseUsecase PurchaseUsecase
	transferUsecase TransferUsecsae
	userUsecase     UserUsecase
}

func NewInfoHandler(purchaseUsecase PurchaseUsecase, transferUsecase TransferUsecsae, userUsecase UserUsecase) *Hanlder {
	return &Hanlder{
		purchaseUsecase: purchaseUsecase,
		transferUsecase: transferUsecase,
		userUsecase:     userUsecase,
	}
}
