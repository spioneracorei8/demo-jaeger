package usecase

import (
	"context"
	"jarger-user-service/proto/proto_models"
	"jarger-user-service/service/auth"
)

type authUsecase struct {
	authRepo auth.GrpcAuthRepository
}

func NewGrpcAuthUsecaseImpl(authRepo auth.GrpcAuthRepository) auth.AuthUsecase {
	return &authUsecase{
		authRepo: authRepo,
	}
}

func (u *authUsecase) FetchAccountByUsername(ctx context.Context,  username, source string) (*proto_models.AuthResponse, error) {
	return u.authRepo.FetchAccountByUsername(ctx, username, source)
}
