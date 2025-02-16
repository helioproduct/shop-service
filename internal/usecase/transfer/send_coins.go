package transfer

import (
	"context"
	"fmt"
	"shop-service/internal/domain"
	"shop-service/internal/repository/user"

	"github.com/rs/zerolog/log"
)

type SendCoinsRequest struct {
	From   domain.UserID
	To     domain.UserID
	Amount uint64
}

func (uc *TransferUsecase) SendCoins(ctx context.Context, req SendCoinsRequest) error {
	caller := "TransferUsecase.SendCoins"

	err := req.Validate()
	if err != nil {
		log.Err(err).Str("caller", caller).Msg("validation failed")
		return err
	}

	return uc.trm.Do(ctx, func(ctx context.Context) error {
		log.Debug().Str("caller", caller).Msg("starting transaction")

		sender, err := uc.userRepo.GetUserByID(ctx, req.From)
		if err != nil {
			err = fmt.Errorf("error getting sender: %w", err)
			log.Err(err).Str("caller", caller).Send()
			return err
		}

		recipient, err := uc.userRepo.GetUserByID(ctx, req.To)
		if err != nil {
			err = fmt.Errorf("error getting recipient: %w", err)
			log.Err(err).Str("caller", caller).Send()
			return err
		}

		if sender.Balance < req.Amount {
			log.Error().Str("caller", caller).Msg("insufficient balance")
			return domain.ErrInsufficientBalance
		}

		newSenderBalance := sender.Balance - req.Amount
		err = uc.userRepo.UpdateUser(ctx, user.UpdateUserRequest{
			UserID:  sender.ID,
			Balance: &newSenderBalance,
		})
		if err != nil {
			err = fmt.Errorf("failed to update sender balance: %w", err)
			log.Err(err).Str("caller", caller).Send()
			return err
		}

		newRecipientBalance := recipient.Balance + req.Amount
		err = uc.userRepo.UpdateUser(ctx, user.UpdateUserRequest{
			UserID:  recipient.ID,
			Balance: &newRecipientBalance,
		})
		if err != nil {
			err = fmt.Errorf("failed to update recipient balance: %w", err)
			log.Err(err).Str("caller", caller).Send()
			return err
		}

		transfer := domain.Transfer{
			From:   req.From,
			To:     req.To,
			Amount: req.Amount,
		}

		_, err = uc.transferRepo.CreateTransfer(ctx, transfer)
		if err != nil {
			err = fmt.Errorf("failed to create transfer record: %w", err)
			log.Err(err).Str("caller", caller).Send()
			return err
		}

		log.Info().
			Str("caller", caller).
			Str("from_user", sender.Username).
			Str("to_user", recipient.Username).
			Uint64("amount", req.Amount).
			Msg("Transfer completed successfully")

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

	if req.From == 0 {
		return domain.ErrMissingFromUser
	}

	if req.To == 0 {
		return domain.ErrMissingToUser
	}

	return nil
}
