package auth

import (
	"context"
	"jarger-user-service/proto/proto_models"
)

type GrpcAuthRepository interface {
	FetchAccountByUsername(ctx context.Context, username, source string) (*proto_models.AuthResponse, error)
}
