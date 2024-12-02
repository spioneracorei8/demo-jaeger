package register

import (
	"context"
	"jarger-user-service/models"
)

type RegisterUsecase interface {
	FetchUserByUsername(ctx context.Context, username, sourse string) (*models.User, error)
}
