package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"koriebruh/uas-ppb/conf"
	"koriebruh/uas-ppb/handler"
	"net/http"
)

func SetupApp() *fiber.App {
	app := fiber.New()
	conn := conf.NewConnection()
	validate := validator.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: false,
	}))

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

	app.Delete("/api/users/:id", userHandler.RemoveUserById)
	app.Get("/api/users", userHandler.FindAllUser)
	app.Get("/api/carts/history", userHandler.HistoryCheckout)

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: false,
	}))

	return app
}

func Handler(w http.ResponseWriter, r *http.Request) {
	app := SetupApp() // Panggil dari package api
	adaptor.FiberApp(app)(w, r)
}
