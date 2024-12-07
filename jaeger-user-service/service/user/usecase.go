package user

import (
	"context"
	"jarger-user-service/models"
)

type RegisterUsecase interface {
	FetchUserByUsername(ctx context.Context, username, source string) (*models.User, error)
}
