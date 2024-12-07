package repository

import (
	"context"
	"jarger-user-service/proto/proto_models"
	"jarger-user-service/service/auth"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcAuthRepository struct {
	grpcAddress string
	timeout     int
}

func NewGrpcAuthRepositoryImpl(grpcAddress string, timeout int) auth.GrpcAuthRepository {
	return &grpcAuthRepository{
		grpcAddress: grpcAddress,
		timeout:     timeout,
	}
}

func (g *grpcAuthRepository) FetchAccountByUsername(ctx context.Context, username, source string) (*proto_models.AuthResponse, error) {
	var (
		request = &proto_models.AuthRequest{
			Username: username,
			Source:   source,
		}
		response = new(proto_models.AuthResponse)
		conn     = new(grpc.ClientConn)
		client   proto_models.AuthClient
		err      error
	)

	if conn, err = grpc.NewClient(g.grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials())); err != nil {
		return nil, err
	}
	defer conn.Close()

	client = proto_models.NewAuthClient(conn)
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Duration(g.timeout*int(time.Second)))
	defer cancel()

	if response, err = client.FetchAccountByUsername(ctxTimeout, request); err != nil {
		return nil, err
	}

	if response == nil {
		return nil, nil
	}

	return response, nil
}
