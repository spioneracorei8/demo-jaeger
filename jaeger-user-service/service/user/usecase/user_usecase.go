package usecase

import (
	"context"
	"jarger-user-service/models"
	"jarger-user-service/service/user"
)

type registerUsecase struct {
	registerRepo user.RegisterRepository
}

func NewRegisterUseaseImpl(registerRepo user.RegisterRepository) user.RegisterUsecase {
	return &registerUsecase{
		registerRepo: registerRepo,
	}
}

func (u *registerUsecase) FetchUserByUsername(ctx context.Context, username, source string) (*models.User, error) {
	return u.registerRepo.FetchUserByUsername(ctx, username, source)
}
