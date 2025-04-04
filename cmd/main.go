package main

import (
	"dmitry-silinsky/go-fiber-hw/internal/pages"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	app.Use(recover.New())

	pages.NewHandler(app)

	app.Listen(":3000")
}
