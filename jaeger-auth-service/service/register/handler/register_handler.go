package handler

import (
	"context"
	"jaeger-auth-service/models"
	"jaeger-auth-service/service/register"

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

func (h *registerHandler) FetchAccountByUsername(c *fiber.Ctx) error {
	var (
		rawJson = c.FormValue("data")
		source  = c.Get("source")
		account = new(models.Account)
		ctx     = context.Background()
		err     error
	)

	if account, err = unmarshal(rawJson); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": "BadRequest",
		})
	}

	if account, err = h.registerUs.FetchAccountByUsername(ctx, account.Username, source); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"status": "InternalServerError",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"account": account,
	})
}
