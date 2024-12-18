package user

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	FetchUserByUsername(c *fiber.Ctx) error
}