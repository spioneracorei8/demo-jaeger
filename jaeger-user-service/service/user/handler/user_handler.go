package handler

import (
	"context"
	"fmt"
	"jarger-user-service/models"
	"jarger-user-service/proto/proto_models"
	"jarger-user-service/service/auth"
	register "jarger-user-service/service/user"

	"github.com/gofiber/fiber/v2"
)

type registerHandler struct {
	registerUs register.RegisterUsecase
	authUs     auth.AuthUsecase
}

func NewRegisterHandlerImpl(registerUs register.RegisterUsecase, authUs auth.AuthUsecase) register.RegisterHandler {
	return &registerHandler{
		registerUs: registerUs,
		authUs:     authUs,
	}
}

func (h *registerHandler) FetchUserByUsername(c *fiber.Ctx) error {
	var (
		rawJson      = c.FormValue("data")
		source       = c.Get("source")
		user         = new(models.User)
		authResponse = new(proto_models.AuthResponse)
		ctx          = context.Background()
		err          error
	)

	if user, err = unmarshalUser(rawJson); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  err.Error(),
			"status": "BadRequest",
		})
	}

	if authResponse, err = h.authUs.FetchAccountByUsername(ctx, user.Username, source); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":  err.Error(),
			"status": "InternalServerError",
		})
	}

	fmt.Println(authResponse)

	// if user, err = h.registerUs.FetchUserByUsername(ctx, user.Username, source); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error":  err.Error(),
	// 		"status": "InternalServerError",
	// 	})
	// }

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user":         user,
		"authResponse": authResponse,
	})
}
