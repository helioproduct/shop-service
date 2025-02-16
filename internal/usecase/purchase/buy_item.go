package purchase

import (
	"context"
	"fmt"
	"shop-service/internal/domain"
	purchaseRepository "shop-service/internal/repository/purchase"
	userRepository "shop-service/internal/repository/user"
)

type BuyItemRequest struct {
	Username    string
	ProductName string
	Quantity    uint64
}

func (uc *PurchaseUsecase) BuyItemByName(ctx context.Context, req BuyItemRequest) error {
	return uc.trm.Do(ctx, func(ctx context.Context) error {
		user, err := uc.userRepo.GetUserByUsername(ctx, req.Username)
		if err != nil {
			return fmt.Errorf("failed to get user: %w", err)
		}

		product, err := uc.productRepo.GetProductByName(ctx, req.ProductName)
		if err != nil {
			return fmt.Errorf("failed to get product: %w", err)
		}

		totalCost := product.Price * req.Quantity
		if user.Balance < totalCost {
			return fmt.Errorf("insufficient balance: have %d, need %d", user.Balance, totalCost)
		}

		purchaseReq := mapBuyItemRequest(req, user, product)
		if _, err = uc.purchaseRepo.CreatePurchase(ctx, purchaseReq); err != nil {
			return fmt.Errorf("failed to create purchase: %w", err)
		}

		updatedBalance := user.Balance - totalCost
		updateReq := userRepository.UpdateUserRequest{
			UserID:  user.ID,
			Balance: &updatedBalance,
		}

		if err := uc.userRepo.UpdateUser(ctx, updateReq); err != nil {
			return fmt.Errorf("failed to update user balance: %w", err)
		}

		return nil
	})
}

func mapBuyItemRequest(req BuyItemRequest, user *domain.User, product *domain.Product) purchaseRepository.CreatePurchaseRequest {
	return purchaseRepository.CreatePurchaseRequest{
		UserID:    user.ID,
		ProductID: product.ID,
		Quantity:  req.Quantity,
	}
}
