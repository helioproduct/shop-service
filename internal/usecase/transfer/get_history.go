package transfer

import (
	"context"
	"fmt"
	"shop-service/internal/repository/transfer"
	"shop-service/pkg/logger"
)

func (uc *TransferUsecase) GetSentCoinsSummary(ctx context.Context, username string) ([]*transfer.SentCoinsSummary, error) {
	caller := "TransferUsecase.GetSentCoinsSummary"

	sentSummary, err := uc.transferRepo.GetSentCoinsSummary(ctx, username)
	if err != nil {
		err = fmt.Errorf("failed to get sent coins summary: %w", err)
		logger.Error(err, caller)
		return nil, err
	}

	return sentSummary, nil
}

func (uc *TransferUsecase) GetReceivedCoinsSummary(ctx context.Context, username string) ([]*transfer.ReceivedCoinsSummary, error) {
	caller := "TransferUsecase.GetReceivedCoinsSummary"
	receivedSummary, err := uc.transferRepo.GetReceivedCoinsSummary(ctx, username)
	if err != nil {
		err = fmt.Errorf("failed to get received coins summary: %w", err)
		logger.Error(err, caller)
		return nil, err
	}

	return receivedSummary, nil
}
