package purchase

import (
	"context"
	"fmt"
	purchaseRepository "shop-service/internal/repository/purchase"
	"shop-service/pkg/logger"
)

func (uc *PurchaseUsecase) GetSummary(ctx context.Context, req purchaseRepository.PurchaseSummaryRequest) ([]*purchaseRepository.PurchaseSummary, error) {
	caller := "PurchaseUsecase.GetSummary"

	summary, err := uc.purchaseRepo.GetPurchaseSummary(ctx, req)
	if err != nil {
		err = fmt.Errorf("failed to get purchase summary: %w", err)
		logger.Error(err, caller)
		return nil, err
	}

	return summary, nil
}
