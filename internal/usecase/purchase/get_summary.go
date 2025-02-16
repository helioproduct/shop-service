package purchase

import (
	"context"
	"fmt"
	purchaseRepository "shop-service/internal/repository/purchase"
)

func (uc *PurchaseUsecase) GetSummary(ctx context.Context, req purchaseRepository.PurchaseSummaryRequest) ([]*purchaseRepository.PurchaseSummary, error) {
	summary, err := uc.purchaseRepo.GetPurchaseSummary(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get purchase summary: %w", err)
	}

	return summary, nil
}
