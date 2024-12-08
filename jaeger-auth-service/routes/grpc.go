package routes

import (
	"jaeger-auth-service/proto/proto_models"

	"google.golang.org/grpc"
)

type grpcRoute struct {
	server *grpc.Server
}

func NewGrpcRoute(server *grpc.Server) *grpcRoute {
	return &grpcRoute{
		server: server,
	}
}

func (g *grpcRoute) RegisterAuthRoutes(handler proto_models.AuthServer) {
	proto_models.RegisterAuthServer(g.server, handler)
}
