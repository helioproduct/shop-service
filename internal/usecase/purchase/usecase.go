package purchase

import (
	"context"
	"shop-service/internal/domain"
	purchaseRepository "shop-service/internal/repository/purchase"
	userRepository "shop-service/internal/repository/user"

	"github.com/avito-tech/go-transaction-manager/trm"
)

type (
	PurchaseRepository interface {
		CreatePurchase(ctx context.Context, req purchaseRepository.CreatePurchaseRequest) (*domain.Purchase, error)
		GetPurchaseSummary(ctx context.Context, req purchaseRepository.PurchaseSummaryRequest) ([]*purchaseRepository.PurchaseSummary, error)
	}

	UserRepository interface {
		GetUserByUsername(ctx context.Context, username string) (*domain.User, error)
		UpdateUser(ctx context.Context, req userRepository.UpdateUserRequest) error
	}

	ProductRepository interface {
		GetProductByName(ctx context.Context, name string) (*domain.Product, error)
	}
)

type PurchaseUsecase struct {
	trm          trm.Manager
	purchaseRepo PurchaseRepository
	userRepo     UserRepository
	productRepo  ProductRepository
}
