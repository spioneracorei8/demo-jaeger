package server

import (
	"fmt"
	"jaeger-auth-service/routes"
	_register_handler "jaeger-auth-service/service/register/handler"
	_register_repo "jaeger-auth-service/service/register/repository"
	_register_usecase "jaeger-auth-service/service/register/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Server struct {
	APP_PORT  string
	ROOT_PATH string

	DRIVER_NAME          string
	PSQL_CONNECTION_AUTH string
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
	)
	app = fiber.New()

	psqlDB = connectPSQL(s.DRIVER_NAME, s.PSQL_CONNECTION_AUTH)

	// # REPOSITORIES
	registerRepo := _register_repo.NewPsqlRegisterRepositoryImpl(psqlDB)

	// # USECASES
	registerUs := _register_usecase.NewRegisterUseaseImpl(registerRepo)

	// # HANDLERS
	registerHandler := _register_handler.NewRegisterHandlerImpl(registerUs)

	// # init fiber

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello jarger.")
	})

	// # init api
	api := routes.NewRoute(app)
	api.RegisterRoutes(registerHandler)

	app.Listen(fmt.Sprintf(":%s", s.APP_PORT))
}
