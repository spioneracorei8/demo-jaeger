package server

import (
	"fmt"
	"jaeger-auth-service/routes"
	_register_grpc "jaeger-auth-service/service/register/grpc"
	_register_handler "jaeger-auth-service/service/register/handler"
	_register_repo "jaeger-auth-service/service/register/repository"
	_register_usecase "jaeger-auth-service/service/register/usecase"
	"net"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

type Server struct {
	SERVERREADY                      chan bool
	ROOT_PATH                        string
	APP_PORT                         string
	GRPC_PORT                        string
	GRPC_TIMEOUT                     int
	GRPC_MAX_RECEIVE_SIZE            int
	SERVICE_SERVER_USER_GRPC_ADDRESS string
	DRIVER_NAME                      string
	PSQL_CONNECTION_AUTH             string
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
	// # init fiber
	app = fiber.New(fiber.Config{
		AppName: "demo-jarger-auth-service",
	})

	// # init postgresql
	psqlDB = connectPSQL(s.DRIVER_NAME, s.PSQL_CONNECTION_AUTH)

	grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(grpcMaxReceiveSize))
	defer grpcServer.GracefulStop()

	//==============================================================
	// # Repositoryies
	//==============================================================
	registerRepo := _register_repo.NewPsqlRegisterRepositoryImpl(psqlDB)

	//==============================================================
	// # Usecases
	//==============================================================
	registerUs := _register_usecase.NewRegisterUseaseImpl(registerRepo)

	//==============================================================
	// # Handlers
	//==============================================================
	registerHandler := _register_handler.NewRegisterHandlerImpl(registerUs)
	registerGrpcHandler := _register_grpc.NewGrpcAuthHandler(registerUs)

	//==============================================================
	// # Fiber Routes
	//==============================================================
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello jarger.")
	})
	api := routes.NewRoute(app)
	api.RegisterRoutes(registerHandler)

	//==============================================================
	// # Grpc Routes
	//==============================================================
	grpcRoute := routes.NewGrpcRoute(grpcServer)
	grpcRoute.RegisterAuthRoutes(registerGrpcHandler)

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
