package user

import "github.com/gofiber/fiber/v2"

type RegisterHandler interface {
	FetchUserByUsername(c *fiber.Ctx) error
}