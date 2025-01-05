package helper

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func SuccessResponse(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(http.StatusOK).JSON(data)
}
