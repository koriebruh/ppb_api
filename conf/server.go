package conf

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"
)

func STARTSERVER(app *fiber.App) {
	config := NewConfig()
	address := fmt.Sprintf(":%s", config.Server.Port)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Atur ini ke domain tertentu jika diperlukan
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type, Authorization",
	}))

	if err := app.Listen(address); err != nil {
		log.Fatalf("Failed to start Fiber server: %v", err)
	}
}
