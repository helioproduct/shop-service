package transfer

import (
	"context"
	"fmt"
	"shop-service/internal/repository/transfer"
)

func (uc *TransferUsecase) GetSentCoinsSummary(ctx context.Context, username string) ([]*transfer.SentCoinsSummary, error) {
	sentSummary, err := uc.transferRepo.GetSentCoinsSummary(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get sent coins summary: %w", err)
	}

	return sentSummary, nil
}

func (uc *TransferUsecase) GetReceivedCoinsSummary(ctx context.Context, username string) ([]*transfer.ReceivedCoinsSummary, error) {
	receivedSummary, err := uc.transferRepo.GetReceivedCoinsSummary(ctx, username)
	if err != nil {
		return nil, fmt.Errorf("failed to get received coins summary: %w", err)
	}

	return receivedSummary, nil
}
