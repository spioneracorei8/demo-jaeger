package register

import (
	"context"
	"jaeger-auth-service/models"
)

type RegisterUsecase interface {
	FetchAccountByUsername(ctx context.Context, username, sourse string) (*models.Account, error)
}
