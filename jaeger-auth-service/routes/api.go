package routes

import (

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

func (r *Route) RegisterRoutes() {
	_ = r.c.Group("/api")

}
