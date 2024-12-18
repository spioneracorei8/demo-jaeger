package usecase

import (
	"context"
	"jaeger-auth-service/models"
	"jaeger-auth-service/service/user"
)

type userUsecase struct {
	userRepo user.UserRepository
}

func NewUserUseaseImpl(userRepo user.UserRepository) user.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) FetchAccountByUsername(ctx context.Context, username, source string) (*models.Account, error) {
	return u.userRepo.FetchAccountByUsername(ctx, username, source)
}
