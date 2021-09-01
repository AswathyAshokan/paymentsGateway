package main

import (
	"paymentGateway/controllers"

	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Post("/api/customer", controllers.AddCustomer)
	app.Post("/api/payments", controllers.AddPayments)
	app.Get("/api/payments/:stripeCustomerID", controllers.GetAllPayments)
}

func main() {
	app := fiber.New()

	setupRoutes(app)
	port := os.Getenv("PORT")

    	if port == "" {
        	log.Fatal("$PORT must be set")
    	}
	app.Listen(port)
}
