package handler

import (
	"context"
	"fmt"
	"jarger-user-service/models"
	"jarger-user-service/proto/proto_models"
	"jarger-user-service/service/auth"
	"jarger-user-service/service/user"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	authUs auth.AuthUsecase
}

func NewUserHandlerImpl(authUs auth.AuthUsecase) user.UserHandler {
	return &userHandler{
		authUs: authUs,
	}
}

func (h *userHandler) FetchUserByUsername(c *fiber.Ctx) error {
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
	if source == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "source cannot be empty.",
			"status": "BadRequest",
		})
	}
	if user.Username == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "username cannot be empty.",
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
