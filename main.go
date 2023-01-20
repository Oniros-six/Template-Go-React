package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	app.Static("/", "./client/dist")

	app.Get("/prueba", func(c *fiber.Ctx) error {
		return c.SendString("Hello")
	})

	app.Listen(":" + port)
	fmt.Println(port)
}
