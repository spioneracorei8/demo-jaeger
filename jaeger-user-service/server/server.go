package server

import (
	"fmt"
	"jarger-user-service/routes"
	_auth_handler "jarger-user-service/service/auth/repository"
	_auth_us "jarger-user-service/service/auth/usecase"
	_user_handler "jarger-user-service/service/user/handler"
	_user_repo "jarger-user-service/service/user/repository"
	_user_us "jarger-user-service/service/user/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Server struct {
	APP_PORT  string
	ROOT_PATH string

	DRIVER_NAME          string
	PSQL_CONNECTION_USER string

	SERVICE_CLIENT_AUTH_GRPC_ADDRESS string
	GRPC_TIMEOUT                     int
	GRPC_MAX_RECEIVE_SIZE            int
}

func connectPSQL(driverName, conn string) *sqlx.DB {
	var (
		connection *sqlx.DB
		err        error
	)
	if connection, err = sqlx.Connect(driverName, conn); err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %s", err.Error()))
	}
	return connection
}

func (s *Server) Start() {
	var (
		app    *fiber.App
		psqlDB *sqlx.DB
		// grpcMaxReceiveSize = (1024 * 1024) * s.GRPC_MAX_RECEIVE_SIZE
	)
	app = fiber.New()

	psqlDB = connectPSQL(s.DRIVER_NAME, s.PSQL_CONNECTION_USER)

	// server := grpc.NewServer(grpc.MaxRecvMsgSize(grpcMaxReceiveSize))
	// defer server.GracefulStop()

	// # REPOSITORIES
	registerRepo := _user_repo.NewPsqlRegisterRepositoryImpl(psqlDB)
	authRepo := _auth_handler.NewGrpcAuthRepositoryImpl(s.SERVICE_CLIENT_AUTH_GRPC_ADDRESS, s.GRPC_TIMEOUT)

	// # USECASES
	registerUs := _user_us.NewRegisterUseaseImpl(registerRepo)
	authUs := _auth_us.NewGrpcAuthUsecaseImpl(authRepo)

	// # HANDLERS
	registerHandler := _user_handler.NewRegisterHandlerImpl(registerUs, authUs)

	// # init fiber

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello jarger.")
	})

	// # init api
	api := routes.NewRoute(app)
	api.RegisterRoutes(registerHandler)

	app.Listen(fmt.Sprintf(":%s", s.APP_PORT))
}
