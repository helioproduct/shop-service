package transfer_test

import (
	"context"
	"shop-service/internal/domain"
	transferMockRepository "shop-service/internal/mocks/repository/transfer"
	userMockRepository "shop-service/internal/mocks/repository/user"
	"shop-service/internal/usecase/transfer"
	"testing"

	trmMock "github.com/avito-tech/go-transaction-manager/trm/v2/drivers/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestTransferUsecase_SendCoins_WithTransaction(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockTRM := trmMock.NewMockManager(ctrl)

	mockUserRepo := userMockRepository.NewUserRepository(t)
	mockTransferRepo := transferMockRepository.NewTransferRepository(t)

	uc := transfer.NewTransferUsecase(mockTRM, mockTransferRepo, mockUserRepo)

	req := transfer.SendCoinsRequest{
		From:   "alice",
		To:     "bob",
		Amount: 300,
	}
	alice := &domain.User{ID: 1, Username: "alice", Balance: 200}
	bob := &domain.User{ID: 2, Username: "bob", Balance: 50}

	mockTRM.EXPECT().
		Do(ctx, gomock.Any()).
		DoAndReturn(func(ctx context.Context, fn func(context.Context) error) error {
			return fn(ctx)
		}).
		Times(1)

	mockUserRepo.On("GetUserByUsername", ctx, "alice").
		Return(alice, nil).
		Once()

	mockUserRepo.On("GetUserByUsername", ctx, "bob").
		Return(bob, nil).
		Once()

	err := uc.SendCoins(context.Background(), req)
	require.ErrorIs(t, err, domain.ErrInsufficientBalance)
}

func TestTransferUsecase_SendCoins(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.Background()
	mockTRM := trmMock.NewMockManager(ctrl)

	mockUserRepo := userMockRepository.NewUserRepository(t)
	mockTransferRepo := transferMockRepository.NewTransferRepository(t)

	uc := transfer.NewTransferUsecase(mockTRM, mockTransferRepo, mockUserRepo)

	req := transfer.SendCoinsRequest{
		From:   "alice",
		To:     "bob",
		Amount: 300,
	}

	// alice := &domain.User{ID: 1, Username: "alice", Balance: 500}
	bob := &domain.User{ID: 2, Username: "bob", Balance: 100}

	mockTRM.EXPECT().
		Do(ctx, gomock.Any()).
		DoAndReturn(func(ctx context.Context, fn func(context.Context) error) error {
			return fn(ctx)
		}).
		Times(1)

	t.Run("InsufficientBalance", func(t *testing.T) {
		lowBalanceUser := &domain.User{ID: 1, Username: "alice", Balance: 100}

		mockUserRepo.On("GetUserByUsername", ctx, "alice").
			Return(lowBalanceUser, nil).
			Once()

		mockUserRepo.On("GetUserByUsername", ctx, "bob").
			Return(bob, nil).
			Once()
		// mockUserRepo.EXPECT().GetUserByUsername(ctx, "alice").Return(lowBalanceUser, nil).Times(1)
		// mockUserRepo.EXPECT().GetUserByUsername(ctx, "bob").Return(bob, nil).Times(1)

		err := uc.SendCoins(ctx, req)
		require.ErrorIs(t, err, domain.ErrInsufficientBalance)
	})

}
