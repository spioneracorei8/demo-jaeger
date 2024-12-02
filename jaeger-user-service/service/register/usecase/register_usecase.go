package usecase

import (
	"context"
	"jarger-user-service/models"
	"jarger-user-service/service/register"
)

type registerUsecase struct {
	registerRepo register.RegisterRepository
}

func NewRegisterUseaseImpl(registerRepo register.RegisterRepository) register.RegisterUsecase {
	return &registerUsecase{
		registerRepo: registerRepo,
	}
}

func (u *registerUsecase) FetchUserByUsername(ctx context.Context, username, sourse string) (*models.User, error) {
	return u.registerRepo.FetchUserByUsername(ctx, username, sourse)
}
