package routes

import (
	_user_handler "jarger-user-service/service/user"

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

func (r *Route) UserRoutes(h _user_handler.UserHandler) {
	api := r.c.Group("/api")

	api.Post("/v1/username", h.FetchUserByUsername)
}
