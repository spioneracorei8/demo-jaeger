package server

import (
	"fmt"
	"jarger-user-service/routes"
	_auth_handler "jarger-user-service/service/auth/repository"
	_auth_us "jarger-user-service/service/auth/usecase"
	_user_handler "jarger-user-service/service/user/handler"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type Server struct {
	SERVERREADY          chan bool
	APP_PORT             string
	ROOT_PATH            string
	GRPC_PORT            string
	DRIVER_NAME          string
	PSQL_CONNECTION_USER string

	SERVICE_SERVER_AUTH_GRPC_ADDRESS string
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

func (s *Server) startGrpcServer(grpcServ *grpc.Server) {
	var (
		listen net.Listener
		err    error
	)

	if listen, err = net.Listen("tcp", fmt.Sprintf(":%s", s.GRPC_PORT)); err != nil {
		panic(fmt.Sprintf("Error while Listening tcp on port %s: %s", s.GRPC_PORT, err.Error()))
	}
	fmt.Printf(`
============================================
|	grpc server running on port:%s   |
============================================`, s.GRPC_PORT)
	fmt.Println()
	if err = grpcServ.Serve(listen); err != nil {
		panic(fmt.Sprintf("Error while grpc serve:%s", err.Error()))
	}
}

func (s *Server) Start() {
	var (
		app                *fiber.App
		psqlDB             *sqlx.DB
		grpcMaxReceiveSize = (1024 * 1024) * s.GRPC_MAX_RECEIVE_SIZE
		err                error
	)
	app = fiber.New(fiber.Config{
		AppName: "demo-jaeger-user-service",
	})

	psqlDB = connectPSQL(s.DRIVER_NAME, s.PSQL_CONNECTION_USER)
	fmt.Println(psqlDB)
	grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(grpcMaxReceiveSize))
	defer grpcServer.GracefulStop()

	//==============================================================
	// # Repositoryies
	//==============================================================
	// userRepo := _user_repo.NewPsqlUserRepositoryImpl(psqlDB)
	authRepo := _auth_handler.NewGrpcAuthRepositoryImpl(s.SERVICE_SERVER_AUTH_GRPC_ADDRESS, s.GRPC_TIMEOUT)

	//==============================================================
	// # Usecases
	//==============================================================
	// userUs := _user_us.NewUserUseaseImpl(userRepo)
	authUs := _auth_us.NewGrpcAuthUsecaseImpl(authRepo)

	//==============================================================
	// # Handlers
	//==============================================================
	userHandler := _user_handler.NewUserHandlerImpl(authUs)

	//==============================================================
	// # Fiber Routes
	//==============================================================
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello jarger.")
	})
	api := routes.NewRoute(app)
	api.UserRoutes(userHandler)

	//==============================================================
	// # Grpc Routes
	//==============================================================
	// grpcRoute := routes.NewGrpcRoute(grpcServer)
	// grpcRoute.RegisterAuthRoutes()

	go func() {
		if r := recover(); r != nil {
			s.SERVERREADY <- false
			panic(fmt.Sprintf("Error while starting grpc on port %s: %s", s.GRPC_PORT, err.Error()))
		} else {
			s.startGrpcServer(grpcServer)
		}
	}()

	if err = app.Listen(fmt.Sprintf(":%s", s.APP_PORT)); err != nil {
		panic(fmt.Sprintf("Failed to listening on port %s: %s", s.APP_PORT, err.Error()))
	}
}
