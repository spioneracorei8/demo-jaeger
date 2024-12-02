package handler

import (
	"context"
	"jarger-user-service/models"
	"jarger-user-service/service/register"

	"github.com/gofiber/fiber/v2"
)

type registerHandler struct {
	registerUs register.RegisterUsecase
}

func NewRegisterHandlerImpl(registerUs register.RegisterUsecase) register.RegisterHandler {
	return &registerHandler{
		registerUs: registerUs,
	}
}

func (h *registerHandler) FetchUserByUsername(c *fiber.Ctx) error {
	var (
		rawJson = c.FormValue("data")
		sourse  = c.Get("sourse")
		account = new(models.User)
		ctx     = context.Background()
		err     error
	)

	if account, err = unmarshal(rawJson); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": "BadRequest",
		})
	}

	if account, err = h.registerUs.FetchUserByUsername(ctx, account.Username, sourse); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"status": "InternalServerError",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"account": account,
	})
}
