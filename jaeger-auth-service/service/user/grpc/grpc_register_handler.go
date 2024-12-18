package grpc

import (
	"context"
	"jaeger-auth-service/proto/proto_models"
	"jaeger-auth-service/service/user"
)

type grpcAuthHandler struct {
	userUs user.UserUsecase
	proto_models.UnimplementedAuthServer
}

func NewGrpcAuthHandler(userUs user.UserUsecase) proto_models.AuthServer {
	return &grpcAuthHandler{
		userUs: userUs,
	}
}

func (g *grpcAuthHandler) FetchAccountByUsername(ctx context.Context, request *proto_models.AuthRequest) (*proto_models.AuthResponse, error) {
	if request == nil {
		return &proto_models.AuthResponse{}, nil
	}

	account, err := g.userUs.FetchAccountByUsername(ctx, request.Username, request.Source)
	if err != nil {
		return nil, err
	}
	if account == nil {
		return &proto_models.AuthResponse{}, nil
	}

	return &proto_models.AuthResponse{
		Id:              account.Id.String(),
		UserId:          account.UserId.String(),
		Status:          account.Status,
		WebAccess:       account.WebAccess,
		RevokeTokenCode: account.RevokeTokenCode,
	}, nil
}
