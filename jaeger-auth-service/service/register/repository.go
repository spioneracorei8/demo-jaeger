package register

import (
	"context"
	"jaeger-auth-service/models"
)

type RegisterRepository interface {
	FetchAccountByUsername(ctx context.Context, username, sourse string) (*models.Account, error)
}
