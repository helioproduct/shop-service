package transfer

import (
	"context"
	"fmt"
	"shop-service/internal/domain"
	"shop-service/internal/repository/user"
	"shop-service/pkg/logger"

	"github.com/rs/zerolog/log"
)

type SendCoinsRequest struct {
	From   string
	To     string
	Amount uint64
}

func (uc *TransferUsecase) SendCoins(ctx context.Context, req SendCoinsRequest) error {
	caller := "TransferUsecase.SendCoins"

	err := req.Validate()
	if err != nil {
		logger.Error(err, caller)
		return err
	}

	return uc.trm.Do(ctx, func(ctx context.Context) error {
		log.Debug().Str("caller", caller).Msg("starting transaction")

		sender, err := uc.userRepo.GetUserByUsername(ctx, req.From)
		if err != nil {
			err = fmt.Errorf("error getting sender: %w", err)
			logger.Error(err, caller)
			return err
		}

		recipient, err := uc.userRepo.GetUserByUsername(ctx, req.To)
		if err != nil {
			err = fmt.Errorf("error getting recipient: %w", err)
			logger.Error(err, caller)
			return err
		}

		if sender.Balance < req.Amount {
			err = domain.ErrInsufficientBalance
			logger.Error(err, caller)
			return err
		}

		newSenderBalance := sender.Balance - req.Amount
		err = uc.userRepo.UpdateUser(ctx, user.UpdateUserRequest{
			UserID:  sender.ID,
			Balance: &newSenderBalance,
		})
		if err != nil {
			err = fmt.Errorf("failed to update sender balance: %w", err)
			logger.Error(err, caller)
			return err
		}

		newRecipientBalance := recipient.Balance + req.Amount
		err = uc.userRepo.UpdateUser(ctx, user.UpdateUserRequest{
			UserID:  recipient.ID,
			Balance: &newRecipientBalance,
		})
		if err != nil {
			err = fmt.Errorf("failed to update recipient balance: %w", err)
			logger.Error(err, caller)
			return err
		}

		transfer := domain.Transfer{
			From:   sender.ID,
			To:     recipient.ID,
			Amount: req.Amount,
		}

		_, err = uc.transferRepo.CreateTransfer(ctx, transfer)
		if err != nil {
			err = fmt.Errorf("failed to create transfer record: %w", err)
			logger.Error(err, caller)
			return err
		}

		logger.Info(caller, "Transfer completed successfully", map[string]interface{}{
			"from_user": sender.Username,
			"to_user":   recipient.Username,
			"amount":    req.Amount,
		})

		return nil
	})
}

func (req *SendCoinsRequest) Validate() error {
	if req.Amount == 0 {
		return domain.ErrZeroAmount
	}

	if req.From == req.To {
		return domain.ErrSameUser
	}

	if req.From == "" {
		return domain.ErrMissingFromUser
	}

	if req.To == "" {
		return domain.ErrMissingToUser
	}

	return nil
}
