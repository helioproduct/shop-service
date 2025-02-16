package transfer_test

import (
	"context"
	"errors"
	transferMock "shop-service/internal/mocks/repository/transfer"
	transferMockRepository "shop-service/internal/mocks/repository/transfer"
	userMockRepository "shop-service/internal/mocks/repository/user"
	transferRepository "shop-service/internal/repository/transfer"
	"shop-service/internal/usecase/transfer"

	// "shop-service/internal/usecase/transfer"
	"testing"

	trmMock "github.com/avito-tech/go-transaction-manager/trm/v2/drivers/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestTransferUsecase_GetSentCoinsSummary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockTRM := trmMock.NewMockManager(ctrl)

	mockUserRepo := userMockRepository.NewUserRepository(t)
	mockTransferRepo := transferMockRepository.NewTransferRepository(t)
	uc := transfer.NewTransferUsecase(mockTRM, mockTransferRepo, mockUserRepo)

	// ctx := context.Background()
	username := "alice"

	expectedSummary := []*transferRepository.SentCoinsSummary{
		{ToUsername: "bob", TotalSent: 100},
		{ToUsername: "charlie", TotalSent: 200},
	}

	// Успешный сценарий
	mockTransferRepo.EXPECT().
		GetSentCoinsSummary(ctx, username).
		Return(expectedSummary, nil).
		Times(1)

	summary, err := uc.GetSentCoinsSummary(ctx, username)
	require.NoError(t, err)
	require.Len(t, summary, 2)
	require.Equal(t, expectedSummary, summary)

	// Ошибка из репозитория
	mockTransferRepo.EXPECT().
		GetSentCoinsSummary(ctx, username).
		Return(nil, errors.New("repository error")).
		Times(1)

	summary, err = uc.GetSentCoinsSummary(ctx, username)
	require.Error(t, err)
	require.Nil(t, summary)
	require.Contains(t, err.Error(), "repository error")
}

func TestTransferUsecase_GetReceivedCoinsSummary(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockTransferRepo := transferMock.NewTransferRepository(t)
	uc := transfer.NewTransferUsecase(nil, mockTransferRepo, nil)

	ctx := context.Background()
	username := "alice"

	expectedSummary := []*transferRepository.ReceivedCoinsSummary{
		{FromUsername: "bob", TotalReceived: 150},
		{FromUsername: "dave", TotalReceived: 250},
	}

	// Успешный сценарий
	mockTransferRepo.EXPECT().
		GetReceivedCoinsSummary(ctx, username).
		Return(expectedSummary, nil).
		Times(1)

	summary, err := uc.GetReceivedCoinsSummary(ctx, username)
	require.NoError(t, err)
	require.Len(t, summary, 2)
	require.Equal(t, expectedSummary, summary)

	// Ошибка из репозитория
	mockTransferRepo.EXPECT().
		GetReceivedCoinsSummary(ctx, username).
		Return(nil, errors.New("repository error")).
		Times(1)

	summary, err = uc.GetReceivedCoinsSummary(ctx, username)
	require.Error(t, err)
	require.Nil(t, summary)
	require.Contains(t, err.Error(), "repository error")
}
