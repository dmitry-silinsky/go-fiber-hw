package main

import (
	"dmitry-silinsky/go-fiber-hw/config"
	"dmitry-silinsky/go-fiber-hw/internal/pages"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.Load()

	appConfig := config.NewAppConfig()

	app := fiber.New()

	app.Use(recover.New())

	pages.NewHandler(app)

	app.Listen(appConfig.ListenAddr)
}
