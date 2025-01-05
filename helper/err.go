package helper

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ErrResponse(ctx *fiber.Ctx, err error) error {
	return ctx.Status(http.StatusBadRequest).JSON(map[string]interface{}{
		"error": err.Error(),
	})
}
