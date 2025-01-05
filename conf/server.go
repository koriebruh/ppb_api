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
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
		ExposeHeaders:    "Set-Cookie",
		MaxAge:           86400,
	}))

	if err := app.Listen(address); err != nil {
		log.Fatalf("Failed to start Fiber server: %v", err)
	}
}
