package register

import (
	"context"
	"jarger-user-service/models"
)

type RegisterRepository interface {
	FetchUserByUsername(ctx context.Context, username, sourse string) (*models.User, error)
}
