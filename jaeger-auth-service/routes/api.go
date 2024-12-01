package routes

import (
	_register_handler "jaeger-auth-service/service/register"

	"github.com/gofiber/fiber/v2"
)

type Route struct {
	c *fiber.App
}

func NewRoute(c *fiber.App) *Route {
	return &Route{
		c: c,
	}
}

func (r *Route) RegisterRoutes(h _register_handler.RegisterHandler) {
	api := r.c.Group("/api")

	api.Post("/v1/username", h.FetchAccountByUsername)
}
