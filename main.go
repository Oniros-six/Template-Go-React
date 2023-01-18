package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	app.Listen(":" + port)
}
