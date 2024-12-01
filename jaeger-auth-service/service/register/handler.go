package register

import "github.com/gofiber/fiber/v2"

type RegisterHandler interface {
	FetchAccountByUsername(c *fiber.Ctx) error
}