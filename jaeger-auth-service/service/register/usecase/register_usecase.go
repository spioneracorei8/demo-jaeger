package usecase

import (
	"context"
	"jaeger-auth-service/models"
	"jaeger-auth-service/service/register"
)

type registerUsecase struct {
	registerRepo register.RegisterRepository
}

func NewRegisterUseaseImpl(registerRepo register.RegisterRepository) register.RegisterUsecase {
	return &registerUsecase{
		registerRepo: registerRepo,
	}
}

func (u *registerUsecase) FetchAccountByUsername(ctx context.Context, username, sourse string) (*models.Account, error) {
	return u.registerRepo.FetchAccountByUsername(ctx, username, sourse)
}
