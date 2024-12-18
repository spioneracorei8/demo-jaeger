package user

import (
	"context"
	"jaeger-auth-service/models"
)

type UserRepository interface {
	FetchAccountByUsername(ctx context.Context, username, source string) (*models.Account, error)
}
