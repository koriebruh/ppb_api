package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"koriebruh/uas-ppb/conf"
	"koriebruh/uas-ppb/handler"
)

func SetupApp() *fiber.App {
	app := fiber.New()
	conn := conf.NewConnection()
	validate := validator.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello from Vercel!")
	})

	productHandler := handler.NewProductHandler(conn, validate)
	app.Get("/api/products", productHandler.FindAllProduct)
	app.Get("/api/products/:id", productHandler.FindByIdProduct)
	app.Put("/api/products/:id", productHandler.UpdateProduct)
	app.Post("/api/products", productHandler.CreateProduct)
	app.Delete("/api/products/:id", productHandler.DeleteProduct)

	userHandler := handler.NewUserHandler(conn, validate)
	app.Post("/api/users", userHandler.CreateUser)
	app.Put("/api/users/:id", userHandler.UpdateUser)
	app.Post("/api/users/login", userHandler.Login)

	app.Post("/api/carts", userHandler.AddProductToCart)
	app.Post("/api/carts/user", userHandler.GetCartItems)
	app.Post("/api/carts/add-shipping", userHandler.AddShippingAndGetTotal)
	app.Post("/api/carts/checkout", userHandler.CheckoutAndClearCart)
	app.Post("/api/carts/remove", userHandler.RemoveProductFromCart)

	return app
}
