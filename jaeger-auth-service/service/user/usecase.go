package user

import (
	"context"
	"jaeger-auth-service/models"
)

type UserUsecase interface {
	FetchAccountByUsername(ctx context.Context, username, source string) (*models.Account, error)
}
