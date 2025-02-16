package user

import "context"

func (uc *UserUsecase) GetBalance(ctx context.Context, username string) (uint64, error) {
	user, err := uc.userRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return 0, err
	}

	return user.Balance, nil
}
